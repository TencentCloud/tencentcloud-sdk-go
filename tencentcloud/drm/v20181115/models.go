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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

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

func (r *AddFairPlayPemRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddFairPlayPemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设置私钥后，后台返回的pem id，用来唯一标识一个私钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
		FairPlayPemId *uint64 `json:"FairPlayPemId,omitempty" name:"FairPlayPemId"`

		// 私钥的优先级，优先级数值越高，优先级越高。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Priority *uint64 `json:"Priority,omitempty" name:"Priority"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddFairPlayPemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddFairPlayPemResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	Tracks []*string `json:"Tracks,omitempty" name:"Tracks" list`

	// 播放策略参数。
	PlaybackPolicy *PlaybackPolicy `json:"PlaybackPolicy,omitempty" name:"PlaybackPolicy"`
}

func (r *CreateLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLicenseRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLicenseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Base64 编码的许可证二进制数据。
		License *string `json:"License,omitempty" name:"License"`

		// 加密内容的内容ID
		ContentId *string `json:"ContentId,omitempty" name:"ContentId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLicenseResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteFairPlayPemRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFairPlayPemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFairPlayPemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFairPlayPemResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeFairPlayPemRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFairPlayPemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该账户下，所有设置的FairPlay私钥摘要信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		FairPlayPems []*FairPlayPemDigestInfo `json:"FairPlayPems,omitempty" name:"FairPlayPems" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFairPlayPemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFairPlayPemResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeysRequest struct {
	*tchttp.BaseRequest

	// 使用的DRM方案类型，接口取值WIDEVINE、FAIRPLAY、NORMALAES。
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// 加密的track列表，接口取值VIDEO、AUDIO。
	Tracks []*string `json:"Tracks,omitempty" name:"Tracks" list`

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

func (r *DescribeKeysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 加密密钥列表
		Keys []*Key `json:"Keys,omitempty" name:"Keys" list`

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
	} `json:"Response"`
}

func (r *DescribeKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

	// 加密track类型。
	Track *string `json:"Track,omitempty" name:"Track"`

	// 密钥ID。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 原始Key使用AES-128 ECB模式和SessionKey加密的后的二进制数据，Base64编码的字符串。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 原始IV使用AES-128 ECB模式和SessionKey加密的后的二进制数据，Base64编码的字符串。
	Iv *string `json:"Iv,omitempty" name:"Iv"`
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

func (r *ModifyFairPlayPemRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyFairPlayPemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设置私钥后，后台返回的pem id，用来唯一标识一个私钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
		FairPlayPemId *uint64 `json:"FairPlayPemId,omitempty" name:"FairPlayPemId"`

		// 私钥的优先级，优先级数值越高，优先级越高。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Priority *uint64 `json:"Priority,omitempty" name:"Priority"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFairPlayPemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFairPlayPemResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PlaybackPolicy struct {

	// 播放许可证的有效期
	LicenseDurationSeconds *uint64 `json:"LicenseDurationSeconds,omitempty" name:"LicenseDurationSeconds"`

	// 开始播放后，允许最长播放时间
	PlaybackDurationSeconds *uint64 `json:"PlaybackDurationSeconds,omitempty" name:"PlaybackDurationSeconds"`
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
	OutputObjects []*DrmOutputObject `json:"OutputObjects,omitempty" name:"OutputObjects" list`
}

func (r *StartEncryptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartEncryptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartEncryptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartEncryptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartEncryptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
