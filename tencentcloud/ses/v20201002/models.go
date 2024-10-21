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

package v20201002

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Attachment struct {
	// 附件名称，最大支持255个字符长度，不支持部分附件类型，详情请参考[附件类型](https://cloud.tencent.com/document/product/1288/51951)。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Base64之后的附件内容，您可以发送的附件大小上限为4M。注意：腾讯云接口请求最大支持 8M 的请求包，附件内容经过 Base64 预期扩大1.5倍。应该控制所有附件的总大小最大在 4M 以内，整体请求超出 8M 接口会返回错误。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type BatchSendEmailRequestParams struct {
	// 发信邮件地址。请填写发件人邮箱地址，例如：noreply@mail.qcloud.com。如需填写发件人说明，请按照
	// 发信人 &lt;邮件地址&gt; 的方式填写，例如：
	// 腾讯云团队 &lt;noreply@mail.qcloud.com&gt;
	FromEmailAddress *string `json:"FromEmailAddress,omitnil,omitempty" name:"FromEmailAddress"`

	// 收件人列表ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 邮件主题
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// 任务类型 1: 立即发送 2: 定时发送 3: 周期（频率）发送
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 邮件的“回复”电子邮件地址。可以填写您能收到邮件的邮箱地址，可以是个人邮箱。如果不填，收件人的回复邮件将会发送失败。
	ReplyToAddresses *string `json:"ReplyToAddresses,omitnil,omitempty" name:"ReplyToAddresses"`

	// 使用模板发送时，填写的模板相关参数
	// <dx-alert infotype="notice" title="注意"> 如您未申请过特殊配置，则该字段为必填 </dx-alert>
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// 已废弃
	// <dx-alert infotype="notice" title="说明"> 仅部分历史上申请了特殊配置的客户需要使用。如您未申请过特殊配置，则不存在该字段。</dx-alert>
	Simple *Simple `json:"Simple,omitnil,omitempty" name:"Simple"`

	// 需要发送附件时，填写附件相关参数（暂未支持）
	Attachments []*Attachment `json:"Attachments,omitnil,omitempty" name:"Attachments"`

	// 周期发送任务的必要参数
	CycleParam *CycleEmailParam `json:"CycleParam,omitnil,omitempty" name:"CycleParam"`

	// 定时发送任务的必要参数
	TimedParam *TimedEmailParam `json:"TimedParam,omitnil,omitempty" name:"TimedParam"`

	// 退订链接选项 0: 不加入退订链接 1: 简体中文 2: 英文 3: 繁体中文 4: 西班牙语 5: 法语 6: 德语 7: 日语 8: 韩语 9: 阿拉伯语 10: 泰语
	Unsubscribe *string `json:"Unsubscribe,omitnil,omitempty" name:"Unsubscribe"`

	// 是否添加广告标识 0:不添加 1:添加到subject前面，2:添加到subject后面
	ADLocation *uint64 `json:"ADLocation,omitnil,omitempty" name:"ADLocation"`
}

type BatchSendEmailRequest struct {
	*tchttp.BaseRequest
	
	// 发信邮件地址。请填写发件人邮箱地址，例如：noreply@mail.qcloud.com。如需填写发件人说明，请按照
	// 发信人 &lt;邮件地址&gt; 的方式填写，例如：
	// 腾讯云团队 &lt;noreply@mail.qcloud.com&gt;
	FromEmailAddress *string `json:"FromEmailAddress,omitnil,omitempty" name:"FromEmailAddress"`

	// 收件人列表ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 邮件主题
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// 任务类型 1: 立即发送 2: 定时发送 3: 周期（频率）发送
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 邮件的“回复”电子邮件地址。可以填写您能收到邮件的邮箱地址，可以是个人邮箱。如果不填，收件人的回复邮件将会发送失败。
	ReplyToAddresses *string `json:"ReplyToAddresses,omitnil,omitempty" name:"ReplyToAddresses"`

	// 使用模板发送时，填写的模板相关参数
	// <dx-alert infotype="notice" title="注意"> 如您未申请过特殊配置，则该字段为必填 </dx-alert>
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// 已废弃
	// <dx-alert infotype="notice" title="说明"> 仅部分历史上申请了特殊配置的客户需要使用。如您未申请过特殊配置，则不存在该字段。</dx-alert>
	Simple *Simple `json:"Simple,omitnil,omitempty" name:"Simple"`

	// 需要发送附件时，填写附件相关参数（暂未支持）
	Attachments []*Attachment `json:"Attachments,omitnil,omitempty" name:"Attachments"`

	// 周期发送任务的必要参数
	CycleParam *CycleEmailParam `json:"CycleParam,omitnil,omitempty" name:"CycleParam"`

	// 定时发送任务的必要参数
	TimedParam *TimedEmailParam `json:"TimedParam,omitnil,omitempty" name:"TimedParam"`

	// 退订链接选项 0: 不加入退订链接 1: 简体中文 2: 英文 3: 繁体中文 4: 西班牙语 5: 法语 6: 德语 7: 日语 8: 韩语 9: 阿拉伯语 10: 泰语
	Unsubscribe *string `json:"Unsubscribe,omitnil,omitempty" name:"Unsubscribe"`

	// 是否添加广告标识 0:不添加 1:添加到subject前面，2:添加到subject后面
	ADLocation *uint64 `json:"ADLocation,omitnil,omitempty" name:"ADLocation"`
}

