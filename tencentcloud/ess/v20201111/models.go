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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Admin struct {
	// 超管名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 超管手机号，打码显示
	// 示例值：138****1569
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`
}

type Agent struct {
	// 代理机构的应用编号,32位字符串，一般不用传
	//
	// Deprecated: AppId is deprecated.
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 被代理机构的应用号，一般不用传
	//
	// Deprecated: ProxyAppId is deprecated.
	ProxyAppId *string `json:"ProxyAppId,omitnil,omitempty" name:"ProxyAppId"`

	// 被代理机构在电子签平台的机构编号，集团代理下场景必传
	ProxyOrganizationId *string `json:"ProxyOrganizationId,omitnil,omitempty" name:"ProxyOrganizationId"`

	// 被代理机构的经办人，一般不用传
	//
	// Deprecated: ProxyOperator is deprecated.
	ProxyOperator *string `json:"ProxyOperator,omitnil,omitempty" name:"ProxyOperator"`
}

type ApproverComponentLimitType struct {
	// 签署方经办人在模板中配置的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 签署方经办人控件类型是个人印章签署控件（SIGN_SIGNATURE） 时，可选的签名方式，可多选
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

type ApproverInfo struct {
	// 在指定签署方时，可选择企业B端或个人C端等不同的参与者类型，可选类型如下:
	// **0**：企业
	// **1**：个人
	// **3**：企业静默签署
	// 注：`类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。静默签署仅进行盖章操作，不能自动签名。`
	// **7**: 个人自动签署，适用于个人自动签场景。
	// 注: `个人自动签场景为白名单功能，使用前请联系对接的客户经理沟通。`
	ApproverType *int64 `json:"ApproverType,omitnil,omitempty" name:"ApproverType"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	ApproverName *string `json:"ApproverName,omitnil,omitempty" name:"ApproverName"`

	// 签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此合同签署方。
	ApproverMobile *string `json:"ApproverMobile,omitnil,omitempty" name:"ApproverMobile"`

	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 如果签署方是企业签署方(approverType = 0 或者 approverType = 3)， 则企业名称必填。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 合同中的签署控件列表，列表中可支持下列多种签署控件,控件的详细定义参考开发者中心的Component结构体
	// <ul><li> 个人签名/印章</li>
	// <li> 企业印章</li>
	// <li> 骑缝章等签署控件</li></ul>
	SignComponents []*Component `json:"SignComponents,omitnil,omitempty" name:"SignComponents"`

	// 签署方经办人的证件类型，支持以下类型，样式可以参考<a href="https://qian.tencent.com/developers/partner/id_card_support/" target="_blank">常见个人证件类型介绍</a>
	// <ul><li>ID_CARD 中国大陆居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证(格式同居民身份证)</li>
	// <li>OTHER_CARD_TYPE 其他证件</li></ul>
	// 
	// 
	// 
	// 
	// 注: 
	// 1. <b>其他证件类型为白名单功能</b>，使用前请联系对接的客户经理沟通。
	// 2. 港澳居民来往内地通行证 和  港澳台居民居住证 类型的签署人<b>至少要过一次大陆的海关</b>才能使用。
	ApproverIdCardType *string `json:"ApproverIdCardType,omitnil,omitempty" name:"ApproverIdCardType"`

	// 签署方经办人的证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitnil,omitempty" name:"ApproverIdCardNumber"`

	// 通知签署方经办人的方式,  有以下途径:
	// <ul><li>  **sms**  :  (默认)短信</li>
	// <li>   **none**   : 不通知</li></ul>
	// 
	// 注意：
	// `如果使用的是通过文件发起合同（CreateFlowByFiles），NotifyType必须 是 sms 才会发送短信`
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// 收据场景设置签署人角色类型, 可以设置如下<b>类型</b>:
	// <ul><li> **1**  :收款人</li>
	// <li>   **2**   :开具人</li>
	// <li>   **3** :见证人</li></ul>
	// 注: `收据场景为白名单功能，使用前请联系对接的客户经理沟通。`
	ApproverRole *int64 `json:"ApproverRole,omitnil,omitempty" name:"ApproverRole"`

	// 可以自定义签署人角色名：收款人、开具人、见证人等，长度不能超过20，只能由中文、字母、数字和下划线组成。
	// 
	// 注: `如果是用模板发起, 优先使用此处上传的, 如果不传则用模板的配置的`
	ApproverRoleName *string `json:"ApproverRoleName,omitnil,omitempty" name:"ApproverRoleName"`

	// 签署意愿确认渠道，默认为WEIXINAPP:人脸识别
	// 
	// 注: 将要废弃, 用ApproverSignTypes签署人签署合同时的认证方式代替, 新客户可请用ApproverSignTypes来设置
	VerifyChannel []*string `json:"VerifyChannel,omitnil,omitempty" name:"VerifyChannel"`

	// 签署方在签署合同之前，需要强制阅读合同的时长，可指定为3秒至300秒之间的任意值。
	// 
	// 若未指定阅读时间，则会按照合同页数大小计算阅读时间，计算规则如下：
	// <ul><li>合同页数少于等于2页，阅读时间为3秒；</li>
	// <li>合同页数为3到5页，阅读时间为5秒；</li>
	// <li>合同页数大于等于6页，阅读时间为10秒。</li></ul>
	PreReadTime *int64 `json:"PreReadTime,omitnil,omitempty" name:"PreReadTime"`

	// 签署人userId，仅支持本企业的员工userid， 可在控制台组织管理处获得
	// 
	// 注： 
	// 如果传进来的<font color="red">UserId已经实名， 则忽略ApproverName，ApproverIdCardType，ApproverIdCardNumber，ApproverMobile这四个入参</font>（会用此UserId实名的身份证和登录的手机号覆盖）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 在企微场景下使用，需设置参数为**WEWORKAPP**，以表明合同来源于企微。
	ApproverSource *string `json:"ApproverSource,omitnil,omitempty" name:"ApproverSource"`

	// 在企业微信场景下，表明该合同流程为或签，其最大长度为64位字符串。
	// 所有参与或签的人员均需具备该标识。
	// 注意，在合同中，不同的或签参与人必须保证其CustomApproverTag唯一。
	// 如果或签签署人为本方企业微信参与人，则需要指定ApproverSource参数为WEWORKAPP。
	CustomApproverTag *string `json:"CustomApproverTag,omitnil,omitempty" name:"CustomApproverTag"`

	// 可以控制签署方在签署合同时能否进行某些操作，例如拒签、转交他人等。
	// 详细操作可以参考开发者中心的ApproverOption结构体。
	ApproverOption *ApproverOption `json:"ApproverOption,omitnil,omitempty" name:"ApproverOption"`

	// 指定个人签署方查看合同的校验方式,可以传值如下:
	// <ul><li>  **1**   : （默认）人脸识别,人脸识别后才能合同内容</li>
	// <li>  **2**  : 手机号验证, 用户手机号和参与方手机号(ApproverMobile)相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// </li></ul>
	// 注: 
	// <ul><li>如果合同流程设置ApproverVerifyType查看合同的校验方式,    则忽略此签署人的查看合同的校验方式</li>
	// <li>此字段可传多个校验方式</li></ul>
	ApproverVerifyTypes []*int64 `json:"ApproverVerifyTypes,omitnil,omitempty" name:"ApproverVerifyTypes"`

	// 您可以指定签署方签署合同的认证校验方式，可传递以下值：
	// <ul><li>**1**：人脸认证，需进行人脸识别成功后才能签署合同；</li>
	// <li>**2**：签署密码，需输入与用户在腾讯电子签设置的密码一致才能校验成功进行合同签署；</li>
	// <li>**3**：运营商三要素，需到运营商处比对手机号实名信息（名字、手机号、证件号）校验一致才能成功进行合同签署。（如果是港澳台客户，建议不要选择这个）</li>
	// <li>**5**：设备指纹识别，需要对比手机机主预留的指纹信息，校验一致才能成功进行合同签署。（iOS系统暂不支持该校验方式）</li>
	// <li>**6**：设备面容识别，需要对比手机机主预留的人脸信息，校验一致才能成功进行合同签署。（Android系统暂不支持该校验方式）</li></ul>
	// 
	// 
	// 默认为：
	// 1(人脸认证 ),2(签署密码),3(运营商三要素),5(设备指纹识别),6(设备面容识别)
	// 
	// 注：
	// 1. 用<font color='red'>模板创建合同场景</font>, 签署人的认证方式需要在配置模板的时候指定, <font color='red'>在创建合同重新指定无效</font>
	// 2. 运营商三要素认证方式对手机号运营商及前缀有限制,可以参考[运营商支持列表类](https://qian.tencent.com/developers/company/mobile_support)得到具体的支持说明
	// 3. 校验方式不允许只包含<font color='red'>设备指纹识别</font>和<font color='red'>设备面容识别</font>，至少需要再增加一种其他校验方式。
	// 4. <font color='red'>设备指纹识别</font>和<font color='red'>设备面容识别</font>只支持小程序使用，其他端暂不支持。
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil,omitempty" name:"ApproverSignTypes"`

	// 此签署人（员工或者个人）签署前，是否需要发起方企业审批，取值如下：
	// <ul><li>**false**：（默认）不需要审批，直接签署。</li>
	// <li>**true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li>如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li>如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 
	// 注：`此功能可用于与发起方企业内部的审批流程进行关联，支持手动、静默签署合同`
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/b14d5188ed0229d1401e74a9a49cab6d.png)
	ApproverNeedSignReview *bool `json:"ApproverNeedSignReview,omitnil,omitempty" name:"ApproverNeedSignReview"`

	// [用PDF文件创建签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowByFiles)时,如果设置了外层参数SignBeanTag=1(允许签署过程中添加签署控件),则可通过此参数明确规定合同所使用的签署控件类型（骑缝章、普通章法人章等）和具体的印章（印章ID或者印章类型）或签名方式。
	// 
	// 注：`限制印章控件或骑缝章控件情况下,仅本企业签署方可以指定具体印章（通过传递ComponentValue,支持多个），他方企业或个人只支持限制控件类型。`
	AddSignComponentsLimits []*ComponentLimit `json:"AddSignComponentsLimits,omitnil,omitempty" name:"AddSignComponentsLimits"`

	// 签署须知：支持传入富文本，最长字数：500个中文字符
	SignInstructionContent *string `json:"SignInstructionContent,omitnil,omitempty" name:"SignInstructionContent"`

	// 签署人的签署截止时间，格式为Unix标准时间戳（秒）
	// 
	// 注: `若不设置此参数，则默认使用合同的截止时间，此参数暂不支持合同组子合同`
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 签署人在合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li>单行文本控件</li>
	// <li>多行文本控件</li>
	// <li>勾选框控件</li>
	// <li>数字控件</li>
	// <li>图片控件</li>
	// </ul>
	// 
	// 具体使用说明可参考[为签署方指定填写控件](https://qian.tencent.cn/developers/company/createFlowByFiles/#指定签署方填写控件)
	// 
	// 注：`此参数仅在通过文件发起合同或者合同组时生效`
	Components []*Component `json:"Components,omitnil,omitempty" name:"Components"`

	// 进入签署流程的限制，目前支持以下选项：
	// <ul><li> <b>空值（默认）</b> :无限制，可在任何场景进入签署流程。</li><li> <b>link</b> :选择此选项后，将无法通过控制台或电子签小程序列表进入填写或签署操作，仅可预览合同。填写或签署流程只能通过短信或发起方提供的专用链接进行。</li></ul>
	SignEndpoints []*string `json:"SignEndpoints,omitnil,omitempty" name:"SignEndpoints"`
}

type ApproverItem struct {
	// 签署方唯一编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 签署方角色编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 签署方角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverRoleName *string `json:"ApproverRoleName,omitnil,omitempty" name:"ApproverRoleName"`
}

type ApproverOption struct {
	// 签署方是否可以拒签
	// 
	// <ul><li> **false** : ( 默认)可以拒签</li>
	// <li> **true** :不可以拒签</li></ul>
	NoRefuse *bool `json:"NoRefuse,omitnil,omitempty" name:"NoRefuse"`

	// 签署方是否可以转他人处理
	// 
	// <ul><li> **false** : ( 默认)可以转他人处理</li>
	// <li> **true** :不可以转他人处理</li></ul>
	NoTransfer *bool `json:"NoTransfer,omitnil,omitempty" name:"NoTransfer"`

	// 允许编辑签署人信息（嵌入式使用） 默认true-可以编辑 false-不可以编辑
	CanEditApprover *bool `json:"CanEditApprover,omitnil,omitempty" name:"CanEditApprover"`

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
	// 指定签署人名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 指定签署人手机号，11位数字
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 指定签署人证件类型，ID_CARD-身份证
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 指定签署人证件号码，字母大写
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`
}

type ArchiveDynamicApproverData struct {
	// 签署方唯一编号，一个全局唯一的标识符，不同的流程不会出现冲突。
	// 
	// 可以使用签署方的唯一编号来生成签署链接（也可以通过RecipientId来生成签署链接）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 签署方角色编号，签署方角色编号是用于区分同一个流程中不同签署方的唯一标识。不同的流程会出现同样的签署方角色编号。
	// 
	// 填写控件和签署控件都与特定的角色编号关联。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`
}

// Predefined struct for user
type ArchiveDynamicFlowRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID, 为32位字符串。
	// 
	// 可登录腾讯电子签控制台，[点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type ArchiveDynamicFlowRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID, 为32位字符串。
	// 
	// 可登录腾讯电子签控制台，[点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
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
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "Agent")
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
	Approvers []*ArchiveDynamicApproverData `json:"Approvers,omitnil,omitempty" name:"Approvers"`

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

type AuthInfoDetail struct {
	// 扩展服务类型，和入参一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 扩展服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 授权员工列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasAuthUserList []*HasAuthUser `json:"HasAuthUserList,omitnil,omitempty" name:"HasAuthUserList"`

	// 授权企业列表（企业自动签时，该字段有值）
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasAuthOrganizationList []*HasAuthOrganization `json:"HasAuthOrganizationList,omitnil,omitempty" name:"HasAuthOrganizationList"`

	// 授权员工列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthUserTotal *int64 `json:"AuthUserTotal,omitnil,omitempty" name:"AuthUserTotal"`

	// 授权企业列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthOrganizationTotal *int64 `json:"AuthOrganizationTotal,omitnil,omitempty" name:"AuthOrganizationTotal"`
}

type AuthRecord struct {
	// 经办人姓名。
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 经办人手机号。
	OperatorMobile *string `json:"OperatorMobile,omitnil,omitempty" name:"OperatorMobile"`

	// 认证授权方式：
	// <ul><li> **0**：未选择授权方式（默认值）</li>
	// <li> **1**：上传授权书</li>
	// <li> **2**：法人授权</li>
	// <li> **3**：法人认证</li></ul>
	AuthType *int64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 企业认证授权书审核状态：
	// <ul><li> **0**：未提交授权书（默认值）</li>
	// <li> **1**：审核通过</li>
	// <li> **2**：审核驳回</li>
	// <li> **3**：审核中</li>
	// <li> **4**：AI识别中</li>
	// <li> **5**：客户确认AI信息</li></ul>
	AuditStatus *int64 `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`
}

type AuthorizedUser struct {
	// 电子签系统中的用户id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type AutoSignConfig struct {
	// 自动签开通个人用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 是否回调证书信息:
	// <ul><li>**false**: 不需要(默认)</li>
	// <li>**true**:需要</li></ul>
	CertInfoCallback *bool `json:"CertInfoCallback,omitnil,omitempty" name:"CertInfoCallback"`

	// 是否支持用户自定义签名印章:
	// <ul><li>**false**: 不能自己定义(默认)</li>
	// <li>**true**: 可以自己定义</li></ul>
	UserDefineSeal *bool `json:"UserDefineSeal,omitnil,omitempty" name:"UserDefineSeal"`

	// 回调中是否需要自动签将要使用的印章(签名) 图片的 base64:
	// <ul><li>**false**: 不需要(默认)</li>
	// <li>**true**: 需要</li></ul>
	SealImgCallback *bool `json:"SealImgCallback,omitnil,omitempty" name:"SealImgCallback"`

	// 执行结果的回调URL，该URL仅支持HTTP或HTTPS协议，建议采用HTTPS协议以保证数据传输的安全性。
	// 腾讯电子签服务器将通过POST方式，application/json格式通知执行结果，请确保外网可以正常访问该URL。
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 开通时候的身份验证方式, 取值为：
	// <ul><li>**WEIXINAPP** : 微信人脸识别</li>
	// <li>**INSIGHT** : 慧眼人脸认别</li>
	// <li>**TELECOM** : 运营商三要素验证</li></ul>
	// 注：
	// <ul><li>如果是小程序开通链接，支持传 WEIXINAPP / TELECOM。为空默认 WEIXINAPP</li>
	// <li>如果是 H5 开通链接，支持传 INSIGHT / TELECOM。为空默认 INSIGHT </li></ul>
	VerifyChannels []*string `json:"VerifyChannels,omitnil,omitempty" name:"VerifyChannels"`

	// 设置用户自动签合同的扣费方式。
	// 
	// <ul><li><b>1</b>: (默认)使用合同份额进行扣减</li></ul>
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
	// 若贵方需要在跳转回时通过链接query参数提示开通成功，JumpUrl中的query应携带如下参数：`appendResult=qian`。这样腾讯电子签H5会在跳转回的url后面会添加query参数提示贵方签署成功，例如： qianapp://YOUR_CUSTOM_URL?action=sign&result=success&from=tencent_ess
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`
}

type BillUsageDetail struct {
	// 合同流程ID，为32位字符串。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 合同经办人名称
	// 如果有多个经办人用分号隔开。
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 发起方组织机构名称
	CreateOrganizationName *string `json:"CreateOrganizationName,omitnil,omitempty" name:"CreateOrganizationName"`

	// 合同流程的名称。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 当前合同状态,如下是状态码对应的状态。
	// <ul>
	// <li>**0**: 还没有发起</li>
	// <li>**1**: 等待签署</li>
	// <li>**2**: 部分签署 </li>
	// <li>**3**: 拒签</li>
	// <li>**4**: 已签署 </li>
	// <li>**5**: 已过期 </li>
	// <li>**6**: 已撤销 </li>
	// <li>**7**: 还没有预发起</li>
	// <li>**8**: 等待填写</li>
	// <li>**9**: 部分填写 </li>
	// <li>**10**: 拒填</li>
	// <li>**11**: 已解除</li>
	// </ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

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
type BindEmployeeUserIdWithClientOpenIdRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写UserId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 
	// 通过<a href="https://qian.tencent.com/developers/companyApis/staffs/DescribeIntegrationEmployees" target="_blank">DescribeIntegrationEmployees</a>接口获取，也可登录腾讯电子签控制台查看
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/97cfffabb0caa61df16999cd395d7850.png)
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 员工在贵司业务系统中的唯一身份标识，用于与腾讯电子签账号进行映射，确保在同一企业内不会出现重复。 该标识最大长度为64位字符串，仅支持包含26个英文字母和数字0-9的字符。
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type BindEmployeeUserIdWithClientOpenIdRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写UserId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 
	// 通过<a href="https://qian.tencent.com/developers/companyApis/staffs/DescribeIntegrationEmployees" target="_blank">DescribeIntegrationEmployees</a>接口获取，也可登录腾讯电子签控制台查看
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/97cfffabb0caa61df16999cd395d7850.png)
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 员工在贵司业务系统中的唯一身份标识，用于与腾讯电子签账号进行映射，确保在同一企业内不会出现重复。 该标识最大长度为64位字符串，仅支持包含26个英文字母和数字0-9的字符。
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *BindEmployeeUserIdWithClientOpenIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEmployeeUserIdWithClientOpenIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "UserId")
	delete(f, "OpenId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindEmployeeUserIdWithClientOpenIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindEmployeeUserIdWithClientOpenIdResponseParams struct {
	// 绑定是否成功。
	// <ul><li>**0**：失败</li><li>**1**：成功</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindEmployeeUserIdWithClientOpenIdResponse struct {
	*tchttp.BaseResponse
	Response *BindEmployeeUserIdWithClientOpenIdResponseParams `json:"Response"`
}

func (r *BindEmployeeUserIdWithClientOpenIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEmployeeUserIdWithClientOpenIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallbackInfo struct {
	// 回调url,。请确保回调地址能够接收并处理 HTTP POST 请求，并返回状态码 200 以表示处理正常。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 回调加密key，已废弃
	//
	// Deprecated: Token is deprecated.
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 回调加密key，用于回调消息加解密。
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 回调验签token，用于回调通知校验。
	CallbackToken *string `json:"CallbackToken,omitnil,omitempty" name:"CallbackToken"`
}

type Caller struct {
	// 应用号
	//
	// Deprecated: ApplicationId is deprecated.
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 主机构ID
	//
	// Deprecated: OrganizationId is deprecated.
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 经办人的用户ID，同UserId
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 下属机构ID
	//
	// Deprecated: SubOrganizationId is deprecated.
	SubOrganizationId *string `json:"SubOrganizationId,omitnil,omitempty" name:"SubOrganizationId"`
}

type CancelFailureFlow struct {
	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 撤销失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

// Predefined struct for user
type CancelFlowRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID, 为32位字符串。
	// 
	// 可登录腾讯电子签控制台，[点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 撤销此合同流程的**撤销理由**，最多支持200个字符长度。只能由中文、字母、数字、中文标点和英文标点组成（不支持表情）。
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/f16cf37dbb3a09d6569877f093b92204/channel_ChannelCancelFlow.png)
	CancelMessage *string `json:"CancelMessage,omitnil,omitempty" name:"CancelMessage"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CancelFlowRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID, 为32位字符串。
	// 
	// 可登录腾讯电子签控制台，[点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 撤销此合同流程的**撤销理由**，最多支持200个字符长度。只能由中文、字母、数字、中文标点和英文标点组成（不支持表情）。
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/f16cf37dbb3a09d6569877f093b92204/channel_ChannelCancelFlow.png)
	CancelMessage *string `json:"CancelMessage,omitnil,omitempty" name:"CancelMessage"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 需要取消的签署码ID，为32位字符串。由[创建一码多签签署码](https://qian.tencent.com/developers/companyApis/startFlows/CreateMultiFlowSignQRCode/)返回
	QrCodeId *string `json:"QrCodeId,omitnil,omitempty" name:"QrCodeId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CancelMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 需要取消的签署码ID，为32位字符串。由[创建一码多签签署码](https://qian.tencent.com/developers/companyApis/startFlows/CreateMultiFlowSignQRCode/)返回
	QrCodeId *string `json:"QrCodeId,omitnil,omitempty" name:"QrCodeId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CancelUserAutoSignEnableUrlRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 预撤销链接的用户信息，包含姓名、证件类型、证件号码等信息。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CancelUserAutoSignEnableUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 预撤销链接的用户信息，包含姓名、证件类型、证件号码等信息。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *CancelUserAutoSignEnableUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelUserAutoSignEnableUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelUserAutoSignEnableUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelUserAutoSignEnableUrlResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelUserAutoSignEnableUrlResponse struct {
	*tchttp.BaseResponse
	Response *CancelUserAutoSignEnableUrlResponseParams `json:"Response"`
}

func (r *CancelUserAutoSignEnableUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelUserAutoSignEnableUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// 通知签署方经办人的方式,  有以下途径:
	// <ul><li> **sms** :  (默认)短信</li>
	// <li> **none** : 不通知</li></ul>
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`
}

type Component struct {
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
	// <li> <b>WATERMARK</b> : 水印控件；只能分配给发起方，必须设置ComponentExtra；</li>
	// <li> <b>DISTRICT</b> : 省市区行政区控件，ComponentValue填写省市区行政区字符串内容；</li></ul>
	// 
	// **如果是SignComponent签署控件类型，
	// 需要根据签署人的类型可选的字段为**
	// * 企业方
	// <ul><li> <b>SIGN_SEAL</b> : 签署印章控件；</li>
	// <li> <b>SIGN_DATE</b> : 签署日期控件；</li>
	// <li> <b>SIGN_SIGNATURE</b> : 用户签名控件；</li>
	// <li> <b>SIGN_PAGING_SEAL</b> : 骑缝章；若文件发起，需要对应填充ComponentPosY、ComponentWidth、ComponentHeight</li>
	// <li> <b>SIGN_OPINION</b> : 签署意见控件，用户需要根据配置的签署意见内容，完成对意见内容的确认；</li>
	// <li> <b>SIGN_VIRTUAL_COMBINATION</b> : 签批控件。内部最多组合4个特定控件（SIGN_SIGNATURE，SIGN_DATA,SIGN_MULTI_LINE_TEXT,SIGN_SELECTOR），本身不填充任何文字内容</li>
	// <li> <b>SIGN_MULTI_LINE_TEXT</b> : 多行文本，<font color="red">仅可用在签批控件内部作为组合控件，单独无法使用</font>，常用作批注附言</li>
	// <li> <b>SIGN_SELECTOR</b> : 选择器，<font color="red">仅可用在签批控件内部作为组合控件，单独无法使用</font>，常用作审批意见的选择</li>
	// <li> <b>SIGN_LEGAL_PERSON_SEAL</b> : 企业法定代表人控件。</li></ul>
	// 
	// * 个人方
	// <ul><li> <b>SIGN_DATE</b> : 签署日期控件；</li>
	// <li> <b>SIGN_SIGNATURE</b> : 用户签名控件；</li>
	// <li> <b>SIGN_VIRTUAL_COMBINATION</b> : 签批控件。内部最多组合4个特定控件（SIGN_SIGNATURE，SIGN_DATA,SIGN_MULTI_LINE_TEXT,SIGN_SELECTOR），本身不填充任何文字内容</li>
	// <li> <b>SIGN_MULTI_LINE_TEXT</b> : 多行文本，<font color="red">仅可用在签批控件内部作为组合控件，单独无法使用</font>，常用作批注附言</li>
	// <li> <b>SIGN_SELECTOR</b> : 选择器，<font color="red">仅可用在签批控件内部作为组合控件，单独无法使用</font>，常用作审批意见的选择</li>
	// <li> <b>SIGN_OPINION</b> : 签署意见控件，用户需要根据配置的签署意见内容，完成对意见内容的确认；</li></ul>
	//  
	// 注：` 表单域的控件不能作为印章和签名控件`
	ComponentType *string `json:"ComponentType,omitnil,omitempty" name:"ComponentType"`

	// **在绝对定位方式和关键字定位方式下**，指定控件的高度， 控件高度是指控件在PDF文件中的高度，单位为pt（点）。
	ComponentHeight *float64 `json:"ComponentHeight,omitnil,omitempty" name:"ComponentHeight"`

	// **在绝对定位方式和关键字定位方式下**，指定控件宽度，控件宽度是指控件在PDF文件中的宽度，单位为pt（点）。
	ComponentWidth *float64 `json:"ComponentWidth,omitnil,omitempty" name:"ComponentWidth"`

	// **在绝对定位方式方式下**，指定控件所在PDF文件上的页码
	// **在使用文件发起的情况下**，绝对定位方式的填写控件和签署控件支持使用负数来指定控件在PDF文件上的页码，使用负数时，页码从最后一页开始。例如：ComponentPage设置为-1，即代表在PDF文件的最后一页，以此类推。
	// 
	// 注：
	// 1. 页码编号是从<font color="red">1</font>开始编号的。
	// 2.  <font color="red">页面编号不能超过PDF文件的页码总数</font>。如果指定的页码超过了PDF文件的页码总数，在填写和签署时会出现错误，导致无法正常进行操作。
	ComponentPage *int64 `json:"ComponentPage,omitnil,omitempty" name:"ComponentPage"`

	// **在绝对定位方式和关键字定位方式下**，可以指定控件横向位置的位置，单位为pt（点）。
	ComponentPosX *float64 `json:"ComponentPosX,omitnil,omitempty" name:"ComponentPosX"`

	// **在绝对定位方式和关键字定位方式下**，可以指定控件纵向位置的位置，单位为pt（点）。
	ComponentPosY *float64 `json:"ComponentPosY,omitnil,omitempty" name:"ComponentPosY"`

	// <font color="red">【暂未使用】</font>控件所属文件的序号（取值为：0-N）。 目前单文件的情况下，值一直为0
	FileIndex *int64 `json:"FileIndex,omitnil,omitempty" name:"FileIndex"`

	// 控件生成的方式：
	// <ul><li> <b>NORMAL</b> : 绝对定位控件</li>
	// <li> <b>FIELD</b> : 表单域</li>
	// <li> <b>KEYWORD</b> : 关键字（设置关键字时，请确保PDF原始文件内是关键字以文字形式保存在PDF文件中，不支持对图片内文字进行关键字查找）</li></ul>
	GenerateMode *string `json:"GenerateMode,omitnil,omitempty" name:"GenerateMode"`

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
	// 
	// <a href="https://qcloudimg.tencent-cloud.cn/raw/93178569d07b4d7dbbe0967ae679e35c.png" target="_blank">点击查看ComponentId在模板编辑页面的位置</a>
	ComponentId *string `json:"ComponentId,omitnil,omitempty" name:"ComponentId"`

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

	// **在所有的定位方式下**，控件的扩展参数，为<font color="red">JSON格式</font>，不同类型的控件会有部分非通用参数。
	// 
	// <font color="red">ComponentType为TEXT、MULTI_LINE_TEXT时</font>，支持以下参数：
	// <ul><li> <b>Font</b>：目前只支持黑体、宋体</li>
	// <li> <b>FontSize</b>： 范围12 :72</li>
	// <li> <b>FontAlign</b>： Left/Right/Center，左对齐/居中/右对齐</li>
	// <li> <b>FontColor</b>：字符串类型，格式为RGB颜色数字</li></ul>
	// <b>参数样例</b>：`{"FontColor":"255,0,0","FontSize":12}`
	// 
	// <font color="red">ComponentType为DATE时</font>，支持以下参数：
	// <ul><li> <b>Font</b>：目前只支持黑体、宋体</li>
	// <li> <b>FontSize</b>： 范围12 :72</li></ul>
	// <b>参数样例</b>：`{"FontColor":"255,0,0","FontSize":12}`
	// 
	// <font color="red">ComponentType为WATERMARK时</font>，支持以下参数：
	// <ul><li> <b>Font</b>：目前只支持黑体、宋体</li>
	// <li> <b>FontSize</b>： 范围6 :24</li>
	// <li> <b>Opacity</b>： 透明度，范围0 :1</li>
	// <li> <b>Density</b>： 水印样式，1-宽松，2-标准（默认值），3-密集，</li>
	// <li> <b>SubType</b>： 水印类型：CUSTOM_WATERMARK-自定义内容，PERSON_INFO_WATERMARK-访问者信息</li></ul>
	// <b>参数样例</b>：`"{\"Font\":\"黑体\",\"FontSize\":20,\"Opacity\":0.1,\"Density\":2,\"SubType\":\"PERSON_INFO_WATERMARK\"}"`
	// 
	// <font color="red">ComponentType为FILL_IMAGE时</font>，支持以下参数：
	// <ul><li> <b>NotMakeImageCenter</b>：bool。是否设置图片居中。false：居中（默认）。 true : 不居中</li>
	// <li> <b>FillMethod</b> : int. 填充方式。0-铺满（默认）；1-等比例缩放</li></ul>
	// 
	// <font color="red">ComponentType为SIGN_SIGNATURE类型时</font>，可以通过**ComponentTypeLimit**参数控制签名方式
	// <ul><li> <b>HANDWRITE</b> :  需要实时手写的手写签名</li>
	// <li> <b>HANDWRITTEN_ESIGN</b> : 长效手写签名， 是使用保存到个人中心的印章列表的手写签名(并且包含HANDWRITE)</li>
	// <li> <b>OCR_ESIGN</b> : AI智能识别手写签名</li>
	// <li> <b>ESIGN</b> : 个人印章类型</li>
	// <li> <b>SYSTEM_ESIGN</b> : 系统签名（该类型可以在用户签署时根据用户姓名一键生成一个签名来进行签署）</li>
	// <li> <b>IMG_ESIGN</b> : 图片印章(该类型支持用户在签署将上传的PNG格式的图片作为签名)</li></ul>
	// <b>参考样例</b>：`{"ComponentTypeLimit": ["SYSTEM_ESIGN"]}`
	// 印章的对应关系参考下图
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/ee0498856c060c065628a0c5ba780d6b.jpg)<br><br>
	// 
	// <font color="red">ComponentType为SIGN_SEAL 或者 SIGN_PAGING_SEAL类型时</font>，可以通过**ComponentTypeLimit**参数控制签署方签署时要使用的印章类型，支持指定以下印章类型
	// <ul><li> <b>OFFICIAL</b> :  企业公章</li>
	// <li> <b>CONTRACT</b> : 合同专用章</li>
	// <li> <b>FINANCE</b> : 财务专用章</li>
	// <li> <b>PERSONNEL</b> : 人事专用章</li></ul>
	// <b>参考样例</b>：`{\"ComponentTypeLimit\":[\"PERSONNEL\",\"FINANCE\"]}` 表示改印章签署区,客户需使用人事专用章或财务专用章盖章签署。<br><br>
	// 
	// <font color="red">ComponentType为SIGN_DATE时</font>，支持以下参数：
	// <ul><li> <b>Font</b> :字符串类型目前只支持"黑体"、"宋体"，如果不填默认为"黑体"</li>
	// <li> <b>FontSize</b> : 数字类型，范围6-72，默认值为12</li>
	// <li> <b>FontAlign</b> : 字符串类型，可取Left/Right/Center，对应左对齐/居中/右对齐</li>
	// <li> <b>Format</b> : 字符串类型，日期格式，必须是以下五种之一 “yyyy m d”，”yyyy年m月d日”，”yyyy/m/d”，”yyyy-m-d”，”yyyy.m.d”。</li>
	// <li> <b>Gaps</b> : 字符串类型，仅在Format为“yyyy m d”时起作用，格式为用逗号分开的两个整数，例如”2,2”，两个数字分别是日期格式的前后两个空隙中的空格个数</li></ul>
	// 如果extra参数为空，默认为”yyyy年m月d日”格式的居中日期
	// 特别地，如果extra中Format字段为空或无法被识别，则extra参数会被当作默认值处理（Font，FontSize，Gaps和FontAlign都不会起效）
	// <b>参数样例</b>： ` "{"Format":"yyyy m d","FontSize":12,"Gaps":"2,2", "FontAlign":"Right"}"`
	// 
	// <font color="red">ComponentType为SIGN_SEAL类型时</font>，支持以下参数：
	// <ul><li> <b>PageRanges</b> :PageRange的数组，通过PageRanges属性设置该印章在PDF所有页面上盖章（适用于标书在所有页面盖章的情况）</li></ul>
	// <b>参数样例</b>：` "{"PageRanges":[{"BeginPage":1,"EndPage":-1}]}"`
	// 
	// <font color="red">签署印章透明度功能设置，</font>当ComponentType为SIGN_SIGNATURE、SIGN_SEAL、SIGN_PAGING_SEAL、SIGN_LEGAL_PERSON_SEAL时，可以通过以下参数设置签署印章的透明度：
	// <ul><li> <b>Opacity</b>：印章透明度，支持范围：0.6-1，0.7表示70%的透明度，1表示无透明度</li></ul>
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
	// <font color="red">ComponentType为SIGN_VIRTUAL_COMBINATION时</font>，支持以下参数：
	// <ul>
	// <li><b>Children:</b> 绝对定位模式下，用来指定此签批控件的组合子控件 </li>
	// <b>参数样例</b>：<br>`{"Children":["ComponentId_29","ComponentId_27","ComponentId_28","ComponentId_30"]}`
	// <li><b>ChildrenComponents:</b> 关键字定位模式下，用来指定此签批控件的组合子控件 </li>
	// ChildrenComponent结构体定义:
	// <table border="1">
	//     <thead>
	//         <tr>
	//             <th>字段名称</th>
	//             <th>类型</th>
	//             <th>描述</th>
	//         </tr>
	//     </thead>
	//     <tbody>
	//         <tr>
	//             <td>ComponentType</td>
	//             <td>string</td>
	//             <td>子控件类型-可选值:SIGN_SIGNATURE,SIGN_DATE,SIGN_SELECTOR,SIGN_MULTI_LINE_TEXT</td>
	//         </tr>
	//         <tr>
	//             <td>ComponentName</td>
	//             <td>string</td>
	//             <td>子控件名称</td>
	//         </tr>
	//         <tr>
	//             <td>Placeholder</td>
	//             <td>string</td>
	//             <td>子控件提示语</td>
	//         </tr>
	//         <tr>
	//             <td>ComponentOffsetX</td>
	//             <td>float</td>
	//             <td>控件偏移位置X（相对于父控件（签批控件的ComponentX））</td>
	//         </tr>
	//         <tr>
	//             <td>ComponentOffsetY</td>
	//             <td>float</td>
	//             <td>控件偏移位置Y 相对于父控件（签批控件的ComponentY））</td>
	//         </tr>
	//         <tr>
	//             <td>ComponentWidth</td>
	//             <td>float</td>
	//             <td>控件宽</td>
	//         </tr>
	//         <tr>
	//             <td>ComponentHeight</td>
	//             <td>float</td>
	//             <td>控件高</td>
	//         </tr>
	//         <tr>
	//             <td>ComponentExtra</td>
	//             <td>string</td>
	//             <td>控件的附属信息，根据ComponentType设置</td>
	//         </tr>
	//     </tbody>
	// </table>
	// <b>参数样例</b>：
	//   
	// ```json
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
	// }
	// ```
	// </ul>
	// 
	ComponentExtra *string `json:"ComponentExtra,omitnil,omitempty" name:"ComponentExtra"`

	// **在通过接口拉取控件信息场景下**，为出参参数，此控件是否通过表单域定位方式生成，默认false-不是，**发起合同时候不要填写此字段留空即可**
	IsFormType *bool `json:"IsFormType,omitnil,omitempty" name:"IsFormType"`

	// 控件填充vaule，ComponentType和传入值类型对应关系：
	// <ul><li> <b>TEXT</b> : 文本内容</li>
	// <li> <b>MULTI_LINE_TEXT</b> : 文本内容，可以用  \n 来控制换行位置 </li>
	// <li> <b>CHECK_BOX</b> : true/false</li>
	// <li> <b>FILL_IMAGE、ATTACHMENT</b> : 附件的FileId，需要通过UploadFiles接口上传获取</li>
	// <li> <b>SELECTOR</b> : 选项值</li>
	// <li> <b>DYNAMIC_TABLE</b>  - 传入json格式的表格内容，详见说明：[数据表格](https://qian.tencent.com/developers/company/dynamic_table)</li>
	// <li> <b>DATE</b> : 格式化为：xxxx年xx月xx日（例如2024年05年28日）</li>
	// <li> <b>SIGN_SEAL</b> : 印章ID，于控制台查询获取， [点击查看在控制台上位置](https://qcloudimg.tencent-cloud.cn/raw/f7b0f2ea4a534aada4b893dbf9671eae.png)</li>
	// <li> <b>SIGN_PAGING_SEAL</b> : 可以指定印章ID，于控制台查询获取， [点击查看在控制台上位置](https://qcloudimg.tencent-cloud.cn/raw/f7b0f2ea4a534aada4b893dbf9671eae.png)</li></ul>
	// 
	// 
	// <b>控件值约束说明</b>：
	// <table> <thead> <tr> <th>特殊控件</th> <th>填写约束</th> </tr> </thead> <tbody> <tr> <td>企业全称控件</td> <td>企业名称中文字符中文括号</td> </tr> <tr> <td>统一社会信用代码控件</td> <td>企业注册的统一社会信用代码</td> </tr> <tr> <td>法人名称控件</td> <td>最大50个字符，2到25个汉字或者1到50个字母</td> </tr> <tr> <td>签署意见控件</td> <td>签署意见最大长度为50字符</td> </tr> <tr> <td>签署人手机号控件</td> <td>国内手机号 13,14,15,16,17,18,19号段长度11位</td> </tr> <tr> <td>签署人身份证控件</td> <td>合法的身份证号码检查</td> </tr> <tr> <td>控件名称</td> <td>控件名称最大长度为20字符，不支持表情</td> </tr> <tr> <td>单行文本控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>多行文本控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>勾选框控件</td> <td>选择填字符串true，不选填字符串false</td> </tr> <tr> <td>选择器控件</td> <td>同单行文本控件约束，填写选择值中的字符串</td> </tr> <tr> <td>数字控件</td> <td>请输入有效的数字(可带小数点)</td> </tr> <tr> <td>日期控件</td> <td>格式：yyyy年mm月dd日</td> </tr> <tr> <td>附件控件</td> <td>JPG或PNG图片，上传数量限制，1到6个，最大6个附件，填写上传的资源ID</td> </tr> <tr> <td>图片控件</td> <td>JPG或PNG图片，填写上传的图片资源ID</td> </tr> <tr> <td>邮箱控件</td> <td>有效的邮箱地址, w3c标准</td> </tr> <tr> <td>地址控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>省市区控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>性别控件</td> <td>选择值中的字符串</td> </tr> <tr> <td>学历控件</td> <td>选择值中的字符串</td> </tr> <tr> <td>水印控件</td> <td>水印控件设置为CUSTOM_WATERMARK类型时的水印内容</td> </tr> </tbody> </table>
	// 注：   `部分特殊控件需要在控制台配置模板形式创建`
	ComponentValue *string `json:"ComponentValue,omitnil,omitempty" name:"ComponentValue"`

	// **如果控件是关键字定位方式**，可以对关键字定位出来的区域进行横坐标方向的调整，单位为pt（点）。例如，如果关键字定位出来的区域偏左或偏右，可以通过调整横坐标方向的参数来使控件位置更加准确。
	// 注意： `向左调整设置为负数， 向右调整设置成正数`
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetX *float64 `json:"OffsetX,omitnil,omitempty" name:"OffsetX"`

	// **如果控件是关键字定位方式**，可以对关键字定位出来的区域进行纵坐标方向的调整，单位为pt（点）。例如，如果关键字定位出来的区域偏上或偏下，可以通过调整纵坐标方向的参数来使控件位置更加准确。
	// 注意： `向上调整设置为负数， 向下调整设置成正数`
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetY *float64 `json:"OffsetY,omitnil,omitempty" name:"OffsetY"`

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

	// **web嵌入发起合同场景下**， 是否锁定填写和签署控件值不允许嵌入页面进行编辑
	// <ul><li>false（默认）：不锁定控件值，允许在页面编辑控件值</li>
	// <li>true：锁定控件值，在页面编辑控件值</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockComponentValue *bool `json:"LockComponentValue,omitnil,omitempty" name:"LockComponentValue"`

	// **web嵌入发起合同场景下**，是否禁止移动和删除填写和签署控件
	// <ul><li> <b>false（默认）</b> :不禁止移动和删除控件</li>
	// <li> <b>true</b> : 可以移动和删除控件</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForbidMoveAndDelete *bool `json:"ForbidMoveAndDelete,omitnil,omitempty" name:"ForbidMoveAndDelete"`

	// <font color="red">【暂未使用】</font>日期签署控件的字号，默认为 12
	ComponentDateFontSize *int64 `json:"ComponentDateFontSize,omitnil,omitempty" name:"ComponentDateFontSize"`

	// <font color="red">【暂未使用】</font>第三方应用集成平台模板控件 ID 标识
	ChannelComponentId *string `json:"ChannelComponentId,omitnil,omitempty" name:"ChannelComponentId"`

	// <font color="red">【暂未使用】</font>第三方应用集成中子客企业控件来源。
	// <ul><li> <b>0</b> :平台指定；</li>
	// <li> <b>1</b> :用户自定义</li></ul>
	ChannelComponentSource *uint64 `json:"ChannelComponentSource,omitnil,omitempty" name:"ChannelComponentSource"`
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
	// <li> <b>PERSONNEL</b> : 人事专用章</li></ul>
	// 
	// **注：`限制印章控件或骑缝章控件情况下,仅本企业签署方可以指定具体印章（通过传递ComponentValue,支持多个),他方企业签署人只能限制类型.若同时指定了印章类型和印章Id,以印章Id为主,印章类型会被忽略`**
	// 
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
type CreateBatchCancelFlowUrlRequestParams struct {
	// 执行本接口操作的员工信息。
	// <br/>注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 需要执行撤回的流程(合同)的编号列表，最多100个.
	// 列表中的流程(合同)编号不要重复.
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 代理企业和员工的信息。
	// <br/>在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CreateBatchCancelFlowUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// <br/>注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 需要执行撤回的流程(合同)的编号列表，最多100个.
	// 列表中的流程(合同)编号不要重复.
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 代理企业和员工的信息。
	// <br/>在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
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
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchCancelFlowUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchCancelFlowUrlResponseParams struct {
	// 批量撤回签署流程链接
	BatchCancelFlowUrl *string `json:"BatchCancelFlowUrl,omitnil,omitempty" name:"BatchCancelFlowUrl"`

	// 签署流程撤回失败信息
	// 数组里边的错误原因与传进来的FlowIds一一对应,如果是空字符串则标识没有出错
	FailMessages []*string `json:"FailMessages,omitnil,omitempty" name:"FailMessages"`

	// 签署连接过期时间字符串：年月日-时分秒
	// 
	// 例如:2023-07-28 17:25:59
	UrlExpireOn *string `json:"UrlExpireOn,omitnil,omitempty" name:"UrlExpireOn"`

	// 批量撤销任务编号，为32位字符串，可用于[查询批量撤销签署流程任务结果](https://qian.tencent.com/developers/companyApis/operateFlows/CreateBatchCancelFlowUrl) 或关联[批量撤销任务结果回调](https://qian.tencent.com/developers/company/callback_types_contracts_sign#%E4%B9%9D-%E6%89%B9%E9%87%8F%E6%92%A4%E9%94%80%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83)
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateBatchInitOrganizationUrlRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 初始化操作类型
	// <ul>
	// <li>CREATE_SEAL : 创建印章</li>
	// <li>AUTH_JOIN_ORGANIZATION_GROUP : 加入集团企业</li>
	// <li>OPEN_AUTO_SIGN :开通企业自动签署</li>
	// <li>PARTNER_AUTO_SIGN_AUTH :合作方企业授权自动签</li>
	// </ul>
	OperateTypes []*string `json:"OperateTypes,omitnil,omitempty" name:"OperateTypes"`

	// 批量操作的企业Id列表，最大支持50个
	OrganizationIds []*string `json:"OrganizationIds,omitnil,omitempty" name:"OrganizationIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 被授权的合作方企业在电子签的企业电子签账号，当操作类型包含 PARTNER_AUTO_SIGN_AUTH （合作方企业授权自动签）时必传。
	// 
	// 企业电子签账号可在[电子签的网页端](https://qian.tencent.com/console/company-settings/company-center) ，于企业设置-企业信息菜单栏下复制获取。
	// 
	// ![企业电子签账号](https://qcloudimg.tencent-cloud.cn/raw/4e6b30ee92f00671f7f1c5bd127c27db.png)
	AuthorizedOrganizationId *string `json:"AuthorizedOrganizationId,omitnil,omitempty" name:"AuthorizedOrganizationId"`
}

type CreateBatchInitOrganizationUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 初始化操作类型
	// <ul>
	// <li>CREATE_SEAL : 创建印章</li>
	// <li>AUTH_JOIN_ORGANIZATION_GROUP : 加入集团企业</li>
	// <li>OPEN_AUTO_SIGN :开通企业自动签署</li>
	// <li>PARTNER_AUTO_SIGN_AUTH :合作方企业授权自动签</li>
	// </ul>
	OperateTypes []*string `json:"OperateTypes,omitnil,omitempty" name:"OperateTypes"`

	// 批量操作的企业Id列表，最大支持50个
	OrganizationIds []*string `json:"OrganizationIds,omitnil,omitempty" name:"OrganizationIds"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 被授权的合作方企业在电子签的企业电子签账号，当操作类型包含 PARTNER_AUTO_SIGN_AUTH （合作方企业授权自动签）时必传。
	// 
	// 企业电子签账号可在[电子签的网页端](https://qian.tencent.com/console/company-settings/company-center) ，于企业设置-企业信息菜单栏下复制获取。
	// 
	// ![企业电子签账号](https://qcloudimg.tencent-cloud.cn/raw/4e6b30ee92f00671f7f1c5bd127c27db.png)
	AuthorizedOrganizationId *string `json:"AuthorizedOrganizationId,omitnil,omitempty" name:"AuthorizedOrganizationId"`
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
	delete(f, "Operator")
	delete(f, "OperateTypes")
	delete(f, "OrganizationIds")
	delete(f, "Agent")
	delete(f, "AuthorizedOrganizationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchInitOrganizationUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchInitOrganizationUrlResponseParams struct {
	// 小程序路径
	MiniAppPath *string `json:"MiniAppPath,omitnil,omitempty" name:"MiniAppPath"`

	// 操作长链
	OperateLongUrl *string `json:"OperateLongUrl,omitnil,omitempty" name:"OperateLongUrl"`

	// 操作短链
	OperateShortUrl *string `json:"OperateShortUrl,omitnil,omitempty" name:"OperateShortUrl"`

	// 操作二维码
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
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 组织机构超管姓名。 在注册流程中，必须是超管本人进行操作。
	// 此参数需要跟[创建企业批量认证链接](https://qian.tencent.com/developers/companyApis/organizations/CreateBatchOrganizationRegistrationTasks)中 AdminName 保持一致。
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 组织机构超管手机号。 在注册流程中，必须是超管本人进行操作。此参数需要跟[创建企业批量认证链接](https://qian.tencent.com/developers/companyApis/organizations/CreateBatchOrganizationRegistrationTasks)中 Admin Mobile保持一致。
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 企业批量认证链接的子任务 SubTaskId，该 SubTaskId 是通过接口 查询企业批量认证链接 DescribeBatchOrganizationRegistrationUrls 获得。此参数需与超管个人三要素（AdminName，AdminMobile，AdminIdCardNumber）配合使用。若 SubTaskId 不属于传入的超级管理员，将进行筛选。
	SubTaskIds []*string `json:"SubTaskIds,omitnil,omitempty" name:"SubTaskIds"`

	// 组织机构超管证件类型支持以下类型
	// - ID_CARD : 中国大陆居民身份证 (默认值)
	// -  HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证
	// - HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)
	// 此参数需要跟[创建企业批量认证链接](https://qian.tencent.com/developers/companyApis/organizations/CreateBatchOrganizationRegistrationTasks)中 AdminIdCardType保持一致。
	AdminIdCardType *string `json:"AdminIdCardType,omitnil,omitempty" name:"AdminIdCardType"`

	// 组织机构超管证件号。 在注册流程中，必须是超管本人进行操作。此参数需要跟[创建企业批量认证链接](https://qian.tencent.com/developers/companyApis/organizations/CreateBatchOrganizationRegistrationTasks)中 AdminIdCardNumber保持一致。
	AdminIdCardNumber *string `json:"AdminIdCardNumber,omitnil,omitempty" name:"AdminIdCardNumber"`

	// 要跳转的链接类型<ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li><li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li><li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li><li>**QR_CODE**： 跳转电子签小程序的http_url的二维码形式,  可以在页面展示适合此类型</li></ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

type CreateBatchOrganizationAuthorizationUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 组织机构超管姓名。 在注册流程中，必须是超管本人进行操作。
	// 此参数需要跟[创建企业批量认证链接](https://qian.tencent.com/developers/companyApis/organizations/CreateBatchOrganizationRegistrationTasks)中 AdminName 保持一致。
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 组织机构超管手机号。 在注册流程中，必须是超管本人进行操作。此参数需要跟[创建企业批量认证链接](https://qian.tencent.com/developers/companyApis/organizations/CreateBatchOrganizationRegistrationTasks)中 Admin Mobile保持一致。
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 企业批量认证链接的子任务 SubTaskId，该 SubTaskId 是通过接口 查询企业批量认证链接 DescribeBatchOrganizationRegistrationUrls 获得。此参数需与超管个人三要素（AdminName，AdminMobile，AdminIdCardNumber）配合使用。若 SubTaskId 不属于传入的超级管理员，将进行筛选。
	SubTaskIds []*string `json:"SubTaskIds,omitnil,omitempty" name:"SubTaskIds"`

	// 组织机构超管证件类型支持以下类型
	// - ID_CARD : 中国大陆居民身份证 (默认值)
	// -  HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证
	// - HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)
	// 此参数需要跟[创建企业批量认证链接](https://qian.tencent.com/developers/companyApis/organizations/CreateBatchOrganizationRegistrationTasks)中 AdminIdCardType保持一致。
	AdminIdCardType *string `json:"AdminIdCardType,omitnil,omitempty" name:"AdminIdCardType"`

	// 组织机构超管证件号。 在注册流程中，必须是超管本人进行操作。此参数需要跟[创建企业批量认证链接](https://qian.tencent.com/developers/companyApis/organizations/CreateBatchOrganizationRegistrationTasks)中 AdminIdCardNumber保持一致。
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
	delete(f, "Operator")
	delete(f, "AdminName")
	delete(f, "AdminMobile")
	delete(f, "Agent")
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
	// 批量企业注册链接-单链接包含多条认证流，根据Endpoint的不同设置，返回不同的链接地址。失效时间：7天
	// 跳转链接, 链接的有效期根据企业,员工状态和终端等有区别, 可以参考下表
	// <table> <thead> <tr> <th>Endpoint</th> <th>示例</th> <th>链接有效期限</th> </tr> </thead>  <tbody>
	//  <tr> <td>HTTP</td> <td>https://res.ess.tencent.cn/cdn/h5-activity-dev/jump-mp.html?to=AUTHORIZATION_ENTERPRISE_FOR_BATCH_SUBMIT&shortKey=yDCHHURDfBxSB2rj2Bfa</td> <td>7天</td> </tr> 
	// <tr> <td>HTTP_SHORT_URL</td> <td>https://test.essurl.cn/8gDKUBAWK8</td> <td>7天</td> </tr> 
	// <tr> <td>APP</td> <td>pages/guide/index?to=AUTHORIZATION_ENTERPRISE_FOR_BATCH_SUBMIT&shortKey=yDCHpURDfR6iEkdpsDde</td> <td>7天</td> </tr><tr> <td>QR_CODE</td> <td>https://dyn.test.ess.tencent.cn/imgs/qrcode_urls/authorization_enterprise_for_batch_submit/yDCHHUUckpbdauq9UEjnoFDCCumAMmv1.png</td> <td>7天</td> </tr> </tbody> </table>
	// 注： 
	// `1.创建的链接应避免被转义，如：&被转义为\u0026；如使用Postman请求后，请选择响应类型为 JSON，否则链接将被转义`
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
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 组织机构注册信息。
	// 一次最多支持10条认证流
	RegistrationOrganizations []*RegistrationOrganizationInfo `json:"RegistrationOrganizations,omitnil,omitempty" name:"RegistrationOrganizations"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要生成链接的类型, 可以选择的值如下: 
	// 
	// <ul>
	// <li>(默认)PC: 生成PC端的链接</li>
	// <li>SHORT_URL: H5跳转到电子签小程序链接的短链形式, 一般用于发送短信中带的链接, 打开后进入腾讯电子签小程序</li>
	// <li>APP：生成小程序跳转链接</li>
	// <li>H5：生成H5跳转长链接</li>
	// <li>SHORT_H5：生成H5跳转短链</li>
	// </ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

type CreateBatchOrganizationRegistrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 组织机构注册信息。
	// 一次最多支持10条认证流
	RegistrationOrganizations []*RegistrationOrganizationInfo `json:"RegistrationOrganizations,omitnil,omitempty" name:"RegistrationOrganizations"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 要生成链接的类型, 可以选择的值如下: 
	// 
	// <ul>
	// <li>(默认)PC: 生成PC端的链接</li>
	// <li>SHORT_URL: H5跳转到电子签小程序链接的短链形式, 一般用于发送短信中带的链接, 打开后进入腾讯电子签小程序</li>
	// <li>APP：生成小程序跳转链接</li>
	// <li>H5：生成H5跳转长链接</li>
	// <li>SHORT_H5：生成H5跳转短链</li>
	// </ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
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
	delete(f, "Operator")
	delete(f, "RegistrationOrganizations")
	delete(f, "Agent")
	delete(f, "Endpoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchOrganizationRegistrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchOrganizationRegistrationTasksResponseParams struct {
	// 生成注册链接的任务Id，
	// 根据这个id， 调用DescribeBatchOrganizationRegistrationUrls 获取生成的链接，进入认证流程
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 批量生成企业认证链接的详细错误信息，
	// 顺序与输入参数保持一致。
	// 若企业认证均成功生成，则不返回错误信息；
	// 若存在任何错误，则返回具体的错误描述。
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
type CreateBatchQuickSignUrlRequestParams struct {
	// 批量签署的流程签署人，其中姓名(ApproverName)、参与人类型(ApproverType)必传，手机号(ApproverMobile)和证件信息(ApproverIdCardType、ApproverIdCardNumber)可任选一种或全部传入。
	// <ul>
	// <li>若为个人参与方：ApproverType=1</li>
	// <li>若为企业参与方：ApproverType=0。同时 OrganizationName 参数需传入参与方企业名称。 </li>
	// </ul>
	// 注:
	// `1. 暂不支持签署人拖动签署控件功能，以及签批控件。`
	// `2. 当需要通过短信验证码签署时，手机号ApproverMobile需要与发起合同时填写的用户手机号一致。`
	FlowApproverInfo *FlowCreateApprover `json:"FlowApproverInfo,omitnil,omitempty" name:"FlowApproverInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId(子企业的组织ID)为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

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

	// <b>只有在生成H5签署链接的情形下</b>（ 如调用<a href="https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateFlowSignUrl" target="_blank">获取H5签署链接</a>、<a href="https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateBatchQuickSignUrl" target="_blank">获取H5批量签署链接</a>等接口），该配置才会生效。  您可以指定H5签署视频核身的意图配置，选择问答模式或点头模式的语音文本。 
	// 
	//  注意： 
	// 1. 视频认证为<b>白名单功能，使用前请联系对接的客户经理沟通</b>。
	// 2. 使用视频认证时，<b>生成H5签署链接的时候必须将签署认证方式指定为人脸</b>（即ApproverSignTypes设置成人脸签署）。 
	// 3. 签署完成后，可以通过<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelDescribeSignFaceVideo" target="_blank">查询签署认证人脸视频</a>获取到当时的视频。
	Intention *Intention `json:"Intention,omitnil,omitempty" name:"Intention"`

	// 缓存签署人信息。在H5签署链接动态领取场景，首次填写后，选择缓存签署人信息，在下次签署人点击领取链接时，会自动将个人信息（姓名、身份证号、手机号）填入，否则需要每次手动填写。
	// 
	// 注: `若参与方为企业员工时，暂不支持对参与方信息进行缓存`
	CacheApproverInfo *bool `json:"CacheApproverInfo,omitnil,omitempty" name:"CacheApproverInfo"`
}

type CreateBatchQuickSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 批量签署的流程签署人，其中姓名(ApproverName)、参与人类型(ApproverType)必传，手机号(ApproverMobile)和证件信息(ApproverIdCardType、ApproverIdCardNumber)可任选一种或全部传入。
	// <ul>
	// <li>若为个人参与方：ApproverType=1</li>
	// <li>若为企业参与方：ApproverType=0。同时 OrganizationName 参数需传入参与方企业名称。 </li>
	// </ul>
	// 注:
	// `1. 暂不支持签署人拖动签署控件功能，以及签批控件。`
	// `2. 当需要通过短信验证码签署时，手机号ApproverMobile需要与发起合同时填写的用户手机号一致。`
	FlowApproverInfo *FlowCreateApprover `json:"FlowApproverInfo,omitnil,omitempty" name:"FlowApproverInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId(子企业的组织ID)为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

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

	// <b>只有在生成H5签署链接的情形下</b>（ 如调用<a href="https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateFlowSignUrl" target="_blank">获取H5签署链接</a>、<a href="https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateBatchQuickSignUrl" target="_blank">获取H5批量签署链接</a>等接口），该配置才会生效。  您可以指定H5签署视频核身的意图配置，选择问答模式或点头模式的语音文本。 
	// 
	//  注意： 
	// 1. 视频认证为<b>白名单功能，使用前请联系对接的客户经理沟通</b>。
	// 2. 使用视频认证时，<b>生成H5签署链接的时候必须将签署认证方式指定为人脸</b>（即ApproverSignTypes设置成人脸签署）。 
	// 3. 签署完成后，可以通过<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelDescribeSignFaceVideo" target="_blank">查询签署认证人脸视频</a>获取到当时的视频。
	Intention *Intention `json:"Intention,omitnil,omitempty" name:"Intention"`

	// 缓存签署人信息。在H5签署链接动态领取场景，首次填写后，选择缓存签署人信息，在下次签署人点击领取链接时，会自动将个人信息（姓名、身份证号、手机号）填入，否则需要每次手动填写。
	// 
	// 注: `若参与方为企业员工时，暂不支持对参与方信息进行缓存`
	CacheApproverInfo *bool `json:"CacheApproverInfo,omitnil,omitempty" name:"CacheApproverInfo"`
}

func (r *CreateBatchQuickSignUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchQuickSignUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowApproverInfo")
	delete(f, "Agent")
	delete(f, "Operator")
	delete(f, "FlowIds")
	delete(f, "FlowGroupId")
	delete(f, "JumpUrl")
	delete(f, "SignatureTypes")
	delete(f, "ApproverSignTypes")
	delete(f, "SignTypeSelector")
	delete(f, "FlowBatchUrlInfo")
	delete(f, "Intention")
	delete(f, "CacheApproverInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchQuickSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchQuickSignUrlResponseParams struct {
	// 签署人签署链接信息
	FlowApproverUrlInfo *FlowApproverUrlInfo `json:"FlowApproverUrlInfo,omitnil,omitempty" name:"FlowApproverUrlInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBatchQuickSignUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchQuickSignUrlResponseParams `json:"Response"`
}

func (r *CreateBatchQuickSignUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchQuickSignUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchSignUrlRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：`请确保和合同中填入的一致`, `除动态签署人场景外，此参数必填`
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此业务通知方。
	// 
	// 注：`请确保和合同中填入的一致,  若无法保持一致，请确保在发起和生成批量签署链接时传入相同的参与方证件信息`，`除动态签署人场景外，此参数必填`
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	// 
	// 注：`请确保和合同中填入的一致`
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	// 
	// 注：`请确保和合同中填入的一致`
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 通知用户方式：
	// <ul>
	// <li>**NONE** : 不通知（默认）</li>
	// <li>**SMS** : 短信通知（发送短信通知到Mobile参数所传的手机号）</li>
	// </ul>
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// 批量签署的合同流程ID数组。
	// 注: `在调用此接口时，请确保合同流程均为本企业发起，且合同数量不超过100个。`
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 目标签署人的企业名称，签署人如果是企业员工身份，需要传此参数。
	// 
	// 注：
	// <ul>
	// <li>请确认该名称与企业营业执照中注册的名称一致。</li>
	// <li>如果名称中包含英文括号()，请使用中文括号（）代替。</li>
	// <li>请确保此企业已完成腾讯电子签企业认证。</li>
	// </ul>
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 是否直接跳转至合同内容页面进行签署
	// <ul>
	// <li>**false**: 会跳转至批量合同流程的列表,  点击需要批量签署合同后进入合同内容页面进行签署(默认)</li>
	// <li>**true**: 跳过合同流程列表, 直接进入合同内容页面进行签署</li>
	// </ul>
	JumpToDetail *bool `json:"JumpToDetail,omitnil,omitempty" name:"JumpToDetail"`

	// 批量签署合同相关信息，指定合同和签署方的信息，用于补充动态签署人。	
	FlowBatchUrlInfo *FlowBatchUrlInfo `json:"FlowBatchUrlInfo,omitnil,omitempty" name:"FlowBatchUrlInfo"`

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
}

type CreateBatchSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：`请确保和合同中填入的一致`, `除动态签署人场景外，此参数必填`
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此业务通知方。
	// 
	// 注：`请确保和合同中填入的一致,  若无法保持一致，请确保在发起和生成批量签署链接时传入相同的参与方证件信息`，`除动态签署人场景外，此参数必填`
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	// 
	// 注：`请确保和合同中填入的一致`
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	// 
	// 注：`请确保和合同中填入的一致`
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 通知用户方式：
	// <ul>
	// <li>**NONE** : 不通知（默认）</li>
	// <li>**SMS** : 短信通知（发送短信通知到Mobile参数所传的手机号）</li>
	// </ul>
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// 批量签署的合同流程ID数组。
	// 注: `在调用此接口时，请确保合同流程均为本企业发起，且合同数量不超过100个。`
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 目标签署人的企业名称，签署人如果是企业员工身份，需要传此参数。
	// 
	// 注：
	// <ul>
	// <li>请确认该名称与企业营业执照中注册的名称一致。</li>
	// <li>如果名称中包含英文括号()，请使用中文括号（）代替。</li>
	// <li>请确保此企业已完成腾讯电子签企业认证。</li>
	// </ul>
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 是否直接跳转至合同内容页面进行签署
	// <ul>
	// <li>**false**: 会跳转至批量合同流程的列表,  点击需要批量签署合同后进入合同内容页面进行签署(默认)</li>
	// <li>**true**: 跳过合同流程列表, 直接进入合同内容页面进行签署</li>
	// </ul>
	JumpToDetail *bool `json:"JumpToDetail,omitnil,omitempty" name:"JumpToDetail"`

	// 批量签署合同相关信息，指定合同和签署方的信息，用于补充动态签署人。	
	FlowBatchUrlInfo *FlowBatchUrlInfo `json:"FlowBatchUrlInfo,omitnil,omitempty" name:"FlowBatchUrlInfo"`

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
}

func (r *CreateBatchSignUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchSignUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Name")
	delete(f, "Mobile")
	delete(f, "Agent")
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "NotifyType")
	delete(f, "FlowIds")
	delete(f, "OrganizationName")
	delete(f, "JumpToDetail")
	delete(f, "FlowBatchUrlInfo")
	delete(f, "AutoJumpBack")
	delete(f, "UrlUseEnv")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchSignUrlResponseParams struct {
	// 批量签署链接，以短链形式返回，短链的有效期参考回参中的 ExpiredTime。
	// 
	// 注: 
	// 1. 非小程序和APP集成使用
	// 2. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	SignUrl *string `json:"SignUrl,omitnil,omitempty" name:"SignUrl"`

	// 链接过期时间以 Unix 时间戳格式表示，默认生成链接时间起，往后7天有效期。过期后短链将失效，无法打开。
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

type CreateBatchSignUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchSignUrlResponseParams `json:"Response"`
}

func (r *CreateBatchSignUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchSignUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConvertTaskApiRequestParams struct {
	// 需要进行转换的资源文件类型
	// 支持的文件类型如下：
	// <ul><li>doc</li>
	// <li>docx</li>
	// <li>xls</li>
	// <li>xlsx</li>
	// <li>jpg</li>
	// <li>jpeg</li>
	// <li>png</li>
	// <li>html</li>
	// <li>bmp</li>
	// <li>txt</li></ul>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 需要进行转换操作的文件资源名称，带资源后缀名。
	// 
	// 注:  `资源名称长度限制为256个字符`
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 需要进行转换操作的文件资源Id，通过<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles" target="_blank">UploadFiles</a>接口获取文件资源Id。
	// 
	// 注:  `目前，此接口仅支持单个文件进行转换。`
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
}

type CreateConvertTaskApiRequest struct {
	*tchttp.BaseRequest
	
	// 需要进行转换的资源文件类型
	// 支持的文件类型如下：
	// <ul><li>doc</li>
	// <li>docx</li>
	// <li>xls</li>
	// <li>xlsx</li>
	// <li>jpg</li>
	// <li>jpeg</li>
	// <li>png</li>
	// <li>html</li>
	// <li>bmp</li>
	// <li>txt</li></ul>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 需要进行转换操作的文件资源名称，带资源后缀名。
	// 
	// 注:  `资源名称长度限制为256个字符`
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 需要进行转换操作的文件资源Id，通过<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles" target="_blank">UploadFiles</a>接口获取文件资源Id。
	// 
	// 注:  `目前，此接口仅支持单个文件进行转换。`
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
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
	// 接口返回的文件转换任务Id，可以调用接口<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/GetTaskResultApi" target="_blank">查询转换任务状态</a>获取转换任务的状态和转换后的文件资源Id。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 调用方用户信息，userId 必填。支持填入集团子公司经办人 userId代发合同。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 用户配置的合同模板ID，会基于此模板创建合同文档，为32位字符串。
	// 
	// [点击查看模板Id在控制台上的位置](https://qcloudimg.tencent-cloud.cn/raw/253071cc2f7becb063c7cf71b37b7861.png)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 合同流程ID，为32位字符串。
	// 此接口的合同流程ID需要由[创建签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow)接口创建得到。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 文件名列表，单个文件名最大长度200个字符，暂时仅支持单文件发起。设置后流程对应的文件名称当前设置的值。
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// 电子文档的填写控件的填充内容。具体方式可以参考[FormField](https://qian.tencent.com/developers/companyApis/dataTypes/#formfield)结构体的定义。
	// <ul>
	// <li>支持自动签传递印章，可通过指定自动签控件id，指定印章id来完成</li>
	// <li>附件控件支持传入图片、文件资源id，并将内容合成到合同文件中。支持的文件类型有doc、docx、xls、xlsx、html、jpg、jpeg、png、bmp、txt、pdf。需要注意如果传入的资源类型都是图片类型，图片资源会放置在合同文件的末尾，如果传入的资源有非图片类型资源，会将资源放置在附件控件所在页面的下一页。</li>
	// </ul>
	// 注：只有在控制台编辑模板时，<font color="red">归属给发起方</font>的填写控件（如下图）才能在创建文档的时候进行内容填充。
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/a54a76a58c454593d06d8e9883ecc9b3.png)
	FormFields []*FormField `json:"FormFields,omitnil,omitempty" name:"FormFields"`

	// 是否为预览模式，取值如下：
	// <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li> 
	// <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。 <font color="red">注意： 以预览模式创建的合同仅供查看，因此参与方无法进行签署操作</font> </li></ul>
	// 注: `当使用的模板中存在动态表格控件时，预览结果中没有动态表格的填写内容，动态表格合成完后会触发文档合成完成的回调通知`
	NeedPreview *bool `json:"NeedPreview,omitnil,omitempty" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子。</li></ul>
	// 注: `1.此参数在NeedPreview 为true时有效`
	// `2.动态表格控件不支持H5链接方式预览`
	PreviewType *int64 `json:"PreviewType,omitnil,omitempty" name:"PreviewType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 已废弃字段，客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

type CreateDocumentRequest struct {
	*tchttp.BaseRequest
	
	// 调用方用户信息，userId 必填。支持填入集团子公司经办人 userId代发合同。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 用户配置的合同模板ID，会基于此模板创建合同文档，为32位字符串。
	// 
	// [点击查看模板Id在控制台上的位置](https://qcloudimg.tencent-cloud.cn/raw/253071cc2f7becb063c7cf71b37b7861.png)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 合同流程ID，为32位字符串。
	// 此接口的合同流程ID需要由[创建签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow)接口创建得到。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 文件名列表，单个文件名最大长度200个字符，暂时仅支持单文件发起。设置后流程对应的文件名称当前设置的值。
	FileNames []*string `json:"FileNames,omitnil,omitempty" name:"FileNames"`

	// 电子文档的填写控件的填充内容。具体方式可以参考[FormField](https://qian.tencent.com/developers/companyApis/dataTypes/#formfield)结构体的定义。
	// <ul>
	// <li>支持自动签传递印章，可通过指定自动签控件id，指定印章id来完成</li>
	// <li>附件控件支持传入图片、文件资源id，并将内容合成到合同文件中。支持的文件类型有doc、docx、xls、xlsx、html、jpg、jpeg、png、bmp、txt、pdf。需要注意如果传入的资源类型都是图片类型，图片资源会放置在合同文件的末尾，如果传入的资源有非图片类型资源，会将资源放置在附件控件所在页面的下一页。</li>
	// </ul>
	// 注：只有在控制台编辑模板时，<font color="red">归属给发起方</font>的填写控件（如下图）才能在创建文档的时候进行内容填充。
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/a54a76a58c454593d06d8e9883ecc9b3.png)
	FormFields []*FormField `json:"FormFields,omitnil,omitempty" name:"FormFields"`

	// 是否为预览模式，取值如下：
	// <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li> 
	// <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。 <font color="red">注意： 以预览模式创建的合同仅供查看，因此参与方无法进行签署操作</font> </li></ul>
	// 注: `当使用的模板中存在动态表格控件时，预览结果中没有动态表格的填写内容，动态表格合成完后会触发文档合成完成的回调通知`
	NeedPreview *bool `json:"NeedPreview,omitnil,omitempty" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子。</li></ul>
	// 注: `1.此参数在NeedPreview 为true时有效`
	// `2.动态表格控件不支持H5链接方式预览`
	PreviewType *int64 `json:"PreviewType,omitnil,omitempty" name:"PreviewType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 已废弃字段，客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
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
	delete(f, "TemplateId")
	delete(f, "FlowId")
	delete(f, "FileNames")
	delete(f, "FormFields")
	delete(f, "NeedPreview")
	delete(f, "PreviewType")
	delete(f, "Agent")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocumentResponseParams struct {
	// 合同流程的底层电子文档ID，为32位字符串。
	// 
	// 注:
	// 后续需用同样的FlowId再次调用[发起签署流程](https://qian.tencent.com/developers/companyApis/startFlows/StartFlow)，合同才能进入签署环节
	DocumentId *string `json:"DocumentId,omitnil,omitempty" name:"DocumentId"`

	// 合同预览链接URL。
	// 
	// 注: `1.如果是预览模式(即NeedPreview设置为true)时, 才会有此预览链接URL`
	// `2.当使用的模板中存在动态表格控件时，预览结果中没有动态表格的填写内容`
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewFileUrl *string `json:"PreviewFileUrl,omitnil,omitempty" name:"PreviewFileUrl"`

	// 签署方信息，如角色ID、角色名称等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Approvers []*ApproverItem `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateDynamicFlowApproverRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。支持填入集团子公司经办人 userId 代发合同。注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 合同流程的参与方列表，最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的ApproverInfo结构体。如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序，请确保列表中参与人的顺序符合实际签署顺序。
	Approvers []*ApproverInfo `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 代理企业和员工的信息。在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：<ul><li> **E_PRESCRIPTION_AUTO_SIGN**：电子处方单（医疗自动签）  </li><li> **OTHER** :  通用场景</li></ul>注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil,omitempty" name:"AutoSignScene"`
}

type CreateDynamicFlowApproverRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。支持填入集团子公司经办人 userId 代发合同。注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 合同流程的参与方列表，最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的ApproverInfo结构体。如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序，请确保列表中参与人的顺序符合实际签署顺序。
	Approvers []*ApproverInfo `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 代理企业和员工的信息。在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：<ul><li> **E_PRESCRIPTION_AUTO_SIGN**：电子处方单（医疗自动签）  </li><li> **OTHER** :  通用场景</li></ul>注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil,omitempty" name:"AutoSignScene"`
}

func (r *CreateDynamicFlowApproverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDynamicFlowApproverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "Approvers")
	delete(f, "Agent")
	delete(f, "AutoSignScene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDynamicFlowApproverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDynamicFlowApproverResponseParams struct {
	// 合同流程ID，为32位字符串
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 补充动态合同签署人的结果数组
	DynamicFlowApproverList []*DynamicFlowApproverResult `json:"DynamicFlowApproverList,omitnil,omitempty" name:"DynamicFlowApproverList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDynamicFlowApproverResponse struct {
	*tchttp.BaseResponse
	Response *CreateDynamicFlowApproverResponseParams `json:"Response"`
}

func (r *CreateDynamicFlowApproverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDynamicFlowApproverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmbedWebUrlRequestParams struct {
	// 执行本接口操作的员工信息。
	// <br/>注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// WEB嵌入资源类型，支持以下类型
	// <ul><li>CREATE_SEAL: 生成创建印章的嵌入页面</li>
	// <li>CREATE_TEMPLATE：生成创建模板的嵌入页面</li>
	// <li>MODIFY_TEMPLATE：生成编辑模板的嵌入页面</li>
	// <li>PREVIEW_TEMPLATE：生成预览模板的嵌入页面</li>
	// <li>PREVIEW_SEAL_LIST：生成预览印章列表的嵌入页面</li>
	// <li>PREVIEW_SEAL_DETAIL：生成预览印章详情的嵌入页面</li>
	// <li>EXTEND_SERVICE：生成拓展服务的嵌入页面</li>
	// <li>PREVIEW_FLOW：生成预览合同的嵌入页面（支持移动端）</li>
	// <li>PREVIEW_FLOW_DETAIL：生成查看合同详情的嵌入页面（仅支持PC端）</li></ul>
	EmbedType *string `json:"EmbedType,omitnil,omitempty" name:"EmbedType"`

	// WEB嵌入的业务资源ID
	// 
	// 当EmbedType取值
	// <ul>
	// <li>为PREVIEW_SEAL_DETAIL，必填，取值为印章id。</li>
	// <li>为CREATE_TEMPLATE，非必填，取值为资源id。*资源Id获取可使用接口[上传文件](https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles)*</li>
	// <li>为MODIFY_TEMPLATE，PREVIEW_TEMPLATE，必填，取值为模板id。</li>
	// <li>为PREVIEW_FLOW，PREVIEW_FLOW_DETAIL，必填，取值为合同id。</li>
	// </ul>
	// 
	// 注意：
	//  1. CREATE_TEMPLATE中的BusinessId仅支持PDF文件类型， 如果您的文件不是PDF， 请使用接口[创建文件转换任务
	// ](https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi) 和[查询转换任务状态](https://qian.tencent.com/developers/companyApis/templatesAndFiles/GetTaskResultApi) 来进行转换成PDF资源。
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 代理企业和员工的信息。
	// <br/>在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 抄送方信息
	Reviewer *ReviewerInfo `json:"Reviewer,omitnil,omitempty" name:"Reviewer"`

	// 个性化参数，用于控制页面展示内容
	Option *EmbedUrlOption `json:"Option,omitnil,omitempty" name:"Option"`

	// <ul> <li>目前仅支持EmbedType=CREATE_TEMPLATE时传入</li> <li>指定后，创建，编辑，删除模板时，回调都会携带该userData</li> <li>支持的格式：json字符串的BASE64编码字符串</li> <li>示例：<ul>                  <li>json字符串：{"ComeFrom":"xxx"}，BASE64编码：eyJDb21lRnJvbSI6Inh4eCJ9</li>                  <li>eyJDb21lRnJvbSI6Inh4eCJ9，为符合要求的userData数据格式</li> </ul> </li> </ul>
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

type CreateEmbedWebUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// <br/>注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// WEB嵌入资源类型，支持以下类型
	// <ul><li>CREATE_SEAL: 生成创建印章的嵌入页面</li>
	// <li>CREATE_TEMPLATE：生成创建模板的嵌入页面</li>
	// <li>MODIFY_TEMPLATE：生成编辑模板的嵌入页面</li>
	// <li>PREVIEW_TEMPLATE：生成预览模板的嵌入页面</li>
	// <li>PREVIEW_SEAL_LIST：生成预览印章列表的嵌入页面</li>
	// <li>PREVIEW_SEAL_DETAIL：生成预览印章详情的嵌入页面</li>
	// <li>EXTEND_SERVICE：生成拓展服务的嵌入页面</li>
	// <li>PREVIEW_FLOW：生成预览合同的嵌入页面（支持移动端）</li>
	// <li>PREVIEW_FLOW_DETAIL：生成查看合同详情的嵌入页面（仅支持PC端）</li></ul>
	EmbedType *string `json:"EmbedType,omitnil,omitempty" name:"EmbedType"`

	// WEB嵌入的业务资源ID
	// 
	// 当EmbedType取值
	// <ul>
	// <li>为PREVIEW_SEAL_DETAIL，必填，取值为印章id。</li>
	// <li>为CREATE_TEMPLATE，非必填，取值为资源id。*资源Id获取可使用接口[上传文件](https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles)*</li>
	// <li>为MODIFY_TEMPLATE，PREVIEW_TEMPLATE，必填，取值为模板id。</li>
	// <li>为PREVIEW_FLOW，PREVIEW_FLOW_DETAIL，必填，取值为合同id。</li>
	// </ul>
	// 
	// 注意：
	//  1. CREATE_TEMPLATE中的BusinessId仅支持PDF文件类型， 如果您的文件不是PDF， 请使用接口[创建文件转换任务
	// ](https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi) 和[查询转换任务状态](https://qian.tencent.com/developers/companyApis/templatesAndFiles/GetTaskResultApi) 来进行转换成PDF资源。
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 代理企业和员工的信息。
	// <br/>在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 抄送方信息
	Reviewer *ReviewerInfo `json:"Reviewer,omitnil,omitempty" name:"Reviewer"`

	// 个性化参数，用于控制页面展示内容
	Option *EmbedUrlOption `json:"Option,omitnil,omitempty" name:"Option"`

	// <ul> <li>目前仅支持EmbedType=CREATE_TEMPLATE时传入</li> <li>指定后，创建，编辑，删除模板时，回调都会携带该userData</li> <li>支持的格式：json字符串的BASE64编码字符串</li> <li>示例：<ul>                  <li>json字符串：{"ComeFrom":"xxx"}，BASE64编码：eyJDb21lRnJvbSI6Inh4eCJ9</li>                  <li>eyJDb21lRnJvbSI6Inh4eCJ9，为符合要求的userData数据格式</li> </ul> </li> </ul>
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

func (r *CreateEmbedWebUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmbedWebUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "EmbedType")
	delete(f, "BusinessId")
	delete(f, "Agent")
	delete(f, "Reviewer")
	delete(f, "Option")
	delete(f, "UserData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmbedWebUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmbedWebUrlResponseParams struct {
	// 嵌入的web链接，有效期：5分钟
	// <br/>EmbedType=PREVIEW_CC_FLOW，该url为h5链接
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEmbedWebUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmbedWebUrlResponseParams `json:"Response"`
}

func (r *CreateEmbedWebUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmbedWebUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmployeeQualificationSealQrCodeRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。 支持填入集团子公司经办人 userId 代发合同。  注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 提示信息，扫码后此信息会展示给扫描用户，用来提示用户授权操作的目的，会在授权界面下面的位置展示。
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/8436ffd78c20605e6b133ff4bc4d2ac7.png)
	HintText *string `json:"HintText,omitnil,omitempty" name:"HintText"`
}

type CreateEmployeeQualificationSealQrCodeRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。 支持填入集团子公司经办人 userId 代发合同。  注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
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
	delete(f, "Operator")
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
type CreateExtendedServiceAuthInfosRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 本企业员工的id，需要已实名，正常在职员工
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 取值
	// <ul><li>OPEN_SERVER_SIGN：企业自动签</li>
	// <li>BATCH_SIGN：批量签署</li>
	// </ul>
	ExtendServiceType *string `json:"ExtendServiceType,omitnil,omitempty" name:"ExtendServiceType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CreateExtendedServiceAuthInfosRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 本企业员工的id，需要已实名，正常在职员工
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 取值
	// <ul><li>OPEN_SERVER_SIGN：企业自动签</li>
	// <li>BATCH_SIGN：批量签署</li>
	// </ul>
	ExtendServiceType *string `json:"ExtendServiceType,omitnil,omitempty" name:"ExtendServiceType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *CreateExtendedServiceAuthInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExtendedServiceAuthInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "UserIds")
	delete(f, "ExtendServiceType")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExtendedServiceAuthInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExtendedServiceAuthInfosResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateExtendedServiceAuthInfosResponse struct {
	*tchttp.BaseResponse
	Response *CreateExtendedServiceAuthInfosResponseParams `json:"Response"`
}

func (r *CreateExtendedServiceAuthInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExtendedServiceAuthInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowApproversRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 补充签署环节签署候选人信息。
	// 
	// 注：` 如果发起方指定的补充签署人是企业微信签署人（ApproverSource=WEWORKAPP），则需要提供企业微信UserId进行补充； 如果不指定，则使用姓名和手机号进行补充。`
	Approvers []*FillApproverInfo `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 合同流程ID，为32位字符串。
	// - 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// - 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	// - <font color="red">不建议继续使用</font>，请使用<a href="https://qian.tencent.com/developers/companyApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId来指定需要补充的合同id
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署人信息补充方式
	// 
	// <ul><li>**0**: <font color="red">或签合同</font>添加签署候选人，或签支持一个节点传多个签署人，不传值默认或签。
	// 注: `或签只支持企业签署方`</li>
	// <li>**1**: <font color="red">动态签署人合同</font>的添加签署候选人，支持企业或个人签署方。</li></ul>
	FillApproverType *int64 `json:"FillApproverType,omitnil,omitempty" name:"FillApproverType"`

	// 在可定制的企业微信通知中，发起人可以根据具体需求进行自定义设置。
	Initiator *string `json:"Initiator,omitnil,omitempty" name:"Initiator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程组的组ID, 在合同流程组场景下，生成合同流程组的签署链接时需要赋值
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`
}

type CreateFlowApproversRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 补充签署环节签署候选人信息。
	// 
	// 注：` 如果发起方指定的补充签署人是企业微信签署人（ApproverSource=WEWORKAPP），则需要提供企业微信UserId进行补充； 如果不指定，则使用姓名和手机号进行补充。`
	Approvers []*FillApproverInfo `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 合同流程ID，为32位字符串。
	// - 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// - 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	// - <font color="red">不建议继续使用</font>，请使用<a href="https://qian.tencent.com/developers/companyApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId来指定需要补充的合同id
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署人信息补充方式
	// 
	// <ul><li>**0**: <font color="red">或签合同</font>添加签署候选人，或签支持一个节点传多个签署人，不传值默认或签。
	// 注: `或签只支持企业签署方`</li>
	// <li>**1**: <font color="red">动态签署人合同</font>的添加签署候选人，支持企业或个人签署方。</li></ul>
	FillApproverType *int64 `json:"FillApproverType,omitnil,omitempty" name:"FillApproverType"`

	// 在可定制的企业微信通知中，发起人可以根据具体需求进行自定义设置。
	Initiator *string `json:"Initiator,omitnil,omitempty" name:"Initiator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程组的组ID, 在合同流程组场景下，生成合同流程组的签署链接时需要赋值
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`
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
	delete(f, "Approvers")
	delete(f, "FlowId")
	delete(f, "FillApproverType")
	delete(f, "Initiator")
	delete(f, "Agent")
	delete(f, "FlowGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowApproversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowApproversResponseParams struct {
	// 批量补充签署人时，补充失败的报错说明
	// 
	// 注:`目前仅补充动态签署人时会返回补充失败的原因`
	// 注意：此字段可能返回 null，表示取不到有效值。
	FillError []*FillError `json:"FillError,omitnil,omitempty" name:"FillError"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateFlowBlockchainEvidenceUrlRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 链接/二维码的有效截止时间，格式为unix时间戳。最长不超过 2099年12月31日（4102415999）。
	// 默认值为有效期为当前时间后7天。
	ExpiredOn *uint64 `json:"ExpiredOn,omitnil,omitempty" name:"ExpiredOn"`
}

type CreateFlowBlockchainEvidenceUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

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
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "Agent")
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
type CreateFlowByFilesRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自定义的合同流程的名称，长度不能超过200个字符，只能由中文汉字、中文标点、英文字母、阿拉伯数字、空格、小括号、中括号、中划线、下划线以及（,）、（;）、（.）、(&)、（+）组成。
	// 
	// 该名称还将用于合同签署完成后文件下载的默认文件名称。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 合同流程的参与方列表，最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的ApproverInfo结构体。
	// 
	// 如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序，请确保列表中参与人的顺序符合实际签署顺序。
	Approvers []*ApproverInfo `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 本合同流程需包含的PDF文件资源编号列表，通过[UploadFiles](https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles)接口获取PDF文件资源编号。
	// 
	// 注:  `目前，此接口仅支持单个文件发起。`
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil,omitempty" name:"FlowDescription"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	// 如果用户已经在控制台创建了自定义合同类型，可以将这里的类型名称传入。 如果没有创建，我们会自动给发起方公司创建此自定义合同类型。
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/36582cea03ae6a2559894844942b5d5c.png)
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 模板或者合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li> 单行文本控件      </li>
	// <li> 多行文本控件      </li>
	// <li> 勾选框控件        </li>
	// <li> 数字控件          </li>
	// <li> 图片控件          </li>
	// <li> 水印控件          </li>
	// <li> 动态表格等填写控件</li></ul>
	Components []*Component `json:"Components,omitnil,omitempty" name:"Components"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// <b>注</b>
	// 1. 抄送人名单中可以包括自然人以及本企业的员工。
	// 2. 请确保抄送人列表中的成员不与任何签署人重复。
	CcInfos []*CcInfo `json:"CcInfos,omitnil,omitempty" name:"CcInfos"`

	// 可以设置以下时间节点来给抄送人发送短信通知来查看合同内容：
	// <ul><li> **0**：合同发起时通知（默认值）</li>
	// <li> **1**：签署完成后通知</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil,omitempty" name:"CcNotifyType"`

	// 是否为预览模式，取值如下：
	// <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li>
	// <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。</li></ul>
	NeedPreview *bool `json:"NeedPreview,omitnil,omitempty" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子</li></ul>
	// 注: `此参数在NeedPreview 为true时有效`
	PreviewType *int64 `json:"PreviewType,omitnil,omitempty" name:"PreviewType"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	Unordered *bool `json:"Unordered,omitnil,omitempty" name:"Unordered"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的[回调通知](https://qian.tencent.com/developers/company/callback_types_v2)模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 合同到期提醒时间，为Unix标准时间戳（秒）格式，支持的范围是从发起时间开始到后10年内。
	// 
	// 到达提醒时间后，腾讯电子签会短信通知发起方企业合同提醒，可用于处理合同到期事务，如合同续签等事宜。
	RemindedOn *int64 `json:"RemindedOn,omitnil,omitempty" name:"RemindedOn"`

	// 指定个人签署方查看合同的校验方式
	// <ul><li>   **VerifyCheck**  :（默认）人脸识别,人脸识别后才能合同内容 </li>
	// <li>   **MobileCheck**  :  手机号验证, 用户手机号和参与方手机号（ApproverMobile）相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）</li></ul>
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil,omitempty" name:"ApproverVerifyType"`

	// 签署方签署控件（印章/签名等）的生成方式：
	// <ul><li> **0**：在合同流程发起时，由发起人指定签署方的签署控件的位置和数量。</li>
	// <li> **1**：签署方在签署时自行添加签署控件，可以拖动位置和控制数量。</li></ul>
	SignBeanTag *int64 `json:"SignBeanTag,omitnil,omitempty" name:"SignBeanTag"`

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
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/628f0928cac15d2e3bfa6088f53f5998.png)
	// 
	CustomShowMap *string `json:"CustomShowMap,omitnil,omitempty" name:"CustomShowMap"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：电子处方单（医疗自动签）  </li><li> **OTHER** :  通用场景</li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil,omitempty" name:"AutoSignScene"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil,omitempty" name:"NeedSignReview"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下：  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li></ul>效果如下:![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"FlowDisplayType,omitnil,omitempty" name:"FlowDisplayType"`

	// 是否开启动态签署合同：
	// <ul><li> **true**：开启动态签署合同，可在签署过程中追加签署人（必须满足：1，发起方企业开启了模块化计费能力；2，发起方企业在企业应用管理中开启了动态签署人2.0能力）    。</li>
	// <li> **false**：不开启动态签署合同。</li></ul>
	OpenDynamicSignFlow *bool `json:"OpenDynamicSignFlow,omitnil,omitempty" name:"OpenDynamicSignFlow"`
}

type CreateFlowByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自定义的合同流程的名称，长度不能超过200个字符，只能由中文汉字、中文标点、英文字母、阿拉伯数字、空格、小括号、中括号、中划线、下划线以及（,）、（;）、（.）、(&)、（+）组成。
	// 
	// 该名称还将用于合同签署完成后文件下载的默认文件名称。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 合同流程的参与方列表，最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的ApproverInfo结构体。
	// 
	// 如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序，请确保列表中参与人的顺序符合实际签署顺序。
	Approvers []*ApproverInfo `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 本合同流程需包含的PDF文件资源编号列表，通过[UploadFiles](https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles)接口获取PDF文件资源编号。
	// 
	// 注:  `目前，此接口仅支持单个文件发起。`
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil,omitempty" name:"FlowDescription"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	// 如果用户已经在控制台创建了自定义合同类型，可以将这里的类型名称传入。 如果没有创建，我们会自动给发起方公司创建此自定义合同类型。
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/36582cea03ae6a2559894844942b5d5c.png)
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 模板或者合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li> 单行文本控件      </li>
	// <li> 多行文本控件      </li>
	// <li> 勾选框控件        </li>
	// <li> 数字控件          </li>
	// <li> 图片控件          </li>
	// <li> 水印控件          </li>
	// <li> 动态表格等填写控件</li></ul>
	Components []*Component `json:"Components,omitnil,omitempty" name:"Components"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// <b>注</b>
	// 1. 抄送人名单中可以包括自然人以及本企业的员工。
	// 2. 请确保抄送人列表中的成员不与任何签署人重复。
	CcInfos []*CcInfo `json:"CcInfos,omitnil,omitempty" name:"CcInfos"`

	// 可以设置以下时间节点来给抄送人发送短信通知来查看合同内容：
	// <ul><li> **0**：合同发起时通知（默认值）</li>
	// <li> **1**：签署完成后通知</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil,omitempty" name:"CcNotifyType"`

	// 是否为预览模式，取值如下：
	// <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li>
	// <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。</li></ul>
	NeedPreview *bool `json:"NeedPreview,omitnil,omitempty" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子</li></ul>
	// 注: `此参数在NeedPreview 为true时有效`
	PreviewType *int64 `json:"PreviewType,omitnil,omitempty" name:"PreviewType"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	Unordered *bool `json:"Unordered,omitnil,omitempty" name:"Unordered"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的[回调通知](https://qian.tencent.com/developers/company/callback_types_v2)模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 合同到期提醒时间，为Unix标准时间戳（秒）格式，支持的范围是从发起时间开始到后10年内。
	// 
	// 到达提醒时间后，腾讯电子签会短信通知发起方企业合同提醒，可用于处理合同到期事务，如合同续签等事宜。
	RemindedOn *int64 `json:"RemindedOn,omitnil,omitempty" name:"RemindedOn"`

	// 指定个人签署方查看合同的校验方式
	// <ul><li>   **VerifyCheck**  :（默认）人脸识别,人脸识别后才能合同内容 </li>
	// <li>   **MobileCheck**  :  手机号验证, 用户手机号和参与方手机号（ApproverMobile）相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）</li></ul>
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil,omitempty" name:"ApproverVerifyType"`

	// 签署方签署控件（印章/签名等）的生成方式：
	// <ul><li> **0**：在合同流程发起时，由发起人指定签署方的签署控件的位置和数量。</li>
	// <li> **1**：签署方在签署时自行添加签署控件，可以拖动位置和控制数量。</li></ul>
	SignBeanTag *int64 `json:"SignBeanTag,omitnil,omitempty" name:"SignBeanTag"`

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
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/628f0928cac15d2e3bfa6088f53f5998.png)
	// 
	CustomShowMap *string `json:"CustomShowMap,omitnil,omitempty" name:"CustomShowMap"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：电子处方单（医疗自动签）  </li><li> **OTHER** :  通用场景</li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil,omitempty" name:"AutoSignScene"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil,omitempty" name:"NeedSignReview"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下：  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li></ul>效果如下:![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"FlowDisplayType,omitnil,omitempty" name:"FlowDisplayType"`

	// 是否开启动态签署合同：
	// <ul><li> **true**：开启动态签署合同，可在签署过程中追加签署人（必须满足：1，发起方企业开启了模块化计费能力；2，发起方企业在企业应用管理中开启了动态签署人2.0能力）    。</li>
	// <li> **false**：不开启动态签署合同。</li></ul>
	OpenDynamicSignFlow *bool `json:"OpenDynamicSignFlow,omitnil,omitempty" name:"OpenDynamicSignFlow"`
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
	delete(f, "FlowDescription")
	delete(f, "FlowType")
	delete(f, "Components")
	delete(f, "CcInfos")
	delete(f, "CcNotifyType")
	delete(f, "NeedPreview")
	delete(f, "PreviewType")
	delete(f, "Deadline")
	delete(f, "Unordered")
	delete(f, "UserData")
	delete(f, "RemindedOn")
	delete(f, "ApproverVerifyType")
	delete(f, "SignBeanTag")
	delete(f, "CustomShowMap")
	delete(f, "Agent")
	delete(f, "AutoSignScene")
	delete(f, "NeedSignReview")
	delete(f, "FlowDisplayType")
	delete(f, "OpenDynamicSignFlow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowByFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowByFilesResponseParams struct {
	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 
	// 注: 如果是预览模式(即NeedPreview设置为true)时, 此处不会有值返回。
	// 
	// [点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 合同预览链接URL。
	// 
	// 注：如果是预览模式(即NeedPreview设置为true)时, 才会有此预览链接URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`

	// 签署方信息，如角色ID、角色名称等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Approvers []*ApproverItem `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定申请的报告类型，可选类型如下：
	// <ul><li> **0** :合同签署报告（默认）</li>
	// <li> **1** :公证处核验报告</li></ul>
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`
}

type CreateFlowEvidenceReportRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定申请的报告类型，可选类型如下：
	// <ul><li> **0** :合同签署报告（默认）</li>
	// <li> **1** :公证处核验报告</li></ul>
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`
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
	delete(f, "Agent")
	delete(f, "ReportType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowEvidenceReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowEvidenceReportResponseParams struct {
	// 出证报告 ID，可用于<a href="https://qian.tencent.com/developers/companyApis/certificate/DescribeFlowEvidenceReport" target="_blank">获取出证报告任务执行结果</a>查询出证任务结果和出证PDF的下载URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// 出证任务执行的状态, 状态含义如下：
	// 
	// <ul><li>**EvidenceStatusExecuting**：  出证任务在执行中</li>
	// <li>**EvidenceStatusSuccess**：  出证任务执行成功</li>
	// <li>**EvidenceStatusFailed** ： 出征任务执行失败</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 此字段已经废除,不再使用.
	// 出证的PDF下载地址请调用DescribeChannelFlowEvidenceReport接口获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ReportUrl is deprecated.
	ReportUrl *string `json:"ReportUrl,omitnil,omitempty" name:"ReportUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateFlowGroupByFilesRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同（流程）组名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil,omitempty" name:"FlowGroupName"`

	// 合同（流程）组的子合同信息，支持2-50个子合同
	FlowGroupInfos []*FlowGroupInfo `json:"FlowGroupInfos,omitnil,omitempty" name:"FlowGroupInfos"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同（流程）组的配置项信息。
	// 其中包括：
	// <ul>
	// <li>是否通知本企业签署方</li>
	// <li>是否通知其他签署方</li>
	// </ul>
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil,omitempty" name:"FlowGroupOptions"`
}

type CreateFlowGroupByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同（流程）组名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil,omitempty" name:"FlowGroupName"`

	// 合同（流程）组的子合同信息，支持2-50个子合同
	FlowGroupInfos []*FlowGroupInfo `json:"FlowGroupInfos,omitnil,omitempty" name:"FlowGroupInfos"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同（流程）组的配置项信息。
	// 其中包括：
	// <ul>
	// <li>是否通知本企业签署方</li>
	// <li>是否通知其他签署方</li>
	// </ul>
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil,omitempty" name:"FlowGroupOptions"`
}

func (r *CreateFlowGroupByFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowGroupByFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowGroupName")
	delete(f, "FlowGroupInfos")
	delete(f, "Agent")
	delete(f, "FlowGroupOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowGroupByFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowGroupByFilesResponseParams struct {
	// 合同(流程)组的合同组Id
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 合同(流程)组中子合同列表.
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 合同组签署方信息。
	Approvers []*FlowGroupApprovers `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFlowGroupByFilesResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowGroupByFilesResponseParams `json:"Response"`
}

func (r *CreateFlowGroupByFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowGroupByFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowGroupByTemplatesRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同（流程）组名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil,omitempty" name:"FlowGroupName"`

	// 合同（流程）组的子合同信息，支持2-50个子合同
	FlowGroupInfos []*FlowGroupInfo `json:"FlowGroupInfos,omitnil,omitempty" name:"FlowGroupInfos"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同（流程）组的配置项信息。
	// 其中包括：
	// <ul>
	// <li>是否通知本企业签署方</li>
	// <li>是否通知其他签署方</li>
	// </ul>
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil,omitempty" name:"FlowGroupOptions"`
}

type CreateFlowGroupByTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同（流程）组名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil,omitempty" name:"FlowGroupName"`

	// 合同（流程）组的子合同信息，支持2-50个子合同
	FlowGroupInfos []*FlowGroupInfo `json:"FlowGroupInfos,omitnil,omitempty" name:"FlowGroupInfos"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同（流程）组的配置项信息。
	// 其中包括：
	// <ul>
	// <li>是否通知本企业签署方</li>
	// <li>是否通知其他签署方</li>
	// </ul>
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil,omitempty" name:"FlowGroupOptions"`
}

func (r *CreateFlowGroupByTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowGroupByTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowGroupName")
	delete(f, "FlowGroupInfos")
	delete(f, "Agent")
	delete(f, "FlowGroupOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowGroupByTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowGroupByTemplatesResponseParams struct {
	// 合同(流程)组的合同组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 合同(流程)组中子合同列表.
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 合同组签署人信息。
	Approvers []*FlowGroupApprovers `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFlowGroupByTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowGroupByTemplatesResponseParams `json:"Response"`
}

func (r *CreateFlowGroupByTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowGroupByTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowGroupSignReviewRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同(流程)组的合同组Id，为32位字符串，通过接口[通过多文件创建合同组签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowGroupByFiles) 或[通过多模板创建合同组签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowGroupByTemplates)创建合同组签署流程时返回。
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

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 审核不通过的原因，该字段的字符串长度不超过200个字符。
	// 
	// 注：`当审核类型（ReviewType）为审核拒绝（REJECT）或拒签（SIGN_REJECT）时，审核结果原因字段必须填写`
	ReviewMessage *string `json:"ReviewMessage,omitnil,omitempty" name:"ReviewMessage"`
}

type CreateFlowGroupSignReviewRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同(流程)组的合同组Id，为32位字符串，通过接口[通过多文件创建合同组签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowGroupByFiles) 或[通过多模板创建合同组签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowGroupByTemplates)创建合同组签署流程时返回。
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

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

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
	delete(f, "Operator")
	delete(f, "FlowGroupId")
	delete(f, "ReviewType")
	delete(f, "ApproverInfo")
	delete(f, "Agent")
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
	// 是否允许修改发起合同时确认弹窗的合同信息（合同名称、合同类型、签署截止时间），若不允许编辑，则表单字段将被禁止输入。
	// <br/>true：允许编辑<br/>false：不允许编辑（默认值）<br/>
	CanEditFlow *bool `json:"CanEditFlow,omitnil,omitempty" name:"CanEditFlow"`

	// 是否允许编辑模板控件
	// <br/>true:允许编辑模板控件信息
	// <br/>false:不允许编辑模板控件信息（默认值）
	// <br/>
	CanEditFormField *bool `json:"CanEditFormField,omitnil,omitempty" name:"CanEditFormField"`

	// 发起页面隐藏合同名称展示
	// <br/>true:发起页面隐藏合同名称展示
	// <br/>false:发起页面不隐藏合同名称展示（默认值）
	// <br/>
	HideShowFlowName *bool `json:"HideShowFlowName,omitnil,omitempty" name:"HideShowFlowName"`

	// 发起页面隐藏合同类型展示
	// <br/>true:发起页面隐藏合同类型展示
	// <br/>false:发起页面不隐藏合同类型展示（默认值）
	// <br/>
	HideShowFlowType *bool `json:"HideShowFlowType,omitnil,omitempty" name:"HideShowFlowType"`

	// 发起页面隐藏合同截止日期展示
	// <br/>true:发起页面隐藏合同截止日期展示
	// <br/>false:发起页面不隐藏合同截止日期展示（默认值）
	// <br/>
	HideShowDeadline *bool `json:"HideShowDeadline,omitnil,omitempty" name:"HideShowDeadline"`

	// 发起页面允许跳过添加签署人环节
	// <br/>true:发起页面允许跳过添加签署人环节
	// <br/>false:发起页面不允许跳过添加签署人环节（默认值）
	// <br/>
	CanSkipAddApprover *bool `json:"CanSkipAddApprover,omitnil,omitempty" name:"CanSkipAddApprover"`

	// 文件发起页面跳过文件上传步骤
	// <br/>true:文件发起页面跳过文件上传步骤
	// <br/>false:文件发起页面不跳过文件上传步骤（默认值）
	// <br/>
	SkipUploadFile *bool `json:"SkipUploadFile,omitnil,omitempty" name:"SkipUploadFile"`

	// 禁止编辑填写控件
	// <br/>true:禁止编辑填写控件
	// <br/>false:允许编辑填写控件（默认值）
	// <br/>
	ForbidEditFillComponent *bool `json:"ForbidEditFillComponent,omitnil,omitempty" name:"ForbidEditFillComponent"`

	// 定制化发起合同弹窗的描述信息，描述信息最长500字符
	CustomCreateFlowDescription *string `json:"CustomCreateFlowDescription,omitnil,omitempty" name:"CustomCreateFlowDescription"`

	//   禁止添加签署方，若为true则在发起流程的可嵌入页面隐藏“添加签署人按钮”
	ForbidAddApprover *bool `json:"ForbidAddApprover,omitnil,omitempty" name:"ForbidAddApprover"`

	//   禁止设置设置签署流程属性 (顺序、合同签署认证方式等)，若为true则在发起流程的可嵌入页面隐藏签署流程设置面板
	ForbidEditFlowProperties *bool `json:"ForbidEditFlowProperties,omitnil,omitempty" name:"ForbidEditFlowProperties"`

	// 在发起流程的可嵌入页面要隐藏的控件列表，和 ShowComponentTypes 参数 只能二选一使用，具体的控件类型如下
	// <ul><li>SIGN_SIGNATURE : 个人签名/印章</li>
	// <li>SIGN_SEAL : 企业印章</li>
	// <li>SIGN_PAGING_SEAL : 骑缝章</li>
	// <li>SIGN_LEGAL_PERSON_SEAL : 法定代表人章</li>
	// <li>SIGN_APPROVE : 签批</li>
	// <li>SIGN_OPINION : 签署意见</li>
	// <li>BUSI-FULL-NAME  : 企业全称</li>
	// <li>BUSI-CREDIT-CODE : 统一社会信用代码</li>
	// <li>BUSI-LEGAL-NAME : 法人/经营者姓名</li>
	// <li>PERSONAL-NAME : 签署人姓名</li>
	// <li>PERSONAL-MOBILE : 签署人手机号</li>
	// <li>PERSONAL-IDCARD-TYPE : 签署人证件类型</li>
	// <li>PERSONAL-IDCARD : 签署人证件号</li>
	// <li>TEXT : 单行文本</li>
	// <li>MULTI_LINE_TEXT : 多行文本</li>
	// <li>CHECK_BOX : 勾选框</li>
	// <li>SELECTOR : 选择器</li>
	// <li>DIGIT : 数字</li>
	// <li>DATE : 日期</li>
	// <li>FILL_IMAGE : 图片</li>
	// <li>ATTACHMENT : 附件</li>
	// <li>EMAIL : 邮箱</li>
	// <li>LOCATION : 地址</li>
	// <li>EDUCATION : 学历</li>
	// <li>GENDER : 性别</li>
	// <li>DISTRICT : 省市区</li></ul>
	HideComponentTypes []*string `json:"HideComponentTypes,omitnil,omitempty" name:"HideComponentTypes"`

	// 在发起流程的可嵌入页面要显示的控件列表，和 HideComponentTypes 参数 只能二选一使用，具体的控件类型如下
	// <ul><li>SIGN_SIGNATURE : 个人签名/印章</li>
	// <li>SIGN_SEAL : 企业印章</li>
	// <li>SIGN_PAGING_SEAL : 骑缝章</li>
	// <li>SIGN_LEGAL_PERSON_SEAL : 法定代表人章</li>
	// <li>SIGN_APPROVE : 签批</li>
	// <li>SIGN_OPINION : 签署意见</li>
	// <li>BUSI-FULL-NAME  : 企业全称</li>
	// <li>BUSI-CREDIT-CODE : 统一社会信用代码</li>
	// <li>BUSI-LEGAL-NAME : 法人/经营者姓名</li>
	// <li>PERSONAL-NAME : 签署人姓名</li>
	// <li>PERSONAL-MOBILE : 签署人手机号</li>
	// <li>PERSONAL-IDCARD-TYPE : 签署人证件类型</li>
	// <li>PERSONAL-IDCARD : 签署人证件号</li>
	// <li>TEXT : 单行文本</li>
	// <li>MULTI_LINE_TEXT : 多行文本</li>
	// <li>CHECK_BOX : 勾选框</li>
	// <li>SELECTOR : 选择器</li>
	// <li>DIGIT : 数字</li>
	// <li>DATE : 日期</li>
	// <li>FILL_IMAGE : 图片</li>
	// <li>ATTACHMENT : 附件</li>
	// <li>EMAIL : 邮箱</li>
	// <li>LOCATION : 地址</li>
	// <li>EDUCATION : 学历</li>
	// <li>GENDER : 性别</li>
	// <li>DISTRICT : 省市区</li></ul>
	ShowComponentTypes []*string `json:"ShowComponentTypes,omitnil,omitempty" name:"ShowComponentTypes"`

	// 发起流程的可嵌入页面结果页配置
	ResultPageConfig []*CreateResultPageConfig `json:"ResultPageConfig,omitnil,omitempty" name:"ResultPageConfig"`
}

// Predefined struct for user
type CreateFlowRemindsRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 需执行催办的签署流程ID数组，最多包含100个。
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CreateFlowRemindsRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 需执行催办的签署流程ID数组，最多包含100个。
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *CreateFlowRemindsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowRemindsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowRemindsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowRemindsResponseParams struct {
	// 合同催办结果的详细信息列表。
	RemindFlowRecords []*RemindFlowRecords `json:"RemindFlowRecords,omitnil,omitempty" name:"RemindFlowRecords"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFlowRemindsResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowRemindsResponseParams `json:"Response"`
}

func (r *CreateFlowRemindsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowRemindsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自定义的合同流程的名称，长度不能超过200个字符，只能由中文汉字、中文标点、英文字母、阿拉伯数字、空格、小括号、中括号、中划线、下划线以及（,）、（;）、（.）、(&)、（+）组成。
	// 
	// 该名称还将用于合同签署完成后文件下载的默认文件名称。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 合同流程的参与方列表，最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的ApproverInfo结构体。
	// 
	// 注:  
	// <font color="red" > <b> 在发起流程时，需要保证 approver 中的顺序与模板定义顺序一致，否则会发起失败。
	// 例如，如果模板中定义的第一个参与人是个人用户，第二个参与人是企业员工，则在 approver 中传参时，第一个也必须是个人用户，第二个参与人必须是企业员工。</b></font>
	// 
	// [点击查看模板参与人顺序定义位置](https://qcloudimg.tencent-cloud.cn/raw/d14457b48cc66b29262ccb9d7b3ed556.png)
	Approvers []*FlowCreateApprover `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil,omitempty" name:"FlowDescription"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	// 此合同类型需要跟模板配置的合同类型保持一致。
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 已经废弃字段，客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	DeadLine *int64 `json:"DeadLine,omitnil,omitempty" name:"DeadLine"`

	// 合同到期提醒时间，为Unix标准时间戳（秒）格式，支持的范围是从发起时间开始到后10年内。
	// 
	// 到达提醒时间后，腾讯电子签会短信通知发起方企业合同提醒，可用于处理合同到期事务，如合同续签等事宜。
	RemindedOn *int64 `json:"RemindedOn,omitnil,omitempty" name:"RemindedOn"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	// 注：`请和模板中的配置保持一致`
	Unordered *bool `json:"Unordered,omitnil,omitempty" name:"Unordered"`

	// 您可以自定义**腾讯电子签小程序合同列表页**展示的合同内容模板，模板中支持以下变量：
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
	// 
	CustomShowMap *string `json:"CustomShowMap,omitnil,omitempty" name:"CustomShowMap"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil,omitempty" name:"NeedSignReview"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// <b>注</b>
	// 1. 抄送人名单中可以包括自然人以及本企业的员工。
	// 2. 请确保抄送人列表中的成员不与任何签署人重复。
	CcInfos []*CcInfo `json:"CcInfos,omitnil,omitempty" name:"CcInfos"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：电子处方单（医疗自动签）  </li><li> **OTHER** :  通用场景</li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil,omitempty" name:"AutoSignScene"`

	// 暂未开放
	//
	// Deprecated: RelatedFlowId is deprecated.
	RelatedFlowId *string `json:"RelatedFlowId,omitnil,omitempty" name:"RelatedFlowId"`

	// 暂未开放
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下： 
	//  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li></ul>
	// 
	// 效果如下:
	// ![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"FlowDisplayType,omitnil,omitempty" name:"FlowDisplayType"`
}

type CreateFlowRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自定义的合同流程的名称，长度不能超过200个字符，只能由中文汉字、中文标点、英文字母、阿拉伯数字、空格、小括号、中括号、中划线、下划线以及（,）、（;）、（.）、(&)、（+）组成。
	// 
	// 该名称还将用于合同签署完成后文件下载的默认文件名称。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 合同流程的参与方列表，最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的ApproverInfo结构体。
	// 
	// 注:  
	// <font color="red" > <b> 在发起流程时，需要保证 approver 中的顺序与模板定义顺序一致，否则会发起失败。
	// 例如，如果模板中定义的第一个参与人是个人用户，第二个参与人是企业员工，则在 approver 中传参时，第一个也必须是个人用户，第二个参与人必须是企业员工。</b></font>
	// 
	// [点击查看模板参与人顺序定义位置](https://qcloudimg.tencent-cloud.cn/raw/d14457b48cc66b29262ccb9d7b3ed556.png)
	Approvers []*FlowCreateApprover `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil,omitempty" name:"FlowDescription"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	// 此合同类型需要跟模板配置的合同类型保持一致。
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 已经废弃字段，客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	DeadLine *int64 `json:"DeadLine,omitnil,omitempty" name:"DeadLine"`

	// 合同到期提醒时间，为Unix标准时间戳（秒）格式，支持的范围是从发起时间开始到后10年内。
	// 
	// 到达提醒时间后，腾讯电子签会短信通知发起方企业合同提醒，可用于处理合同到期事务，如合同续签等事宜。
	RemindedOn *int64 `json:"RemindedOn,omitnil,omitempty" name:"RemindedOn"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	// 注：`请和模板中的配置保持一致`
	Unordered *bool `json:"Unordered,omitnil,omitempty" name:"Unordered"`

	// 您可以自定义**腾讯电子签小程序合同列表页**展示的合同内容模板，模板中支持以下变量：
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
	// 
	CustomShowMap *string `json:"CustomShowMap,omitnil,omitempty" name:"CustomShowMap"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil,omitempty" name:"NeedSignReview"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// <b>注</b>
	// 1. 抄送人名单中可以包括自然人以及本企业的员工。
	// 2. 请确保抄送人列表中的成员不与任何签署人重复。
	CcInfos []*CcInfo `json:"CcInfos,omitnil,omitempty" name:"CcInfos"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：电子处方单（医疗自动签）  </li><li> **OTHER** :  通用场景</li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil,omitempty" name:"AutoSignScene"`

	// 暂未开放
	RelatedFlowId *string `json:"RelatedFlowId,omitnil,omitempty" name:"RelatedFlowId"`

	// 暂未开放
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下： 
	//  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li></ul>
	// 
	// 效果如下:
	// ![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"FlowDisplayType,omitnil,omitempty" name:"FlowDisplayType"`
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
	delete(f, "FlowDescription")
	delete(f, "FlowType")
	delete(f, "ClientToken")
	delete(f, "DeadLine")
	delete(f, "RemindedOn")
	delete(f, "UserData")
	delete(f, "Unordered")
	delete(f, "CustomShowMap")
	delete(f, "NeedSignReview")
	delete(f, "Agent")
	delete(f, "CcInfos")
	delete(f, "AutoSignScene")
	delete(f, "RelatedFlowId")
	delete(f, "CallbackUrl")
	delete(f, "FlowDisplayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowResponseParams struct {
	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 
	// 注:
	// 此返回的合同流程ID，需再次调用<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateDocument" target="_blank">创建电子文档</a>和<a href="https://qian.tencent.com/developers/companyApis/startFlows/StartFlow" target="_blank">发起签署流程</a>接口将合同开始后，合同才能进入签署环节，[点击查看FlowId在控制台中的位置（只在进入签署环节后有效）](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)
	// 
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// <ul><li>建议开发者妥善保存此流程ID，以便于顺利进行后续操作。</li>
	// <li>可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。</li></ul>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 企业内部审核结果
	// <ul><li>PASS: 审核通过</li>
	// <li>REJECT: 审核拒绝</li>
	// <li>SIGN_REJECT:拒签(流程结束)</li></ul>
	ReviewType *string `json:"ReviewType,omitnil,omitempty" name:"ReviewType"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 审核节点的签署人标志，用于指定当前审核的签署方
	// <ul><li>**如果签署审核节点是个人， 此参数必填**。</li></ul>
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 操作类型：（接口通过该字段区分不同的操作类型）
	// 
	// <ul><li>SignReview: 签署审核（默认）</li>
	// <li>CreateReview: 创建审核</li></ul>
	// 
	// 如果审核节点是个人，则操作类型只能为SignReview。
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 审核结果原因
	// <ul><li>字符串长度不超过200</li>
	// <li>当ReviewType 是拒绝（REJECT） 时此字段必填。</li>
	// <li>当ReviewType 是拒绝（SIGN_REJECT） 时此字段必填。</li></ul>
	// 
	ReviewMessage *string `json:"ReviewMessage,omitnil,omitempty" name:"ReviewMessage"`
}

type CreateFlowSignReviewRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// <ul><li>建议开发者妥善保存此流程ID，以便于顺利进行后续操作。</li>
	// <li>可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。</li></ul>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 企业内部审核结果
	// <ul><li>PASS: 审核通过</li>
	// <li>REJECT: 审核拒绝</li>
	// <li>SIGN_REJECT:拒签(流程结束)</li></ul>
	ReviewType *string `json:"ReviewType,omitnil,omitempty" name:"ReviewType"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 审核节点的签署人标志，用于指定当前审核的签署方
	// <ul><li>**如果签署审核节点是个人， 此参数必填**。</li></ul>
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 操作类型：（接口通过该字段区分不同的操作类型）
	// 
	// <ul><li>SignReview: 签署审核（默认）</li>
	// <li>CreateReview: 创建审核</li></ul>
	// 
	// 如果审核节点是个人，则操作类型只能为SignReview。
	OperateType *string `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 审核结果原因
	// <ul><li>字符串长度不超过200</li>
	// <li>当ReviewType 是拒绝（REJECT） 时此字段必填。</li>
	// <li>当ReviewType 是拒绝（SIGN_REJECT） 时此字段必填。</li></ul>
	// 
	ReviewMessage *string `json:"ReviewMessage,omitnil,omitempty" name:"ReviewMessage"`
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
	delete(f, "Agent")
	delete(f, "RecipientId")
	delete(f, "OperateType")
	delete(f, "ReviewMessage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowSignReviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowSignReviewResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateFlowSignUrlRequestParams struct {
	// 合同流程ID为32位字符串。
	// 
	// 您可以登录腾讯电子签控制台，在 "合同" -> "合同中心" 中查看某个合同的FlowId（在页面中展示为合同ID）。[点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 流程签署人列表中，结构体的ApproverName、ApproverMobile和ApproverType为必传字段。如果是企业签署人，还需传递OrganizationName。
	// 
	// 注：
	// 1. 签署人<b>只能使用手写签名、时间类型、印章类型、签批类型的签署控件和内容填写控件</b>，其他类型的签署控件暂时不支持。
	// 2. 生成发起方预览链接时，该字段（FlowApproverInfos）可以传空或者不传。
	// 
	FlowApproverInfos []*FlowCreateApprover `json:"FlowApproverInfos,omitnil,omitempty" name:"FlowApproverInfos"`

	// 机构信息，暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 签署完之后的H5页面的跳转链接，最大长度1000个字符。链接类型请参考 <a href="https://qian.tencent.com/developers/company/openqianh5" target="_blank">跳转电子签H5</a>
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 链接类型支持以下指定类型：
	// 
	// <ul><li><b>0</b>: 签署链接（默认值），进入后可以填写或签署合同。</li><li><b>1 </b>: 预览链接，进入后可以预览合同当前的样子。</li></ul>
	// 
	// 注：
	// 
	// 1. 当指定链接类型为1时，链接为预览链接，打开链接后无法进行签署操作，仅支持预览和查看当前合同状态。
	// 2. 如需生成发起方预览链接，则签署方信息应传空，即FlowApproverInfos传空或者不传。
	UrlType *int64 `json:"UrlType,omitnil,omitempty" name:"UrlType"`
}

type CreateFlowSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 合同流程ID为32位字符串。
	// 
	// 您可以登录腾讯电子签控制台，在 "合同" -> "合同中心" 中查看某个合同的FlowId（在页面中展示为合同ID）。[点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 流程签署人列表中，结构体的ApproverName、ApproverMobile和ApproverType为必传字段。如果是企业签署人，还需传递OrganizationName。
	// 
	// 注：
	// 1. 签署人<b>只能使用手写签名、时间类型、印章类型、签批类型的签署控件和内容填写控件</b>，其他类型的签署控件暂时不支持。
	// 2. 生成发起方预览链接时，该字段（FlowApproverInfos）可以传空或者不传。
	// 
	FlowApproverInfos []*FlowCreateApprover `json:"FlowApproverInfos,omitnil,omitempty" name:"FlowApproverInfos"`

	// 机构信息，暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 签署完之后的H5页面的跳转链接，最大长度1000个字符。链接类型请参考 <a href="https://qian.tencent.com/developers/company/openqianh5" target="_blank">跳转电子签H5</a>
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 链接类型支持以下指定类型：
	// 
	// <ul><li><b>0</b>: 签署链接（默认值），进入后可以填写或签署合同。</li><li><b>1 </b>: 预览链接，进入后可以预览合同当前的样子。</li></ul>
	// 
	// 注：
	// 
	// 1. 当指定链接类型为1时，链接为预览链接，打开链接后无法进行签署操作，仅支持预览和查看当前合同状态。
	// 2. 如需生成发起方预览链接，则签署方信息应传空，即FlowApproverInfos传空或者不传。
	UrlType *int64 `json:"UrlType,omitnil,omitempty" name:"UrlType"`
}

func (r *CreateFlowSignUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowSignUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Operator")
	delete(f, "Agent")
	delete(f, "FlowApproverInfos")
	delete(f, "Organization")
	delete(f, "JumpUrl")
	delete(f, "UrlType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowSignUrlResponseParams struct {
	// 签署人签署链接信息
	FlowApproverUrlInfos []*FlowApproverUrlInfo `json:"FlowApproverUrlInfos,omitnil,omitempty" name:"FlowApproverUrlInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFlowSignUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowSignUrlResponseParams `json:"Response"`
}

func (r *CreateFlowSignUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowSignUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationDepartmentRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 部门名称，最大长度为50个字符。
	DeptName *string `json:"DeptName,omitnil,omitempty" name:"DeptName"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 电子签父部门ID。
	// 注：`如果同时指定了ParentDeptId与ParentDeptOpenId参数，系统将优先使用ParentDeptId。当二者都未指定时，创建的新部门将自动填充至根节点部门下。`
	ParentDeptId *string `json:"ParentDeptId,omitnil,omitempty" name:"ParentDeptId"`

	// 第三方平台中父部门ID。
	// 注：`如果同时指定了ParentDeptId与ParentDeptOpenId参数，系统将优先使用ParentDeptId。当二者都未指定时，创建的新部门将自动填充至根节点部门下。`
	ParentDeptOpenId *string `json:"ParentDeptOpenId,omitnil,omitempty" name:"ParentDeptOpenId"`

	// 客户系统部门ID，最大长度为64个字符。
	DeptOpenId *string `json:"DeptOpenId,omitnil,omitempty" name:"DeptOpenId"`

	// 排序号，支持设置的数值范围为1~30000。同一父部门下，排序号越大，部门顺序越靠前。
	OrderNo *uint64 `json:"OrderNo,omitnil,omitempty" name:"OrderNo"`
}

type CreateIntegrationDepartmentRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 部门名称，最大长度为50个字符。
	DeptName *string `json:"DeptName,omitnil,omitempty" name:"DeptName"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 电子签父部门ID。
	// 注：`如果同时指定了ParentDeptId与ParentDeptOpenId参数，系统将优先使用ParentDeptId。当二者都未指定时，创建的新部门将自动填充至根节点部门下。`
	ParentDeptId *string `json:"ParentDeptId,omitnil,omitempty" name:"ParentDeptId"`

	// 第三方平台中父部门ID。
	// 注：`如果同时指定了ParentDeptId与ParentDeptOpenId参数，系统将优先使用ParentDeptId。当二者都未指定时，创建的新部门将自动填充至根节点部门下。`
	ParentDeptOpenId *string `json:"ParentDeptOpenId,omitnil,omitempty" name:"ParentDeptOpenId"`

	// 客户系统部门ID，最大长度为64个字符。
	DeptOpenId *string `json:"DeptOpenId,omitnil,omitempty" name:"DeptOpenId"`

	// 排序号，支持设置的数值范围为1~30000。同一父部门下，排序号越大，部门顺序越靠前。
	OrderNo *uint64 `json:"OrderNo,omitnil,omitempty" name:"OrderNo"`
}

func (r *CreateIntegrationDepartmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationDepartmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "DeptName")
	delete(f, "Agent")
	delete(f, "ParentDeptId")
	delete(f, "ParentDeptOpenId")
	delete(f, "DeptOpenId")
	delete(f, "OrderNo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrationDepartmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationDepartmentResponseParams struct {
	// 电子签部门ID。建议开发者保存此部门ID，方便后续查询或修改部门信息。
	DeptId *string `json:"DeptId,omitnil,omitempty" name:"DeptId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIntegrationDepartmentResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrationDepartmentResponseParams `json:"Response"`
}

func (r *CreateIntegrationDepartmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationEmployeesRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 待创建员工的信息最多不超过20个。
	// 
	// **1. 在创建企业微信员工的场景下** :  只需传入下面的参数，其他信息不支持设置。
	// <table> <thead> <tr> <th>参数</th> <th>是否必填</th> <th>含义</th> </tr> </thead> <tbody> <tr> <td>WeworkOpenId</td> <td>是</td> <td>企业微信用户账号ID</td> </tr> </tbody> </table>
	// 
	// **2. 在其他场景下** :   只需传入下面的参数，其他信息不支持设置。
	// <table> <thead> <tr> <th>参数</th> <th>是否必填</th> <th>含义</th> </tr> </thead> <tbody> <tr> <td>DisplayName</td> <td>是</td> <td>用户的真实名字</td> </tr> <tr> <td>Mobile</td> <td>是</td> <td>用户手机号码</td> </tr> <tr> <td>OpenId</td> <td>否</td> <td>用户的自定义ID</td> </tr> <tr> <td>Email</td> <td>否</td> <td>用户的邮箱</td> </tr> <tr> <td>Department.DepartmentId</td> <td>否</td> <td>用户加入后的部门ID</td> </tr> </tbody> </table>
	// 
	// 
	// 注: `每个手机号每天最多使用3次`
	Employees []*Staff `json:"Employees,omitnil,omitempty" name:"Employees"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 员工邀请方式
	// 可通过以下途径进行设置：
	// <ul><li>**SMS（默认）**：邀请将通过短信或企业微信消息发送。若场景非企业微信，则采用企业微信消息；其他情境下则使用短信通知。短信内含链接，点击后将进入微信小程序进行认证并加入企业的流程。</li><li>**H5**：将生成H5链接，用户点击链接后可进入H5页面进行认证并加入企业的流程。</li><li>**NONE**：系统会根据Endpoint生成签署链接，业务方需获取链接并通知客户。</li></ul>
	InvitationNotifyType *string `json:"InvitationNotifyType,omitnil,omitempty" name:"InvitationNotifyType"`

	// 回跳地址，为认证成功后页面进行回跳的URL，请确保回跳地址的可用性。
	// 
	// 注：`只有在员工邀请方式（InvitationNotifyType参数）为H5场景下才生效， 其他方式下设置无效。`
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 要跳转的链接类型<ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li><li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li><li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li><li>**H5**： 第三方移动端浏览器进行嵌入，不支持小程序嵌入，过期时间一个月</li></ul>注意：InvitationNotifyType 和 Endpoint 的关系图<table><tbody><tr><td>通知类型（InvitationNotifyType）</td><td>Endpoint</td></tr><tr><td>SMS（默认）</td><td>不需要传递，会将 Endpoint 默认设置为HTTP_SHORT_URL</td></tr><tr><td>H5</td><td>不需要传递，会将 Endpoint 默认设置为 H5</td></tr><tr><td>NONE</td><td>所有 Endpoint 都支持（HTTP_URL/HTTP_SHORT_URL/H5/APP）默认为HTTP_SHORT_URL</td></tr></tbody></table>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

type CreateIntegrationEmployeesRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 待创建员工的信息最多不超过20个。
	// 
	// **1. 在创建企业微信员工的场景下** :  只需传入下面的参数，其他信息不支持设置。
	// <table> <thead> <tr> <th>参数</th> <th>是否必填</th> <th>含义</th> </tr> </thead> <tbody> <tr> <td>WeworkOpenId</td> <td>是</td> <td>企业微信用户账号ID</td> </tr> </tbody> </table>
	// 
	// **2. 在其他场景下** :   只需传入下面的参数，其他信息不支持设置。
	// <table> <thead> <tr> <th>参数</th> <th>是否必填</th> <th>含义</th> </tr> </thead> <tbody> <tr> <td>DisplayName</td> <td>是</td> <td>用户的真实名字</td> </tr> <tr> <td>Mobile</td> <td>是</td> <td>用户手机号码</td> </tr> <tr> <td>OpenId</td> <td>否</td> <td>用户的自定义ID</td> </tr> <tr> <td>Email</td> <td>否</td> <td>用户的邮箱</td> </tr> <tr> <td>Department.DepartmentId</td> <td>否</td> <td>用户加入后的部门ID</td> </tr> </tbody> </table>
	// 
	// 
	// 注: `每个手机号每天最多使用3次`
	Employees []*Staff `json:"Employees,omitnil,omitempty" name:"Employees"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 员工邀请方式
	// 可通过以下途径进行设置：
	// <ul><li>**SMS（默认）**：邀请将通过短信或企业微信消息发送。若场景非企业微信，则采用企业微信消息；其他情境下则使用短信通知。短信内含链接，点击后将进入微信小程序进行认证并加入企业的流程。</li><li>**H5**：将生成H5链接，用户点击链接后可进入H5页面进行认证并加入企业的流程。</li><li>**NONE**：系统会根据Endpoint生成签署链接，业务方需获取链接并通知客户。</li></ul>
	InvitationNotifyType *string `json:"InvitationNotifyType,omitnil,omitempty" name:"InvitationNotifyType"`

	// 回跳地址，为认证成功后页面进行回跳的URL，请确保回跳地址的可用性。
	// 
	// 注：`只有在员工邀请方式（InvitationNotifyType参数）为H5场景下才生效， 其他方式下设置无效。`
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 要跳转的链接类型<ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li><li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li><li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li><li>**H5**： 第三方移动端浏览器进行嵌入，不支持小程序嵌入，过期时间一个月</li></ul>注意：InvitationNotifyType 和 Endpoint 的关系图<table><tbody><tr><td>通知类型（InvitationNotifyType）</td><td>Endpoint</td></tr><tr><td>SMS（默认）</td><td>不需要传递，会将 Endpoint 默认设置为HTTP_SHORT_URL</td></tr><tr><td>H5</td><td>不需要传递，会将 Endpoint 默认设置为 H5</td></tr><tr><td>NONE</td><td>所有 Endpoint 都支持（HTTP_URL/HTTP_SHORT_URL/H5/APP）默认为HTTP_SHORT_URL</td></tr></tbody></table>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

func (r *CreateIntegrationEmployeesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationEmployeesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Employees")
	delete(f, "Agent")
	delete(f, "InvitationNotifyType")
	delete(f, "JumpUrl")
	delete(f, "Endpoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrationEmployeesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationEmployeesResponseParams struct {
	// 创建员工的结果。包含创建成功的数据与创建失败数据。
	CreateEmployeeResult *CreateStaffResult `json:"CreateEmployeeResult,omitnil,omitempty" name:"CreateEmployeeResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIntegrationEmployeesResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrationEmployeesResponseParams `json:"Response"`
}

func (r *CreateIntegrationEmployeesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationEmployeesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationRoleRequestParams struct {
	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 角色类型，0:saas角色，1:集团角色
	// 默认0，saas角色
	IsGroupRole *int64 `json:"IsGroupRole,omitnil,omitempty" name:"IsGroupRole"`

	// 权限树
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil,omitempty" name:"PermissionGroups"`

	// 集团角色的话，需要传递集团子企业列表，如果是全选，则传1
	SubOrganizationIds *string `json:"SubOrganizationIds,omitnil,omitempty" name:"SubOrganizationIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CreateIntegrationRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 角色类型，0:saas角色，1:集团角色
	// 默认0，saas角色
	IsGroupRole *int64 `json:"IsGroupRole,omitnil,omitempty" name:"IsGroupRole"`

	// 权限树
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil,omitempty" name:"PermissionGroups"`

	// 集团角色的话，需要传递集团子企业列表，如果是全选，则传1
	SubOrganizationIds *string `json:"SubOrganizationIds,omitnil,omitempty" name:"SubOrganizationIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *CreateIntegrationRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Operator")
	delete(f, "Description")
	delete(f, "IsGroupRole")
	delete(f, "PermissionGroups")
	delete(f, "SubOrganizationIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrationRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationRoleResponseParams struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIntegrationRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrationRoleResponseParams `json:"Response"`
}

func (r *CreateIntegrationRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationSubOrganizationActiveRecordRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。 支持填入集团子公司经办人 userId 代发合同。  注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 待激活成员企业ID集合
	SubOrganizationIds []*string `json:"SubOrganizationIds,omitnil,omitempty" name:"SubOrganizationIds"`
}

type CreateIntegrationSubOrganizationActiveRecordRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。 支持填入集团子公司经办人 userId 代发合同。  注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 待激活成员企业ID集合
	SubOrganizationIds []*string `json:"SubOrganizationIds,omitnil,omitempty" name:"SubOrganizationIds"`
}

func (r *CreateIntegrationSubOrganizationActiveRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationSubOrganizationActiveRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SubOrganizationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrationSubOrganizationActiveRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationSubOrganizationActiveRecordResponseParams struct {
	// 激活失败的成员企业ID集合
	FailedSubOrganizationIds []*string `json:"FailedSubOrganizationIds,omitnil,omitempty" name:"FailedSubOrganizationIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIntegrationSubOrganizationActiveRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrationSubOrganizationActiveRecordResponseParams `json:"Response"`
}

func (r *CreateIntegrationSubOrganizationActiveRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationSubOrganizationActiveRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationUserRolesRequestParams struct {
	// 执行本接口操作的员工信息。 注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 绑定角色的用户id列表，不能重复，不能大于 100 个
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 绑定角色的角色id列表，不能重复，不能大于 100，可以通过DescribeIntegrationRoles接口获取角色信息
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CreateIntegrationUserRolesRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。 注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 绑定角色的用户id列表，不能重复，不能大于 100 个
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 绑定角色的角色id列表，不能重复，不能大于 100，可以通过DescribeIntegrationRoles接口获取角色信息
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *CreateIntegrationUserRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationUserRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "UserIds")
	delete(f, "RoleIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrationUserRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationUserRolesResponseParams struct {
	// 绑定角色失败列表信息
	FailedCreateRoleData []*FailedCreateRoleData `json:"FailedCreateRoleData,omitnil,omitempty" name:"FailedCreateRoleData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIntegrationUserRolesResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrationUserRolesResponseParams `json:"Response"`
}

func (r *CreateIntegrationUserRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationUserRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLegalSealQrCodeRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 机构信息，暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
}

type CreateLegalSealQrCodeRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 机构信息，暂未开放
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
	delete(f, "Operator")
	delete(f, "Agent")
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
type CreateMultiFlowSignQRCodeRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同模板ID，为32位字符串。
	// 可登录腾讯电子签控制台，在 "模板"->"模板中心"->"列表展示设置"选中模板 ID 中查看某个模板的TemplateId(在页面中展示为模板ID)。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 通过此二维码可发起的流程最大限额，如未明确指定，默认为5份。
	// 一旦发起流程数超越该限制，该二维码将自动失效。
	MaxFlowNum *int64 `json:"MaxFlowNum,omitnil,omitempty" name:"MaxFlowNum"`

	// 二维码的有效期限，默认为7天，最高设定不得超过90天。
	// 一旦超过二维码的有效期限，该二维码将自动失效。
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitnil,omitempty" name:"QrEffectiveDay"`

	// 合同流程的签署有效期限，若未设定签署截止日期，则默认为自合同流程创建起的7天内截止。
	// 若在签署截止日期前未完成签署，合同状态将变更为已过期，从而导致合同无效。
	// 最长设定期限不得超过30天。
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitnil,omitempty" name:"FlowEffectiveDay"`

	// 指定签署人信息。
	// 在指定签署人后，仅允许特定签署人通过扫描二维码进行签署。
	Restrictions []*ApproverRestriction `json:"Restrictions,omitnil,omitempty" name:"Restrictions"`

	// 调用方自定义的个性化字段(可自定义此字段的值)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 已废弃，回调配置统一使用企业应用管理-应用集成-企业版应用中的配置 
	// <br/> 通过一码多扫二维码发起的合同，回调消息可参考文档 https://qian.tencent.com/developers/company/callback_types_contracts_sign
	// <br/> 用户通过签署二维码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档 https://qian.tencent.com/developers/company/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 限制二维码用户条件（已弃用）
	//
	// Deprecated: ApproverRestrictions is deprecated.
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitnil,omitempty" name:"ApproverRestrictions"`

	// 指定签署方在使用个人印章签署控件（SIGN_SIGNATURE） 时可使用的签署方式：自由书写、正楷临摹、系统签名、个人印章。
	ApproverComponentLimitTypes []*ApproverComponentLimitType `json:"ApproverComponentLimitTypes,omitnil,omitempty" name:"ApproverComponentLimitTypes"`

	// 禁止个人用户重复签署，默认不禁止，即同一用户可多次扫码签署多份合同。若要求同一用户仅能扫码签署一份合同，请传入true。
	ForbidPersonalMultipleSign *bool `json:"ForbidPersonalMultipleSign,omitnil,omitempty" name:"ForbidPersonalMultipleSign"`

	// 合同流程名称是否应包含扫码签署人的信息，且遵循特定格式（flowname-姓名-手机号后四位）。
	// 例如，通过参数FlowName设定的扫码发起合同名称为“员工入职合同”，当扫码人张三（手机号18800009527）扫码签署时，合同名称将自动生成为“员工入职合同-张三-9527”。
	FlowNameAppendScannerInfo *bool `json:"FlowNameAppendScannerInfo,omitnil,omitempty" name:"FlowNameAppendScannerInfo"`
}

type CreateMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同模板ID，为32位字符串。
	// 可登录腾讯电子签控制台，在 "模板"->"模板中心"->"列表展示设置"选中模板 ID 中查看某个模板的TemplateId(在页面中展示为模板ID)。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 通过此二维码可发起的流程最大限额，如未明确指定，默认为5份。
	// 一旦发起流程数超越该限制，该二维码将自动失效。
	MaxFlowNum *int64 `json:"MaxFlowNum,omitnil,omitempty" name:"MaxFlowNum"`

	// 二维码的有效期限，默认为7天，最高设定不得超过90天。
	// 一旦超过二维码的有效期限，该二维码将自动失效。
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitnil,omitempty" name:"QrEffectiveDay"`

	// 合同流程的签署有效期限，若未设定签署截止日期，则默认为自合同流程创建起的7天内截止。
	// 若在签署截止日期前未完成签署，合同状态将变更为已过期，从而导致合同无效。
	// 最长设定期限不得超过30天。
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitnil,omitempty" name:"FlowEffectiveDay"`

	// 指定签署人信息。
	// 在指定签署人后，仅允许特定签署人通过扫描二维码进行签署。
	Restrictions []*ApproverRestriction `json:"Restrictions,omitnil,omitempty" name:"Restrictions"`

	// 调用方自定义的个性化字段(可自定义此字段的值)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 已废弃，回调配置统一使用企业应用管理-应用集成-企业版应用中的配置 
	// <br/> 通过一码多扫二维码发起的合同，回调消息可参考文档 https://qian.tencent.com/developers/company/callback_types_contracts_sign
	// <br/> 用户通过签署二维码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档 https://qian.tencent.com/developers/company/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 限制二维码用户条件（已弃用）
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitnil,omitempty" name:"ApproverRestrictions"`

	// 指定签署方在使用个人印章签署控件（SIGN_SIGNATURE） 时可使用的签署方式：自由书写、正楷临摹、系统签名、个人印章。
	ApproverComponentLimitTypes []*ApproverComponentLimitType `json:"ApproverComponentLimitTypes,omitnil,omitempty" name:"ApproverComponentLimitTypes"`

	// 禁止个人用户重复签署，默认不禁止，即同一用户可多次扫码签署多份合同。若要求同一用户仅能扫码签署一份合同，请传入true。
	ForbidPersonalMultipleSign *bool `json:"ForbidPersonalMultipleSign,omitnil,omitempty" name:"ForbidPersonalMultipleSign"`

	// 合同流程名称是否应包含扫码签署人的信息，且遵循特定格式（flowname-姓名-手机号后四位）。
	// 例如，通过参数FlowName设定的扫码发起合同名称为“员工入职合同”，当扫码人张三（手机号18800009527）扫码签署时，合同名称将自动生成为“员工入职合同-张三-9527”。
	FlowNameAppendScannerInfo *bool `json:"FlowNameAppendScannerInfo,omitnil,omitempty" name:"FlowNameAppendScannerInfo"`
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
	delete(f, "Operator")
	delete(f, "TemplateId")
	delete(f, "FlowName")
	delete(f, "MaxFlowNum")
	delete(f, "QrEffectiveDay")
	delete(f, "FlowEffectiveDay")
	delete(f, "Restrictions")
	delete(f, "UserData")
	delete(f, "CallbackUrl")
	delete(f, "Agent")
	delete(f, "ApproverRestrictions")
	delete(f, "ApproverComponentLimitTypes")
	delete(f, "ForbidPersonalMultipleSign")
	delete(f, "FlowNameAppendScannerInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMultiFlowSignQRCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultiFlowSignQRCodeResponseParams struct {
	// 一码多签签署码的基本信息，用户可扫描该二维码进行签署操作。
	QrCode *SignQrCode `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 一码多签签署码的链接信息，适用于客户系统整合二维码功能。通过链接，用户可直接访问电子签名小程序并签署合同。
	SignUrls *SignUrl `json:"SignUrls,omitnil,omitempty" name:"SignUrls"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateOrganizationAuthFileRequestParams struct {
	// 企业授权书信息参数， 需要自行保证这些参数跟营业执照中的信息一致。
	OrganizationCommonInfo *OrganizationCommonInfo `json:"OrganizationCommonInfo,omitnil,omitempty" name:"OrganizationCommonInfo"`

	// 代理企业和员工的信息。<br/>在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 执行本接口操作的员工信息。<br/>注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 授权书类型：
	// - 0: 企业认证超管授权书
	// - 1: 超管变更授权书
	// - 2: 企业注销授权书
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type CreateOrganizationAuthFileRequest struct {
	*tchttp.BaseRequest
	
	// 企业授权书信息参数， 需要自行保证这些参数跟营业执照中的信息一致。
	OrganizationCommonInfo *OrganizationCommonInfo `json:"OrganizationCommonInfo,omitnil,omitempty" name:"OrganizationCommonInfo"`

	// 代理企业和员工的信息。<br/>在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 执行本接口操作的员工信息。<br/>注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 授权书类型：
	// - 0: 企业认证超管授权书
	// - 1: 超管变更授权书
	// - 2: 企业注销授权书
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
	delete(f, "OrganizationCommonInfo")
	delete(f, "Agent")
	delete(f, "Operator")
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
type CreateOrganizationAuthUrlRequestParams struct {
	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定授权方式 支持多选:
	// 
	// <ul>
	// <li><strong>1</strong>:上传授权书方式</li>
	// <li><strong>2</strong>: 法人授权方式</li>
	// <li><strong>3</strong>: 法人身份认证方式</li>
	// </ul>
	AuthorizationTypes []*uint64 `json:"AuthorizationTypes,omitnil,omitempty" name:"AuthorizationTypes"`

	// 认证企业名称，请确认该名称与企业营业执照中注册的名称一致。
	// 
	// 注：
	// 
	// 1. `如果名称中包含英文括号()，请使用中文括号（）代替。`
	// 
	// 2. `EndPointType=“H5”或者"SHORT_H5"时，该参数必填`
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 企业统一社会信用代码
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// 企业法人的姓名
	LegalName *string `json:"LegalName,omitnil,omitempty" name:"LegalName"`

	// 认证完成跳回的链接，最长500个字符
	AutoJumpUrl *string `json:"AutoJumpUrl,omitnil,omitempty" name:"AutoJumpUrl"`

	// 营业执照企业地址
	OrganizationAddress *string `json:"OrganizationAddress,omitnil,omitempty" name:"OrganizationAddress"`

	// 认证人姓名
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 认证人手机号
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`

	// 认证人身份证号
	AdminIdCardNumber *string `json:"AdminIdCardNumber,omitnil,omitempty" name:"AdminIdCardNumber"`

	// 认证人证件类型， 支持以下类型
	// <ul><li><b>ID_CARD</b> : 中国大陆居民身份证  (默认值)</li>
	// <li><b>HONGKONG_AND_MACAO</b>  : 中国港澳居民来往内地通行证</li>
	// <li><b>HONGKONG_MACAO_AND_TAIWAN</b>  : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	AdminIdCardType *string `json:"AdminIdCardType,omitnil,omitempty" name:"AdminIdCardType"`

	// 对方打开链接认证时，对方填写的营业执照的社会信用代码是否与接口上传上来的要保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>
	UniformSocialCreditCodeSame *bool `json:"UniformSocialCreditCodeSame,omitnil,omitempty" name:"UniformSocialCreditCodeSame"`

	// 对方打开链接认证时，法人姓名是否要与接口传递上来的保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>
	// 
	// p.s. 仅在法人姓名不为空时有效
	LegalNameSame *bool `json:"LegalNameSame,omitnil,omitempty" name:"LegalNameSame"`

	// 对方打开链接认证时，认证人姓名是否要与接口传递上来的保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>
	// 
	// p.s. 仅在认证人姓名不为空时有效
	AdminNameSame *bool `json:"AdminNameSame,omitnil,omitempty" name:"AdminNameSame"`

	// 对方打开链接认证时，认证人居民身份证件号是否要与接口传递上来的保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>
	// 
	// p.s. 仅在认证人身份证号不为空时有效
	AdminIdCardNumberSame *bool `json:"AdminIdCardNumberSame,omitnil,omitempty" name:"AdminIdCardNumberSame"`

	// 对方打开链接认证时，认证人手机号是否要与接口传递上来的保持一致。<ul>
	// <li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li>
	// <li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li>
	// </ul>
	// 
	// p.s. 仅在认证人手机号不为空时有效
	AdminMobileSame *bool `json:"AdminMobileSame,omitnil,omitempty" name:"AdminMobileSame"`

	// 对方打开链接认证时，企业名称是否要与接口传递上来的保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>
	// 
	// 
	// p.s. 仅在企业名称不为空时有效
	OrganizationNameSame *bool `json:"OrganizationNameSame,omitnil,omitempty" name:"OrganizationNameSame"`

	// 营业执照正面照（支持PNG或JPG格式）需以base64格式提供，且文件大小不得超过5MB。
	BusinessLicense *string `json:"BusinessLicense,omitnil,omitempty" name:"BusinessLicense"`

	// 跳转链接类型：
	// 
	// <ul>
	// <li><b>PC</b>：适用于PC端的认证链接</li>
	// <li><b>APP</b>：用于全屏或半屏跳转的小程序链接</li>
	// <li><b>SHORT_URL</b>：跳转小程序的链接的短链形式</li>
	// <li><b>H5</b>：适用于H5页面的认证链接</li>
	// <li><b>SHORT_H5</b>：H5认证链接的短链形式</li>
	// </ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 指定企业初始化引导，现在可以配置如下的选项：
	// 
	// <b>1</b>: 启用此选项后，在企业认证的最终步骤将添加创建印章的引导。如下图的位置
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/88e0b45095a5c589de8995462ad755dc.jpg)
	Initialization []*uint64 `json:"Initialization,omitnil,omitempty" name:"Initialization"`

	// 授权书(PNG或JPG或PDF) base64格式, 大小不超过8M 。 
	// 授权书可以通过接口[生成企业授权书](https://qian.tencent.com/developers/companyApis/organizations/CreateOrganizationAuthFile) 来获得。
	// p.s. 如果上传授权书 ，需遵循以下条件 
	// 1.  超管的信息（超管姓名，超管手机号）必须为必填参数。
	// 2.  认证方式AuthorizationTypes必须只能是上传授权书方式 
	PowerOfAttorneys []*string `json:"PowerOfAttorneys,omitnil,omitempty" name:"PowerOfAttorneys"`
}

type CreateOrganizationAuthUrlRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定授权方式 支持多选:
	// 
	// <ul>
	// <li><strong>1</strong>:上传授权书方式</li>
	// <li><strong>2</strong>: 法人授权方式</li>
	// <li><strong>3</strong>: 法人身份认证方式</li>
	// </ul>
	AuthorizationTypes []*uint64 `json:"AuthorizationTypes,omitnil,omitempty" name:"AuthorizationTypes"`

	// 认证企业名称，请确认该名称与企业营业执照中注册的名称一致。
	// 
	// 注：
	// 
	// 1. `如果名称中包含英文括号()，请使用中文括号（）代替。`
	// 
	// 2. `EndPointType=“H5”或者"SHORT_H5"时，该参数必填`
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 企业统一社会信用代码
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// 企业法人的姓名
	LegalName *string `json:"LegalName,omitnil,omitempty" name:"LegalName"`

	// 认证完成跳回的链接，最长500个字符
	AutoJumpUrl *string `json:"AutoJumpUrl,omitnil,omitempty" name:"AutoJumpUrl"`

	// 营业执照企业地址
	OrganizationAddress *string `json:"OrganizationAddress,omitnil,omitempty" name:"OrganizationAddress"`

	// 认证人姓名
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 认证人手机号
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`

	// 认证人身份证号
	AdminIdCardNumber *string `json:"AdminIdCardNumber,omitnil,omitempty" name:"AdminIdCardNumber"`

	// 认证人证件类型， 支持以下类型
	// <ul><li><b>ID_CARD</b> : 中国大陆居民身份证  (默认值)</li>
	// <li><b>HONGKONG_AND_MACAO</b>  : 中国港澳居民来往内地通行证</li>
	// <li><b>HONGKONG_MACAO_AND_TAIWAN</b>  : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	AdminIdCardType *string `json:"AdminIdCardType,omitnil,omitempty" name:"AdminIdCardType"`

	// 对方打开链接认证时，对方填写的营业执照的社会信用代码是否与接口上传上来的要保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>
	UniformSocialCreditCodeSame *bool `json:"UniformSocialCreditCodeSame,omitnil,omitempty" name:"UniformSocialCreditCodeSame"`

	// 对方打开链接认证时，法人姓名是否要与接口传递上来的保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>
	// 
	// p.s. 仅在法人姓名不为空时有效
	LegalNameSame *bool `json:"LegalNameSame,omitnil,omitempty" name:"LegalNameSame"`

	// 对方打开链接认证时，认证人姓名是否要与接口传递上来的保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>
	// 
	// p.s. 仅在认证人姓名不为空时有效
	AdminNameSame *bool `json:"AdminNameSame,omitnil,omitempty" name:"AdminNameSame"`

	// 对方打开链接认证时，认证人居民身份证件号是否要与接口传递上来的保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>
	// 
	// p.s. 仅在认证人身份证号不为空时有效
	AdminIdCardNumberSame *bool `json:"AdminIdCardNumberSame,omitnil,omitempty" name:"AdminIdCardNumberSame"`

	// 对方打开链接认证时，认证人手机号是否要与接口传递上来的保持一致。<ul>
	// <li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li>
	// <li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li>
	// </ul>
	// 
	// p.s. 仅在认证人手机号不为空时有效
	AdminMobileSame *bool `json:"AdminMobileSame,omitnil,omitempty" name:"AdminMobileSame"`

	// 对方打开链接认证时，企业名称是否要与接口传递上来的保持一致。<ul><li><b>false（默认值）</b>：关闭状态，实际认证时允许与接口传递的信息存在不一致。</li><li><b>true</b>：启用状态，实际认证时必须与接口传递的信息完全相符。</li></ul>
	// 
	// 
	// p.s. 仅在企业名称不为空时有效
	OrganizationNameSame *bool `json:"OrganizationNameSame,omitnil,omitempty" name:"OrganizationNameSame"`

	// 营业执照正面照（支持PNG或JPG格式）需以base64格式提供，且文件大小不得超过5MB。
	BusinessLicense *string `json:"BusinessLicense,omitnil,omitempty" name:"BusinessLicense"`

	// 跳转链接类型：
	// 
	// <ul>
	// <li><b>PC</b>：适用于PC端的认证链接</li>
	// <li><b>APP</b>：用于全屏或半屏跳转的小程序链接</li>
	// <li><b>SHORT_URL</b>：跳转小程序的链接的短链形式</li>
	// <li><b>H5</b>：适用于H5页面的认证链接</li>
	// <li><b>SHORT_H5</b>：H5认证链接的短链形式</li>
	// </ul>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 指定企业初始化引导，现在可以配置如下的选项：
	// 
	// <b>1</b>: 启用此选项后，在企业认证的最终步骤将添加创建印章的引导。如下图的位置
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/88e0b45095a5c589de8995462ad755dc.jpg)
	Initialization []*uint64 `json:"Initialization,omitnil,omitempty" name:"Initialization"`

	// 授权书(PNG或JPG或PDF) base64格式, 大小不超过8M 。 
	// 授权书可以通过接口[生成企业授权书](https://qian.tencent.com/developers/companyApis/organizations/CreateOrganizationAuthFile) 来获得。
	// p.s. 如果上传授权书 ，需遵循以下条件 
	// 1.  超管的信息（超管姓名，超管手机号）必须为必填参数。
	// 2.  认证方式AuthorizationTypes必须只能是上传授权书方式 
	PowerOfAttorneys []*string `json:"PowerOfAttorneys,omitnil,omitempty" name:"PowerOfAttorneys"`
}

func (r *CreateOrganizationAuthUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationAuthUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "AuthorizationTypes")
	delete(f, "OrganizationName")
	delete(f, "UniformSocialCreditCode")
	delete(f, "LegalName")
	delete(f, "AutoJumpUrl")
	delete(f, "OrganizationAddress")
	delete(f, "AdminName")
	delete(f, "AdminMobile")
	delete(f, "AdminIdCardNumber")
	delete(f, "AdminIdCardType")
	delete(f, "UniformSocialCreditCodeSame")
	delete(f, "LegalNameSame")
	delete(f, "AdminNameSame")
	delete(f, "AdminIdCardNumberSame")
	delete(f, "AdminMobileSame")
	delete(f, "OrganizationNameSame")
	delete(f, "BusinessLicense")
	delete(f, "Endpoint")
	delete(f, "Initialization")
	delete(f, "PowerOfAttorneys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationAuthUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationAuthUrlResponseParams struct {
	// 生成的认证链接。
	// 
	// 注： `链接有效期统一30天`
	AuthUrl *string `json:"AuthUrl,omitnil,omitempty" name:"AuthUrl"`

	// 链接过期时间，格式为Unix标准时间戳（秒）
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationAuthUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationAuthUrlResponseParams `json:"Response"`
}

func (r *CreateOrganizationAuthUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationAuthUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationBatchSignUrlRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 请指定需执行批量签署的流程ID，数量范围为1-100。
	// 您可登录腾讯电子签控制台，浏览 "合同"->"合同中心" 以查阅某一合同的FlowId（在页面中显示为合同ID）。
	// 用户将利用链接对这些合同实施批量操作。
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 员工在腾讯电子签平台的独特身份标识，为32位字符串。
	// 您可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查阅某位员工的UserId（在页面中显示为用户ID）。
	// UserId必须是传入合同（FlowId）中的签署人。
	// 
	// <ul>
	// <li>1. 若UserId为空，Name和Mobile 必须提供。</li>
	// <li>2. 若UserId 与 Name，Mobile均存在，将优先采用UserId对应的员工。</li>
	// </ul>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 员工姓名，必须与手机号码一起使用。
	// 如果UserId为空，则此字段不能为空。同时，姓名和手机号码必须与传入合同（FlowId）中的签署人信息一致。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 员工手机号，必须与姓名一起使用。
	//  如果UserId为空，则此字段不能为空。同时，姓名和手机号码必须与传入合同（FlowId）中的签署人信息一致。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 为签署方经办人在签署合同中的参与方ID，必须与参数FlowIds数组一一对应。
	// 您可以通过查询合同接口（DescribeFlowInfo）查询此参数。
	// 若传了此参数，则可以不传 UserId, Name, Mobile等参数
	RecipientIds []*string `json:"RecipientIds,omitnil,omitempty" name:"RecipientIds"`
}

type CreateOrganizationBatchSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 请指定需执行批量签署的流程ID，数量范围为1-100。
	// 您可登录腾讯电子签控制台，浏览 "合同"->"合同中心" 以查阅某一合同的FlowId（在页面中显示为合同ID）。
	// 用户将利用链接对这些合同实施批量操作。
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 员工在腾讯电子签平台的独特身份标识，为32位字符串。
	// 您可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查阅某位员工的UserId（在页面中显示为用户ID）。
	// UserId必须是传入合同（FlowId）中的签署人。
	// 
	// <ul>
	// <li>1. 若UserId为空，Name和Mobile 必须提供。</li>
	// <li>2. 若UserId 与 Name，Mobile均存在，将优先采用UserId对应的员工。</li>
	// </ul>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 员工姓名，必须与手机号码一起使用。
	// 如果UserId为空，则此字段不能为空。同时，姓名和手机号码必须与传入合同（FlowId）中的签署人信息一致。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 员工手机号，必须与姓名一起使用。
	//  如果UserId为空，则此字段不能为空。同时，姓名和手机号码必须与传入合同（FlowId）中的签署人信息一致。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 为签署方经办人在签署合同中的参与方ID，必须与参数FlowIds数组一一对应。
	// 您可以通过查询合同接口（DescribeFlowInfo）查询此参数。
	// 若传了此参数，则可以不传 UserId, Name, Mobile等参数
	RecipientIds []*string `json:"RecipientIds,omitnil,omitempty" name:"RecipientIds"`
}

func (r *CreateOrganizationBatchSignUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationBatchSignUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowIds")
	delete(f, "Agent")
	delete(f, "UserId")
	delete(f, "Name")
	delete(f, "Mobile")
	delete(f, "RecipientIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationBatchSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationBatchSignUrlResponseParams struct {
	// 批量签署入口链接，用户可使用这个链接跳转到控制台页面对合同进行签署操作。
	SignUrl *string `json:"SignUrl,omitnil,omitempty" name:"SignUrl"`

	// 链接过期截止时间，格式为Unix标准时间戳（秒），默认为7天后截止。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationBatchSignUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationBatchSignUrlResponseParams `json:"Response"`
}

func (r *CreateOrganizationBatchSignUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationBatchSignUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationGroupInvitationLinkRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。 支持填入集团子公司经办人 userId 代发合同。  注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 到期时间（以秒为单位的时间戳），其上限为30天的有效期限。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type CreateOrganizationGroupInvitationLinkRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。 支持填入集团子公司经办人 userId 代发合同。  注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 到期时间（以秒为单位的时间戳），其上限为30天的有效期限。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

func (r *CreateOrganizationGroupInvitationLinkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationGroupInvitationLinkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "ExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationGroupInvitationLinkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationGroupInvitationLinkResponseParams struct {
	// 加入集团二维码链接，子企业的管理员可以直接扫码进入。
	// 注意:1. 该链接有效期时间为ExpireTime，同时需要注意保密，不要外泄给无关用户。2. 该链接不支持小程序嵌入，仅支持<b>移动端浏览器</b>打开。3. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	Link *string `json:"Link,omitnil,omitempty" name:"Link"`

	// 到期时间（以秒为单位的时间戳）
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 加入集团短链接。
	// 注意:
	// 1. 该链接有效期时间为ExpireTime，同时需要注意保密，不要外泄给无关用户。
	// 2. 该链接不支持小程序嵌入，仅支持<b>移动端浏览器</b>打开。
	// 3. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 腾讯电子签小程序加入集团链接。
	// 
	// <li>小程序和APP集成使用</li>
	// <li>得到的链接类似于`pages/guide?shortKey=yDw***k1xFc5`, 用法可以参考：<a href="https://qian.tencent.com/developers/company/openwxminiprogram" target="_blank">跳转电子签小程序</a></li>
	// 
	// 
	// 注： <font color="red">生成的链路后面不能再增加参数</font>
	MiniAppPath *string `json:"MiniAppPath,omitnil,omitempty" name:"MiniAppPath"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationGroupInvitationLinkResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationGroupInvitationLinkResponseParams `json:"Response"`
}

func (r *CreateOrganizationGroupInvitationLinkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationGroupInvitationLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationInfoChangeUrlRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 企业信息变更类型，可选类型如下：
	// <ul><li>**1**：企业超管变更</li><li>**2**：企业基础信息变更</li></ul>
	ChangeType *uint64 `json:"ChangeType,omitnil,omitempty" name:"ChangeType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CreateOrganizationInfoChangeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 企业信息变更类型，可选类型如下：
	// <ul><li>**1**：企业超管变更</li><li>**2**：企业基础信息变更</li></ul>
	ChangeType *uint64 `json:"ChangeType,omitnil,omitempty" name:"ChangeType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *CreateOrganizationInfoChangeUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationInfoChangeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "ChangeType")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationInfoChangeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationInfoChangeUrlResponseParams struct {
	// 创建的企业信息变更链接。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 链接过期时间。链接7天有效。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationInfoChangeUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationInfoChangeUrlResponseParams `json:"Response"`
}

func (r *CreateOrganizationInfoChangeUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationInfoChangeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePartnerAutoSignAuthUrlRequestParams struct {
	// 代理企业和员工的信息。<br/>在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 执行本接口操作的员工信息。<br/>注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 被授企业id/授权方企业id（即OrganizationId），和AuthorizedOrganizationName二选一传入
	AuthorizedOrganizationId *string `json:"AuthorizedOrganizationId,omitnil,omitempty" name:"AuthorizedOrganizationId"`

	// 被授企业名称/授权方企业的名字，和AuthorizedOrganizationId二选一传入即可。请确认该名称与企业营业执照中注册的名称一致。
	// 注: `如果名称中包含英文括号()，请使用中文括号（）代替。`
	AuthorizedOrganizationName *string `json:"AuthorizedOrganizationName,omitnil,omitempty" name:"AuthorizedOrganizationName"`

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
	
	// 代理企业和员工的信息。<br/>在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 执行本接口操作的员工信息。<br/>注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 被授企业id/授权方企业id（即OrganizationId），和AuthorizedOrganizationName二选一传入
	AuthorizedOrganizationId *string `json:"AuthorizedOrganizationId,omitnil,omitempty" name:"AuthorizedOrganizationId"`

	// 被授企业名称/授权方企业的名字，和AuthorizedOrganizationId二选一传入即可。请确认该名称与企业营业执照中注册的名称一致。
	// 注: `如果名称中包含英文括号()，请使用中文括号（）代替。`
	AuthorizedOrganizationName *string `json:"AuthorizedOrganizationName,omitnil,omitempty" name:"AuthorizedOrganizationName"`

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
	delete(f, "Operator")
	delete(f, "AuthorizedOrganizationId")
	delete(f, "AuthorizedOrganizationName")
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
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 个人用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 证件类型，支持以下类型
	// <ul><li> ID_CARD  : 中国大陆居民身份证 (默认值)</li>
	// <li> HONGKONG_AND_MACAO  : 中国港澳居民来往内地通行证</li>
	// <li> HONGKONG_MACAO_AND_TAIWAN  : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	// 
	// 注: `不传默认为处方单场景，即E_PRESCRIPTION_AUTO_SIGN`
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`
}

type CreatePersonAuthCertificateImageRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 个人用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 证件类型，支持以下类型
	// <ul><li> ID_CARD  : 中国大陆居民身份证 (默认值)</li>
	// <li> HONGKONG_AND_MACAO  : 中国港澳居民来往内地通行证</li>
	// <li> HONGKONG_MACAO_AND_TAIWAN  : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	// 
	// 注: `不传默认为处方单场景，即E_PRESCRIPTION_AUTO_SIGN`
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
	delete(f, "Operator")
	delete(f, "UserName")
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "Agent")
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

	// 个人用户认证证书的编号, 为20位数字组成的字符串,  由腾讯电子签下发此编号 。
	// 该编号会合成到个人用户证书证明图片。
	// 
	// 注: `个人用户认证证书的编号和证明图片绑定, 获取新的证明图片编号会变动`
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageCertId *string `json:"ImageCertId,omitnil,omitempty" name:"ImageCertId"`

	// 在数字证书申请过程中，系统会自动生成一个独一无二的序列号。请注意，当证书到期并自动续期时，该序列号将会发生变化。值得注意的是，此序列号不会被合成至个人用户证书的证明图片中。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerialNumber *string `json:"SerialNumber,omitnil,omitempty" name:"SerialNumber"`

	// CA证书颁发时间，格式为Unix标准时间戳（秒）   
	// 该时间格式化后会合成到个人用户证书证明图片
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidFrom *uint64 `json:"ValidFrom,omitnil,omitempty" name:"ValidFrom"`

	// CA证书有效截止时间，格式为Unix标准时间戳（秒）
	// 该时间格式化后会合成到个人用户证书证明图片
	// 注意：此字段可能返回 null，表示取不到有效值。
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
type CreatePrepareFlowRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 资源id，与ResourceType相对应，取值范围：
	// <ul>
	// <li>文件Id（通过UploadFiles获取文件资源Id）</li>
	// <li>模板Id（通过控制台创建模板后获取模板Id）</li>
	// </ul>
	// 注意：需要同时设置 ResourceType 参数指定资源类型
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 自定义的合同流程的名称，长度不能超过200个字符，只能由中文汉字、中文标点、英文字母、阿拉伯数字、空格、小括号、中括号、中划线、下划线以及（,）、（;）、（.）、(&)、（+）组成。
	// 
	// 该名称还将用于合同签署完成后文件下载的默认文件名称。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 资源类型，取值有：
	// <ul><li> **1**：模板</li>
	// <li> **2**：文件（默认值）</li></ul>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	Unordered *bool `json:"Unordered,omitnil,omitempty" name:"Unordered"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 用户自定义合同类型Id
	// 
	// 该id为电子签企业内的合同类型id， 可以在控制台-合同-自定义合同类型处获取
	// 注: `该参数如果和FlowType同时传，以该参数优先生效`
	UserFlowTypeId *string `json:"UserFlowTypeId,omitnil,omitempty" name:"UserFlowTypeId"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 合同流程的参与方列表，最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的ApproverInfo结构体。
	// 
	// 如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序，请确保列表中参与人的顺序符合实际签署顺序。
	Approvers []*FlowCreateApprover `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 开启或者关闭智能添加填写区：
	// <ul><li> **OPEN**：开启（默认值）</li>
	// <li> **CLOSE**：关闭</li></ul>
	IntelligentStatus *string `json:"IntelligentStatus,omitnil,omitempty" name:"IntelligentStatus"`

	// 该字段已废弃，请使用InitiatorComponents
	Components *Component `json:"Components,omitnil,omitempty" name:"Components"`

	// 发起合同个性化参数
	// 用于满足创建及页面操作过程中的个性化要求
	// 具体定制化内容详见数据接口说明
	FlowOption *CreateFlowOption `json:"FlowOption,omitnil,omitempty" name:"FlowOption"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil,omitempty" name:"NeedSignReview"`

	// 发起方企业的签署人进行发起操作是否需要企业内部审批。使用此功能需要发起方企业有参与签署。
	// 
	// 若设置为true，发起审核结果需通过接口 CreateFlowSignReview 通知电子签，审核通过后，发起方企业签署人方可进行发起操作，否则会阻塞其发起操作。
	// 
	NeedCreateReview *bool `json:"NeedCreateReview,omitnil,omitempty" name:"NeedCreateReview"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	CcInfos []*CcInfo `json:"CcInfos,omitnil,omitempty" name:"CcInfos"`

	// 合同Id：用于通过一个已发起的合同快速生成一个发起流程web链接
	// 注: `该参数必须是一个待发起审核的合同id，并且还未审核通过`
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 模板或者合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	InitiatorComponents []*Component `json:"InitiatorComponents,omitnil,omitempty" name:"InitiatorComponents"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下：  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li></ul>效果如下:![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"FlowDisplayType,omitnil,omitempty" name:"FlowDisplayType"`
}

type CreatePrepareFlowRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 资源id，与ResourceType相对应，取值范围：
	// <ul>
	// <li>文件Id（通过UploadFiles获取文件资源Id）</li>
	// <li>模板Id（通过控制台创建模板后获取模板Id）</li>
	// </ul>
	// 注意：需要同时设置 ResourceType 参数指定资源类型
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 自定义的合同流程的名称，长度不能超过200个字符，只能由中文汉字、中文标点、英文字母、阿拉伯数字、空格、小括号、中括号、中划线、下划线以及（,）、（;）、（.）、(&)、（+）组成。
	// 
	// 该名称还将用于合同签署完成后文件下载的默认文件名称。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 资源类型，取值有：
	// <ul><li> **1**：模板</li>
	// <li> **2**：文件（默认值）</li></ul>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	Unordered *bool `json:"Unordered,omitnil,omitempty" name:"Unordered"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 用户自定义合同类型Id
	// 
	// 该id为电子签企业内的合同类型id， 可以在控制台-合同-自定义合同类型处获取
	// 注: `该参数如果和FlowType同时传，以该参数优先生效`
	UserFlowTypeId *string `json:"UserFlowTypeId,omitnil,omitempty" name:"UserFlowTypeId"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 合同流程的参与方列表，最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的ApproverInfo结构体。
	// 
	// 如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序，请确保列表中参与人的顺序符合实际签署顺序。
	Approvers []*FlowCreateApprover `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 开启或者关闭智能添加填写区：
	// <ul><li> **OPEN**：开启（默认值）</li>
	// <li> **CLOSE**：关闭</li></ul>
	IntelligentStatus *string `json:"IntelligentStatus,omitnil,omitempty" name:"IntelligentStatus"`

	// 该字段已废弃，请使用InitiatorComponents
	Components *Component `json:"Components,omitnil,omitempty" name:"Components"`

	// 发起合同个性化参数
	// 用于满足创建及页面操作过程中的个性化要求
	// 具体定制化内容详见数据接口说明
	FlowOption *CreateFlowOption `json:"FlowOption,omitnil,omitempty" name:"FlowOption"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil,omitempty" name:"NeedSignReview"`

	// 发起方企业的签署人进行发起操作是否需要企业内部审批。使用此功能需要发起方企业有参与签署。
	// 
	// 若设置为true，发起审核结果需通过接口 CreateFlowSignReview 通知电子签，审核通过后，发起方企业签署人方可进行发起操作，否则会阻塞其发起操作。
	// 
	NeedCreateReview *bool `json:"NeedCreateReview,omitnil,omitempty" name:"NeedCreateReview"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	CcInfos []*CcInfo `json:"CcInfos,omitnil,omitempty" name:"CcInfos"`

	// 合同Id：用于通过一个已发起的合同快速生成一个发起流程web链接
	// 注: `该参数必须是一个待发起审核的合同id，并且还未审核通过`
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 模板或者合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	InitiatorComponents []*Component `json:"InitiatorComponents,omitnil,omitempty" name:"InitiatorComponents"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下：  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li></ul>效果如下:![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"FlowDisplayType,omitnil,omitempty" name:"FlowDisplayType"`
}

func (r *CreatePrepareFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrepareFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "ResourceId")
	delete(f, "FlowName")
	delete(f, "ResourceType")
	delete(f, "Unordered")
	delete(f, "Deadline")
	delete(f, "UserFlowTypeId")
	delete(f, "FlowType")
	delete(f, "Approvers")
	delete(f, "IntelligentStatus")
	delete(f, "Components")
	delete(f, "FlowOption")
	delete(f, "NeedSignReview")
	delete(f, "NeedCreateReview")
	delete(f, "UserData")
	delete(f, "CcInfos")
	delete(f, "FlowId")
	delete(f, "Agent")
	delete(f, "InitiatorComponents")
	delete(f, "FlowDisplayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrepareFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrepareFlowResponseParams struct {
	// 发起流程的web页面链接，有效期5分钟
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 创建的合同id（还未实际发起），每次调用会生成新的id，用户可以记录此字段对应后续页面发起的合同，若在页面上未成功发起，则此字段无效。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrepareFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrepareFlowResponseParams `json:"Response"`
}

func (r *CreatePrepareFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrepareFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePreparedPersonalEsignRequestParams struct {
	// 个人用户姓名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 证件号码，应符合以下规则
	// <ul><li> 中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 印章名称，长度1-50个字。
	SealName *string `json:"SealName,omitnil,omitempty" name:"SealName"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同 中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 印章图片的base64
	// 注：已废弃
	// 请先通过UploadFiles接口上传文件，获取 FileId
	//
	// Deprecated: SealImage is deprecated.
	SealImage *string `json:"SealImage,omitnil,omitempty" name:"SealImage"`

	// 是否开启印章图片压缩处理，默认不开启，如需开启请设置为 true。当印章超过 2M 时建议开启，开启后图片的 hash 将发生变化。
	SealImageCompress *bool `json:"SealImageCompress,omitnil,omitempty" name:"SealImageCompress"`

	// 手机号码；当需要开通自动签时，该参数必传
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 此字段已废弃，请勿继续使用。
	EnableAutoSign *bool `json:"EnableAutoSign,omitnil,omitempty" name:"EnableAutoSign"`

	// 印章颜色（参数ProcessSeal=true时生效）
	// 默认值：BLACK黑色
	// 取值: 
	// BLACK 黑色,
	// RED 红色,
	// BLUE 蓝色。
	SealColor *string `json:"SealColor,omitnil,omitempty" name:"SealColor"`

	// 是否处理印章，默认不做印章处理。
	// 取值如下：
	// <ul>
	// <li>false：不做任何处理；</li>
	// <li>true：做透明化处理和颜色增强。</li>
	// </ul>
	ProcessSeal *bool `json:"ProcessSeal,omitnil,omitempty" name:"ProcessSeal"`

	// 印章图片文件 id
	// 取值：
	// 填写的FileId通过UploadFiles接口上传文件获取。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。一旦绑定后，将扣减购买的个人自动签账号许可一次（1年有效期），不可解绑释放。不传默认为绑定自动签账号许可。 0-绑定个人自动签账号许可，开通后将扣减购买的个人自动签账号许可一次 1-不绑定，发起合同时将按标准合同套餐进行扣减	
	LicenseType *int64 `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	// 
	// 注: `不传默认为处方单场景，即E_PRESCRIPTION_AUTO_SIGN`
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`
}

type CreatePreparedPersonalEsignRequest struct {
	*tchttp.BaseRequest
	
	// 个人用户姓名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 证件号码，应符合以下规则
	// <ul><li> 中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 印章名称，长度1-50个字。
	SealName *string `json:"SealName,omitnil,omitempty" name:"SealName"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同 中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 印章图片的base64
	// 注：已废弃
	// 请先通过UploadFiles接口上传文件，获取 FileId
	SealImage *string `json:"SealImage,omitnil,omitempty" name:"SealImage"`

	// 是否开启印章图片压缩处理，默认不开启，如需开启请设置为 true。当印章超过 2M 时建议开启，开启后图片的 hash 将发生变化。
	SealImageCompress *bool `json:"SealImageCompress,omitnil,omitempty" name:"SealImageCompress"`

	// 手机号码；当需要开通自动签时，该参数必传
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 此字段已废弃，请勿继续使用。
	EnableAutoSign *bool `json:"EnableAutoSign,omitnil,omitempty" name:"EnableAutoSign"`

	// 印章颜色（参数ProcessSeal=true时生效）
	// 默认值：BLACK黑色
	// 取值: 
	// BLACK 黑色,
	// RED 红色,
	// BLUE 蓝色。
	SealColor *string `json:"SealColor,omitnil,omitempty" name:"SealColor"`

	// 是否处理印章，默认不做印章处理。
	// 取值如下：
	// <ul>
	// <li>false：不做任何处理；</li>
	// <li>true：做透明化处理和颜色增强。</li>
	// </ul>
	ProcessSeal *bool `json:"ProcessSeal,omitnil,omitempty" name:"ProcessSeal"`

	// 印章图片文件 id
	// 取值：
	// 填写的FileId通过UploadFiles接口上传文件获取。
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。一旦绑定后，将扣减购买的个人自动签账号许可一次（1年有效期），不可解绑释放。不传默认为绑定自动签账号许可。 0-绑定个人自动签账号许可，开通后将扣减购买的个人自动签账号许可一次 1-不绑定，发起合同时将按标准合同套餐进行扣减	
	LicenseType *int64 `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	// 
	// 注: `不传默认为处方单场景，即E_PRESCRIPTION_AUTO_SIGN`
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`
}

func (r *CreatePreparedPersonalEsignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePreparedPersonalEsignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "IdCardNumber")
	delete(f, "SealName")
	delete(f, "Operator")
	delete(f, "IdCardType")
	delete(f, "SealImage")
	delete(f, "SealImageCompress")
	delete(f, "Mobile")
	delete(f, "EnableAutoSign")
	delete(f, "SealColor")
	delete(f, "ProcessSeal")
	delete(f, "FileId")
	delete(f, "Agent")
	delete(f, "LicenseType")
	delete(f, "SceneKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePreparedPersonalEsignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePreparedPersonalEsignResponseParams struct {
	// 导入生成的印章ID，为32位字符串。
	// 建议开发者保存此印章ID，开头实名认证后，通过此 ID查询导入的印章。
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePreparedPersonalEsignResponse struct {
	*tchttp.BaseResponse
	Response *CreatePreparedPersonalEsignResponseParams `json:"Response"`
}

func (r *CreatePreparedPersonalEsignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePreparedPersonalEsignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseFlowRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 待解除的签署流程编号（即原签署流程的编号）。
	NeedRelievedFlowId *string `json:"NeedRelievedFlowId,omitnil,omitempty" name:"NeedRelievedFlowId"`

	// 解除协议内容, 包括解除理由等信息。
	ReliveInfo *RelieveInfo `json:"ReliveInfo,omitnil,omitempty" name:"ReliveInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 替换解除协议的签署人， 如不指定新的签署人，将继续使用原流程的签署人作为本解除协议的参与方。 <br/>
	// 如需更换原合同中的企业端签署人，可通过指定该签署人的RecipientId编号更换此企业端签署人。(可通过接口<a href="https://qian.tencent.com/developers/companyApis/queryFlows/DescribeFlowInfo/">DescribeFlowInfo</a>查询签署人的RecipientId编号)<br/>
	// 
	// 注：
	// 1. 支持更换企业的签署人，不支持更换个人类型的签署人。
	// 2. 己方企业支持自动签署，他方企业不支持自动签署。
	// 3. <b>仅将需要替换的签署人添加至此列表</b>，无需替换的签署人无需添加进来。
	ReleasedApprovers []*ReleasedApprover `json:"ReleasedApprovers,omitnil,omitempty" name:"ReleasedApprovers"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的7天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 调用方自定义的个性化字段，该字段的值可以是字符串JSON或其他字符串形式，客户可以根据自身需求自定义数据格式并在需要时进行解析。该字段的信息将以Base64编码的形式传输，支持的最大数据大小为20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

type CreateReleaseFlowRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 待解除的签署流程编号（即原签署流程的编号）。
	NeedRelievedFlowId *string `json:"NeedRelievedFlowId,omitnil,omitempty" name:"NeedRelievedFlowId"`

	// 解除协议内容, 包括解除理由等信息。
	ReliveInfo *RelieveInfo `json:"ReliveInfo,omitnil,omitempty" name:"ReliveInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 替换解除协议的签署人， 如不指定新的签署人，将继续使用原流程的签署人作为本解除协议的参与方。 <br/>
	// 如需更换原合同中的企业端签署人，可通过指定该签署人的RecipientId编号更换此企业端签署人。(可通过接口<a href="https://qian.tencent.com/developers/companyApis/queryFlows/DescribeFlowInfo/">DescribeFlowInfo</a>查询签署人的RecipientId编号)<br/>
	// 
	// 注：
	// 1. 支持更换企业的签署人，不支持更换个人类型的签署人。
	// 2. 己方企业支持自动签署，他方企业不支持自动签署。
	// 3. <b>仅将需要替换的签署人添加至此列表</b>，无需替换的签署人无需添加进来。
	ReleasedApprovers []*ReleasedApprover `json:"ReleasedApprovers,omitnil,omitempty" name:"ReleasedApprovers"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的7天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 调用方自定义的个性化字段，该字段的值可以是字符串JSON或其他字符串形式，客户可以根据自身需求自定义数据格式并在需要时进行解析。该字段的信息将以Base64编码的形式传输，支持的最大数据大小为20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

func (r *CreateReleaseFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "NeedRelievedFlowId")
	delete(f, "ReliveInfo")
	delete(f, "Agent")
	delete(f, "ReleasedApprovers")
	delete(f, "Deadline")
	delete(f, "UserData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReleaseFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseFlowResponseParams struct {
	// 解除协议流程编号
	// `注意：这里的流程编号对应的合同是本次发起的解除协议。`
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReleaseFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreateReleaseFlowResponseParams `json:"Response"`
}

func (r *CreateReleaseFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResultPageConfig struct {
	// <ul>
	//   <li>0 : 发起审批成功页面（通过接口<a href="https://qian.tencent.com/developers/companyApis/embedPages/CreatePrepareFlow/" target="_blank">创建发起流程web页面</a>发起时设置了NeedCreateReview参数为true）</li>
	// </ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 结果页标题，不超过50字
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 结果页描述，不超过200字
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type CreateSchemeUrlRequestParams struct {
	// 执行本接口操作的员工信息, userId 必填。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程签署方的组织机构名称。如果名称中包含英文括号()，请使用中文括号（）代替。注: `获取B端动态签署人领取链接时,可指定此字段来预先设定签署人的企业,预设后只能以该企业身份去领取合同并完成签署`
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 合同流程里边签署方经办人的姓名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 合同流程里边签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成(如存在X，请大写)。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 要跳转的链接类型
	// 
	// <ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li>
	// <li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li>
	// <li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li></ul>
	EndPoint *string `json:"EndPoint,omitnil,omitempty" name:"EndPoint"`

	// 合同流程ID 
	// 注: `如果准备跳转到合同流程签署的详情页面(即PathType=1时)必传,   跳转其他页面可不传`
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 合同流程组的组ID, 在合同流程组场景下，生成合同流程组的签署链接时需要赋值
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 要跳转到的页面类型 <ul><li> **0** : 腾讯电子签小程序个人首页 (默认) <a href="https://qcloudimg.tencent-cloud.cn/raw/a2667ea84ec993cc060321afe3191d65.jpg" target="_blank" >点击查看示例页面</a></li><li> **1** : 腾讯电子签小程序流程合同的详情页,即合同签署页面(注:`仅生成短链或者小程序Path时支持跳转合同详情页,即EndPoint传HTTP_SHORT_URL或者APP`)<a href="https://qcloudimg.tencent-cloud.cn/raw/446a679f09b1b7f40eb84e67face8acc.jpg" target="_blank" >点击查看示例页面</a></li><li> **2** : 腾讯电子签小程序合同列表页 <a href="https://qcloudimg.tencent-cloud.cn/raw/c7b80e44c1d68ae3270a6fc4939c7214.jpg" target="_blank" >点击查看示例页面</a> </li><li> **3** : 腾讯电子签小程序合同封面页  （注：`生成动态签署人补充链接时，必须指定为封面页`）<a href="https://qcloudimg.tencent-cloud.cn/raw/0d22cc587be4bf084877c151350c3bf7.jpg" target="_blank" >点击查看示例页面</a></li></ul>
	PathType *uint64 `json:"PathType,omitnil,omitempty" name:"PathType"`

	// 签署完成后是否自动回跳
	// <ul><li>**false**：否, 签署完成不会自动跳转回来(默认)</li><li>**true**：是, 签署完成会自动跳转回来</li></ul>
	// 
	// 注: 
	// 1. 该参数<font color="red">只针对APP类型（电子签小程序跳转贵方小程序）场景</font> 的签署链接有效
	// 2. <font color="red">手机应用APP 或 微信小程序需要监控界面的返回走后序逻辑</font>, 微信小程序的文档可以参考[这个](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onShow-Object-object)
	// 3. <font color="red">电子签小程序跳转贵方APP，不支持自动跳转，必需用户手动点击完成按钮（微信的限制）</font> 
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil,omitempty" name:"AutoJumpBack"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 生成的签署链接在签署页面隐藏的按钮列表，可设置如下：
	// 
	// <ul><li> **0** :合同签署页面更多操作按钮</li>
	// <li> **1** :合同签署页面更多操作的拒绝签署按钮</li>
	// <li> **2** :合同签署页面更多操作的转他人处理按钮</li>
	// <li> **3** :签署成功页的查看详情按钮</li></ul>
	// 
	// 注:  `字段为数组, 可以传值隐藏多个按钮`
	Hides []*int64 `json:"Hides,omitnil,omitempty" name:"Hides"`

	// 签署节点ID，用于生成动态签署人链接完成领取。
	// 
	// 注：`生成动态签署人补充链接时必传。`
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 合同组相关信息，指定合同组子合同和签署方的信息，用于补充动态签署人。
	FlowGroupUrlInfo *FlowGroupUrlInfo `json:"FlowGroupUrlInfo,omitnil,omitempty" name:"FlowGroupUrlInfo"`

	// <font color="red">仅公众号 H5 跳转电子签小程序时</font>，如需签署完成的“返回应用”功能，在获取签署链接接口的 UrlUseEnv 参数需设置为 **WeChatOfficialAccounts**，小程序签署成功的结果页面中才会出现“返回应用”按钮。在用户点击“返回应用”按钮之后，会返回到公众号 H5。 
	// 
	// 参考 [公众号 H5 跳转电子签小程序](https://qian.tencent.com/developers/company/openwxminiprogram/#23-%E5%85%AC%E4%BC%97%E5%8F%B7-h5-%E4%B8%AD%E8%B7%B3%E8%BD%AC)。
	UrlUseEnv *string `json:"UrlUseEnv,omitnil,omitempty" name:"UrlUseEnv"`
}

type CreateSchemeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息, userId 必填。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程签署方的组织机构名称。如果名称中包含英文括号()，请使用中文括号（）代替。注: `获取B端动态签署人领取链接时,可指定此字段来预先设定签署人的企业,预设后只能以该企业身份去领取合同并完成签署`
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 合同流程里边签署方经办人的姓名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 合同流程里边签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成(如存在X，请大写)。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 要跳转的链接类型
	// 
	// <ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li>
	// <li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li>
	// <li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li></ul>
	EndPoint *string `json:"EndPoint,omitnil,omitempty" name:"EndPoint"`

	// 合同流程ID 
	// 注: `如果准备跳转到合同流程签署的详情页面(即PathType=1时)必传,   跳转其他页面可不传`
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 合同流程组的组ID, 在合同流程组场景下，生成合同流程组的签署链接时需要赋值
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 要跳转到的页面类型 <ul><li> **0** : 腾讯电子签小程序个人首页 (默认) <a href="https://qcloudimg.tencent-cloud.cn/raw/a2667ea84ec993cc060321afe3191d65.jpg" target="_blank" >点击查看示例页面</a></li><li> **1** : 腾讯电子签小程序流程合同的详情页,即合同签署页面(注:`仅生成短链或者小程序Path时支持跳转合同详情页,即EndPoint传HTTP_SHORT_URL或者APP`)<a href="https://qcloudimg.tencent-cloud.cn/raw/446a679f09b1b7f40eb84e67face8acc.jpg" target="_blank" >点击查看示例页面</a></li><li> **2** : 腾讯电子签小程序合同列表页 <a href="https://qcloudimg.tencent-cloud.cn/raw/c7b80e44c1d68ae3270a6fc4939c7214.jpg" target="_blank" >点击查看示例页面</a> </li><li> **3** : 腾讯电子签小程序合同封面页  （注：`生成动态签署人补充链接时，必须指定为封面页`）<a href="https://qcloudimg.tencent-cloud.cn/raw/0d22cc587be4bf084877c151350c3bf7.jpg" target="_blank" >点击查看示例页面</a></li></ul>
	PathType *uint64 `json:"PathType,omitnil,omitempty" name:"PathType"`

	// 签署完成后是否自动回跳
	// <ul><li>**false**：否, 签署完成不会自动跳转回来(默认)</li><li>**true**：是, 签署完成会自动跳转回来</li></ul>
	// 
	// 注: 
	// 1. 该参数<font color="red">只针对APP类型（电子签小程序跳转贵方小程序）场景</font> 的签署链接有效
	// 2. <font color="red">手机应用APP 或 微信小程序需要监控界面的返回走后序逻辑</font>, 微信小程序的文档可以参考[这个](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onShow-Object-object)
	// 3. <font color="red">电子签小程序跳转贵方APP，不支持自动跳转，必需用户手动点击完成按钮（微信的限制）</font> 
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil,omitempty" name:"AutoJumpBack"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 生成的签署链接在签署页面隐藏的按钮列表，可设置如下：
	// 
	// <ul><li> **0** :合同签署页面更多操作按钮</li>
	// <li> **1** :合同签署页面更多操作的拒绝签署按钮</li>
	// <li> **2** :合同签署页面更多操作的转他人处理按钮</li>
	// <li> **3** :签署成功页的查看详情按钮</li></ul>
	// 
	// 注:  `字段为数组, 可以传值隐藏多个按钮`
	Hides []*int64 `json:"Hides,omitnil,omitempty" name:"Hides"`

	// 签署节点ID，用于生成动态签署人链接完成领取。
	// 
	// 注：`生成动态签署人补充链接时必传。`
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 合同组相关信息，指定合同组子合同和签署方的信息，用于补充动态签署人。
	FlowGroupUrlInfo *FlowGroupUrlInfo `json:"FlowGroupUrlInfo,omitnil,omitempty" name:"FlowGroupUrlInfo"`

	// <font color="red">仅公众号 H5 跳转电子签小程序时</font>，如需签署完成的“返回应用”功能，在获取签署链接接口的 UrlUseEnv 参数需设置为 **WeChatOfficialAccounts**，小程序签署成功的结果页面中才会出现“返回应用”按钮。在用户点击“返回应用”按钮之后，会返回到公众号 H5。 
	// 
	// 参考 [公众号 H5 跳转电子签小程序](https://qian.tencent.com/developers/company/openwxminiprogram/#23-%E5%85%AC%E4%BC%97%E5%8F%B7-h5-%E4%B8%AD%E8%B7%B3%E8%BD%AC)。
	UrlUseEnv *string `json:"UrlUseEnv,omitnil,omitempty" name:"UrlUseEnv"`
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
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "EndPoint")
	delete(f, "FlowId")
	delete(f, "FlowGroupId")
	delete(f, "PathType")
	delete(f, "AutoJumpBack")
	delete(f, "Agent")
	delete(f, "Hides")
	delete(f, "RecipientId")
	delete(f, "FlowGroupUrlInfo")
	delete(f, "UrlUseEnv")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchemeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchemeUrlResponseParams struct {
	// 腾讯电子签小程序的签署链接。
	// 
	// <ul><li>如果EndPoint是**APP**，得到的链接类似于`pages/guide?from=default&where=mini&id=yDwJSUUirqauh***7jNSxwdirTSGuH&to=CONTRACT_DETAIL&name=&phone=&shortKey=yDw***k1xFc5`, 用法可以参加接口描述中的"跳转到小程序的实现"</li>
	// <li>如果EndPoint是**HTTP**，得到的链接类似于 `https://res.ess.tencent.cn/cdn/h5-activity/jump-mp.html?where=mini&from=SFY&id=yDwfEUUw**4rV6Avz&to=MVP_CONTRACT_COVER&name=%E9%83%**5%86%9B`，点击后会跳转到腾讯电子签小程序进行签署</li>
	// <li>如果EndPoint是**HTTP_SHORT_URL**，得到的链接类似于 `https://essurl.cn/2n**42Nd`，点击后会跳转到腾讯电子签小程序进行签署</li></ul>
	// 
	// 
	// 注： <font color="red">生成的链路后面不能再增加参数</font>
	SchemeUrl *string `json:"SchemeUrl,omitnil,omitempty" name:"SchemeUrl"`

	// 二维码，在生成动态签署人跳转封面页链接时返回
	SchemeQrcodeUrl *string `json:"SchemeQrcodeUrl,omitnil,omitempty" name:"SchemeQrcodeUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateSealPolicyRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 用户在电子文件签署平台标识信息，具体参考UserInfo结构体。可跟下面的UserIds可叠加起作用
	Users []*UserInfo `json:"Users,omitnil,omitempty" name:"Users"`

	// 电子印章ID，为32位字符串。
	// 建议开发者保留此印章ID，后续指定签署区印章或者操作印章需此印章ID。
	// 可登录腾讯电子签控制台，在 "印章"->"印章中心"选择查看的印章，在"印章详情" 中查看某个印章的SealId(在页面中展示为印章ID)。
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 授权有效期，时间戳秒级。可以传0，代表有效期到2099年12月12日23点59分59秒。
	Expired *int64 `json:"Expired,omitnil,omitempty" name:"Expired"`

	// 需要授权的用户UserId集合。跟上面的SealId参数配合使用。选填，跟上面的Users同时起作用
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 印章授权内容
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CreateSealPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 用户在电子文件签署平台标识信息，具体参考UserInfo结构体。可跟下面的UserIds可叠加起作用
	Users []*UserInfo `json:"Users,omitnil,omitempty" name:"Users"`

	// 电子印章ID，为32位字符串。
	// 建议开发者保留此印章ID，后续指定签署区印章或者操作印章需此印章ID。
	// 可登录腾讯电子签控制台，在 "印章"->"印章中心"选择查看的印章，在"印章详情" 中查看某个印章的SealId(在页面中展示为印章ID)。
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 授权有效期，时间戳秒级。可以传0，代表有效期到2099年12月12日23点59分59秒。
	Expired *int64 `json:"Expired,omitnil,omitempty" name:"Expired"`

	// 需要授权的用户UserId集合。跟上面的SealId参数配合使用。选填，跟上面的Users同时起作用
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 印章授权内容
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *CreateSealPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Users")
	delete(f, "SealId")
	delete(f, "Expired")
	delete(f, "UserIds")
	delete(f, "Policy")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSealPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealPolicyResponseParams struct {
	// 最终授权成功的用户ID，在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查看某位员工的UserId(在页面中展示为用户ID)。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSealPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateSealPolicyResponseParams `json:"Response"`
}

func (r *CreateSealPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 电子印章名字，1-50个中文字符
	// 注:`同一企业下电子印章名字不能相同`
	SealName *string `json:"SealName,omitnil,omitempty" name:"SealName"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 电子印章生成方式
	// <ul>
	// <li><strong>空值</strong>:(默认)使用上传的图片生成印章, 此时需要上传SealImage图片</li>
	// <li><strong>SealGenerateSourceSystem</strong>: 系统生成印章, 无需上传SealImage图片</li>
	// </ul>
	GenerateSource *string `json:"GenerateSource,omitnil,omitempty" name:"GenerateSource"`

	// 电子印章类型 , 可选类型如下: <ul><li>**OFFICIAL**: (默认)公章</li><li>**CONTRACT**: 合同专用章;</li><li>**FINANCE**: 财务专用章;</li><li>**PERSONNEL**: 人事专用章</li><li>**INVOICE**: 发票专用章</li><li>**OTHER**: 其他</li></ul>注: 同企业下只能有<font color="red">一个</font>公章, 重复创建会报错
	SealType *string `json:"SealType,omitnil,omitempty" name:"SealType"`

	// 电子印章图片文件名称，1-50个中文字符。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 电子印章图片base64编码，大小不超过10M（原始图片不超过5M），只支持PNG或JPG图片格式
	// 
	// 注: `通过图片创建的电子印章，需电子签平台人工审核`
	// 
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 电子印章宽度,单位px
	// 参数不再启用，系统会设置印章大小为标准尺寸。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 电子印章高度,单位px
	// 参数不再启用，系统会设置印章大小为标准尺寸。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 电子印章印章颜色(默认红色RED),RED-红色
	// 
	// 系统目前只支持红色印章创建。
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`

	// 企业印章横向文字，最多可填15个汉字  (若超过印章最大宽度，优先压缩字间距，其次缩小字号)
	// 横向文字的位置如下图中的"印章横向文字在这里"
	// 
	// ![image](https://dyn.ess.tencent.cn/guide/capi/CreateSealByImage2.png)
	SealHorizontalText *string `json:"SealHorizontalText,omitnil,omitempty" name:"SealHorizontalText"`

	// 暂时不支持下弦文字设置
	SealChordText *string `json:"SealChordText,omitnil,omitempty" name:"SealChordText"`

	// 系统生成的印章只支持STAR
	SealCentralType *string `json:"SealCentralType,omitnil,omitempty" name:"SealCentralType"`

	// 通过文件上传时，服务端生成的电子印章上传图片的token
	FileToken *string `json:"FileToken,omitnil,omitempty" name:"FileToken"`

	// 印章样式, 可以选择的样式如下: 
	// <ul><li>**circle**:(默认)圆形印章</li>
	// <li>**ellipse**:椭圆印章</li></ul>
	SealStyle *string `json:"SealStyle,omitnil,omitempty" name:"SealStyle"`

	// 印章尺寸取值描述, 可以选择的尺寸如下: <ul><li> **38_38**: 圆形企业公章直径38mm, 当SealStyle是圆形的时候才有效</li> <li> **40_40**: 圆形企业公章直径40mm, 当SealStyle是圆形的时候才有效</li> <li> **42_42**（默认）: 圆形企业公章直径42mm, 当SealStyle是圆形的时候才有效</li> <li> **45_45**: 圆形企业印章直径45mm, 当SealStyle是圆形的时候才有效</li> <li> **50_50**: 圆形企业印章直径45mm, 当SealStyle是圆形的时候才有效</li> <li> **58_58**: 圆形企业印章直径45mm, 当SealStyle是圆形的时候才有效</li>  <li> **40_30**: 椭圆形印章40mm x 30mm, 当SealStyle是椭圆的时候才有效</li> <li> **45_30**: 椭圆形印章45mm x 30mm, 当SealStyle是椭圆的时候才有效</li> </ul>
	SealSize *string `json:"SealSize,omitnil,omitempty" name:"SealSize"`

	// 企业税号
	// 注:
	// <ul>
	// <li>1.印章类型SealType是INVOICE类型时，此参数才会生效</li>
	// <li>2.印章类型SealType是INVOICE类型，且该字段没有传入值或传入空时，会取该企业对应的统一社会信用代码作为默认的企业税号（<font color="red">如果是通过授权书授权方式认证的企业，此参数必传不能为空</font>）</li>
	// </ul>
	TaxIdentifyCode *string `json:"TaxIdentifyCode,omitnil,omitempty" name:"TaxIdentifyCode"`
}

type CreateSealRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 电子印章名字，1-50个中文字符
	// 注:`同一企业下电子印章名字不能相同`
	SealName *string `json:"SealName,omitnil,omitempty" name:"SealName"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 电子印章生成方式
	// <ul>
	// <li><strong>空值</strong>:(默认)使用上传的图片生成印章, 此时需要上传SealImage图片</li>
	// <li><strong>SealGenerateSourceSystem</strong>: 系统生成印章, 无需上传SealImage图片</li>
	// </ul>
	GenerateSource *string `json:"GenerateSource,omitnil,omitempty" name:"GenerateSource"`

	// 电子印章类型 , 可选类型如下: <ul><li>**OFFICIAL**: (默认)公章</li><li>**CONTRACT**: 合同专用章;</li><li>**FINANCE**: 财务专用章;</li><li>**PERSONNEL**: 人事专用章</li><li>**INVOICE**: 发票专用章</li><li>**OTHER**: 其他</li></ul>注: 同企业下只能有<font color="red">一个</font>公章, 重复创建会报错
	SealType *string `json:"SealType,omitnil,omitempty" name:"SealType"`

	// 电子印章图片文件名称，1-50个中文字符。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 电子印章图片base64编码，大小不超过10M（原始图片不超过5M），只支持PNG或JPG图片格式
	// 
	// 注: `通过图片创建的电子印章，需电子签平台人工审核`
	// 
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 电子印章宽度,单位px
	// 参数不再启用，系统会设置印章大小为标准尺寸。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 电子印章高度,单位px
	// 参数不再启用，系统会设置印章大小为标准尺寸。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 电子印章印章颜色(默认红色RED),RED-红色
	// 
	// 系统目前只支持红色印章创建。
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`

	// 企业印章横向文字，最多可填15个汉字  (若超过印章最大宽度，优先压缩字间距，其次缩小字号)
	// 横向文字的位置如下图中的"印章横向文字在这里"
	// 
	// ![image](https://dyn.ess.tencent.cn/guide/capi/CreateSealByImage2.png)
	SealHorizontalText *string `json:"SealHorizontalText,omitnil,omitempty" name:"SealHorizontalText"`

	// 暂时不支持下弦文字设置
	SealChordText *string `json:"SealChordText,omitnil,omitempty" name:"SealChordText"`

	// 系统生成的印章只支持STAR
	SealCentralType *string `json:"SealCentralType,omitnil,omitempty" name:"SealCentralType"`

	// 通过文件上传时，服务端生成的电子印章上传图片的token
	FileToken *string `json:"FileToken,omitnil,omitempty" name:"FileToken"`

	// 印章样式, 可以选择的样式如下: 
	// <ul><li>**circle**:(默认)圆形印章</li>
	// <li>**ellipse**:椭圆印章</li></ul>
	SealStyle *string `json:"SealStyle,omitnil,omitempty" name:"SealStyle"`

	// 印章尺寸取值描述, 可以选择的尺寸如下: <ul><li> **38_38**: 圆形企业公章直径38mm, 当SealStyle是圆形的时候才有效</li> <li> **40_40**: 圆形企业公章直径40mm, 当SealStyle是圆形的时候才有效</li> <li> **42_42**（默认）: 圆形企业公章直径42mm, 当SealStyle是圆形的时候才有效</li> <li> **45_45**: 圆形企业印章直径45mm, 当SealStyle是圆形的时候才有效</li> <li> **50_50**: 圆形企业印章直径45mm, 当SealStyle是圆形的时候才有效</li> <li> **58_58**: 圆形企业印章直径45mm, 当SealStyle是圆形的时候才有效</li>  <li> **40_30**: 椭圆形印章40mm x 30mm, 当SealStyle是椭圆的时候才有效</li> <li> **45_30**: 椭圆形印章45mm x 30mm, 当SealStyle是椭圆的时候才有效</li> </ul>
	SealSize *string `json:"SealSize,omitnil,omitempty" name:"SealSize"`

	// 企业税号
	// 注:
	// <ul>
	// <li>1.印章类型SealType是INVOICE类型时，此参数才会生效</li>
	// <li>2.印章类型SealType是INVOICE类型，且该字段没有传入值或传入空时，会取该企业对应的统一社会信用代码作为默认的企业税号（<font color="red">如果是通过授权书授权方式认证的企业，此参数必传不能为空</font>）</li>
	// </ul>
	TaxIdentifyCode *string `json:"TaxIdentifyCode,omitnil,omitempty" name:"TaxIdentifyCode"`
}

func (r *CreateSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SealName")
	delete(f, "Agent")
	delete(f, "GenerateSource")
	delete(f, "SealType")
	delete(f, "FileName")
	delete(f, "Image")
	delete(f, "Width")
	delete(f, "Height")
	delete(f, "Color")
	delete(f, "SealHorizontalText")
	delete(f, "SealChordText")
	delete(f, "SealCentralType")
	delete(f, "FileToken")
	delete(f, "SealStyle")
	delete(f, "SealSize")
	delete(f, "TaxIdentifyCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealResponseParams struct {
	// 电子印章ID，为32位字符串。
	// 建议开发者保留此印章ID，后续指定签署区印章或者操作印章需此印章ID。
	// 可登录腾讯电子签控制台，在 "印章"->"印章中心"选择查看的印章，在"印章详情" 中查看某个印章的SealId(在页面中展示为印章ID)。
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSealResponse struct {
	*tchttp.BaseResponse
	Response *CreateSealResponseParams `json:"Response"`
}

func (r *CreateSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStaffResult struct {
	// 创建员工的成功列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessEmployeeData []*SuccessCreateStaffData `json:"SuccessEmployeeData,omitnil,omitempty" name:"SuccessEmployeeData"`

	// 创建员工的失败列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedEmployeeData []*FailedCreateStaffData `json:"FailedEmployeeData,omitnil,omitempty" name:"FailedEmployeeData"`
}

// Predefined struct for user
type CreateUserAutoSignEnableUrlRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

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

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 调用方自定义的个性化字段(可自定义此字段的值)，并以base64方式编码，支持的最大数据大小为 20480长度。 在个人自动签的开通、关闭等回调信息场景中，该字段的信息将原封不动地透传给贵方。 
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

type CreateUserAutoSignEnableUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

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

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 调用方自定义的个性化字段(可自定义此字段的值)，并以base64方式编码，支持的最大数据大小为 20480长度。 在个人自动签的开通、关闭等回调信息场景中，该字段的信息将原封不动地透传给贵方。 
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

func (r *CreateUserAutoSignEnableUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserAutoSignEnableUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SceneKey")
	delete(f, "AutoSignConfig")
	delete(f, "UrlType")
	delete(f, "NotifyType")
	delete(f, "NotifyAddress")
	delete(f, "ExpiredTime")
	delete(f, "Agent")
	delete(f, "UserData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserAutoSignEnableUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserAutoSignEnableUrlResponseParams struct {
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

type CreateUserAutoSignEnableUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserAutoSignEnableUrlResponseParams `json:"Response"`
}

func (r *CreateUserAutoSignEnableUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserAutoSignEnableUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserAutoSignSealUrlRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 自动签开通个人用户信息, 包括名字,身份证等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为当前时间往后30天。`如果不传，默认过期时间为当前时间往后7天。`
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`
}

type CreateUserAutoSignSealUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 自动签开通个人用户信息, 包括名字,身份证等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为当前时间往后30天。`如果不传，默认过期时间为当前时间往后7天。`
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`
}

func (r *CreateUserAutoSignSealUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserAutoSignSealUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	delete(f, "Agent")
	delete(f, "ExpiredTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserAutoSignSealUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserAutoSignSealUrlResponseParams struct {
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

type CreateUserAutoSignSealUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserAutoSignSealUrlResponseParams `json:"Response"`
}

func (r *CreateUserAutoSignSealUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserAutoSignSealUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserMobileChangeUrlRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。 支持填入集团子公司经办人 userId 代发合同。  注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 如果您要修改企业员工用户ID，传递此用户ID即可，其他参数（Name，UserAccountType，IdCardType，IdCardNumber）将被忽略。如果不传此用户ID，则会使用其他参数来进行链接生成。
	// 
	// [点击查看用户ID的获取方式](https://res.ess.tencent.cn/cdn/tsign-developer-center/assets/images/%E7%BB%84%E7%BB%87%E6%9E%B6%E6%9E%84-47eb7105dd300e6dc0c502fba22688ae.png)
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 要修改手机号用户的类型。
	// <ul><li>0：员工 （默认）</li><li>1：个人</li>
	// </ul>
	// 如果是员工类型，<b>只能修改本方员工，而不能修改其他企业的员工</b>。
	// 如果是个人类型，可<b>不指定用户身份，生成的是固定的链接，当前登录电子签小程序的用户可进行换绑。</b>
	UserAccountType *uint64 `json:"UserAccountType,omitnil,omitempty" name:"UserAccountType"`

	// 要修改手机号用户的姓名，请确保填写的姓名为对方的真实姓名，而非昵称等代名。
	// 
	// 如果没有传递 userId且 userAccountType 是 0 或者没有传递， 此参数为<b>必填项。</b>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 要修改手机号用户的证件类型，
	// 目前支持的账号类型如下：
	// 
	// <ul><li><b>ID_CARD </b>: （默认）中国大陆居民身份证 </li>
	// <li><b>HONGKONG_AND_MACAO</b> : 港澳居民来往内地通行证</li>
	// <li><b>HONGKONG_MACAO_AND_TAIWAN </b>: 港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 要修改手机号用户的身份证号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	// 如果没有传递 userId且 userAccountType 是 0 或者没有传递， 此参数为<b>必填项。</b>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 要跳转的链接类型
	// 
	// <ul>
	// <li><b>HTTP</b>：（默认）跳转电子签小程序的http_url,短信通知或者H5跳转适合此类型 ，此时返回长链 (默认类型)</li>
	// <li><b>HTTP_SHORT_URL</b>：跳转电子签小程序的http_url,短信通知或者H5跳转适合此类型，此时返回短链</li>
	// <li><b>APP</b>：第三方APP或小程序跳转电子签小程序的path, APP或者小程序跳转适合此类型</li>
	// </ul>
	// 
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 在用户完成实名认证后，其自定义数据将通过[手机号换绑回调](https://qian.tencent.com/developers/company/callback_types_staffs/#%E5%8D%81%E4%B8%89-%E4%B8%AA%E4%BA%BA%E5%91%98%E5%B7%A5%E6%89%8B%E6%9C%BA%E5%8F%B7%E4%BF%AE%E6%94%B9%E5%90%8E%E5%9B%9E%E8%B0%83)返回，以便用户确认其个人数据信息。请注意，自定义数据的字符长度上限为1000，且必须采用base64编码格式。
	// 
	// 请注意：
	// 此参数仅支持通过[获取c端用户实名链接](https://qian.tencent.com/developers/companyApis/users/CreateUserVerifyUrl)接口实名的用户生效。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

type CreateUserMobileChangeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。 支持填入集团子公司经办人 userId 代发合同。  注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 如果您要修改企业员工用户ID，传递此用户ID即可，其他参数（Name，UserAccountType，IdCardType，IdCardNumber）将被忽略。如果不传此用户ID，则会使用其他参数来进行链接生成。
	// 
	// [点击查看用户ID的获取方式](https://res.ess.tencent.cn/cdn/tsign-developer-center/assets/images/%E7%BB%84%E7%BB%87%E6%9E%B6%E6%9E%84-47eb7105dd300e6dc0c502fba22688ae.png)
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 要修改手机号用户的类型。
	// <ul><li>0：员工 （默认）</li><li>1：个人</li>
	// </ul>
	// 如果是员工类型，<b>只能修改本方员工，而不能修改其他企业的员工</b>。
	// 如果是个人类型，可<b>不指定用户身份，生成的是固定的链接，当前登录电子签小程序的用户可进行换绑。</b>
	UserAccountType *uint64 `json:"UserAccountType,omitnil,omitempty" name:"UserAccountType"`

	// 要修改手机号用户的姓名，请确保填写的姓名为对方的真实姓名，而非昵称等代名。
	// 
	// 如果没有传递 userId且 userAccountType 是 0 或者没有传递， 此参数为<b>必填项。</b>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 要修改手机号用户的证件类型，
	// 目前支持的账号类型如下：
	// 
	// <ul><li><b>ID_CARD </b>: （默认）中国大陆居民身份证 </li>
	// <li><b>HONGKONG_AND_MACAO</b> : 港澳居民来往内地通行证</li>
	// <li><b>HONGKONG_MACAO_AND_TAIWAN </b>: 港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 要修改手机号用户的身份证号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	// 如果没有传递 userId且 userAccountType 是 0 或者没有传递， 此参数为<b>必填项。</b>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 要跳转的链接类型
	// 
	// <ul>
	// <li><b>HTTP</b>：（默认）跳转电子签小程序的http_url,短信通知或者H5跳转适合此类型 ，此时返回长链 (默认类型)</li>
	// <li><b>HTTP_SHORT_URL</b>：跳转电子签小程序的http_url,短信通知或者H5跳转适合此类型，此时返回短链</li>
	// <li><b>APP</b>：第三方APP或小程序跳转电子签小程序的path, APP或者小程序跳转适合此类型</li>
	// </ul>
	// 
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 在用户完成实名认证后，其自定义数据将通过[手机号换绑回调](https://qian.tencent.com/developers/company/callback_types_staffs/#%E5%8D%81%E4%B8%89-%E4%B8%AA%E4%BA%BA%E5%91%98%E5%B7%A5%E6%89%8B%E6%9C%BA%E5%8F%B7%E4%BF%AE%E6%94%B9%E5%90%8E%E5%9B%9E%E8%B0%83)返回，以便用户确认其个人数据信息。请注意，自定义数据的字符长度上限为1000，且必须采用base64编码格式。
	// 
	// 请注意：
	// 此参数仅支持通过[获取c端用户实名链接](https://qian.tencent.com/developers/companyApis/users/CreateUserVerifyUrl)接口实名的用户生效。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

func (r *CreateUserMobileChangeUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserMobileChangeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Agent")
	delete(f, "UserId")
	delete(f, "UserAccountType")
	delete(f, "Name")
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "Endpoint")
	delete(f, "UserData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserMobileChangeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserMobileChangeUrlResponseParams struct {
	// 腾讯电子签小程序的实名认证链接。 如果没有传递，默认值是 HTTP。 链接的有效期均是 7 天。
	// 
	// <b>1.如果EndPoint是APP</b>，
	// 得到的链接类似于<a href="">pages/guide/index?to=MOBILE_CHANGE_INTENTION&shortKey=yDCZHUyOcExAlcOvNod0</a>, 用法可以参考描述中的"跳转到小程序的实现"
	// 
	// <b>2.如果EndPoint是HTTP</b>，
	// 得到的链接类似于<a href="">https://res.ess.tencent.cn/cdn/h5-activity/jump-mp.html?to=MOBILE_CHANGE_INTENTION&shortKey=yDCZHUyOcChrfpaswT0d</a>，点击后会跳转到腾讯电子签小程序进行签署
	// 
	// <b>3.如果EndPoint是HTTP_SHORT_URL</b>，
	// 得到的链接类似于<a href="">https://essurl.cn/2n**42Nd</a>，点击后会跳转到腾讯电子签小程序进行签署
	// 
	// 注： <font color="red">生成的链路后面不能再增加参数</font>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 链接失效期限，为Unix时间戳（单位秒），有如下规则：
	// 
	// <ul>
	// <li>如果指定更换绑定手机号的用户(指定用户ID或姓名等信息)，则设定的链接失效期限为7天后。</li>
	// <li>如果没有指定更换绑定手机号的用户，则生成通用跳转到个人换手机号的界面，链接不会过期。</li>
	// </ul>
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserMobileChangeUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserMobileChangeUrlResponseParams `json:"Response"`
}

func (r *CreateUserMobileChangeUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserMobileChangeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserVerifyUrlRequestParams struct {
	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 要实名的姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 要实名的身份证号码，
	// 身份证号码如果有x的话，统一传大写X
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 证件类型，目前只支持身份证类型：ID_CARD
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 要实名的手机号，兼容带+86的格式
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 实名完之后的跳转链接，最大长度1000个字符。
	// 链接类型请参考 <a href="https://qian.tencent.com/developers/company/openqianh5" target="_blank">跳转电子签H5</a>。
	// 
	// 注：此参数仅支持 Endpoint 为 <font color="red">H5 或 H5_SHORT_URL </font>的时候传递
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 要跳转的链接类型
	// 
	// <ul>
	// <li><strong>HTTP</strong>：适用于短信通知或H5跳转的电子签小程序HTTP长链接</li>
	// <li><strong>HTTP_SHORT_URL</strong>：适用于短信通知或H5跳转的电子签小程序HTTP短链接</li>
	// <li><strong>APP</strong>：（默认类型）适用于第三方APP或小程序跳转的电子签小程序路径</li>
	// <li><strong>H5</strong>：适用于跳转至电子签H5实名页面的长链接</li>
	// <li><strong>H5_SHORT_URL</strong>：适用于跳转至电子签H5实名页面的短链接</li>
	// </ul>
	// 
	// 注：如果不传递，默认值是 <font color="red"> APP </font>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 签署完成后是否自动回跳
	// <ul><li>false：否, 实名完成不会自动跳转回来(默认)</li><li>true：是, 实名完成会自动跳转回来</li></ul>
	// 
	// 注: 
	// 1. 该参数<font color="red">只针对APP类型（第三方APP或小程序跳转电子签小程序）场景</font> 的实名链接有效
	// 2. <font color="red">手机应用APP 或 微信小程序需要监控界面的返回走后序逻辑</font>, 微信小程序的文档可以参考[这个](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onShow-Object-object)
	// 3. <font color="red">电子签小程序跳转贵方APP，不支持自动跳转，必需用户手动点击完成按钮（微信的限制）</font> 
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil,omitempty" name:"AutoJumpBack"`

	// 在用户完成实名认证后，其自定义数据将通过[企业引导个人实名认证后回调](https://qian.tencent.com/developers/company/callback_types_staffs/#%E5%8D%81%E4%BA%8C-%E4%BC%81%E4%B8%9A%E5%BC%95%E5%AF%BC%E4%B8%AA%E4%BA%BA%E5%AE%9E%E5%90%8D%E8%AE%A4%E8%AF%81%E5%90%8E%E5%9B%9E%E8%B0%83)返回，以便用户确认其个人数据信息。请注意，自定义数据的字符长度上限为1000，且必须采用base64编码格式。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

type CreateUserVerifyUrlRequest struct {
	*tchttp.BaseRequest
	
	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 要实名的姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 要实名的身份证号码，
	// 身份证号码如果有x的话，统一传大写X
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 证件类型，目前只支持身份证类型：ID_CARD
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`

	// 要实名的手机号，兼容带+86的格式
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 实名完之后的跳转链接，最大长度1000个字符。
	// 链接类型请参考 <a href="https://qian.tencent.com/developers/company/openqianh5" target="_blank">跳转电子签H5</a>。
	// 
	// 注：此参数仅支持 Endpoint 为 <font color="red">H5 或 H5_SHORT_URL </font>的时候传递
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 要跳转的链接类型
	// 
	// <ul>
	// <li><strong>HTTP</strong>：适用于短信通知或H5跳转的电子签小程序HTTP长链接</li>
	// <li><strong>HTTP_SHORT_URL</strong>：适用于短信通知或H5跳转的电子签小程序HTTP短链接</li>
	// <li><strong>APP</strong>：（默认类型）适用于第三方APP或小程序跳转的电子签小程序路径</li>
	// <li><strong>H5</strong>：适用于跳转至电子签H5实名页面的长链接</li>
	// <li><strong>H5_SHORT_URL</strong>：适用于跳转至电子签H5实名页面的短链接</li>
	// </ul>
	// 
	// 注：如果不传递，默认值是 <font color="red"> APP </font>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 签署完成后是否自动回跳
	// <ul><li>false：否, 实名完成不会自动跳转回来(默认)</li><li>true：是, 实名完成会自动跳转回来</li></ul>
	// 
	// 注: 
	// 1. 该参数<font color="red">只针对APP类型（第三方APP或小程序跳转电子签小程序）场景</font> 的实名链接有效
	// 2. <font color="red">手机应用APP 或 微信小程序需要监控界面的返回走后序逻辑</font>, 微信小程序的文档可以参考[这个](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onShow-Object-object)
	// 3. <font color="red">电子签小程序跳转贵方APP，不支持自动跳转，必需用户手动点击完成按钮（微信的限制）</font> 
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil,omitempty" name:"AutoJumpBack"`

	// 在用户完成实名认证后，其自定义数据将通过[企业引导个人实名认证后回调](https://qian.tencent.com/developers/company/callback_types_staffs/#%E5%8D%81%E4%BA%8C-%E4%BC%81%E4%B8%9A%E5%BC%95%E5%AF%BC%E4%B8%AA%E4%BA%BA%E5%AE%9E%E5%90%8D%E8%AE%A4%E8%AF%81%E5%90%8E%E5%9B%9E%E8%B0%83)返回，以便用户确认其个人数据信息。请注意，自定义数据的字符长度上限为1000，且必须采用base64编码格式。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`
}

func (r *CreateUserVerifyUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserVerifyUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Name")
	delete(f, "IdCardNumber")
	delete(f, "IdCardType")
	delete(f, "Mobile")
	delete(f, "JumpUrl")
	delete(f, "Endpoint")
	delete(f, "AutoJumpBack")
	delete(f, "UserData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserVerifyUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserVerifyUrlResponseParams struct {
	// 腾讯电子签小程序的实名认证链接。
	// 如果没有传递，默认值是 HTTP。 链接的有效期均是 7 天。
	// 
	// <strong>1.如果EndPoint是APP</strong>：
	// 得到的链接类似于<a href="">pages/guide/index?to=MP_PERSONAL_VERIFY&shortKey=yDCZHUyOcExAlcOvNod0</a>, 用法可以参考描述中的"跳转到小程序的实现"
	// 
	// <strong>2.如果EndPoint是HTTP</strong>：
	// 得到的链接类似于 <a href="">https://res.ess.tencent.cn/cdn/h5-activity/jump-mp.html?to=TAG_VERIFY&shortKey=yDCZHUyOcChrfpaswT0d</a>，点击后会跳转到腾讯电子签小程序进行签署
	// 
	// <strong>3.如果EndPoint是HTTP_SHORT_URL</strong>：
	// 得到的链接类似于<a href="">https://essurl.cn/2n**42Nd</a>，点击后会跳转到腾讯电子签小程序进行签署
	// 
	// <strong>4.如果EndPoint是H5</strong>：
	// 得到的链接类似于 <a href="">https://quick.test.qian.tencent.cn/guide?Code=yDU****VJhsS5q&CodeType=xxx&shortKey=yD*****frcb</a>，点击后会跳转到腾讯电子签H5页面进行签署
	// 
	// <strong>5.如果EndPoint是H5_SHORT_URL</strong>：
	// 得到的链接类似于<a href="">https://essurl.cn/2n**42Nd</a>，点击后会跳转到腾讯电子签H5页面进行签署
	// 
	// 
	// `注：` <font color="red">生成的链路后面不能再增加参数，防止出错重复参数覆盖原有的参数</font>
	// 示例值：https://essurl.cn/2n**42Nd
	UserVerifyUrl *string `json:"UserVerifyUrl,omitnil,omitempty" name:"UserVerifyUrl"`

	// 链接过期时间，为Unix时间戳（单位为秒）。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 小程序appid，用于半屏拉起电子签小程序， 仅在 Endpoint 设置为 APP 的时候返回
	MiniAppId *string `json:"MiniAppId,omitnil,omitempty" name:"MiniAppId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserVerifyUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserVerifyUrlResponseParams `json:"Response"`
}

func (r *CreateUserVerifyUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserVerifyUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebThemeConfigRequestParams struct {
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 主题类型，取值如下：
	// <ul><li> **EMBED_WEB_THEME**：嵌入式主题（默认），web页面嵌入的主题风格配置</li>
	// </ul>
	ThemeType *string `json:"ThemeType,omitnil,omitempty" name:"ThemeType"`

	// 电子签logo是否展示，主体颜色等配置项
	WebThemeConfig *WebThemeConfig `json:"WebThemeConfig,omitnil,omitempty" name:"WebThemeConfig"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CreateWebThemeConfigRequest struct {
	*tchttp.BaseRequest
	
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 主题类型，取值如下：
	// <ul><li> **EMBED_WEB_THEME**：嵌入式主题（默认），web页面嵌入的主题风格配置</li>
	// </ul>
	ThemeType *string `json:"ThemeType,omitnil,omitempty" name:"ThemeType"`

	// 电子签logo是否展示，主体颜色等配置项
	WebThemeConfig *WebThemeConfig `json:"WebThemeConfig,omitnil,omitempty" name:"WebThemeConfig"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *CreateWebThemeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebThemeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "ThemeType")
	delete(f, "WebThemeConfig")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWebThemeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebThemeConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWebThemeConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateWebThemeConfigResponseParams `json:"Response"`
}

func (r *CreateWebThemeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebThemeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExtendedServiceAuthInfosRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 本企业员工的id，需要已实名，正常在职员工
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 取值如下所示：
	// <ul><li>OPEN_SERVER_SIGN：企业自动签</li>
	// <li>BATCH_SIGN：批量签署</li>
	// </ul>
	ExtendServiceType *string `json:"ExtendServiceType,omitnil,omitempty" name:"ExtendServiceType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DeleteExtendedServiceAuthInfosRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 本企业员工的id，需要已实名，正常在职员工
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 取值如下所示：
	// <ul><li>OPEN_SERVER_SIGN：企业自动签</li>
	// <li>BATCH_SIGN：批量签署</li>
	// </ul>
	ExtendServiceType *string `json:"ExtendServiceType,omitnil,omitempty" name:"ExtendServiceType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *DeleteExtendedServiceAuthInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExtendedServiceAuthInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "UserIds")
	delete(f, "ExtendServiceType")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteExtendedServiceAuthInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExtendedServiceAuthInfosResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteExtendedServiceAuthInfosResponse struct {
	*tchttp.BaseResponse
	Response *DeleteExtendedServiceAuthInfosResponseParams `json:"Response"`
}

func (r *DeleteExtendedServiceAuthInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExtendedServiceAuthInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationDepartmentRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 电子签中的部门ID，通过<a href="https://qian.tencent.com/developers/companyApis/organizations/DescribeIntegrationDepartments" target="_blank">DescribeIntegrationDepartments</a>接口可获得。
	DeptId *string `json:"DeptId,omitnil,omitempty" name:"DeptId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 交接部门ID。
	// 待删除部门中的合同、印章和模板数据，将会被交接至该部门ID下；若未填写则交接至公司根部门。
	ReceiveDeptId *string `json:"ReceiveDeptId,omitnil,omitempty" name:"ReceiveDeptId"`
}

type DeleteIntegrationDepartmentRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 电子签中的部门ID，通过<a href="https://qian.tencent.com/developers/companyApis/organizations/DescribeIntegrationDepartments" target="_blank">DescribeIntegrationDepartments</a>接口可获得。
	DeptId *string `json:"DeptId,omitnil,omitempty" name:"DeptId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 交接部门ID。
	// 待删除部门中的合同、印章和模板数据，将会被交接至该部门ID下；若未填写则交接至公司根部门。
	ReceiveDeptId *string `json:"ReceiveDeptId,omitnil,omitempty" name:"ReceiveDeptId"`
}

func (r *DeleteIntegrationDepartmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationDepartmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "DeptId")
	delete(f, "Agent")
	delete(f, "ReceiveDeptId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIntegrationDepartmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationDepartmentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIntegrationDepartmentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIntegrationDepartmentResponseParams `json:"Response"`
}

func (r *DeleteIntegrationDepartmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationEmployeesRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写UserId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 待离职员工的信息最多不超过100个。应符合以下规则：
	// 
	// 1. UserId和OpenId不可同时为空，必须填写其中一个，优先使用UserId。
	// 
	// 2. **若需要进行离职交接**，交接人信息ReceiveUserId和ReceiveOpenId不可同时为空，必须填写其中一个，优先使用ReceiveUserId。
	Employees []*Staff `json:"Employees,omitnil,omitempty" name:"Employees"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DeleteIntegrationEmployeesRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写UserId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 待离职员工的信息最多不超过100个。应符合以下规则：
	// 
	// 1. UserId和OpenId不可同时为空，必须填写其中一个，优先使用UserId。
	// 
	// 2. **若需要进行离职交接**，交接人信息ReceiveUserId和ReceiveOpenId不可同时为空，必须填写其中一个，优先使用ReceiveUserId。
	Employees []*Staff `json:"Employees,omitnil,omitempty" name:"Employees"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *DeleteIntegrationEmployeesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationEmployeesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Employees")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIntegrationEmployeesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationEmployeesResponseParams struct {
	// 员工删除结果。包含成功数据与失败数据。
	// <ul><li>**成功数据**：展示员工姓名、手机号与电子签平台UserId</li>
	// <li>**失败数据**：展示员工电子签平台UserId、第三方平台OpenId和失败原因</li></ul>
	DeleteEmployeeResult *DeleteStaffsResult `json:"DeleteEmployeeResult,omitnil,omitempty" name:"DeleteEmployeeResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIntegrationEmployeesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIntegrationEmployeesResponseParams `json:"Response"`
}

func (r *DeleteIntegrationEmployeesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationEmployeesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationRoleUsersRequestParams struct {
	// 执行本接口操作的员工信息。 注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 角色id，可以通过DescribeIntegrationRoles接口获取角色信息
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 用户信息,最多 200 个用户，并且 UserId 和 OpenId 二选一，其他字段不需要传
	Users []*UserInfo `json:"Users,omitnil,omitempty" name:"Users"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DeleteIntegrationRoleUsersRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。 注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 角色id，可以通过DescribeIntegrationRoles接口获取角色信息
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 用户信息,最多 200 个用户，并且 UserId 和 OpenId 二选一，其他字段不需要传
	Users []*UserInfo `json:"Users,omitnil,omitempty" name:"Users"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *DeleteIntegrationRoleUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationRoleUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "RoleId")
	delete(f, "Users")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIntegrationRoleUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationRoleUsersResponseParams struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIntegrationRoleUsersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIntegrationRoleUsersResponseParams `json:"Response"`
}

func (r *DeleteIntegrationRoleUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationRoleUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrganizationAuthorizationInfo struct {
	// 认证流 Id 是指在企业认证过程中，当前操作人的认证流程的唯一标识。每个企业在认证过程中只能有一条认证流认证成功。这意味着在同一认证过程内，一个企业只能有一个认证流程处于成功状态，以确保认证的唯一性和有效性。
	AuthorizationId *string `json:"AuthorizationId,omitnil,omitempty" name:"AuthorizationId"`

	// 认证的企业名称
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 清除认证流产生的错误信息
	Errormessage *string `json:"Errormessage,omitnil,omitempty" name:"Errormessage"`
}

// Predefined struct for user
type DeleteOrganizationAuthorizationsRequestParams struct {
	// 执行本接口操作的员工信息, userId 必填。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 认证流Ids数组
	// 认证流 Id 是指在企业认证过程中，当前操作人的认证流程的唯一标识。每个企业在认证过程中只能有一条认证流认证成功。这意味着在同一认证过程内，一个企业只能有一个认证流程处于成功状态，以确保认证的唯一性和有效性。
	// 
	// 
	// 认证流 Id可以通过回调 [授权书认证审核结果回调](https://qian.tencent.com/developers/company/callback_types_staffs/#%E5%8D%81%E5%9B%9B-%E6%8E%88%E6%9D%83%E4%B9%A6%E8%AE%A4%E8%AF%81%E5%AE%A1%E6%A0%B8%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83) 获取
	// 
	// 注意：
	// 如果传递了认证流Id，则下面的参数 超管二要素不会生效
	AuthorizationIds []*string `json:"AuthorizationIds,omitnil,omitempty" name:"AuthorizationIds"`

	// 认证人姓名，组织机构超管姓名。 
	// 在注册流程中，必须是超管本人进行操作。 
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 认证人手机号，组织机构超管手机号。 在注册流程中，必须是超管本人进行操作。 
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`

	// 代理企业和员工的信息。在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DeleteOrganizationAuthorizationsRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息, userId 必填。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 认证流Ids数组
	// 认证流 Id 是指在企业认证过程中，当前操作人的认证流程的唯一标识。每个企业在认证过程中只能有一条认证流认证成功。这意味着在同一认证过程内，一个企业只能有一个认证流程处于成功状态，以确保认证的唯一性和有效性。
	// 
	// 
	// 认证流 Id可以通过回调 [授权书认证审核结果回调](https://qian.tencent.com/developers/company/callback_types_staffs/#%E5%8D%81%E5%9B%9B-%E6%8E%88%E6%9D%83%E4%B9%A6%E8%AE%A4%E8%AF%81%E5%AE%A1%E6%A0%B8%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83) 获取
	// 
	// 注意：
	// 如果传递了认证流Id，则下面的参数 超管二要素不会生效
	AuthorizationIds []*string `json:"AuthorizationIds,omitnil,omitempty" name:"AuthorizationIds"`

	// 认证人姓名，组织机构超管姓名。 
	// 在注册流程中，必须是超管本人进行操作。 
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 认证人手机号，组织机构超管手机号。 在注册流程中，必须是超管本人进行操作。 
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`

	// 代理企业和员工的信息。在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
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
	delete(f, "Operator")
	delete(f, "AuthorizationIds")
	delete(f, "AdminName")
	delete(f, "AdminMobile")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationAuthorizationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationAuthorizationsResponseParams struct {
	// 清理的认证流的详细信息，其中包括企业名称，认证流唯一 Id 以及清理过程中产生的错误信息
	DeleteOrganizationAuthorizationInfos []*DeleteOrganizationAuthorizationInfo `json:"DeleteOrganizationAuthorizationInfos,omitnil,omitempty" name:"DeleteOrganizationAuthorizationInfos"`

	// 批量清理认证流返回的状态值
	// 其中包括
	// - 1 全部成功
	// - 2 部分成功
	// - 3 全部失败
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

// Predefined struct for user
type DeleteSealPoliciesRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 印章授权编码数组。这个参数跟下面的SealId其中一个必填，另外一个可选填
	PolicyIds []*string `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// 电子印章ID，为32位字符串。
	// 建议开发者保留此印章ID，后续指定签署区印章或者操作印章需此印章ID。
	// 可登录腾讯电子签控制台，在 "印章"->"印章中心"选择查看的印章，在"印章详情" 中查看某个印章的SealId(在页面中展示为印章ID)。
	// 注：印章ID。这个参数跟上面的PolicyIds其中一个必填，另外一个可选填。
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 待授权的员工ID，员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查看某位员工的UserId(在页面中展示为用户ID)。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DeleteSealPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 印章授权编码数组。这个参数跟下面的SealId其中一个必填，另外一个可选填
	PolicyIds []*string `json:"PolicyIds,omitnil,omitempty" name:"PolicyIds"`

	// 电子印章ID，为32位字符串。
	// 建议开发者保留此印章ID，后续指定签署区印章或者操作印章需此印章ID。
	// 可登录腾讯电子签控制台，在 "印章"->"印章中心"选择查看的印章，在"印章详情" 中查看某个印章的SealId(在页面中展示为印章ID)。
	// 注：印章ID。这个参数跟上面的PolicyIds其中一个必填，另外一个可选填。
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 待授权的员工ID，员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查看某位员工的UserId(在页面中展示为用户ID)。
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *DeleteSealPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSealPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "PolicyIds")
	delete(f, "SealId")
	delete(f, "UserIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSealPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSealPoliciesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSealPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSealPoliciesResponseParams `json:"Response"`
}

func (r *DeleteSealPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSealPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStaffsResult struct {
	// 删除员工的成功数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessEmployeeData []*SuccessDeleteStaffData `json:"SuccessEmployeeData,omitnil,omitempty" name:"SuccessEmployeeData"`

	// 删除员工的失败数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedEmployeeData []*FailedDeleteStaffData `json:"FailedEmployeeData,omitnil,omitempty" name:"FailedEmployeeData"`
}

type Department struct {
	// 部门ID。
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 部门名称。
	DepartmentName *string `json:"DepartmentName,omitnil,omitempty" name:"DepartmentName"`
}

// Predefined struct for user
type DescribeBatchOrganizationRegistrationUrlsRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 通过接口CreateBatchOrganizationRegistrationTasks创建企业批量认证链接任得到的任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DescribeBatchOrganizationRegistrationUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 通过接口CreateBatchOrganizationRegistrationTasks创建企业批量认证链接任得到的任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
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
	delete(f, "Operator")
	delete(f, "TaskId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchOrganizationRegistrationUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOrganizationRegistrationUrlsResponseParams struct {
	// 企业批量注册链接信息
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
type DescribeBillUsageDetailRequestParams struct {
	// 查询开始时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定分页每页返回的数据条数，如果不传默认为 50，单页最大支持 50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

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
	// <li>**NoAuthSign**: 形式签</li>
	// </ul>
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DescribeBillUsageDetailRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定分页每页返回的数据条数，如果不传默认为 50，单页最大支持 50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

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
	// <li>**NoAuthSign**: 形式签</li>
	// </ul>
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *DescribeBillUsageDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillUsageDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "QuotaType")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillUsageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillUsageDetailResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 消耗详情
	Details []*BillUsageDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillUsageDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillUsageDetailResponseParams `json:"Response"`
}

func (r *DescribeBillUsageDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillUsageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillUsageRequestParams struct {
	// 查询开始时间字符串，格式为yyyymmdd,时间跨度不能大于90天
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间字符串，格式为yyyymmdd,时间跨度不能大于90天
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的套餐类型 （选填 ）不传则查询所有套餐；目前支持:<ul><li>**CloudEnterprise**: 企业版合同</li><li>**SingleSignature**: 单方签章</li><li>**CloudProve**: 签署报告</li><li>**CloudOnlineSign**: 腾讯会议在线签约</li><li>**ChannelWeCard**: 微工卡</li><li>**SignFlow**: 合同套餐</li><li>**SignFace**: 签署意愿（人脸识别）</li><li>**SignPassword**: 签署意愿（密码）</li><li>**SignSMS**: 签署意愿（短信）</li><li>**PersonalEssAuth**: 签署人实名（腾讯电子签认证）</li><li>**PersonalThirdAuth**: 签署人实名（信任第三方认证）</li><li>**OrgEssAuth**: 签署企业实名</li><li>**FlowNotify**: 短信通知</li><li>**AuthService**: 企业工商信息查询</li></ul>
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// 是否展示集团子企业的明细，默认否
	DisplaySubEnterprise *bool `json:"DisplaySubEnterprise,omitnil,omitempty" name:"DisplaySubEnterprise"`
}

type DescribeBillUsageRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间字符串，格式为yyyymmdd,时间跨度不能大于90天
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间字符串，格式为yyyymmdd,时间跨度不能大于90天
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的套餐类型 （选填 ）不传则查询所有套餐；目前支持:<ul><li>**CloudEnterprise**: 企业版合同</li><li>**SingleSignature**: 单方签章</li><li>**CloudProve**: 签署报告</li><li>**CloudOnlineSign**: 腾讯会议在线签约</li><li>**ChannelWeCard**: 微工卡</li><li>**SignFlow**: 合同套餐</li><li>**SignFace**: 签署意愿（人脸识别）</li><li>**SignPassword**: 签署意愿（密码）</li><li>**SignSMS**: 签署意愿（短信）</li><li>**PersonalEssAuth**: 签署人实名（腾讯电子签认证）</li><li>**PersonalThirdAuth**: 签署人实名（信任第三方认证）</li><li>**OrgEssAuth**: 签署企业实名</li><li>**FlowNotify**: 短信通知</li><li>**AuthService**: 企业工商信息查询</li></ul>
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// 是否展示集团子企业的明细，默认否
	DisplaySubEnterprise *bool `json:"DisplaySubEnterprise,omitnil,omitempty" name:"DisplaySubEnterprise"`
}

func (r *DescribeBillUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QuotaType")
	delete(f, "DisplaySubEnterprise")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillUsageResponseParams struct {
	// 企业套餐余额及使用情况
	Summary []*OrgBillSummary `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 集团子企业套餐使用情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubOrgSummary []*SubOrgBillSummary `json:"SubOrgSummary,omitnil,omitempty" name:"SubOrgSummary"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillUsageResponseParams `json:"Response"`
}

func (r *DescribeBillUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCancelFlowsTaskRequestParams struct {
	// 执行本接口操作的员工信息。
	// <br/>注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 批量撤销任务编号，为32位字符串，通过接口[获取批量撤销签署流程腾讯电子签小程序链接](https://qian.tencent.com/developers/companyApis/operateFlows/CreateBatchCancelFlowUrl)获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 代理企业和员工的信息。
	// <br/>在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DescribeCancelFlowsTaskRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// <br/>注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 批量撤销任务编号，为32位字符串，通过接口[获取批量撤销签署流程腾讯电子签小程序链接](https://qian.tencent.com/developers/companyApis/operateFlows/CreateBatchCancelFlowUrl)获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 代理企业和员工的信息。
	// <br/>在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
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
	delete(f, "Operator")
	delete(f, "TaskId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCancelFlowsTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCancelFlowsTaskResponseParams struct {
	// 批量撤销任务编号，为32位字符串，通过接口[获取批量撤销签署流程腾讯电子签小程序链接](https://qian.tencent.com/developers/companyApis/operateFlows/CreateBatchCancelFlowUrl)获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态，需要关注的状态
	// <ul><li>**PROCESSING**  - 任务执行中</li>
	// <li>**END** - 任务处理完成</li>
	// <li>**TIMEOUT** 任务超时未处理完成，用户未在批量撤销链接有效期内操作</li></ul>
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
type DescribeExtendedServiceAuthDetailRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 要查询的扩展服务类型。
	// 如下所示：
	// <ul><li>OPEN_SERVER_SIGN：企业静默签署</li>
	// <li>BATCH_SIGN：批量签署</li>
	// </ul>
	ExtendServiceType *string `json:"ExtendServiceType,omitnil,omitempty" name:"ExtendServiceType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定每页返回的数据条数，和Offset参数配合使用。 注：`1.默认值为20，单页做大值为200。`	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。 注：`1.offset从0开始，即第一页为0。` `2.默认从第一页返回。`	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeExtendedServiceAuthDetailRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 要查询的扩展服务类型。
	// 如下所示：
	// <ul><li>OPEN_SERVER_SIGN：企业静默签署</li>
	// <li>BATCH_SIGN：批量签署</li>
	// </ul>
	ExtendServiceType *string `json:"ExtendServiceType,omitnil,omitempty" name:"ExtendServiceType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

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
	delete(f, "Operator")
	delete(f, "ExtendServiceType")
	delete(f, "Agent")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtendedServiceAuthDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtendedServiceAuthDetailResponseParams struct {
	// 服务授权的信息列表，根据查询类型返回特定扩展服务的授权状况。
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
type DescribeExtendedServiceAuthInfosRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 要查询的扩展服务类型。
	// 默认为空，即查询当前支持的所有扩展服务信息。
	// 若需查询单个扩展服务的开通情况，请传递相应的值，如下所示：
	// <ul><li>OPEN_SERVER_SIGN：企业自动签署</li>
	// <li>AUTO_SIGN_CAN_FILL_IN：本企业自动签合同支持签前内容补充</li>
	// <li>BATCH_SIGN：批量签署</li>
	// <li>OVERSEA_SIGN：企业与港澳台居民签署合同</li>
	// <li>AGE_LIMIT_EXPANSION：拓宽签署方年龄限制</li>
	// <li>MOBILE_CHECK_APPROVER：个人签署方仅校验手机号</li>
	// <li>HIDE_OPERATOR_DISPLAY：隐藏合同经办人姓名</li>
	// <li>ORGANIZATION_OCR_FALLBACK：正楷临摹签名失败后更换其他签名类型</li>
	// <li>ORGANIZATION_FLOW_NOTIFY_TYPE：短信通知签署方</li>
	// <li>HIDE_ONE_KEY_SIGN：个人签署方手动签字</li>
	// <li>ORGANIZATION_FLOW_EMAIL_NOTIFY：邮件通知签署方</li>
	// <li>FLOW_APPROVAL：合同审批强制开启</li>
	// <li>ORGANIZATION_FLOW_PASSWD_NOTIFY：签署密码开通引导</li></ul>
	ExtendServiceType *string `json:"ExtendServiceType,omitnil,omitempty" name:"ExtendServiceType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DescribeExtendedServiceAuthInfosRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 要查询的扩展服务类型。
	// 默认为空，即查询当前支持的所有扩展服务信息。
	// 若需查询单个扩展服务的开通情况，请传递相应的值，如下所示：
	// <ul><li>OPEN_SERVER_SIGN：企业自动签署</li>
	// <li>AUTO_SIGN_CAN_FILL_IN：本企业自动签合同支持签前内容补充</li>
	// <li>BATCH_SIGN：批量签署</li>
	// <li>OVERSEA_SIGN：企业与港澳台居民签署合同</li>
	// <li>AGE_LIMIT_EXPANSION：拓宽签署方年龄限制</li>
	// <li>MOBILE_CHECK_APPROVER：个人签署方仅校验手机号</li>
	// <li>HIDE_OPERATOR_DISPLAY：隐藏合同经办人姓名</li>
	// <li>ORGANIZATION_OCR_FALLBACK：正楷临摹签名失败后更换其他签名类型</li>
	// <li>ORGANIZATION_FLOW_NOTIFY_TYPE：短信通知签署方</li>
	// <li>HIDE_ONE_KEY_SIGN：个人签署方手动签字</li>
	// <li>ORGANIZATION_FLOW_EMAIL_NOTIFY：邮件通知签署方</li>
	// <li>FLOW_APPROVAL：合同审批强制开启</li>
	// <li>ORGANIZATION_FLOW_PASSWD_NOTIFY：签署密码开通引导</li></ul>
	ExtendServiceType *string `json:"ExtendServiceType,omitnil,omitempty" name:"ExtendServiceType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *DescribeExtendedServiceAuthInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtendedServiceAuthInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "ExtendServiceType")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtendedServiceAuthInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtendedServiceAuthInfosResponseParams struct {
	// 服务开通和授权的信息列表，根据查询类型返回所有支持的扩展服务开通和授权状况，或者返回特定扩展服务的开通和授权状况。
	AuthInfoList []*ExtendAuthInfo `json:"AuthInfoList,omitnil,omitempty" name:"AuthInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExtendedServiceAuthInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtendedServiceAuthInfosResponseParams `json:"Response"`
}

func (r *DescribeExtendedServiceAuthInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtendedServiceAuthInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileUrlsRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 文件对应的业务类型，目前支持：
	// <ul>
	// <li>**FLOW ** : <font color="red">如需下载合同文件请选择此项</font></li>
	// <li>**TEMPLATE ** : 如需下载模板文件请选择此项</li>
	// <li>**DOCUMENT  **: 如需下载文档文件请选择此项</li>
	// <li>**SEAL  **: 如需下载印章图片请选择此项</li>
	// </ul>
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// 业务编号的数组，取值如下：
	// <ul>
	// <li>流程编号</li>
	// <li>模板编号</li>
	// <li>文档编号</li>
	// <li>印章编号</li>
	// <li>如需下载合同文件请传入FlowId，最大支持20个资源</li>
	// </ul>
	BusinessIds []*string `json:"BusinessIds,omitnil,omitempty" name:"BusinessIds"`

	// 下载后的文件命名，只有FileType为zip的时候生效
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 要下载的文件类型，取值如下：
	// <ul>
	// <li>JPG</li>
	// <li>PDF</li>
	// <li>ZIP</li>
	// </ul>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0，最大 1000。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定分页每页返回的数据条数，如果不传默认为 20，单页最大支持 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 下载url过期时间，单位秒。0: 按默认值5分钟，允许范围：1s~24x60x60s(1天)
	UrlTtl *int64 `json:"UrlTtl,omitnil,omitempty" name:"UrlTtl"`

	// 暂不开放
	//
	// Deprecated: CcToken is deprecated.
	CcToken *string `json:"CcToken,omitnil,omitempty" name:"CcToken"`

	// 暂不开放
	//
	// Deprecated: Scene is deprecated.
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DescribeFileUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 文件对应的业务类型，目前支持：
	// <ul>
	// <li>**FLOW ** : <font color="red">如需下载合同文件请选择此项</font></li>
	// <li>**TEMPLATE ** : 如需下载模板文件请选择此项</li>
	// <li>**DOCUMENT  **: 如需下载文档文件请选择此项</li>
	// <li>**SEAL  **: 如需下载印章图片请选择此项</li>
	// </ul>
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// 业务编号的数组，取值如下：
	// <ul>
	// <li>流程编号</li>
	// <li>模板编号</li>
	// <li>文档编号</li>
	// <li>印章编号</li>
	// <li>如需下载合同文件请传入FlowId，最大支持20个资源</li>
	// </ul>
	BusinessIds []*string `json:"BusinessIds,omitnil,omitempty" name:"BusinessIds"`

	// 下载后的文件命名，只有FileType为zip的时候生效
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 要下载的文件类型，取值如下：
	// <ul>
	// <li>JPG</li>
	// <li>PDF</li>
	// <li>ZIP</li>
	// </ul>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0，最大 1000。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定分页每页返回的数据条数，如果不传默认为 20，单页最大支持 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 下载url过期时间，单位秒。0: 按默认值5分钟，允许范围：1s~24x60x60s(1天)
	UrlTtl *int64 `json:"UrlTtl,omitnil,omitempty" name:"UrlTtl"`

	// 暂不开放
	CcToken *string `json:"CcToken,omitnil,omitempty" name:"CcToken"`

	// 暂不开放
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
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
	delete(f, "Operator")
	delete(f, "BusinessType")
	delete(f, "BusinessIds")
	delete(f, "FileName")
	delete(f, "FileType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "UrlTtl")
	delete(f, "CcToken")
	delete(f, "Scene")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileUrlsResponseParams struct {
	// 文件URL信息；
	// 链接不是永久链接,  过期时间受UrlTtl入参的影响,  默认有效期5分钟后,  到期后链接失效。
	FileUrls []*FileUrl `json:"FileUrls,omitnil,omitempty" name:"FileUrls"`

	// URL数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 查询的合同流程ID列表最多支持100个流程ID。 
	// 
	// 如果某个合同流程ID不存在，系统会跳过此ID的查询，继续查询剩余存在的合同流程。
	// 
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。[点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DescribeFlowBriefsRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 查询的合同流程ID列表最多支持100个流程ID。 
	// 
	// 如果某个合同流程ID不存在，系统会跳过此ID的查询，继续查询剩余存在的合同流程。
	// 
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。[点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
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
	// 合同流程基础信息列表，包含流程的名称、状态、创建日期等基本信息。 
	// 注：`与入参 FlowIds 的顺序可能存在不一致的情况。`
	FlowBriefs []*FlowBrief `json:"FlowBriefs,omitnil,omitempty" name:"FlowBriefs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeFlowComponentsRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 
	// [点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DescribeFlowComponentsRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 
	// [点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *DescribeFlowComponentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowComponentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowComponentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowComponentsResponseParams struct {
	// 合同流程关联的填写控件信息，包括填写控件的归属方以及是否填写等内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientComponentInfos []*RecipientComponentInfo `json:"RecipientComponentInfos,omitnil,omitempty" name:"RecipientComponentInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFlowComponentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowComponentsResponseParams `json:"Response"`
}

func (r *DescribeFlowComponentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowEvidenceReportRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 签署报告编号, 由<a href="https://qian.tencent.com/developers/companyApis/certificate/CreateFlowEvidenceReport" target="_blank">提交申请出证报告任务</a>产生
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定申请的报告类型，可选类型如下：
	// <ul><li> **0** :合同签署报告（默认）</li>
	// <li> **1** :公证处核验报告</li></ul>
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`
}

type DescribeFlowEvidenceReportRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 签署报告编号, 由<a href="https://qian.tencent.com/developers/companyApis/certificate/CreateFlowEvidenceReport" target="_blank">提交申请出证报告任务</a>产生
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 指定申请的报告类型，可选类型如下：
	// <ul><li> **0** :合同签署报告（默认）</li>
	// <li> **1** :公证处核验报告</li></ul>
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`
}

func (r *DescribeFlowEvidenceReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowEvidenceReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "ReportId")
	delete(f, "Agent")
	delete(f, "ReportType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowEvidenceReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowEvidenceReportResponseParams struct {
	// 出证报告PDF的下载 URL，`有效期为5分钟`，超过有效期后将无法再下载。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportUrl *string `json:"ReportUrl,omitnil,omitempty" name:"ReportUrl"`

	// 出证任务执行的状态, 状态含义如下：
	// 
	// <ul><li>**EvidenceStatusExecuting**：  出证任务在执行中</li>
	// <li>**EvidenceStatusSuccess**：  出证任务执行成功</li>
	// <li>**EvidenceStatusFailed** ： 出征任务执行失败</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFlowEvidenceReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowEvidenceReportResponseParams `json:"Response"`
}

func (r *DescribeFlowEvidenceReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowEvidenceReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowInfoRequestParams struct {
	// 执行本接口操作的员工信息。 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`	
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 需要查询的流程ID列表，最多可传入100个ID。
	// 如果要查询合同组的信息，则不需要传入此参数，只需传入 FlowGroupId 参数即可。
	// 
	// 
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	// 
	// [点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。	
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要查询的流程组ID，如果传入此参数，则会忽略 FlowIds 参数。该合同组由<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowGroupByFiles" target="_blank">通过多文件创建合同组签署流程</a>等接口创建。
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`
}

type DescribeFlowInfoRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`	
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 需要查询的流程ID列表，最多可传入100个ID。
	// 如果要查询合同组的信息，则不需要传入此参数，只需传入 FlowGroupId 参数即可。
	// 
	// 
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	// 
	// [点击查看FlowId在控制台中的位置](https://qcloudimg.tencent-cloud.cn/raw/0a83015166cfe1cb043d14f9ec4bd75e.png)
	FlowIds []*string `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。	
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需要查询的流程组ID，如果传入此参数，则会忽略 FlowIds 参数。该合同组由<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowGroupByFiles" target="_blank">通过多文件创建合同组签署流程</a>等接口创建。
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`
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
	delete(f, "Operator")
	delete(f, "FlowIds")
	delete(f, "Agent")
	delete(f, "FlowGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowInfoResponseParams struct {
	// 合同流程的详细信息。
	// 如果查询的是合同组信息，则返回的是组内所有子合同流程的详细信息。
	FlowDetailInfos []*FlowDetailInfo `json:"FlowDetailInfos,omitnil,omitempty" name:"FlowDetailInfos"`

	// 合同组ID，只有在查询合同组信息时才会返回。
	FlowGroupId *string `json:"FlowGroupId,omitnil,omitempty" name:"FlowGroupId"`

	// 合同组名称，只有在查询合同组信息时才会返回。
	FlowGroupName *string `json:"FlowGroupName,omitnil,omitempty" name:"FlowGroupName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 查询内容控制
	// 
	// <ul><li>**0**：模板列表及详情（默认）</li>
	// <li>**1**：仅模板列表</li></ul>
	ContentType *int64 `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 搜索过滤的条件，本字段允许您通过指定模板 ID 或模板名称来进行查询。
	// <ul><li><strong>模板 ID</strong>：<strong>Key</strong>设置为 <code>template-id</code> ，<strong>Values</strong>为您想要查询的 <a href="https://qcloudimg.tencent-cloud.cn/raw/5c27b917b2bbe8c341566c78ca6f8782.png" target="_blank">模板 ID </a>列表。</li>  <li><strong>主企业模板 ID</strong>：<strong>Key</strong>设置为 <code>share-template-id</code> ，<strong>Values</strong>为您想要查询的 <a href="https://qcloudimg.tencent-cloud.cn/raw/5c27b917b2bbe8c341566c78ca6f8782.png" target="_blank">主企业模板 ID </a>列表。用来查询主企业分享模板到子企业场景下，子企业的模板信息，在此情境下，参数 <strong>Agent.ProxyOrganizationId</strong>（子企业的组织ID）为必填项。</li> <li><strong>模板名称</strong>：<strong>Key</strong>设置为 <code>template-name</code> ，<strong>Values</strong>为您想要查询的<a href="https://qcloudimg.tencent-cloud.cn/raw/03a924ee0a53d86575f8067d1c97876d.png" target="_blank">模板名称</a>列表。</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。
	// 
	// 注：`1.offset从0开始，即第一页为0。`
	// `2.默认从第一页返回。`
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定每页返回的数据条数，和Offset参数配合使用。
	// 
	// 注：`1.默认值为20，单页做大值为200。`
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 通过指定[第三方应用的应用号（ApplicationId）](https://qcloudimg.tencent-cloud.cn/raw/60efa1e9049732e5246b20a268882b1a.png)，您可以查询<a href="https://qcloudimg.tencent-cloud.cn/raw/18319e5e77f7d47eab493d43d47827d3.png" target="_blank">【应用模板库管理】</a>中某个第三方应用下的模板。
	// 
	// <p><strong>注意事项：</strong></p>
	// <ul><li>当 <strong>ApplicationId</strong> 为空时（默认），系统将查询<a href="https://qcloudimg.tencent-cloud.cn/raw/376943a1d472393dd5388592f2e85ee5.png" target="_blank">平台企业的所有模板</a>（自建应用使用的模板）。</li><li>当 <strong>ApplicationId</strong> 不为空时，系统将从<a href="https://qcloudimg.tencent-cloud.cn/raw/18319e5e77f7d47eab493d43d47827d3.png" target="_blank">【应用模板库管理】</a>中查询该特定应用下的模板（分享给第三方应用子企业的模板）。</li></ul>
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 默认为false，查询SaaS模板库列表；
	// 为true，查询第三方应用集成平台企业模板库管理列表
	//
	// Deprecated: IsChannel is deprecated.
	IsChannel *bool `json:"IsChannel,omitnil,omitempty" name:"IsChannel"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 暂未开放
	//
	// Deprecated: GenerateSource is deprecated.
	GenerateSource *uint64 `json:"GenerateSource,omitnil,omitempty" name:"GenerateSource"`

	// 是否获取模板预览链接。
	// 
	// <ul><li><strong>false</strong>：不获取（默认）</li><li><strong>true</strong>：需要获取</li></ul>
	// 设置为true之后， 返回参数PreviewUrl，为模板的H5预览链接, 有效期5分钟。可以通过浏览器打开此链接预览模板，或者嵌入到iframe中预览模板。
	WithPreviewUrl *bool `json:"WithPreviewUrl,omitnil,omitempty" name:"WithPreviewUrl"`
}

type DescribeFlowTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 查询内容控制
	// 
	// <ul><li>**0**：模板列表及详情（默认）</li>
	// <li>**1**：仅模板列表</li></ul>
	ContentType *int64 `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 搜索过滤的条件，本字段允许您通过指定模板 ID 或模板名称来进行查询。
	// <ul><li><strong>模板 ID</strong>：<strong>Key</strong>设置为 <code>template-id</code> ，<strong>Values</strong>为您想要查询的 <a href="https://qcloudimg.tencent-cloud.cn/raw/5c27b917b2bbe8c341566c78ca6f8782.png" target="_blank">模板 ID </a>列表。</li>  <li><strong>主企业模板 ID</strong>：<strong>Key</strong>设置为 <code>share-template-id</code> ，<strong>Values</strong>为您想要查询的 <a href="https://qcloudimg.tencent-cloud.cn/raw/5c27b917b2bbe8c341566c78ca6f8782.png" target="_blank">主企业模板 ID </a>列表。用来查询主企业分享模板到子企业场景下，子企业的模板信息，在此情境下，参数 <strong>Agent.ProxyOrganizationId</strong>（子企业的组织ID）为必填项。</li> <li><strong>模板名称</strong>：<strong>Key</strong>设置为 <code>template-name</code> ，<strong>Values</strong>为您想要查询的<a href="https://qcloudimg.tencent-cloud.cn/raw/03a924ee0a53d86575f8067d1c97876d.png" target="_blank">模板名称</a>列表。</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。
	// 
	// 注：`1.offset从0开始，即第一页为0。`
	// `2.默认从第一页返回。`
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定每页返回的数据条数，和Offset参数配合使用。
	// 
	// 注：`1.默认值为20，单页做大值为200。`
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 通过指定[第三方应用的应用号（ApplicationId）](https://qcloudimg.tencent-cloud.cn/raw/60efa1e9049732e5246b20a268882b1a.png)，您可以查询<a href="https://qcloudimg.tencent-cloud.cn/raw/18319e5e77f7d47eab493d43d47827d3.png" target="_blank">【应用模板库管理】</a>中某个第三方应用下的模板。
	// 
	// <p><strong>注意事项：</strong></p>
	// <ul><li>当 <strong>ApplicationId</strong> 为空时（默认），系统将查询<a href="https://qcloudimg.tencent-cloud.cn/raw/376943a1d472393dd5388592f2e85ee5.png" target="_blank">平台企业的所有模板</a>（自建应用使用的模板）。</li><li>当 <strong>ApplicationId</strong> 不为空时，系统将从<a href="https://qcloudimg.tencent-cloud.cn/raw/18319e5e77f7d47eab493d43d47827d3.png" target="_blank">【应用模板库管理】</a>中查询该特定应用下的模板（分享给第三方应用子企业的模板）。</li></ul>
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 默认为false，查询SaaS模板库列表；
	// 为true，查询第三方应用集成平台企业模板库管理列表
	IsChannel *bool `json:"IsChannel,omitnil,omitempty" name:"IsChannel"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 暂未开放
	GenerateSource *uint64 `json:"GenerateSource,omitnil,omitempty" name:"GenerateSource"`

	// 是否获取模板预览链接。
	// 
	// <ul><li><strong>false</strong>：不获取（默认）</li><li><strong>true</strong>：需要获取</li></ul>
	// 设置为true之后， 返回参数PreviewUrl，为模板的H5预览链接, 有效期5分钟。可以通过浏览器打开此链接预览模板，或者嵌入到iframe中预览模板。
	WithPreviewUrl *bool `json:"WithPreviewUrl,omitnil,omitempty" name:"WithPreviewUrl"`
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
	delete(f, "Agent")
	delete(f, "ContentType")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ApplicationId")
	delete(f, "IsChannel")
	delete(f, "Organization")
	delete(f, "GenerateSource")
	delete(f, "WithPreviewUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowTemplatesResponseParams struct {
	// 模板详情列表数据
	Templates []*TemplateInfo `json:"Templates,omitnil,omitempty" name:"Templates"`

	// 查询到的模板总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeIntegrationDepartmentsRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 查询类型，支持以下类型：
	// <ul><li>**0**：查询单个部门节点列表，不包含子节点部门信息</li>
	// <li>**1**：查询单个部门节点级一级子节点部门信息列表</li></ul>
	QueryType *uint64 `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 查询的部门ID。
	// 注：`如果同时指定了DeptId与DeptOpenId参数，系统将优先使用DeptId参数进行查询。当二者都未指定时，系统将返回根节点部门数据。`
	DeptId *string `json:"DeptId,omitnil,omitempty" name:"DeptId"`

	// 查询的客户系统部门ID。
	// 注：`如果同时指定了DeptId与DeptOpenId参数，系统将优先使用DeptId参数进行查询。当二者都未指定时，系统将返回根节点部门数据。`
	DeptOpenId *string `json:"DeptOpenId,omitnil,omitempty" name:"DeptOpenId"`
}

type DescribeIntegrationDepartmentsRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 查询类型，支持以下类型：
	// <ul><li>**0**：查询单个部门节点列表，不包含子节点部门信息</li>
	// <li>**1**：查询单个部门节点级一级子节点部门信息列表</li></ul>
	QueryType *uint64 `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 查询的部门ID。
	// 注：`如果同时指定了DeptId与DeptOpenId参数，系统将优先使用DeptId参数进行查询。当二者都未指定时，系统将返回根节点部门数据。`
	DeptId *string `json:"DeptId,omitnil,omitempty" name:"DeptId"`

	// 查询的客户系统部门ID。
	// 注：`如果同时指定了DeptId与DeptOpenId参数，系统将优先使用DeptId参数进行查询。当二者都未指定时，系统将返回根节点部门数据。`
	DeptOpenId *string `json:"DeptOpenId,omitnil,omitempty" name:"DeptOpenId"`
}

func (r *DescribeIntegrationDepartmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationDepartmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "QueryType")
	delete(f, "Agent")
	delete(f, "DeptId")
	delete(f, "DeptOpenId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationDepartmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationDepartmentsResponseParams struct {
	// 部门信息列表。部门信息根据部门排序号OrderNo降序排列，根据部门创建时间升序排列。
	Departments []*IntegrationDepartment `json:"Departments,omitnil,omitempty" name:"Departments"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIntegrationDepartmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationDepartmentsResponseParams `json:"Response"`
}

func (r *DescribeIntegrationDepartmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationDepartmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationEmployeesRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写UserId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定分页每页返回的数据条数，单页最大支持 20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 查询的关键字段，支持Key-Values查询。可选键值如下：
	// <ul>
	//   <li>Key:**"Status"**，根据实名状态查询员工，Values可选：
	//     <ul><li>**["IsVerified"]**：查询已实名的员工</li><li>**["NotVerified"]**：查询未实名的员工</li></ul></li>
	//   <li>Key:**"DepartmentId"**，根据部门ID查询部门下的员工，Values为指定的部门ID：**["DepartmentId"]**</li>
	//   <li>Key:**"UserId"**，根据用户ID查询员工，Values为指定的用户ID：**["UserId"]**</li>
	//   <li>Key:**"UserWeWorkOpenId"**，根据用户企微账号ID查询员工，Values为指定用户的企微账号ID：**["UserWeWorkOpenId"]**</li>
	//   <li>Key:**"StaffOpenId"**，根据第三方系统用户OpenId查询员工，Values为第三方系统用户OpenId列表：**["OpenId1","OpenId2",...]**</li>
	//   <li>Key:**"RoleId"**，根据电子签角色ID查询员工，Values为指定的角色ID，满足其中任意一个角色即可：**["RoleId1","RoleId2",...]**</li>
	// </ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页。页码从 0 开始，即首页为 0，最大20000。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeIntegrationEmployeesRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写UserId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定分页每页返回的数据条数，单页最大支持 20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 查询的关键字段，支持Key-Values查询。可选键值如下：
	// <ul>
	//   <li>Key:**"Status"**，根据实名状态查询员工，Values可选：
	//     <ul><li>**["IsVerified"]**：查询已实名的员工</li><li>**["NotVerified"]**：查询未实名的员工</li></ul></li>
	//   <li>Key:**"DepartmentId"**，根据部门ID查询部门下的员工，Values为指定的部门ID：**["DepartmentId"]**</li>
	//   <li>Key:**"UserId"**，根据用户ID查询员工，Values为指定的用户ID：**["UserId"]**</li>
	//   <li>Key:**"UserWeWorkOpenId"**，根据用户企微账号ID查询员工，Values为指定用户的企微账号ID：**["UserWeWorkOpenId"]**</li>
	//   <li>Key:**"StaffOpenId"**，根据第三方系统用户OpenId查询员工，Values为第三方系统用户OpenId列表：**["OpenId1","OpenId2",...]**</li>
	//   <li>Key:**"RoleId"**，根据电子签角色ID查询员工，Values为指定的角色ID，满足其中任意一个角色即可：**["RoleId1","RoleId2",...]**</li>
	// </ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页。页码从 0 开始，即首页为 0，最大20000。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeIntegrationEmployeesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationEmployeesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Limit")
	delete(f, "Agent")
	delete(f, "Filters")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationEmployeesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationEmployeesResponseParams struct {
	// 员工信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Employees []*Staff `json:"Employees,omitnil,omitempty" name:"Employees"`

	// 指定分页返回第几页的数据。页码从 0 开始，即首页为 0，最大20000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定分页每页返回的数据条数，单页最大支持 20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 符合条件的员工数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIntegrationEmployeesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationEmployeesResponseParams `json:"Response"`
}

func (r *DescribeIntegrationEmployeesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationEmployeesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationRolesRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写UserId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定分页每页返回的数据条数，单页最大支持 200。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 查询的关键字段，支持Key-Value单值查询。可选键值对如下：
	// <ul>
	//   <li>Key:"RoleType"，查询角色类型，Values可选：
	//     <ul><li>**"1"**：查询系统角色</li><li>**"2"**：查询自定义角色</li></ul>
	//   </li><li>Key:"RoleStatus"，查询角色状态，Values可选：
	//     <ul><li>**"1"**：查询启用角色</li><li>**"2"**：查询禁用角色</li></ul>
	//   </li><li>Key:"IsGroupRole"，是否查询集团角色，Values可选：
	//     <ul><li>**"0"**：查询非集团角色</li><li>**"1"**：查询集团角色</li></ul>
	//   </li><li>Key:"IsReturnPermissionGroup"，是否返回角色对应权限树，Values可选：
	//     <ul><li>**"0"**：接口不返回角色对应的权限树字段</li><li>**"1"**：接口返回角色对应的权限树字段</li></ul>
	//   </li>
	// </ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// OFFSET 用于指定查询结果的偏移量，如果不传默认偏移为0,最大2000。
	// 分页参数, 需要limit, offset 配合使用
	// 例如:
	// 您希望得到第三页的数据, 且每页限制最多10条
	// 您可以使用 LIMIT 10 OFFSET 20
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeIntegrationRolesRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写UserId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定分页每页返回的数据条数，单页最大支持 200。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 查询的关键字段，支持Key-Value单值查询。可选键值对如下：
	// <ul>
	//   <li>Key:"RoleType"，查询角色类型，Values可选：
	//     <ul><li>**"1"**：查询系统角色</li><li>**"2"**：查询自定义角色</li></ul>
	//   </li><li>Key:"RoleStatus"，查询角色状态，Values可选：
	//     <ul><li>**"1"**：查询启用角色</li><li>**"2"**：查询禁用角色</li></ul>
	//   </li><li>Key:"IsGroupRole"，是否查询集团角色，Values可选：
	//     <ul><li>**"0"**：查询非集团角色</li><li>**"1"**：查询集团角色</li></ul>
	//   </li><li>Key:"IsReturnPermissionGroup"，是否返回角色对应权限树，Values可选：
	//     <ul><li>**"0"**：接口不返回角色对应的权限树字段</li><li>**"1"**：接口返回角色对应的权限树字段</li></ul>
	//   </li>
	// </ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// OFFSET 用于指定查询结果的偏移量，如果不传默认偏移为0,最大2000。
	// 分页参数, 需要limit, offset 配合使用
	// 例如:
	// 您希望得到第三页的数据, 且每页限制最多10条
	// 您可以使用 LIMIT 10 OFFSET 20
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeIntegrationRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Limit")
	delete(f, "Agent")
	delete(f, "Filters")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationRolesResponseParams struct {
	// OFFSET 用于指定查询结果的偏移量，如果不传默认偏移为0, 最大为2000
	// 分页参数, 需要limit, offset 配合使用
	// 例如:
	// 您希望得到第三页的数据, 且每页限制最多10条
	// 您可以使用 LIMIT 10 OFFSET 20
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定分页每页返回的数据条数，单页最大支持 200。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 符合查询条件的总角色数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 企业角色信息列表。
	IntegrateRoles []*IntegrateRole `json:"IntegrateRoles,omitnil,omitempty" name:"IntegrateRoles"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIntegrationRolesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationRolesResponseParams `json:"Response"`
}

func (r *DescribeIntegrationRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationAuthStatusRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。 支持填入集团子公司经办人 userId 代发合同。  注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 组织机构名称。 请确认该名称与企业营业执照中注册的名称一致。 如果名称中包含英文括号()，请使用中文括号（）代替。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 企业统一社会信用代码
	// 注意：OrganizationName和UniformSocialCreditCode不能同时为空
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// 法人姓名
	LegalName *string `json:"LegalName,omitnil,omitempty" name:"LegalName"`
}

type DescribeOrganizationAuthStatusRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。 支持填入集团子公司经办人 userId 代发合同。  注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 组织机构名称。 请确认该名称与企业营业执照中注册的名称一致。 如果名称中包含英文括号()，请使用中文括号（）代替。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 企业统一社会信用代码
	// 注意：OrganizationName和UniformSocialCreditCode不能同时为空
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// 法人姓名
	LegalName *string `json:"LegalName,omitnil,omitempty" name:"LegalName"`
}

func (r *DescribeOrganizationAuthStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationAuthStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "OrganizationName")
	delete(f, "UniformSocialCreditCode")
	delete(f, "LegalName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationAuthStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationAuthStatusResponseParams struct {
	// 企业是否已认证
	IsVerified *bool `json:"IsVerified,omitnil,omitempty" name:"IsVerified"`

	// 企业认证状态 0-未认证 1-认证中 2-已认证
	AuthStatus *int64 `json:"AuthStatus,omitnil,omitempty" name:"AuthStatus"`

	// 企业认证信息
	AuthRecords []*AuthRecord `json:"AuthRecords,omitnil,omitempty" name:"AuthRecords"`

	// 企业在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 可登录腾讯电子签控制台，在 "更多"->"企业设置"->"企业中心"- 中查看企业电子签账号。
	// p.s. 只有当前企业认证成功的时候返回
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationAuthStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationAuthStatusResponseParams `json:"Response"`
}

func (r *DescribeOrganizationAuthStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationAuthStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationGroupOrganizationsRequestParams struct {
	// 执行本接口操作的员工信息,userId必填。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定分页每页返回的数据条数，单页最大1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询成员企业的企业名，模糊匹配
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 成员企业加入集团的当前状态
	// <ul><li> **1**：待授权</li>
	// <li> **2**：已授权待激活</li>
	// <li> **3**：拒绝授权</li>
	// <li> **4**：已解除</li>
	// <li> **5**：已加入</li>
	// </ul>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否导出当前成员企业数据
	// <ul><li> **false**：不导出（默认值）</li>
	// <li> **true**：导出</li></ul>
	Export *bool `json:"Export,omitnil,omitempty" name:"Export"`

	// 成员企业机构 ID，32 位字符串，在PC控制台 集团管理可获取
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeOrganizationGroupOrganizationsRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息,userId必填。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定分页每页返回的数据条数，单页最大1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询成员企业的企业名，模糊匹配
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 成员企业加入集团的当前状态
	// <ul><li> **1**：待授权</li>
	// <li> **2**：已授权待激活</li>
	// <li> **3**：拒绝授权</li>
	// <li> **4**：已解除</li>
	// <li> **5**：已加入</li>
	// </ul>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否导出当前成员企业数据
	// <ul><li> **false**：不导出（默认值）</li>
	// <li> **true**：导出</li></ul>
	Export *bool `json:"Export,omitnil,omitempty" name:"Export"`

	// 成员企业机构 ID，32 位字符串，在PC控制台 集团管理可获取
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeOrganizationGroupOrganizationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationGroupOrganizationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Name")
	delete(f, "Status")
	delete(f, "Export")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationGroupOrganizationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationGroupOrganizationsResponseParams struct {
	// 符合查询条件的资源实例总数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 已授权待激活的子企业总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinedTotal *uint64 `json:"JoinedTotal,omitnil,omitempty" name:"JoinedTotal"`

	// 已加入的企业数量(废弃,请使用ActivatedTotal)
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ActivedTotal is deprecated.
	ActivedTotal *uint64 `json:"ActivedTotal,omitnil,omitempty" name:"ActivedTotal"`

	// 如果入参Export为 true 时使用，表示导出Excel的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExportUrl *string `json:"ExportUrl,omitnil,omitempty" name:"ExportUrl"`

	// 成员企业信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*GroupOrganization `json:"List,omitnil,omitempty" name:"List"`

	// 已加入的子企业总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivatedTotal *uint64 `json:"ActivatedTotal,omitnil,omitempty" name:"ActivatedTotal"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationGroupOrganizationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationGroupOrganizationsResponseParams `json:"Response"`
}

func (r *DescribeOrganizationGroupOrganizationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationGroupOrganizationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationSealsRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定分页每页返回的数据条数，如果不传默认为 20，单页最大支持 200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0，最大 20000。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询授权用户信息类型，取值如下：
	// 
	// <ul> <li><b>0</b>：（默认）不返回授权用户信息</li> <li><b>1</b>：返回授权用户的信息</li> </ul>
	InfoType *int64 `json:"InfoType,omitnil,omitempty" name:"InfoType"`

	// 印章id，是否查询特定的印章（没有输入返回所有）
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 印章种类列表（均为组织机构印章）。 若无特定需求，将展示所有类型的印章。 目前支持以下几种：<ul> <li><strong>OFFICIAL</strong>：企业公章；</li> <li><strong>CONTRACT</strong>：合同专用章；</li> <li><strong>ORGANIZATION_SEAL</strong>：企业印章（通过图片上传创建）；</li> <li><strong>LEGAL_PERSON_SEAL</strong>：法定代表人章。</li> <li><strong>EMPLOYEE_QUALIFICATION_SEAL</strong>：员工执业章。</li> </ul>
	SealTypes []*string `json:"SealTypes,omitnil,omitempty" name:"SealTypes"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需查询的印章状态列表。
	// <ul>
	// <li>空：（默认）仅查询启用状态的印章；</li>
	// <li><strong>ALL</strong>：查询所有状态的印章；</li>
	// <li><strong>CHECKING</strong>：查询待审核的印章；</li>
	// <li><strong>SUCCESS</strong>：查询启用状态的印章；</li>
	// <li><strong>FAIL</strong>：查询印章审核拒绝的印章；</li>
	// <li><strong>DISABLE</strong>：查询已停用的印章；</li>
	// <li><strong>STOPPED</strong>：查询已终止的印章；</li>
	// <li><strong>VOID</strong>：查询已作废的印章；</li>
	// <li><strong>INVALID</strong>：查询已失效的印章。</li>
	// </ul>
	SealStatuses []*string `json:"SealStatuses,omitnil,omitempty" name:"SealStatuses"`
}

type DescribeOrganizationSealsRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 指定分页每页返回的数据条数，如果不传默认为 20，单页最大支持 200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0，最大 20000。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询授权用户信息类型，取值如下：
	// 
	// <ul> <li><b>0</b>：（默认）不返回授权用户信息</li> <li><b>1</b>：返回授权用户的信息</li> </ul>
	InfoType *int64 `json:"InfoType,omitnil,omitempty" name:"InfoType"`

	// 印章id，是否查询特定的印章（没有输入返回所有）
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 印章种类列表（均为组织机构印章）。 若无特定需求，将展示所有类型的印章。 目前支持以下几种：<ul> <li><strong>OFFICIAL</strong>：企业公章；</li> <li><strong>CONTRACT</strong>：合同专用章；</li> <li><strong>ORGANIZATION_SEAL</strong>：企业印章（通过图片上传创建）；</li> <li><strong>LEGAL_PERSON_SEAL</strong>：法定代表人章。</li> <li><strong>EMPLOYEE_QUALIFICATION_SEAL</strong>：员工执业章。</li> </ul>
	SealTypes []*string `json:"SealTypes,omitnil,omitempty" name:"SealTypes"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 需查询的印章状态列表。
	// <ul>
	// <li>空：（默认）仅查询启用状态的印章；</li>
	// <li><strong>ALL</strong>：查询所有状态的印章；</li>
	// <li><strong>CHECKING</strong>：查询待审核的印章；</li>
	// <li><strong>SUCCESS</strong>：查询启用状态的印章；</li>
	// <li><strong>FAIL</strong>：查询印章审核拒绝的印章；</li>
	// <li><strong>DISABLE</strong>：查询已停用的印章；</li>
	// <li><strong>STOPPED</strong>：查询已终止的印章；</li>
	// <li><strong>VOID</strong>：查询已作废的印章；</li>
	// <li><strong>INVALID</strong>：查询已失效的印章。</li>
	// </ul>
	SealStatuses []*string `json:"SealStatuses,omitnil,omitempty" name:"SealStatuses"`
}

func (r *DescribeOrganizationSealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationSealsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InfoType")
	delete(f, "SealId")
	delete(f, "SealTypes")
	delete(f, "Agent")
	delete(f, "SealStatuses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationSealsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationSealsResponseParams struct {
	// 在设定了SealId时，返回值为0或1；若未设定SealId，则返回公司的总印章数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 查询到的印章结果数组
	Seals []*OccupiedSeal `json:"Seals,omitnil,omitempty" name:"Seals"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationSealsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationSealsResponseParams `json:"Response"`
}

func (r *DescribeOrganizationSealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationSealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationVerifyStatusRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。 支持填入集团子公司经办人 userId 代发合同。  注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。	
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DescribeOrganizationVerifyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。 支持填入集团子公司经办人 userId 代发合同。  注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。	
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *DescribeOrganizationVerifyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationVerifyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationVerifyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationVerifyStatusResponseParams struct {
	// 当前企业认证状态
	// <ul>
	// <li> <b>0 </b>:未认证</li>
	// <li> <b>1 </b> : 认证中</li>
	// <li> <b>2 </b> : 已认证</li>
	// </ul>
	VerifyStatus *uint64 `json:"VerifyStatus,omitnil,omitempty" name:"VerifyStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationVerifyStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationVerifyStatusResponseParams `json:"Response"`
}

func (r *DescribeOrganizationVerifyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationVerifyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonCertificateRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 个人用户的三要素信息：
	// <ul><li>姓名</li>
	// <li>证件号</li>
	// <li>证件类型</li></ul>
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 证书使用场景，可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`
}

type DescribePersonCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 个人用户的三要素信息：
	// <ul><li>姓名</li>
	// <li>证件号</li>
	// <li>证件类型</li></ul>
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 证书使用场景，可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`
}

func (r *DescribePersonCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "UserInfo")
	delete(f, "Agent")
	delete(f, "SceneKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePersonCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonCertificateResponseParams struct {
	// 证书的Base64
	Cert *string `json:"Cert,omitnil,omitempty" name:"Cert"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePersonCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DescribePersonCertificateResponseParams `json:"Response"`
}

func (r *DescribePersonCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSignFaceVideoRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。<br/>注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署参与人在本流程中的编号ID(每个流程不同)，可用此ID来定位签署参与人在本流程的签署节点，也可用于后续创建签署链接等操作。
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DescribeSignFaceVideoRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。<br/>注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署参与人在本流程中的编号ID(每个流程不同)，可用此ID来定位签署参与人在本流程的签署节点，也可用于后续创建签署链接等操作。
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *DescribeSignFaceVideoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSignFaceVideoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "SignId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSignFaceVideoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSignFaceVideoResponseParams struct {
	// 核身视频结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoData *DetectInfoVideoData `json:"VideoData,omitnil,omitempty" name:"VideoData"`

	// 意愿核身问答模式结果。若未使用该意愿核身功能，该字段返回值可以不处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionQuestionResult *IntentionQuestionResult `json:"IntentionQuestionResult,omitnil,omitempty" name:"IntentionQuestionResult"`

	// 意愿核身点头确认模式的结果信息，若未使用该意愿核身功能，该字段返回值可以不处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionActionResult *IntentionActionResult `json:"IntentionActionResult,omitnil,omitempty" name:"IntentionActionResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSignFaceVideoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSignFaceVideoResponseParams `json:"Response"`
}

func (r *DescribeSignFaceVideoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSignFaceVideoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeThirdPartyAuthCodeRequestParams struct {
	// 腾讯电子签小程序跳转客户企业小程序时携带的授权查看码，AuthCode由腾讯电子签小程序生成。
	AuthCode *string `json:"AuthCode,omitnil,omitempty" name:"AuthCode"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DescribeThirdPartyAuthCodeRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯电子签小程序跳转客户企业小程序时携带的授权查看码，AuthCode由腾讯电子签小程序生成。
	AuthCode *string `json:"AuthCode,omitnil,omitempty" name:"AuthCode"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
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
	delete(f, "Operator")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeThirdPartyAuthCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeThirdPartyAuthCodeResponseParams struct {
	// AuthCode 中对应个人用户是否实名
	// <ul>
	// <li> **VERIFIED** : 此个人已实名</li>
	// <li> **UNVERIFIED**: 此个人未实名</li></ul>
	VerifyStatus *string `json:"VerifyStatus,omitnil,omitempty" name:"VerifyStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeUserAutoSignStatusRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 要查询状态的用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DescribeUserAutoSignStatusRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 要查询状态的用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *DescribeUserAutoSignStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserAutoSignStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserAutoSignStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserAutoSignStatusResponseParams struct {
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

type DescribeUserAutoSignStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserAutoSignStatusResponseParams `json:"Response"`
}

func (r *DescribeUserAutoSignStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserAutoSignStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserVerifyStatusRequestParams struct {
	// 用户信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`
}

type DescribeUserVerifyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 用户信息
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil,omitempty" name:"IdCardType"`
}

func (r *DescribeUserVerifyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserVerifyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Name")
	delete(f, "IdCardNumber")
	delete(f, "IdCardType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserVerifyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserVerifyStatusResponseParams struct {
	// true:表示已实名
	// false：表示未实名
	VerifyStatus *bool `json:"VerifyStatus,omitnil,omitempty" name:"VerifyStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserVerifyStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserVerifyStatusResponseParams `json:"Response"`
}

func (r *DescribeUserVerifyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserVerifyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectInfoVideoData struct {
	// 活体视频的base64编码，mp4格式
	// 
	// 注:`需进行base64解码获取活体视频文件`
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveNessVideo *string `json:"LiveNessVideo,omitnil,omitempty" name:"LiveNessVideo"`
}

// Predefined struct for user
type DisableUserAutoSignRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 需要关闭自动签的个人的信息，如姓名，证件信息等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type DisableUserAutoSignRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 需要关闭自动签的个人的信息，如姓名，证件信息等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *DisableUserAutoSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableUserAutoSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableUserAutoSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableUserAutoSignResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableUserAutoSignResponse struct {
	*tchttp.BaseResponse
	Response *DisableUserAutoSignResponseParams `json:"Response"`
}

func (r *DisableUserAutoSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableUserAutoSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DynamicFlowApproverResult struct {
	// 签署方角色编号，签署方角色编号是用于区分同一个流程中不同签署方的唯一标识。不同的流程会出现同样的签署方角色编号。
	// 
	// 填写控件和签署控件都与特定的角色编号关联。
	// 
	// 在进行新增签署方操作时，建议记录下该签署方的角色编号。后续可以拉取流程信息，用来判断该签署方的当前状态。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 签署方唯一编号，一个全局唯一的标识符，不同的流程不会出现冲突。
	// 
	// 可以使用签署方的唯一编号来生成签署链接（也可以通过RecipientId来生成签署链接）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 签署方当前状态，会出现下面的状态
	// 
	// 2：待签署
	// 3：已签署
	// 4：已拒绝
	// 5：已过期
	// 6：已撤销
	// 8：待填写
	// 9：因为各种原因（签署人改名等）而终止
	// 10：填写完成
	// 15：已解除
	// 19：转他人处理
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverStatus *int64 `json:"ApproverStatus,omitnil,omitempty" name:"ApproverStatus"`
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
	SkipUploadFile *bool `json:"SkipUploadFile,omitnil,omitempty" name:"SkipUploadFile"`
}

type ExtendAuthInfo struct {
	// 扩展服务的类型，可能是以下值：
	// <ul><li>OPEN_SERVER_SIGN：企业自动签署</li>
	// <li>BATCH_SIGN：批量签署</li>
	// <li>OVERSEA_SIGN：企业与港澳台居民签署合同</li>
	// <li>AGE_LIMIT_EXPANSION：拓宽签署方年龄限制</li>
	// <li>MOBILE_CHECK_APPROVER：个人签署方仅校验手机号</li>
	// <li>HIDE_OPERATOR_DISPLAY：隐藏合同经办人姓名</li>
	// <li>ORGANIZATION_OCR_FALLBACK：正楷临摹签名失败后更换其他签名类型</li>
	// <li>ORGANIZATION_FLOW_NOTIFY_TYPE：短信通知签署方</li>
	// <li>HIDE_ONE_KEY_SIGN：个人签署方手动签字</li>
	// <li>PAGING_SEAL：骑缝章</li>
	// <li>ORGANIZATION_FLOW_PASSWD_NOTIFY：签署密码开通引导</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 扩展服务的名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 扩展服务的开通状态：
	// <ul>
	// <li>ENABLE : 已开通</li>
	// <li>DISABLE : 未开通</li>
	// </ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作扩展服务的操作人UserId，员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUserId *string `json:"OperatorUserId,omitnil,omitempty" name:"OperatorUserId"`

	// 扩展服务的操作时间，格式为Unix标准时间戳（秒）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateOn *int64 `json:"OperateOn,omitnil,omitempty" name:"OperateOn"`

	// 该扩展服务若可以授权，此参数对应授权人员的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasAuthUserList []*HasAuthUser `json:"HasAuthUserList,omitnil,omitempty" name:"HasAuthUserList"`
}

type ExtendScene struct {
	// 印章来源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GenerateType *string `json:"GenerateType,omitnil,omitempty" name:"GenerateType"`

	// 印章来源类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	GenerateTypeDesc *string `json:"GenerateTypeDesc,omitnil,omitempty" name:"GenerateTypeDesc"`

	// 印章来源logo
	// 注意：此字段可能返回 null，表示取不到有效值。
	GenerateTypeLogo *string `json:"GenerateTypeLogo,omitnil,omitempty" name:"GenerateTypeLogo"`
}

type FailedCreateRoleData struct {
	// 用户userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 角色id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type FailedCreateStaffData struct {
	// 员工名
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 员工手机号
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 传入的企微账号id
	WeworkOpenId *string `json:"WeworkOpenId,omitnil,omitempty" name:"WeworkOpenId"`

	// 失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type FailedDeleteStaffData struct {
	// 员工在电子签的userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 员工在第三方平台的openId
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type FailedUpdateStaffData struct {
	// 用户传入的名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户传入的手机号，明文展示
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查看某位员工的UserId(在页面中展示为用户ID)。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 员工在第三方平台的openId
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`
}

type FileInfo struct {
	// 文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件大小，单位为Byte
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 文件上传时间，格式为Unix标准时间戳（秒）
	CreatedOn *int64 `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`
}

type FileUrl struct {
	// 下载文件的URL，有效期为输入的UrlTtl，默认5分钟
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 下载文件的附加信息。如果是pdf文件，会返回pdf文件每页的有效高宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	Option *string `json:"Option,omitnil,omitempty" name:"Option"`
}

type FillApproverInfo struct {
	// 签署方经办人在模板中配置的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。
	// 模板发起合同时，该参数为必填项。
	// 文件发起合同是，该参数无需传值。
	// 如果开发者后序用合同模板发起合同，建议保存此值，在用合同模板发起合同中需此值绑定对应的签署经办人 。
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 签署人来源
	// WEWORKAPP: 企业微信
	// <br/>仅【企微或签】时指定WEWORKAPP
	ApproverSource *string `json:"ApproverSource,omitnil,omitempty" name:"ApproverSource"`

	// 企业微信UserId
	// <br/>当ApproverSource为WEWORKAPP的企微或签场景下，必须指企业自有应用获取企业微信的UserId
	CustomUserId *string `json:"CustomUserId,omitnil,omitempty" name:"CustomUserId"`

	// 补充企业签署人员工姓名
	ApproverName *string `json:"ApproverName,omitnil,omitempty" name:"ApproverName"`

	// 补充企业签署人员工手机号
	ApproverMobile *string `json:"ApproverMobile,omitnil,omitempty" name:"ApproverMobile"`

	// 补充企业动态签署人时，需要指定对应企业名称
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

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
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	// 
	// 注：`补充个人签署方时，若该用户已在电子签完成实名则可通过指定姓名和证件类型、证件号码完成补充。`
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitnil,omitempty" name:"ApproverIdCardNumber"`

	// 合同流程ID
	// - 补充合同组子合同动态签署人时必传。
	// - 补充普通合同时，请阅读：<a href="https://qian.tencent.com/developers/companyApis/operateFlows/CreateFlowApprovers/" target="_blank">补充签署人接口</a>的接口使用说明
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type FillError struct {
	// 为签署方经办人在签署合同中的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。与入参中补充的签署人角色ID对应，批量补充部分失败返回对应的错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 补充失败错误说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitnil,omitempty" name:"ErrMessage"`

	// 合同流程ID，为32位字符串。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type FilledComponent struct {
	// 控件Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentId *string `json:"ComponentId,omitnil,omitempty" name:"ComponentId"`

	// 控件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// 控件填写状态；0-未填写；1-已填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentFillStatus *string `json:"ComponentFillStatus,omitnil,omitempty" name:"ComponentFillStatus"`

	// 控件填写内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentValue *string `json:"ComponentValue,omitnil,omitempty" name:"ComponentValue"`

	// 控件所属参与方Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentRecipientId *string `json:"ComponentRecipientId,omitnil,omitempty" name:"ComponentRecipientId"`

	// 图片填充控件下载链接，如果是图片填充控件时，这里返回图片的下载链接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type Filter struct {
	// 查询过滤条件的Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 查询过滤条件的Value列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FlowApproverDetail struct {
	// 签署时的相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveMessage *string `json:"ApproveMessage,omitnil,omitempty" name:"ApproveMessage"`

	// 签署方姓名
	ApproveName *string `json:"ApproveName,omitnil,omitempty" name:"ApproveName"`

	// 签署方的签署状态
	// 0：还没有发起
	// 1：流程中 没有开始处理
	// 2：待签署
	// 3：已签署
	// 4：已拒绝
	// 5：已过期
	// 6：已撤销
	// 7：还没有预发起
	// 8：待填写
	// 9：因为各种原因而终止
	// 10：填写完成
	// 15：已解除
	// 19：转他人处理
	ApproveStatus *int64 `json:"ApproveStatus,omitnil,omitempty" name:"ApproveStatus"`

	// 模板配置中的参与方ID,与控件绑定
	ReceiptId *string `json:"ReceiptId,omitnil,omitempty" name:"ReceiptId"`

	// 客户自定义的用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomUserId *string `json:"CustomUserId,omitnil,omitempty" name:"CustomUserId"`

	// 签署人手机号
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 签署顺序，如果是有序签署，签署顺序从小到大
	SignOrder *int64 `json:"SignOrder,omitnil,omitempty" name:"SignOrder"`

	// 签署人签署时间，时间戳，单位秒
	ApproveTime *int64 `json:"ApproveTime,omitnil,omitempty" name:"ApproveTime"`

	// 签署方类型，ORGANIZATION-企业员工，PERSON-个人，ENTERPRISESERVER-企业静默签
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveType *string `json:"ApproveType,omitnil,omitempty" name:"ApproveType"`

	// 签署方侧用户来源，如WEWORKAPP-企业微信等
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverSource *string `json:"ApproverSource,omitnil,omitempty" name:"ApproverSource"`

	// 客户自定义签署方标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomApproverTag *string `json:"CustomApproverTag,omitnil,omitempty" name:"CustomApproverTag"`

	// 签署方企业Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 签署方企业名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 签署参与人在本流程中的编号ID（每个流程不同），可用此ID来定位签署参与人在本流程的签署节点，也可用于后续创建签署链接等操作。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 自定义签署人角色
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverRoleName *string `json:"ApproverRoleName,omitnil,omitempty" name:"ApproverRoleName"`
}

type FlowApproverUrlInfo struct {
	// 签署短链接。
	// 注意:
	// 1. 该链接有效期为<b>30分钟</b>，同时需要注意保密，不要外泄给无关用户。
	// 2. 该链接不支持小程序嵌入，仅支持<b>移动端浏览器</b>打开。
	// 3. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	SignUrl *string `json:"SignUrl,omitnil,omitempty" name:"SignUrl"`

	// 签署人类型。
	// - **1**: 个人
	ApproverType *int64 `json:"ApproverType,omitnil,omitempty" name:"ApproverType"`

	// 签署人姓名。
	ApproverName *string `json:"ApproverName,omitnil,omitempty" name:"ApproverName"`

	// 签署人手机号。
	ApproverMobile *string `json:"ApproverMobile,omitnil,omitempty" name:"ApproverMobile"`

	// 签署长链接。
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

type FlowBrief struct {
	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 合同流程的名称。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 合同流程描述信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowDescription *string `json:"FlowDescription,omitnil,omitempty" name:"FlowDescription"`

	// 合同流程的类别分类（如销售合同/入职合同等）。
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 合同流程当前的签署状态, 会存在下列的状态值
	// <ul><li> **0** : 未开启流程(合同中不存在填写环节)</li>
	// <li> **1** : 待签署</li>
	// <li> **2** : 部分签署</li>
	// <li> **3** : 已拒签</li>
	// <li> **4** : 已签署</li>
	// <li> **5** : 已过期</li>
	// <li> **6** : 已撤销</li>
	// <li> **7** : 未开启流程(合同中存在填写环节)</li>
	// <li> **8** : 等待填写</li>
	// <li> **9** : 部分填写</li>
	// <li> **10** : 已拒填</li>
	// <li> **21** : 已解除</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowStatus *int64 `json:"FlowStatus,omitnil,omitempty" name:"FlowStatus"`

	// 合同流程创建时间，格式为Unix标准时间戳（秒）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedOn *int64 `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 当合同流程状态为已拒签（即 FlowStatus=3）或已撤销（即 FlowStatus=6）时，此字段 FlowMessage 为拒签或撤销原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMessage *string `json:"FlowMessage,omitnil,omitempty" name:"FlowMessage"`

	//  合同流程发起方的员工编号, 即员工在腾讯电子签平台的唯一身份标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`
}

type FlowCreateApprover struct {
	// 在指定签署方时，可以选择企业B端或个人C端等不同的参与者类型，可选类型如下：
	// 
	// <ul><li> <b>0</b> :企业B端。</li>
	// <li> <b>1</b> :个人C端。</li>
	// <li> <b>3</b> :企业B端静默（自动）签署，无需签署人参与，自动签署可以参考<a href="https://qian.tencent.com/developers/company/autosign_guide" target="_blank" rel="noopener noreferrer">自动签署使用说明</a>文档。</li>
	// <li> <b>7</b> :个人C端自动签署，适用于个人自动签场景。注: <b>个人自动签场景为白名单功能，使用前请联系对接的客户经理沟通。</b> </li></ul>
	ApproverType *int64 `json:"ApproverType,omitnil,omitempty" name:"ApproverType"`

	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 
	// 注: `当approverType=0(企业签署方) 或 approverType=3(企业静默签署)时，必须指定`
	// 
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 在未指定签署人电子签UserId情况下，为必填参数
	ApproverName *string `json:"ApproverName,omitnil,omitempty" name:"ApproverName"`

	// 签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。 此手机号用于通知和用户的实名认证等环境，请确认手机号所有方为此合同签署方。
	// 
	// 注：`在未指定签署人电子签UserId情况下，为必填参数`
	ApproverMobile *string `json:"ApproverMobile,omitnil,omitempty" name:"ApproverMobile"`

	// 证件类型，支持以下类型
	// <ul><li><b>ID_CARD</b>: 居民身份证 (默认值)</li>
	// <li><b>HONGKONG_AND_MACAO</b> : 港澳居民来往内地通行证</li>
	// <li><b>HONGKONG_MACAO_AND_TAIWAN</b> : 港澳台居民居住证(格式同居民身份证)</li></ul>
	ApproverIdCardType *string `json:"ApproverIdCardType,omitnil,omitempty" name:"ApproverIdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitnil,omitempty" name:"ApproverIdCardNumber"`

	// 签署方经办人在模板中配置的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。
	// 
	// <b>模板发起合同时，该参数为必填项，可以通过[查询模板信息接口](https://qian.tencent.com/developers/companyApis/templatesAndFiles/DescribeFlowTemplates)获得。</b>
	// <b>文件发起合同时，该参数无需传值。</b>
	// 
	// 如果开发者后续用合同模板发起合同，建议保存此值，在用合同模板发起合同中需此值绑定对应的签署经办人 。
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 签署意愿确认渠道，默认为WEIXINAPP:人脸识别
	// 
	// 注: <font color="red">将要废弃</font >, `用ApproverSignTypes签署人签署合同时的认证方式代替, 新客户可请用ApproverSignTypes来设置`
	VerifyChannel []*string `json:"VerifyChannel,omitnil,omitempty" name:"VerifyChannel"`

	// 通知签署方经办人的方式,  有以下途径:
	// <ul><li>  **sms**  :  (默认)短信</li>
	// <li>   **none**   : 不通知</li></ul>
	// 
	// 注: `既是发起方又是签署方时，不给此签署方发送短信`
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// 合同强制需要阅读全文，无需传此参数
	IsFullText *bool `json:"IsFullText,omitnil,omitempty" name:"IsFullText"`

	// 签署方在签署合同之前，需要强制阅读合同的时长，可指定为3秒至300秒之间的任意值。
	// 
	// 若未指定阅读时间，则会按照合同页数大小计算阅读时间，计算规则如下：
	// <ul>
	// <li>合同页数少于等于2页，阅读时间为3秒；</li>
	// <li>合同页数为3到5页，阅读时间为5秒；</li>
	// <li>合同页数大于等于6页，阅读时间为10秒。</li>
	// </ul>
	PreReadTime *uint64 `json:"PreReadTime,omitnil,omitempty" name:"PreReadTime"`

	// 签署人userId，仅支持本企业的员工userid， 可在控制台组织管理处获得
	// 
	// 注： 
	// 如果传进来的<font color="red">UserId已经实名， 则忽略ApproverName，ApproverIdCardType，ApproverIdCardNumber，ApproverMobile这四个入参</font>（会用此UserId实名的身份证和登录的手机号覆盖）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <font color="red">字段已经废弃</font>，当前只支持true，默认为true
	Required *bool `json:"Required,omitnil,omitempty" name:"Required"`

	// 在企微场景下使用，需设置参数为**WEWORKAPP**，以表明合同来源于企微。
	ApproverSource *string `json:"ApproverSource,omitnil,omitempty" name:"ApproverSource"`

	// 在企业微信场景下，表明该合同流程为或签，其最大长度为64位字符串。
	// 所有参与或签的人员均需具备该标识。
	// 注意，在合同中，不同的或签参与人必须保证其CustomApproverTag唯一。
	// 如果或签签署人为本方企业微信参与人，则需要指定ApproverSource参数为WEWORKAPP。
	CustomApproverTag *string `json:"CustomApproverTag,omitnil,omitempty" name:"CustomApproverTag"`

	// 已经废弃, 快速注册相关信息
	RegisterInfo *RegisterInfo `json:"RegisterInfo,omitnil,omitempty" name:"RegisterInfo"`

	// 签署人个性化能力值，如是否可以转发他人处理、是否可以拒签、是否为动态补充签署人等功能开关。
	ApproverOption *ApproverOption `json:"ApproverOption,omitnil,omitempty" name:"ApproverOption"`

	// 签署完前端跳转的url，暂未使用
	//
	// Deprecated: JumpUrl is deprecated.
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 签署人的签署ID
	// 
	// <ul>
	// <li>在CreateFlow、CreatePrepareFlow等发起流程时不需要传入此参数，电子签后台系统会自动生成。</li>
	// <li>在CreateFlowSignUrl、CreateBatchQuickSignUrl等生成签署链接时，可以通过查询详情接口获取签署人的SignId，然后可以将此值传入，为该签署人创建签署链接。这样可以避免重复传输姓名、手机号、证件号等其他信息。</li>
	// </ul>
	SignId *string `json:"SignId,omitnil,omitempty" name:"SignId"`

	// 此签署人(员工或者个人)签署时，是否需要发起方企业审批，取值如下：
	// <ul><li>**false**：（默认）不需要审批，直接签署。</li>
	// <li>**true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li>如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li>如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 
	// 注：`此功能可用于与发起方企业内部的审批流程进行关联，支持手动、静默签署合同`
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/b14d5188ed0229d1401e74a9a49cab6d.png)
	ApproverNeedSignReview *bool `json:"ApproverNeedSignReview,omitnil,omitempty" name:"ApproverNeedSignReview"`

	// 签署人签署控件， 此参数仅针对文件发起（CreateFlowByFiles）生效
	// 
	// 合同中的签署控件列表，列表中可支持下列多种签署控件,控件的详细定义参考开发者中心的Component结构体
	// <ul><li> 个人签名/印章</li>
	// <li> 企业印章</li>
	// <li> 骑缝章等签署控件</li></ul>
	// 
	// `此参数仅针对文件发起设置生效,模板发起合同签署流程, 请以模板配置为主`
	SignComponents []*Component `json:"SignComponents,omitnil,omitempty" name:"SignComponents"`

	// 签署人填写控件 此参数仅针对文件发起（CreateFlowByFiles）生效
	// 
	// 合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li>单行文本控件</li>
	// <li>多行文本控件</li>
	// <li>勾选框控件</li>
	// <li>数字控件</li>
	// <li>图片控件</li>
	// <li>动态表格等填写控件</li></ul>
	// 
	// `此参数仅针对文件发起设置生效,模板发起合同签署流程, 请以模板配置为主`
	Components []*Component `json:"Components,omitnil,omitempty" name:"Components"`

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

	// 指定个人签署方查看合同的校验方式,可以传值如下:
	// <ul><li>  **1**   : （默认）人脸识别,人脸识别后才能合同内容</li>
	// <li>  **2**  : 手机号验证, 用户手机号和参与方手机号(ApproverMobile)相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// </li></ul>
	// 注: 
	// <ul><li>如果合同流程设置ApproverVerifyType查看合同的校验方式,    则忽略此签署人的查看合同的校验方式</li>
	// <li>此字段可传多个校验方式</li></ul>
	// 
	// `此参数仅针对文件发起设置生效,模板发起合同签署流程, 请以模板配置为主`
	// 
	// .
	ApproverVerifyTypes []*int64 `json:"ApproverVerifyTypes,omitnil,omitempty" name:"ApproverVerifyTypes"`

	// 您可以指定签署方签署合同的认证校验方式，可传递以下值：
	// <ul><li>**1**：人脸认证，需进行人脸识别成功后才能签署合同；</li>
	// <li>**2**：签署密码，需输入与用户在腾讯电子签设置的密码一致才能校验成功进行合同签署；</li>
	// <li>**3**：运营商三要素，需到运营商处比对手机号实名信息（名字、手机号、证件号）校验一致才能成功进行合同签署。（如果是港澳台客户，建议不要选择这个）</li>
	// <li>**5**：设备指纹识别，需要对比手机机主预留的指纹信息，校验一致才能成功进行合同签署。（iOS系统暂不支持该校验方式）</li>
	// <li>**6**：设备面容识别，需要对比手机机主预留的人脸信息，校验一致才能成功进行合同签署。（Android系统暂不支持该校验方式）</li></ul>
	// 
	// 注：
	// <ul><li>默认情况下，认证校验方式为人脸认证和签署密码两种形式；</li>
	// <li>您可以传递多种值，表示可用多种认证校验方式。</li>
	// <li>校验方式不允许只包含设备指纹识别和设备面容识别，至少需要再增加一种其他校验方式。</li>
	// <li>设备指纹识别和设备面容识别只支持小程序使用，其他端暂不支持。</li></ul>
	// 
	// 注:
	// `此参数仅针对文件发起设置生效,模板发起合同签署流程, 请以模板配置为主`
	ApproverSignTypes []*uint64 `json:"ApproverSignTypes,omitnil,omitempty" name:"ApproverSignTypes"`

	// 生成H5签署链接时，您可以指定签署方签署合同的认证校验方式的选择模式，可传递一下值：
	// <ul><li>**0**：签署方自行选择，签署方可以从预先指定的认证方式中自由选择；</li>
	// <li>**1**：自动按顺序首位推荐，签署方无需选择，系统会优先推荐使用第一种认证方式。</li></ul>
	// 注：
	// `不指定该值时，默认为签署方自行选择。`
	SignTypeSelector *uint64 `json:"SignTypeSelector,omitnil,omitempty" name:"SignTypeSelector"`

	// 签署人的签署截止时间，格式为Unix标准时间戳（秒）, 超过此时间未签署的合同变成已过期状态，不能在继续签署
	// 
	// 注: `若不设置此参数，则默认使用合同的截止时间，此参数暂不支持合同组子合同`
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// <b>只有在生成H5签署链接的情形下</b>（ 如调用<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowSignUrl" target="_blank">获取H5签署链接</a>、<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateBatchQuickSignUrl" target="_blank">获取H5批量签署链接</a>等接口），该配置才会生效。
	// 
	// 您可以指定H5签署视频核身的意图配置，选择问答模式或点头模式的语音文本。
	// 
	// 注意：
	// 1. 视频认证为<b>白名单功能，使用前请联系对接的客户经理沟通</b>。
	// 2. 使用视频认证时，<b>生成H5签署链接必须将签署认证方式指定为人脸</b>（即ApproverSignTypes设置成人脸签署）。
	// 3. 签署完成后，可以通过<a href="https://qian.tencent.com/developers/companyApis/queryFlows/DescribeSignFaceVideo" target="_blank">查询签署认证人脸视频</a>获取到当时的视频。
	Intention *Intention `json:"Intention,omitnil,omitempty" name:"Intention"`

	// 进入签署流程的限制，目前支持以下选项：
	// <ul><li> <b>空值（默认）</b> :无限制，可在任何场景进入签署流程。</li><li> <b>link</b> :选择此选项后，将无法通过控制台或电子签小程序列表进入填写或签署操作，仅可预览合同。填写或签署流程只能通过短信或发起方提供的专用链接进行。</li></ul>
	SignEndpoints []*string `json:"SignEndpoints,omitnil,omitempty" name:"SignEndpoints"`
}

type FlowDetailInfo struct {
	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 合同流程的类别分类（如销售合同/入职合同等）。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 合同流程当前的签署状态, 会存在下列的状态值 <ul><li> **0** : 未开启流程(合同中不存在填写环节)</li> <li> **1** : 待签署</li> <li> **2** : 部分签署</li> <li> **3** : 已拒签</li> <li> **4** : 已签署</li> <li> **5** : 已过期</li> <li> **6** : 已撤销</li> <li> **7** : 未开启流程(合同中存在填写环节)</li> <li> **8** : 等待填写</li> <li> **9** : 部分填写</li> <li> **10** : 已拒填</li> <li> **21** : 已解除</li></ul>	
	FlowStatus *int64 `json:"FlowStatus,omitnil,omitempty" name:"FlowStatus"`

	// 当合同流程状态为已拒签（即 FlowStatus=3）或已撤销（即 FlowStatus=6）时，此字段 FlowMessage 为拒签或撤销原因。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMessage *string `json:"FlowMessage,omitnil,omitempty" name:"FlowMessage"`

	// 合同流程描述信息。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowDescription *string `json:"FlowDescription,omitnil,omitempty" name:"FlowDescription"`

	// 合同流程的创建时间戳，格式为Unix标准时间戳（秒）。	
	CreatedOn *int64 `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 合同流程的签署方数组
	FlowApproverInfos []*FlowApproverDetail `json:"FlowApproverInfos,omitnil,omitempty" name:"FlowApproverInfos"`

	// 合同流程的关注方信息数组
	CcInfos []*FlowApproverDetail `json:"CcInfos,omitnil,omitempty" name:"CcInfos"`

	// 合同流程发起方的员工编号, 即员工在腾讯电子签平台的唯一身份标识。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`
}

type FlowGroupApproverInfo struct {
	// 合同流程ID。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署节点ID，用于生成动态签署人链接完成领取。注：`生成动态签署人补充链接时必传。`
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`
}

type FlowGroupApprovers struct {
	// 合同流程ID 
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署方信息，包含合同ID和角色ID用于定位RecipientId。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Approvers []*ApproverItem `json:"Approvers,omitnil,omitempty" name:"Approvers"`
}

type FlowGroupInfo struct {
	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 签署流程参与者信息，最大限制50方
	// 注意 approver中的顺序需要和模板中的顺序保持一致， 否则会导致模板中配置的信息无效。
	Approvers []*ApproverInfo `json:"Approvers,omitnil,omitempty" name:"Approvers"`

	// 文件资源ID，通过多文件上传[UploadFiles](https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles)接口获得，为32位字符串。
	// 建议开发者保存此资源ID，后续创建合同或创建合同流程需此资源ID。
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 合同模板ID，为32位字符串。
	// 建议开发者保存此模板ID，后续用此模板发起合同流程需要此参数。
	// 可登录腾讯电子签控制台，在 "模板"->"模板中心"->"列表展示设置"选中模板 ID 中查看某个模板的TemplateId(在页面中展示为模板ID)。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 签署流程的类型(如销售合同/入职合同等)，最大长度200个字符
	// 示例值：劳务合同
	FlowType *string `json:"FlowType,omitnil,omitempty" name:"FlowType"`

	// 签署流程描述,最大长度1000个字符
	FlowDescription *string `json:"FlowDescription,omitnil,omitempty" name:"FlowDescription"`

	// 签署流程的签署截止时间。
	// 
	// 值为unix时间戳,精确到秒,不传默认为当前时间一年后
	// 示例值：1604912664
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 合同（流程）的回调地址
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 调用方自定义的个性化字段(可自定义此字段的值)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 发送类型：
	// true：无序签
	// false：有序签
	// 注：默认为false（有序签），请和模板中的配置保持一致
	// 示例值：true
	Unordered *bool `json:"Unordered,omitnil,omitempty" name:"Unordered"`

	// 模板或者合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li>单行文本控件</li>
	// <li>多行文本控件</li>
	// <li>勾选框控件</li>
	// <li>数字控件</li>
	// <li>图片控件</li>
	// <li>动态表格等填写控件</li></ul>
	Components []*Component `json:"Components,omitnil,omitempty" name:"Components"`

	// 发起方企业的签署人进行签署操作是否需要企业内部审批。使用此功能需要发起方企业有参与签署。
	// 若设置为true，审核结果需通过接口 [CreateFlowSignReview](https://qian.tencent.com/developers/companyApis/operateFlows/CreateFlowSignReview) 通知电子签，审核通过后，发起方企业签署人方可进行签署操作，否则会阻塞其签署操作。
	// 
	// 注：企业可以通过此功能与企业内部的审批流程进行关联，支持手动、静默签署合同。
	// 示例值：true
	NeedSignReview *bool `json:"NeedSignReview,omitnil,omitempty" name:"NeedSignReview"`

	// 个人自动签场景。发起自动签署时，需设置对应自动签署场景，目前仅支持场景：处方单-E_PRESCRIPTION_AUTO_SIGN
	// 示例值：E_PRESCRIPTION_AUTO_SIGN
	AutoSignScene *string `json:"AutoSignScene,omitnil,omitempty" name:"AutoSignScene"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下：  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li></ul>效果如下:![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"FlowDisplayType,omitnil,omitempty" name:"FlowDisplayType"`
}

type FlowGroupOptions struct {
	// 签署人校验方式,支持以下类型
	// <ul><li>VerifyCheck : 人脸识别 (默认值)</li>
	// <li>MobileCheck : 手机号验证</li></ul>
	// 参数说明：此参数仅在合同组文件发起有效，可选人脸识别或手机号验证两种方式，若选择后者，未实名个人签署方在签署合同时，无需经过实名认证和意愿确认两次人脸识别，该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil,omitempty" name:"ApproverVerifyType"`

	// 发起合同（流程）组本方企业经办人通知方式
	// 签署通知类型，支持以下类型
	// <ul><li>sms : 短信 (默认值)</li><li>none : 不通知</li></ul>
	SelfOrganizationApproverNotifyType *string `json:"SelfOrganizationApproverNotifyType,omitnil,omitempty" name:"SelfOrganizationApproverNotifyType"`

	// 发起合同（流程）组他方经办人通知方式
	// 签署通知类型，支持以下类型
	// <ul><li>sms : 短信 (默认值)</li><li>none : 不通知</li></ul>
	OtherApproverNotifyType *string `json:"OtherApproverNotifyType,omitnil,omitempty" name:"OtherApproverNotifyType"`
}

type FlowGroupUrlInfo struct {
	// 合同组子合同和签署方的信息，用于补充动态签署人。
	FlowGroupApproverInfos []*FlowGroupApproverInfo `json:"FlowGroupApproverInfos,omitnil,omitempty" name:"FlowGroupApproverInfos"`
}

type FormField struct {
	// 控件填充vaule，ComponentType和传入值类型对应关系：
	// <ul><li> <b>TEXT</b> : 文本内容</li>
	// <li> <b>MULTI_LINE_TEXT</b> : 文本内容， 可以用  \n 来控制换行位置</li>
	// <li> <b>CHECK_BOX</b> : true/false</li>
	// <li> <b>FILL_IMAGE、ATTACHMENT</b> : 附件的FileId，需要通过UploadFiles接口上传获取</li>
	// <li> <b>SELECTOR</b> : 选项值</li>
	// <li> <b>DYNAMIC_TABLE</b>  - 传入json格式的表格内容，详见说明：[数据表格](https://qian.tencent.com/developers/company/dynamic_table)</li>
	// <li> <b>DATE</b> : 格式化：xxxx年xx月xx日（例如：2024年05月28日）</li>
	// </ul>
	// 
	// 
	// <b>控件值约束说明</b>：
	// <table> <thead> <tr> <th>特殊控件</th> <th>填写约束</th> </tr> </thead> <tbody> <tr> <td>企业全称控件</td> <td>企业名称中文字符中文括号</td> </tr> <tr> <td>统一社会信用代码控件</td> <td>企业注册的统一社会信用代码</td> </tr> <tr> <td>法人名称控件</td> <td>最大50个字符，2到25个汉字或者1到50个字母</td> </tr> <tr> <td>签署意见控件</td> <td>签署意见最大长度为50字符</td> </tr> <tr> <td>签署人手机号控件</td> <td>国内手机号 13,14,15,16,17,18,19号段长度11位</td> </tr> <tr> <td>签署人身份证控件</td> <td>合法的身份证号码检查</td> </tr> <tr> <td>控件名称</td> <td>控件名称最大长度为20字符，不支持表情</td> </tr> <tr> <td>单行文本控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>多行文本控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>勾选框控件</td> <td>选择填字符串true，不选填字符串false</td> </tr> <tr> <td>选择器控件</td> <td>同单行文本控件约束，填写选择值中的字符串</td> </tr> <tr> <td>数字控件</td> <td>请输入有效的数字(可带小数点)</td> </tr> <tr> <td>日期控件</td> <td>格式：yyyy年mm月dd日</td> </tr> <tr> <td>附件控件</td> <td>JPG或PNG图片，上传数量限制，1到6个，最大6个附件，填写上传的资源ID</td> </tr> <tr> <td>图片控件</td> <td>JPG或PNG图片，填写上传的图片资源ID</td> </tr> <tr> <td>邮箱控件</td> <td>有效的邮箱地址, w3c标准</td> </tr> <tr> <td>地址控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>省市区控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>性别控件</td> <td>选择值中的字符串</td> </tr> <tr> <td>学历控件</td> <td>选择值中的字符串</td> </tr> </tbody> </table>
	ComponentValue *string `json:"ComponentValue,omitnil,omitempty" name:"ComponentValue"`

	// 控件id，和ComponentName选择一项传入即可
	// 
	// <a href="https://dyn.ess.tencent.cn/guide/apivideo/component_name.mp4" target="_blank">点击查看在模板中找到控件ID的方式</a>
	ComponentId *string `json:"ComponentId,omitnil,omitempty" name:"ComponentId"`

	// 控件名字，最大长度不超过30字符，和ComponentId选择一项传入即可
	// 
	// <a href="https://dyn.ess.tencent.cn/guide/apivideo/component_name.mp4" target="_blank">点击查看在模板中找到控件名字的方式</a>
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`
}

// Predefined struct for user
type GetTaskResultApiRequestParams struct {
	// 转换任务Id，通过接口<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>或<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateMergeFileTask" target="_blank">创建多文件转换任务接口</a>
	// 得到的转换任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
}

type GetTaskResultApiRequest struct {
	*tchttp.BaseRequest
	
	// 转换任务Id，通过接口<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>或<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateMergeFileTask" target="_blank">创建多文件转换任务接口</a>
	// 得到的转换任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil,omitempty" name:"Organization"`
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

	// 资源Id（即FileId），用于[用PDF文件创建签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowByFiles)
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type GroupOrganization struct {
	// 成员企业名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 成员企业别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 成员企业id，为 32 位字符串，可在电子签PC 控制台，企业设置->企业电子签账号 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 记录更新时间， unix时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 成员企业加入集团的当前状态
	// <ul><li> **1**：待授权</li>
	// <li> **2**：已授权待激活</li>
	// <li> **3**：拒绝授权</li>
	// <li> **4**：已解除</li>
	// <li> **5**：已加入</li>
	// </ul>
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否为集团主企业
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMainOrganization *bool `json:"IsMainOrganization,omitnil,omitempty" name:"IsMainOrganization"`

	// 企业社会信用代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`

	// 企业超管信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminInfo *Admin `json:"AdminInfo,omitnil,omitempty" name:"AdminInfo"`

	// 企业许可证Id，此字段暂时不需要关注
	// 注意：此字段可能返回 null，表示取不到有效值。
	License *string `json:"License,omitnil,omitempty" name:"License"`

	// 企业许可证过期时间，unix时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseExpireTime *uint64 `json:"LicenseExpireTime,omitnil,omitempty" name:"LicenseExpireTime"`

	// 成员企业加入集团时间，unix时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinTime *uint64 `json:"JoinTime,omitnil,omitempty" name:"JoinTime"`

	// 是否使用自建审批流引擎（即不是企微审批流引擎）
	// <ul><li> **true**：是</li>
	// <li> **false**：否</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowEngineEnable *bool `json:"FlowEngineEnable,omitnil,omitempty" name:"FlowEngineEnable"`
}

type HasAuthOrganization struct {
	// 授权企业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 授权企业名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 被授权企业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizedOrganizationId *string `json:"AuthorizedOrganizationId,omitnil,omitempty" name:"AuthorizedOrganizationId"`

	// 被授权企业名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizedOrganizationName *string `json:"AuthorizedOrganizationName,omitnil,omitempty" name:"AuthorizedOrganizationName"`

	// 授权模板id（仅当授权方式为模板授权时有值）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 授权模板名称（仅当授权方式为模板授权时有值）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 授权时间，格式为时间戳，单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizeTime *int64 `json:"AuthorizeTime,omitnil,omitempty" name:"AuthorizeTime"`
}

type HasAuthUser struct {
	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 当前员工的归属情况，可能值是：
	// MainOrg：在集团企业的场景下，返回此值代表是归属主企业
	// CurrentOrg：在普通企业场景下返回此值；或者在集团企业的场景下，返回此值代表归属子企业
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongTo *string `json:"BelongTo,omitnil,omitempty" name:"BelongTo"`

	// 集团主企业id，当前企业为集团子企业时，该字段有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainOrganizationId *string `json:"MainOrganizationId,omitnil,omitempty" name:"MainOrganizationId"`
}

type IntegrateRole struct {
	// 角色id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色状态，1-启用，2-禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleStatus *uint64 `json:"RoleStatus,omitnil,omitempty" name:"RoleStatus"`

	// 是否是集团角色，true-是，false-否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsGroupRole *bool `json:"IsGroupRole,omitnil,omitempty" name:"IsGroupRole"`

	// 管辖的子企业列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubOrgIdList []*string `json:"SubOrgIdList,omitnil,omitempty" name:"SubOrgIdList"`

	// 权限树
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil,omitempty" name:"PermissionGroups"`
}

type IntegrationDepartment struct {
	// 部门ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeptId *string `json:"DeptId,omitnil,omitempty" name:"DeptId"`

	// 部门名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeptName *string `json:"DeptName,omitnil,omitempty" name:"DeptName"`

	// 父部门ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentDeptId *string `json:"ParentDeptId,omitnil,omitempty" name:"ParentDeptId"`

	// 客户系统部门ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeptOpenId *string `json:"DeptOpenId,omitnil,omitempty" name:"DeptOpenId"`

	// 序列号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderNo *uint64 `json:"OrderNo,omitnil,omitempty" name:"OrderNo"`
}

type Intention struct {
	// 视频认证类型，支持以下类型
	// <ul><li>1 : 问答模式</li>
	// <li>2 : 点头模式</li></ul>
	// 
	// 注: `视频认证为白名单功能，使用前请联系对接的客户经理沟通。`
	IntentionType *int64 `json:"IntentionType,omitnil,omitempty" name:"IntentionType"`

	// 意愿核身语音问答模式（即语音播报+语音回答）使用的文案，包括：系统语音播报的文本、需要核验的标准文本。当前仅支持1轮问答。
	// 
	// 注：`选择问答模式时，此字段可不传，不传则使用默认语音文本：请问，您是否同意签署本协议？可语音回复“同意”或“不同意”。`
	IntentionQuestions []*IntentionQuestion `json:"IntentionQuestions,omitnil,omitempty" name:"IntentionQuestions"`

	// 意愿核身（点头确认模式）使用的文案，若未使用意愿核身（点头确认模式），则该字段无需传入。当前仅支持一个提示文本。
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*IntentionActionResultDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type IntentionActionResultDetail struct {
	// 视频base64编码（其中包含全程提示文本和点头音频，mp4格式）
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Video *string `json:"Video,omitnil,omitempty" name:"Video"`

	//  和答案匹配结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultCode []*string `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`

	// 回答问题语音识别结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrResult []*string `json:"AsrResult,omitnil,omitempty" name:"AsrResult"`
}

// Predefined struct for user
type ModifyApplicationCallbackInfoRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 操作类型：
	// 1-新增
	// 2-删除
	OperateType *int64 `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 企业应用回调信息
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitnil,omitempty" name:"CallbackInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type ModifyApplicationCallbackInfoRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 操作类型：
	// 1-新增
	// 2-删除
	OperateType *int64 `json:"OperateType,omitnil,omitempty" name:"OperateType"`

	// 企业应用回调信息
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitnil,omitempty" name:"CallbackInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *ModifyApplicationCallbackInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationCallbackInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "OperateType")
	delete(f, "CallbackInfo")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationCallbackInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationCallbackInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApplicationCallbackInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationCallbackInfoResponseParams `json:"Response"`
}

func (r *ModifyApplicationCallbackInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationCallbackInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyExtendedServiceRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 要管理的拓展服务类型。
	// <ul><li>OPEN_SERVER_SIGN：企业自动签署</li>
	// <li>AUTO_SIGN_CAN_FILL_IN：本企业自动签合同支持签前内容补充</li>
	// <li>OVERSEA_SIGN：企业与港澳台居民签署合同</li>
	// <li>AGE_LIMIT_EXPANSION：拓宽签署方年龄限制</li>
	// <li>MOBILE_CHECK_APPROVER：个人签署方仅校验手机号</li>
	// <li>HIDE_OPERATOR_DISPLAY：隐藏合同经办人姓名</li>
	// <li>ORGANIZATION_OCR_FALLBACK：正楷临摹签名失败后更换其他签名类型</li>
	// <li>ORGANIZATION_FLOW_NOTIFY_TYPE：短信通知签署方</li>
	// <li>HIDE_ONE_KEY_SIGN：个人签署方手动签字</li>
	// <li>ORGANIZATION_FLOW_EMAIL_NOTIFY：邮件通知签署方</li>
	// <li>FLOW_APPROVAL：合同审批强制开启</li>
	// <li>ORGANIZATION_FLOW_PASSWD_NOTIFY：签署密码开通引导</li></ul>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 操作类型
	// <ul>
	// <li>OPEN : 开通</li>
	// <li>CLOSE : 关闭</li>
	// </ul>
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

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
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 要管理的拓展服务类型。
	// <ul><li>OPEN_SERVER_SIGN：企业自动签署</li>
	// <li>AUTO_SIGN_CAN_FILL_IN：本企业自动签合同支持签前内容补充</li>
	// <li>OVERSEA_SIGN：企业与港澳台居民签署合同</li>
	// <li>AGE_LIMIT_EXPANSION：拓宽签署方年龄限制</li>
	// <li>MOBILE_CHECK_APPROVER：个人签署方仅校验手机号</li>
	// <li>HIDE_OPERATOR_DISPLAY：隐藏合同经办人姓名</li>
	// <li>ORGANIZATION_OCR_FALLBACK：正楷临摹签名失败后更换其他签名类型</li>
	// <li>ORGANIZATION_FLOW_NOTIFY_TYPE：短信通知签署方</li>
	// <li>HIDE_ONE_KEY_SIGN：个人签署方手动签字</li>
	// <li>ORGANIZATION_FLOW_EMAIL_NOTIFY：邮件通知签署方</li>
	// <li>FLOW_APPROVAL：合同审批强制开启</li>
	// <li>ORGANIZATION_FLOW_PASSWD_NOTIFY：签署密码开通引导</li></ul>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 操作类型
	// <ul>
	// <li>OPEN : 开通</li>
	// <li>CLOSE : 关闭</li>
	// </ul>
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

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
	delete(f, "Operator")
	delete(f, "ServiceType")
	delete(f, "Operate")
	delete(f, "Agent")
	delete(f, "Endpoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyExtendedServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyExtendedServiceResponseParams struct {
	// 操作跳转链接
	// <ul>
	// <li><strong>有效期：</strong> 跳转链接的有效期为24小时。</li>
	// <li><strong>无跳转链接返回的情况：</strong> 如果在操作过程中没有返回跳转链接，这意味着无需进行跳转操作。在这种情况下，服务将会直接被开通或关闭。
	// <li><strong>有跳转链接返回的情况：</strong> 当操作类型为“OPEN”（开通服务），并且扩展服务类型为以下之一时，  系统将返回一个操作链接。当前操作人（超级管理员或法人）需要点击此链接，以完成服务的开通操作。
	// 
	// <ul>
	// <li><strong>OPEN_SERVER_SIGN</strong>（企业自动签署）</li>
	// </ul></li></li>
	// </ul>
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
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// <ul><li>建议开发者妥善保存此流程ID，以便于顺利进行后续操作。</li>
	// <li>可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。</li></ul>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署流程或签署人新的签署截止时间，格式为Unix标准时间戳（秒）
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 签署方角色编号，为32位字符串
	// <ul><li>若指定了此参数，则只调整签署流程中此签署人的签署截止时间，否则调整合同整体的签署截止时间（合同截止时间+发起时未设置签署人截止时间的参与人的签署截止时间）</li>
	// <li>通过[用PDF文件创建签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowByFiles)发起合同，或通过[模板发起合同-创建电子文档](https://qian.tencent.com/developers/companyApis/startFlows/CreateDocument)时，返回参数[Approvers](https://qian.tencent.com/developers/companyApis/dataTypes/#approveritem)会返回此信息，建议开发者妥善保存</li>
	// <li>也可通过[查询合同流程的详情信息](https://qian.tencent.com/developers/companyApis/queryFlows/DescribeFlowInfo)接口查询签署人的RecipientId编号</li></ul>
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`
}

type ModifyFlowDeadlineRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// <ul><li>建议开发者妥善保存此流程ID，以便于顺利进行后续操作。</li>
	// <li>可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。</li></ul>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 签署流程或签署人新的签署截止时间，格式为Unix标准时间戳（秒）
	Deadline *int64 `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 代理相关应用信息，如集团主企业代子企业操作的场景中ProxyOrganizationId必填
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 签署方角色编号，为32位字符串
	// <ul><li>若指定了此参数，则只调整签署流程中此签署人的签署截止时间，否则调整合同整体的签署截止时间（合同截止时间+发起时未设置签署人截止时间的参与人的签署截止时间）</li>
	// <li>通过[用PDF文件创建签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowByFiles)发起合同，或通过[模板发起合同-创建电子文档](https://qian.tencent.com/developers/companyApis/startFlows/CreateDocument)时，返回参数[Approvers](https://qian.tencent.com/developers/companyApis/dataTypes/#approveritem)会返回此信息，建议开发者妥善保存</li>
	// <li>也可通过[查询合同流程的详情信息](https://qian.tencent.com/developers/companyApis/queryFlows/DescribeFlowInfo)接口查询签署人的RecipientId编号</li></ul>
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
	delete(f, "Operator")
	delete(f, "FlowId")
	delete(f, "Deadline")
	delete(f, "Agent")
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
type ModifyIntegrationDepartmentRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 电子签部门ID，通过<a href="https://qian.tencent.com/developers/companyApis/organizations/DescribeIntegrationDepartments" target="_blank">DescribeIntegrationDepartments</a>接口获得。
	DeptId *string `json:"DeptId,omitnil,omitempty" name:"DeptId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 电子签父部门ID，通过<a href="https://qian.tencent.com/developers/companyApis/organizations/DescribeIntegrationDepartments" target="_blank">DescribeIntegrationDepartments</a>接口获得。
	ParentDeptId *string `json:"ParentDeptId,omitnil,omitempty" name:"ParentDeptId"`

	// 部门名称，最大长度为50个字符。
	DeptName *string `json:"DeptName,omitnil,omitempty" name:"DeptName"`

	// 客户系统部门ID，最大长度为64个字符。
	DeptOpenId *string `json:"DeptOpenId,omitnil,omitempty" name:"DeptOpenId"`

	// 排序号，支持设置的数值范围为1~30000。同一父部门下，排序号越大，部门顺序越靠前。
	OrderNo *uint64 `json:"OrderNo,omitnil,omitempty" name:"OrderNo"`
}

type ModifyIntegrationDepartmentRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得组织架构管理权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 电子签部门ID，通过<a href="https://qian.tencent.com/developers/companyApis/organizations/DescribeIntegrationDepartments" target="_blank">DescribeIntegrationDepartments</a>接口获得。
	DeptId *string `json:"DeptId,omitnil,omitempty" name:"DeptId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 电子签父部门ID，通过<a href="https://qian.tencent.com/developers/companyApis/organizations/DescribeIntegrationDepartments" target="_blank">DescribeIntegrationDepartments</a>接口获得。
	ParentDeptId *string `json:"ParentDeptId,omitnil,omitempty" name:"ParentDeptId"`

	// 部门名称，最大长度为50个字符。
	DeptName *string `json:"DeptName,omitnil,omitempty" name:"DeptName"`

	// 客户系统部门ID，最大长度为64个字符。
	DeptOpenId *string `json:"DeptOpenId,omitnil,omitempty" name:"DeptOpenId"`

	// 排序号，支持设置的数值范围为1~30000。同一父部门下，排序号越大，部门顺序越靠前。
	OrderNo *uint64 `json:"OrderNo,omitnil,omitempty" name:"OrderNo"`
}

func (r *ModifyIntegrationDepartmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntegrationDepartmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "DeptId")
	delete(f, "Agent")
	delete(f, "ParentDeptId")
	delete(f, "DeptName")
	delete(f, "DeptOpenId")
	delete(f, "OrderNo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIntegrationDepartmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntegrationDepartmentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIntegrationDepartmentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIntegrationDepartmentResponseParams `json:"Response"`
}

func (r *ModifyIntegrationDepartmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntegrationDepartmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntegrationRoleRequestParams struct {
	// 角色Id，可通过接口 DescribeIntegrationRoles 查询获取
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 权限树
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil,omitempty" name:"PermissionGroups"`

	// 集团角色的话，需要传递集团子企业列表，如果是全选，则传1
	SubOrganizationIds []*string `json:"SubOrganizationIds,omitnil,omitempty" name:"SubOrganizationIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type ModifyIntegrationRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色Id，可通过接口 DescribeIntegrationRoles 查询获取
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 执行本接口操作的员工信息。使用此接口时，必须填写userId。
	// 支持填入集团子公司经办人 userId 代发合同。
	// 
	// 注: 在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 权限树
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil,omitempty" name:"PermissionGroups"`

	// 集团角色的话，需要传递集团子企业列表，如果是全选，则传1
	SubOrganizationIds []*string `json:"SubOrganizationIds,omitnil,omitempty" name:"SubOrganizationIds"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *ModifyIntegrationRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntegrationRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleId")
	delete(f, "Name")
	delete(f, "Operator")
	delete(f, "Description")
	delete(f, "PermissionGroups")
	delete(f, "SubOrganizationIds")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIntegrationRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntegrationRoleResponseParams struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIntegrationRoleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIntegrationRoleResponseParams `json:"Response"`
}

func (r *ModifyIntegrationRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntegrationRoleResponse) FromJsonString(s string) error {
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
	// <li>HONGKONG_MACAO_AND_TAIWAN 中国港澳台居民居住证(格式同居民身份证)</li>
	// <li>OTHER_CARD_TYPE 其他证件</li></ul>
	// 
	// 注: `其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	ApproverIdCardType *string `json:"ApproverIdCardType,omitnil,omitempty" name:"ApproverIdCardType"`

	// 签署方经办人的证件号码，应符合以下规则
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
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

	// 电子印章授权人的UserId
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 电子印章策略Id
	SealPolicyId *string `json:"SealPolicyId,omitnil,omitempty" name:"SealPolicyId"`

	// 印章状态，有以下六种：CHECKING（审核中）SUCCESS（已启用）FAIL（审核拒绝）CHECKING-SADM（待超管审核）DISABLE（已停用）STOPPED（已终止）
	SealStatus *string `json:"SealStatus,omitnil,omitempty" name:"SealStatus"`

	// 审核失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// 印章图片url，5分钟内有效
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 印章类型,OFFICIAL-企业公章, CONTRACT-合同专用章,ORGANIZATIONSEAL-企业印章(本地上传印章类型),LEGAL_PERSON_SEAL-法人印章
	SealType *string `json:"SealType,omitnil,omitempty" name:"SealType"`

	// 用印申请是否为永久授权，true-是，false-否
	IsAllTime *bool `json:"IsAllTime,omitnil,omitempty" name:"IsAllTime"`

	// 授权人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizedUsers []*AuthorizedUser `json:"AuthorizedUsers,omitnil,omitempty" name:"AuthorizedUsers"`

	// 印章扩展数据信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtendScene *ExtendScene `json:"ExtendScene,omitnil,omitempty" name:"ExtendScene"`
}

type OrgBillSummary struct {
	// 套餐总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 套餐使用数
	Used *int64 `json:"Used,omitnil,omitempty" name:"Used"`

	// 套餐剩余数
	Available *int64 `json:"Available,omitnil,omitempty" name:"Available"`

	// 套餐类型
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
}

type OrganizationAuthUrl struct {
	// 企业批量注册链接，根据Endpoint的不同设置，返回不同的链接地址。失效时间：7天
	// 跳转链接, 链接的有效期根据企业,员工状态和终端等有区别, 可以参考下表
	// <table> <thead> <tr> <th>Endpoint</th> <th>示例</th> <th>链接有效期限</th> </tr> </thead>  <tbody>
	//  <tr> <td>PC</td> <td>https://qian.tencent.com/console/batch-register?token=yDSx0UUgtjuaf4UEfd2MjCnfI1iuXFE6&orgName=批量认证线上测试企业9</td> <td>7天</td> </tr> 
	// <tr> <td>PC_SHORT_URL</td> <td>https://test.essurl.cn/8gDKUBAWK8</td> <td>7天</td> </tr> 
	// <tr> <td>APP</td> <td>/pages/guide/index?to=REGISTER_ENTERPRISE_FOR_BATCH&urlAuthToken=yDSx0UUgtjuaf4UEfd2MjCnfI1iuXFE6&orgName=批量认证线上测试企业9</td> <td>7天</td> </tr> </tbody> </table>
	// 注： 
	// `1.创建的链接应避免被转义，如：&被转义为\u0026；如使用Postman请求后，请选择响应类型为 JSON，否则链接将被转义`
	AuthUrl *string `json:"AuthUrl,omitnil,omitempty" name:"AuthUrl"`

	// 企业批量注册的错误信息，例如：企业三要素不通过	
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 企业批量注册的唯一 Id， 此 Id 可以用在[创建企业批量认证链接-单链接](https://qian.tencent.com/developers/companyApis/organizations/CreateBatchOrganizationAuthorizationUrl)。
	SubTaskId *string `json:"SubTaskId,omitnil,omitempty" name:"SubTaskId"`
}

type OrganizationCommonInfo struct {
	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 组织机构企业统一社会信用代码。
	// 请确认该企业统一社会信用代码与企业营业执照中注册的统一社会信用代码一致。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil,omitempty" name:"UniformSocialCreditCode"`

	// 组织机构法人的姓名。
	// 请确认该企业统一社会信用代码与企业营业执照中注册的法人姓名一致。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LegalName *string `json:"LegalName,omitnil,omitempty" name:"LegalName"`

	// 组织机构法人的证件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LegalIdCardType *string `json:"LegalIdCardType,omitnil,omitempty" name:"LegalIdCardType"`

	// 组织机构法人的证件号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	LegalIdCardNumber *string `json:"LegalIdCardNumber,omitnil,omitempty" name:"LegalIdCardNumber"`

	// 组织机构超管姓名。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 组织机构超管手机号。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminMobile *string `json:"AdminMobile,omitnil,omitempty" name:"AdminMobile"`

	// 组织机构超管证件类型
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminIdCardType *string `json:"AdminIdCardType,omitnil,omitempty" name:"AdminIdCardType"`

	// 组织机构超管证件号码
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminIdCardNumber *string `json:"AdminIdCardNumber,omitnil,omitempty" name:"AdminIdCardNumber"`

	// 原超管姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldAdminName *string `json:"OldAdminName,omitnil,omitempty" name:"OldAdminName"`

	// 原超管手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldAdminMobile *string `json:"OldAdminMobile,omitnil,omitempty" name:"OldAdminMobile"`

	// 原超管证件类型
	OldAdminIdCardType *string `json:"OldAdminIdCardType,omitnil,omitempty" name:"OldAdminIdCardType"`

	// 原超管证件号码
	OldAdminIdCardNumber *string `json:"OldAdminIdCardNumber,omitnil,omitempty" name:"OldAdminIdCardNumber"`
}

type OrganizationInfo struct {
	// 机构在平台的编号，内部字段，暂未开放
	//
	// Deprecated: OrganizationId is deprecated.
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 用户渠道，内部字段，暂未开放
	//
	// Deprecated: Channel is deprecated.
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// 用户在渠道的机构编号，内部字段，暂未开放
	//
	// Deprecated: OrganizationOpenId is deprecated.
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil,omitempty" name:"OrganizationOpenId"`

	// 用户真实的IP，内部字段，暂未开放
	//
	// Deprecated: ClientIp is deprecated.
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 机构的代理IP，内部字段，暂未开放
	//
	// Deprecated: ProxyIp is deprecated.
	ProxyIp *string `json:"ProxyIp,omitnil,omitempty" name:"ProxyIp"`
}

type PdfVerifyResult struct {
	// 验签结果。0-签名域未签名；1-验签成功； 3-验签失败；4-未找到签名域：文件内没有签名域；5-签名值格式不正确。
	VerifyResult *int64 `json:"VerifyResult,omitnil,omitempty" name:"VerifyResult"`

	// 签署平台
	// 如果文件是在腾讯电子签平台签署，则为**腾讯电子签**，
	// 如果文件不在腾讯电子签平台签署，则为**其他平台**。
	SignPlatform *string `json:"SignPlatform,omitnil,omitempty" name:"SignPlatform"`

	// 申请证书的主体的名字
	// 
	// 如果是在腾讯电子签平台签署, 则对应的主体的名字个数如下
	// **企业**:  ESS@企业名称@编码
	// **个人**: ESS@个人姓名@证件号@808854
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 权限key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 权限类型 1前端，2后端
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否隐藏
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hide *int64 `json:"Hide,omitnil,omitempty" name:"Hide"`

	// 数据权限标签 1:表示根节点，2:表示叶子结点
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataLabel *int64 `json:"DataLabel,omitnil,omitempty" name:"DataLabel"`

	// 数据权限独有，1:关联其他模块鉴权，2:表示关联自己模块鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// 数据权限独有，表示数据范围，1：全公司，2:部门及下级部门，3:自己
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataRange *int64 `json:"DataRange,omitnil,omitempty" name:"DataRange"`

	// 关联权限, 表示这个功能权限要受哪个数据权限管控
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataTo *string `json:"DataTo,omitnil,omitempty" name:"DataTo"`

	// 父级权限key
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentKey *string `json:"ParentKey,omitnil,omitempty" name:"ParentKey"`

	// 是否选中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsChecked *bool `json:"IsChecked,omitnil,omitempty" name:"IsChecked"`

	// 子权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*Permission `json:"Children,omitnil,omitempty" name:"Children"`
}

type PermissionGroup struct {
	// 权限组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 权限组key
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupKey *string `json:"GroupKey,omitnil,omitempty" name:"GroupKey"`

	// 是否隐藏分组，0否1是
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hide *int64 `json:"Hide,omitnil,omitempty" name:"Hide"`

	// 权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions []*Permission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

type Recipient struct {
	// 签署参与者ID，唯一标识
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 参与者类型。
	// 默认为空。
	// ENTERPRISE-企业；
	// INDIVIDUAL-个人；
	// PROMOTER-发起方
	RecipientType *string `json:"RecipientType,omitnil,omitempty" name:"RecipientType"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 是否需要验证，
	// 默认为false-不需要验证
	RequireValidation *bool `json:"RequireValidation,omitnil,omitempty" name:"RequireValidation"`

	// 是否需要签署，
	// 默认为true-需要签署
	RequireSign *bool `json:"RequireSign,omitnil,omitempty" name:"RequireSign"`

	// 此参与方添加的顺序，从0～N
	RoutingOrder *int64 `json:"RoutingOrder,omitnil,omitempty" name:"RoutingOrder"`

	// 是否需要发送，
	// 默认为true-需要发送
	RequireDelivery *bool `json:"RequireDelivery,omitnil,omitempty" name:"RequireDelivery"`

	// 邮箱地址
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 电话号码
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 关联的用户ID，电子签系统的用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 发送方式，默认为EMAIL。
	// EMAIL-邮件；
	// MOBILE-手机短信；
	// WECHAT-微信通知
	DeliveryMethod *string `json:"DeliveryMethod,omitnil,omitempty" name:"DeliveryMethod"`

	// 参与方的一些附属信息，json格式
	RecipientExtra *string `json:"RecipientExtra,omitnil,omitempty" name:"RecipientExtra"`

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
	// 签署方经办人在合同流程中的参与方ID，与控件绑定，是控件的归属方
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"RecipientId,omitnil,omitempty" name:"RecipientId"`

	// 参与方填写状态
	// <ul>
	// <li>**空值** : 此参与方没有填写控件</li>
	// <li>**0**:  未填写, 表示此参与方还没有填写合同的填写控件</li>
	// <li>**1**:  已填写, 表示此参与方已经填写所有的填写控件</li></ul>
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientFillStatus *string `json:"RecipientFillStatus,omitnil,omitempty" name:"RecipientFillStatus"`

	// 是否为发起方
	// <ul><li>true-发起方</li>
	// <li>false-参与方</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPromoter *bool `json:"IsPromoter,omitnil,omitempty" name:"IsPromoter"`

	// 改参与方填写控件信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components []*FilledComponent `json:"Components,omitnil,omitempty" name:"Components"`
}

type RegisterInfo struct {
	// 法人姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	LegalName *string `json:"LegalName,omitnil,omitempty" name:"LegalName"`

	// 社会统一信用代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Uscc is deprecated.
	Uscc *string `json:"Uscc,omitnil,omitempty" name:"Uscc"`

	// 社会统一信用代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnifiedSocialCreditCode *string `json:"UnifiedSocialCreditCode,omitnil,omitempty" name:"UnifiedSocialCreditCode"`
}

type RegistrationOrganizationInfo struct {
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

	// 组织机构企业注册地址。
	// 请确认该企业注册地址与企业营业执照中注册的地址一致。
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 组织机构超管姓名。
	// 在注册流程中，必须是超管本人进行操作。
	// 如果法人做为超管管理组织机构,超管姓名就是法人姓名
	// 如果入参中传递超管授权书PowerOfAttorneys，则此参数为必填参数。
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// 组织机构超管手机号。
	// 在注册流程中，这个手机号必须跟操作人在电子签注册的个人手机号一致。
	// 如果入参中传递超管授权书PowerOfAttorneys，则此参数为必填参数
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

	// 认证人身份证号，如果入参中传递超管授权书PowerOfAttorneys，则此参数为必填参数
	AdminIdCardNumber *string `json:"AdminIdCardNumber,omitnil,omitempty" name:"AdminIdCardNumber"`

	// 认证人证件类型 
	// 支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 中国港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 中国港澳台居民居住证(格式同中国大陆居民身份证)</li></ul>
	AdminIdCardType *string `json:"AdminIdCardType,omitnil,omitempty" name:"AdminIdCardType"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitnil,omitempty" name:"BusinessLicense"`

	// 授权书(PNG或JPG或PDF) base64格式, 大小不超过8M 。
	// p.s. 如果上传授权书 ，需遵循以下条件
	// 1. 超管的信息（超管姓名，超管手机号）必须为必填参数。
	// 2. 超管的个人身份必须在电子签已经实名。
	// 2. 认证方式AuthorizationTypes必须只能是上传授权书方式 
	PowerOfAttorneys []*string `json:"PowerOfAttorneys,omitnil,omitempty" name:"PowerOfAttorneys"`
}

type ReleasedApprover struct {
	// 签署人姓名，最大长度50个字。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 签署人手机号。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 要更换的原合同参与人RecipientId编号。(可通过接口<a href="https://qian.tencent.com/developers/companyApis/queryFlows/DescribeFlowInfo/">DescribeFlowInfo</a>查询签署人的RecipientId编号)<br/>
	RelievedApproverReceiptId *string `json:"RelievedApproverReceiptId,omitnil,omitempty" name:"RelievedApproverReceiptId"`

	// 指定签署人类型，目前仅支持
	// <ul><li> **ORGANIZATION**：企业（默认值）</li>
	// <li> **ENTERPRISESERVER**：企业静默签</li></ul>
	ApproverType *string `json:"ApproverType,omitnil,omitempty" name:"ApproverType"`

	// 签署控件类型，支持自定义企业签署方的签署控件类型
	// <ul><li> **SIGN_SEAL**：默认为印章控件类型（默认值）</li>
	// <li> **SIGN_SIGNATURE**：手写签名控件类型</li></ul>
	ApproverSignComponentType *string `json:"ApproverSignComponentType,omitnil,omitempty" name:"ApproverSignComponentType"`

	// 参与方在合同中的角色是按照创建合同的时候来排序的，解除协议默认会将第一个参与人叫`甲方`,第二个叫`乙方`,  第三个叫`丙方`，以此类推。
	// 
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
	// 合同流程是否可以催办：
	// true - 可以，false - 不可以。
	// 若无法催办，将返回RemindMessage以解释原因。
	CanRemind *bool `json:"CanRemind,omitnil,omitempty" name:"CanRemind"`

	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 在合同流程无法催办的情况下，系统将返回RemindMessage以阐述原因。
	RemindMessage *string `json:"RemindMessage,omitnil,omitempty" name:"RemindMessage"`
}

// Predefined struct for user
type RenewAutoSignLicenseRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 需要续期自动签的个人的信息，如姓名，证件信息等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type RenewAutoSignLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil,omitempty" name:"SceneKey"`

	// 需要续期自动签的个人的信息，如姓名，证件信息等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *RenewAutoSignLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewAutoSignLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "SceneKey")
	delete(f, "UserInfo")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewAutoSignLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewAutoSignLicenseResponseParams struct {
	// 续期成功后新的自动签许可到期时间。当且仅当已通过许可开通自动签时有值。
	// 
	// 值为unix时间戳,单位为秒。
	LicenseTo *int64 `json:"LicenseTo,omitnil,omitempty" name:"LicenseTo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewAutoSignLicenseResponse struct {
	*tchttp.BaseResponse
	Response *RenewAutoSignLicenseResponseParams `json:"Response"`
}

func (r *RenewAutoSignLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewAutoSignLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReviewerInfo struct {
	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`
}

type SealInfo struct {
	// 印章ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SealId *string `json:"SealId,omitnil,omitempty" name:"SealId"`

	// 印章类型。LEGAL_PERSON_SEAL: 法定代表人章；
	// ORGANIZATIONSEAL：企业印章；
	// OFFICIAL：企业公章；
	// CONTRACT：合同专用章
	// 注意：此字段可能返回 null，表示取不到有效值。
	SealType *string `json:"SealType,omitnil,omitempty" name:"SealType"`

	// 印章名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SealName *string `json:"SealName,omitnil,omitempty" name:"SealName"`
}

type SignQrCode struct {
	// 二维码ID，为32位字符串。
	QrCodeId *string `json:"QrCodeId,omitnil,omitempty" name:"QrCodeId"`

	// 二维码URL，可通过转换二维码的工具或代码组件将此URL转化为二维码，以便用户扫描进行流程签署。
	QrCodeUrl *string `json:"QrCodeUrl,omitnil,omitempty" name:"QrCodeUrl"`

	// 二维码的有截止时间，格式为Unix标准时间戳（秒）。
	// 一旦超过二维码的有效期限，该二维码将自动失效。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 微信小程序二维码
	WeixinQrCodeUrl *string `json:"WeixinQrCodeUrl,omitnil,omitempty" name:"WeixinQrCodeUrl"`
}

type SignUrl struct {
	// 跳转至电子签名小程序签署的链接地址。
	// 适用于客户端APP及小程序直接唤起电子签名小程序。
	AppSignUrl *string `json:"AppSignUrl,omitnil,omitempty" name:"AppSignUrl"`

	// 签署链接有效时间，格式类似"2022-08-05 15:55:01"
	EffectiveTime *string `json:"EffectiveTime,omitnil,omitempty" name:"EffectiveTime"`

	// 跳转至电子签名小程序签署的链接地址，格式类似于https://essurl.cn/xxx。
	// 打开此链接将会展示H5中间页面，随后唤起电子签名小程序以进行合同签署。
	HttpSignUrl *string `json:"HttpSignUrl,omitnil,omitempty" name:"HttpSignUrl"`
}

type Staff struct {
	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 注：`创建和更新场景无需填写。`
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 显示的用户名/昵称。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 用户邮箱。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户在第三方平台ID。
	// 注：`如需在此接口提醒员工实名，该参数不传。`
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 员工角色信息。
	// 注：`创建和更新场景无需填写。`
	// 注意：此字段可能返回 null，表示取不到有效值。
	Roles []*StaffRole `json:"Roles,omitnil,omitempty" name:"Roles"`

	// 员工部门信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitnil,omitempty" name:"Department"`

	// 员工是否实名。
	// 注：`创建和更新场景无需填写。`
	Verified *bool `json:"Verified,omitnil,omitempty" name:"Verified"`

	// 员工创建时间戳，单位秒。
	// 注：`创建和更新场景无需填写。`
	CreatedOn *int64 `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 员工实名时间戳，单位秒。
	// 注：`创建和更新场景无需填写。`
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifiedOn *int64 `json:"VerifiedOn,omitnil,omitempty" name:"VerifiedOn"`

	// 员工是否离职：
	// <ul><li>**0**：未离职</li><li>**1**：离职</li></ul>
	// 注：`创建和更新场景无需填写。`
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuiteJob *int64 `json:"QuiteJob,omitnil,omitempty" name:"QuiteJob"`

	// 员工离职交接人用户ID。
	// 注：`创建和更新场景无需填写。`
	ReceiveUserId *string `json:"ReceiveUserId,omitnil,omitempty" name:"ReceiveUserId"`

	// 员工离职交接人用户OpenId。
	// 注：`创建和更新场景无需填写。`
	ReceiveOpenId *string `json:"ReceiveOpenId,omitnil,omitempty" name:"ReceiveOpenId"`

	// 企业微信用户账号ID。
	// 注：`仅企微类型的企业创建员工接口支持该字段。`
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeworkOpenId *string `json:"WeworkOpenId,omitnil,omitempty" name:"WeworkOpenId"`
}

type StaffRole struct {
	// 角色ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

// Predefined struct for user
type StartFlowRequestParams struct {
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 此处需要传入[创建签署流程接口](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow)得到的FlowId。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 客户端Token，保持接口幂等性,最大长度64个字符
	//
	// Deprecated: ClientToken is deprecated.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 若在创建签署流程时指定了关注人CcInfos，此参数可设定向关注人发送短信通知的类型：
	// <ul><li> **0** :合同发起时通知通知对方来查看合同（默认）</li>
	// <li> **1** : 签署完成后通知对方来查看合同</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil,omitempty" name:"CcNotifyType"`
}

type StartFlowRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 合同流程ID，为32位字符串。
	// 此处需要传入[创建签署流程接口](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow)得到的FlowId。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 客户端Token，保持接口幂等性,最大长度64个字符
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 若在创建签署流程时指定了关注人CcInfos，此参数可设定向关注人发送短信通知的类型：
	// <ul><li> **0** :合同发起时通知通知对方来查看合同（默认）</li>
	// <li> **1** : 签署完成后通知对方来查看合同</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil,omitempty" name:"CcNotifyType"`
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
	delete(f, "CcNotifyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartFlowResponseParams struct {
	// 发起成功后返回的状态，根据合同流程的不同，返回不同状态：
	// <ul><li> **START** : 发起成功, 合同进入签署环节</li>
	// <li> **REVIEW** : 提交审核成功, 合同需要发起审核, 发起方企业通过接口审核通过后合同才进入签署环境  `白名单功能，使用前请联系对接的客户经理沟通。`</li>
	// <li> **EXECUTING** : 已提交发起任务且PDF合同正在合成中, 等PDF合同合成成功后进入签署环节</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type SubOrgBillSummary struct {
	// 子企业名称
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	//  
	Usage []*SubOrgBillUsage `json:"Usage,omitnil,omitempty" name:"Usage"`
}

type SubOrgBillUsage struct {
	// 套餐使用数
	Used *int64 `json:"Used,omitnil,omitempty" name:"Used"`

	// 套餐类型
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
}

type SuccessCreateStaffData struct {
	// 员工名
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 员工手机号
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 员工在电子签平台的id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 提示，当创建已存在未实名用户时，该字段有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 传入的企微账号id
	WeworkOpenId *string `json:"WeworkOpenId,omitnil,omitempty" name:"WeworkOpenId"`

	// 员工邀请返回链接 根据入参的 InvitationNotifyType 和 Endpoint 返回链接 <table><tbody><tr><td>链接类型</td><td>有效期</td><td>示例</td></tr><tr><td>HTTP_SHORT_URL（短链）</td><td>一天</td><td>https://test.essurl.cn/fvG7UBEd0F</td></tr><tr><td>HTTP（长链）</td><td>一天</td><td>https://res.ess.tencent.cn/cdn/h5-activity-dev/jump-mp.html?where=mini&from=MSG&to=USER_VERIFY&verifyToken=yDCVbUUckpwocmfpUySko7IS83LTV0u0&expireTime=1710840183</td></tr><tr><td>H5</td><td>30 天</td><td>https://quick.test.qian.tencent.cn/guide?Code=yDCVbUUckpwtvxqoUbTw4VBBjLbfAtW7&CodeType=QUICK&shortKey=yDCVbUY7lhqV7mZlCL2d</td></tr><tr><td>APP</td><td>一天</td><td>/pages/guide/index?to=USER_VERIFY&verifyToken=yDCVbUUckpwocm96UySko7ISvEIZH7Yz&expireTime=1710840455 </td></tr></tbody></table>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type SuccessDeleteStaffData struct {
	// 员工名
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 员工手机号
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 员工在电子签平台的id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type SuccessUpdateStaffData struct {
	// 传入的用户名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 传入的手机号，没有打码
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查看某位员工的UserId(在页面中展示为用户ID)。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// H5端员工实名链接
	// 
	// 只有入参 InvitationNotifyType = H5的时候才会进行返回。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type TemplateInfo struct {
	// 模板ID，模板的唯一标识
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板的名字
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 此模块需要签署的各个参与方的角色列表。RecipientId标识每个参与方角色对应的唯一标识符，用于确定此角色的信息。
	// 
	// [点击查看在模板中配置的签署参与方角色列表的样子](https://qcloudimg.tencent-cloud.cn/raw/e082bbcc0d923f8cb723d98382410aa2.png)
	// 
	Recipients []*Recipient `json:"Recipients,omitnil,omitempty" name:"Recipients"`

	// 模板的填充控件列表
	// 
	// [点击查看在模板中配置的填充控件的样子](https://qcloudimg.tencent-cloud.cn/raw/cb2f58529fca8d909258f9d45a56f7f4.png)
	Components []*Component `json:"Components,omitnil,omitempty" name:"Components"`

	// 此模板中的签署控件列表
	// 
	// [点击查看在模板中配置的签署控件的样子](https://qcloudimg.tencent-cloud.cn/raw/29bc6ed753a5a0fce4a3ab02e2c0d955.png)
	SignComponents []*Component `json:"SignComponents,omitnil,omitempty" name:"SignComponents"`

	// 模板描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 此模板的资源ID
	DocumentResourceIds []*string `json:"DocumentResourceIds,omitnil,omitempty" name:"DocumentResourceIds"`

	// 生成模板的文件基础信息
	FileInfos []*FileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 此模板里边附件的资源ID
	AttachmentResourceIds []*string `json:"AttachmentResourceIds,omitnil,omitempty" name:"AttachmentResourceIds"`

	// 签署人参与签署的顺序，可以分为以下两种方式：
	// 
	// <b>无序</b>：不限定签署人的签署顺序，签署人可以在任何时间签署。此种方式值为 ：｛-1｝
	// <b>有序</b>：通过序列数字标识签署顺序，从0开始编码，数字越大签署顺序越靠后，签署人按照指定的顺序依次签署。此种方式值为： ｛0，1，2，3………｝
	SignOrder []*int64 `json:"SignOrder,omitnil,omitempty" name:"SignOrder"`

	// 此模板的状态可以分为以下几种：
	// 
	// <b>-1</b>：不可用状态。
	// <b>0</b>：草稿态，即模板正在编辑或未发布状态。
	// <b>1</b>：正式态，只有正式态的模板才可以发起合同。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 模板的创建者信息，用户的名字
	// 
	// 注： `是创建者的名字，而非创建者的用户ID`
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 模板创建的时间戳，格式为Unix标准时间戳（秒）
	CreatedOn *int64 `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 此模板创建方角色信息。
	// 
	// [点击查看在模板中配置的创建方角色的样子](https://qcloudimg.tencent-cloud.cn/raw/e082bbcc0d923f8cb723d98382410aa2.png)
	Promoter *Recipient `json:"Promoter,omitnil,omitempty" name:"Promoter"`

	// 模板类型可以分为以下两种：
	// 
	// <b>1</b>：带有本企业自动签署的模板，即签署过程无需签署人手动操作，系统自动完成签署。
	// <b>3</b>：普通模板，即签署人需要手动进行签署操作。
	TemplateType *int64 `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// 模板可用状态可以分为以下两种：
	// 
	// <b>1</b>：（默认）启用状态，即模板可以正常使用。
	// <b>2</b>：停用状态，即模板暂时无法使用。
	// 
	// 可到控制台启停模板
	Available *int64 `json:"Available,omitnil,omitempty" name:"Available"`

	// 创建模板的企业ID，电子签的机构ID
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 模板创建人用户ID
	CreatorId *string `json:"CreatorId,omitnil,omitempty" name:"CreatorId"`

	// 模板的 H5 预览链接，有效期为 5 分钟。
	// 您可以通过浏览器直接打开此链接预览模板，或将其嵌入到 iframe 中进行预览。
	// 
	// 注意：只有在请求接口时将 <b>WithPreviewUrl </b>参数设置为 true，才会生成预览链接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`

	// 用户自定义合同类型。
	// 
	// 返回配置模板的时候选择的合同类型。[点击查看配置的位置](https://qcloudimg.tencent-cloud.cn/raw/4a766f0540253bf2a05d50c58bd14990.png)
	// 
	// 自定义合同类型配置的地方如链接图所示。[点击查看自定义合同类型管理的位置](https://qcloudimg.tencent-cloud.cn/raw/36582cea03ae6a2559894844942b5d5c.png)
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserFlowType *UserFlowType `json:"UserFlowType,omitnil,omitempty" name:"UserFlowType"`

	// 模板版本的编号，旨在标识其独特的版本信息，通常呈现为一串字符串，由日期和递增的数字组成
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateVersion *string `json:"TemplateVersion,omitnil,omitempty" name:"TemplateVersion"`

	// 模板是否已发布可以分为以下两种状态：
	// 
	// <b>true</b>：已发布状态，表示该模板已经发布并可以正常使用。
	// <b>false</b>：未发布状态，表示该模板还未发布，无法使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Published *bool `json:"Published,omitnil,omitempty" name:"Published"`

	// <b>集体账号场景下</b>： 集团账号分享给子企业的模板的来源模板ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShareTemplateId *string `json:"ShareTemplateId,omitnil,omitempty" name:"ShareTemplateId"`

	// 此模板配置的预填印章列表（包括自动签署指定的印章）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateSeals []*SealInfo `json:"TemplateSeals,omitnil,omitempty" name:"TemplateSeals"`

	// 模板内部指定的印章列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Seals is deprecated.
	Seals []*SealInfo `json:"Seals,omitnil,omitempty" name:"Seals"`
}

// Predefined struct for user
type UnbindEmployeeUserIdWithClientOpenIdRequestParams struct {
	// 执行本接口操作的员工信息。使用此接口时，必须填写UserId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 
	// 通过<a href="https://qian.tencent.com/developers/companyApis/staffs/DescribeIntegrationEmployees" target="_blank">DescribeIntegrationEmployees</a>接口获取，也可登录腾讯电子签控制台查看
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/97cfffabb0caa61df16999cd395d7850.png)
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 员工在贵司业务系统中的唯一身份标识，用于与腾讯电子签账号进行映射，确保在同一企业内不会出现重复。
	// 该标识最大长度为64位字符串，仅支持包含26个英文字母和数字0-9的字符。
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type UnbindEmployeeUserIdWithClientOpenIdRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息。使用此接口时，必须填写UserId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 
	// 通过<a href="https://qian.tencent.com/developers/companyApis/staffs/DescribeIntegrationEmployees" target="_blank">DescribeIntegrationEmployees</a>接口获取，也可登录腾讯电子签控制台查看
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/97cfffabb0caa61df16999cd395d7850.png)
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 员工在贵司业务系统中的唯一身份标识，用于与腾讯电子签账号进行映射，确保在同一企业内不会出现重复。
	// 该标识最大长度为64位字符串，仅支持包含26个英文字母和数字0-9的字符。
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *UnbindEmployeeUserIdWithClientOpenIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindEmployeeUserIdWithClientOpenIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "UserId")
	delete(f, "OpenId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindEmployeeUserIdWithClientOpenIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindEmployeeUserIdWithClientOpenIdResponseParams struct {
	// 解绑是否成功。
	// <ul><li> **0**：失败 </li>
	// <li> **1**：成功 </li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindEmployeeUserIdWithClientOpenIdResponse struct {
	*tchttp.BaseResponse
	Response *UnbindEmployeeUserIdWithClientOpenIdResponseParams `json:"Response"`
}

func (r *UnbindEmployeeUserIdWithClientOpenIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindEmployeeUserIdWithClientOpenIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIntegrationEmployeesRequestParams struct {
	// 执行本接口操作的员工信息,UserId必填。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 需要更新的员工信息，最多不超过100个。根据UserId或OpenId更新员工信息，必须填写其中一个，优先使用UserId。
	// 
	// 可更新以下字段，其他字段暂不支持
	// <table> <thead> <tr> <th>参数</th> <th>含义</th> </tr> </thead> <tbody> <tr> <td>DisplayName</td> <td>用户的真实名字</td> </tr> <tr> <td>Mobile</td> <td>用户手机号码</td> </tr> <tr> <td>Email</td> <td>用户的邮箱</td> </tr> <tr> <td>Department.DepartmentId</td> <td>用户进入后的部门ID</td> </tr> </tbody> </table>
	Employees []*Staff `json:"Employees,omitnil,omitempty" name:"Employees"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 员工邀请方式可通过以下途径进行设置：<ul><li>**SMS（默认）**：邀请将通过短信或企业微信消息发送。若场景非企业微信，则采用企业微信消息；其他情境下则使用短信通知。短信内含链接，点击后将进入微信小程序进行认证并加入企业的流程。</li><li>**H5**：将生成H5链接，用户点击链接后可进入H5页面进行认证并加入企业的流程。</li><li>**NONE**：系统会根据Endpoint生成签署链接，业务方需获取链接并通知客户。</li></ul>
	InvitationNotifyType *string `json:"InvitationNotifyType,omitnil,omitempty" name:"InvitationNotifyType"`

	// 回跳地址，为认证成功后页面进行回跳的URL，请确保回跳地址的可用性。注：`只有在员工邀请方式（InvitationNotifyType参数）为H5场景下才生效， 其他方式下设置无效。`
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 要跳转的链接类型<ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li><li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li><li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li><li>**H5**： 第三方移动端浏览器进行嵌入，不支持小程序嵌入，过期时间一个月</li></ul>注意：InvitationNotifyType 和 Endpoint 的关系图<table><tbody><tr><td>通知类型（InvitationNotifyType）</td><td>Endpoint</td></tr><tr><td>SMS（默认）</td><td>不需要传递，会将 Endpoint 默认设置为HTTP_SHORT_URL</td></tr><tr><td>H5</td><td>不需要传递，会将 Endpoint 默认设置为 H5</td></tr><tr><td>NONE</td><td>所有 Endpoint 都支持（HTTP_URL/HTTP_SHORT_URL/H5/APP）默认为HTTP_SHORT_URL</td></tr></tbody></table>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

type UpdateIntegrationEmployeesRequest struct {
	*tchttp.BaseRequest
	
	// 执行本接口操作的员工信息,UserId必填。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 需要更新的员工信息，最多不超过100个。根据UserId或OpenId更新员工信息，必须填写其中一个，优先使用UserId。
	// 
	// 可更新以下字段，其他字段暂不支持
	// <table> <thead> <tr> <th>参数</th> <th>含义</th> </tr> </thead> <tbody> <tr> <td>DisplayName</td> <td>用户的真实名字</td> </tr> <tr> <td>Mobile</td> <td>用户手机号码</td> </tr> <tr> <td>Email</td> <td>用户的邮箱</td> </tr> <tr> <td>Department.DepartmentId</td> <td>用户进入后的部门ID</td> </tr> </tbody> </table>
	Employees []*Staff `json:"Employees,omitnil,omitempty" name:"Employees"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 员工邀请方式可通过以下途径进行设置：<ul><li>**SMS（默认）**：邀请将通过短信或企业微信消息发送。若场景非企业微信，则采用企业微信消息；其他情境下则使用短信通知。短信内含链接，点击后将进入微信小程序进行认证并加入企业的流程。</li><li>**H5**：将生成H5链接，用户点击链接后可进入H5页面进行认证并加入企业的流程。</li><li>**NONE**：系统会根据Endpoint生成签署链接，业务方需获取链接并通知客户。</li></ul>
	InvitationNotifyType *string `json:"InvitationNotifyType,omitnil,omitempty" name:"InvitationNotifyType"`

	// 回跳地址，为认证成功后页面进行回跳的URL，请确保回跳地址的可用性。注：`只有在员工邀请方式（InvitationNotifyType参数）为H5场景下才生效， 其他方式下设置无效。`
	JumpUrl *string `json:"JumpUrl,omitnil,omitempty" name:"JumpUrl"`

	// 要跳转的链接类型<ul><li> **HTTP**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型  ，此时返回长链 (默认类型)</li><li>**HTTP_SHORT_URL**：跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型，此时返回短链</li><li>**APP**： 第三方APP或小程序跳转电子签小程序的path,  APP或者小程序跳转适合此类型</li><li>**H5**： 第三方移动端浏览器进行嵌入，不支持小程序嵌入，过期时间一个月</li></ul>注意：InvitationNotifyType 和 Endpoint 的关系图<table><tbody><tr><td>通知类型（InvitationNotifyType）</td><td>Endpoint</td></tr><tr><td>SMS（默认）</td><td>不需要传递，会将 Endpoint 默认设置为HTTP_SHORT_URL</td></tr><tr><td>H5</td><td>不需要传递，会将 Endpoint 默认设置为 H5</td></tr><tr><td>NONE</td><td>所有 Endpoint 都支持（HTTP_URL/HTTP_SHORT_URL/H5/APP）默认为HTTP_SHORT_URL</td></tr></tbody></table>
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

func (r *UpdateIntegrationEmployeesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateIntegrationEmployeesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operator")
	delete(f, "Employees")
	delete(f, "Agent")
	delete(f, "InvitationNotifyType")
	delete(f, "JumpUrl")
	delete(f, "Endpoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateIntegrationEmployeesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIntegrationEmployeesResponseParams struct {
	// 更新成功的用户列表
	SuccessEmployeeData []*SuccessUpdateStaffData `json:"SuccessEmployeeData,omitnil,omitempty" name:"SuccessEmployeeData"`

	// 更新失败的用户列表
	FailedEmployeeData []*FailedUpdateStaffData `json:"FailedEmployeeData,omitnil,omitempty" name:"FailedEmployeeData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateIntegrationEmployeesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateIntegrationEmployeesResponseParams `json:"Response"`
}

func (r *UpdateIntegrationEmployeesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateIntegrationEmployeesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	// 文件对应业务类型,可以选择的类型如下
	// <ul><li> **TEMPLATE** : 此上传的文件用户生成合同模板，文件类型支持.pdf/.doc/.docx/.html格式，如果非pdf文件需要通过<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi" target="_blank">创建文件转换任务</a>转换后才能使用</li>
	// <li> **DOCUMENT** : 此文件用来发起合同流程，文件类型支持.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html，如果非pdf文件需要通过<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi" target="_blank">创建文件转换任务</a>转换后才能使用</li>
	// <li> **SEAL** : 此文件用于印章的生成，文件类型支持.jpg/.jpeg/.png</li></ul>
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// 执行本接口操作的员工信息。其中OperatorId为必填字段，即用户的UserId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Caller *Caller `json:"Caller,omitnil,omitempty" name:"Caller"`

	// 请上传文件内容数组，最多可上传20个文件。
	// 
	// <b>所有文件必须符合<font color="red">FileType</font>指定的文件类型。</b>
	FileInfos []*UploadFile `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 文件类型， 默认通过文件内容和文件后缀一起解析得到文件类型，调用接口时可以显示的指定上传文件的类型。
	// 可支持的指定类型如下:
	// <ul><li>pdf</li>
	// <li>doc</li>
	// <li>docx</li>
	// <li>xls</li>
	// <li>xlsx</li>
	// <li>html</li>
	// <li>jpg</li>
	// <li>jpeg</li>
	// <li>png</li></ul>
	// 如：pdf 表示上传的文件 张三入职合同.pdf的文件类型是 pdf
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 此参数仅对上传的PDF文件有效。其主要作用是确定是否将PDF中的灰色矩阵置为白色。
	// <ul><li>**true**：将灰色矩阵置为白色。</li>
	// <li>**false**：无需处理，不会将灰色矩阵置为白色（默认）。</li></ul>
	// 
	// 注: `该参数仅在关键字定位时，需要去除关键字所在的灰框场景下使用。`
	CoverRect *bool `json:"CoverRect,omitnil,omitempty" name:"CoverRect"`

	// 用户自定义ID数组，与上传文件一一对应
	// 
	// 注: `历史遗留问题，已经废弃，调用接口时不用赋值`
	CustomIds []*string `json:"CustomIds,omitnil,omitempty" name:"CustomIds"`

	// 不再使用，上传文件链接数组，最多支持20个URL
	//
	// Deprecated: FileUrls is deprecated.
	FileUrls *string `json:"FileUrls,omitnil,omitempty" name:"FileUrls"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type UploadFilesRequest struct {
	*tchttp.BaseRequest
	
	// 文件对应业务类型,可以选择的类型如下
	// <ul><li> **TEMPLATE** : 此上传的文件用户生成合同模板，文件类型支持.pdf/.doc/.docx/.html格式，如果非pdf文件需要通过<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi" target="_blank">创建文件转换任务</a>转换后才能使用</li>
	// <li> **DOCUMENT** : 此文件用来发起合同流程，文件类型支持.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html，如果非pdf文件需要通过<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi" target="_blank">创建文件转换任务</a>转换后才能使用</li>
	// <li> **SEAL** : 此文件用于印章的生成，文件类型支持.jpg/.jpeg/.png</li></ul>
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// 执行本接口操作的员工信息。其中OperatorId为必填字段，即用户的UserId。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Caller *Caller `json:"Caller,omitnil,omitempty" name:"Caller"`

	// 请上传文件内容数组，最多可上传20个文件。
	// 
	// <b>所有文件必须符合<font color="red">FileType</font>指定的文件类型。</b>
	FileInfos []*UploadFile `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 文件类型， 默认通过文件内容和文件后缀一起解析得到文件类型，调用接口时可以显示的指定上传文件的类型。
	// 可支持的指定类型如下:
	// <ul><li>pdf</li>
	// <li>doc</li>
	// <li>docx</li>
	// <li>xls</li>
	// <li>xlsx</li>
	// <li>html</li>
	// <li>jpg</li>
	// <li>jpeg</li>
	// <li>png</li></ul>
	// 如：pdf 表示上传的文件 张三入职合同.pdf的文件类型是 pdf
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 此参数仅对上传的PDF文件有效。其主要作用是确定是否将PDF中的灰色矩阵置为白色。
	// <ul><li>**true**：将灰色矩阵置为白色。</li>
	// <li>**false**：无需处理，不会将灰色矩阵置为白色（默认）。</li></ul>
	// 
	// 注: `该参数仅在关键字定位时，需要去除关键字所在的灰框场景下使用。`
	CoverRect *bool `json:"CoverRect,omitnil,omitempty" name:"CoverRect"`

	// 用户自定义ID数组，与上传文件一一对应
	// 
	// 注: `历史遗留问题，已经废弃，调用接口时不用赋值`
	CustomIds []*string `json:"CustomIds,omitnil,omitempty" name:"CustomIds"`

	// 不再使用，上传文件链接数组，最多支持20个URL
	FileUrls *string `json:"FileUrls,omitnil,omitempty" name:"FileUrls"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
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
	delete(f, "FileType")
	delete(f, "CoverRect")
	delete(f, "CustomIds")
	delete(f, "FileUrls")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFilesResponseParams struct {
	// 文件资源ID数组，每个文件资源ID为32位字符串。
	// 建议开发者保存此资源ID，后续创建合同或创建合同流程需此资源ID。
	// 注:`有效期一个小时（超过一小时后系统不定期清理，会有部分时间差）, 有效期内此文件id可以反复使用, 超过有效期无法使用`
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 上传成功文件数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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

type UserFlowType struct {
	// 合同类型ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserFlowTypeId *string `json:"UserFlowTypeId,omitnil,omitempty" name:"UserFlowTypeId"`

	// 合同类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 合同类型说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UserInfo struct {
	// 用户在平台的编号
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户的来源渠道，一般不用传，特定场景根据接口说明传值
	//
	// Deprecated: Channel is deprecated.
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// 用户在渠道的编号，一般不用传，特定场景根据接口说明传值
	//
	// Deprecated: OpenId is deprecated.
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 用户真实IP，内部字段，暂未开放
	//
	// Deprecated: ClientIp is deprecated.
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 用户代理IP，内部字段，暂未开放
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
	// <ul><li>中国大陆居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>中国港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给中国香港居民，“M”字头签发给中国澳门居民；第2位至第11位为数字。</li>
	// <li>中国港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil,omitempty" name:"IdCardNumber"`
}

// Predefined struct for user
type VerifyPdfRequestParams struct {
	// 合同流程ID，为32位字符串。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type VerifyPdfRequest struct {
	*tchttp.BaseRequest
	
	// 合同流程ID，为32位字符串。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。
	Operator *UserInfo `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
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
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyPdfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyPdfResponseParams struct {
	// 验签结果代码，代码的含义如下：
	// 
	// <ul><li>**1**：文件未被篡改，全部签名在腾讯电子签完成。</li>
	// <li>**2**：文件未被篡改，部分签名在腾讯电子签完成。</li>
	// <li>**3**：文件被篡改。</li>
	// <li>**4**：异常：文件内没有签名域。</li>
	// <li>**5**：异常：文件签名格式错误。</li></ul>
	VerifyResult *int64 `json:"VerifyResult,omitnil,omitempty" name:"VerifyResult"`

	// 验签结果详情，每个签名域对应的验签结果。状态值如下
	// <ul><li> **1** :验签成功，在电子签签署</li>
	// <li> **2** :验签成功，在其他平台签署</li>
	// <li> **3** :验签失败</li>
	// <li> **4** :pdf文件没有签名域</li>
	// <li> **5** :文件签名格式错误</li></ul>
	PdfVerifyResults []*PdfVerifyResult `json:"PdfVerifyResults,omitnil,omitempty" name:"PdfVerifyResults"`

	// 验签序列号, 为11为数组组成的字符串
	VerifySerialNo *string `json:"VerifySerialNo,omitnil,omitempty" name:"VerifySerialNo"`

	// 合同文件MD5哈希值
	PdfResourceMd5 *string `json:"PdfResourceMd5,omitnil,omitempty" name:"PdfResourceMd5"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type WebThemeConfig struct {
	// 是否显示页面底部电子签logo，取值如下：
	// <ul><li> **true**：页面底部显示电子签logo</li>
	// <li> **false**：页面底部不显示电子签logo（默认）</li></ul>
	DisplaySignBrandLogo *bool `json:"DisplaySignBrandLogo,omitnil,omitempty" name:"DisplaySignBrandLogo"`

	// 主题颜色：
	// 支持十六进制颜色值以及RGB格式颜色值，例如：#D54941，rgb(213, 73, 65)
	// <br/>
	WebEmbedThemeColor *string `json:"WebEmbedThemeColor,omitnil,omitempty" name:"WebEmbedThemeColor"`
}