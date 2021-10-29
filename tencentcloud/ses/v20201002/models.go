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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Attachment struct {

	// 附件名称，最大支持255个字符长度，不支持部分附件类型，详情请参考[附件类型](https://cloud.tencent.com/document/product/1288/51951)。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// base64之后的附件内容，您可以发送的附件大小上限为5 MB。 注意：腾讯云api目前要求请求包大小不得超过10 MB。如果您要发送多个附件，那么这些附件的总大小不能超过10 MB。
	Content *string `json:"Content,omitempty" name:"Content"`
}

type BatchSendEmailRequest struct {
	*tchttp.BaseRequest

	// 发信邮件地址。请填写发件人邮箱地址，例如：noreply@mail.qcloud.com。如需填写发件人说明，请按照
	// 发信人 <邮件地址> 的方式填写，例如：
	// 腾讯云团队 <noreply@mail.qcloud.com>
	FromEmailAddress *string `json:"FromEmailAddress,omitempty" name:"FromEmailAddress"`

	// 收件人列表ID
	ReceiverId *uint64 `json:"ReceiverId,omitempty" name:"ReceiverId"`

	// 邮件主题
	Subject *string `json:"Subject,omitempty" name:"Subject"`

	// 任务类型 1即时 2 定时 3 周期
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`

	// 邮件的“回复”电子邮件地址。可以填写您能收到邮件的邮箱地址，可以是个人邮箱。如果不填，收件人将会回复到腾讯云。
	ReplyToAddresses *string `json:"ReplyToAddresses,omitempty" name:"ReplyToAddresses"`

	// 使用模板发送时，填写的模板相关参数
	Template *Template `json:"Template,omitempty" name:"Template"`

	// 使用API直接发送内容时，填写的邮件内容
	Simple *Simple `json:"Simple,omitempty" name:"Simple"`

	// 需要发送附件时，填写附件相关参数。
	Attachments []*Attachment `json:"Attachments,omitempty" name:"Attachments"`

	// 周期发送任务的必要参数
	CycleParam *CycleEmailParam `json:"CycleParam,omitempty" name:"CycleParam"`

	// 定时发送任务的必要参数
	TimedParam *TimedEmailParam `json:"TimedParam,omitempty" name:"TimedParam"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchSendEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BatchSendEmailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发送任务ID
		TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type BlackEmailAddress struct {

	// 邮箱被拉黑时间
	BounceTime *string `json:"BounceTime,omitempty" name:"BounceTime"`

	// 被拉黑的邮箱地址
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`
}

type CreateEmailAddressRequest struct {
	*tchttp.BaseRequest

	// 您的发信地址（发信地址总数上限为10个）
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// 发件人别名
	EmailSenderName *string `json:"EmailSenderName,omitempty" name:"EmailSenderName"`
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

type CreateEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateEmailIdentityRequest struct {
	*tchttp.BaseRequest

	// 您的发信域名，建议使用三级以上域名。例如：mail.qcloud.com。
	EmailIdentity *string `json:"EmailIdentity,omitempty" name:"EmailIdentity"`
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

type CreateEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证类型。固定值：DOMAIN
		IdentityType *string `json:"IdentityType,omitempty" name:"IdentityType"`

		// 是否已通过验证
		VerifiedForSendingStatus *bool `json:"VerifiedForSendingStatus,omitempty" name:"VerifiedForSendingStatus"`

		// 需要配置的DNS信息
		Attributes []*DNSAttributes `json:"Attributes,omitempty" name:"Attributes"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateEmailTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模板内容
	TemplateContent *TemplateContent `json:"TemplateContent,omitempty" name:"TemplateContent"`
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

type CreateEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CycleEmailParam struct {

	// 任务开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 任务周期 小时维度
	IntervalTime *uint64 `json:"IntervalTime,omitempty" name:"IntervalTime"`
}

type DNSAttributes struct {

	// 记录类型 CNAME | A | TXT | MX
	Type *string `json:"Type,omitempty" name:"Type"`

	// 域名
	SendDomain *string `json:"SendDomain,omitempty" name:"SendDomain"`

	// 需要配置的值
	ExpectedValue *string `json:"ExpectedValue,omitempty" name:"ExpectedValue"`

	// 腾讯云目前检测到的值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 检测是否通过，创建时默认为false
	Status *bool `json:"Status,omitempty" name:"Status"`
}

type DeleteBlackListRequest struct {
	*tchttp.BaseRequest

	// 需要清除的黑名单邮箱列表，数组长度至少为1
	EmailAddressList []*string `json:"EmailAddressList,omitempty" name:"EmailAddressList"`
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

type DeleteBlackListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteEmailAddressRequest struct {
	*tchttp.BaseRequest

	// 发信地址
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`
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

type DeleteEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteEmailIdentityRequest struct {
	*tchttp.BaseRequest

	// 发信域名
	EmailIdentity *string `json:"EmailIdentity,omitempty" name:"EmailIdentity"`
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

type DeleteEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteEmailTemplateRequest struct {
	*tchttp.BaseRequest

	// 模版ID
	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`
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

type DeleteEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type EmailIdentity struct {

	// 发信域名
	IdentityName *string `json:"IdentityName,omitempty" name:"IdentityName"`

	// 验证类型，固定为DOMAIN
	IdentityType *string `json:"IdentityType,omitempty" name:"IdentityType"`

	// 是否已通过验证
	SendingEnabled *bool `json:"SendingEnabled,omitempty" name:"SendingEnabled"`
}

type EmailSender struct {

	// 发信地址
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// 发信人别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmailSenderName *string `json:"EmailSenderName,omitempty" name:"EmailSenderName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTimestamp *uint64 `json:"CreatedTimestamp,omitempty" name:"CreatedTimestamp"`
}

type GetEmailIdentityRequest struct {
	*tchttp.BaseRequest

	// 发信域名
	EmailIdentity *string `json:"EmailIdentity,omitempty" name:"EmailIdentity"`
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

type GetEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证类型。固定值：DOMAIN
		IdentityType *string `json:"IdentityType,omitempty" name:"IdentityType"`

		// 是否已通过验证
		VerifiedForSendingStatus *bool `json:"VerifiedForSendingStatus,omitempty" name:"VerifiedForSendingStatus"`

		// DNS配置详情
		Attributes []*DNSAttributes `json:"Attributes,omitempty" name:"Attributes"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type GetEmailTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板ID
	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`
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

type GetEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板内容数据
		TemplateContent *TemplateContent `json:"TemplateContent,omitempty" name:"TemplateContent"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type GetSendEmailStatusRequest struct {
	*tchttp.BaseRequest

	// 发送的日期，必填。仅支持查询某个日期，不支持范围查询。
	RequestDate *string `json:"RequestDate,omitempty" name:"RequestDate"`

	// 偏移量。默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取最大条数，最多 100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// SendMail接口返回的MessageId字段。
	MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

	// 收件人邮箱。
	ToEmailAddress *string `json:"ToEmailAddress,omitempty" name:"ToEmailAddress"`
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

type GetSendEmailStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 邮件发送状态列表
		EmailStatusList []*SendEmailStatus `json:"EmailStatusList,omitempty" name:"EmailStatusList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type GetStatisticsReportRequest struct {
	*tchttp.BaseRequest

	// 开始日期
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 发信域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 收件方邮箱类型，例如gmail.com
	ReceivingMailboxType *string `json:"ReceivingMailboxType,omitempty" name:"ReceivingMailboxType"`
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

type GetStatisticsReportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发信统计报告，按天
		DailyVolumes []*Volume `json:"DailyVolumes,omitempty" name:"DailyVolumes"`

		// 发信统计报告，总览
		OverallVolume *Volume `json:"OverallVolume,omitempty" name:"OverallVolume"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ListBlackEmailAddressRequest struct {
	*tchttp.BaseRequest

	// 开始日期，格式为YYYY-MM-DD
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式为YYYY-MM-DD
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 规范，配合Offset使用
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 规范，配合Limit使用，Limit最大取值为100
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 可以指定邮箱进行查询
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// 可以指定任务ID进行查询
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
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

type ListBlackEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 黑名单列表
		BlackList []*BlackEmailAddress `json:"BlackList,omitempty" name:"BlackList"`

		// 黑名单总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ListEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发信地址列表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		EmailSenders []*EmailSender `json:"EmailSenders,omitempty" name:"EmailSenders"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ListEmailIdentitiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发信域名列表
		EmailIdentities []*EmailIdentity `json:"EmailIdentities,omitempty" name:"EmailIdentities"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ListEmailTemplatesRequest struct {
	*tchttp.BaseRequest

	// 获取模版数据量，用于分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 获取模版偏移值，用于分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
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

type ListEmailTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 邮件模板列表
		TemplatesMetadata []*TemplatesMetadata `json:"TemplatesMetadata,omitempty" name:"TemplatesMetadata"`

		// 模版总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type SendEmailRequest struct {
	*tchttp.BaseRequest

	// 发信邮件地址。请填写发件人邮箱地址，例如：noreply@mail.qcloud.com。如需填写发件人说明，请按照 
	// 发信人 <邮件地址> 的方式填写，例如：
	// 腾讯云团队 <noreply@mail.qcloud.com>
	FromEmailAddress *string `json:"FromEmailAddress,omitempty" name:"FromEmailAddress"`

	// 收信人邮箱地址，最多支持群发50人。注意：邮件内容会显示所有收件人地址，非群发邮件请多次调用API发送。
	Destination []*string `json:"Destination,omitempty" name:"Destination"`

	// 邮件主题
	Subject *string `json:"Subject,omitempty" name:"Subject"`

	// 邮件的“回复”电子邮件地址。可以填写您能收到邮件的邮箱地址，可以是个人邮箱。如果不填，收件人将会回复到腾讯云。
	ReplyToAddresses *string `json:"ReplyToAddresses,omitempty" name:"ReplyToAddresses"`

	// 使用模板发送时，填写的模板相关参数
	Template *Template `json:"Template,omitempty" name:"Template"`

	// 使用API直接发送内容时，填写的邮件内容
	Simple *Simple `json:"Simple,omitempty" name:"Simple"`

	// 需要发送附件时，填写附件相关参数。
	Attachments []*Attachment `json:"Attachments,omitempty" name:"Attachments"`
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
	delete(f, "Template")
	delete(f, "Simple")
	delete(f, "Attachments")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendEmailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接受消息生成的唯一消息标识符。
		MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

	// 收件人邮箱
	ToEmailAddress *string `json:"ToEmailAddress,omitempty" name:"ToEmailAddress"`

	// 发件人邮箱
	FromEmailAddress *string `json:"FromEmailAddress,omitempty" name:"FromEmailAddress"`

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
	// 1009: 内部系统异常
	// 1010: 超出了每日发送限制
	// 1011: 无发送自定义内容权限，必须使用模板
	// 2001: 找不到相关记录
	// 3007: 模板ID无效或者不可用
	// 3008: 模板状态异常
	// 3009: 无权限使用该模板
	// 3010: TemplateData字段格式不正确 
	// 3014: 发件域名没有经过认证，无法发送
	// 3020: 收件方邮箱类型在黑名单
	// 3024: 邮箱地址格式预检查失败
	// 3030: 退信率过高，临时限制发送
	// 3033: 余额不足，账号欠费等
	SendStatus *int64 `json:"SendStatus,omitempty" name:"SendStatus"`

	// 收件方处理状态
	// 0: 请求成功被腾讯云接受，进入发送队列
	// 1: 邮件递送成功，DeliverTime表示递送成功的时间
	// 2: 邮件因某种原因被丢弃，DeliverMessage表示丢弃原因
	// 3: 收件方ESP拒信，一般原因为邮箱地址不存在，或其它原因
	// 8: 邮件被ESP因某些原因延迟递送，DeliverMessage表示延迟原因
	DeliverStatus *int64 `json:"DeliverStatus,omitempty" name:"DeliverStatus"`

	// 收件方处理状态描述
	DeliverMessage *string `json:"DeliverMessage,omitempty" name:"DeliverMessage"`

	// 请求到达腾讯云时间戳
	RequestTime *int64 `json:"RequestTime,omitempty" name:"RequestTime"`

	// 腾讯云执行递送时间戳
	DeliverTime *int64 `json:"DeliverTime,omitempty" name:"DeliverTime"`

	// 用户是否打开该邮件
	UserOpened *bool `json:"UserOpened,omitempty" name:"UserOpened"`

	// 用户是否点击该邮件中的链接
	UserClicked *bool `json:"UserClicked,omitempty" name:"UserClicked"`

	// 用户是否取消该发送者的订阅
	UserUnsubscribed *bool `json:"UserUnsubscribed,omitempty" name:"UserUnsubscribed"`

	// 用户是否举报该发送者
	UserComplainted *bool `json:"UserComplainted,omitempty" name:"UserComplainted"`
}

type Simple struct {

	// base64之后的Html代码。需要包含所有的代码信息，不要包含外部css，否则会导致显示格式错乱
	Html *string `json:"Html,omitempty" name:"Html"`

	// base64之后的纯文本信息，如果没有Html，邮件中会直接显示纯文本；如果有Html，它代表邮件的纯文本样式
	Text *string `json:"Text,omitempty" name:"Text"`
}

type Template struct {

	// 模板ID。如果没有模板，请先新建一个
	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`

	// 模板中的变量参数，请使用json.dump将json对象格式化为string类型。该对象是一组键值对，每个Key代表模板中的一个变量，模板中的变量使用{{键}}表示，相应的值在发送时会被替换为{{值}}。
	TemplateData *string `json:"TemplateData,omitempty" name:"TemplateData"`
}

type TemplateContent struct {

	// base64之后的Html代码
	Html *string `json:"Html,omitempty" name:"Html"`

	// base64之后的文本内容
	Text *string `json:"Text,omitempty" name:"Text"`
}

type TemplatesMetadata struct {

	// 创建时间
	CreatedTimestamp *uint64 `json:"CreatedTimestamp,omitempty" name:"CreatedTimestamp"`

	// 模板名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模板状态。1-审核中|0-已通过|2-拒绝|其它-不可用
	TemplateStatus *int64 `json:"TemplateStatus,omitempty" name:"TemplateStatus"`

	// 模板ID
	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`

	// 审核原因
	ReviewReason *string `json:"ReviewReason,omitempty" name:"ReviewReason"`
}

type TimedEmailParam struct {

	// 定时发送邮件的开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
}

type UpdateEmailIdentityRequest struct {
	*tchttp.BaseRequest

	// 请求验证的域名
	EmailIdentity *string `json:"EmailIdentity,omitempty" name:"EmailIdentity"`
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

type UpdateEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证类型。固定值：DOMAIN
		IdentityType *string `json:"IdentityType,omitempty" name:"IdentityType"`

		// 是否已通过验证
		VerifiedForSendingStatus *bool `json:"VerifiedForSendingStatus,omitempty" name:"VerifiedForSendingStatus"`

		// 需要配置的DNS信息
		Attributes []*DNSAttributes `json:"Attributes,omitempty" name:"Attributes"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type UpdateEmailTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板内容
	TemplateContent *TemplateContent `json:"TemplateContent,omitempty" name:"TemplateContent"`

	// 模板ID
	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`

	// 模版名字
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
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

type UpdateEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	SendDate *string `json:"SendDate,omitempty" name:"SendDate"`

	// 邮件请求数量
	RequestCount *uint64 `json:"RequestCount,omitempty" name:"RequestCount"`

	// 腾讯云通过数量
	AcceptedCount *uint64 `json:"AcceptedCount,omitempty" name:"AcceptedCount"`

	// 送达数量
	DeliveredCount *uint64 `json:"DeliveredCount,omitempty" name:"DeliveredCount"`

	// 打开邮件的用户数量，根据收件人去重
	OpenedCount *uint64 `json:"OpenedCount,omitempty" name:"OpenedCount"`

	// 点击了邮件中的链接数量用户数量
	ClickedCount *uint64 `json:"ClickedCount,omitempty" name:"ClickedCount"`

	// 退信数量
	BounceCount *uint64 `json:"BounceCount,omitempty" name:"BounceCount"`

	// 取消订阅的用户数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnsubscribeCount *uint64 `json:"UnsubscribeCount,omitempty" name:"UnsubscribeCount"`
}
