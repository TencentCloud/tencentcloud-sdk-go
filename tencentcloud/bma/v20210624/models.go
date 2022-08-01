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

	// 是否启用存证0：不启用 1：启用 默认为1
	IsCert *string `json:"IsCert,omitempty" name:"IsCert"`

	// 存证回调地址
	CertUrl *string `json:"CertUrl,omitempty" name:"CertUrl"`

	// 监测回调地址
	MonitorUrl *string `json:"MonitorUrl,omitempty" name:"MonitorUrl"`

	// 创作性质（原创,改编,翻译,汇编,注释,整理,其他)
	ProduceType *string `json:"ProduceType,omitempty" name:"ProduceType"`

	// xxx
	WhiteLists []*string `json:"WhiteLists,omitempty" name:"WhiteLists"`

	// x
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// xx
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// xx
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// x
	Authorization *string `json:"Authorization,omitempty" name:"Authorization"`

	// x
	AuthorizationStartTime *string `json:"AuthorizationStartTime,omitempty" name:"AuthorizationStartTime"`

	// x
	AuthorizationEndTime *string `json:"AuthorizationEndTime,omitempty" name:"AuthorizationEndTime"`

	// x
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// x
	Content *string `json:"Content,omitempty" name:"Content"`
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

	// 是否启用存证0：不启用 1：启用 默认为1
	IsCert *string `json:"IsCert,omitempty" name:"IsCert"`

	// 存证回调地址
	CertUrl *string `json:"CertUrl,omitempty" name:"CertUrl"`

	// 监测回调地址
	MonitorUrl *string `json:"MonitorUrl,omitempty" name:"MonitorUrl"`

	// 创作性质（原创,改编,翻译,汇编,注释,整理,其他)
	ProduceType *string `json:"ProduceType,omitempty" name:"ProduceType"`

	// xxx
	WhiteLists []*string `json:"WhiteLists,omitempty" name:"WhiteLists"`

	// x
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// xx
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// xx
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// x
	Authorization *string `json:"Authorization,omitempty" name:"Authorization"`

	// x
	AuthorizationStartTime *string `json:"AuthorizationStartTime,omitempty" name:"AuthorizationStartTime"`

	// x
	AuthorizationEndTime *string `json:"AuthorizationEndTime,omitempty" name:"AuthorizationEndTime"`

	// x
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// x
	Content *string `json:"Content,omitempty" name:"Content"`
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

type Filter struct {

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

	// 监测状态 1-开启监测 2-关闭监测 默认为1
	MonitorStatus *string `json:"MonitorStatus,omitempty" name:"MonitorStatus"`
}

type ModifyCRMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 作品ID
	WorkId *int64 `json:"WorkId,omitempty" name:"WorkId"`

	// 监测状态 1-开启监测 2-关闭监测 默认为1
	MonitorStatus *string `json:"MonitorStatus,omitempty" name:"MonitorStatus"`
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