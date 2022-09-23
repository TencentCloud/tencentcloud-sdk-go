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

package v20201111

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Agent struct {

}

type ApproverInfo struct {
	// 参与者类型：
	// 0：企业
	// 1：个人
	// 3：企业静默签署
	// 注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。静默签署仅进行盖章操作，不能自动签名。
	ApproverType *int64 `json:"ApproverType,omitempty" name:"ApproverType"`

	// 本环节需要操作人的名字
	ApproverName *string `json:"ApproverName,omitempty" name:"ApproverName"`

	// 本环节需要操作人的手机号
	ApproverMobile *string `json:"ApproverMobile,omitempty" name:"ApproverMobile"`

	// 本环节操作人签署控件配置
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 如果是企业,则为企业的名字
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 身份证号
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitempty" name:"ApproverIdCardNumber"`

	// 证件类型 
	// ID_CARD 身份证
	// HONGKONG_AND_MACAO 港澳居民来往内地通行证
	// HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证(格式同居民身份证)
	ApproverIdCardType *string `json:"ApproverIdCardType,omitempty" name:"ApproverIdCardType"`

	// sms--短信，none--不通知
	NotifyType *string `json:"NotifyType,omitempty" name:"NotifyType"`

	// 1--收款人、2--开具人、3--见证人
	ApproverRole *int64 `json:"ApproverRole,omitempty" name:"ApproverRole"`

	// 签署意愿确认渠道,WEIXINAPP:人脸识别
	VerifyChannel []*string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 合同的强制预览时间：3~300s，未指定则按合同页数计算
	PreReadTime *int64 `json:"PreReadTime,omitempty" name:"PreReadTime"`

	// 签署人userId，非企微场景不使用此字段
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 签署人用户来源,企微侧用户请传入：WEWORKAPP
	ApproverSource *string `json:"ApproverSource,omitempty" name:"ApproverSource"`

	// 客户自定义签署人标识，64位长度，保证唯一，非企微场景不使用此字段
	CustomApproverTag *string `json:"CustomApproverTag,omitempty" name:"CustomApproverTag"`
}

type ApproverRestriction struct {
	// 指定签署人名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 指定签署人手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 指定签署人证件类型
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 指定签署人证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`
}

type Caller struct {
	// 应用号
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 主机构ID
	OrganizationId *string `json:"OrganizationId,omitempty" name:"OrganizationId"`

	// 经办人的用户ID
	OperatorId *string `json:"OperatorId,omitempty" name:"OperatorId"`

	// 下属机构ID
	SubOrganizationId *string `json:"SubOrganizationId,omitempty" name:"SubOrganizationId"`
}

// Predefined struct for user
type CancelFlowRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 撤销原因，最长200个字符；
	CancelMessage *string `json:"CancelMessage,omitempty" name:"CancelMessage"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

type CancelFlowRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 撤销原因，最长200个字符；
	CancelMessage *string `json:"CancelMessage,omitempty" name:"CancelMessage"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

func (r *CancelFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "CancelMessage")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelFlowResponse struct {
	*tchttp.BaseResponse
	Response *CancelFlowResponseParams `json:"Response"`
}

func (r *CancelFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelMultiFlowSignQRCodeRequestParams struct {
	// 用户信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 二维码id
	QrCodeId *string `json:"QrCodeId,omitempty" name:"QrCodeId"`

	// 应用信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

type CancelMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 用户信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 二维码id
	QrCodeId *string `json:"QrCodeId,omitempty" name:"QrCodeId"`

	// 应用信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

func (r *CancelMultiFlowSignQRCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelMultiFlowSignQRCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "QrCodeId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelMultiFlowSignQRCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelMultiFlowSignQRCodeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelMultiFlowSignQRCodeResponse struct {
	*tchttp.BaseResponse
	Response *CancelMultiFlowSignQRCodeResponseParams `json:"Response"`
}

func (r *CancelMultiFlowSignQRCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelMultiFlowSignQRCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcInfo struct {
	// 被抄送人手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
}

type Component struct {
	// 如果是Component控件类型，则可选的字段为：
	// TEXT - 普通文本控件；
	// MULTI_LINE_TEXT - 多行文本控件；
	// CHECK_BOX - 勾选框控件；
	// FILL_IMAGE - 图片控件；
	// DYNAMIC_TABLE - 动态表格控件；
	// ATTACHMENT - 附件控件；
	// SELECTOR - 选择器控件；
	// 
	// 如果是SignComponent控件类型，则可选的字段为
	// SIGN_SEAL - 签署印章控件；
	// SIGN_DATE - 签署日期控件；
	// SIGN_SIGNATURE - 用户签名控件；
	// SIGN_PERSONAL_SEAL - 个人签署印章控件（使用文件发起暂不支持此类型）；
	// 
	// 表单域的控件不能作为印章和签名控件
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// 参数控件宽度，单位pt
	ComponentWidth *float64 `json:"ComponentWidth,omitempty" name:"ComponentWidth"`

	// 参数控件高度，单位pt
	ComponentHeight *float64 `json:"ComponentHeight,omitempty" name:"ComponentHeight"`

	// 参数控件所在页码，取值为：1-N
	ComponentPage *int64 `json:"ComponentPage,omitempty" name:"ComponentPage"`

	// 参数控件X位置，单位pt
	ComponentPosX *float64 `json:"ComponentPosX,omitempty" name:"ComponentPosX"`

	// 参数控件Y位置，单位pt
	ComponentPosY *float64 `json:"ComponentPosY,omitempty" name:"ComponentPosY"`

	// 控件所属文件的序号（模板中的resourceId排列序号，取值为：0-N）
	FileIndex *int64 `json:"FileIndex,omitempty" name:"FileIndex"`

	// GenerateMode==KEYWORD 指定关键字
	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`

	// GenerateMode==FIELD 指定表单域名称
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// 是否必选，默认为false
	ComponentRequired *bool `json:"ComponentRequired,omitempty" name:"ComponentRequired"`

	// 扩展参数：
	// ComponentType为SIGN_SIGNATURE类型可以控制签署方式
	// {“ComponentTypeLimit”: [“xxx”]}
	// xxx可以为：
	// HANDWRITE – 手写签名
	// BORDERLESS_ESIGN – 自动生成无边框腾讯体
	// OCR_ESIGN -- AI智能识别手写签名
	// ESIGN -- 个人印章类型
	// 如：{“ComponentTypeLimit”: [“BORDERLESS_ESIGN”]}
	ComponentExtra *string `json:"ComponentExtra,omitempty" name:"ComponentExtra"`

	// 控件关联的签署人ID
	ComponentRecipientId *string `json:"ComponentRecipientId,omitempty" name:"ComponentRecipientId"`

	// 控件填充vaule，ComponentType和传入值类型对应关系：
	// TEXT - 文本内容
	// MULTI_LINE_TEXT - 文本内容
	// CHECK_BOX - true/false
	// FILL_IMAGE、ATTACHMENT - 附件的FileId，需要通过UploadFiles接口上传获取
	// SELECTOR - 选项值
	// DYNAMIC_TABLE - 传入json格式的表格内容，具体见数据结构FlowInfo：https://cloud.tencent.com/document/api/1420/61525#FlowInfo
	ComponentValue *string `json:"ComponentValue,omitempty" name:"ComponentValue"`

	// 是否是表单域类型，默认不存在
	IsFormType *bool `json:"IsFormType,omitempty" name:"IsFormType"`

	// NORMAL 正常模式，使用坐标制定签署控件位置
	// FIELD 表单域，需使用ComponentName指定表单域名称
	// KEYWORD 关键字，使用ComponentId指定关键字
	GenerateMode *string `json:"GenerateMode,omitempty" name:"GenerateMode"`

	// 日期控件类型字号
	ComponentDateFontSize *int64 `json:"ComponentDateFontSize,omitempty" name:"ComponentDateFontSize"`

	// 指定关键字时横坐标偏移量，单位pt
	OffsetX *float64 `json:"OffsetX,omitempty" name:"OffsetX"`

	// 指定关键字时纵坐标偏移量，单位pt
	OffsetY *float64 `json:"OffsetY,omitempty" name:"OffsetY"`
}

