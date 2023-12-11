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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Agent struct {
	// 应用的唯一标识(由电子签平台自动生成)。不同的业务系统可以采用不同的AppId，不同AppId下的数据是隔离的。可以由控制台开发者中心-应用集成自主生成。位置如下:
	// 
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/fac77e0d3f28aaec56669f67e65c8db8.png)
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 第三方应用平台自定义，对应第三方平台子客企业的唯一标识。一个第三方平台子客企业主体与子客企业ProxyOrganizationOpenId是一一对应的，不可更改，不可重复使用。（例如，可以使用企业名称的hash值，或者社会统一信用代码的hash值，或者随机hash值，需要第三方应用平台保存），最大64位字符串
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil" name:"ProxyOrganizationOpenId"`

	// 第三方平台子客企业中的员工/经办人，通过第三方应用平台进入电子签完成实名、且被赋予相关权限后，可以参与到企业资源的管理或签署流程中。
	ProxyOperator *UserInfo `json:"ProxyOperator,omitnil" name:"ProxyOperator"`

	// **不用填写**，在第三方平台子客企业开通电子签后，会生成唯一的子客应用Id（ProxyAppId）用于代理调用时的鉴权，在子客开通的回调中获取。
	ProxyAppId *string `json:"ProxyAppId,omitnil" name:"ProxyAppId"`

	// 内部参数，暂未开放使用
	//
	// Deprecated: ProxyOrganizationId is deprecated.
	ProxyOrganizationId *string `json:"ProxyOrganizationId,omitnil" name:"ProxyOrganizationId"`
}

