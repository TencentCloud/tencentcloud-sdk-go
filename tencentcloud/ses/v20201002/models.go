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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BlackEmailAddress struct {

	// 邮箱被拉黑时间
	BounceTime *string `json:"BounceTime,omitempty" name:"BounceTime"`

	// 被拉黑的邮箱地址
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`
}

type CreateEmailAddressRequest struct {
	*tchttp.BaseRequest

	// 您的发信地址，上限为10个
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// 发件人别名
	EmailSenderName *string `json:"EmailSenderName,omitempty" name:"EmailSenderName"`
}

func (r *CreateEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEmailAddressRequest) FromJsonString(s string) error {
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

func (r *CreateEmailIdentityRequest) FromJsonString(s string) error {
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
		Attributes []*DNSAttributes `json:"Attributes,omitempty" name:"Attributes" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

func (r *CreateEmailTemplateRequest) FromJsonString(s string) error {
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

func (r *CreateEmailTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	EmailAddressList []*string `json:"EmailAddressList,omitempty" name:"EmailAddressList" list`
}

func (r *DeleteBlackListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBlackListRequest) FromJsonString(s string) error {
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

func (r *DeleteEmailAddressRequest) FromJsonString(s string) error {
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

func (r *DeleteEmailIdentityRequest) FromJsonString(s string) error {
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

func (r *DeleteEmailIdentityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteEmailTemplateRequest struct {
	*tchttp.BaseRequest

	// 删除发信模版
	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`
}

func (r *DeleteEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEmailTemplateRequest) FromJsonString(s string) error {
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

func (r *GetEmailIdentityRequest) FromJsonString(s string) error {
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
		Attributes []*DNSAttributes `json:"Attributes,omitempty" name:"Attributes" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

func (r *GetEmailTemplateRequest) FromJsonString(s string) error {
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

func (r *GetEmailTemplateResponse) FromJsonString(s string) error {
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

func (r *GetStatisticsReportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetStatisticsReportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发信统计报告，按天
		DailyVolumes []*Volume `json:"DailyVolumes,omitempty" name:"DailyVolumes" list`

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

func (r *GetStatisticsReportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListBlackEmailAddressRequest struct {
	*tchttp.BaseRequest

	// 开始日期
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 规范，配合Offset使用
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 规范，配合Limit使用
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

func (r *ListBlackEmailAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListBlackEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 黑名单列表
		BlackList []*BlackEmailAddress `json:"BlackList,omitempty" name:"BlackList" list`

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

func (r *ListEmailAddressRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发信地址列表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		EmailSenders []*EmailSender `json:"EmailSenders,omitempty" name:"EmailSenders" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

func (r *ListEmailIdentitiesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListEmailIdentitiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发信域名列表
		EmailIdentities []*EmailIdentity `json:"EmailIdentities,omitempty" name:"EmailIdentities" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEmailIdentitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

func (r *ListEmailTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListEmailTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 邮件模板列表
		TemplatesMetadata []*TemplatesMetadata `json:"TemplatesMetadata,omitempty" name:"TemplatesMetadata" list`

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

func (r *ListEmailTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendEmailRequest struct {
	*tchttp.BaseRequest

	// 发信邮件地址。例如：noreply@mail.qcloud.com。
	FromEmailAddress *string `json:"FromEmailAddress,omitempty" name:"FromEmailAddress"`

	// 收信人邮箱地址
	Destination []*string `json:"Destination,omitempty" name:"Destination" list`

	// 邮件主题
	Subject *string `json:"Subject,omitempty" name:"Subject"`

	// 邮件的“回复”电子邮件地址。可以填写您能收到邮件的邮箱地址，可以是个人邮箱。如果不填，收件人将会回复到腾讯云。
	ReplyToAddresses *string `json:"ReplyToAddresses,omitempty" name:"ReplyToAddresses"`

	// 使用模板发送时，填写的模板相关参数
	Template *Template `json:"Template,omitempty" name:"Template"`

	// 使用API直接发送内容时，填写的邮件内容
	Simple *Simple `json:"Simple,omitempty" name:"Simple"`
}

func (r *SendEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendEmailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendEmailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 接受消息时生成的消息的唯一标识符。
		MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendEmailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type UpdateEmailIdentityRequest struct {
	*tchttp.BaseRequest

	// 请求验证的域名
	EmailIdentity *string `json:"EmailIdentity,omitempty" name:"EmailIdentity"`
}

func (r *UpdateEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateEmailIdentityRequest) FromJsonString(s string) error {
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
		Attributes []*DNSAttributes `json:"Attributes,omitempty" name:"Attributes" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

func (r *UpdateEmailTemplateRequest) FromJsonString(s string) error {
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