// Predefined struct for user
type CreateBatchCancelFlowUrlRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 需要执行撤回的签署流程id数组，最多100个
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`
}

type CreateBatchCancelFlowUrlRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 需要执行撤回的签署流程id数组，最多100个
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`
}

func (r *CreateBatchCancelFlowUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchCancelFlowUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchCancelFlowUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchCancelFlowUrlResponseParams struct {
	// 批量撤回签署流程链接
	BatchCancelFlowUrl *string `json:"BatchCancelFlowUrl,omitempty" name:"BatchCancelFlowUrl"`

	// 签署流程撤回失败信息
	FailMessages []*string `json:"FailMessages,omitempty" name:"FailMessages"`

	// 签署连接过期时间字符串：年月日-时分秒
	UrlExpireOn *string `json:"UrlExpireOn,omitempty" name:"UrlExpireOn"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBatchCancelFlowUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchCancelFlowUrlResponseParams `json:"Response"`
}

func (r *CreateBatchCancelFlowUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchCancelFlowUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConvertTaskApiRequestParams struct {
	// 资源类型 取值范围doc,docx,html之一
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，长度限制为256字符
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源Id，通过UploadFiles获取
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 操作者信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 应用号信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`
}

type CreateConvertTaskApiRequest struct {
	*tchttp.BaseRequest
	
	// 资源类型 取值范围doc,docx,html之一
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，长度限制为256字符
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源Id，通过UploadFiles获取
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 操作者信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 应用号信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`
}

func (r *CreateConvertTaskApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConvertTaskApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	delete(f, "ResourceName")
	delete(f, "ResourceId")
	delete(f, "Operator")
	delete(f, "Agent")
	delete(f, "Organization")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConvertTaskApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConvertTaskApiResponseParams struct {
	// 转换任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateConvertTaskApiResponse struct {
	*tchttp.BaseResponse
	Response *CreateConvertTaskApiResponseParams `json:"Response"`
}

func (r *CreateConvertTaskApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConvertTaskApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocumentRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程编号,由CreateFlow接口返回
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 用户上传的模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 文件名列表，单个文件名最大长度200个字符，暂时仅支持单文件发起
	FileNames []*string `json:"FileNames,omitempty" name:"FileNames"`

	// 内容控件信息数组
	FormFields []*FormField `json:"FormFields,omitempty" name:"FormFields"`

	// 是否需要生成预览文件 默认不生成；
	// 预览链接有效期300秒；
	NeedPreview *bool `json:"NeedPreview,omitempty" name:"NeedPreview"`

	// 客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

type CreateDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程编号,由CreateFlow接口返回
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 用户上传的模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 文件名列表，单个文件名最大长度200个字符，暂时仅支持单文件发起
	FileNames []*string `json:"FileNames,omitempty" name:"FileNames"`

	// 内容控件信息数组
	FormFields []*FormField `json:"FormFields,omitempty" name:"FormFields"`

	// 是否需要生成预览文件 默认不生成；
	// 预览链接有效期300秒；
	NeedPreview *bool `json:"NeedPreview,omitempty" name:"NeedPreview"`

	// 客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

func (r *CreateDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "TemplateId")
	delete(f, "FileNames")
	delete(f, "FormFields")
	delete(f, "NeedPreview")
	delete(f, "ClientToken")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocumentResponseParams struct {
	// 签署流程电子文档ID
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

	// 签署流程文件的预览地址, 5分钟内有效。仅当NeedPreview为true 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewFileUrl *string `json:"PreviewFileUrl,omitempty" name:"PreviewFileUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDocumentResponse struct {
	*tchttp.BaseResponse
	Response *CreateDocumentResponseParams `json:"Response"`
}

func (r *CreateDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowApproversRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 补充签署人信息
	Approvers []*FillApproverInfo `json:"Approvers,omitempty" name:"Approvers"`
}

type CreateFlowApproversRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 补充签署人信息
	Approvers []*FillApproverInfo `json:"Approvers,omitempty" name:"Approvers"`
}

func (r *CreateFlowApproversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowApproversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "Approvers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowApproversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowApproversResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFlowApproversResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowApproversResponseParams `json:"Response"`
}