type ApproverComponentLimitType struct {
	// 签署方经办人在模板中配置的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 签署方经办人控件类型是个人印章签署控件（SIGN_SIGNATURE） 时，可选的签名方式。
	// 
	// 签名方式：
	// <ul>
	// <li>HANDWRITE-手写签名</li>
	// <li>ESIGN-个人印章类型</li>
	// <li>OCR_ESIGN-AI智能识别手写签名</li>
	// <li>SYSTEM_ESIGN-系统签名</li>
	// </ul>
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type ApproverItem struct {
	// 签署方唯一编号
	// 
	// 在<a href="https://qian.tencent.com/developers/company/dynamic_signer" target="_blank">动态补充签署人</a>场景下，可以用此编号确定参与方
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"SignId,omitnil" name:"SignId"`

	// 签署方角色编号
	// 
	// 在<a href="https://qian.tencent.com/developers/company/dynamic_signer" target="_blank">动态补充签署人</a>场景下，可以用此编号确定参与方
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 签署方角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverRoleName *string `json:"ApproverRoleName,omitnil" name:"ApproverRoleName"`
}

type ApproverOption struct {
	// 是否可以拒签 默认false-可以拒签 true-不可以拒签
	NoRefuse *bool `json:"NoRefuse,omitnil" name:"NoRefuse"`

	// 是否可以转发 默认false-可以转发 true-不可以转发
	NoTransfer *bool `json:"NoTransfer,omitnil" name:"NoTransfer"`

	// 是否隐藏一键签署 默认false-不隐藏true-隐藏
	HideOneKeySign *bool `json:"HideOneKeySign,omitnil" name:"HideOneKeySign"`

	// 签署人信息补充类型，默认无需补充。
	// 
	// <ul><li> **1** : ( 动态签署人（可发起合同后再补充签署人信息）注：`企业自动签不支持动态补充`</li>
	// </ul>
	FillType *int64 `json:"FillType,omitnil" name:"FillType"`

	// 签署人阅读合同限制参数
	//  <br/>取值：
	// <ul>
	// <li> LimitReadTimeAndBottom，阅读合同必须限制阅读时长并且必须阅读到底</li>
	// <li> LimitReadTime，阅读合同仅限制阅读时长</li>
	// <li> LimitBottom，阅读合同仅限制必须阅读到底</li>
	// <li> NoReadTimeAndBottom，阅读合同不限制阅读时长且不限制阅读到底（白名单功能，请联系客户经理开白使用）</li>
	// </ul>
	FlowReadLimit *string `json:"FlowReadLimit,omitnil" name:"FlowReadLimit"`
}

type ApproverRestriction struct {
	// 指定签署人姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 指定签署人手机号，11位数字
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 指定签署人证件类型，ID_CARD-身份证，HONGKONG_AND_MACAO-港澳居民来往内地通行证，HONGKONG_MACAO_AND_TAIWAN-港澳台居民居住证
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 指定签署人证件号码，其中字母大写
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`
}

type AuthFailMessage struct {
	// 第三方平台子客企业的唯一标识，长度不能超过64，只能由字母和数字组成。开发者可自定义此字段的值，并需要保存此 ID 以便进行后续操作。
	// 
	// 一个第三方平台子客企业主体与子客企业 ProxyOrganizationOpenId 是一一对应的，不可更改，不可重复使用。例如，可以使用企业名称的哈希值，或者社会统一信用代码的哈希值，或者随机哈希值。
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil" name:"ProxyOrganizationOpenId"`

	// 错误信息
	Message *string `json:"Message,omitnil" name:"Message"`
}

type AuthorizedUser struct {
	// 第三方应用平台的用户openid
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`
}

type AutoSignConfig struct {
	// 自动签开通个人用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 是否回调证书信息:
	// <ul><li>**false**: 不需要(默认)</li>
	// <li>**true**:需要</li></ul>
	CertInfoCallback *bool `json:"CertInfoCallback,omitnil" name:"CertInfoCallback"`

	// 是否支持用户自定义签名印章:
	// <ul><li>**false**: 不能自己定义(默认)</li>
	// <li>**true**: 可以自己定义</li></ul>
	UserDefineSeal *bool `json:"UserDefineSeal,omitnil" name:"UserDefineSeal"`

	// 回调中是否需要自动签将要使用的印章（签名）图片的 base64:
	// <ul><li>**false**: 不需要(默认)</li>
	// <li>**true**: 需要</li></ul>
	SealImgCallback *bool `json:"SealImgCallback,omitnil" name:"SealImgCallback"`

	// 回调链接，如果渠道已经配置了，可以不传
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 开通时候的身份验证方式, 取值为：
	// <ul><li>**WEIXINAPP** : 微信人脸识别</li>
	// <li>**INSIGHT** : 慧眼人脸认别</li>
	// <li>**TELECOM** : 运营商三要素验证</li></ul>
	// 注：
	// <ul><li>如果是小程序开通链接，支持传 WEIXINAPP / TELECOM。为空默认 WEIXINAPP</li>
	// <li>如果是 H5 开通链接，支持传 INSIGHT / TELECOM。为空默认 INSIGHT </li></ul>
	VerifyChannels []*string `json:"VerifyChannels,omitnil" name:"VerifyChannels"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。
	// 
	// <ul><li>**0**: (默认) 使用个人自动签账号许可进行开通，个人自动签账号许可有效期1年，注: `不可解绑释放更换他人`</li>
	// <li>**1**: 不绑定自动签账号许可开通，后续使用合同份额进行合同发起</li></ul>
	LicenseType *int64 `json:"LicenseType,omitnil" name:"LicenseType"`
}

type BaseFlowInfo struct {
	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 合同流程的签署顺序类型：
	// **false**：(默认)有序签署, 本合同多个参与人需要依次签署
	// **true**：无序签署, 本合同多个参与人没有先后签署限制
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 是否打开智能添加填写区(默认开启，打开:"OPEN" 关闭："CLOSE")
	IntelligentStatus *string `json:"IntelligentStatus,omitnil" name:"IntelligentStatus"`

	// 填写控件内容， 填写的控制的ID-填写的内容对列表
	FormFields []*FormField `json:"FormFields,omitnil" name:"FormFields"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为1000长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的回调通知模块。
	UserData *string `json:"UserData,omitnil" name:"UserData"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	// 
	// 注:`此功能为白名单功能，使用前请联系对接的客户经理沟通。`
	CcInfos []*CcInfo `json:"CcInfos,omitnil" name:"CcInfos"`

	// 发起方企业的签署人进行发起操作是否需要企业内部审批。使用此功能需要发起方企业有参与签署。
	// 
	// 若设置为true，发起审核结果需通过接口 [提交企业签署流程审批结果](https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateFlowSignReview)通知电子签，审核通过后，发起方企业签署人方可进行发起操作，否则会阻塞其发起操作。
	// 
	NeedCreateReview *bool `json:"NeedCreateReview,omitnil" name:"NeedCreateReview"`

	// 填写控件：文件发起使用
	Components []*Component `json:"Components,omitnil" name:"Components"`
}

type BillUsageDetail struct {
	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 合同经办人名称
	// 如果有多个经办人用分号隔开。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorName *string `json:"OperatorName,omitnil" name:"OperatorName"`

	// 发起方组织机构名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateOrganizationName *string `json:"CreateOrganizationName,omitnil" name:"CreateOrganizationName"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 该名称还将用于合同签署完成后的下载文件名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 当前合同状态,如下是状态码对应的状态。
	// 0-还没有发起
	// 1-等待签署
	// 2-部分签署 
	// 3-拒签
	// 4-已签署 
	// 5-已过期 
	// 6-已撤销 
	// 7-还没有预发起
	// 8-等待填写
	// 9-部分填写 
	// 10-拒填
	// 11-已解除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 套餐类型
	// 对应关系如下
	// CloudEnterprise-企业版合同
	// SingleSignature-单方签章
	// CloudProve-签署报告
	// CloudOnlineSign-腾讯会议在线签约
	// ChannelWeCard-微工卡
	// SignFlow-合同套餐
	// SignFace-签署意愿（人脸识别）
	// SignPassword-签署意愿（密码）
	// SignSMS-签署意愿（短信）
	// PersonalEssAuth-签署人实名（腾讯电子签认证）
	// PersonalThirdAuth-签署人实名（信任第三方认证）
	// OrgEssAuth-签署企业实名
	// FlowNotify-短信通知
	// AuthService-企业工商信息查询
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaType *string `json:"QuotaType,omitnil" name:"QuotaType"`

	// 合同使用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseCount *int64 `json:"UseCount,omitnil" name:"UseCount"`

	// 消耗的时间戳，格式为Unix标准时间戳（秒）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *int64 `json:"CostTime,omitnil" name:"CostTime"`

	// 消耗的套餐名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaName *string `json:"QuotaName,omitnil" name:"QuotaName"`

	// 消耗类型
	// 1.扣费 2.撤销返还
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostType *int64 `json:"CostType,omitnil" name:"CostType"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type CcInfo struct {
	// 被抄送人手机号，大陆11位手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 被抄送人姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 被抄送人类型
	// 0--个人. 1--员工
	CcType *int64 `json:"CcType,omitnil" name:"CcType"`

	// 被抄送人权限
	// 0--可查看
	// 1--可查看也可下载
	CcPermission *int64 `json:"CcPermission,omitnil" name:"CcPermission"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 要撤销的合同流程ID列表，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 撤回原因，长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 
	// 备注:`如果不传递撤回原因，那么默认撤回原因是 "自动撤销（通过接口实现）"`
	CancelMessage *string `json:"CancelMessage,omitnil" name:"CancelMessage"`

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
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitnil" name:"CancelMessageFormat"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 要撤销的合同流程ID列表，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 撤回原因，长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 
	// 备注:`如果不传递撤回原因，那么默认撤回原因是 "自动撤销（通过接口实现）"`
	CancelMessage *string `json:"CancelMessage,omitnil" name:"CancelMessage"`

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
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitnil" name:"CancelMessageFormat"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	FailMessages []*string `json:"FailMessages,omitnil" name:"FailMessages"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 合同经办人名称
	// 如果有多个经办人用分号隔开。
	OperatorName *string `json:"OperatorName,omitnil" name:"OperatorName"`

	// 发起方组织机构名称
	CreateOrganizationName *string `json:"CreateOrganizationName,omitnil" name:"CreateOrganizationName"`

	// 合同流程的名称。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

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
	FlowStatus *string `json:"FlowStatus,omitnil" name:"FlowStatus"`

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
	QuotaType *string `json:"QuotaType,omitnil" name:"QuotaType"`

	// 合同使用量
	// 注: `如果消耗类型是撤销返还，此值为负值代表返还的合同数量`
	UseCount *int64 `json:"UseCount,omitnil" name:"UseCount"`

	// 消耗的时间戳，格式为Unix标准时间戳（秒）。
	CostTime *int64 `json:"CostTime,omitnil" name:"CostTime"`

	// 消耗的套餐名称
	QuotaName *string `json:"QuotaName,omitnil" name:"QuotaName"`

	// 消耗类型
	// **1**.扣费 
	// **2**.撤销返还
	CostType *int64 `json:"CostType,omitnil" name:"CostType"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

// Predefined struct for user
type ChannelCancelFlowRequestParams struct {
	// 要撤销的合同流程ID
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 撤回原因，长度不能超过200，只能由中文、字母、数字和下划线组成。
	CancelMessage *string `json:"CancelMessage,omitnil" name:"CancelMessage"`

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
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitnil" name:"CancelMessageFormat"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCancelFlowRequest struct {
	*tchttp.BaseRequest
	
	// 要撤销的合同流程ID
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 撤回原因，长度不能超过200，只能由中文、字母、数字和下划线组成。
	CancelMessage *string `json:"CancelMessage,omitnil" name:"CancelMessage"`

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
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitnil" name:"CancelMessageFormat"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 需要取消签署的二维码ID，为32位字符串。由[创建一码多扫流程签署二维码](https://qian.tencent.com/developers/partnerApis/templates/ChannelCreateMultiFlowSignQRCode)返回
	QrCodeId *string `json:"QrCodeId,omitnil" name:"QrCodeId"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 需要取消签署的二维码ID，为32位字符串。由[创建一码多扫流程签署二维码](https://qian.tencent.com/developers/partnerApis/templates/ChannelCreateMultiFlowSignQRCode)返回
	QrCodeId *string `json:"QrCodeId,omitnil" name:"QrCodeId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 指定撤销链接的用户信息，包含姓名、证件类型、证件号码。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`
}

type ChannelCancelUserAutoSignEnableUrlRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 指定撤销链接的用户信息，包含姓名、证件类型、证件号码。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 要撤销的合同流程ID列表，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCreateBatchCancelFlowUrlRequest struct {
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 要撤销的合同流程ID列表，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	BatchCancelFlowUrl *string `json:"BatchCancelFlowUrl,omitnil" name:"BatchCancelFlowUrl"`

	// 与入参的FlowIds数组一致,   成功生成到撤销链接中,则为"",   不能撤销合同则为失败原因
	FailMessages []*string `json:"FailMessages,omitnil" name:"FailMessages"`

	// 签署撤销链接的过期时间(格式为:年-月-日 时:分:秒), 默认是生成链接的24小时后失效
	// 
	UrlExpireOn *string `json:"UrlExpireOn,omitnil" name:"UrlExpireOn"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 注:
	// `1. ApproverType目前只支持个人类型的签署人。`
	// `2. 签署人只能有手写签名和时间类型的签署控件，其他类型的填写控件和签署控件暂时都未支持。`
	// `3. 当需要通过短信验证码签署时，手机号ApproverMobile需要与发起合同时填写的用户手机号一致。`
	FlowApproverInfo *FlowApproverInfo `json:"FlowApproverInfo,omitnil" name:"FlowApproverInfo"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 批量签署的合同流程ID数组。
	// 注: `在调用此接口时，请确保合同流程均为本企业发起，且合同数量不超过100个。`
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 合同组编号
	// 注：`该参数和合同流程ID数组必须二选一`
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 签署完之后的H5页面的跳转链接，此链接及支持http://和https://，最大长度1000个字符。(建议https协议)
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`

	// 指定批量签署合同的签名类型，可传递以下值：
	// <ul><li>**0**：手写签名(默认)</li>
	// <li>**1**：OCR楷体</li></ul>
	// 注：
	// <ul><li>默认情况下，签名类型为手写签名</li>
	// <li>您可以传递多种值，表示可用多种签名类型。</li></ul>
	SignatureTypes []*int64 `json:"SignatureTypes,omitnil" name:"SignatureTypes"`

	// 指定批量签署合同的认证校验方式，可传递以下值：
	// <ul><li>**1**：人脸认证(默认)，需进行人脸识别成功后才能签署合同</li>
	// <li>**2**：密码认证(默认)，需输入与用户在腾讯电子签设置的密码一致才能校验成功进行合同签署</li>
	// <li>**3**：运营商三要素，需到运营商处比对手机号实名信息(名字、手机号、证件号)校验一致才能成功进行合同签署。</li></ul>
	// 注：
	// <ul><li>默认情况下，认证校验方式为人脸和密码认证</li>
	// <li>您可以传递多种值，表示可用多种认证校验方式。</li></ul>
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil" name:"ApproverSignTypes"`
}

type ChannelCreateBatchQuickSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 批量签署的流程签署人，其中姓名(ApproverName)、参与人类型(ApproverType)必传，手机号(ApproverMobile)和证件信息(ApproverIdCardType、ApproverIdCardNumber)可任选一种或全部传入。
	// 注:
	// `1. ApproverType目前只支持个人类型的签署人。`
	// `2. 签署人只能有手写签名和时间类型的签署控件，其他类型的填写控件和签署控件暂时都未支持。`
	// `3. 当需要通过短信验证码签署时，手机号ApproverMobile需要与发起合同时填写的用户手机号一致。`
	FlowApproverInfo *FlowApproverInfo `json:"FlowApproverInfo,omitnil" name:"FlowApproverInfo"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 批量签署的合同流程ID数组。
	// 注: `在调用此接口时，请确保合同流程均为本企业发起，且合同数量不超过100个。`
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 合同组编号
	// 注：`该参数和合同流程ID数组必须二选一`
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 签署完之后的H5页面的跳转链接，此链接及支持http://和https://，最大长度1000个字符。(建议https协议)
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`

	// 指定批量签署合同的签名类型，可传递以下值：
	// <ul><li>**0**：手写签名(默认)</li>
	// <li>**1**：OCR楷体</li></ul>
	// 注：
	// <ul><li>默认情况下，签名类型为手写签名</li>
	// <li>您可以传递多种值，表示可用多种签名类型。</li></ul>
	SignatureTypes []*int64 `json:"SignatureTypes,omitnil" name:"SignatureTypes"`

	// 指定批量签署合同的认证校验方式，可传递以下值：
	// <ul><li>**1**：人脸认证(默认)，需进行人脸识别成功后才能签署合同</li>
	// <li>**2**：密码认证(默认)，需输入与用户在腾讯电子签设置的密码一致才能校验成功进行合同签署</li>
	// <li>**3**：运营商三要素，需到运营商处比对手机号实名信息(名字、手机号、证件号)校验一致才能成功进行合同签署。</li></ul>
	// 注：
	// <ul><li>默认情况下，认证校验方式为人脸和密码认证</li>
	// <li>您可以传递多种值，表示可用多种认证校验方式。</li></ul>
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil" name:"ApproverSignTypes"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateBatchQuickSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateBatchQuickSignUrlResponseParams struct {
	// 签署人签署链接信息
	FlowApproverUrlInfo *FlowApproverUrlInfo `json:"FlowApproverUrlInfo,omitnil" name:"FlowApproverUrlInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 关于渠道应用的相关信息，包括子客企业及应用编、号等详细内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：`请确保和合同中填入的一致`
	Name *string `json:"Name,omitnil" name:"Name"`

	// 手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此业务通知方。
	// 
	// 注：`请确保和合同中填入的一致,  若无法保持一致，请确保在发起和生成批量签署链接时传入相同的参与方证件信息`
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 证件类型，支持以下类型
	// <ul><li>**ID_CARD** : 居民身份证 (默认值)</li>
	// <li>**HONGKONG_AND_MACAO** : 港澳居民来往内地通行证</li>
	// <li>**HONGKONG_MACAO_AND_TAIWAN** : 港澳台居民居住证(格式同居民身份证)</li></ul>
	// 
	// 注：`请确保和合同中填入的一致`
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母（但“I”、“O”除外），后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	// 
	// 注：`请确保和合同中填入的一致`
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 通知用户方式：
	// <ul>
	// <li>**NONE** : 不通知（默认）</li>
	// <li>**SMS** : 短信通知（发送短信通知到Mobile参数所传的手机号）</li>
	// </ul>
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 本次需要批量签署的合同流程ID列表。
	// 可以不传,  如不传则是发给对方的所有待签署合同流程。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 目标签署人的企业名称，签署人如果是企业员工身份，需要传此参数。
	// 
	// 注：
	// <ul>
	// <li>请确认该名称与企业营业执照中注册的名称一致。</li>
	// <li>如果名称中包含英文括号()，请使用中文括号（）代替。</li>
	// <li>请确保此企业已完成腾讯电子签企业认证。</li>
	// <li>若为子客企业，请确保员工已经加入企业。</li>
	// </ul>
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 是否直接跳转至合同内容页面进行签署
	// <ul>
	// <li>**false**: 会跳转至批量合同流程的列表,  点击需要批量签署合同后进入合同内容页面进行签署(默认)</li>
	// <li>**true**: 跳过合同流程列表, 直接进入合同内容页面进行签署</li>
	// </ul>
	JumpToDetail *bool `json:"JumpToDetail,omitnil" name:"JumpToDetail"`
}

type ChannelCreateBatchSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括子客企业及应用编、号等详细内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：`请确保和合同中填入的一致`
	Name *string `json:"Name,omitnil" name:"Name"`

	// 手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此业务通知方。
	// 
	// 注：`请确保和合同中填入的一致,  若无法保持一致，请确保在发起和生成批量签署链接时传入相同的参与方证件信息`
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 证件类型，支持以下类型
	// <ul><li>**ID_CARD** : 居民身份证 (默认值)</li>
	// <li>**HONGKONG_AND_MACAO** : 港澳居民来往内地通行证</li>
	// <li>**HONGKONG_MACAO_AND_TAIWAN** : 港澳台居民居住证(格式同居民身份证)</li></ul>
	// 
	// 注：`请确保和合同中填入的一致`
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母（但“I”、“O”除外），后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	// 
	// 注：`请确保和合同中填入的一致`
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 通知用户方式：
	// <ul>
	// <li>**NONE** : 不通知（默认）</li>
	// <li>**SMS** : 短信通知（发送短信通知到Mobile参数所传的手机号）</li>
	// </ul>
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 本次需要批量签署的合同流程ID列表。
	// 可以不传,  如不传则是发给对方的所有待签署合同流程。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 目标签署人的企业名称，签署人如果是企业员工身份，需要传此参数。
	// 
	// 注：
	// <ul>
	// <li>请确认该名称与企业营业执照中注册的名称一致。</li>
	// <li>如果名称中包含英文括号()，请使用中文括号（）代替。</li>
	// <li>请确保此企业已完成腾讯电子签企业认证。</li>
	// <li>若为子客企业，请确保员工已经加入企业。</li>
	// </ul>
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 是否直接跳转至合同内容页面进行签署
	// <ul>
	// <li>**false**: 会跳转至批量合同流程的列表,  点击需要批量签署合同后进入合同内容页面进行签署(默认)</li>
	// <li>**true**: 跳过合同流程列表, 直接进入合同内容页面进行签署</li>
	// </ul>
	JumpToDetail *bool `json:"JumpToDetail,omitnil" name:"JumpToDetail"`
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
	delete(f, "JumpToDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateBatchSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateBatchSignUrlResponseParams struct {
	// 批量签署链接，以短链形式返回，短链的有效期参考回参中的 ExpiredTime。
	// 
	// 注: `非小程序和APP集成使用`
	SignUrl *string `json:"SignUrl,omitnil" name:"SignUrl"`

	// 链接过期时间以 Unix 时间戳格式表示，从生成链接时间起，往后7天有效期。过期后短链将失效，无法打开。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 从客户小程序或者客户APP跳转至腾讯电子签小程序进行批量签署的跳转路径
	// 
	// 注: `小程序和APP集成使用`
	MiniAppPath *string `json:"MiniAppPath,omitnil" name:"MiniAppPath"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 需要领取的合同流程的ID列表
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 需要领取的合同流程的ID列表
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

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
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 需要进行转换操作的文件资源名称，带资源后缀名。
	// 
	// 注:  `资源名称长度限制为256个字符`
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 需要进行转换操作的文件资源Id，通过<a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">UploadFiles</a>接口获取文件资源Id。
	// 
	// 注:  `目前，此接口仅支持单个文件进行转换。`
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 调用方用户信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

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
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 需要进行转换操作的文件资源名称，带资源后缀名。
	// 
	// 注:  `资源名称长度限制为256个字符`
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 需要进行转换操作的文件资源Id，通过<a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">UploadFiles</a>接口获取文件资源Id。
	// 
	// 注:  `目前，此接口仅支持单个文件进行转换。`
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 调用方用户信息，不用传
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`
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
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 要生成WEB嵌入界面的类型, 可以选择的值如下: 
	// 
	// <ul>
	// <li>CREATE_SEAL: 生成创建印章的嵌入页面</li>
	// <li>CREATE_TEMPLATE：生成创建模板的嵌入页面</li>
	// <li>MODIFY_TEMPLATE：生成修改模板的嵌入页面</li>
	// <li>PREVIEW_TEMPLATE：生成预览模板的嵌入页面</li>
	// <li>PREVIEW_FLOW：生成预览合同文档的嵌入页面</li>
	// <li>PREVIEW_FLOW_DETAIL：生成预览合同详情的嵌入页面</li>
	// <li>PREVIEW_SEAL_LIST：生成预览印章列表的嵌入页面</li>
	// <li>PREVIEW_SEAL_DETAIL：生成预览印章详情的嵌入页面</li>
	// <li>EXTEND_SERVICE：生成扩展服务的嵌入页面</li>
	// </ul>
	EmbedType *string `json:"EmbedType,omitnil" name:"EmbedType"`

	// WEB嵌入的业务资源ID
	// 
	// <ul>
	// <li>当EmbedType取值MODIFY_TEMPLATE，PREVIEW_TEMPLATE时需要填写模板id作为BusinessId</li>
	// <li>当EmbedType取值PREVIEW_FLOW，PREVIEW_FLOW_DETAIL时需要填写合同id作为BusinessId</li>
	// <li>当EmbedType取值PREVIEW_SEAL_DETAIL需要填写印章id作为BusinessId</li>
	// </ul>
	BusinessId *string `json:"BusinessId,omitnil" name:"BusinessId"`

	// 是否隐藏控件，只有预览模板时生效
	HiddenComponents *bool `json:"HiddenComponents,omitnil" name:"HiddenComponents"`

	// 渠道操作者信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

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
	UserData *string `json:"UserData,omitnil" name:"UserData"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 要生成WEB嵌入界面的类型, 可以选择的值如下: 
	// 
	// <ul>
	// <li>CREATE_SEAL: 生成创建印章的嵌入页面</li>
	// <li>CREATE_TEMPLATE：生成创建模板的嵌入页面</li>
	// <li>MODIFY_TEMPLATE：生成修改模板的嵌入页面</li>
	// <li>PREVIEW_TEMPLATE：生成预览模板的嵌入页面</li>
	// <li>PREVIEW_FLOW：生成预览合同文档的嵌入页面</li>
	// <li>PREVIEW_FLOW_DETAIL：生成预览合同详情的嵌入页面</li>
	// <li>PREVIEW_SEAL_LIST：生成预览印章列表的嵌入页面</li>
	// <li>PREVIEW_SEAL_DETAIL：生成预览印章详情的嵌入页面</li>
	// <li>EXTEND_SERVICE：生成扩展服务的嵌入页面</li>
	// </ul>
	EmbedType *string `json:"EmbedType,omitnil" name:"EmbedType"`

	// WEB嵌入的业务资源ID
	// 
	// <ul>
	// <li>当EmbedType取值MODIFY_TEMPLATE，PREVIEW_TEMPLATE时需要填写模板id作为BusinessId</li>
	// <li>当EmbedType取值PREVIEW_FLOW，PREVIEW_FLOW_DETAIL时需要填写合同id作为BusinessId</li>
	// <li>当EmbedType取值PREVIEW_SEAL_DETAIL需要填写印章id作为BusinessId</li>
	// </ul>
	BusinessId *string `json:"BusinessId,omitnil" name:"BusinessId"`

	// 是否隐藏控件，只有预览模板时生效
	HiddenComponents *bool `json:"HiddenComponents,omitnil" name:"HiddenComponents"`

	// 渠道操作者信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

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
	UserData *string `json:"UserData,omitnil" name:"UserData"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateEmbedWebUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateEmbedWebUrlResponseParams struct {
	// 嵌入的web链接，5分钟有效
	WebUrl *string `json:"WebUrl,omitnil" name:"WebUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程ID，为32位字符串。 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 补充企业签署人信息。
	// 
	// - 如果发起方指定的补充签署人是企业签署人，则需要提供企业名称或者企业OpenId；
	// 
	// - 如果不指定，则使用姓名和手机号进行补充。
	Approvers []*FillApproverInfo `json:"Approvers,omitnil" name:"Approvers"`

	// 签署人信息补充方式
	// 
	// <ul><li>**1**: 表示往未指定签署人的节点，添加一个明确的签署人，支持企业或个人签署方。</li></ul>
	FillApproverType *int64 `json:"FillApproverType,omitnil" name:"FillApproverType"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程ID，为32位字符串。 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 补充企业签署人信息。
	// 
	// - 如果发起方指定的补充签署人是企业签署人，则需要提供企业名称或者企业OpenId；
	// 
	// - 如果不指定，则使用姓名和手机号进行补充。
	Approvers []*FillApproverInfo `json:"Approvers,omitnil" name:"Approvers"`

	// 签署人信息补充方式
	// 
	// <ul><li>**1**: 表示往未指定签署人的节点，添加一个明确的签署人，支持企业或个人签署方。</li></ul>
	FillApproverType *int64 `json:"FillApproverType,omitnil" name:"FillApproverType"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	delete(f, "FlowId")
	delete(f, "Approvers")
	delete(f, "FillApproverType")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowApproversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowApproversResponseParams struct {
	// 批量补充签署人时，补充失败的报错说明 
	// 注:`目前仅补充动态签署人时会返回补充失败的原因`	
	// 注意：此字段可能返回 null，表示取不到有效值。
	FillError []*FillError `json:"FillError,omitnil" name:"FillError"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业标识: Agent. ProxyOperator.OpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 合同流程的参与方列表, 最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#flowapproverinfo" target="_blank">FlowApproverInfo结构体</a>。
	// 
	// 如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序, 请确保列表中参与人的顺序符合实际签署顺序。
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitnil" name:"FlowApprovers"`

	// 本合同流程需包含的PDF文件资源编号列表，通过<a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">UploadFiles</a>接口获取PDF文件资源编号。
	// 
	// 注: `目前，此接口仅支持单个文件发起。`
	FileIds []*string `json:"FileIds,omitnil" name:"FileIds"`

	// 模板或者合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li>单行文本控件</li>
	// <li>多行文本控件</li>
	// <li>勾选框控件</li>
	// <li>数字控件</li>
	// <li>图片控件</li>
	// <li>数据表格等填写控件</li></ul>
	Components []*Component `json:"Components,omitnil" name:"Components"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 执行结果的回调URL，长度不超过255个字符，该URL仅支持HTTP或HTTPS协议，建议采用HTTPS协议以保证数据传输的安全性。
	// 腾讯电子签服务器将通过POST方式，application/json格式通知执行结果，请确保外网可以正常访问该URL。
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/partner/callback_data_types" target="_blank">回调通知</a>模块。
	// 
	// 注:
	// `如果不传递回调地址， 则默认是配置应用号时候使用的回调地址`
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	// **注**: `有序签署时以传入FlowApprovers数组的顺序作为签署顺序`
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为255个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

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
	CustomShowMap *string `json:"CustomShowMap,omitnil" name:"CustomShowMap"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 1000长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/partner/callback_types_contracts_sign" target="_blank">回调通知</a>模块。
	CustomerData *string `json:"CustomerData,omitnil" name:"CustomerData"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过ChannelCreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证，用户手机号和参与方手机号（ApproverMobile）相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// 参数说明：可选人脸识别或手机号验证两种方式，若选择后者，未实名个人签署方在签署合同时，无需经过实名认证和意愿确认两次人脸识别，该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil" name:"ApproverVerifyType"`

	// 签署方签署控件（印章/签名等）的生成方式：
	// <ul><li> **0**：在合同流程发起时，由发起人指定签署方的签署控件的位置和数量。</li>
	// <li> **1**：签署方在签署时自行添加签署控件，可以拖动位置和控制数量。</li></ul>
	// **注**: `发起后添加控件功能不支持添加签批控件`
	SignBeanTag *int64 `json:"SignBeanTag,omitnil" name:"SignBeanTag"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	CcInfos []*CcInfo `json:"CcInfos,omitnil" name:"CcInfos"`

	// 可以设置以下时间节点来给抄送人发送短信通知来查看合同内容：
	// <ul><li> **0**：合同发起时通知（默认值）</li>
	// <li> **1**：签署完成后通知</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil" name:"CcNotifyType"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：处方单（医疗自动签）  </li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil" name:"AutoSignScene"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCreateFlowByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业标识: Agent. ProxyOperator.OpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 合同流程的参与方列表, 最多可支持50个参与方，可在列表中指定企业B端签署方和个人C端签署方的联系和认证方式等信息，具体定义可以参考开发者中心的<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#flowapproverinfo" target="_blank">FlowApproverInfo结构体</a>。
	// 
	// 如果合同流程是有序签署，Approvers列表中参与人的顺序就是默认的签署顺序, 请确保列表中参与人的顺序符合实际签署顺序。
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitnil" name:"FlowApprovers"`

	// 本合同流程需包含的PDF文件资源编号列表，通过<a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">UploadFiles</a>接口获取PDF文件资源编号。
	// 
	// 注: `目前，此接口仅支持单个文件发起。`
	FileIds []*string `json:"FileIds,omitnil" name:"FileIds"`

	// 模板或者合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li>单行文本控件</li>
	// <li>多行文本控件</li>
	// <li>勾选框控件</li>
	// <li>数字控件</li>
	// <li>图片控件</li>
	// <li>数据表格等填写控件</li></ul>
	Components []*Component `json:"Components,omitnil" name:"Components"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 执行结果的回调URL，长度不超过255个字符，该URL仅支持HTTP或HTTPS协议，建议采用HTTPS协议以保证数据传输的安全性。
	// 腾讯电子签服务器将通过POST方式，application/json格式通知执行结果，请确保外网可以正常访问该URL。
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/partner/callback_data_types" target="_blank">回调通知</a>模块。
	// 
	// 注:
	// `如果不传递回调地址， 则默认是配置应用号时候使用的回调地址`
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 合同流程的签署顺序类型：
	// <ul><li> **false**：(默认)有序签署, 本合同多个参与人需要依次签署 </li>
	// <li> **true**：无序签署, 本合同多个参与人没有先后签署限制</li></ul>
	// **注**: `有序签署时以传入FlowApprovers数组的顺序作为签署顺序`
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为255个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

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
	CustomShowMap *string `json:"CustomShowMap,omitnil" name:"CustomShowMap"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为 1000长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/partner/callback_types_contracts_sign" target="_blank">回调通知</a>模块。
	CustomerData *string `json:"CustomerData,omitnil" name:"CustomerData"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过ChannelCreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证，用户手机号和参与方手机号（ApproverMobile）相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// 参数说明：可选人脸识别或手机号验证两种方式，若选择后者，未实名个人签署方在签署合同时，无需经过实名认证和意愿确认两次人脸识别，该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil" name:"ApproverVerifyType"`

	// 签署方签署控件（印章/签名等）的生成方式：
	// <ul><li> **0**：在合同流程发起时，由发起人指定签署方的签署控件的位置和数量。</li>
	// <li> **1**：签署方在签署时自行添加签署控件，可以拖动位置和控制数量。</li></ul>
	// **注**: `发起后添加控件功能不支持添加签批控件`
	SignBeanTag *int64 `json:"SignBeanTag,omitnil" name:"SignBeanTag"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	CcInfos []*CcInfo `json:"CcInfos,omitnil" name:"CcInfos"`

	// 可以设置以下时间节点来给抄送人发送短信通知来查看合同内容：
	// <ul><li> **0**：合同发起时通知（默认值）</li>
	// <li> **1**：签署完成后通知</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil" name:"CcNotifyType"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：处方单（医疗自动签）  </li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil" name:"AutoSignScene"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowByFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowByFilesResponseParams struct {
	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 签署方信息，如角色ID、角色名称等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Approvers []*ApproverItem `json:"Approvers,omitnil" name:"Approvers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	FlowFileInfos []*FlowFileInfo `json:"FlowFileInfos,omitnil" name:"FlowFileInfos"`

	// 合同组的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同组中签署人校验和认证的方式：
	// <ul><li>**VerifyCheck**：人脸识别（默认）</li>
	// <li>**MobileCheck**：手机号验证</li></ul>
	// 注意：
	// `1. MobileCheck 方式，未实名的个人/自然人签署方无需进行人脸识别实名认证即可查看合同（但签署合同时仍然需要人脸实名），企业签署方需经过人脸认证。`
	// `2. 合同组的校验和认证的方式会优先使用，会覆盖合同组中单个合同和合同签署方认证方式的限制配置。`
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil" name:"ApproverVerifyType"`

	// 合同组的签署配置项信息，例如在合同组签署过程中，是否需要对每个子合同进行独立的意愿确认。
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil" name:"FlowGroupOptions"`

	// 操作者的信息，此参数不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCreateFlowGroupByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 合同组中每个合同签署流程的信息，合同组中最少包含2个合同，不能超过50个合同。
	FlowFileInfos []*FlowFileInfo `json:"FlowFileInfos,omitnil" name:"FlowFileInfos"`

	// 合同组的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同组中签署人校验和认证的方式：
	// <ul><li>**VerifyCheck**：人脸识别（默认）</li>
	// <li>**MobileCheck**：手机号验证</li></ul>
	// 注意：
	// `1. MobileCheck 方式，未实名的个人/自然人签署方无需进行人脸识别实名认证即可查看合同（但签署合同时仍然需要人脸实名），企业签署方需经过人脸认证。`
	// `2. 合同组的校验和认证的方式会优先使用，会覆盖合同组中单个合同和合同签署方认证方式的限制配置。`
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil" name:"ApproverVerifyType"`

	// 合同组的签署配置项信息，例如在合同组签署过程中，是否需要对每个子合同进行独立的意愿确认。
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil" name:"FlowGroupOptions"`

	// 操作者的信息，此参数不用传
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 合同组中每个合同流程ID，每个ID均为32位字符串。
	// 
	// 注:
	// `此数组的顺序和入参中的FlowGroupInfos顺序一致`
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同组中每个合同签署流程的信息，合同组中最少包含2个合同，不能超过50个合同。
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil" name:"FlowInfos"`

	// 合同组的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`
}

type ChannelCreateFlowGroupByTemplatesRequest struct {
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同组中每个合同签署流程的信息，合同组中最少包含2个合同，不能超过50个合同。
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil" name:"FlowInfos"`

	// 合同组的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`
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
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 合同组中每个合同流程ID，每个ID均为32位字符串。
	// 
	// 注:
	// `此数组的顺序和入参中的FlowInfos顺序一致`
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 复杂文档合成任务（如，包含动态表格的预览任务）的任务信息数组；
	// 如果文档需要异步合成，此字段会返回该异步任务的任务信息，后续可以通过ChannelGetTaskResultApi接口查询任务详情；
	TaskInfos []*TaskInfo `json:"TaskInfos,omitnil" name:"TaskInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 需执行催办的合同流程ID数组，最多支持100个。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 需执行催办的合同流程ID数组，最多支持100个。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`
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
	RemindFlowRecords []*RemindFlowRecords `json:"RemindFlowRecords,omitnil" name:"RemindFlowRecords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 企业内部审核结果
	// PASS: 通过
	// REJECT: 拒绝
	// SIGN_REJECT:拒签(流程结束)
	ReviewType *string `json:"ReviewType,omitnil" name:"ReviewType"`

	// 审核原因 
	// 当ReviewType 是REJECT 时此字段必填,字符串长度不超过200
	ReviewMessage *string `json:"ReviewMessage,omitnil" name:"ReviewMessage"`

	// 签署节点审核时需要指定，给个人审核时必填。
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 操作类型，默认：SignReview；SignReview:签署审核，CreateReview：发起审核
	// 注：接口通过该字段区分操作类型
	// 该字段不传或者为空，则默认为SignReview签署审核，走签署审核流程
	// 若想使用发起审核，请指定该字段为：CreateReview
	// 若发起个人审核，则指定该字段为：SignReview
	OperateType *string `json:"OperateType,omitnil" name:"OperateType"`
}

type ChannelCreateFlowSignReviewRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 企业内部审核结果
	// PASS: 通过
	// REJECT: 拒绝
	// SIGN_REJECT:拒签(流程结束)
	ReviewType *string `json:"ReviewType,omitnil" name:"ReviewType"`

	// 审核原因 
	// 当ReviewType 是REJECT 时此字段必填,字符串长度不超过200
	ReviewMessage *string `json:"ReviewMessage,omitnil" name:"ReviewMessage"`

	// 签署节点审核时需要指定，给个人审核时必填。
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 操作类型，默认：SignReview；SignReview:签署审核，CreateReview：发起审核
	// 注：接口通过该字段区分操作类型
	// 该字段不传或者为空，则默认为SignReview签署审核，走签署审核流程
	// 若想使用发起审核，请指定该字段为：CreateReview
	// 若发起个人审核，则指定该字段为：SignReview
	OperateType *string `json:"OperateType,omitnil" name:"OperateType"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 流程签署人列表，其中结构体的Name，Mobile和ApproverType必传，其他可不传。
	// 注:
	// `1. ApproverType目前只支持个人(PERSON)类型的签署人。`
	// `2. 签署人只能有手写签名和时间类型的签署控件，其他类型的填写控件和签署控件暂时都未支持。`
	FlowApproverInfos []*FlowApproverInfo `json:"FlowApproverInfos,omitnil" name:"FlowApproverInfos"`

	// 用户信息，暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 机构信息，暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 签署完之后的H5页面的跳转链接，此链接及支持http://和https://，最大长度1000个字符。(建议https协议)
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 流程签署人列表，其中结构体的Name，Mobile和ApproverType必传，其他可不传。
	// 注:
	// `1. ApproverType目前只支持个人(PERSON)类型的签署人。`
	// `2. 签署人只能有手写签名和时间类型的签署控件，其他类型的填写控件和签署控件暂时都未支持。`
	FlowApproverInfos []*FlowApproverInfo `json:"FlowApproverInfos,omitnil" name:"FlowApproverInfos"`

	// 用户信息，暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 机构信息，暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 签署完之后的H5页面的跳转链接，此链接及支持http://和https://，最大长度1000个字符。(建议https协议)
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowSignUrlResponseParams struct {
	// 签署人签署链接信息
	FlowApproverUrlInfos []*FlowApproverUrlInfo `json:"FlowApproverUrlInfos,omitnil" name:"FlowApproverUrlInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同模板ID，为32位字符串。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 通过此二维码可发起的流程最大限额，如未明确指定，默认为5份。 一旦发起流程数超越该限制，该二维码将自动失效。	
	MaxFlowNum *int64 `json:"MaxFlowNum,omitnil" name:"MaxFlowNum"`

	// 合同流程的签署有效期限，若未设定签署截止日期，则默认为自合同流程创建起的7天内截止。 若在签署截止日期前未完成签署，合同状态将变更为已过期，从而导致合同无效。 最长设定期限不得超过30天。	
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitnil" name:"FlowEffectiveDay"`

	// 二维码的有效期限，默认为7天，最高设定不得超过90天。 一旦超过二维码的有效期限，该二维码将自动失效。	
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitnil" name:"QrEffectiveDay"`

	// 指定签署人信息。 在指定签署人后，仅允许特定签署人通过扫描二维码进行签署。	
	Restrictions []*ApproverRestriction `json:"Restrictions,omitnil" name:"Restrictions"`

	// 指定签署方经办人控件类型是个人印章签署控件（SIGN_SIGNATURE） 时，可选的签名方式。
	ApproverComponentLimitTypes []*ApproverComponentLimitType `json:"ApproverComponentLimitTypes,omitnil" name:"ApproverComponentLimitTypes"`

	// 已废弃，回调配置统一使用企业应用管理-应用集成-第三方应用中的配置
	// <br/> 通过一码多扫二维码发起的合同，回调消息可参考文档 https://qian.tencent.com/developers/partner/callback_types_contracts_sign
	// <br/> 用户通过签署二维码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档 https://qian.tencent.com/developers/partner/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 限制二维码用户条件（已弃用）
	//
	// Deprecated: ApproverRestrictions is deprecated.
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitnil" name:"ApproverRestrictions"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同模板ID，为32位字符串。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 通过此二维码可发起的流程最大限额，如未明确指定，默认为5份。 一旦发起流程数超越该限制，该二维码将自动失效。	
	MaxFlowNum *int64 `json:"MaxFlowNum,omitnil" name:"MaxFlowNum"`

	// 合同流程的签署有效期限，若未设定签署截止日期，则默认为自合同流程创建起的7天内截止。 若在签署截止日期前未完成签署，合同状态将变更为已过期，从而导致合同无效。 最长设定期限不得超过30天。	
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitnil" name:"FlowEffectiveDay"`

	// 二维码的有效期限，默认为7天，最高设定不得超过90天。 一旦超过二维码的有效期限，该二维码将自动失效。	
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitnil" name:"QrEffectiveDay"`

	// 指定签署人信息。 在指定签署人后，仅允许特定签署人通过扫描二维码进行签署。	
	Restrictions []*ApproverRestriction `json:"Restrictions,omitnil" name:"Restrictions"`

	// 指定签署方经办人控件类型是个人印章签署控件（SIGN_SIGNATURE） 时，可选的签名方式。
	ApproverComponentLimitTypes []*ApproverComponentLimitType `json:"ApproverComponentLimitTypes,omitnil" name:"ApproverComponentLimitTypes"`

	// 已废弃，回调配置统一使用企业应用管理-应用集成-第三方应用中的配置
	// <br/> 通过一码多扫二维码发起的合同，回调消息可参考文档 https://qian.tencent.com/developers/partner/callback_types_contracts_sign
	// <br/> 用户通过签署二维码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档 https://qian.tencent.com/developers/partner/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 限制二维码用户条件（已弃用）
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitnil" name:"ApproverRestrictions"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateMultiFlowSignQRCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateMultiFlowSignQRCodeResponseParams struct {
	// 签署二维码的基本信息，用于创建二维码，用户可扫描该二维码进行签署操作。	
	QrCode *SignQrCode `json:"QrCode,omitnil" name:"QrCode"`

	// 流程签署二维码的签署信息，适用于客户系统整合二维码功能。通过链接，用户可直接访问电子签名小程序并签署合同。	
	SignUrls *SignUrl `json:"SignUrls,omitnil" name:"SignUrls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 请指定需执行批量签署的流程ID，数量范围为1-100。 您可登录腾讯电子签控制台，浏览 "合同"->"合同中心" 以查阅某一合同的FlowId（在页面中显示为合同ID）。 用户将利用链接对这些合同实施批量操作。	
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 第三方应用平台的用户openid。 您可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查阅某位员工的OpenId。 OpenId必须是传入合同（FlowId）中的签署人。 - 1. 若OpenId为空，Name和Mobile 必须提供。 - 2. 若OpenId 与 Name，Mobile均存在，将优先采用OpenId对应的员工。	
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：`请确保和合同中填入的一致`
	Name *string `json:"Name,omitnil" name:"Name"`

	// 员工手机号，必须与姓名一起使用。 如果OpenId为空，则此字段不能为空。同时，姓名和手机号码必须与传入合同（FlowId）中的签署人信息一致。	
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`
}

type ChannelCreateOrganizationBatchSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括子客企业及应用编、号等详细内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 请指定需执行批量签署的流程ID，数量范围为1-100。 您可登录腾讯电子签控制台，浏览 "合同"->"合同中心" 以查阅某一合同的FlowId（在页面中显示为合同ID）。 用户将利用链接对这些合同实施批量操作。	
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 第三方应用平台的用户openid。 您可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查阅某位员工的OpenId。 OpenId必须是传入合同（FlowId）中的签署人。 - 1. 若OpenId为空，Name和Mobile 必须提供。 - 2. 若OpenId 与 Name，Mobile均存在，将优先采用OpenId对应的员工。	
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：`请确保和合同中填入的一致`
	Name *string `json:"Name,omitnil" name:"Name"`

	// 员工手机号，必须与姓名一起使用。 如果OpenId为空，则此字段不能为空。同时，姓名和手机号码必须与传入合同（FlowId）中的签署人信息一致。	
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateOrganizationBatchSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateOrganizationBatchSignUrlResponseParams struct {
	// 批量签署入口链接，用户可使用这个链接跳转到控制台页面对合同进行签署操作。	
	SignUrl *string `json:"SignUrl,omitnil" name:"SignUrl"`

	// 链接过期时间以 Unix 时间戳格式表示，从生成链接时间起，往后7天有效期。过期后短链将失效，无法打开。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
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
	QrCodeUrl *string `json:"QrCodeUrl,omitnil" name:"QrCodeUrl"`

	// 二维码失效时间 UNIX 时间戳 精确到秒
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type ChannelCreatePrepareFlowRequestParams struct {
	// 资源id，与ResourceType相对应，取值范围：
	// <ul>
	// <li>文件Id（通过UploadFiles获取文件资源Id）</li>
	// <li>模板Id</li>
	// </ul>
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源类型，取值有：
	// <ul><li> **1**：模板</li>
	// <li> **2**：文件（默认值）</li></ul>
	ResourceType *int64 `json:"ResourceType,omitnil" name:"ResourceType"`

	// 要创建的合同信息
	FlowInfo *BaseFlowInfo `json:"FlowInfo,omitnil" name:"FlowInfo"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程配置信息，用于配置发起合同时定制化如是否允许修改，某些按钮的隐藏等逻辑
	FlowOption *CreateFlowOption `json:"FlowOption,omitnil" name:"FlowOption"`

	// 合同签署人信息
	FlowApproverList []*CommonFlowApprover `json:"FlowApproverList,omitnil" name:"FlowApproverList"`

	// 合同Id：用于通过一个已发起的合同快速生成一个发起流程web链接
	// 注: `该参数必须是一个待发起审核的合同id，并且还未审核通过`
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 该参数不可用，请通过获取 web 可嵌入接口获取合同流程预览 URL
	//
	// Deprecated: NeedPreview is deprecated.
	NeedPreview *bool `json:"NeedPreview,omitnil" name:"NeedPreview"`

	// 企业机构信息，不用传
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 操作人（用户）信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCreatePrepareFlowRequest struct {
	*tchttp.BaseRequest
	
	// 资源id，与ResourceType相对应，取值范围：
	// <ul>
	// <li>文件Id（通过UploadFiles获取文件资源Id）</li>
	// <li>模板Id</li>
	// </ul>
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源类型，取值有：
	// <ul><li> **1**：模板</li>
	// <li> **2**：文件（默认值）</li></ul>
	ResourceType *int64 `json:"ResourceType,omitnil" name:"ResourceType"`

	// 要创建的合同信息
	FlowInfo *BaseFlowInfo `json:"FlowInfo,omitnil" name:"FlowInfo"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程配置信息，用于配置发起合同时定制化如是否允许修改，某些按钮的隐藏等逻辑
	FlowOption *CreateFlowOption `json:"FlowOption,omitnil" name:"FlowOption"`

	// 合同签署人信息
	FlowApproverList []*CommonFlowApprover `json:"FlowApproverList,omitnil" name:"FlowApproverList"`

	// 合同Id：用于通过一个已发起的合同快速生成一个发起流程web链接
	// 注: `该参数必须是一个待发起审核的合同id，并且还未审核通过`
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 该参数不可用，请通过获取 web 可嵌入接口获取合同流程预览 URL
	NeedPreview *bool `json:"NeedPreview,omitnil" name:"NeedPreview"`

	// 企业机构信息，不用传
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 操作人（用户）信息，不用传
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	delete(f, "ResourceId")
	delete(f, "ResourceType")
	delete(f, "FlowInfo")
	delete(f, "Agent")
	delete(f, "FlowOption")
	delete(f, "FlowApproverList")
	delete(f, "FlowId")
	delete(f, "NeedPreview")
	delete(f, "Organization")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreatePrepareFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreatePrepareFlowResponseParams struct {
	// 发起的合同嵌入链接， 可以直接点击进入进行合同发起， 有效期为5分钟
	PrepareFlowUrl *string `json:"PrepareFlowUrl,omitnil" name:"PrepareFlowUrl"`

	// 合同发起后预览链接， 注意此时合同并未发起，仅只是展示效果， 有效期为5分钟
	PreviewFlowUrl *string `json:"PreviewFlowUrl,omitnil" name:"PreviewFlowUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 个人用户姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 证件号码, 应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母（但“I”、“O”除外），后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 电子印章名字，1-50个中文字符
	// 注:`同一企业下电子印章名字不能相同`
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 电子印章图片base64编码，大小不超过10M（原始图片不超过5M），只支持PNG或JPG图片格式。
	// 
	SealImage *string `json:"SealImage,omitnil" name:"SealImage"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li>
	// <li>OTHER_CARD_TYPE : 其他</li></ul>
	// 
	// 注: `其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 是否开启印章图片压缩处理，默认不开启，如需开启请设置为 true。当印章超过 2M 时建议开启，开启后图片的 hash 将发生变化。
	SealImageCompress *bool `json:"SealImageCompress,omitnil" name:"SealImageCompress"`

	// 手机号码；当需要开通自动签时，该参数必传
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 是否开通自动签，该功能需联系运营工作人员开通后使用
	EnableAutoSign *bool `json:"EnableAutoSign,omitnil" name:"EnableAutoSign"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。一旦绑定后，将扣减购买的个人自动签账号许可一次（1年有效期），不可解绑释放。不传默认为绑定自动签账号许可。 0-绑定个人自动签账号许可，开通后将扣减购买的个人自动签账号许可一次 1-不绑定，发起合同时将按标准合同套餐进行扣减	
	LicenseType *int64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`
}

type ChannelCreatePreparedPersonalEsignRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 个人用户姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 证件号码, 应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母（但“I”、“O”除外），后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 电子印章名字，1-50个中文字符
	// 注:`同一企业下电子印章名字不能相同`
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 电子印章图片base64编码，大小不超过10M（原始图片不超过5M），只支持PNG或JPG图片格式。
	// 
	SealImage *string `json:"SealImage,omitnil" name:"SealImage"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li>
	// <li>OTHER_CARD_TYPE : 其他</li></ul>
	// 
	// 注: `其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 是否开启印章图片压缩处理，默认不开启，如需开启请设置为 true。当印章超过 2M 时建议开启，开启后图片的 hash 将发生变化。
	SealImageCompress *bool `json:"SealImageCompress,omitnil" name:"SealImageCompress"`

	// 手机号码；当需要开通自动签时，该参数必传
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 是否开通自动签，该功能需联系运营工作人员开通后使用
	EnableAutoSign *bool `json:"EnableAutoSign,omitnil" name:"EnableAutoSign"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。一旦绑定后，将扣减购买的个人自动签账号许可一次（1年有效期），不可解绑释放。不传默认为绑定自动签账号许可。 0-绑定个人自动签账号许可，开通后将扣减购买的个人自动签账号许可一次 1-不绑定，发起合同时将按标准合同套餐进行扣减	
	LicenseType *int64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`
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
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 待解除的签署流程编号(即原签署流程的编号)。
	NeedRelievedFlowId *string `json:"NeedRelievedFlowId,omitnil" name:"NeedRelievedFlowId"`

	// 解除协议内容, 包括解除理由等信息。
	ReliveInfo *RelieveInfo `json:"ReliveInfo,omitnil" name:"ReliveInfo"`

	// 指定解除协议的签署人，如不指定，则默认使用原流程的签署人。 <br/>
	// 如需更换原合同中的企业端签署人，可通过指定该签署人在原合同列表中的ApproverNumber编号来更换此企业端签署人。(可通过接口<a href="https://qian.tencent.com/developers/partnerApis/flows/DescribeFlowDetailInfo/">DescribeFlowDetailInfo</a>查询签署人的ApproverNumber编号，默认从0开始，顺序递增)<br/>
	// 
	// 注意：
	// <ul>
	// <li>只能更换自己企业的签署人，不支持更换个人类型或者其他企业的签署人</li>
	// <li>可以不指定替换签署人，使用原流程的签署人</li>
	// </ul>
	ReleasedApprovers []*ReleasedApprover `json:"ReleasedApprovers,omitnil" name:"ReleasedApprovers"`

	// 签署完回调url，最大长度1000个字符
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳(秒)，如果未设置签署截止时间，则默认为合同流程创建后的7天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 调用方自定义的个性化字段，该字段的值可以是字符串JSON或其他字符串形式，客户可以根据自身需求自定义数据格式并在需要时进行解析。该字段的信息将以Base64编码的形式传输，支持的最大数据大小为20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil" name:"UserData"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 待解除的签署流程编号(即原签署流程的编号)。
	NeedRelievedFlowId *string `json:"NeedRelievedFlowId,omitnil" name:"NeedRelievedFlowId"`

	// 解除协议内容, 包括解除理由等信息。
	ReliveInfo *RelieveInfo `json:"ReliveInfo,omitnil" name:"ReliveInfo"`

	// 指定解除协议的签署人，如不指定，则默认使用原流程的签署人。 <br/>
	// 如需更换原合同中的企业端签署人，可通过指定该签署人在原合同列表中的ApproverNumber编号来更换此企业端签署人。(可通过接口<a href="https://qian.tencent.com/developers/partnerApis/flows/DescribeFlowDetailInfo/">DescribeFlowDetailInfo</a>查询签署人的ApproverNumber编号，默认从0开始，顺序递增)<br/>
	// 
	// 注意：
	// <ul>
	// <li>只能更换自己企业的签署人，不支持更换个人类型或者其他企业的签署人</li>
	// <li>可以不指定替换签署人，使用原流程的签署人</li>
	// </ul>
	ReleasedApprovers []*ReleasedApprover `json:"ReleasedApprovers,omitnil" name:"ReleasedApprovers"`

	// 签署完回调url，最大长度1000个字符
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳(秒)，如果未设置签署截止时间，则默认为合同流程创建后的7天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 调用方自定义的个性化字段，该字段的值可以是字符串JSON或其他字符串形式，客户可以根据自身需求自定义数据格式并在需要时进行解析。该字段的信息将以Base64编码的形式传输，支持的最大数据大小为20480长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"UserData,omitnil" name:"UserData"`
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
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil" name:"Description"`

	// 权限树，权限树内容 PermissionGroups 可参考接口 ChannelDescribeRoles 的输出
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil" name:"PermissionGroups"`
}

type ChannelCreateRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil" name:"Description"`

	// 权限树，权限树内容 PermissionGroups 可参考接口 ChannelDescribeRoles 的输出
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil" name:"PermissionGroups"`
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
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 电子印章ID，为32位字符串。
	// 建议开发者保留此印章ID，后续指定签署区印章或者操作印章需此印章ID。
	// 可登录腾讯电子签控制台，在 "印章"->"印章中心"选择查看的印章，在"印章详情" 中查看某个印章的SealId(在页面中展示为印章ID)。
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 
	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查看某位员工的UserId(在页面中展示为用户ID)。
	// 员工在贵司业务系统中的唯一身份标识，用于与腾讯电子签账号进行映射，确保在同一企业内不会出现重复。
	// 该标识最大长度为64位字符串，仅支持包含26个英文字母和数字0-9的字符。
	// 指定待授权的用户ID数组,电子签的用户ID
	// 可以填写OpenId，系统会通过组织+渠道+OpenId查询得到UserId进行授权。
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 操作人（用户）信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 企业机构信息，不用传
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 电子印章ID，为32位字符串。
	// 建议开发者保留此印章ID，后续指定签署区印章或者操作印章需此印章ID。
	// 可登录腾讯电子签控制台，在 "印章"->"印章中心"选择查看的印章，在"印章详情" 中查看某个印章的SealId(在页面中展示为印章ID)。
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 
	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查看某位员工的UserId(在页面中展示为用户ID)。
	// 员工在贵司业务系统中的唯一身份标识，用于与腾讯电子签账号进行映射，确保在同一企业内不会出现重复。
	// 该标识最大长度为64位字符串，仅支持包含26个英文字母和数字0-9的字符。
	// 指定待授权的用户ID数组,电子签的用户ID
	// 可以填写OpenId，系统会通过组织+渠道+OpenId查询得到UserId进行授权。
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 操作人（用户）信息，不用传
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 企业机构信息，不用传
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`
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
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签开通配置信息, 包括开通的人员的信息等
	AutoSignConfig *AutoSignConfig `json:"AutoSignConfig,omitnil" name:"AutoSignConfig"`

	// 生成的链接类型：
	// <ul><li> 不传(即为空值) 则会生成小程序端开通链接(默认)</li>
	// <li> **H5SIGN** : 生成H5端开通链接</li></ul>
	UrlType *string `json:"UrlType,omitnil" name:"UrlType"`

	// 是否通知开通方，通知类型:
	// <ul><li>默认不设置为不通知开通方</li>
	// <li>**SMS** :  短信通知 ,如果需要短信通知则NotifyAddress填写对方的手机号</li></ul>
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 如果通知类型NotifyType选择为SMS，则此处为手机号, 其他通知类型不需要设置此项
	NotifyAddress *string `json:"NotifyAddress,omitnil" name:"NotifyAddress"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为当前时间往后30天。`如果不传，默认过期时间为当前时间往后7天。`
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`
}

type ChannelCreateUserAutoSignEnableUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签开通配置信息, 包括开通的人员的信息等
	AutoSignConfig *AutoSignConfig `json:"AutoSignConfig,omitnil" name:"AutoSignConfig"`

	// 生成的链接类型：
	// <ul><li> 不传(即为空值) 则会生成小程序端开通链接(默认)</li>
	// <li> **H5SIGN** : 生成H5端开通链接</li></ul>
	UrlType *string `json:"UrlType,omitnil" name:"UrlType"`

	// 是否通知开通方，通知类型:
	// <ul><li>默认不设置为不通知开通方</li>
	// <li>**SMS** :  短信通知 ,如果需要短信通知则NotifyAddress填写对方的手机号</li></ul>
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 如果通知类型NotifyType选择为SMS，则此处为手机号, 其他通知类型不需要设置此项
	NotifyAddress *string `json:"NotifyAddress,omitnil" name:"NotifyAddress"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为当前时间往后30天。`如果不传，默认过期时间为当前时间往后7天。`
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateUserAutoSignEnableUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateUserAutoSignEnableUrlResponseParams struct {
	// 个人用户自动签的开通链接, 短链形式。过期时间受 `ExpiredTime` 参数控制。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 腾讯电子签小程序的 AppID，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用
	// 
	// 注: `如果获取的是H5链接, 则不会返回此值`
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 腾讯电子签小程序的原始 Id,  ，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用
	// 
	// 注: `如果获取的是H5链接, 则不会返回此值`
	AppOriginalId *string `json:"AppOriginalId,omitnil" name:"AppOriginalId"`

	// 腾讯电子签小程序的跳转路径，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用
	// 
	// 注: `如果获取的是H5链接, 则不会返回此值`
	Path *string `json:"Path,omitnil" name:"Path"`

	// base64 格式的跳转二维码图片，可通过微信扫描后跳转到腾讯电子签小程序的开通界面。
	// 
	// 注: `如果获取的是H5链接, 则不会返回此二维码图片`
	QrCode *string `json:"QrCode,omitnil" name:"QrCode"`

	// 返回的链接类型
	// <ul><li> 空: 默认小程序端链接</li>
	// <li> **H5SIGN** : h5端链接</li></ul>
	UrlType *string `json:"UrlType,omitnil" name:"UrlType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 自动签开通个人用户信息，包括名字，身份证等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为当前时间往后30天。`如果不传，默认过期时间为当前时间往后7天。`
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`
}

type ChannelCreateUserAutoSignSealUrlRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 自动签开通个人用户信息，包括名字，身份证等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为当前时间往后30天。`如果不传，默认过期时间为当前时间往后7天。`
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`
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
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 腾讯电子签小程序的原始Id，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用。
	AppOriginalId *string `json:"AppOriginalId,omitnil" name:"AppOriginalId"`

	// 个人用户自动签的开通链接, 短链形式。过期时间受 `ExpiredTime` 参数控制。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 腾讯电子签小程序的跳转路径，用于其他小程序/APP等应用跳转至腾讯电子签小程序使用。
	Path *string `json:"Path,omitnil" name:"Path"`

	// base64格式的跳转二维码图片，可通过微信扫描后跳转到腾讯电子签小程序的开通界面。
	QrCode *string `json:"QrCode,omitnil" name:"QrCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 绑定角色的角色id列表，最多 100 个
	RoleIds []*string `json:"RoleIds,omitnil" name:"RoleIds"`

	// 电子签用户ID列表，与OpenIds参数二选一,优先UserIds参数，最多 100 个
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 客户系统用户ID列表，与UserIds参数二选一,优先UserIds参数，最多 100 个
	OpenIds []*string `json:"OpenIds,omitnil" name:"OpenIds"`

	// 操作者信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCreateUserRolesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 绑定角色的角色id列表，最多 100 个
	RoleIds []*string `json:"RoleIds,omitnil" name:"RoleIds"`

	// 电子签用户ID列表，与OpenIds参数二选一,优先UserIds参数，最多 100 个
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 客户系统用户ID列表，与UserIds参数二选一,优先UserIds参数，最多 100 个
	OpenIds []*string `json:"OpenIds,omitnil" name:"OpenIds"`

	// 操作者信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	FailedCreateRoleData []*FailedCreateRoleData `json:"FailedCreateRoleData,omitnil" name:"FailedCreateRoleData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 主题类型<br/>EMBED_WEB_THEME：嵌入式主题
	// <br/>目前只支持EMBED_WEB_THEME，web页面嵌入的主题风格配置
	ThemeType *string `json:"ThemeType,omitnil" name:"ThemeType"`

	// 主题配置
	WebThemeConfig *WebThemeConfig `json:"WebThemeConfig,omitnil" name:"WebThemeConfig"`
}

type ChannelCreateWebThemeConfigRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 主题类型<br/>EMBED_WEB_THEME：嵌入式主题
	// <br/>目前只支持EMBED_WEB_THEME，web页面嵌入的主题风格配置
	ThemeType *string `json:"ThemeType,omitnil" name:"ThemeType"`

	// 主题配置
	WebThemeConfig *WebThemeConfig `json:"WebThemeConfig,omitnil" name:"WebThemeConfig"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 角色id，最多20个
	RoleIds []*string `json:"RoleIds,omitnil" name:"RoleIds"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 角色id，最多20个
	RoleIds []*string `json:"RoleIds,omitnil" name:"RoleIds"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 角色Id（非超管或法人角色Id）
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 电子签用户ID列表，与OpenIds参数二选一,优先UserIds参数，最多两百
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 操作人信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 客户系统用户ID列表，与UserIds参数二选一,优先UserIds参数，最多两百
	OpenIds []*string `json:"OpenIds,omitnil" name:"OpenIds"`
}

type ChannelDeleteRoleUsersRequest struct {
	*tchttp.BaseRequest
	
	// 代理信息此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 角色Id（非超管或法人角色Id）
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 电子签用户ID列表，与OpenIds参数二选一,优先UserIds参数，最多两百
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 客户系统用户ID列表，与UserIds参数二选一,优先UserIds参数，最多两百
	OpenIds []*string `json:"OpenIds,omitnil" name:"OpenIds"`
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
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作的印章ID
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 需要删除授权的用户ID数组，可以传入电子签系统用户ID或OpenId。
	// 注: 
	// 1. `填写OpenId时，系统会通过组织+渠道+OpenId查询得到对应的UserId进行授权取消操作`
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 组织机构信息，不用传
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 操作人（用户）信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelDeleteSealPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作的印章ID
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 需要删除授权的用户ID数组，可以传入电子签系统用户ID或OpenId。
	// 注: 
	// 1. `填写OpenId时，系统会通过组织+渠道+OpenId查询得到对应的UserId进行授权取消操作`
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// 组织机构信息，不用传
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 操作人（用户）信息，不用传
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type ChannelDescribeBillUsageDetailRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询开始时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

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
	QuotaType *string `json:"QuotaType,omitnil" name:"QuotaType"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 指定分页每页返回的数据条数，如果不传默认为 50，单页最大支持 50。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type ChannelDescribeBillUsageDetailRequest struct {
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询开始时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

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
	QuotaType *string `json:"QuotaType,omitnil" name:"QuotaType"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 指定分页每页返回的数据条数，如果不传默认为 50，单页最大支持 50。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 消耗记录详情
	Details []*ChannelBillUsageDetail `json:"Details,omitnil" name:"Details"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定分页每页返回的数据条数，单页最大支持 20。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询的关键字段，支持Key-Values查询。可选键值如下：
	// <ul>
	//   <li>Key:**"Status"**，Values: **["IsVerified"]**, 查询已实名的员工</li>
	//   <li>Key:**"Status"**，Values: **["QuiteJob"]**, 查询离职员工</li>
	//   <li>Key:**"StaffOpenId"**，Values: **["OpenId1","OpenId2",...]**, 根据第三方系统用户OpenId查询员工</li>
	// </ul>
	// 注: `同名字的Key的过滤条件会冲突,  只能填写一个`
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页。
	// 页码从 0 开始，即首页为 0，最大20000。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定分页每页返回的数据条数，单页最大支持 20。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询的关键字段，支持Key-Values查询。可选键值如下：
	// <ul>
	//   <li>Key:**"Status"**，Values: **["IsVerified"]**, 查询已实名的员工</li>
	//   <li>Key:**"Status"**，Values: **["QuiteJob"]**, 查询离职员工</li>
	//   <li>Key:**"StaffOpenId"**，Values: **["OpenId1","OpenId2",...]**, 根据第三方系统用户OpenId查询员工</li>
	// </ul>
	// 注: `同名字的Key的过滤条件会冲突,  只能填写一个`
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页。
	// 页码从 0 开始，即首页为 0，最大20000。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Employees []*Staff `json:"Employees,omitnil" name:"Employees"`

	// 指定分页返回第几页的数据。页码从 0 开始，即首页为 0，最大20000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 指定分页每页返回的数据条数，单页最大支持 20。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 符合条件的员工数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 需要获取填写控件填写内容的合同流程ID
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 需要获取填写控件填写内容的合同流程ID
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`
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
	RecipientComponentInfos []*RecipientComponentInfo `json:"RecipientComponentInfos,omitnil" name:"RecipientComponentInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 返回最大数量，最大为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页查询偏移量，默认为0，最大为20000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询信息类型
	// 支持的值如下：
	// <ul><li>0-默认，不返回授权用户信息</li>
	// <li>1-返回授权用户信息</li>
	// </ul>
	InfoType *int64 `json:"InfoType,omitnil" name:"InfoType"`

	// 印章id（没有输入返回所有）
	// 
	// 注:  `没有输入返回所有记录，最大返回100条。`
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 印章类型列表，目前支持传入以下类型：
	// <ul><li>OFFICIAL-企业公章</li>
	// <li>CONTRACT-合同专用章</li>
	// <li>ORGANIZATION_SEAL-企业印章(图片上传创建)</li>
	// <li>LEGAL_PERSON_SEAL-法定代表人章</li>
	// </ul>
	// 
	// 注:  `为空时查询所有类型的印章。`
	SealTypes []*string `json:"SealTypes,omitnil" name:"SealTypes"`
}

type ChannelDescribeOrganizationSealsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 返回最大数量，最大为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页查询偏移量，默认为0，最大为20000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询信息类型
	// 支持的值如下：
	// <ul><li>0-默认，不返回授权用户信息</li>
	// <li>1-返回授权用户信息</li>
	// </ul>
	InfoType *int64 `json:"InfoType,omitnil" name:"InfoType"`

	// 印章id（没有输入返回所有）
	// 
	// 注:  `没有输入返回所有记录，最大返回100条。`
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 印章类型列表，目前支持传入以下类型：
	// <ul><li>OFFICIAL-企业公章</li>
	// <li>CONTRACT-合同专用章</li>
	// <li>ORGANIZATION_SEAL-企业印章(图片上传创建)</li>
	// <li>LEGAL_PERSON_SEAL-法定代表人章</li>
	// </ul>
	// 
	// 注:  `为空时查询所有类型的印章。`
	SealTypes []*string `json:"SealTypes,omitnil" name:"SealTypes"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelDescribeOrganizationSealsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelDescribeOrganizationSealsResponseParams struct {
	// 在设置了SealId时返回0或1，没有设置时返回公司的总印章数量，可能比返回的印章数组数量多
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 查询到的印章结果数组
	Seals []*OccupiedSeal `json:"Seals,omitnil" name:"Seals"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定每页返回的数据条数，和Offset参数配合使用，单页最大200。
	// 
	// 注: `因为历史原因, 此字段为字符串类型`
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 查询的关键字段:
	// Key:"**RoleType**",Values:["**1**"]查询系统角色，
	// Key:"**RoleType**",Values:["**2**"]查询自定义角色
	// Key:"**RoleStatus**",Values:["**1**"]查询启用角色
	// Key:"**RoleStatus**",Values:["**2**"]查询禁用角色
	// Key:"**IsReturnPermissionGroup**"，Values:["**0**"]表示接口不返回角色对应的权限树字段
	// Key:"**IsReturnPermissionGroup**"，Values:["**1**"]表示接口返回角色对应的权限树字段
	// 
	// 注: `同名字的Key的过滤条件会冲突, 只能填写一个`
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用，最大2000条。
	// 
	// 注：
	// 1.`offset从0开始，即第一页为0。`
	// 2.`默认从第一页返回。`
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 操作人信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定每页返回的数据条数，和Offset参数配合使用，单页最大200。
	// 
	// 注: `因为历史原因, 此字段为字符串类型`
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 查询的关键字段:
	// Key:"**RoleType**",Values:["**1**"]查询系统角色，
	// Key:"**RoleType**",Values:["**2**"]查询自定义角色
	// Key:"**RoleStatus**",Values:["**1**"]查询启用角色
	// Key:"**RoleStatus**",Values:["**2**"]查询禁用角色
	// Key:"**IsReturnPermissionGroup**"，Values:["**0**"]表示接口不返回角色对应的权限树字段
	// Key:"**IsReturnPermissionGroup**"，Values:["**1**"]表示接口返回角色对应的权限树字段
	// 
	// 注: `同名字的Key的过滤条件会冲突, 只能填写一个`
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用，最大2000条。
	// 
	// 注：
	// 1.`offset从0开始，即第一页为0。`
	// 2.`默认从第一页返回。`
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 指定每页返回的数据条数，和Offset参数配合使用，单页最大200。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询角色的总数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 查询的角色信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelRoles []*ChannelRole `json:"ChannelRoles,omitnil" name:"ChannelRoles"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type ChannelDescribeUserAutoSignStatusRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 要查询状态的用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelDescribeUserAutoSignStatusRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 要查询状态的用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	IsOpen *bool `json:"IsOpen,omitnil" name:"IsOpen"`

	// 自动签许可生效时间。当且仅当已通过许可开通自动签时有值。
	// 
	// 值为unix时间戳,单位为秒。
	LicenseFrom *int64 `json:"LicenseFrom,omitnil" name:"LicenseFrom"`

	// 自动签许可到期时间。当且仅当已通过许可开通自动签时有值。
	// 
	// 值为unix时间戳,单位为秒。
	LicenseTo *int64 `json:"LicenseTo,omitnil" name:"LicenseTo"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。
	// 
	// <ul><li>**0**: 使用个人自动签账号许可进行开通，个人自动签账号许可有效期1年，注: `不可解绑释放更换他人`</li></ul>
	LicenseType *int64 `json:"LicenseType,omitnil" name:"LicenseType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 需要关闭自动签的个人的信息，如姓名，证件信息等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelDisableUserAutoSignRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签使用的场景值, 可以选择的场景值如下:
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN** :  电子处方场景</li><li> **OTHER** :  通用场景</li></ul>
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 需要关闭自动签的个人的信息，如姓名，证件信息等。
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 执行本接口操作的员工信息。
	// 注: `在调用此接口时，请确保指定的员工已获得所需的接口调用权限，并具备接口传入的相应资源的数据权限。`
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 转换任务Id，通过接口<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelCreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>得到的转换任务id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 转换任务Id，通过接口<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelCreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>得到的转换任务id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`
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
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务状态，需要关注的状态
	// <ul><li>**0**  :NeedTranform   - 任务已提交</li>
	// <li>**4**  :Processing     - 文档转换中</li>
	// <li>**8**  :TaskEnd        - 任务处理完成</li>
	// <li>**-2** :DownloadFailed - 下载失败</li>
	// <li>**-6** :ProcessFailed  - 转换失败</li>
	// <li>**-13**:ProcessTimeout - 转换文件超时</li></ul>
	TaskStatus *int64 `json:"TaskStatus,omitnil" name:"TaskStatus"`

	// 状态描述，需要关注的状态
	// <ul><li> **NeedTranform** : 任务已提交</li>
	// <li> **Processing** : 文档转换中</li>
	// <li> **TaskEnd** : 任务处理完成</li>
	// <li> **DownloadFailed** : 下载失败</li>
	// <li> **ProcessFailed** : 转换失败</li>
	// <li> **ProcessTimeout** : 转换文件超时</li></ul>
	TaskMessage *string `json:"TaskMessage,omitnil" name:"TaskMessage"`

	// 资源Id，也是FileId，用于文件发起时使用
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 预览文件Url，有效期30分钟 
	// 当前字段返回为空，发起的时候，将ResourceId 放入发起即可
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: PreviewUrl is deprecated.
	PreviewUrl *string `json:"PreviewUrl,omitnil" name:"PreviewUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 角色Id，可通过接口 ChannelDescribeRoles 查询获取
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil" name:"Description"`

	// 权限树，权限树内容 PermissionGroups 可参考接口 ChannelDescribeRoles的输出
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil" name:"PermissionGroups"`
}

type ChannelModifyRoleRequest struct {
	*tchttp.BaseRequest
	
	// 代理企业和员工的信息。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 角色Id，可通过接口 ChannelDescribeRoles 查询获取
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil" name:"Description"`

	// 权限树，权限树内容 PermissionGroups 可参考接口 ChannelDescribeRoles的输出
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil" name:"PermissionGroups"`
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
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 第三方平台子客企业的唯一标识
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 第三方平台子客企业名称
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 企业的统一社会信用代码
	UnifiedSocialCreditCode *string `json:"UnifiedSocialCreditCode,omitnil" name:"UnifiedSocialCreditCode"`

	// 企业法定代表人的姓名
	LegalName *string `json:"LegalName,omitnil" name:"LegalName"`

	// 企业法定代表人作为第三方平台子客企业员工的唯一标识
	LegalOpenId *string `json:"LegalOpenId,omitnil" name:"LegalOpenId"`

	// 企业超级管理员的姓名
	AdminName *string `json:"AdminName,omitnil" name:"AdminName"`

	// 企业超级管理员作为第三方平台子客企业员工的唯一标识
	AdminOpenId *string `json:"AdminOpenId,omitnil" name:"AdminOpenId"`

	// 企业超级管理员的手机号码
	// **注**：`手机号码脱敏（隐藏部分用*替代）`
	AdminMobile *string `json:"AdminMobile,omitnil" name:"AdminMobile"`

	// 企业认证状态字段。值如下：
	// <ul>
	//   <li>**"UNVERIFIED"**： 未认证的企业</li>
	//   <li>**"VERIFYINGLEGALPENDINGAUTHORIZATION"**： 认证中待法人授权的企业</li>
	//   <li>**"VERIFYINGAUTHORIZATIONFILEPENDING"**： 认证中授权书审核中的企业</li>
	//   <li>**"VERIFYINGAUTHORIZATIONFILEREJECT"**： 认证中授权书已驳回的企业</li>
	//   <li>**"VERIFYING"**： 认证中的企业</li>
	//   <li>**"VERIFIED"**： 已认证的企业</li>
	// </ul>
	AuthorizationStatus *string `json:"AuthorizationStatus,omitnil" name:"AuthorizationStatus"`

	// 企业认证方式字段。值如下：
	// <ul>
	//   <li>**"AuthorizationInit"**： 暂未选择授权方式</li>
	//   <li>**"AuthorizationFile"**： 授权书</li>
	//   <li>**"AuthorizationLegalPerson"**： 法人授权超管</li>
	//   <li>**"AuthorizationLegalIdentity"**： 法人直接认证</li>
	// </ul>
	AuthorizationType *string `json:"AuthorizationType,omitnil" name:"AuthorizationType"`
}

type ChannelRole struct {
	// 角色ID,为32位字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 角色的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// 此角色状态
	// 1: 已经启用
	// 2: 已经禁用
	RoleStatus *uint64 `json:"RoleStatus,omitnil" name:"RoleStatus"`

	// 此角色对应的权限列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil" name:"PermissionGroups"`
}

// Predefined struct for user
type ChannelUpdateSealStatusRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 印章状态，目前支持传入以下类型：
	// <ul><li>DISABLE-停用印章</li></ul>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 印章ID
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 更新印章状态原因说明
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 操作者的信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelUpdateSealStatusRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 印章状态，目前支持传入以下类型：
	// <ul><li>DISABLE-停用印章</li></ul>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 印章ID
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 更新印章状态原因说明
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelVerifyPdfRequest struct {
	*tchttp.BaseRequest
	
	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	VerifyResult *int64 `json:"VerifyResult,omitnil" name:"VerifyResult"`

	// 验签结果详情，所有签署区(包括签名区, 印章区, 日期签署区,骑缝章等)的签署验签结果
	PdfVerifyResults []*PdfVerifyResult `json:"PdfVerifyResults,omitnil" name:"PdfVerifyResults"`

	// 验签序列号, 为11为数组组成的字符串
	VerifySerialNo *string `json:"VerifySerialNo,omitnil" name:"VerifySerialNo"`

	// 合同文件MD5哈希值
	PdfResourceMd5 *string `json:"PdfResourceMd5,omitnil" name:"PdfResourceMd5"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	CanEditApprover *bool `json:"CanEditApprover,omitnil" name:"CanEditApprover"`
}

type CommonFlowApprover struct {
	// 指定签署人非第三方平台子客企业下员工还是SaaS平台企业，在ApproverType为ORGANIZATION时指定。
	// <ul>
	// <li>false: 默认值，第三方平台子客企业下员工</li>
	// <li>true: SaaS平台企业下的员工</li>
	// </ul>
	NotChannelOrganization *bool `json:"NotChannelOrganization,omitnil" name:"NotChannelOrganization"`

	// 在指定签署方时，可选择企业B端或个人C端等不同的参与者类型，可选类型如下:
	// 
	//  **0** :企业/企业员工（企业签署方或模板发起时的企业静默签）
	//  **1** :个人/自然人
	// **3** :企业/企业员工自动签（他方企业自动签署或文件发起时的本方企业自动签）
	// 
	// 注：类型为3（企业/企业员工自动签）时，此接口会默认完成该签署方的签署。静默签署仅进行盖章操作，不能自动签名。
	ApproverType *int64 `json:"ApproverType,omitnil" name:"ApproverType"`

	// 电子签平台给企业生成的企业id
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 企业OpenId，第三方应用集成非静默签子客企业签署人发起合同必传
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 企业名称，第三方应用集成非静默签子客企业签署人必传，saas企业签署人必传
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 电子签平台给企业员工或者自热人生成的用户id
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 第三方平台子客企业员工的唯一标识
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	ApproverName *string `json:"ApproverName,omitnil" name:"ApproverName"`

	// 签署人手机号，saas企业签署人，个人签署人必传
	ApproverMobile *string `json:"ApproverMobile,omitnil" name:"ApproverMobile"`

	// 签署方经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li>
	// <li>OTHER_CARD_TYPE : 其他证件</li></ul>
	// 
	// 注: `其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	ApproverIdCardType *string `json:"ApproverIdCardType,omitnil" name:"ApproverIdCardType"`

	// 签署方经办人的证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母（但“I”、“O”除外），后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	ApproverIdCardNumber *string `json:"ApproverIdCardNumber,omitnil" name:"ApproverIdCardNumber"`

	// 签署人Id，使用模板发起是，对应模板配置中的签署人RecipientId
	// 注意：模板发起时该字段必填
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 签署前置条件：阅读时长限制，不传默认10s,最大300s，最小3s
	PreReadTime *int64 `json:"PreReadTime,omitnil" name:"PreReadTime"`

	// 签署前置条件：阅读全文限制
	IsFullText *bool `json:"IsFullText,omitnil" name:"IsFullText"`

	// 通知签署方经办人的方式, 有以下途径:
	// <ul><li> **SMS** :(默认)短信</li>
	// <li> **NONE** : 不通知</li></ul>
	// 
	// 注: `签署方为第三方子客企业时会被置为NONE,   不会发短信通知`
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 签署人配置
	ApproverOption *CommonApproverOption `json:"ApproverOption,omitnil" name:"ApproverOption"`

	// 使用PDF文件直接发起合同时，签署人指定的签署控件；<br/>使用模板发起合同时，指定本企业印章签署控件的印章ID: <br/>通过ComponentId或ComponenetName指定签署控件，ComponentValue为印章ID。
	SignComponents []*Component `json:"SignComponents,omitnil" name:"SignComponents"`

	// 指定个人签署方查看合同的校验方式,可以传值如下:
	// <ul><li>  **1**   : （默认）人脸识别,人脸识别后才能合同内容</li>
	// <li>  **2**  : 手机号验证, 用户手机号和参与方手机号(ApproverMobile)相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// </li></ul>
	// 注: 
	// <ul><li>如果合同流程设置ApproverVerifyType查看合同的校验方式,    则忽略此签署人的查看合同的校验方式</li>
	// <li>此字段可传多个校验方式</li></ul>
	ApproverVerifyTypes []*int64 `json:"ApproverVerifyTypes,omitnil" name:"ApproverVerifyTypes"`

	// 签署人签署合同时的认证方式
	// <ul><li> **1** :人脸认证</li>
	// <li> **2** :签署密码</li>
	// <li> **3** :运营商三要素</li></ul>
	// 
	// 默认为1(人脸认证 ),2(签署密码)
	// 
	// 注: 
	// 1. 用<font color='red'>模板创建合同场景</font>, 签署人的认证方式需要在配置模板的时候指定, <font color='red'>在创建合同重新指定无效</font>
	// 2. 运营商三要素认证方式对手机号运营商及前缀有限制,可以参考[运营商支持列表类](https://qian.tencent.com/developers/partner/mobile_support)得到具体的支持说明
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil" name:"ApproverSignTypes"`
}

type Component struct {
	// 控件编号
	// 
	// CreateFlowByTemplates发起合同时优先以ComponentId（不为空）填充；否则以ComponentName填充
	// 
	// 注：
	// 当GenerateMode=KEYWORD时，通过"^"来决定是否使用关键字整词匹配能力。
	// 例：当GenerateMode=KEYWORD时，如果传入关键字"^甲方签署^"，则会在PDF文件中有且仅有"甲方签署"关键字的地方进行对应操作。
	// 如传入的关键字为"甲方签署"，则PDF文件中每个出现关键字的位置都会执行相应操作。
	// 
	// 创建控件时，此值为空
	// 查询时返回完整结构
	ComponentId *string `json:"ComponentId,omitnil" name:"ComponentId"`

	// 如果是Component控件类型，则可选的字段为：
	// TEXT - 普通文本控件，输入文本字符串；
	// MULTI_LINE_TEXT - 多行文本控件，输入文本字符串；
	// CHECK_BOX - 勾选框控件，若选中填写ComponentValue 填写 true或者 false 字符串；
	// FILL_IMAGE - 图片控件，ComponentValue 填写图片的资源 ID；
	// DYNAMIC_TABLE - 动态表格控件；
	// ATTACHMENT - 附件控件,ComponentValue 填写附件图片的资源 ID列表，以逗号分割；
	// SELECTOR - 选择器控件，ComponentValue填写选择的字符串内容；
	// DATE - 日期控件；默认是格式化为xxxx年xx月xx日字符串；
	// DISTRICT - 省市区行政区控件，ComponentValue填写省市区行政区字符串内容；
	// 
	// 如果是SignComponent控件类型，则可选的字段为
	// SIGN_SEAL - 签署印章控件；
	// SIGN_DATE - 签署日期控件；
	// SIGN_SIGNATURE - 用户签名控件；
	// SIGN_PERSONAL_SEAL - 个人签署印章控件（使用文件发起暂不支持此类型）；
	// SIGN_PAGING_SEAL - 骑缝章；若文件发起，需要对应填充ComponentPosY、ComponentWidth、ComponentHeight
	// SIGN_OPINION - 签署意见控件，用户需要根据配置的签署意见内容，完成对意见内容的确认;
	// SIGN_LEGAL_PERSON_SEAL - 企业法定代表人控件。
	// 
	// 表单域的控件不能作为印章和签名控件
	ComponentType *string `json:"ComponentType,omitnil" name:"ComponentType"`

	// 控件简称，不超过30个字符
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// 控件是否为必填项，
	// 默认为false-非必填
	ComponentRequired *bool `json:"ComponentRequired,omitnil" name:"ComponentRequired"`

	// 控件关联的参与方ID，对应Recipient结构体中的RecipientId	
	ComponentRecipientId *string `json:"ComponentRecipientId,omitnil" name:"ComponentRecipientId"`

	// 控件所属文件的序号 (文档中文件的排列序号，从0开始)
	FileIndex *int64 `json:"FileIndex,omitnil" name:"FileIndex"`

	// 控件生成的方式：
	// NORMAL - 普通控件
	// FIELD - 表单域
	// KEYWORD - 关键字（设置关键字时，请确保PDF原始文件内是关键字以文字形式保存在PDF文件中，不支持对图片内文字进行关键字查找）
	GenerateMode *string `json:"GenerateMode,omitnil" name:"GenerateMode"`

	// 参数控件宽度，默认100，单位px
	// 表单域和关键字转换控件不用填
	ComponentWidth *float64 `json:"ComponentWidth,omitnil" name:"ComponentWidth"`

	// 参数控件高度，默认100，单位px
	// 表单域和关键字转换控件不用填
	ComponentHeight *float64 `json:"ComponentHeight,omitnil" name:"ComponentHeight"`

	// 参数控件所在页码，从1开始
	ComponentPage *int64 `json:"ComponentPage,omitnil" name:"ComponentPage"`

	// 参数控件X位置，单位px
	ComponentPosX *float64 `json:"ComponentPosX,omitnil" name:"ComponentPosX"`

	// 参数控件Y位置，单位px
	ComponentPosY *float64 `json:"ComponentPosY,omitnil" name:"ComponentPosY"`

	// 扩展参数：
	// 为JSON格式。
	// 不同类型的控件会有部分非通用参数
	// 
	// ComponentType为TEXT、MULTI_LINE_TEXT时，支持以下参数：
	// 1 Font：目前只支持黑体、宋体
	// 2 FontSize： 范围12-72
	// 3 FontAlign： Left/Right/Center，左对齐/居中/右对齐
	// 4 FontColor：字符串类型，格式为RGB颜色数字
	// 参数样例：{\"FontColor\":\"255,0,0\",\"FontSize\":12}
	// 
	// ComponentType为FILL_IMAGE时，支持以下参数：
	// NotMakeImageCenter：bool。是否设置图片居中。false：居中（默认）。 true: 不居中
	// FillMethod: int. 填充方式。0-铺满（默认）；1-等比例缩放
	// 
	// ComponentType为SIGN_SIGNATURE类型可以控制签署方式
	// {“ComponentTypeLimit”: [“xxx”]}
	// xxx可以为：
	// HANDWRITE – 手写签名
	// OCR_ESIGN -- AI智能识别手写签名
	// ESIGN -- 个人印章类型
	// SYSTEM_ESIGN -- 系统签名（该类型可以在用户签署时根据用户姓名一键生成一个签名来进行签署）
	// 如：{“ComponentTypeLimit”: [“SYSTEM_ESIGN”]}
	// 
	// ComponentType为SIGN_DATE时，支持以下参数：
	// 1 Font：字符串类型目前只支持"黑体"、"宋体"，如果不填默认为"黑体"
	// 2 FontSize： 数字类型，范围6-72，默认值为12
	// 3 FontAlign： 字符串类型，可取Left/Right/Center，对应左对齐/居中/右对齐
	// 4 Format： 字符串类型，日期格式，必须是以下五种之一 “yyyy m d”，”yyyy年m月d日”，”yyyy/m/d”，”yyyy-m-d”，”yyyy.m.d”。
	// 5 Gaps:： 字符串类型，仅在Format为“yyyy m d”时起作用，格式为用逗号分开的两个整数，例如”2,2”，两个数字分别是日期格式的前后两个空隙中的空格个数
	// 如果extra参数为空，默认为”yyyy年m月d日”格式的居中日期
	// 特别地，如果extra中Format字段为空或无法被识别，则extra参数会被当作默认值处理（Font，FontSize，Gaps和FontAlign都不会起效）
	// 参数样例： "ComponentExtra": "{"Format":“yyyy m d”,"FontSize":12,"Gaps":"2,2", "FontAlign":"Right"}"
	// 
	// ComponentType为SIGN_SEAL类型时，支持以下参数：
	// 1.PageRanges：PageRange的数组，通过PageRanges属性设置该印章在PDF所有页面上盖章（适用于标书在所有页面盖章的情况）
	// 参数样例： "ComponentExtra":"{"PageRanges":[{"BeginPage":1,"EndPage":-1}]}"
	ComponentExtra *string `json:"ComponentExtra,omitnil" name:"ComponentExtra"`

	// 控件填充vaule，ComponentType和传入值类型对应关系：
	// TEXT - 文本内容
	// MULTI_LINE_TEXT - 文本内容
	// CHECK_BOX - true/false
	// FILL_IMAGE、ATTACHMENT - 附件的FileId，需要通过UploadFiles接口上传获取
	// SELECTOR - 选项值
	// DATE - 默认是格式化为xxxx年xx月xx日
	// DYNAMIC_TABLE - 传入json格式的表格内容，具体见数据结构FlowInfo：https://cloud.tencent.com/document/api/1420/61525#FlowInfo
	// SIGN_SEAL - 印章ID
	// SIGN_PAGING_SEAL - 可以指定印章ID
	// 
	// 控件值约束说明：
	// 企业全称控件：
	//   约束：企业名称中文字符中文括号
	//   检查正则表达式：/^[\u3400-\u4dbf\u4e00-\u9fa5（）]+$/
	// 
	// 统一社会信用代码控件：
	//   检查正则表达式：/^[A-Z0-9]{1,18}$/
	// 
	// 法人名称控件：
	//   约束：最大50个字符，2到25个汉字或者1到50个字母
	//   检查正则表达式：/^([\u3400-\u4dbf\u4e00-\u9fa5.·]{2,25}|[a-zA-Z·,\s-]{1,50})$/
	// 
	// 签署意见控件：
	//   约束：签署意见最大长度为50字符
	// 
	// 签署人手机号控件：
	//   约束：国内手机号 13,14,15,16,17,18,19号段长度11位
	// 
	// 签署人身份证控件：
	//   约束：合法的身份证号码检查
	// 
	// 控件名称：
	//   约束：控件名称最大长度为20字符
	// 
	// 单行文本控件：
	//   约束：只允许输入中文，英文，数字，中英文标点符号
	// 
	// 多行文本控件：
	//   约束：只允许输入中文，英文，数字，中英文标点符号
	// 
	// 勾选框控件：
	//   约束：选择填字符串true，不选填字符串false
	// 
	// 选择器控件：
	//   约束：同单行文本控件约束，填写选择值中的字符串
	// 
	// 数字控件：
	//   约束：请输入有效的数字(可带小数点) 
	//   检查正则表达式：/^(-|\+)?\d+(\.\d+)?$/
	// 
	// 日期控件：
	//   约束：格式：yyyy年mm月dd日
	// 
	// 附件控件：
	//   约束：JPG或PNG图片，上传数量限制，1到6个，最大6个附件
	// 
	// 图片控件：
	//   约束：JPG或PNG图片，填写上传的图片资源ID
	// 
	// 邮箱控件：
	//   约束：请输入有效的邮箱地址, w3c标准
	//   检查正则表达式：/^([A-Za-z0-9_\-.!#$%&])+@([A-Za-z0-9_\-.])+\.([A-Za-z]{2,4})$/
	//   参考：https://emailregex.com/
	// 
	// 地址控件：
	//   同单行文本控件约束
	// 
	// 省市区控件：
	//   同单行文本控件约束
	// 
	// 性别控件：
	//   同单行文本控件约束，填写选择值中的字符串
	// 
	// 学历控件：
	//   同单行文本控件约束，填写选择值中的字符串
	ComponentValue *string `json:"ComponentValue,omitnil" name:"ComponentValue"`

	// 日期签署控件的字号，默认为 12
	// 
	// 签署区日期控件会转换成图片格式并带存证，需要通过字体决定图片大小
	ComponentDateFontSize *int64 `json:"ComponentDateFontSize,omitnil" name:"ComponentDateFontSize"`

	// 控件所属文档的Id, 模板相关接口为空值
	DocumentId *string `json:"DocumentId,omitnil" name:"DocumentId"`

	// 控件描述，不超过30个字符
	ComponentDescription *string `json:"ComponentDescription,omitnil" name:"ComponentDescription"`

	// 指定关键字时横坐标偏移量，单位pt
	OffsetX *float64 `json:"OffsetX,omitnil" name:"OffsetX"`

	// 指定关键字时纵坐标偏移量，单位pt
	OffsetY *float64 `json:"OffsetY,omitnil" name:"OffsetY"`

	// 平台企业控件ID。
	// 如果不为空，属于平台企业预设控件；
	ChannelComponentId *string `json:"ChannelComponentId,omitnil" name:"ChannelComponentId"`

	// 指定关键字排序规则，
	// Positive-正序，
	// Reverse-倒序。
	// 传入Positive时会根据关键字在PDF文件内的顺序进行排列。在指定KeywordIndexes时，0代表在PDF内查找内容时，查找到的第一个关键字。
	// 传入Reverse时会根据关键字在PDF文件内的反序进行排列。在指定KeywordIndexes时，0代表在PDF内查找内容时，查找到的最后一个关键字。
	KeywordOrder *string `json:"KeywordOrder,omitnil" name:"KeywordOrder"`

	// 指定关键字页码。
	// 指定页码后，将只在指定的页码内查找关键字，非该页码的关键字将不会查询出来
	KeywordPage *int64 `json:"KeywordPage,omitnil" name:"KeywordPage"`

	// 关键字位置模式，
	// Middle-居中，
	// Below-正下方，
	// Right-正右方，
	// LowerRight-右上角，
	// UpperRight-右下角。
	// 示例：如果设置Middle的关键字盖章，则印章的中心会和关键字的中心重合，如果设置Below，则印章在关键字的正下方
	RelativeLocation *string `json:"RelativeLocation,omitnil" name:"RelativeLocation"`

	// 关键字索引，如果一个关键字在PDF文件中存在多个，可以通过关键字索引指定使用第几个关键字作为最后的结果，可指定多个索引。
	// 示例[0,2]，说明使用PDF文件内第1个和第3个关键字位置。
	KeywordIndexes []*int64 `json:"KeywordIndexes,omitnil" name:"KeywordIndexes"`

	// 填写提示的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Placeholder *string `json:"Placeholder,omitnil" name:"Placeholder"`

	// 是否锁定控件值不允许编辑（嵌入式发起使用） <br/>默认false：不锁定控件值，允许在页面编辑控件值	
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockComponentValue *bool `json:"LockComponentValue,omitnil" name:"LockComponentValue"`

	// 是否禁止移动和删除控件 <br/>默认false，不禁止移动和删除控件	
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForbidMoveAndDelete *bool `json:"ForbidMoveAndDelete,omitnil" name:"ForbidMoveAndDelete"`
}

type ComponentLimit struct {
	// 控件类型，支持以下类型
	// <ul><li>SIGN_SEAL : 印章控件</li>
	// <li>SIGN_PAGING_SEAL : 骑缝章控件</li>
	// <li>SIGN_LEGAL_PERSON_SEAL : 企业法定代表人控件</li>
	// <li>SIGN_SIGNATURE : 用户签名控件</li></ul>
	ComponentType *string `json:"ComponentType,omitnil" name:"ComponentType"`

	// 签署控件类型的值(可选)，用与限制签署时印章或者签名的选择范围
	// 
	// 1.当ComponentType 是 SIGN_SEAL 或者 SIGN_PAGING_SEAL 时可传入企业印章Id（支持多个）
	// 
	// 2.当ComponentType 是 SIGN_SIGNATURE 时可传入以下类型（支持多个）
	// 
	// <ul><li>HANDWRITE : 手写签名</li>
	// <li>OCR_ESIGN : OCR印章（智慧手写签名）</li>
	// <li>ESIGN : 个人印章</li>
	// <li>SYSTEM_ESIGN : 系统印章</li></ul>
	// 
	// 3.当ComponentType 是 SIGN_LEGAL_PERSON_SEAL 时无需传递此参数。
	ComponentValue []*string `json:"ComponentValue,omitnil" name:"ComponentValue"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程ID，为32位字符串。
	// 建议开发者妥善保存此流程ID，以便于顺利进行后续操作。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChannelFlowEvidenceReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChannelFlowEvidenceReportResponseParams struct {
	// 出证报告 ID，可用于<a href="https://qian.tencent.com/developers/partnerApis/certificate/DescribeChannelFlowEvidenceReport" target="_blank">获取出证报告任务执行结果</a>查询出证任务结果和出证PDF的下载URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportId *string `json:"ReportId,omitnil" name:"ReportId"`

	// 出证任务执行的状态, 状态含义如下：
	// 
	// <ul><li>**EvidenceStatusExecuting**：  出证任务在执行中</li>
	// <li>**EvidenceStatusSuccess**：  出证任务执行成功</li>
	// <li>**EvidenceStatusFailed** ： 出征任务执行失败</li></ul>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 废除，字段无效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportUrl *string `json:"ReportUrl,omitnil" name:"ReportUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// <li>第三方平台子客企业标识: Agent. ProxyOperator.OpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOrganizationOpenId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 企业信息变更类型，可选类型如下：
	// <ul><li>**1**：企业超管变更, 可以将超管换成同企业的其他员工</li>
	// <li>**2**：企业基础信息变更, 可以改企业名称 , 所在地址 , 法人名字等信息</li></ul>
	ChangeType *uint64 `json:"ChangeType,omitnil" name:"ChangeType"`
}

type CreateChannelOrganizationInfoChangeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent. ProxyOperator.OpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOrganizationOpenId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 企业信息变更类型，可选类型如下：
	// <ul><li>**1**：企业超管变更, 可以将超管换成同企业的其他员工</li>
	// <li>**2**：企业基础信息变更, 可以改企业名称 , 所在地址 , 法人名字等信息</li></ul>
	ChangeType *uint64 `json:"ChangeType,omitnil" name:"ChangeType"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChannelOrganizationInfoChangeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChannelOrganizationInfoChangeUrlResponseParams struct {
	// 创建的企业信息变更链接。需要在移动端打开，会跳转到微信腾讯电子签小程序进行更换。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 链接过期时间。链接7天有效。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type CreateConsoleLoginUrlRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>注:
	// `1. 企业激活时,  此时的Agent.ProxyOrganizationOpenId将会是企业激活后企业的唯一标识, 建议开发者保存企业ProxyOrganizationOpenId，后续各项接口调用皆需要此参数。 `
	// `2. 员工认证时,  此时的Agent.ProxyOrganizationOpenId将会是员工认证加入企业后的唯一标识, 建议开发者保存此员工的penId, 后续各项接口调用皆需要此参数。 `
	// `3. 同渠道应用(Agent.AppId)下,企业唯一标识ProxyOrganizationOpenId需要保持唯一, 员工唯一标识OpenId也要保持唯一 (而不是企业下唯一)`
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 第三方平台子客企业名称，请确认该名称与企业营业执照中注册的名称一致。
	// 
	// 注:
	//  `1. 如果名称中包含英文括号()，请使用中文括号（）代替。`
	//  `2、该名称需要与Agent.ProxyOrganizationOpenId相匹配,  企业激活后Agent.ProxyOrganizationOpenId会跟此企业名称一一绑定; 如果您的企业已经在认证授权中或者激活完成，这里修改子客企业名字将不会生效。 `
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil" name:"ProxyOrganizationName"`

	// 子客企业统一社会信用代码，最大长度200个字符
	// 注意：`如果您的企业已经在认证授权中或者激活完成，这里修改子客企业名字将不会生效`。
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil" name:"UniformSocialCreditCode"`

	// 子客企业员工的姓名，最大长度50个字符,  员工的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：`该姓名需要和Agent.ProxyOperator.OpenId相匹配,  当员工完成认证后该姓名会和Agent.ProxyOperator.OpenId一一绑定, 若员工已认证加入企业，这里修改经办人名字传入将不会生效`
	ProxyOperatorName *string `json:"ProxyOperatorName,omitnil" name:"ProxyOperatorName"`

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
	Module *string `json:"Module,omitnil" name:"Module"`

	// 该参数和Module参数配合使用，用于指定模块下的资源Id，指定后链接登录将展示该资源的详情。
	// 
	// 根据Module参数的不同所代表的含义不同(ModuleId需要和Module对应，ModuleId可以通过API或者控制台获取到)。当前支持：
	// <table> <thead> <tr> <th>Module传值</th> <th>ModuleId传值</th> <th>进入的目标页面</th> </tr> </thead> <tbody> <tr> <td>SEAL</td> <td>印章ID</td> <td>查看指定印章的详情页面</td> </tr> <tr> <td>TEMPLATE</td> <td>合同模板ID</td> <td>指定模板的详情页面</td> </tr> <tr> <td>DOCUMENT</td> <td>合同ID</td> <td>指定合同的详情页面</td> </tr> </tbody> </table>
	// 注意：该参数**仅在企业和员工激活完成，登录控制台场景才生效**。
	ModuleId *string `json:"ModuleId,omitnil" name:"ModuleId"`

	// 是否展示左侧菜单栏 
	// <ul><li> **ENABLE** : (默认)进入web控制台展示左侧菜单栏</li>
	// <li> **DISABLE** : 进入web控制台不展示左侧菜单栏</li></ul>
	// 注：该参数**仅在企业和员工激活完成，登录控制台场景才生效**。
	MenuStatus *string `json:"MenuStatus,omitnil" name:"MenuStatus"`

	// 生成链接的类型：
	// 生成链接的类型
	// <ul><li>**PC**：(默认)web控制台链接, 需要在PC浏览器中打开</li>
	// <li>**CHANNEL**：H5跳转到电子签小程序链接,  一般用于发送短信中带的链接,  打开后进入腾讯电子签小程序</li>
	// <li>**APP**：第三方APP或小程序跳转电子签小程序链接, 一般用于贵方小程序或者APP跳转过来,  打开后进入腾讯电子签小程序</li></ul>
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// 触发自动跳转事件，仅对EndPoint为App类型有效，可选值包括：
	// <ul><li> **VERIFIED** :企业认证完成/员工认证完成后跳回原App/小程序</li></ul>
	AutoJumpBackEvent *string `json:"AutoJumpBackEvent,omitnil" name:"AutoJumpBackEvent"`

	// 可选的此企业允许的授权方式, 可以设置的方式有:
	// <ul><li>1：上传授权书</li>
	// <li>2：转法定代表人授权</li>
	// <li>4：企业实名认证（信任第三方认证源）（此项有排他性, 选择后不能增添其他的方式）</li></ul>
	// 注:<ul>
	// <li>未选择信任第三方认证源时，如果是法人进行企业激活，仅支持法人扫脸直接授权，该配置不对此法人生效`</li>
	// <li>选择信任第三方认证源时，请先通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/SyncProxyOrganization" target="_blank">同步企业信息</a>接口同步信息。</li>
	// <li>该参数仅在企业未激活时生效</li>
	// </ul>
	AuthorizationTypes []*int64 `json:"AuthorizationTypes,omitnil" name:"AuthorizationTypes"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type CreateConsoleLoginUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>注:
	// `1. 企业激活时,  此时的Agent.ProxyOrganizationOpenId将会是企业激活后企业的唯一标识, 建议开发者保存企业ProxyOrganizationOpenId，后续各项接口调用皆需要此参数。 `
	// `2. 员工认证时,  此时的Agent.ProxyOrganizationOpenId将会是员工认证加入企业后的唯一标识, 建议开发者保存此员工的penId, 后续各项接口调用皆需要此参数。 `
	// `3. 同渠道应用(Agent.AppId)下,企业唯一标识ProxyOrganizationOpenId需要保持唯一, 员工唯一标识OpenId也要保持唯一 (而不是企业下唯一)`
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 第三方平台子客企业名称，请确认该名称与企业营业执照中注册的名称一致。
	// 
	// 注:
	//  `1. 如果名称中包含英文括号()，请使用中文括号（）代替。`
	//  `2、该名称需要与Agent.ProxyOrganizationOpenId相匹配,  企业激活后Agent.ProxyOrganizationOpenId会跟此企业名称一一绑定; 如果您的企业已经在认证授权中或者激活完成，这里修改子客企业名字将不会生效。 `
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil" name:"ProxyOrganizationName"`

	// 子客企业统一社会信用代码，最大长度200个字符
	// 注意：`如果您的企业已经在认证授权中或者激活完成，这里修改子客企业名字将不会生效`。
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil" name:"UniformSocialCreditCode"`

	// 子客企业员工的姓名，最大长度50个字符,  员工的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 
	// 注：`该姓名需要和Agent.ProxyOperator.OpenId相匹配,  当员工完成认证后该姓名会和Agent.ProxyOperator.OpenId一一绑定, 若员工已认证加入企业，这里修改经办人名字传入将不会生效`
	ProxyOperatorName *string `json:"ProxyOperatorName,omitnil" name:"ProxyOperatorName"`

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
	Module *string `json:"Module,omitnil" name:"Module"`

	// 该参数和Module参数配合使用，用于指定模块下的资源Id，指定后链接登录将展示该资源的详情。
	// 
	// 根据Module参数的不同所代表的含义不同(ModuleId需要和Module对应，ModuleId可以通过API或者控制台获取到)。当前支持：
	// <table> <thead> <tr> <th>Module传值</th> <th>ModuleId传值</th> <th>进入的目标页面</th> </tr> </thead> <tbody> <tr> <td>SEAL</td> <td>印章ID</td> <td>查看指定印章的详情页面</td> </tr> <tr> <td>TEMPLATE</td> <td>合同模板ID</td> <td>指定模板的详情页面</td> </tr> <tr> <td>DOCUMENT</td> <td>合同ID</td> <td>指定合同的详情页面</td> </tr> </tbody> </table>
	// 注意：该参数**仅在企业和员工激活完成，登录控制台场景才生效**。
	ModuleId *string `json:"ModuleId,omitnil" name:"ModuleId"`

	// 是否展示左侧菜单栏 
	// <ul><li> **ENABLE** : (默认)进入web控制台展示左侧菜单栏</li>
	// <li> **DISABLE** : 进入web控制台不展示左侧菜单栏</li></ul>
	// 注：该参数**仅在企业和员工激活完成，登录控制台场景才生效**。
	MenuStatus *string `json:"MenuStatus,omitnil" name:"MenuStatus"`

	// 生成链接的类型：
	// 生成链接的类型
	// <ul><li>**PC**：(默认)web控制台链接, 需要在PC浏览器中打开</li>
	// <li>**CHANNEL**：H5跳转到电子签小程序链接,  一般用于发送短信中带的链接,  打开后进入腾讯电子签小程序</li>
	// <li>**APP**：第三方APP或小程序跳转电子签小程序链接, 一般用于贵方小程序或者APP跳转过来,  打开后进入腾讯电子签小程序</li></ul>
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// 触发自动跳转事件，仅对EndPoint为App类型有效，可选值包括：
	// <ul><li> **VERIFIED** :企业认证完成/员工认证完成后跳回原App/小程序</li></ul>
	AutoJumpBackEvent *string `json:"AutoJumpBackEvent,omitnil" name:"AutoJumpBackEvent"`

	// 可选的此企业允许的授权方式, 可以设置的方式有:
	// <ul><li>1：上传授权书</li>
	// <li>2：转法定代表人授权</li>
	// <li>4：企业实名认证（信任第三方认证源）（此项有排他性, 选择后不能增添其他的方式）</li></ul>
	// 注:<ul>
	// <li>未选择信任第三方认证源时，如果是法人进行企业激活，仅支持法人扫脸直接授权，该配置不对此法人生效`</li>
	// <li>选择信任第三方认证源时，请先通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/SyncProxyOrganization" target="_blank">同步企业信息</a>接口同步信息。</li>
	// <li>该参数仅在企业未激活时生效</li>
	// </ul>
	AuthorizationTypes []*int64 `json:"AuthorizationTypes,omitnil" name:"AuthorizationTypes"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	delete(f, "Module")
	delete(f, "ModuleId")
	delete(f, "MenuStatus")
	delete(f, "Endpoint")
	delete(f, "AutoJumpBackEvent")
	delete(f, "AuthorizationTypes")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsoleLoginUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsoleLoginUrlResponseParams struct {
	// 跳转链接, 链接的有效期根据企业,员工状态和终端等有区别, 可以参考下表
	// <table> <thead> <tr> <th>子客企业状态</th> <th>子客企业员工状态</th> <th>Endpoint</th> <th>链接有效期限</th> </tr> </thead>  <tbody> <tr> <td>企业未激活</td> <td>员工未认证</td> <td>PC/CHANNEL/APP</td> <td>一年</td>  </tr>  <tr> <td>企业已激活</td> <td>员工未认证</td> <td>PC/CHANNEL/APP</td> <td>一年</td>  </tr>  <tr> <td>企业已激活</td> <td>员工已认证</td> <td>PC</td> <td>5分钟</td>  </tr>  <tr> <td>企业已激活</td> <td>员工已认证</td> <td>CHANNEL/APP</td> <td>一年</td>  </tr> </tbody> </table>
	// 注： 
	// `1.链接仅单次有效，每次登录需要需要重新创建新的链接`
	// `2.创建的链接应避免被转义，如：&被转义为\u0026；如使用Postman请求后，请选择响应类型为 JSON，否则链接将被转义`
	ConsoleUrl *string `json:"ConsoleUrl,omitnil" name:"ConsoleUrl"`

	// 子客企业是否已开通腾讯电子签，
	// <ul><li> **true** :已经开通腾讯电子签</li>
	// <li> **false** :还未开通腾讯电子签</li></ul>
	// 
	// 注：`企业是否实名根据传参Agent.ProxyOrganizationOpenId进行判断，非企业名称或者社会信用代码`
	IsActivated *bool `json:"IsActivated,omitnil" name:"IsActivated"`

	// 当前经办人是否已认证并加入功能
	// <ul><li> **true** : 已经认证加入公司</li>
	// <li> **false** : 还未认证加入公司</li></ul>
	// 注意：**员工是否实名是根据Agent.ProxyOperator.OpenId判断，非经办人姓名**
	ProxyOperatorIsVerified *bool `json:"ProxyOperatorIsVerified,omitnil" name:"ProxyOperatorIsVerified"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type CreateFlowOption struct {
	// 是否允许修改合同信息，
	// **true**：可以
	// **false**：（默认）不可以
	CanEditFlow *bool `json:"CanEditFlow,omitnil" name:"CanEditFlow"`

	// 是否允许发起合同弹窗隐藏合同名称
	// **true**：允许
	// **false**：（默认）不允许
	HideShowFlowName *bool `json:"HideShowFlowName,omitnil" name:"HideShowFlowName"`

	// 是否允许发起合同弹窗隐藏合同类型，
	// **true**：允许
	// **false**：（默认）不允许
	HideShowFlowType *bool `json:"HideShowFlowType,omitnil" name:"HideShowFlowType"`

	// 是否允许发起合同弹窗隐藏合同到期时间
	// **true**：允许
	// **false**：（默认）不允许
	HideShowDeadline *bool `json:"HideShowDeadline,omitnil" name:"HideShowDeadline"`

	// 是否允许发起合同步骤跳过指定签署方步骤
	// **true**：允许
	// **false**：（默认）不允许
	CanSkipAddApprover *bool `json:"CanSkipAddApprover,omitnil" name:"CanSkipAddApprover"`

	// 定制化发起合同弹窗的描述信息，长度不能超过500，只能由中文、字母、数字和标点组成。
	CustomCreateFlowDescription *string `json:"CustomCreateFlowDescription,omitnil" name:"CustomCreateFlowDescription"`

	// 禁止编辑填写控件
	// 
	// **true**：禁止编辑填写控件
	// **false**：（默认）允许编辑填写控件
	ForbidEditFillComponent *bool `json:"ForbidEditFillComponent,omitnil" name:"ForbidEditFillComponent"`

	// 跳过上传文件步骤
	// 
	// **true**：跳过
	// **false**：（默认）不跳过，需要传ResourceId
	SkipUploadFile *bool `json:"SkipUploadFile,omitnil" name:"SkipUploadFile"`
}

// Predefined struct for user
type CreateFlowsByTemplatesRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 要创建的合同信息列表，最多支持一次创建20个合同
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil" name:"FlowInfos"`

	// 是否为预览模式，取值如下：
	// <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li>
	// <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。</li></ul>
	// 
	// 注:
	// 
	// `如果预览的文件中指定了动态表格控件，此时此接口返回的是合成前的文档预览链接，合成完成后的文档预览链接需要通过回调通知的方式或使用返回的TaskInfo中的TaskId通过ChannelGetTaskResultApi接口查询得到`
	NeedPreview *bool `json:"NeedPreview,omitnil" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子</li></ul>
	// 注: `此参数在NeedPreview 为true时有效`
	PreviewType *int64 `json:"PreviewType,omitnil" name:"PreviewType"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type CreateFlowsByTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 要创建的合同信息列表，最多支持一次创建20个合同
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil" name:"FlowInfos"`

	// 是否为预览模式，取值如下：
	// <ul><li> **false**：非预览模式（默认），会产生合同流程并返回合同流程编号FlowId。</li>
	// <li> **true**：预览模式，不产生合同流程，不返回合同流程编号FlowId，而是返回预览链接PreviewUrl，有效期为300秒，用于查看真实发起后合同的样子。</li></ul>
	// 
	// 注:
	// 
	// `如果预览的文件中指定了动态表格控件，此时此接口返回的是合成前的文档预览链接，合成完成后的文档预览链接需要通过回调通知的方式或使用返回的TaskInfo中的TaskId通过ChannelGetTaskResultApi接口查询得到`
	NeedPreview *bool `json:"NeedPreview,omitnil" name:"NeedPreview"`

	// 预览模式下产生的预览链接类型 
	// <ul><li> **0** :(默认) 文件流 ,点开后下载预览的合同PDF文件 </li>
	// <li> **1** :H5链接 ,点开后在浏览器中展示合同的样子</li></ul>
	// 注: `此参数在NeedPreview 为true时有效`
	PreviewType *int64 `json:"PreviewType,omitnil" name:"PreviewType"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 第三方应用平台的业务信息, 与创建合同的FlowInfos数组中的CustomerData一一对应
	CustomerData []*string `json:"CustomerData,omitnil" name:"CustomerData"`

	// 创建消息，对应多个合同ID，
	// 成功为“”,创建失败则对应失败消息
	ErrorMessages []*string `json:"ErrorMessages,omitnil" name:"ErrorMessages"`

	// 合同预览链接URL数组。
	// 
	// 注：如果是预览模式(即NeedPreview设置为true)时, 才会有此预览链接URL
	PreviewUrls []*string `json:"PreviewUrls,omitnil" name:"PreviewUrls"`

	// 复杂文档合成任务（如，包含动态表格的预览任务）的任务信息数组；
	// 如果文档需要异步合成，此字段会返回该异步任务的任务信息，后续可以通过ChannelGetTaskResultApi接口查询任务详情；
	TaskInfos []*TaskInfo `json:"TaskInfos,omitnil" name:"TaskInfos"`

	// 签署方信息，如角色ID、角色名称等
	FlowApprovers []*FlowApproverItem `json:"FlowApprovers,omitnil" name:"FlowApprovers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 电子印章名字，1-50个中文字符
	// 注:`同一企业下电子印章名字不能相同`
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 电子印章图片base64编码，大小不超过10M（原始图片不超过5M），只支持PNG或JPG图片格式
	// 
	// 注: `通过图片创建的电子印章，需电子签平台人工审核`
	// 
	SealImage *string `json:"SealImage,omitnil" name:"SealImage"`

	// 操作者的信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 电子印章生成方式
	// <ul>
	// <li><strong>空值</strong>:(默认)使用上传的图片生成印章, 此时需要上传SealImage图片</li>
	// <li><strong>SealGenerateSourceSystem</strong>: 系统生成印章, 无需上传SealImage图片</li>
	// </ul>
	GenerateSource *string `json:"GenerateSource,omitnil" name:"GenerateSource"`

	// 电子印章类型 , 可选类型如下: 
	// <ul><li>**OFFICIAL**: (默认)公章</li>
	// <li>**CONTRACT**: 合同专用章;</li>
	// <li>**FINANCE**: 财务专用章;</li>
	// <li>**PERSONNEL**: 人事专用章</li>
	// </ul>
	// 注: `同企业下只能有一个公章, 重复创建会报错`
	SealType *string `json:"SealType,omitnil" name:"SealType"`

	// 企业印章横向文字，最多可填15个汉字  (若超过印章最大宽度，优先压缩字间距，其次缩小字号)
	// 横向文字的位置如下图中的"印章横向文字在这里"
	// 
	// ![image](https://dyn.ess.tencent.cn/guide/capi/CreateSealByImage2.png)
	SealHorizontalText *string `json:"SealHorizontalText,omitnil" name:"SealHorizontalText"`

	// 印章样式, 可以选择的样式如下: 
	// <ul><li>**circle**:(默认)圆形印章</li>
	// <li>**ellipse**:椭圆印章</li></ul>
	SealStyle *string `json:"SealStyle,omitnil" name:"SealStyle"`

	// 印章尺寸取值描述, 可以选择的尺寸如下: 
	// <ul><li> **42_42**: 圆形企业公章直径42mm, 当SealStyle是圆形的时候才有效</li>
	// <li> **40_40**: 圆形企业印章直径40mm, 当SealStyle是圆形的时候才有效</li>
	// <li> **45_30**: 椭圆形印章45mm x 30mm, 当SealStyle是椭圆的时候才有效</li></ul>
	SealSize *string `json:"SealSize,omitnil" name:"SealSize"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 电子印章名字，1-50个中文字符
	// 注:`同一企业下电子印章名字不能相同`
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 电子印章图片base64编码，大小不超过10M（原始图片不超过5M），只支持PNG或JPG图片格式
	// 
	// 注: `通过图片创建的电子印章，需电子签平台人工审核`
	// 
	SealImage *string `json:"SealImage,omitnil" name:"SealImage"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 电子印章生成方式
	// <ul>
	// <li><strong>空值</strong>:(默认)使用上传的图片生成印章, 此时需要上传SealImage图片</li>
	// <li><strong>SealGenerateSourceSystem</strong>: 系统生成印章, 无需上传SealImage图片</li>
	// </ul>
	GenerateSource *string `json:"GenerateSource,omitnil" name:"GenerateSource"`

	// 电子印章类型 , 可选类型如下: 
	// <ul><li>**OFFICIAL**: (默认)公章</li>
	// <li>**CONTRACT**: 合同专用章;</li>
	// <li>**FINANCE**: 财务专用章;</li>
	// <li>**PERSONNEL**: 人事专用章</li>
	// </ul>
	// 注: `同企业下只能有一个公章, 重复创建会报错`
	SealType *string `json:"SealType,omitnil" name:"SealType"`

	// 企业印章横向文字，最多可填15个汉字  (若超过印章最大宽度，优先压缩字间距，其次缩小字号)
	// 横向文字的位置如下图中的"印章横向文字在这里"
	// 
	// ![image](https://dyn.ess.tencent.cn/guide/capi/CreateSealByImage2.png)
	SealHorizontalText *string `json:"SealHorizontalText,omitnil" name:"SealHorizontalText"`

	// 印章样式, 可以选择的样式如下: 
	// <ul><li>**circle**:(默认)圆形印章</li>
	// <li>**ellipse**:椭圆印章</li></ul>
	SealStyle *string `json:"SealStyle,omitnil" name:"SealStyle"`

	// 印章尺寸取值描述, 可以选择的尺寸如下: 
	// <ul><li> **42_42**: 圆形企业公章直径42mm, 当SealStyle是圆形的时候才有效</li>
	// <li> **40_40**: 圆形企业印章直径40mm, 当SealStyle是圆形的时候才有效</li>
	// <li> **45_30**: 椭圆形印章45mm x 30mm, 当SealStyle是椭圆的时候才有效</li></ul>
	SealSize *string `json:"SealSize,omitnil" name:"SealSize"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSealByImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealByImageResponseParams struct {
	// 电子印章ID，为32位字符串。
	// 建议开发者保留此印章ID，后续指定签署区印章或者操作印章需此印章ID。
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 电子印章预览链接地址，地址默认失效时间为24小时。
	// 
	// 注:`图片上传生成的电子印章无预览链接地址`
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程ID数组，最多支持100个。
	// 注: `该参数和合同组编号必须二选一`
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 合同组编号
	// 注：`该参数和合同流程ID数组必须二选一`
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 签署链接类型,可以设置的参数如下
	// <ul><li> **WEIXINAPP** :(默认)跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型 ，此时返回短链</li>
	// <li> **CHANNEL** :带有H5引导页的跳转电子签小程序的链接</li>
	// <li> **APP** :第三方App或小程序跳转电子签小程序的path, App或者小程序跳转适合此类型</li>
	// <li> **LONGURL2WEIXINAPP** :跳转电子签小程序的链接, H5跳转适合此类型，此时返回长链</li></ul>
	// 
	// 详细使用场景可以参数接口说明中的 **主要使用场景可以更加EndPoint分类如下**
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// 签署链接生成类型，可以选择的类型如下
	// 
	// <ul><li>**ALL**：(默认)全部签署方签署链接，此时不会给自动签署(静默签署)的签署方创建签署链接</li>
	// <li>**CHANNEL**：第三方子企业员工签署方</li>
	// <li>**NOT_CHANNEL**：SaaS平台企业员工签署方</li>
	// <li>**PERSON**：个人/自然人签署方</li>
	// <li>**FOLLOWER**：关注方，目前是合同抄送方</li>
	// <li>**RECIPIENT**：获取RecipientId对应的签署链接，可用于生成动态签署人补充链接</li></ul>
	GenerateType *string `json:"GenerateType,omitnil" name:"GenerateType"`

	// SaaS平台企业员工签署方的企业名称
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 
	// 注: `GenerateType为"NOT_CHANNEL"时必填`
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 合同流程里边参与方的姓名。
	// 注: `GenerateType为"PERSON"(即个人签署方)时必填`
	Name *string `json:"Name,omitnil" name:"Name"`

	// 合同流程里边签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 注:  `GenerateType为"PERSON"或"FOLLOWER"时必填`
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成(如存在X，请大写)。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母(但“I”、“O”除外)，后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 第三方平台子客企业的企业的标识, 即OrganizationOpenId
	// 注: `GenerateType为"CHANNEL"时必填`
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 第三方平台子客企业员工的标识OpenId，GenerateType为"CHANNEL"时可用，指定到具体参与人, 仅展示已经实名的经办人信息
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// Endpoint为"APP" 类型的签署链接，可以设置此值；支持调用方小程序打开签署链接，在电子签小程序完成签署后自动回跳至调用方小程序
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil" name:"AutoJumpBack"`

	// 签署完之后的H5页面的跳转链接，针对Endpoint为CHANNEL时有效，最大长度1000个字符。
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 生成的签署链接在签署页面隐藏的按钮列表，可设置如下：
	// 
	// <ul><li> **0** :合同签署页面更多操作按钮</li>
	// <li> **1** :合同签署页面更多操作的拒绝签署按钮</li>
	// <li> **2** :合同签署页面更多操作的转他人处理按钮</li>
	// <li> **3** :签署成功页的查看详情按钮</li></ul>
	// 
	// 注:  `字段为数组, 可以传值隐藏多个按钮`
	Hides []*int64 `json:"Hides,omitnil" name:"Hides"`

	// 参与方角色ID，用于生成动态签署人链接完成领取。
	// 
	// 注：`使用此参数需要与flow_ids数量一致并且一一对应, 表示在对应同序号的流程中的参与角色ID`，
	RecipientIds []*string `json:"RecipientIds,omitnil" name:"RecipientIds"`
}

type CreateSignUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// <li>第三方平台子客企业标识: Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent. ProxyOperator.OpenId</li>
	// </ul>
	// 第三方平台子客企业和员工必须已经过实名认证
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程ID数组，最多支持100个。
	// 注: `该参数和合同组编号必须二选一`
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 合同组编号
	// 注：`该参数和合同流程ID数组必须二选一`
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 签署链接类型,可以设置的参数如下
	// <ul><li> **WEIXINAPP** :(默认)跳转电子签小程序的http_url, 短信通知或者H5跳转适合此类型 ，此时返回短链</li>
	// <li> **CHANNEL** :带有H5引导页的跳转电子签小程序的链接</li>
	// <li> **APP** :第三方App或小程序跳转电子签小程序的path, App或者小程序跳转适合此类型</li>
	// <li> **LONGURL2WEIXINAPP** :跳转电子签小程序的链接, H5跳转适合此类型，此时返回长链</li></ul>
	// 
	// 详细使用场景可以参数接口说明中的 **主要使用场景可以更加EndPoint分类如下**
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// 签署链接生成类型，可以选择的类型如下
	// 
	// <ul><li>**ALL**：(默认)全部签署方签署链接，此时不会给自动签署(静默签署)的签署方创建签署链接</li>
	// <li>**CHANNEL**：第三方子企业员工签署方</li>
	// <li>**NOT_CHANNEL**：SaaS平台企业员工签署方</li>
	// <li>**PERSON**：个人/自然人签署方</li>
	// <li>**FOLLOWER**：关注方，目前是合同抄送方</li>
	// <li>**RECIPIENT**：获取RecipientId对应的签署链接，可用于生成动态签署人补充链接</li></ul>
	GenerateType *string `json:"GenerateType,omitnil" name:"GenerateType"`

	// SaaS平台企业员工签署方的企业名称
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 
	// 注: `GenerateType为"NOT_CHANNEL"时必填`
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 合同流程里边参与方的姓名。
	// 注: `GenerateType为"PERSON"(即个人签署方)时必填`
	Name *string `json:"Name,omitnil" name:"Name"`

	// 合同流程里边签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 注:  `GenerateType为"PERSON"或"FOLLOWER"时必填`
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成(如存在X，请大写)。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母(但“I”、“O”除外)，后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 第三方平台子客企业的企业的标识, 即OrganizationOpenId
	// 注: `GenerateType为"CHANNEL"时必填`
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 第三方平台子客企业员工的标识OpenId，GenerateType为"CHANNEL"时可用，指定到具体参与人, 仅展示已经实名的经办人信息
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// Endpoint为"APP" 类型的签署链接，可以设置此值；支持调用方小程序打开签署链接，在电子签小程序完成签署后自动回跳至调用方小程序
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil" name:"AutoJumpBack"`

	// 签署完之后的H5页面的跳转链接，针对Endpoint为CHANNEL时有效，最大长度1000个字符。
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 生成的签署链接在签署页面隐藏的按钮列表，可设置如下：
	// 
	// <ul><li> **0** :合同签署页面更多操作按钮</li>
	// <li> **1** :合同签署页面更多操作的拒绝签署按钮</li>
	// <li> **2** :合同签署页面更多操作的转他人处理按钮</li>
	// <li> **3** :签署成功页的查看详情按钮</li></ul>
	// 
	// 注:  `字段为数组, 可以传值隐藏多个按钮`
	Hides []*int64 `json:"Hides,omitnil" name:"Hides"`

	// 参与方角色ID，用于生成动态签署人链接完成领取。
	// 
	// 注：`使用此参数需要与flow_ids数量一致并且一一对应, 表示在对应同序号的流程中的参与角色ID`，
	RecipientIds []*string `json:"RecipientIds,omitnil" name:"RecipientIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSignUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignUrlsResponseParams struct {
	// 签署参与者签署H5链接信息数组
	SignUrlInfos []*SignUrlInfo `json:"SignUrlInfos,omitnil" name:"SignUrlInfos"`

	// 生成失败时的错误信息，成功返回”“，顺序和出参SignUrlInfos保持一致
	ErrorMessages []*string `json:"ErrorMessages,omitnil" name:"ErrorMessages"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type Department struct {
	// 部门id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 部门名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentName *string `json:"DepartmentName,omitnil" name:"DepartmentName"`
}

// Predefined struct for user
type DescribeBillUsageDetailRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询开始时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的套餐类型 （选填 ）不传则查询所有套餐；
	// 对应关系如下
	// CloudEnterprise-企业版合同
	// SingleSignature-单方签章
	// CloudProve-签署报告
	// CloudOnlineSign-腾讯会议在线签约
	// ChannelWeCard-微工卡
	// SignFlow-合同套餐
	// SignFace-签署意愿（人脸识别）
	// SignPassword-签署意愿（密码）
	// SignSMS-签署意愿（短信）
	// PersonalEssAuth-签署人实名（腾讯电子签认证）
	// PersonalThirdAuth-签署人实名（信任第三方认证）
	// OrgEssAuth-签署企业实名
	// FlowNotify-短信通知
	// AuthService-企业工商信息查询
	QuotaType *string `json:"QuotaType,omitnil" name:"QuotaType"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 指定分页每页返回的数据条数，如果不传默认为 50，单页最大支持 50。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeBillUsageDetailRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询开始时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间字符串，格式为yyyymmdd,时间跨度不能大于31天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的套餐类型 （选填 ）不传则查询所有套餐；
	// 对应关系如下
	// CloudEnterprise-企业版合同
	// SingleSignature-单方签章
	// CloudProve-签署报告
	// CloudOnlineSign-腾讯会议在线签约
	// ChannelWeCard-微工卡
	// SignFlow-合同套餐
	// SignFace-签署意愿（人脸识别）
	// SignPassword-签署意愿（密码）
	// SignSMS-签署意愿（短信）
	// PersonalEssAuth-签署人实名（腾讯电子签认证）
	// PersonalThirdAuth-签署人实名（信任第三方认证）
	// OrgEssAuth-签署企业实名
	// FlowNotify-短信通知
	// AuthService-企业工商信息查询
	QuotaType *string `json:"QuotaType,omitnil" name:"QuotaType"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页，页码从 0 开始，即首页为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 指定分页每页返回的数据条数，如果不传默认为 50，单页最大支持 50。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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
	delete(f, "Agent")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QuotaType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillUsageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillUsageDetailResponseParams struct {
	// 返回查询记录总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 消耗记录详情
	Details []*BillUsageDetail `json:"Details,omitnil" name:"Details"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署报告编号, 由<a href="https://qian.tencent.com/developers/partnerApis/certificate/CreateChannelFlowEvidenceReport" target="_blank">提交申请出证报告任务</a>产生
	ReportId *string `json:"ReportId,omitnil" name:"ReportId"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署报告编号, 由<a href="https://qian.tencent.com/developers/partnerApis/certificate/CreateChannelFlowEvidenceReport" target="_blank">提交申请出证报告任务</a>产生
	ReportId *string `json:"ReportId,omitnil" name:"ReportId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChannelFlowEvidenceReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelFlowEvidenceReportResponseParams struct {
	// 出证报告PDF的下载 URL，有效期为5分钟，超过有效期后将无法再下载。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportUrl *string `json:"ReportUrl,omitnil" name:"ReportUrl"`

	// 出证任务执行的状态, 状态含义如下：
	// 
	// <ul><li>**EvidenceStatusExecuting**：  出证任务在执行中</li>
	// <li>**EvidenceStatusSuccess**：  出证任务执行成功</li>
	// <li>**EvidenceStatusFailed** ： 出征任务执行失败</li></ul>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定分页每页返回的数据条数，单页最大支持 200。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 该字段是指第三方平台子客企业的唯一标识，用于查询单独某个子客的企业数据。
	// 
	// **注**：`如果需要批量查询本应用下的所有企业的信息，则该字段不需要赋值`
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 可以按照当前企业的认证状态进行过滤。可值如下：
	// <ul><li>**"UNVERIFIED"**： 未认证的企业</li>
	//   <li>**"VERIFYINGLEGALPENDINGAUTHORIZATION"**： 认证中待法人授权的企业</li>
	//   <li>**"VERIFYINGAUTHORIZATIONFILEPENDING"**： 认证中授权书审核中的企业</li>
	//   <li>**"VERIFYINGAUTHORIZATIONFILEREJECT"**： 认证中授权书已驳回的企业</li>
	//   <li>**"VERIFYING"**： 认证进行中的企业</li>
	//   <li>**"VERIFIED"**： 已认证完成的企业</li></ul>
	AuthorizationStatusList []*string `json:"AuthorizationStatusList,omitnil" name:"AuthorizationStatusList"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页。 页码从 0 开始，即首页为 0，最大20000。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定分页每页返回的数据条数，单页最大支持 200。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 该字段是指第三方平台子客企业的唯一标识，用于查询单独某个子客的企业数据。
	// 
	// **注**：`如果需要批量查询本应用下的所有企业的信息，则该字段不需要赋值`
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 可以按照当前企业的认证状态进行过滤。可值如下：
	// <ul><li>**"UNVERIFIED"**： 未认证的企业</li>
	//   <li>**"VERIFYINGLEGALPENDINGAUTHORIZATION"**： 认证中待法人授权的企业</li>
	//   <li>**"VERIFYINGAUTHORIZATIONFILEPENDING"**： 认证中授权书审核中的企业</li>
	//   <li>**"VERIFYINGAUTHORIZATIONFILEREJECT"**： 认证中授权书已驳回的企业</li>
	//   <li>**"VERIFYING"**： 认证进行中的企业</li>
	//   <li>**"VERIFIED"**： 已认证完成的企业</li></ul>
	AuthorizationStatusList []*string `json:"AuthorizationStatusList,omitnil" name:"AuthorizationStatusList"`

	// 指定分页返回第几页的数据，如果不传默认返回第一页。 页码从 0 开始，即首页为 0，最大20000。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
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
	ChannelOrganizationInfos []*ChannelOrganizationInfo `json:"ChannelOrganizationInfos,omitnil" name:"ChannelOrganizationInfos"`

	// 指定分页返回第几页的数据。页码从 0 开始，即首页为 0，最大20000。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 指定分页每页返回的数据条数，单页最大支持 200。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 满足查询条件的企业总数量。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 用印审批单的ID，可通过用印申请回调获取。
	WorkflowInstanceId *string `json:"WorkflowInstanceId,omitnil" name:"WorkflowInstanceId"`

	// 生成链接的类型：
	// 生成链接的类型
	// <ul><li>**LongLink**：(默认)长链接，H5跳转到电子签小程序链接，链接有效期为1年</li>
	// <li>**ShortLink**：H5跳转到电子签小程序链接，一般用于发送短信中带的链接，打开后进入腾讯电子签小程序，链接有效期为7天</li>
	// <li>**App**：第三方APP或小程序跳转电子签小程序链接，一般用于贵方小程序或者APP跳转过来，打开后进入腾讯电子签小程序，链接有效期为1年</li></ul>
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 用印审批单的ID，可通过用印申请回调获取。
	WorkflowInstanceId *string `json:"WorkflowInstanceId,omitnil" name:"WorkflowInstanceId"`

	// 生成链接的类型：
	// 生成链接的类型
	// <ul><li>**LongLink**：(默认)长链接，H5跳转到电子签小程序链接，链接有效期为1年</li>
	// <li>**ShortLink**：H5跳转到电子签小程序链接，一般用于发送短信中带的链接，打开后进入腾讯电子签小程序，链接有效期为7天</li>
	// <li>**App**：第三方APP或小程序跳转电子签小程序链接，一般用于贵方小程序或者APP跳转过来，打开后进入腾讯电子签小程序，链接有效期为1年</li></ul>
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`
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
	WorkflowUrl *string `json:"WorkflowUrl,omitnil" name:"WorkflowUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeExtendedServiceAuthInfoRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业标识: Agent. ProxyOperator.OpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DescribeExtendedServiceAuthInfoRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业标识: Agent. ProxyOperator.OpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthInfo []*ExtentServiceAuthInfo `json:"AuthInfo,omitnil" name:"AuthInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 需要查询的流程ID列表，最多可传入100个ID。
	// 如果要查询合同组的信息，则不需要传入此参数，只需传入 FlowGroupId 参数即可。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 需要查询的流程组ID，如果传入此参数，则会忽略 FlowIds 参数。
	// 
	// 合同组由<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByTemplates" target="_blank">通过多模板创建合同组签署流程</a>和<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByFiles" target="_blank">通过多文件创建合同组签署流程</a>等接口创建。
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 需要查询的流程ID列表，最多可传入100个ID。
	// 如果要查询合同组的信息，则不需要传入此参数，只需传入 FlowGroupId 参数即可。
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 需要查询的流程组ID，如果传入此参数，则会忽略 FlowIds 参数。
	// 
	// 合同组由<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByTemplates" target="_blank">通过多模板创建合同组签署流程</a>和<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByFiles" target="_blank">通过多文件创建合同组签署流程</a>等接口创建。
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 合同归属的第三方平台子客企业OpenId
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil" name:"ProxyOrganizationOpenId"`

	// 合同流程的详细信息。
	// 如果查询的是合同组信息，则返回的是组内所有子合同流程的详细信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowInfo []*FlowDetailInfo `json:"FlowInfo,omitnil" name:"FlowInfo"`

	// 合同组ID，只有在查询合同组信息时才会返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 合同组名称，只有在查询合同组信息时才会返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 需要下载的合同流程的ID,  至少需要1个,  做多50个
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 需要下载的合同流程的ID,  至少需要1个,  做多50个
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	FlowResourceUrlInfos []*FlowResourceUrlInfo `json:"FlowResourceUrlInfos,omitnil" name:"FlowResourceUrlInfos"`

	// 如果某个序号的合同流程生成PDF下载链接成功, 对应序号的值为空
	// 如果某个序号的合同流程生成PDF下载链接失败, 对应序号的值为错误的原因
	ErrorMessages []*string `json:"ErrorMessages,omitnil" name:"ErrorMessages"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同模板ID，为32位字符串。
	// 
	// 可以通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>登录企业控制台, 在企业模板中得到合同模板ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 查询模板的内容
	// 
	// <ul><li>**0**：（默认）模板列表及详情</li>
	// <li>**1**：仅模板列表, 不会返回模板中的签署控件, 填写控件, 参与方角色列表等信息</li></ul>
	ContentType *int64 `json:"ContentType,omitnil" name:"ContentType"`

	// 合同模板ID数组，每一个合同模板ID为32位字符串,  最多支持200个模板的批量查询。
	// 
	// 注意: 
	// 1.` 此参数TemplateIds与TemplateId互为独立，若两者均传入，以TemplateId为准。`
	// 2. `请确保每个模板均正确且属于当前企业，若有任一模板不存在，则返回错误。`
	// 4. `若传递此参数，分页参数(Limit,Offset)无效`
	TemplateIds []*string `json:"TemplateIds,omitnil" name:"TemplateIds"`

	// 指定每页返回的数据条数，和Offset参数配合使用。
	// 
	// 注：`1.默认值为20，单页做大值为200。`
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。
	// 
	// 注：`1.offset从0开始，即第一页为0。`
	// `2.默认从第一页返回。`
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 模糊搜索的模板名称，注意是模板名的连续部分，长度不能超过200，可支持由中文、字母、数字和下划线组成字符串。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 对应第三方应用平台企业的模板ID，通过此值可以搜索由第三方应用平台模板ID下发或领取得到的子客模板列表。
	ChannelTemplateId *string `json:"ChannelTemplateId,omitnil" name:"ChannelTemplateId"`

	// 返回控件的范围, 是只返回发起方自己的还是所有参与方的
	// 
	// <ul><li>**false**：（默认）只返回发起方控件</li>
	// <li>**true**：返回所有参与方(包括发起方和签署方)控件</li></ul>
	QueryAllComponents *bool `json:"QueryAllComponents,omitnil" name:"QueryAllComponents"`

	// 是否获取模板预览链接。
	// 
	// <ul><li>**false**：不获取（默认）</li>
	// <li>**true**：获取</li></ul>
	// 
	// 设置为true之后， 返回参数PreviewUrl，为模板的H5预览链接,  有效期5分钟。可以通过浏览器打开此链接预览模板，或者嵌入到iframe中预览模板。
	// 
	// 注: `此功能为白名单功能，使用前请联系对接的客户经理沟通。`
	WithPreviewUrl *bool `json:"WithPreviewUrl,omitnil" name:"WithPreviewUrl"`

	// 是否获取模板的PDF文件链接。
	// 
	// <ul><li>**false**：不获取（默认）</li>
	// <li>**true**：获取</li></ul>
	// 
	// 设置为true之后， 返回参数PdfUrl，为模板PDF文件链接，有效期5分钟, 可以用于将PDF文件下载到本地
	// 
	// 注: `此功能为白名单功能，使用前请联系对接的客户经理沟通。`
	WithPdfUrl *bool `json:"WithPdfUrl,omitnil" name:"WithPdfUrl"`

	// 操作者的信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同模板ID，为32位字符串。
	// 
	// 可以通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>登录企业控制台, 在企业模板中得到合同模板ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 查询模板的内容
	// 
	// <ul><li>**0**：（默认）模板列表及详情</li>
	// <li>**1**：仅模板列表, 不会返回模板中的签署控件, 填写控件, 参与方角色列表等信息</li></ul>
	ContentType *int64 `json:"ContentType,omitnil" name:"ContentType"`

	// 合同模板ID数组，每一个合同模板ID为32位字符串,  最多支持200个模板的批量查询。
	// 
	// 注意: 
	// 1.` 此参数TemplateIds与TemplateId互为独立，若两者均传入，以TemplateId为准。`
	// 2. `请确保每个模板均正确且属于当前企业，若有任一模板不存在，则返回错误。`
	// 4. `若传递此参数，分页参数(Limit,Offset)无效`
	TemplateIds []*string `json:"TemplateIds,omitnil" name:"TemplateIds"`

	// 指定每页返回的数据条数，和Offset参数配合使用。
	// 
	// 注：`1.默认值为20，单页做大值为200。`
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。
	// 
	// 注：`1.offset从0开始，即第一页为0。`
	// `2.默认从第一页返回。`
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 模糊搜索的模板名称，注意是模板名的连续部分，长度不能超过200，可支持由中文、字母、数字和下划线组成字符串。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 对应第三方应用平台企业的模板ID，通过此值可以搜索由第三方应用平台模板ID下发或领取得到的子客模板列表。
	ChannelTemplateId *string `json:"ChannelTemplateId,omitnil" name:"ChannelTemplateId"`

	// 返回控件的范围, 是只返回发起方自己的还是所有参与方的
	// 
	// <ul><li>**false**：（默认）只返回发起方控件</li>
	// <li>**true**：返回所有参与方(包括发起方和签署方)控件</li></ul>
	QueryAllComponents *bool `json:"QueryAllComponents,omitnil" name:"QueryAllComponents"`

	// 是否获取模板预览链接。
	// 
	// <ul><li>**false**：不获取（默认）</li>
	// <li>**true**：获取</li></ul>
	// 
	// 设置为true之后， 返回参数PreviewUrl，为模板的H5预览链接,  有效期5分钟。可以通过浏览器打开此链接预览模板，或者嵌入到iframe中预览模板。
	// 
	// 注: `此功能为白名单功能，使用前请联系对接的客户经理沟通。`
	WithPreviewUrl *bool `json:"WithPreviewUrl,omitnil" name:"WithPreviewUrl"`

	// 是否获取模板的PDF文件链接。
	// 
	// <ul><li>**false**：不获取（默认）</li>
	// <li>**true**：获取</li></ul>
	// 
	// 设置为true之后， 返回参数PdfUrl，为模板PDF文件链接，有效期5分钟, 可以用于将PDF文件下载到本地
	// 
	// 注: `此功能为白名单功能，使用前请联系对接的客户经理沟通。`
	WithPdfUrl *bool `json:"WithPdfUrl,omitnil" name:"WithPdfUrl"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplatesResponseParams struct {
	// 模板详情列表数据
	Templates []*TemplateInfo `json:"Templates,omitnil" name:"Templates"`

	// 查询到的模板总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 每页返回的数据条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果分页返回，此处指定第几页。页码从0开始，即首页为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询日期范围的开始时间, 查询会包含此日期的数据 , 格式为yyyy-mm-dd (例如：2021-03-21)
	// 
	// 注: `查询日期范围区间长度大于90天`。
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 查询日期范围的结束时间, 查询会包含此日期的数据 , 格式为yyyy-mm-dd (例如：2021-04-21)
	// 
	// 注: `查询日期范围区间长度大于90天`。
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 是否汇总数据，默认不汇总。
	// <ul><li> **true** :  汇总数据,  即每个企业一条数据, 对日志范围内的数据相加</li>
	// <li> **false** :  不会总数据,  返回企业每日明细,   按日期返回每个企业的数据(如果企业对应天数没有操作则无此企业此日期的数据)</li></ul>
	NeedAggregate *bool `json:"NeedAggregate,omitnil" name:"NeedAggregate"`

	// 指定每页返回的数据条数，和Offset参数配合使用。
	// 
	// 注: `默认值为1000，单页做大值为1000`
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。
	// 
	// 注：`offset从0开始，即第一页为0。`
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type DescribeUsageRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询日期范围的开始时间, 查询会包含此日期的数据 , 格式为yyyy-mm-dd (例如：2021-03-21)
	// 
	// 注: `查询日期范围区间长度大于90天`。
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 查询日期范围的结束时间, 查询会包含此日期的数据 , 格式为yyyy-mm-dd (例如：2021-04-21)
	// 
	// 注: `查询日期范围区间长度大于90天`。
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 是否汇总数据，默认不汇总。
	// <ul><li> **true** :  汇总数据,  即每个企业一条数据, 对日志范围内的数据相加</li>
	// <li> **false** :  不会总数据,  返回企业每日明细,   按日期返回每个企业的数据(如果企业对应天数没有操作则无此企业此日期的数据)</li></ul>
	NeedAggregate *bool `json:"NeedAggregate,omitnil" name:"NeedAggregate"`

	// 指定每页返回的数据条数，和Offset参数配合使用。
	// 
	// 注: `默认值为1000，单页做大值为1000`
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果分页返回，指定从第几页返回数据，和Limit参数配合使用。
	// 
	// 注：`offset从0开始，即第一页为0。`
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 用量明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*UsageDetail `json:"Details,omitnil" name:"Details"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 签署流程的标识数组
	FlowIdList []*string `json:"FlowIdList,omitnil" name:"FlowIdList"`
}

type ExtentServiceAuthInfo struct {
	// 扩展服务类型
	// <ul>
	// <li>AUTO_SIGN             企业自动签（自动签署）</li>
	// <li>  OVERSEA_SIGN          企业与港澳台居民*签署合同</li>
	// <li>  MOBILE_CHECK_APPROVER 使用手机号验证签署方身份</li>
	// <li> PAGING_SEAL           骑缝章</li>
	// <li> DOWNLOAD_FLOW         授权渠道下载合同 </li>
	// <li>AGE_LIMIT_EXPANSION 拓宽签署方年龄限制</li>
	// </ul>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 扩展服务名称 
	Name *string `json:"Name,omitnil" name:"Name"`

	// 扩展服务的开通状态
	// **ENABLE**：开通 
	// **DISABLE**：未开通	
	Status *string `json:"Status,omitnil" name:"Status"`

	// 操作扩展服务的操作人第三方应用平台的用户openid
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorOpenId *string `json:"OperatorOpenId,omitnil" name:"OperatorOpenId"`

	// 扩展服务的操作时间，格式为Unix标准时间戳（秒）。	
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateOn *int64 `json:"OperateOn,omitnil" name:"OperateOn"`
}

type FailedCreateRoleData struct {
	// 用户userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 角色RoleId列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleIds []*string `json:"RoleIds,omitnil" name:"RoleIds"`
}

type FillApproverInfo struct {
	// 签署方经办人在模板中配置的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 指定企业经办签署人OpenId
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 签署人姓名
	ApproverName *string `json:"ApproverName,omitnil" name:"ApproverName"`

	// 签署人手机号码
	ApproverMobile *string `json:"ApproverMobile,omitnil" name:"ApproverMobile"`

	// 企业名称
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 企业OpenId
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 签署企业非渠道子客，默认为false，即表示同一渠道下的企业；如果为true，则目前表示接收方企业为SaaS企业, 为渠道子客时，OrganizationOpenId 必传
	NotChannelOrganization *bool `json:"NotChannelOrganization,omitnil" name:"NotChannelOrganization"`
}

type FillError struct {
	// 为签署方经办人在签署合同中的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。与入参中补充的签署人角色ID对应，批量补充部分失败返回对应的错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 补充失败错误说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitnil" name:"ErrMessage"`
}

type FilledComponent struct {
	// 填写控件ID
	ComponentId *string `json:"ComponentId,omitnil" name:"ComponentId"`

	// 控件名称
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// 此填写控件的填写状态
	//  **0** : 此填写控件**未填写**
	// **1** : 此填写控件**已填写**
	ComponentFillStatus *string `json:"ComponentFillStatus,omitnil" name:"ComponentFillStatus"`

	// 控件填写内容
	ComponentValue *string `json:"ComponentValue,omitnil" name:"ComponentValue"`

	// 图片填充控件下载链接，如果是图片填充控件时，这里返回图片的下载链接。
	// 
	// 注: `链接不是永久链接,  默认有效期5分钟后, 到期后链接失效`
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type Filter struct {
	// 查询过滤条件的Key
	Key *string `json:"Key,omitnil" name:"Key"`

	// 查询过滤条件的Value列表
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type FlowApproverDetail struct {
	// 模板配置时候的签署人角色ID(用PDF文件发起也可以指定,如果不指定则自动生成此角色ID), 所有的填写控件和签署控件都归属不同的角色
	ReceiptId *string `json:"ReceiptId,omitnil" name:"ReceiptId"`

	// 第三方平台子客企业的唯一标识，定义Agent中的ProxyOrganizationOpenId一样, 可以参考<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#agent" target="_blank">Agent结构体</a>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil" name:"ProxyOrganizationOpenId"`

	// 第三方平台子客企业员工的唯一标识
	ProxyOperatorOpenId *string `json:"ProxyOperatorOpenId,omitnil" name:"ProxyOperatorOpenId"`

	// 第三方平台子客企业名称，与企业营业执照中注册的名称一致。
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil" name:"ProxyOrganizationName"`

	// 签署人手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 签署顺序，如果是有序签署，签署顺序从小到大
	SignOrder *int64 `json:"SignOrder,omitnil" name:"SignOrder"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveName *string `json:"ApproveName,omitnil" name:"ApproveName"`

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
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveStatus *string `json:"ApproveStatus,omitnil" name:"ApproveStatus"`

	// 签署人拒签等情况的时候填写的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveMessage *string `json:"ApproveMessage,omitnil" name:"ApproveMessage"`

	// 签署人签署时间戳，单位秒
	ApproveTime *int64 `json:"ApproveTime,omitnil" name:"ApproveTime"`

	// 参与者类型 
	// <ul><li> **ORGANIZATION** :企业签署人</li>
	// <li> **PERSON** :个人签署人</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveType *string `json:"ApproveType,omitnil" name:"ApproveType"`

	// 自定义签署人的角色名, 如: 收款人、开具人、见证人等
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverRoleName *string `json:"ApproverRoleName,omitnil" name:"ApproverRoleName"`
}

type FlowApproverInfo struct {
	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 签署方经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li>
	// <li>OTHER_CARD_TYPE : 其他证件</li></ul>
	// 
	// 注: `其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 签署方经办人的证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母（但“I”、“O”除外），后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)， 不支持海外手机号。
	// 请确认手机号所有方为此合同签署方。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 指定签署人非第三方平台子客企业下员工还是SaaS平台企业，在ApproverType为ORGANIZATION时指定。
	// <ul>
	// <li>false: 默认值，第三方平台子客企业下员工</li>
	// <li>true: SaaS平台企业下的员工</li>
	// </ul>
	NotChannelOrganization *bool `json:"NotChannelOrganization,omitnil" name:"NotChannelOrganization"`

	// 第三方平台子客企业员工的唯一标识，长度不能超过64，只能由字母和数字组成
	// 
	// 当签署方为同一第三方平台下的员工时，该字段若不指定，则发起【待领取】的流程
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 同应用下第三方平台子客企业的唯一标识，定义Agent中的ProxyOrganizationOpenId一样，签署方为非发起方企业场景下必传，最大长度64个字符
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

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
	ApproverType *string `json:"ApproverType,omitnil" name:"ApproverType"`

	// 签署流程签署人在模板中对应的签署人Id；在非单方签署、以及非B2C签署的场景下必传，用于指定当前签署方在签署流程中的位置；
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 本签署人在此合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 签署完回调url，最大长度1000个字符
	//
	// Deprecated: CallbackUrl is deprecated.
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 使用PDF文件直接发起合同时，签署人指定的签署控件；<br/>使用模板发起合同时，指定本企业印章签署控件的印章ID: <br/>通过ComponentId或ComponenetName指定签署控件，ComponentValue为印章ID。
	SignComponents []*Component `json:"SignComponents,omitnil" name:"SignComponents"`

	// 签署方控件类型为 SIGN_SIGNATURE时，可以指定签署方签名方式
	// 	HANDWRITE – 手写签名
	// 	OCR_ESIGN -- AI智能识别手写签名
	// 	ESIGN -- 个人印章类型
	// 	SYSTEM_ESIGN -- 系统签名（该类型可以在用户签署时根据用户姓名一键生成一个签名来进行签署）
	ComponentLimitType []*string `json:"ComponentLimitType,omitnil" name:"ComponentLimitType"`

	// 签署方在签署合同之前，需要强制阅读合同的时长，可指定为3秒至300秒之间的任意值。
	// 
	// 若未指定阅读时间，则会按照合同页数大小计算阅读时间，计算规则如下：
	// <ul>
	// <li>合同页数少于等于2页，阅读时间为3秒；</li>
	// <li>合同页数为3到5页，阅读时间为5秒；</li>
	// <li>合同页数大于等于6页，阅读时间为10秒。</li>
	// </ul>
	PreReadTime *int64 `json:"PreReadTime,omitnil" name:"PreReadTime"`

	// 签署完前端跳转的url，此字段的用法场景请联系客户经理确认
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`

	// 可以控制签署方在签署合同时能否进行某些操作，例如拒签、转交他人、是否为动态补充签署人等。
	// 详细操作可以参考开发者中心的ApproverOption结构体。
	ApproverOption *ApproverOption `json:"ApproverOption,omitnil" name:"ApproverOption"`

	// 当前签署方进行签署操作是否需要企业内部审批，true 则为需要
	ApproverNeedSignReview *bool `json:"ApproverNeedSignReview,omitnil" name:"ApproverNeedSignReview"`

	// 指定个人签署方查看合同的校验方式,可以传值如下:
	// <ul><li>  **1**   : （默认）人脸识别,人脸识别后才能合同内容</li>
	// <li>  **2**  : 手机号验证, 用户手机号和参与方手机号(ApproverMobile)相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// </li></ul>
	// 注: 
	// <ul><li>如果合同流程设置ApproverVerifyType查看合同的校验方式,    则忽略此签署人的查看合同的校验方式</li>
	// <li>此字段可传多个校验方式</li></ul>
	ApproverVerifyTypes []*int64 `json:"ApproverVerifyTypes,omitnil" name:"ApproverVerifyTypes"`

	// 签署人签署合同时的认证方式
	// <ul><li> **1** :人脸认证</li>
	// <li> **2** :签署密码</li>
	// <li> **3** :运营商三要素</li></ul>
	// 
	// 默认为1(人脸认证 ),2(签署密码)
	// 
	// 注: 
	// 1. 用<font color='red'>模板创建合同场景</font>, 签署人的认证方式需要在配置模板的时候指定, <font color='red'>在创建合同重新指定无效</font>
	// 2. 运营商三要素认证方式对手机号运营商及前缀有限制,可以参考[运营商支持列表类](https://qian.tencent.com/developers/partner/mobile_support)得到具体的支持说明
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil" name:"ApproverSignTypes"`

	// 签署ID
	// - 发起流程时系统自动补充
	// - 创建签署链接时，可以通过查询详情接口获得签署人的SignId，然后可传入此值为该签署人创建签署链接，无需再传姓名、手机号、证件号等其他信息
	SignId *string `json:"SignId,omitnil" name:"SignId"`

	// 通知签署方经办人的方式, 有以下途径:
	// <ul><li> **SMS** :(默认)短信</li>
	// <li> **NONE** : 不通知</li></ul>
	// 
	// 注: `签署方为第三方子客企业时会被置为NONE,   不会发短信通知`
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// [通过文件创建签署流程](https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowByFiles)时,如果设置了外层参数SignBeanTag=1(允许签署过程中添加签署控件),则可通过此参数明确规定合同所使用的签署控件类型（骑缝章、普通章法人章等）和具体的印章（印章ID）或签名方式。
	// 
	// 注：`限制印章控件或骑缝章控件情况下,仅本企业签署方可以指定具体印章（通过传递ComponentValue,支持多个），他方企业或个人只支持限制控件类型。`
	AddSignComponentsLimits []*ComponentLimit `json:"AddSignComponentsLimits,omitnil" name:"AddSignComponentsLimits"`

	// 可以自定义签署人角色名：收款人、开具人、见证人等，长度不能超过20，只能由中文、字母、数字和下划线组成。
	// 
	// 注: `如果是用模板发起, 优先使用此处上传的, 如果不传则用模板的配置的`
	ApproverRoleName *string `json:"ApproverRoleName,omitnil" name:"ApproverRoleName"`
}

type FlowApproverItem struct {
	// 合同编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 签署方信息，如角色ID、角色名称等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Approvers []*ApproverItem `json:"Approvers,omitnil" name:"Approvers"`
}

type FlowApproverUrlInfo struct {
	// 签署短链接。</br>
	// 注意:
	// - 该链接有效期为**30分钟**，同时需要注意保密，不要外泄给无关用户。
	// - 该链接不支持小程序嵌入，仅支持**移动端浏览器**打开。
	SignUrl *string `json:"SignUrl,omitnil" name:"SignUrl"`

	// 签署人类型。
	// - **PERSON**: 个人
	ApproverType *string `json:"ApproverType,omitnil" name:"ApproverType"`

	// 签署人姓名。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 签署人手机号。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 签署长链接。</br>
	// 注意:
	// - 该链接有效期为**30分钟**，同时需要注意保密，不要外泄给无关用户。
	// - 该链接不支持小程序嵌入，仅支持**移动端浏览器**打开。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LongUrl *string `json:"LongUrl,omitnil" name:"LongUrl"`
}

type FlowDetailInfo struct {
	// 合同流程ID，为32位字符串。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同流程的类别分类（如销售合同/入职合同等）。
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 合同流程当前的签署状态, 会存在下列的状态值
	// <ul><li> **INIT** :合同创建</li>
	// <li> **PART** :合同签署中(至少有一个签署方已经签署)</li>
	// <li> **REJECT** :合同拒签</li>
	// <li> **ALL** :合同签署完成</li>
	// <li> **DEADLINE** :合同流签(合同过期)</li>
	// <li> **CANCEL** :合同撤回</li>
	// <li> **RELIEVED** :解除协议（已解除）</li></ul>
	//  
	FlowStatus *string `json:"FlowStatus,omitnil" name:"FlowStatus"`

	// 当合同流程状态为已拒签（即 FlowStatus=REJECT）或已撤销（即 FlowStatus=CANCEL ）时，此字段 FlowMessage 为拒签或撤销原因。
	FlowMessage *string `json:"FlowMessage,omitnil" name:"FlowMessage"`

	// 合同流程的创建时间戳，格式为Unix标准时间戳（秒）。
	CreateOn *int64 `json:"CreateOn,omitnil" name:"CreateOn"`

	// 签署流程的签署截止时间, 值为unix时间戳, 精确到秒。
	DeadLine *int64 `json:"DeadLine,omitnil" name:"DeadLine"`

	// 调用方自定义的个性化字段(可自定义此字段的值)，并以base64方式编码，支持的最大数据大小为 1000长度。
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	CustomData *string `json:"CustomData,omitnil" name:"CustomData"`

	// 合同流程的签署方数组
	FlowApproverInfos []*FlowApproverDetail `json:"FlowApproverInfos,omitnil" name:"FlowApproverInfos"`

	// 合同流程的关注方信息数组
	CcInfos []*FlowApproverDetail `json:"CcInfos,omitnil" name:"CcInfos"`

	// 是否需要发起前审批
	// <ul><li>当NeedCreateReview为true，表明当前流程是需要发起前审核的合同，可能无法进行查看，签署操作，需要等审核完成后，才可以继续后续流程</li>
	// <li>当NeedCreateReview为false，不需要发起前审核的合同</li></ul>
	NeedCreateReview *bool `json:"NeedCreateReview,omitnil" name:"NeedCreateReview"`
}

type FlowFileInfo struct {
	// 签署文件资源Id列表，目前仅支持单个文件
	FileIds []*string `json:"FileIds,omitnil" name:"FileIds"`

	// 签署流程名称，长度不超过200个字符
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 签署流程签约方列表，最多不超过5个参与方
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitnil" name:"FlowApprovers"`

	// 签署流程截止时间，十位数时间戳，最大值为33162419560，即3020年
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 签署流程的描述，长度不超过1000个字符
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 签署流程的类型，长度不超过255个字符
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 签署流程回调地址，长度不超过255个字符
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 第三方应用的业务信息，最大长度1000个字符。发起自动签署时，需设置对应自动签署场景，目前仅支持场景：处方单-E_PRESCRIPTION_AUTO_SIGN
	CustomerData *string `json:"CustomerData,omitnil" name:"CustomerData"`

	// 合同签署顺序类型(无序签,顺序签)，默认为false，即有序签署
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 签署文件中的发起方的填写控件，需要在发起的时候进行填充
	Components []*Component `json:"Components,omitnil" name:"Components"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitnil" name:"CustomShowMap"`

	// 本企业(发起方企业)是否需要签署审批
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`
}

type FlowGroupOptions struct {
	// 发起方企业经办人（即签署人为发起方企业员工）是否需要对子合同进行独立的意愿确认
	// <ul><li>**false**（默认）：发起方企业经办人签署时对所有子合同进行统一的意愿确认。</li>
	// <li>**true**：发起方企业经办人签署时需要对子合同进行独立的意愿确认。</li></ul>
	SelfOrganizationApproverSignEach *bool `json:"SelfOrganizationApproverSignEach,omitnil" name:"SelfOrganizationApproverSignEach"`

	// 非发起方企业经办人（即：签署人为个人或者不为发起方企业的员工）是否需要对子合同进行独立的意愿确认
	// <ul><li>**false**（默认）：非发起方企业经办人签署时对所有子合同进行统一的意愿确认。</li>
	// <li>**true**：非发起方企业经办人签署时需要对子合同进行独立的意愿确认。</li></ul>
	OtherApproverSignEach *bool `json:"OtherApproverSignEach,omitnil" name:"OtherApproverSignEach"`
}

type FlowInfo struct {
	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒），如果未设置签署截止时间，则默认为合同流程创建后的365天时截止。
	// 如果在签署截止时间前未完成签署，则合同状态会变为已过期，导致合同作废。
	// 示例值：1604912664
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 用户配置的合同模板ID，会基于此模板创建合同文档，为32位字符串。
	// 如果使用模板发起接口，此参数为必填。
	// 
	// 可以通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>登录企业控制台, 在**企业模板**中得到合同模板ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 多个签署人信息，最大支持50个签署方
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitnil" name:"FlowApprovers"`

	// 表单K-V对列表
	FormFields []*FormField `json:"FormFields,omitnil" name:"FormFields"`

	// 合同状态变动结的通知回调URL，该URL仅支持HTTP或HTTPS协议，建议采用HTTPS协议以保证数据传输的安全性，最大长度1000个字符。
	// 
	// 腾讯电子签服务器将通过POST方式，application/json格式通知执行结果，请确保外网可以正常访问该URL。
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/partner/callback_data_types" target="_blank">回调通知</a>模块
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 合同流程的类别分类（可自定义名称，如销售合同/入职合同等），最大长度为200个字符，仅限中文、字母、数字和下划线组成。
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 合同流程描述信息(可自定义此描述)，最大长度1000个字符。
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 调用方自定义的个性化字段(可自定义此名称)，并以base64方式编码，支持的最大数据大小为1000长度。
	// 
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。回调的相关说明可参考开发者中心的回调通知模块。
	CustomerData *string `json:"CustomerData,omitnil" name:"CustomerData"`

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
	CustomShowMap *string `json:"CustomShowMap,omitnil" name:"CustomShowMap"`

	// 合同流程的抄送人列表，最多可支持50个抄送人，抄送人可查看合同内容及签署进度，但无需参与合同签署。
	CcInfos []*CcInfo `json:"CcInfos,omitnil" name:"CcInfos"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li> **false**：（默认）不需要审批，直接签署。</li>
	// <li> **true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li> 如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li> 如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 若在创建签署流程时指定了关注人CcInfos，此参数可设定向关注人发送短信通知的类型：
	// <ul><li> **0** :合同发起时通知通知对方来查看合同（默认）</li>
	// <li> **1** : 签署完成后通知对方来查看合同</li></ul>
	CcNotifyType *int64 `json:"CcNotifyType,omitnil" name:"CcNotifyType"`

	// 个人自动签名的使用场景包括以下, 个人自动签署(即ApproverType设置成个人自动签署时)业务此值必传：
	// <ul><li> **E_PRESCRIPTION_AUTO_SIGN**：处方单（医疗自动签）  </li></ul>
	// 注: `个人自动签名场景是白名单功能，使用前请与对接的客户经理联系沟通。`
	AutoSignScene *string `json:"AutoSignScene,omitnil" name:"AutoSignScene"`
}

type FlowResourceUrlInfo struct {
	// 合同流程的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 对应的合同流程的PDF下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceUrlInfos []*ResourceUrlInfo `json:"ResourceUrlInfos,omitnil" name:"ResourceUrlInfos"`
}

type FormField struct {
	// 控件填充值，ComponentType和传入值格式对应关系如下：
	// <ul>
	// <li>TEXT - 普通文本控件，需输入文本字符串；</li>
	// <li>MULTI_LINE_TEXT - 多行文本控件，需输入文本字符串；</li>
	// <li>CHECK_BOX - 勾选框控件，若选中需填写ComponentValue，填写 true或者 false 字符串；</li>
	// <li>FILL_IMAGE - 图片控件，需填写ComponentValue为图片的资源 ID；</li>
	// <li>DYNAMIC_TABLE - 动态表格控件；</li>
	// <li>ATTACHMENT - 附件控件，需填写ComponentValue为附件图片的资源 ID列表，以逗号分割；</li>
	// <li>DATE - 日期控件；格式为 <b>xxxx年xx月xx日</b> 字符串；</li>
	// <li>DISTRICT - 省市区行政区控件，需填写ComponentValue为省市区行政区字符串内容；</li>
	// </ul>
	ComponentValue *string `json:"ComponentValue,omitnil" name:"ComponentValue"`

	// 表单域或控件的ID，跟ComponentName二选一，不能全为空；
	// CreateFlowsByTemplates 接口不使用此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentId *string `json:"ComponentId,omitnil" name:"ComponentId"`

	// 控件的名字，跟ComponentId二选一，不能全为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// 是否锁定模板控件值，锁定后无法修改（用于嵌入式发起合同），true-锁定，false-不锁定
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockComponentValue *bool `json:"LockComponentValue,omitnil" name:"LockComponentValue"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 流程合同ID列表,  可将这些流程ID组织成合同组的形式, 下载时候每个文件夹会是一个zip压缩包,  每个文件夹对多20个合同, 所有文件夹最多50个合同
	// 
	// 如下列组织形式,  控制台下载页面点击下载按钮后, 会生成**2023采购合同.zip**和**2023入职合同.zip** 两个下载任务(注:`部分浏览器需要授权或不支持创建多下载任务`)
	// 
	// **2023采购合同.zip**压缩包会有`yDwivUUckpor6wtoUuogwQHCAB0ES0pQ`和`yDwi8UUckpo5fz9cUqI6nGwcuTvt9YSh`两个合同的文件
	// **2023入职合同.zip** 压缩包会有`yDwivUUckpor6wobUuogwQHvdGfvDi5K`的文件
	// 
	// ![image](	https://dyn.ess.tencent.cn/guide/capi/channel_GetDownloadFlowUrl_DownLoadFlows.png)
	DownLoadFlows []*DownloadFlowInfo `json:"DownLoadFlows,omitnil" name:"DownLoadFlows"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 流程合同ID列表,  可将这些流程ID组织成合同组的形式, 下载时候每个文件夹会是一个zip压缩包,  每个文件夹对多20个合同, 所有文件夹最多50个合同
	// 
	// 如下列组织形式,  控制台下载页面点击下载按钮后, 会生成**2023采购合同.zip**和**2023入职合同.zip** 两个下载任务(注:`部分浏览器需要授权或不支持创建多下载任务`)
	// 
	// **2023采购合同.zip**压缩包会有`yDwivUUckpor6wtoUuogwQHCAB0ES0pQ`和`yDwi8UUckpo5fz9cUqI6nGwcuTvt9YSh`两个合同的文件
	// **2023入职合同.zip** 压缩包会有`yDwivUUckpor6wobUuogwQHvdGfvDi5K`的文件
	// 
	// ![image](	https://dyn.ess.tencent.cn/guide/capi/channel_GetDownloadFlowUrl_DownLoadFlows.png)
	DownLoadFlows []*DownloadFlowInfo `json:"DownLoadFlows,omitnil" name:"DownLoadFlows"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	DownLoadUrl *string `json:"DownLoadUrl,omitnil" name:"DownLoadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type ModifyExtendedServiceRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业标识: Agent. ProxyOperator.OpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	//   扩展服务类型
	// <ul>
	// <li>AUTO_SIGN             企业自动签（自动签署）</li>
	// <li>  OVERSEA_SIGN          企业与港澳台居民*签署合同</li>
	// <li>  MOBILE_CHECK_APPROVER 使用手机号验证签署方身份</li>
	// <li> PAGING_SEAL           骑缝章</li>
	// <li> DOWNLOAD_FLOW         授权渠道下载合同 </li>
	// <li>AGE_LIMIT_EXPANSION 拓宽签署方年龄限制</li>
	// </ul>
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 操作类型 
	// OPEN:开通 
	// CLOSE:关闭
	Operate *string `json:"Operate,omitnil" name:"Operate"`

	// 链接跳转类型，支持以下类型
	// <ul>
	// <li>WEIXINAPP : 短链直接跳转到电子签小程序  (默认值)</li>
	// <li>APP : 第三方APP或小程序跳转电子签小程序</li>
	// </ul>
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`
}

type ModifyExtendedServiceRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>渠道应用标识:  Agent.ProxyOrganizationOpenId</li>
	// <li>第三方平台子客企业标识: Agent. ProxyOperator.OpenId</li>
	// <li>第三方平台子客企业中的员工标识: Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	//   扩展服务类型
	// <ul>
	// <li>AUTO_SIGN             企业自动签（自动签署）</li>
	// <li>  OVERSEA_SIGN          企业与港澳台居民*签署合同</li>
	// <li>  MOBILE_CHECK_APPROVER 使用手机号验证签署方身份</li>
	// <li> PAGING_SEAL           骑缝章</li>
	// <li> DOWNLOAD_FLOW         授权渠道下载合同 </li>
	// <li>AGE_LIMIT_EXPANSION 拓宽签署方年龄限制</li>
	// </ul>
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 操作类型 
	// OPEN:开通 
	// CLOSE:关闭
	Operate *string `json:"Operate,omitnil" name:"Operate"`

	// 链接跳转类型，支持以下类型
	// <ul>
	// <li>WEIXINAPP : 短链直接跳转到电子签小程序  (默认值)</li>
	// <li>APP : 第三方APP或小程序跳转电子签小程序</li>
	// </ul>
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`
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
	// 操作跳转链接，有效期24小时
	// 若操作时没有返回跳转链接，表示无需跳转操作，此时会直接开通/关闭服务。
	// 
	// 当操作类型是 OPEN 且 扩展服务类型是  AUTO_SIGN 或 DOWNLOAD_FLOW 或者 OVERSEA_SIGN 时返回操作链接，
	// 返回的链接需要平台方自行触达超管或法人，超管或法人点击链接完成服务开通操作。
	OperateUrl *string `json:"OperateUrl,omitnil" name:"OperateUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type OccupiedSeal struct {
	// 电子印章编号
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 电子印章名称
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 电子印章授权时间戳，单位秒
	CreateOn *int64 `json:"CreateOn,omitnil" name:"CreateOn"`

	// 电子印章授权人，电子签的UserId
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// 电子印章策略Id
	SealPolicyId *string `json:"SealPolicyId,omitnil" name:"SealPolicyId"`

	// 印章状态，有以下六种：CHECKING（审核中）SUCCESS（已启用）FAIL（审核拒绝）CHECKING-SADM（待超管审核）DISABLE（已停用）STOPPED（已终止）
	SealStatus *string `json:"SealStatus,omitnil" name:"SealStatus"`

	// 审核失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitnil" name:"FailReason"`

	// 印章图片url，5分钟内有效
	Url *string `json:"Url,omitnil" name:"Url"`

	// 印章类型，OFFICIAL-企业公章，CONTRACT-合同专用章，LEGAL_PERSON_SEAL-法人章
	SealType *string `json:"SealType,omitnil" name:"SealType"`

	// 用印申请是否为永久授权
	IsAllTime *bool `json:"IsAllTime,omitnil" name:"IsAllTime"`

	// 授权人列表
	AuthorizedUsers []*AuthorizedUser `json:"AuthorizedUsers,omitnil" name:"AuthorizedUsers"`
}

// Predefined struct for user
type OperateChannelTemplateRequestParams struct {
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>第三方平台子客企业中的员工标识: Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作类型，可取值如下:
	// <ul>
	// <li>SELECT:  查询</li>
	// <li>DELETE:  删除</li>
	// <li>UPDATE: 更新</li>
	// </ul>
	OperateType *string `json:"OperateType,omitnil" name:"OperateType"`

	// 合同模板ID，为32位字符串。
	// 注: ` 此处为第三方应用平台模板库模板ID，非子客模板ID`
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 第三方平台子客企业的唯一标识，支持批量(用,分割)，
	ProxyOrganizationOpenIds *string `json:"ProxyOrganizationOpenIds,omitnil" name:"ProxyOrganizationOpenIds"`

	// 模板可见范围, 可以设置的值如下:
	// 
	// **all**: 所有本第三方应用合作企业可见
	// **part**: 指定的本第三方应用合作企业
	// 
	// 对应控制台的位置
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/68b97812c68d6af77a5991e3bff5c790.png)
	AuthTag *string `json:"AuthTag,omitnil" name:"AuthTag"`

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
	Available *int64 `json:"Available,omitnil" name:"Available"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type OperateChannelTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括渠道应用标识、第三方平台子客企业标识及第三方平台子客企业中的员工标识等内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	// 
	// 此接口下面信息必填。
	// <ul>
	// <li>第三方平台子客企业中的员工标识: Agent.AppId</li>
	// </ul>
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作类型，可取值如下:
	// <ul>
	// <li>SELECT:  查询</li>
	// <li>DELETE:  删除</li>
	// <li>UPDATE: 更新</li>
	// </ul>
	OperateType *string `json:"OperateType,omitnil" name:"OperateType"`

	// 合同模板ID，为32位字符串。
	// 注: ` 此处为第三方应用平台模板库模板ID，非子客模板ID`
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 第三方平台子客企业的唯一标识，支持批量(用,分割)，
	ProxyOrganizationOpenIds *string `json:"ProxyOrganizationOpenIds,omitnil" name:"ProxyOrganizationOpenIds"`

	// 模板可见范围, 可以设置的值如下:
	// 
	// **all**: 所有本第三方应用合作企业可见
	// **part**: 指定的本第三方应用合作企业
	// 
	// 对应控制台的位置
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/68b97812c68d6af77a5991e3bff5c790.png)
	AuthTag *string `json:"AuthTag,omitnil" name:"AuthTag"`

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
	Available *int64 `json:"Available,omitnil" name:"Available"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 合同模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 描述模板可见性更改的结果。
	// <ul>
	// <li>all-success: 全部成功</li>
	// <li>part-success: 部分成功,失败的会在FailMessageList中展示</li>
	// <li>fail:全部失败, 失败的会在FailMessageList中展示</li>
	// </ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateResult *string `json:"OperateResult,omitnil" name:"OperateResult"`

	// 模板可见范围:
	// **all**: 所有本第三方应用合作企业可见
	// **part**: 指定的本第三方应用合作企业
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthTag *string `json:"AuthTag,omitnil" name:"AuthTag"`

	// 第三方平台子客企业标识列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyOrganizationOpenIds []*string `json:"ProxyOrganizationOpenIds,omitnil" name:"ProxyOrganizationOpenIds"`

	// 操作失败信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailMessageList []*AuthFailMessage `json:"FailMessageList,omitnil" name:"FailMessageList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type OrganizationInfo struct {
	// 用户在渠道的机构编号
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 机构在平台的编号
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 用户渠道
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 用户真实的IP
	//
	// Deprecated: ClientIp is deprecated.
	ClientIp *string `json:"ClientIp,omitnil" name:"ClientIp"`

	// 机构的代理IP
	//
	// Deprecated: ProxyIp is deprecated.
	ProxyIp *string `json:"ProxyIp,omitnil" name:"ProxyIp"`
}

type PdfVerifyResult struct {
	// 验签结果详情，每个签名域对应的验签结果。状态值如下
	// <ul><li> **1** :验签成功，在电子签签署</li>
	// <li> **2** :验签成功，在其他平台签署</li>
	// <li> **3** :验签失败</li>
	// <li> **4** :pdf文件没有签名域</li>
	// <li> **5** :文件签名格式错误</li></ul>
	VerifyResult *int64 `json:"VerifyResult,omitnil" name:"VerifyResult"`

	// 签署平台
	// 如果文件是在腾讯电子签平台签署，则为**腾讯电子签**，
	// 如果文件不在腾讯电子签平台签署，则为**其他平台**。
	SignPlatform *string `json:"SignPlatform,omitnil" name:"SignPlatform"`

	// 申请证书的主体的名字
	// 
	// 如果是在腾讯电子签平台签署, 则对应的主体的名字个数如下
	// **企业**:  ESS@企业名称@平台生成的数字编码
	// **个人**: ESS@个人姓名@证件号@平台生成的数字编码
	// 
	// 如果在其他平台签署的, 主体的名字参考其他平台的说明
	SignerName *string `json:"SignerName,omitnil" name:"SignerName"`

	// 签署时间的Unix时间戳，单位毫秒
	SignTime *int64 `json:"SignTime,omitnil" name:"SignTime"`

	// 证书签名算法,  如SHA1withRSA等算法
	SignAlgorithm *string `json:"SignAlgorithm,omitnil" name:"SignAlgorithm"`

	// CA供应商下发给用户的证书编号
	// 
	// 注意：`腾讯电子签接入多家CA供应商以提供容灾能力，不同CA下发的证书编号区别较大，但基本都是由数字和字母组成，长度在200以下`。
	CertSn *string `json:"CertSn,omitnil" name:"CertSn"`

	// 证书起始时间的Unix时间戳，单位毫秒
	CertNotBefore *int64 `json:"CertNotBefore,omitnil" name:"CertNotBefore"`

	// 证书过期时间的时间戳，单位毫秒
	CertNotAfter *int64 `json:"CertNotAfter,omitnil" name:"CertNotAfter"`

	// 签名类型, 保留字段, 现在全部为0
	// 
	SignType *int64 `json:"SignType,omitnil" name:"SignType"`

	// 签名域横坐标，单位px
	ComponentPosX *float64 `json:"ComponentPosX,omitnil" name:"ComponentPosX"`

	// 签名域纵坐标，单位px
	ComponentPosY *float64 `json:"ComponentPosY,omitnil" name:"ComponentPosY"`

	// 签名域宽度，单位px
	ComponentWidth *float64 `json:"ComponentWidth,omitnil" name:"ComponentWidth"`

	// 签名域高度，单位px
	ComponentHeight *float64 `json:"ComponentHeight,omitnil" name:"ComponentHeight"`

	// 签名域所在页码，1～N
	ComponentPage *int64 `json:"ComponentPage,omitnil" name:"ComponentPage"`
}

type Permission struct {
	// 权限名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 权限key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 权限类型 1前端，2后端
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 是否隐藏
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hide *int64 `json:"Hide,omitnil" name:"Hide"`

	// 数据权限标签 1:表示根节点，2:表示叶子结点
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataLabel *int64 `json:"DataLabel,omitnil" name:"DataLabel"`

	// 数据权限独有，1:关联其他模块鉴权，2:表示关联自己模块鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *int64 `json:"DataType,omitnil" name:"DataType"`

	// 数据权限独有，表示数据范围，1：全公司，2:部门及下级部门，3:自己
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataRange *int64 `json:"DataRange,omitnil" name:"DataRange"`

	// 关联权限, 表示这个功能权限要受哪个数据权限管控
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataTo *string `json:"DataTo,omitnil" name:"DataTo"`

	// 父级权限key
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentKey *string `json:"ParentKey,omitnil" name:"ParentKey"`

	// 是否选中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsChecked *bool `json:"IsChecked,omitnil" name:"IsChecked"`

	// 子权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*Permission `json:"Children,omitnil" name:"Children"`
}

type PermissionGroup struct {
	// 权限组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 权限组key
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupKey *string `json:"GroupKey,omitnil" name:"GroupKey"`

	// 是否隐藏分组，0否1是
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hide *int64 `json:"Hide,omitnil" name:"Hide"`

	// 权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions []*Permission `json:"Permissions,omitnil" name:"Permissions"`
}

// Predefined struct for user
type PrepareFlowsRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 多个合同（签署流程）信息，最大支持20个签署流程。
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil" name:"FlowInfos"`

	// 操作完成后的跳转地址，最大长度200
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type PrepareFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 多个合同（签署流程）信息，最大支持20个签署流程。
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil" name:"FlowInfos"`

	// 操作完成后的跳转地址，最大长度200
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	ConfirmUrl *string `json:"ConfirmUrl,omitnil" name:"ConfirmUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 员工的唯一标识(即OpenId),  定义Agent中的OpenId一样, 可以参考<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#agent" target="_blank">Agent结构体</a>
	Id *string `json:"Id,omitnil" name:"Id"`

	// 员工的姓名，最大长度50个字符
	// 员工的姓名将用于身份认证和电子签名，请确保填写的姓名为真实姓名，而非昵称等代名。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 签署方经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 经办人证件号
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 员工的手机号，支持国内手机号11位数字(无需加+86前缀或其他字符)，不支持海外手机号。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 预先分配员工的角色, 可以分配的角色如下:
	// <table> <thead> <tr> <th>可以分配的角色</th> <th>角色名称</th> <th>角色描述</th> </tr> </thead> <tbody> <tr> <td>admin</td> <td>业务管理员（IT 系统负责人，e.g. CTO）</td> <td>有企业合同模块、印章模块、模板模块等全量功能及数据权限。</td> </tr> <tr> <td>channel-normal-operator</td> <td>经办人（企业法务负责人）</td> <td>有发起合同、签署合同（含填写、拒签）、撤销合同、持有印章等权限能力，可查看企业所有合同数据。</td> </tr> <tr> <td>channel-sales-man</td> <td>业务员（一般为销售员、采购员）</td> <td>有发起合同、签署合同（含填写、拒签）、撤销合同、持有印章等权限能力，可查看自己相关所有合同数据。</td> </tr> </tbody> </table>
	DefaultRole *string `json:"DefaultRole,omitnil" name:"DefaultRole"`
}

type Recipient struct {
	// 合同参与方的角色ID
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 参与者类型, 可以选择的类型如下:
	// <ul><li> **ENTERPRISE** :此角色为企业参与方</li>
	// <li> **INDIVIDUAL** :此角色为个人参与方</li>
	// <li> **PROMOTER** :此角色是发起方</li></ul>
	RecipientType *string `json:"RecipientType,omitnil" name:"RecipientType"`

	// 合同参与方的角色描述，长度不能超过100，只能由中文、字母、数字和下划线组成。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 合同参与方的角色名字，长度不能超过20，只能由中文、字母、数字和下划线组成。
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// 是否需要校验，
	// true-是，
	// false-否
	RequireValidation *bool `json:"RequireValidation,omitnil" name:"RequireValidation"`

	// 是否必须填写，
	// true-是，
	// false-否
	RequireSign *bool `json:"RequireSign,omitnil" name:"RequireSign"`

	// 内部字段，签署类型
	SignType *int64 `json:"SignType,omitnil" name:"SignType"`

	// 签署顺序：数字越小优先级越高
	RoutingOrder *int64 `json:"RoutingOrder,omitnil" name:"RoutingOrder"`

	// 是否是发起方，
	// true-是 
	// false-否
	IsPromoter *bool `json:"IsPromoter,omitnil" name:"IsPromoter"`

	// 签署人查看合同校验方式, 支持的类型如下:
	// <ul><li> 1 :实名认证查看</li>
	// <li> 2 :手机号校验查看</li></ul>
	ApproverVerifyTypes []*int64 `json:"ApproverVerifyTypes,omitnil" name:"ApproverVerifyTypes"`

	// 签署人进行合同签署时的认证方式，支持的类型如下:
	// <ul><li> 1 :人脸认证</li>
	// <li> 2 :签署密码</li>
	// <li> 3 :运营商三要素认证</li>
	// <li> 4 :UKey认证</li></ul>
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil" name:"ApproverSignTypes"`
}

