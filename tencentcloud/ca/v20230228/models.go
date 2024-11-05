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

package v20230228

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CertificateIdentityUser struct {
	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 唯一身份id
	IdentityUniqueId *string `json:"IdentityUniqueId,omitnil,omitempty" name:"IdentityUniqueId"`

	// 身份证号
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 身份鉴别类型
	// 1：授权金融机构身份鉴别
	IdentificationType *string `json:"IdentificationType,omitnil,omitempty" name:"IdentificationType"`

	// 身份鉴别措施
	// 1、身份证鉴别
	// 2、银行卡鉴别
	// 3、支付账户密码验证
	// 4、人脸识别验证
	IdentificationMeasures []*string `json:"IdentificationMeasures,omitnil,omitempty" name:"IdentificationMeasures"`
}

// Predefined struct for user
type CreateVerifyReportRequestParams struct {
	// 申请者类型 1:个人，2:企业
	ApplyCustomerType *string `json:"ApplyCustomerType,omitnil,omitempty" name:"ApplyCustomerType"`

	// 申请企业 or 自然人名称
	ApplyCustomerName *string `json:"ApplyCustomerName,omitnil,omitempty" name:"ApplyCustomerName"`

	// 验签申请经办人姓名
	ApplyName *string `json:"ApplyName,omitnil,omitempty" name:"ApplyName"`

	// 验签申请经办人电话
	ApplyMobile *string `json:"ApplyMobile,omitnil,omitempty" name:"ApplyMobile"`

	// 验签文件id
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 验签申请经办人邮箱
	ApplyEmail *string `json:"ApplyEmail,omitnil,omitempty" name:"ApplyEmail"`

	// 证书用户身份及身份鉴别信息
	CertificateIdentityUsers []*CertificateIdentityUser `json:"CertificateIdentityUsers,omitnil,omitempty" name:"CertificateIdentityUsers"`
}

type CreateVerifyReportRequest struct {
	*tchttp.BaseRequest
	
	// 申请者类型 1:个人，2:企业
	ApplyCustomerType *string `json:"ApplyCustomerType,omitnil,omitempty" name:"ApplyCustomerType"`

	// 申请企业 or 自然人名称
	ApplyCustomerName *string `json:"ApplyCustomerName,omitnil,omitempty" name:"ApplyCustomerName"`

	// 验签申请经办人姓名
	ApplyName *string `json:"ApplyName,omitnil,omitempty" name:"ApplyName"`

	// 验签申请经办人电话
	ApplyMobile *string `json:"ApplyMobile,omitnil,omitempty" name:"ApplyMobile"`

	// 验签文件id
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 验签申请经办人邮箱
	ApplyEmail *string `json:"ApplyEmail,omitnil,omitempty" name:"ApplyEmail"`

	// 证书用户身份及身份鉴别信息
	CertificateIdentityUsers []*CertificateIdentityUser `json:"CertificateIdentityUsers,omitnil,omitempty" name:"CertificateIdentityUsers"`
}

func (r *CreateVerifyReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVerifyReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplyCustomerType")
	delete(f, "ApplyCustomerName")
	delete(f, "ApplyName")
	delete(f, "ApplyMobile")
	delete(f, "FileId")
	delete(f, "ApplyEmail")
	delete(f, "CertificateIdentityUsers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVerifyReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVerifyReportResponseParams struct {
	// 签名id
	SignatureId *string `json:"SignatureId,omitnil,omitempty" name:"SignatureId"`

	// code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// message
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVerifyReportResponse struct {
	*tchttp.BaseResponse
	Response *CreateVerifyReportResponseParams `json:"Response"`
}

func (r *CreateVerifyReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVerifyReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVerifyReportRequestParams struct {
	// 签名id
	SignatureId *string `json:"SignatureId,omitnil,omitempty" name:"SignatureId"`
}

type DescribeVerifyReportRequest struct {
	*tchttp.BaseRequest
	
	// 签名id
	SignatureId *string `json:"SignatureId,omitnil,omitempty" name:"SignatureId"`
}

func (r *DescribeVerifyReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVerifyReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SignatureId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVerifyReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVerifyReportResponseParams struct {
	// 下载url
	ReportUrl *string `json:"ReportUrl,omitnil,omitempty" name:"ReportUrl"`

	// code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// message
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVerifyReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVerifyReportResponseParams `json:"Response"`
}

func (r *DescribeVerifyReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVerifyReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileInfo struct {
	// BASE64编码后的文件内容
	FileBody *string `json:"FileBody,omitnil,omitempty" name:"FileBody"`

	// 文件名及类型，最大长度不超过200字符
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

// Predefined struct for user
type UploadFileRequestParams struct {
	// 验签源文件信息列表
	FileInfos []*FileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`
}

type UploadFileRequest struct {
	*tchttp.BaseRequest
	
	// 验签源文件信息列表
	FileInfos []*FileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`
}

func (r *UploadFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFileResponseParams struct {
	// 文件id列表
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 文件id总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadFileResponse struct {
	*tchttp.BaseResponse
	Response *UploadFileResponseParams `json:"Response"`
}

func (r *UploadFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}