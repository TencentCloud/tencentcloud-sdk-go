// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Agent struct {
	// 应用的唯一标识(由电子签平台自动生成)。不同的业务系统可以采用不同的AppId，不同AppId下的数据是隔离的。可以由控制台开发者中心-应用集成自主生成。位置如下:
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/fac77e0d3f28aaec56669f67e65c8db8.png)
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 第三方应用平台自定义，对应第三方平台子客企业的唯一标识。一个第三方平台子客企业主体与子客企业ProxyOrganizationOpenId是一一对应的，不可更改，不可重复使用。（例如，可以使用企业名称的hash值，或者社会统一信用代码的hash值，或者随机hash值，需要第三方应用平台保存），最大64位字符串
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil,omitempty" name:"ProxyOrganizationOpenId"`

	// 第三方平台子客企业中的员工/经办人，通过第三方应用平台进入电子签完成实名、且被赋予相关权限后，可以参与到企业资源的管理或签署流程中。
	ProxyOperator *UserInfo `json:"ProxyOperator,omitnil,omitempty" name:"ProxyOperator"`

	// **不用填写**，在第三方平台子客企业开通电子签后，会生成唯一的子客应用Id（ProxyAppId）用于代理调用时的鉴权，在子客开通的回调中获取。
	ProxyAppId *string `json:"ProxyAppId,omitnil,omitempty" name:"ProxyAppId"`

	// 内部参数，暂未开放使用
	//
	// Deprecated: ProxyOrganizationId is deprecated.
	ProxyOrganizationId *string `json:"ProxyOrganizationId,omitnil,omitempty" name:"ProxyOrganizationId"`
}

type ApproverComponentLimitType struct {
	// 签署方经办人在模板中配置的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 签署方经办人控件类型是个人印章签署控件（SIGN_SIGNATURE） 时，可选的签名方式。
	// 
	// 签名方式：
	// <ul>
	// <li>HANDWRITE-手写签名</li>
	// <li>ESIGN-个人印章类型</li>
	// <li>OCR_ESIGN-AI智能识别手写签名</li>
	// <li>SYSTEM_ESIGN-系统签名</li>
	// </ul>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type ApproverItem struct {
	// 签署方唯一编号
	// 
	// 在<a href="https://qian.tencent.com/developers/company/dynamic_signer" target="_blank">动态补充签署人</a>场景下，可以用此编号确定参与方
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 签署方角色编号
	// 
	// 在<a href="https://qian.tencent.com/developers/company/dynamic_signer" target="_blank">动态补充签署人</a>场景下，可以用此编号确定参与方
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 签署方角色名称
	ApproverRoleName *string `json:"ApproverRoleName,omitnil,omitempty" name:"ApproverRoleName"`
}

type ApproverOption struct {
	// 是否可以拒签 默认false-可以拒签 true-不可以拒签
	NoRefuse *bool `json:"NoRefuse,omitnil,omitempty" name:"NoRefuse"`

	// 是否可以转发 默认false-可以转发 true-不可以转发
	NoTransfer *bool `json:"NoTransfer,omitnil,omitempty" name:"NoTransfer"`

	// 当签署方有多个签署区时候，是否隐藏一键所有的签署区
	// 
	// false：（默认）不隐藏
	// true：隐藏，每个签署区要单独选择印章或者签名
	HideOneKeySign *bool `json:"HideOneKeySign,omitnil,omitempty" name:"HideOneKeySign"`

	// 签署人信息补充类型，默认无需补充。
	// 
	// <ul><li> **1** :  动态签署人（可发起合同后再补充签署人信息）注：`企业自动签不支持动态补充`</li></ul>
	// 
	// 注：
	// `使用动态签署人能力前，需登陆腾讯电子签控制台打开服务开关`
	FillType *int64 `json:"FillType,omitnil,omitempty" name:"FillType"`

	// 签署人阅读合同限制参数
	//  <br/>取值：
	// <ul>
	// <li> LimitReadTimeAndBottom，阅读合同必须限制阅读时长并且必须阅读到底</li>
	// <li> LimitReadTime，阅读合同仅限制阅读时长</li>
	// <li> LimitBottom，阅读合同仅限制必须阅读到底</li>
	// <li> NoReadTimeAndBottom，阅读合同不限制阅读时长且不限制阅读到底（白名单功能，请联系客户经理开白使用）</li>
	// </ul>
	FlowReadLimit *string `json:"FlowReadLimit,omitnil,omitempty" name:"FlowReadLimit"`

	// 禁止在签署过程中添加签署日期控件
	//  <br/>前置条件：文件发起合同时，指定SignBeanTag=1（可以在签署过程中添加签署控件）：
	// <ul>
	// <li> 默认值：false，在开启：签署过程中添加签署控件时，添加签署控件会默认自带签署日期控件</li>
	// <li> 可选值：true，在开启：签署过程中添加签署控件时，添加签署控件不会自带签署日期控件</li>
	// </ul>
	ForbidAddSignDate *bool `json:"ForbidAddSignDate,omitnil,omitempty" name:"ForbidAddSignDate"`
}

type ApproverRestriction struct {
	// 指定签署人姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 指定签署人手机号，11位数字
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 指定签署人证件类型，ID_CARD-身份证，HONGKONG_AND_MACAO-港澳居民来往内地通行证，HONGKONG_MACAO_AND_TAIWAN-港澳台居民居住证
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 指定签署人证件号码，其中字母大写
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`
}

// Predefined struct for user
type ArchiveDynamicFlowRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type ArchiveDynamicFlowRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

func (r *ArchiveDynamicFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ArchiveDynamicFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ArchiveDynamicFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ArchiveDynamicFlowResponseParams struct {
	// 合同流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 动态签署人的参与人信息
	Approvers []*ChannelArchiveDynamicApproverData `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ArchiveDynamicFlowResponse struct {
	*tchttp.BaseResponse
	Response *ArchiveDynamicFlowResponseParams `json:"Response"`
}

func (r *ArchiveDynamicFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ArchiveDynamicFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthFailMessage struct {
	// 第三方平台子客企业的唯一标识，长度不能超过64，只能由字母和数字组成。开发者可自定义此字段的值，并需要保存此 ID 以便进行后续操作。
	// 
	// 一个第三方平台子客企业主体与子客企业 ProxyOrganizationOpenId 是一一对应的，不可更改，不可重复使用。例如，可以使用企业名称的哈希值，或者社会统一信用代码的哈希值，或者随机哈希值。
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil,omitempty" name:"ProxyOrganizationOpenId"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type AuthInfoDetail struct {
	// 扩展服务类型，和入参一致	
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 扩展服务名称	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 授权员工列表	
	HasAuthUserList []*HasAuthUser `json:"HasAuthUserList,omitnil,omitempty" name:"HasAuthUserList"`

	// 授权企业列表（企业自动签时，该字段有值）	
	HasAuthOrganizationList []*HasAuthOrganization `json:"HasAuthOrganizationList,omitnil,omitempty" name:"HasAuthOrganizationList"`

	// 授权员工列表总数	
	AuthUserTotal *int64 `json:"AuthUserTotal,omitnil,omitempty" name:"AuthUserTotal"`

	// 授权企业列表总数	
	AuthOrganizationTotal *int64 `json:"AuthOrganizationTotal,omitnil,omitempty" name:"AuthOrganizationTotal"`
}

type AuthorizedUser struct {
	// 第三方应用平台的用户openid
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`
}

type AutoSignConfig struct {
	// 自动签开通个人用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 是否回调证书信息:
	// <ul><li>**false**: 不需要(默认)</li>
	// <li>**true**:需要</li></ul>
	// 
	// 
	// 注：`该字段已经失效，请勿设置此参数。`
	CertInfoCallback *bool `json:"CertInfoCallback,omitnil,omitempty" name:"CertInfoCallback"`

	// 是否支持用户自定义签名印章:
	// <ul><li>**false**: 不能自己定义(默认)</li>
	// <li>**true**: 可以自己定义</li></ul>
	UserDefineSeal *bool `json:"UserDefineSeal,omitnil,omitempty" name:"UserDefineSeal"`

	// 回调中是否需要自动签将要使用的印章（签名）图片的 base64:
	// <ul><li>**false**: 不需要(默认)</li>
	// <li>**true**: 需要</li></ul>
	SealImgCallback *bool `json:"SealImgCallback,omitnil,omitempty" name:"SealImgCallback"`

	// 该字段已废弃，请使用【应用号配置】中的回调地址统一接口消息
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 开通时候的身份验证方式, 取值为：
	// <ul><li>**WEIXINAPP** : 微信人脸识别</li>
	// <li>**INSIGHT** : 慧眼人脸识别</li>
	// <li>**TELECOM** : 运营商三要素验证</li></ul>
	// 注：
	// <ul><li>如果是小程序开通链接，仅支持传 WEIXINAPP。为空默认 WEIXINAPP</li>
	// <li>如果是 H5 开通链接，支持传 INSIGHT / TELECOM。为空默认 INSIGHT </li></ul>
	VerifyChannels []*string `json:"VerifyChannels,omitnil,omitempty" name:"VerifyChannels"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。
	// 
	// <ul><li><b>1</b>: (默认)不绑定自动签账号许可开通，开通后一直有效,   后续使用合同份额进行合同发起</li></ul>
	// 
	// 注：`该字段已经失效，请勿设置此参数。`
	LicenseType *int64 `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 开通成功后前端页面跳转的url，此字段的用法场景请联系客户经理确认。
	// 
	// 注：`仅支持H5开通场景`, `跳转链接仅支持 https:// , qianapp:// 开头`
	// 
	// 跳转场景：
	// <ul><li>**贵方H5 -> 腾讯电子签H5 -> 贵方H5** : JumpUrl格式: https://YOUR_CUSTOM_URL/xxxx，只需满足 https:// 开头的正确且合规的网址即可。</li>
	// <li>**贵方原生App -> 腾讯电子签H5 -> 贵方原生App** : JumpUrl格式: qianapp://YOUR_CUSTOM_URL，只需满足 qianapp:// 开头的URL即可。`APP实现方，需要拦截Webview地址跳转，发现url是qianapp:// 开头时跳转到原生页面。`APP拦截地址跳转可参考：<a href='https://stackoverflow.com/questions/41693263/android-webview-err-unknown-url-scheme'>Android</a>，<a href='https://razorpay.com/docs/payments/payment-gateway/web-integration/standard/webview/upi-intent-ios/'>IOS</a> </li></ul>
	// 
	// 成功结果返回：
	// 若贵方需要在跳转回时通过链接query参数提示开通成功，JumpUrl中的query应携带如下参数：`appendResult=qian`。这样腾讯电子签H5会在跳转回的url后面会添加query参数提示贵方签署成功，例如：qianapp://YOUR_CUSTOM_URL?action=sign&result=success&from=tencent_ess
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`
}

type BaseFlowInfo struct {
	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil,omitempty" name:"FlowDescription"`

	// 合同流程的签署顺序类型：
	// **false**：(默认)有序签署, 本合同多个参与人需要依次签署
	// **true**：无序签署, 本合同多个参与人没有先后签署限制
	Unordered *bool `json:"Unordered,omitnil,omitempty" name:"Unordered"`

	// 是否打开智能添加填写区(默认开启，打开:"OPEN" 关闭："CLOSE")
	IntelligentStatus *string `json:"IntelligentStatus,omitnil,omitempty" name:"IntelligentStatus"`

	// 填写控件内容， 填写的控制的ID-填写的内容对列表
	FormFields []*FormField `json:"FormFields,omitnil,omitempty" name:"FormFields"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil,omitempty" name:"NeedSignReview"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为1000长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的回调通知模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// 注
	// 1. 抄送人名单中可以包括自然人以及本企业的员工（本企业员工必须已经完成认证并加入企业）。
	// 2. 请确保抄送人列表中的成员不与任何签署人重复。
	CcInfos []*CcInfo `json:"CcInfos,omitnil,omitempty" name:"CcInfos"`

	// 发起方企业的签署人进行发起操作是否需要企业内部审批。使用此功能需要发起方企业有参与签署。
	// 
	// 若设置为true，发起审核结果需通过接口 [提交企业签署流程审批结果](https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateFlowSignReview)通知电子签，审核通过后，发起方企业签署人方可进行发起操作，否则会阻塞其发起操作。
	// 
	NeedCreateReview *bool `json:"NeedCreateReview,omitnil,omitempty" name:"NeedCreateReview"`

	// 填写控件：文件发起使用
	Components []*Component `json:"Components,omitnil,omitempty" name:"Components"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下：  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li><li> <b>3</b> :文书</li></ul>效果如下:![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"FlowDisplayType,omitnil,omitempty" name:"FlowDisplayType"`

	// 签署文件资源Id列表，目前仅支持单个文件
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 合同签署人信息
	Approvers []*CommonFlowApprover `json:"Approvers,omitnil,omitempty" name:"Approvers"`
}

type BatchOrganizationRegistrationTasksDetails struct {
	// 生成注册链接的任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 批量创建企业任务的状态
	// <ul>
	// <li>Processing</li>
	// <li>Create</li>
	// <li>Submit</li>
	// <li>Authorization</li>
	// <li>Failed</li>
	// </ul>
	// 
	// 各个状态所代表的含义如下表格所示：
	// <table>
	// <thead align="center" valign="center">
	// <tr><th>任务状态名称</th><th>任务状态详情</th></tr>
	// </thead>
	// <tbody>
	// <tr><th align="center" valign="center">Processing</th><th>企业认证任务处理中，用户调用了<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks">CreateBatchOrganizationRegistrationTasks</a>接口，但是任务还在处理中的状态</th></tr>
	// <tr><th align="center" valign="center">Create</th><th>创建企业认证链接任务完成，可以调用生成任务链接接口</th></tr>
	// <tr><th align="center" valign="center">Submit</th><th>企业认证任务已提交,到如下界面之后，会变为这个状态
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/acbcec8c7a71de14d9c041e3b8ca8b3f.png)</th></tr>
	// <tr><th align="center" valign="center">Authorization</th><th>企业认证任务认证成功,点击下图下一步，进入到授权书上传或者法人认证，则会变为这个状态
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/c52448354871cffa729da8db4e3a6f18.png)</th></tr>
	// <tr><th align="center" valign="center">Failed</th><th>企业认证任务失败</th></tr>
	// </tbody>
	// </table>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 如果任务失败,会返回错误信息
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

type CancelFailureFlow struct {
	// 签署流程编号，为32位字符串
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 撤销失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type CcInfo struct {
	// 被抄送方手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此业务通知方。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 被抄送方姓名。
	// 抄送方的姓名将用于身份认证，请确保填写的姓名为抄送方的真实姓名，而非昵称等代名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 被抄送方类型, 可设置以下类型:
	// <ul><li> **0** :个人抄送方</li>
	// <li> **1** :企业员工抄送方</li></ul>
	CcType *int64 `json:"CcType,omitnil,omitempty" name:"CcType"`

	// 被抄送方权限, 可设置如下权限:
	// <ul><li> **0** :可查看合同内容</li>
	// <li> **1** :可查看合同内容也可下载原文</li></ul>
	CcPermission *int64 `json:"CcPermission,omitnil,omitempty" name:"CcPermission"`
}

type ChannelArchiveDynamicApproverData struct {
	// 签署方唯一编号，一个全局唯一的标识符，不同的流程不会出现冲突。 可以使用签署方的唯一编号来生成签署链接（也可以通过RecipientId来生成签署链接）。	
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 签署方角色编号，签署方角色编号是用于区分同一个流程中不同签署方的唯一标识。不同的流程会出现同样的签署方角色编号。 填写控件和签署控件都与特定的角色编号关联。	
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`
}

// Predefined struct for user
type ChannelBatchCancelFlowsRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要撤销的合同流程ID列表，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 撤回原因，长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 
	// 备注:`如果不传递撤回原因，那么默认撤回原因是 "自动撤销（通过接口实现）"`
	CancelMessage *string `json:"CancelMessage,omitnil,omitempty" name:"CancelMessage"`

	// 撤销理由自定义格式,  会展示在合同预览的界面中,  可以选择下面的组合方式：
	// 
	// **0** : 默认格式,  合同封面页面会展示为: 发起方-企业名称-撤销的经办人名字以**CancelMessage**的理由撤销当前合同
	// **1** :  合同封面页面会展示为:  发起方以**CancelMessage**的理由撤销当前合同
	// **2** : 保留企业名称,  合同封面页面会展示为:  发起方-企业名称以**CancelMessage**的理由撤销当前合同
	// **3** : 保留企业名称+经办人名字,  合同封面页面会展示为: 发起方-企业名称-撤销的经办人名字以**CancelMessage**的理由撤销当前合同
	// 
	// 注: `CancelMessage为撤销当前合同的理由`
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/f16cf37dbb3a09d6569877f093b92204/channel_ChannelCancelFlow.png)
	// 
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitnil,omitempty" name:"CancelMessageFormat"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelBatchCancelFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要撤销的合同流程ID列表，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 撤回原因，长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 
	// 备注:`如果不传递撤回原因，那么默认撤回原因是 "自动撤销（通过接口实现）"`
	CancelMessage *string `json:"CancelMessage,omitnil,omitempty" name:"CancelMessage"`

	// 撤销理由自定义格式,  会展示在合同预览的界面中,  可以选择下面的组合方式：
	// 
	// **0** : 默认格式,  合同封面页面会展示为: 发起方-企业名称-撤销的经办人名字以**CancelMessage**的理由撤销当前合同
	// **1** :  合同封面页面会展示为:  发起方以**CancelMessage**的理由撤销当前合同
	// **2** : 保留企业名称,  合同封面页面会展示为:  发起方-企业名称以**CancelMessage**的理由撤销当前合同
	// **3** : 保留企业名称+经办人名字,  合同封面页面会展示为: 发起方-企业名称-撤销的经办人名字以**CancelMessage**的理由撤销当前合同
	// 
	// 注: `CancelMessage为撤销当前合同的理由`
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/f16cf37dbb3a09d6569877f093b92204/channel_ChannelCancelFlow.png)
	// 
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitnil,omitempty" name:"CancelMessageFormat"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelBatchCancelFlowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelBatchCancelFlowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowIds")
	delete(f, "CancelMessage")
	delete(f, "CancelMessageFormat")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelBatchCancelFlowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelBatchCancelFlowsResponseParams struct {
	// 签署流程批量撤销失败原因，错误信息与流程Id一一对应，成功为"", 失败则对应失败原因
	// 
	// 注:  `如果全部撤销成功, 此数组为空数组`
	FailMessages []*string `json:"FailMessages,omitnil,omitempty" name:"FailMessages"`

	// 批量撤销任务编号，为32位字符串，可用于[查询批量撤销合同结果](https://qian.tencent.com/developers/partnerApis/operateFlows/DescribeCancelFlowsTask) 或关联[批量撤销任务结果回调](https://qian.tencent.com/developers/partner/callback_types_contracts_sign#%E4%B9%9D-%E6%89%B9%E9%87%8F%E6%92%A4%E9%94%80%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83)
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelBatchCancelFlowsResponse struct {
	*tchttp.BaseResponse
	Response *ChannelBatchCancelFlowsResponseParams `json:"Response"`
}

func (r *ChannelBatchCancelFlowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelBatchCancelFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChannelBillUsageDetail struct {
	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 合同经办人名称
	// 如果有多个经办人用分号隔开。
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 发起方组织机构名称
	CreateOrganizationName *string `json:"CreateOrganizationName,omitnil,omitempty" name:"CreateOrganizationName"`

	// 合同流程的名称。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 合同流程当前的签署状态, 会存在下列的状态值
	// <ul>
	// <li>**INIT**: 合同创建</li>
	// <li>**PART**: 合同签署中(至少有一个签署方已经签署)</li>
	// <li>**REJECT**: 合同拒签</li>
	// <li>**ALL**: 合同签署完成</li>
	// <li>**DEADLINE**: 合同流签(合同过期)</li>
	// <li>**CANCEL**: 合同撤回</li>
	// <li>**RELIEVED**: 解除协议（已解除）</li>
	// <li>**WILLEXPIRE**: 合同即将过期</li>
	// <li>**EXCEPTION**: 合同异常</li>
	// </ul>
	FlowStatus *string `json:"FlowStatus,omitnil,omitempty" name:"FlowStatus"`

	// 查询的套餐类型
	// 对应关系如下:
	// <ul>
	// <li>**CloudEnterprise**: 企业版合同</li>
	// <li>**SingleSignature**: 单方签章</li>
	// <li>**CloudProve**: 签署报告</li>
	// <li>**CloudOnlineSign**: 腾讯会议在线签约</li>
	// <li>**ChannelWeCard**: 微工卡</li>
	// <li>**SignFlow**: 合同套餐</li>
	// <li>**SignFace**: 签署意愿（人脸识别）</li>
	// <li>**SignPassword**: 签署意愿（密码）</li>
	// <li>**SignSMS**: 签署意愿（短信）</li>
	// <li>**PersonalEssAuth**: 签署人实名（腾讯电子签认证）</li>
	// <li>**PersonalThirdAuth**: 签署人实名（信任第三方认证）</li>
	// <li>**OrgEssAuth**: 签署企业实名</li>
	// <li>**FlowNotify**: 短信通知</li>
	// <li>**AuthService**: 企业工商信息查询</li>
	// </ul>
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// 合同使用量
	// 注: `如果消耗类型是撤销返还，此值为负值代表返还的合同数量`
	UseCount *int64 `json:"UseCount,omitnil,omitempty" name:"UseCount"`

	// 消耗的时间戳，格式为Unix标准时间戳（秒）。
	CostTime *int64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// 消耗的套餐名称
	QuotaName *string `json:"QuotaName,omitnil,omitempty" name:"QuotaName"`

	// 消耗类型
	// **1**.扣费 
	// **2**.撤销返还
	CostType *int64 `json:"CostType,omitnil,omitempty" name:"CostType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

// Predefined struct for user
type ChannelCancelFlowRequestParams struct {
	// 要撤销的合同流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 撤回原因，长度不能超过200，只能由中文、字母、数字和下划线组成。
	CancelMessage *string `json:"CancelMessage,omitnil,omitempty" name:"CancelMessage"`

	// 撤销理由自定义格式,  会展示在合同预览的界面中,  可以选择下面的组合方式：
	// 
	// **0** : 默认格式,  合同封面页面会展示为: 发起方-企业名称-撤销的经办人名字以**CancelMessage**的理由撤销当前合同
	// **1** :  合同封面页面会展示为:  发起方以**CancelMessage**的理由撤销当前合同
	// **2** : 保留企业名称,  合同封面页面会展示为:  发起方-企业名称以**CancelMessage**的理由撤销当前合同
	// **3** : 保留企业名称+经办人名字,  合同封面页面会展示为: 发起方-企业名称-撤销的经办人名字以**CancelMessage**的理由撤销当前合同
	// 
	// 注: `CancelMessage为撤销当前合同的理由`
	// 
	// ![image](https://dyn.ess.tencent.cn/guide/capi/channel_ChannelCancelFlow.png)
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitnil,omitempty" name:"CancelMessageFormat"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelCancelFlowRequest struct {
	*tchttp.BaseRequest
	
	// 要撤销的合同流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 撤回原因，长度不能超过200，只能由中文、字母、数字和下划线组成。
	CancelMessage *string `json:"CancelMessage,omitnil,omitempty" name:"CancelMessage"`

	// 撤销理由自定义格式,  会展示在合同预览的界面中,  可以选择下面的组合方式：
	// 
	// **0** : 默认格式,  合同封面页面会展示为: 发起方-企业名称-撤销的经办人名字以**CancelMessage**的理由撤销当前合同
	// **1** :  合同封面页面会展示为:  发起方以**CancelMessage**的理由撤销当前合同
	// **2** : 保留企业名称,  合同封面页面会展示为:  发起方-企业名称以**CancelMessage**的理由撤销当前合同
	// **3** : 保留企业名称+经办人名字,  合同封面页面会展示为: 发起方-企业名称-撤销的经办人名字以**CancelMessage**的理由撤销当前合同
	// 
	// 注: `CancelMessage为撤销当前合同的理由`
	// 
	// ![image](https://dyn.ess.tencent.cn/guide/capi/channel_ChannelCancelFlow.png)
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitnil,omitempty" name:"CancelMessageFormat"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelCancelFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCancelFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Agent")
	delete(f, "CancelMessage")
	delete(f, "CancelMessageFormat")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCancelFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCancelFlowResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCancelFlowResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCancelFlowResponseParams `json:"Response"`
}

func (r *ChannelCancelFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCancelFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCancelMultiFlowSignQRCodeRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要取消的签署码ID，为32位字符串。由[创建一码多签签署码](https://qian.tencent.com/developers/partnerApis/templates/ChannelCreateMultiFlowSignQRCode)返回
	QrCodeId *string `json:"QrCodeId,omitnil,omitempty" name:"QrCodeId"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelCancelMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要取消的签署码ID，为32位字符串。由[创建一码多签签署码](https://qian.tencent.com/developers/partnerApis/templates/ChannelCreateMultiFlowSignQRCode)返回
	QrCodeId *string `json:"QrCodeId,omitnil,omitempty" name:"QrCodeId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ChannelCancelUserAutoSignEnableUrlRequestParams struct {
	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 指定撤销链接的用户信息，包含姓名、证件类型、证件号码。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`
}

type ChannelCancelUserAutoSignEnableUrlRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 指定撤销链接的用户信息，包含姓名、证件类型、证件号码。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`
}

func (r *ChannelCancelUserAutoSignEnableUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCancelUserAutoSignEnableUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "Operator")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCancelUserAutoSignEnableUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCancelUserAutoSignEnableUrlResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCancelUserAutoSignEnableUrlResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCancelUserAutoSignEnableUrlResponseParams `json:"Response"`
}

func (r *ChannelCancelUserAutoSignEnableUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCancelUserAutoSignEnableUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateBatchCancelFlowUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要撤销的合同流程ID列表，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelCreateBatchCancelFlowUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要撤销的合同流程ID列表，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelCreateBatchCancelFlowUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateBatchCancelFlowUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowIds")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateBatchCancelFlowUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateBatchCancelFlowUrlResponseParams struct {
	// 批量撤销合同的URL链接, 需要在手机端打开, 有效期24小时
	// 
	// 注：<font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	BatchCancelFlowUrl *string `json:"BatchCancelFlowUrl,omitnil,omitempty" name:"BatchCancelFlowUrl"`

	// 与入参的FlowIds数组一致,   成功生成到撤销链接中,则为"",   不能撤销合同则为失败原因
	FailMessages []*string `json:"FailMessages,omitnil,omitempty" name:"FailMessages"`

	// 签署撤销链接的过期时间(格式为:年-月-日 时:分:秒), 默认是生成链接的24小时后失效
	// 
	UrlExpireOn *string `json:"UrlExpireOn,omitnil,omitempty" name:"UrlExpireOn"`

	// 批量撤销任务编号，为32位字符串，可用于[查询批量撤销合同结果](https://qian.tencent.com/developers/partnerApis/operateFlows/DescribeCancelFlowsTask) 或关联[批量撤销任务结果回调](https://qian.tencent.com/developers/partner/callback_types_contracts_sign#%E4%B9%9D-%E6%89%B9%E9%87%8F%E6%92%A4%E9%94%80%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83)
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateBatchCancelFlowUrlResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateBatchCancelFlowUrlResponseParams `json:"Response"`
}

func (r *ChannelCreateBatchCancelFlowUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateBatchCancelFlowUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateBatchQuickSignUrlRequestParams struct {
	// 批量签署的流程签署人，其中姓名(ApproverName)、参与人类型(ApproverType)必传，手机号(ApproverMobile)和证件信息(ApproverIdCardType、ApproverIdCardNumber)可任选一种或全部传入。
	// <ul>
	// <li>若为个人参与方：ApproverType:"PERSON"</li>
	// <li>若为企业参与方：ApproverType:"ORGANIZATION"。同时若签署方为saas企业员工， OrganizationName 参数需传入参与方企业名称。若签署方为渠道子客企业员工，除了 OrganizationName 还需要传 OpenId、OrganizationOpenId。（如果OrganizationOpenId对应子客企业已经认证激活，则可以省略OrganizationName参数）</li>
	// </ul>
	// 
	// 注:
	// `1. 暂不支持签署人拖动签署控件功能，以及签批控件。`
	// `2. 当需要通过短信验证码签署时，手机号ApproverMobile需要与发起合同时填写的用户手机号一致。`
	FlowApproverInfo *FlowApproverInfo `json:"FlowApproverInfo,omitnil,omitempty" name:"FlowApproverInfo"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 批量签署的合同流程ID数组。
	// 注: `在调用此接口时，请确保合同流程均为本企业发起，且合同数量不超过100个。`
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 合同组编号
	// 注：`该参数和合同流程ID数组必须二选一`
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 签署完之后的H5页面的跳转链接，此链接及支持http://和https://，最大长度1000个字符。(建议https协议)
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 指定批量签署合同的签名类型，可传递以下值：
	// <ul><li>**0**：手写签名(默认)</li>
	// <li>**1**：OCR楷体</li>
	// <li>**2**：姓名印章</li>
	// <li>**3**：图片印章</li>
	// <li>**4**：系统签名</li></ul>
	// 注：
	// <ul><li>默认情况下，签名类型为手写签名</li>
	// <li>您可以传递多种值，表示可用多种签名类型。</li>
	// <li>该参数会覆盖您合同中的签名类型，若您在发起合同时限定了签名类型(赋值签名类型给ComponentTypeLimit)，请将这些签名类型赋予此参数</li>
	// <li>若签署方为企业员工，此参数无效，签名方式将以合同中为准。</li>
	// </ul>
	SignatureTypes []*int64 `json:"SignatureTypes,omitnil,omitempty" name:"SignatureTypes"`

	// 指定批量签署合同的认证校验方式，可传递以下值：
	// <ul><li>**1**：人脸认证(默认)，需进行人脸识别成功后才能签署合同</li>
	// <li>**2**：密码认证(默认)，需输入与用户在腾讯电子签设置的密码一致才能校验成功进行合同签署</li>
	// <li>**3**：运营商三要素，需到运营商处比对手机号实名信息(名字、手机号、证件号)校验一致才能成功进行合同签署。</li></ul>
	// 注：
	// <ul><li>默认情况下，认证校验方式为人脸和密码认证</li>
	// <li>您可以传递多种值，表示可用多种认证校验方式。</li></ul>
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil,omitempty" name:"ApproverSignTypes"`

	// 生成H5签署链接时，您可以指定签署方签署合同的认证校验方式的选择模式，可传递一下值：
	// <ul><li>**0**：签署方自行选择，签署方可以从预先指定的认证方式中自由选择；</li>
	// <li>**1**：自动按顺序首位推荐，签署方无需选择，系统会优先推荐使用第一种认证方式。</li></ul>
	// 注：
	// `不指定该值时，默认为签署方自行选择。`
	SignTypeSelector *uint64 `json:"SignTypeSelector,omitnil,omitempty" name:"SignTypeSelector"`

	// 批量签署合同相关信息，指定合同和签署方的信息，用于补充动态签署人。	
	// 
	// 注: `若签署方为企业员工，暂不支持通过H5端进行动态签署人的补充`
	FlowBatchUrlInfo *FlowBatchUrlInfo `json:"FlowBatchUrlInfo,omitnil,omitempty" name:"FlowBatchUrlInfo"`

	// <b>只有在生成H5签署链接的情形下</b>（ 如调用<a href="https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateFlowSignUrl" target="_blank">获取H5签署链接</a>、<a href="https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateBatchQuickSignUrl" target="_blank">获取H5批量签署链接</a>等接口），该配置才会生效。  您可以指定H5签署视频核身的意图配置，选择问答模式或点头模式的语音文本。  注意： 1. 视频认证为<b>白名单功能，使用前请联系对接的客户经理沟通</b>。 2. 使用视频认证时，<b>生成H5签署链接的时候必须将签署认证方式指定为人脸</b>（即ApproverSignTypes设置成人脸签署）。 3. 签署完成后，可以通过<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelDescribeSignFaceVideo" target="_blank">查询签署认证人脸视频</a>获取到当时的视频。
	Intention *Intention `json:"Intention,omitnil,omitempty" name:"Intention"`

	// 缓存签署人信息。在H5签署链接动态领取场景，首次填写后，选择缓存签署人信息，在下次签署人点击领取链接时，会自动将个人信息（姓名、身份证号、手机号）填入，否则需要每次手动填写。
	// 
	// 注: `若参与方为企业员工时，暂不支持对参与方信息进行缓存`
	CacheApproverInfo *bool `json:"CacheApproverInfo,omitnil,omitempty" name:"CacheApproverInfo"`

	// 是否允许此链接中签署方批量拒签。 <ul><li>false (默认): 不允许批量拒签</li> <li>true : 允许批量拒签。</li></ul>注：`合同组暂不支持批量拒签功能。`
	CanBatchReject *bool `json:"CanBatchReject,omitnil,omitempty" name:"CanBatchReject"`

	// 预设的动态签署方的补充信息，仅匹配对应信息的签署方才能领取合同。暂时仅对个人参与方生效。
	PresetApproverInfo *PresetApproverInfo `json:"PresetApproverInfo,omitnil,omitempty" name:"PresetApproverInfo"`
}

type ChannelCreateBatchQuickSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 批量签署的流程签署人，其中姓名(ApproverName)、参与人类型(ApproverType)必传，手机号(ApproverMobile)和证件信息(ApproverIdCardType、ApproverIdCardNumber)可任选一种或全部传入。
	// <ul>
	// <li>若为个人参与方：ApproverType:"PERSON"</li>
	// <li>若为企业参与方：ApproverType:"ORGANIZATION"。同时若签署方为saas企业员工， OrganizationName 参数需传入参与方企业名称。若签署方为渠道子客企业员工，除了 OrganizationName 还需要传 OpenId、OrganizationOpenId。（如果OrganizationOpenId对应子客企业已经认证激活，则可以省略OrganizationName参数）</li>
	// </ul>
	// 
	// 注:
	// `1. 暂不支持签署人拖动签署控件功能，以及签批控件。`
	// `2. 当需要通过短信验证码签署时，手机号ApproverMobile需要与发起合同时填写的用户手机号一致。`
	FlowApproverInfo *FlowApproverInfo `json:"FlowApproverInfo,omitnil,omitempty" name:"FlowApproverInfo"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 批量签署的合同流程ID数组。
	// 注: `在调用此接口时，请确保合同流程均为本企业发起，且合同数量不超过100个。`
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 合同组编号
	// 注：`该参数和合同流程ID数组必须二选一`
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 签署完之后的H5页面的跳转链接，此链接及支持http://和https://，最大长度1000个字符。(建议https协议)
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 指定批量签署合同的签名类型，可传递以下值：
	// <ul><li>**0**：手写签名(默认)</li>
	// <li>**1**：OCR楷体</li>
	// <li>**2**：姓名印章</li>
	// <li>**3**：图片印章</li>
	// <li>**4**：系统签名</li></ul>
	// 注：
	// <ul><li>默认情况下，签名类型为手写签名</li>
	// <li>您可以传递多种值，表示可用多种签名类型。</li>
	// <li>该参数会覆盖您合同中的签名类型，若您在发起合同时限定了签名类型(赋值签名类型给ComponentTypeLimit)，请将这些签名类型赋予此参数</li>
	// <li>若签署方为企业员工，此参数无效，签名方式将以合同中为准。</li>
	// </ul>
	SignatureTypes []*int64 `json:"SignatureTypes,omitnil,omitempty" name:"SignatureTypes"`

	// 指定批量签署合同的认证校验方式，可传递以下值：
	// <ul><li>**1**：人脸认证(默认)，需进行人脸识别成功后才能签署合同</li>
	// <li>**2**：密码认证(默认)，需输入与用户在腾讯电子签设置的密码一致才能校验成功进行合同签署</li>
	// <li>**3**：运营商三要素，需到运营商处比对手机号实名信息(名字、手机号、证件号)校验一致才能成功进行合同签署。</li></ul>
	// 注：
	// <ul><li>默认情况下，认证校验方式为人脸和密码认证</li>
	// <li>您可以传递多种值，表示可用多种认证校验方式。</li></ul>
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil,omitempty" name:"ApproverSignTypes"`

	// 生成H5签署链接时，您可以指定签署方签署合同的认证校验方式的选择模式，可传递一下值：
	// <ul><li>**0**：签署方自行选择，签署方可以从预先指定的认证方式中自由选择；</li>
	// <li>**1**：自动按顺序首位推荐，签署方无需选择，系统会优先推荐使用第一种认证方式。</li></ul>
	// 注：
	// `不指定该值时，默认为签署方自行选择。`
	SignTypeSelector *uint64 `json:"SignTypeSelector,omitnil,omitempty" name:"SignTypeSelector"`

	// 批量签署合同相关信息，指定合同和签署方的信息，用于补充动态签署人。	
	// 
	// 注: `若签署方为企业员工，暂不支持通过H5端进行动态签署人的补充`
	FlowBatchUrlInfo *FlowBatchUrlInfo `json:"FlowBatchUrlInfo,omitnil,omitempty" name:"FlowBatchUrlInfo"`

	// <b>只有在生成H5签署链接的情形下</b>（ 如调用<a href="https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateFlowSignUrl" target="_blank">获取H5签署链接</a>、<a href="https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateBatchQuickSignUrl" target="_blank">获取H5批量签署链接</a>等接口），该配置才会生效。  您可以指定H5签署视频核身的意图配置，选择问答模式或点头模式的语音文本。  注意： 1. 视频认证为<b>白名单功能，使用前请联系对接的客户经理沟通</b>。 2. 使用视频认证时，<b>生成H5签署链接的时候必须将签署认证方式指定为人脸</b>（即ApproverSignTypes设置成人脸签署）。 3. 签署完成后，可以通过<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelDescribeSignFaceVideo" target="_blank">查询签署认证人脸视频</a>获取到当时的视频。
	Intention *Intention `json:"Intention,omitnil,omitempty" name:"Intention"`

	// 缓存签署人信息。在H5签署链接动态领取场景，首次填写后，选择缓存签署人信息，在下次签署人点击领取链接时，会自动将个人信息（姓名、身份证号、手机号）填入，否则需要每次手动填写。
	// 
	// 注: `若参与方为企业员工时，暂不支持对参与方信息进行缓存`
	CacheApproverInfo *bool `json:"CacheApproverInfo,omitnil,omitempty" name:"CacheApproverInfo"`

	// 是否允许此链接中签署方批量拒签。 <ul><li>false (默认): 不允许批量拒签</li> <li>true : 允许批量拒签。</li></ul>注：`合同组暂不支持批量拒签功能。`
	CanBatchReject *bool `json:"CanBatchReject,omitnil,omitempty" name:"CanBatchReject"`

	// 预设的动态签署方的补充信息，仅匹配对应信息的签署方才能领取合同。暂时仅对个人参与方生效。
	PresetApproverInfo *PresetApproverInfo `json:"PresetApproverInfo,omitnil,omitempty" name:"PresetApproverInfo"`
}

func (r *ChannelCreateBatchQuickSignUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateBatchQuickSignUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowApproverInfo")
	delete(f, "Agent")
	delete(f, "FlowIds")
	delete(f, "FlowGroupId")
	delete(f, "JumpUrl")
	delete(f, "SignatureTypes")
	delete(f, "ApproverSignTypes")
	delete(f, "SignTypeSelector")
	delete(f, "FlowBatchUrlInfo")
	delete(f, "Intention")
	delete(f, "CacheApproverInfo")
	delete(f, "CanBatchReject")
	delete(f, "PresetApproverInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateBatchQuickSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateBatchQuickSignUrlResponseParams struct {
	// 签署人签署链接信息
	FlowApproverUrlInfo *FlowApproverUrlInfo `json:"FlowApproverUrlInfo,omitnil,omitempty" name:"FlowApproverUrlInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateBatchQuickSignUrlResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateBatchQuickSignUrlResponseParams `json:"Response"`
}

func (r *ChannelCreateBatchQuickSignUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateBatchQuickSignUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateBatchSignUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：
	// <ul>
	// <li>请确保和合同中填入的一致。</li>
	// <li>在动态签署人补充链接场景中，可以通过传入这个值，对补充的个人参与方信息进行限制。仅匹配传入姓名的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方。</li>
	// </ul>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此业务通知方。
	// 
	// 注：
	// <ul>
	// <li>请确保和合同中填入的一致,  若无法保持一致，请确保在发起和生成批量签署链接时传入相同的参与方证件信息。</li><li>在生成动态签署人补充链接场景中，可以通过传入此值，对补充的个人参与方信息进行限制。仅匹配传入手机号的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方。 </li>
	// </ul>
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 证件类型，支持以下类型
	// <ul><li>**ID_CARD** : 中国大陆居民身份证 (默认值)</li>
	// <li>**HONGKONG_AND_MACAO** : 中国港澳居民来往内地通行证</li>
	// <li>**HONGKONG_MACAO_AND_TAIWAN** : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	// 
	// 注：
	// 1. `请确保和合同中填入的一致`。
	// 2. `在生成动态签署人补充链接场景中，可以通过传入此值，对补充的个人参与方信息进行限制。仅匹配传入证件类型的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方，且需要和证件号参数一同传递，不能单独进行限制。`
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	// 
	// 注：
	// 1. `请确保和合同中填入的一致`。
	// 2. `在生成动态签署人补充链接场景中，可以通过传入此值，对补充的个人参与方信息进行限制。仅匹配传入证件号的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方。`
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 通知用户方式：
	// <ul>
	// <li>**NONE** : 不通知（默认）</li>
	// <li>**SMS** : 短信通知（发送短信通知到Mobile参数所传的手机号）</li>
	// </ul>
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// 批量签署的合同流程ID数组。<font color="red">此参数必传。</font>
	// 注: `在调用此接口时，请确保合同流程均为本企业发起，且合同数量不超过100个。`
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// SaaS平台企业员工签署方的企业名称。目标签署人如果为saas应用企业员工身份，此参数必填。
	// 
	// 注：
	// <ul>
	// <li>请确认该名称与企业营业执照中注册的名称一致。</li>
	// <li>如果名称中包含英文括号()，请使用中文括号（）代替。</li>
	// <li>请确保此企业已完成腾讯电子签企业认证。</li>
	// <li>**若为子客企业员工，请使用OpenId，OrganizationOpenId参数，此参数留空即可**</li>
	// </ul>
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 指定批量签署合同的签名类型，可传递以下值：<ul><li>**0**：手写签名</li><li>**1**：OCR楷体</li><li>**2**：姓名印章</li><li>**3**：图片印章</li><li>**4**：系统签名</li><li>**5**：长效手写签名（包含手写签名）</li></ul>注：<ul><li>不传值的情况则计算所有合同中个人签署区的签名类型，规则如下：<ul><li>1.如果所有合同中所有的个人签署区方式包含多种则是手写</li><li>2.如果所有合同中所有个人签名区签名类型仅为一种则就是那一种签名方式（例如合同1有多个签署区都是指定OCR楷体，合同2中也是多个签署区都是指定OCR楷体...则使用OCR楷体）</li></ul></li><li>该参数会覆盖您合同中的签名类型，若您在发起合同时限定了签名类型(赋值签名类型给ComponentTypeLimit)，请将这些签名类型赋予此参数</li><li>若签署方为企业员工，此参数无效，签名方式将以合同中为准。</li></ul>
	SignatureTypes []*int64 `json:"SignatureTypes,omitnil,omitempty" name:"SignatureTypes"`

	// 是否直接跳转至合同内容页面进行签署
	// <ul>
	// <li>**false**: 会跳转至批量合同流程的列表,  点击需要批量签署合同后进入合同内容页面进行签署(默认)</li>
	// <li>**true**: 跳过合同流程列表, 直接进入合同内容页面进行签署</li>
	// </ul>
	JumpToDetail *bool `json:"JumpToDetail,omitnil,omitempty" name:"JumpToDetail"`

	// 批量签署合同相关信息，指定合同和签署方的信息，用于补充动态签署人。	
	FlowBatchUrlInfo *FlowBatchUrlInfo `json:"FlowBatchUrlInfo,omitnil,omitempty" name:"FlowBatchUrlInfo"`

	// 第三方平台子客企业员工的标识OpenId，批签合同经办人为子客员工的情况下为必填。
	// 
	// 注：
	// <ul>
	// <li>传入的OpenId对应员工在此子客企业下必须已经实名</li>
	// <li>传递了此参数可以无需传递Name，Mobile，IdCardNumber，IdCardType参数。系统会根据员工OpenId自动拉取实名信息。</li>
	// </ul>
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 第三方平台子客企业的企业的标识, 即OrganizationOpenId，批签合同经办人为子客企业员工是为必填。
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 签署完成后是否自动回跳
	// <ul><li>false：否, 签署完成不会自动跳转回来(默认)</li><li>true：是, 签署完成会自动跳转回来</li></ul>
	// 
	// 注: 
	// 1. 该参数<font color="red">只针对APP类型（电子签小程序跳转贵方小程序）场景</font> 的签署链接有效
	// 2. <font color="red">手机应用APP 或 微信小程序需要监控界面的返回走后序逻辑</font>, 微信小程序的文档可以参考[这个](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onShow-Object-object)
	// 3. <font color="red">电子签小程序跳转贵方APP，不支持自动跳转，必需用户手动点击完成按钮（微信的限制）</font> 
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil,omitempty" name:"AutoJumpBack"`

	// <font color="red">仅公众号 H5 跳转电子签小程序时</font>，如需签署完成的“返回应用”功能，在获取签署链接接口的 UrlUseEnv 参数需设置为 **WeChatOfficialAccounts**，小程序签署成功的结果页面中才会出现“返回应用”按钮。在用户点击“返回应用”按钮之后，会返回到公众号 H5。 
	// 
	// 参考 [公众号 H5 跳转电子签小程序](https://qian.tencent.com/developers/company/openwxminiprogram/#23-%E5%85%AC%E4%BC%97%E5%8F%B7-h5-%E4%B8%AD%E8%B7%B3%E8%BD%AC)。
	UrlUseEnv *string `json:"UrlUseEnv,omitnil,omitempty" name:"UrlUseEnv"`

	// 是否允许此链接中签署方批量拒签。 <ul><li>false (默认): 不允许批量拒签</li> <li>true : 允许批量拒签。</li></ul>
	// 注：`1. 合同组暂不支持批量拒签功能。2. 如果是链接直接跳转至详情页（JumpToDetail参数为true），也不支持批量拒签功能`
	CanBatchReject *bool `json:"CanBatchReject,omitnil,omitempty" name:"CanBatchReject"`

	// 是否允许此链接中签署方批量确认已读文件。 <ul><li>false (默认): 不允许批量确认已读文件。</li> <li>true : 允许批量确认已读文件。</li></ul>
	// 注：`1. 此功能为白名单功能，使用前请联系对应客户经理进行开通。2. 使用此功能时，FlowIds参数必传。3. 对于企业签署方，如果对印章/签名控件有限制要求，需要保证所有印章/签名控件的限制要求(印章id或印章/签名类型限制)一致，否则无法使用此功能。`
	CanSkipReadFlow *bool `json:"CanSkipReadFlow,omitnil,omitempty" name:"CanSkipReadFlow"`
}

type ChannelCreateBatchSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：
	// <ul>
	// <li>请确保和合同中填入的一致。</li>
	// <li>在动态签署人补充链接场景中，可以通过传入这个值，对补充的个人参与方信息进行限制。仅匹配传入姓名的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方。</li>
	// </ul>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此业务通知方。
	// 
	// 注：
	// <ul>
	// <li>请确保和合同中填入的一致,  若无法保持一致，请确保在发起和生成批量签署链接时传入相同的参与方证件信息。</li><li>在生成动态签署人补充链接场景中，可以通过传入此值，对补充的个人参与方信息进行限制。仅匹配传入手机号的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方。 </li>
	// </ul>
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 证件类型，支持以下类型
	// <ul><li>**ID_CARD** : 中国大陆居民身份证 (默认值)</li>
	// <li>**HONGKONG_AND_MACAO** : 中国港澳居民来往内地通行证</li>
	// <li>**HONGKONG_MACAO_AND_TAIWAN** : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	// 
	// 注：
	// 1. `请确保和合同中填入的一致`。
	// 2. `在生成动态签署人补充链接场景中，可以通过传入此值，对补充的个人参与方信息进行限制。仅匹配传入证件类型的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方，且需要和证件号参数一同传递，不能单独进行限制。`
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	// 
	// 注：
	// 1. `请确保和合同中填入的一致`。
	// 2. `在生成动态签署人补充链接场景中，可以通过传入此值，对补充的个人参与方信息进行限制。仅匹配传入证件号的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方。`
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 通知用户方式：
	// <ul>
	// <li>**NONE** : 不通知（默认）</li>
	// <li>**SMS** : 短信通知（发送短信通知到Mobile参数所传的手机号）</li>
	// </ul>
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// 批量签署的合同流程ID数组。<font color="red">此参数必传。</font>
	// 注: `在调用此接口时，请确保合同流程均为本企业发起，且合同数量不超过100个。`
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// SaaS平台企业员工签署方的企业名称。目标签署人如果为saas应用企业员工身份，此参数必填。
	// 
	// 注：
	// <ul>
	// <li>请确认该名称与企业营业执照中注册的名称一致。</li>
	// <li>如果名称中包含英文括号()，请使用中文括号（）代替。</li>
	// <li>请确保此企业已完成腾讯电子签企业认证。</li>
	// <li>**若为子客企业员工，请使用OpenId，OrganizationOpenId参数，此参数留空即可**</li>
	// </ul>
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 指定批量签署合同的签名类型，可传递以下值：<ul><li>**0**：手写签名</li><li>**1**：OCR楷体</li><li>**2**：姓名印章</li><li>**3**：图片印章</li><li>**4**：系统签名</li><li>**5**：长效手写签名（包含手写签名）</li></ul>注：<ul><li>不传值的情况则计算所有合同中个人签署区的签名类型，规则如下：<ul><li>1.如果所有合同中所有的个人签署区方式包含多种则是手写</li><li>2.如果所有合同中所有个人签名区签名类型仅为一种则就是那一种签名方式（例如合同1有多个签署区都是指定OCR楷体，合同2中也是多个签署区都是指定OCR楷体...则使用OCR楷体）</li></ul></li><li>该参数会覆盖您合同中的签名类型，若您在发起合同时限定了签名类型(赋值签名类型给ComponentTypeLimit)，请将这些签名类型赋予此参数</li><li>若签署方为企业员工，此参数无效，签名方式将以合同中为准。</li></ul>
	SignatureTypes []*int64 `json:"SignatureTypes,omitnil,omitempty" name:"SignatureTypes"`

	// 是否直接跳转至合同内容页面进行签署
	// <ul>
	// <li>**false**: 会跳转至批量合同流程的列表,  点击需要批量签署合同后进入合同内容页面进行签署(默认)</li>
	// <li>**true**: 跳过合同流程列表, 直接进入合同内容页面进行签署</li>
	// </ul>
	JumpToDetail *bool `json:"JumpToDetail,omitnil,omitempty" name:"JumpToDetail"`

	// 批量签署合同相关信息，指定合同和签署方的信息，用于补充动态签署人。	
	FlowBatchUrlInfo *FlowBatchUrlInfo `json:"FlowBatchUrlInfo,omitnil,omitempty" name:"FlowBatchUrlInfo"`

	// 第三方平台子客企业员工的标识OpenId，批签合同经办人为子客员工的情况下为必填。
	// 
	// 注：
	// <ul>
	// <li>传入的OpenId对应员工在此子客企业下必须已经实名</li>
	// <li>传递了此参数可以无需传递Name，Mobile，IdCardNumber，IdCardType参数。系统会根据员工OpenId自动拉取实名信息。</li>
	// </ul>
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 第三方平台子客企业的企业的标识, 即OrganizationOpenId，批签合同经办人为子客企业员工是为必填。
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 签署完成后是否自动回跳
	// <ul><li>false：否, 签署完成不会自动跳转回来(默认)</li><li>true：是, 签署完成会自动跳转回来</li></ul>
	// 
	// 注: 
	// 1. 该参数<font color="red">只针对APP类型（电子签小程序跳转贵方小程序）场景</font> 的签署链接有效
	// 2. <font color="red">手机应用APP 或 微信小程序需要监控界面的返回走后序逻辑</font>, 微信小程序的文档可以参考[这个](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onShow-Object-object)
	// 3. <font color="red">电子签小程序跳转贵方APP，不支持自动跳转，必需用户手动点击完成按钮（微信的限制）</font> 
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil,omitempty" name:"AutoJumpBack"`

	// <font color="red">仅公众号 H5 跳转电子签小程序时</font>，如需签署完成的“返回应用”功能，在获取签署链接接口的 UrlUseEnv 参数需设置为 **WeChatOfficialAccounts**，小程序签署成功的结果页面中才会出现“返回应用”按钮。在用户点击“返回应用”按钮之后，会返回到公众号 H5。 
	// 
	// 参考 [公众号 H5 跳转电子签小程序](https://qian.tencent.com/developers/company/openwxminiprogram/#23-%E5%85%AC%E4%BC%97%E5%8F%B7-h5-%E4%B8%AD%E8%B7%B3%E8%BD%AC)。
	UrlUseEnv *string `json:"UrlUseEnv,omitnil,omitempty" name:"UrlUseEnv"`

	// 是否允许此链接中签署方批量拒签。 <ul><li>false (默认): 不允许批量拒签</li> <li>true : 允许批量拒签。</li></ul>
	// 注：`1. 合同组暂不支持批量拒签功能。2. 如果是链接直接跳转至详情页（JumpToDetail参数为true），也不支持批量拒签功能`
	CanBatchReject *bool `json:"CanBatchReject,omitnil,omitempty" name:"CanBatchReject"`

	// 是否允许此链接中签署方批量确认已读文件。 <ul><li>false (默认): 不允许批量确认已读文件。</li> <li>true : 允许批量确认已读文件。</li></ul>
	// 注：`1. 此功能为白名单功能，使用前请联系对应客户经理进行开通。2. 使用此功能时，FlowIds参数必传。3. 对于企业签署方，如果对印章/签名控件有限制要求，需要保证所有印章/签名控件的限制要求(印章id或印章/签名类型限制)一致，否则无法使用此功能。`
	CanSkipReadFlow *bool `json:"CanSkipReadFlow,omitnil,omitempty" name:"CanSkipReadFlow"`
}

func (r *ChannelCreateBatchSignUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateBatchSignUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "Name")
	delete(f, "Mobile")
	delete(f, "Operator")
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "NotifyType")
	delete(f, "FlowIds")
	delete(f, "OrganizationName")
	delete(f, "SignatureTypes")
	delete(f, "JumpToDetail")
	delete(f, "FlowBatchUrlInfo")
	delete(f, "OpenId")
	delete(f, "OrganizationOpenId")
	delete(f, "AutoJumpBack")
	delete(f, "UrlUseEnv")
	delete(f, "CanBatchReject")
	delete(f, "CanSkipReadFlow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateBatchSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateBatchSignUrlResponseParams struct {
	// 批量签署链接，以短链形式返回，短链的有效期参考回参中的 ExpiredTime。
	// 
	// 注: 
	// 1. 非小程序和APP集成使用
	// 2. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	SignUrl *string `json:"SignUrl,omitnil,omitempty" name:"SignUrl"`

	// 链接过期时间以 Unix 时间戳格式表示，从生成链接时间起，往后7天有效期。过期后短链将失效，无法打开。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 从客户小程序或者客户APP跳转至腾讯电子签小程序进行批量签署的跳转路径
	// 
	// 注: 
	// 1. 小程序和APP集成使用
	// 2. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	MiniAppPath *string `json:"MiniAppPath,omitnil,omitempty" name:"MiniAppPath"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateBatchSignUrlResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateBatchSignUrlResponseParams `json:"Response"`
}

func (r *ChannelCreateBatchSignUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateBatchSignUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateBoundFlowsRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证,  合同会领取给对应的Agent.ProxyOperator.OpenId指定的员工来处理
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要领取的合同流程的ID列表
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelCreateBoundFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证,  合同会领取给对应的Agent.ProxyOperator.OpenId指定的员工来处理
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要领取的合同流程的ID列表
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelCreateBoundFlowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateBoundFlowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowIds")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateBoundFlowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateBoundFlowsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateBoundFlowsResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateBoundFlowsResponseParams `json:"Response"`
}

func (r *ChannelCreateBoundFlowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateBoundFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateConvertTaskApiRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要进行转换的资源文件类型
	// 支持的文件类型如下：
	// <ul><li>doc</li>
	// <li>docx</li>
	// <li>xls</li>
	// <li>xlsx</li>
	// <li>jpg</li>
	// <li>jpeg</li>
	// <li>png</li>
	// <li>bmp</li>
	// <li>html</li>
	// <li>txt</li></ul>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 需要进行转换操作的文件资源名称，带资源后缀名。
	// 
	// 注:  `资源名称长度限制为256个字符`
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 需要进行转换操作的文件资源Id，通过<a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">UploadFiles</a>接口获取文件资源Id。
	// 
	// 注:  `目前，此接口仅支持单个文件进行转换。`
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 调用方用户信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
}

type ChannelCreateConvertTaskApiRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要进行转换的资源文件类型
	// 支持的文件类型如下：
	// <ul><li>doc</li>
	// <li>docx</li>
	// <li>xls</li>
	// <li>xlsx</li>
	// <li>jpg</li>
	// <li>jpeg</li>
	// <li>png</li>
	// <li>bmp</li>
	// <li>html</li>
	// <li>txt</li></ul>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 需要进行转换操作的文件资源名称，带资源后缀名。
	// 
	// 注:  `资源名称长度限制为256个字符`
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 需要进行转换操作的文件资源Id，通过<a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">UploadFiles</a>接口获取文件资源Id。
	// 
	// 注:  `目前，此接口仅支持单个文件进行转换。`
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 调用方用户信息，不用传
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
}

func (r *ChannelCreateConvertTaskApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateConvertTaskApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "ResourceType")
	delete(f, "ResourceName")
	delete(f, "ResourceId")
	delete(f, "Operator")
	delete(f, "Organization")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateConvertTaskApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateConvertTaskApiResponseParams struct {
	// 接口返回的文件转换任务Id，可以调用接口<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi" target="_blank">查询转换任务状态</a>获取转换任务的状态和转换后的文件资源Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateConvertTaskApiResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateConvertTaskApiResponseParams `json:"Response"`
}

func (r *ChannelCreateConvertTaskApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateConvertTaskApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateDynamicFlowApproverRequestParams struct {
	// 动态合同信息
	FillDynamicFlowList []*DynamicFlowInfo `json:"FillDynamicFlowList,omitnil,omitempty" name:"FillDynamicFlowList"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。 此接口下面信息必填。 <ul> <li>渠道应用标识: Agent.AppId</li> <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li> <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li> </ul> 第三方平台子客企业和员工必须已经经过实名认证	
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type ChannelCreateDynamicFlowApproverRequest struct {
	*tchttp.BaseRequest
	
	// 动态合同信息
	FillDynamicFlowList []*DynamicFlowInfo `json:"FillDynamicFlowList,omitnil,omitempty" name:"FillDynamicFlowList"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。 此接口下面信息必填。 <ul> <li>渠道应用标识: Agent.AppId</li> <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li> <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li> </ul> 第三方平台子客企业和员工必须已经经过实名认证	
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *ChannelCreateDynamicFlowApproverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateDynamicFlowApproverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FillDynamicFlowList")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateDynamicFlowApproverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateDynamicFlowApproverResponseParams struct {
	// 动态合同补充结果列表
	DynamicFlowResultList []*DynamicFlowResult `json:"DynamicFlowResultList,omitnil,omitempty" name:"DynamicFlowResultList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateDynamicFlowApproverResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateDynamicFlowApproverResponseParams `json:"Response"`
}

func (r *ChannelCreateDynamicFlowApproverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateDynamicFlowApproverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateEmbedWebUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要生成WEB嵌入界面的类型, 可以选择的值如下: 
	// 
	// <ul>
	// <li>CREATE_SEAL: 生成创建印章的嵌入页面</li>
	// <li>CREATE_TEMPLATE：生成创建模板的嵌入页面</li>
	// <li>MODIFY_TEMPLATE：生成修改模板的嵌入页面</li>
	// <li>PREVIEW_TEMPLATE：生成预览模板的嵌入页面</li>
	// <li>PREVIEW_FLOW：生成预览合同文档的嵌入页面（H5链接，支持移动端的浏览器中打开）</li>
	// <li>PREVIEW_FLOW_DETAIL：生成预览合同详情的嵌入页面（仅支持PC的浏览器中打开）</li>
	// <li>PREVIEW_SEAL_LIST：生成预览印章列表的嵌入页面</li>
	// <li>PREVIEW_SEAL_DETAIL：生成预览印章详情的嵌入页面</li>
	// <li>EXTEND_SERVICE：生成扩展服务的嵌入页面</li>
	// </ul>
	EmbedType *string `json:"EmbedType,omitnil,omitempty" name:"EmbedType"`

	// WEB嵌入的业务资源ID
	// 
	// 当EmbedType取值
	// <ul>
	// <li>为MODIFY_TEMPLATE，PREVIEW_TEMPLATE必填，取值为模板id</li>
	// <li>为CREATE_TEMPLATE，非必填，取值为资源id。*资源Id获取可使用接口[上传文件](https://qian.tencent.com/developers/partnerApis/files/UploadFiles)*</li>
	// <li>为PREVIEW_FLOW，PREVIEW_FLOW_DETAIL必填，取值为合同id</li>
	// <li>为PREVIEW_SEAL_DETAIL必填，取值为印章id</li>
	// </ul>
	// 
	// 
	// 注意：
	//  1. CREATE_TEMPLATE中的BusinessId仅支持PDF文件类型， 如果您的文件不是PDF， 请使用接口[创建文件转换任务](https://qian.tencent.com/developers/partnerApis/files/ChannelCreateConvertTaskApi) 和[查询转换任务状态](https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi) 来进行转换成PDF资源。
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 是否隐藏控件，只有预览模板时生效
	HiddenComponents *bool `json:"HiddenComponents,omitnil,omitempty" name:"HiddenComponents"`

	// 渠道操作者信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 用户自定义参数
	// <ul>
	// <li>目前仅支持EmbedType=CREATE_TEMPLATE时传入</li>
	// <li>指定后，创建，编辑，删除模板时，回调都会携带该userData</li>
	// <li>支持的格式：json字符串的BASE64编码字符串</li>
	// <li>示例：<ul>
	//                  <li>json字符串：{"ComeFrom":"xxx"}，BASE64编码：eyJDb21lRnJvbSI6Inh4eCJ9</li>
	//                  <li>eyJDb21lRnJvbSI6Inh4eCJ9，为符合要求的userData数据格式</li>
	// </ul>
	// </li>
	// </ul>
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 个性化参数，用于控制页面展示内容
	Option *EmbedUrlOption `json:"Option,omitnil,omitempty" name:"Option"`
}

type ChannelCreateEmbedWebUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要生成WEB嵌入界面的类型, 可以选择的值如下: 
	// 
	// <ul>
	// <li>CREATE_SEAL: 生成创建印章的嵌入页面</li>
	// <li>CREATE_TEMPLATE：生成创建模板的嵌入页面</li>
	// <li>MODIFY_TEMPLATE：生成修改模板的嵌入页面</li>
	// <li>PREVIEW_TEMPLATE：生成预览模板的嵌入页面</li>
	// <li>PREVIEW_FLOW：生成预览合同文档的嵌入页面（H5链接，支持移动端的浏览器中打开）</li>
	// <li>PREVIEW_FLOW_DETAIL：生成预览合同详情的嵌入页面（仅支持PC的浏览器中打开）</li>
	// <li>PREVIEW_SEAL_LIST：生成预览印章列表的嵌入页面</li>
	// <li>PREVIEW_SEAL_DETAIL：生成预览印章详情的嵌入页面</li>
	// <li>EXTEND_SERVICE：生成扩展服务的嵌入页面</li>
	// </ul>
	EmbedType *string `json:"EmbedType,omitnil,omitempty" name:"EmbedType"`

	// WEB嵌入的业务资源ID
	// 
	// 当EmbedType取值
	// <ul>
	// <li>为MODIFY_TEMPLATE，PREVIEW_TEMPLATE必填，取值为模板id</li>
	// <li>为CREATE_TEMPLATE，非必填，取值为资源id。*资源Id获取可使用接口[上传文件](https://qian.tencent.com/developers/partnerApis/files/UploadFiles)*</li>
	// <li>为PREVIEW_FLOW，PREVIEW_FLOW_DETAIL必填，取值为合同id</li>
	// <li>为PREVIEW_SEAL_DETAIL必填，取值为印章id</li>
	// </ul>
	// 
	// 
	// 注意：
	//  1. CREATE_TEMPLATE中的BusinessId仅支持PDF文件类型， 如果您的文件不是PDF， 请使用接口[创建文件转换任务](https://qian.tencent.com/developers/partnerApis/files/ChannelCreateConvertTaskApi) 和[查询转换任务状态](https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi) 来进行转换成PDF资源。
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 是否隐藏控件，只有预览模板时生效
	HiddenComponents *bool `json:"HiddenComponents,omitnil,omitempty" name:"HiddenComponents"`

	// 渠道操作者信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 用户自定义参数
	// <ul>
	// <li>目前仅支持EmbedType=CREATE_TEMPLATE时传入</li>
	// <li>指定后，创建，编辑，删除模板时，回调都会携带该userData</li>
	// <li>支持的格式：json字符串的BASE64编码字符串</li>
	// <li>示例：<ul>
	//                  <li>json字符串：{"ComeFrom":"xxx"}，BASE64编码：eyJDb21lRnJvbSI6Inh4eCJ9</li>
	//                  <li>eyJDb21lRnJvbSI6Inh4eCJ9，为符合要求的userData数据格式</li>
	// </ul>
	// </li>
	// </ul>
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 个性化参数，用于控制页面展示内容
	Option *EmbedUrlOption `json:"Option,omitnil,omitempty" name:"Option"`
}

func (r *ChannelCreateEmbedWebUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateEmbedWebUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "EmbedType")
	delete(f, "BusinessId")
	delete(f, "HiddenComponents")
	delete(f, "Operator")
	delete(f, "UserData")
	delete(f, "Option")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateEmbedWebUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateEmbedWebUrlResponseParams struct {
	// 嵌入的web链接，5分钟有效
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateEmbedWebUrlResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateEmbedWebUrlResponseParams `json:"Response"`
}

func (r *ChannelCreateEmbedWebUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateEmbedWebUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowApproversRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 补充企业签署人信息。
	// 
	// - 如果发起方指定的补充签署人是企业签署人，则需要提供企业名称或者企业OpenId；
	// 
	// - 如果不指定，则使用姓名和手机号进行补充。
	Approvers []*FillApproverInfo `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 合同流程ID，为32位字符串。 
	// - 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// - 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	// - <font color="red">不建议继续使用</font>，请使用<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId指定合同
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署人信息补充方式
	// 
	// <ul><li>**1**: 表示往未指定签署人的节点，添加一个明确的签署人，支持企业或个人签署方。</li></ul>
	FillApproverType *int64 `json:"FillApproverType,omitnil,omitempty" name:"FillApproverType"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程组的组ID, 在合同流程组场景下，生成合同流程组的签署链接时需要赋值
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`
}

type ChannelCreateFlowApproversRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 补充企业签署人信息。
	// 
	// - 如果发起方指定的补充签署人是企业签署人，则需要提供企业名称或者企业OpenId；
	// 
	// - 如果不指定，则使用姓名和手机号进行补充。
	Approvers []*FillApproverInfo `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 合同流程ID，为32位字符串。 
	// - 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// - 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	// - <font color="red">不建议继续使用</font>，请使用<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId指定合同
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署人信息补充方式
	// 
	// <ul><li>**1**: 表示往未指定签署人的节点，添加一个明确的签署人，支持企业或个人签署方。</li></ul>
	FillApproverType *int64 `json:"FillApproverType,omitnil,omitempty" name:"FillApproverType"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程组的组ID, 在合同流程组场景下，生成合同流程组的签署链接时需要赋值
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`
}

func (r *ChannelCreateFlowApproversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowApproversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "Approvers")
	delete(f, "FlowId")
	delete(f, "FillApproverType")
	delete(f, "Operator")
	delete(f, "FlowGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowApproversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowApproversResponseParams struct {
	// 批量补充签署人时，补充失败的报错说明 
	// 注:`目前仅补充动态签署人时会返回补充失败的原因`	
	FillError []*FillError `json:"FillError,omitnil,omitempty" name:"FillError"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateFlowApproversResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateFlowApproversResponseParams `json:"Response"`
}

func (r *ChannelCreateFlowApproversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowApproversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowByFilesRequestParams struct {
	// 合同的发起企业和发起人信息，<a href="https://qcloudimg.tencent-cloud.cn/raw/b69f8aad306c40b7b78d096e39b2edbb.png" target="_blank">点击查看合同发起企业和人展示的位置</a>
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId（合同的发起企业）</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId （合同的发起人）</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil,omitempty" name:"FlowDescription"`

	// 合同流程的参与方列表, 最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，不同类型的签署方传参方式可以参考文档 [签署方入参指引](https://qian.tencent.com/developers/partner/flow_approver)。
	// 
	// 如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序, 请确保列表中参与人的顺序符合实际签署顺序。
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitnil,omitempty" name:"FlowApprovers"`

	// 本合同流程需包含的PDF文件资源编号列表，通过<a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">UploadFiles</a>接口获取PDF文件资源编号。
	// 
	// 注: `目前，此接口仅支持单个文件发起。`
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 模板或者合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li>单行文本控件</li>
	// <li>多行文本控件</li>
	// <li>勾选框控件</li>
	// <li>数字控件</li>
	// <li>图片控件</li>
	// <li>数据表格等填写控件</li></ul>
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/e004195ee4cb98a7f9bc12eb4a0a0b77.png)
	Components []*Component `json:"Components,omitnil,omitempty" name:"Components"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 该字段已废弃，请使用【应用号配置】中的回调地址
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	// **注**: `有序签署时以传入FlowApprovers数组的顺序作为签署顺序`
	Unordered *bool `json:"Unordered,omitnil,omitempty" name:"Unordered"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为255个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 您可以自定义腾讯电子签小程序合同列表页展示的合同内容模板，模板中支持以下变量：
	// <ul><li>{合同名称}   </li>
	// <li>{发起方企业} </li>
	// <li>{发起方姓名} </li>
	// <li>{签署方N企业}</li>
	// <li>{签署方N姓名}</li></ul>
	// 其中，N表示签署方的编号，从1开始，不能超过签署人的数量。
	// 
	// 例如，如果是腾讯公司张三发给李四名称为“租房合同”的合同，您可以将此字段设置为：`合同名称:{合同名称};发起方: {发起方企业}({发起方姓名});签署方:{签署方1姓名}`，则小程序中列表页展示此合同为以下样子
	// 
	// 合同名称：租房合同 
	// 发起方：腾讯公司(张三) 
	// 签署方：李四
	// 
	CustomShowMap *string `json:"CustomShowMap,omitnil,omitempty" name:"CustomShowMap"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 1000长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/partner/callback_types_contracts_sign" target="_blank">回调通知</a>模块。
	CustomerData *string `json:"CustomerData,omitnil,omitempty" name:"CustomerData"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过ChannelCreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil,omitempty" name:"NeedSignReview"`

	// 签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证，用户手机号和参与方手机号（ApproverMobile）相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// 参数说明：可选人脸识别或手机号验证两种方式，若选择后者，未实名个人签署方在签署合同时，无需经过实名认证和意愿确认两次人脸识别，该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil,omitempty" name:"ApproverVerifyType"`

	// 签署方签署控件（印章/签名等）的生成方式：
	// <ul><li> **0**：在合同流程发起时，由发起人指定签署方的签署控件的位置和数量。</li>
	// <li> **1**：签署方在签署时自行添加签署控件，可以拖动位置和控制数量。</li></ul>
	// 
	// **注**: 
	// 1.发起后添加控件功能不支持添加签批控件 
	// 2.签署方在签署时自行添加签署控件仅支持电子签小程序或web控制台签署，不支持H5
	SignBeanTag *int64 `json:"SignBeanTag,omitnil,omitempty" name:"SignBeanTag"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// <b>注</b>
	// 1. 抄送人名单中可以包括自然人以及本企业的员工（本企业员工必须已经完成认证并加入企业）。
	// 2. 请确保抄送人列表中的成员不与任何签署人重复。
	CcInfos []*CcInfo `json:"CcInfos,omitnil,omitempty" name:"CcInfos"`

	// 可以设置以下时间节点来给抄送人发送短信通知来查看合同内容：
	// <ul><li> **0**：合同发起时通知（默认值）</li>
	// <li> **1**：签署完成后通知</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil,omitempty" name:"CcNotifyType"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：电子处方单（医疗自动签）  </li><li> **OTHER** :  通用场景</li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil,omitempty" name:"AutoSignScene"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下：  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li><li> <b>3</b> :文书</li></ul>效果如下:![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"FlowDisplayType,omitnil,omitempty" name:"FlowDisplayType"`

	// 是否为预览模式，取值如下： <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li> <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。</li></ul>
	NeedPreview *bool `json:"NeedPreview,omitnil,omitempty" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子</li></ul>
	// 注: `此参数在NeedPreview 为true时有效`
	PreviewType *int64 `json:"PreviewType,omitnil,omitempty" name:"PreviewType"`

	// 是否开启动态合同（动态签署人2.0）
	// <ul><li> **false** :(默认) 不开启动态合同（动态签署人2.0）</li>
	// <li> **true** :开启动态合同（动态签署人2.0）,发起后可继续追加合同签署人</li></ul>
	//
	// Deprecated: OpenDynamicFlow is deprecated.
	OpenDynamicFlow *bool `json:"OpenDynamicFlow,omitnil,omitempty" name:"OpenDynamicFlow"`

	// 是否开启动态合同（动态签署人2.0）<ul><li> **false** :(默认) 不开启动态合同（动态签署人2.0）</li><li> **true** :开启动态合同（动态签署人2.0）,发起后可继续追加合同签署人</li></ul>
	OpenDynamicSignFlow *bool `json:"OpenDynamicSignFlow,omitnil,omitempty" name:"OpenDynamicSignFlow"`
}

type ChannelCreateFlowByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 合同的发起企业和发起人信息，<a href="https://qcloudimg.tencent-cloud.cn/raw/b69f8aad306c40b7b78d096e39b2edbb.png" target="_blank">点击查看合同发起企业和人展示的位置</a>
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId（合同的发起企业）</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId （合同的发起人）</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil,omitempty" name:"FlowDescription"`

	// 合同流程的参与方列表, 最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，不同类型的签署方传参方式可以参考文档 [签署方入参指引](https://qian.tencent.com/developers/partner/flow_approver)。
	// 
	// 如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序, 请确保列表中参与人的顺序符合实际签署顺序。
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitnil,omitempty" name:"FlowApprovers"`

	// 本合同流程需包含的PDF文件资源编号列表，通过<a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">UploadFiles</a>接口获取PDF文件资源编号。
	// 
	// 注: `目前，此接口仅支持单个文件发起。`
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 模板或者合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li>单行文本控件</li>
	// <li>多行文本控件</li>
	// <li>勾选框控件</li>
	// <li>数字控件</li>
	// <li>图片控件</li>
	// <li>数据表格等填写控件</li></ul>
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/e004195ee4cb98a7f9bc12eb4a0a0b77.png)
	Components []*Component `json:"Components,omitnil,omitempty" name:"Components"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 该字段已废弃，请使用【应用号配置】中的回调地址
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	// **注**: `有序签署时以传入FlowApprovers数组的顺序作为签署顺序`
	Unordered *bool `json:"Unordered,omitnil,omitempty" name:"Unordered"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为255个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 您可以自定义腾讯电子签小程序合同列表页展示的合同内容模板，模板中支持以下变量：
	// <ul><li>{合同名称}   </li>
	// <li>{发起方企业} </li>
	// <li>{发起方姓名} </li>
	// <li>{签署方N企业}</li>
	// <li>{签署方N姓名}</li></ul>
	// 其中，N表示签署方的编号，从1开始，不能超过签署人的数量。
	// 
	// 例如，如果是腾讯公司张三发给李四名称为“租房合同”的合同，您可以将此字段设置为：`合同名称:{合同名称};发起方: {发起方企业}({发起方姓名});签署方:{签署方1姓名}`，则小程序中列表页展示此合同为以下样子
	// 
	// 合同名称：租房合同 
	// 发起方：腾讯公司(张三) 
	// 签署方：李四
	// 
	CustomShowMap *string `json:"CustomShowMap,omitnil,omitempty" name:"CustomShowMap"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 1000长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/partner/callback_types_contracts_sign" target="_blank">回调通知</a>模块。
	CustomerData *string `json:"CustomerData,omitnil,omitempty" name:"CustomerData"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过ChannelCreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil,omitempty" name:"NeedSignReview"`

	// 签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证，用户手机号和参与方手机号（ApproverMobile）相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// 参数说明：可选人脸识别或手机号验证两种方式，若选择后者，未实名个人签署方在签署合同时，无需经过实名认证和意愿确认两次人脸识别，该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil,omitempty" name:"ApproverVerifyType"`

	// 签署方签署控件（印章/签名等）的生成方式：
	// <ul><li> **0**：在合同流程发起时，由发起人指定签署方的签署控件的位置和数量。</li>
	// <li> **1**：签署方在签署时自行添加签署控件，可以拖动位置和控制数量。</li></ul>
	// 
	// **注**: 
	// 1.发起后添加控件功能不支持添加签批控件 
	// 2.签署方在签署时自行添加签署控件仅支持电子签小程序或web控制台签署，不支持H5
	SignBeanTag *int64 `json:"SignBeanTag,omitnil,omitempty" name:"SignBeanTag"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// <b>注</b>
	// 1. 抄送人名单中可以包括自然人以及本企业的员工（本企业员工必须已经完成认证并加入企业）。
	// 2. 请确保抄送人列表中的成员不与任何签署人重复。
	CcInfos []*CcInfo `json:"CcInfos,omitnil,omitempty" name:"CcInfos"`

	// 可以设置以下时间节点来给抄送人发送短信通知来查看合同内容：
	// <ul><li> **0**：合同发起时通知（默认值）</li>
	// <li> **1**：签署完成后通知</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil,omitempty" name:"CcNotifyType"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：电子处方单（医疗自动签）  </li><li> **OTHER** :  通用场景</li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil,omitempty" name:"AutoSignScene"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下：  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li><li> <b>3</b> :文书</li></ul>效果如下:![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"FlowDisplayType,omitnil,omitempty" name:"FlowDisplayType"`

	// 是否为预览模式，取值如下： <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li> <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。</li></ul>
	NeedPreview *bool `json:"NeedPreview,omitnil,omitempty" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子</li></ul>
	// 注: `此参数在NeedPreview 为true时有效`
	PreviewType *int64 `json:"PreviewType,omitnil,omitempty" name:"PreviewType"`

	// 是否开启动态合同（动态签署人2.0）
	// <ul><li> **false** :(默认) 不开启动态合同（动态签署人2.0）</li>
	// <li> **true** :开启动态合同（动态签署人2.0）,发起后可继续追加合同签署人</li></ul>
	OpenDynamicFlow *bool `json:"OpenDynamicFlow,omitnil,omitempty" name:"OpenDynamicFlow"`

	// 是否开启动态合同（动态签署人2.0）<ul><li> **false** :(默认) 不开启动态合同（动态签署人2.0）</li><li> **true** :开启动态合同（动态签署人2.0）,发起后可继续追加合同签署人</li></ul>
	OpenDynamicSignFlow *bool `json:"OpenDynamicSignFlow,omitnil,omitempty" name:"OpenDynamicSignFlow"`
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
	delete(f, "FlowDescription")
	delete(f, "FlowApprovers")
	delete(f, "FileIds")
	delete(f, "Components")
	delete(f, "Deadline")
	delete(f, "CallbackUrl")
	delete(f, "Unordered")
	delete(f, "FlowType")
	delete(f, "CustomShowMap")
	delete(f, "CustomerData")
	delete(f, "NeedSignReview")
	delete(f, "ApproverVerifyType")
	delete(f, "SignBeanTag")
	delete(f, "CcInfos")
	delete(f, "CcNotifyType")
	delete(f, "AutoSignScene")
	delete(f, "Operator")
	delete(f, "FlowDisplayType")
	delete(f, "NeedPreview")
	delete(f, "PreviewType")
	delete(f, "OpenDynamicFlow")
	delete(f, "OpenDynamicSignFlow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowByFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowByFilesResponseParams struct {
	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 
	// [点击查看FlowId在控制台上的位置](https://qcloudimg.tencent-cloud.cn/raw/05af26573d5106763b4cfbb9f7c64b41.png)
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署方信息，如角色ID、角色名称等
	Approvers []*ApproverItem `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 预览链接，有效期5分钟
	// 注：如果是预览模式(即NeedPreview设置为true)时, 才会有此预览链接URL
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ChannelCreateFlowGroupByFilesRequestParams struct {
	// 合同组中每个合同签署流程的信息，合同组中最少包含2个合同，不能超过50个合同。
	FlowFileInfos []*FlowFileInfo `json:"FlowFileInfos,omitnil,omitempty" name:"FlowFileInfos"`

	// 合同组的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil,omitempty" name:"FlowGroupName"`

	// 合同的发起企业和发起人信息，<a href="https://qcloudimg.tencent-cloud.cn/raw/b69f8aad306c40b7b78d096e39b2edbb.png" target="_blank">点击查看合同发起企业和人展示的位置</a>
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识: <a href="https://qcloudimg.tencent-cloud.cn/raw/a71872de3d540d55451e3e73a2ad1a6e.png" target="_blank">Agent.AppId</a></li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId（合同的发起企业）</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId （合同的发起人）</li>
	// </ul>
	// 
	// 合同的发起企业和发起人必需已经完成实名，并加入企业
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同组中签署人校验和认证的方式：
	// <ul><li>**VerifyCheck**：人脸识别（默认）</li>
	// <li>**MobileCheck**：手机号验证</li></ul>
	// 注意：
	// `1. MobileCheck 方式，未实名的个人/自然人签署方无需进行人脸识别实名认证即可查看合同（但签署合同时仍然需要人脸实名），企业签署方需经过人脸认证。`
	// `2. 合同组的校验和认证的方式会优先使用，会覆盖合同组中单个合同和合同签署方认证方式的限制配置。`
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil,omitempty" name:"ApproverVerifyType"`

	// 合同组的签署配置项信息，例如在合同组签署过程中，是否需要对每个子合同进行独立的意愿确认。
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil,omitempty" name:"FlowGroupOptions"`

	// 操作者的信息，此参数不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelCreateFlowGroupByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 合同组中每个合同签署流程的信息，合同组中最少包含2个合同，不能超过50个合同。
	FlowFileInfos []*FlowFileInfo `json:"FlowFileInfos,omitnil,omitempty" name:"FlowFileInfos"`

	// 合同组的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil,omitempty" name:"FlowGroupName"`

	// 合同的发起企业和发起人信息，<a href="https://qcloudimg.tencent-cloud.cn/raw/b69f8aad306c40b7b78d096e39b2edbb.png" target="_blank">点击查看合同发起企业和人展示的位置</a>
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识: <a href="https://qcloudimg.tencent-cloud.cn/raw/a71872de3d540d55451e3e73a2ad1a6e.png" target="_blank">Agent.AppId</a></li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId（合同的发起企业）</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId （合同的发起人）</li>
	// </ul>
	// 
	// 合同的发起企业和发起人必需已经完成实名，并加入企业
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同组中签署人校验和认证的方式：
	// <ul><li>**VerifyCheck**：人脸识别（默认）</li>
	// <li>**MobileCheck**：手机号验证</li></ul>
	// 注意：
	// `1. MobileCheck 方式，未实名的个人/自然人签署方无需进行人脸识别实名认证即可查看合同（但签署合同时仍然需要人脸实名），企业签署方需经过人脸认证。`
	// `2. 合同组的校验和认证的方式会优先使用，会覆盖合同组中单个合同和合同签署方认证方式的限制配置。`
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil,omitempty" name:"ApproverVerifyType"`

	// 合同组的签署配置项信息，例如在合同组签署过程中，是否需要对每个子合同进行独立的意愿确认。
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil,omitempty" name:"FlowGroupOptions"`

	// 操作者的信息，此参数不用传
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelCreateFlowGroupByFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowGroupByFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowFileInfos")
	delete(f, "FlowGroupName")
	delete(f, "Agent")
	delete(f, "ApproverVerifyType")
	delete(f, "FlowGroupOptions")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowGroupByFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowGroupByFilesResponseParams struct {
	// 合同组ID，为32位字符串。
	// 建议开发者妥善保存此合同组ID，以便于顺利进行后续操作。
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 合同组中每个合同流程ID，每个ID均为32位字符串。
	// 
	// 注:
	// `此数组的顺序和入参中的FlowGroupInfos顺序一致`
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 合同组签署方信息。
	Approvers []*FlowGroupApprovers `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateFlowGroupByFilesResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateFlowGroupByFilesResponseParams `json:"Response"`
}

func (r *ChannelCreateFlowGroupByFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowGroupByFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowGroupByTemplatesRequestParams struct {
	// 合同的发起企业和发起人信息，<a href="https://qcloudimg.tencent-cloud.cn/raw/b69f8aad306c40b7b78d096e39b2edbb.png" target="_blank">点击查看合同发起企业和人展示的位置</a>
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  <a href="https://qcloudimg.tencent-cloud.cn/raw/a71872de3d540d55451e3e73a2ad1a6e.png" target="_blank">Agent.AppId</a></li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId（合同的发起企业）</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId （合同的发起人）</li>
	// </ul>
	// 
	// 合同的发起企业和发起人必需已经完成实名，并加入企业
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同组中每个合同签署流程的信息，合同组中最少包含2个合同，不能超过50个合同。
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil,omitempty" name:"FlowInfos"`

	// 合同组的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil,omitempty" name:"FlowGroupName"`
}

type ChannelCreateFlowGroupByTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 合同的发起企业和发起人信息，<a href="https://qcloudimg.tencent-cloud.cn/raw/b69f8aad306c40b7b78d096e39b2edbb.png" target="_blank">点击查看合同发起企业和人展示的位置</a>
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  <a href="https://qcloudimg.tencent-cloud.cn/raw/a71872de3d540d55451e3e73a2ad1a6e.png" target="_blank">Agent.AppId</a></li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId（合同的发起企业）</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId （合同的发起人）</li>
	// </ul>
	// 
	// 合同的发起企业和发起人必需已经完成实名，并加入企业
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同组中每个合同签署流程的信息，合同组中最少包含2个合同，不能超过50个合同。
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil,omitempty" name:"FlowInfos"`

	// 合同组的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil,omitempty" name:"FlowGroupName"`
}

func (r *ChannelCreateFlowGroupByTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowGroupByTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowInfos")
	delete(f, "FlowGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowGroupByTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowGroupByTemplatesResponseParams struct {
	// 合同组ID，为32位字符串。
	// 建议开发者妥善保存此合同组ID，以便于顺利进行后续操作。
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 合同组中每个合同流程ID，每个ID均为32位字符串。
	// 
	// 注:
	// `此数组的顺序和入参中的FlowInfos顺序一致`
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 复杂文档合成任务（如，包含动态表格的预览任务）的任务信息数组；
	// 如果文档需要异步合成，此字段会返回该异步任务的任务信息，后续可以通过ChannelGetTaskResultApi接口查询任务详情；
	TaskInfos []*TaskInfo `json:"TaskInfos,omitnil,omitempty" name:"TaskInfos"`

	// 合同组签署方信息
	Approvers []*FlowGroupApprovers `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateFlowGroupByTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateFlowGroupByTemplatesResponseParams `json:"Response"`
}

func (r *ChannelCreateFlowGroupByTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowGroupByTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowRemindsRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需执行催办的合同流程ID数组，最多支持100个。
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`
}

type ChannelCreateFlowRemindsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需执行催办的合同流程ID数组，最多支持100个。
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`
}

func (r *ChannelCreateFlowRemindsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowRemindsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowRemindsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowRemindsResponseParams struct {
	// 合同催办结果的详细信息列表。
	RemindFlowRecords []*RemindFlowRecords `json:"RemindFlowRecords,omitnil,omitempty" name:"RemindFlowRecords"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateFlowRemindsResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateFlowRemindsResponseParams `json:"Response"`
}

func (r *ChannelCreateFlowRemindsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowRemindsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowSignReviewRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 企业内部审核结果
	// <ul><li>PASS: 审核通过（流程可以继续签署或者发起）</li>
	// <li>REJECT: 审核拒绝（流程状态不变，可以继续调用审核接口通过审核）</li>
	// <li>SIGN_REJECT:拒签(流程终止，流程状态变为拒签状态)</li></ul>
	ReviewType *string `json:"ReviewType,omitnil,omitempty" name:"ReviewType"`

	// 审核结果原因
	// <ul><li>字符串长度不超过200</li>
	// <li>当ReviewType 是拒绝（REJECT） 时此字段必填。</li>
	// <li>当ReviewType 是拒绝（SIGN_REJECT） 时此字段必填。</li></ul>
	ReviewMessage *string `json:"ReviewMessage,omitnil,omitempty" name:"ReviewMessage"`

	// 审核节点的签署人标志，用于指定当前审核的签署方。
	// <font color= "red">注意：以下情况必须传递RecipientId</font>
	// <ul><li> **发起签署流程时，指定签署人需要审批（即签署人需要审批
	// <a href="https://qian.tencent.com/developers/partnerApis/dataTypes#flowapproverinfo" target="_blank">ApproverNeedSignReview</a>为true），则必须指定RecipientId**</li><li>**如果签署审核节点是个人， 此参数必填**。</li></ul>
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 流程审核操作类型，取值如下：
	// <ul><li>**SignReview**：（默认）签署审核</li>
	// <li>**CreateReview**：发起审核</li>
	// <li>注意：`该字段不传或者为空，则默认为SignReview签署审核，走签署审核流程`</li></ul>
	// 
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`
}

type ChannelCreateFlowSignReviewRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 企业内部审核结果
	// <ul><li>PASS: 审核通过（流程可以继续签署或者发起）</li>
	// <li>REJECT: 审核拒绝（流程状态不变，可以继续调用审核接口通过审核）</li>
	// <li>SIGN_REJECT:拒签(流程终止，流程状态变为拒签状态)</li></ul>
	ReviewType *string `json:"ReviewType,omitnil,omitempty" name:"ReviewType"`

	// 审核结果原因
	// <ul><li>字符串长度不超过200</li>
	// <li>当ReviewType 是拒绝（REJECT） 时此字段必填。</li>
	// <li>当ReviewType 是拒绝（SIGN_REJECT） 时此字段必填。</li></ul>
	ReviewMessage *string `json:"ReviewMessage,omitnil,omitempty" name:"ReviewMessage"`

	// 审核节点的签署人标志，用于指定当前审核的签署方。
	// <font color= "red">注意：以下情况必须传递RecipientId</font>
	// <ul><li> **发起签署流程时，指定签署人需要审批（即签署人需要审批
	// <a href="https://qian.tencent.com/developers/partnerApis/dataTypes#flowapproverinfo" target="_blank">ApproverNeedSignReview</a>为true），则必须指定RecipientId**</li><li>**如果签署审核节点是个人， 此参数必填**。</li></ul>
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 流程审核操作类型，取值如下：
	// <ul><li>**SignReview**：（默认）签署审核</li>
	// <li>**CreateReview**：发起审核</li>
	// <li>注意：`该字段不传或者为空，则默认为SignReview签署审核，走签署审核流程`</li></ul>
	// 
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`
}

func (r *ChannelCreateFlowSignReviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowSignReviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowId")
	delete(f, "ReviewType")
	delete(f, "ReviewMessage")
	delete(f, "RecipientId")
	delete(f, "OperateType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowSignReviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowSignReviewResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateFlowSignReviewResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateFlowSignReviewResponseParams `json:"Response"`
}

func (r *ChannelCreateFlowSignReviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowSignReviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowSignUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 流程签署人列表，其中结构体的ApproverType必传。
	// 若为个人签署方或saas企业签署方，则Name，Mobile必传。OrganizationName 传对应企业名称。
	// 若为子客企业签署方则需传OpenId、OrganizationOpenId、OrganizationName， 其他可不传。（如果对应OrganizationOpenId 子客已经认证激活了，则可以省去OrganizationName）
	// 
	// 此结构体和发起接口参与方结构体复用，除了上述参数外，可传递的参数有：
	// 1. RecipientId: 发起合同会返回，可以直接用于指定需要生成链接的签署方。
	// 2. ApproverSignTypes: 指定签署方签署时候的认证方式，仅此链接生效。
	// 3. SignTypeSelector: 可以指定签署方签署合同的认证校验方式的选择模式。
	// 4. Intention: 指定H5签署视频核身的意图配置，仅视频签署需要使用。
	// 
	// 注:
	// `1. 签署人只能有手写签名、时间类型、印章类型、签批类型的签署控件和内容填写控件，其他类型的签署控件暂时未支持。`
	// `2. 生成发起方预览链接时，该字段（FlowApproverInfos）传空或者不传`
	FlowApproverInfos []*FlowApproverInfo `json:"FlowApproverInfos,omitnil,omitempty" name:"FlowApproverInfos"`

	// 用户信息，暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 机构信息，暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 签署完之后的H5页面的跳转链接，最大长度1000个字符。链接类型请参考 <a href="https://qian.tencent.com/developers/company/openqianh5" target="_blank">跳转电子签H5</a>
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 链接类型，支持指定以下类型
	// <ul><li>0 : 签署链接 (默认值)</li>
	// <li>1 : 预览链接</li></ul>
	// 注:
	// `1. 当指定链接类型为1时，链接为预览链接，打开链接无法签署仅支持预览以及查看当前合同状态。`
	// `2. 如需生成发起方预览链接，则签署方信息传空，即FlowApproverInfos传空或者不传。`
	UrlType *int64 `json:"UrlType,omitnil,omitempty" name:"UrlType"`
}

type ChannelCreateFlowSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 流程签署人列表，其中结构体的ApproverType必传。
	// 若为个人签署方或saas企业签署方，则Name，Mobile必传。OrganizationName 传对应企业名称。
	// 若为子客企业签署方则需传OpenId、OrganizationOpenId、OrganizationName， 其他可不传。（如果对应OrganizationOpenId 子客已经认证激活了，则可以省去OrganizationName）
	// 
	// 此结构体和发起接口参与方结构体复用，除了上述参数外，可传递的参数有：
	// 1. RecipientId: 发起合同会返回，可以直接用于指定需要生成链接的签署方。
	// 2. ApproverSignTypes: 指定签署方签署时候的认证方式，仅此链接生效。
	// 3. SignTypeSelector: 可以指定签署方签署合同的认证校验方式的选择模式。
	// 4. Intention: 指定H5签署视频核身的意图配置，仅视频签署需要使用。
	// 
	// 注:
	// `1. 签署人只能有手写签名、时间类型、印章类型、签批类型的签署控件和内容填写控件，其他类型的签署控件暂时未支持。`
	// `2. 生成发起方预览链接时，该字段（FlowApproverInfos）传空或者不传`
	FlowApproverInfos []*FlowApproverInfo `json:"FlowApproverInfos,omitnil,omitempty" name:"FlowApproverInfos"`

	// 用户信息，暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 机构信息，暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 签署完之后的H5页面的跳转链接，最大长度1000个字符。链接类型请参考 <a href="https://qian.tencent.com/developers/company/openqianh5" target="_blank">跳转电子签H5</a>
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 链接类型，支持指定以下类型
	// <ul><li>0 : 签署链接 (默认值)</li>
	// <li>1 : 预览链接</li></ul>
	// 注:
	// `1. 当指定链接类型为1时，链接为预览链接，打开链接无法签署仅支持预览以及查看当前合同状态。`
	// `2. 如需生成发起方预览链接，则签署方信息传空，即FlowApproverInfos传空或者不传。`
	UrlType *int64 `json:"UrlType,omitnil,omitempty" name:"UrlType"`
}

func (r *ChannelCreateFlowSignUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowSignUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowId")
	delete(f, "FlowApproverInfos")
	delete(f, "Operator")
	delete(f, "Organization")
	delete(f, "JumpUrl")
	delete(f, "UrlType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowSignUrlResponseParams struct {
	// 签署人签署链接信息
	FlowApproverUrlInfos []*FlowApproverUrlInfo `json:"FlowApproverUrlInfos,omitnil,omitempty" name:"FlowApproverUrlInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateFlowSignUrlResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateFlowSignUrlResponseParams `json:"Response"`
}

func (r *ChannelCreateFlowSignUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateFlowSignUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateMultiFlowSignQRCodeRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同模板ID，为32位字符串。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 通过此二维码可发起的流程最大限额，如未明确指定，默认为5份。 一旦发起流程数超越该限制，该二维码将自动失效。	
	MaxFlowNum *int64 `json:"MaxFlowNum,omitnil,omitempty" name:"MaxFlowNum"`

	// 合同流程的签署有效期限，若未设定签署截止日期，则默认为自合同流程创建起的7天内截止。 若在签署截止日期前未完成签署，合同状态将变更为已过期，从而导致合同无效。 最长设定期限不得超过30天。	
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitnil,omitempty" name:"FlowEffectiveDay"`

	// 二维码的有效期限，默认为7天，最高设定不得超过90天。 一旦超过二维码的有效期限，该二维码将自动失效。	
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitnil,omitempty" name:"QrEffectiveDay"`

	// 指定签署人信息。 在指定签署人后，仅允许特定签署人通过扫描二维码进行签署。	
	Restrictions []*ApproverRestriction `json:"Restrictions,omitnil,omitempty" name:"Restrictions"`

	// 指定签署方经办人控件类型是个人印章签署控件（SIGN_SIGNATURE） 时，可选的签名方式。
	ApproverComponentLimitTypes []*ApproverComponentLimitType `json:"ApproverComponentLimitTypes,omitnil,omitempty" name:"ApproverComponentLimitTypes"`

	// 已废弃，回调配置统一使用企业应用管理-应用集成-第三方应用中的配置
	// <br/> 通过一码多扫二维码发起的合同，回调消息可参考文档 https://qian.tencent.com/developers/partner/callback_types_contracts_sign
	// <br/> 用户通过签署二维码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档 https://qian.tencent.com/developers/partner/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 限制二维码用户条件（已弃用）
	//
	// Deprecated: ApproverRestrictions is deprecated.
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitnil,omitempty" name:"ApproverRestrictions"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 禁止个人用户重复签署，默认不禁止，即同一用户可多次扫码签署多份合同。若要求同一用户仅能扫码签署一份合同，请传入true。
	ForbidPersonalMultipleSign *bool `json:"ForbidPersonalMultipleSign,omitnil,omitempty" name:"ForbidPersonalMultipleSign"`

	// 合同流程名称是否应包含扫码签署人的信息，且遵循特定格式（flowname-姓名-手机号后四位）。 例如，通过参数FlowName设定的扫码发起合同名称为“员工入职合同”，当扫码人张三（手机号18800009527）扫码签署时，合同名称将自动生成为“员工入职合同-张三-9527”。
	FlowNameAppendScannerInfo *bool `json:"FlowNameAppendScannerInfo,omitnil,omitempty" name:"FlowNameAppendScannerInfo"`
}

type ChannelCreateMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同模板ID，为32位字符串。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 通过此二维码可发起的流程最大限额，如未明确指定，默认为5份。 一旦发起流程数超越该限制，该二维码将自动失效。	
	MaxFlowNum *int64 `json:"MaxFlowNum,omitnil,omitempty" name:"MaxFlowNum"`

	// 合同流程的签署有效期限，若未设定签署截止日期，则默认为自合同流程创建起的7天内截止。 若在签署截止日期前未完成签署，合同状态将变更为已过期，从而导致合同无效。 最长设定期限不得超过30天。	
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitnil,omitempty" name:"FlowEffectiveDay"`

	// 二维码的有效期限，默认为7天，最高设定不得超过90天。 一旦超过二维码的有效期限，该二维码将自动失效。	
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitnil,omitempty" name:"QrEffectiveDay"`

	// 指定签署人信息。 在指定签署人后，仅允许特定签署人通过扫描二维码进行签署。	
	Restrictions []*ApproverRestriction `json:"Restrictions,omitnil,omitempty" name:"Restrictions"`

	// 指定签署方经办人控件类型是个人印章签署控件（SIGN_SIGNATURE） 时，可选的签名方式。
	ApproverComponentLimitTypes []*ApproverComponentLimitType `json:"ApproverComponentLimitTypes,omitnil,omitempty" name:"ApproverComponentLimitTypes"`

	// 已废弃，回调配置统一使用企业应用管理-应用集成-第三方应用中的配置
	// <br/> 通过一码多扫二维码发起的合同，回调消息可参考文档 https://qian.tencent.com/developers/partner/callback_types_contracts_sign
	// <br/> 用户通过签署二维码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档 https://qian.tencent.com/developers/partner/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 限制二维码用户条件（已弃用）
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitnil,omitempty" name:"ApproverRestrictions"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 禁止个人用户重复签署，默认不禁止，即同一用户可多次扫码签署多份合同。若要求同一用户仅能扫码签署一份合同，请传入true。
	ForbidPersonalMultipleSign *bool `json:"ForbidPersonalMultipleSign,omitnil,omitempty" name:"ForbidPersonalMultipleSign"`

	// 合同流程名称是否应包含扫码签署人的信息，且遵循特定格式（flowname-姓名-手机号后四位）。 例如，通过参数FlowName设定的扫码发起合同名称为“员工入职合同”，当扫码人张三（手机号18800009527）扫码签署时，合同名称将自动生成为“员工入职合同-张三-9527”。
	FlowNameAppendScannerInfo *bool `json:"FlowNameAppendScannerInfo,omitnil,omitempty" name:"FlowNameAppendScannerInfo"`
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
	delete(f, "Restrictions")
	delete(f, "ApproverComponentLimitTypes")
	delete(f, "CallbackUrl")
	delete(f, "ApproverRestrictions")
	delete(f, "Operator")
	delete(f, "ForbidPersonalMultipleSign")
	delete(f, "FlowNameAppendScannerInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateMultiFlowSignQRCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateMultiFlowSignQRCodeResponseParams struct {
	// 一码多签签署码的基本信息，用户可扫描该二维码进行签署操作。	
	QrCode *SignQrCode `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 一码多签签署码链接信息，适用于客户系统整合二维码功能。通过链接，用户可直接访问电子签名小程序并签署合同。	
	SignUrls *SignUrl `json:"SignUrls,omitnil,omitempty" name:"SignUrls"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ChannelCreateOrganizationBatchSignUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括子客企业及应用编、号等详细内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 请指定需执行批量签署的流程ID，数量范围为1-100。 您可登录腾讯电子签控制台，浏览 "合同"->"合同中心" 以查阅某一合同的FlowId（在页面中显示为合同ID）。 用户将利用链接对这些合同实施批量操作。	
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 第三方应用平台的用户openid。 您可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查阅某位员工的OpenId。 OpenId必须是传入合同（FlowId）中的签署人。
	// 
	// <ul>
	// <li>1. 若OpenId为空，Name和Mobile 必须提供。</li>
	// <li>2. 若OpenId 与 Name，Mobile均存在，将优先采用OpenId对应的员工。	</li>
	// </ul>
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：`请确保和合同中填入的一致`
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 员工手机号，必须与姓名一起使用。 如果OpenId为空，则此字段不能为空。同时，姓名和手机号码必须与传入合同（FlowId）中的签署人信息一致。	
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 合同组Id，传入此参数则可以不传FlowIds
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`
}

type ChannelCreateOrganizationBatchSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括子客企业及应用编、号等详细内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 请指定需执行批量签署的流程ID，数量范围为1-100。 您可登录腾讯电子签控制台，浏览 "合同"->"合同中心" 以查阅某一合同的FlowId（在页面中显示为合同ID）。 用户将利用链接对这些合同实施批量操作。	
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 第三方应用平台的用户openid。 您可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查阅某位员工的OpenId。 OpenId必须是传入合同（FlowId）中的签署人。
	// 
	// <ul>
	// <li>1. 若OpenId为空，Name和Mobile 必须提供。</li>
	// <li>2. 若OpenId 与 Name，Mobile均存在，将优先采用OpenId对应的员工。	</li>
	// </ul>
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：`请确保和合同中填入的一致`
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 员工手机号，必须与姓名一起使用。 如果OpenId为空，则此字段不能为空。同时，姓名和手机号码必须与传入合同（FlowId）中的签署人信息一致。	
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 合同组Id，传入此参数则可以不传FlowIds
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`
}

func (r *ChannelCreateOrganizationBatchSignUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateOrganizationBatchSignUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowIds")
	delete(f, "OpenId")
	delete(f, "Name")
	delete(f, "Mobile")
	delete(f, "FlowGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateOrganizationBatchSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateOrganizationBatchSignUrlResponseParams struct {
	// 批量签署入口链接，用户可使用这个链接跳转到控制台页面对合同进行签署操作。	
	SignUrl *string `json:"SignUrl,omitnil,omitempty" name:"SignUrl"`

	// 链接过期时间以 Unix 时间戳格式表示，从生成链接时间起，往后7天有效期。过期后短链将失效，无法打开。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateOrganizationBatchSignUrlResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateOrganizationBatchSignUrlResponseParams `json:"Response"`
}

func (r *ChannelCreateOrganizationBatchSignUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateOrganizationBatchSignUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateOrganizationModifyQrCodeRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// 
	// 渠道应用标识: Agent.AppId
	// 第三方平台子客企业标识: Agent.ProxyOrganizationOpenId
	// 第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type ChannelCreateOrganizationModifyQrCodeRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// 
	// 渠道应用标识: Agent.AppId
	// 第三方平台子客企业标识: Agent.ProxyOrganizationOpenId
	// 第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *ChannelCreateOrganizationModifyQrCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateOrganizationModifyQrCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateOrganizationModifyQrCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateOrganizationModifyQrCodeResponseParams struct {
	// 二维码下载链接
	QrCodeUrl *string `json:"QrCodeUrl,omitnil,omitempty" name:"QrCodeUrl"`

	// 二维码失效时间 UNIX 时间戳 精确到秒
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateOrganizationModifyQrCodeResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateOrganizationModifyQrCodeResponseParams `json:"Response"`
}

func (r *ChannelCreateOrganizationModifyQrCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateOrganizationModifyQrCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreatePrepareFlowGroupRequestParams struct {
	// 合同组中每个合同签署流程的信息，合同组中最少包含2个合同，不能超过50个合同。
	BaseFlowInfos []*BaseFlowInfo `json:"BaseFlowInfos,omitnil,omitempty" name:"BaseFlowInfos"`

	// 合同组的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil,omitempty" name:"FlowGroupName"`

	// 资源类型，取值有： <ul><li> **1**：模板</li> <li> **2**：文件</li></ul>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 合同的发起企业和发起人信息，<a href="https://qcloudimg.tencent-cloud.cn/raw/b69f8aad306c40b7b78d096e39b2edbb.png" target="_blank">点击查看合同发起企业和人展示的位置</a>
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识: <a href="https://qcloudimg.tencent-cloud.cn/raw/a71872de3d540d55451e3e73a2ad1a6e.png" target="_blank">Agent.AppId</a></li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId（合同的发起企业）</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId （合同的发起人）</li>
	// </ul>
	// 
	// 合同的发起企业和发起人必需已经完成实名，并加入企业
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type ChannelCreatePrepareFlowGroupRequest struct {
	*tchttp.BaseRequest
	
	// 合同组中每个合同签署流程的信息，合同组中最少包含2个合同，不能超过50个合同。
	BaseFlowInfos []*BaseFlowInfo `json:"BaseFlowInfos,omitnil,omitempty" name:"BaseFlowInfos"`

	// 合同组的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil,omitempty" name:"FlowGroupName"`

	// 资源类型，取值有： <ul><li> **1**：模板</li> <li> **2**：文件</li></ul>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 合同的发起企业和发起人信息，<a href="https://qcloudimg.tencent-cloud.cn/raw/b69f8aad306c40b7b78d096e39b2edbb.png" target="_blank">点击查看合同发起企业和人展示的位置</a>
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识: <a href="https://qcloudimg.tencent-cloud.cn/raw/a71872de3d540d55451e3e73a2ad1a6e.png" target="_blank">Agent.AppId</a></li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId（合同的发起企业）</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId （合同的发起人）</li>
	// </ul>
	// 
	// 合同的发起企业和发起人必需已经完成实名，并加入企业
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *ChannelCreatePrepareFlowGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreatePrepareFlowGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BaseFlowInfos")
	delete(f, "FlowGroupName")
	delete(f, "ResourceType")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreatePrepareFlowGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreatePrepareFlowGroupResponseParams struct {
	// 合同组ID，为32位字符串。
	// 建议开发者妥善保存此合同组ID，以便于顺利进行后续操作。
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 嵌入式发起链接
	PrepareUrl *string `json:"PrepareUrl,omitnil,omitempty" name:"PrepareUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreatePrepareFlowGroupResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreatePrepareFlowGroupResponseParams `json:"Response"`
}

func (r *ChannelCreatePrepareFlowGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreatePrepareFlowGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreatePrepareFlowRequestParams struct {
	// 资源类型，取值有：
	// <ul><li> **1**：模板</li>
	// <li> **2**：文件（默认值）</li></ul>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 要创建的合同信息
	FlowInfo *BaseFlowInfo `json:"FlowInfo,omitnil,omitempty" name:"FlowInfo"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 资源id，与ResourceType相对应，取值范围：
	// <ul>
	// <li>文件Id（通过UploadFiles获取文件资源Id）</li>
	// <li>模板Id（通过控制台创建模板后获取模板Id）</li>
	// </ul>
	// 注意：需要同时设置 ResourceType 参数指定资源类型
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 合同流程配置信息，用于配置发起合同时定制化如是否允许修改，某些按钮的隐藏等逻辑
	FlowOption *CreateFlowOption `json:"FlowOption,omitnil,omitempty" name:"FlowOption"`

	// 已废弃，请用FlowInfo.Approvers
	//
	// Deprecated: FlowApproverList is deprecated.
	FlowApproverList []*CommonFlowApprover `json:"FlowApproverList,omitnil,omitempty" name:"FlowApproverList"`

	// 合同Id：用于通过一个已发起的合同快速生成一个发起流程web链接
	// 注: `该参数必须是一个待发起审核的合同id，并且还未审核通过`
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 该参数不可用，请通过获取 web 可嵌入接口获取合同流程预览 URL
	//
	// Deprecated: NeedPreview is deprecated.
	NeedPreview *bool `json:"NeedPreview,omitnil,omitempty" name:"NeedPreview"`

	// 企业机构信息，不用传
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 操作人（用户）信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <font color="red">此参数已经废弃，请使用 CreateFlowOption 里面的 SignComponentConfig</font>
	// 签署控件的配置信息，用在嵌入式发起的页面配置，包括
	// 
	// - 签署控件 是否默认展示日期.
	//
	// Deprecated: SignComponentConfig is deprecated.
	SignComponentConfig *SignComponentConfig `json:"SignComponentConfig,omitnil,omitempty" name:"SignComponentConfig"`
}

type ChannelCreatePrepareFlowRequest struct {
	*tchttp.BaseRequest
	
	// 资源类型，取值有：
	// <ul><li> **1**：模板</li>
	// <li> **2**：文件（默认值）</li></ul>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 要创建的合同信息
	FlowInfo *BaseFlowInfo `json:"FlowInfo,omitnil,omitempty" name:"FlowInfo"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 资源id，与ResourceType相对应，取值范围：
	// <ul>
	// <li>文件Id（通过UploadFiles获取文件资源Id）</li>
	// <li>模板Id（通过控制台创建模板后获取模板Id）</li>
	// </ul>
	// 注意：需要同时设置 ResourceType 参数指定资源类型
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 合同流程配置信息，用于配置发起合同时定制化如是否允许修改，某些按钮的隐藏等逻辑
	FlowOption *CreateFlowOption `json:"FlowOption,omitnil,omitempty" name:"FlowOption"`

	// 已废弃，请用FlowInfo.Approvers
	FlowApproverList []*CommonFlowApprover `json:"FlowApproverList,omitnil,omitempty" name:"FlowApproverList"`

	// 合同Id：用于通过一个已发起的合同快速生成一个发起流程web链接
	// 注: `该参数必须是一个待发起审核的合同id，并且还未审核通过`
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 该参数不可用，请通过获取 web 可嵌入接口获取合同流程预览 URL
	NeedPreview *bool `json:"NeedPreview,omitnil,omitempty" name:"NeedPreview"`

	// 企业机构信息，不用传
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 操作人（用户）信息，不用传
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <font color="red">此参数已经废弃，请使用 CreateFlowOption 里面的 SignComponentConfig</font>
	// 签署控件的配置信息，用在嵌入式发起的页面配置，包括
	// 
	// - 签署控件 是否默认展示日期.
	SignComponentConfig *SignComponentConfig `json:"SignComponentConfig,omitnil,omitempty" name:"SignComponentConfig"`
}

func (r *ChannelCreatePrepareFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreatePrepareFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	delete(f, "FlowInfo")
	delete(f, "Agent")
	delete(f, "ResourceId")
	delete(f, "FlowOption")
	delete(f, "FlowApproverList")
	delete(f, "FlowId")
	delete(f, "NeedPreview")
	delete(f, "Organization")
	delete(f, "Operator")
	delete(f, "SignComponentConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreatePrepareFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreatePrepareFlowResponseParams struct {
	// 发起的合同嵌入链接， 可以直接点击进入进行合同发起， 有效期为5分钟
	PrepareFlowUrl *string `json:"PrepareFlowUrl,omitnil,omitempty" name:"PrepareFlowUrl"`

	// 合同发起后预览链接， 注意此时合同并未发起，仅只是展示效果， 有效期为5分钟
	PreviewFlowUrl *string `json:"PreviewFlowUrl,omitnil,omitempty" name:"PreviewFlowUrl"`

	// 发起的合同临时Id， 只有当点击进入链接，成功发起合同后， 此Id才有效
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreatePrepareFlowResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreatePrepareFlowResponseParams `json:"Response"`
}

func (r *ChannelCreatePrepareFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreatePrepareFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreatePreparedPersonalEsignRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 个人用户姓名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 证件号码, 应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 电子印章名字，1-50个中文字符
	// 注:`同一企业下电子印章名字不能相同`
	SealName *string `json:"SealName,omitnil,omitempty" name:"SealName"`

	// 电子印章图片base64编码，大小不超过10M（原始图片不超过5M），只支持PNG或JPG图片格式。
	// 
	SealImage *string `json:"SealImage,omitnil,omitempty" name:"SealImage"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li>
	// <li>OTHER_CARD_TYPE : 其他</li></ul>
	// 
	// 注: `其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 是否开启印章图片压缩处理，默认不开启，如需开启请设置为 true。当印章超过 2M 时建议开启，开启后图片的 hash 将发生变化。
	SealImageCompress *bool `json:"SealImageCompress,omitnil,omitempty" name:"SealImageCompress"`

	// 手机号码；当需要开通自动签时，该参数必传
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 该字段已不再使用
	EnableAutoSign *bool `json:"EnableAutoSign,omitnil,omitempty" name:"EnableAutoSign"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。一旦绑定后，将扣减购买的个人自动签账号许可一次（1年有效期），不可解绑释放。不传默认为绑定自动签账号许可。 0-绑定个人自动签账号许可，开通后将扣减购买的个人自动签账号许可一次 1-不绑定，发起合同时将按标准合同套餐进行扣减	
	LicenseType *int64 `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`
}

type ChannelCreatePreparedPersonalEsignRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 个人用户姓名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 证件号码, 应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 电子印章名字，1-50个中文字符
	// 注:`同一企业下电子印章名字不能相同`
	SealName *string `json:"SealName,omitnil,omitempty" name:"SealName"`

	// 电子印章图片base64编码，大小不超过10M（原始图片不超过5M），只支持PNG或JPG图片格式。
	// 
	SealImage *string `json:"SealImage,omitnil,omitempty" name:"SealImage"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li>
	// <li>OTHER_CARD_TYPE : 其他</li></ul>
	// 
	// 注: `其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 是否开启印章图片压缩处理，默认不开启，如需开启请设置为 true。当印章超过 2M 时建议开启，开启后图片的 hash 将发生变化。
	SealImageCompress *bool `json:"SealImageCompress,omitnil,omitempty" name:"SealImageCompress"`

	// 手机号码；当需要开通自动签时，该参数必传
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 该字段已不再使用
	EnableAutoSign *bool `json:"EnableAutoSign,omitnil,omitempty" name:"EnableAutoSign"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。一旦绑定后，将扣减购买的个人自动签账号许可一次（1年有效期），不可解绑释放。不传默认为绑定自动签账号许可。 0-绑定个人自动签账号许可，开通后将扣减购买的个人自动签账号许可一次 1-不绑定，发起合同时将按标准合同套餐进行扣减	
	LicenseType *int64 `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`
}

func (r *ChannelCreatePreparedPersonalEsignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreatePreparedPersonalEsignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "UserName")
	delete(f, "IdCardNumber")
	delete(f, "SealName")
	delete(f, "SealImage")
	delete(f, "Operator")
	delete(f, "IdCardType")
	delete(f, "SealImageCompress")
	delete(f, "Mobile")
	delete(f, "EnableAutoSign")
	delete(f, "LicenseType")
	delete(f, "SceneKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreatePreparedPersonalEsignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreatePreparedPersonalEsignResponseParams struct {
	// 电子印章ID，为32位字符串。
	// 建议开发者保留此印章ID，后续指定签署区印章或者操作印章需此印章ID。
	// 可登录腾讯电子签控制台，在 "印章"->"印章中心"选择查看的印章，在"印章详情" 中查看某个印章的SealId(在页面中展示为印章ID)。
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreatePreparedPersonalEsignResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreatePreparedPersonalEsignResponseParams `json:"Response"`
}

func (r *ChannelCreatePreparedPersonalEsignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreatePreparedPersonalEsignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateReleaseFlowRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 待解除的签署流程编号(即原签署流程的编号)。
	// 
	// 
	// [点击查看流程编号在控制台上的位置](https://qcloudimg.tencent-cloud.cn/raw/05af26573d5106763b4cfbb9f7c64b41.png)
	NeedRelievedFlowId *string `json:"NeedRelievedFlowId,omitnil,omitempty" name:"NeedRelievedFlowId"`

	// 解除协议内容, 包括解除理由等信息。
	ReliveInfo *RelieveInfo `json:"ReliveInfo,omitnil,omitempty" name:"ReliveInfo"`

	// 替换解除协议的签署人， 如不指定新的签署人，将继续使用原流程的签署人作为本解除协议的参与方。 <br/>
	// 如需更换原合同中的企业端签署人，可通过指定该签署人在原合同列表中的ApproverNumber编号来更换此企业端签署人。(可通过接口<a href="https://qian.tencent.com/developers/partnerApis/flows/DescribeFlowDetailInfo/">DescribeFlowDetailInfo</a>查询签署人的ApproverNumber编号，默认从0开始，顺序递增)<br/>
	// 
	// 注：
	// 1. 支持更换企业的签署人，不支持更换个人类型的签署人。
	// 2. 己方企业支持自动签署，他方企业不支持自动签署。
	// 3. <b>仅将需要替换的签署人添加至此列表</b>，无需替换的签署人无需添加进来。
	ReleasedApprovers []*ReleasedApprover `json:"ReleasedApprovers,omitnil,omitempty" name:"ReleasedApprovers"`

	// 签署完回调url，最大长度1000个字符
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳(秒)，如果未设置签署截止时间，则默认为合同流程创建后的7天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 调用方自定义的个性化字段，该字段的值可以是字符串JSON或其他字符串形式，客户可以根据自身需求自定义数据格式并在需要时进行解析。该字段的信息将以Base64编码的形式传输，支持的最大数据大小为20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

type ChannelCreateReleaseFlowRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 待解除的签署流程编号(即原签署流程的编号)。
	// 
	// 
	// [点击查看流程编号在控制台上的位置](https://qcloudimg.tencent-cloud.cn/raw/05af26573d5106763b4cfbb9f7c64b41.png)
	NeedRelievedFlowId *string `json:"NeedRelievedFlowId,omitnil,omitempty" name:"NeedRelievedFlowId"`

	// 解除协议内容, 包括解除理由等信息。
	ReliveInfo *RelieveInfo `json:"ReliveInfo,omitnil,omitempty" name:"ReliveInfo"`

	// 替换解除协议的签署人， 如不指定新的签署人，将继续使用原流程的签署人作为本解除协议的参与方。 <br/>
	// 如需更换原合同中的企业端签署人，可通过指定该签署人在原合同列表中的ApproverNumber编号来更换此企业端签署人。(可通过接口<a href="https://qian.tencent.com/developers/partnerApis/flows/DescribeFlowDetailInfo/">DescribeFlowDetailInfo</a>查询签署人的ApproverNumber编号，默认从0开始，顺序递增)<br/>
	// 
	// 注：
	// 1. 支持更换企业的签署人，不支持更换个人类型的签署人。
	// 2. 己方企业支持自动签署，他方企业不支持自动签署。
	// 3. <b>仅将需要替换的签署人添加至此列表</b>，无需替换的签署人无需添加进来。
	ReleasedApprovers []*ReleasedApprover `json:"ReleasedApprovers,omitnil,omitempty" name:"ReleasedApprovers"`

	// 签署完回调url，最大长度1000个字符
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳(秒)，如果未设置签署截止时间，则默认为合同流程创建后的7天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 调用方自定义的个性化字段，该字段的值可以是字符串JSON或其他字符串形式，客户可以根据自身需求自定义数据格式并在需要时进行解析。该字段的信息将以Base64编码的形式传输，支持的最大数据大小为20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

func (r *ChannelCreateReleaseFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateReleaseFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "NeedRelievedFlowId")
	delete(f, "ReliveInfo")
	delete(f, "ReleasedApprovers")
	delete(f, "CallbackUrl")
	delete(f, "Organization")
	delete(f, "Operator")
	delete(f, "Deadline")
	delete(f, "UserData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateReleaseFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateReleaseFlowResponseParams struct {
	// 解除协议流程编号
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateReleaseFlowResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateReleaseFlowResponseParams `json:"Response"`
}

func (r *ChannelCreateReleaseFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateReleaseFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateRoleRequestParams struct {
	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 权限树，权限树内容 PermissionGroups 可参考接口 ChannelDescribeRoles 的输出
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil,omitempty" name:"PermissionGroups"`
}

type ChannelCreateRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 权限树，权限树内容 PermissionGroups 可参考接口 ChannelDescribeRoles 的输出
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil,omitempty" name:"PermissionGroups"`
}

func (r *ChannelCreateRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Agent")
	delete(f, "Description")
	delete(f, "PermissionGroups")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateRoleResponseParams struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateRoleResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateRoleResponseParams `json:"Response"`
}

func (r *ChannelCreateRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateSealPolicyRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// 
	// 渠道应用标识: Agent.AppId
	// 第三方平台子客企业标识: Agent.ProxyOrganizationOpenId
	// 第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId
	// 第三方平台子客企业和员工必须已经经过实名认证。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 电子印章ID，为32位字符串。
	// 建议开发者保留此印章ID，后续指定签署区印章或者操作印章需此印章ID。
	// 可登录腾讯电子签控制台，在 "印章"->"印章中心"选择查看的印章，在"印章详情" 中查看某个印章的SealId(在页面中展示为印章ID)。
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 
	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查看某位员工的UserId(在页面中展示为用户ID)。
	// 员工在贵司业务系统中的唯一身份标识，用于与腾讯电子签账号进行映射，确保在同一企业内不会出现重复。
	// 该标识最大长度为64位字符串，仅支持包含26个英文字母和数字0-9的字符。
	// 指定待授权的用户ID数组,电子签的用户ID
	// 可以填写OpenId，系统会通过组织+渠道+OpenId查询得到UserId进行授权。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 操作人（用户）信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 企业机构信息，不用传
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
}

type ChannelCreateSealPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// 
	// 渠道应用标识: Agent.AppId
	// 第三方平台子客企业标识: Agent.ProxyOrganizationOpenId
	// 第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId
	// 第三方平台子客企业和员工必须已经经过实名认证。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 电子印章ID，为32位字符串。
	// 建议开发者保留此印章ID，后续指定签署区印章或者操作印章需此印章ID。
	// 可登录腾讯电子签控制台，在 "印章"->"印章中心"选择查看的印章，在"印章详情" 中查看某个印章的SealId(在页面中展示为印章ID)。
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 
	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查看某位员工的UserId(在页面中展示为用户ID)。
	// 员工在贵司业务系统中的唯一身份标识，用于与腾讯电子签账号进行映射，确保在同一企业内不会出现重复。
	// 该标识最大长度为64位字符串，仅支持包含26个英文字母和数字0-9的字符。
	// 指定待授权的用户ID数组,电子签的用户ID
	// 可以填写OpenId，系统会通过组织+渠道+OpenId查询得到UserId进行授权。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 操作人（用户）信息，不用传
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 企业机构信息，不用传
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
}

func (r *ChannelCreateSealPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateSealPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "SealId")
	delete(f, "UserIds")
	delete(f, "Operator")
	delete(f, "Organization")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateSealPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateSealPolicyResponseParams struct {
	// 最终授权成功的电子签系统用户ID数组。其他的跳过的是已经授权了的。
	// 请求参数填写OpenId时，返回授权成功的 Openid。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateSealPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateSealPolicyResponseParams `json:"Response"`
}

func (r *ChannelCreateSealPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateSealPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateUserAutoSignEnableUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签开通配置信息, 包括开通的人员的信息等
	AutoSignConfig *AutoSignConfig `json:"AutoSignConfig,omitnil,omitempty" name:"AutoSignConfig"`

	// 生成的链接类型：
	// <ul><li> 不传(即为空值) 则会生成小程序端开通链接(默认)</li>
	// <li> **H5SIGN** : 生成H5端开通链接</li></ul>
	UrlType *string `json:"UrlType,omitnil,omitempty" name:"UrlType"`

	// 是否通知开通方，通知类型:<ul><li>默认为不通知开通方</li><li>**SMS** :  短信通知 ,如果需要短信通知则NotifyAddress填写对方的手机号</li></ul>
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// 如果通知类型NotifyType选择为SMS，则此处为手机号, 其他通知类型不需要设置此项
	NotifyAddress *string `json:"NotifyAddress,omitnil,omitempty" name:"NotifyAddress"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为当前时间往后30天。`如果不传，默认过期时间为当前时间往后7天。`
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 调用方自定义的个性化字段(可自定义此字段的值)，并以base64方式编码，支持的最大数据大小为 20480长度。 在个人自动签的开通、关闭等回调信息场景中，该字段的信息将原封不动地透传给贵方。 
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

type ChannelCreateUserAutoSignEnableUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签开通配置信息, 包括开通的人员的信息等
	AutoSignConfig *AutoSignConfig `json:"AutoSignConfig,omitnil,omitempty" name:"AutoSignConfig"`

	// 生成的链接类型：
	// <ul><li> 不传(即为空值) 则会生成小程序端开通链接(默认)</li>
	// <li> **H5SIGN** : 生成H5端开通链接</li></ul>
	UrlType *string `json:"UrlType,omitnil,omitempty" name:"UrlType"`

	// 是否通知开通方，通知类型:<ul><li>默认为不通知开通方</li><li>**SMS** :  短信通知 ,如果需要短信通知则NotifyAddress填写对方的手机号</li></ul>
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// 如果通知类型NotifyType选择为SMS，则此处为手机号, 其他通知类型不需要设置此项
	NotifyAddress *string `json:"NotifyAddress,omitnil,omitempty" name:"NotifyAddress"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为当前时间往后30天。`如果不传，默认过期时间为当前时间往后7天。`
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 调用方自定义的个性化字段(可自定义此字段的值)，并以base64方式编码，支持的最大数据大小为 20480长度。 在个人自动签的开通、关闭等回调信息场景中，该字段的信息将原封不动地透传给贵方。 
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

func (r *ChannelCreateUserAutoSignEnableUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateUserAutoSignEnableUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "SceneKey")
	delete(f, "Operator")
	delete(f, "AutoSignConfig")
	delete(f, "UrlType")
	delete(f, "NotifyType")
	delete(f, "NotifyAddress")
	delete(f, "ExpiredTime")
	delete(f, "UserData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateUserAutoSignEnableUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateUserAutoSignEnableUrlResponseParams struct {
	// 个人用户自动签的开通链接, 短链形式。过期时间受 `ExpiredTime` 参数控制。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 腾讯电子签小程序的 AppID，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用
	// 
	// 注: `如果获取的是H5链接, 则不会返回此值`
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 腾讯电子签小程序的原始 Id,  ，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用
	// 
	// 注: `如果获取的是H5链接, 则不会返回此值`
	AppOriginalId *string `json:"AppOriginalId,omitnil,omitempty" name:"AppOriginalId"`

	// 腾讯电子签小程序的跳转路径，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用
	// 
	// 注: `如果获取的是H5链接, 则不会返回此值`
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// base64 格式的跳转二维码图片，可通过微信扫描后跳转到腾讯电子签小程序的开通界面。
	// 
	// 注: `如果获取的是H5链接, 则不会返回此二维码图片`
	QrCode *string `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 返回的链接类型
	// <ul><li> 空: 默认小程序端链接</li>
	// <li> **H5SIGN** : h5端链接</li></ul>
	UrlType *string `json:"UrlType,omitnil,omitempty" name:"UrlType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateUserAutoSignEnableUrlResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateUserAutoSignEnableUrlResponseParams `json:"Response"`
}

func (r *ChannelCreateUserAutoSignEnableUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateUserAutoSignEnableUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateUserAutoSignSealUrlRequestParams struct {
	// 渠道应用相关信息。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 自动签开通个人用户信息，包括名字，身份证等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为当前时间往后30天。`如果不传，默认过期时间为当前时间往后7天。`
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`
}

type ChannelCreateUserAutoSignSealUrlRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 自动签开通个人用户信息，包括名字，身份证等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为当前时间往后30天。`如果不传，默认过期时间为当前时间往后7天。`
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`
}

func (r *ChannelCreateUserAutoSignSealUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateUserAutoSignSealUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	delete(f, "Operator")
	delete(f, "ExpiredTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateUserAutoSignSealUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateUserAutoSignSealUrlResponseParams struct {
	// 腾讯电子签小程序的AppId，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 腾讯电子签小程序的原始Id，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用。
	AppOriginalId *string `json:"AppOriginalId,omitnil,omitempty" name:"AppOriginalId"`

	// 个人用户自动签的开通链接, 短链形式。过期时间受 `ExpiredTime` 参数控制。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 腾讯电子签小程序的跳转路径，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// base64格式的跳转二维码图片，可通过微信扫描后跳转到腾讯电子签小程序的开通界面。
	QrCode *string `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateUserAutoSignSealUrlResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateUserAutoSignSealUrlResponseParams `json:"Response"`
}

func (r *ChannelCreateUserAutoSignSealUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateUserAutoSignSealUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateUserRolesRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 绑定角色的角色id列表，最多 100 个
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`

	// 电子签用户ID列表，与OpenIds参数二选一,优先UserIds参数，最多 100 个
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 客户系统用户ID列表，与UserIds参数二选一,优先UserIds参数，最多 100 个
	OpenIds []*string `json:"OpenIds,omitnil,omitempty" name:"OpenIds"`

	// 操作者信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelCreateUserRolesRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 绑定角色的角色id列表，最多 100 个
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`

	// 电子签用户ID列表，与OpenIds参数二选一,优先UserIds参数，最多 100 个
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 客户系统用户ID列表，与UserIds参数二选一,优先UserIds参数，最多 100 个
	OpenIds []*string `json:"OpenIds,omitnil,omitempty" name:"OpenIds"`

	// 操作者信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelCreateUserRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateUserRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "RoleIds")
	delete(f, "UserIds")
	delete(f, "OpenIds")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateUserRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateUserRolesResponseParams struct {
	// 绑定失败的用户角色列表
	FailedCreateRoleData []*FailedCreateRoleData `json:"FailedCreateRoleData,omitnil,omitempty" name:"FailedCreateRoleData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateUserRolesResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateUserRolesResponseParams `json:"Response"`
}

func (r *ChannelCreateUserRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateUserRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateWebThemeConfigRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 主题类型<br/>EMBED_WEB_THEME：嵌入式主题
	// <ul>
	// <li>EMBED_WEB_THEME，web页面嵌入的主题风格配置</li>
	// <li>COMPANY_AUTHENTICATE，子客认证主题配置， 对当前第三方应用号生效，
	// 目前支持的有，背景图替换，隐藏企业认证页面导航栏和隐藏企业认证顶部logo</li>
	// </ul>
	ThemeType *string `json:"ThemeType,omitnil,omitempty" name:"ThemeType"`

	// 主题配置
	WebThemeConfig *WebThemeConfig `json:"WebThemeConfig,omitnil,omitempty" name:"WebThemeConfig"`
}

type ChannelCreateWebThemeConfigRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 主题类型<br/>EMBED_WEB_THEME：嵌入式主题
	// <ul>
	// <li>EMBED_WEB_THEME，web页面嵌入的主题风格配置</li>
	// <li>COMPANY_AUTHENTICATE，子客认证主题配置， 对当前第三方应用号生效，
	// 目前支持的有，背景图替换，隐藏企业认证页面导航栏和隐藏企业认证顶部logo</li>
	// </ul>
	ThemeType *string `json:"ThemeType,omitnil,omitempty" name:"ThemeType"`

	// 主题配置
	WebThemeConfig *WebThemeConfig `json:"WebThemeConfig,omitnil,omitempty" name:"WebThemeConfig"`
}

func (r *ChannelCreateWebThemeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateWebThemeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "ThemeType")
	delete(f, "WebThemeConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateWebThemeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateWebThemeConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelCreateWebThemeConfigResponse struct {
	*tchttp.BaseResponse
	Response *ChannelCreateWebThemeConfigResponseParams `json:"Response"`
}

func (r *ChannelCreateWebThemeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelCreateWebThemeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDeleteRoleRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 角色id，最多20个
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type ChannelDeleteRoleRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 角色id，最多20个
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

func (r *ChannelDeleteRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDeleteRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDeleteRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDeleteRoleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelDeleteRoleResponse struct {
	*tchttp.BaseResponse
	Response *ChannelDeleteRoleResponseParams `json:"Response"`
}

func (r *ChannelDeleteRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDeleteRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDeleteRoleUsersRequestParams struct {
	// 代理信息此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 角色Id（非超管或法人角色Id）
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 电子签用户ID列表，与OpenIds参数二选一,优先UserIds参数，最多两百
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 操作人信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 客户系统用户ID列表，与UserIds参数二选一,优先UserIds参数，最多两百
	OpenIds []*string `json:"OpenIds,omitnil,omitempty" name:"OpenIds"`
}

type ChannelDeleteRoleUsersRequest struct {
	*tchttp.BaseRequest
	
	// 代理信息此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 角色Id（非超管或法人角色Id）
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 电子签用户ID列表，与OpenIds参数二选一,优先UserIds参数，最多两百
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 客户系统用户ID列表，与UserIds参数二选一,优先UserIds参数，最多两百
	OpenIds []*string `json:"OpenIds,omitnil,omitempty" name:"OpenIds"`
}

func (r *ChannelDeleteRoleUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDeleteRoleUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "RoleId")
	delete(f, "UserIds")
	delete(f, "Operator")
	delete(f, "OpenIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDeleteRoleUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDeleteRoleUsersResponseParams struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelDeleteRoleUsersResponse struct {
	*tchttp.BaseResponse
	Response *ChannelDeleteRoleUsersResponseParams `json:"Response"`
}

func (r *ChannelDeleteRoleUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDeleteRoleUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDeleteSealPoliciesRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 操作的印章ID
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 需要删除授权的用户ID数组，可以传入电子签系统用户ID或OpenId。
	// 注: 
	// 1. `填写OpenId时，系统会通过组织+渠道+OpenId查询得到对应的UserId进行授权取消操作`
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 组织机构信息，不用传
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 操作人（用户）信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelDeleteSealPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 操作的印章ID
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 需要删除授权的用户ID数组，可以传入电子签系统用户ID或OpenId。
	// 注: 
	// 1. `填写OpenId时，系统会通过组织+渠道+OpenId查询得到对应的UserId进行授权取消操作`
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 组织机构信息，不用传
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 操作人（用户）信息，不用传
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelDeleteSealPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDeleteSealPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "SealId")
	delete(f, "UserIds")
	delete(f, "Organization")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDeleteSealPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDeleteSealPoliciesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelDeleteSealPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ChannelDeleteSealPoliciesResponseParams `json:"Response"`
}

func (r *ChannelDeleteSealPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDeleteSealPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeAccountBillDetailRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	// 第三方平台子客企业必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type ChannelDescribeAccountBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	// 第三方平台子客企业必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *ChannelDescribeAccountBillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeAccountBillDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDescribeAccountBillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeAccountBillDetailResponseParams struct {
	// 当前绑定中账号数量
	BoundAccountsNumber *int64 `json:"BoundAccountsNumber,omitnil,omitempty" name:"BoundAccountsNumber"`

	// 剩余可绑定账号数量
	RemainAvailableAccountsNumber *int64 `json:"RemainAvailableAccountsNumber,omitnil,omitempty" name:"RemainAvailableAccountsNumber"`

	// 已失效账号数量
	InvalidAccountsNumber *int64 `json:"InvalidAccountsNumber,omitnil,omitempty" name:"InvalidAccountsNumber"`

	// 购买数量
	TotalBuyAccountsNumber *int64 `json:"TotalBuyAccountsNumber,omitnil,omitempty" name:"TotalBuyAccountsNumber"`

	// 赠送数量
	TotalGiftAccountsNumber *int64 `json:"TotalGiftAccountsNumber,omitnil,omitempty" name:"TotalGiftAccountsNumber"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelDescribeAccountBillDetailResponse struct {
	*tchttp.BaseResponse
	Response *ChannelDescribeAccountBillDetailResponseParams `json:"Response"`
}

func (r *ChannelDescribeAccountBillDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeAccountBillDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeBillUsageDetailRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	// 第三方平台子客企业必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 查询开始时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的套餐类型 （选填 ）不传则查询所有套餐；
	// 目前支持:
	// <ul>
	// <li>**CloudEnterprise**: 企业版合同</li>
	// <li>**SingleSignature**: 单方签章</li>
	// <li>**CloudProve**: 签署报告</li>
	// <li>**CloudOnlineSign**: 腾讯会议在线签约</li>
	// <li>**ChannelWeCard**: 微工卡</li>
	// <li>**SignFlow**: 合同套餐</li>
	// <li>**SignFace**: 签署意愿（人脸识别）</li>
	// <li>**SignPassword**: 签署意愿（密码）</li>
	// <li>**SignSMS**: 签署意愿（短信）</li>
	// <li>**PersonalEssAuth**: 签署人实名（腾讯电子签认证）</li>
	// <li>**PersonalThirdAuth**: 签署人实名（信任第三方认证）</li>
	// <li>**OrgEssAuth**: 签署企业实名</li>
	// <li>**FlowNotify**: 短信通知</li>
	// <li>**AuthService**: 企业工商信息查询</li>
	// </ul>
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定分页每页返回的数据条数，如果不传默认为 50，单页最大支持 50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type ChannelDescribeBillUsageDetailRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	// 第三方平台子客企业必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 查询开始时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的套餐类型 （选填 ）不传则查询所有套餐；
	// 目前支持:
	// <ul>
	// <li>**CloudEnterprise**: 企业版合同</li>
	// <li>**SingleSignature**: 单方签章</li>
	// <li>**CloudProve**: 签署报告</li>
	// <li>**CloudOnlineSign**: 腾讯会议在线签约</li>
	// <li>**ChannelWeCard**: 微工卡</li>
	// <li>**SignFlow**: 合同套餐</li>
	// <li>**SignFace**: 签署意愿（人脸识别）</li>
	// <li>**SignPassword**: 签署意愿（密码）</li>
	// <li>**SignSMS**: 签署意愿（短信）</li>
	// <li>**PersonalEssAuth**: 签署人实名（腾讯电子签认证）</li>
	// <li>**PersonalThirdAuth**: 签署人实名（信任第三方认证）</li>
	// <li>**OrgEssAuth**: 签署企业实名</li>
	// <li>**FlowNotify**: 短信通知</li>
	// <li>**AuthService**: 企业工商信息查询</li>
	// </ul>
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定分页每页返回的数据条数，如果不传默认为 50，单页最大支持 50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *ChannelDescribeBillUsageDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeBillUsageDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QuotaType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDescribeBillUsageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeBillUsageDetailResponseParams struct {
	// 返回查询记录总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 消耗记录详情
	Details []*ChannelBillUsageDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelDescribeBillUsageDetailResponse struct {
	*tchttp.BaseResponse
	Response *ChannelDescribeBillUsageDetailResponseParams `json:"Response"`
}

func (r *ChannelDescribeBillUsageDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeBillUsageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeEmployeesRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定分页每页返回的数据条数，单页最大支持 20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询的关键字段，支持Key-Values查询。可选键值如下：
	// <ul>
	//   <li>Key:**"Status"**，Values: **["IsVerified"]**, 查询已实名的员工</li>
	//   <li>Key:**"Status"**，Values: **["NotVerified"]**, 查询未实名的员工</li>
	//   <li>Key:**"Status"**，Values: **["QuiteJob"]**, 查询离职员工</li>
	//   <li>Key:**"ExcludeQuiteJob"**，Values: **["true"]**, 查询排除离职员工</li>
	//   <li>Key:**"StaffOpenId"**，Values: **["OpenId1","OpenId2",...]**, 根据第三方系统用户OpenId查询员工</li>
	// </ul>
	// 注: `同名字的Key的过滤条件会冲突,  只能填写一个`
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量:从 0 开始，最大20000。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelDescribeEmployeesRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定分页每页返回的数据条数，单页最大支持 20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询的关键字段，支持Key-Values查询。可选键值如下：
	// <ul>
	//   <li>Key:**"Status"**，Values: **["IsVerified"]**, 查询已实名的员工</li>
	//   <li>Key:**"Status"**，Values: **["NotVerified"]**, 查询未实名的员工</li>
	//   <li>Key:**"Status"**，Values: **["QuiteJob"]**, 查询离职员工</li>
	//   <li>Key:**"ExcludeQuiteJob"**，Values: **["true"]**, 查询排除离职员工</li>
	//   <li>Key:**"StaffOpenId"**，Values: **["OpenId1","OpenId2",...]**, 根据第三方系统用户OpenId查询员工</li>
	// </ul>
	// 注: `同名字的Key的过滤条件会冲突,  只能填写一个`
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量:从 0 开始，最大20000。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelDescribeEmployeesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeEmployeesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDescribeEmployeesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeEmployeesResponseParams struct {
	// 员工信息列表。
	Employees []*Staff `json:"Employees,omitnil,omitempty" name:"Employees"`

	// 偏移量，默认为0，最大20000。关于<code>Offset</code>的更进一步介绍请参考 API <a href="https://cloud.tencent.com/document/api/213/15688" target="_blank">简介</a>中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定分页每页返回的数据条数，单页最大支持 20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 符合条件的员工数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelDescribeEmployeesResponse struct {
	*tchttp.BaseResponse
	Response *ChannelDescribeEmployeesResponseParams `json:"Response"`
}

func (r *ChannelDescribeEmployeesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeEmployeesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeFlowComponentsRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要获取填写控件填写内容的合同流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type ChannelDescribeFlowComponentsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要获取填写控件填写内容的合同流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

func (r *ChannelDescribeFlowComponentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeFlowComponentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDescribeFlowComponentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeFlowComponentsResponseParams struct {
	// 合同填写控件信息列表，填写控件会按照参与方角色进行分类。
	RecipientComponentInfos []*RecipientComponentInfo `json:"RecipientComponentInfos,omitnil,omitempty" name:"RecipientComponentInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelDescribeFlowComponentsResponse struct {
	*tchttp.BaseResponse
	Response *ChannelDescribeFlowComponentsResponseParams `json:"Response"`
}

func (r *ChannelDescribeFlowComponentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeFlowComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeOrganizationSealsRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定分页每页返回的数据条数，单页最大支持 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量，默认为0，最大为20000
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询授权用户信息类型，取值如下：
	// 
	// <ul> <li><b>0</b>：（默认）不返回授权用户信息</li> <li><b>1</b>：返回授权用户的信息</li> </ul>
	InfoType *int64 `json:"InfoType,omitnil,omitempty" name:"InfoType"`

	// 印章id，是否查询特定的印章（没有输入返回所有）
	// 
	// 注:  `没有输入返回所有记录，最大返回100条。`
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 电子印章类型 , 可选类型如下: <ul><li>**OFFICIAL**: 公章</li><li>**CONTRACT**: 合同专用章;</li><li>**FINANCE**: 财务专用章;</li><li>**PERSONNEL**: 人事专用章</li><li>**INVOICE**: 发票专用章</li><li>**LEGAL_PERSON_SEAL**: 法定代表人章;</li><li>**EMPLOYEE_QUALIFICATION_SEAL**: 员工执业章</li></ul>注:  `1.为空时查询所有类型的印章。`
	SealTypes []*string `json:"SealTypes,omitnil,omitempty" name:"SealTypes"`

	// 
	// 需查询的印章状态列表。
	// 
	// <ul> <li>空，()仅查询启用状态的印章；</li> <li><strong>ALL</strong>，查询所有状态的印章；</li> <li><strong>CHECKING</strong>，查询待审核的印章；</li> <li><strong>SUCCESS</strong>，查询启用状态的印章；</li> <li><strong>FAIL</strong>，查询印章审核拒绝的印章；</li> <li><strong>DISABLE</strong>，查询已停用的印章；</li> <li><strong>STOPPED</strong>，查询已终止的印章；</li> <li><strong>VOID</strong>，查询已作废的印章；</li> <li><strong>INVALID</strong>，查询已失效的印章。</li> </ul>
	SealStatuses []*string `json:"SealStatuses,omitnil,omitempty" name:"SealStatuses"`
}

type ChannelDescribeOrganizationSealsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定分页每页返回的数据条数，单页最大支持 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询偏移量，默认为0，最大为20000
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询授权用户信息类型，取值如下：
	// 
	// <ul> <li><b>0</b>：（默认）不返回授权用户信息</li> <li><b>1</b>：返回授权用户的信息</li> </ul>
	InfoType *int64 `json:"InfoType,omitnil,omitempty" name:"InfoType"`

	// 印章id，是否查询特定的印章（没有输入返回所有）
	// 
	// 注:  `没有输入返回所有记录，最大返回100条。`
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 电子印章类型 , 可选类型如下: <ul><li>**OFFICIAL**: 公章</li><li>**CONTRACT**: 合同专用章;</li><li>**FINANCE**: 财务专用章;</li><li>**PERSONNEL**: 人事专用章</li><li>**INVOICE**: 发票专用章</li><li>**LEGAL_PERSON_SEAL**: 法定代表人章;</li><li>**EMPLOYEE_QUALIFICATION_SEAL**: 员工执业章</li></ul>注:  `1.为空时查询所有类型的印章。`
	SealTypes []*string `json:"SealTypes,omitnil,omitempty" name:"SealTypes"`

	// 
	// 需查询的印章状态列表。
	// 
	// <ul> <li>空，()仅查询启用状态的印章；</li> <li><strong>ALL</strong>，查询所有状态的印章；</li> <li><strong>CHECKING</strong>，查询待审核的印章；</li> <li><strong>SUCCESS</strong>，查询启用状态的印章；</li> <li><strong>FAIL</strong>，查询印章审核拒绝的印章；</li> <li><strong>DISABLE</strong>，查询已停用的印章；</li> <li><strong>STOPPED</strong>，查询已终止的印章；</li> <li><strong>VOID</strong>，查询已作废的印章；</li> <li><strong>INVALID</strong>，查询已失效的印章。</li> </ul>
	SealStatuses []*string `json:"SealStatuses,omitnil,omitempty" name:"SealStatuses"`
}

func (r *ChannelDescribeOrganizationSealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeOrganizationSealsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InfoType")
	delete(f, "SealId")
	delete(f, "SealTypes")
	delete(f, "SealStatuses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDescribeOrganizationSealsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeOrganizationSealsResponseParams struct {
	// 在设定了SealId时，返回值为0或1；若未设定SealId，则返回公司的总印章数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 查询到的印章结果数组
	Seals []*OccupiedSeal `json:"Seals,omitnil,omitempty" name:"Seals"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelDescribeOrganizationSealsResponse struct {
	*tchttp.BaseResponse
	Response *ChannelDescribeOrganizationSealsResponseParams `json:"Response"`
}

func (r *ChannelDescribeOrganizationSealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeOrganizationSealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeRolesRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定每页返回的数据条数，和Offset参数配合使用，单页最大200。
	// 
	// 注: `因为历史原因, 此字段为字符串类型`
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询的关键字段:
	// Key:"**RoleType**",Values:["**1**"]查询系统角色，
	// Key:"**RoleType**",Values:["**2**"]查询自定义角色
	// Key:"**RoleStatus**",Values:["**1**"]查询启用角色
	// Key:"**RoleStatus**",Values:["**2**"]查询禁用角色
	// Key:"**IsReturnPermissionGroup**"，Values:["**0**"]表示接口不返回角色对应的权限树字段
	// Key:"**IsReturnPermissionGroup**"，Values:["**1**"]表示接口返回角色对应的权限树字段
	// 
	// 注: `同名字的Key的过滤条件会冲突, 只能填写一个`
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用，最大2000条。
	// 
	// 注：
	// 1.`offset从0开始，即第一页为0。`
	// 2.`默认从第一页返回。`
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 操作人信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelDescribeRolesRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定每页返回的数据条数，和Offset参数配合使用，单页最大200。
	// 
	// 注: `因为历史原因, 此字段为字符串类型`
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询的关键字段:
	// Key:"**RoleType**",Values:["**1**"]查询系统角色，
	// Key:"**RoleType**",Values:["**2**"]查询自定义角色
	// Key:"**RoleStatus**",Values:["**1**"]查询启用角色
	// Key:"**RoleStatus**",Values:["**2**"]查询禁用角色
	// Key:"**IsReturnPermissionGroup**"，Values:["**0**"]表示接口不返回角色对应的权限树字段
	// Key:"**IsReturnPermissionGroup**"，Values:["**1**"]表示接口返回角色对应的权限树字段
	// 
	// 注: `同名字的Key的过滤条件会冲突, 只能填写一个`
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用，最大2000条。
	// 
	// 注：
	// 1.`offset从0开始，即第一页为0。`
	// 2.`默认从第一页返回。`
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelDescribeRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDescribeRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeRolesResponseParams struct {
	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用，最大2000条。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定每页返回的数据条数，和Offset参数配合使用，单页最大200。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询角色的总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 查询的角色信息列表
	ChannelRoles []*ChannelRole `json:"ChannelRoles,omitnil,omitempty" name:"ChannelRoles"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelDescribeRolesResponse struct {
	*tchttp.BaseResponse
	Response *ChannelDescribeRolesResponseParams `json:"Response"`
}

func (r *ChannelDescribeRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeSignFaceVideoRequestParams struct {
	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署参与人在本流程中的编号ID(每个流程不同)，可用此ID来定位签署参与人在本流程的签署节点，也可用于后续创建签署链接等操作。
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type ChannelDescribeSignFaceVideoRequest struct {
	*tchttp.BaseRequest
	
	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署参与人在本流程中的编号ID(每个流程不同)，可用此ID来定位签署参与人在本流程的签署节点，也可用于后续创建签署链接等操作。
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *ChannelDescribeSignFaceVideoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeSignFaceVideoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "SignId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDescribeSignFaceVideoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeSignFaceVideoResponseParams struct {
	// 核身视频结果。
	VideoData *DetectInfoVideoData `json:"VideoData,omitnil,omitempty" name:"VideoData"`

	// 意愿核身问答模式结果。若未使用该意愿核身功能，该字段返回值可以不处理。
	IntentionQuestionResult *IntentionQuestionResult `json:"IntentionQuestionResult,omitnil,omitempty" name:"IntentionQuestionResult"`

	// 意愿核身点头确认模式的结果信息，若未使用该意愿核身功能，该字段返回值可以不处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionActionResult *IntentionActionResult `json:"IntentionActionResult,omitnil,omitempty" name:"IntentionActionResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelDescribeSignFaceVideoResponse struct {
	*tchttp.BaseResponse
	Response *ChannelDescribeSignFaceVideoResponseParams `json:"Response"`
}

func (r *ChannelDescribeSignFaceVideoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeSignFaceVideoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeUserAutoSignStatusRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 要查询状态的用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelDescribeUserAutoSignStatusRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 要查询状态的用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelDescribeUserAutoSignStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeUserAutoSignStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDescribeUserAutoSignStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeUserAutoSignStatusResponseParams struct {
	// 查询用户是否已开通自动签
	IsOpen *bool `json:"IsOpen,omitnil,omitempty" name:"IsOpen"`

	// 自动签许可生效时间。当且仅当已通过许可开通自动签时有值。
	// 
	// 值为unix时间戳,单位为秒。
	LicenseFrom *int64 `json:"LicenseFrom,omitnil,omitempty" name:"LicenseFrom"`

	// 自动签许可到期时间。当且仅当已通过许可开通自动签时有值。
	// 
	// 值为unix时间戳,单位为秒。
	LicenseTo *int64 `json:"LicenseTo,omitnil,omitempty" name:"LicenseTo"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。<ul><li>**0**: 使用个人自动签账号许可进行开通，个人自动签账号许可有效期1年，注: `不可解绑释放更换他人`</li><li>**1**: 不绑定自动签账号许可开通，后续使用合同份额进行合同发起</li></ul>
	LicenseType *int64 `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 用户开通自动签指定使用的印章，为空则未设置印章，需重新进入开通链接设置印章。
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelDescribeUserAutoSignStatusResponse struct {
	*tchttp.BaseResponse
	Response *ChannelDescribeUserAutoSignStatusResponseParams `json:"Response"`
}

func (r *ChannelDescribeUserAutoSignStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDescribeUserAutoSignStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDisableUserAutoSignRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 需要关闭自动签的个人的信息，如姓名，证件信息等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelDisableUserAutoSignRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 需要关闭自动签的个人的信息，如姓名，证件信息等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelDisableUserAutoSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDisableUserAutoSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDisableUserAutoSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDisableUserAutoSignResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelDisableUserAutoSignResponse struct {
	*tchttp.BaseResponse
	Response *ChannelDisableUserAutoSignResponseParams `json:"Response"`
}

func (r *ChannelDisableUserAutoSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelDisableUserAutoSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelGetTaskResultApiRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 转换任务Id，通过接口<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelCreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>得到的转换任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
}

type ChannelGetTaskResultApiRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 转换任务Id，通过接口<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelCreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>得到的转换任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
}

func (r *ChannelGetTaskResultApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelGetTaskResultApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "TaskId")
	delete(f, "Operator")
	delete(f, "Organization")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelGetTaskResultApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelGetTaskResultApiResponseParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态，需要关注的状态
	// <ul><li>**0**  :NeedTranform   - 任务已提交</li>
	// <li>**4**  :Processing     - 文档转换中</li>
	// <li>**8**  :TaskEnd        - 任务处理完成</li>
	// <li>**-2** :DownloadFailed - 下载失败</li>
	// <li>**-6** :ProcessFailed  - 转换失败</li>
	// <li>**-13**:ProcessTimeout - 转换文件超时</li></ul>
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 状态描述，需要关注的状态
	// <ul><li> **NeedTranform** : 任务已提交</li>
	// <li> **Processing** : 文档转换中</li>
	// <li> **TaskEnd** : 任务处理完成</li>
	// <li> **DownloadFailed** : 下载失败</li>
	// <li> **ProcessFailed** : 转换失败</li>
	// <li> **ProcessTimeout** : 转换文件超时</li></ul>
	TaskMessage *string `json:"TaskMessage,omitnil,omitempty" name:"TaskMessage"`

	// 资源Id（即FileId），用于[用PDF文件创建签署流程](https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowByFiles)
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 预览文件Url，有效期30分钟 
	// 当前字段返回为空，发起的时候，将ResourceId 放入发起即可
	//
	// Deprecated: PreviewUrl is deprecated.
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelGetTaskResultApiResponse struct {
	*tchttp.BaseResponse
	Response *ChannelGetTaskResultApiResponseParams `json:"Response"`
}

func (r *ChannelGetTaskResultApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelGetTaskResultApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelModifyRoleRequestParams struct {
	// 代理企业和员工的信息。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 角色Id，可通过接口 ChannelDescribeRoles 查询获取
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 权限树，权限树内容 PermissionGroups 可参考接口 ChannelDescribeRoles的输出
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil,omitempty" name:"PermissionGroups"`
}

type ChannelModifyRoleRequest struct {
	*tchttp.BaseRequest
	
	// 代理企业和员工的信息。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 角色Id，可通过接口 ChannelDescribeRoles 查询获取
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 权限树，权限树内容 PermissionGroups 可参考接口 ChannelDescribeRoles的输出
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil,omitempty" name:"PermissionGroups"`
}

func (r *ChannelModifyRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelModifyRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "Name")
	delete(f, "RoleId")
	delete(f, "Description")
	delete(f, "PermissionGroups")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelModifyRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelModifyRoleResponseParams struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelModifyRoleResponse struct {
	*tchttp.BaseResponse
	Response *ChannelModifyRoleResponseParams `json:"Response"`
}

func (r *ChannelModifyRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelModifyRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChannelOrganizationInfo struct {
	// 电子签平台给企业分配的ID（在不同应用下同一个企业会分配通用的ID）
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 第三方平台子客企业的唯一标识
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 第三方平台子客企业名称
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 企业的统一社会信用代码
	UnifiedSocialCreditCode *string `json:"UnifiedSocialCreditCode,omitnil,omitempty" name:"UnifiedSocialCreditCode"`

	// 企业法定代表人的姓名
	LegalName *string `json:"LegalName,omitnil,omitempty" name:"LegalName"`

	// 企业法定代表人作为第三方平台子客企业员工的唯一标识
	LegalOpenId *string `json:"LegalOpenId,omitnil,omitempty" name:"LegalOpenId"`

	// 企业超级管理员的姓名
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 企业超级管理员作为第三方平台子客企业员工的唯一标识
	AdminOpenId *string `json:"AdminOpenId,omitnil,omitempty" name:"AdminOpenId"`

	// 企业超级管理员的手机号码
	// **注**：`手机号码脱敏（隐藏部分用*替代）`
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`

	// 企业认证状态字段。值如下：
	// <ul>
	//   <li>**"UNVERIFIED"**： 未认证的企业</li>
	//   <li>**"VERIFYINGLEGALPENDINGAUTHORIZATION"**： 认证中待法人授权的企业</li>
	//   <li>**"VERIFYINGAUTHORIZATIONFILEPENDING"**： 认证中授权书审核中的企业</li>
	//   <li>**"VERIFYINGAUTHORIZATIONFILEREJECT"**： 认证中授权书已驳回的企业</li>
	//   <li>**"VERIFYING"**： 认证中的企业</li>
	//   <li>**"VERIFIED"**： 已认证的企业</li>
	// </ul>
	AuthorizationStatus *string `json:"AuthorizationStatus,omitnil,omitempty" name:"AuthorizationStatus"`

	// 企业认证方式字段。值如下：
	// <ul>
	//   <li>**"AuthorizationInit"**： 暂未选择授权方式</li>
	//   <li>**"AuthorizationFile"**： 授权书</li>
	//   <li>**"AuthorizationLegalPerson"**： 法人授权超管</li>
	//   <li>**"AuthorizationLegalIdentity"**： 法人直接认证</li>
	// </ul>
	AuthorizationType *string `json:"AuthorizationType,omitnil,omitempty" name:"AuthorizationType"`

	// 子企业激活状态。值如下：
	// <ul>
	//   <li>**0**： 未激活</li>
	//   <li>**1**： 已激活</li>
	// </ul>
	ActiveStatus *int64 `json:"ActiveStatus,omitnil,omitempty" name:"ActiveStatus"`

	// 账号到期时间，时间戳
	LicenseExpireTime *int64 `json:"LicenseExpireTime,omitnil,omitempty" name:"LicenseExpireTime"`
}

// Predefined struct for user
type ChannelRenewAutoSignLicenseRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 要查询状态的用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelRenewAutoSignLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 要查询状态的用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelRenewAutoSignLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelRenewAutoSignLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelRenewAutoSignLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelRenewAutoSignLicenseResponseParams struct {
	// 续期成功后自动签许可到期时间。当且仅当已通过许可开通自动签时有值。
	// 
	// 值为unix时间戳,单位为秒。
	LicenseTo *int64 `json:"LicenseTo,omitnil,omitempty" name:"LicenseTo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelRenewAutoSignLicenseResponse struct {
	*tchttp.BaseResponse
	Response *ChannelRenewAutoSignLicenseResponseParams `json:"Response"`
}

func (r *ChannelRenewAutoSignLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelRenewAutoSignLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChannelRole struct {
	// 角色ID,为32位字符串
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色的名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 此角色状态
	// 1: 已经启用
	// 2: 已经禁用
	RoleStatus *uint64 `json:"RoleStatus,omitnil,omitempty" name:"RoleStatus"`

	// 此角色对应的权限列表
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil,omitempty" name:"PermissionGroups"`
}

// Predefined struct for user
type ChannelUpdateSealStatusRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 印章状态，目前支持传入以下类型：
	// <ul><li>DISABLE-停用印章</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 印章ID
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 更新印章状态原因说明
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 操作者的信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelUpdateSealStatusRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 印章状态，目前支持传入以下类型：
	// <ul><li>DISABLE-停用印章</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 印章ID
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 更新印章状态原因说明
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelUpdateSealStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelUpdateSealStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "Status")
	delete(f, "SealId")
	delete(f, "Reason")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelUpdateSealStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelUpdateSealStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelUpdateSealStatusResponse struct {
	*tchttp.BaseResponse
	Response *ChannelUpdateSealStatusResponseParams `json:"Response"`
}

func (r *ChannelUpdateSealStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelUpdateSealStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelVerifyPdfRequestParams struct {
	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type ChannelVerifyPdfRequest struct {
	*tchttp.BaseRequest
	
	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

func (r *ChannelVerifyPdfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelVerifyPdfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Agent")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelVerifyPdfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelVerifyPdfResponseParams struct {
	// 验签结果代码，代码的含义如下：
	// 
	// <ul><li>**1**：文件未被篡改，全部签名在腾讯电子签完成。</li>
	// <li>**2**：文件未被篡改，部分签名在腾讯电子签完成。</li>
	// <li>**3**：文件被篡改。</li>
	// <li>**4**：异常：文件内没有签名域。(如果合同还没有签署也会返回此代码)</li>
	// <li>**5**：异常：文件签名格式错误。</li></ul>
	VerifyResult *int64 `json:"VerifyResult,omitnil,omitempty" name:"VerifyResult"`

	// 验签结果详情，所有签署区(包括签名区, 印章区, 日期签署区,骑缝章等)的签署验签结果
	PdfVerifyResults []*PdfVerifyResult `json:"PdfVerifyResults,omitnil,omitempty" name:"PdfVerifyResults"`

	// 验签序列号, 为11为数组组成的字符串
	VerifySerialNo *string `json:"VerifySerialNo,omitnil,omitempty" name:"VerifySerialNo"`

	// 合同文件MD5哈希值
	PdfResourceMd5 *string `json:"PdfResourceMd5,omitnil,omitempty" name:"PdfResourceMd5"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChannelVerifyPdfResponse struct {
	*tchttp.BaseResponse
	Response *ChannelVerifyPdfResponseParams `json:"Response"`
}

func (r *ChannelVerifyPdfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChannelVerifyPdfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonApproverOption struct {
	// 是否允许修改签署人信息
	CanEditApprover *bool `json:"CanEditApprover,omitnil,omitempty" name:"CanEditApprover"`

	// 是否可以拒签 默认false-可以拒签 true-不可以拒签
	NoRefuse *bool `json:"NoRefuse,omitnil,omitempty" name:"NoRefuse"`

	// 是否可以转发 默认false-可以转发 true-不可以转发
	NoTransfer *bool `json:"NoTransfer,omitnil,omitempty" name:"NoTransfer"`

	// 当签署方有多个签署区时候，是否隐藏一键所有的签署区
	// 
	// false：（默认）不隐藏
	// true：隐藏，每个签署区要单独选择印章或者签名
	HideOneKeySign *bool `json:"HideOneKeySign,omitnil,omitempty" name:"HideOneKeySign"`

	// 签署人阅读合同限制参数
	//  <br/>取值：
	// <ul>
	// <li> LimitReadTimeAndBottom，阅读合同必须限制阅读时长并且必须阅读到底</li>
	// <li> LimitReadTime，阅读合同仅限制阅读时长</li>
	// <li> LimitBottom，阅读合同仅限制必须阅读到底</li>
	// <li> NoReadTimeAndBottom，阅读合同不限制阅读时长且不限制阅读到底（白名单功能，请联系客户经理开白使用）</li>
	// </ul>
	FlowReadLimit *string `json:"FlowReadLimit,omitnil,omitempty" name:"FlowReadLimit"`

	// 禁止在签署过程中添加签署日期控件
	//  <br/>前置条件：文件发起合同时，指定SignBeanTag=1（可以在签署过程中添加签署控件）：
	// <ul>
	// <li> 默认值：false，在开启：签署过程中添加签署控件时，添加签署控件会默认自带签署日期控件</li>
	// <li> 可选值：true，在开启：签署过程中添加签署控件时，添加签署控件不会自带签署日期控件</li>
	// </ul>
	ForbidAddSignDate *bool `json:"ForbidAddSignDate,omitnil,omitempty" name:"ForbidAddSignDate"`
}

type CommonFlowApprover struct {
	// 指定签署人非第三方平台子客企业下员工还是SaaS平台企业，在ApproverType为ORGANIZATION时指定。
	// <ul>
	// <li>false: 默认值，第三方平台子客企业下员工</li>
	// <li>true: SaaS平台企业下的员工</li>
	// </ul>
	NotChannelOrganization *bool `json:"NotChannelOrganization,omitnil,omitempty" name:"NotChannelOrganization"`

	// 在指定签署方时，可选择企业B端或个人C端等不同的参与者类型，可选类型如下:
	// 
	//  **0** :企业/企业员工（企业签署方或模板发起时的企业静默签）
	//  **1** :个人/自然人
	// **3** :企业/企业员工自动签（他方企业自动签署或文件发起时的本方企业自动签）
	// 
	// 注：类型为3（企业/企业员工自动签）时，此接口会默认完成该签署方的签署。静默签署仅进行盖章操作，不能自动签名。
	// 使用自动签时，请确保企业已经开通自动签功能，开通方式：控制台 -> 企业设置 -> 扩展服务 -> 企业自动签。
	// 使用文件发起自动签时使用前请联系对接的客户经理沟通。
	ApproverType *int64 `json:"ApproverType,omitnil,omitempty" name:"ApproverType"`

	// 电子签平台给企业生成的企业id
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 企业OpenId，第三方应用集成非静默签子客企业签署人发起合同必传
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 企业名称，第三方应用集成非静默签子客企业签署人必传，saas企业签署人必传
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 电子签平台给企业员工或者自热人生成的用户id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 第三方平台子客企业员工的唯一标识
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	ApproverName *string `json:"ApproverName,omitnil,omitempty" name:"ApproverName"`

	// 签署人手机号，saas企业签署人，个人签署人必传
	ApproverMobile *string `json:"ApproverMobile,omitnil,omitempty" name:"ApproverMobile"`

	// 签署方经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li>
	// <li>OTHER_CARD_TYPE : 其他证件</li></ul>
	// 
	// 注: `其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	ApproverIdCardType *string `json:"ApproverIdCardType,omitnil,omitempty" name:"ApproverIdCardType"`

	// 签署方经办人的证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。
	// </li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitnil,omitempty" name:"ApproverIdCardNumber"`

	// 签署人Id，使用模板发起是，对应模板配置中的签署人RecipientId
	// 注意：模板发起时该字段必填
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 签署前置条件：阅读时长限制，不传默认10s,最大300s，最小3s
	PreReadTime *int64 `json:"PreReadTime,omitnil,omitempty" name:"PreReadTime"`

	// 签署前置条件：阅读全文限制
	IsFullText *bool `json:"IsFullText,omitnil,omitempty" name:"IsFullText"`

	// 通知签署方经办人的方式, 有以下途径:
	// <ul><li> **SMS** :(默认)短信</li>
	// <li> **NONE** : 不通知</li></ul>
	// 
	// 注: `签署方为第三方子客企业时会被置为NONE,   不会发短信通知`
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// 签署人配置，用于控制签署人相关属性
	ApproverOption *CommonApproverOption `json:"ApproverOption,omitnil,omitempty" name:"ApproverOption"`

	// 使用PDF文件直接发起合同时，签署人指定的签署控件；<br/>使用模板发起合同时，指定本企业印章签署控件的印章ID: <br/>通过ComponentId或ComponenetName指定签署控件，ComponentValue为印章ID。
	SignComponents []*Component `json:"SignComponents,omitnil,omitempty" name:"SignComponents"`

	// 指定个人签署方查看合同的校验方式,可以传值如下:
	// <ul><li>  **1**   : （默认）人脸识别,人脸识别后才能合同内容</li>
	// <li>  **2**  : 手机号验证, 用户手机号和参与方手机号(ApproverMobile)相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// </li></ul>
	// 注: 
	// <ul><li>如果合同流程设置ApproverVerifyType查看合同的校验方式,    则忽略此签署人的查看合同的校验方式</li>
	// <li>此字段可传多个校验方式</li></ul>
	ApproverVerifyTypes []*int64 `json:"ApproverVerifyTypes,omitnil,omitempty" name:"ApproverVerifyTypes"`

	// 签署人签署合同时的认证方式
	// <ul><li> **1** :人脸认证</li>
	// <li> **2** :签署密码</li>
	// <li> **3** :运营商三要素</li>
	// <li> **5** :设备指纹识别</li>
	// <li> **6** :设备面容识别</li></ul>
	// 
	// 默认为1(人脸认证 ),2(签署密码),3(运营商三要素),5(设备指纹识别),6(设备面容识别)
	// 
	// 注: 
	// 1. 用<font color='red'>模板创建合同场景</font>, 签署人的认证方式需要在配置模板的时候指定, <font color='red'>在创建合同重新指定无效</font>
	// 2. 运营商三要素认证方式对手机号运营商及前缀有限制,可以参考[运营商支持列表类](https://qian.tencent.com/developers/partner/mobile_support)得到具体的支持说明
	// 3. 校验方式不允许只包含<font color='red'>设备指纹识别</font>和<font color='red'>设备面容识别</font>，至少需要再增加一种其他校验方式。
	// 4. <font color='red'>设备指纹识别</font>和<font color='red'>设备面容识别</font>只支持小程序使用，其他端暂不支持。
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil,omitempty" name:"ApproverSignTypes"`
}

type Component struct {
	// 控件唯一ID。
	// 
	// **在绝对定位方式方式下**，ComponentId为控件的ID，长度不能超过30，只能由中文、字母、数字和下划线组成，可以在后续的操作中使用该名称来引用控件。
	// 
	// **在关键字定位方式下**，ComponentId不仅为控件的ID，也是关键字整词。此方式下可以通过"^"来决定是否使用关键字整词匹配能力。
	// 
	// 例：
	// 
	// - 如传入的关键字<font color="red">"^甲方签署^"</font >，则会在PDF文件中有且仅有"甲方签署"关键字的地方（<font color="red">前后不能有其他字符</font >）进行对应操作。
	// - 如传入的关键字为<font color="red">"甲方签署</font >"，则PDF文件中每个出现关键字的位置（<font color="red">前后可以有其他字符</font >）都会执行相应操作。
	// 
	// 
	// 注：`控件ID可以在一个PDF中不可重复`
	// <a href="https://qcloudimg.tencent-cloud.cn/raw/93178569d07b4d7dbbe0967ae679e35c.png" target="_blank">点击查看ComponentId在模板页面的位置</a>
	ComponentId *string `json:"ComponentId,omitnil,omitempty" name:"ComponentId"`

	// **如果是Component填写控件类型，则可选的字段为**：
	// 
	// <ul><li> <b>TEXT</b> : 普通文本控件，输入文本字符串；</li>
	// <li> <b>MULTI_LINE_TEXT</b> : 多行文本控件，输入文本字符串；</li>
	// <li> <b>CHECK_BOX</b> : 勾选框控件，若选中填写ComponentValue 填写 true或者 false 字符串；</li>
	// <li> <b>FILL_IMAGE</b> : 图片控件，ComponentValue 填写图片的资源 ID；</li>
	// <li> <b>DYNAMIC_TABLE</b> : 动态表格控件；</li>
	// <li> <b>ATTACHMENT</b> : 附件控件,ComponentValue 填写附件图片的资源 ID列表，以逗号分隔；</li>
	// <li> <b>SELECTOR</b> : 选择器控件，ComponentValue填写选择的字符串内容；</li>
	// <li> <b>DATE</b> : 日期控件；默认是格式化为xxxx年xx月xx日字符串；</li>
	// <li> <b>DISTRICT</b> : 省市区行政区控件，ComponentValue填写省市区行政区字符串内容；</li></ul>
	// 
	// **如果是SignComponent签署控件类型，
	// 需要根据签署人的类型可选的字段为**
	// * 企业方
	// <ul><li> <b>SIGN_SEAL</b> : 签署印章控件；</li>
	// <li> <b>SIGN_DATE</b> : 签署日期控件；</li>
	// <li> <b>SIGN_SIGNATURE</b> : 用户签名控件；</li>
	// <li> <b>SIGN_PAGING_SIGNATURE</b> : 用户签名骑缝章控件；若文件发起，需要对应填充ComponentPosY、ComponentWidth、ComponentHeight</li>
	// <li> <b>SIGN_PAGING_SEAL</b> : 骑缝章；若文件发起，需要对应填充ComponentPosY、ComponentWidth、ComponentHeight</li>
	// <li> <b>SIGN_OPINION</b> : 签署意见控件，用户需要根据配置的签署意见内容，完成对意见内容的确认；</li>
	// <li> <b>SIGN_VIRTUAL_COMBINATION</b> : 签批控件。内部最多组合4个特定控件（SIGN_SIGNATURE，SIGN_DATA,SIGN_MULTI_LINE_TEXT,SIGN_SELECTOR），本身不填充任何文字内容</li>
	// <li> <b>SIGN_MULTI_LINE_TEXT</b> : 多行文本，<font color="red">仅可用在签批控件内部作为组合控件，单独无法使用</font>，常用作批注附言</li>
	// <li> <b>SIGN_SELECTOR</b> : 选择器，<font color="red">仅可用在签批控件内部作为组合控件，单独无法使用</font>，常用作审批意见的选择</li>
	// <li> <b>SIGN_LEGAL_PERSON_SEAL</b> : 企业法定代表人控件。</li></ul>
	// 
	// 
	// * 个人方
	// <ul><li> <b>SIGN_DATE</b> : 签署日期控件；</li>
	// <li> <b>SIGN_SIGNATURE</b> : 用户签名控件；</li>
	// <li> <b>SIGN_PAGING_SIGNATURE</b> : 用户签名骑缝章控件；</li>
	// <li> <b>SIGN_OPINION</b> : 签署意见控件，用户需要根据配置的签署意见内容，完成对意见内容的确认；</li>
	// <li> <b>SIGN_VIRTUAL_COMBINATION</b> : 签批控件。内部包含最多4个特定控件（SIGN_SIGNATURE，SIGN_DATA,SIGN_MULTI_LINE_TEXT,SIGN_SELECTOR），本身不填充任何文字内容</li>
	// <li> <b>SIGN_MULTI_LINE_TEXT</b> : 多行文本，<font color="red">仅可用在签批控件内部作为组合控件，单独无法使用</font>，常用作批注附言</li>
	// <li> <b>SIGN_SELECTOR</b> : 选择器，<font color="red">仅可用在签批控件内部作为组合控件，单独无法使用</font>，常用作审批意见的选择</li>
	// </ul>
	//  
	// 注：` 表单域的控件不能作为印章和签名控件`
	ComponentType *string `json:"ComponentType,omitnil,omitempty" name:"ComponentType"`

	// **在绝对定位方式方式下**，ComponentName为控件名，长度不能超过20，只能由中文、字母、数字和下划线组成，可以在后续的操作中使用该名称来引用控件。
	// 
	// **在表单域定位方式下**，ComponentName不仅为控件名，也是表单域名称。
	// 
	// 注：`控件名可以在一个PDF中可以重复`
	// 
	// <a href="https://qcloudimg.tencent-cloud.cn/raw/93178569d07b4d7dbbe0967ae679e35c.png" target="_blank">点击查看ComponentName在模板页面的位置</a>
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// 如果是<b>填写控件</b>，ComponentRequired表示在填写页面此控件是否必填
	// <ul><li>false（默认）：可以不填写</li>
	// <li>true ：必须填写此填写控件</li></ul>
	// 如果是<b>签署控件</b>，签批控件中签署意见等可以不填写， 其他签署控件不受此字段影响
	ComponentRequired *bool `json:"ComponentRequired,omitnil,omitempty" name:"ComponentRequired"`

	// **在通过接口拉取控件信息场景下**，为出参参数，此控件归属的参与方的角色ID角色（即RecipientId），**发起合同时候不要填写此字段留空即可**
	ComponentRecipientId *string `json:"ComponentRecipientId,omitnil,omitempty" name:"ComponentRecipientId"`

	// <font color="red">【暂未使用】</font>控件所属文件的序号（取值为：0-N）。 目前单文件的情况下，值一直为0
	FileIndex *int64 `json:"FileIndex,omitnil,omitempty" name:"FileIndex"`

	// 控件生成的方式：
	// <ul><li> <b>NORMAL</b> : 绝对定位控件</li>
	// <li> <b>FIELD</b> : 表单域</li>
	// <li> <b>KEYWORD</b> : 关键字（设置关键字时，请确保PDF原始文件内是关键字以文字形式保存在PDF文件中，不支持对图片内文字进行关键字查找）</li></ul>
	GenerateMode *string `json:"GenerateMode,omitnil,omitempty" name:"GenerateMode"`

	// **在绝对定位方式和关键字定位方式下**，指定控件宽度，控件宽度是指控件在PDF文件中的宽度，单位为pt（点）。
	ComponentWidth *float64 `json:"ComponentWidth,omitnil,omitempty" name:"ComponentWidth"`

	// **在绝对定位方式和关键字定位方式下**，指定控件的高度， 控件高度是指控件在PDF文件中的高度，单位为pt（点）。
	ComponentHeight *float64 `json:"ComponentHeight,omitnil,omitempty" name:"ComponentHeight"`

	// **在绝对定位方式方式下**，指定控件所在PDF文件上的页码
	// **在使用文件发起的情况下**，绝对定位方式的填写控件和签署控件支持使用负数来指定控件在PDF文件上的页码，使用负数时，页码从最后一页开始。例如：ComponentPage设置为-1，即代表在PDF文件的最后一页，以此类推。
	// 
	// 注：
	// 1. 页码编号是从<font color="red">1</font>开始编号的。
	// 2.  <font color="red">页面编号不能超过PDF文件的页码总数</font>。如果指定的页码超过了PDF文件的页码总数，在填写和签署时会出现错误，导致无法正常进行操作。
	ComponentPage *int64 `json:"ComponentPage,omitnil,omitempty" name:"ComponentPage"`

	// **在绝对定位方式下**，可以指定控件横向位置的位置，单位为pt（点）。
	ComponentPosX *float64 `json:"ComponentPosX,omitnil,omitempty" name:"ComponentPosX"`

	// **在绝对定位方式下**，可以指定控件纵向位置的位置，单位为pt（点）。
	ComponentPosY *float64 `json:"ComponentPosY,omitnil,omitempty" name:"ComponentPosY"`

	// **在所有的定位方式下**，控件的扩展参数，为<font color="red">JSON格式</font>，不同类型的控件会有部分非通用参数。
	// 
	// <font color="red">ComponentType为TEXT、MULTI_LINE_TEXT时</font>，支持以下参数：
	// <ul><li> <b>Font</b>：目前只支持黑体、宋体、仿宋</li>
	// <li> <b>FontSize</b>： 范围6 :72</li>
	// <li> <b>FontAlign</b>： Left/Right/Center，左对齐/居中/右对齐</li>
	// <li> <b>FontColor</b>：字符串类型，格式为RGB颜色数字</li></ul>
	// <b>参数样例</b>：`{"FontColor":"255,0,0","FontSize":12}`
	// 
	// <font color="red">ComponentType为DATE时</font>，支持以下参数：
	// <ul><li> <b>Font</b>：目前只支持黑体、宋体、仿宋</li>
	// <li> <b>FontSize</b>： 范围6 :72</li></ul>
	// <b>参数样例</b>：`{"FontColor":"255,0,0","FontSize":12}`
	// 
	// <font color="red">ComponentType为WATERMARK时</font>，支持以下参数：
	// <ul><li> <b>Font</b>：目前只支持黑体、宋体、仿宋</li>
	// <li> <b>FontSize</b>： 范围6 :72</li>
	// <li> <b>Opacity</b>： 透明度，范围0 :1</li>
	// <li> <b>Rotate</b>： 水印旋转角度，范围0 :359</li>
	// <li> <b>Density</b>： 水印样式，1-宽松，2-标准（默认值），3-密集，</li>
	// <li> <b>Position</b>： 水印位置，None-平铺（默认值），LeftTop-左上，LeftBottom-左下，RightTop-右上，RightBottom-右下，Center-居中</li>
	// <li> <b>SubType</b>： 水印类型：CUSTOM_WATERMARK-自定义内容，PERSON_INFO_WATERMARK-访问者信息</li></ul>
	// <b>参数样例</b>：`"{\"Font\":\"黑体\",\"FontSize\":20,\"Opacity\":0.1,\"Density\":2,\"SubType\":\"PERSON_INFO_WATERMARK\"}"`
	// 
	// <font color="red">ComponentType为FILL_IMAGE时</font>，支持以下参数：
	// <ul><li> <b>NotMakeImageCenter</b>：bool。是否设置图片居中。false：居中（默认）。 true : 不居中</li>
	// <li> <b>FillMethod</b> : int. 填充方式。0-铺满（默认）；1-等比例缩放</li></ul>
	// 
	// <font color="red">ComponentType为SELECTOR时</font>，支持以下参数：
	// <ul><li> <b>WordWrap</b>：bool。是否支持选择控件内容自动折行合成。false：不支持（默认）。 true : 支持自动折行合成</li>
	// </ul>
	// 
	// <font color="red">ComponentType为SIGN_SIGNATURE、SIGN_PAGING_SIGNATURE类型时</font>，可以**ComponentTypeLimit**参数控制签署方式
	// <ul><li> <b>HANDWRITE</b> : 需要实时手写的手写签名</li>
	// <li> <b>HANDWRITTEN_ESIGN</b> : 长效手写签名， 是使用保存到个人中心的印章列表的手写签名(并且包含HANDWRITE)</li>
	// <li> <b>OCR_ESIGN</b> : AI智能识别手写签名</li>
	// <li> <b>ESIGN</b> : 个人印章类型</li>
	// <li> <b>SYSTEM_ESIGN</b> : 系统签名（该类型可以在用户签署时根据用户姓名一键生成一个签名来进行签署）</li>
	// <li> <b>IMG_ESIGN</b> : 图片印章(该类型支持用户在签署将上传的PNG格式的图片作为签名)</li></ul>
	// <b>参考样例</b>：`{"ComponentTypeLimit": ["SYSTEM_ESIGN"]}`
	// 印章的对应关系参考下图
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/ee0498856c060c065628a0c5ba780d6b.jpg)<br><br>
	// <font color="red">ComponentType为SIGN_SEAL 或者 SIGN_PAGING_SEAL类型时</font>，可以通过**ComponentTypeLimit**参数控制签署方签署时要使用的印章类型，支持指定以下印章类型
	// <ul><li> <b>OFFICIAL</b> :  企业公章</li>
	// <li> <b>CONTRACT</b> : 合同专用章</li>
	// <li> <b>FINANCE</b> : 财务专用章</li>
	// <li> <b>PERSONNEL</b> : 人事专用章</li>
	// <li> <b>OTHER</b> : 其他</li>
	// </ul>
	// <b>参考样例</b>：`{\"ComponentTypeLimit\":[\"PERSONNEL\",\"FINANCE\"]}` 表示改印章签署区,客户需使用人事专用章或财务专用章盖章签署。<br><br>
	// 
	// <font color="red">ComponentType为SIGN_DATE时</font>，支持以下参数：
	// <ul><li> <b>Font</b> :字符串类型目前只支持"黑体"、"宋体"、仿宋，如果不填默认为"黑体"</li>
	// <li> <b>FontSize</b> : 数字类型，范围6-72，默认值为12</li>
	// <li> <b>FontAlign</b> : 字符串类型，可取Left/Right/Center，对应左对齐/居中/右对齐</li>
	// <li> <b>Format</b> : 字符串类型，日期格式，必须是以下五种之一 “yyyy m d”，”yyyy年m月d日”，”yyyy/m/d”，”yyyy-m-d”，”yyyy.m.d”。</li>
	// <li> <b>Gaps</b> : 字符串类型，仅在Format为“yyyy m d”时起作用，格式为用逗号分开的两个整数，例如”2,2”，两个数字分别是日期格式的前后两个空隙中的空格个数</li></ul>
	// 如果extra参数为空，默认为”yyyy年m月d日”格式的居中日期
	// 特别地，如果extra中Format字段为空或无法被识别，则extra参数会被当作默认值处理（Font，FontSize，Gaps和FontAlign都不会起效）
	// <b>参数样例</b>： ` "{"Format":"yyyy m d","FontSize":12,"Gaps":"2,2", "FontAlign":"Right"}"`
	// 
	// <font color="red">ComponentType为SIGN_SEAL、SIGN_SIGNATURE类型时</font>，支持以下参数：
	// <ul><li> <b>PageRanges</b> :PageRange的数组，通过PageRanges属性设置该印章在PDF所有页面上盖章（适用于标书在所有页面盖章的情况）</li></ul>
	// <b>参数样例</b>：` "{"PageRanges":[{"BeginPage":1,"EndPage":-1}]}"`
	// 
	// <font color="red">签署印章旋转功能，当ComponentType为SIGN_SIGNATURE、SIGN_DATE、SIGN_SEAL时</font>，可以通过以下参数设置签署图片的旋转角度：
	// <ul><li> <b>Rotate</b>：旋转角度，支持范围：-360：360，为正整数时，为顺时针旋转；为负整数时，为逆时针旋转。</li>
	// <li> <b>RotateRelation</b>：旋转关联控件，用于指定关联旋转的控件。例如：让印章控件和签署日期控件按照印章控件为中心旋转（此时，设置印章控件的RotateRelation为日期控件的ComponentId，设置日期签署控件的RotateRelation为印章控件的ComponentId）。</li></ul>
	// <b>参数样例</b>：`{"Rotate":-30,"RotateRelation":"Component_Id1"}`
	// 
	// <font color="red">签署印章透明度功能设置，</font>当ComponentType为SIGN_SIGNATURE、SIGN_SEAL、SIGN_PAGING_SEAL、SIGN_LEGAL_PERSON_SEAL时，可以通过以下参数设置签署印章的透明度：
	// <ul><li> <b>Opacity</b>：印章透明度，支持范围：0-1，0.7表示70%的透明度，1表示无透明度</li></ul>
	// <b>参数样例</b>：`{"Opacity":0.7}`
	// 
	// <font color="red">签署印章大小功能设置，</font>当ComponentType为SIGN_SEAL、SIGN_PAGING_SEAL、SIGN_LEGAL_PERSON_SEAL时，可以通过以下参数设置签署时按照实际印章的大小进行签署，如果印章没有设置大小，那么默认会是4.2cm的印章大小：
	// <ul><li> <b>UseSealSize</b>：使用印章设置的大小盖章，true表示使用印章设置的大小盖章，false表示使用签署控件的大小进行盖章；不传则为false</li></ul>
	// <b>参数样例</b>：`{"UseSealSize":true}`
	// 
	// <font color="red">关键字模式下支持关键字找不到的情况下不进行报错的设置</font>
	// <ul><li> <b>IgnoreKeywordError</b> :1-关键字查找不到时不进行报错</li></ul>
	// 场景说明：如果使用关键字进行定位，但是指定的PDF文件中又没有设置的关键字时，发起合同会进行关键字是否存在的校验，如果关键字不存在，会进行报错返回。如果不希望进行报错，可以设置"IgnoreKeywordError"来忽略错误。请注意，如果关键字签署控件对应的签署方在整个PDF文件中一个签署控件都没有，还是会触发报错逻辑。
	// <b>参数样例</b>：` "{"IgnoreKeywordError":1}"`
	// 
	// 
	// <font color="red">ComponentType为SIGN_VIRTUAL_COMBINATION时</font>，支持以下参数：
	// <ul>
	// <li><b>Children:</b> 绝对定位模式下，用来指定此签批控件的组合子控件 </li>
	// <b>参数样例</b>：<br>`{"Children":["ComponentId_29","ComponentId_27","ComponentId_28","ComponentId_30"]}`
	// <li><b>ChildrenComponents:</b> 关键字定位模式下，用来指定此签批控件的组合子控件 </li>
	// ChildrenComponent结构体定义:
	// <table border="1">     <thead>         <tr>             <th>字段名称</th>             <th>类型</th>             <th>描述</th>         </tr>     </thead>     <tbody>         <tr>             <td>ComponentType</td>             <td>string</td>             <td>子控件类型-可选值:SIGN_SIGNATURE,SIGN_DATE,SIGN_SELECTOR,SIGN_MULTI_LINE_TEXT</td>         </tr>         <tr>             <td>ComponentName</td>             <td>string</td>             <td>子控件名称</td>         </tr>         <tr>             <td>Placeholder</td>             <td>string</td>             <td>子控件提示语</td>         </tr>         <tr>             <td>ComponentOffsetX</td>             <td>float</td>             <td>控件偏移位置X（相对于父控件（签批控件的ComponentX））</td>         </tr>         <tr>             <td>ComponentOffsetY</td>             <td>float</td>             <td>控件偏移位置Y 相对于父控件（签批控件的ComponentY））</td>         </tr>         <tr>             <td>ComponentWidth</td>             <td>float</td>             <td>控件宽</td>         </tr>         <tr>             <td>ComponentHeight</td>             <td>float</td>             <td>控件高</td>         </tr>         <tr>             <td>ComponentExtra</td>             <td>string</td>             <td>控件的附属信息，根据ComponentType设置</td>         </tr>     </tbody> </table>
	// <b>参数样例</b>：
	// 
	// <pre>
	// {
	//     "ChildrenComponents": [
	//         {
	//             "ComponentType": "SIGN_SIGNATURE",
	//             "ComponentName": "个人签名",
	//             "Placeholder": "请签名",
	//             "ComponentOffsetX": 10,
	//             "ComponentOffsetY": 30,
	//             "ComponentWidth": 119,
	//             "ComponentHeight": 43,
	//             "ComponentExtra": "{\"ComponentTypeLimit\":[\"SYSTEM_ESIGN\"]}"
	//         },
	//         {
	//             "ComponentType": "SIGN_SELECTOR",
	//             "ComponentName": "是否同意此协议",
	//             "Placeholder": "",
	//             "ComponentOffsetX": 50,
	//             "ComponentOffsetY": 130,
	//             "ComponentWidth": 120,
	//             "ComponentHeight": 43,
	//             "ComponentExtra": "{\"Values\":[\"同意\",\"不同意\",\"再想想\"],\"FontSize\":12,\"FontAlign\":\"Left\",\"Font\":\"黑体\",\"MultiSelect\":false}"
	//         },
	//         {
	//             "ComponentType": "SIGN_MULTI_LINE_TEXT",
	//             "ComponentName": "批注附言",
	//             "Placeholder": "",
	//             "ComponentOffsetX": 150,
	//             "ComponentOffsetY": 300,
	//             "ComponentWidth": 200,
	//             "ComponentHeight": 86,
	//             "ComponentExtra": ""
	//         }
	//     ]
	// }</pre>
	// </ul>
	// 
	ComponentExtra *string `json:"ComponentExtra,omitnil,omitempty" name:"ComponentExtra"`

	// 控件填充vaule，ComponentType和传入值类型对应关系：
	// <ul><li> <b>TEXT</b> : 文本内容</li>
	// <li> <b>MULTI_LINE_TEXT</b> : 文本内容， 可以用  \n 来控制换行位置</li>
	// <li> <b>CHECK_BOX</b> : true/false</li>
	// <li> <b>FILL_IMAGE、ATTACHMENT</b> : 附件的FileId，需要通过UploadFiles接口上传获取</li>
	// <li> <b>SELECTOR</b> : 选项值</li>
	// <li> <b>DYNAMIC_TABLE</b>  - 传入json格式的表格内容，详见说明：[数据表格](https://qian.tencent.com/developers/company/dynamic_table)</li>
	// <li> <b>DATE</b> : 格式化：xxxx年xx月xx日（例如：2024年05月28日）</li>
	// <li> <b>SIGN_SEAL</b> : 印章ID，于控制台查询获取，[点击查看在控制上的位置](https://qcloudimg.tencent-cloud.cn/raw/cd403a5b949fce197fd9e88bb6db1517.png)</li>
	// <li> <b>SIGN_PAGING_SEAL</b> : 可以指定印章ID，于控制台查询获取，[点击查看在控制上的位置](https://qcloudimg.tencent-cloud.cn/raw/cd403a5b949fce197fd9e88bb6db1517.png)</li></ul>
	// 
	// 
	// <b>控件值约束说明</b>：
	// <table> <thead> <tr> <th>特殊控件</th> <th>填写约束</th> </tr> </thead> <tbody> <tr> <td>企业全称控件</td> <td>企业名称中文字符中文括号</td> </tr> <tr> <td>统一社会信用代码控件</td> <td>企业注册的统一社会信用代码</td> </tr> <tr> <td>法人名称控件</td> <td>最大50个字符，2到25个汉字或者1到50个字母</td> </tr> <tr> <td>签署意见控件</td> <td>签署意见最大长度为50字符</td> </tr> <tr> <td>签署人手机号控件</td> <td>国内手机号 13,14,15,16,17,18,19号段长度11位</td> </tr> <tr> <td>签署人身份证控件</td> <td>合法的身份证号码检查</td> </tr> <tr> <td>控件名称</td> <td>控件名称最大长度为20字符，不支持表情</td> </tr> <tr> <td>单行文本控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>多行文本控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>勾选框控件</td> <td>选择填字符串true，不选填字符串false</td> </tr> <tr> <td>选择器控件</td> <td>同单行文本控件约束，填写选择值中的字符串</td> </tr> <tr> <td>数字控件</td> <td>请输入有效的数字(可带小数点)</td> </tr> <tr> <td>日期控件</td> <td>格式：yyyy年mm月dd日</td> </tr> <tr> <td>附件控件</td> <td>JPG或PNG图片，上传数量限制，1到6个，最大6个附件，填写上传的资源ID</td> </tr> <tr> <td>图片控件</td> <td>JPG或PNG图片，填写上传的图片资源ID</td> </tr> <tr> <td>邮箱控件</td> <td>有效的邮箱地址, w3c标准</td> </tr> <tr> <td>地址控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>省市区控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>性别控件</td> <td>选择值中的字符串</td> </tr> <tr> <td>学历控件</td> <td>选择值中的字符串</td> </tr><tr> <td>水印控件</td> <td>水印控件设置为CUSTOM_WATERMARK类型时的水印内容</td> </tr> </tbody> </table>
	// 注：   `部分特殊控件需要在控制台配置模板形式创建`
	ComponentValue *string `json:"ComponentValue,omitnil,omitempty" name:"ComponentValue"`

	// <font color="red">【暂未使用】</font>日期签署控件的字号，默认为 12
	ComponentDateFontSize *int64 `json:"ComponentDateFontSize,omitnil,omitempty" name:"ComponentDateFontSize"`

	// <font color="red">【暂未使用】</font>控件归属的文档的ID， **发起合同时候不要填写此字段留空即可**
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`

	// <font color="red">【暂未使用】</font>控件描述，用户自定义，不影响合同发起流程
	ComponentDescription *string `json:"ComponentDescription,omitnil,omitempty" name:"ComponentDescription"`

	// **如果控件是关键字定位方式**，可以对关键字定位出来的区域进行横坐标方向的调整，单位为pt（点）。例如，如果关键字定位出来的区域偏左或偏右，可以通过调整横坐标方向的参数来使控件位置更加准确。
	// 注意： `向左调整设置为负数， 向右调整设置成正数`
	OffsetX *float64 `json:"OffsetX,omitnil,omitempty" name:"OffsetX"`

	// **如果控件是关键字定位方式**，可以对关键字定位出来的区域进行纵坐标方向的调整，单位为pt（点）。例如，如果关键字定位出来的区域偏上或偏下，可以通过调整纵坐标方向的参数来使控件位置更加准确。
	// 注意： `向上调整设置为负数， 向下调整设置成正数`
	OffsetY *float64 `json:"OffsetY,omitnil,omitempty" name:"OffsetY"`

	// <font color="red">【暂未使用】</font>第三方应用集成平台模板控件 ID 标识
	ChannelComponentId *string `json:"ChannelComponentId,omitnil,omitempty" name:"ChannelComponentId"`

	// **如果控件是关键字定位方式**，指定关键字排序规则时，可以选择Positive或Reverse两种排序方式。
	// <ul><li> <b>Positive</b> :表示正序，即根据关键字在PDF文件内的顺序进行排列</li>
	// <li> <b>Reverse</b> :表示倒序，即根据关键字在PDF文件内的反序进行排列</li></ul>
	// 
	// 在指定KeywordIndexes时，如果使用Positive排序方式，0代表在PDF内查找内容时，查找到的第一个关键字；如果使用Reverse排序方式，0代表在PDF内查找内容时，查找到的最后一个关键字。
	KeywordOrder *string `json:"KeywordOrder,omitnil,omitempty" name:"KeywordOrder"`

	// **如果控件是关键字定位方式**，在KeywordPage中指定关键字页码时，将只会在该页码中查找关键字，非该页码的关键字将不会查询出来。如果不设置查找所有页面中的关键字。
	KeywordPage *int64 `json:"KeywordPage,omitnil,omitempty" name:"KeywordPage"`

	// **如果控件是关键字定位方式**，关键字生成的区域的对齐方式， 可以设置下面的值
	// <ul><li> <b>Middle</b> :居中</li>
	// <li> <b>Below</b> :正下方</li>
	// <li> <b>Right</b> :正右方</li>
	// <li> <b>LowerRight</b> :右下角</li>
	// <li> <b>UpperRight</b> :右上角。</li></ul>
	// 示例：如果设置Middle的关键字盖章，则印章的中心会和关键字的中心重合，如果设置Below，则印章在关键字的正下方
	RelativeLocation *string `json:"RelativeLocation,omitnil,omitempty" name:"RelativeLocation"`

	// **如果控件是关键字定位方式**，关键字索引是指在PDF文件中存在多个相同的关键字时，通过索引指定使用哪一个关键字作为最后的结果。可以通过指定多个索引来同时使用多个关键字。例如，[0,2]表示使用PDF文件内第1个和第3个关键字位置作为最后的结果。
	// 
	// 注意：关键字索引是从0开始计数的
	KeywordIndexes []*int64 `json:"KeywordIndexes,omitnil,omitempty" name:"KeywordIndexes"`

	// 填写控件在腾讯电子签小程序填写界面展示的提示信息，例如，在身份证号码填写控件中，提示信息可以设置成“请输入18位身份证号码”。
	// 注：`签署控件设置此字段无效`
	Placeholder *string `json:"Placeholder,omitnil,omitempty" name:"Placeholder"`

	// **web嵌入发起合同场景下**， 是否锁定填写和签署控件值不允许嵌入页面进行编辑
	// <ul><li>false（默认）：不锁定控件值，允许在页面编辑控件值</li>
	// <li>true：锁定控件值，在页面无法编辑控件值</li></ul>
	LockComponentValue *bool `json:"LockComponentValue,omitnil,omitempty" name:"LockComponentValue"`

	// **web嵌入发起合同场景下**，是否禁止移动和删除填写和签署控件
	// <ul><li> <b>false（默认）</b> :可以移动和删除控件</li>
	// <li> <b>true</b> : 禁止移动和删除控件</li></ul>
	ForbidMoveAndDelete *bool `json:"ForbidMoveAndDelete,omitnil,omitempty" name:"ForbidMoveAndDelete"`
}

type ComponentLimit struct {
	// 控件类型，支持以下类型
	// <ul><li>SIGN_SEAL : 印章控件</li>
	// <li>SIGN_PAGING_SEAL : 骑缝章控件</li>
	// <li>SIGN_LEGAL_PERSON_SEAL : 企业法定代表人控件</li>
	// <li>SIGN_SIGNATURE : 用户签名控件</li></ul>
	ComponentType *string `json:"ComponentType,omitnil,omitempty" name:"ComponentType"`

	// 签署控件类型的值(可选)，用与限制签署时印章或者签名的选择范围
	// 
	// 1.当ComponentType 是 SIGN_SEAL 或者 SIGN_PAGING_SEAL 时可传入企业印章Id（支持多个）或者以下印章类型
	// 
	// <ul><li> <b>OFFICIAL</b> :  企业公章</li>
	// <li> <b>CONTRACT</b> : 合同专用章</li>
	// <li> <b>FINANCE</b> : 财务专用章</li>
	// <li> <b>PERSONNEL</b> : 人事专用章</li>
	// <li> <b>OTHER</b> : 其他</li>
	// </ul>
	// 
	// 注：`限制印章控件或骑缝章控件情况下,仅本企业签署方可以指定具体印章（通过传递ComponentValue,支持多个),他方企业签署人只能限制类型.若同时指定了印章类型和印章Id,以印章Id为主,印章类型会被忽略`
	// 
	// 2.当ComponentType 是 SIGN_SIGNATURE 时可传入以下类型（支持多个）
	// 
	// <ul><li>HANDWRITE : 需要实时手写的手写签名</li>
	// <li>HANDWRITTEN_ESIGN : 长效手写签名， 是使用保存到个人中心的印章列表的手写签名(并且包含HANDWRITE)</li>
	// <li>OCR_ESIGN : OCR印章（智慧手写签名）</li>
	// <li>ESIGN : 个人印章</li>
	// <li>SYSTEM_ESIGN : 系统印章</li></ul>
	// 
	// 3.当ComponentType 是 SIGN_LEGAL_PERSON_SEAL 时无需传递此参数。
	ComponentValue []*string `json:"ComponentValue,omitnil,omitempty" name:"ComponentValue"`
}

// Predefined struct for user
type CreateBatchInitOrganizationUrlRequestParams struct {
	// 应用相关信息。 此接口Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 初始化操作类型
	// <ul>
	// <li>CREATE_SEAL : 创建印章</li>
	// <li>OPEN_AUTO_SIGN :开通企业自动签署</li>
	// <li>PARTNER_AUTO_SIGN_AUTH :合作方企业或应用平台方授权自动签</li>
	// </ul>
	OperateTypes []*string `json:"OperateTypes,omitnil,omitempty" name:"OperateTypes"`

	// 批量操作的企业列表在第三方平台的企业Id列表，即ProxyOrganizationOpenId列表,最大支持50个
	ProxyOrganizationOpenIds []*string `json:"ProxyOrganizationOpenIds,omitnil,omitempty" name:"ProxyOrganizationOpenIds"`

	// 当操作类型包含 PARTNER_AUTO_SIGN_AUTH 且是给应用平台方授权自动签时传true。
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/f9aba7c999a6d79ada20b4384520e120.png)
	IsAuthorizePlatformApplication *bool `json:"IsAuthorizePlatformApplication,omitnil,omitempty" name:"IsAuthorizePlatformApplication"`

	// 被授权的合作方企业在第三方平台子客企业标识，即ProxyOrganizationOpenId，当操作类型包含 PARTNER_AUTO_SIGN_AUTH 且要进行合作方企业授权自动签时必传。
	// 
	// 
	AuthorizedProxyOrganizationOpenId *string `json:"AuthorizedProxyOrganizationOpenId,omitnil,omitempty" name:"AuthorizedProxyOrganizationOpenId"`
}

type CreateBatchInitOrganizationUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 初始化操作类型
	// <ul>
	// <li>CREATE_SEAL : 创建印章</li>
	// <li>OPEN_AUTO_SIGN :开通企业自动签署</li>
	// <li>PARTNER_AUTO_SIGN_AUTH :合作方企业或应用平台方授权自动签</li>
	// </ul>
	OperateTypes []*string `json:"OperateTypes,omitnil,omitempty" name:"OperateTypes"`

	// 批量操作的企业列表在第三方平台的企业Id列表，即ProxyOrganizationOpenId列表,最大支持50个
	ProxyOrganizationOpenIds []*string `json:"ProxyOrganizationOpenIds,omitnil,omitempty" name:"ProxyOrganizationOpenIds"`

	// 当操作类型包含 PARTNER_AUTO_SIGN_AUTH 且是给应用平台方授权自动签时传true。
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/f9aba7c999a6d79ada20b4384520e120.png)
	IsAuthorizePlatformApplication *bool `json:"IsAuthorizePlatformApplication,omitnil,omitempty" name:"IsAuthorizePlatformApplication"`

	// 被授权的合作方企业在第三方平台子客企业标识，即ProxyOrganizationOpenId，当操作类型包含 PARTNER_AUTO_SIGN_AUTH 且要进行合作方企业授权自动签时必传。
	// 
	// 
	AuthorizedProxyOrganizationOpenId *string `json:"AuthorizedProxyOrganizationOpenId,omitnil,omitempty" name:"AuthorizedProxyOrganizationOpenId"`
}

func (r *CreateBatchInitOrganizationUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchInitOrganizationUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "OperateTypes")
	delete(f, "ProxyOrganizationOpenIds")
	delete(f, "IsAuthorizePlatformApplication")
	delete(f, "AuthorizedProxyOrganizationOpenId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchInitOrganizationUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchInitOrganizationUrlResponseParams struct {
	// 小程序路径，有效时间为7天
	MiniAppPath *string `json:"MiniAppPath,omitnil,omitempty" name:"MiniAppPath"`

	// 操作长链，有效时间为7天
	OperateLongUrl *string `json:"OperateLongUrl,omitnil,omitempty" name:"OperateLongUrl"`

	// 操作短链，有效时间为7天
	OperateShortUrl *string `json:"OperateShortUrl,omitnil,omitempty" name:"OperateShortUrl"`

	// 操作二维码，有效时间为7天
	QRCodeUrl *string `json:"QRCodeUrl,omitnil,omitempty" name:"QRCodeUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBatchInitOrganizationUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchInitOrganizationUrlResponseParams `json:"Response"`
}

func (r *CreateBatchInitOrganizationUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchInitOrganizationUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchOrganizationAuthorizationUrlRequestParams struct {
	// 应用相关信息。 此接口Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 组织机构超管姓名。 在注册流程中，必须是超管本人进行操作。此参数需要跟[创建子企业批量认证链接](https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks)中 AdminName 保持一致。
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 组织机构超管手机号。 在注册流程中，必须是超管本人进行操作。此参数需要跟[创建子企业批量认证链接](https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks)中 Admin Mobile保持一致。
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`

	// 企业批量认证链接的子任务 SubTaskId，该 SubTaskId 是通过接口 查询企业批量认证链接 DescribeBatchOrganizationRegistrationUrls 获得。此参数需与超管个人三要素（AdminName，AdminMobile，AdminIdCardNumber）配合使用。若 SubTaskId 不属于传入的超级管理员，将进行筛选。
	SubTaskIds []*string `json:"SubTaskIds,omitnil,omitempty" name:"SubTaskIds"`

	// 组织机构超管证件类型支持以下类型
	// - ID_CARD : 居民身份证 (默认值)
	// -  HONGKONG_AND_MACAO : 港澳居民来往内地通行证
	// - HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)
	// 此参数需要跟[创建子企业批量认证链接](https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks)中 AdminIdCardType保持一致。
	AdminIdCardType *string `json:"AdminIdCardType,omitnil,omitempty" name:"AdminIdCardType"`

	// 组织机构超管证件号。 在注册流程中，必须是超管本人进行操作。此参数需要跟[创建子企业批量认证链接](https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks)中 AdminIdCardNumber保持一致。
	AdminIdCardNumber *string `json:"AdminIdCardNumber,omitnil,omitempty" name:"AdminIdCardNumber"`

	// 要跳转的链接类型<ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li><li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li><li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li><li>**QR_CODE**： 跳转电子签小程序的http_url的二维码形式,  可以在页面展示适合此类型</li></ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

type CreateBatchOrganizationAuthorizationUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 组织机构超管姓名。 在注册流程中，必须是超管本人进行操作。此参数需要跟[创建子企业批量认证链接](https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks)中 AdminName 保持一致。
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 组织机构超管手机号。 在注册流程中，必须是超管本人进行操作。此参数需要跟[创建子企业批量认证链接](https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks)中 Admin Mobile保持一致。
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`

	// 企业批量认证链接的子任务 SubTaskId，该 SubTaskId 是通过接口 查询企业批量认证链接 DescribeBatchOrganizationRegistrationUrls 获得。此参数需与超管个人三要素（AdminName，AdminMobile，AdminIdCardNumber）配合使用。若 SubTaskId 不属于传入的超级管理员，将进行筛选。
	SubTaskIds []*string `json:"SubTaskIds,omitnil,omitempty" name:"SubTaskIds"`

	// 组织机构超管证件类型支持以下类型
	// - ID_CARD : 居民身份证 (默认值)
	// -  HONGKONG_AND_MACAO : 港澳居民来往内地通行证
	// - HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)
	// 此参数需要跟[创建子企业批量认证链接](https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks)中 AdminIdCardType保持一致。
	AdminIdCardType *string `json:"AdminIdCardType,omitnil,omitempty" name:"AdminIdCardType"`

	// 组织机构超管证件号。 在注册流程中，必须是超管本人进行操作。此参数需要跟[创建子企业批量认证链接](https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks)中 AdminIdCardNumber保持一致。
	AdminIdCardNumber *string `json:"AdminIdCardNumber,omitnil,omitempty" name:"AdminIdCardNumber"`

	// 要跳转的链接类型<ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li><li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li><li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li><li>**QR_CODE**： 跳转电子签小程序的http_url的二维码形式,  可以在页面展示适合此类型</li></ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

func (r *CreateBatchOrganizationAuthorizationUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchOrganizationAuthorizationUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "AdminName")
	delete(f, "AdminMobile")
	delete(f, "SubTaskIds")
	delete(f, "AdminIdCardType")
	delete(f, "AdminIdCardNumber")
	delete(f, "Endpoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchOrganizationAuthorizationUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchOrganizationAuthorizationUrlResponseParams struct {
	// 批量企业注册链接-单链接包含多条认证流，根据Endpoint的不同设置，返回不同的链接地址。失效时间：7天跳转链接, 链接的有效期根据企业,员工状态和终端等有区别, 可以参考下表<table> <thead> <tr> <th>Endpoint</th> <th>示例</th> <th>链接有效期限</th> </tr> </thead>  <tbody> <tr> <td>HTTP</td> <td>https://res.ess.tencent.cn/cdn/h5-activity-dev/jump-mp.html?to=AUTHORIZATION_ENTERPRISE_FOR_BATCH_SUBMIT&shortKey=yDCHHURDfBxSB2rj2Bfa</td> <td>7天</td> </tr> <tr> <td>HTTP_SHORT_URL</td> <td>https://test.essurl.cn/8gDKUBAWK8</td> <td>7天</td> </tr> <tr> <td>APP</td> <td>pages/guide/index?to=AUTHORIZATION_ENTERPRISE_FOR_BATCH_SUBMIT&shortKey=yDCHpURDfR6iEkdpsDde</td> <td>7天</td> </tr><tr> <td>QR_CODE</td> <td>https://dyn.test.ess.tencent.cn/imgs/qrcode_urls/authorization_enterprise_for_batch_submit/yDCHHUUckpbdauq9UEjnoFDCCumAMmv1.png</td> <td>7天</td> </tr> </tbody> </table>注： `1.创建的链接应避免被转义，如：&被转义为\u0026；如使用Postman请求后，请选择响应类型为 JSON，否则链接将被转义`
	AuthUrl *string `json:"AuthUrl,omitnil,omitempty" name:"AuthUrl"`

	// 认证流认证失败信息
	ErrorMessages []*string `json:"ErrorMessages,omitnil,omitempty" name:"ErrorMessages"`

	// 链接过期时间，为 7 天后，创建时间，格式为Unix标准时间戳（秒）。
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBatchOrganizationAuthorizationUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchOrganizationAuthorizationUrlResponseParams `json:"Response"`
}

func (r *CreateBatchOrganizationAuthorizationUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchOrganizationAuthorizationUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchOrganizationRegistrationTasksRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 当前应用下子企业的组织机构注册信息。最多支持查询10子企业。
	RegistrationOrganizations []*RegistrationOrganizationInfo `json:"RegistrationOrganizations,omitnil,omitempty" name:"RegistrationOrganizations"`

	// 生成链接的类型：
	// <ul><li>**PC**：(默认)web控制台链接, 需要在PC浏览器中打开</li>
	// <li>**CHANNEL**：H5跳转到电子签小程序链接, 一般用于发送短信中带的链接, 打开后进入腾讯电子签小程序</li>
	// <li>**SHORT_URL**：H5跳转到电子签小程序链接的短链形式, 一般用于发送短信中带的链接, 打开后进入腾讯电子签小程序</li>
	// <li>**APP**：第三方APP或小程序跳转电子签小程序链接, 一般用于贵方小程序或者APP跳转过来,  打开后进入腾讯电子签小程序</li>
	// <li>**H5**：第三方H5跳转到电子签H5长链接, 一般用于贵方H5跳转过来,  打开后进入腾讯电子签H5页面</li>
	// <li>**SHORT_H5**：第三方H5跳转到电子签H5短链接, 一般用于贵方H5跳转过来,  打开后进入腾讯电子签H5页面</li></ul>
	// 
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 认证链接使用单链接还是多链接模式<ul><li>0 - 多链接(默认)，指批量生成链接， 每一个企业会拥有一个认证链接，然后分别认证</li><li>1 - 单链接 ， 指批量生成链接，然后会将多个链接聚合成一个链接，进行认证</li></ul>p.s.请注意， 如果使用单链接的模式并且认证方式是授权书方式的时候，必须在接口中传递超管授权书。
	BatchAuthMethod *uint64 `json:"BatchAuthMethod,omitnil,omitempty" name:"BatchAuthMethod"`
}

type CreateBatchOrganizationRegistrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 当前应用下子企业的组织机构注册信息。最多支持查询10子企业。
	RegistrationOrganizations []*RegistrationOrganizationInfo `json:"RegistrationOrganizations,omitnil,omitempty" name:"RegistrationOrganizations"`

	// 生成链接的类型：
	// <ul><li>**PC**：(默认)web控制台链接, 需要在PC浏览器中打开</li>
	// <li>**CHANNEL**：H5跳转到电子签小程序链接, 一般用于发送短信中带的链接, 打开后进入腾讯电子签小程序</li>
	// <li>**SHORT_URL**：H5跳转到电子签小程序链接的短链形式, 一般用于发送短信中带的链接, 打开后进入腾讯电子签小程序</li>
	// <li>**APP**：第三方APP或小程序跳转电子签小程序链接, 一般用于贵方小程序或者APP跳转过来,  打开后进入腾讯电子签小程序</li>
	// <li>**H5**：第三方H5跳转到电子签H5长链接, 一般用于贵方H5跳转过来,  打开后进入腾讯电子签H5页面</li>
	// <li>**SHORT_H5**：第三方H5跳转到电子签H5短链接, 一般用于贵方H5跳转过来,  打开后进入腾讯电子签H5页面</li></ul>
	// 
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 认证链接使用单链接还是多链接模式<ul><li>0 - 多链接(默认)，指批量生成链接， 每一个企业会拥有一个认证链接，然后分别认证</li><li>1 - 单链接 ， 指批量生成链接，然后会将多个链接聚合成一个链接，进行认证</li></ul>p.s.请注意， 如果使用单链接的模式并且认证方式是授权书方式的时候，必须在接口中传递超管授权书。
	BatchAuthMethod *uint64 `json:"BatchAuthMethod,omitnil,omitempty" name:"BatchAuthMethod"`
}

func (r *CreateBatchOrganizationRegistrationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchOrganizationRegistrationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "RegistrationOrganizations")
	delete(f, "Endpoint")
	delete(f, "BatchAuthMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchOrganizationRegistrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchOrganizationRegistrationTasksResponseParams struct {
	// 生成注册链接的任务ID，后序根据这个任务ID， 调用<a href="https://qian.tencent.com/developers/partnerApis/accounts/DescribeBatchOrganizationRegistrationUrls" target="_blank">查询子企业批量认证链接</a>获取生成的链接，发给对应的客户使用。
	// 
	// 注：`如果有错误，则不会返回任务ID`
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 批量生成企业认证链接的详细错误信息，顺序与输入参数子企业列表顺序一致。
	// <ul>
	// <li>如果所有企业认证链接都成功生成，将不返回错误信息</li>
	// <li>如果存在任何错误，将返回具体的错误描述。（没有错误的企业返回空字符串）</li>
	// </ul>
	ErrorMessages []*string `json:"ErrorMessages,omitnil,omitempty" name:"ErrorMessages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBatchOrganizationRegistrationTasksResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchOrganizationRegistrationTasksResponseParams `json:"Response"`
}

func (r *CreateBatchOrganizationRegistrationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchOrganizationRegistrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChannelFlowEvidenceReportRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定申请的报告类型，可选类型如下：
	// <ul><li> **0** :合同签署报告（默认）</li>
	// <li> **1** :公证处核验报告</li></ul>
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`
}

type CreateChannelFlowEvidenceReportRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定申请的报告类型，可选类型如下：
	// <ul><li> **0** :合同签署报告（默认）</li>
	// <li> **1** :公证处核验报告</li></ul>
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`
}

func (r *CreateChannelFlowEvidenceReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChannelFlowEvidenceReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowId")
	delete(f, "Operator")
	delete(f, "ReportType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChannelFlowEvidenceReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChannelFlowEvidenceReportResponseParams struct {
	// 出证报告 ID，可用于<a href="https://qian.tencent.com/developers/partnerApis/certificate/DescribeChannelFlowEvidenceReport" target="_blank">获取出证报告任务执行结果</a>查询出证任务结果和出证PDF的下载URL
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// 出证任务执行的状态, 状态含义如下：
	// 
	// <ul><li>**EvidenceStatusExecuting**：  出证任务在执行中</li>
	// <li>**EvidenceStatusSuccess**：  出证任务执行成功</li>
	// <li>**EvidenceStatusFailed** ： 出证任务执行失败</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 废除，字段无效
	ReportUrl *string `json:"ReportUrl,omitnil,omitempty" name:"ReportUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateChannelFlowEvidenceReportResponse struct {
	*tchttp.BaseResponse
	Response *CreateChannelFlowEvidenceReportResponseParams `json:"Response"`
}

func (r *CreateChannelFlowEvidenceReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChannelFlowEvidenceReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChannelOrganizationInfoChangeUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 企业信息变更类型，可选类型如下：
	// <ul><li>**1**：企业超管变更, 可以将超管换成同企业的其他员工</li>
	// <li>**2**：企业基础信息变更, 可以改企业名称 , 所在地址 , 法人名字等信息</li></ul>
	ChangeType *uint64 `json:"ChangeType,omitnil,omitempty" name:"ChangeType"`

	// 变更链接类型：
	// <ul>
	// <li>**WEIXINAPP** : 创建变更短链。需要在移动端打开，会跳转到微信腾讯电子签小程序进行更换。（默认）</li>
	// <li>**APP** : 创建变更小程序链接，可从第三方App或者小程序跳转到微信腾讯电子签小程序进行更换。</li>
	// </ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

type CreateChannelOrganizationInfoChangeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 企业信息变更类型，可选类型如下：
	// <ul><li>**1**：企业超管变更, 可以将超管换成同企业的其他员工</li>
	// <li>**2**：企业基础信息变更, 可以改企业名称 , 所在地址 , 法人名字等信息</li></ul>
	ChangeType *uint64 `json:"ChangeType,omitnil,omitempty" name:"ChangeType"`

	// 变更链接类型：
	// <ul>
	// <li>**WEIXINAPP** : 创建变更短链。需要在移动端打开，会跳转到微信腾讯电子签小程序进行更换。（默认）</li>
	// <li>**APP** : 创建变更小程序链接，可从第三方App或者小程序跳转到微信腾讯电子签小程序进行更换。</li>
	// </ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

func (r *CreateChannelOrganizationInfoChangeUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChannelOrganizationInfoChangeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "ChangeType")
	delete(f, "Endpoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChannelOrganizationInfoChangeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChannelOrganizationInfoChangeUrlResponseParams struct {
	// 创建的企业信息变更链接。需要在移动端打开，会跳转到微信腾讯电子签小程序进行更换。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 链接过期时间。链接7天有效。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateChannelOrganizationInfoChangeUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateChannelOrganizationInfoChangeUrlResponseParams `json:"Response"`
}

func (r *CreateChannelOrganizationInfoChangeUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChannelOrganizationInfoChangeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChannelSubOrganizationActiveRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要进行激活或者续期的子客企业OrganizationOpenId列表，请确保所有列出的子客企业均已完成认证。
	SubOrganizationOpenIds []*string `json:"SubOrganizationOpenIds,omitnil,omitempty" name:"SubOrganizationOpenIds"`

	// 操作类型，可以选择如下：
	// 
	// **false**：（默认）激活子客企业
	// **true**：续期子客企业
	Renew *bool `json:"Renew,omitnil,omitempty" name:"Renew"`
}

type CreateChannelSubOrganizationActiveRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要进行激活或者续期的子客企业OrganizationOpenId列表，请确保所有列出的子客企业均已完成认证。
	SubOrganizationOpenIds []*string `json:"SubOrganizationOpenIds,omitnil,omitempty" name:"SubOrganizationOpenIds"`

	// 操作类型，可以选择如下：
	// 
	// **false**：（默认）激活子客企业
	// **true**：续期子客企业
	Renew *bool `json:"Renew,omitnil,omitempty" name:"Renew"`
}

func (r *CreateChannelSubOrganizationActiveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChannelSubOrganizationActiveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "SubOrganizationOpenIds")
	delete(f, "Renew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChannelSubOrganizationActiveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChannelSubOrganizationActiveResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateChannelSubOrganizationActiveResponse struct {
	*tchttp.BaseResponse
	Response *CreateChannelSubOrganizationActiveResponseParams `json:"Response"`
}

func (r *CreateChannelSubOrganizationActiveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChannelSubOrganizationActiveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloseOrganizationUrlRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CreateCloseOrganizationUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *CreateCloseOrganizationUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloseOrganizationUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloseOrganizationUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloseOrganizationUrlResponseParams struct {
	// 链接有效期，unix时间戳，精确到秒
	ExpiredOn *int64 `json:"ExpiredOn,omitnil,omitempty" name:"ExpiredOn"`

	// H5跳转到电子签小程序链接, 一般用于发送短信中带的链接, 打开后进入腾讯电子签小程序
	LongUrl *string `json:"LongUrl,omitnil,omitempty" name:"LongUrl"`

	// H5跳转到电子签小程序链接的短链形式, 一般用于发送短信中带的链接, 打开后进入腾讯电子签小程序
	ShortUrl *string `json:"ShortUrl,omitnil,omitempty" name:"ShortUrl"`

	// APP或小程序跳转电子签小程序链接, 一般用于客户小程序或者APP跳转过来, 打开后进入腾讯电子签小程序
	MiniAppPath *string `json:"MiniAppPath,omitnil,omitempty" name:"MiniAppPath"`

	// 二维码链接
	QrcodeUrl *string `json:"QrcodeUrl,omitnil,omitempty" name:"QrcodeUrl"`

	// 直接跳转至电子签小程序的二维码链接，无需通过中转页。您需要自行将其转换为二维码，使用微信扫码后可直接进入。请注意，直接点击链接是无效的。
	WeixinQrcodeUrl *string `json:"WeixinQrcodeUrl,omitnil,omitempty" name:"WeixinQrcodeUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloseOrganizationUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloseOrganizationUrlResponseParams `json:"Response"`
}

func (r *CreateCloseOrganizationUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloseOrganizationUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsoleLoginUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  <a href="https://qcloudimg.tencent-cloud.cn/raw/a71872de3d540d55451e3e73a2ad1a6e.png" target="_blank">Agent.AppId</a></li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>注:
	// `1. 企业激活时， 此时的Agent.ProxyOrganizationOpenId将会是企业激活后企业的唯一标识，建议开发者保存企业ProxyOrganizationOpenId，后续各项接口调用皆需要此参数。 `
	// `2. 员工认证时， 此时的Agent.ProxyOperator.OpenId将会是员工认证加入企业后的唯一标识，建议开发者保存此员工的OpenId，后续各项接口调用皆需要此参数。 `
	// `3. 同渠道应用(Agent.AppId)下，企业唯一标识ProxyOrganizationOpenId需要保持唯一，员工唯一标识OpenId也要保持唯一 (而不是企业下唯一)。 `
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 第三方平台子客的企业名称，请确认该企业名称与企业营业执照中注册的名称完全一致。
	// <font color="red">
	// 在测试环境联调的过程中，企业名称请统一加上“测试”二字，如：典子谦示例企业测试，否则将无法审核通过。
	// 企业名称请使用以下名称, 以下名称可以不用走收录。
	// **子客测试专用企业1 - 子客测试专用企业9**</font>
	// 
	// 注:
	//  `1. 如果名称中包含英文括号()，请使用中文括号（）代替。`
	//  `2、该名称需要与Agent.ProxyOrganizationOpenId相匹配,  企业激活后Agent.ProxyOrganizationOpenId会跟此企业名称一一绑定; 如果您的企业已经在认证授权中或者激活完成，这里修改子客企业名字将不会生效。 `
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil,omitempty" name:"ProxyOrganizationName"`

	// 子客企业统一社会信用代码，最大长度200个字符
	// 注意：`如果您的企业已经在认证授权中或者激活完成，这里修改子客企业名字将不会生效`。
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// 子客企业员工的姓名，最大长度50个字符,  员工的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：`该姓名需要和Agent.ProxyOperator.OpenId相匹配,  当员工完成认证后该姓名会和Agent.ProxyOperator.OpenId一一绑定, 若员工已认证加入企业，这里修改经办人名字传入将不会生效`
	ProxyOperatorName *string `json:"ProxyOperatorName,omitnil,omitempty" name:"ProxyOperatorName"`

	// 子客企业员工的手机码,  支持国内手机号11位数字(无需加+86前缀或其他字符)。注：`该手机号需要和Agent.ProxyOperator.OpenId相匹配,  当员工完成认证后该手机号会和Agent.ProxyOperator.OpenId一一绑定, 若员工已认证加入企业，这里修改经办人手机号传入将不会生效`
	ProxyOperatorMobile *string `json:"ProxyOperatorMobile,omitnil,omitempty" name:"ProxyOperatorMobile"`

	// Web控制台登录后进入的功能模块,  支持的模块包括：
	// <ul>
	// <li> **空值** :(默认)企业中心模块</li>
	// <li> **DOCUMENT** :合同管理模块</li>
	// <li> **TEMPLATE** :企业模板管理模块</li>
	// <li> **SEAL** :印章管理模块</li>
	// <li> **OPERATOR** :组织管理模块</li></ul>
	// 注意：
	// 1、如果EndPoint选择"CHANNEL"或"APP"，该参数仅支持传递"SEAL"，进入印章管理模块
	// 2、该参数**仅在企业和员工激活已经完成，登录控制台场景才生效**。
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 该参数和Module参数配合使用，用于指定模块下的资源Id，指定后链接登录将展示该资源的详情。
	// 
	// 根据Module参数的不同所代表的含义不同(ModuleId需要和Module对应，ModuleId可以通过API或者控制台获取到)。当前支持：
	// <table> <thead> <tr> <th>Module传值</th> <th>ModuleId传值</th> <th>进入的目标页面</th> </tr> </thead> <tbody> <tr> <td>SEAL</td> <td>印章ID</td> <td>查看指定印章的详情页面</td> </tr> <tr> <td>TEMPLATE</td> <td>合同模板ID</td> <td>指定模板的详情页面</td> </tr> <tr> <td>DOCUMENT</td> <td>合同ID</td> <td>指定合同的详情页面</td> </tr> </tbody> </table>
	// 注意：该参数**仅在企业和员工激活完成，登录控制台场景才生效**。
	ModuleId *string `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// 是否展示左侧菜单栏 
	// <ul><li> **ENABLE** : (默认)进入web控制台展示左侧菜单栏</li>
	// <li> **DISABLE** : 进入web控制台不展示左侧菜单栏</li></ul>
	// 注：该参数**仅在企业和员工激活完成，登录控制台场景才生效**。
	MenuStatus *string `json:"MenuStatus,omitnil,omitempty" name:"MenuStatus"`

	// 生成链接的类型：
	// <ul><li>**PC**：(默认)<font color="red">web控制台</font>链接, 需要在PC浏览器中打开</li>
	// <li>**CHANNEL**：H5跳转到电子签小程序链接, 一般用于发送短信中带的链接, 打开后进入腾讯电子签小程序</li>
	// <li>**SHORT_URL**：<font color="red">H5</font>跳转到电子签小程序链接的短链形式, 一般用于发送短信中带的链接, 打开后进入腾讯电子签小程序</li>
	// <li>**WEIXIN_QRCODE_URL**：直接跳转至电子签小程序的二维码链接，无需通过中转页。<font color="red">您需要自行将其转换为二维码，使用微信扫码后可直接进入。请注意，直接点击链接是无效的。</font></li>
	// <li>**APP**：<font color="red">APP或小程序</font>跳转电子签小程序链接, 一般用于贵方小程序或者APP跳转过来,  打开后进入腾讯电子签小程序</li>
	// <li>**H5**：<font color="red">H5长链接</font>跳转H5链接, 一般用于贵方H5跳转过来,  打开后进入腾讯电子签H5页面</li>
	// <li>**SHORT_H5**：<font color="red">H5短链</font>跳转H5的短链形式, 一般用于发送短信中带的链接, 打开后进入腾讯电子签H5页面</li></ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 触发自动跳转事件，仅对EndPoint为App类型有效，可选值包括：
	// <ul><li> **VERIFIED** :企业认证完成/员工认证完成后跳回原App/小程序</li></ul>
	AutoJumpBackEvent *string `json:"AutoJumpBackEvent,omitnil,omitempty" name:"AutoJumpBackEvent"`

	// 可选的此企业允许的授权方式, 可以设置的方式有:
	// <ul>
	// <li>2：转法定代表人授权</li>
	// <li>5：授权书+对公打款</li>
	// </ul>
	AuthorizationTypes []*int64 `json:"AuthorizationTypes,omitnil,omitempty" name:"AuthorizationTypes"`

	// 子客经办人身份证
	// 注意：`如果已同步，这里非空会更新同步的经办人身份证号，暂时只支持中国大陆居民身份证类型`。
	ProxyOperatorIdCardNumber *string `json:"ProxyOperatorIdCardNumber,omitnil,omitempty" name:"ProxyOperatorIdCardNumber"`

	// 认证完成跳转链接。
	// 注意：`此功能仅在Endpoint参数设置成 H5 或 PC时才有效`。
	AutoJumpUrl *string `json:"AutoJumpUrl,omitnil,omitempty" name:"AutoJumpUrl"`

	// 是否展示头顶导航栏  <ul><li> **ENABLE** : (默认)进入web控制台展示头顶导航栏</li> <li> **DISABLE** : 进入web控制台不展示头顶导航栏</li></ul> 注：该参数**仅在企业和员工激活完成，登录控制台场景才生效**。
	// 
	// <a href="https://qcloudimg.tencent-cloud.cn/raw/dd54f333140c711cf6a88e3801bcd178.png" target="_blank">点击查看头顶导航栏位置</a>
	TopNavigationStatus *string `json:"TopNavigationStatus,omitnil,omitempty" name:"TopNavigationStatus"`

	// 是否自动激活子客企业，有下面两种选项：
	// 
	// **false（默认设置）**：不自动激活子客户。您需要通过控制台或调用[激活或者续期子企业](https://qian.tencent.com/developers/partnerApis/accounts/CreateChannelSubOrganizationActive)接口手动完成激活过程。
	// 
	// **true**：若持有的许可证充足，子客户企业注册完成后将自动激活，无需手动操作或访问控制台。
	// 
	// <b>注</b>：如果<b>应用扩展服务</b>中的<b>自动激活子客企业</b>为打开态， 则忽略本接口的AutoActive这个参数（若持有的许可证充足，子客户企业注册完成后将自动激活），具体位置参考下图：
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/c3639b05503d3735bac483d17aa6b0a3.png)
	AutoActive *bool `json:"AutoActive,omitnil,omitempty" name:"AutoActive"`

	// 营业执照正面照（支持PNG或JPG格式）需以base64格式提供，且文件大小不得超过5MB。
	BusinessLicense *string `json:"BusinessLicense,omitnil,omitempty" name:"BusinessLicense"`

	// 组织机构企业注册地址。 请确认该企业注册地址与企业营业执照中注册的地址一致。	
	ProxyAddress *string `json:"ProxyAddress,omitnil,omitempty" name:"ProxyAddress"`

	// 组织机构法人的姓名。 请确认该企业统一社会信用代码与企业营业执照中注册的法人姓名一致。	
	ProxyLegalName *string `json:"ProxyLegalName,omitnil,omitempty" name:"ProxyLegalName"`

	// 授权书(PNG或JPG或PDF) base64格式, 大小不超过8M 。
	//  p.s. 如果上传授权书 ，需遵循以下条件 
	// 1. 超管的信息（超管姓名，超管手机号）必须为必填参数。 
	// 2. 认证方式AuthorizationTypes必须只能是上传授权书方式	
	PowerOfAttorneys []*string `json:"PowerOfAttorneys,omitnil,omitempty" name:"PowerOfAttorneys"`

	// 企业认证时个性化能力信息
	OrganizationAuthorizationOptions *OrganizationAuthorizationOptions `json:"OrganizationAuthorizationOptions,omitnil,omitempty" name:"OrganizationAuthorizationOptions"`

	// 组织机构对公打款 账号，账户名跟企业名称一致。
	// 
	// p.s.
	// 只有认证方式是授权书+对公打款时才生效。
	BankAccountNumber *string `json:"BankAccountNumber,omitnil,omitempty" name:"BankAccountNumber"`

	// 无
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type CreateConsoleLoginUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  <a href="https://qcloudimg.tencent-cloud.cn/raw/a71872de3d540d55451e3e73a2ad1a6e.png" target="_blank">Agent.AppId</a></li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>注:
	// `1. 企业激活时， 此时的Agent.ProxyOrganizationOpenId将会是企业激活后企业的唯一标识，建议开发者保存企业ProxyOrganizationOpenId，后续各项接口调用皆需要此参数。 `
	// `2. 员工认证时， 此时的Agent.ProxyOperator.OpenId将会是员工认证加入企业后的唯一标识，建议开发者保存此员工的OpenId，后续各项接口调用皆需要此参数。 `
	// `3. 同渠道应用(Agent.AppId)下，企业唯一标识ProxyOrganizationOpenId需要保持唯一，员工唯一标识OpenId也要保持唯一 (而不是企业下唯一)。 `
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 第三方平台子客的企业名称，请确认该企业名称与企业营业执照中注册的名称完全一致。
	// <font color="red">
	// 在测试环境联调的过程中，企业名称请统一加上“测试”二字，如：典子谦示例企业测试，否则将无法审核通过。
	// 企业名称请使用以下名称, 以下名称可以不用走收录。
	// **子客测试专用企业1 - 子客测试专用企业9**</font>
	// 
	// 注:
	//  `1. 如果名称中包含英文括号()，请使用中文括号（）代替。`
	//  `2、该名称需要与Agent.ProxyOrganizationOpenId相匹配,  企业激活后Agent.ProxyOrganizationOpenId会跟此企业名称一一绑定; 如果您的企业已经在认证授权中或者激活完成，这里修改子客企业名字将不会生效。 `
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil,omitempty" name:"ProxyOrganizationName"`

	// 子客企业统一社会信用代码，最大长度200个字符
	// 注意：`如果您的企业已经在认证授权中或者激活完成，这里修改子客企业名字将不会生效`。
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// 子客企业员工的姓名，最大长度50个字符,  员工的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：`该姓名需要和Agent.ProxyOperator.OpenId相匹配,  当员工完成认证后该姓名会和Agent.ProxyOperator.OpenId一一绑定, 若员工已认证加入企业，这里修改经办人名字传入将不会生效`
	ProxyOperatorName *string `json:"ProxyOperatorName,omitnil,omitempty" name:"ProxyOperatorName"`

	// 子客企业员工的手机码,  支持国内手机号11位数字(无需加+86前缀或其他字符)。注：`该手机号需要和Agent.ProxyOperator.OpenId相匹配,  当员工完成认证后该手机号会和Agent.ProxyOperator.OpenId一一绑定, 若员工已认证加入企业，这里修改经办人手机号传入将不会生效`
	ProxyOperatorMobile *string `json:"ProxyOperatorMobile,omitnil,omitempty" name:"ProxyOperatorMobile"`

	// Web控制台登录后进入的功能模块,  支持的模块包括：
	// <ul>
	// <li> **空值** :(默认)企业中心模块</li>
	// <li> **DOCUMENT** :合同管理模块</li>
	// <li> **TEMPLATE** :企业模板管理模块</li>
	// <li> **SEAL** :印章管理模块</li>
	// <li> **OPERATOR** :组织管理模块</li></ul>
	// 注意：
	// 1、如果EndPoint选择"CHANNEL"或"APP"，该参数仅支持传递"SEAL"，进入印章管理模块
	// 2、该参数**仅在企业和员工激活已经完成，登录控制台场景才生效**。
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 该参数和Module参数配合使用，用于指定模块下的资源Id，指定后链接登录将展示该资源的详情。
	// 
	// 根据Module参数的不同所代表的含义不同(ModuleId需要和Module对应，ModuleId可以通过API或者控制台获取到)。当前支持：
	// <table> <thead> <tr> <th>Module传值</th> <th>ModuleId传值</th> <th>进入的目标页面</th> </tr> </thead> <tbody> <tr> <td>SEAL</td> <td>印章ID</td> <td>查看指定印章的详情页面</td> </tr> <tr> <td>TEMPLATE</td> <td>合同模板ID</td> <td>指定模板的详情页面</td> </tr> <tr> <td>DOCUMENT</td> <td>合同ID</td> <td>指定合同的详情页面</td> </tr> </tbody> </table>
	// 注意：该参数**仅在企业和员工激活完成，登录控制台场景才生效**。
	ModuleId *string `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// 是否展示左侧菜单栏 
	// <ul><li> **ENABLE** : (默认)进入web控制台展示左侧菜单栏</li>
	// <li> **DISABLE** : 进入web控制台不展示左侧菜单栏</li></ul>
	// 注：该参数**仅在企业和员工激活完成，登录控制台场景才生效**。
	MenuStatus *string `json:"MenuStatus,omitnil,omitempty" name:"MenuStatus"`

	// 生成链接的类型：
	// <ul><li>**PC**：(默认)<font color="red">web控制台</font>链接, 需要在PC浏览器中打开</li>
	// <li>**CHANNEL**：H5跳转到电子签小程序链接, 一般用于发送短信中带的链接, 打开后进入腾讯电子签小程序</li>
	// <li>**SHORT_URL**：<font color="red">H5</font>跳转到电子签小程序链接的短链形式, 一般用于发送短信中带的链接, 打开后进入腾讯电子签小程序</li>
	// <li>**WEIXIN_QRCODE_URL**：直接跳转至电子签小程序的二维码链接，无需通过中转页。<font color="red">您需要自行将其转换为二维码，使用微信扫码后可直接进入。请注意，直接点击链接是无效的。</font></li>
	// <li>**APP**：<font color="red">APP或小程序</font>跳转电子签小程序链接, 一般用于贵方小程序或者APP跳转过来,  打开后进入腾讯电子签小程序</li>
	// <li>**H5**：<font color="red">H5长链接</font>跳转H5链接, 一般用于贵方H5跳转过来,  打开后进入腾讯电子签H5页面</li>
	// <li>**SHORT_H5**：<font color="red">H5短链</font>跳转H5的短链形式, 一般用于发送短信中带的链接, 打开后进入腾讯电子签H5页面</li></ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 触发自动跳转事件，仅对EndPoint为App类型有效，可选值包括：
	// <ul><li> **VERIFIED** :企业认证完成/员工认证完成后跳回原App/小程序</li></ul>
	AutoJumpBackEvent *string `json:"AutoJumpBackEvent,omitnil,omitempty" name:"AutoJumpBackEvent"`

	// 可选的此企业允许的授权方式, 可以设置的方式有:
	// <ul>
	// <li>2：转法定代表人授权</li>
	// <li>5：授权书+对公打款</li>
	// </ul>
	AuthorizationTypes []*int64 `json:"AuthorizationTypes,omitnil,omitempty" name:"AuthorizationTypes"`

	// 子客经办人身份证
	// 注意：`如果已同步，这里非空会更新同步的经办人身份证号，暂时只支持中国大陆居民身份证类型`。
	ProxyOperatorIdCardNumber *string `json:"ProxyOperatorIdCardNumber,omitnil,omitempty" name:"ProxyOperatorIdCardNumber"`

	// 认证完成跳转链接。
	// 注意：`此功能仅在Endpoint参数设置成 H5 或 PC时才有效`。
	AutoJumpUrl *string `json:"AutoJumpUrl,omitnil,omitempty" name:"AutoJumpUrl"`

	// 是否展示头顶导航栏  <ul><li> **ENABLE** : (默认)进入web控制台展示头顶导航栏</li> <li> **DISABLE** : 进入web控制台不展示头顶导航栏</li></ul> 注：该参数**仅在企业和员工激活完成，登录控制台场景才生效**。
	// 
	// <a href="https://qcloudimg.tencent-cloud.cn/raw/dd54f333140c711cf6a88e3801bcd178.png" target="_blank">点击查看头顶导航栏位置</a>
	TopNavigationStatus *string `json:"TopNavigationStatus,omitnil,omitempty" name:"TopNavigationStatus"`

	// 是否自动激活子客企业，有下面两种选项：
	// 
	// **false（默认设置）**：不自动激活子客户。您需要通过控制台或调用[激活或者续期子企业](https://qian.tencent.com/developers/partnerApis/accounts/CreateChannelSubOrganizationActive)接口手动完成激活过程。
	// 
	// **true**：若持有的许可证充足，子客户企业注册完成后将自动激活，无需手动操作或访问控制台。
	// 
	// <b>注</b>：如果<b>应用扩展服务</b>中的<b>自动激活子客企业</b>为打开态， 则忽略本接口的AutoActive这个参数（若持有的许可证充足，子客户企业注册完成后将自动激活），具体位置参考下图：
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/c3639b05503d3735bac483d17aa6b0a3.png)
	AutoActive *bool `json:"AutoActive,omitnil,omitempty" name:"AutoActive"`

	// 营业执照正面照（支持PNG或JPG格式）需以base64格式提供，且文件大小不得超过5MB。
	BusinessLicense *string `json:"BusinessLicense,omitnil,omitempty" name:"BusinessLicense"`

	// 组织机构企业注册地址。 请确认该企业注册地址与企业营业执照中注册的地址一致。	
	ProxyAddress *string `json:"ProxyAddress,omitnil,omitempty" name:"ProxyAddress"`

	// 组织机构法人的姓名。 请确认该企业统一社会信用代码与企业营业执照中注册的法人姓名一致。	
	ProxyLegalName *string `json:"ProxyLegalName,omitnil,omitempty" name:"ProxyLegalName"`

	// 授权书(PNG或JPG或PDF) base64格式, 大小不超过8M 。
	//  p.s. 如果上传授权书 ，需遵循以下条件 
	// 1. 超管的信息（超管姓名，超管手机号）必须为必填参数。 
	// 2. 认证方式AuthorizationTypes必须只能是上传授权书方式	
	PowerOfAttorneys []*string `json:"PowerOfAttorneys,omitnil,omitempty" name:"PowerOfAttorneys"`

	// 企业认证时个性化能力信息
	OrganizationAuthorizationOptions *OrganizationAuthorizationOptions `json:"OrganizationAuthorizationOptions,omitnil,omitempty" name:"OrganizationAuthorizationOptions"`

	// 组织机构对公打款 账号，账户名跟企业名称一致。
	// 
	// p.s.
	// 只有认证方式是授权书+对公打款时才生效。
	BankAccountNumber *string `json:"BankAccountNumber,omitnil,omitempty" name:"BankAccountNumber"`

	// 无
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	delete(f, "UniformSocialCreditCode")
	delete(f, "ProxyOperatorName")
	delete(f, "ProxyOperatorMobile")
	delete(f, "Module")
	delete(f, "ModuleId")
	delete(f, "MenuStatus")
	delete(f, "Endpoint")
	delete(f, "AutoJumpBackEvent")
	delete(f, "AuthorizationTypes")
	delete(f, "ProxyOperatorIdCardNumber")
	delete(f, "AutoJumpUrl")
	delete(f, "TopNavigationStatus")
	delete(f, "AutoActive")
	delete(f, "BusinessLicense")
	delete(f, "ProxyAddress")
	delete(f, "ProxyLegalName")
	delete(f, "PowerOfAttorneys")
	delete(f, "OrganizationAuthorizationOptions")
	delete(f, "BankAccountNumber")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsoleLoginUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsoleLoginUrlResponseParams struct {
	// 跳转链接, 链接的有效期根据企业,员工状态和终端等有区别, 可以参考下表
	// <table> <thead> <tr> <th>子客企业状态</th> <th>子客企业员工状态</th> <th>Endpoint</th> <th>链接有效期限</th> </tr> </thead>  <tbody> <tr> <td>企业未激活</td> <td>员工未认证</td> <td>PC/PC_SHORT_URL</td> <td>5分钟</td>  </tr>  <tr> <td>企业未激活</td> <td>员工未认证</td> <td>CHANNEL/APP/H5/SHORT_H5/WEIXIN_QRCODE_URL</td> <td>30天</td>  </tr>  <tr> <td>企业已激活</td> <td>员工未认证</td> <td>PC/PC_SHORT_URL</td> <td>5分钟</td>  </tr> <tr> <td>企业已激活</td> <td>员工未认证</td> <td>CHANNEL/APP/H5/SHORT_H5/WEIXIN_QRCODE_URL</td> <td>30天</td>  </tr>  <tr> <td>企业已激活</td> <td>员工已认证</td> <td>PC</td> <td>5分钟</td>  </tr>  <tr> <td>企业已激活</td> <td>员工已认证</td> <td>CHANNEL/APP/H5/SHORT_H5/WEIXIN_QRCODE_URL</td> <td>30天</td>  </tr> </tbody> </table>
	// 
	// 注： 
	// 1. <font color="red">链接仅单次有效</font>，每次登录需要需要重新创建新的链接
	// 2. 创建的链接应避免被转义，如：&被转义为\u0026；如使用Postman请求后，请选择响应类型为 JSON，否则链接将被转义
	// 3. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	ConsoleUrl *string `json:"ConsoleUrl,omitnil,omitempty" name:"ConsoleUrl"`

	// 子客企业是否已开通腾讯电子签，
	// <ul><li> **true** :已经开通腾讯电子签</li>
	// <li> **false** :还未开通腾讯电子签</li></ul>
	// 
	// 注：`企业是否实名根据传参Agent.ProxyOrganizationOpenId进行判断，非企业名称或者社会信用代码`
	IsActivated *bool `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 当前经办人是否已认证并加入功能
	// <ul><li> **true** : 已经认证加入公司</li>
	// <li> **false** : 还未认证加入公司</li></ul>
	// 注意：**员工是否实名是根据Agent.ProxyOperator.OpenId判断，非经办人姓名**
	ProxyOperatorIsVerified *bool `json:"ProxyOperatorIsVerified,omitnil,omitempty" name:"ProxyOperatorIsVerified"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateEmployeeChangeUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 待修改的员工OpenId
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 待修改的员工手机号，支持海外格式
	NewMobile *string `json:"NewMobile,omitnil,omitempty" name:"NewMobile"`
}

type CreateEmployeeChangeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 待修改的员工OpenId
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 待修改的员工手机号，支持海外格式
	NewMobile *string `json:"NewMobile,omitnil,omitempty" name:"NewMobile"`
}

func (r *CreateEmployeeChangeUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmployeeChangeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "OpenId")
	delete(f, "NewMobile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmployeeChangeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmployeeChangeUrlResponseParams struct {
	// 修改员工信息的小程序链接<br>跳转到腾讯电子签小程序的实现可以参考微信的官方文档:<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/launchApp.html" target="_blank">开放能力/打开 App</a> 
	MiniAppPath *string `json:"MiniAppPath,omitnil,omitempty" name:"MiniAppPath"`

	// 链接过期时间以 Unix 时间戳格式表示，从生成链接时间起，往后7天有效期。过期后短链将失效，无法打开。
	// 
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEmployeeChangeUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmployeeChangeUrlResponseParams `json:"Response"`
}

func (r *CreateEmployeeChangeUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmployeeChangeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmployeeQualificationSealQrCodeRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。此接口下面信息必填。<ul><li>渠道应用标识:  Agent.AppId</li><li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li><li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li></ul>第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 提示信息，扫码后此信息会展示给扫描用户，用来提示用户授权操作的目的，会在授权界面下面的位置展示。
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/8436ffd78c20605e6b133ff4bc4d2ac7.png)
	HintText *string `json:"HintText,omitnil,omitempty" name:"HintText"`
}

type CreateEmployeeQualificationSealQrCodeRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。此接口下面信息必填。<ul><li>渠道应用标识:  Agent.AppId</li><li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li><li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li></ul>第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 提示信息，扫码后此信息会展示给扫描用户，用来提示用户授权操作的目的，会在授权界面下面的位置展示。
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/8436ffd78c20605e6b133ff4bc4d2ac7.png)
	HintText *string `json:"HintText,omitnil,omitempty" name:"HintText"`
}

func (r *CreateEmployeeQualificationSealQrCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmployeeQualificationSealQrCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "HintText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmployeeQualificationSealQrCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmployeeQualificationSealQrCodeResponseParams struct {
	// 二维码图片的Base64  注:  `此二维码的有效时间为7天，过期后需要重新生成新的二维码图片`
	QrcodeBase64 *string `json:"QrcodeBase64,omitnil,omitempty" name:"QrcodeBase64"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEmployeeQualificationSealQrCodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmployeeQualificationSealQrCodeResponseParams `json:"Response"`
}

func (r *CreateEmployeeQualificationSealQrCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmployeeQualificationSealQrCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowBlockchainEvidenceUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。  此接口下面信息必填。 <ul> <li>渠道应用标识:  Agent.AppId</li> <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li> <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li> </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID，为32位字符串。建议开发者妥善保存此流程ID，以便于顺利进行后续操作。可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 链接/二维码的有效截止时间，格式为unix时间戳。最长不超过 2099年12月31日（4102415999）。
	// 默认值为有效期为当前时间后7天。
	ExpiredOn *uint64 `json:"ExpiredOn,omitnil,omitempty" name:"ExpiredOn"`
}

type CreateFlowBlockchainEvidenceUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。  此接口下面信息必填。 <ul> <li>渠道应用标识:  Agent.AppId</li> <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li> <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li> </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID，为32位字符串。建议开发者妥善保存此流程ID，以便于顺利进行后续操作。可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 链接/二维码的有效截止时间，格式为unix时间戳。最长不超过 2099年12月31日（4102415999）。
	// 默认值为有效期为当前时间后7天。
	ExpiredOn *uint64 `json:"ExpiredOn,omitnil,omitempty" name:"ExpiredOn"`
}

func (r *CreateFlowBlockchainEvidenceUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowBlockchainEvidenceUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowId")
	delete(f, "ExpiredOn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowBlockchainEvidenceUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowBlockchainEvidenceUrlResponseParams struct {
	// 二维码图片下载链接，下载链接有效时间5分钟，请尽快下载保存。
	QrCode *string `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 查看短链，可直接点击短链查看证书。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 二维码和短链的过期时间戳，过期时间默认为生成链接后7天。
	ExpiredOn *uint64 `json:"ExpiredOn,omitnil,omitempty" name:"ExpiredOn"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFlowBlockchainEvidenceUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowBlockchainEvidenceUrlResponseParams `json:"Response"`
}

func (r *CreateFlowBlockchainEvidenceUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowBlockchainEvidenceUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowForwardsRequestParams struct {
	// 合同对应参与方需要修改的目标经办人对应的OpenId。
	// 
	// 注意：`需要保证目标经办人已经加入企业且已实名`
	TargetOpenId *string `json:"TargetOpenId,omitnil,omitempty" name:"TargetOpenId"`

	// 企业签署方的合同及对应签署方
	FlowForwardInfos []*FlowForwardInfo `json:"FlowForwardInfos,omitnil,omitempty" name:"FlowForwardInfos"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。此接口下面信息必填。<ul><li>渠道应用标识:  Agent.AppId</li><li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li><li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li></ul>第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CreateFlowForwardsRequest struct {
	*tchttp.BaseRequest
	
	// 合同对应参与方需要修改的目标经办人对应的OpenId。
	// 
	// 注意：`需要保证目标经办人已经加入企业且已实名`
	TargetOpenId *string `json:"TargetOpenId,omitnil,omitempty" name:"TargetOpenId"`

	// 企业签署方的合同及对应签署方
	FlowForwardInfos []*FlowForwardInfo `json:"FlowForwardInfos,omitnil,omitempty" name:"FlowForwardInfos"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。此接口下面信息必填。<ul><li>渠道应用标识:  Agent.AppId</li><li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li><li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li></ul>第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *CreateFlowForwardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowForwardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetOpenId")
	delete(f, "FlowForwardInfos")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowForwardsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowForwardsResponseParams struct {
	// 失败的合同id以及错误详情
	FailedFlows []*FlowForwardResult `json:"FailedFlows,omitnil,omitempty" name:"FailedFlows"`

	// 成功的合同id
	SuccessFlows []*string `json:"SuccessFlows,omitnil,omitempty" name:"SuccessFlows"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFlowForwardsResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowForwardsResponseParams `json:"Response"`
}

func (r *CreateFlowForwardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowForwardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowGroupSignReviewRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	//   合同(流程)组的合同组Id，为32位字符串，通过接口[通过多文件创建合同组签署流程](https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByFiles) 或[通过多模板创建合同组签署流程](https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByTemplates)创建合同组签署流程时返回。
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 提交的审核结果，审核结果有下面三种情况
	// <ul><li><b>PASS</b>: 审核通过，合同流程可以继续执行签署等操作</li>
	// <li><b>REJECT</b>: 审核拒绝，合同流程不会变动</li>
	// <li><b>SIGN_REJECT</b>:拒签，合同流程直接结束，合同状态变为**合同拒签**</li></ul>
	ReviewType *string `json:"ReviewType,omitnil,omitempty" name:"ReviewType"`

	// 需要进行签署审核的签署人的个人信息或企业信息，签署方的匹配方式按照以下规则:
	// 
	// 个人：二选一（选择其中任意信息组合即可）
	// <ul><li>姓名+证件类型+证件号</li>
	// <li>姓名+手机号</li></ul>
	// 
	// 企业：二选一  （选择其中任意信息组合即可）
	// <ul><li>企业名+姓名+证件类型+证件号</li>
	// <li>企业名+姓名+手机号</li></ul>
	ApproverInfo *NeedReviewApproverInfo `json:"ApproverInfo,omitnil,omitempty" name:"ApproverInfo"`

	// 审核不通过的原因，该字段的字符串长度不超过200个字符。
	// 
	// 注：`当审核类型（ReviewType）为审核拒绝（REJECT）或拒签（SIGN_REJECT）时，审核结果原因字段必须填写`
	ReviewMessage *string `json:"ReviewMessage,omitnil,omitempty" name:"ReviewMessage"`
}

type CreateFlowGroupSignReviewRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	//   合同(流程)组的合同组Id，为32位字符串，通过接口[通过多文件创建合同组签署流程](https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByFiles) 或[通过多模板创建合同组签署流程](https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByTemplates)创建合同组签署流程时返回。
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 提交的审核结果，审核结果有下面三种情况
	// <ul><li><b>PASS</b>: 审核通过，合同流程可以继续执行签署等操作</li>
	// <li><b>REJECT</b>: 审核拒绝，合同流程不会变动</li>
	// <li><b>SIGN_REJECT</b>:拒签，合同流程直接结束，合同状态变为**合同拒签**</li></ul>
	ReviewType *string `json:"ReviewType,omitnil,omitempty" name:"ReviewType"`

	// 需要进行签署审核的签署人的个人信息或企业信息，签署方的匹配方式按照以下规则:
	// 
	// 个人：二选一（选择其中任意信息组合即可）
	// <ul><li>姓名+证件类型+证件号</li>
	// <li>姓名+手机号</li></ul>
	// 
	// 企业：二选一  （选择其中任意信息组合即可）
	// <ul><li>企业名+姓名+证件类型+证件号</li>
	// <li>企业名+姓名+手机号</li></ul>
	ApproverInfo *NeedReviewApproverInfo `json:"ApproverInfo,omitnil,omitempty" name:"ApproverInfo"`

	// 审核不通过的原因，该字段的字符串长度不超过200个字符。
	// 
	// 注：`当审核类型（ReviewType）为审核拒绝（REJECT）或拒签（SIGN_REJECT）时，审核结果原因字段必须填写`
	ReviewMessage *string `json:"ReviewMessage,omitnil,omitempty" name:"ReviewMessage"`
}

func (r *CreateFlowGroupSignReviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowGroupSignReviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowGroupId")
	delete(f, "ReviewType")
	delete(f, "ApproverInfo")
	delete(f, "ReviewMessage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowGroupSignReviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowGroupSignReviewResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFlowGroupSignReviewResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowGroupSignReviewResponseParams `json:"Response"`
}

func (r *CreateFlowGroupSignReviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowGroupSignReviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlowOption struct {
	// 是否允许修改合同信息，
	// **true**：可以
	// **false**：（默认）不可以
	CanEditFlow *bool `json:"CanEditFlow,omitnil,omitempty" name:"CanEditFlow"`

	// 是否允许发起合同弹窗隐藏合同名称
	// **true**：允许
	// **false**：（默认）不允许
	HideShowFlowName *bool `json:"HideShowFlowName,omitnil,omitempty" name:"HideShowFlowName"`

	// 是否允许发起合同弹窗隐藏合同类型，
	// **true**：允许
	// **false**：（默认）不允许
	HideShowFlowType *bool `json:"HideShowFlowType,omitnil,omitempty" name:"HideShowFlowType"`

	// 是否允许发起合同弹窗隐藏合同到期时间
	// **true**：允许
	// **false**：（默认）不允许
	HideShowDeadline *bool `json:"HideShowDeadline,omitnil,omitempty" name:"HideShowDeadline"`

	// 是否允许发起合同步骤跳过指定签署方步骤
	// **true**：允许
	// **false**：（默认）不允许
	CanSkipAddApprover *bool `json:"CanSkipAddApprover,omitnil,omitempty" name:"CanSkipAddApprover"`

	// 是否可以编辑签署人包括新增，修改，删除 
	// <ul><li>（默认） false -可以编辑签署人</li> <li> true - 禁止编辑签署人</li></ul>
	// 
	// 
	// 
	// 注意：
	// * 如果设置参数为 true， 则 参数签署人 [FlowApproverList](https://qian.tencent.com/developers/partnerApis/embedPages/ChannelCreatePrepareFlow) 不能为空
	// * 此参数对子客和自动签无效，不允许进行修改。
	ForbidEditApprover *bool `json:"ForbidEditApprover,omitnil,omitempty" name:"ForbidEditApprover"`

	// 定制化发起合同弹窗的描述信息，长度不能超过500，只能由中文、字母、数字和标点组成。
	CustomCreateFlowDescription *string `json:"CustomCreateFlowDescription,omitnil,omitempty" name:"CustomCreateFlowDescription"`

	// 禁止编辑填写控件
	// 
	// **true**：禁止编辑填写控件
	// **false**：（默认）允许编辑填写控件
	ForbidEditFillComponent *bool `json:"ForbidEditFillComponent,omitnil,omitempty" name:"ForbidEditFillComponent"`

	// 跳过上传文件步骤
	// 
	// **true**：跳过
	// **false**：（默认）不跳过，需要传ResourceId
	SkipUploadFile *bool `json:"SkipUploadFile,omitnil,omitempty" name:"SkipUploadFile"`

	// 签署控件的配置信息，用在嵌入式发起的页面配置，包括 
	//  - 签署控件 是否默认展示日期.
	SignComponentConfig *SignComponentConfig `json:"SignComponentConfig,omitnil,omitempty" name:"SignComponentConfig"`

	// 是否禁止编辑（展示）水印控件属性
	// <ul><li>（默认） false -否</li> <li> true - 禁止编辑</li></ul>
	ForbidEditWatermark *bool `json:"ForbidEditWatermark,omitnil,omitempty" name:"ForbidEditWatermark"`
}

// Predefined struct for user
type CreateFlowsByTemplatesRequestParams struct {
	// 合同的发起企业和发起人信息，<a href="https://qcloudimg.tencent-cloud.cn/raw/b69f8aad306c40b7b78d096e39b2edbb.png" target="_blank">点击查看合同发起企业和人展示的位置</a>
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  <a href="https://qcloudimg.tencent-cloud.cn/raw/a71872de3d540d55451e3e73a2ad1a6e.png" target="_blank">Agent.AppId</a></li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId（合同的发起企业）</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId （合同的发起人）</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要创建的合同信息列表，最多支持一次创建20个合同
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil,omitempty" name:"FlowInfos"`

	// 是否为预览模式，取值如下：
	// <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li>
	// <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。 <font color="red">注意： 以预览模式创建的合同仅供查看，因此参与方无法进行签署操作</font></li></ul>
	// 
	// 注:
	// 
	// `如果预览的文件中指定了动态表格控件，此时此接口返回的是合成前的文档预览链接，合成完成后的文档预览链接需要通过回调通知的方式或使用返回的TaskInfo中的TaskId通过ChannelGetTaskResultApi接口查询得到`
	NeedPreview *bool `json:"NeedPreview,omitnil,omitempty" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子</li></ul>
	// 注: `此参数在NeedPreview 为true时有效`
	PreviewType *int64 `json:"PreviewType,omitnil,omitempty" name:"PreviewType"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type CreateFlowsByTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 合同的发起企业和发起人信息，<a href="https://qcloudimg.tencent-cloud.cn/raw/b69f8aad306c40b7b78d096e39b2edbb.png" target="_blank">点击查看合同发起企业和人展示的位置</a>
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  <a href="https://qcloudimg.tencent-cloud.cn/raw/a71872de3d540d55451e3e73a2ad1a6e.png" target="_blank">Agent.AppId</a></li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId（合同的发起企业）</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId （合同的发起人）</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要创建的合同信息列表，最多支持一次创建20个合同
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil,omitempty" name:"FlowInfos"`

	// 是否为预览模式，取值如下：
	// <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li>
	// <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。 <font color="red">注意： 以预览模式创建的合同仅供查看，因此参与方无法进行签署操作</font></li></ul>
	// 
	// 注:
	// 
	// `如果预览的文件中指定了动态表格控件，此时此接口返回的是合成前的文档预览链接，合成完成后的文档预览链接需要通过回调通知的方式或使用返回的TaskInfo中的TaskId通过ChannelGetTaskResultApi接口查询得到`
	NeedPreview *bool `json:"NeedPreview,omitnil,omitempty" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子</li></ul>
	// 注: `此参数在NeedPreview 为true时有效`
	PreviewType *int64 `json:"PreviewType,omitnil,omitempty" name:"PreviewType"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	delete(f, "PreviewType")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowsByTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowsByTemplatesResponseParams struct {
	// 生成的合同流程ID数组，合同流程ID为32位字符串。
	// 建议开发者妥善保存此流程ID数组，以便于顺利进行后续操作。
	// 
	// [点击产看FlowId在控制台上的位置](https://qcloudimg.tencent-cloud.cn/raw/05af26573d5106763b4cfbb9f7c64b41.png)
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 第三方应用平台的业务信息, 与创建合同的FlowInfos数组中的CustomerData一一对应
	CustomerData []*string `json:"CustomerData,omitnil,omitempty" name:"CustomerData"`

	// 创建消息，对应多个合同ID，
	// 成功为“”,创建失败则对应失败消息
	ErrorMessages []*string `json:"ErrorMessages,omitnil,omitempty" name:"ErrorMessages"`

	// 合同预览链接URL数组。
	// 
	// 注：如果是预览模式(即NeedPreview设置为true)时, 才会有此预览链接URL
	// 如果预览的文件中指定了动态表格控件，此时此接口返回的是合成前的文档预览链接，合成完成后的文档预览链接需要通过[合同文档合成完成回调](https://qian.tencent.com/developers/partner/callback_types_contracts_sign#%E5%8D%81%E4%B8%80-%E5%90%88%E5%90%8C%E6%96%87%E6%A1%A3%E5%90%88%E6%88%90%E5%AE%8C%E6%88%90%E5%9B%9E%E8%B0%83)获取或使用返回的TaskInfo中的TaskId通过[查询转换任务状态
	// ](https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi)接口查询得到
	PreviewUrls []*string `json:"PreviewUrls,omitnil,omitempty" name:"PreviewUrls"`

	// 复杂文档合成任务（如，包含动态表格的预览任务）的任务信息数组；
	// 如果文档需要异步合成，此字段会返回该异步任务的任务信息，后续可以通过ChannelGetTaskResultApi接口查询任务详情；
	TaskInfos []*TaskInfo `json:"TaskInfos,omitnil,omitempty" name:"TaskInfos"`

	// 签署方信息，如角色ID、角色名称等
	FlowApprovers []*FlowApproverItem `json:"FlowApprovers,omitnil,omitempty" name:"FlowApprovers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateLegalSealQrCodeRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>注:
	// `1. 企业激活时， 此时的Agent.ProxyOrganizationOpenId将会是企业激活后企业的唯一标识，建议开发者保存企业ProxyOrganizationOpenId，后续各项接口调用皆需要此参数。 `
	// `2. 员工认证时， 此时的Agent.ProxyOperator.OpenId将会是员工认证加入企业后的唯一标识，建议开发者保存此员工的OpenId，后续各项接口调用皆需要此参数。 `
	// `3. 同渠道应用(Agent.AppId)下，企业唯一标识ProxyOrganizationOpenId需要保持唯一，员工唯一标识OpenId也要保持唯一 (而不是企业下唯一)。 `
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 操作人信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 企业信息
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
}

type CreateLegalSealQrCodeRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>注:
	// `1. 企业激活时， 此时的Agent.ProxyOrganizationOpenId将会是企业激活后企业的唯一标识，建议开发者保存企业ProxyOrganizationOpenId，后续各项接口调用皆需要此参数。 `
	// `2. 员工认证时， 此时的Agent.ProxyOperator.OpenId将会是员工认证加入企业后的唯一标识，建议开发者保存此员工的OpenId，后续各项接口调用皆需要此参数。 `
	// `3. 同渠道应用(Agent.AppId)下，企业唯一标识ProxyOrganizationOpenId需要保持唯一，员工唯一标识OpenId也要保持唯一 (而不是企业下唯一)。 `
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 企业信息
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
}

func (r *CreateLegalSealQrCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLegalSealQrCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "Operator")
	delete(f, "Organization")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLegalSealQrCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLegalSealQrCodeResponseParams struct {
	// 二维码图片base64值，二维码有效期7天（604800秒）
	// 
	// 二维码图片的样式如下图：
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/7ec2478761158a35a9c623882839a5df.png)
	QrcodeBase64 *string `json:"QrcodeBase64,omitnil,omitempty" name:"QrcodeBase64"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLegalSealQrCodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateLegalSealQrCodeResponseParams `json:"Response"`
}

func (r *CreateLegalSealQrCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLegalSealQrCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModifyAdminAuthorizationUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 企业认证流Id，可以通过回调[授权书认证审核结果回调](https://qian.tencent.com/developers/company/callback_types_staffs#%E5%8D%81%E5%85%AD-%E6%8E%88%E6%9D%83%E4%B9%A6%E8%AE%A4%E8%AF%81%E5%AE%A1%E6%A0%B8%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83)得到
	AuthorizationId *string `json:"AuthorizationId,omitnil,omitempty" name:"AuthorizationId"`

	// 要跳转的链接类型<ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li><li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li><li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li><li>**PC**： 跳转电子签web 端控制台的链接。</li></ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

type CreateModifyAdminAuthorizationUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 企业认证流Id，可以通过回调[授权书认证审核结果回调](https://qian.tencent.com/developers/company/callback_types_staffs#%E5%8D%81%E5%85%AD-%E6%8E%88%E6%9D%83%E4%B9%A6%E8%AE%A4%E8%AF%81%E5%AE%A1%E6%A0%B8%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83)得到
	AuthorizationId *string `json:"AuthorizationId,omitnil,omitempty" name:"AuthorizationId"`

	// 要跳转的链接类型<ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li><li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li><li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li><li>**PC**： 跳转电子签web 端控制台的链接。</li></ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

func (r *CreateModifyAdminAuthorizationUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModifyAdminAuthorizationUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "AuthorizationId")
	delete(f, "Endpoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModifyAdminAuthorizationUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModifyAdminAuthorizationUrlResponseParams struct {
	// 变更企业超管授权书链接。没有有效期限制。注意：此链接仅能由当时认证企业的认证人使用。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateModifyAdminAuthorizationUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateModifyAdminAuthorizationUrlResponseParams `json:"Response"`
}

func (r *CreateModifyAdminAuthorizationUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModifyAdminAuthorizationUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationAuthFileRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 企业授权书信息参数， 需要自行保证这些参数跟营业执照中的信息一致。
	OrganizationCommonInfo *OrganizationCommonInfo `json:"OrganizationCommonInfo,omitnil,omitempty" name:"OrganizationCommonInfo"`

	// 授权书类型：
	// 
	// <ul><li>0: 企业认证超管授权书</li><li>1: 超管变更授权书</li><li>2: 企业注销授权书</li></ul>
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type CreateOrganizationAuthFileRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 企业授权书信息参数， 需要自行保证这些参数跟营业执照中的信息一致。
	OrganizationCommonInfo *OrganizationCommonInfo `json:"OrganizationCommonInfo,omitnil,omitempty" name:"OrganizationCommonInfo"`

	// 授权书类型：
	// 
	// <ul><li>0: 企业认证超管授权书</li><li>1: 超管变更授权书</li><li>2: 企业注销授权书</li></ul>
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *CreateOrganizationAuthFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationAuthFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "OrganizationCommonInfo")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationAuthFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationAuthFileResponseParams struct {
	// 授权书链接，有效期5分钟。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationAuthFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationAuthFileResponseParams `json:"Response"`
}

func (r *CreateOrganizationAuthFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationAuthFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePartnerAutoSignAuthUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 被授企业id/授权方企业id（即OrganizationId），如果是企业之间授权和AuthorizedOrganizationName二选一传入。
	// 
	// 注：`被授权企业必须和当前企业在同一应用号下`
	AuthorizedOrganizationId *string `json:"AuthorizedOrganizationId,omitnil,omitempty" name:"AuthorizedOrganizationId"`

	// 被授企业名称/授权方企业的名字，如果是企业之间授权和AuthorizedOrganizationId二选一传入即可。请确认该名称与企业营业执照中注册的名称一致。
	// 
	// 注: 
	// 1. 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 2. 被授权企业必须和当前企业在同一应用号下
	AuthorizedOrganizationName *string `json:"AuthorizedOrganizationName,omitnil,omitempty" name:"AuthorizedOrganizationName"`

	// 是否给平台应用授权
	// 
	// <ul>
	// <li><strong>true</strong>: 表示是，授权平台应用。在此情况下，无需设置<code>AuthorizedOrganizationId</code>和<code>AuthorizedOrganizationName</code>。</li>
	// <li><strong>false</strong>: （默认）表示否，不是授权平台应用。</li>
	// </ul>
	// 
	//  注：授权给平台应用需要开通【基于子客授权第三方应用可文件发起子客自动签署】白名单，请联系运营经理开通。
	PlatformAppAuthorization *bool `json:"PlatformAppAuthorization,omitnil,omitempty" name:"PlatformAppAuthorization"`

	// 在设置印章授权时，可以指定特定的印章类型，以确保在授权过程中只使用相应类型的印章。支持的印章类型包括：
	// 
	// <ul>
	// <li><strong>OFFICIAL</strong>：企业公章，用于代表企业对外的正式文件和重要事务的认证。</li>
	// <li><strong>CONTRACT</strong>：合同专用章，专门用于签署各类合同。</li>
	// <li><strong>FINANCE</strong>：财务专用章，用于企业的财务相关文件，如发票、收据等财务凭证的认证。</li>
	// <li><strong>PERSONNEL</strong>：人事专用章，用于人事管理相关文件，如劳动合同、人事任命等。</li>
	// </ul>
	SealTypes []*string `json:"SealTypes,omitnil,omitempty" name:"SealTypes"`

	// 在处理授权关系时，授权的方向
	// <ul>
	// <li><strong>false</strong>（默认值）：表示我方授权他方。在这种情况下，<code>AuthorizedOrganizationName</code> 代表的是【被授权方】的企业名称，即接收授权的企业。</li>
	// <li><strong>true</strong>：表示他方授权我方。在这种情况下，<code>AuthorizedOrganizationName</code> 代表的是【授权方】的企业名称，即提供授权的企业。</li>
	// </ul>
	AuthToMe *bool `json:"AuthToMe,omitnil,omitempty" name:"AuthToMe"`
}

type CreatePartnerAutoSignAuthUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 被授企业id/授权方企业id（即OrganizationId），如果是企业之间授权和AuthorizedOrganizationName二选一传入。
	// 
	// 注：`被授权企业必须和当前企业在同一应用号下`
	AuthorizedOrganizationId *string `json:"AuthorizedOrganizationId,omitnil,omitempty" name:"AuthorizedOrganizationId"`

	// 被授企业名称/授权方企业的名字，如果是企业之间授权和AuthorizedOrganizationId二选一传入即可。请确认该名称与企业营业执照中注册的名称一致。
	// 
	// 注: 
	// 1. 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 2. 被授权企业必须和当前企业在同一应用号下
	AuthorizedOrganizationName *string `json:"AuthorizedOrganizationName,omitnil,omitempty" name:"AuthorizedOrganizationName"`

	// 是否给平台应用授权
	// 
	// <ul>
	// <li><strong>true</strong>: 表示是，授权平台应用。在此情况下，无需设置<code>AuthorizedOrganizationId</code>和<code>AuthorizedOrganizationName</code>。</li>
	// <li><strong>false</strong>: （默认）表示否，不是授权平台应用。</li>
	// </ul>
	// 
	//  注：授权给平台应用需要开通【基于子客授权第三方应用可文件发起子客自动签署】白名单，请联系运营经理开通。
	PlatformAppAuthorization *bool `json:"PlatformAppAuthorization,omitnil,omitempty" name:"PlatformAppAuthorization"`

	// 在设置印章授权时，可以指定特定的印章类型，以确保在授权过程中只使用相应类型的印章。支持的印章类型包括：
	// 
	// <ul>
	// <li><strong>OFFICIAL</strong>：企业公章，用于代表企业对外的正式文件和重要事务的认证。</li>
	// <li><strong>CONTRACT</strong>：合同专用章，专门用于签署各类合同。</li>
	// <li><strong>FINANCE</strong>：财务专用章，用于企业的财务相关文件，如发票、收据等财务凭证的认证。</li>
	// <li><strong>PERSONNEL</strong>：人事专用章，用于人事管理相关文件，如劳动合同、人事任命等。</li>
	// </ul>
	SealTypes []*string `json:"SealTypes,omitnil,omitempty" name:"SealTypes"`

	// 在处理授权关系时，授权的方向
	// <ul>
	// <li><strong>false</strong>（默认值）：表示我方授权他方。在这种情况下，<code>AuthorizedOrganizationName</code> 代表的是【被授权方】的企业名称，即接收授权的企业。</li>
	// <li><strong>true</strong>：表示他方授权我方。在这种情况下，<code>AuthorizedOrganizationName</code> 代表的是【授权方】的企业名称，即提供授权的企业。</li>
	// </ul>
	AuthToMe *bool `json:"AuthToMe,omitnil,omitempty" name:"AuthToMe"`
}

func (r *CreatePartnerAutoSignAuthUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePartnerAutoSignAuthUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "AuthorizedOrganizationId")
	delete(f, "AuthorizedOrganizationName")
	delete(f, "PlatformAppAuthorization")
	delete(f, "SealTypes")
	delete(f, "AuthToMe")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePartnerAutoSignAuthUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePartnerAutoSignAuthUrlResponseParams struct {
	// 授权链接，以短链形式返回，短链的有效期参考回参中的 ExpiredTime。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 从客户小程序或者客户APP跳转至腾讯电子签小程序进行批量签署的跳转路径
	MiniAppPath *string `json:"MiniAppPath,omitnil,omitempty" name:"MiniAppPath"`

	// 链接过期时间以 Unix 时间戳格式表示，从生成链接时间起，往后7天有效期。过期后短链将失效，无法打开。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePartnerAutoSignAuthUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreatePartnerAutoSignAuthUrlResponseParams `json:"Response"`
}

func (r *CreatePartnerAutoSignAuthUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePartnerAutoSignAuthUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonAuthCertificateImageRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 个人用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 证件类型，支持以下类型<ul><li> ID_CARD  : 中国大陆居民身份证 (默认值)</li><li> HONGKONG_AND_MACAO  : 中国港澳居民来往内地通行证</li><li> HONGKONG_MACAO_AND_TAIWAN  : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则<ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li><li>港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li><li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 自动签使用的场景值, 可以选择的场景值如下:<ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>注: `不传默认为处方单场景，即E_PRESCRIPTION_AUTO_SIGN`
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`
}

type CreatePersonAuthCertificateImageRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 个人用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 证件类型，支持以下类型<ul><li> ID_CARD  : 中国大陆居民身份证 (默认值)</li><li> HONGKONG_AND_MACAO  : 中国港澳居民来往内地通行证</li><li> HONGKONG_MACAO_AND_TAIWAN  : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则<ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li><li>港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li><li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 自动签使用的场景值, 可以选择的场景值如下:<ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>注: `不传默认为处方单场景，即E_PRESCRIPTION_AUTO_SIGN`
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`
}

func (r *CreatePersonAuthCertificateImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonAuthCertificateImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "UserName")
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "SceneKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePersonAuthCertificateImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonAuthCertificateImageResponseParams struct {
	// 个人用户认证证书图片下载URL，`有效期为5分钟`，超过有效期后将无法再下载。
	AuthCertUrl *string `json:"AuthCertUrl,omitnil,omitempty" name:"AuthCertUrl"`

	// 个人用户认证证书的编号, 为20位数字组成的字符串,  由腾讯电子签下发此编号 。该编号会合成到个人用户证书证明图片。注: `个人用户认证证书的编号和证明图片绑定, 获取新的证明图片编号会变动`
	ImageCertId *string `json:"ImageCertId,omitnil,omitempty" name:"ImageCertId"`

	// 在数字证书申请过程中，系统会自动生成一个独一无二的序列号。请注意，当证书到期并自动续期时，该序列号将会发生变化。值得注意的是，此序列号不会被合成至个人用户证书的证明图片中。
	SerialNumber *string `json:"SerialNumber,omitnil,omitempty" name:"SerialNumber"`

	// CA证书颁发时间，格式为Unix标准时间戳（秒）   该时间格式化后会合成到个人用户证书证明图片
	ValidFrom *uint64 `json:"ValidFrom,omitnil,omitempty" name:"ValidFrom"`

	// CA证书有效截止时间，格式为Unix标准时间戳（秒）该时间格式化后会合成到个人用户证书证明图片
	ValidTo *uint64 `json:"ValidTo,omitnil,omitempty" name:"ValidTo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePersonAuthCertificateImageResponse struct {
	*tchttp.BaseResponse
	Response *CreatePersonAuthCertificateImageResponseParams `json:"Response"`
}

func (r *CreatePersonAuthCertificateImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonAuthCertificateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealByImageRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 电子印章名字，1-50个中文字符
	// 注:`同一企业下电子印章名字不能相同`
	SealName *string `json:"SealName,omitnil,omitempty" name:"SealName"`

	// 电子印章图片base64编码，大小不超过10M（原始图片不超过5M），只支持PNG或JPG图片格式
	// 
	// 注: `通过图片创建的电子印章，需电子签平台人工审核`
	// 
	SealImage *string `json:"SealImage,omitnil,omitempty" name:"SealImage"`

	// 操作者的信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 电子印章生成方式
	// <ul>
	// <li><strong>空值</strong>:(默认)使用上传的图片生成印章, 此时需要上传SealImage图片</li>
	// <li><strong>SealGenerateSourceSystem</strong>: 系统生成印章, 无需上传SealImage图片</li>
	// </ul>
	GenerateSource *string `json:"GenerateSource,omitnil,omitempty" name:"GenerateSource"`

	// 电子印章类型 , 可选类型如下: <ul><li>**OFFICIAL**: (默认)公章</li><li>**CONTRACT**: 合同专用章;</li><li>**FINANCE**: 财务专用章;</li><li>**PERSONNEL**: 人事专用章</li><li>**INVOICE**: 发票专用章</li><li>**OTHER**: 其他</li></ul>注: 同企业下只能有<font color="red">一个</font>公章, 重复创建会报错
	SealType *string `json:"SealType,omitnil,omitempty" name:"SealType"`

	// 企业印章横向文字，最多可填15个汉字  (若超过印章最大宽度，优先压缩字间距，其次缩小字号)
	// 横向文字的位置如下图中的"印章横向文字在这里"
	// 
	// ![image](https://dyn.ess.tencent.cn/guide/capi/CreateSealByImage2.png)
	SealHorizontalText *string `json:"SealHorizontalText,omitnil,omitempty" name:"SealHorizontalText"`

	// 印章样式, 可以选择的样式如下: 
	// <ul><li>**circle**:(默认)圆形印章</li>
	// <li>**ellipse**:椭圆印章</li></ul>
	SealStyle *string `json:"SealStyle,omitnil,omitempty" name:"SealStyle"`

	// 印章尺寸取值描述, 可以选择的尺寸如下: <ul><li> **38_38**: 圆形企业公章直径38mm, 当SealStyle是圆形的时候才有效</li> <li> **40_40**: 圆形企业公章直径40mm, 当SealStyle是圆形的时候才有效</li> <li> **42_42**（默认）: 圆形企业公章直径42mm, 当SealStyle是圆形的时候才有效</li> <li> **45_45**: 圆形企业印章直径45mm, 当SealStyle是圆形的时候才有效</li> <li> **50_50**: 圆形企业印章直径45mm, 当SealStyle是圆形的时候才有效</li> <li> **58_58**: 圆形企业印章直径45mm, 当SealStyle是圆形的时候才有效</li>  <li> **40_30**: 椭圆形印章40mm x 30mm, 当SealStyle是椭圆的时候才有效</li> <li> **45_30**: 椭圆形印章45mm x 30mm, 当SealStyle是椭圆的时候才有效</li> </ul>
	SealSize *string `json:"SealSize,omitnil,omitempty" name:"SealSize"`

	// 企业税号
	// 
	// 注:
	// <ul>
	// <li>1.印章类型SealType是INVOICE类型时，此参数才会生效</li>
	// <li>2.印章类型SealType是INVOICE类型，且该字段没有传入值或传入空时，会取该企业对应的统一社会信用代码作为默认的企业税号（<font color="red">如果是通过授权书授权方式认证的企业，此参数必传不能为空</font>）</li>
	// </ul>
	TaxIdentifyCode *string `json:"TaxIdentifyCode,omitnil,omitempty" name:"TaxIdentifyCode"`
}

type CreateSealByImageRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 电子印章名字，1-50个中文字符
	// 注:`同一企业下电子印章名字不能相同`
	SealName *string `json:"SealName,omitnil,omitempty" name:"SealName"`

	// 电子印章图片base64编码，大小不超过10M（原始图片不超过5M），只支持PNG或JPG图片格式
	// 
	// 注: `通过图片创建的电子印章，需电子签平台人工审核`
	// 
	SealImage *string `json:"SealImage,omitnil,omitempty" name:"SealImage"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 电子印章生成方式
	// <ul>
	// <li><strong>空值</strong>:(默认)使用上传的图片生成印章, 此时需要上传SealImage图片</li>
	// <li><strong>SealGenerateSourceSystem</strong>: 系统生成印章, 无需上传SealImage图片</li>
	// </ul>
	GenerateSource *string `json:"GenerateSource,omitnil,omitempty" name:"GenerateSource"`

	// 电子印章类型 , 可选类型如下: <ul><li>**OFFICIAL**: (默认)公章</li><li>**CONTRACT**: 合同专用章;</li><li>**FINANCE**: 财务专用章;</li><li>**PERSONNEL**: 人事专用章</li><li>**INVOICE**: 发票专用章</li><li>**OTHER**: 其他</li></ul>注: 同企业下只能有<font color="red">一个</font>公章, 重复创建会报错
	SealType *string `json:"SealType,omitnil,omitempty" name:"SealType"`

	// 企业印章横向文字，最多可填15个汉字  (若超过印章最大宽度，优先压缩字间距，其次缩小字号)
	// 横向文字的位置如下图中的"印章横向文字在这里"
	// 
	// ![image](https://dyn.ess.tencent.cn/guide/capi/CreateSealByImage2.png)
	SealHorizontalText *string `json:"SealHorizontalText,omitnil,omitempty" name:"SealHorizontalText"`

	// 印章样式, 可以选择的样式如下: 
	// <ul><li>**circle**:(默认)圆形印章</li>
	// <li>**ellipse**:椭圆印章</li></ul>
	SealStyle *string `json:"SealStyle,omitnil,omitempty" name:"SealStyle"`

	// 印章尺寸取值描述, 可以选择的尺寸如下: <ul><li> **38_38**: 圆形企业公章直径38mm, 当SealStyle是圆形的时候才有效</li> <li> **40_40**: 圆形企业公章直径40mm, 当SealStyle是圆形的时候才有效</li> <li> **42_42**（默认）: 圆形企业公章直径42mm, 当SealStyle是圆形的时候才有效</li> <li> **45_45**: 圆形企业印章直径45mm, 当SealStyle是圆形的时候才有效</li> <li> **50_50**: 圆形企业印章直径45mm, 当SealStyle是圆形的时候才有效</li> <li> **58_58**: 圆形企业印章直径45mm, 当SealStyle是圆形的时候才有效</li>  <li> **40_30**: 椭圆形印章40mm x 30mm, 当SealStyle是椭圆的时候才有效</li> <li> **45_30**: 椭圆形印章45mm x 30mm, 当SealStyle是椭圆的时候才有效</li> </ul>
	SealSize *string `json:"SealSize,omitnil,omitempty" name:"SealSize"`

	// 企业税号
	// 
	// 注:
	// <ul>
	// <li>1.印章类型SealType是INVOICE类型时，此参数才会生效</li>
	// <li>2.印章类型SealType是INVOICE类型，且该字段没有传入值或传入空时，会取该企业对应的统一社会信用代码作为默认的企业税号（<font color="red">如果是通过授权书授权方式认证的企业，此参数必传不能为空</font>）</li>
	// </ul>
	TaxIdentifyCode *string `json:"TaxIdentifyCode,omitnil,omitempty" name:"TaxIdentifyCode"`
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
	delete(f, "GenerateSource")
	delete(f, "SealType")
	delete(f, "SealHorizontalText")
	delete(f, "SealStyle")
	delete(f, "SealSize")
	delete(f, "TaxIdentifyCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSealByImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealByImageResponseParams struct {
	// 电子印章ID，为32位字符串。
	// 建议开发者保留此印章ID，后续指定签署区印章或者操作印章需此印章ID。
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 电子印章预览链接地址，地址默认失效时间为24小时。
	// 
	// 注:`图片上传生成的电子印章无预览链接地址`
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID数组，最多支持100个。
	// 
	// 注: 
	// 1. 必须选择提供此参数或合同组编号中的一个。
	// 2. 当生成类型（GenerateType）设为“ALL”时，不可提供多个流程ID。
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 合同组编号
	// 注：`该参数和合同流程ID数组必须二选一`
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 签署链接类型,可以设置的参数如下
	// <ul><li> **WEIXINAPP** :(默认)跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型 ，此时返回短链</li>
	// <li> **CHANNEL** :带有H5引导页的跳转电子签小程序的链接(<b>GenerateType非ALL时候不能设置成CHANNEL</b>)</li>
	// <li> **APP** :第三方App或小程序跳转电子签小程序的path, App或者小程序跳转适合此类型</li>
	// <li> **LONGURL2WEIXINAPP** :跳转电子签小程序的链接, H5跳转适合此类型，此时返回长链</li></ul>
	// 
	// **注：**动态签署人场景，如果签署链接类型设置为`APP`，则仅支持跳转到封面页。
	// 
	// 详细使用场景可以参考接口描述说明中的 **主要使用场景EndPoint分类**
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 签署链接生成类型，可以选择的类型如下
	// 
	// <ul><li><strong>ALL</strong>：（默认）为所有签署方生成签署链接，但不包括自动签署（静默签署）的签署方。注意：<strong>此中类型不支持多个合同ID（FlowIds）</strong>。</li>
	// <li><strong>CHANNEL</strong>：适用于第三方子企业的员工签署方。</li>
	// <li><strong>NOT_CHANNEL</strong>：适用于SaaS平台企业的员工签署方。</li>
	// <li><strong>PERSON</strong>：适用于个人或自然人签署方。</li>
	// <li><strong>FOLLOWER</strong>：适用于关注方，目前指合同的抄送方。</li>
	// <li><strong>RECIPIENT</strong>：根据RecipientId生成对应的签署链接，适用于动态添加签署人的情况。</li></ul>
	GenerateType *string `json:"GenerateType,omitnil,omitempty" name:"GenerateType"`

	// SaaS平台企业员工签署方的企业名称如果名称中包含英文括号()，请使用中文括号（）代替。  注:  `1.GenerateType为"NOT_CHANNEL"时必填` `2.获取B端动态签署人领取链接时,可指定此字段来预先设定签署人的企业,预设后只能以该企业身份去领取合同并完成签署`
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 合同流程里边参与方的姓名。
	// 注: 
	// 1. `GenerateType为"PERSON"(即个人签署方)时必填`。
	// 2. `在动态签署人补充链接场景中，可以通过传入这个值，对补充的个人参与方信息进行限制。仅匹配传入姓名的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方。`
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 合同流程里边签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 注:  
	// 1. `GenerateType为"PERSON"或"FOLLOWER"时必填。`
	// 2. `在动态签署人补充链接场景中，可以通过传入此值，对补充的个人参与方信息进行限制。仅匹配传入手机号的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方。`
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	// 
	// `注：在动态签署人补充链接场景中，可以通过传入此值，对补充的个人参与方信息进行限制。仅匹配传入证件类型的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方，且需要和证件号参数一同传递，不能单独进行限制。`
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成(如存在X，请大写)。</li>
	// <li>港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	// 
	// `注：在动态签署人补充链接场景中，可以通过传入此值，对补充的个人参与方信息进行限制。仅匹配传入证件号的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方。`
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 第三方平台子客企业的企业的标识, 即OrganizationOpenId。 注:  `1.GenerateType为"CHANNEL"时必填` `2.获取B端动态签署人领取链接时,可指定此字段来预先设定签署人的平台子客企业,预设后只能以该平台子客企业身份去领取合同并完成签署`
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 第三方平台子客企业员工的标识OpenId，GenerateType为"CHANNEL"时可用，指定到具体参与人, 仅展示已经实名的经办人信息
	// 
	// 注： 
	// 如果传进来的<font color="red">OpenId已经实名并且加入企业， 则忽略Name，IdCardType，IdCardNumber，Mobile这四个入参</font>（会用此OpenId实名的身份证和登录的手机号覆盖）
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 签署完成后是否自动回跳
	// <ul><li>false：否, 签署完成不会自动跳转回来(默认)</li><li>true：是, 签署完成会自动跳转回来</li></ul>
	// 
	// 注: 
	// 1. 该参数<font color="red">只针对APP类型（电子签小程序跳转贵方小程序）场景</font> 的签署链接有效
	// 2. <font color="red">手机应用APP 或 微信小程序需要监控界面的返回走后序逻辑</font>, 微信小程序的文档可以参考[这个](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onShow-Object-object)
	// 3. <font color="red">电子签小程序跳转贵方APP，不支持自动跳转，必需用户手动点击完成按钮（微信的限制）</font> 
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil,omitempty" name:"AutoJumpBack"`

	// 签署完之后的H5页面的跳转链接，针对Endpoint为CHANNEL时有效，最大长度1000个字符。
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 生成的签署链接在签署页面隐藏的按钮列表，可设置如下：
	// 
	// <ul><li> **0** :合同签署页面更多操作按钮</li>
	// <li> **1** :合同签署页面更多操作的拒绝签署按钮</li>
	// <li> **2** :合同签署页面更多操作的转他人处理按钮</li>
	// <li> **3** :签署成功页的查看详情按钮</li></ul>
	// 
	// 注:  `字段为数组, 可以传值隐藏多个按钮`
	Hides []*int64 `json:"Hides,omitnil,omitempty" name:"Hides"`

	// 参与方角色ID，用于生成动态签署人链接完成领取。
	// 
	// 注：`使用此参数需要与flow_ids数量一致并且一一对应, 表示在对应同序号的流程中的参与角色ID`，
	RecipientIds []*string `json:"RecipientIds,omitnil,omitempty" name:"RecipientIds"`

	// 合同组相关信息，指定合同组子合同和签署方的信息，用于补充动态签署人。
	FlowGroupUrlInfo *FlowGroupUrlInfo `json:"FlowGroupUrlInfo,omitnil,omitempty" name:"FlowGroupUrlInfo"`

	// <font color="red">仅公众号 H5 跳转电子签小程序时</font>，如需签署完成的“返回应用”功能，在获取签署链接接口的 UrlUseEnv 参数需设置为 **WeChatOfficialAccounts**，小程序签署成功的结果页面中才会出现“返回应用”按钮。在用户点击“返回应用”按钮之后，会返回到公众号 H5。 
	// 
	// 参考 [公众号 H5 跳转电子签小程序](https://qian.tencent.com/developers/company/openwxminiprogram/#23-%E5%85%AC%E4%BC%97%E5%8F%B7-h5-%E4%B8%AD%E8%B7%B3%E8%BD%AC)。
	UrlUseEnv *string `json:"UrlUseEnv,omitnil,omitempty" name:"UrlUseEnv"`
}

type CreateSignUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID数组，最多支持100个。
	// 
	// 注: 
	// 1. 必须选择提供此参数或合同组编号中的一个。
	// 2. 当生成类型（GenerateType）设为“ALL”时，不可提供多个流程ID。
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 合同组编号
	// 注：`该参数和合同流程ID数组必须二选一`
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 签署链接类型,可以设置的参数如下
	// <ul><li> **WEIXINAPP** :(默认)跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型 ，此时返回短链</li>
	// <li> **CHANNEL** :带有H5引导页的跳转电子签小程序的链接(<b>GenerateType非ALL时候不能设置成CHANNEL</b>)</li>
	// <li> **APP** :第三方App或小程序跳转电子签小程序的path, App或者小程序跳转适合此类型</li>
	// <li> **LONGURL2WEIXINAPP** :跳转电子签小程序的链接, H5跳转适合此类型，此时返回长链</li></ul>
	// 
	// **注：**动态签署人场景，如果签署链接类型设置为`APP`，则仅支持跳转到封面页。
	// 
	// 详细使用场景可以参考接口描述说明中的 **主要使用场景EndPoint分类**
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 签署链接生成类型，可以选择的类型如下
	// 
	// <ul><li><strong>ALL</strong>：（默认）为所有签署方生成签署链接，但不包括自动签署（静默签署）的签署方。注意：<strong>此中类型不支持多个合同ID（FlowIds）</strong>。</li>
	// <li><strong>CHANNEL</strong>：适用于第三方子企业的员工签署方。</li>
	// <li><strong>NOT_CHANNEL</strong>：适用于SaaS平台企业的员工签署方。</li>
	// <li><strong>PERSON</strong>：适用于个人或自然人签署方。</li>
	// <li><strong>FOLLOWER</strong>：适用于关注方，目前指合同的抄送方。</li>
	// <li><strong>RECIPIENT</strong>：根据RecipientId生成对应的签署链接，适用于动态添加签署人的情况。</li></ul>
	GenerateType *string `json:"GenerateType,omitnil,omitempty" name:"GenerateType"`

	// SaaS平台企业员工签署方的企业名称如果名称中包含英文括号()，请使用中文括号（）代替。  注:  `1.GenerateType为"NOT_CHANNEL"时必填` `2.获取B端动态签署人领取链接时,可指定此字段来预先设定签署人的企业,预设后只能以该企业身份去领取合同并完成签署`
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 合同流程里边参与方的姓名。
	// 注: 
	// 1. `GenerateType为"PERSON"(即个人签署方)时必填`。
	// 2. `在动态签署人补充链接场景中，可以通过传入这个值，对补充的个人参与方信息进行限制。仅匹配传入姓名的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方。`
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 合同流程里边签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 注:  
	// 1. `GenerateType为"PERSON"或"FOLLOWER"时必填。`
	// 2. `在动态签署人补充链接场景中，可以通过传入此值，对补充的个人参与方信息进行限制。仅匹配传入手机号的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方。`
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	// 
	// `注：在动态签署人补充链接场景中，可以通过传入此值，对补充的个人参与方信息进行限制。仅匹配传入证件类型的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方，且需要和证件号参数一同传递，不能单独进行限制。`
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成(如存在X，请大写)。</li>
	// <li>港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	// 
	// `注：在动态签署人补充链接场景中，可以通过传入此值，对补充的个人参与方信息进行限制。仅匹配传入证件号的参与方才能补充合同。此参数预设信息功能暂时仅支持个人动态参与方。`
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 第三方平台子客企业的企业的标识, 即OrganizationOpenId。 注:  `1.GenerateType为"CHANNEL"时必填` `2.获取B端动态签署人领取链接时,可指定此字段来预先设定签署人的平台子客企业,预设后只能以该平台子客企业身份去领取合同并完成签署`
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 第三方平台子客企业员工的标识OpenId，GenerateType为"CHANNEL"时可用，指定到具体参与人, 仅展示已经实名的经办人信息
	// 
	// 注： 
	// 如果传进来的<font color="red">OpenId已经实名并且加入企业， 则忽略Name，IdCardType，IdCardNumber，Mobile这四个入参</font>（会用此OpenId实名的身份证和登录的手机号覆盖）
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 签署完成后是否自动回跳
	// <ul><li>false：否, 签署完成不会自动跳转回来(默认)</li><li>true：是, 签署完成会自动跳转回来</li></ul>
	// 
	// 注: 
	// 1. 该参数<font color="red">只针对APP类型（电子签小程序跳转贵方小程序）场景</font> 的签署链接有效
	// 2. <font color="red">手机应用APP 或 微信小程序需要监控界面的返回走后序逻辑</font>, 微信小程序的文档可以参考[这个](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onShow-Object-object)
	// 3. <font color="red">电子签小程序跳转贵方APP，不支持自动跳转，必需用户手动点击完成按钮（微信的限制）</font> 
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil,omitempty" name:"AutoJumpBack"`

	// 签署完之后的H5页面的跳转链接，针对Endpoint为CHANNEL时有效，最大长度1000个字符。
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 生成的签署链接在签署页面隐藏的按钮列表，可设置如下：
	// 
	// <ul><li> **0** :合同签署页面更多操作按钮</li>
	// <li> **1** :合同签署页面更多操作的拒绝签署按钮</li>
	// <li> **2** :合同签署页面更多操作的转他人处理按钮</li>
	// <li> **3** :签署成功页的查看详情按钮</li></ul>
	// 
	// 注:  `字段为数组, 可以传值隐藏多个按钮`
	Hides []*int64 `json:"Hides,omitnil,omitempty" name:"Hides"`

	// 参与方角色ID，用于生成动态签署人链接完成领取。
	// 
	// 注：`使用此参数需要与flow_ids数量一致并且一一对应, 表示在对应同序号的流程中的参与角色ID`，
	RecipientIds []*string `json:"RecipientIds,omitnil,omitempty" name:"RecipientIds"`

	// 合同组相关信息，指定合同组子合同和签署方的信息，用于补充动态签署人。
	FlowGroupUrlInfo *FlowGroupUrlInfo `json:"FlowGroupUrlInfo,omitnil,omitempty" name:"FlowGroupUrlInfo"`

	// <font color="red">仅公众号 H5 跳转电子签小程序时</font>，如需签署完成的“返回应用”功能，在获取签署链接接口的 UrlUseEnv 参数需设置为 **WeChatOfficialAccounts**，小程序签署成功的结果页面中才会出现“返回应用”按钮。在用户点击“返回应用”按钮之后，会返回到公众号 H5。 
	// 
	// 参考 [公众号 H5 跳转电子签小程序](https://qian.tencent.com/developers/company/openwxminiprogram/#23-%E5%85%AC%E4%BC%97%E5%8F%B7-h5-%E4%B8%AD%E8%B7%B3%E8%BD%AC)。
	UrlUseEnv *string `json:"UrlUseEnv,omitnil,omitempty" name:"UrlUseEnv"`
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
	delete(f, "FlowGroupId")
	delete(f, "Endpoint")
	delete(f, "GenerateType")
	delete(f, "OrganizationName")
	delete(f, "Name")
	delete(f, "Mobile")
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "OrganizationOpenId")
	delete(f, "OpenId")
	delete(f, "AutoJumpBack")
	delete(f, "JumpUrl")
	delete(f, "Operator")
	delete(f, "Hides")
	delete(f, "RecipientIds")
	delete(f, "FlowGroupUrlInfo")
	delete(f, "UrlUseEnv")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSignUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignUrlsResponseParams struct {
	// 生成的签署参与者的签署链接信息数组。
	SignUrlInfos []*SignUrlInfo `json:"SignUrlInfos,omitnil,omitempty" name:"SignUrlInfos"`

	// 生成失败时的错误信息，成功返回”“，顺序和出参SignUrlInfos保持一致
	ErrorMessages []*string `json:"ErrorMessages,omitnil,omitempty" name:"ErrorMessages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DeleteOrganizationAuthorizationInfo struct {
	// 认证流 Id 是指在企业认证过程中，当前操作人的认证流程的唯一标识。每个企业在认证过程中只能有一条认证流认证成功。这意味着在同一认证过程内，一个企业只能有一个认证流程处于成功状态，以确保认证的唯一性和有效性。	
	AuthorizationId *string `json:"AuthorizationId,omitnil,omitempty" name:"AuthorizationId"`

	// 认证的企业名称	
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 第三方平台子客企业的唯一标识，定义Agent中的ProxyOrganizationOpenId一样, 可以参考<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#agent" target="_blank">Agent结构体</a>
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 清除认证流产生的错误信息
	Errormessage *string `json:"Errormessage,omitnil,omitempty" name:"Errormessage"`
}

// Predefined struct for user
type DeleteOrganizationAuthorizationsRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 认证人姓名，组织机构超管姓名。 在注册流程中，必须是超管本人进行操作。 
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 认证人手机号，组织机构超管手机号。 在注册流程中，必须是超管本人进行操作。 
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`
}

type DeleteOrganizationAuthorizationsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 认证人姓名，组织机构超管姓名。 在注册流程中，必须是超管本人进行操作。 
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 认证人手机号，组织机构超管手机号。 在注册流程中，必须是超管本人进行操作。 
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`
}

func (r *DeleteOrganizationAuthorizationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationAuthorizationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "AdminName")
	delete(f, "AdminMobile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationAuthorizationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationAuthorizationsResponseParams struct {
	// 清理认证流的详细信息，包括企业名称、认证流唯一 ID 以及清理认证流过程中产生的错误信息。
	DeleteOrganizationAuthorizationInfos []*DeleteOrganizationAuthorizationInfo `json:"DeleteOrganizationAuthorizationInfos,omitnil,omitempty" name:"DeleteOrganizationAuthorizationInfos"`

	// 批量清理认证流返回的状态值其中包括- 1 全部成功- 2 部分成功- 3 全部失败
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationAuthorizationsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationAuthorizationsResponseParams `json:"Response"`
}

func (r *DeleteOrganizationAuthorizationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationAuthorizationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Department struct {
	// 部门id
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 部门名称
	DepartmentName *string `json:"DepartmentName,omitnil,omitempty" name:"DepartmentName"`
}

// Predefined struct for user
type DescribeBatchOrganizationRegistrationTasksRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 企业批量认证链接的子任务 SubTaskId，该 SubTaskId 是通过接口[查询企业批量认证链接](https://qian.tencent.com/developers/companyApis/organizations/DescribeBatchOrganizationRegistrationUrls)可以得到。
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

type DescribeBatchOrganizationRegistrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 企业批量认证链接的子任务 SubTaskId，该 SubTaskId 是通过接口[查询企业批量认证链接](https://qian.tencent.com/developers/companyApis/organizations/DescribeBatchOrganizationRegistrationUrls)可以得到。
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

func (r *DescribeBatchOrganizationRegistrationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOrganizationRegistrationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "TaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchOrganizationRegistrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOrganizationRegistrationTasksResponseParams struct {
	// 企业批量任务状态明细
	Details []*BatchOrganizationRegistrationTasksDetails `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBatchOrganizationRegistrationTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchOrganizationRegistrationTasksResponseParams `json:"Response"`
}

func (r *DescribeBatchOrganizationRegistrationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOrganizationRegistrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOrganizationRegistrationUrlsRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 通过接口<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks" target="_blank">提交子企业批量认证链接创建任务</a>调用得到的任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeBatchOrganizationRegistrationUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 通过接口<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks" target="_blank">提交子企业批量认证链接创建任务</a>调用得到的任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeBatchOrganizationRegistrationUrlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOrganizationRegistrationUrlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchOrganizationRegistrationUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOrganizationRegistrationUrlsResponseParams struct {
	// 子企业注册认证的链接列表
	OrganizationAuthUrls []*OrganizationAuthUrl `json:"OrganizationAuthUrls,omitnil,omitempty" name:"OrganizationAuthUrls"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBatchOrganizationRegistrationUrlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchOrganizationRegistrationUrlsResponseParams `json:"Response"`
}

func (r *DescribeBatchOrganizationRegistrationUrlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOrganizationRegistrationUrlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCancelFlowsTaskRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 批量撤销任务编号，为32位字符串，通过接口[批量撤销合同流程](https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelBatchCancelFlows)或者[获取批量撤销签署流程腾讯电子签小程序链接](https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateBatchCancelFlowUrl)获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeCancelFlowsTaskRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 批量撤销任务编号，为32位字符串，通过接口[批量撤销合同流程](https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelBatchCancelFlows)或者[获取批量撤销签署流程腾讯电子签小程序链接](https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateBatchCancelFlowUrl)获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeCancelFlowsTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCancelFlowsTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCancelFlowsTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCancelFlowsTaskResponseParams struct {
	// 批量撤销任务编号，为32位字符串。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态，需要关注的状态<ul><li>**PROCESSING**  - 任务执行中</li><li>**END** - 任务处理完成</li><li>**TIMEOUT** 任务超时未处理完成，用户未在批量撤销链接有效期内操作</li></ul>
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 批量撤销成功的签署流程编号
	SuccessFlowIds []*string `json:"SuccessFlowIds,omitnil,omitempty" name:"SuccessFlowIds"`

	// 批量撤销失败的签署流程信息
	FailureFlows []*CancelFailureFlow `json:"FailureFlows,omitnil,omitempty" name:"FailureFlows"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCancelFlowsTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCancelFlowsTaskResponseParams `json:"Response"`
}

func (r *DescribeCancelFlowsTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCancelFlowsTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelFlowEvidenceReportRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 签署报告编号, 由<a href="https://qian.tencent.com/developers/partnerApis/certificate/CreateChannelFlowEvidenceReport" target="_blank">提交申请出证报告任务</a>产生
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定申请的报告类型，可选类型如下：
	// <ul><li> **0** :合同签署报告（默认）</li>
	// <li> **1** :公证处核验报告</li></ul>
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`
}

type DescribeChannelFlowEvidenceReportRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 签署报告编号, 由<a href="https://qian.tencent.com/developers/partnerApis/certificate/CreateChannelFlowEvidenceReport" target="_blank">提交申请出证报告任务</a>产生
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定申请的报告类型，可选类型如下：
	// <ul><li> **0** :合同签署报告（默认）</li>
	// <li> **1** :公证处核验报告</li></ul>
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`
}

func (r *DescribeChannelFlowEvidenceReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelFlowEvidenceReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "ReportId")
	delete(f, "Operator")
	delete(f, "ReportType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChannelFlowEvidenceReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelFlowEvidenceReportResponseParams struct {
	// 出证报告PDF的下载 URL，有效期为5分钟，超过有效期后将无法再下载。
	ReportUrl *string `json:"ReportUrl,omitnil,omitempty" name:"ReportUrl"`

	// 出证任务执行的状态, 状态含义如下：
	// 
	// <ul><li>**EvidenceStatusExecuting**：  出证任务在执行中</li>
	// <li>**EvidenceStatusSuccess**：  出证任务执行成功</li>
	// <li>**EvidenceStatusFailed** ： 出证任务执行失败</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChannelFlowEvidenceReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChannelFlowEvidenceReportResponseParams `json:"Response"`
}

func (r *DescribeChannelFlowEvidenceReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelFlowEvidenceReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelOrganizationsRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// 
	// 渠道应用标识: Agent.AppId
	// 第三方平台子客企业标识: Agent.ProxyOrganizationOpenId
	// 第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId
	// 
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定分页每页返回的数据条数，单页最大支持 200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 该字段是指第三方平台子客企业的唯一标识，用于查询单独某个子客的企业数据。
	// 
	// **注**：`如果需要批量查询本应用下的所有企业的信息，则该字段不需要赋值`
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 可以按照当前企业的认证状态进行过滤。可值如下：
	// <ul><li>**"UNVERIFIED"**： 未认证的企业</li>
	//   <li>**"VERIFYINGLEGALPENDINGAUTHORIZATION"**： 认证中待法人授权的企业</li>
	//   <li>**"VERIFYINGAUTHORIZATIONFILEPENDING"**： 认证中授权书审核中的企业</li>
	//   <li>**"VERIFYINGAUTHORIZATIONFILEREJECT"**： 认证中授权书已驳回的企业</li>
	//   <li>**"VERIFYING"**： 认证进行中的企业</li>
	//   <li>**"VERIFIED"**： 已认证完成的企业</li></ul>
	AuthorizationStatusList []*string `json:"AuthorizationStatusList,omitnil,omitempty" name:"AuthorizationStatusList"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页。 页码从 0 开始，即首页为 0，最大20000。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeChannelOrganizationsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// 
	// 渠道应用标识: Agent.AppId
	// 第三方平台子客企业标识: Agent.ProxyOrganizationOpenId
	// 第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId
	// 
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定分页每页返回的数据条数，单页最大支持 200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 该字段是指第三方平台子客企业的唯一标识，用于查询单独某个子客的企业数据。
	// 
	// **注**：`如果需要批量查询本应用下的所有企业的信息，则该字段不需要赋值`
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 可以按照当前企业的认证状态进行过滤。可值如下：
	// <ul><li>**"UNVERIFIED"**： 未认证的企业</li>
	//   <li>**"VERIFYINGLEGALPENDINGAUTHORIZATION"**： 认证中待法人授权的企业</li>
	//   <li>**"VERIFYINGAUTHORIZATIONFILEPENDING"**： 认证中授权书审核中的企业</li>
	//   <li>**"VERIFYINGAUTHORIZATIONFILEREJECT"**： 认证中授权书已驳回的企业</li>
	//   <li>**"VERIFYING"**： 认证进行中的企业</li>
	//   <li>**"VERIFIED"**： 已认证完成的企业</li></ul>
	AuthorizationStatusList []*string `json:"AuthorizationStatusList,omitnil,omitempty" name:"AuthorizationStatusList"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页。 页码从 0 开始，即首页为 0，最大20000。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeChannelOrganizationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelOrganizationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "Limit")
	delete(f, "OrganizationOpenId")
	delete(f, "AuthorizationStatusList")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChannelOrganizationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelOrganizationsResponseParams struct {
	// 满足查询条件的企业信息列表。
	ChannelOrganizationInfos []*ChannelOrganizationInfo `json:"ChannelOrganizationInfos,omitnil,omitempty" name:"ChannelOrganizationInfos"`

	// 指定分页返回第几页的数据。页码从 0 开始，即首页为 0，最大20000。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定分页每页返回的数据条数，单页最大支持 200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 满足查询条件的企业总数量。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChannelOrganizationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChannelOrganizationsResponseParams `json:"Response"`
}

func (r *DescribeChannelOrganizationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelOrganizationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelSealPolicyWorkflowUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// 
	// 渠道应用标识: Agent.AppId
	// 第三方平台子客企业标识: Agent.ProxyOrganizationOpenId
	// 第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 用印审批单的ID，可通过用印[申请回调](https://qian.tencent.com/developers/partner/callback_types_seals#%E4%B8%89-%E7%94%A8%E5%8D%B0%E7%94%B3%E8%AF%B7%E5%AE%A1%E6%89%B9%E7%8A%B6%E6%80%81%E9%80%9A%E7%9F%A5)获取。
	WorkflowInstanceId *string `json:"WorkflowInstanceId,omitnil,omitempty" name:"WorkflowInstanceId"`

	// 生成链接的类型：
	// 生成链接的类型
	// <ul><li>**LongLink**：(默认)长链接，H5跳转到电子签小程序链接，链接有效期为1年</li>
	// <li>**ShortLink**：H5跳转到电子签小程序链接，一般用于发送短信中带的链接，打开后进入腾讯电子签小程序，链接有效期为7天</li>
	// <li>**App**：第三方APP或小程序跳转电子签小程序链接，一般用于贵方小程序或者APP跳转过来，打开后进入腾讯电子签小程序，链接有效期为1年</li></ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

type DescribeChannelSealPolicyWorkflowUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// 
	// 渠道应用标识: Agent.AppId
	// 第三方平台子客企业标识: Agent.ProxyOrganizationOpenId
	// 第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 用印审批单的ID，可通过用印[申请回调](https://qian.tencent.com/developers/partner/callback_types_seals#%E4%B8%89-%E7%94%A8%E5%8D%B0%E7%94%B3%E8%AF%B7%E5%AE%A1%E6%89%B9%E7%8A%B6%E6%80%81%E9%80%9A%E7%9F%A5)获取。
	WorkflowInstanceId *string `json:"WorkflowInstanceId,omitnil,omitempty" name:"WorkflowInstanceId"`

	// 生成链接的类型：
	// 生成链接的类型
	// <ul><li>**LongLink**：(默认)长链接，H5跳转到电子签小程序链接，链接有效期为1年</li>
	// <li>**ShortLink**：H5跳转到电子签小程序链接，一般用于发送短信中带的链接，打开后进入腾讯电子签小程序，链接有效期为7天</li>
	// <li>**App**：第三方APP或小程序跳转电子签小程序链接，一般用于贵方小程序或者APP跳转过来，打开后进入腾讯电子签小程序，链接有效期为1年</li></ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

func (r *DescribeChannelSealPolicyWorkflowUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelSealPolicyWorkflowUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "WorkflowInstanceId")
	delete(f, "Endpoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChannelSealPolicyWorkflowUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelSealPolicyWorkflowUrlResponseParams struct {
	// 用印审批小程序链接，链接类型（通过H5唤起小程序或通过APP跳转方式查看）。
	WorkflowUrl *string `json:"WorkflowUrl,omitnil,omitempty" name:"WorkflowUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChannelSealPolicyWorkflowUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChannelSealPolicyWorkflowUrlResponseParams `json:"Response"`
}

func (r *DescribeChannelSealPolicyWorkflowUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelSealPolicyWorkflowUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtendedServiceAuthDetailRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要查询的扩展服务类型。
	// 如下所示：
	// <ul><li> AUTO_SIGN：企业静默签署</li>
	// <li>BATCH_SIGN：批量签署</li>
	// </ul>
	ExtendServiceType *string `json:"ExtendServiceType,omitnil,omitempty" name:"ExtendServiceType"`

	// 指定每页返回的数据条数，和Offset参数配合使用。 注：`1.默认值为20，单页做大值为200。`	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。 注：`1.offset从0开始，即第一页为0。` `2.默认从第一页返回。`	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeExtendedServiceAuthDetailRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要查询的扩展服务类型。
	// 如下所示：
	// <ul><li> AUTO_SIGN：企业静默签署</li>
	// <li>BATCH_SIGN：批量签署</li>
	// </ul>
	ExtendServiceType *string `json:"ExtendServiceType,omitnil,omitempty" name:"ExtendServiceType"`

	// 指定每页返回的数据条数，和Offset参数配合使用。 注：`1.默认值为20，单页做大值为200。`	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。 注：`1.offset从0开始，即第一页为0。` `2.默认从第一页返回。`	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeExtendedServiceAuthDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtendedServiceAuthDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "ExtendServiceType")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtendedServiceAuthDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtendedServiceAuthDetailResponseParams struct {
	// 服务授权的信息列表，根据查询类型返回特定扩展服务的开通和授权状况。
	AuthInfoDetail *AuthInfoDetail `json:"AuthInfoDetail,omitnil,omitempty" name:"AuthInfoDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExtendedServiceAuthDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtendedServiceAuthDetailResponseParams `json:"Response"`
}

func (r *DescribeExtendedServiceAuthDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtendedServiceAuthDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtendedServiceAuthInfoRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DescribeExtendedServiceAuthInfoRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *DescribeExtendedServiceAuthInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtendedServiceAuthInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtendedServiceAuthInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtendedServiceAuthInfoResponseParams struct {
	// 服务开通和授权的信息列表，根据查询类型返回所有支持的扩展服务开通和授权状况，或者返回特定扩展服务的开通和授权状况。
	AuthInfo []*ExtentServiceAuthInfo `json:"AuthInfo,omitnil,omitempty" name:"AuthInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExtendedServiceAuthInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtendedServiceAuthInfoResponseParams `json:"Response"`
}

func (r *DescribeExtendedServiceAuthInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtendedServiceAuthInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowDetailInfoRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要查询的流程ID列表，最多可传入100个ID。
	// 如果要查询合同组的信息，则不需要传入此参数，只需传入 FlowGroupId 参数即可。
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 需要查询的流程组ID，如果传入此参数，则会忽略 FlowIds 参数。
	// 
	// 合同组由<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByTemplates" target="_blank">通过多模板创建合同组签署流程</a>和<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByFiles" target="_blank">通过多文件创建合同组签署流程</a>等接口创建。
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type DescribeFlowDetailInfoRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要查询的流程ID列表，最多可传入100个ID。
	// 如果要查询合同组的信息，则不需要传入此参数，只需传入 FlowGroupId 参数即可。
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 需要查询的流程组ID，如果传入此参数，则会忽略 FlowIds 参数。
	// 
	// 合同组由<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByTemplates" target="_blank">通过多模板创建合同组签署流程</a>和<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByFiles" target="_blank">通过多文件创建合同组签署流程</a>等接口创建。
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	delete(f, "FlowGroupId")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowDetailInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowDetailInfoResponseParams struct {
	// 合同归属的第三方平台应用号ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 合同归属的第三方平台子客企业OpenId
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil,omitempty" name:"ProxyOrganizationOpenId"`

	// 合同流程的详细信息。
	// 如果查询的是合同组信息，则返回的是组内所有子合同流程的详细信息。
	FlowInfo []*FlowDetailInfo `json:"FlowInfo,omitnil,omitempty" name:"FlowInfo"`

	// 合同组ID，只有在查询合同组信息时才会返回。
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 合同组名称，只有在查询合同组信息时才会返回。
	FlowGroupName *string `json:"FlowGroupName,omitnil,omitempty" name:"FlowGroupName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要下载的合同流程的ID,  至少需要1个,  做多50个
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type DescribeResourceUrlsByFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要下载的合同流程的ID,  至少需要1个,  做多50个
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	// 合同流程PDF下载链接
	FlowResourceUrlInfos []*FlowResourceUrlInfo `json:"FlowResourceUrlInfos,omitnil,omitempty" name:"FlowResourceUrlInfos"`

	// 如果某个序号的合同流程生成PDF下载链接成功, 对应序号的值为空
	// 如果某个序号的合同流程生成PDF下载链接失败, 对应序号的值为错误的原因
	ErrorMessages []*string `json:"ErrorMessages,omitnil,omitempty" name:"ErrorMessages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同模板ID，为32位字符串。
	// 
	// 可以通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>登录企业控制台, 在企业模板中得到合同模板ID。
	// 
	// [点击查看模板Id在控制台上的位置](https://qcloudimg.tencent-cloud.cn/raw/e988be12bf28a89b4716aed4502c2e02.png)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 查询模板的内容
	// 
	// <ul><li>**0**：（默认）模板列表及详情</li>
	// <li>**1**：仅模板列表, 不会返回模板中的签署控件, 填写控件, 参与方角色列表等信息</li></ul>
	ContentType *int64 `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 合同模板ID数组，每一个合同模板ID为32位字符串,  最多支持100个模板的批量查询。
	// 
	// 注意: 
	// 1.` 此参数TemplateIds与TemplateId互为独立，若两者均传入，以TemplateId为准。`
	// 2. `请确保每个模板均正确且属于当前企业，若有任一模板不存在，则返回错误。`
	// 4. `若传递此参数，分页参数(Limit,Offset)无效`
	// 
	// 
	// [点击查看模板Id在控制台上的位置](https://qcloudimg.tencent-cloud.cn/raw/e988be12bf28a89b4716aed4502c2e02.png)
	TemplateIds []*string `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// 指定每页返回的数据条数，和Offset参数配合使用。
	// 
	// 注：`1.默认值为20，单页做大值为100。`
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。
	// 
	// 注：`1.offset从0开始，即第一页为0。`
	// `2.默认从第一页返回。`
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 模糊搜索的模板名称，注意是模板名的连续部分，长度不能超过200，可支持由中文、字母、数字和下划线组成字符串。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 对应第三方应用平台企业的模板ID，通过此值可以搜索由第三方应用平台模板ID下发或领取得到的子客模板列表。
	ChannelTemplateId *string `json:"ChannelTemplateId,omitnil,omitempty" name:"ChannelTemplateId"`

	// 返回控件的范围, 是只返回发起方自己的还是所有参与方的
	// 
	// <ul><li>**false**：（默认）只返回发起方控件</li>
	// <li>**true**：返回所有参与方(包括发起方和签署方)控件</li></ul>
	QueryAllComponents *bool `json:"QueryAllComponents,omitnil,omitempty" name:"QueryAllComponents"`

	// 是否获取模板预览链接。
	// 
	// <ul><li>**false**：不获取（默认）</li>
	// <li>**true**：获取</li></ul>
	// 
	// 设置为true之后， 返回参数PreviewUrl，为模板的H5预览链接,  有效期5分钟。可以通过浏览器打开此链接预览模板，或者嵌入到iframe中预览模板。
	WithPreviewUrl *bool `json:"WithPreviewUrl,omitnil,omitempty" name:"WithPreviewUrl"`

	// 是否获取模板的PDF文件链接。
	// 
	// <ul><li>**false**：不获取（默认）</li>
	// <li>**true**：获取</li></ul>
	// 
	// 设置为true之后， 返回参数PdfUrl，为模板PDF文件链接，有效期5分钟, 可以用于将PDF文件下载到本地
	// 
	// 注: `此功能需要开通功能白名单【第三方应用集成企业获取模版PDF下载链接】，使用前请联系对接的客户经理沟通。`
	WithPdfUrl *bool `json:"WithPdfUrl,omitnil,omitempty" name:"WithPdfUrl"`

	// 操作者的信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 用户合同类型id
	UserFlowTypeId *string `json:"UserFlowTypeId,omitnil,omitempty" name:"UserFlowTypeId"`
}

type DescribeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同模板ID，为32位字符串。
	// 
	// 可以通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>登录企业控制台, 在企业模板中得到合同模板ID。
	// 
	// [点击查看模板Id在控制台上的位置](https://qcloudimg.tencent-cloud.cn/raw/e988be12bf28a89b4716aed4502c2e02.png)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 查询模板的内容
	// 
	// <ul><li>**0**：（默认）模板列表及详情</li>
	// <li>**1**：仅模板列表, 不会返回模板中的签署控件, 填写控件, 参与方角色列表等信息</li></ul>
	ContentType *int64 `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 合同模板ID数组，每一个合同模板ID为32位字符串,  最多支持100个模板的批量查询。
	// 
	// 注意: 
	// 1.` 此参数TemplateIds与TemplateId互为独立，若两者均传入，以TemplateId为准。`
	// 2. `请确保每个模板均正确且属于当前企业，若有任一模板不存在，则返回错误。`
	// 4. `若传递此参数，分页参数(Limit,Offset)无效`
	// 
	// 
	// [点击查看模板Id在控制台上的位置](https://qcloudimg.tencent-cloud.cn/raw/e988be12bf28a89b4716aed4502c2e02.png)
	TemplateIds []*string `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`

	// 指定每页返回的数据条数，和Offset参数配合使用。
	// 
	// 注：`1.默认值为20，单页做大值为100。`
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。
	// 
	// 注：`1.offset从0开始，即第一页为0。`
	// `2.默认从第一页返回。`
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 模糊搜索的模板名称，注意是模板名的连续部分，长度不能超过200，可支持由中文、字母、数字和下划线组成字符串。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 对应第三方应用平台企业的模板ID，通过此值可以搜索由第三方应用平台模板ID下发或领取得到的子客模板列表。
	ChannelTemplateId *string `json:"ChannelTemplateId,omitnil,omitempty" name:"ChannelTemplateId"`

	// 返回控件的范围, 是只返回发起方自己的还是所有参与方的
	// 
	// <ul><li>**false**：（默认）只返回发起方控件</li>
	// <li>**true**：返回所有参与方(包括发起方和签署方)控件</li></ul>
	QueryAllComponents *bool `json:"QueryAllComponents,omitnil,omitempty" name:"QueryAllComponents"`

	// 是否获取模板预览链接。
	// 
	// <ul><li>**false**：不获取（默认）</li>
	// <li>**true**：获取</li></ul>
	// 
	// 设置为true之后， 返回参数PreviewUrl，为模板的H5预览链接,  有效期5分钟。可以通过浏览器打开此链接预览模板，或者嵌入到iframe中预览模板。
	WithPreviewUrl *bool `json:"WithPreviewUrl,omitnil,omitempty" name:"WithPreviewUrl"`

	// 是否获取模板的PDF文件链接。
	// 
	// <ul><li>**false**：不获取（默认）</li>
	// <li>**true**：获取</li></ul>
	// 
	// 设置为true之后， 返回参数PdfUrl，为模板PDF文件链接，有效期5分钟, 可以用于将PDF文件下载到本地
	// 
	// 注: `此功能需要开通功能白名单【第三方应用集成企业获取模版PDF下载链接】，使用前请联系对接的客户经理沟通。`
	WithPdfUrl *bool `json:"WithPdfUrl,omitnil,omitempty" name:"WithPdfUrl"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 用户合同类型id
	UserFlowTypeId *string `json:"UserFlowTypeId,omitnil,omitempty" name:"UserFlowTypeId"`
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
	delete(f, "TemplateIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "TemplateName")
	delete(f, "ChannelTemplateId")
	delete(f, "QueryAllComponents")
	delete(f, "WithPreviewUrl")
	delete(f, "WithPdfUrl")
	delete(f, "Operator")
	delete(f, "UserFlowTypeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplatesResponseParams struct {
	// 模板详情列表数据
	Templates []*TemplateInfo `json:"Templates,omitnil,omitempty" name:"Templates"`

	// 查询到的模板总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 每页返回的数据条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结果分页返回，此处指定第几页。页码从0开始，即首页为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 查询日期范围的开始时间, 查询会包含此日期的数据 , 格式为yyyy-mm-dd (例如：2021-03-21)
	// 
	// 注: `查询日期范围区间长度大于90天`。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 查询日期范围的结束时间, 查询会包含此日期的数据 , 格式为yyyy-mm-dd (例如：2021-04-21)
	// 
	// 注: `查询日期范围区间长度大于90天`。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 是否汇总数据，默认不汇总。
	// <ul><li> **true** :  汇总数据,  即每个企业一条数据, 对日志范围内的数据相加</li>
	// <li> **false** :  不会总数据,  返回企业每日明细,   按日期返回每个企业的数据(如果企业对应天数没有操作则无此企业此日期的数据)</li></ul>
	NeedAggregate *bool `json:"NeedAggregate,omitnil,omitempty" name:"NeedAggregate"`

	// 指定每页返回的数据条数，和Offset参数配合使用。
	// 
	// 注: `默认值为1000，单页做大值为1000`
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。
	// 
	// 注：`offset从0开始，即第一页为0。`
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type DescribeUsageRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 查询日期范围的开始时间, 查询会包含此日期的数据 , 格式为yyyy-mm-dd (例如：2021-03-21)
	// 
	// 注: `查询日期范围区间长度大于90天`。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 查询日期范围的结束时间, 查询会包含此日期的数据 , 格式为yyyy-mm-dd (例如：2021-04-21)
	// 
	// 注: `查询日期范围区间长度大于90天`。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 是否汇总数据，默认不汇总。
	// <ul><li> **true** :  汇总数据,  即每个企业一条数据, 对日志范围内的数据相加</li>
	// <li> **false** :  不会总数据,  返回企业每日明细,   按日期返回每个企业的数据(如果企业对应天数没有操作则无此企业此日期的数据)</li></ul>
	NeedAggregate *bool `json:"NeedAggregate,omitnil,omitempty" name:"NeedAggregate"`

	// 指定每页返回的数据条数，和Offset参数配合使用。
	// 
	// 注: `默认值为1000，单页做大值为1000`
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。
	// 
	// 注：`offset从0开始，即第一页为0。`
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 用量明细
	Details []*UsageDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeUserFlowTypeRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。 此接口下面信息必填。 <ul> <li>渠道应用标识: Agent.AppId</li> <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li> <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li> </ul> 第三方平台子客企业和员工必须已经经过实名认证	
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 搜索过滤的条件，本字段允许您通过指定模板 ID 或模板名称来进行查询。 <ul><li><strong>模板的用户合同类型</strong>：<strong>Key</strong>设置为 <code>user-flow-type-id</code> ，<strong>Values</strong>为您想要查询的用户模版类型id列表。</li></ul>	
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询绑定了模版的用户合同类型
	// <ul>
	// <li>false（默认值），查询用户合同类型</li>
	// <li>true，查询绑定了模版的用户合同类型</li>
	// </ul>
	QueryBindTemplate *bool `json:"QueryBindTemplate,omitnil,omitempty" name:"QueryBindTemplate"`
}

type DescribeUserFlowTypeRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。 此接口下面信息必填。 <ul> <li>渠道应用标识: Agent.AppId</li> <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li> <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li> </ul> 第三方平台子客企业和员工必须已经经过实名认证	
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 搜索过滤的条件，本字段允许您通过指定模板 ID 或模板名称来进行查询。 <ul><li><strong>模板的用户合同类型</strong>：<strong>Key</strong>设置为 <code>user-flow-type-id</code> ，<strong>Values</strong>为您想要查询的用户模版类型id列表。</li></ul>	
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询绑定了模版的用户合同类型
	// <ul>
	// <li>false（默认值），查询用户合同类型</li>
	// <li>true，查询绑定了模版的用户合同类型</li>
	// </ul>
	QueryBindTemplate *bool `json:"QueryBindTemplate,omitnil,omitempty" name:"QueryBindTemplate"`
}

func (r *DescribeUserFlowTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserFlowTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "Filters")
	delete(f, "QueryBindTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserFlowTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserFlowTypeResponseParams struct {
	// 查询到的所有用户合同类型列表
	AllUserFlowTypes []*TemplateUserFlowType `json:"AllUserFlowTypes,omitnil,omitempty" name:"AllUserFlowTypes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserFlowTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserFlowTypeResponseParams `json:"Response"`
}

func (r *DescribeUserFlowTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserFlowTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectInfoVideoData struct {
	// 活体视频的base64编码，mp4格式
	// 
	// 注:`需进行base64解码获取活体视频文件`
	LiveNessVideo *string `json:"LiveNessVideo,omitnil,omitempty" name:"LiveNessVideo"`
}

type DownloadFlowInfo struct {
	// 文件夹名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 签署流程的标识数组
	FlowIdList []*string `json:"FlowIdList,omitnil,omitempty" name:"FlowIdList"`
}

type DynamicFlowApproverResult struct {
	// 签署流程签署人在模板中对应的签署人Id；在非单方签署、以及非B2C签署的场景下必传，用于指定当前签署方在签署流程中的位置；	
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 签署ID - 发起流程时系统自动补充 - 创建签署链接时，可以通过查询详情接口获得签署人的SignId，然后可传入此值为该签署人创建签署链接，无需再传姓名、手机号、证件号等其他信息	
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 签署人状态信息
	ApproverStatus *int64 `json:"ApproverStatus,omitnil,omitempty" name:"ApproverStatus"`
}

type DynamicFlowInfo struct {
	// 合同流程ID，为32位字符串。 - 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。 - 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。 - <font color="red">不建议继续使用</font>，请使用<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId指定合同	
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 合同流程的参与方列表, 最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，不同类型的签署方传参方式可以参考文档 [签署方入参指引](https://qian.tencent.com/developers/partner/flow_approver)。 如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序, 请确保列表中参与人的顺序符合实际签署顺序。	
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitnil,omitempty" name:"FlowApprovers"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传： <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：电子处方单（医疗自动签） </li><li> **OTHER** : 通用场景</li></ul> 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`	
	AutoSignScene *string `json:"AutoSignScene,omitnil,omitempty" name:"AutoSignScene"`

	// 签署人校验方式 VerifyCheck: 人脸识别（默认） MobileCheck：手机号验证，用户手机号和参与方手机号（ApproverMobile）相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证） 参数说明：可选人脸识别或手机号验证两种方式，若选择后者，未实名个人签署方在签署合同时，无需经过实名认证和意愿确认两次人脸识别，该能力仅适用于个人签署方。	
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil,omitempty" name:"ApproverVerifyType"`
}

type DynamicFlowResult struct {
	// 合同流程ID，为32位字符串。 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。 [点击查看FlowId在控制台上的位置](https://qcloudimg.tencent-cloud.cn/raw/05af26573d5106763b4cfbb9f7c64b41.png)	
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 动态合同签署人补充结果信息列表
	DynamicFlowApproverList []*DynamicFlowApproverResult `json:"DynamicFlowApproverList,omitnil,omitempty" name:"DynamicFlowApproverList"`
}

type EmbedUrlOption struct {
	// 合同详情预览，允许展示控件信息
	// <ul>
	// <li><b>true</b>：允许在合同详情页展示控件</li>
	// <li><b>false</b>：（默认）不允许在合同详情页展示控件</li>
	// </ul>
	ShowFlowDetailComponent *bool `json:"ShowFlowDetailComponent,omitnil,omitempty" name:"ShowFlowDetailComponent"`

	// 模板预览，允许展示模板控件信息
	// <ul><li> <b>true</b> :允许在模板预览页展示控件</li>
	// <li> <b>false</b> :（默认）不允许在模板预览页展示控件</li></ul>
	ShowTemplateComponent *bool `json:"ShowTemplateComponent,omitnil,omitempty" name:"ShowTemplateComponent"`

	// 跳过上传文件，默认为false(展示上传文件页）![image](https://qcloudimg.tencent-cloud.cn/raw/8ca33745cf772e79831dbe5a70e82400.png)
	// - false: 展示上传文件页
	// - true: 不展示上传文件页
	//  
	// 
	// 注意: 此参数仅针对**EmbedType=CREATE_TEMPLATE(创建模板)有效**，
	SkipUploadFile *string `json:"SkipUploadFile,omitnil,omitempty" name:"SkipUploadFile"`

	// 是否禁止编辑（展示）水印控件属性
	// <ul><li>（默认） false -否</li> <li> true - 禁止编辑</li></ul>
	ForbidEditWatermark *bool `json:"ForbidEditWatermark,omitnil,omitempty" name:"ForbidEditWatermark"`
}

type ExtentServiceAuthInfo struct {
	// 扩展服务类型
	// <ul>
	// <li>AUTO_SIGN             企业自动签（自动签署）</li>
	// <li>  OVERSEA_SIGN          企业与港澳台居民签署合同</li>
	// <li>  MOBILE_CHECK_APPROVER 使用手机号验证签署方身份</li>
	// <li> DOWNLOAD_FLOW         授权渠道下载合同 </li>
	// <li>AGE_LIMIT_EXPANSION 拓宽签署方年龄限制</li>
	// <li>HIDE_OPERATOR_DISPLAY 隐藏合同经办人姓名</li>
	// </ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 扩展服务名称 
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 扩展服务的开通状态
	// **ENABLE**：开通 
	// **DISABLE**：未开通	
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作扩展服务的操作人第三方应用平台的用户openid
	OperatorOpenId *string `json:"OperatorOpenId,omitnil,omitempty" name:"OperatorOpenId"`

	// 扩展服务的操作时间，格式为Unix标准时间戳（秒）。	
	OperateOn *int64 `json:"OperateOn,omitnil,omitempty" name:"OperateOn"`
}

type FailedCreateRoleData struct {
	// 用户userId
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 角色RoleId列表
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type FillApproverInfo struct {
	// 签署方经办人在模板中配置的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 指定企业经办签署人OpenId
	// 
	// 注: `签署人OpenId未实名时，需要传入签署人姓名以及手机号码。`
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 签署人姓名
	ApproverName *string `json:"ApproverName,omitnil,omitempty" name:"ApproverName"`

	// 签署人手机号码
	ApproverMobile *string `json:"ApproverMobile,omitnil,omitempty" name:"ApproverMobile"`

	// 企业名称
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 企业OpenId
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 签署企业非渠道子客，默认为false，即表示同一渠道下的企业；如果为true，则目前表示接收方企业为SaaS企业, 为渠道子客时，OrganizationOpenId 必传
	NotChannelOrganization *bool `json:"NotChannelOrganization,omitnil,omitempty" name:"NotChannelOrganization"`

	// 签署方经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD 中国大陆居民身份证</li>
	// <li>HONGKONG_AND_MACAO 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN 中国港澳台居民居住证(格式同中国大陆居民身份证)</li>
	// <li>OTHER_CARD_TYPE 其他证件</li></ul>
	// 
	// 注: `1.其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	// `2.补充个人签署方时，若该用户已在电子签完成实名则可通过指定姓名和证件类型、证件号码完成补充。`
	ApproverIdCardType *string `json:"ApproverIdCardType,omitnil,omitempty" name:"ApproverIdCardType"`

	// 签署方经办人的证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	// 
	// 注：`补充个人签署方时，若该用户已在电子签完成实名则可通过指定姓名和证件类型、证件号码完成补充。`
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitnil,omitempty" name:"ApproverIdCardNumber"`

	// 合同流程ID
	// - 补充合同组子合同动态签署人时必传。
	// - 补充正常合同，请阅读：<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelCreateFlowApprovers/" target="_blank">补充签署人接口</a>接口使用说明
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type FillError struct {
	// 为签署方经办人在签署合同中的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。与入参中补充的签署人角色ID对应，批量补充部分失败返回对应的错误信息。
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 补充失败错误说明
	ErrMessage *string `json:"ErrMessage,omitnil,omitempty" name:"ErrMessage"`

	// 合同流程ID，为32位字符串。	
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type FilledComponent struct {
	// 填写控件ID
	ComponentId *string `json:"ComponentId,omitnil,omitempty" name:"ComponentId"`

	// 控件名称
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// 此填写控件的填写状态
	//  **0** : 此填写控件**未填写**
	// **1** : 此填写控件**已填写**
	ComponentFillStatus *string `json:"ComponentFillStatus,omitnil,omitempty" name:"ComponentFillStatus"`

	// 控件填写内容
	ComponentValue *string `json:"ComponentValue,omitnil,omitempty" name:"ComponentValue"`

	// 图片填充控件下载链接，如果是图片填充控件时，这里返回图片的下载链接。
	// 
	// 注: `链接不是永久链接,  默认有效期5分钟后, 到期后链接失效`
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type Filter struct {
	// 查询过滤条件的Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 查询过滤条件的Value列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FlowApproverDetail struct {
	// 模板配置时候的签署人角色ID(用PDF文件发起也可以指定,如果不指定则自动生成此角色ID), 所有的填写控件和签署控件都归属不同的角色
	//
	// Deprecated: ReceiptId is deprecated.
	ReceiptId *string `json:"ReceiptId,omitnil,omitempty" name:"ReceiptId"`

	// 第三方平台子客企业的唯一标识，定义Agent中的ProxyOrganizationOpenId一样, 可以参考<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#agent" target="_blank">Agent结构体</a>
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil,omitempty" name:"ProxyOrganizationOpenId"`

	// 第三方平台子客企业员工的唯一标识
	ProxyOperatorOpenId *string `json:"ProxyOperatorOpenId,omitnil,omitempty" name:"ProxyOperatorOpenId"`

	// 第三方平台子客企业名称，与企业营业执照中注册的名称一致。
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil,omitempty" name:"ProxyOrganizationName"`

	// 签署人手机号
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 签署顺序，如果是有序签署，签署顺序从小到大
	SignOrder *int64 `json:"SignOrder,omitnil,omitempty" name:"SignOrder"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	ApproveName *string `json:"ApproveName,omitnil,omitempty" name:"ApproveName"`

	// 当前签署人的状态, 状态如下
	// <ul><li> **PENDING** :待签署</li>
	// <li> **FILLPENDING** :待填写</li>
	// <li> **FILLACCEPT** :填写完成</li>
	// <li> **FILLREJECT** :拒绝填写</li>
	// <li> **WAITPICKUP** :待领取</li>
	// <li> **ACCEPT** :已签署</li>
	// <li> **REJECT** :拒签</li>
	// <li> **DEADLINE** :过期没人处理</li>
	// <li> **CANCEL** :流程已撤回</li>
	// <li> **FORWARD** :已经转他人处理</li>
	// <li> **STOP** :流程已终止</li>
	// <li> **RELIEVED** :解除协议（已解除）</li></ul>
	ApproveStatus *string `json:"ApproveStatus,omitnil,omitempty" name:"ApproveStatus"`

	// 签署人拒签等情况的时候填写的原因
	ApproveMessage *string `json:"ApproveMessage,omitnil,omitempty" name:"ApproveMessage"`

	// 签署人签署时间戳，单位秒
	ApproveTime *int64 `json:"ApproveTime,omitnil,omitempty" name:"ApproveTime"`

	// 参与者类型 
	// <ul><li> **ORGANIZATION** :企业签署人</li>
	// <li> **PERSON** :个人签署人</li></ul>
	ApproveType *string `json:"ApproveType,omitnil,omitempty" name:"ApproveType"`

	// 自定义签署人的角色名, 如: 收款人、开具人、见证人等
	ApproverRoleName *string `json:"ApproverRoleName,omitnil,omitempty" name:"ApproverRoleName"`

	// 签署参与人在本流程中的编号ID（每个流程不同），可用此ID来定位签署参与人在本流程的签署节点。
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 模板配置时候的签署人角色ID(用PDF文件发起也可以指定,如果不指定则自动生成此角色ID), 所有的填写控件和签署控件都归属不同的角色
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`
}

type FlowApproverInfo struct {
	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 签署方经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li>
	// <li>OTHER_CARD_TYPE : 其他证件</li></ul>
	// 
	// 注: `其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 签署方经办人的证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)， 不支持海外手机号。
	// 请确认手机号所有方为此合同签署方。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 指定签署人非第三方平台子客企业下员工还是SaaS平台企业，在ApproverType为ORGANIZATION时指定。
	// <ul>
	// <li>false: 默认值，第三方平台子客企业下员工</li>
	// <li>true: SaaS平台企业下的员工</li>
	// </ul>
	NotChannelOrganization *bool `json:"NotChannelOrganization,omitnil,omitempty" name:"NotChannelOrganization"`

	// 第三方平台子客企业员工的唯一标识，长度不能超过64，只能由字母和数字组成
	// 
	// 当签署方为同一第三方平台下的员工时，该字段若不指定，则发起【待领取】的流程
	// 
	// 注： 
	// 如果传进来的<font color="red">OpenId已经实名并且加入企业， 则忽略Name，IdCardType，IdCardNumber，Mobile这四个入参</font>（会用此OpenId实名的身份证和登录的手机号覆盖）
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 同应用下第三方平台子客企业的唯一标识，定义Agent中的ProxyOrganizationOpenId一样，签署方为非发起方企业场景下必传，最大长度64个字符
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 在指定签署方时，可选择企业B端或个人C端等不同的参与者类型，可选类型如下:
	// <ul><li> **PERSON** :个人/自然人</li>
	// <li> **PERSON_AUTO_SIGN** :个人/自然人自动签署，适用于个人自动签场景</li>
	// <li> **ORGANIZATION** :企业/企业员工（企业签署方或模板发起时的企业静默签）</li>
	// <li> **ENTERPRISESERVER** :企业/企业员工自动签（他方企业自动签署或文件发起时的本方企业自动签）</li></ul>
	// 
	// 注:  
	// `1. 个人自动签场景(PERSON_AUTO_SIGN)为白名单功能, 使用前请联系对接的客户经理沟通。`
	// `2. 若要实现他方企业（同一应用下）自动签，需要满足3个条件：`
	// <ul><li>条件1：ApproverType 设置为ENTERPRISESERVER</li>
	// <li>条件2：子客之间完成授权</li>
	// <li>条件3：联系对接的客户经理沟通如何使用</li></ul>
	ApproverType *string `json:"ApproverType,omitnil,omitempty" name:"ApproverType"`

	// 签署流程签署人在模板中对应的签署人Id；在非单方签署、以及非B2C签署的场景下必传，用于指定当前签署方在签署流程中的位置；
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 签署人的签署截止时间，格式为Unix标准时间戳（秒）
	// 
	// 注: `若不设置此参数，则默认使用合同的截止时间，此参数暂不支持合同组子合同`
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 签署完回调url，最大长度1000个字符
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 使用PDF文件直接发起合同时，签署人指定的签署控件；<br/>使用模板发起合同时，指定本企业印章签署控件的印章ID:注意：(如果模板里面指定了印章，默认使用模板里面配置的印章，不能进行变更) <br/>通过ComponentId或ComponenetName指定签署控件，ComponentValue为印章ID。
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/91757a7f9188ccf3057a4a8979cf3f93.png)
	SignComponents []*Component `json:"SignComponents,omitnil,omitempty" name:"SignComponents"`

	// 当签署方控件类型为 <b>SIGN_SIGNATURE</b> 时，可以指定签署方签名方式。如果不指定，签署人可以使用所有的签名类型，可指定的签名类型包括：
	// 
	// <ul><li> <b>HANDWRITE</b> :需要实时手写的手写签名。</li>
	// <li> <b>HANDWRITTEN_ESIGN</b> :长效手写签名， 是使用保存到个人中心的印章列表的手写签名。(并且包含HANDWRITE)</li>
	// <li> <b>OCR_ESIGN</b> :AI智能识别手写签名。</li>
	// <li> <b>ESIGN</b> :个人印章类型。</li>
	// <li> <b>IMG_ESIGN</b>  : 图片印章。该类型支持用户在签署将上传的PNG格式的图片作为签名。</li>
	// <li> <b>SYSTEM_ESIGN</b> :系统签名。该类型可以在用户签署时根据用户姓名一键生成一个签名来进行签署。</li></ul>
	// 
	// 各种签名的样式可以参考下图：
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/ee0498856c060c065628a0c5ba780d6b.jpg)
	ComponentLimitType []*string `json:"ComponentLimitType,omitnil,omitempty" name:"ComponentLimitType"`

	// 签署方在签署合同之前，需要强制阅读合同的时长，可指定为3秒至300秒之间的任意值。
	// 
	// 若未指定阅读时间，则会按照合同页数大小计算阅读时间，计算规则如下：
	// <ul>
	// <li>合同页数少于等于2页，阅读时间为3秒；</li>
	// <li>合同页数为3到5页，阅读时间为5秒；</li>
	// <li>合同页数大于等于6页，阅读时间为10秒。</li>
	// </ul>
	PreReadTime *int64 `json:"PreReadTime,omitnil,omitempty" name:"PreReadTime"`

	// 签署完前端跳转的url，此字段的用法场景请联系客户经理确认
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 可以控制签署方在签署合同时能否进行某些操作，例如拒签、转交他人、是否为动态补充签署人等。
	// 详细操作可以参考开发者中心的ApproverOption结构体。
	ApproverOption *ApproverOption `json:"ApproverOption,omitnil,omitempty" name:"ApproverOption"`

	// 此签署人（员工或者个人）签署前，是否需要发起方企业进行审批，取值如下：
	// <ul><li>**false**：（默认）不需要审批，直接签署。</li>
	// <li>**true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待发起方企业内部审批完成。</li></ul>
	// 企业可以通过ChannelCreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li>如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li>如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 
	// 注：`此功能可用于与发起方企业内部的审批流程进行关联，支持手动、静默签署合同`
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/b14d5188ed0229d1401e74a9a49cab6d.png)
	ApproverNeedSignReview *bool `json:"ApproverNeedSignReview,omitnil,omitempty" name:"ApproverNeedSignReview"`

	// 指定个人签署方查看合同的校验方式,可以传值如下:
	// <ul><li>  **1**   : （默认）人脸识别,人脸识别后才能合同内容</li>
	// <li>  **2**  : 手机号验证, 用户手机号和参与方手机号(ApproverMobile)相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// </li></ul>
	// 注: 
	// <ul><li>如果合同流程设置ApproverVerifyType查看合同的校验方式,    则忽略此签署人的查看合同的校验方式</li>
	// <li>此字段可传多个校验方式</li></ul>
	ApproverVerifyTypes []*int64 `json:"ApproverVerifyTypes,omitnil,omitempty" name:"ApproverVerifyTypes"`

	// 签署人签署合同时的认证方式
	// <ul><li> **1** :人脸认证</li>
	// <li> **2** :签署密码</li>
	// <li> **3** :运营商三要素（如果是港澳台客户，建议不要选择这个）</li>
	// <li>**5**：设备指纹识别，需要对比手机机主预留的指纹信息，校验一致才能成功进行合同签署。（iOS系统暂不支持该校验方式）</li>
	// <li>**6**：设备面容识别，需要对比手机机主预留的人脸信息，校验一致才能成功进行合同签署。（Android系统暂不支持该校验方式）</li></ul>
	// 
	// 默认为：
	// 1(人脸认证 ),2(签署密码),3(运营商三要素),5(设备指纹识别),6(设备面容识别)
	// 
	// 注: 
	// 1. 用<font color='red'>模板创建合同场景</font>, 签署人的认证方式需要在配置模板的时候指定, <font color='red'>在创建合同重新指定无效</font>
	// 2. 运营商三要素认证方式对手机号运营商及前缀有限制,可以参考[运营商支持列表类](https://qian.tencent.com/developers/partner/mobile_support)得到具体的支持说明
	// 3. 校验方式不允许只包含<font color='red'>设备指纹识别</font>和<font color='red'>设备面容识别</font>，至少需要再增加一种其他校验方式。
	// 4. <font color='red'>设备指纹识别</font>和<font color='red'>设备面容识别</font>只支持小程序使用，其他端暂不支持。
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil,omitempty" name:"ApproverSignTypes"`

	// 签署ID
	// - 发起流程时系统自动补充
	// - 创建签署链接时，可以通过查询详情接口获得签署人的SignId，然后可传入此值为该签署人创建签署链接，无需再传姓名、手机号、证件号等其他信息
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 通知签署方经办人的方式, 有以下途径:
	// <ul><li> **SMS** :(默认)短信</li>
	// <li> **NONE** : 不通知</li></ul>
	// 
	// 注: `签署方为第三方子客企业时会被置为NONE,   不会发短信通知`
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// [通过文件创建签署流程](https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowByFiles)时,如果设置了外层参数SignBeanTag=1(允许签署过程中添加签署控件),则可通过此参数明确规定合同所使用的签署控件类型（骑缝章、普通章法人章等）和具体的印章（印章ID,或者印章类型）或签名方式。
	// 
	// 注：`限制印章控件或骑缝章控件情况下,仅本企业签署方可以指定具体印章（通过传递ComponentValue,支持多个），他方企业或个人只支持限制控件类型。`
	AddSignComponentsLimits []*ComponentLimit `json:"AddSignComponentsLimits,omitnil,omitempty" name:"AddSignComponentsLimits"`

	// 可以自定义签署人角色名：收款人、开具人、见证人等，长度不能超过20，只能由中文、字母、数字和下划线组成。
	// 
	// 注: `如果是用模板发起, 优先使用此处上传的, 如果不传则用模板的配置的`
	ApproverRoleName *string `json:"ApproverRoleName,omitnil,omitempty" name:"ApproverRoleName"`

	// 生成H5签署链接时，您可以指定签署方签署合同的认证校验方式的选择模式，可传递一下值：
	// <ul><li>**0**：签署方自行选择，签署方可以从预先指定的认证方式中自由选择；</li>
	// <li>**1**：自动按顺序首位推荐，签署方无需选择，系统会优先推荐使用第一种认证方式。</li></ul>
	// 注：
	// `不指定该值时，默认为签署方自行选择。`
	SignTypeSelector *uint64 `json:"SignTypeSelector,omitnil,omitempty" name:"SignTypeSelector"`

	// 签署人在合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li>单行文本控件</li>
	// <li>多行文本控件</li>
	// <li>勾选框控件</li>
	// <li>数字控件</li>
	// <li>图片控件</li>
	// <li>数据表格等填写控件</li></ul>
	// 
	// 具体使用说明可参考[为签署方指定填写控件](https://qian.tencent.cn/developers/partner/createFlowByFiles#为签署方指定填写控件)
	// 
	// 注：`此参数仅在通过文件发起合同或者合同组时生效`
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/e004195ee4cb98a7f9bc12eb4a0a0b77.png)
	Components []*Component `json:"Components,omitnil,omitempty" name:"Components"`

	// <b>只有在生成H5签署链接的情形下</b>（ 如调用<a href="https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateFlowSignUrl" target="_blank">获取H5签署链接</a>、<a href="https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateBatchQuickSignUrl" target="_blank">获取H5批量签署链接</a>等接口），该配置才会生效。
	// 
	// 您可以指定H5签署视频核身的意图配置，选择问答模式或点头模式的语音文本。
	// 
	// 注意：
	// 1. 视频认证为<b>白名单功能，使用前请联系对接的客户经理沟通</b>。
	// 2. 使用视频认证时，<b>生成H5签署链接的时候必须将签署认证方式指定为人脸</b>（即ApproverSignTypes设置成人脸签署）。
	// 3. 签署完成后，可以通过<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelDescribeSignFaceVideo" target="_blank">查询签署认证人脸视频</a>获取到当时的视频。
	Intention *Intention `json:"Intention,omitnil,omitempty" name:"Intention"`

	// 进入签署流程的限制，目前支持以下选项：
	// <ul><li> <b>空值（默认）</b> :无限制，可在任何场景进入签署流程。</li><li> <b>link</b> :选择此选项后，将无法通过控制台或电子签小程序列表进入填写或签署操作，仅可预览合同。填写或签署流程只能通过短信或发起方提供的专用链接进行。</li></ul>
	SignEndpoints []*string `json:"SignEndpoints,omitnil,omitempty" name:"SignEndpoints"`
}

type FlowApproverItem struct {
	// 合同编号
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署方信息，如角色ID、角色名称等
	Approvers []*ApproverItem `json:"Approvers,omitnil,omitempty" name:"Approvers"`
}

type FlowApproverUrlInfo struct {
	// 签署短链接。
	// 
	// 注意:
	// 1. 该链接有效期为**30分钟**，同时需要注意保密，不要外泄给无关用户。
	// 2. 该链接不支持小程序嵌入，仅支持**移动端浏览器**打开。
	// 3. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	SignUrl *string `json:"SignUrl,omitnil,omitempty" name:"SignUrl"`

	// 签署人类型。
	// - **PERSON**: 个人
	ApproverType *string `json:"ApproverType,omitnil,omitempty" name:"ApproverType"`

	// 签署人姓名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 签署人手机号。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 签署长链接。
	// 
	// 注意:
	// 1. 该链接有效期为**30分钟**，同时需要注意保密，不要外泄给无关用户。
	// 2. 该链接不支持小程序嵌入，仅支持**移动端浏览器**打开。
	// 3. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	LongUrl *string `json:"LongUrl,omitnil,omitempty" name:"LongUrl"`
}

type FlowBatchApproverInfo struct {
	// 合同流程ID。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署节点ID，用于生成动态签署人链接完成领取。注：`生成动态签署人补充链接时必传。`
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`
}

type FlowBatchUrlInfo struct {
	// 批量签署合同和签署方的信息，用于补充动态签署人。
	FlowBatchApproverInfos []*FlowBatchApproverInfo `json:"FlowBatchApproverInfos,omitnil,omitempty" name:"FlowBatchApproverInfos"`
}

type FlowDetailInfo struct {
	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 合同流程的类别分类（如销售合同/入职合同等）。
	// 该字段将被废弃，不建议使用。	请使用 UserFlowType
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 合同流程当前的签署状态, 会存在下列的状态值
	// <ul><li> **INIT** :合同创建</li>
	// <li> **PART** :合同签署中(至少有一个签署方已经签署)</li>
	// <li> **REJECT** :合同拒签</li>
	// <li> **ALL** :合同签署完成</li>
	// <li> **DEADLINE** :合同流签(合同过期)</li>
	// <li> **CANCEL** :合同撤回</li>
	// <li> **RELIEVED** :解除协议（已解除）</li></ul>
	//  
	FlowStatus *string `json:"FlowStatus,omitnil,omitempty" name:"FlowStatus"`

	// 当合同流程状态为已拒签（即 FlowStatus=REJECT）或已撤销（即 FlowStatus=CANCEL ）时，此字段 FlowMessage 为拒签或撤销原因。
	FlowMessage *string `json:"FlowMessage,omitnil,omitempty" name:"FlowMessage"`

	// 合同流程的创建时间戳，格式为Unix标准时间戳（秒）。
	CreateOn *int64 `json:"CreateOn,omitnil,omitempty" name:"CreateOn"`

	// 签署流程的签署截止时间, 值为unix时间戳, 精确到秒。
	DeadLine *int64 `json:"DeadLine,omitnil,omitempty" name:"DeadLine"`

	// 调用方自定义的个性化字段(可自定义此字段的值)，并以base64方式编码，支持的最大数据大小为 1000长度。
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	CustomData *string `json:"CustomData,omitnil,omitempty" name:"CustomData"`

	// 合同流程的签署方数组
	FlowApproverInfos []*FlowApproverDetail `json:"FlowApproverInfos,omitnil,omitempty" name:"FlowApproverInfos"`

	// 合同流程的关注方信息数组
	CcInfos []*FlowApproverDetail `json:"CcInfos,omitnil,omitempty" name:"CcInfos"`

	// 是否需要发起前审批
	// <ul><li>当NeedCreateReview为true，表明当前流程是需要发起前审核的合同，可能无法进行查看，签署操作，需要等审核完成后，才可以继续后续流程</li>
	// <li>当NeedCreateReview为false，不需要发起前审核的合同</li></ul>
	NeedCreateReview *bool `json:"NeedCreateReview,omitnil,omitempty" name:"NeedCreateReview"`

	// 用户合同的自定义分类。
	// 
	// 自定义合同类型的位置，在下图所示地方:
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/37138cc5f3c38e6f788f4b82f695cebf.png)
	UserFlowType *UserFlowType `json:"UserFlowType,omitnil,omitempty" name:"UserFlowType"`

	// 发起模板时,使用的模板Id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type FlowFileInfo struct {
	// 签署文件资源Id列表，目前仅支持单个文件
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 签署流程名称，长度不超过200个字符
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 签署流程签约方列表，最多不超过5个参与方
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitnil,omitempty" name:"FlowApprovers"`

	// 签署流程截止时间，十位数时间戳，最大值为33162419560，即3020年
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 签署流程的描述，长度不超过1000个字符
	FlowDescription *string `json:"FlowDescription,omitnil,omitempty" name:"FlowDescription"`

	// 签署流程的类型，长度不超过255个字符
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 已废弃，请使用【应用号配置】中的回调地址统一接收消息
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 第三方应用的业务信息，最大长度1000个字符。
	CustomerData *string `json:"CustomerData,omitnil,omitempty" name:"CustomerData"`

	// 合同签署顺序类型(无序签,顺序签)，默认为false，即有序签署
	Unordered *bool `json:"Unordered,omitnil,omitempty" name:"Unordered"`

	// 签署文件中的发起方的填写控件，需要在发起的时候进行填充
	Components []*Component `json:"Components,omitnil,omitempty" name:"Components"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitnil,omitempty" name:"CustomShowMap"`

	// 本企业(发起方企业)是否需要签署审批
	NeedSignReview *bool `json:"NeedSignReview,omitnil,omitempty" name:"NeedSignReview"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下：  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li><li> <b>3</b> :文书</li></ul>效果如下:![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"FlowDisplayType,omitnil,omitempty" name:"FlowDisplayType"`
}

type FlowForwardInfo struct {
	// 合同流程ID，为32位字符串。此接口的合同流程ID需要由[创建签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow)接口创建得到。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署方经办人在合同中的参与方ID，为32位字符串。
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`
}

type FlowForwardResult struct {
	// 合同流程ID为32位字符串。您可以登录腾讯电子签控制台，在 "合同" -> "合同中心" 中查看某个合同的FlowId（在页面中展示为合同ID）。[点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 如果失败，返回的错误细节。
	ErrorDetail *string `json:"ErrorDetail,omitnil,omitempty" name:"ErrorDetail"`
}

type FlowGroupApproverInfo struct {
	// 合同流程ID。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署节点ID，用于生成动态签署人链接完成领取。注：`生成动态签署人补充链接时必传。`
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`
}

type FlowGroupApprovers struct {
	// 合同流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署方信息，包含合同ID和角色ID用于定位RecipientId。
	Approvers []*ApproverItem `json:"Approvers,omitnil,omitempty" name:"Approvers"`
}

type FlowGroupOptions struct {
	// 发起方企业经办人（即签署人为发起方企业员工）是否需要对子合同进行独立的意愿确认
	// <ul><li>**false**（默认）：发起方企业经办人签署时对所有子合同进行统一的意愿确认。</li>
	// <li>**true**：发起方企业经办人签署时需要对子合同进行独立的意愿确认。</li></ul>
	SelfOrganizationApproverSignEach *bool `json:"SelfOrganizationApproverSignEach,omitnil,omitempty" name:"SelfOrganizationApproverSignEach"`

	// 非发起方企业经办人（即：签署人为个人或者不为发起方企业的员工）是否需要对子合同进行独立的意愿确认
	// <ul><li>**false**（默认）：非发起方企业经办人签署时对所有子合同进行统一的意愿确认。</li>
	// <li>**true**：非发起方企业经办人签署时需要对子合同进行独立的意愿确认。</li></ul>
	OtherApproverSignEach *bool `json:"OtherApproverSignEach,omitnil,omitempty" name:"OtherApproverSignEach"`
}

type FlowGroupUrlInfo struct {
	// 合同组子合同和签署方的信息，用于补充动态签署人。	
	FlowGroupApproverInfos []*FlowGroupApproverInfo `json:"FlowGroupApproverInfos,omitnil,omitempty" name:"FlowGroupApproverInfos"`
}

type FlowInfo struct {
	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	// 示例值：1604912664
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 用户配置的合同模板ID，会基于此模板创建合同文档，为32位字符串。
	// 如果使用模板发起接口，此参数为必填。
	// 
	// 可以通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>登录企业控制台, 在**企业模板**中得到合同模板ID。
	// 
	// [点击产看模板Id在控制台上的位置](https://qcloudimg.tencent-cloud.cn/raw/e988be12bf28a89b4716aed4502c2e02.png)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 合同流程的参与方列表，最多可支持50个参与方。对应不同签署人的传参方式可以参考文档 [签署方入参指引](https://qian.tencent.com/developers/partner/flow_approver)
	// 
	// 注:  
	// <font color="red" > <b> 在发起流程时，需要保证 FlowApprovers中的顺序与模板定义顺序一致，否则会发起失败。
	// 例如，如果模板中定义的第一个参与人是个人用户，第二个参与人是企业员工，则在 approver 中传参时，第一个也必须是个人用户，第二个参与人必须是企业员工。</b></font>
	// 
	// [点击查看模板参与人顺序定义位置](https://qcloudimg.tencent-cloud.cn/raw/c50e0a204fc5c66aaa2ca70e451ef2d6.png)
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitnil,omitempty" name:"FlowApprovers"`

	// 发起方角色的填写控件的填充内容。
	// 
	// 注：只有在控制台编辑模板时，<font color="red">归属给发起方</font>的填写控件（如下图）才能在创建文档的时候进行内容填充。(<font color="red">白名单功能需要联系对接经理开通，否则模板编辑时无法将填写控件分配给发起方</font>)。
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/b1d3978140ee2b44e2c9fdc96e467a5d.png)
	FormFields []*FormField `json:"FormFields,omitnil,omitempty" name:"FormFields"`

	// 该字段已废弃，请使用【应用号配置】中的回调地址统一接口消息
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil,omitempty" name:"FlowDescription"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为1000长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的回调通知模块。
	CustomerData *string `json:"CustomerData,omitnil,omitempty" name:"CustomerData"`

	// 您可以自定义腾讯电子签小程序合同列表页展示的合同内容模板，模板中支持以下变量：
	// <ul><li>{合同名称}   </li>
	// <li>{发起方企业} </li>
	// <li>{发起方姓名} </li>
	// <li>{签署方N企业}</li>
	// <li>{签署方N姓名}</li></ul>
	// 其中，N表示签署方的编号，从1开始，不能超过签署人的数量。
	// 
	// 例如，如果是腾讯公司张三发给李四名称为“租房合同”的合同，您可以将此字段设置为：`合同名称:{合同名称};发起方: {发起方企业}({发起方姓名});签署方:{签署方1姓名}`，则小程序中列表页展示此合同为以下样子
	// 
	// 合同名称：租房合同 
	// 发起方：腾讯公司(张三) 
	// 签署方：李四
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/628f0928cac15d2e3bfa6088f53f5998.png)
	CustomShowMap *string `json:"CustomShowMap,omitnil,omitempty" name:"CustomShowMap"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// <b>注</b>
	// 1. 抄送人名单中可以包括自然人以及本企业的员工（本企业员工必须已经完成认证并加入企业）。
	// 2. 请确保抄送人列表中的成员不与任何签署人重复。
	CcInfos []*CcInfo `json:"CcInfos,omitnil,omitempty" name:"CcInfos"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil,omitempty" name:"NeedSignReview"`

	// 若在创建签署流程时指定了关注人CcInfos，此参数可设定向关注人发送短信通知的类型：
	// <ul><li> **0** :合同发起时通知通知对方来查看合同（默认）</li>
	// <li> **1** : 签署完成后通知对方来查看合同</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil,omitempty" name:"CcNotifyType"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：电子处方单（医疗自动签）  </li><li> **OTHER** :  通用场景</li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil,omitempty" name:"AutoSignScene"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下：  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li><li> <b>3</b> :文书</li></ul>效果如下:![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"FlowDisplayType,omitnil,omitempty" name:"FlowDisplayType"`
}

type FlowResourceUrlInfo struct {
	// 合同流程的ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 对应的合同流程的PDF下载链接
	ResourceUrlInfos []*ResourceUrlInfo `json:"ResourceUrlInfos,omitnil,omitempty" name:"ResourceUrlInfos"`
}

type FormField struct {
	// 控件填充值，ComponentType和传入值格式对应关系如下：
	// <ul><li> <b>TEXT</b> : 文本内容</li>
	// <li> <b>MULTI_LINE_TEXT</b> : 文本内容， 可以用  \n 来控制换行位置</li>
	// <li> <b>CHECK_BOX</b> : true/false</li>
	// <li> <b>FILL_IMAGE、ATTACHMENT</b> : 附件的FileId，需要通过UploadFiles接口上传获取</li>
	// <li> <b>SELECTOR</b> : 选项值</li>
	// <li> <b>DYNAMIC_TABLE</b>  - 传入json格式的表格内容，详见说明：[数据表格](https://qian.tencent.com/developers/partner/dynamic_table)</li>
	// <li> <b>DATE</b> : 格式化：xxxx年xx月xx日（例如：2024年05月28日）</li>
	// <li> <b>DISTRICT </b> : 省市区行政区控件，需填写ComponentValue为省市区行政区字符串内容</li>
	// </ul>
	// 
	// 
	// <b>控件值约束说明</b>：
	// <table> <thead> <tr> <th>特殊控件</th> <th>填写约束</th> </tr> </thead> <tbody> <tr> <td>企业全称控件</td> <td>企业名称中文字符中文括号</td> </tr> <tr> <td>统一社会信用代码控件</td> <td>企业注册的统一社会信用代码</td> </tr> <tr> <td>法人名称控件</td> <td>最大50个字符，2到25个汉字或者1到50个字母</td> </tr> <tr> <td>签署意见控件</td> <td>签署意见最大长度为50字符</td> </tr> <tr> <td>签署人手机号控件</td> <td>国内手机号 13,14,15,16,17,18,19号段长度11位</td> </tr> <tr> <td>签署人身份证控件</td> <td>合法的身份证号码检查</td> </tr> <tr> <td>控件名称</td> <td>控件名称最大长度为20字符，不支持表情</td> </tr> <tr> <td>单行文本控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>多行文本控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>勾选框控件</td> <td>选择填字符串true，不选填字符串false</td> </tr> <tr> <td>选择器控件</td> <td>同单行文本控件约束，填写选择值中的字符串</td> </tr> <tr> <td>数字控件</td> <td>请输入有效的数字(可带小数点)</td> </tr> <tr> <td>日期控件</td> <td>格式：yyyy年mm月dd日</td> </tr> <tr> <td>附件控件</td> <td>JPG或PNG图片，上传数量限制，1到6个，最大6个附件，填写上传的资源ID</td> </tr> <tr> <td>图片控件</td> <td>JPG或PNG图片，填写上传的图片资源ID</td> </tr> <tr> <td>邮箱控件</td> <td>有效的邮箱地址, w3c标准</td> </tr> <tr> <td>地址控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>省市区控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>性别控件</td> <td>选择值中的字符串</td> </tr> <tr> <td>学历控件</td> <td>选择值中的字符串</td> </tr> </tbody> </table>
	ComponentValue *string `json:"ComponentValue,omitnil,omitempty" name:"ComponentValue"`

	// 表单域或控件的ID，跟ComponentName二选一，不能全为空；
	// CreateFlowsByTemplates 接口不使用此字段。
	// 
	// <a href="https://dyn.ess.tencent.cn/guide/apivideo/channel_component_name.mp4" target="_blank">点击此处查看模板上控件ID的获取方式</a>
	ComponentId *string `json:"ComponentId,omitnil,omitempty" name:"ComponentId"`

	// 控件的名字，跟ComponentId二选一，不能全为空
	// 
	// <a href="https://dyn.ess.tencent.cn/guide/apivideo/channel_component_name.mp4" target="_blank">点击此处查看模板上控件名字的获取方式</a>
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// 是否锁定模板控件值，锁定后无法修改（用于嵌入式发起合同），true-锁定，false-不锁定
	LockComponentValue *bool `json:"LockComponentValue,omitnil,omitempty" name:"LockComponentValue"`
}

// Predefined struct for user
type GetDownloadFlowUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 程合同ID列表,  可将这些流程ID组织成合同组的形式, 下载时候每个文件夹会是一个zip压缩包,  每个文件夹最多20个合同, 所有文件夹最多50个合同
	// 如下列组织形式,  控制台下载页面点击下载按钮后, 会生成**2023采购合同.zip**和**2023入职合同.zip** 两个下载任务(注:`部分浏览器需要授权或不支持创建多下载任务`)
	// 
	// **2023采购合同.zip**压缩包会有`yDwivUUckpor6wtoUuogwQHCAB0ES0pQ`和`yDwi8UUckpo5fz9cUqI6nGwcuTvt9YSh`两个合同的文件
	// **2023入职合同.zip** 压缩包会有`yDwivUUckpor6wobUuogwQHvdGfvDi5K`的文件
	// 
	// ![image](	https://dyn.ess.tencent.cn/guide/capi/channel_GetDownloadFlowUrl_DownLoadFlows.png)
	DownLoadFlows []*DownloadFlowInfo `json:"DownLoadFlows,omitnil,omitempty" name:"DownLoadFlows"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type GetDownloadFlowUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 程合同ID列表,  可将这些流程ID组织成合同组的形式, 下载时候每个文件夹会是一个zip压缩包,  每个文件夹最多20个合同, 所有文件夹最多50个合同
	// 如下列组织形式,  控制台下载页面点击下载按钮后, 会生成**2023采购合同.zip**和**2023入职合同.zip** 两个下载任务(注:`部分浏览器需要授权或不支持创建多下载任务`)
	// 
	// **2023采购合同.zip**压缩包会有`yDwivUUckpor6wtoUuogwQHCAB0ES0pQ`和`yDwi8UUckpo5fz9cUqI6nGwcuTvt9YSh`两个合同的文件
	// **2023入职合同.zip** 压缩包会有`yDwivUUckpor6wobUuogwQHvdGfvDi5K`的文件
	// 
	// ![image](	https://dyn.ess.tencent.cn/guide/capi/channel_GetDownloadFlowUrl_DownLoadFlows.png)
	DownLoadFlows []*DownloadFlowInfo `json:"DownLoadFlows,omitnil,omitempty" name:"DownLoadFlows"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	// 跳转控制台合同下载页面链接 , 5分钟之内有效，且只能访问一次
	DownLoadUrl *string `json:"DownLoadUrl,omitnil,omitempty" name:"DownLoadUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type HasAuthOrganization struct {
	// 授权企业openid，
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 授权企业名称	
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 被授权企业openid，
	AuthorizedOrganizationOpenId *string `json:"AuthorizedOrganizationOpenId,omitnil,omitempty" name:"AuthorizedOrganizationOpenId"`

	// 被授权企业名称	
	AuthorizedOrganizationName *string `json:"AuthorizedOrganizationName,omitnil,omitempty" name:"AuthorizedOrganizationName"`

	// 授权时间，格式为时间戳，单位s	
	AuthorizeTime *int64 `json:"AuthorizeTime,omitnil,omitempty" name:"AuthorizeTime"`
}

type HasAuthUser struct {
	// 第三方应用平台自定义，对应第三方平台子客企业员工的唯一标识。
	// 
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`
}

type Intention struct {
	// 视频认证类型，支持以下类型
	// <ul><li>1 : 问答模式</li>
	// <li>2 : 点头模式</li></ul>
	// 
	// 注: `视频认证为白名单功能，使用前请联系对接的客户经理沟通。`
	IntentionType *int64 `json:"IntentionType,omitnil,omitempty" name:"IntentionType"`

	// 意愿核身语音问答模式（即语音播报+语音回答）使用的文案，包括：系统语音播报的文本、需要核验的标准文本。支持传入1～10轮问答，最多支持10轮。
	// 
	// 注：`选择问答模式时，此字段可不传，不传则使用默认语音文本：请问，您是否同意签署本协议？可语音回复“同意”或“不同意”。`
	IntentionQuestions []*IntentionQuestion `json:"IntentionQuestions,omitnil,omitempty" name:"IntentionQuestions"`

	// 意愿核身（点头确认模式）使用的文案，若未使用意愿核身（点头确认模式），则该字段无需传入。支持传入1～10轮点头确认文本，最多支持10轮。
	// 
	// 注：`选择点头模式时，此字段可不传，不传则使用默认语音文本：请问，您是否同意签署本协议？可点头同意。`
	IntentionActions []*IntentionAction `json:"IntentionActions,omitnil,omitempty" name:"IntentionActions"`
}

type IntentionAction struct {
	// 点头确认模式下，系统语音播报使用的问题文本，问题最大长度为150个字符。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type IntentionActionResult struct {
	// 意愿核身结果详细数据，与每段点头确认过程一一对应
	Details []*IntentionActionResultDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type IntentionActionResultDetail struct {
	// 视频base64编码（其中包含全程提示文本和点头音频，mp4格式）
	Video *string `json:"Video,omitnil,omitempty" name:"Video"`
}

type IntentionQuestion struct {
	// 当选择语音问答模式时，系统自动播报的问题文本，最大长度为150个字符。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	//  当选择语音问答模式时，用于判断用户回答是否通过的标准答案列表，传入后可自动判断用户回答文本是否在标准文本列表中。
	Answers []*string `json:"Answers,omitnil,omitempty" name:"Answers"`
}

type IntentionQuestionResult struct {
	// 视频base64（其中包含全程问题和回答音频，mp4格式）
	// 
	// 注：`需进行base64解码获取视频文件`
	Video *string `json:"Video,omitnil,omitempty" name:"Video"`

	//  和答案匹配结果列表
	ResultCode []*string `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`

	// 回答问题语音识别结果列表
	AsrResult []*string `json:"AsrResult,omitnil,omitempty" name:"AsrResult"`
}

// Predefined struct for user
type ModifyExtendedServiceRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	//   扩展服务类型
	// <ul>
	// <li>AUTO_SIGN             企业自动签（自动签署）</li>
	// <li>  OVERSEA_SIGN          企业与港澳台居民签署合同</li>
	// <li>  MOBILE_CHECK_APPROVER 使用手机号验证签署方身份</li>
	// <li> DOWNLOAD_FLOW         授权渠道下载合同 </li>
	// <li>AGE_LIMIT_EXPANSION 拓宽签署方年龄限制</li>
	// <li>HIDE_OPERATOR_DISPLAY 隐藏合同经办人姓名</li>
	// </ul>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 操作类型
	// <ul>
	// <li>OPEN : 开通</li>
	// <li>CLOSE : 关闭</li>
	// </ul>
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 链接跳转类型，支持以下类型
	// <ul>
	// <li>WEIXINAPP : 短链直接跳转到电子签小程序  (默认值)</li>
	// <li>APP : 第三方APP或小程序跳转电子签小程序</li>
	// <li>WEIXIN_QRCODE_URL：直接跳转至电子签小程序的二维码链接，无需通过中转页。<font color="red">您需要自行将其转换为二维码，使用微信扫码后可直接进入。请注意，直接点击链接是无效的。</font></li>
	// </ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

type ModifyExtendedServiceRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	//   扩展服务类型
	// <ul>
	// <li>AUTO_SIGN             企业自动签（自动签署）</li>
	// <li>  OVERSEA_SIGN          企业与港澳台居民签署合同</li>
	// <li>  MOBILE_CHECK_APPROVER 使用手机号验证签署方身份</li>
	// <li> DOWNLOAD_FLOW         授权渠道下载合同 </li>
	// <li>AGE_LIMIT_EXPANSION 拓宽签署方年龄限制</li>
	// <li>HIDE_OPERATOR_DISPLAY 隐藏合同经办人姓名</li>
	// </ul>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 操作类型
	// <ul>
	// <li>OPEN : 开通</li>
	// <li>CLOSE : 关闭</li>
	// </ul>
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 链接跳转类型，支持以下类型
	// <ul>
	// <li>WEIXINAPP : 短链直接跳转到电子签小程序  (默认值)</li>
	// <li>APP : 第三方APP或小程序跳转电子签小程序</li>
	// <li>WEIXIN_QRCODE_URL：直接跳转至电子签小程序的二维码链接，无需通过中转页。<font color="red">您需要自行将其转换为二维码，使用微信扫码后可直接进入。请注意，直接点击链接是无效的。</font></li>
	// </ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

func (r *ModifyExtendedServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyExtendedServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "ServiceType")
	delete(f, "Operate")
	delete(f, "Endpoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyExtendedServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyExtendedServiceResponseParams struct {
	// 操作跳转链接
	// <ul><li><strong>链接有效期：</strong> 跳转链接的有效期为24小时。</li>
	// <li><strong>没有返回链接的情形：</strong> 如果在操作时没有返回跳转链接，说明此次操作无需进行跳转，服务将会直接被开通或关闭。</li>
	// <li><strong>返回链接的情形：</strong> 当操作类型为<b>OPEN（开通服务）</b>，并且扩展服务类型为<b>AUTO_SIGN（ 企业自动签署）</b>、<b>DOWNLOAD_FLOW（授权渠道下载合同）</b>时，系统将返回一个操作链接。收到操作链接后，贵方需主动联系超级管理员（超管）或法人。<font color="red"><b>由超管或法人点击链接</b></font>，以完成服务的开通操作。</li>
	// </ul>
	// 
	OperateUrl *string `json:"OperateUrl,omitnil,omitempty" name:"OperateUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyExtendedServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyExtendedServiceResponseParams `json:"Response"`
}

func (r *ModifyExtendedServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyExtendedServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFlowDeadlineRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID，为32位字符串。
	// <ul><li>建议开发者妥善保存此流程ID，以便于顺利进行后续操作。</li>
	// <li>可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。</li></ul>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署流程或签署人新的签署截止时间，格式为Unix标准时间戳（秒）
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 签署方角色编号，为32位字符串
	// <ul><li>若指定了此参数，则只调整签署流程中此签署人的签署截止时间，否则调整合同整体的签署截止时间（合同截止时间+发起时未设置签署人截止时间的参与人的签署截止时间）</li>
	// <li>通过[用PDF文件创建签署流程](https://test.qian.tencent.cn/developers/partnerApis/startFlows/ChannelCreateFlowByFiles)发起合同，或通过[用模板创建签署流程](https://test.qian.tencent.cn/developers/partnerApis/startFlows/CreateFlowsByTemplates)时，返回参数[FlowApprovers](https://test.qian.tencent.cn/developers/partnerApis/dataTypes/#approveritem)会返回此信息，建议开发者妥善保存</li>
	// <li>也可通过[获取合同信息](https://test.qian.tencent.cn/developers/partnerApis/flows/DescribeFlowDetailInfo)接口查询签署人的RecipientId编号</li></ul>
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`
}

type ModifyFlowDeadlineRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程ID，为32位字符串。
	// <ul><li>建议开发者妥善保存此流程ID，以便于顺利进行后续操作。</li>
	// <li>可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。</li></ul>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署流程或签署人新的签署截止时间，格式为Unix标准时间戳（秒）
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 签署方角色编号，为32位字符串
	// <ul><li>若指定了此参数，则只调整签署流程中此签署人的签署截止时间，否则调整合同整体的签署截止时间（合同截止时间+发起时未设置签署人截止时间的参与人的签署截止时间）</li>
	// <li>通过[用PDF文件创建签署流程](https://test.qian.tencent.cn/developers/partnerApis/startFlows/ChannelCreateFlowByFiles)发起合同，或通过[用模板创建签署流程](https://test.qian.tencent.cn/developers/partnerApis/startFlows/CreateFlowsByTemplates)时，返回参数[FlowApprovers](https://test.qian.tencent.cn/developers/partnerApis/dataTypes/#approveritem)会返回此信息，建议开发者妥善保存</li>
	// <li>也可通过[获取合同信息](https://test.qian.tencent.cn/developers/partnerApis/flows/DescribeFlowDetailInfo)接口查询签署人的RecipientId编号</li></ul>
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`
}

func (r *ModifyFlowDeadlineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFlowDeadlineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "FlowId")
	delete(f, "Deadline")
	delete(f, "RecipientId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFlowDeadlineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFlowDeadlineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFlowDeadlineResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFlowDeadlineResponseParams `json:"Response"`
}

func (r *ModifyFlowDeadlineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFlowDeadlineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPartnerAutoSignAuthUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 被授企业id/授权方企业id（即OrganizationId），如果是企业之间授权和AuthorizedOrganizationName二选一传入。
	// 
	// 注：`被授权企业必须和当前企业在同一应用号下`
	AuthorizedOrganizationId *string `json:"AuthorizedOrganizationId,omitnil,omitempty" name:"AuthorizedOrganizationId"`

	// 被授企业名称/授权方企业的名字，如果是企业之间授权和AuthorizedOrganizationId二选一传入即可。请确认该名称与企业营业执照中注册的名称一致。
	// 
	// 注: 
	// 1. 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 2. 被授权企业必须和当前企业在同一应用号下
	AuthorizedOrganizationName *string `json:"AuthorizedOrganizationName,omitnil,omitempty" name:"AuthorizedOrganizationName"`

	// 是否给平台应用授权
	// 
	// <ul>
	// <li><strong>true</strong>: 表示是，授权平台应用。在此情况下，无需设置<code>AuthorizedOrganizationId</code>和<code>AuthorizedOrganizationName</code>。</li>
	// <li><strong>false</strong>: （默认）表示否，不是授权平台应用。</li>
	// </ul>
	// 
	//  注：授权给平台应用需要开通【基于子客授权第三方应用可文件发起子客自动签署】白名单，请联系运营经理开通。
	PlatformAppAuthorization *bool `json:"PlatformAppAuthorization,omitnil,omitempty" name:"PlatformAppAuthorization"`

	// 在处理授权关系时，授权的方向
	// <ul>
	// <li><strong>false</strong>（默认值）：表示我方授权他方。在这种情况下，<code>AuthorizedOrganizationName</code> 代表的是【被授权方】的企业名称，即接收授权的企业。</li>
	// <li><strong>true</strong>：表示他方授权我方。在这种情况下，<code>AuthorizedOrganizationName</code> 代表的是【授权方】的企业名称，即提供授权的企业。</li>
	// </ul>
	AuthToMe *bool `json:"AuthToMe,omitnil,omitempty" name:"AuthToMe"`
}

type ModifyPartnerAutoSignAuthUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 被授企业id/授权方企业id（即OrganizationId），如果是企业之间授权和AuthorizedOrganizationName二选一传入。
	// 
	// 注：`被授权企业必须和当前企业在同一应用号下`
	AuthorizedOrganizationId *string `json:"AuthorizedOrganizationId,omitnil,omitempty" name:"AuthorizedOrganizationId"`

	// 被授企业名称/授权方企业的名字，如果是企业之间授权和AuthorizedOrganizationId二选一传入即可。请确认该名称与企业营业执照中注册的名称一致。
	// 
	// 注: 
	// 1. 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 2. 被授权企业必须和当前企业在同一应用号下
	AuthorizedOrganizationName *string `json:"AuthorizedOrganizationName,omitnil,omitempty" name:"AuthorizedOrganizationName"`

	// 是否给平台应用授权
	// 
	// <ul>
	// <li><strong>true</strong>: 表示是，授权平台应用。在此情况下，无需设置<code>AuthorizedOrganizationId</code>和<code>AuthorizedOrganizationName</code>。</li>
	// <li><strong>false</strong>: （默认）表示否，不是授权平台应用。</li>
	// </ul>
	// 
	//  注：授权给平台应用需要开通【基于子客授权第三方应用可文件发起子客自动签署】白名单，请联系运营经理开通。
	PlatformAppAuthorization *bool `json:"PlatformAppAuthorization,omitnil,omitempty" name:"PlatformAppAuthorization"`

	// 在处理授权关系时，授权的方向
	// <ul>
	// <li><strong>false</strong>（默认值）：表示我方授权他方。在这种情况下，<code>AuthorizedOrganizationName</code> 代表的是【被授权方】的企业名称，即接收授权的企业。</li>
	// <li><strong>true</strong>：表示他方授权我方。在这种情况下，<code>AuthorizedOrganizationName</code> 代表的是【授权方】的企业名称，即提供授权的企业。</li>
	// </ul>
	AuthToMe *bool `json:"AuthToMe,omitnil,omitempty" name:"AuthToMe"`
}

func (r *ModifyPartnerAutoSignAuthUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPartnerAutoSignAuthUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "AuthorizedOrganizationId")
	delete(f, "AuthorizedOrganizationName")
	delete(f, "PlatformAppAuthorization")
	delete(f, "AuthToMe")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPartnerAutoSignAuthUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPartnerAutoSignAuthUrlResponseParams struct {
	// 授权链接，以短链形式返回，短链的有效期参考回参中的 ExpiredTime。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 从客户小程序或者客户APP跳转至腾讯电子签小程序进行批量签署的跳转路径
	MiniAppPath *string `json:"MiniAppPath,omitnil,omitempty" name:"MiniAppPath"`

	// 链接过期时间以 Unix 时间戳格式表示，从生成链接时间起，往后7天有效期。过期后短链将失效，无法打开。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPartnerAutoSignAuthUrlResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPartnerAutoSignAuthUrlResponseParams `json:"Response"`
}

func (r *ModifyPartnerAutoSignAuthUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPartnerAutoSignAuthUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NeedReviewApproverInfo struct {
	// 签署方经办人的类型，支持以下类型
	// <ul><li> ORGANIZATION 企业（含企业自动签）</li>
	// <li>PERSON 个人（含个人自动签）</li></ul>
	ApproverType *string `json:"ApproverType,omitnil,omitempty" name:"ApproverType"`

	// 签署方经办人的姓名。 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	ApproverName *string `json:"ApproverName,omitnil,omitempty" name:"ApproverName"`

	// 签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。 请确认手机号所有方为此合同签署方。
	ApproverMobile *string `json:"ApproverMobile,omitnil,omitempty" name:"ApproverMobile"`

	// 签署方经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD 中国大陆居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN 中国港澳台居民居住证(格式同中国大陆居民身份证)</li>
	// <li>OTHER_CARD_TYPE 其他证件</li></ul>
	// 
	// 注: `其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	ApproverIdCardType *string `json:"ApproverIdCardType,omitnil,omitempty" name:"ApproverIdCardType"`

	// 签署方经办人的证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitnil,omitempty" name:"ApproverIdCardNumber"`

	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 如果签署方是企业签署方(approverType = 0 或者 approverType = 3)， 则企业名称必填。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`
}

type OccupiedSeal struct {
	// 电子印章编号
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 电子印章名称
	SealName *string `json:"SealName,omitnil,omitempty" name:"SealName"`

	// 电子印章授权时间戳，单位秒
	CreateOn *int64 `json:"CreateOn,omitnil,omitempty" name:"CreateOn"`

	// 电子印章授权人，电子签的UserId
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 电子印章策略Id
	SealPolicyId *string `json:"SealPolicyId,omitnil,omitempty" name:"SealPolicyId"`

	// 印章状态，有以下六种：CHECKING（审核中）SUCCESS（已启用）FAIL（审核拒绝）CHECKING-SADM（待超管审核）DISABLE（已停用）STOPPED（已终止）
	SealStatus *string `json:"SealStatus,omitnil,omitempty" name:"SealStatus"`

	// 审核失败原因
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// 印章图片url，5分钟内有效
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 电子印章类型 , 可选类型如下: 
	// <ul><li>**OFFICIAL**: (默认)公章</li>
	// <li>**CONTRACT**: 合同专用章;</li>
	// <li>**FINANCE**: 财务专用章;</li>
	// <li>**PERSONNEL**: 人事专用章</li>
	// <li>**INVOICE**: 发票专用章</li>
	// </ul>
	SealType *string `json:"SealType,omitnil,omitempty" name:"SealType"`

	// 用印申请是否为永久授权
	IsAllTime *bool `json:"IsAllTime,omitnil,omitempty" name:"IsAllTime"`

	// 授权人列表
	AuthorizedUsers []*AuthorizedUser `json:"AuthorizedUsers,omitnil,omitempty" name:"AuthorizedUsers"`
}

// Predefined struct for user
type OperateChannelTemplateRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识: Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 操作类型，可取值如下:
	// <ul>
	// <li>SELECT:  查询</li>
	// <li>DELETE:  删除</li>
	// <li>UPDATE: 更新</li>
	// </ul>
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 合同模板ID，为32位字符串。
	// 注: ` 此处为第三方应用平台模板库模板ID，非子客模板ID`
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 第三方平台子客企业的唯一标识，支持批量(用,分割)，
	ProxyOrganizationOpenIds *string `json:"ProxyOrganizationOpenIds,omitnil,omitempty" name:"ProxyOrganizationOpenIds"`

	// 模板可见范围, 可以设置的值如下:
	// 
	// **all**: 所有本第三方应用合作企业可见
	// **part**: 指定的本第三方应用合作企业
	// 
	// 对应控制台的位置
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/68b97812c68d6af77a5991e3bff5c790.png)
	AuthTag *string `json:"AuthTag,omitnil,omitempty" name:"AuthTag"`

	// 当OperateType=UPDATE时，可以通过设置此字段对模板启停用状态进行操作。
	// <ul>
	// <li>0: 不修改模板可用状态</li>
	// <li>1:  启用模板</li>
	// <li>2: 停用模板</li>
	// </ul>
	// 启用后模板可以正常领取。
	// 
	// 停用后，推送方式为【自动推送】的模板则无法被子客使用，推送方式为【手动领取】的模板则无法出现被模板库被子客领用。
	// 如果Available更新失败，会直接返回错误。
	Available *int64 `json:"Available,omitnil,omitempty" name:"Available"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type OperateChannelTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识: Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 操作类型，可取值如下:
	// <ul>
	// <li>SELECT:  查询</li>
	// <li>DELETE:  删除</li>
	// <li>UPDATE: 更新</li>
	// </ul>
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 合同模板ID，为32位字符串。
	// 注: ` 此处为第三方应用平台模板库模板ID，非子客模板ID`
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 第三方平台子客企业的唯一标识，支持批量(用,分割)，
	ProxyOrganizationOpenIds *string `json:"ProxyOrganizationOpenIds,omitnil,omitempty" name:"ProxyOrganizationOpenIds"`

	// 模板可见范围, 可以设置的值如下:
	// 
	// **all**: 所有本第三方应用合作企业可见
	// **part**: 指定的本第三方应用合作企业
	// 
	// 对应控制台的位置
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/68b97812c68d6af77a5991e3bff5c790.png)
	AuthTag *string `json:"AuthTag,omitnil,omitempty" name:"AuthTag"`

	// 当OperateType=UPDATE时，可以通过设置此字段对模板启停用状态进行操作。
	// <ul>
	// <li>0: 不修改模板可用状态</li>
	// <li>1:  启用模板</li>
	// <li>2: 停用模板</li>
	// </ul>
	// 启用后模板可以正常领取。
	// 
	// 停用后，推送方式为【自动推送】的模板则无法被子客使用，推送方式为【手动领取】的模板则无法出现被模板库被子客领用。
	// 如果Available更新失败，会直接返回错误。
	Available *int64 `json:"Available,omitnil,omitempty" name:"Available"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	delete(f, "Available")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OperateChannelTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OperateChannelTemplateResponseParams struct {
	// 第三方应用平台的应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 合同模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 描述模板可见性更改的结果。
	// <ul>
	// <li>all-success: 全部成功</li>
	// <li>part-success: 部分成功,失败的会在FailMessageList中展示</li>
	// <li>fail:全部失败, 失败的会在FailMessageList中展示</li>
	// </ul>
	OperateResult *string `json:"OperateResult,omitnil,omitempty" name:"OperateResult"`

	// 模板可见范围:
	// **all**: 所有本第三方应用合作企业可见
	// **part**: 指定的本第三方应用合作企业
	AuthTag *string `json:"AuthTag,omitnil,omitempty" name:"AuthTag"`

	// 第三方平台子客企业标识列表
	ProxyOrganizationOpenIds []*string `json:"ProxyOrganizationOpenIds,omitnil,omitempty" name:"ProxyOrganizationOpenIds"`

	// 操作失败信息数组
	FailMessageList []*AuthFailMessage `json:"FailMessageList,omitnil,omitempty" name:"FailMessageList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type OperateTemplateRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId（模板的归属企业）</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId （操作人）</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 模板ID，为32位字符串。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 操作类型，可取值如下:
	// <ul>
	// <li>DELETE:  删除</li>
	// <li>ENABLE: 启用</li>
	// <li>DISABLE: 停用</li>
	// <li>COPY: 复制新建</li>
	// </ul>
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 模板名称，长度不超过64字符。<br>
	// 模板复制时指定有效，若为空，则复制后模板名称为 **原模板名称_副本**。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`
}

type OperateTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId（模板的归属企业）</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId （操作人）</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 模板ID，为32位字符串。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 操作类型，可取值如下:
	// <ul>
	// <li>DELETE:  删除</li>
	// <li>ENABLE: 启用</li>
	// <li>DISABLE: 停用</li>
	// <li>COPY: 复制新建</li>
	// </ul>
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 模板名称，长度不超过64字符。<br>
	// 模板复制时指定有效，若为空，则复制后模板名称为 **原模板名称_副本**。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`
}

func (r *OperateTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OperateTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agent")
	delete(f, "TemplateId")
	delete(f, "OperateType")
	delete(f, "TemplateName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OperateTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OperateTemplateResponseParams struct {
	// 模板ID，为32位字符串，模板复制新建时返回
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称，模板复制新建时返回
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OperateTemplateResponse struct {
	*tchttp.BaseResponse
	Response *OperateTemplateResponseParams `json:"Response"`
}

func (r *OperateTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OperateTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrganizationAuthUrl struct {
	// 跳转链接, 链接的有效期根据企业,员工状态和终端等有区别, 可以参考下表
	// <table> <thead> <tr> <th>子客企业状态</th> <th>子客企业员工状态</th> <th>Endpoint</th> <th>链接有效期限</th> </tr> </thead>  <tbody> <tr> <td>企业未激活</td> <td>员工未认证</td> <td>PC</td> <td>5分钟</td>  </tr>  <tr> <td>企业未激活</td> <td>员工未认证</td> <td>CHANNEL/SHORT_URL/APP</td> <td>一年</td>  </tr>  <tr> <td>企业已激活</td> <td>员工未认证</td> <td>PC</td> <td>5分钟</td>  </tr> <tr> <td>企业已激活</td> <td>员工未认证</td> <td>CHANNEL/SHORT_URL/APP</td> <td>一年</td>  </tr>  <tr> <td>企业已激活</td> <td>员工已认证</td> <td>PC</td> <td>5分钟</td>  </tr>  <tr> <td>企业已激活</td> <td>员工已认证</td> <td>CHANNEL/SHORT_URL/APP</td> <td>一年</td>  </tr> </tbody> </table>
	// 注： 
	// `1.链接仅单次有效，每次登录需要需要重新创建新的链接`
	// `2.创建的链接应避免被转义，如：&被转义为\u0026；如使用Postman请求后，请选择响应类型为 JSON，否则链接将被转义`
	AuthUrl *string `json:"AuthUrl,omitnil,omitempty" name:"AuthUrl"`

	// 企业批量注册的错误信息，例如：企业三要素不通过
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

type OrganizationAuthorizationOptions struct {
	// 对方打开链接认证时，对方填写的营业执照的社会信用代码是否与接口上传上来的要保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>
	UniformSocialCreditCodeSame *bool `json:"UniformSocialCreditCodeSame,omitnil,omitempty" name:"UniformSocialCreditCodeSame"`

	// 对方打开链接认证时，企业名称是否要与接口传递上来的保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>p.s. 仅在企业名称不为空时有效
	OrganizationNameSame *bool `json:"OrganizationNameSame,omitnil,omitempty" name:"OrganizationNameSame"`

	// 对方打开链接认证时，法人姓名是否要与接口传递上来的保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>p.s. 仅在法人姓名不为空时有效
	LegalNameSame *bool `json:"LegalNameSame,omitnil,omitempty" name:"LegalNameSame"`

	// 对方打开链接认证时，对公打款账号是否要与接口传递上来的保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>p.s. 仅在对公打款账号不为空时有效
	BankAccountNumberSame *bool `json:"BankAccountNumberSame,omitnil,omitempty" name:"BankAccountNumberSame"`
}

type OrganizationCommonInfo struct {
	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 组织机构企业统一社会信用代码。
	// 请确认该企业统一社会信用代码与企业营业执照中注册的统一社会信用代码一致。
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// 组织机构法人的姓名。
	// 请确认该企业统一社会信用代码与企业营业执照中注册的法人姓名一致。
	LegalName *string `json:"LegalName,omitnil,omitempty" name:"LegalName"`

	// 组织机构法人的证件类型
	LegalIdCardType *string `json:"LegalIdCardType,omitnil,omitempty" name:"LegalIdCardType"`

	// 组织机构法人的证件号码
	LegalIdCardNumber *string `json:"LegalIdCardNumber,omitnil,omitempty" name:"LegalIdCardNumber"`

	// 组织机构超管姓名。
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 组织机构超管手机号。
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`

	// 组织机构超管证件类型
	AdminIdCardType *string `json:"AdminIdCardType,omitnil,omitempty" name:"AdminIdCardType"`

	// 组织机构超管证件号码
	AdminIdCardNumber *string `json:"AdminIdCardNumber,omitnil,omitempty" name:"AdminIdCardNumber"`

	// 原超管姓名
	OldAdminName *string `json:"OldAdminName,omitnil,omitempty" name:"OldAdminName"`

	// 原超管手机号
	OldAdminMobile *string `json:"OldAdminMobile,omitnil,omitempty" name:"OldAdminMobile"`

	// 原超管证件类型
	OldAdminIdCardType *string `json:"OldAdminIdCardType,omitnil,omitempty" name:"OldAdminIdCardType"`

	// 原超管证件号码
	OldAdminIdCardNumber *string `json:"OldAdminIdCardNumber,omitnil,omitempty" name:"OldAdminIdCardNumber"`
}

type OrganizationInfo struct {
	// 用户在渠道的机构编号
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 机构在平台的编号
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 用户渠道
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// 用户真实的IP
	//
	// Deprecated: ClientIp is deprecated.
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 机构的代理IP
	//
	// Deprecated: ProxyIp is deprecated.
	ProxyIp *string `json:"ProxyIp,omitnil,omitempty" name:"ProxyIp"`
}

type PdfVerifyResult struct {
	// 验签结果详情，每个签名域对应的验签结果。状态值如下
	// <ul><li> **1** :验签成功，在电子签签署</li>
	// <li> **2** :验签成功，在其他平台签署</li>
	// <li> **3** :验签失败</li>
	// <li> **4** :pdf文件没有签名域</li>
	// <li> **5** :文件签名格式错误</li></ul>
	VerifyResult *int64 `json:"VerifyResult,omitnil,omitempty" name:"VerifyResult"`

	// 签署平台
	// 如果文件是在腾讯电子签平台签署，则为**腾讯电子签**，
	// 如果文件不在腾讯电子签平台签署，则为**其他平台**。
	SignPlatform *string `json:"SignPlatform,omitnil,omitempty" name:"SignPlatform"`

	// 申请证书的主体的名字
	// 
	// 如果是在腾讯电子签平台签署, 则对应的主体的名字个数如下
	// **企业**:  ESS@企业名称@平台生成的数字编码
	// **个人**: ESS@个人姓名@证件号@平台生成的数字编码
	// 
	// 如果在其他平台签署的, 主体的名字参考其他平台的说明
	SignerName *string `json:"SignerName,omitnil,omitempty" name:"SignerName"`

	// 签署时间的Unix时间戳，单位毫秒
	SignTime *int64 `json:"SignTime,omitnil,omitempty" name:"SignTime"`

	// 证书签名算法,  如SHA1withRSA等算法
	SignAlgorithm *string `json:"SignAlgorithm,omitnil,omitempty" name:"SignAlgorithm"`

	// 在数字证书申请过程中，系统会自动生成一个独一无二的序列号。
	CertSn *string `json:"CertSn,omitnil,omitempty" name:"CertSn"`

	// 证书起始时间的Unix时间戳，单位毫秒
	CertNotBefore *int64 `json:"CertNotBefore,omitnil,omitempty" name:"CertNotBefore"`

	// 证书过期时间的时间戳，单位毫秒
	CertNotAfter *int64 `json:"CertNotAfter,omitnil,omitempty" name:"CertNotAfter"`

	// 签名类型, 保留字段, 现在全部为0
	// 
	SignType *int64 `json:"SignType,omitnil,omitempty" name:"SignType"`

	// 签名域横坐标，单位px
	ComponentPosX *float64 `json:"ComponentPosX,omitnil,omitempty" name:"ComponentPosX"`

	// 签名域纵坐标，单位px
	ComponentPosY *float64 `json:"ComponentPosY,omitnil,omitempty" name:"ComponentPosY"`

	// 签名域宽度，单位px
	ComponentWidth *float64 `json:"ComponentWidth,omitnil,omitempty" name:"ComponentWidth"`

	// 签名域高度，单位px
	ComponentHeight *float64 `json:"ComponentHeight,omitnil,omitempty" name:"ComponentHeight"`

	// 签名域所在页码，1～N
	ComponentPage *int64 `json:"ComponentPage,omitnil,omitempty" name:"ComponentPage"`
}

type Permission struct {
	// 权限名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 权限key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 权限类型 1前端，2后端
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否隐藏
	Hide *int64 `json:"Hide,omitnil,omitempty" name:"Hide"`

	// 数据权限标签 1:表示根节点，2:表示叶子结点
	DataLabel *int64 `json:"DataLabel,omitnil,omitempty" name:"DataLabel"`

	// 数据权限独有，1:关联其他模块鉴权，2:表示关联自己模块鉴权
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// 数据权限独有，表示数据范围，1：全公司，2:部门及下级部门，3:自己
	DataRange *int64 `json:"DataRange,omitnil,omitempty" name:"DataRange"`

	// 关联权限, 表示这个功能权限要受哪个数据权限管控
	DataTo *string `json:"DataTo,omitnil,omitempty" name:"DataTo"`

	// 父级权限key
	ParentKey *string `json:"ParentKey,omitnil,omitempty" name:"ParentKey"`

	// 是否选中
	IsChecked *bool `json:"IsChecked,omitnil,omitempty" name:"IsChecked"`

	// 子权限集合
	Children []*Permission `json:"Children,omitnil,omitempty" name:"Children"`
}

type PermissionGroup struct {
	// 权限组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 权限组key
	GroupKey *string `json:"GroupKey,omitnil,omitempty" name:"GroupKey"`

	// 是否隐藏分组，0否1是
	Hide *int64 `json:"Hide,omitnil,omitempty" name:"Hide"`

	// 权限集合
	Permissions []*Permission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

// Predefined struct for user
type PrepareFlowsRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 多个合同（签署流程）信息，最大支持20个签署流程。
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil,omitempty" name:"FlowInfos"`

	// 操作完成后的跳转地址，最大长度200
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type PrepareFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 多个合同（签署流程）信息，最大支持20个签署流程。
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil,omitempty" name:"FlowInfos"`

	// 操作完成后的跳转地址，最大长度200
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	ConfirmUrl *string `json:"ConfirmUrl,omitnil,omitempty" name:"ConfirmUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type PresetApproverInfo struct {
	// 预设参与方姓名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 预设参与方手机号。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 预设参与方证件号，需要和IdCardType同时传入。
	// 
	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成(如存在X，请大写)。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 预设参与方的证件类型，需要与IdCardNumber同时传入。
	// 
	// 证件类型，支持以下类型
	// <ul><li><b>ID_CARD</b>: 居民身份证</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`
}

type ProxyOrganizationOperator struct {
	// 员工的唯一标识(即OpenId),  定义Agent中的OpenId一样, 可以参考<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#agent" target="_blank">Agent结构体</a>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 员工的姓名，最大长度50个字符
	// 员工的姓名将用于身份认证和电子签名，请确保填写的姓名为真实姓名，而非昵称等代名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 签署方经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO :  中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN :  中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 经办人证件号
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 员工的手机号，支持国内手机号11位数字(无需加+86前缀或其他字符)，不支持海外手机号。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 预先分配员工的角色, 可以分配的角色如下:
	// <table> <thead> <tr> <th>可以分配的角色</th> <th>角色名称</th> <th>角色描述</th> </tr> </thead> <tbody> <tr> <td>admin</td> <td>业务管理员（IT 系统负责人，e.g. CTO）</td> <td>有企业合同模块、印章模块、模板模块等全量功能及数据权限。</td> </tr> <tr> <td>channel-normal-operator</td> <td>经办人（企业法务负责人）</td> <td>有发起合同、签署合同（含填写、拒签）、撤销合同、持有印章等权限能力，可查看企业所有合同数据。</td> </tr> <tr> <td>channel-sales-man</td> <td>业务员（一般为销售员、采购员）</td> <td>有发起合同、签署合同（含填写、拒签）、撤销合同、持有印章等权限能力，可查看自己相关所有合同数据。</td> </tr> </tbody> </table>
	DefaultRole *string `json:"DefaultRole,omitnil,omitempty" name:"DefaultRole"`
}

type Recipient struct {
	// 合同参与方的角色ID
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 参与者类型, 可以选择的类型如下:
	// <ul><li> **ENTERPRISE** :此角色为企业参与方</li>
	// <li> **INDIVIDUAL** :此角色为个人参与方</li>
	// <li> **PROMOTER** :此角色是发起方</li></ul>
	RecipientType *string `json:"RecipientType,omitnil,omitempty" name:"RecipientType"`

	// 合同参与方的角色描述，长度不能超过100，只能由中文、字母、数字和下划线组成。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 合同参与方的角色名字，长度不能超过20，只能由中文、字母、数字和下划线组成。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 是否需要校验，
	// true-是，
	// false-否
	RequireValidation *bool `json:"RequireValidation,omitnil,omitempty" name:"RequireValidation"`

	// 是否必须填写，
	// true-是，
	// false-否
	RequireSign *bool `json:"RequireSign,omitnil,omitempty" name:"RequireSign"`

	// 内部字段，签署类型
	SignType *int64 `json:"SignType,omitnil,omitempty" name:"SignType"`

	// 签署顺序：数字越小优先级越高
	RoutingOrder *int64 `json:"RoutingOrder,omitnil,omitempty" name:"RoutingOrder"`

	// 是否是发起方，
	// true-是 
	// false-否
	IsPromoter *bool `json:"IsPromoter,omitnil,omitempty" name:"IsPromoter"`

	// 签署人查看合同校验方式, 支持的类型如下:
	// <ul><li> 1 :实名认证查看</li>
	// <li> 2 :手机号校验查看</li></ul>
	ApproverVerifyTypes []*int64 `json:"ApproverVerifyTypes,omitnil,omitempty" name:"ApproverVerifyTypes"`

	// 签署人进行合同签署时的认证方式，支持的类型如下:
	// <ul><li> 1 :人脸认证</li>
	// <li> 2 :签署密码</li>
	// <li> 3 :运营商三要素认证</li>
	// <li> 4 :UKey认证</li>
	// <li> 5 :设备指纹识别</li>
	// <li> 6 :设备面容识别</li></ul>
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil,omitempty" name:"ApproverSignTypes"`

	// 签署方是否可以转他人处理
	// 
	// <ul><li> **false** : ( 默认)可以转他人处理</li>
	// <li> **true** :不可以转他人处理</li></ul>
	NoTransfer *bool `json:"NoTransfer,omitnil,omitempty" name:"NoTransfer"`
}

type RecipientComponentInfo struct {
	// 参与方的角色ID
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 参与方填写状态
	// 
	// <ul><li> **0** : 还没有填写</li>
	// <li> **1** : 已经填写</li></ul>
	RecipientFillStatus *string `json:"RecipientFillStatus,omitnil,omitempty" name:"RecipientFillStatus"`

	// 此角色是否是发起方角色
	// 
	// <ul><li> **true** : 是发起方角色</li>
	// <li> **false** : 不是发起方角色</li></ul>
	IsPromoter *bool `json:"IsPromoter,omitnil,omitempty" name:"IsPromoter"`

	// 此角色的填写控件列表
	Components []*FilledComponent `json:"Components,omitnil,omitempty" name:"Components"`
}

type RegistrationOrganizationInfo struct {
	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 机构在贵司业务系统中的唯一标识，用于与腾讯电子签企业账号进行映射，确保在同一应用内不会出现重复。
	// 该标识最大长度为64位字符串，仅支持包含26个英文字母和数字0-9的字符。
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 员工在贵司业务系统中的唯一身份标识，用于与腾讯电子签账号进行映射，确保在同一应用内不会出现重复。
	// 该标识最大长度为64位字符串，仅支持包含26个英文字母和数字0-9的字符。
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 组织机构企业统一社会信用代码。
	// 请确认该企业统一社会信用代码与企业营业执照中注册的统一社会信用代码一致。
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// 组织机构法人的姓名。
	// 请确认该企业统一社会信用代码与企业营业执照中注册的法人姓名一致。
	LegalName *string `json:"LegalName,omitnil,omitempty" name:"LegalName"`

	// 组织机构企业注册地址。
	// 请确认该企业注册地址与企业营业执照中注册的地址一致。
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 组织机构超管姓名。
	// 在注册流程中，必须是超管本人进行操作。
	// 如果法人作为超管管理组织机构,超管姓名就是法人姓名
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 组织机构超管手机号。
	// 在注册流程中，这个手机号必须跟操作人在电子签注册的个人手机号一致。
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`

	// 可选的此企业允许的授权方式, 可以设置的方式有:
	// 1：上传授权书
	// 2：法人授权超管
	// 5：授权书+对公打款
	// 
	// 
	// 注:
	// `1. 当前仅支持一种认证方式`
	// `2. 如果当前的企业类型是政府/事业单位, 则只支持上传授权书+对公打款`
	// `3. 如果当前操作人是法人,则是法人认证`
	AuthorizationTypes []*uint64 `json:"AuthorizationTypes,omitnil,omitempty" name:"AuthorizationTypes"`

	// 经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	AdminIdCardType *string `json:"AdminIdCardType,omitnil,omitempty" name:"AdminIdCardType"`

	// 经办人的证件号
	AdminIdCardNumber *string `json:"AdminIdCardNumber,omitnil,omitempty" name:"AdminIdCardNumber"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitnil,omitempty" name:"BusinessLicense"`

	// 授权书(PNG或JPG或PDF) base64格式, 大小不超过8M 。
	// p.s. 如果上传授权书 ，需遵循以下条件
	// 1. 超管的信息（超管姓名，超管身份证，超管手机号）必须为必填参数。
	// 2. 超管的个人身份必须在电子签已经实名。
	// 2. 认证方式AuthorizationTypes必须只能是上传授权书方式 
	PowerOfAttorneys []*string `json:"PowerOfAttorneys,omitnil,omitempty" name:"PowerOfAttorneys"`

	// 认证完之后的H5页面的跳转链接，最大长度1000个字符。链接类型请参考 [跳转电子签H5](https://qian.tencent.com/developers/company/openqianh5/)
	AutoJumpUrl *string `json:"AutoJumpUrl,omitnil,omitempty" name:"AutoJumpUrl"`
}

type ReleasedApprover struct {
	// 签署人在原合同签署人列表中的顺序序号(从0开始，按顺序依次递增)。
	// 可以通过<a href="https://qian.tencent.com/developers/partnerApis/flows/DescribeFlowDetailInfo" target="_blank">DescribeFlowDetailInfo</a>接口查看原流程中的签署人列表。
	ApproverNumber *uint64 `json:"ApproverNumber,omitnil,omitempty" name:"ApproverNumber"`

	// 指定签署人类型，目前支持
	// <ul><li> **ORGANIZATION**：企业(默认值)</li>
	// <li> **ENTERPRISESERVER**：企业静默签</li></ul>
	ApproverType *string `json:"ApproverType,omitnil,omitempty" name:"ApproverType"`

	// 签署人在原合同中的RecipientId，可以通过<a href="https://qian.tencent.com/developers/partnerApis/flows/DescribeFlowDetailInfo" target="_blank">DescribeFlowDetailInfo</a>接口查看原流程中的签署人信息，可参考返回结构体<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#flowapproverdetail" target="_blank">FlowApproverDetail</a>中的RecipientId。
	// **注意**：当指定了此参数后，ApproverNumber即失效，会以本参数作为原合同参与人的选取。与ApproverNumber二选一。
	ReleasedApproverRecipientId *string `json:"ReleasedApproverRecipientId,omitnil,omitempty" name:"ReleasedApproverRecipientId"`

	// 签署人姓名，最大长度50个字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 签署方经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证(默认值)</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成(如存在X，请大写)。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。
	// </li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 签署人手机号。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 如果签署方是企业签署方(approverType = 0 或者 approverType = 3)， 则企业名称必填。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 第三方平台子客企业的唯一标识，定义Agent中的ProxyOrganizationOpenId一样, 可以参考<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#agent" target="_blank">Agent结构体</a>。
	// <font color="red">当为子客企业指定经办人时，此OrganizationOpenId必传。</font>
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 第三方平台子客企业员工的唯一标识，长度不能超过64，只能由字母和数字组成。
	// <font color="red">当签署方为同一第三方平台下的员工时，此OpenId必传。</font>
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 签署控件类型，支持自定义企业签署方的签署控件类型
	// <ul><li> **SIGN_SEAL**：默认为印章控件类型(默认值)</li>
	// <li> **SIGN_SIGNATURE**：手写签名控件类型</li></ul>
	ApproverSignComponentType *string `json:"ApproverSignComponentType,omitnil,omitempty" name:"ApproverSignComponentType"`

	// 参与方在合同中的角色是按照创建合同的时候来排序的，解除协议默认会将第一个参与人叫`甲方`,第二个叫`乙方`,  第三个叫`丙方`，以此类推。
	// 如果需改动此参与人的角色名字，可用此字段指定，由汉字,英文字符,数字组成，最大20个字。
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/973a820ab66d1ce57082c160c2b2d44a.png)
	ApproverSignRole *string `json:"ApproverSignRole,omitnil,omitempty" name:"ApproverSignRole"`

	// 印章Id，签署控件类型为印章时，用于指定本企业签署方在解除协议中使用那个印章进行签署
	ApproverSignSealId *string `json:"ApproverSignSealId,omitnil,omitempty" name:"ApproverSignSealId"`
}

type RelieveInfo struct {
	// 解除理由，长度不能超过200，只能由中文、字母、数字、中文标点和英文标点组成(不支持表情)。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 解除后仍然有效的条款，保留条款，长度不能超过200，只能由中文、字母、数字、中文标点和英文标点组成(不支持表情)。
	RemainInForceItem *string `json:"RemainInForceItem,omitnil,omitempty" name:"RemainInForceItem"`

	// 原合同事项处理-费用结算，长度不能超过200，只能由中文、字母、数字、中文标点和英文标点组成(不支持表情)。
	OriginalExpenseSettlement *string `json:"OriginalExpenseSettlement,omitnil,omitempty" name:"OriginalExpenseSettlement"`

	// 原合同事项处理-其他事项，长度不能超过200，只能由中文、字母、数字、中文标点和英文标点组成(不支持表情)。
	OriginalOtherSettlement *string `json:"OriginalOtherSettlement,omitnil,omitempty" name:"OriginalOtherSettlement"`

	// 其他约定（如约定的与解除协议存在冲突的，以【其他约定】为准），最大支持200个字，只能由中文、字母、数字、中文标点和英文标点组成(不支持表情)。
	OtherDeals *string `json:"OtherDeals,omitnil,omitempty" name:"OtherDeals"`
}

type RemindFlowRecords struct {
	// 合同流程是否可以催办： true - 可以，false - 不可以。 若无法催办，将返回RemindMessage以解释原因。	
	CanRemind *bool `json:"CanRemind,omitnil,omitempty" name:"CanRemind"`

	// 合同流程ID，为32位字符串。	
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 在合同流程无法催办的情况下，系统将返回RemindMessage以阐述原因。	
	RemindMessage *string `json:"RemindMessage,omitnil,omitempty" name:"RemindMessage"`
}

type ResourceUrlInfo struct {
	// 资源链接地址，过期时间5分钟
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 资源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type SignComponentConfig struct {
	// 签署控件默认属性配置，是否默认展示签署日期， 在页面中可以进行修改。
	// 
	// - false 展示签署日期（默认）
	// - true 不展示签署日期 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/448514412e2f69f6129425beda4ff568.png)。
	HideDate *bool `json:"HideDate,omitnil,omitempty" name:"HideDate"`
}

type SignQrCode struct {
	// 二维码ID，为32位字符串。	
	// 
	// 注: 需要保留此二维码ID, 用于后序通过<a href="https://qian.tencent.com/developers/partnerApis/templates/ChannelCancelMultiFlowSignQRCode" target="_blank">取消一码多扫二维码</a>关闭这个二维码的签署功能。	
	QrCodeId *string `json:"QrCodeId,omitnil,omitempty" name:"QrCodeId"`

	// 二维码URL，可通过转换二维码的工具或代码组件将此URL转化为二维码，以便用户扫描进行流程签署。	
	QrCodeUrl *string `json:"QrCodeUrl,omitnil,omitempty" name:"QrCodeUrl"`

	// 二维码的有截止时间，格式为Unix标准时间戳（秒），可以通过入参的QrEffectiveDay来设置有效期，默认为7天有效期。 
	// 一旦超过二维码的有效期限，该二维码将自动失效。	
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 微信小程序二维码
	WeixinQrCodeUrl *string `json:"WeixinQrCodeUrl,omitnil,omitempty" name:"WeixinQrCodeUrl"`
}

type SignUrl struct {
	// 跳转至电子签名小程序签署的链接地址。 适用于客户端APP及小程序直接唤起电子签名小程序。	
	AppSignUrl *string `json:"AppSignUrl,omitnil,omitempty" name:"AppSignUrl"`

	// 签署链接有效时间，格式类似"2022-08-05 15:55:01"	
	EffectiveTime *string `json:"EffectiveTime,omitnil,omitempty" name:"EffectiveTime"`

	// 跳转至电子签名小程序签署的链接地址，格式类似于https://essurl.cn/xxx。 打开此链接将会展示H5中间页面，随后唤起电子签名小程序以进行合同签署。	
	HttpSignUrl *string `json:"HttpSignUrl,omitnil,omitempty" name:"HttpSignUrl"`
}

type SignUrlInfo struct {
	// 签署链接，过期时间为90天
	// 
	// 注：<font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	SignUrl *string `json:"SignUrl,omitnil,omitempty" name:"SignUrl"`

	// 合同过期时间戳，单位秒
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 当流程为顺序签署此参数有效时，数字越小优先级越高，暂不支持并行签署 可选
	SignOrder *int64 `json:"SignOrder,omitnil,omitempty" name:"SignOrder"`

	// 签署人编号
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 自定义用户编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: CustomUserId is deprecated.
	CustomUserId *string `json:"CustomUserId,omitnil,omitempty" name:"CustomUserId"`

	// 用户姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户手机号码
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 签署参与者机构名字
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 参与者类型, 类型如下:
	// **ORGANIZATION**:企业经办人
	// **PERSON**: 自然人
	ApproverType *string `json:"ApproverType,omitnil,omitempty" name:"ApproverType"`

	// 经办人身份证号
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 签署链接对应流程Id
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 企业经办人 用户在渠道的编号
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 合同组签署链接对应的合同组id
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 二维码，在生成动态签署人跳转封面页链接时返回
	// 
	// 注：`此二维码下载链接有效期为5分钟，可下载二维码后本地保存，二维码有效期为90天。`
	SignQrcodeUrl *string `json:"SignQrcodeUrl,omitnil,omitempty" name:"SignQrcodeUrl"`
}

type Staff struct {
	// 员工在电子签平台的用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 显示的员工名
	// 注意：2024-07-08 及之后创建的应用号，该字段返回的是打码信息
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 员工手机号
	// 注意：2024-07-08 及之后创建的应用号，该字段返回的是打码信息
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 员工邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 员工在第三方应用平台的用户ID
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 员工角色
	Roles []*StaffRole `json:"Roles,omitnil,omitempty" name:"Roles"`

	// 员工部门
	Department *Department `json:"Department,omitnil,omitempty" name:"Department"`

	// 员工是否实名
	Verified *bool `json:"Verified,omitnil,omitempty" name:"Verified"`

	// 员工创建时间戳，单位秒
	CreatedOn *int64 `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 员工实名时间戳，单位秒
	VerifiedOn *int64 `json:"VerifiedOn,omitnil,omitempty" name:"VerifiedOn"`

	// 员工是否离职：0-未离职，1-离职
	QuiteJob *int64 `json:"QuiteJob,omitnil,omitempty" name:"QuiteJob"`
}

type StaffRole struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type SyncFailReason struct {
	// 企业员工标识(即OpenId)
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 新增员工或者员工离职失败原因, 可能存证ID不符合规范、证件号码不合法等原因
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type SyncProxyOrganizationOperatorsRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// </ul>
	// 第三方平台子客企业必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 操作类型，对应的操作
	// <ul><li> **CREATE** :新增员工</li>
	// <li> **UPDATE** :修改员工（仅支持修改未实名员工的信息，如果已经实名并加入企业的员工基础信息修改需要到小程序中进行）</li>
	// <li> **RESIGN** :离职员工</li></ul>
	OperatorType *string `json:"OperatorType,omitnil,omitempty" name:"OperatorType"`

	// 员工信息列表，最多支持200个
	ProxyOrganizationOperators []*ProxyOrganizationOperator `json:"ProxyOrganizationOperators,omitnil,omitempty" name:"ProxyOrganizationOperators"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type SyncProxyOrganizationOperatorsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// </ul>
	// 第三方平台子客企业必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 操作类型，对应的操作
	// <ul><li> **CREATE** :新增员工</li>
	// <li> **UPDATE** :修改员工（仅支持修改未实名员工的信息，如果已经实名并加入企业的员工基础信息修改需要到小程序中进行）</li>
	// <li> **RESIGN** :离职员工</li></ul>
	OperatorType *string `json:"OperatorType,omitnil,omitempty" name:"OperatorType"`

	// 员工信息列表，最多支持200个
	ProxyOrganizationOperators []*ProxyOrganizationOperator `json:"ProxyOrganizationOperators,omitnil,omitempty" name:"ProxyOrganizationOperators"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	//  同步的状态,  全部同步失败接口是接口会直接报错
	// 
	// <ul><li> **1** :全部成功</li>
	// <li> **2** :部分成功</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 同步失败员工ID及其失败原因
	FailedList []*SyncFailReason `json:"FailedList,omitnil,omitempty" name:"FailedList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 第三方平台子客企业名称，请确认该名称与企业营业执照中注册的名称一致。
	// 注: `如果名称中包含英文括号()，请使用中文括号（）代替。`
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil,omitempty" name:"ProxyOrganizationName"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitnil,omitempty" name:"BusinessLicense"`

	// 第三方平台子客企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// 第三方平台子客企业法定代表人的名字
	ProxyLegalName *string `json:"ProxyLegalName,omitnil,omitempty" name:"ProxyLegalName"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 第三方平台子客企业法定代表人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证 (默认值)</li></ul>
	// 注: `现在仅支持ID_CARD中国大陆居民身份证类型`
	ProxyLegalIdCardType *string `json:"ProxyLegalIdCardType,omitnil,omitempty" name:"ProxyLegalIdCardType"`

	// 第三方平台子客企业法定代表人的证件号码, 应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li></ul>
	ProxyLegalIdCardNumber *string `json:"ProxyLegalIdCardNumber,omitnil,omitempty" name:"ProxyLegalIdCardNumber"`

	// 第三方平台子客企业详细住所，最大长度500个字符
	// 
	// 注：`需要符合省市区详情的格式例如： XX省XX市XX区街道具体地址`
	ProxyAddress *string `json:"ProxyAddress,omitnil,omitempty" name:"ProxyAddress"`
}

type SyncProxyOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 第三方平台子客企业名称，请确认该名称与企业营业执照中注册的名称一致。
	// 注: `如果名称中包含英文括号()，请使用中文括号（）代替。`
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil,omitempty" name:"ProxyOrganizationName"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitnil,omitempty" name:"BusinessLicense"`

	// 第三方平台子客企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// 第三方平台子客企业法定代表人的名字
	ProxyLegalName *string `json:"ProxyLegalName,omitnil,omitempty" name:"ProxyLegalName"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 第三方平台子客企业法定代表人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证 (默认值)</li></ul>
	// 注: `现在仅支持ID_CARD中国大陆居民身份证类型`
	ProxyLegalIdCardType *string `json:"ProxyLegalIdCardType,omitnil,omitempty" name:"ProxyLegalIdCardType"`

	// 第三方平台子客企业法定代表人的证件号码, 应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li></ul>
	ProxyLegalIdCardNumber *string `json:"ProxyLegalIdCardNumber,omitnil,omitempty" name:"ProxyLegalIdCardNumber"`

	// 第三方平台子客企业详细住所，最大长度500个字符
	// 
	// 注：`需要符合省市区详情的格式例如： XX省XX市XX区街道具体地址`
	ProxyAddress *string `json:"ProxyAddress,omitnil,omitempty" name:"ProxyAddress"`
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
	delete(f, "ProxyLegalName")
	delete(f, "Operator")
	delete(f, "ProxyLegalIdCardType")
	delete(f, "ProxyLegalIdCardNumber")
	delete(f, "ProxyAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncProxyOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncProxyOrganizationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type TaskInfo struct {
	// 合成任务Id，可以通过 ChannelGetTaskResultApi 接口获取任务信息
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态：READY - 任务已完成；NOTREADY - 任务未完成；
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`
}

type TemplateInfo struct {
	// 模板ID，模板的唯一标识
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 模板描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 模板的填充控件列表
	// 
	// [点击查看在模板中配置的填充控件的样子](https://qcloudimg.tencent-cloud.cn/raw/cb2f58529fca8d909258f9d45a56f7f4.png)
	Components []*Component `json:"Components,omitnil,omitempty" name:"Components"`

	// 此模块需要签署的各个参与方的角色列表。RecipientId标识每个参与方角色对应的唯一标识符，用于确定此角色的信息。
	// 
	// [点击查看在模板中配置的签署参与方角色列表的样子](https://qcloudimg.tencent-cloud.cn/raw/e082bbcc0d923f8cb723d98382410aa2.png)
	// 
	Recipients []*Recipient `json:"Recipients,omitnil,omitempty" name:"Recipients"`

	// 此模板中的签署控件列表
	// 
	// [点击查看在模板中配置的签署控件的样子](https://qcloudimg.tencent-cloud.cn/raw/29bc6ed753a5a0fce4a3ab02e2c0d955.png)
	SignComponents []*Component `json:"SignComponents,omitnil,omitempty" name:"SignComponents"`

	// 模板类型可以分为以下两种：
	// 
	// <b>1</b>：带有<b>本企业自动签署</b>的模板，即签署过程无需签署人手动操作，系统自动完成签署。
	// <b>3</b>：普通模板，即签署人需要手动进行签署操作。
	TemplateType *int64 `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// 是否是发起人 ,已弃用
	//
	// Deprecated: IsPromoter is deprecated.
	IsPromoter *bool `json:"IsPromoter,omitnil,omitempty" name:"IsPromoter"`

	// 模板的创建者名字
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 模板创建的时间戳，格式为Unix标准时间戳（秒）
	CreatedOn *int64 `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 模板的 H5 预览链接，有效期为 5 分钟。
	// 您可以通过浏览器直接打开此链接预览模板，或将其嵌入到 iframe 中进行预览。
	// 
	// 注意：只有在请求接口时将 <b>WithPreviewUrl </b>参数设置为 true，才会生成预览链接。
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`

	// 第三方应用集成-模板PDF文件链接，有效期5分钟。
	// 请求参数WithPdfUrl=true时返回
	// （此功能开放需要联系客户经理）。
	PdfUrl *string `json:"PdfUrl,omitnil,omitempty" name:"PdfUrl"`

	// 本模板关联的第三方应用平台企业模板ID
	ChannelTemplateId *string `json:"ChannelTemplateId,omitnil,omitempty" name:"ChannelTemplateId"`

	// 本模板关联的三方应用平台平台企业模板名称
	ChannelTemplateName *string `json:"ChannelTemplateName,omitnil,omitempty" name:"ChannelTemplateName"`

	// 0-需要子客企业手动领取平台企业的模板(默认); 
	// 1-平台自动设置子客模板
	ChannelAutoSave *int64 `json:"ChannelAutoSave,omitnil,omitempty" name:"ChannelAutoSave"`

	// 模板版本，由全数字字符组成。
	// 默认为空，模板版本号由日期和序号组成，初始版本为yyyyMMdd001，yyyyMMdd002表示第二个版本，以此类推。
	TemplateVersion *string `json:"TemplateVersion,omitnil,omitempty" name:"TemplateVersion"`

	// 模板可用状态的取值通常为以下两种：
	// 
	// <ul><li>1：启用（默认），表示模板处于启用状态，可以被用户正常使用。</li>
	// <li>2：停用，表示模板处于停用状态，禁止用户使用该模板。</li></ul>
	Available *int64 `json:"Available,omitnil,omitempty" name:"Available"`

	// 模版的用户合同类型
	UserFlowType *UserFlowType `json:"UserFlowType,omitnil,omitempty" name:"UserFlowType"`
}

type TemplateUserFlowType struct {
	// 合同类型id
	UserFlowTypeId *string `json:"UserFlowTypeId,omitnil,omitempty" name:"UserFlowTypeId"`

	// 用户合同类型名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 每个合同类型绑定的模版数量	
	TemplateNum *int64 `json:"TemplateNum,omitnil,omitempty" name:"TemplateNum"`

	// 合同类型的具体描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UploadFile struct {
	// Base64编码后的文件内容
	FileBody *string `json:"FileBody,omitnil,omitempty" name:"FileBody"`

	// 文件的名字。
	// 文件名的最大长度应不超过200个字符，并且文件名的后缀必须反映其文件类型。
	// 例如，PDF文件应以“.pdf”结尾，如“XXX.pdf”，而Word文件应以“.doc”或“.docx”结尾，如“XXX.doc”或“XXX.docx”。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

// Predefined struct for user
type UploadFilesRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 
	// 文件对应业务类型,可以选择的类型如下
	// <ul><li> **TEMPLATE** : 此上传的文件用户生成合同模板，文件类型支持.pdf/.doc/.docx/.html格式，如果非pdf文件需要通过<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi" target="_blank">创建文件转换任务</a>转换后才能使用</li>
	// <li> **DOCUMENT** : 此文件用来发起合同流程，文件类型支持.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html，如果非pdf文件需要通过<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi" target="_blank">创建文件转换任务</a>转换后才能使用</li></ul>
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// 上传文件内容数组，一次最多可上传20个文件。
	// 
	// <b>若上传多个文件，所有文件必须为相同类型</b>，例如全部为PDF或全部为Word文件。不支持混合文件类型的上传。
	FileInfos []*UploadFile `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 操作者的信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type UploadFilesRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 
	// 文件对应业务类型,可以选择的类型如下
	// <ul><li> **TEMPLATE** : 此上传的文件用户生成合同模板，文件类型支持.pdf/.doc/.docx/.html格式，如果非pdf文件需要通过<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi" target="_blank">创建文件转换任务</a>转换后才能使用</li>
	// <li> **DOCUMENT** : 此文件用来发起合同流程，文件类型支持.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html，如果非pdf文件需要通过<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi" target="_blank">创建文件转换任务</a>转换后才能使用</li></ul>
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// 上传文件内容数组，一次最多可上传20个文件。
	// 
	// <b>若上传多个文件，所有文件必须为相同类型</b>，例如全部为PDF或全部为Word文件。不支持混合文件类型的上传。
	FileInfos []*UploadFile `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	// 上传成功文件数量
	// 注: `如果一个文件上传失败, 则全部文件皆上传失败`
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 文件资源ID数组，每个文件资源ID为32位字符串。
	// 建议开发者保存此资源ID，后续创建合同或创建合同流程需此资源ID。
	// 注:`有效期一个小时（超过一小时后系统不定期清理，会有部分时间差）, 有效期内此文件id可以反复使用, 超过有效期无法使用`
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 对应上传文件的下载链接，过期时间5分钟
	FileUrls []*string `json:"FileUrls,omitnil,omitempty" name:"FileUrls"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 子客企业标识
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil,omitempty" name:"ProxyOrganizationOpenId"`

	// 子客企业名
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil,omitempty" name:"ProxyOrganizationName"`

	// 对应的消耗日期, **如果是汇总数据则为1970-01-01**
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 消耗合同数量
	Usage *uint64 `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 撤回合同数量
	Cancel *uint64 `json:"Cancel,omitnil,omitempty" name:"Cancel"`

	// 消耗渠道
	FlowChannel *string `json:"FlowChannel,omitnil,omitempty" name:"FlowChannel"`
}

type UserFlowType struct {
	// 用户合同类型id
	UserFlowTypeId *string `json:"UserFlowTypeId,omitnil,omitempty" name:"UserFlowTypeId"`

	// 用户合同类型名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户合同类型的描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UserInfo struct {
	// 第三方应用平台自定义，对应第三方平台子客企业员工的唯一标识。
	// 
	// 
	// 注意:
	// 1. OpenId在子客企业对应一个真实员工，**本应用唯一, 不可重复使用**，最大64位字符串
	// 2. 可使用用户在贵方企业系统中的Userid或者hash值作为子客企业的员工OpenId
	// 3. **员工加入企业后**, 可以通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>登录子客控制台后, 在**组织架构**模块查看员工们的OpenId, 样式如下图
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/bb67fb66c926759df3a0af5838fdafd5.png)
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 内部参数，暂未开放使用
	//
	// Deprecated: Channel is deprecated.
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// 内部参数，暂未开放使用
	//
	// Deprecated: CustomUserId is deprecated.
	CustomUserId *string `json:"CustomUserId,omitnil,omitempty" name:"CustomUserId"`

	// 内部参数，暂未开放使用
	//
	// Deprecated: ClientIp is deprecated.
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 内部参数，暂未开放使用
	//
	// Deprecated: ProxyIp is deprecated.
	ProxyIp *string `json:"ProxyIp,omitnil,omitempty" name:"ProxyIp"`
}

type UserThreeFactor struct {
	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`
}

type WebThemeConfig struct {
	// 是否显示页面底部电子签logo，取值如下：
	// <ul><li> **true**：页面底部显示电子签logo</li>
	// <li> **false**：页面底部不显示电子签logo（默认）</li></ul>
	DisplaySignBrandLogo *bool `json:"DisplaySignBrandLogo,omitnil,omitempty" name:"DisplaySignBrandLogo"`

	// 主题颜色：
	// 支持十六进制颜色值以及RGB格式颜色值，例如：#D54941，rgb(213, 73, 65)
	// <br/>
	WebEmbedThemeColor *string `json:"WebEmbedThemeColor,omitnil,omitempty" name:"WebEmbedThemeColor"`

	// 企业认证页背景图（base64图片）
	AuthenticateBackground *string `json:"AuthenticateBackground,omitnil,omitempty" name:"AuthenticateBackground"`

	// 隐藏企业认证页面导航栏，取值如下：
	// <ul><li> **true**：隐藏企业认证页面导航栏</li>
	// <li> **false**：显示企业认证页面导航栏（默认）</li></ul>
	HideAuthenticateNavigationBar *bool `json:"HideAuthenticateNavigationBar,omitnil,omitempty" name:"HideAuthenticateNavigationBar"`

	// 隐藏企业认证顶部logo，取值如下：
	// <ul><li> **true**：隐藏企业认证顶部logo</li>
	// <li> **false**：显示企业认证顶部logo（默认）</li></ul>
	HideAuthenticateTopLogo *bool `json:"HideAuthenticateTopLogo,omitnil,omitempty" name:"HideAuthenticateTopLogo"`
}