type RecipientComponentInfo struct {
	// 参与方的角色ID
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 参与方填写状态
	// 
	// <ul><li> **0** : 还没有填写</li>
	// <li> **1** : 已经填写</li></ul>
	RecipientFillStatus *string `json:"RecipientFillStatus,omitnil" name:"RecipientFillStatus"`

	// 此角色是否是发起方角色
	// 
	// <ul><li> **true** : 是发起方角色</li>
	// <li> **false** : 不是发起方角色</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPromoter *bool `json:"IsPromoter,omitnil" name:"IsPromoter"`

	// 此角色的填写控件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components []*FilledComponent `json:"Components,omitnil" name:"Components"`
}

type ReleasedApprover struct {
	// 签署人在原合同签署人列表中的顺序序号(从0开始，按顺序依次递增)。</br>
	// 可以通过<a href="https://qian.tencent.com/developers/partnerApis/flows/DescribeFlowDetailInfo" target="_blank">DescribeFlowDetailInfo</a>接口查看原流程中的签署人列表。
	ApproverNumber *uint64 `json:"ApproverNumber,omitnil" name:"ApproverNumber"`

	// 指定签署人类型，目前支持
	// <ul><li> **ORGANIZATION**：企业(默认值)</li>
	// <li> **ENTERPRISESERVER**：企业静默签</li></ul>
	ApproverType *string `json:"ApproverType,omitnil" name:"ApproverType"`

	// 签署人姓名，最大长度50个字。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 签署方经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证(默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成(如存在X，请大写)。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母(但“I”、“O”除外)，后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 签署人手机号。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 如果签署方是企业签署方(approverType = 0 或者 approverType = 3)， 则企业名称必填。
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 第三方平台子客企业的唯一标识，定义Agent中的ProxyOrganizationOpenId一样, 可以参考<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#agent" target="_blank">Agent结构体</a>。</br>
	// 当为子客企业指定经办人时，此OrganizationOpenId必传。
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 第三方平台子客企业员工的唯一标识，长度不能超过64，只能由字母和数字组成。</br>
	// 当签署方为同一第三方平台下的员工时，此OpenId必传。
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 签署控件类型，支持自定义企业签署方的签署控件类型
	// <ul><li> **SIGN_SEAL**：默认为印章控件类型(默认值)</li>
	// <li> **SIGN_SIGNATURE**：手写签名控件类型</li></ul>
	ApproverSignComponentType *string `json:"ApproverSignComponentType,omitnil" name:"ApproverSignComponentType"`

	// 参与方在合同中的角色是按照创建合同的时候来排序的，解除协议默认会将第一个参与人叫`甲方`,第二个叫`乙方`,  第三个叫`丙方`，以此类推。</br>
	// 如果需改动此参与人的角色名字，可用此字段指定，由汉字,英文字符,数字组成，最大20个字。
	ApproverSignRole *string `json:"ApproverSignRole,omitnil" name:"ApproverSignRole"`
}