func (r *BatchSendEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchSendEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromEmailAddress")
	delete(f, "ReceiverId")
	delete(f, "Subject")
	delete(f, "TaskType")
	delete(f, "ReplyToAddresses")
	delete(f, "Template")
	delete(f, "Simple")
	delete(f, "Attachments")
	delete(f, "CycleParam")
	delete(f, "TimedParam")
	delete(f, "Unsubscribe")
	delete(f, "ADLocation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchSendEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchSendEmailResponseParams struct {
	// 发送任务ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchSendEmailResponse struct {
	*tchttp.BaseResponse
	Response *BatchSendEmailResponseParams `json:"Response"`
}

func (r *BatchSendEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchSendEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlackAddressDetail struct {
	// 黑名单地址id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 邮箱地址
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 过期时间
	ExpireDate *string `json:"ExpireDate,omitnil,omitempty" name:"ExpireDate"`

	// 黑名单状态，0:已过期，1:生效中
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type BlackEmailAddress struct {
	// 邮箱被拉黑时间
	BounceTime *string `json:"BounceTime,omitnil,omitempty" name:"BounceTime"`

	// 被拉黑的邮箱地址
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// 被拉黑的理由
	// 注意：此字段可能返回 null，表示取不到有效值。
	IspDesc *string `json:"IspDesc,omitnil,omitempty" name:"IspDesc"`
}

// Predefined struct for user
type CreateCustomBlacklistRequestParams struct {
	// 添加到黑名单的邮件地址
	Emails []*string `json:"Emails,omitnil,omitempty" name:"Emails"`

	// 过期日期
	ExpireDate *string `json:"ExpireDate,omitnil,omitempty" name:"ExpireDate"`
}

type CreateCustomBlacklistRequest struct {
	*tchttp.BaseRequest
	
	// 添加到黑名单的邮件地址
	Emails []*string `json:"Emails,omitnil,omitempty" name:"Emails"`

	// 过期日期
	ExpireDate *string `json:"ExpireDate,omitnil,omitempty" name:"ExpireDate"`
}

func (r *CreateCustomBlacklistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomBlacklistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Emails")
	delete(f, "ExpireDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomBlacklistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomBlacklistResponseParams struct {
	// 收件人总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实际上传数量
	ValidCount *uint64 `json:"ValidCount,omitnil,omitempty" name:"ValidCount"`

	// 数据过长数量
	TooLongCount *uint64 `json:"TooLongCount,omitnil,omitempty" name:"TooLongCount"`

	// 重复数量
	RepeatCount *uint64 `json:"RepeatCount,omitnil,omitempty" name:"RepeatCount"`

	// 格式不正确数量
	InvalidCount *uint64 `json:"InvalidCount,omitnil,omitempty" name:"InvalidCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomBlacklistResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomBlacklistResponseParams `json:"Response"`
}

func (r *CreateCustomBlacklistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomBlacklistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmailAddressRequestParams struct {
	// 您的发信地址（发信地址总数上限为10个）
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// 发件人别名
	EmailSenderName *string `json:"EmailSenderName,omitnil,omitempty" name:"EmailSenderName"`
}

type CreateEmailAddressRequest struct {
	*tchttp.BaseRequest
	
	// 您的发信地址（发信地址总数上限为10个）
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// 发件人别名
	EmailSenderName *string `json:"EmailSenderName,omitnil,omitempty" name:"EmailSenderName"`
}

func (r *CreateEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailAddress")
	delete(f, "EmailSenderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmailAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmailAddressResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmailAddressResponseParams `json:"Response"`
}

func (r *CreateEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmailIdentityRequestParams struct {
	// 您的发信域名，建议使用三级以上域名。例如：mail.qcloud.com。
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

type CreateEmailIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 您的发信域名，建议使用三级以上域名。例如：mail.qcloud.com。
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

func (r *CreateEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailIdentity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmailIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmailIdentityResponseParams struct {
	// 验证类型。固定值：DOMAIN
	IdentityType *string `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// 是否已通过验证
	VerifiedForSendingStatus *bool `json:"VerifiedForSendingStatus,omitnil,omitempty" name:"VerifiedForSendingStatus"`

	// 需要配置的DNS信息
	Attributes []*DNSAttributes `json:"Attributes,omitnil,omitempty" name:"Attributes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmailIdentityResponseParams `json:"Response"`
}

func (r *CreateEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmailTemplateRequestParams struct {
	// 模板名称
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 模板内容
	TemplateContent *TemplateContent `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`
}

type CreateEmailTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 模板内容
	TemplateContent *TemplateContent `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`
}

func (r *CreateEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "TemplateContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmailTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmailTemplateResponseParams struct {
	// 模板id
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmailTemplateResponseParams `json:"Response"`
}

func (r *CreateEmailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReceiverDetailRequestParams struct {
	// 收件人列表ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 邮箱
	Emails []*string `json:"Emails,omitnil,omitempty" name:"Emails"`
}

type CreateReceiverDetailRequest struct {
	*tchttp.BaseRequest
	
	// 收件人列表ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 邮箱
	Emails []*string `json:"Emails,omitnil,omitempty" name:"Emails"`
}

func (r *CreateReceiverDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReceiverDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReceiverId")
	delete(f, "Emails")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReceiverDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReceiverDetailResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReceiverDetailResponse struct {
	*tchttp.BaseResponse
	Response *CreateReceiverDetailResponseParams `json:"Response"`
}

func (r *CreateReceiverDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReceiverDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReceiverDetailWithDataRequestParams struct {
	// 收件人列表ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 收信人邮箱以及模板参数，数组形式。收件人个数限制20000个以内。
	Datas []*ReceiverInputData `json:"Datas,omitnil,omitempty" name:"Datas"`
}

type CreateReceiverDetailWithDataRequest struct {
	*tchttp.BaseRequest
	
	// 收件人列表ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 收信人邮箱以及模板参数，数组形式。收件人个数限制20000个以内。
	Datas []*ReceiverInputData `json:"Datas,omitnil,omitempty" name:"Datas"`
}

func (r *CreateReceiverDetailWithDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReceiverDetailWithDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReceiverId")
	delete(f, "Datas")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReceiverDetailWithDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReceiverDetailWithDataResponseParams struct {
	// 收件人总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实际上传数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidCount *uint64 `json:"ValidCount,omitnil,omitempty" name:"ValidCount"`

	// 数据过长数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TooLongCount *uint64 `json:"TooLongCount,omitnil,omitempty" name:"TooLongCount"`

	// 邮件地址为空数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmptyEmailCount *uint64 `json:"EmptyEmailCount,omitnil,omitempty" name:"EmptyEmailCount"`

	// 重复数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepeatCount *uint64 `json:"RepeatCount,omitnil,omitempty" name:"RepeatCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReceiverDetailWithDataResponse struct {
	*tchttp.BaseResponse
	Response *CreateReceiverDetailWithDataResponseParams `json:"Response"`
}

func (r *CreateReceiverDetailWithDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReceiverDetailWithDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReceiverRequestParams struct {
	// 收件人列表名称
	ReceiversName *string `json:"ReceiversName,omitnil,omitempty" name:"ReceiversName"`

	// 收件人列表描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type CreateReceiverRequest struct {
	*tchttp.BaseRequest
	
	// 收件人列表名称
	ReceiversName *string `json:"ReceiversName,omitnil,omitempty" name:"ReceiversName"`

	// 收件人列表描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

func (r *CreateReceiverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReceiverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReceiversName")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReceiverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReceiverResponseParams struct {
	// 收件人列表id，后续根据收件人列表id上传收件人地址
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReceiverResponse struct {
	*tchttp.BaseResponse
	Response *CreateReceiverResponseParams `json:"Response"`
}

func (r *CreateReceiverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CycleEmailParam struct {
	// 任务开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 任务周期 小时维度
	IntervalTime *uint64 `json:"IntervalTime,omitnil,omitempty" name:"IntervalTime"`

	// 是否终止周期，用于任务更新 0否1是
	TermCycle *uint64 `json:"TermCycle,omitnil,omitempty" name:"TermCycle"`
}

type DNSAttributes struct {
	// 记录类型 CNAME | A | TXT | MX
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 域名
	SendDomain *string `json:"SendDomain,omitnil,omitempty" name:"SendDomain"`

	// 需要配置的值
	ExpectedValue *string `json:"ExpectedValue,omitnil,omitempty" name:"ExpectedValue"`

	// 腾讯云目前检测到的值
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 检测是否通过，创建时默认为false
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteBlackListRequestParams struct {
	// 需要清除的黑名单邮箱列表，数组长度至少为1
	EmailAddressList []*string `json:"EmailAddressList,omitnil,omitempty" name:"EmailAddressList"`
}

type DeleteBlackListRequest struct {
	*tchttp.BaseRequest
	
	// 需要清除的黑名单邮箱列表，数组长度至少为1
	EmailAddressList []*string `json:"EmailAddressList,omitnil,omitempty" name:"EmailAddressList"`
}

func (r *DeleteBlackListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlackListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailAddressList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBlackListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBlackListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBlackListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBlackListResponseParams `json:"Response"`
}

func (r *DeleteBlackListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomBlackListRequestParams struct {
	// 需要删除的邮箱地址
	Emails []*string `json:"Emails,omitnil,omitempty" name:"Emails"`
}

type DeleteCustomBlackListRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的邮箱地址
	Emails []*string `json:"Emails,omitnil,omitempty" name:"Emails"`
}

func (r *DeleteCustomBlackListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomBlackListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Emails")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomBlackListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomBlackListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCustomBlackListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomBlackListResponseParams `json:"Response"`
}

func (r *DeleteCustomBlackListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomBlackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEmailAddressRequestParams struct {
	// 发信地址
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`
}

type DeleteEmailAddressRequest struct {
	*tchttp.BaseRequest
	
	// 发信地址
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`
}

func (r *DeleteEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEmailAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEmailAddressResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEmailAddressResponseParams `json:"Response"`
}

func (r *DeleteEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEmailIdentityRequestParams struct {
	// 发信域名
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

type DeleteEmailIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 发信域名
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

func (r *DeleteEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailIdentity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEmailIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEmailIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEmailIdentityResponseParams `json:"Response"`
}

func (r *DeleteEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEmailTemplateRequestParams struct {
	// 模板ID
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`
}

type DeleteEmailTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`
}

func (r *DeleteEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEmailTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEmailTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEmailTemplateResponseParams `json:"Response"`
}

func (r *DeleteEmailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReceiverRequestParams struct {
	// 收件人列表id，创建收件人列表时会返回
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`
}

type DeleteReceiverRequest struct {
	*tchttp.BaseRequest
	
	// 收件人列表id，创建收件人列表时会返回
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`
}

func (r *DeleteReceiverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReceiverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReceiverId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReceiverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReceiverResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReceiverResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReceiverResponseParams `json:"Response"`
}

func (r *DeleteReceiverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmailIdentity struct {
	// 发信域名
	IdentityName *string `json:"IdentityName,omitnil,omitempty" name:"IdentityName"`

	// 验证类型，固定为DOMAIN
	IdentityType *string `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// 是否已通过验证
	SendingEnabled *bool `json:"SendingEnabled,omitnil,omitempty" name:"SendingEnabled"`

	// 当前信誉等级
	CurrentReputationLevel *uint64 `json:"CurrentReputationLevel,omitnil,omitempty" name:"CurrentReputationLevel"`

	// 当日最高发信量
	DailyQuota *uint64 `json:"DailyQuota,omitnil,omitempty" name:"DailyQuota"`
}

type EmailSender struct {
	// 发信地址
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// 发信人别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmailSenderName *string `json:"EmailSenderName,omitnil,omitempty" name:"EmailSenderName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTimestamp *uint64 `json:"CreatedTimestamp,omitnil,omitempty" name:"CreatedTimestamp"`
}

// Predefined struct for user
type GetEmailIdentityRequestParams struct {
	// 发信域名
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

type GetEmailIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 发信域名
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

func (r *GetEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmailIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailIdentity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetEmailIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEmailIdentityResponseParams struct {
	// 验证类型。固定值：DOMAIN
	IdentityType *string `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// 是否已通过验证
	VerifiedForSendingStatus *bool `json:"VerifiedForSendingStatus,omitnil,omitempty" name:"VerifiedForSendingStatus"`

	// DNS配置详情
	Attributes []*DNSAttributes `json:"Attributes,omitnil,omitempty" name:"Attributes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *GetEmailIdentityResponseParams `json:"Response"`
}

func (r *GetEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmailIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEmailTemplateRequestParams struct {
	// 模板ID
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`
}

type GetEmailTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`
}

func (r *GetEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetEmailTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEmailTemplateResponseParams struct {
	// 模板内容数据
	TemplateContent *TemplateContent `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`

	// 模板状态 0-审核通过 1-待审核 2-审核拒绝
	TemplateStatus *uint64 `json:"TemplateStatus,omitnil,omitempty" name:"TemplateStatus"`

	// 模板名称
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *GetEmailTemplateResponseParams `json:"Response"`
}

func (r *GetEmailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSendEmailStatusRequestParams struct {
	// 发送的日期，必填。仅支持查询某个日期，不支持范围查询。
	RequestDate *string `json:"RequestDate,omitnil,omitempty" name:"RequestDate"`

	// 偏移量。默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// SendMail接口返回的MessageId字段。
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 收件人邮箱。
	ToEmailAddress *string `json:"ToEmailAddress,omitnil,omitempty" name:"ToEmailAddress"`
}

type GetSendEmailStatusRequest struct {
	*tchttp.BaseRequest
	
	// 发送的日期，必填。仅支持查询某个日期，不支持范围查询。
	RequestDate *string `json:"RequestDate,omitnil,omitempty" name:"RequestDate"`

	// 偏移量。默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// SendMail接口返回的MessageId字段。
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 收件人邮箱。
	ToEmailAddress *string `json:"ToEmailAddress,omitnil,omitempty" name:"ToEmailAddress"`
}

func (r *GetSendEmailStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSendEmailStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestDate")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MessageId")
	delete(f, "ToEmailAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSendEmailStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSendEmailStatusResponseParams struct {
	// 邮件发送状态列表
	EmailStatusList []*SendEmailStatus `json:"EmailStatusList,omitnil,omitempty" name:"EmailStatusList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSendEmailStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetSendEmailStatusResponseParams `json:"Response"`
}

func (r *GetSendEmailStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSendEmailStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetStatisticsReportRequestParams struct {
	// 开始日期
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 发信域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 收件方邮箱类型，例如gmail.com
	ReceivingMailboxType *string `json:"ReceivingMailboxType,omitnil,omitempty" name:"ReceivingMailboxType"`
}

type GetStatisticsReportRequest struct {
	*tchttp.BaseRequest
	
	// 开始日期
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 发信域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 收件方邮箱类型，例如gmail.com
	ReceivingMailboxType *string `json:"ReceivingMailboxType,omitnil,omitempty" name:"ReceivingMailboxType"`
}

func (r *GetStatisticsReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetStatisticsReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Domain")
	delete(f, "ReceivingMailboxType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetStatisticsReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetStatisticsReportResponseParams struct {
	// 发信统计报告，按天
	DailyVolumes []*Volume `json:"DailyVolumes,omitnil,omitempty" name:"DailyVolumes"`

	// 发信统计报告，总览
	OverallVolume *Volume `json:"OverallVolume,omitnil,omitempty" name:"OverallVolume"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetStatisticsReportResponse struct {
	*tchttp.BaseResponse
	Response *GetStatisticsReportResponseParams `json:"Response"`
}

func (r *GetStatisticsReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetStatisticsReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListBlackEmailAddressRequestParams struct {
	// 开始日期，格式为YYYY-MM-DD
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期，格式为YYYY-MM-DD
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 规范，配合Offset使用
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 规范，配合Limit使用，Limit最大取值为100
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 可以指定邮箱进行查询
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// 已废弃
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

type ListBlackEmailAddressRequest struct {
	*tchttp.BaseRequest
	
	// 开始日期，格式为YYYY-MM-DD
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期，格式为YYYY-MM-DD
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 规范，配合Offset使用
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 规范，配合Limit使用，Limit最大取值为100
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 可以指定邮箱进行查询
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// 已废弃
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

func (r *ListBlackEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListBlackEmailAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EmailAddress")
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListBlackEmailAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListBlackEmailAddressResponseParams struct {
	// 黑名单列表
	BlackList []*BlackEmailAddress `json:"BlackList,omitnil,omitempty" name:"BlackList"`

	// 黑名单总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListBlackEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *ListBlackEmailAddressResponseParams `json:"Response"`
}

func (r *ListBlackEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListBlackEmailAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCustomBlacklistRequestParams struct {
	// 偏移量，整型，从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目，整型,不超过100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选黑名单的状态，0:已过期，1:生效中, 2:全部
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 黑名单中的邮箱地址
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
}

type ListCustomBlacklistRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，整型，从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目，整型,不超过100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选黑名单的状态，0:已过期，1:生效中, 2:全部
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 黑名单中的邮箱地址
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
}

func (r *ListCustomBlacklistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCustomBlacklistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Status")
	delete(f, "Email")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListCustomBlacklistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCustomBlacklistResponseParams struct {
	// 列表总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 黑名单列表详情
	Data []*BlackAddressDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListCustomBlacklistResponse struct {
	*tchttp.BaseResponse
	Response *ListCustomBlacklistResponseParams `json:"Response"`
}

func (r *ListCustomBlacklistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCustomBlacklistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEmailAddressRequestParams struct {

}

type ListEmailAddressRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEmailAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEmailAddressResponseParams struct {
	// 发信地址列表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmailSenders []*EmailSender `json:"EmailSenders,omitnil,omitempty" name:"EmailSenders"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *ListEmailAddressResponseParams `json:"Response"`
}

func (r *ListEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEmailIdentitiesRequestParams struct {

}

type ListEmailIdentitiesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListEmailIdentitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailIdentitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEmailIdentitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEmailIdentitiesResponseParams struct {
	// 发信域名列表
	EmailIdentities []*EmailIdentity `json:"EmailIdentities,omitnil,omitempty" name:"EmailIdentities"`

	// 最大信誉等级
	MaxReputationLevel *uint64 `json:"MaxReputationLevel,omitnil,omitempty" name:"MaxReputationLevel"`

	// 单域名最高日发送量
	MaxDailyQuota *uint64 `json:"MaxDailyQuota,omitnil,omitempty" name:"MaxDailyQuota"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListEmailIdentitiesResponse struct {
	*tchttp.BaseResponse
	Response *ListEmailIdentitiesResponseParams `json:"Response"`
}

func (r *ListEmailIdentitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailIdentitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEmailTemplatesRequestParams struct {
	// 获取模板数据量，用于分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 获取模板偏移值，用于分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListEmailTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 获取模板数据量，用于分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 获取模板偏移值，用于分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListEmailTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEmailTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEmailTemplatesResponseParams struct {
	// 邮件模板列表
	TemplatesMetadata []*TemplatesMetadata `json:"TemplatesMetadata,omitnil,omitempty" name:"TemplatesMetadata"`

	// 模板总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListEmailTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *ListEmailTemplatesResponseParams `json:"Response"`
}

func (r *ListEmailTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReceiverDetailsRequestParams struct {
	// 收件人列表ID,CreateReceiver接口创建收件人列表时会返回该值
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 偏移量，整型，从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目，整型,不超过100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 收件人地址，长度0-50，示例：xxx@te.com，支持模糊查询
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 搜索开始时间
	CreateTimeBegin *string `json:"CreateTimeBegin,omitnil,omitempty" name:"CreateTimeBegin"`

	// 搜索结束时间
	CreateTimeEnd *string `json:"CreateTimeEnd,omitnil,omitempty" name:"CreateTimeEnd"`

	// 1:有效，2:无效
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ListReceiverDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 收件人列表ID,CreateReceiver接口创建收件人列表时会返回该值
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 偏移量，整型，从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目，整型,不超过100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 收件人地址，长度0-50，示例：xxx@te.com，支持模糊查询
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 搜索开始时间
	CreateTimeBegin *string `json:"CreateTimeBegin,omitnil,omitempty" name:"CreateTimeBegin"`

	// 搜索结束时间
	CreateTimeEnd *string `json:"CreateTimeEnd,omitnil,omitempty" name:"CreateTimeEnd"`

	// 1:有效，2:无效
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ListReceiverDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReceiverDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReceiverId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Email")
	delete(f, "CreateTimeBegin")
	delete(f, "CreateTimeEnd")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReceiverDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReceiverDetailsResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据记录
	Data []*ReceiverDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 有效邮件地址数
	ValidCount *uint64 `json:"ValidCount,omitnil,omitempty" name:"ValidCount"`

	// 无效邮件地址数
	InvalidCount *uint64 `json:"InvalidCount,omitnil,omitempty" name:"InvalidCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReceiverDetailsResponse struct {
	*tchttp.BaseResponse
	Response *ListReceiverDetailsResponseParams `json:"Response"`
}

func (r *ListReceiverDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReceiverDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReceiversRequestParams struct {
	// 偏移量，整型，从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目，整型，不超过100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列表状态(1 待上传 2 上传中  3传完成)，若查询所有就不传这个字段
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 列表名称的关键字，模糊查询
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`
}

type ListReceiversRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，整型，从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目，整型，不超过100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列表状态(1 待上传 2 上传中  3传完成)，若查询所有就不传这个字段
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 列表名称的关键字，模糊查询
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`
}

func (r *ListReceiversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReceiversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Status")
	delete(f, "KeyWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReceiversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReceiversResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据记录
	Data []*ReceiverData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReceiversResponse struct {
	*tchttp.BaseResponse
	Response *ListReceiversResponseParams `json:"Response"`
}

func (r *ListReceiversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReceiversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSendTasksRequestParams struct {
	// 偏移量，整型，从0开始，0代表跳过0行
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目，整型,不超过100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 任务状态 1 待开始 5 发送中 6 今日暂停发送  7 发信异常 10 发送完成。查询所有状态，则不传这个字段
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 收件人列表ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 任务类型 1即时 2定时 3周期，查询所有类型则不传这个字段
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type ListSendTasksRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，整型，从0开始，0代表跳过0行
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目，整型,不超过100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 任务状态 1 待开始 5 发送中 6 今日暂停发送  7 发信异常 10 发送完成。查询所有状态，则不传这个字段
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 收件人列表ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 任务类型 1即时 2定时 3周期，查询所有类型则不传这个字段
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

func (r *ListSendTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSendTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Status")
	delete(f, "ReceiverId")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSendTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSendTasksResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据记录
	Data []*SendTaskData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSendTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListSendTasksResponseParams `json:"Response"`
}

func (r *ListSendTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSendTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReceiverData struct {
	// 收件人列表ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 收件人列表名称
	ReceiversName *string `json:"ReceiversName,omitnil,omitempty" name:"ReceiversName"`

	// 收件人地址总数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 收件人列表描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 列表状态(1 待上传 2 上传中 3 上传完成)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiversStatus *uint64 `json:"ReceiversStatus,omitnil,omitempty" name:"ReceiversStatus"`

	// 创建时间,如:2021-09-28 16:40:35
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 无效收件人数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvalidCount *uint64 `json:"InvalidCount,omitnil,omitempty" name:"InvalidCount"`
}

type ReceiverDetail struct {
	// 收件人地址
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模板参数
	TemplateData *string `json:"TemplateData,omitnil,omitempty" name:"TemplateData"`

	// 无效原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 1:有效，2:无效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 收件人地址id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmailId *uint64 `json:"EmailId,omitnil,omitempty" name:"EmailId"`
}

type ReceiverInputData struct {
	// 收件人邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 模板中的变量参数，请使用json.dump将json对象格式化为string类型。该对象是一组键值对，每个Key代表模板中的一个变量，模板中的变量使用{{键}}表示，相应的值在发送时会被替换为{{值}}。
	// 注意：参数值不能是html等复杂类型的数据。TemplateData (整个 JSON 结构) 总长度限制为 800 bytes。
	TemplateData *string `json:"TemplateData,omitnil,omitempty" name:"TemplateData"`
}

// Predefined struct for user
type SendEmailRequestParams struct {
	// 发件人邮箱地址。不使用别名时请直接填写发件人邮箱地址，例如：noreply@mail.qcloud.com如需填写发件人别名时，请按照如下方式（注意别名与邮箱地址之间必须使用一个空格隔开）：别名+一个空格+<邮箱地址>，别名中不能带有冒号(:)。
	FromEmailAddress *string `json:"FromEmailAddress,omitnil,omitempty" name:"FromEmailAddress"`

	// 收信人邮箱地址，最多支持群发50人。注意：邮件内容会显示所有收件人地址，非群发邮件请多次调用API发送。
	Destination []*string `json:"Destination,omitnil,omitempty" name:"Destination"`

	// 邮件主题
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// 邮件的“回复”电子邮件地址。可以填写您能收到邮件的邮箱地址，可以是个人邮箱。如果不填，收件人的回复邮件将会发送失败。
	ReplyToAddresses *string `json:"ReplyToAddresses,omitnil,omitempty" name:"ReplyToAddresses"`

	// 抄送人邮箱地址，最多支持抄送20人。
	Cc []*string `json:"Cc,omitnil,omitempty" name:"Cc"`

	// 密送人邮箱地址，最多支持抄送20人。
	Bcc []*string `json:"Bcc,omitnil,omitempty" name:"Bcc"`

	// 使用模板发送时，填写模板相关参数。
	// <dx-alert infotype="notice" title="注意"> 如您未申请过特殊配置，则该字段为必填 </dx-alert>
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// 已废弃
	// <dx-alert infotype="notice" title="说明"> 仅部分历史上申请了特殊配置的客户需要使用。如您未申请过特殊配置，则不存在该字段。</dx-alert>
	Simple *Simple `json:"Simple,omitnil,omitempty" name:"Simple"`

	// 需要发送附件时，填写附件相关参数。腾讯云接口请求最大支持 8M 的请求包，附件内容经过 Base64 预期扩大1.5倍，应该控制所有附件的总大小最大在 4M 以内，整体请求超出 8M 时接口会返回错误
	Attachments []*Attachment `json:"Attachments,omitnil,omitempty" name:"Attachments"`

	// 退订链接选项 0: 不加入退订链接 1: 简体中文 2: 英文 3: 繁体中文 4: 西班牙语 5: 法语 6: 德语 7: 日语 8: 韩语 9: 阿拉伯语 10: 泰语
	Unsubscribe *string `json:"Unsubscribe,omitnil,omitempty" name:"Unsubscribe"`

	// 邮件触发类型 0:非触发类，默认类型，营销类邮件、非即时类邮件等选择此类型  1:触发类，验证码等即时发送类邮件，若邮件超过一定大小，系统会自动选择非触发类型通道
	TriggerType *uint64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// smtp头中的Message-Id字段
	SmtpMessageId *string `json:"SmtpMessageId,omitnil,omitempty" name:"SmtpMessageId"`

	// smtp头中可以设置的其它字段
	SmtpHeaders *string `json:"SmtpHeaders,omitnil,omitempty" name:"SmtpHeaders"`
}

type SendEmailRequest struct {
	*tchttp.BaseRequest
	
	// 发件人邮箱地址。不使用别名时请直接填写发件人邮箱地址，例如：noreply@mail.qcloud.com如需填写发件人别名时，请按照如下方式（注意别名与邮箱地址之间必须使用一个空格隔开）：别名+一个空格+<邮箱地址>，别名中不能带有冒号(:)。
	FromEmailAddress *string `json:"FromEmailAddress,omitnil,omitempty" name:"FromEmailAddress"`

	// 收信人邮箱地址，最多支持群发50人。注意：邮件内容会显示所有收件人地址，非群发邮件请多次调用API发送。
	Destination []*string `json:"Destination,omitnil,omitempty" name:"Destination"`

	// 邮件主题
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// 邮件的“回复”电子邮件地址。可以填写您能收到邮件的邮箱地址，可以是个人邮箱。如果不填，收件人的回复邮件将会发送失败。
	ReplyToAddresses *string `json:"ReplyToAddresses,omitnil,omitempty" name:"ReplyToAddresses"`

	// 抄送人邮箱地址，最多支持抄送20人。
	Cc []*string `json:"Cc,omitnil,omitempty" name:"Cc"`

	// 密送人邮箱地址，最多支持抄送20人。
	Bcc []*string `json:"Bcc,omitnil,omitempty" name:"Bcc"`

	// 使用模板发送时，填写模板相关参数。
	// <dx-alert infotype="notice" title="注意"> 如您未申请过特殊配置，则该字段为必填 </dx-alert>
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// 已废弃
	// <dx-alert infotype="notice" title="说明"> 仅部分历史上申请了特殊配置的客户需要使用。如您未申请过特殊配置，则不存在该字段。</dx-alert>
	Simple *Simple `json:"Simple,omitnil,omitempty" name:"Simple"`

	// 需要发送附件时，填写附件相关参数。腾讯云接口请求最大支持 8M 的请求包，附件内容经过 Base64 预期扩大1.5倍，应该控制所有附件的总大小最大在 4M 以内，整体请求超出 8M 时接口会返回错误
	Attachments []*Attachment `json:"Attachments,omitnil,omitempty" name:"Attachments"`

	// 退订链接选项 0: 不加入退订链接 1: 简体中文 2: 英文 3: 繁体中文 4: 西班牙语 5: 法语 6: 德语 7: 日语 8: 韩语 9: 阿拉伯语 10: 泰语
	Unsubscribe *string `json:"Unsubscribe,omitnil,omitempty" name:"Unsubscribe"`

	// 邮件触发类型 0:非触发类，默认类型，营销类邮件、非即时类邮件等选择此类型  1:触发类，验证码等即时发送类邮件，若邮件超过一定大小，系统会自动选择非触发类型通道
	TriggerType *uint64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// smtp头中的Message-Id字段
	SmtpMessageId *string `json:"SmtpMessageId,omitnil,omitempty" name:"SmtpMessageId"`

	// smtp头中可以设置的其它字段
	SmtpHeaders *string `json:"SmtpHeaders,omitnil,omitempty" name:"SmtpHeaders"`
}

func (r *SendEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromEmailAddress")
	delete(f, "Destination")
	delete(f, "Subject")
	delete(f, "ReplyToAddresses")
	delete(f, "Cc")
	delete(f, "Bcc")
	delete(f, "Template")
	delete(f, "Simple")
	delete(f, "Attachments")
	delete(f, "Unsubscribe")
	delete(f, "TriggerType")
	delete(f, "SmtpMessageId")
	delete(f, "SmtpHeaders")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendEmailResponseParams struct {
	// 接受消息生成的唯一消息标识符。
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendEmailResponse struct {
	*tchttp.BaseResponse
	Response *SendEmailResponseParams `json:"Response"`
}

func (r *SendEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendEmailStatus struct {
	// SendEmail返回的MessageId
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 收件人邮箱
	ToEmailAddress *string `json:"ToEmailAddress,omitnil,omitempty" name:"ToEmailAddress"`

	// 发件人邮箱
	FromEmailAddress *string `json:"FromEmailAddress,omitnil,omitempty" name:"FromEmailAddress"`

	// 腾讯云处理状态
	// 0: 处理成功
	// 1001: 内部系统异常
	// 1002: 内部系统异常
	// 1003: 内部系统异常
	// 1003: 内部系统异常
	// 1004: 发信超时
	// 1005: 内部系统异常
	// 1006: 触发频率控制，短时间内对同一地址发送过多邮件
	// 1007: 邮件地址在黑名单中
	// 1008: 域名被收件人拒收
	// 1009: 内部系统异常
	// 1010: 超出了每日发送限制
	// 1011: 无发送自定义内容权限，必须使用模板
	// 1013: 域名被收件人取消订阅
	// 2001: 找不到相关记录
	// 3007: 模板ID无效或者不可用
	// 3008: 被收信域名临时封禁
	// 3009: 无权限使用该模板
	// 3010: TemplateData字段格式不正确 
	// 3014: 发件域名没有经过认证，无法发送
	// 3020: 收件方邮箱类型在黑名单
	// 3024: 邮箱地址格式预检查失败
	// 3030: 退信率过高，临时限制发送
	// 3033: 余额不足，账号欠费等
	SendStatus *int64 `json:"SendStatus,omitnil,omitempty" name:"SendStatus"`

	// 收件方处理状态
	// 0: 请求成功被腾讯云接受，进入发送队列
	// 1: 邮件递送成功，DeliverTime表示递送成功的时间
	// 2: 邮件因某种原因被丢弃，DeliverMessage表示丢弃原因
	// 3: 收件方ESP拒信，一般原因为邮箱地址不存在，或其它原因
	// 8: 邮件被ESP因某些原因延迟递送，DeliverMessage表示延迟原因
	DeliverStatus *int64 `json:"DeliverStatus,omitnil,omitempty" name:"DeliverStatus"`

	// 收件方处理状态描述
	DeliverMessage *string `json:"DeliverMessage,omitnil,omitempty" name:"DeliverMessage"`

	// 请求到达腾讯云时间戳
	RequestTime *int64 `json:"RequestTime,omitnil,omitempty" name:"RequestTime"`

	// 腾讯云执行递送时间戳
	DeliverTime *int64 `json:"DeliverTime,omitnil,omitempty" name:"DeliverTime"`

	// 用户是否打开该邮件
	UserOpened *bool `json:"UserOpened,omitnil,omitempty" name:"UserOpened"`

	// 用户是否点击该邮件中的链接
	UserClicked *bool `json:"UserClicked,omitnil,omitempty" name:"UserClicked"`

	// 用户是否取消该发送者的订阅
	UserUnsubscribed *bool `json:"UserUnsubscribed,omitnil,omitempty" name:"UserUnsubscribed"`

	// 用户是否举报该发送者
	UserComplainted *bool `json:"UserComplainted,omitnil,omitempty" name:"UserComplainted"`
}

type SendTaskData struct {
	// 任务id
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 发信地址
	FromEmailAddress *string `json:"FromEmailAddress,omitnil,omitempty" name:"FromEmailAddress"`

	// 收件人列表Id
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// 任务状态 1 待开始 5 发送中 6 今日暂停发送  7 发信异常 10 发送完成
	TaskStatus *uint64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 任务类型 1 即时 2 定时 3 周期
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务请求发信数量
	RequestCount *uint64 `json:"RequestCount,omitnil,omitempty" name:"RequestCount"`

	// 已经发送数量
	SendCount *uint64 `json:"SendCount,omitnil,omitempty" name:"SendCount"`

	// 缓存数量
	CacheCount *uint64 `json:"CacheCount,omitnil,omitempty" name:"CacheCount"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 邮件主题
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// 模板和模板数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// 周期任务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleParam *CycleEmailParam `json:"CycleParam,omitnil,omitempty" name:"CycleParam"`

	// 定时任务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimedParam *TimedEmailParam `json:"TimedParam,omitnil,omitempty" name:"TimedParam"`

	// 任务异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 收件人列表名称
	ReceiversName *string `json:"ReceiversName,omitnil,omitempty" name:"ReceiversName"`
}

type Simple struct {
	// base64之后的Html代码。需要包含所有的代码信息，不要包含外部css，否则会导致显示格式错乱
	Html *string `json:"Html,omitnil,omitempty" name:"Html"`

	// base64之后的纯文本信息，如果没有Html，邮件中会直接显示纯文本；如果有Html，它代表邮件的纯文本样式
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type Template struct {
	// 模板ID。如果没有模板，请先新建一个
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// 模板中的变量参数，请使用json.dump将json对象格式化为string类型。该对象是一组键值对，每个Key代表模板中的一个变量，模板中的变量使用{{键}}表示，相应的值在发送时会被替换为{{值}}。
	// 注意：参数值不能是html等复杂类型的数据。
	// 示例：{"name":"xxx","age":"xx"}
	TemplateData *string `json:"TemplateData,omitnil,omitempty" name:"TemplateData"`
}

type TemplateContent struct {
	// base64之后的Html代码
	Html *string `json:"Html,omitnil,omitempty" name:"Html"`

	// base64之后的文本内容
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type TemplatesMetadata struct {
	// 创建时间
	CreatedTimestamp *uint64 `json:"CreatedTimestamp,omitnil,omitempty" name:"CreatedTimestamp"`

	// 模板名称
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 模板状态。1-审核中|0-已通过|2-拒绝|其它-不可用
	TemplateStatus *int64 `json:"TemplateStatus,omitnil,omitempty" name:"TemplateStatus"`

	// 模板ID
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// 审核原因
	ReviewReason *string `json:"ReviewReason,omitnil,omitempty" name:"ReviewReason"`
}

type TimedEmailParam struct {
	// 定时发送邮件的开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`
}

// Predefined struct for user
type UpdateCustomBlackListRequestParams struct {
	// 需要更改的黑名单id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 修改后的邮件地址
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 过期时间，为空则表示永久有效
	ExpireDate *string `json:"ExpireDate,omitnil,omitempty" name:"ExpireDate"`
}

type UpdateCustomBlackListRequest struct {
	*tchttp.BaseRequest
	
	// 需要更改的黑名单id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 修改后的邮件地址
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 过期时间，为空则表示永久有效
	ExpireDate *string `json:"ExpireDate,omitnil,omitempty" name:"ExpireDate"`
}

func (r *UpdateCustomBlackListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCustomBlackListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Email")
	delete(f, "ExpireDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCustomBlackListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCustomBlackListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCustomBlackListResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCustomBlackListResponseParams `json:"Response"`
}

func (r *UpdateCustomBlackListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCustomBlackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEmailIdentityRequestParams struct {
	// 请求验证的域名
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

type UpdateEmailIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 请求验证的域名
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

func (r *UpdateEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailIdentity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEmailIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEmailIdentityResponseParams struct {
	// 验证类型。固定值：DOMAIN
	IdentityType *string `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// 是否已通过验证
	VerifiedForSendingStatus *bool `json:"VerifiedForSendingStatus,omitnil,omitempty" name:"VerifiedForSendingStatus"`

	// 需要配置的DNS信息
	Attributes []*DNSAttributes `json:"Attributes,omitnil,omitempty" name:"Attributes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEmailIdentityResponseParams `json:"Response"`
}

func (r *UpdateEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEmailSmtpPassWordRequestParams struct {
	// smtp密码，长度限制64
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 发信邮箱,长度限制128
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`
}

type UpdateEmailSmtpPassWordRequest struct {
	*tchttp.BaseRequest
	
	// smtp密码，长度限制64
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 发信邮箱,长度限制128
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`
}

func (r *UpdateEmailSmtpPassWordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailSmtpPassWordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Password")
	delete(f, "EmailAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEmailSmtpPassWordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEmailSmtpPassWordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateEmailSmtpPassWordResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEmailSmtpPassWordResponseParams `json:"Response"`
}

func (r *UpdateEmailSmtpPassWordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailSmtpPassWordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEmailTemplateRequestParams struct {
	// 模板内容
	TemplateContent *TemplateContent `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`

	// 模板ID
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// 模板名字
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`
}

type UpdateEmailTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板内容
	TemplateContent *TemplateContent `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`

	// 模板ID
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// 模板名字
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`
}

func (r *UpdateEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateContent")
	delete(f, "TemplateID")
	delete(f, "TemplateName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEmailTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEmailTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEmailTemplateResponseParams `json:"Response"`
}

func (r *UpdateEmailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Volume struct {
	// 日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	SendDate *string `json:"SendDate,omitnil,omitempty" name:"SendDate"`

	// 邮件请求数量
	RequestCount *uint64 `json:"RequestCount,omitnil,omitempty" name:"RequestCount"`

	// 腾讯云通过数量
	AcceptedCount *uint64 `json:"AcceptedCount,omitnil,omitempty" name:"AcceptedCount"`

	// 送达数量
	DeliveredCount *uint64 `json:"DeliveredCount,omitnil,omitempty" name:"DeliveredCount"`

	// 打开邮件的用户数量，根据收件人去重
	OpenedCount *uint64 `json:"OpenedCount,omitnil,omitempty" name:"OpenedCount"`

	// 点击了邮件中的链接数量用户数量
	ClickedCount *uint64 `json:"ClickedCount,omitnil,omitempty" name:"ClickedCount"`

	// 退信数量
	BounceCount *uint64 `json:"BounceCount,omitnil,omitempty" name:"BounceCount"`

	// 取消订阅的用户数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnsubscribeCount *uint64 `json:"UnsubscribeCount,omitnil,omitempty" name:"UnsubscribeCount"`
}