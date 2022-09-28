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

type BrandData struct {
	// xxx
	BrandName *string `json:"BrandName,omitempty" name:"BrandName"`

	// xxx
	BrandCertificateName *string `json:"BrandCertificateName,omitempty" name:"BrandCertificateName"`

	// xxx
	BrandStatus *int64 `json:"BrandStatus,omitempty" name:"BrandStatus"`

	// xxx
	BrandNote *string `json:"BrandNote,omitempty" name:"BrandNote"`

	// xxx
	TransferName *string `json:"TransferName,omitempty" name:"TransferName"`

	// xxx
	TransferStatus *int64 `json:"TransferStatus,omitempty" name:"TransferStatus"`

	// xxx
	TransferNote *string `json:"TransferNote,omitempty" name:"TransferNote"`
}

// Predefined struct for user
type CreateBPFakeURLRequestParams struct {
	// xxx
	ProtectURLId *int64 `json:"ProtectURLId,omitempty" name:"ProtectURLId"`

	// xxx
	FakeURL *string `json:"FakeURL,omitempty" name:"FakeURL"`

	// xxx
	SnapshotNames []*string `json:"SnapshotNames,omitempty" name:"SnapshotNames"`

	// xxx
	Note *string `json:"Note,omitempty" name:"Note"`
}

type CreateBPFakeURLRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	ProtectURLId *int64 `json:"ProtectURLId,omitempty" name:"ProtectURLId"`

	// xxx
	FakeURL *string `json:"FakeURL,omitempty" name:"FakeURL"`

	// xxx
	SnapshotNames []*string `json:"SnapshotNames,omitempty" name:"SnapshotNames"`

	// xxx
	Note *string `json:"Note,omitempty" name:"Note"`
}

func (r *CreateBPFakeURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFakeURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProtectURLId")
	delete(f, "FakeURL")
	delete(f, "SnapshotNames")
	delete(f, "Note")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPFakeURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFakeURLResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBPFakeURLResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPFakeURLResponseParams `json:"Response"`
}

func (r *CreateBPFakeURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFakeURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFalseTicketRequestParams struct {
	// xxx
	FakeURL *string `json:"FakeURL,omitempty" name:"FakeURL"`
}

type CreateBPFalseTicketRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	FakeURL *string `json:"FakeURL,omitempty" name:"FakeURL"`
}

func (r *CreateBPFalseTicketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFalseTicketRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FakeURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPFalseTicketRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFalseTicketResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBPFalseTicketResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPFalseTicketResponseParams `json:"Response"`
}

func (r *CreateBPFalseTicketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFalseTicketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPOfflineAttachmentRequestParams struct {
	// xxx
	BrandName *string `json:"BrandName,omitempty" name:"BrandName"`

	// xx
	BrandCertificateName *string `json:"BrandCertificateName,omitempty" name:"BrandCertificateName"`

	// xx
	TransferName *string `json:"TransferName,omitempty" name:"TransferName"`

	// xx
	AuthorizationName *string `json:"AuthorizationName,omitempty" name:"AuthorizationName"`
}

type CreateBPOfflineAttachmentRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	BrandName *string `json:"BrandName,omitempty" name:"BrandName"`

	// xx
	BrandCertificateName *string `json:"BrandCertificateName,omitempty" name:"BrandCertificateName"`

	// xx
	TransferName *string `json:"TransferName,omitempty" name:"TransferName"`

	// xx
	AuthorizationName *string `json:"AuthorizationName,omitempty" name:"AuthorizationName"`
}

func (r *CreateBPOfflineAttachmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPOfflineAttachmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BrandName")
	delete(f, "BrandCertificateName")
	delete(f, "TransferName")
	delete(f, "AuthorizationName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPOfflineAttachmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPOfflineAttachmentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBPOfflineAttachmentResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPOfflineAttachmentResponseParams `json:"Response"`
}

func (r *CreateBPOfflineAttachmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPOfflineAttachmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPOfflineTicketRequestParams struct {
	// xxx
	FakeURLId *int64 `json:"FakeURLId,omitempty" name:"FakeURLId"`
}

type CreateBPOfflineTicketRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	FakeURLId *int64 `json:"FakeURLId,omitempty" name:"FakeURLId"`
}

func (r *CreateBPOfflineTicketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPOfflineTicketRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FakeURLId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPOfflineTicketRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPOfflineTicketResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBPOfflineTicketResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPOfflineTicketResponseParams `json:"Response"`
}

func (r *CreateBPOfflineTicketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPOfflineTicketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPProtectURLsRequestParams struct {
	// xxx
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// xxx
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// xxx
	LicenseName *string `json:"LicenseName,omitempty" name:"LicenseName"`

	// xxx
	ProtectURLs []*string `json:"ProtectURLs,omitempty" name:"ProtectURLs"`

	// xxx
	ProtectWebs []*string `json:"ProtectWebs,omitempty" name:"ProtectWebs"`
}

type CreateBPProtectURLsRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// xxx
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// xxx
	LicenseName *string `json:"LicenseName,omitempty" name:"LicenseName"`

	// xxx
	ProtectURLs []*string `json:"ProtectURLs,omitempty" name:"ProtectURLs"`

	// xxx
	ProtectWebs []*string `json:"ProtectWebs,omitempty" name:"ProtectWebs"`
}

func (r *CreateBPProtectURLsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPProtectURLsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyName")
	delete(f, "Phone")
	delete(f, "LicenseName")
	delete(f, "ProtectURLs")
	delete(f, "ProtectWebs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPProtectURLsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPProtectURLsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBPProtectURLsResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPProtectURLsResponseParams `json:"Response"`
}

func (r *CreateBPProtectURLsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPProtectURLsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

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

	// x
	EvidenceFileUrl *string `json:"EvidenceFileUrl,omitempty" name:"EvidenceFileUrl"`

	// x
	EvidenceValidStartDate *string `json:"EvidenceValidStartDate,omitempty" name:"EvidenceValidStartDate"`

	// x
	EvidenceValidEndDate *string `json:"EvidenceValidEndDate,omitempty" name:"EvidenceValidEndDate"`
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

	// x
	EvidenceFileUrl *string `json:"EvidenceFileUrl,omitempty" name:"EvidenceFileUrl"`

	// x
	EvidenceValidStartDate *string `json:"EvidenceValidStartDate,omitempty" name:"EvidenceValidStartDate"`

	// x
	EvidenceValidEndDate *string `json:"EvidenceValidEndDate,omitempty" name:"EvidenceValidEndDate"`
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
	delete(f, "EvidenceFileUrl")
	delete(f, "EvidenceValidStartDate")
	delete(f, "EvidenceValidEndDate")
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

	// 企业认证号码类型 1：社会信用代码 2：组织机构代码 3：企业工商注册码 4：其他 默认为1
	CompanyIDType *string `json:"CompanyIDType,omitempty" name:"CompanyIDType"`

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

	// xxx
	Type *string `json:"Type,omitempty" name:"Type"`
}

type CreateCRCompanyVerifyRequest struct {
	*tchttp.BaseRequest
	
	// 企业名称
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 企业认证号码类型 1：社会信用代码 2：组织机构代码 3：企业工商注册码 4：其他 默认为1
	CompanyIDType *string `json:"CompanyIDType,omitempty" name:"CompanyIDType"`

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
	delete(f, "CompanyIDType")
	delete(f, "CompanyID")
	delete(f, "CompanyLegalName")
	delete(f, "ManagerName")
	delete(f, "ManagerPhone")
	delete(f, "VerificationCode")
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
type CreateCRRightFileRequestParams struct {
	// xxx
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// xxx
	FileList []*File `json:"FileList,omitempty" name:"FileList"`
}

type CreateCRRightFileRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// xxx
	FileList []*File `json:"FileList,omitempty" name:"FileList"`
}

func (r *CreateCRRightFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRRightFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "FileList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRRightFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRRightFileResponseParams struct {
	// xxx
	FileIds []*int64 `json:"FileIds,omitempty" name:"FileIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCRRightFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRRightFileResponseParams `json:"Response"`
}

func (r *CreateCRRightFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRRightFileResponse) FromJsonString(s string) error {
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

	// x
	EvidenceFileUrl *string `json:"EvidenceFileUrl,omitempty" name:"EvidenceFileUrl"`

	// x
	EvidenceValidStartDate *string `json:"EvidenceValidStartDate,omitempty" name:"EvidenceValidStartDate"`

	// x
	EvidenceValidEndDate *string `json:"EvidenceValidEndDate,omitempty" name:"EvidenceValidEndDate"`
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

	// x
	EvidenceFileUrl *string `json:"EvidenceFileUrl,omitempty" name:"EvidenceFileUrl"`

	// x
	EvidenceValidStartDate *string `json:"EvidenceValidStartDate,omitempty" name:"EvidenceValidStartDate"`

	// x
	EvidenceValidEndDate *string `json:"EvidenceValidEndDate,omitempty" name:"EvidenceValidEndDate"`
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
	delete(f, "EvidenceFileUrl")
	delete(f, "EvidenceValidStartDate")
	delete(f, "EvidenceValidEndDate")
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
type CreateCRTortRequestParams struct {
	// xx
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// xx
	TortURL *string `json:"TortURL,omitempty" name:"TortURL"`

	// xx
	TortPlat *string `json:"TortPlat,omitempty" name:"TortPlat"`

	// xx
	TortTitle *string `json:"TortTitle,omitempty" name:"TortTitle"`
}

type CreateCRTortRequest struct {
	*tchttp.BaseRequest
	
	// xx
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// xx
	TortURL *string `json:"TortURL,omitempty" name:"TortURL"`

	// xx
	TortPlat *string `json:"TortPlat,omitempty" name:"TortPlat"`

	// xx
	TortTitle *string `json:"TortTitle,omitempty" name:"TortTitle"`
}

func (r *CreateCRTortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRTortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "TortURL")
	delete(f, "TortPlat")
	delete(f, "TortTitle")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRTortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRTortResponseParams struct {
	// xx
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// xx
	TortId *int64 `json:"TortId,omitempty" name:"TortId"`

	// xx
	TortTitle *string `json:"TortTitle,omitempty" name:"TortTitle"`

	// xx
	TortPlat *string `json:"TortPlat,omitempty" name:"TortPlat"`

	// xx
	TortURL *string `json:"TortURL,omitempty" name:"TortURL"`

	// xx
	TortDomain *string `json:"TortDomain,omitempty" name:"TortDomain"`

	// xx
	TortBodyName *string `json:"TortBodyName,omitempty" name:"TortBodyName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCRTortResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRTortResponseParams `json:"Response"`
}

func (r *CreateCRTortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRTortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRUserVerifyRequestParams struct {
	// 用户真实姓名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户身份证ID
	UserID *string `json:"UserID,omitempty" name:"UserID"`

	// 用户手机号码
	UserPhone *string `json:"UserPhone,omitempty" name:"UserPhone"`

	// 短信验证码
	VerificationCode *string `json:"VerificationCode,omitempty" name:"VerificationCode"`

	// xxx
	Type *string `json:"Type,omitempty" name:"Type"`
}

type CreateCRUserVerifyRequest struct {
	*tchttp.BaseRequest
	
	// 用户真实姓名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户身份证ID
	UserID *string `json:"UserID,omitempty" name:"UserID"`

	// 用户手机号码
	UserPhone *string `json:"UserPhone,omitempty" name:"UserPhone"`

	// 短信验证码
	VerificationCode *string `json:"VerificationCode,omitempty" name:"VerificationCode"`

	// xxx
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateCRUserVerifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRUserVerifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "UserID")
	delete(f, "UserPhone")
	delete(f, "VerificationCode")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRUserVerifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRUserVerifyResponseParams struct {
	// 认证状态 0-认证成功 1-认证失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 认证结果返回
	Note *string `json:"Note,omitempty" name:"Note"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCRUserVerifyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRUserVerifyResponseParams `json:"Response"`
}

func (r *CreateCRUserVerifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRUserVerifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRWorkRequestParams struct {
	// 作品名称
	WorkName *string `json:"WorkName,omitempty" name:"WorkName"`

	// 作品分类
	WorkCategory *string `json:"WorkCategory,omitempty" name:"WorkCategory"`

	// 作品内容类型
	WorkType *string `json:"WorkType,omitempty" name:"WorkType"`

	// 作品标签
	WorkSign *string `json:"WorkSign,omitempty" name:"WorkSign"`

	// 作品图片
	WorkPic *string `json:"WorkPic,omitempty" name:"WorkPic"`

	// 创作描述
	WorkDesc *string `json:"WorkDesc,omitempty" name:"WorkDesc"`

	// 是否原创 0:否 1:是
	IsOriginal *string `json:"IsOriginal,omitempty" name:"IsOriginal"`

	// 是否发布 0：未发布 1：已发布
	IsRelease *string `json:"IsRelease,omitempty" name:"IsRelease"`

	// 著作权人ID
	ProducerID *int64 `json:"ProducerID,omitempty" name:"ProducerID"`

	// 创作时间
	ProduceTime *string `json:"ProduceTime,omitempty" name:"ProduceTime"`

	// 样品文件路径
	SampleContentURL *string `json:"SampleContentURL,omitempty" name:"SampleContentURL"`

	// 样本下载Url
	SampleDownloadURL *string `json:"SampleDownloadURL,omitempty" name:"SampleDownloadURL"`

	// 授予类型
	GrantType *string `json:"GrantType,omitempty" name:"GrantType"`

	// 作品发布Url
	SamplePublicURL *string `json:"SamplePublicURL,omitempty" name:"SamplePublicURL"`

	// 是否启用监测 0：不启用 1：启用 默认为0
	IsMonitor *string `json:"IsMonitor,omitempty" name:"IsMonitor"`

	// 是否启用存证0：不存证  2：存证 默认为0
	IsCert *string `json:"IsCert,omitempty" name:"IsCert"`

	// 存证回调地址
	CertUrl *string `json:"CertUrl,omitempty" name:"CertUrl"`

	// 监测回调地址
	MonitorUrl *string `json:"MonitorUrl,omitempty" name:"MonitorUrl"`

	// 创作性质（原创,改编,翻译,汇编,注释,整理,其他)
	ProduceType *string `json:"ProduceType,omitempty" name:"ProduceType"`

	// 白名单
	WhiteLists []*string `json:"WhiteLists,omitempty" name:"WhiteLists"`

	// 作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 著作权人姓名
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// 作者
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 授权书
	Authorization *string `json:"Authorization,omitempty" name:"Authorization"`

	// 授权书开始时间
	AuthorizationStartTime *string `json:"AuthorizationStartTime,omitempty" name:"AuthorizationStartTime"`

	// 授权书结束时间
	AuthorizationEndTime *string `json:"AuthorizationEndTime,omitempty" name:"AuthorizationEndTime"`

	// 内容格式
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// 文件内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 监测结束时间
	MonitorEndTime *string `json:"MonitorEndTime,omitempty" name:"MonitorEndTime"`

	// 申请人ID
	ApplierId *string `json:"ApplierId,omitempty" name:"ApplierId"`

	// 申请人姓名
	ApplierName *string `json:"ApplierName,omitempty" name:"ApplierName"`
}

type CreateCRWorkRequest struct {
	*tchttp.BaseRequest
	
	// 作品名称
	WorkName *string `json:"WorkName,omitempty" name:"WorkName"`

	// 作品分类
	WorkCategory *string `json:"WorkCategory,omitempty" name:"WorkCategory"`

	// 作品内容类型
	WorkType *string `json:"WorkType,omitempty" name:"WorkType"`

	// 作品标签
	WorkSign *string `json:"WorkSign,omitempty" name:"WorkSign"`

	// 作品图片
	WorkPic *string `json:"WorkPic,omitempty" name:"WorkPic"`

	// 创作描述
	WorkDesc *string `json:"WorkDesc,omitempty" name:"WorkDesc"`

	// 是否原创 0:否 1:是
	IsOriginal *string `json:"IsOriginal,omitempty" name:"IsOriginal"`

	// 是否发布 0：未发布 1：已发布
	IsRelease *string `json:"IsRelease,omitempty" name:"IsRelease"`

	// 著作权人ID
	ProducerID *int64 `json:"ProducerID,omitempty" name:"ProducerID"`

	// 创作时间
	ProduceTime *string `json:"ProduceTime,omitempty" name:"ProduceTime"`

	// 样品文件路径
	SampleContentURL *string `json:"SampleContentURL,omitempty" name:"SampleContentURL"`

	// 样本下载Url
	SampleDownloadURL *string `json:"SampleDownloadURL,omitempty" name:"SampleDownloadURL"`

	// 授予类型
	GrantType *string `json:"GrantType,omitempty" name:"GrantType"`

	// 作品发布Url
	SamplePublicURL *string `json:"SamplePublicURL,omitempty" name:"SamplePublicURL"`

	// 是否启用监测 0：不启用 1：启用 默认为0
	IsMonitor *string `json:"IsMonitor,omitempty" name:"IsMonitor"`

	// 是否启用存证0：不存证  2：存证 默认为0
	IsCert *string `json:"IsCert,omitempty" name:"IsCert"`

	// 存证回调地址
	CertUrl *string `json:"CertUrl,omitempty" name:"CertUrl"`

	// 监测回调地址
	MonitorUrl *string `json:"MonitorUrl,omitempty" name:"MonitorUrl"`

	// 创作性质（原创,改编,翻译,汇编,注释,整理,其他)
	ProduceType *string `json:"ProduceType,omitempty" name:"ProduceType"`

	// 白名单
	WhiteLists []*string `json:"WhiteLists,omitempty" name:"WhiteLists"`

	// 作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 著作权人姓名
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// 作者
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 授权书
	Authorization *string `json:"Authorization,omitempty" name:"Authorization"`

	// 授权书开始时间
	AuthorizationStartTime *string `json:"AuthorizationStartTime,omitempty" name:"AuthorizationStartTime"`

	// 授权书结束时间
	AuthorizationEndTime *string `json:"AuthorizationEndTime,omitempty" name:"AuthorizationEndTime"`

	// 内容格式
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// 文件内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 监测结束时间
	MonitorEndTime *string `json:"MonitorEndTime,omitempty" name:"MonitorEndTime"`

	// 申请人ID
	ApplierId *string `json:"ApplierId,omitempty" name:"ApplierId"`

	// 申请人姓名
	ApplierName *string `json:"ApplierName,omitempty" name:"ApplierName"`
}

func (r *CreateCRWorkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRWorkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkName")
	delete(f, "WorkCategory")
	delete(f, "WorkType")
	delete(f, "WorkSign")
	delete(f, "WorkPic")
	delete(f, "WorkDesc")
	delete(f, "IsOriginal")
	delete(f, "IsRelease")
	delete(f, "ProducerID")
	delete(f, "ProduceTime")
	delete(f, "SampleContentURL")
	delete(f, "SampleDownloadURL")
	delete(f, "GrantType")
	delete(f, "SamplePublicURL")
	delete(f, "IsMonitor")
	delete(f, "IsCert")
	delete(f, "CertUrl")
	delete(f, "MonitorUrl")
	delete(f, "ProduceType")
	delete(f, "WhiteLists")
	delete(f, "WorkId")
	delete(f, "ProducerName")
	delete(f, "Nickname")
	delete(f, "Authorization")
	delete(f, "AuthorizationStartTime")
	delete(f, "AuthorizationEndTime")
	delete(f, "ContentType")
	delete(f, "Content")
	delete(f, "MonitorEndTime")
	delete(f, "ApplierId")
	delete(f, "ApplierName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCRWorkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCRWorkResponseParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// x
	EvidenceId *int64 `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCRWorkResponse struct {
	*tchttp.BaseResponse
	Response *CreateCRWorkResponseParams `json:"Response"`
}

func (r *CreateCRWorkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCRWorkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPCompanyInfoRequestParams struct {

}

type DescribeBPCompanyInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBPCompanyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPCompanyInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBPCompanyInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPCompanyInfoResponseParams struct {
	// xx
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// xx
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// xx
	LicenseName *string `json:"LicenseName,omitempty" name:"LicenseName"`

	// xx
	LicenseStatus *int64 `json:"LicenseStatus,omitempty" name:"LicenseStatus"`

	// xx
	LicenseNote *string `json:"LicenseNote,omitempty" name:"LicenseNote"`

	// xx
	AuthorizationName *string `json:"AuthorizationName,omitempty" name:"AuthorizationName"`

	// xx
	AuthorizationStatus *int64 `json:"AuthorizationStatus,omitempty" name:"AuthorizationStatus"`

	// xx
	AuthorizationNote *string `json:"AuthorizationNote,omitempty" name:"AuthorizationNote"`

	// xx
	BrandDatas []*BrandData `json:"BrandDatas,omitempty" name:"BrandDatas"`

	// xx
	CompanyId *int64 `json:"CompanyId,omitempty" name:"CompanyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBPCompanyInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBPCompanyInfoResponseParams `json:"Response"`
}

func (r *DescribeBPCompanyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPCompanyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPFakeURLsRequestParams struct {
	// xxx
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// xxx
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// xxx
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

type DescribeBPFakeURLsRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// xxx
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// xxx
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

func (r *DescribeBPFakeURLsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPFakeURLsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBPFakeURLsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPFakeURLsResponseParams struct {
	// xxx
	FakeURLInfos []*FakeURLInfo `json:"FakeURLInfos,omitempty" name:"FakeURLInfos"`

	// xxx
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// xxx
	ExportURL *string `json:"ExportURL,omitempty" name:"ExportURL"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBPFakeURLsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBPFakeURLsResponseParams `json:"Response"`
}

func (r *DescribeBPFakeURLsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPFakeURLsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPProtectURLsRequestParams struct {
	// xxx
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// xxx
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

type DescribeBPProtectURLsRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// xxx
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

func (r *DescribeBPProtectURLsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPProtectURLsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBPProtectURLsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPProtectURLsResponseParams struct {
	// xxx
	ProtectURLInfos []*ProtectURLInfo `json:"ProtectURLInfos,omitempty" name:"ProtectURLInfos"`

	// xxx
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBPProtectURLsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBPProtectURLsResponseParams `json:"Response"`
}

func (r *DescribeBPProtectURLsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPProtectURLsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPReportFakeURLsRequestParams struct {
	// xxx
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// xxx
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// xxx
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

type DescribeBPReportFakeURLsRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// xxx
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// xxx
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

func (r *DescribeBPReportFakeURLsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPReportFakeURLsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBPReportFakeURLsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPReportFakeURLsResponseParams struct {
	// xxx
	ReportFakeURLInfos []*ReportFakeURLInfo `json:"ReportFakeURLInfos,omitempty" name:"ReportFakeURLInfos"`

	// xxx
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBPReportFakeURLsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBPReportFakeURLsResponseParams `json:"Response"`
}

func (r *DescribeBPReportFakeURLsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPReportFakeURLsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRMonitorDetailRequestParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 页数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// x
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeCRMonitorDetailRequest struct {
	*tchttp.BaseRequest
	
	// 作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 页数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// x
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCRMonitorDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRMonitorDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCRMonitorDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRMonitorDetailResponseParams struct {
	// MonitorTort数组
	Torts []*MonitorTort `json:"Torts,omitempty" name:"Torts"`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// x
	MonitorStatus *int64 `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCRMonitorDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCRMonitorDetailResponseParams `json:"Response"`
}

func (r *DescribeCRMonitorDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRMonitorDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRMonitorsRequestParams struct {
	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

type DescribeCRMonitorsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

func (r *DescribeCRMonitorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRMonitorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCRMonitorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCRMonitorsResponseParams struct {
	// 监测结果
	Monitors []*Monitor `json:"Monitors,omitempty" name:"Monitors"`

	// 记录总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCRMonitorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCRMonitorsResponseParams `json:"Response"`
}

func (r *DescribeCRMonitorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCRMonitorsResponse) FromJsonString(s string) error {
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

	// xxx
	EvidenceStatus *int64 `json:"EvidenceStatus,omitempty" name:"EvidenceStatus"`

	// xxx
	WorkCategory *string `json:"WorkCategory,omitempty" name:"WorkCategory"`

	// xxx
	IsOriginal *string `json:"IsOriginal,omitempty" name:"IsOriginal"`

	// xxx
	IsRelease *string `json:"IsRelease,omitempty" name:"IsRelease"`

	// xxx
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// xxx
	ProduceTime *string `json:"ProduceTime,omitempty" name:"ProduceTime"`

	// xxx
	WhiteLists []*string `json:"WhiteLists,omitempty" name:"WhiteLists"`

	// xxx
	WorkDesc *string `json:"WorkDesc,omitempty" name:"WorkDesc"`

	// xxx
	Authorization *string `json:"Authorization,omitempty" name:"Authorization"`

	// xxx
	AuthorizationStartTime *string `json:"AuthorizationStartTime,omitempty" name:"AuthorizationStartTime"`

	// xxx
	AuthorizationEndTime *string `json:"AuthorizationEndTime,omitempty" name:"AuthorizationEndTime"`

	// xxx
	Commission *string `json:"Commission,omitempty" name:"Commission"`

	// xxx
	CommissionStartTime *string `json:"CommissionStartTime,omitempty" name:"CommissionStartTime"`

	// xxx
	CommissionEndTime *string `json:"CommissionEndTime,omitempty" name:"CommissionEndTime"`

	// xxx
	EvidenceUrl *string `json:"EvidenceUrl,omitempty" name:"EvidenceUrl"`

	// xxx
	EvidenceStartTime *string `json:"EvidenceStartTime,omitempty" name:"EvidenceStartTime"`

	// xxx
	EvidenceEndTime *string `json:"EvidenceEndTime,omitempty" name:"EvidenceEndTime"`

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

type FakeURLInfo struct {
	// xxx
	FakeURLId *int64 `json:"FakeURLId,omitempty" name:"FakeURLId"`

	// xxx
	ProtectWeb *string `json:"ProtectWeb,omitempty" name:"ProtectWeb"`

	// xxx
	DetectTime *string `json:"DetectTime,omitempty" name:"DetectTime"`

	// xxx
	FakeURL *string `json:"FakeURL,omitempty" name:"FakeURL"`

	// xxx
	Snapshot *string `json:"Snapshot,omitempty" name:"Snapshot"`

	// xxx
	IP *string `json:"IP,omitempty" name:"IP"`

	// xxx
	IPLoc *string `json:"IPLoc,omitempty" name:"IPLoc"`

	// xxx
	Heat *int64 `json:"Heat,omitempty" name:"Heat"`

	// xxx
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// xxx
	Note *string `json:"Note,omitempty" name:"Note"`

	// xxx
	FakeURLCompany *string `json:"FakeURLCompany,omitempty" name:"FakeURLCompany"`

	// xxx
	FakeURLAttr *string `json:"FakeURLAttr,omitempty" name:"FakeURLAttr"`

	// xxx
	FakeURLName *string `json:"FakeURLName,omitempty" name:"FakeURLName"`

	// xxx
	FakeURLICP *string `json:"FakeURLICP,omitempty" name:"FakeURLICP"`

	// xxx
	FakeURLCreateTime *string `json:"FakeURLCreateTime,omitempty" name:"FakeURLCreateTime"`

	// xxx
	FakeURLExpireTime *string `json:"FakeURLExpireTime,omitempty" name:"FakeURLExpireTime"`
}

type File struct {
	// xxx
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// xxx
	FileType *int64 `json:"FileType,omitempty" name:"FileType"`

	// xxx
	ValidStartDate *string `json:"ValidStartDate,omitempty" name:"ValidStartDate"`

	// xxx
	ValidEndDate *string `json:"ValidEndDate,omitempty" name:"ValidEndDate"`
}

type Filter struct {

}

// Predefined struct for user
type ModifyBPOfflineAttachmentRequestParams struct {
	// xx
	LicenseName *string `json:"LicenseName,omitempty" name:"LicenseName"`

	// xx
	AuthorizationName *string `json:"AuthorizationName,omitempty" name:"AuthorizationName"`

	// xx
	BrandName *string `json:"BrandName,omitempty" name:"BrandName"`

	// xx
	BrandCertificateName *string `json:"BrandCertificateName,omitempty" name:"BrandCertificateName"`

	// xx
	TransferName *string `json:"TransferName,omitempty" name:"TransferName"`
}

type ModifyBPOfflineAttachmentRequest struct {
	*tchttp.BaseRequest
	
	// xx
	LicenseName *string `json:"LicenseName,omitempty" name:"LicenseName"`

	// xx
	AuthorizationName *string `json:"AuthorizationName,omitempty" name:"AuthorizationName"`

	// xx
	BrandName *string `json:"BrandName,omitempty" name:"BrandName"`

	// xx
	BrandCertificateName *string `json:"BrandCertificateName,omitempty" name:"BrandCertificateName"`

	// xx
	TransferName *string `json:"TransferName,omitempty" name:"TransferName"`
}

func (r *ModifyBPOfflineAttachmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBPOfflineAttachmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseName")
	delete(f, "AuthorizationName")
	delete(f, "BrandName")
	delete(f, "BrandCertificateName")
	delete(f, "TransferName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBPOfflineAttachmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBPOfflineAttachmentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBPOfflineAttachmentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBPOfflineAttachmentResponseParams `json:"Response"`
}

func (r *ModifyBPOfflineAttachmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBPOfflineAttachmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRBlockStatusRequestParams struct {
	// 侵权ID
	TortId *int64 `json:"TortId,omitempty" name:"TortId"`

	// 拦截结果回调地址
	BlockUrl *string `json:"BlockUrl,omitempty" name:"BlockUrl"`
}

type ModifyCRBlockStatusRequest struct {
	*tchttp.BaseRequest
	
	// 侵权ID
	TortId *int64 `json:"TortId,omitempty" name:"TortId"`

	// 拦截结果回调地址
	BlockUrl *string `json:"BlockUrl,omitempty" name:"BlockUrl"`
}

func (r *ModifyCRBlockStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRBlockStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TortId")
	delete(f, "BlockUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCRBlockStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRBlockStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCRBlockStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCRBlockStatusResponseParams `json:"Response"`
}

func (r *ModifyCRBlockStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRBlockStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRMonitorRequestParams struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 监测状态 1-开启监测 2-关闭监测
	MonitorStatus *string `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

	// 默认不停止，支持续期
	MonitorEnd *string `json:"MonitorEnd,omitempty" name:"MonitorEnd"`
}

type ModifyCRMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 监测状态 1-开启监测 2-关闭监测
	MonitorStatus *string `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

	// 默认不停止，支持续期
	MonitorEnd *string `json:"MonitorEnd,omitempty" name:"MonitorEnd"`
}

func (r *ModifyCRMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "MonitorStatus")
	delete(f, "MonitorEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCRMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRMonitorResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCRMonitorResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCRMonitorResponseParams `json:"Response"`
}

func (r *ModifyCRMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRObtainStatusRequestParams struct {
	// xxx
	TortId *int64 `json:"TortId,omitempty" name:"TortId"`

	// xxx
	ObtainType *int64 `json:"ObtainType,omitempty" name:"ObtainType"`

	// xxx
	ObtainDuration *int64 `json:"ObtainDuration,omitempty" name:"ObtainDuration"`

	// xxx
	ObtainUrl *string `json:"ObtainUrl,omitempty" name:"ObtainUrl"`
}

type ModifyCRObtainStatusRequest struct {
	*tchttp.BaseRequest
	
	// xxx
	TortId *int64 `json:"TortId,omitempty" name:"TortId"`

	// xxx
	ObtainType *int64 `json:"ObtainType,omitempty" name:"ObtainType"`

	// xxx
	ObtainDuration *int64 `json:"ObtainDuration,omitempty" name:"ObtainDuration"`

	// xxx
	ObtainUrl *string `json:"ObtainUrl,omitempty" name:"ObtainUrl"`
}

func (r *ModifyCRObtainStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRObtainStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TortId")
	delete(f, "ObtainType")
	delete(f, "ObtainDuration")
	delete(f, "ObtainUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCRObtainStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRObtainStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCRObtainStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCRObtainStatusResponseParams `json:"Response"`
}

func (r *ModifyCRObtainStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRObtainStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRRightStatusRequestParams struct {
	// 侵权ID
	TortId *int64 `json:"TortId,omitempty" name:"TortId"`

	// 发函结果回调地址
	RightUrl *string `json:"RightUrl,omitempty" name:"RightUrl"`
}

type ModifyCRRightStatusRequest struct {
	*tchttp.BaseRequest
	
	// 侵权ID
	TortId *int64 `json:"TortId,omitempty" name:"TortId"`

	// 发函结果回调地址
	RightUrl *string `json:"RightUrl,omitempty" name:"RightUrl"`
}

func (r *ModifyCRRightStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRRightStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TortId")
	delete(f, "RightUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCRRightStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRRightStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCRRightStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCRRightStatusResponseParams `json:"Response"`
}

func (r *ModifyCRRightStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRRightStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRWhiteListRequestParams struct {
	// 白名单ID
	WhiteListId *int64 `json:"WhiteListId,omitempty" name:"WhiteListId"`

	// 平台名称
	PlatForm *string `json:"PlatForm,omitempty" name:"PlatForm"`

	// 平台站点链接
	PlatUrl *string `json:"PlatUrl,omitempty" name:"PlatUrl"`

	// 作者ID
	AuthorId *string `json:"AuthorId,omitempty" name:"AuthorId"`

	// 作品ID
	WorksId *int64 `json:"WorksId,omitempty" name:"WorksId"`

	// xxx
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// xxx
	WhiteSites *string `json:"WhiteSites,omitempty" name:"WhiteSites"`
}

type ModifyCRWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 白名单ID
	WhiteListId *int64 `json:"WhiteListId,omitempty" name:"WhiteListId"`

	// 平台名称
	PlatForm *string `json:"PlatForm,omitempty" name:"PlatForm"`

	// 平台站点链接
	PlatUrl *string `json:"PlatUrl,omitempty" name:"PlatUrl"`

	// 作者ID
	AuthorId *string `json:"AuthorId,omitempty" name:"AuthorId"`

	// 作品ID
	WorksId *int64 `json:"WorksId,omitempty" name:"WorksId"`

	// xxx
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// xxx
	WhiteSites *string `json:"WhiteSites,omitempty" name:"WhiteSites"`
}

func (r *ModifyCRWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WhiteListId")
	delete(f, "PlatForm")
	delete(f, "PlatUrl")
	delete(f, "AuthorId")
	delete(f, "WorksId")
	delete(f, "WorkId")
	delete(f, "WhiteSites")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCRWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCRWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCRWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCRWhiteListResponseParams `json:"Response"`
}

func (r *ModifyCRWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCRWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Monitor struct {
	// 作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 作品名称
	WorkName *string `json:"WorkName,omitempty" name:"WorkName"`

	// 作品内容类型 01-视频 02-音频 03-文本 04-图片
	WorkType *string `json:"WorkType,omitempty" name:"WorkType"`

	// 侵权平台数量
	TortPlatNum *int64 `json:"TortPlatNum,omitempty" name:"TortPlatNum"`

	// 侵权链接数量
	TortURLNum *int64 `json:"TortURLNum,omitempty" name:"TortURLNum"`

	// 监测时间
	MonitorTime *string `json:"MonitorTime,omitempty" name:"MonitorTime"`

	// 0-待监测 1-监测中 2-不监测 3-暂停监测
	MonitorStatus *int64 `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

	// 作品类型
	WorkCategory *string `json:"WorkCategory,omitempty" name:"WorkCategory"`

	// xx
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

	// xx
	MonitorNote *string `json:"MonitorNote,omitempty" name:"MonitorNote"`
}

type MonitorTort struct {
	// 侵权信息ID
	TortId *int64 `json:"TortId,omitempty" name:"TortId"`

	// 侵权标题
	TortTitle *string `json:"TortTitle,omitempty" name:"TortTitle"`

	// 侵权平台
	TortPlat *string `json:"TortPlat,omitempty" name:"TortPlat"`

	// 侵权链接
	TortURL *string `json:"TortURL,omitempty" name:"TortURL"`

	// 侵权链接发布时间
	PubTime *string `json:"PubTime,omitempty" name:"PubTime"`

	// 作者
	Author *string `json:"Author,omitempty" name:"Author"`

	// xxx
	DetectTime *string `json:"DetectTime,omitempty" name:"DetectTime"`

	// 1
	ObtainStatus *int64 `json:"ObtainStatus,omitempty" name:"ObtainStatus"`

	// 1
	RightStatus *int64 `json:"RightStatus,omitempty" name:"RightStatus"`

	// 1
	BlockStatus *int64 `json:"BlockStatus,omitempty" name:"BlockStatus"`

	// 1
	TortNum *string `json:"TortNum,omitempty" name:"TortNum"`

	// 1
	ObtainNote *string `json:"ObtainNote,omitempty" name:"ObtainNote"`

	// 1
	WorkTitle *string `json:"WorkTitle,omitempty" name:"WorkTitle"`

	// 1
	TortSite *string `json:"TortSite,omitempty" name:"TortSite"`

	// 1
	ICP *string `json:"ICP,omitempty" name:"ICP"`

	// 1
	RightNote *string `json:"RightNote,omitempty" name:"RightNote"`

	// 1
	ObtainType *int64 `json:"ObtainType,omitempty" name:"ObtainType"`

	// 1
	BlockNote *string `json:"BlockNote,omitempty" name:"BlockNote"`

	// 1
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 1
	WorkName *string `json:"WorkName,omitempty" name:"WorkName"`

	// 1
	AuthStatus *int64 `json:"AuthStatus,omitempty" name:"AuthStatus"`

	// 1
	CommStatus *int64 `json:"CommStatus,omitempty" name:"CommStatus"`

	// 1
	EvidenceStatus *int64 `json:"EvidenceStatus,omitempty" name:"EvidenceStatus"`

	// 1
	IsProducer *int64 `json:"IsProducer,omitempty" name:"IsProducer"`

	// 1
	IsOverseas *int64 `json:"IsOverseas,omitempty" name:"IsOverseas"`
}

type ProtectURLInfo struct {
	// xxx
	ProtectURLId *int64 `json:"ProtectURLId,omitempty" name:"ProtectURLId"`

	// xxx
	ProtectURL *string `json:"ProtectURL,omitempty" name:"ProtectURL"`

	// xxx
	ProtectWeb *string `json:"ProtectWeb,omitempty" name:"ProtectWeb"`

	// xxx
	ProtectURLStatus *int64 `json:"ProtectURLStatus,omitempty" name:"ProtectURLStatus"`

	// xxx
	ProtectURLNote *string `json:"ProtectURLNote,omitempty" name:"ProtectURLNote"`

	// xxx
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ReportFakeURLInfo struct {
	// xxx
	FakeURLId *int64 `json:"FakeURLId,omitempty" name:"FakeURLId"`

	// xxx
	DetectTime *string `json:"DetectTime,omitempty" name:"DetectTime"`

	// xxx
	ProtectURL *string `json:"ProtectURL,omitempty" name:"ProtectURL"`

	// xxx
	ProtectWeb *string `json:"ProtectWeb,omitempty" name:"ProtectWeb"`

	// xxx
	FakeURL *string `json:"FakeURL,omitempty" name:"FakeURL"`

	// xxx
	Snapshot *string `json:"Snapshot,omitempty" name:"Snapshot"`

	// xxx
	IP *string `json:"IP,omitempty" name:"IP"`

	// xxx
	IPLoc *string `json:"IPLoc,omitempty" name:"IPLoc"`

	// xxx
	Heat *int64 `json:"Heat,omitempty" name:"Heat"`

	// xxx
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// xxx
	Note *string `json:"Note,omitempty" name:"Note"`

	// xxx
	FakeURLCompany *string `json:"FakeURLCompany,omitempty" name:"FakeURLCompany"`

	// xxx
	FakeURLAttr *string `json:"FakeURLAttr,omitempty" name:"FakeURLAttr"`

	// xxx
	FakeURLName *string `json:"FakeURLName,omitempty" name:"FakeURLName"`

	// xxx
	FakeURLICP *string `json:"FakeURLICP,omitempty" name:"FakeURLICP"`

	// xxx
	FakeURLCreateTime *string `json:"FakeURLCreateTime,omitempty" name:"FakeURLCreateTime"`

	// xxx
	FakeURLExpireTime *string `json:"FakeURLExpireTime,omitempty" name:"FakeURLExpireTime"`

	// xxx
	BlockTime *string `json:"BlockTime,omitempty" name:"BlockTime"`
}

// Predefined struct for user
type UpdateCRWorkRequestParams struct {
	// xx
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// xx
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// xx
	Content *string `json:"Content,omitempty" name:"Content"`

	// xx
	CertType *string `json:"CertType,omitempty" name:"CertType"`
}

type UpdateCRWorkRequest struct {
	*tchttp.BaseRequest
	
	// xx
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// xx
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// xx
	Content *string `json:"Content,omitempty" name:"Content"`

	// xx
	CertType *string `json:"CertType,omitempty" name:"CertType"`
}

func (r *UpdateCRWorkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCRWorkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkId")
	delete(f, "ContentType")
	delete(f, "Content")
	delete(f, "CertType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCRWorkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCRWorkResponseParams struct {
	// xx
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// xx
	EvidenceId *int64 `json:"EvidenceId,omitempty" name:"EvidenceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateCRWorkResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCRWorkResponseParams `json:"Response"`
}

func (r *UpdateCRWorkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCRWorkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}