type RelieveInfo struct {
	// 解除理由，最大支持200个字
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 解除后仍然有效的条款，保留条款，最大支持200个字
	RemainInForceItem *string `json:"RemainInForceItem,omitnil" name:"RemainInForceItem"`

	// 原合同事项处理-费用结算，最大支持200个字
	OriginalExpenseSettlement *string `json:"OriginalExpenseSettlement,omitnil" name:"OriginalExpenseSettlement"`

	// 原合同事项处理-其他事项，最大支持200个字
	OriginalOtherSettlement *string `json:"OriginalOtherSettlement,omitnil" name:"OriginalOtherSettlement"`

	// 其他约定，最大支持200个字
	OtherDeals *string `json:"OtherDeals,omitnil" name:"OtherDeals"`
}

type RemindFlowRecords struct {
	// 合同流程是否可以催办： true - 可以，false - 不可以。 若无法催办，将返回RemindMessage以解释原因。	
	CanRemind *bool `json:"CanRemind,omitnil" name:"CanRemind"`

	// 合同流程ID，为32位字符串。	
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 在合同流程无法催办的情况下，系统将返回RemindMessage以阐述原因。	
	RemindMessage *string `json:"RemindMessage,omitnil" name:"RemindMessage"`
}

type ResourceUrlInfo struct {
	// 资源链接地址，过期时间5分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type SignQrCode struct {
	// 二维码ID，为32位字符串。	
	// 
	// 注: 需要保留此二维码ID, 用于后序通过<a href="https://qian.tencent.com/developers/partnerApis/templates/ChannelCancelMultiFlowSignQRCode" target="_blank">取消一码多扫二维码</a>关闭这个二维码的签署功能。	
	QrCodeId *string `json:"QrCodeId,omitnil" name:"QrCodeId"`

	// 二维码URL，可通过转换二维码的工具或代码组件将此URL转化为二维码，以便用户扫描进行流程签署。	
	QrCodeUrl *string `json:"QrCodeUrl,omitnil" name:"QrCodeUrl"`

	// 二维码的有截止时间，格式为Unix标准时间戳（秒），可以通过入参的QrEffectiveDay来设置有效期，默认为7天有效期。 
	// 一旦超过二维码的有效期限，该二维码将自动失效。	
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`
}

type SignUrl struct {
	// 跳转至电子签名小程序签署的链接地址。 适用于客户端APP及小程序直接唤起电子签名小程序。	
	AppSignUrl *string `json:"AppSignUrl,omitnil" name:"AppSignUrl"`

	// 签署链接有效时间，格式类似"2022-08-05 15:55:01"	
	EffectiveTime *string `json:"EffectiveTime,omitnil" name:"EffectiveTime"`

	// 跳转至电子签名小程序签署的链接地址，格式类似于https://essurl.cn/xxx。 打开此链接将会展示H5中间页面，随后唤起电子签名小程序以进行合同签署。	
	HttpSignUrl *string `json:"HttpSignUrl,omitnil" name:"HttpSignUrl"`
}

type SignUrlInfo struct {
	// 签署链接，过期时间为90天
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignUrl *string `json:"SignUrl,omitnil" name:"SignUrl"`

	// 合同过期时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 当流程为顺序签署此参数有效时，数字越小优先级越高，暂不支持并行签署 可选
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignOrder *int64 `json:"SignOrder,omitnil" name:"SignOrder"`

	// 签署人编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"SignId,omitnil" name:"SignId"`

	// 自定义用户编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: CustomUserId is deprecated.
	CustomUserId *string `json:"CustomUserId,omitnil" name:"CustomUserId"`

	// 用户姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 用户手机号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 签署参与者机构名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 参与者类型, 类型如下:
	// **ORGANIZATION**:企业经办人
	// **PERSON**: 自然人
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverType *string `json:"ApproverType,omitnil" name:"ApproverType"`

	// 经办人身份证号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 签署链接对应流程Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 企业经办人 用户在渠道的编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 合同组签署链接对应的合同组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 二维码，在生成动态签署人跳转封面页链接时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignQrcodeUrl *string `json:"SignQrcodeUrl,omitnil" name:"SignQrcodeUrl"`
}

