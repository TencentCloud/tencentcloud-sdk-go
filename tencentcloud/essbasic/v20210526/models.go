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

package v20210526

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Agent struct {
	// 腾讯电子签颁发给渠道的应用ID，32位字符串
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 渠道/平台合作企业的企业ID
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitempty" name:"ProxyOrganizationOpenId"`

	// 渠道/平台合作企业经办人（操作员）
	ProxyOperator *UserInfo `json:"ProxyOperator,omitempty" name:"ProxyOperator"`

	// 腾讯电子签颁发给渠道侧合作企业的应用ID
	ProxyAppId *string `json:"ProxyAppId,omitempty" name:"ProxyAppId"`

	// 腾讯电子签颁发给渠道侧合作企业的企业ID
	ProxyOrganizationId *string `json:"ProxyOrganizationId,omitempty" name:"ProxyOrganizationId"`
}

type AuthFailMessage struct {
	// 合作企业Id
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitempty" name:"ProxyOrganizationOpenId"`

	// 出错信息
	Message *string `json:"Message,omitempty" name:"Message"`
}

type CcInfo struct {
	// 被抄送人手机号，大陆11位手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
}

// Predefined struct for user
type ChannelCancelMultiFlowSignQRCodeRequestParams struct {
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 二维码id
	QrCodeId *string `json:"QrCodeId,omitempty" name:"QrCodeId"`

	// 用户信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelCancelMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 二维码id
	QrCodeId *string `json:"QrCodeId,omitempty" name:"QrCodeId"`

	// 用户信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *ChannelCancelMultiFlowSignQRCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCancelMultiFlowSignQRCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "QrCodeId")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCancelMultiFlowSignQRCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCancelMultiFlowSignQRCodeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChannelCancelMultiFlowSignQRCodeResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCancelMultiFlowSignQRCodeResponseParams `json:"Response"`
}