func (r *CreateFlowApproversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowApproversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowByFilesRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程名称,最大长度200个字符
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署参与者信息，最大限制50方
	Approvers []*ApproverInfo `json:"Approvers,omitempty" name:"Approvers"`

	// 签署pdf文件的资源编号列表，通过UploadFiles接口获取，暂时仅支持单文件发起
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 签署流程的类型(如销售合同/入职合同等)，最大长度200个字符
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 经办人内容控件配置
	Components []*Component `json:"Components,omitempty" name:"Components"`

	// 被抄送人的信息列表。
	// 注:此功能为白名单功能，若有需要，请联系电子签客服开白使用
	CcInfos []*CcInfo `json:"CcInfos,omitempty" name:"CcInfos"`

	// 是否需要预览，true：预览模式，false：非预览（默认）；
	// 预览链接有效期300秒；
	// 
	// 注：如果使用“预览模式”，出参会返回合同预览链接 PreviewUrl，不会正式发起合同，且出参不会返回签署流程编号 FlowId；如果使用“非预览”，则会正常返回签署流程编号 FlowId，不会生成合同预览链接 PreviewUrl。
	NeedPreview *bool `json:"NeedPreview,omitempty" name:"NeedPreview"`

	// 签署流程描述,最大长度1000个字符
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 签署流程的签署截止时间。
	// 值为unix时间戳,精确到秒,不传默认为当前时间一年后
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 发送类型：
	// true：无序签
	// false：有序签
	// 注：默认为false（有序签）
	Unordered *bool `json:"Unordered,omitempty" name:"Unordered"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitempty" name:"CustomShowMap"`

	// 发起方企业的签署人进行签署操作是否需要企业内部审批。使用此功能需要发起方企业有参与签署。
	// 若设置为true，审核结果需通过接口 CreateFlowSignReview 通知电子签，审核通过后，发起方企业签署人方可进行签署操作，否则会阻塞其签署操作。
	// 
	// 注：企业可以通过此功能与企业内部的审批流程进行关联，支持手动、静默签署合同。
	NeedSignReview *bool `json:"NeedSignReview,omitempty" name:"NeedSignReview"`

	// 应用号信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

type CreateFlowByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程名称,最大长度200个字符
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署参与者信息，最大限制50方
	Approvers []*ApproverInfo `json:"Approvers,omitempty" name:"Approvers"`

	// 签署pdf文件的资源编号列表，通过UploadFiles接口获取，暂时仅支持单文件发起
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 签署流程的类型(如销售合同/入职合同等)，最大长度200个字符
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 经办人内容控件配置
	Components []*Component `json:"Components,omitempty" name:"Components"`

	// 被抄送人的信息列表。
	// 注:此功能为白名单功能，若有需要，请联系电子签客服开白使用
	CcInfos []*CcInfo `json:"CcInfos,omitempty" name:"CcInfos"`

	// 是否需要预览，true：预览模式，false：非预览（默认）；
	// 预览链接有效期300秒；
	// 
	// 注：如果使用“预览模式”，出参会返回合同预览链接 PreviewUrl，不会正式发起合同，且出参不会返回签署流程编号 FlowId；如果使用“非预览”，则会正常返回签署流程编号 FlowId，不会生成合同预览链接 PreviewUrl。
	NeedPreview *bool `json:"NeedPreview,omitempty" name:"NeedPreview"`

	// 签署流程描述,最大长度1000个字符
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 签署流程的签署截止时间。
	// 值为unix时间戳,精确到秒,不传默认为当前时间一年后
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 发送类型：
	// true：无序签
	// false：有序签
	// 注：默认为false（有序签）
	Unordered *bool `json:"Unordered,omitempty" name:"Unordered"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitempty" name:"CustomShowMap"`

	// 发起方企业的签署人进行签署操作是否需要企业内部审批。使用此功能需要发起方企业有参与签署。
	// 若设置为true，审核结果需通过接口 CreateFlowSignReview 通知电子签，审核通过后，发起方企业签署人方可进行签署操作，否则会阻塞其签署操作。
	// 
	// 注：企业可以通过此功能与企业内部的审批流程进行关联，支持手动、静默签署合同。
	NeedSignReview *bool `json:"NeedSignReview,omitempty" name:"NeedSignReview"`

	// 应用号信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

func (r *CreateFlowByFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowByFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowName")
	delete(f, "Approvers")
	delete(f, "FileIds")
	delete(f, "FlowType")
	delete(f, "Components")
	delete(f, "CcInfos")
	delete(f, "NeedPreview")
	delete(f, "FlowDescription")
	delete(f, "Deadline")
	delete(f, "Unordered")
	delete(f, "CustomShowMap")
	delete(f, "NeedSignReview")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowByFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowByFilesResponseParams struct {
	// 签署流程编号。
	// 
	// 注：如入参 是否需要预览 NeedPreview 设置为 true，不会正式发起合同，此处不会有值返回；如入参 是否需要预览 NeedPreview 设置为 false，此处会正常返回签署流程编号 FlowId。
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 合同预览链接。
	// 
	// 注：如入参 是否需要预览 NeedPreview 设置为 true，会开启“预览模式”，此处会返回预览链接；如入参 是否需要预览 NeedPreview 设置为 false，此处不会有值返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewUrl *string `json:"PreviewUrl,omitempty" name:"PreviewUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFlowByFilesResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowByFilesResponseParams `json:"Response"`
}

func (r *CreateFlowByFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowByFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowEvidenceReportRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type CreateFlowEvidenceReportRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *CreateFlowEvidenceReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowEvidenceReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowEvidenceReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowEvidenceReportResponseParams struct {
	// 出证报告 URL（有效期5分钟）
	ReportUrl *string `json:"ReportUrl,omitempty" name:"ReportUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFlowEvidenceReportResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowEvidenceReportResponseParams `json:"Response"`
}

func (r *CreateFlowEvidenceReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowEvidenceReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程名称,最大长度200个字符
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署流程参与者信息，最大限制50方
	Approvers []*FlowCreateApprover `json:"Approvers,omitempty" name:"Approvers"`

	// 签署流程的类型(如销售合同/入职合同等)，最大长度200个字符
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 暂未开放
	RelatedFlowId *string `json:"RelatedFlowId,omitempty" name:"RelatedFlowId"`

	// 签署流程的签署截止时间。
	// 值为unix时间戳,精确到秒,不传默认为当前时间一年后
	DeadLine *int64 `json:"DeadLine,omitempty" name:"DeadLine"`

	// 用户自定义字段(需进行base64 encode),回调的时候会进行透传, 长度需要小于20480
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// 签署流程描述,最大长度1000个字符
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 发送类型：
	// true：无序签
	// false：有序签
	// 注：默认为false（有序签），请和模板中的配置保持一致
	Unordered *bool `json:"Unordered,omitempty" name:"Unordered"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitempty" name:"CustomShowMap"`

	// 发起方企业的签署人进行签署操作是否需要企业内部审批。使用此功能需要发起方企业有参与签署。
	// 若设置为true，审核结果需通过接口 CreateFlowSignReview 通知电子签，审核通过后，发起方企业签署人方可进行签署操作，否则会阻塞其签署操作。
	// 
	// 注：企业可以通过此功能与企业内部的审批流程进行关联，支持手动、静默签署合同。
	NeedSignReview *bool `json:"NeedSignReview,omitempty" name:"NeedSignReview"`

	// 暂未开放
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

type CreateFlowRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程名称,最大长度200个字符
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署流程参与者信息，最大限制50方
	Approvers []*FlowCreateApprover `json:"Approvers,omitempty" name:"Approvers"`

	// 签署流程的类型(如销售合同/入职合同等)，最大长度200个字符
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 暂未开放
	RelatedFlowId *string `json:"RelatedFlowId,omitempty" name:"RelatedFlowId"`

	// 签署流程的签署截止时间。
	// 值为unix时间戳,精确到秒,不传默认为当前时间一年后
	DeadLine *int64 `json:"DeadLine,omitempty" name:"DeadLine"`

	// 用户自定义字段(需进行base64 encode),回调的时候会进行透传, 长度需要小于20480
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// 签署流程描述,最大长度1000个字符
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 发送类型：
	// true：无序签
	// false：有序签
	// 注：默认为false（有序签），请和模板中的配置保持一致
	Unordered *bool `json:"Unordered,omitempty" name:"Unordered"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitempty" name:"CustomShowMap"`

	// 发起方企业的签署人进行签署操作是否需要企业内部审批。使用此功能需要发起方企业有参与签署。
	// 若设置为true，审核结果需通过接口 CreateFlowSignReview 通知电子签，审核通过后，发起方企业签署人方可进行签署操作，否则会阻塞其签署操作。
	// 
	// 注：企业可以通过此功能与企业内部的审批流程进行关联，支持手动、静默签署合同。
	NeedSignReview *bool `json:"NeedSignReview,omitempty" name:"NeedSignReview"`

	// 暂未开放
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

func (r *CreateFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowName")
	delete(f, "Approvers")
	delete(f, "FlowType")
	delete(f, "ClientToken")
	delete(f, "RelatedFlowId")
	delete(f, "DeadLine")
	delete(f, "UserData")
	delete(f, "FlowDescription")
	delete(f, "Unordered")
	delete(f, "CustomShowMap")
	delete(f, "NeedSignReview")
	delete(f, "CallbackUrl")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowResponseParams struct {
	// 签署流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowResponseParams `json:"Response"`
}

func (r *CreateFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowSignReviewRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 企业内部审核结果
	// PASS: 通过 
	// REJECT: 拒绝
	ReviewType *string `json:"ReviewType,omitempty" name:"ReviewType"`

	// 审核原因 
	// 当ReviewType 是REJECT 时此字段必填,字符串长度不超过200
	ReviewMessage *string `json:"ReviewMessage,omitempty" name:"ReviewMessage"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

type CreateFlowSignReviewRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 企业内部审核结果
	// PASS: 通过 
	// REJECT: 拒绝
	ReviewType *string `json:"ReviewType,omitempty" name:"ReviewType"`

	// 审核原因 
	// 当ReviewType 是REJECT 时此字段必填,字符串长度不超过200
	ReviewMessage *string `json:"ReviewMessage,omitempty" name:"ReviewMessage"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

func (r *CreateFlowSignReviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowSignReviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "ReviewType")
	delete(f, "ReviewMessage")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowSignReviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowSignReviewResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFlowSignReviewResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowSignReviewResponseParams `json:"Response"`
}

func (r *CreateFlowSignReviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowSignReviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiFlowSignQRCodeRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 签署流程名称，最大长度不超过200字符
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 用户信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 应用信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 回调地址,最大长度1000字符串
	// 回调时机：
	// 用户通过签署二维码发起签署流程时，企业额度不足导致失败
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 最大可发起签署流程份数，默认5份 
	// 发起流程数量超过此上限后二维码自动失效
	MaxFlowNum *int64 `json:"MaxFlowNum,omitempty" name:"MaxFlowNum"`

	// 签署流程有效天数 默认7天 最高设置不超过30天
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitempty" name:"FlowEffectiveDay"`

	// 二维码有效天数 默认7天 最高设置不超过90天
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitempty" name:"QrEffectiveDay"`

	// 限制二维码用户条件
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitempty" name:"ApproverRestrictions"`
}

type CreateMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 签署流程名称，最大长度不超过200字符
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 用户信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 应用信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 回调地址,最大长度1000字符串
	// 回调时机：
	// 用户通过签署二维码发起签署流程时，企业额度不足导致失败
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 最大可发起签署流程份数，默认5份 
	// 发起流程数量超过此上限后二维码自动失效
	MaxFlowNum *int64 `json:"MaxFlowNum,omitempty" name:"MaxFlowNum"`

	// 签署流程有效天数 默认7天 最高设置不超过30天
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitempty" name:"FlowEffectiveDay"`

	// 二维码有效天数 默认7天 最高设置不超过90天
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitempty" name:"QrEffectiveDay"`

	// 限制二维码用户条件
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitempty" name:"ApproverRestrictions"`
}

func (r *CreateMultiFlowSignQRCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiFlowSignQRCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "FlowName")
	delete(f, "Operator")
	delete(f, "Agent")
	delete(f, "CallbackUrl")
	delete(f, "MaxFlowNum")
	delete(f, "FlowEffectiveDay")
	delete(f, "QrEffectiveDay")
	delete(f, "ApproverRestrictions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMultiFlowSignQRCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiFlowSignQRCodeResponseParams struct {
	// 签署二维码对象
	QrCode *SignQrCode `json:"QrCode,omitempty" name:"QrCode"`

	// 签署链接对象
	SignUrls *SignUrl `json:"SignUrls,omitempty" name:"SignUrls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMultiFlowSignQRCodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateMultiFlowSignQRCodeResponseParams `json:"Response"`
}

func (r *CreateMultiFlowSignQRCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultiFlowSignQRCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchemeUrlRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 企业名称
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 姓名,最大长度50个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 手机号，大陆手机号11位
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 链接类型
	// HTTP：跳转电子签小程序的http_url，
	// APP：第三方APP或小程序跳转电子签小程序的path。
	// 默认为HTTP类型
	EndPoint *string `json:"EndPoint,omitempty" name:"EndPoint"`

	// 签署流程编号 (PathType=1时必传)
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 跳转页面 1: 小程序合同详情 2: 小程序合同列表页 0: 不传, 默认主页
	PathType *uint64 `json:"PathType,omitempty" name:"PathType"`

	// 是否自动回跳 true：是， false：否。该参数只针对"APP" 类型的签署链接有效
	AutoJumpBack *bool `json:"AutoJumpBack,omitempty" name:"AutoJumpBack"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

type CreateSchemeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 企业名称
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 姓名,最大长度50个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 手机号，大陆手机号11位
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 链接类型
	// HTTP：跳转电子签小程序的http_url，
	// APP：第三方APP或小程序跳转电子签小程序的path。
	// 默认为HTTP类型
	EndPoint *string `json:"EndPoint,omitempty" name:"EndPoint"`

	// 签署流程编号 (PathType=1时必传)
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 跳转页面 1: 小程序合同详情 2: 小程序合同列表页 0: 不传, 默认主页
	PathType *uint64 `json:"PathType,omitempty" name:"PathType"`

	// 是否自动回跳 true：是， false：否。该参数只针对"APP" 类型的签署链接有效
	AutoJumpBack *bool `json:"AutoJumpBack,omitempty" name:"AutoJumpBack"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

func (r *CreateSchemeUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchemeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "OrganizationName")
	delete(f, "Name")
	delete(f, "Mobile")
	delete(f, "EndPoint")
	delete(f, "FlowId")
	delete(f, "PathType")
	delete(f, "AutoJumpBack")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchemeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchemeUrlResponseParams struct {
	// 小程序链接地址
	SchemeUrl *string `json:"SchemeUrl,omitempty" name:"SchemeUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSchemeUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateSchemeUrlResponseParams `json:"Response"`
}

func (r *CreateSchemeUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchemeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileUrlsRequestParams struct {
	// 文件对应的业务类型，目前支持：
	// - 模板 "TEMPLATE"
	// - 文档 "DOCUMENT"
	// - 印章  “SEAL”
	// - 流程 "FLOW"
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 业务编号的数组，如模板编号、文档编号、印章编号
	// 最大支持20个资源
	BusinessIds []*string `json:"BusinessIds,omitempty" name:"BusinessIds"`

	// 操作者信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 下载后的文件命名，只有fileType为zip的时候生效
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件类型，"JPG", "PDF","ZIP"等
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 指定资源起始偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 指定资源数量，查询全部资源则传入-1
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 下载url过期时间，单位秒。0: 按默认值5分钟，允许范围：1s~24x60x60s(1天)
	UrlTtl *int64 `json:"UrlTtl,omitempty" name:"UrlTtl"`

	// 暂不开放
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 暂不开放
	CcToken *string `json:"CcToken,omitempty" name:"CcToken"`
}

type DescribeFileUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 文件对应的业务类型，目前支持：
	// - 模板 "TEMPLATE"
	// - 文档 "DOCUMENT"
	// - 印章  “SEAL”
	// - 流程 "FLOW"
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 业务编号的数组，如模板编号、文档编号、印章编号
	// 最大支持20个资源
	BusinessIds []*string `json:"BusinessIds,omitempty" name:"BusinessIds"`

	// 操作者信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 下载后的文件命名，只有fileType为zip的时候生效
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件类型，"JPG", "PDF","ZIP"等
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 指定资源起始偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 指定资源数量，查询全部资源则传入-1
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 下载url过期时间，单位秒。0: 按默认值5分钟，允许范围：1s~24x60x60s(1天)
	UrlTtl *int64 `json:"UrlTtl,omitempty" name:"UrlTtl"`

	// 暂不开放
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 暂不开放
	CcToken *string `json:"CcToken,omitempty" name:"CcToken"`
}

func (r *DescribeFileUrlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileUrlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessType")
	delete(f, "BusinessIds")
	delete(f, "Operator")
	delete(f, "Agent")
	delete(f, "FileName")
	delete(f, "FileType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "UrlTtl")
	delete(f, "Scene")
	delete(f, "CcToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileUrlsResponseParams struct {
	// URL信息
	FileUrls []*FileUrl `json:"FileUrls,omitempty" name:"FileUrls"`

	// URL数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFileUrlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileUrlsResponseParams `json:"Response"`
}

func (r *DescribeFileUrlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileUrlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowBriefsRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 需要查询的流程ID列表，限制最大20个
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

type DescribeFlowBriefsRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 需要查询的流程ID列表，限制最大20个
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

func (r *DescribeFlowBriefsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowBriefsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowBriefsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowBriefsResponseParams struct {
	// 流程列表
	FlowBriefs []*FlowBrief `json:"FlowBriefs,omitempty" name:"FlowBriefs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowBriefsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowBriefsResponseParams `json:"Response"`
}

func (r *DescribeFlowBriefsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowBriefsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowInfoRequestParams struct {
	// 需要查询的流程ID列表，限制最大100个
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 调用方用户信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type DescribeFlowInfoRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的流程ID列表，限制最大100个
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 调用方用户信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeFlowInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowIds")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowInfoResponseParams struct {
	// 签署流程信息
	FlowDetailInfos []*FlowDetailInfo `json:"FlowDetailInfos,omitempty" name:"FlowDetailInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowInfoResponseParams `json:"Response"`
}

func (r *DescribeFlowInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowTemplatesRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 搜索条件，具体参考Filter结构体。本接口取值：template-id：按照【 **模板唯一标识** 】进行过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询个数，默认20，最大200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移位置，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询内容：0-模板列表及详情（默认），1-仅模板列表
	ContentType *int64 `json:"ContentType,omitempty" name:"ContentType"`

	// 暂未开放
	GenerateSource *uint64 `json:"GenerateSource,omitempty" name:"GenerateSource"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

type DescribeFlowTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 搜索条件，具体参考Filter结构体。本接口取值：template-id：按照【 **模板唯一标识** 】进行过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询个数，默认20，最大200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移位置，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询内容：0-模板列表及详情（默认），1-仅模板列表
	ContentType *int64 `json:"ContentType,omitempty" name:"ContentType"`

	// 暂未开放
	GenerateSource *uint64 `json:"GenerateSource,omitempty" name:"GenerateSource"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

func (r *DescribeFlowTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ContentType")
	delete(f, "GenerateSource")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowTemplatesResponseParams struct {
	// 模板详情列表
	Templates []*TemplateInfo `json:"Templates,omitempty" name:"Templates"`

	// 查询到的总个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowTemplatesResponseParams `json:"Response"`
}

func (r *DescribeFlowTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeThirdPartyAuthCodeRequestParams struct {
	// 电子签小程序跳转客户小程序时携带的授权查看码
	AuthCode *string `json:"AuthCode,omitempty" name:"AuthCode"`
}

type DescribeThirdPartyAuthCodeRequest struct {
	*tchttp.BaseRequest
	
	// 电子签小程序跳转客户小程序时携带的授权查看码
	AuthCode *string `json:"AuthCode,omitempty" name:"AuthCode"`
}

func (r *DescribeThirdPartyAuthCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeThirdPartyAuthCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuthCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeThirdPartyAuthCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeThirdPartyAuthCodeResponseParams struct {
	// 用户是否实名，VERIFIED 为实名，UNVERIFIED 未实名
	VerifyStatus *string `json:"VerifyStatus,omitempty" name:"VerifyStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeThirdPartyAuthCodeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeThirdPartyAuthCodeResponseParams `json:"Response"`
}

func (r *DescribeThirdPartyAuthCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeThirdPartyAuthCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileInfo struct {
	// 文件Id
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件大小，单位为Byte
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件上传时间，10位时间戳（精确到秒）
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`
}

type FileUrl struct {
	// 下载文件的URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// 下载文件的附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Option *string `json:"Option,omitempty" name:"Option"`
}

type FillApproverInfo struct {
	// 签署人签署Id
	RecipientId *string `json:"RecipientId,omitempty" name:"RecipientId"`

	// 签署人来源
	// WEWORKAPP: 企业微信
	ApproverSource *string `json:"ApproverSource,omitempty" name:"ApproverSource"`

	// 企业自定义账号Id
	// WEWORKAPP场景下指企业自有应用获取企微明文的userid
	CustomUserId *string `json:"CustomUserId,omitempty" name:"CustomUserId"`
}

type Filter struct {
	// 查询过滤条件的Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 查询过滤条件的Value列表
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FlowApproverDetail struct {
	// 签署人信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveMessage *string `json:"ApproveMessage,omitempty" name:"ApproveMessage"`

	// 签署人名字
	ApproveName *string `json:"ApproveName,omitempty" name:"ApproveName"`

	// 签署人的状态
	// 0：还没有发起
	// 1：流程中 没有开始处理
	// 2：待处理
	// 3：签署态
	// 4：拒绝态
	// 5：过期没人处理
	// 6：取消态
	// 7：还没有预发起
	// 8：待填写
	// 9：因为各种原因而终止
	ApproveStatus *int64 `json:"ApproveStatus,omitempty" name:"ApproveStatus"`

	// 模板配置时候的签署人id,与控件绑定
	ReceiptId *string `json:"ReceiptId,omitempty" name:"ReceiptId"`

	// 客户自定义userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomUserId *string `json:"CustomUserId,omitempty" name:"CustomUserId"`

	// 签署人手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 签署顺序
	SignOrder *int64 `json:"SignOrder,omitempty" name:"SignOrder"`

	// 签署人签署时间
	ApproveTime *int64 `json:"ApproveTime,omitempty" name:"ApproveTime"`

	// 参与者类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveType *string `json:"ApproveType,omitempty" name:"ApproveType"`

	// 签署人侧用户来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverSource *string `json:"ApproverSource,omitempty" name:"ApproverSource"`

	// 客户自定义签署人标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomApproverTag *string `json:"CustomApproverTag,omitempty" name:"CustomApproverTag"`

	// 签署人企业Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"OrganizationId,omitempty" name:"OrganizationId"`

	// 签署人企业名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`
}

type FlowBrief struct {
	// 流程的编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 流程的名称
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 流程的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 流程的类型
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 流程状态
	// - `0`  还没有发起
	// - `1`  未签署
	// - `2`  部分签署
	// - `3`  已退回
	// - `4`  完成签署
	// - `5`  已过期
	// - `6`  已取消
	// - `7`  还没有预发起
	// - `8`  等待填写
	// - `9`  部分填写
	// - `10`  拒填
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowStatus *int64 `json:"FlowStatus,omitempty" name:"FlowStatus"`

	// 流程创建的时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 拒签或者取消的原因描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMessage *string `json:"FlowMessage,omitempty" name:"FlowMessage"`
}

type FlowCreateApprover struct {
	// 参与者类型：
	// 0：企业
	// 1：个人
	// 3：企业静默签署
	// 注：类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。静默签署仅进行盖章操作，不能自动签名。
	ApproverType *int64 `json:"ApproverType,omitempty" name:"ApproverType"`

	// 如果签署方为企业，需要填入企业全称
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 签署方经办人姓名
	ApproverName *string `json:"ApproverName,omitempty" name:"ApproverName"`

	// 签署方经办人手机号码
	ApproverMobile *string `json:"ApproverMobile,omitempty" name:"ApproverMobile"`

	// 签署方经办人证件类型ID_CARD 身份证
	// HONGKONG_AND_MACAO 港澳居民来往内地通行证
	// HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证(格式同居民身份证)
	ApproverIdCardType *string `json:"ApproverIdCardType,omitempty" name:"ApproverIdCardType"`

	// 签署方经办人证件号码
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitempty" name:"ApproverIdCardNumber"`

	// 签署方经办人在模板中的角色ID
	RecipientId *string `json:"RecipientId,omitempty" name:"RecipientId"`

	// 签署意愿确认渠道,WEIXINAPP:人脸识别
	VerifyChannel []*string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 是否发送短信，sms--短信通知，none--不通知，默认为sms；发起方=签署方时不发送短信
	NotifyType *string `json:"NotifyType,omitempty" name:"NotifyType"`

	// 签署前置条件：是否需要阅读全文，默认为不需要
	IsFullText *bool `json:"IsFullText,omitempty" name:"IsFullText"`

	// 签署前置条件：阅读时长限制，默认为不需要
	PreReadTime *uint64 `json:"PreReadTime,omitempty" name:"PreReadTime"`

	// 签署方经办人的用户ID,和签署方经办人姓名+手机号+证件必须有一个。非企微场景不使用此字段
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 当前只支持true，默认为true
	Required *bool `json:"Required,omitempty" name:"Required"`

	// 签署人用户来源,企微侧用户请传入：WEWORKAPP
	ApproverSource *string `json:"ApproverSource,omitempty" name:"ApproverSource"`

	// 客户自定义签署人标识，64位长度，保证唯一。非企微场景不使用此字段
	CustomApproverTag *string `json:"CustomApproverTag,omitempty" name:"CustomApproverTag"`

	// 快速注册相关信息，目前暂未开放！
	RegisterInfo *RegisterInfo `json:"RegisterInfo,omitempty" name:"RegisterInfo"`
}

type FlowDetailInfo struct {
	// 合同(流程)的Id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 合同(流程)的名字
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 合同(流程)的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 合同(流程)的状态
	// 1：未签署
	// 2：部分签署
	// 3：已退回
	// 4：完成签署
	// 5：已过期
	// 6：已取消
	FlowStatus *int64 `json:"FlowStatus,omitempty" name:"FlowStatus"`

	// 合同(流程)的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMessage *string `json:"FlowMessage,omitempty" name:"FlowMessage"`

	// 流程的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 合同(流程)的创建时间戳
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 合同(流程)的签署人数组
	FlowApproverInfos []*FlowApproverDetail `json:"FlowApproverInfos,omitempty" name:"FlowApproverInfos"`
}

type FormField struct {
	// 控件填充vaule，ComponentType和传入值类型对应关系：
	// TEXT - 文本内容
	// MULTI_LINE_TEXT - 文本内容
	// CHECK_BOX - true/false
	// FILL_IMAGE、ATTACHMENT - 附件的FileId，需要通过UploadFiles接口上传获取
	// SELECTOR - 选项值
	// DYNAMIC_TABLE - 传入json格式的表格内容，具体见数据结构FlowInfo：https://cloud.tencent.com/document/api/1420/61525#FlowInfo
	ComponentValue *string `json:"ComponentValue,omitempty" name:"ComponentValue"`

	// 控件id，和ComponentName选择一项传入即可
	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`

	// 控件名字，最大长度不超过30字符，和ComponentId选择一项传入即可
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

// Predefined struct for user
type GetTaskResultApiRequestParams struct {
	// 任务Id，通过CreateConvertTaskApi得到
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 应用号信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`
}

type GetTaskResultApiRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id，通过CreateConvertTaskApi得到
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 应用号信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`
}

func (r *GetTaskResultApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskResultApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Operator")
	delete(f, "Agent")
	delete(f, "Organization")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskResultApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskResultApiResponseParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态，需要关注的状态
	// 0  :NeedTranform   - 任务已提交
	// 4  :Processing     - 文档转换中
	// 8  :TaskEnd        - 任务处理完成
	// -2 :DownloadFailed - 下载失败
	// -6 :ProcessFailed  - 转换失败
	// -13:ProcessTimeout - 转换文件超时
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 状态描述，需要关注的状态
	// NeedTranform   - 任务已提交
	// Processing     - 文档转换中
	// TaskEnd        - 任务处理完成
	// DownloadFailed - 下载失败
	// ProcessFailed  - 转换失败
	// ProcessTimeout - 转换文件超时
	TaskMessage *string `json:"TaskMessage,omitempty" name:"TaskMessage"`

	// 资源Id，也是FileId，用于文件发起使用
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTaskResultApiResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskResultApiResponseParams `json:"Response"`
}

func (r *GetTaskResultApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskResultApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrganizationInfo struct {
	// 机构在平台的编号
	OrganizationId *string `json:"OrganizationId,omitempty" name:"OrganizationId"`

	// 用户渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 用户在渠道的机构编号
	OrganizationOpenId *string `json:"OrganizationOpenId,omitempty" name:"OrganizationOpenId"`

	// 用户真实的IP
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 机构的代理IP
	ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
}

type PdfVerifyResult struct {
	// 验签结果
	VerifyResult *int64 `json:"VerifyResult,omitempty" name:"VerifyResult"`

	// 签署平台
	SignPlatform *string `json:"SignPlatform,omitempty" name:"SignPlatform"`

	// 签署人名称
	SignerName *string `json:"SignerName,omitempty" name:"SignerName"`

	// 签署时间
	SignTime *int64 `json:"SignTime,omitempty" name:"SignTime"`

	// 签名算法
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" name:"SignAlgorithm"`

	// 签名证书序列号
	CertSn *string `json:"CertSn,omitempty" name:"CertSn"`

	// 证书起始时间
	CertNotBefore *int64 `json:"CertNotBefore,omitempty" name:"CertNotBefore"`

	// 证书过期时间
	CertNotAfter *int64 `json:"CertNotAfter,omitempty" name:"CertNotAfter"`

	// 签名域横坐标
	ComponentPosX *float64 `json:"ComponentPosX,omitempty" name:"ComponentPosX"`

	// 签名域纵坐标
	ComponentPosY *float64 `json:"ComponentPosY,omitempty" name:"ComponentPosY"`

	// 签名域宽度
	ComponentWidth *float64 `json:"ComponentWidth,omitempty" name:"ComponentWidth"`

	// 签名域高度
	ComponentHeight *float64 `json:"ComponentHeight,omitempty" name:"ComponentHeight"`

	// 签名域所在页码
	ComponentPage *int64 `json:"ComponentPage,omitempty" name:"ComponentPage"`
}

type Recipient struct {
	// 签署参与者ID
	RecipientId *string `json:"RecipientId,omitempty" name:"RecipientId"`

	// 参与者类型（ENTERPRISE/INDIVIDUAL）
	RecipientType *string `json:"RecipientType,omitempty" name:"RecipientType"`

	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 角色名称
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 是否需要验证，默认为false
	RequireValidation *bool `json:"RequireValidation,omitempty" name:"RequireValidation"`

	// 是否需要签署，默认为true
	RequireSign *bool `json:"RequireSign,omitempty" name:"RequireSign"`

	// 添加序列
	RoutingOrder *int64 `json:"RoutingOrder,omitempty" name:"RoutingOrder"`

	// 是否需要发送，默认为true
	RequireDelivery *bool `json:"RequireDelivery,omitempty" name:"RequireDelivery"`

	// 邮箱地址
	Email *string `json:"Email,omitempty" name:"Email"`

	// 电话号码
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 关联的用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 发送方式（EMAIL/MOBILE）
	DeliveryMethod *string `json:"DeliveryMethod,omitempty" name:"DeliveryMethod"`

	// 附属信息
	RecipientExtra *string `json:"RecipientExtra,omitempty" name:"RecipientExtra"`
}

type RegisterInfo struct {
	// 法人姓名
	LegalName *string `json:"LegalName,omitempty" name:"LegalName"`

	// 社会统一信用代码
	Uscc *string `json:"Uscc,omitempty" name:"Uscc"`
}

type SignQrCode struct {
	// 二维码id
	QrCodeId *string `json:"QrCodeId,omitempty" name:"QrCodeId"`

	// 二维码url
	QrCodeUrl *string `json:"QrCodeUrl,omitempty" name:"QrCodeUrl"`

	// 二维码过期时间
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
}

type SignUrl struct {
	// 小程序签署链接
	AppSignUrl *string `json:"AppSignUrl,omitempty" name:"AppSignUrl"`

	// 签署链接有效时间
	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`

	// 移动端签署链接
	HttpSignUrl *string `json:"HttpSignUrl,omitempty" name:"HttpSignUrl"`
}

// Predefined struct for user
type StartFlowRequestParams struct {
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程编号，由CreateFlow接口返回
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

type StartFlowRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署流程编号，由CreateFlow接口返回
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

func (r *StartFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "ClientToken")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartFlowResponseParams struct {
	// 返回描述，START-发起成功， REVIEW-提交审核成功，EXECUTING-已提交发起任务
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartFlowResponse struct {
	*tchttp.BaseResponse
	Response *StartFlowResponseParams `json:"Response"`
}

func (r *StartFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateInfo struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名字
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模板描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 模板关联的资源IDs
	DocumentResourceIds []*string `json:"DocumentResourceIds,omitempty" name:"DocumentResourceIds"`

	// 返回的文件信息结构
	FileInfos []*FileInfo `json:"FileInfos,omitempty" name:"FileInfos"`

	// 附件关联的资源ID是
	AttachmentResourceIds []*string `json:"AttachmentResourceIds,omitempty" name:"AttachmentResourceIds"`

	// 签署顺序
	SignOrder []*int64 `json:"SignOrder,omitempty" name:"SignOrder"`

	// 签署参与者的信息
	Recipients []*Recipient `json:"Recipients,omitempty" name:"Recipients"`

	// 模板信息结构
	Components []*Component `json:"Components,omitempty" name:"Components"`

	// 签署区模板信息结构
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 模板状态(-1:不可用；0:草稿态；1:正式态)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 模板的创建人
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 模板创建的时间戳（精确到秒）
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 发起人角色信息
	Promoter *Recipient `json:"Promoter,omitempty" name:"Promoter"`
}

type UploadFile struct {
	// Base64编码后的文件内容
	FileBody *string `json:"FileBody,omitempty" name:"FileBody"`

	// 文件名，最大长度不超过200字符
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

// Predefined struct for user
type UploadFilesRequestParams struct {
	// 文件对应业务类型，用于区分文件存储路径：
	// 1. TEMPLATE - 模板； 文件类型：.pdf .doc .docx .html
	// 2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.jpg/.png
	// 3. SEAL - 印章； 文件类型：.jpg/.jpeg/.png
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 上传文件内容数组，最多支持20个文件
	FileInfos []*UploadFile `json:"FileInfos,omitempty" name:"FileInfos"`

	// 不再使用，上传文件链接数组，最多支持20个URL
	FileUrls *string `json:"FileUrls,omitempty" name:"FileUrls"`

	// 此参数只对 PDF 文件有效。是否将pdf灰色矩阵置白
	// true--是，处理置白
	// false--否，不处理
	CoverRect *bool `json:"CoverRect,omitempty" name:"CoverRect"`

	// 文件类型， 默认通过文件内容解析得到文件类型，客户可以显示的说明上传文件的类型。
	// 如：PDF 表示上传的文件 xxx.pdf的文件类型是 PDF
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 用户自定义ID数组，与上传文件一一对应
	CustomIds []*string `json:"CustomIds,omitempty" name:"CustomIds"`
}

type UploadFilesRequest struct {
	*tchttp.BaseRequest
	
	// 文件对应业务类型，用于区分文件存储路径：
	// 1. TEMPLATE - 模板； 文件类型：.pdf .doc .docx .html
	// 2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.jpg/.png
	// 3. SEAL - 印章； 文件类型：.jpg/.jpeg/.png
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 上传文件内容数组，最多支持20个文件
	FileInfos []*UploadFile `json:"FileInfos,omitempty" name:"FileInfos"`

	// 不再使用，上传文件链接数组，最多支持20个URL
	FileUrls *string `json:"FileUrls,omitempty" name:"FileUrls"`

	// 此参数只对 PDF 文件有效。是否将pdf灰色矩阵置白
	// true--是，处理置白
	// false--否，不处理
	CoverRect *bool `json:"CoverRect,omitempty" name:"CoverRect"`

	// 文件类型， 默认通过文件内容解析得到文件类型，客户可以显示的说明上传文件的类型。
	// 如：PDF 表示上传的文件 xxx.pdf的文件类型是 PDF
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 用户自定义ID数组，与上传文件一一对应
	CustomIds []*string `json:"CustomIds,omitempty" name:"CustomIds"`
}

func (r *UploadFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessType")
	delete(f, "Caller")
	delete(f, "FileInfos")
	delete(f, "FileUrls")
	delete(f, "CoverRect")
	delete(f, "FileType")
	delete(f, "CustomIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFilesResponseParams struct {
	// 文件id数组
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 上传成功文件数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadFilesResponse struct {
	*tchttp.BaseResponse
	Response *UploadFilesResponseParams `json:"Response"`
}

func (r *UploadFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {
	// 用户在平台的编号
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户的来源渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 用户在渠道的编号
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 用户真实IP
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 用户代理IP
	ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
}

// Predefined struct for user
type VerifyPdfRequestParams struct {
	// 合同Id，流程Id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type VerifyPdfRequest struct {
	*tchttp.BaseRequest
	
	// 合同Id，流程Id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 调用方用户信息，userId 必填
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *VerifyPdfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPdfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyPdfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyPdfResponseParams struct {
	// 验签结果，1-文件未被篡改，全部签名在腾讯电子签完成； 2-文件未被篡改，部分签名在腾讯电子签完成；3-文件被篡改；4-异常：文件内没有签名域；5-异常：文件签名格式错误
	VerifyResult *int64 `json:"VerifyResult,omitempty" name:"VerifyResult"`

	// 验签结果详情,内部状态1-验签成功，在电子签签署；2-验签成功，在其他平台签署；3-验签失败；4-pdf文件没有签名域
	// ；5-文件签名格式错误
	PdfVerifyResults []*PdfVerifyResult `json:"PdfVerifyResults,omitempty" name:"PdfVerifyResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyPdfResponse struct {
	*tchttp.BaseResponse
	Response *VerifyPdfResponseParams `json:"Response"`
}

func (r *VerifyPdfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPdfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}