type Staff struct {
	// 员工在电子签平台的用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 显示的员工名
	DisplayName *string `json:"DisplayName,omitnil" name:"DisplayName"`

	// 员工手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 员工邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil" name:"Email"`

	// 员工在第三方应用平台的用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 员工角色
	// 注意：此字段可能返回 null，表示取不到有效值。
	Roles []*StaffRole `json:"Roles,omitnil" name:"Roles"`

	// 员工部门
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitnil" name:"Department"`

	// 员工是否实名
	Verified *bool `json:"Verified,omitnil" name:"Verified"`

	// 员工创建时间戳，单位秒
	CreatedOn *int64 `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 员工实名时间戳，单位秒
	VerifiedOn *int64 `json:"VerifiedOn,omitnil" name:"VerifiedOn"`

	// 员工是否离职：0-未离职，1-离职
	QuiteJob *int64 `json:"QuiteJob,omitnil" name:"QuiteJob"`
}

type StaffRole struct {
	// 角色id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`
}

type SyncFailReason struct {
	// 企业员工标识(即OpenId)
	Id *string `json:"Id,omitnil" name:"Id"`

	// 新增员工或者员工离职失败原因, 可能存证ID不符合规范、证件号码不合法等原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作类型，对应的操作
	// <ul><li> **CREATE** :新增员工</li>
	// <li> **UPDATE** :修改员工</li>
	// <li> **RESIGN** :离职员工</li></ul>
	OperatorType *string `json:"OperatorType,omitnil" name:"OperatorType"`

	// 员工信息列表，最多支持200个
	ProxyOrganizationOperators []*ProxyOrganizationOperator `json:"ProxyOrganizationOperators,omitnil" name:"ProxyOrganizationOperators"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作类型，对应的操作
	// <ul><li> **CREATE** :新增员工</li>
	// <li> **UPDATE** :修改员工</li>
	// <li> **RESIGN** :离职员工</li></ul>
	OperatorType *string `json:"OperatorType,omitnil" name:"OperatorType"`

	// 员工信息列表，最多支持200个
	ProxyOrganizationOperators []*ProxyOrganizationOperator `json:"ProxyOrganizationOperators,omitnil" name:"ProxyOrganizationOperators"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 同步失败员工ID及其失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedList []*SyncFailReason `json:"FailedList,omitnil" name:"FailedList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 第三方平台子客企业名称，请确认该名称与企业营业执照中注册的名称一致。
	// 注: `如果名称中包含英文括号()，请使用中文括号（）代替。`
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil" name:"ProxyOrganizationName"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitnil" name:"BusinessLicense"`

	// 第三方平台子客企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil" name:"UniformSocialCreditCode"`

	// 第三方平台子客企业法定代表人的名字
	ProxyLegalName *string `json:"ProxyLegalName,omitnil" name:"ProxyLegalName"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 第三方平台子客企业法定代表人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证 (默认值)</li></ul>
	// 注: `现在仅支持ID_CARD居民身份证类型`
	ProxyLegalIdCardType *string `json:"ProxyLegalIdCardType,omitnil" name:"ProxyLegalIdCardType"`

	// 第三方平台子客企业法定代表人的证件号码, 应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li></ul>
	ProxyLegalIdCardNumber *string `json:"ProxyLegalIdCardNumber,omitnil" name:"ProxyLegalIdCardNumber"`

	// 第三方平台子客企业详细住所，最大长度500个字符
	// 
	// 注：`需要符合省市区详情的格式例如： XX省XX市XX区街道具体地址`
	ProxyAddress *string `json:"ProxyAddress,omitnil" name:"ProxyAddress"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 第三方平台子客企业名称，请确认该名称与企业营业执照中注册的名称一致。
	// 注: `如果名称中包含英文括号()，请使用中文括号（）代替。`
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil" name:"ProxyOrganizationName"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitnil" name:"BusinessLicense"`

	// 第三方平台子客企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil" name:"UniformSocialCreditCode"`

	// 第三方平台子客企业法定代表人的名字
	ProxyLegalName *string `json:"ProxyLegalName,omitnil" name:"ProxyLegalName"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 第三方平台子客企业法定代表人的证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证 (默认值)</li></ul>
	// 注: `现在仅支持ID_CARD居民身份证类型`
	ProxyLegalIdCardType *string `json:"ProxyLegalIdCardType,omitnil" name:"ProxyLegalIdCardType"`

	// 第三方平台子客企业法定代表人的证件号码, 应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li></ul>
	ProxyLegalIdCardNumber *string `json:"ProxyLegalIdCardNumber,omitnil" name:"ProxyLegalIdCardNumber"`

	// 第三方平台子客企业详细住所，最大长度500个字符
	// 
	// 注：`需要符合省市区详情的格式例如： XX省XX市XX区街道具体地址`
	ProxyAddress *string `json:"ProxyAddress,omitnil" name:"ProxyAddress"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 任务状态：READY - 任务已完成；NOTREADY - 任务未完成；
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *string `json:"TaskStatus,omitnil" name:"TaskStatus"`
}