func (r *ChannelCancelMultiFlowSignQRCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCancelMultiFlowSignQRCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowByFilesRequestParams struct {
	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程名称，长度不超过200个字符
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署流程签约方列表，最多不超过5个参与方
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitempty" name:"FlowApprovers"`

	// 签署文件资源Id列表，目前仅支持单个文件
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 签署文件中的控件，如：填写控件等
	Components []*Component `json:"Components,omitempty" name:"Components"`

	// 签署流程截止时间，十位数时间戳，最大值为33162419560，即3020年
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 签署流程回调地址，长度不超过255个字符
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 合同签署顺序类型(无序签,顺序签)，默认为false，即有序签署
	Unordered *bool `json:"Unordered,omitempty" name:"Unordered"`

	// 签署流程的类型，长度不超过255个字符
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 签署流程的描述，长度不超过1000个字符
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitempty" name:"CustomShowMap"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelCreateFlowByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程名称，长度不超过200个字符
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署流程签约方列表，最多不超过5个参与方
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitempty" name:"FlowApprovers"`

	// 签署文件资源Id列表，目前仅支持单个文件
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 签署文件中的控件，如：填写控件等
	Components []*Component `json:"Components,omitempty" name:"Components"`

	// 签署流程截止时间，十位数时间戳，最大值为33162419560，即3020年
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 签署流程回调地址，长度不超过255个字符
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 合同签署顺序类型(无序签,顺序签)，默认为false，即有序签署
	Unordered *bool `json:"Unordered,omitempty" name:"Unordered"`

	// 签署流程的类型，长度不超过255个字符
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 签署流程的描述，长度不超过1000个字符
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitempty" name:"CustomShowMap"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *ChannelCreateFlowByFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowByFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowName")
	delete(f, "FlowApprovers")
	delete(f, "FileIds")
	delete(f, "Components")
	delete(f, "Deadline")
	delete(f, "CallbackUrl")
	delete(f, "Unordered")
	delete(f, "FlowType")
	delete(f, "FlowDescription")
	delete(f, "CustomShowMap")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowByFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowByFilesResponseParams struct {
	// 合同签署流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChannelCreateFlowByFilesResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateFlowByFilesResponseParams `json:"Response"`
}

func (r *ChannelCreateFlowByFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowByFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateMultiFlowSignQRCodeRequestParams struct {
	// 渠道应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 模版ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 签署流程名称，最大长度200个字符。
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 最大可发起签署流程份数，默认5份；发起签署流程数量超过此上限后，二维码自动失效。
	MaxFlowNum *int64 `json:"MaxFlowNum,omitempty" name:"MaxFlowNum"`

	// 签署流程有效天数 默认7天 最高设置不超过30天
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitempty" name:"FlowEffectiveDay"`

	// 二维码有效天数 默认7天 最高设置不超过90天
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitempty" name:"QrEffectiveDay"`

	// 回调地址，最大长度1000个字符
	// 不传默认使用渠道应用号配置的回调地址
	// 回调时机:用户通过签署二维码发起合同时，企业额度不足导致失败
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 用户信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelCreateMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 模版ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 签署流程名称，最大长度200个字符。
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 最大可发起签署流程份数，默认5份；发起签署流程数量超过此上限后，二维码自动失效。
	MaxFlowNum *int64 `json:"MaxFlowNum,omitempty" name:"MaxFlowNum"`

	// 签署流程有效天数 默认7天 最高设置不超过30天
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitempty" name:"FlowEffectiveDay"`

	// 二维码有效天数 默认7天 最高设置不超过90天
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitempty" name:"QrEffectiveDay"`

	// 回调地址，最大长度1000个字符
	// 不传默认使用渠道应用号配置的回调地址
	// 回调时机:用户通过签署二维码发起合同时，企业额度不足导致失败
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 用户信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *ChannelCreateMultiFlowSignQRCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateMultiFlowSignQRCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "TemplateId")
	delete(f, "FlowName")
	delete(f, "MaxFlowNum")
	delete(f, "FlowEffectiveDay")
	delete(f, "QrEffectiveDay")
	delete(f, "CallbackUrl")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateMultiFlowSignQRCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateMultiFlowSignQRCodeResponseParams struct {
	// 签署二维码对象
	QrCode *SignQrCode `json:"QrCode,omitempty" name:"QrCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChannelCreateMultiFlowSignQRCodeResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateMultiFlowSignQRCodeResponseParams `json:"Response"`
}

func (r *ChannelCreateMultiFlowSignQRCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateMultiFlowSignQRCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Component struct {
	// 控件编号
	// 
	// 注：
	// 当GenerateMode=3时，通过"^"来决定是否使用关键字整词匹配能力。
	// 例：
	// 当GenerateMode=3时，如果传入关键字"^甲方签署^"，则会在PDF文件中有且仅有"甲方签署"关键字的地方进行对应操作。
	// 如传入的关键字为"甲方签署"，则PDF文件中每个出现关键字的位置都会执行相应操作。
	// 
	// 创建控件时，此值为空
	// 查询时返回完整结构
	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`

	// 如果是Component控件类型，则可选的字段为：
	// TEXT - 普通文本控件；
	// DATE - 普通日期控件；跟TEXT相比会有校验逻辑
	// 如果是SignComponent控件类型，则可选的字段为
	// SIGN_SEAL - 签署印章控件；
	// SIGN_DATE - 签署日期控件；
	// SIGN_SIGNATURE - 用户签名控件；
	// SIGN_PERSONAL_SEAL - 个人签署印章控件；
	// 
	// 表单域的控件不能作为印章和签名控件
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// 控件简称，不能超过30个字符
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// 定义控件是否为必填项，默认为false
	ComponentRequired *bool `json:"ComponentRequired,omitempty" name:"ComponentRequired"`

	// 控件所属文件的序号 (文档中文件的排列序号，从0开始)
	FileIndex *int64 `json:"FileIndex,omitempty" name:"FileIndex"`

	// 控件生成的方式：
	// NORMAL - 普通控件
	// FIELD - 表单域
	// KEYWORD - 关键字
	GenerateMode *string `json:"GenerateMode,omitempty" name:"GenerateMode"`

	// 参数控件宽度，默认100，单位px
	// 表单域和关键字转换控件不用填
	ComponentWidth *float64 `json:"ComponentWidth,omitempty" name:"ComponentWidth"`

	// 参数控件高度，默认100，单位px
	// 表单域和关键字转换控件不用填
	ComponentHeight *float64 `json:"ComponentHeight,omitempty" name:"ComponentHeight"`

	// 参数控件所在页码，从1开始
	ComponentPage *int64 `json:"ComponentPage,omitempty" name:"ComponentPage"`

	// 参数控件X位置，单位px
	ComponentPosX *float64 `json:"ComponentPosX,omitempty" name:"ComponentPosX"`

	// 参数控件Y位置，单位px
	ComponentPosY *float64 `json:"ComponentPosY,omitempty" name:"ComponentPosY"`

	// 参数控件样式，json格式表述
	// 不同类型的控件会有部分非通用参数
	// TEXT控件可以指定字体
	// 例如：{"FontSize":12}
	ComponentExtra *string `json:"ComponentExtra,omitempty" name:"ComponentExtra"`

	// 印章 ID，传参 DEFAULT_COMPANY_SEAL 表示使用默认印章。
	// 控件填入内容，印章控件里面，如果是手写签名内容为PNG图片格式的base64编码
	ComponentValue *string `json:"ComponentValue,omitempty" name:"ComponentValue"`

	// 日期签署控件的字号，默认为 12
	// 
	// 签署区日期控件会转换成图片格式并带存证，需要通过字体决定图片大小
	ComponentDateFontSize *int64 `json:"ComponentDateFontSize,omitempty" name:"ComponentDateFontSize"`

	// 控件所属文档的Id, 模块相关接口为空值
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`

	// 控件描述，不能超过30个字符
	ComponentDescription *string `json:"ComponentDescription,omitempty" name:"ComponentDescription"`

	// 指定关键字时横坐标偏移量，单位pt
	OffsetX *float64 `json:"OffsetX,omitempty" name:"OffsetX"`

	// 指定关键字时纵坐标偏移量，单位pt
	OffsetY *float64 `json:"OffsetY,omitempty" name:"OffsetY"`
}

// Predefined struct for user
type CreateConsoleLoginUrlRequestParams struct {
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 渠道侧合作企业名称，最大长度64个字符
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitempty" name:"ProxyOrganizationName"`

	// 渠道侧合作企业经办人的姓名，最大长度50个字符
	ProxyOperatorName *string `json:"ProxyOperatorName,omitempty" name:"ProxyOperatorName"`

	// 控制台指定模块，文件/合同管理:"DOCUMENT"，模板管理:"TEMPLATE"，印章管理:"SEAL"，组织架构/人员:"OPERATOR"，空字符串："账号信息"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 控制台指定模块Id
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 渠道侧合作企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`

	// 是否展示左侧菜单栏 是：ENABLE（默认） 否：DISABLE
	MenuStatus *string `json:"MenuStatus,omitempty" name:"MenuStatus"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type CreateConsoleLoginUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 渠道侧合作企业名称，最大长度64个字符
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitempty" name:"ProxyOrganizationName"`

	// 渠道侧合作企业经办人的姓名，最大长度50个字符
	ProxyOperatorName *string `json:"ProxyOperatorName,omitempty" name:"ProxyOperatorName"`

	// 控制台指定模块，文件/合同管理:"DOCUMENT"，模板管理:"TEMPLATE"，印章管理:"SEAL"，组织架构/人员:"OPERATOR"，空字符串："账号信息"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 控制台指定模块Id
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 渠道侧合作企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`

	// 是否展示左侧菜单栏 是：ENABLE（默认） 否：DISABLE
	MenuStatus *string `json:"MenuStatus,omitempty" name:"MenuStatus"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *CreateConsoleLoginUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsoleLoginUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "ProxyOrganizationName")
	delete(f, "ProxyOperatorName")
	delete(f, "Module")
	delete(f, "ModuleId")
	delete(f, "UniformSocialCreditCode")
	delete(f, "MenuStatus")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsoleLoginUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsoleLoginUrlResponseParams struct {
	// 控制台url，此链接5分钟内有效，且只能访问一次
	ConsoleUrl *string `json:"ConsoleUrl,omitempty" name:"ConsoleUrl"`

	// 渠道合作企业是否认证开通腾讯电子签。
	// 当渠道合作企业未完成认证开通腾讯电子签,建议先调用同步企业信息(SyncProxyOrganization)和同步经办人信息(SyncProxyOrganizationOperators)接口成功后再跳转到登录页面。
	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateConsoleLoginUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateConsoleLoginUrlResponseParams `json:"Response"`
}

func (r *CreateConsoleLoginUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsoleLoginUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowsByTemplatesRequestParams struct {
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 多个合同（签署流程）信息，最多支持20个
	FlowInfos []*FlowInfo `json:"FlowInfos,omitempty" name:"FlowInfos"`

	// 是否为预览模式；默认为false，即非预览模式，此时发起合同并返回FlowIds；若为预览模式，则返回PreviewUrls；
	// 预览链接有效期300秒；
	NeedPreview *bool `json:"NeedPreview,omitempty" name:"NeedPreview"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type CreateFlowsByTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 多个合同（签署流程）信息，最多支持20个
	FlowInfos []*FlowInfo `json:"FlowInfos,omitempty" name:"FlowInfos"`

	// 是否为预览模式；默认为false，即非预览模式，此时发起合同并返回FlowIds；若为预览模式，则返回PreviewUrls；
	// 预览链接有效期300秒；
	NeedPreview *bool `json:"NeedPreview,omitempty" name:"NeedPreview"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *CreateFlowsByTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowsByTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowInfos")
	delete(f, "NeedPreview")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowsByTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowsByTemplatesResponseParams struct {
	// 多个合同ID
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 渠道的业务信息，限制1024字符
	CustomerData []*string `json:"CustomerData,omitempty" name:"CustomerData"`

	// 创建消息，对应多个合同ID，
	// 成功为“”,创建失败则对应失败消息
	ErrorMessages []*string `json:"ErrorMessages,omitempty" name:"ErrorMessages"`

	// 预览模式下返回的预览文件url数组
	PreviewUrls []*string `json:"PreviewUrls,omitempty" name:"PreviewUrls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFlowsByTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowsByTemplatesResponseParams `json:"Response"`
}

func (r *CreateFlowsByTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowsByTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealByImageRequestParams struct {
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 印章名称，最大长度不超过30字符
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 印章图片base64
	SealImage *string `json:"SealImage,omitempty" name:"SealImage"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type CreateSealByImageRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 印章名称，最大长度不超过30字符
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 印章图片base64
	SealImage *string `json:"SealImage,omitempty" name:"SealImage"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *CreateSealByImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealByImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "SealName")
	delete(f, "SealImage")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSealByImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealByImageResponseParams struct {
	// 印章id
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSealByImageResponse struct {
	*tchttp.BaseResponse
	Response *CreateSealByImageResponseParams `json:"Response"`
}

func (r *CreateSealByImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealByImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignUrlsRequestParams struct {
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程编号数组，最多支持100个。
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 签署链接类型：“WEIXINAPP”-直接跳小程序；“CHANNEL”-跳转H5页面；“APP”-第三方APP或小程序跳转电子签小程序；默认“WEIXINAPP”类型，即跳转至小程序；
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 签署链接生成类型，默认是 "ALL"；
	// "ALL"：全部签署方签署链接；
	// "CHANNEL"：渠道合作企业；
	// "NOT_CHANNEL"：非渠道合作企业；
	// "PERSON"：个人；
	// "FOLLOWER"：关注方，目前是合同抄送方；
	GenerateType *string `json:"GenerateType,omitempty" name:"GenerateType"`

	// 非渠道合作企业参与方的企业名称，GenerateType为"NOT_CHANNEL"时必填
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 参与人姓名，GenerateType为"PERSON"时必填
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参与人手机号；
	// GenerateType为"PERSON"或"FOLLOWER"时必填
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 渠道合作企业的企业Id，GenerateType为"CHANNEL"时必填
	OrganizationOpenId *string `json:"OrganizationOpenId,omitempty" name:"OrganizationOpenId"`

	// 渠道合作企业参与人OpenId，GenerateType为"CHANNEL"时可用，指定到具体参与人
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// Endpoint为"APP" 类型的签署链接，可以设置此值；支持调用方小程序打开签署链接，在电子签小程序完成签署后自动回跳至调用方小程序
	AutoJumpBack *bool `json:"AutoJumpBack,omitempty" name:"AutoJumpBack"`

	// 签署完之后的H5页面的跳转链接，针对Endpoint为CHANNEL时有效，最大长度1000个字符。
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type CreateSignUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程编号数组，最多支持100个。
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 签署链接类型：“WEIXINAPP”-直接跳小程序；“CHANNEL”-跳转H5页面；“APP”-第三方APP或小程序跳转电子签小程序；默认“WEIXINAPP”类型，即跳转至小程序；
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 签署链接生成类型，默认是 "ALL"；
	// "ALL"：全部签署方签署链接；
	// "CHANNEL"：渠道合作企业；
	// "NOT_CHANNEL"：非渠道合作企业；
	// "PERSON"：个人；
	// "FOLLOWER"：关注方，目前是合同抄送方；
	GenerateType *string `json:"GenerateType,omitempty" name:"GenerateType"`

	// 非渠道合作企业参与方的企业名称，GenerateType为"NOT_CHANNEL"时必填
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 参与人姓名，GenerateType为"PERSON"时必填
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参与人手机号；
	// GenerateType为"PERSON"或"FOLLOWER"时必填
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 渠道合作企业的企业Id，GenerateType为"CHANNEL"时必填
	OrganizationOpenId *string `json:"OrganizationOpenId,omitempty" name:"OrganizationOpenId"`

	// 渠道合作企业参与人OpenId，GenerateType为"CHANNEL"时可用，指定到具体参与人
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// Endpoint为"APP" 类型的签署链接，可以设置此值；支持调用方小程序打开签署链接，在电子签小程序完成签署后自动回跳至调用方小程序
	AutoJumpBack *bool `json:"AutoJumpBack,omitempty" name:"AutoJumpBack"`

	// 签署完之后的H5页面的跳转链接，针对Endpoint为CHANNEL时有效，最大长度1000个字符。
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *CreateSignUrlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignUrlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowIds")
	delete(f, "Endpoint")
	delete(f, "GenerateType")
	delete(f, "OrganizationName")
	delete(f, "Name")
	delete(f, "Mobile")
	delete(f, "OrganizationOpenId")
	delete(f, "OpenId")
	delete(f, "AutoJumpBack")
	delete(f, "JumpUrl")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSignUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignUrlsResponseParams struct {
	// 签署参与者签署H5链接信息数组
	SignUrlInfos []*SignUrlInfo `json:"SignUrlInfos,omitempty" name:"SignUrlInfos"`

	// 生成失败时的错误信息，成功返回”“，顺序和出参SignUrlInfos保持一致
	ErrorMessages []*string `json:"ErrorMessages,omitempty" name:"ErrorMessages"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSignUrlsResponse struct {
	*tchttp.BaseResponse
	Response *CreateSignUrlsResponseParams `json:"Response"`
}

func (r *CreateSignUrlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignUrlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowDetailInfoRequestParams struct {
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 合同(流程)编号数组，最多支持100个。
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type DescribeFlowDetailInfoRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 合同(流程)编号数组，最多支持100个。
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeFlowDetailInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowDetailInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowIds")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowDetailInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowDetailInfoResponseParams struct {
	// 渠道侧应用号Id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 渠道侧企业第三方Id
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitempty" name:"ProxyOrganizationOpenId"`

	// 合同(签署流程)的具体详细描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowInfo []*FlowDetailInfo `json:"FlowInfo,omitempty" name:"FlowInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowDetailInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowDetailInfoResponseParams `json:"Response"`
}

func (r *DescribeFlowDetailInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowDetailInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUrlsByFlowsRequestParams struct {
	// 渠道应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 查询资源所对应的签署流程Id，最多支持50个。
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type DescribeResourceUrlsByFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 查询资源所对应的签署流程Id，最多支持50个。
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeResourceUrlsByFlowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceUrlsByFlowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowIds")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceUrlsByFlowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUrlsByFlowsResponseParams struct {
	// 签署流程资源对应链接信息
	FlowResourceUrlInfos []*FlowResourceUrlInfo `json:"FlowResourceUrlInfos,omitempty" name:"FlowResourceUrlInfos"`

	// 创建消息，对应多个合同ID，
	// 成功为“”,创建失败则对应失败消息
	ErrorMessages []*string `json:"ErrorMessages,omitempty" name:"ErrorMessages"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceUrlsByFlowsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceUrlsByFlowsResponseParams `json:"Response"`
}

func (r *DescribeResourceUrlsByFlowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceUrlsByFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplatesRequestParams struct {
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 模板唯一标识
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 查询内容：0-模板列表及详情（默认），1-仅模板列表
	ContentType *int64 `json:"ContentType,omitempty" name:"ContentType"`

	// 查询个数，默认20，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移位置，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type DescribeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 模板唯一标识
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 查询内容：0-模板列表及详情（默认），1-仅模板列表
	ContentType *int64 `json:"ContentType,omitempty" name:"ContentType"`

	// 查询个数，默认20，最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移位置，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "TemplateId")
	delete(f, "ContentType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplatesResponseParams struct {
	// 模板详情
	Templates []*TemplateInfo `json:"Templates,omitempty" name:"Templates"`

	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 查询数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询起始偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTemplatesResponseParams `json:"Response"`
}

func (r *DescribeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageRequestParams struct {
	// 应用信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 开始时间，例如：2021-03-21
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束时间，例如：2021-06-21；
	// 开始时间到结束时间的区间长度小于等于90天。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 是否汇总数据，默认不汇总。
	// 不汇总：返回在统计区间内渠道下所有企业的每日明细，即每个企业N条数据，N为统计天数；
	// 汇总：返回在统计区间内渠道下所有企业的汇总后数据，即每个企业一条数据；
	NeedAggregate *bool `json:"NeedAggregate,omitempty" name:"NeedAggregate"`

	// 单次返回的最多条目数量。默认为1000，且不能超过1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认是0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type DescribeUsageRequest struct {
	*tchttp.BaseRequest
	
	// 应用信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 开始时间，例如：2021-03-21
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束时间，例如：2021-06-21；
	// 开始时间到结束时间的区间长度小于等于90天。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 是否汇总数据，默认不汇总。
	// 不汇总：返回在统计区间内渠道下所有企业的每日明细，即每个企业N条数据，N为统计天数；
	// 汇总：返回在统计区间内渠道下所有企业的汇总后数据，即每个企业一条数据；
	NeedAggregate *bool `json:"NeedAggregate,omitempty" name:"NeedAggregate"`

	// 单次返回的最多条目数量。默认为1000，且不能超过1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认是0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "NeedAggregate")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageResponseParams struct {
	// 用量明细条数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 用量明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*UsageDetail `json:"Details,omitempty" name:"Details"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsageResponseParams `json:"Response"`
}

func (r *DescribeUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadFlowInfo struct {
	// 文件夹名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 签署流程的标识数组
	FlowIdList []*string `json:"FlowIdList,omitempty" name:"FlowIdList"`
}

type FlowApproverDetail struct {
	// 模板配置时候的签署人id,与控件绑定
	ReceiptId *string `json:"ReceiptId,omitempty" name:"ReceiptId"`

	// 渠道侧企业的第三方id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitempty" name:"ProxyOrganizationOpenId"`

	// 渠道侧企业操作人的第三方id
	ProxyOperatorOpenId *string `json:"ProxyOperatorOpenId,omitempty" name:"ProxyOperatorOpenId"`

	// 渠道侧企业名称
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitempty" name:"ProxyOrganizationName"`

	// 签署人手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 签署人签署顺序
	SignOrder *int64 `json:"SignOrder,omitempty" name:"SignOrder"`

	// 签署人姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveName *string `json:"ApproveName,omitempty" name:"ApproveName"`

	// 当前签署人的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveStatus *string `json:"ApproveStatus,omitempty" name:"ApproveStatus"`

	// 签署人信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveMessage *string `json:"ApproveMessage,omitempty" name:"ApproveMessage"`

	// 签署人签署时间
	ApproveTime *int64 `json:"ApproveTime,omitempty" name:"ApproveTime"`

	// 参与者类型 (ORGANIZATION企业/PERSON个人)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveType *string `json:"ApproveType,omitempty" name:"ApproveType"`
}

type FlowApproverInfo struct {
	// 签署人姓名，最大长度50个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 经办人身份证件类型
	// 1.ID_CARD 居民身份证
	// 2.HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证
	// 3.HONGKONG_AND_MACAO 港澳居民来往内地通行证
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 经办人证件号
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 签署人手机号，脱敏显示。大陆手机号为11位，暂不支持海外手机号。
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 企业签署方工商营业执照上的企业名称，签署方为非发起方企业场景下必传，最大长度64个字符；
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 指定签署人非渠道企业下员工，在ApproverType为ORGANIZATION时指定。
	// 默认为false，即签署人位于同一个渠道应用号下；
	NotChannelOrganization *bool `json:"NotChannelOrganization,omitempty" name:"NotChannelOrganization"`

	// 用户侧第三方id，最大长度64个字符
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 企业签署方在同一渠道下的其他合作企业OpenId，签署方为非发起方企业场景下必传，最大长度64个字符；
	OrganizationOpenId *string `json:"OrganizationOpenId,omitempty" name:"OrganizationOpenId"`

	// 签署人类型，PERSON-个人；
	// PERSON_AUTO_SIGN-个人自动签；
	// ORGANIZATION-企业；
	// ENTERPRISESERVER-企业静默签;
	// 注：ENTERPRISESERVER 类型仅用于使用文件创建签署流程（ChannelCreateFlowByFiles）接口；并且仅能指定发起方企业签署方为静默签署；
	ApproverType *string `json:"ApproverType,omitempty" name:"ApproverType"`

	// 签署流程签署人在模板中对应的签署人Id；在非单方签署、以及非B2C签署的场景下必传，用于指定当前签署方在签署流程中的位置；
	RecipientId *string `json:"RecipientId,omitempty" name:"RecipientId"`

	// 签署截止时间，默认一年
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 签署完回调url，最大长度1000个字符
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 使用PDF文件直接发起合同时，签署人指定的签署控件
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 个人签署方指定签署控件类型，目前仅支持：OCR_ESIGN(AI智慧手写签名)
	ComponentLimitType []*string `json:"ComponentLimitType,omitempty" name:"ComponentLimitType"`

	// 合同的强制预览时间：3~300s，未指定则按合同页数计算
	PreReadTime *int64 `json:"PreReadTime,omitempty" name:"PreReadTime"`

	// 签署完前端跳转的url，暂未使用
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`
}

type FlowDetailInfo struct {
	// 合同(流程)的Id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 合同(流程)的名字
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 合同(流程)的类型
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 合同(流程)的状态
	FlowStatus *string `json:"FlowStatus,omitempty" name:"FlowStatus"`

	// 合同(流程)的信息
	FlowMessage *string `json:"FlowMessage,omitempty" name:"FlowMessage"`

	// 合同(流程)的创建时间戳
	CreateOn *int64 `json:"CreateOn,omitempty" name:"CreateOn"`

	// 合同(流程)的签署截止时间戳
	DeadLine *int64 `json:"DeadLine,omitempty" name:"DeadLine"`

	// 用户自定义数据
	CustomData *string `json:"CustomData,omitempty" name:"CustomData"`

	// 合同(流程)的签署人数组
	FlowApproverInfos []*FlowApproverDetail `json:"FlowApproverInfos,omitempty" name:"FlowApproverInfos"`
}

type FlowInfo struct {
	// 合同名字，最大长度200个字符
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署截止时间戳，超过有效签署时间则该签署流程失败，默认一年
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 多个签署人信息，最大支持50个签署方
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitempty" name:"FlowApprovers"`

	// 表单K-V对列表
	FormFields []*FormField `json:"FormFields,omitempty" name:"FormFields"`

	// 回调地址，最大长度1000个字符
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 合同类型，如：1. “劳务”；2. “销售”；3. “租赁”；4. “其他”，最大长度200个字符
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 合同描述，最大长度1000个字符
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 渠道的业务信息，最大长度1000个字符。发起自动签署时，需设置对应自动签署场景，目前仅支持场景：处方单-E_PRESCRIPTION_AUTO_SIGN
	CustomerData *string `json:"CustomerData,omitempty" name:"CustomerData"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitempty" name:"CustomShowMap"`

	// 被抄送人的信息列表，抄送功能暂不开放
	CcInfos []*CcInfo `json:"CcInfos,omitempty" name:"CcInfos"`
}

type FlowResourceUrlInfo struct {
	// 流程对应Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 流程对应资源链接信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceUrlInfos []*ResourceUrlInfo `json:"ResourceUrlInfos,omitempty" name:"ResourceUrlInfos"`
}

type FormField struct {
	// 表单域或控件的Value
	ComponentValue *string `json:"ComponentValue,omitempty" name:"ComponentValue"`

	// 表单域或控件的ID，跟ComponentName二选一，不能全为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`

	// 控件的名字，跟ComponentId二选一，不能全为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

// Predefined struct for user
type GetDownloadFlowUrlRequestParams struct {
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 文件夹数组，签署流程总数不能超过50个，一个文件夹下，不能超过20个签署流程
	DownLoadFlows []*DownloadFlowInfo `json:"DownLoadFlows,omitempty" name:"DownLoadFlows"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type GetDownloadFlowUrlRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 文件夹数组，签署流程总数不能超过50个，一个文件夹下，不能超过20个签署流程
	DownLoadFlows []*DownloadFlowInfo `json:"DownLoadFlows,omitempty" name:"DownLoadFlows"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *GetDownloadFlowUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDownloadFlowUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "DownLoadFlows")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDownloadFlowUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDownloadFlowUrlResponseParams struct {
	// 合同（流程）下载地址
	DownLoadUrl *string `json:"DownLoadUrl,omitempty" name:"DownLoadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDownloadFlowUrlResponse struct {
	*tchttp.BaseResponse
	Response *GetDownloadFlowUrlResponseParams `json:"Response"`
}

func (r *GetDownloadFlowUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDownloadFlowUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OperateChannelTemplateRequestParams struct {
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 操作类型，查询:"SELECT"，删除:"DELETE"，更新:"UPDATE"
	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`

	// 渠道方模板库模板唯一标识
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 合作企业方第三方机构唯一标识数据，支持多个， 用","进行分隔
	ProxyOrganizationOpenIds *string `json:"ProxyOrganizationOpenIds,omitempty" name:"ProxyOrganizationOpenIds"`

	// 模板可见性, 全部可见-"all", 部分可见-"part"
	AuthTag *string `json:"AuthTag,omitempty" name:"AuthTag"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type OperateChannelTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 操作类型，查询:"SELECT"，删除:"DELETE"，更新:"UPDATE"
	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`

	// 渠道方模板库模板唯一标识
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 合作企业方第三方机构唯一标识数据，支持多个， 用","进行分隔
	ProxyOrganizationOpenIds *string `json:"ProxyOrganizationOpenIds,omitempty" name:"ProxyOrganizationOpenIds"`

	// 模板可见性, 全部可见-"all", 部分可见-"part"
	AuthTag *string `json:"AuthTag,omitempty" name:"AuthTag"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *OperateChannelTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OperateChannelTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "OperateType")
	delete(f, "TemplateId")
	delete(f, "ProxyOrganizationOpenIds")
	delete(f, "AuthTag")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OperateChannelTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OperateChannelTemplateResponseParams struct {
	// 腾讯电子签颁发给渠道的应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 渠道方模板库模板唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 全部成功-"all-success",部分成功-"part-success", 全部失败-"fail"失败的会在FailMessageList中展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateResult *string `json:"OperateResult,omitempty" name:"OperateResult"`

	// 模板可见性, 全部可见-"all", 部分可见-"part"
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthTag *string `json:"AuthTag,omitempty" name:"AuthTag"`

	// 合作企业方第三方机构唯一标识数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyOrganizationOpenIds []*string `json:"ProxyOrganizationOpenIds,omitempty" name:"ProxyOrganizationOpenIds"`

	// 操作失败信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailMessageList []*AuthFailMessage `json:"FailMessageList,omitempty" name:"FailMessageList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OperateChannelTemplateResponse struct {
	*tchttp.BaseResponse
	Response *OperateChannelTemplateResponseParams `json:"Response"`
}

func (r *OperateChannelTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OperateChannelTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PrepareFlowsRequestParams struct {
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 多个合同（签署流程）信息，最大支持20个签署流程。
	FlowInfos []*FlowInfo `json:"FlowInfos,omitempty" name:"FlowInfos"`

	// 操作完成后的跳转地址，最大长度200
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type PrepareFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 多个合同（签署流程）信息，最大支持20个签署流程。
	FlowInfos []*FlowInfo `json:"FlowInfos,omitempty" name:"FlowInfos"`

	// 操作完成后的跳转地址，最大长度200
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *PrepareFlowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PrepareFlowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowInfos")
	delete(f, "JumpUrl")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PrepareFlowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PrepareFlowsResponseParams struct {
	// 待发起文件确认页
	ConfirmUrl *string `json:"ConfirmUrl,omitempty" name:"ConfirmUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PrepareFlowsResponse struct {
	*tchttp.BaseResponse
	Response *PrepareFlowsResponseParams `json:"Response"`
}

func (r *PrepareFlowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PrepareFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProxyOrganizationOperator struct {
	// 经办人ID（渠道颁发），最大长度64个字符
	Id *string `json:"Id,omitempty" name:"Id"`

	// 经办人姓名，最大长度50个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 经办人身份证件类型
	// 1.ID_CARD 居民身份证
	// 2.HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证
	// 3.HONGKONG_AND_MACAO 港澳居民来往内地通行证
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 经办人证件号
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 经办人手机号，大陆手机号输入11位，暂不支持海外手机号。
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
}

type Recipient struct {
	// 签署人唯一标识
	RecipientId *string `json:"RecipientId,omitempty" name:"RecipientId"`

	// 签署方类型：ENTERPRISE-企业INDIVIDUAL-自然人
	RecipientType *string `json:"RecipientType,omitempty" name:"RecipientType"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 签署方备注信息
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 是否需要校验
	RequireValidation *bool `json:"RequireValidation,omitempty" name:"RequireValidation"`

	// 是否必须填写
	RequireSign *bool `json:"RequireSign,omitempty" name:"RequireSign"`

	// 签署类型
	SignType *int64 `json:"SignType,omitempty" name:"SignType"`

	// 签署顺序：数字越小优先级越高
	RoutingOrder *int64 `json:"RoutingOrder,omitempty" name:"RoutingOrder"`

	// 是否是发起方
	IsPromoter *bool `json:"IsPromoter,omitempty" name:"IsPromoter"`
}

type ResourceUrlInfo struct {
	// 资源链接地址，过期时间5分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type SignQrCode struct {
	// 二维码id
	QrCodeId *string `json:"QrCodeId,omitempty" name:"QrCodeId"`

	// 二维码url
	QrCodeUrl *string `json:"QrCodeUrl,omitempty" name:"QrCodeUrl"`

	// 二维码过期时间
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
}

type SignUrlInfo struct {
	// 签署链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignUrl *string `json:"SignUrl,omitempty" name:"SignUrl"`

	// 链接失效时间,默认30分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 当流程为顺序签署此参数有效时，数字越小优先级越高，暂不支持并行签署 可选
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignOrder *int64 `json:"SignOrder,omitempty" name:"SignOrder"`

	// 签署人编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"SignId,omitempty" name:"SignId"`

	// 自定义用户编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomUserId *string `json:"CustomUserId,omitempty" name:"CustomUserId"`

	// 用户姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户手机号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 签署参与者机构名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 参与者类型:
	// ORGANIZATION 企业经办人
	// PERSON 自然人
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverType *string `json:"ApproverType,omitempty" name:"ApproverType"`

	// 经办人身份证号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 签署链接对应流程Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 企业经办人 用户在渠道的编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
}

type SyncFailReason struct {
	// 经办人Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 失败原因
	// 例如：Id不符合规范、证件号码不合法等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

// Predefined struct for user
type SyncProxyOrganizationOperatorsRequestParams struct {
	// 渠道应用相关信息。 此接口Agent.AppId 和 Agent.ProxyOrganizationOpenId必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 操作类型，新增:"CREATE"，修改:"UPDATE"，离职:"RESIGN"
	OperatorType *string `json:"OperatorType,omitempty" name:"OperatorType"`

	// 经办人信息列表，最大长度200
	ProxyOrganizationOperators []*ProxyOrganizationOperator `json:"ProxyOrganizationOperators,omitempty" name:"ProxyOrganizationOperators"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type SyncProxyOrganizationOperatorsRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。 此接口Agent.AppId 和 Agent.ProxyOrganizationOpenId必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 操作类型，新增:"CREATE"，修改:"UPDATE"，离职:"RESIGN"
	OperatorType *string `json:"OperatorType,omitempty" name:"OperatorType"`

	// 经办人信息列表，最大长度200
	ProxyOrganizationOperators []*ProxyOrganizationOperator `json:"ProxyOrganizationOperators,omitempty" name:"ProxyOrganizationOperators"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *SyncProxyOrganizationOperatorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncProxyOrganizationOperatorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "OperatorType")
	delete(f, "ProxyOrganizationOperators")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncProxyOrganizationOperatorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncProxyOrganizationOperatorsResponseParams struct {
	// Status 同步状态,全部同步失败接口会直接报错
	// 1-成功 
	// 2-部分成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 同步失败经办人及其失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedList []*SyncFailReason `json:"FailedList,omitempty" name:"FailedList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SyncProxyOrganizationOperatorsResponse struct {
	*tchttp.BaseResponse
	Response *SyncProxyOrganizationOperatorsResponseParams `json:"Response"`
}

func (r *SyncProxyOrganizationOperatorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncProxyOrganizationOperatorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncProxyOrganizationRequestParams struct {
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 渠道侧合作企业名称，最大长度64个字符
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitempty" name:"ProxyOrganizationName"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitempty" name:"BusinessLicense"`

	// 渠道侧合作企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type SyncProxyOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 渠道侧合作企业名称，最大长度64个字符
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitempty" name:"ProxyOrganizationName"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitempty" name:"BusinessLicense"`

	// 渠道侧合作企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

func (r *SyncProxyOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncProxyOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "ProxyOrganizationName")
	delete(f, "BusinessLicense")
	delete(f, "UniformSocialCreditCode")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncProxyOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncProxyOrganizationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SyncProxyOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *SyncProxyOrganizationResponseParams `json:"Response"`
}

func (r *SyncProxyOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncProxyOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateInfo struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 模板名字
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 模板描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 模板控件信息结构
	Components []*Component `json:"Components,omitempty" name:"Components"`

	// 签署区模板信息结构
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 模板中的流程参与人信息
	Recipients []*Recipient `json:"Recipients,omitempty" name:"Recipients"`

	// 模板类型：1-静默签；3-普通模板
	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`

	// 是否是发起人
	IsPromoter *bool `json:"IsPromoter,omitempty" name:"IsPromoter"`

	// 模板的创建者信息
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 模板创建的时间戳（精确到秒）
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`
}

type UploadFile struct {
	// Base64编码后的文件内容
	FileBody *string `json:"FileBody,omitempty" name:"FileBody"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

// Predefined struct for user
type UploadFilesRequestParams struct {
	// 应用相关信息，若是渠道版调用 appid 和proxyappid 必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 文件对应业务类型，用于区分文件存储路径：
	// 1. TEMPLATE - 模板； 文件类型：.pdf
	// 2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.jpg/.png
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 上传文件内容数组，最多支持20个文件
	FileInfos []*UploadFile `json:"FileInfos,omitempty" name:"FileInfos"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type UploadFilesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息，若是渠道版调用 appid 和proxyappid 必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 文件对应业务类型，用于区分文件存储路径：
	// 1. TEMPLATE - 模板； 文件类型：.pdf
	// 2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.jpg/.png
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 上传文件内容数组，最多支持20个文件
	FileInfos []*UploadFile `json:"FileInfos,omitempty" name:"FileInfos"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	delete(f, "Agent")
	delete(f, "BusinessType")
	delete(f, "FileInfos")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFilesResponseParams struct {
	// 文件id数组，有效期一个小时
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 上传成功文件数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 文件Url
	FileUrls []*string `json:"FileUrls,omitempty" name:"FileUrls"`

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

type UsageDetail struct {
	// 渠道侧合作企业唯一标识
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitempty" name:"ProxyOrganizationOpenId"`

	// 渠道侧合作企业名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitempty" name:"ProxyOrganizationName"`

	// 日期，当需要汇总数据时日期为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 消耗数量
	Usage *uint64 `json:"Usage,omitempty" name:"Usage"`

	// 撤回数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cancel *uint64 `json:"Cancel,omitempty" name:"Cancel"`
}

type UserInfo struct {
	// 用户在渠道的编号
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 用户的来源渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 自定义用户编号
	CustomUserId *string `json:"CustomUserId,omitempty" name:"CustomUserId"`

	// 用户真实IP
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 用户代理IP
	ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
}