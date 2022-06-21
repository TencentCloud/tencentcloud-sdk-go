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

package v20181115

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddFairPlayPemRequestParams struct {
	// 加密后的fairplay方案申请时使用的私钥。
	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对私钥文件中的字段进行加密，并对加密结果进行base64编码。
	Pem *string `json:"Pem,omitempty" name:"Pem"`

	// 加密后的fairplay方案申请返回的ask数据。
	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对Ask字符串进行加密，并对加密结果进行base64编码。
	Ask *string `json:"Ask,omitempty" name:"Ask"`

	// 私钥的解密密钥。
	// openssl在生成rsa时，可能会需要设置加密密钥，请记住设置的密钥。
	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对解密密钥进行加密，并对加密结果进行base64编码。
	PemDecryptKey *string `json:"PemDecryptKey,omitempty" name:"PemDecryptKey"`

	// 委托者Id,适用于托管自身证书的客户。普通客户无需填该字段。
	BailorId *uint64 `json:"BailorId,omitempty" name:"BailorId"`

	// 私钥的优先级，优先级数值越高，优先级越高。
	// 该值可以不传，后台将自动分配一个优先级。
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
}

type AddFairPlayPemRequest struct {
	*tchttp.BaseRequest
	
	// 加密后的fairplay方案申请时使用的私钥。
	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对私钥文件中的字段进行加密，并对加密结果进行base64编码。
	Pem *string `json:"Pem,omitempty" name:"Pem"`

	// 加密后的fairplay方案申请返回的ask数据。
	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对Ask字符串进行加密，并对加密结果进行base64编码。
	Ask *string `json:"Ask,omitempty" name:"Ask"`

	// 私钥的解密密钥。
	// openssl在生成rsa时，可能会需要设置加密密钥，请记住设置的密钥。
	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对解密密钥进行加密，并对加密结果进行base64编码。
	PemDecryptKey *string `json:"PemDecryptKey,omitempty" name:"PemDecryptKey"`

	// 委托者Id,适用于托管自身证书的客户。普通客户无需填该字段。
	BailorId *uint64 `json:"BailorId,omitempty" name:"BailorId"`

	// 私钥的优先级，优先级数值越高，优先级越高。
	// 该值可以不传，后台将自动分配一个优先级。
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *AddFairPlayPemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddFairPlayPemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Pem")
	delete(f, "Ask")
	delete(f, "PemDecryptKey")
	delete(f, "BailorId")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddFairPlayPemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddFairPlayPemResponseParams struct {
	// 设置私钥后，后台返回的pem id，用来唯一标识一个私钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FairPlayPemId *uint64 `json:"FairPlayPemId,omitempty" name:"FairPlayPemId"`

	// 私钥的优先级，优先级数值越高，优先级越高。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddFairPlayPemResponse struct {
	*tchttp.BaseResponse
	Response *AddFairPlayPemResponseParams `json:"Response"`
}

func (r *AddFairPlayPemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddFairPlayPemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEncryptKeysRequestParams struct {
	// 使用的DRM方案类型，接口取值WIDEVINE、FAIRPLAY、NORMALAES。
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// 设置的加密密钥列表。
	Keys []*KeyParam `json:"Keys,omitempty" name:"Keys"`

	// 一个加密内容的唯一标识。
	ContentId *string `json:"ContentId,omitempty" name:"ContentId"`

	// 内容类型。接口取值VodVideo,LiveVideo。
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
}

type CreateEncryptKeysRequest struct {
	*tchttp.BaseRequest
	
	// 使用的DRM方案类型，接口取值WIDEVINE、FAIRPLAY、NORMALAES。
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// 设置的加密密钥列表。
	Keys []*KeyParam `json:"Keys,omitempty" name:"Keys"`

	// 一个加密内容的唯一标识。
	ContentId *string `json:"ContentId,omitempty" name:"ContentId"`

	// 内容类型。接口取值VodVideo,LiveVideo。
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
}

func (r *CreateEncryptKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEncryptKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DrmType")
	delete(f, "Keys")
	delete(f, "ContentId")
	delete(f, "ContentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEncryptKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEncryptKeysResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEncryptKeysResponse struct {
	*tchttp.BaseResponse
	Response *CreateEncryptKeysResponseParams `json:"Response"`
}

func (r *CreateEncryptKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEncryptKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLicenseRequestParams struct {
	// DRM方案类型，接口取值：WIDEVINE，FAIRPLAY。
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// Base64编码的终端设备License Request数据。
	LicenseRequest *string `json:"LicenseRequest,omitempty" name:"LicenseRequest"`

	// 内容类型，接口取值：VodVideo,LiveVideo。
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// 授权播放的Track列表。
	// 该值为空时，默认授权所有track播放。
	Tracks []*string `json:"Tracks,omitempty" name:"Tracks"`

	// 播放策略参数。
	PlaybackPolicy *PlaybackPolicy `json:"PlaybackPolicy,omitempty" name:"PlaybackPolicy"`
}

type CreateLicenseRequest struct {
	*tchttp.BaseRequest
	
	// DRM方案类型，接口取值：WIDEVINE，FAIRPLAY。
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// Base64编码的终端设备License Request数据。
	LicenseRequest *string `json:"LicenseRequest,omitempty" name:"LicenseRequest"`

	// 内容类型，接口取值：VodVideo,LiveVideo。
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// 授权播放的Track列表。
	// 该值为空时，默认授权所有track播放。
	Tracks []*string `json:"Tracks,omitempty" name:"Tracks"`

	// 播放策略参数。
	PlaybackPolicy *PlaybackPolicy `json:"PlaybackPolicy,omitempty" name:"PlaybackPolicy"`
}

func (r *CreateLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DrmType")
	delete(f, "LicenseRequest")
	delete(f, "ContentType")
	delete(f, "Tracks")
	delete(f, "PlaybackPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLicenseResponseParams struct {
	// Base64 编码的许可证二进制数据。
	License *string `json:"License,omitempty" name:"License"`

	// 加密内容的内容ID
	ContentId *string `json:"ContentId,omitempty" name:"ContentId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLicenseResponse struct {
	*tchttp.BaseResponse
	Response *CreateLicenseResponseParams `json:"Response"`
}

func (r *CreateLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFairPlayPemRequestParams struct {
	// 委托者Id,适用于托管自身证书的客户。普通客户无需填该字段。
	BailorId *uint64 `json:"BailorId,omitempty" name:"BailorId"`

	// 要删除的pem id。
	// 当未传入该值时，将删除所有的私钥。
	FairPlayPemId *uint64 `json:"FairPlayPemId,omitempty" name:"FairPlayPemId"`
}

type DeleteFairPlayPemRequest struct {
	*tchttp.BaseRequest
	
	// 委托者Id,适用于托管自身证书的客户。普通客户无需填该字段。
	BailorId *uint64 `json:"BailorId,omitempty" name:"BailorId"`

	// 要删除的pem id。
	// 当未传入该值时，将删除所有的私钥。
	FairPlayPemId *uint64 `json:"FairPlayPemId,omitempty" name:"FairPlayPemId"`
}

func (r *DeleteFairPlayPemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFairPlayPemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BailorId")
	delete(f, "FairPlayPemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFairPlayPemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFairPlayPemResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteFairPlayPemResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFairPlayPemResponseParams `json:"Response"`
}

func (r *DeleteFairPlayPemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFairPlayPemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllKeysRequestParams struct {
	// 使用的DRM方案类型，接口取值WIDEVINE、FAIRPLAY、NORMALAES。
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// Base64编码的Rsa公钥，用来加密出参中的SessionKey。
	// 如果该参数为空，则出参中SessionKey为明文。
	RsaPublicKey *string `json:"RsaPublicKey,omitempty" name:"RsaPublicKey"`

	// 一个加密内容的唯一标识。
	ContentId *string `json:"ContentId,omitempty" name:"ContentId"`

	// 内容类型。接口取值VodVideo,LiveVideo。
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
}

type DescribeAllKeysRequest struct {
	*tchttp.BaseRequest
	
	// 使用的DRM方案类型，接口取值WIDEVINE、FAIRPLAY、NORMALAES。
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// Base64编码的Rsa公钥，用来加密出参中的SessionKey。
	// 如果该参数为空，则出参中SessionKey为明文。
	RsaPublicKey *string `json:"RsaPublicKey,omitempty" name:"RsaPublicKey"`

	// 一个加密内容的唯一标识。
	ContentId *string `json:"ContentId,omitempty" name:"ContentId"`

	// 内容类型。接口取值VodVideo,LiveVideo。
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
}

func (r *DescribeAllKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DrmType")
	delete(f, "RsaPublicKey")
	delete(f, "ContentId")
	delete(f, "ContentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllKeysResponseParams struct {
	// 加密密钥列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*Key `json:"Keys,omitempty" name:"Keys"`

	// 用来加密密钥。
	// 如果入参中带有RsaPublicKey，则SessionKey为使用Rsa公钥加密后的二进制数据，Base64编码字符串。
	// 如果入参中没有RsaPublicKey，则SessionKey为原始数据的字符串形式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionKey *string `json:"SessionKey,omitempty" name:"SessionKey"`

	// 内容ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentId *string `json:"ContentId,omitempty" name:"ContentId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAllKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllKeysResponseParams `json:"Response"`
}

func (r *DescribeAllKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFairPlayPemRequestParams struct {
	// 委托者Id,适用于托管自身证书的客户。普通客户无需填该字段。
	BailorId *uint64 `json:"BailorId,omitempty" name:"BailorId"`

	// 需要查询的pem id。
	// 当该值未填入时，将返回所有的私钥信息。
	FairPlayPemId *uint64 `json:"FairPlayPemId,omitempty" name:"FairPlayPemId"`
}

type DescribeFairPlayPemRequest struct {
	*tchttp.BaseRequest
	
	// 委托者Id,适用于托管自身证书的客户。普通客户无需填该字段。
	BailorId *uint64 `json:"BailorId,omitempty" name:"BailorId"`

	// 需要查询的pem id。
	// 当该值未填入时，将返回所有的私钥信息。
	FairPlayPemId *uint64 `json:"FairPlayPemId,omitempty" name:"FairPlayPemId"`
}

func (r *DescribeFairPlayPemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFairPlayPemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BailorId")
	delete(f, "FairPlayPemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFairPlayPemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFairPlayPemResponseParams struct {
	// 该账户下，所有设置的FairPlay私钥摘要信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FairPlayPems []*FairPlayPemDigestInfo `json:"FairPlayPems,omitempty" name:"FairPlayPems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFairPlayPemResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFairPlayPemResponseParams `json:"Response"`
}

func (r *DescribeFairPlayPemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFairPlayPemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeysRequestParams struct {
	// 使用的DRM方案类型，接口取值WIDEVINE、FAIRPLAY、NORMALAES。
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// 加密的track列表，接口取值VIDEO、AUDIO。
	Tracks []*string `json:"Tracks,omitempty" name:"Tracks"`

	// 内容类型。接口取值VodVideo,LiveVideo
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// Base64编码的Rsa公钥，用来加密出参中的SessionKey。
	// 如果该参数为空，则出参中SessionKey为明文。
	RsaPublicKey *string `json:"RsaPublicKey,omitempty" name:"RsaPublicKey"`

	// 一个加密内容的唯一标识。
	// 如果该参数为空，则后台自动生成
	ContentId *string `json:"ContentId,omitempty" name:"ContentId"`
}

type DescribeKeysRequest struct {
	*tchttp.BaseRequest
	
	// 使用的DRM方案类型，接口取值WIDEVINE、FAIRPLAY、NORMALAES。
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// 加密的track列表，接口取值VIDEO、AUDIO。
	Tracks []*string `json:"Tracks,omitempty" name:"Tracks"`

	// 内容类型。接口取值VodVideo,LiveVideo
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// Base64编码的Rsa公钥，用来加密出参中的SessionKey。
	// 如果该参数为空，则出参中SessionKey为明文。
	RsaPublicKey *string `json:"RsaPublicKey,omitempty" name:"RsaPublicKey"`

	// 一个加密内容的唯一标识。
	// 如果该参数为空，则后台自动生成
	ContentId *string `json:"ContentId,omitempty" name:"ContentId"`
}

func (r *DescribeKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DrmType")
	delete(f, "Tracks")
	delete(f, "ContentType")
	delete(f, "RsaPublicKey")
	delete(f, "ContentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeysResponseParams struct {
	// 加密密钥列表
	Keys []*Key `json:"Keys,omitempty" name:"Keys"`

	// 用来加密密钥。
	// 如果入参中带有RsaPublicKey，则SessionKey为使用Rsa公钥加密后的二进制数据，Base64编码字符串。
	// 如果入参中没有RsaPublicKey，则SessionKey为原始数据的字符串形式。
	SessionKey *string `json:"SessionKey,omitempty" name:"SessionKey"`

	// 内容ID
	ContentId *string `json:"ContentId,omitempty" name:"ContentId"`

	// Widevine方案的Pssh数据，Base64编码。
	// Fairplay方案无该值。
	Pssh *string `json:"Pssh,omitempty" name:"Pssh"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeysResponseParams `json:"Response"`
}

func (r *DescribeKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrmOutputObject struct {
	// 输出的桶名称。
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 输出的对象名称。
	ObjectName *string `json:"ObjectName,omitempty" name:"ObjectName"`

	// 输出对象参数。
	Para *DrmOutputPara `json:"Para,omitempty" name:"Para"`
}

type DrmOutputPara struct {
	// 内容类型。例:video，audio，mpd，m3u8
	Type *string `json:"Type,omitempty" name:"Type"`

	// 语言,例: en, zh-cn
	Language *string `json:"Language,omitempty" name:"Language"`
}

type DrmSourceObject struct {
	// 输入的桶名称。
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 输入对象名称。
	ObjectName *string `json:"ObjectName,omitempty" name:"ObjectName"`
}

type FairPlayPemDigestInfo struct {
	// fairplay 私钥pem id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FairPlayPemId *uint64 `json:"FairPlayPemId,omitempty" name:"FairPlayPemId"`

	// 私钥的优先级。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`

	// 私钥的md5 信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5Pem *string `json:"Md5Pem,omitempty" name:"Md5Pem"`

	// ASK的md5信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5Ask *string `json:"Md5Ask,omitempty" name:"Md5Ask"`

	// 私钥解密密钥的md5值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5PemDecryptKey *string `json:"Md5PemDecryptKey,omitempty" name:"Md5PemDecryptKey"`
}

type Key struct {
	// 加密track类型。Widevine支持SD、HD、UHD1、UHD2、AUDIO。Fairplay只支持HD。
	Track *string `json:"Track,omitempty" name:"Track"`

	// 密钥ID。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 原始Key使用AES-128 ECB模式和SessionKey加密的后的二进制数据，Base64编码的字符串。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 原始IV使用AES-128 ECB模式和SessionKey加密的后的二进制数据，Base64编码的字符串。
	Iv *string `json:"Iv,omitempty" name:"Iv"`

	// 该key生成时的时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTimestamp *uint64 `json:"InsertTimestamp,omitempty" name:"InsertTimestamp"`
}

type KeyParam struct {
	// 加密track类型。取值范围：
	// SD、HD、UHD1、UHD2、AUDIO
	Track *string `json:"Track,omitempty" name:"Track"`

	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对解密密钥进行加密，并对加密结果进行base64编码。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 密钥ID。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对解密密钥进行加密，并对加密结果进行base64编码。
	Iv *string `json:"Iv,omitempty" name:"Iv"`
}

// Predefined struct for user
type ModifyFairPlayPemRequestParams struct {
	// 加密后的fairplay方案申请时使用的私钥。
	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对私钥文件中的字段进行加密，并对加密结果进行base64编码。
	Pem *string `json:"Pem,omitempty" name:"Pem"`

	// 加密后的fairplay方案申请返回的ask数据。
	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对Ask字符串进行加密，并对加密结果进行base64编码。
	Ask *string `json:"Ask,omitempty" name:"Ask"`

	// 要修改的私钥id
	FairPlayPemId *uint64 `json:"FairPlayPemId,omitempty" name:"FairPlayPemId"`

	// 私钥的解密密钥。
	// openssl在生成rsa时，可能会需要设置加密密钥，请记住设置的密钥。
	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对解密密钥进行加密，并对加密结果进行base64编码。
	PemDecryptKey *string `json:"PemDecryptKey,omitempty" name:"PemDecryptKey"`

	// 委托者Id,适用于托管自身证书的客户。普通客户无需填该字段。
	BailorId *uint64 `json:"BailorId,omitempty" name:"BailorId"`

	// 私钥的优先级，优先级数值越高，优先级越高。
	// 该值可以不传，后台将自动分配一个优先级。
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
}

type ModifyFairPlayPemRequest struct {
	*tchttp.BaseRequest
	
	// 加密后的fairplay方案申请时使用的私钥。
	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对私钥文件中的字段进行加密，并对加密结果进行base64编码。
	Pem *string `json:"Pem,omitempty" name:"Pem"`

	// 加密后的fairplay方案申请返回的ask数据。
	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对Ask字符串进行加密，并对加密结果进行base64编码。
	Ask *string `json:"Ask,omitempty" name:"Ask"`

	// 要修改的私钥id
	FairPlayPemId *uint64 `json:"FairPlayPemId,omitempty" name:"FairPlayPemId"`

	// 私钥的解密密钥。
	// openssl在生成rsa时，可能会需要设置加密密钥，请记住设置的密钥。
	// 请使用腾讯云DRM 提供的公钥，使用rsa加密算法，PKCS1填充方式对解密密钥进行加密，并对加密结果进行base64编码。
	PemDecryptKey *string `json:"PemDecryptKey,omitempty" name:"PemDecryptKey"`

	// 委托者Id,适用于托管自身证书的客户。普通客户无需填该字段。
	BailorId *uint64 `json:"BailorId,omitempty" name:"BailorId"`

	// 私钥的优先级，优先级数值越高，优先级越高。
	// 该值可以不传，后台将自动分配一个优先级。
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`
}

func (r *ModifyFairPlayPemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFairPlayPemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Pem")
	delete(f, "Ask")
	delete(f, "FairPlayPemId")
	delete(f, "PemDecryptKey")
	delete(f, "BailorId")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFairPlayPemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFairPlayPemResponseParams struct {
	// 设置私钥后，后台返回的pem id，用来唯一标识一个私钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FairPlayPemId *uint64 `json:"FairPlayPemId,omitempty" name:"FairPlayPemId"`

	// 私钥的优先级，优先级数值越高，优先级越高。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyFairPlayPemResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFairPlayPemResponseParams `json:"Response"`
}

func (r *ModifyFairPlayPemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFairPlayPemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PlaybackPolicy struct {
	// 播放许可证的有效期
	LicenseDurationSeconds *uint64 `json:"LicenseDurationSeconds,omitempty" name:"LicenseDurationSeconds"`

	// 开始播放后，允许最长播放时间
	PlaybackDurationSeconds *uint64 `json:"PlaybackDurationSeconds,omitempty" name:"PlaybackDurationSeconds"`
}

// Predefined struct for user
type StartEncryptionRequestParams struct {
	// cos的end point。
	CosEndPoint *string `json:"CosEndPoint,omitempty" name:"CosEndPoint"`

	// cos api密钥id。
	CosSecretId *string `json:"CosSecretId,omitempty" name:"CosSecretId"`

	// cos api密钥。
	CosSecretKey *string `json:"CosSecretKey,omitempty" name:"CosSecretKey"`

	// 使用的DRM方案类型，接口取值WIDEVINE,FAIRPLAY
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// 存储在COS上的原始内容信息
	SourceObject *DrmSourceObject `json:"SourceObject,omitempty" name:"SourceObject"`

	// 加密后的内容存储到COS的对象
	OutputObjects []*DrmOutputObject `json:"OutputObjects,omitempty" name:"OutputObjects"`
}

type StartEncryptionRequest struct {
	*tchttp.BaseRequest
	
	// cos的end point。
	CosEndPoint *string `json:"CosEndPoint,omitempty" name:"CosEndPoint"`

	// cos api密钥id。
	CosSecretId *string `json:"CosSecretId,omitempty" name:"CosSecretId"`

	// cos api密钥。
	CosSecretKey *string `json:"CosSecretKey,omitempty" name:"CosSecretKey"`

	// 使用的DRM方案类型，接口取值WIDEVINE,FAIRPLAY
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// 存储在COS上的原始内容信息
	SourceObject *DrmSourceObject `json:"SourceObject,omitempty" name:"SourceObject"`

	// 加密后的内容存储到COS的对象
	OutputObjects []*DrmOutputObject `json:"OutputObjects,omitempty" name:"OutputObjects"`
}

func (r *StartEncryptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartEncryptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CosEndPoint")
	delete(f, "CosSecretId")
	delete(f, "CosSecretKey")
	delete(f, "DrmType")
	delete(f, "SourceObject")
	delete(f, "OutputObjects")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartEncryptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartEncryptionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartEncryptionResponse struct {
	*tchttp.BaseResponse
	Response *StartEncryptionResponseParams `json:"Response"`
}

func (r *StartEncryptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartEncryptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}