type TemplateInfo struct {
	// 模板ID，模板的唯一标识
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 模板名
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 模板描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 模板的填充控件列表
	Components []*Component `json:"Components,omitnil" name:"Components"`

	// 模板中的签署参与方列表
	Recipients []*Recipient `json:"Recipients,omitnil" name:"Recipients"`

	// 模板中的签署控件列表
	SignComponents []*Component `json:"SignComponents,omitnil" name:"SignComponents"`

	// 模板类型：1-静默签；3-普通模板
	TemplateType *int64 `json:"TemplateType,omitnil" name:"TemplateType"`

	// 是否是发起人 ,已弃用
	//
	// Deprecated: IsPromoter is deprecated.
	IsPromoter *bool `json:"IsPromoter,omitnil" name:"IsPromoter"`

	// 模板的创建者信息，电子签系统用户ID
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// 模板创建的时间戳，格式为Unix标准时间戳（秒）
	CreatedOn *int64 `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 模板的H5预览链接,有效期5分钟。
	// 可以通过浏览器打开此链接预览模板，或者嵌入到iframe中预览模板。
	// （此功能开放需要联系客户经理）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewUrl *string `json:"PreviewUrl,omitnil" name:"PreviewUrl"`

	// 第三方应用集成-模板PDF文件链接，有效期5分钟。
	// 请求参数WithPdfUrl=true时返回
	// （此功能开放需要联系客户经理）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PdfUrl *string `json:"PdfUrl,omitnil" name:"PdfUrl"`

	// 本模板关联的第三方应用平台企业模板ID
	ChannelTemplateId *string `json:"ChannelTemplateId,omitnil" name:"ChannelTemplateId"`

	// 本模板关联的三方应用平台平台企业模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelTemplateName *string `json:"ChannelTemplateName,omitnil" name:"ChannelTemplateName"`

	// 0-需要子客企业手动领取平台企业的模板(默认); 
	// 1-平台自动设置子客模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelAutoSave *int64 `json:"ChannelAutoSave,omitnil" name:"ChannelAutoSave"`

	// 模板版本，全数字字符。
	// 默认为空，初始版本为yyyyMMdd001。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateVersion *string `json:"TemplateVersion,omitnil" name:"TemplateVersion"`

	// 模板可用状态：
	// 1启用（默认）
	// 2停用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Available *int64 `json:"Available,omitnil" name:"Available"`
}

type UploadFile struct {
	// Base64编码后的文件内容
	FileBody *string `json:"FileBody,omitnil" name:"FileBody"`

	// 文件名
	FileName *string `json:"FileName,omitnil" name:"FileName"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 
	// 文件对应业务类型,可以选择的类型如下
	// <ul><li> **TEMPLATE** : 此上传的文件用户生成合同模板，文件类型支持.pdf/.doc/.docx/.html格式，如果非pdf文件需要通过<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi" target="_blank">创建文件转换任务</a>转换后才能使用</li>
	// <li> **DOCUMENT** : 此文件用来发起合同流程，文件类型支持.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html，如果非pdf文件需要通过<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi" target="_blank">创建文件转换任务</a>转换后才能使用</li></ul>
	BusinessType *string `json:"BusinessType,omitnil" name:"BusinessType"`

	// 上传文件内容数组，最多支持上传20个文件。
	FileInfos []*UploadFile `json:"FileInfos,omitnil" name:"FileInfos"`

	// 操作者的信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 
	// 文件对应业务类型,可以选择的类型如下
	// <ul><li> **TEMPLATE** : 此上传的文件用户生成合同模板，文件类型支持.pdf/.doc/.docx/.html格式，如果非pdf文件需要通过<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi" target="_blank">创建文件转换任务</a>转换后才能使用</li>
	// <li> **DOCUMENT** : 此文件用来发起合同流程，文件类型支持.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html，如果非pdf文件需要通过<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi" target="_blank">创建文件转换任务</a>转换后才能使用</li></ul>
	BusinessType *string `json:"BusinessType,omitnil" name:"BusinessType"`

	// 上传文件内容数组，最多支持上传20个文件。
	FileInfos []*UploadFile `json:"FileInfos,omitnil" name:"FileInfos"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 文件资源ID数组，每个文件资源ID为32位字符串。
	// 建议开发者保存此资源ID，后续创建合同或创建合同流程需此资源ID。
	// 注:`有效期一个小时, 有效期内此文件id可以反复使用, 超过有效期无法使用`
	FileIds []*string `json:"FileIds,omitnil" name:"FileIds"`

	// 对应上传文件的下载链接，过期时间5分钟
	FileUrls []*string `json:"FileUrls,omitnil" name:"FileUrls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil" name:"ProxyOrganizationOpenId"`

	// 子客企业名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil" name:"ProxyOrganizationName"`

	// 对应的消耗日期, **如果是汇总数据则为1970-01-01**
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitnil" name:"Date"`

	// 消耗合同数量
	Usage *uint64 `json:"Usage,omitnil" name:"Usage"`

	// 撤回合同数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cancel *uint64 `json:"Cancel,omitnil" name:"Cancel"`

	// 消耗渠道
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowChannel *string `json:"FlowChannel,omitnil" name:"FlowChannel"`
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
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 内部参数，暂未开放使用
	//
	// Deprecated: Channel is deprecated.
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 内部参数，暂未开放使用
	//
	// Deprecated: CustomUserId is deprecated.
	CustomUserId *string `json:"CustomUserId,omitnil" name:"CustomUserId"`

	// 内部参数，暂未开放使用
	//
	// Deprecated: ClientIp is deprecated.
	ClientIp *string `json:"ClientIp,omitnil" name:"ClientIp"`

	// 内部参数，暂未开放使用
	//
	// Deprecated: ProxyIp is deprecated.
	ProxyIp *string `json:"ProxyIp,omitnil" name:"ProxyIp"`
}

type UserThreeFactor struct {
	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li></ul>
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码应为9位字符串，第1位为“C”，第2位为英文字母（但“I”、“O”除外），后7位为阿拉伯数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`
}

type WebThemeConfig struct {
	// 是否显示页面底部电子签logo，取值如下：
	// <ul><li> **true**：页面底部显示电子签logo</li>
	// <li> **false**：页面底部不显示电子签logo（默认）</li></ul>
	DisplaySignBrandLogo *bool `json:"DisplaySignBrandLogo,omitnil" name:"DisplaySignBrandLogo"`

	// 主题颜色：
	// 支持十六进制颜色值以及RGB格式颜色值，例如：#D54941，rgb(213, 73, 65)
	// <br/>
	WebEmbedThemeColor *string `json:"WebEmbedThemeColor,omitnil" name:"WebEmbedThemeColor"`
}