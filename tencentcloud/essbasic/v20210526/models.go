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
	// 应用的唯一标识。不同的业务系统可以采用不同的AppId，不同AppId下的数据是隔离的。可以由控制台开发者中心-应用集成自主生成。
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 第三方应用平台自定义，对应第三方平台子客企业的唯一标识。一个第三方平台子客企业主体与子客企业ProxyOrganizationOpenId是一一对应的，不可更改，不可重复使用。（例如，可以使用企业名称的hash值，或者社会统一信用代码的hash值，或者随机hash值，需要第三方应用平台保存），最大64位字符串
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil" name:"ProxyOrganizationOpenId"`

	// 第三方平台子客企业中的员工/经办人，通过第三方应用平台进入电子签完成实名、且被赋予相关权限后，可以参与到企业资源的管理或签署流程中。
	ProxyOperator *UserInfo `json:"ProxyOperator,omitnil" name:"ProxyOperator"`

	// 非必需参数，在第三方平台子客企业开通电子签后，会生成唯一的子客应用Id（ProxyAppId）用于代理调用时的鉴权，在子客开通的回调中获取。
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"SignId,omitnil" name:"SignId"`

	// 签署方角色编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 签署方角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverRoleName *string `json:"ApproverRoleName,omitnil" name:"ApproverRoleName"`
}

type ApproverOption struct {
	// 是否隐藏一键签署 默认false-不隐藏true-隐藏
	HideOneKeySign *bool `json:"HideOneKeySign,omitnil" name:"HideOneKeySign"`

	// 签署人信息补充类型，默认无需补充。
	// 
	// <ul><li> **1** : ( 动态签署人（可发起合同后再补充签署人信息）</li>
	// </ul>
	FillType *int64 `json:"FillType,omitnil" name:"FillType"`
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
	// 第三方应用平台的子客企业OpenId
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil" name:"ProxyOrganizationOpenId"`

	// 错误信息
	Message *string `json:"Message,omitnil" name:"Message"`
}

type AuthorizedUser struct {
	// 第三方应用平台的用户openid
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`
}

type AutoSignConfig struct {
	// 自动签开通个人用户的三要素
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 是否回调证书信息
	CertInfoCallback *bool `json:"CertInfoCallback,omitnil" name:"CertInfoCallback"`

	// 是否支持用户自定义签名印章
	UserDefineSeal *bool `json:"UserDefineSeal,omitnil" name:"UserDefineSeal"`

	// 是否需要回调的时候返回印章(签名) 图片的 base64
	SealImgCallback *bool `json:"SealImgCallback,omitnil" name:"SealImgCallback"`

	// 回调链接，如果渠道已经配置了，可以不传
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 开通时候的验证方式，取值：WEIXINAPP（微信人脸识别），INSIGHT（慧眼人脸认别），TELECOM（运营商三要素验证）。如果是小程序开通链接，支持传 WEIXINAPP / TELECOM。如果是 H5 开通链接，支持传 INSIGHT / TELECOM。默认值 WEIXINAPP / INSIGHT。
	VerifyChannels []*string `json:"VerifyChannels,omitnil" name:"VerifyChannels"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。一旦绑定后，将扣减购买的个人自动签账号许可一次（1年有效期），不可解绑释放。不传默认为绑定自动签账号许可。 0-绑定个人自动签账号许可，开通后将扣减购买的个人自动签账号许可一次
	LicenseType *int64 `json:"LicenseType,omitnil" name:"LicenseType"`
}

type BaseFlowInfo struct {
	// 合同流程名称
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同流程类型
	// <br/>客户自定义，用于合同分类展示
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 合同流程描述信息
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 合同流程截止时间，unix时间戳，单位秒
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 是否顺序签署(true:无序签,false:顺序签)
	// <br/>默认false，有序签署合同
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 是否打开智能添加填写区(默认开启，打开:"OPEN" 关闭："CLOSE")
	IntelligentStatus *string `json:"IntelligentStatus,omitnil" name:"IntelligentStatus"`

	// 填写控件内容
	FormFields []*FormField `json:"FormFields,omitnil" name:"FormFields"`

	// 本企业(发起方企业)是否需要签署审批
	// <br/>true：开启发起方签署审批
	// <br/>false：不开启发起方签署审批
	// <br/>开启后，使用ChannelCreateFlowSignReview接口提交审批结果，才能继续完成签署
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 用户流程自定义数据参数
	UserData *string `json:"UserData,omitnil" name:"UserData"`

	// 抄送人信息
	CcInfos []*CcInfo `json:"CcInfos,omitnil" name:"CcInfos"`

	// 是否需要开启发起方发起前审核
	// <br/>true：开启发起方发起前审核
	// <br/>false：不开启发起方发起前审核
	// <br/>当指定NeedCreateReview=true，则提交审核后，需要使用接口：ChannelCreateFlowSignReview，来完成发起前审核，审核通过后，可以继续查看，签署合同
	NeedCreateReview *bool `json:"NeedCreateReview,omitnil" name:"NeedCreateReview"`

	// 填写控件：文件发起使用
	Components []*Component `json:"Components,omitnil" name:"Components"`
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署流程Id数组，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 撤销理由,不超过200个字符
	CancelMessage *string `json:"CancelMessage,omitnil" name:"CancelMessage"`

	// 撤销理由自定义格式；选项：
	// 
	// - 0 默认格式
	// - 1 只保留身份信息：展示为【发起方】
	// - 2 保留身份信息+企业名称：展示为【发起方xxx公司】
	// - 3 保留身份信息+企业名称+经办人名称：展示为【发起方xxxx公司-经办人姓名】
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitnil" name:"CancelMessageFormat"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelBatchCancelFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署流程Id数组，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 撤销理由,不超过200个字符
	CancelMessage *string `json:"CancelMessage,omitnil" name:"CancelMessage"`

	// 撤销理由自定义格式；选项：
	// 
	// - 0 默认格式
	// - 1 只保留身份信息：展示为【发起方】
	// - 2 保留身份信息+企业名称：展示为【发起方xxx公司】
	// - 3 保留身份信息+企业名称+经办人名称：展示为【发起方xxxx公司-经办人姓名】
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
	// 签署流程批量撤销失败原因，错误信息与流程Id一一对应，成功为“”,失败则对应失败消息
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

// Predefined struct for user
type ChannelCancelFlowRequestParams struct {
	// 签署流程编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 撤回原因，最大不超过200字符
	CancelMessage *string `json:"CancelMessage,omitnil" name:"CancelMessage"`

	// 撤销理由自定义格式；选项：
	// 0 默认格式
	// 1 只保留身份信息：展示为【发起方】
	// 2 保留身份信息+企业名称：展示为【发起方xxx公司】
	// 3 保留身份信息+企业名称+经办人名称：展示为【发起方xxxx公司-经办人姓名】
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitnil" name:"CancelMessageFormat"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCancelFlowRequest struct {
	*tchttp.BaseRequest
	
	// 签署流程编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 撤回原因，最大不超过200字符
	CancelMessage *string `json:"CancelMessage,omitnil" name:"CancelMessage"`

	// 撤销理由自定义格式；选项：
	// 0 默认格式
	// 1 只保留身份信息：展示为【发起方】
	// 2 保留身份信息+企业名称：展示为【发起方xxx公司】
	// 3 保留身份信息+企业名称+经办人名称：展示为【发起方xxxx公司-经办人姓名】
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 二维码id
	QrCodeId *string `json:"QrCodeId,omitnil" name:"QrCodeId"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCancelMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 二维码id
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

	// 自动签场景: E_PRESCRIPTION_AUTO_SIGN 电子处方
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

	// 自动签场景: E_PRESCRIPTION_AUTO_SIGN 电子处方
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署流程Id数组
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCreateBatchCancelFlowUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署流程Id数组
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
	// 批量撤销url
	BatchCancelFlowUrl *string `json:"BatchCancelFlowUrl,omitnil" name:"BatchCancelFlowUrl"`

	// 签署流程批量撤销失败原因
	FailMessages []*string `json:"FailMessages,omitnil" name:"FailMessages"`

	// 签署撤销url过期时间-年月日-时分秒
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
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 领取的合同id列表
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCreateBoundFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 领取的合同id列表
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 资源类型 支持doc,docx,html,xls,xlsx,jpg,jpeg,png,bmp文件类型
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源名称，长度限制为256字符
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 文件Id，通过UploadFiles获取
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
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 资源类型 支持doc,docx,html,xls,xlsx,jpg,jpeg,png,bmp文件类型
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源名称，长度限制为256字符
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 文件Id，通过UploadFiles获取
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
	// 任务id
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
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 要生成WEB嵌入界面的类型, 可以选择的值如下: 
	// 
	// - CREATE_SEAL: 生成创建印章的嵌入页面
	// - CREATE_TEMPLATE：生成创建模板的嵌入页面
	// - MODIFY_TEMPLATE：生成修改模板的嵌入页面
	// - PREVIEW_TEMPLATE：生成预览模板的嵌入页面
	// - PREVIEW_FLOW：生成预览合同文档的嵌入页面
	// - PREVIEW_FLOW_DETAIL：生成预览合同详情的嵌入页面
	// - PREVIEW_SEAL_LIST：生成预览印章列表的嵌入页面
	// - PREVIEW_SEAL_DETAIL：生成预览印章详情的嵌入页面
	// - EXTEND_SERVICE：生成扩展服务的嵌入页面
	EmbedType *string `json:"EmbedType,omitnil" name:"EmbedType"`

	// WEB嵌入的业务资源ID
	// 
	// - 当EmbedType取值MODIFY_TEMPLATE，PREVIEW_TEMPLATE时需要填写模板id作为BusinessId
	// - 当EmbedType取值PREVIEW_FLOW，PREVIEW_FLOW_DETAIL时需要填写合同id作为BusinessId
	// - 当EmbedType取值PREVIEW_SEAL_DETAIL需要填写印章id作为BusinessId
	BusinessId *string `json:"BusinessId,omitnil" name:"BusinessId"`

	// 是否隐藏控件，只有预览模板时生效
	HiddenComponents *bool `json:"HiddenComponents,omitnil" name:"HiddenComponents"`

	// 渠道操作者信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCreateEmbedWebUrlRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 要生成WEB嵌入界面的类型, 可以选择的值如下: 
	// 
	// - CREATE_SEAL: 生成创建印章的嵌入页面
	// - CREATE_TEMPLATE：生成创建模板的嵌入页面
	// - MODIFY_TEMPLATE：生成修改模板的嵌入页面
	// - PREVIEW_TEMPLATE：生成预览模板的嵌入页面
	// - PREVIEW_FLOW：生成预览合同文档的嵌入页面
	// - PREVIEW_FLOW_DETAIL：生成预览合同详情的嵌入页面
	// - PREVIEW_SEAL_LIST：生成预览印章列表的嵌入页面
	// - PREVIEW_SEAL_DETAIL：生成预览印章详情的嵌入页面
	// - EXTEND_SERVICE：生成扩展服务的嵌入页面
	EmbedType *string `json:"EmbedType,omitnil" name:"EmbedType"`

	// WEB嵌入的业务资源ID
	// 
	// - 当EmbedType取值MODIFY_TEMPLATE，PREVIEW_TEMPLATE时需要填写模板id作为BusinessId
	// - 当EmbedType取值PREVIEW_FLOW，PREVIEW_FLOW_DETAIL时需要填写合同id作为BusinessId
	// - 当EmbedType取值PREVIEW_SEAL_DETAIL需要填写印章id作为BusinessId
	BusinessId *string `json:"BusinessId,omitnil" name:"BusinessId"`

	// 是否隐藏控件，只有预览模板时生效
	HiddenComponents *bool `json:"HiddenComponents,omitnil" name:"HiddenComponents"`

	// 渠道操作者信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateEmbedWebUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateEmbedWebUrlResponseParams struct {
	// 嵌入的web链接
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
type ChannelCreateFlowByFilesRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 均必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署流程名称，长度不超过200个字符
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 签署流程的描述，长度不超过1000个字符
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 签署流程签约方列表，最多不超过50个参与方
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitnil" name:"FlowApprovers"`

	// 签署文件资源Id列表，目前仅支持单个文件
	FileIds []*string `json:"FileIds,omitnil" name:"FileIds"`

	// 签署文件中的发起方的填写控件，需要在发起的时候进行填充
	Components []*Component `json:"Components,omitnil" name:"Components"`

	// 签署流程的签署截止时间。
	// 值为unix时间戳,精确到秒,不传默认为当前时间一年后
	// 不能早于当前时间
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 签署流程回调地址，长度不超过255个字符
	// 如果不传递回调地址， 则默认是配置应用号时候使用的回调地址
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 合同签署顺序类型
	// true - 无序签,
	// false - 顺序签，
	// 默认为false，即有序签署。
	// 有序签署时以传入FlowApprovers数组的顺序作为签署顺序
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 签署流程的类型，长度不超过255个字符
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitnil" name:"CustomShowMap"`

	// 业务信息，最大长度1000个字符。
	CustomerData *string `json:"CustomerData,omitnil" name:"CustomerData"`

	// 发起方企业的签署人进行签署操作是否需要企业内部审批。 若设置为true,审核结果需通过接口 ChannelCreateFlowSignReview 通知电子签，审核通过后，发起方企业签署人方可进行签署操作，否则会阻塞其签署操作。  注：企业可以通过此功能与企业内部的审批流程进行关联，支持手动、静默签署合同。
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证，用户手机号和参与方手机号（ApproverMobile）相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// 参数说明：可选人脸识别或手机号验证两种方式，若选择后者，未实名个人签署方在签署合同时，无需经过实名认证和意愿确认两次人脸识别，该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil" name:"ApproverVerifyType"`

	// 标识是否允许发起后添加控件。
	// 0为不允许
	// 1为允许。
	// 如果为1，创建的时候不能有签署控件，只能创建后添加。注意发起后添加控件功能不支持添加骑缝章和签批控件
	SignBeanTag *int64 `json:"SignBeanTag,omitnil" name:"SignBeanTag"`

	// 被抄送人信息列表
	CcInfos []*CcInfo `json:"CcInfos,omitnil" name:"CcInfos"`

	// 给关注人发送短信通知的类型，
	// 0-合同发起时通知 
	// 1-签署完成后通知
	CcNotifyType *int64 `json:"CcNotifyType,omitnil" name:"CcNotifyType"`

	// 个人自动签场景。发起自动签署时，需设置对应自动签署场景，目前仅支持场景：处方单-E_PRESCRIPTION_AUTO_SIGN
	AutoSignScene *string `json:"AutoSignScene,omitnil" name:"AutoSignScene"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCreateFlowByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 均必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署流程名称，长度不超过200个字符
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 签署流程的描述，长度不超过1000个字符
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	// 签署流程签约方列表，最多不超过50个参与方
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitnil" name:"FlowApprovers"`

	// 签署文件资源Id列表，目前仅支持单个文件
	FileIds []*string `json:"FileIds,omitnil" name:"FileIds"`

	// 签署文件中的发起方的填写控件，需要在发起的时候进行填充
	Components []*Component `json:"Components,omitnil" name:"Components"`

	// 签署流程的签署截止时间。
	// 值为unix时间戳,精确到秒,不传默认为当前时间一年后
	// 不能早于当前时间
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 签署流程回调地址，长度不超过255个字符
	// 如果不传递回调地址， 则默认是配置应用号时候使用的回调地址
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 合同签署顺序类型
	// true - 无序签,
	// false - 顺序签，
	// 默认为false，即有序签署。
	// 有序签署时以传入FlowApprovers数组的顺序作为签署顺序
	Unordered *bool `json:"Unordered,omitnil" name:"Unordered"`

	// 签署流程的类型，长度不超过255个字符
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitnil" name:"CustomShowMap"`

	// 业务信息，最大长度1000个字符。
	CustomerData *string `json:"CustomerData,omitnil" name:"CustomerData"`

	// 发起方企业的签署人进行签署操作是否需要企业内部审批。 若设置为true,审核结果需通过接口 ChannelCreateFlowSignReview 通知电子签，审核通过后，发起方企业签署人方可进行签署操作，否则会阻塞其签署操作。  注：企业可以通过此功能与企业内部的审批流程进行关联，支持手动、静默签署合同。
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证，用户手机号和参与方手机号（ApproverMobile）相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// 参数说明：可选人脸识别或手机号验证两种方式，若选择后者，未实名个人签署方在签署合同时，无需经过实名认证和意愿确认两次人脸识别，该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil" name:"ApproverVerifyType"`

	// 标识是否允许发起后添加控件。
	// 0为不允许
	// 1为允许。
	// 如果为1，创建的时候不能有签署控件，只能创建后添加。注意发起后添加控件功能不支持添加骑缝章和签批控件
	SignBeanTag *int64 `json:"SignBeanTag,omitnil" name:"SignBeanTag"`

	// 被抄送人信息列表
	CcInfos []*CcInfo `json:"CcInfos,omitnil" name:"CcInfos"`

	// 给关注人发送短信通知的类型，
	// 0-合同发起时通知 
	// 1-签署完成后通知
	CcNotifyType *int64 `json:"CcNotifyType,omitnil" name:"CcNotifyType"`

	// 个人自动签场景。发起自动签署时，需设置对应自动签署场景，目前仅支持场景：处方单-E_PRESCRIPTION_AUTO_SIGN
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
	// 合同签署流程ID
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
	// 每个子合同的发起所需的信息，数量限制2-50
	FlowFileInfos []*FlowFileInfo `json:"FlowFileInfos,omitnil" name:"FlowFileInfos"`

	// 合同组名称，长度不超过200个字符
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证
	// 参数说明：若选择后者，未实名的个人签署方查看合同时，无需进行人脸识别实名认证（但签署合同时仍然需要人脸实名），该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil" name:"ApproverVerifyType"`

	// 合同组的配置项信息包括：在合同组签署过程中，是否需要对每个子合同进行独立的意愿确认。
	FlowGroupOptions *FlowGroupOptions `json:"FlowGroupOptions,omitnil" name:"FlowGroupOptions"`

	// 操作者的信息，此参数不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelCreateFlowGroupByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 每个子合同的发起所需的信息，数量限制2-50
	FlowFileInfos []*FlowFileInfo `json:"FlowFileInfos,omitnil" name:"FlowFileInfos"`

	// 合同组名称，长度不超过200个字符
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证
	// 参数说明：若选择后者，未实名的个人签署方查看合同时，无需进行人脸识别实名认证（但签署合同时仍然需要人脸实名），该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitnil" name:"ApproverVerifyType"`

	// 合同组的配置项信息包括：在合同组签署过程中，是否需要对每个子合同进行独立的意愿确认。
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
	// 合同组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 子合同ID列表
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 均必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 每个子合同的发起所需的信息，数量限制2-50（合同组暂不支持抄送功能）
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil" name:"FlowInfos"`

	// 合同组名称，长度不超过200个字符
	FlowGroupName *string `json:"FlowGroupName,omitnil" name:"FlowGroupName"`
}

type ChannelCreateFlowGroupByTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 均必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 每个子合同的发起所需的信息，数量限制2-50（合同组暂不支持抄送功能）
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil" name:"FlowInfos"`

	// 合同组名称，长度不超过200个字符
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
	// 合同组ID
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 子合同ID列表
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署流程Id数组，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`
}

type ChannelCreateFlowRemindsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署流程Id数组，最多100个，超过100不处理
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
	// 合同催办详情信息
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 流程编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 流程签署人，其中Name和Mobile必传，其他可不传，ApproverType目前只支持PERSON类型的签署人，如果不传默认为该值。还需注意签署人只能有手写签名和时间类型的签署控件，其他类型的填写控件和签署控件暂时都未支持。
	FlowApproverInfos []*FlowApproverInfo `json:"FlowApproverInfos,omitnil" name:"FlowApproverInfos"`

	// 用户信息，暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 机构信息，暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 签署完之后的H5页面的跳转链接，此链接支持http://和https://，最大长度1000个字符。
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`
}

type ChannelCreateFlowSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 流程编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 流程签署人，其中Name和Mobile必传，其他可不传，ApproverType目前只支持PERSON类型的签署人，如果不传默认为该值。还需注意签署人只能有手写签名和时间类型的签署控件，其他类型的填写控件和签署控件暂时都未支持。
	FlowApproverInfos []*FlowApproverInfo `json:"FlowApproverInfos,omitnil" name:"FlowApproverInfos"`

	// 用户信息，暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 机构信息，暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 签署完之后的H5页面的跳转链接，此链接支持http://和https://，最大长度1000个字符。
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
	// 应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 模版ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 签署流程名称，最大长度200个字符。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 最大可发起签署流程份数
	// <br/>默认5份
	// <br/>备注：发起签署流程数量超过此上限后，二维码自动失效。
	MaxFlowNum *int64 `json:"MaxFlowNum,omitnil" name:"MaxFlowNum"`

	// 签署流程有效天数 默认7天 最高设置不超过30天
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitnil" name:"FlowEffectiveDay"`

	// 二维码有效天数 默认7天 最高设置不超过90天
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitnil" name:"QrEffectiveDay"`

	// 指定的签署二维码签署人
	// <br/>指定后，只允许知道的人操作和签署
	Restrictions []*ApproverRestriction `json:"Restrictions,omitnil" name:"Restrictions"`

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

	// 指定签署方经办人控件类型是个人印章签署控件（SIGN_SIGNATURE） 时，可选的签名方式。
	ApproverComponentLimitTypes []*ApproverComponentLimitType `json:"ApproverComponentLimitTypes,omitnil" name:"ApproverComponentLimitTypes"`
}

type ChannelCreateMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 模版ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 签署流程名称，最大长度200个字符。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 最大可发起签署流程份数
	// <br/>默认5份
	// <br/>备注：发起签署流程数量超过此上限后，二维码自动失效。
	MaxFlowNum *int64 `json:"MaxFlowNum,omitnil" name:"MaxFlowNum"`

	// 签署流程有效天数 默认7天 最高设置不超过30天
	FlowEffectiveDay *int64 `json:"FlowEffectiveDay,omitnil" name:"FlowEffectiveDay"`

	// 二维码有效天数 默认7天 最高设置不超过90天
	QrEffectiveDay *int64 `json:"QrEffectiveDay,omitnil" name:"QrEffectiveDay"`

	// 指定的签署二维码签署人
	// <br/>指定后，只允许知道的人操作和签署
	Restrictions []*ApproverRestriction `json:"Restrictions,omitnil" name:"Restrictions"`

	// 已废弃，回调配置统一使用企业应用管理-应用集成-第三方应用中的配置
	// <br/> 通过一码多扫二维码发起的合同，回调消息可参考文档 https://qian.tencent.com/developers/partner/callback_types_contracts_sign
	// <br/> 用户通过签署二维码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档 https://qian.tencent.com/developers/partner/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 限制二维码用户条件（已弃用）
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitnil" name:"ApproverRestrictions"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 指定签署方经办人控件类型是个人印章签署控件（SIGN_SIGNATURE） 时，可选的签名方式。
	ApproverComponentLimitTypes []*ApproverComponentLimitType `json:"ApproverComponentLimitTypes,omitnil" name:"ApproverComponentLimitTypes"`
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
	delete(f, "CallbackUrl")
	delete(f, "ApproverRestrictions")
	delete(f, "Operator")
	delete(f, "ApproverComponentLimitTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateMultiFlowSignQRCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateMultiFlowSignQRCodeResponseParams struct {
	// 签署二维码对象
	QrCode *SignQrCode `json:"QrCode,omitnil" name:"QrCode"`

	// 签署链接对象
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
type ChannelCreateOrganizationModifyQrCodeRequestParams struct {
	// 应用相关信息。 此接口Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type ChannelCreateOrganizationModifyQrCodeRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.AppId 必填。
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
	// 资源id，与ResourceType对应
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源类型，与ResourceId对应1：模板   2: 文件
	ResourceType *int64 `json:"ResourceType,omitnil" name:"ResourceType"`

	// 合同流程基础信息
	FlowInfo *BaseFlowInfo `json:"FlowInfo,omitnil" name:"FlowInfo"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程配置信息，用于配置发起合同时定制化
	FlowOption *CreateFlowOption `json:"FlowOption,omitnil" name:"FlowOption"`

	// 合同签署人信息
	FlowApproverList []*CommonFlowApprover `json:"FlowApproverList,omitnil" name:"FlowApproverList"`

	// 通过flowid快速获得之前成功通过页面发起的合同生成链接
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
	
	// 资源id，与ResourceType对应
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源类型，与ResourceId对应1：模板   2: 文件
	ResourceType *int64 `json:"ResourceType,omitnil" name:"ResourceType"`

	// 合同流程基础信息
	FlowInfo *BaseFlowInfo `json:"FlowInfo,omitnil" name:"FlowInfo"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同流程配置信息，用于配置发起合同时定制化
	FlowOption *CreateFlowOption `json:"FlowOption,omitnil" name:"FlowOption"`

	// 合同签署人信息
	FlowApproverList []*CommonFlowApprover `json:"FlowApproverList,omitnil" name:"FlowApproverList"`

	// 通过flowid快速获得之前成功通过页面发起的合同生成链接
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
	// 预发起的合同链接， 可以直接点击进入进行合同发起
	PrepareFlowUrl *string `json:"PrepareFlowUrl,omitnil" name:"PrepareFlowUrl"`

	// 合同发起后预览链接， 注意此时合同并未发起，仅只是展示效果
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 个人用户姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 印章名称
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 印章图片的base64，最大不超过 8M
	SealImage *string `json:"SealImage,omitnil" name:"SealImage"`

	// 操作者信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 身份证件类型
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 是否开启印章图片压缩处理，默认不开启，如需开启请设置为 true。当印章超过 2M 时建议开启，开启后图片的 hash 将发生变化。
	SealImageCompress *bool `json:"SealImageCompress,omitnil" name:"SealImageCompress"`

	// 手机号码；当需要开通自动签时，该参数必传
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 是否开通自动签，该功能需联系运营工作人员开通后使用
	EnableAutoSign *bool `json:"EnableAutoSign,omitnil" name:"EnableAutoSign"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。一旦绑定后，将扣减购买的个人自动签账号许可一次（1年有效期），不可解绑释放。不传默认为绑定自动签账号许可。 0-绑定个人自动签账号许可，开通后将扣减购买的个人自动签账号许可一次 1-不绑定，发起合同时将按标准合同套餐进行扣减	
	LicenseType *int64 `json:"LicenseType,omitnil" name:"LicenseType"`
}

type ChannelCreatePreparedPersonalEsignRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 个人用户姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 印章名称
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 印章图片的base64，最大不超过 8M
	SealImage *string `json:"SealImage,omitnil" name:"SealImage"`

	// 操作者信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 身份证件类型
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 是否开启印章图片压缩处理，默认不开启，如需开启请设置为 true。当印章超过 2M 时建议开启，开启后图片的 hash 将发生变化。
	SealImageCompress *bool `json:"SealImageCompress,omitnil" name:"SealImageCompress"`

	// 手机号码；当需要开通自动签时，该参数必传
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 是否开通自动签，该功能需联系运营工作人员开通后使用
	EnableAutoSign *bool `json:"EnableAutoSign,omitnil" name:"EnableAutoSign"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。一旦绑定后，将扣减购买的个人自动签账号许可一次（1年有效期），不可解绑释放。不传默认为绑定自动签账号许可。 0-绑定个人自动签账号许可，开通后将扣减购买的个人自动签账号许可一次 1-不绑定，发起合同时将按标准合同套餐进行扣减	
	LicenseType *int64 `json:"LicenseType,omitnil" name:"LicenseType"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreatePreparedPersonalEsignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreatePreparedPersonalEsignResponseParams struct {
	// 导入的印章 ID
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 待解除的流程编号（即原流程的编号）
	NeedRelievedFlowId *string `json:"NeedRelievedFlowId,omitnil" name:"NeedRelievedFlowId"`

	// 解除协议内容
	ReliveInfo *RelieveInfo `json:"ReliveInfo,omitnil" name:"ReliveInfo"`

	// 非必须，解除协议的本企业签署人列表，默认使用原流程的签署人列表；当解除协议的签署人与原流程的签署人不能相同时（例如原流程签署人离职了），需要指定本企业的其他签署人来替换原流程中的原签署人，注意需要指明ApproverNumber来代表需要替换哪一个签署人，已转发的签署人不包含在内，解除协议的签署人数量不能多于原流程的签署人数量
	ReleasedApprovers []*ReleasedApprover `json:"ReleasedApprovers,omitnil" name:"ReleasedApprovers"`

	// 签署完回调url，最大长度1000个字符
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 暂未开放
	//
	// Deprecated: Organization is deprecated.
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 签署流程的签署截止时间。 值为unix时间戳,精确到秒,不传默认为当前时间七天后
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`
}

type ChannelCreateReleaseFlowRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 待解除的流程编号（即原流程的编号）
	NeedRelievedFlowId *string `json:"NeedRelievedFlowId,omitnil" name:"NeedRelievedFlowId"`

	// 解除协议内容
	ReliveInfo *RelieveInfo `json:"ReliveInfo,omitnil" name:"ReliveInfo"`

	// 非必须，解除协议的本企业签署人列表，默认使用原流程的签署人列表；当解除协议的签署人与原流程的签署人不能相同时（例如原流程签署人离职了），需要指定本企业的其他签署人来替换原流程中的原签署人，注意需要指明ApproverNumber来代表需要替换哪一个签署人，已转发的签署人不包含在内，解除协议的签署人数量不能多于原流程的签署人数量
	ReleasedApprovers []*ReleasedApprover `json:"ReleasedApprovers,omitnil" name:"ReleasedApprovers"`

	// 签署完回调url，最大长度1000个字符
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitnil" name:"Organization"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 签署流程的签署截止时间。 值为unix时间戳,精确到秒,不传默认为当前时间七天后
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`
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

	// 代理企业和员工的信息。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil" name:"Description"`

	// 权限树，权限树内容 PermissionGroups 可参考接口 DescribeIntegrationRoles 的输出
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil" name:"PermissionGroups"`
}

type ChannelCreateRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色名称，最大长度为20个字符，仅限中文、字母、数字和下划线组成。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 代理企业和员工的信息。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 角色描述，最大长度为50个字符
	Description *string `json:"Description,omitnil" name:"Description"`

	// 权限树，权限树内容 PermissionGroups 可参考接口 DescribeIntegrationRoles 的输出
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定印章ID
	SealId *string `json:"SealId,omitnil" name:"SealId"`

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
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定印章ID
	SealId *string `json:"SealId,omitnil" name:"SealId"`

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
	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签场景:
	// E_PRESCRIPTION_AUTO_SIGN 电子处方
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签开通，签署相关配置
	AutoSignConfig *AutoSignConfig `json:"AutoSignConfig,omitnil" name:"AutoSignConfig"`

	// 链接类型，空-默认小程序端链接，H5SIGN-h5端链接
	UrlType *string `json:"UrlType,omitnil" name:"UrlType"`

	// 通知类型，默认不填为不通知开通方，填写 SMS 为短信通知。
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 若上方填写为 SMS，则此处为手机号
	NotifyAddress *string `json:"NotifyAddress,omitnil" name:"NotifyAddress"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为30天。如果不传，默认有效期为7天。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`
}

type ChannelCreateUserAutoSignEnableUrlRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签场景:
	// E_PRESCRIPTION_AUTO_SIGN 电子处方
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 自动签开通，签署相关配置
	AutoSignConfig *AutoSignConfig `json:"AutoSignConfig,omitnil" name:"AutoSignConfig"`

	// 链接类型，空-默认小程序端链接，H5SIGN-h5端链接
	UrlType *string `json:"UrlType,omitnil" name:"UrlType"`

	// 通知类型，默认不填为不通知开通方，填写 SMS 为短信通知。
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 若上方填写为 SMS，则此处为手机号
	NotifyAddress *string `json:"NotifyAddress,omitnil" name:"NotifyAddress"`

	// 链接的过期时间，格式为Unix时间戳，不能早于当前时间，且最大为30天。如果不传，默认有效期为7天。
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
	// 跳转短链
	Url *string `json:"Url,omitnil" name:"Url"`

	// 小程序AppId
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 小程序 原始 Id
	AppOriginalId *string `json:"AppOriginalId,omitnil" name:"AppOriginalId"`

	// 跳转路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// base64格式跳转二维码
	QrCode *string `json:"QrCode,omitnil" name:"QrCode"`

	// 链接类型，空-默认小程序端链接，H5SIGN-h5端链接
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
	// 代理企业和员工的信息。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 角色id，最多20个
	RoleIds []*string `json:"RoleIds,omitnil" name:"RoleIds"`
}

type ChannelDeleteRoleRequest struct {
	*tchttp.BaseRequest
	
	// 代理企业和员工的信息。
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定印章ID
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 指定用户ID数组，电子签系统用户ID
	// 可以填写OpenId，系统会通过组织+渠道+OpenId查询得到UserId进行授权取消。
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
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定印章ID
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 指定用户ID数组，电子签系统用户ID
	// 可以填写OpenId，系统会通过组织+渠道+OpenId查询得到UserId进行授权取消。
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
type ChannelDescribeEmployeesRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定每页多少条数据，单页最大20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询过滤实名用户，Key为Status，Values为["IsVerified"]
	// 根据第三方系统openId过滤查询员工时,Key为StaffOpenId,Values为["OpenId","OpenId",...]
	// 查询离职员工时，Key为Status，Values为["QuiteJob"]
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0,最大为20000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelDescribeEmployeesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定每页多少条数据，单页最大20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询过滤实名用户，Key为Status，Values为["IsVerified"]
	// 根据第三方系统openId过滤查询员工时,Key为StaffOpenId,Values为["OpenId","OpenId",...]
	// 查询离职员工时，Key为Status，Values为["QuiteJob"]
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0,最大为20000
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
	// 员工数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Employees []*Staff `json:"Employees,omitnil" name:"Employees"`

	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0，最大20000
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 指定每页多少条数据，单页最大20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 符合条件的员工数量
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
	// 应用相关信息。此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 电子签流程的Id
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`
}

type ChannelDescribeFlowComponentsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 电子签流程的Id
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
	// 流程关联的填写控件信息，控件会按照参与方进行分类。
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 返回最大数量，最大为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0，最大为20000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询信息类型，为1时返回授权用户，为其他值时不返回
	InfoType *int64 `json:"InfoType,omitnil" name:"InfoType"`

	// 印章id（没有输入返回所有）
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 印章类型列表（都是组织机构印章）。
	// 为空时查询所有类型的印章。
	// 目前支持以下类型：
	// OFFICIAL：企业公章；
	// CONTRACT：合同专用章；
	// ORGANIZATION_SEAL：企业印章(图片上传创建)；
	// LEGAL_PERSON_SEAL：法定代表人章
	SealTypes []*string `json:"SealTypes,omitnil" name:"SealTypes"`
}

type ChannelDescribeOrganizationSealsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 返回最大数量，最大为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0，最大为20000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询信息类型，为1时返回授权用户，为其他值时不返回
	InfoType *int64 `json:"InfoType,omitnil" name:"InfoType"`

	// 印章id（没有输入返回所有）
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 印章类型列表（都是组织机构印章）。
	// 为空时查询所有类型的印章。
	// 目前支持以下类型：
	// OFFICIAL：企业公章；
	// CONTRACT：合同专用章；
	// ORGANIZATION_SEAL：企业印章(图片上传创建)；
	// LEGAL_PERSON_SEAL：法定代表人章
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定每页多少条数据，单页最大200
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 查询的关键字段:
	// Key:"RoleType",Values:["1"]查询系统角色，Values:["2"]查询自定义角色
	// Key:"RoleStatus",Values:["1"]查询启用角色，Values:["2"]查询禁用角色
	// Key:"IsReturnPermissionGroup"，Values:["0"]:表示接口不返回角色对应的权限树字段，Values:["1"]表示接口返回角色对应的权限树字段
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0，最大2000
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 操作人信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelDescribeRolesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 指定每页多少条数据，单页最大200
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 查询的关键字段:
	// Key:"RoleType",Values:["1"]查询系统角色，Values:["2"]查询自定义角色
	// Key:"RoleStatus",Values:["1"]查询启用角色，Values:["2"]查询禁用角色
	// Key:"IsReturnPermissionGroup"，Values:["0"]:表示接口不返回角色对应的权限树字段，Values:["1"]表示接口返回角色对应的权限树字段
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0，最大2000
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
	// 查询结果分页返回，此处指定第几页，如果不传默认从第一页返回。页码从 0 开始，即首页为 0，最大2000
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 指定每页多少条数据，单页最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询角色的总数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 角色信息
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
	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签场景:
	// E_PRESCRIPTION_AUTO_SIGN 电子处方
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 查询开启状态的用户信息
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelDescribeUserAutoSignStatusRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签场景:
	// E_PRESCRIPTION_AUTO_SIGN 电子处方
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 查询开启状态的用户信息
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 操作人信息
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
	// 是否开通
	IsOpen *bool `json:"IsOpen,omitnil" name:"IsOpen"`

	// 自动签许可生效时间。当且仅当已开通自动签时有值。
	LicenseFrom *int64 `json:"LicenseFrom,omitnil" name:"LicenseFrom"`

	// 自动签许可到期时间。当且仅当已开通自动签时有值。
	LicenseTo *int64 `json:"LicenseTo,omitnil" name:"LicenseTo"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。一旦绑定后，将扣减购买的个人自动签账号许可一次（1年有效期），不可解绑释放。不传默认为绑定自动签账号许可。 0-绑定个人自动签账号许可，开通后将扣减购买的个人自动签账号许可一次
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
	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签场景:
	// E_PRESCRIPTION_AUTO_SIGN 电子处方
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 关闭自动签的个人的三要素
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 操作人信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelDisableUserAutoSignRequest struct {
	*tchttp.BaseRequest
	
	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 自动签场景:
	// E_PRESCRIPTION_AUTO_SIGN 电子处方
	SceneKey *string `json:"SceneKey,omitnil" name:"SceneKey"`

	// 关闭自动签的个人的三要素
	UserInfo *UserThreeFactor `json:"UserInfo,omitnil" name:"UserInfo"`

	// 操作人信息
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 均必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 任务Id，通过ChannelCreateConvertTaskApi接口获得
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
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 均必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 任务Id，通过ChannelCreateConvertTaskApi接口获得
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
	// 0  :NeedTranform   - 任务已提交
	// 4  :Processing     - 文档转换中
	// 8  :TaskEnd        - 任务处理完成
	// -2 :DownloadFailed - 下载失败
	// -6 :ProcessFailed  - 转换失败
	// -13:ProcessTimeout - 转换文件超时
	TaskStatus *int64 `json:"TaskStatus,omitnil" name:"TaskStatus"`

	// 状态描述，需要关注的状态
	// NeedTranform   - 任务已提交
	// Processing     - 文档转换中
	// TaskEnd        - 任务处理完成
	// DownloadFailed - 下载失败
	// ProcessFailed  - 转换失败
	// ProcessTimeout - 转换文件超时
	TaskMessage *string `json:"TaskMessage,omitnil" name:"TaskMessage"`

	// 资源Id，也是FileId，用于文件发起使用
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

	// 权限树，权限树内容 PermissionGroups 可参考接口 DescribeIntegrationRoles 的输出
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

	// 权限树，权限树内容 PermissionGroups 可参考接口 DescribeIntegrationRoles 的输出
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

type ChannelRole struct {
	// 角色id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *string `json:"RoleId,omitnil" name:"RoleId"`

	// 角色名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// 角色状态：1-启用；2-禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleStatus *uint64 `json:"RoleStatus,omitnil" name:"RoleStatus"`

	// 权限树
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermissionGroups []*PermissionGroup `json:"PermissionGroups,omitnil" name:"PermissionGroups"`
}

// Predefined struct for user
type ChannelUpdateSealStatusRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作的印章状态，DISABLE-停用印章
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
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作的印章状态，DISABLE-停用印章
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
	// 流程ID
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type ChannelVerifyPdfRequest struct {
	*tchttp.BaseRequest
	
	// 流程ID
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
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
	// 验签结果，1-文件未被篡改，全部签名在腾讯电子签完成； 2-文件未被篡改，部分签名在腾讯电子签完成；3-文件被篡改；4-异常：文件内没有签名域；5-异常：文件签名格式错误
	VerifyResult *int64 `json:"VerifyResult,omitnil" name:"VerifyResult"`

	// 验签结果详情,内部状态1-验签成功，在电子签签署；2-验签成功，在其他平台签署；3-验签失败；4-pdf文件没有签名域；5-文件签名格式错误
	PdfVerifyResults []*PdfVerifyResult `json:"PdfVerifyResults,omitnil" name:"PdfVerifyResults"`

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
	// 指定当前签署人为第三方应用集成子客，默认false：当前签署人为第三方应用集成子客，true：当前签署人为saas企业用户
	NotChannelOrganization *bool `json:"NotChannelOrganization,omitnil" name:"NotChannelOrganization"`

	// 签署人类型,目前支持：0-企业签署人，1-个人签署人，3-企业静默签署人
	ApproverType *int64 `json:"ApproverType,omitnil" name:"ApproverType"`

	// 企业id
	OrganizationId *string `json:"OrganizationId,omitnil" name:"OrganizationId"`

	// 企业OpenId，第三方应用集成非静默签子客企业签署人发起合同必传
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 企业名称，第三方应用集成非静默签子客企业签署人必传，saas企业签署人必传
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 用户id
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户openId，第三方应用集成非静默签子客企业签署人必传
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 签署人名称，saas企业签署人，个人签署人必传
	ApproverName *string `json:"ApproverName,omitnil" name:"ApproverName"`

	// 签署人手机号，saas企业签署人，个人签署人必传
	ApproverMobile *string `json:"ApproverMobile,omitnil" name:"ApproverMobile"`

	// 签署人Id，使用模板发起是，对应模板配置中的签署人RecipientId
	// 注意：模板发起时该字段必填
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 签署前置条件：阅读时长限制，不传默认10s,最大300s，最小3s
	PreReadTime *int64 `json:"PreReadTime,omitnil" name:"PreReadTime"`

	// 签署前置条件：阅读全文限制
	IsFullText *bool `json:"IsFullText,omitnil" name:"IsFullText"`

	// 通知类型：SMS（短信） NONE（不做通知）, 不传 默认SMS
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// 签署人配置
	ApproverOption *CommonApproverOption `json:"ApproverOption,omitnil" name:"ApproverOption"`

	// 签署控件：文件发起使用
	SignComponents []*Component `json:"SignComponents,omitnil" name:"SignComponents"`

	// 签署人查看合同时认证方式, 1-实名查看 2-短信验证码查看(企业签署方不支持该方式) 如果不传默认为1 查看合同的认证方式 Flow层级的优先于approver层级的 （当手写签名方式为OCR_ESIGN时，合同认证方式2无效，因为这种签名方式依赖实名认证）
	ApproverVerifyTypes []*int64 `json:"ApproverVerifyTypes,omitnil" name:"ApproverVerifyTypes"`

	// 签署人签署合同时的认证方式 1-人脸认证 2-签署密码 3-运营商三要素(默认为1,2)	
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
	// KEYWORD - 关键字
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
	// TEXT/MULTI_LINE_TEXT控件可以指定
	// 1 Font：目前只支持黑体、宋体
	// 2 FontSize： 范围12-72
	// 3 FontAlign： Left/Right/Center，左对齐/居中/右对齐
	// 例如：{"FontSize":12}
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
	// 参数样例： "ComponentExtra":"{"PageRange":[{"BeginPage":1,"EndPage":-1}]}"
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type CreateChannelFlowEvidenceReportRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 签署流程编号
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
	// 出证报告 ID，可用户DescribeChannelFlowEvidenceReport接口查询出证PDF的下载地址
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportId *string `json:"ReportId,omitnil" name:"ReportId"`

	// 出征任务的执行状态,状态列表如下
	// 
	// - EvidenceStatusExecuting : 出征任务正在执行中
	// - EvidenceStatusSuccess : 出征任务执行成功
	// - EvidenceStatusFailed : 出征任务执行失败
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
	// 关于渠道应用的相关信息，包括子客企业及应用编、号等详细内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 企业信息变更类型，可选类型如下：
	// <ul><li>**1**：企业超管变更</li><li>**2**：企业基础信息变更</li></ul>
	ChangeType *uint64 `json:"ChangeType,omitnil" name:"ChangeType"`
}

type CreateChannelOrganizationInfoChangeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 关于渠道应用的相关信息，包括子客企业及应用编、号等详细内容，您可以参阅开发者中心所提供的 Agent 结构体以获取详细定义。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 企业信息变更类型，可选类型如下：
	// <ul><li>**1**：企业超管变更</li><li>**2**：企业基础信息变更</li></ul>
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
	// 创建的企业信息变更链接。
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
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 子客企业名称，最大长度64个字符
	// 注意：
	// 1、如果您的企业已经在认证授权中或者激活完成，这里修改子客企业名字将不会生效。
	// 2、该名称需要与Agent.ProxyOrganizationOpenId相匹配。
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil" name:"ProxyOrganizationName"`

	// 子客企业统一社会信用代码，最大长度200个字符
	// 注意：
	// 1、如果您的企业已经在认证授权中或者激活完成，这里修改子客企业名字将不会生效。
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil" name:"UniformSocialCreditCode"`

	// 子客企业经办人的姓名，最大长度50个字符
	// 注意：
	// 1、若经办人已经实名，这里修改经办人名字传入将不会生效。
	// 2、该名称需要和Agent. ProxyOperator.OpenId相匹配
	ProxyOperatorName *string `json:"ProxyOperatorName,omitnil" name:"ProxyOperatorName"`

	// PC控制台登录后进入该参数指定的模块，如果不传递，将默认进入控制台首页。支持的模块包括：
	// 1、DOCUMENT:合同管理模块，
	// 2、TEMPLATE:企业模板管理模块，
	// 3、SEAL:印章管理模块，
	// 4、OPERATOR:组织管理模块，
	// 默认将进入企业中心模块
	// 注意：
	// 1、如果EndPoint选择"CHANNEL"或"APP"，该参数仅支持传递"SEAL"，进入印章管理模块
	// 2、该参数仅在企业和员工激活完成，登录控制台场景才生效。
	Module *string `json:"Module,omitnil" name:"Module"`

	// 该参数和Module参数配合使用，用于指定模块下的资源Id，指定后链接登录将展示该资源的详情。
	// 根据Module参数的不同所代表的含义不同。当前支持：
	// 1、如果Module="SEAL"，ModuleId代表印章Id, 登录链接将直接查看指定印章的详情。
	// 2、如果Module="TEMPLATE"，ModuleId代表模版Id，登录链接将直接查看指定模版的详情。
	// 3、如果Module="1、DOCUMENT"，ModuleId代表合同Id，登录链接将直接查看指定合同的详情。
	// 注意：
	// 1、该参数仅在企业和员工激活完成，登录控制台场景才生效。
	// 2、ModuleId需要和Module对应，ModuleId可以通过API或者控制台获取到。
	ModuleId *string `json:"ModuleId,omitnil" name:"ModuleId"`

	// 是否展示左侧菜单栏 
	// "ENABLE": 是，展示 
	// “DISABLE”: 否，不展示
	// 默认值为ENABLE
	// 注意：
	// 1、该参数仅在企业和员工激活完成，登录控制台场景才生效。
	MenuStatus *string `json:"MenuStatus,omitnil" name:"MenuStatus"`

	// 生成链接的类型：
	// "PC"：PC控制台链接
	// "CHANNEL"：H5跳转到电子签小程序链接
	// "APP"：第三方APP或小程序跳转电子签小程序链接
	// 默认将生成PC控制台链接
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// 触发自动跳转事件，仅对EndPoint为App类型有效，可选值包括：
	// "VERIFIED":企业认证完成/员工认证完成后跳回原App/小程序
	AutoJumpBackEvent *string `json:"AutoJumpBackEvent,omitnil" name:"AutoJumpBackEvent"`

	// 可选的企业授权方式: 
	// 1：上传授权书 
	// 2：转法定代表人授权
	// 4：企业实名认证（信任第三方认证源）（此项仅支持单选）
	// 未选择信任第三方认证源时，如果是法人进行企业激活，仅支持法人扫脸直接授权，该配置不生效；选择信任第三方认证源时，请先通过“同步企业信息”接口同步信息。
	// 该参数仅在企业激活场景生效
	AuthorizationTypes []*int64 `json:"AuthorizationTypes,omitnil" name:"AuthorizationTypes"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type CreateConsoleLoginUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 子客企业名称，最大长度64个字符
	// 注意：
	// 1、如果您的企业已经在认证授权中或者激活完成，这里修改子客企业名字将不会生效。
	// 2、该名称需要与Agent.ProxyOrganizationOpenId相匹配。
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil" name:"ProxyOrganizationName"`

	// 子客企业统一社会信用代码，最大长度200个字符
	// 注意：
	// 1、如果您的企业已经在认证授权中或者激活完成，这里修改子客企业名字将不会生效。
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil" name:"UniformSocialCreditCode"`

	// 子客企业经办人的姓名，最大长度50个字符
	// 注意：
	// 1、若经办人已经实名，这里修改经办人名字传入将不会生效。
	// 2、该名称需要和Agent. ProxyOperator.OpenId相匹配
	ProxyOperatorName *string `json:"ProxyOperatorName,omitnil" name:"ProxyOperatorName"`

	// PC控制台登录后进入该参数指定的模块，如果不传递，将默认进入控制台首页。支持的模块包括：
	// 1、DOCUMENT:合同管理模块，
	// 2、TEMPLATE:企业模板管理模块，
	// 3、SEAL:印章管理模块，
	// 4、OPERATOR:组织管理模块，
	// 默认将进入企业中心模块
	// 注意：
	// 1、如果EndPoint选择"CHANNEL"或"APP"，该参数仅支持传递"SEAL"，进入印章管理模块
	// 2、该参数仅在企业和员工激活完成，登录控制台场景才生效。
	Module *string `json:"Module,omitnil" name:"Module"`

	// 该参数和Module参数配合使用，用于指定模块下的资源Id，指定后链接登录将展示该资源的详情。
	// 根据Module参数的不同所代表的含义不同。当前支持：
	// 1、如果Module="SEAL"，ModuleId代表印章Id, 登录链接将直接查看指定印章的详情。
	// 2、如果Module="TEMPLATE"，ModuleId代表模版Id，登录链接将直接查看指定模版的详情。
	// 3、如果Module="1、DOCUMENT"，ModuleId代表合同Id，登录链接将直接查看指定合同的详情。
	// 注意：
	// 1、该参数仅在企业和员工激活完成，登录控制台场景才生效。
	// 2、ModuleId需要和Module对应，ModuleId可以通过API或者控制台获取到。
	ModuleId *string `json:"ModuleId,omitnil" name:"ModuleId"`

	// 是否展示左侧菜单栏 
	// "ENABLE": 是，展示 
	// “DISABLE”: 否，不展示
	// 默认值为ENABLE
	// 注意：
	// 1、该参数仅在企业和员工激活完成，登录控制台场景才生效。
	MenuStatus *string `json:"MenuStatus,omitnil" name:"MenuStatus"`

	// 生成链接的类型：
	// "PC"：PC控制台链接
	// "CHANNEL"：H5跳转到电子签小程序链接
	// "APP"：第三方APP或小程序跳转电子签小程序链接
	// 默认将生成PC控制台链接
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// 触发自动跳转事件，仅对EndPoint为App类型有效，可选值包括：
	// "VERIFIED":企业认证完成/员工认证完成后跳回原App/小程序
	AutoJumpBackEvent *string `json:"AutoJumpBackEvent,omitnil" name:"AutoJumpBackEvent"`

	// 可选的企业授权方式: 
	// 1：上传授权书 
	// 2：转法定代表人授权
	// 4：企业实名认证（信任第三方认证源）（此项仅支持单选）
	// 未选择信任第三方认证源时，如果是法人进行企业激活，仅支持法人扫脸直接授权，该配置不生效；选择信任第三方认证源时，请先通过“同步企业信息”接口同步信息。
	// 该参数仅在企业激活场景生效
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
	// 子客企业Web控制台url注意事项：
	// 1. 所有类型的链接在企业未认证/员工未认证完成时，只要在有效期内（一年）都可以访问
	// 2. 若企业认证完成且员工认证完成后，重新获取pc端的链接在5分钟之内有效，且只能访问一次
	// 3. 若企业认证完成且员工认证完成后，重新获取CHANNEL/APP的链接只要在有效期内（一年）都可以访问
	// 4. 此链接仅单次有效，每次登录需要需要重新创建新的链接，尽量不要做链接存储，多次使用。
	// 5. 创建的链接应避免被转义，如：&被转义为\u0026；如使用Postman请求后，请选择响应类型为 JSON，否则链接将被转义
	ConsoleUrl *string `json:"ConsoleUrl,omitnil" name:"ConsoleUrl"`

	// 子客企业是否已开通腾讯电子签，true-是，false-否
	// 注意：
	// 1、企业是否实名根据传参Agent.ProxyOrganizationOpenId进行判断，非企业名称或者社会信用代码
	IsActivated *bool `json:"IsActivated,omitnil" name:"IsActivated"`

	// 当前经办人是否已认证，true-是，false-否
	// 注意：
	// 1、经办人是否实名是根据Agent.ProxyOperator.OpenId判断，非经办人姓名
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
	// 是否允许修改合同信息，true-是，false-否
	CanEditFlow *bool `json:"CanEditFlow,omitnil" name:"CanEditFlow"`

	// 是否允许发起合同弹窗隐藏合同名称，true-允许，false-不允许
	HideShowFlowName *bool `json:"HideShowFlowName,omitnil" name:"HideShowFlowName"`

	// 是否允许发起合同弹窗隐藏合同类型，true-允许，false-不允许
	HideShowFlowType *bool `json:"HideShowFlowType,omitnil" name:"HideShowFlowType"`

	// 是否允许发起合同弹窗隐藏合同到期时间，true-允许，false-不允许
	HideShowDeadline *bool `json:"HideShowDeadline,omitnil" name:"HideShowDeadline"`

	// 是否允许发起合同步骤跳过指定签署方步骤，true-允许，false-不允许
	CanSkipAddApprover *bool `json:"CanSkipAddApprover,omitnil" name:"CanSkipAddApprover"`

	// 定制化发起合同弹窗的描述信息，描述信息最长500
	CustomCreateFlowDescription *string `json:"CustomCreateFlowDescription,omitnil" name:"CustomCreateFlowDescription"`
}

// Predefined struct for user
type CreateFlowsByTemplatesRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 均必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 要创建的合同信息列表，最多支持一次创建20个合同
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil" name:"FlowInfos"`

	// 是否为预览模式；默认为false，即非预览模式，此时发起合同并返回FlowIds；若为预览模式，不会发起合同，会返回PreviewUrls；
	// 预览链接有效期300秒；
	// 同时，如果预览的文件中指定了动态表格控件，需要进行异步合成；此时此接口返回的是合成前的文档预览链接，而合成完成后的文档预览链接会通过：回调通知的方式、或使用返回的TaskInfo中的TaskId通过ChannelGetTaskResultApi接口查询；
	NeedPreview *bool `json:"NeedPreview,omitnil" name:"NeedPreview"`

	// 预览链接类型 默认:0-文件流, 1- H5链接 注意:此参数在NeedPreview 为true 时有效,
	PreviewType *int64 `json:"PreviewType,omitnil" name:"PreviewType"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type CreateFlowsByTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 均必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 要创建的合同信息列表，最多支持一次创建20个合同
	FlowInfos []*FlowInfo `json:"FlowInfos,omitnil" name:"FlowInfos"`

	// 是否为预览模式；默认为false，即非预览模式，此时发起合同并返回FlowIds；若为预览模式，不会发起合同，会返回PreviewUrls；
	// 预览链接有效期300秒；
	// 同时，如果预览的文件中指定了动态表格控件，需要进行异步合成；此时此接口返回的是合成前的文档预览链接，而合成完成后的文档预览链接会通过：回调通知的方式、或使用返回的TaskInfo中的TaskId通过ChannelGetTaskResultApi接口查询；
	NeedPreview *bool `json:"NeedPreview,omitnil" name:"NeedPreview"`

	// 预览链接类型 默认:0-文件流, 1- H5链接 注意:此参数在NeedPreview 为true 时有效,
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
	// 多个合同ID
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 第三方应用平台的业务信息, 与创建合同的FlowInfos数组中的CustomerData一一对应
	CustomerData []*string `json:"CustomerData,omitnil" name:"CustomerData"`

	// 创建消息，对应多个合同ID，
	// 成功为“”,创建失败则对应失败消息
	ErrorMessages []*string `json:"ErrorMessages,omitnil" name:"ErrorMessages"`

	// 预览模式下返回的预览文件url数组
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
	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 印章名称，最大长度不超过50字符
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 印章图片base64，大小不超过10M（原始图片不超过7.6M）
	SealImage *string `json:"SealImage,omitnil" name:"SealImage"`

	// 操作者的信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 本接口支持上传图片印章及系统直接生成印章； 如果要使用系统生成印章，此值传：SealGenerateSourceSystem； 如果要使用图片上传请传字段 SealImage
	GenerateSource *string `json:"GenerateSource,omitnil" name:"GenerateSource"`

	// 电子印章类型：
	// <ul><li>OFFICIAL-公章</li>
	// <li>CONTRACT-合同专用章;</li>
	// <li>FINANCE-合财务专用章;</li>
	// <li>PERSONNEL-人事专用章
	// </li>
	// <li>默认：OFFICIAL</li>
	// </ul>
	SealType *string `json:"SealType,omitnil" name:"SealType"`

	// 企业印章横向文字，最多可填15个汉字（若超过印章最大宽度，优先压缩字间距，其次缩小字号
	SealHorizontalText *string `json:"SealHorizontalText,omitnil" name:"SealHorizontalText"`

	// 印章样式:
	// 
	// <ul><li>cycle:圆形印章</li>
	// <li>ellipse:椭圆印章</li>
	// <li> 注：默认圆形印章</li></ul>
	SealStyle *string `json:"SealStyle,omitnil" name:"SealStyle"`

	// 印章尺寸取值描述：<ul><li> 42_42 圆形企业公章直径42mm</li>
	// <li> 40_40 圆形企业印章直径40mm</li>
	// <li> 45_30 椭圆形印章45mm x 30mm</li>
	// </ul>
	SealSize *string `json:"SealSize,omitnil" name:"SealSize"`
}

type CreateSealByImageRequest struct {
	*tchttp.BaseRequest
	
	// 代理企业和员工的信息。
	// 在集团企业代理子企业操作的场景中，需设置此参数。在此情境下，ProxyOrganizationId（子企业的组织ID）为必填项。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 印章名称，最大长度不超过50字符
	SealName *string `json:"SealName,omitnil" name:"SealName"`

	// 印章图片base64，大小不超过10M（原始图片不超过7.6M）
	SealImage *string `json:"SealImage,omitnil" name:"SealImage"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 本接口支持上传图片印章及系统直接生成印章； 如果要使用系统生成印章，此值传：SealGenerateSourceSystem； 如果要使用图片上传请传字段 SealImage
	GenerateSource *string `json:"GenerateSource,omitnil" name:"GenerateSource"`

	// 电子印章类型：
	// <ul><li>OFFICIAL-公章</li>
	// <li>CONTRACT-合同专用章;</li>
	// <li>FINANCE-合财务专用章;</li>
	// <li>PERSONNEL-人事专用章
	// </li>
	// <li>默认：OFFICIAL</li>
	// </ul>
	SealType *string `json:"SealType,omitnil" name:"SealType"`

	// 企业印章横向文字，最多可填15个汉字（若超过印章最大宽度，优先压缩字间距，其次缩小字号
	SealHorizontalText *string `json:"SealHorizontalText,omitnil" name:"SealHorizontalText"`

	// 印章样式:
	// 
	// <ul><li>cycle:圆形印章</li>
	// <li>ellipse:椭圆印章</li>
	// <li> 注：默认圆形印章</li></ul>
	SealStyle *string `json:"SealStyle,omitnil" name:"SealStyle"`

	// 印章尺寸取值描述：<ul><li> 42_42 圆形企业公章直径42mm</li>
	// <li> 40_40 圆形企业印章直径40mm</li>
	// <li> 45_30 椭圆形印章45mm x 30mm</li>
	// </ul>
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
	// 可登录腾讯电子签控制台，在 "印章"->"印章中心"选择查看的印章，在"印章详情" 中查看某个印章的SealId(在页面中展示为印章ID)。
	SealId *string `json:"SealId,omitnil" name:"SealId"`

	// 电子印章预览链接地址，地址默认失效时间为24小时。
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 流程(合同)的编号列表，最多支持100个。(备注：该参数和合同组编号必须二选一)
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 合同组编号(备注：该参数和合同(流程)编号数组必须二选一)
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 签署链接类型,可以设置的参数如下
	// 
	// - WEIXINAPP:短链直接跳小程序 (默认类型)
	// - CHANNEL:跳转H5页面
	// - APP:第三方APP或小程序跳转电子签小程序
	// - LONGURL2WEIXINAPP:长链接跳转小程序
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// 签署链接生成类型，可以选择的类型如下
	// 
	// - ALL：全部签署方签署链接，此时不会给自动签署的签署方创建签署链接(默认类型)
	// - CHANNEL：第三方平台子客企业企业
	// - NOT_CHANNEL：非第三方平台子客企业企业
	// - PERSON：个人
	// - FOLLOWER：关注方，目前是合同抄送方
	GenerateType *string `json:"GenerateType,omitnil" name:"GenerateType"`

	// 非第三方平台子客企业参与方的企业名称，GenerateType为"NOT_CHANNEL"时必填
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 参与人姓名
	// GenerateType为"PERSON"(即个人签署方)时必填
	Name *string `json:"Name,omitnil" name:"Name"`

	// 参与人手机号
	// GenerateType为"PERSON"或"FOLLOWER"时必填
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 第三方平台子客企业的企业OpenId，GenerateType为"CHANNEL"时必填
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 第三方平台子客企业参与人OpenId，GenerateType为"CHANNEL"时可用，指定到具体参与人, 仅展示已经实名的经办人信息
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// Endpoint为"APP" 类型的签署链接，可以设置此值；支持调用方小程序打开签署链接，在电子签小程序完成签署后自动回跳至调用方小程序
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil" name:"AutoJumpBack"`

	// 签署完之后的H5页面的跳转链接，针对Endpoint为CHANNEL时有效，最大长度1000个字符。
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 生成的签署链接在签署过程隐藏的按钮列表, 可以设置隐藏的按钮列表如下
	// 
	// - 0:合同签署页面更多操作按钮
	// - 1:合同签署页面更多操作的拒绝签署按钮
	// - 2:合同签署页面更多操作的转他人处理按钮
	// - 3:签署成功页的查看详情按钮
	Hides []*int64 `json:"Hides,omitnil" name:"Hides"`

	// 签署节点ID，用于补充动态签署人，使用此参数需要与flow_ids数量一致
	RecipientIds []*string `json:"RecipientIds,omitnil" name:"RecipientIds"`
}

type CreateSignUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 流程(合同)的编号列表，最多支持100个。(备注：该参数和合同组编号必须二选一)
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 合同组编号(备注：该参数和合同(流程)编号数组必须二选一)
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 签署链接类型,可以设置的参数如下
	// 
	// - WEIXINAPP:短链直接跳小程序 (默认类型)
	// - CHANNEL:跳转H5页面
	// - APP:第三方APP或小程序跳转电子签小程序
	// - LONGURL2WEIXINAPP:长链接跳转小程序
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// 签署链接生成类型，可以选择的类型如下
	// 
	// - ALL：全部签署方签署链接，此时不会给自动签署的签署方创建签署链接(默认类型)
	// - CHANNEL：第三方平台子客企业企业
	// - NOT_CHANNEL：非第三方平台子客企业企业
	// - PERSON：个人
	// - FOLLOWER：关注方，目前是合同抄送方
	GenerateType *string `json:"GenerateType,omitnil" name:"GenerateType"`

	// 非第三方平台子客企业参与方的企业名称，GenerateType为"NOT_CHANNEL"时必填
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 参与人姓名
	// GenerateType为"PERSON"(即个人签署方)时必填
	Name *string `json:"Name,omitnil" name:"Name"`

	// 参与人手机号
	// GenerateType为"PERSON"或"FOLLOWER"时必填
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 第三方平台子客企业的企业OpenId，GenerateType为"CHANNEL"时必填
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 第三方平台子客企业参与人OpenId，GenerateType为"CHANNEL"时可用，指定到具体参与人, 仅展示已经实名的经办人信息
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// Endpoint为"APP" 类型的签署链接，可以设置此值；支持调用方小程序打开签署链接，在电子签小程序完成签署后自动回跳至调用方小程序
	AutoJumpBack *bool `json:"AutoJumpBack,omitnil" name:"AutoJumpBack"`

	// 签署完之后的H5页面的跳转链接，针对Endpoint为CHANNEL时有效，最大长度1000个字符。
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 生成的签署链接在签署过程隐藏的按钮列表, 可以设置隐藏的按钮列表如下
	// 
	// - 0:合同签署页面更多操作按钮
	// - 1:合同签署页面更多操作的拒绝签署按钮
	// - 2:合同签署页面更多操作的转他人处理按钮
	// - 3:签署成功页的查看详情按钮
	Hides []*int64 `json:"Hides,omitnil" name:"Hides"`

	// 签署节点ID，用于补充动态签署人，使用此参数需要与flow_ids数量一致
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
type DescribeChannelFlowEvidenceReportRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 出证报告编号
	ReportId *string `json:"ReportId,omitnil" name:"ReportId"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type DescribeChannelFlowEvidenceReportRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 出证报告编号
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
	// 出证报告下载 URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportUrl *string `json:"ReportUrl,omitnil" name:"ReportUrl"`

	// 出征任务的执行状态,状态列表如下
	// 
	// - EvidenceStatusExecuting : 出征任务正在执行中
	// - EvidenceStatusSuccess : 出征任务执行成功
	// - EvidenceStatusFailed : 出征任务执行失败
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
type DescribeExtendedServiceAuthInfoRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
	// 
	// 注: 此接口 参数Agent. ProxyOperator.OpenId 需要传递超管或者法人的OpenId
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`
}

type DescribeExtendedServiceAuthInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
	// 
	// 注: 此接口 参数Agent. ProxyOperator.OpenId 需要传递超管或者法人的OpenId
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
	// 企业扩展服务授权信息
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同(流程)编号数组，最多支持100个。
	// <br/>备注：该参数和合同组编号必须二选一, 如果填写FlowGroupId则忽略此FlowIds的入参
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 合同组编号
	// <br/>备注：该参数和合同(流程)编号数组必须二选一
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type DescribeFlowDetailInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 合同(流程)编号数组，最多支持100个。
	// <br/>备注：该参数和合同组编号必须二选一, 如果填写FlowGroupId则忽略此FlowIds的入参
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 合同组编号
	// <br/>备注：该参数和合同(流程)编号数组必须二选一
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
	// 第三方平台应用号Id
	ApplicationId *string `json:"ApplicationId,omitnil" name:"ApplicationId"`

	// 第三方平台子客企业OpenId
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil" name:"ProxyOrganizationOpenId"`

	// 合同(签署流程)的具体详细描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowInfo []*FlowDetailInfo `json:"FlowInfo,omitnil" name:"FlowInfo"`

	// 合同组编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowGroupId *string `json:"FlowGroupId,omitnil" name:"FlowGroupId"`

	// 合同组名称
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
	// 应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询资源所对应的签署流程Id，最多支持50个
	FlowIds []*string `json:"FlowIds,omitnil" name:"FlowIds"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type DescribeResourceUrlsByFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 查询资源所对应的签署流程Id，最多支持50个
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
	// 签署流程资源对应链接信息
	FlowResourceUrlInfos []*FlowResourceUrlInfo `json:"FlowResourceUrlInfos,omitnil" name:"FlowResourceUrlInfos"`

	// 创建消息，对应多个合同ID，
	// 成功为“”,创建失败则对应失败消息
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
	// 应用相关信息。 
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 模板唯一标识，查询单个模板时使用
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 查询内容：
	// 0-模板列表及详情（默认），
	// 1-仅模板列表
	ContentType *int64 `json:"ContentType,omitnil" name:"ContentType"`

	// 指定每页多少条数据，如果不传默认为20，单页最大100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果分页返回，此处指定第几页，如果不传默从第一页返回。页码从0开始，即首页为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 是否返回所有组件信息。
	// 默认false，只返回发起方控件；
	// true，返回所有签署方控件
	QueryAllComponents *bool `json:"QueryAllComponents,omitnil" name:"QueryAllComponents"`

	// 模糊搜索模板名称，最大长度200
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 是否获取模板预览链接，
	// 默认false-不获取
	// true-获取
	WithPreviewUrl *bool `json:"WithPreviewUrl,omitnil" name:"WithPreviewUrl"`

	// 是否获取模板的PDF文件链接。
	// 默认false-不获取
	// true-获取
	// 请联系客户经理开白后使用。
	WithPdfUrl *bool `json:"WithPdfUrl,omitnil" name:"WithPdfUrl"`

	// 对应第三方应用平台企业的模板ID
	ChannelTemplateId *string `json:"ChannelTemplateId,omitnil" name:"ChannelTemplateId"`

	// 操作者的信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type DescribeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 模板唯一标识，查询单个模板时使用
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 查询内容：
	// 0-模板列表及详情（默认），
	// 1-仅模板列表
	ContentType *int64 `json:"ContentType,omitnil" name:"ContentType"`

	// 指定每页多少条数据，如果不传默认为20，单页最大100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果分页返回，此处指定第几页，如果不传默从第一页返回。页码从0开始，即首页为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 是否返回所有组件信息。
	// 默认false，只返回发起方控件；
	// true，返回所有签署方控件
	QueryAllComponents *bool `json:"QueryAllComponents,omitnil" name:"QueryAllComponents"`

	// 模糊搜索模板名称，最大长度200
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 是否获取模板预览链接，
	// 默认false-不获取
	// true-获取
	WithPreviewUrl *bool `json:"WithPreviewUrl,omitnil" name:"WithPreviewUrl"`

	// 是否获取模板的PDF文件链接。
	// 默认false-不获取
	// true-获取
	// 请联系客户经理开白后使用。
	WithPdfUrl *bool `json:"WithPdfUrl,omitnil" name:"WithPdfUrl"`

	// 对应第三方应用平台企业的模板ID
	ChannelTemplateId *string `json:"ChannelTemplateId,omitnil" name:"ChannelTemplateId"`

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
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "QueryAllComponents")
	delete(f, "TemplateName")
	delete(f, "WithPreviewUrl")
	delete(f, "WithPdfUrl")
	delete(f, "ChannelTemplateId")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplatesResponseParams struct {
	// 模板列表
	Templates []*TemplateInfo `json:"Templates,omitnil" name:"Templates"`

	// 查询到的总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 每页多少条数据
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
	// 应用信息，此接口Agent.AppId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 开始时间，例如：2021-03-21
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 结束时间，例如：2021-06-21；
	// 开始时间到结束时间的区间长度小于等于90天。
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 是否汇总数据，默认不汇总。
	// 不汇总：返回在统计区间内第三方平台下所有企业的每日明细，即每个企业N条数据，N为统计天数；
	// 汇总：返回在统计区间内第三方平台下所有企业的汇总后数据，即每个企业一条数据；
	NeedAggregate *bool `json:"NeedAggregate,omitnil" name:"NeedAggregate"`

	// 单次返回的最多条目数量。默认为1000，且不能超过1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认是0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type DescribeUsageRequest struct {
	*tchttp.BaseRequest
	
	// 应用信息，此接口Agent.AppId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 开始时间，例如：2021-03-21
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 结束时间，例如：2021-06-21；
	// 开始时间到结束时间的区间长度小于等于90天。
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 是否汇总数据，默认不汇总。
	// 不汇总：返回在统计区间内第三方平台下所有企业的每日明细，即每个企业N条数据，N为统计天数；
	// 汇总：返回在统计区间内第三方平台下所有企业的汇总后数据，即每个企业一条数据；
	NeedAggregate *bool `json:"NeedAggregate,omitnil" name:"NeedAggregate"`

	// 单次返回的最多条目数量。默认为1000，且不能超过1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认是0。
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
	//   AUTO_SIGN             企业静默签（自动签署）
	//   OVERSEA_SIGN          企业与港澳台居民*签署合同
	//   MOBILE_CHECK_APPROVER 使用手机号验证签署方身份
	//   PAGING_SEAL           骑缝章
	//   DOWNLOAD_FLOW         授权平台企业下载合同 
	Type *string `json:"Type,omitnil" name:"Type"`

	// 扩展服务名称 
	Name *string `json:"Name,omitnil" name:"Name"`

	// 服务状态 
	// ENABLE 开启 
	// DISABLE 关闭
	Status *string `json:"Status,omitnil" name:"Status"`

	// 最近操作人第三方应用平台的用户openid
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorOpenId *string `json:"OperatorOpenId,omitnil" name:"OperatorOpenId"`

	// 最近操作时间戳，单位秒
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

type FilledComponent struct {
	// 控件Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentId *string `json:"ComponentId,omitnil" name:"ComponentId"`

	// 控件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// 控件填写状态；0-未填写；1-已填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentFillStatus *string `json:"ComponentFillStatus,omitnil" name:"ComponentFillStatus"`

	// 控件填写内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentValue *string `json:"ComponentValue,omitnil" name:"ComponentValue"`

	// 图片填充控件下载链接，如果是图片填充控件时，这里返回图片的下载链接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type Filter struct {
	// 查询过滤条件的Key
	Key *string `json:"Key,omitnil" name:"Key"`

	// 查询过滤条件的Value列表
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type FlowApproverDetail struct {
	// 模板配置时候的签署人id,与控件绑定
	ReceiptId *string `json:"ReceiptId,omitnil" name:"ReceiptId"`

	// 平台企业的第三方id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil" name:"ProxyOrganizationOpenId"`

	// 平台企业操作人的第三方id
	ProxyOperatorOpenId *string `json:"ProxyOperatorOpenId,omitnil" name:"ProxyOperatorOpenId"`

	// 平台企业名称
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil" name:"ProxyOrganizationName"`

	// 签署人手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 签署人签署顺序
	SignOrder *int64 `json:"SignOrder,omitnil" name:"SignOrder"`

	// 签署人姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveName *string `json:"ApproveName,omitnil" name:"ApproveName"`

	// 当前签署人的状态, 状态如下
	// <br/>PENDING 待签署	
	// <br/>FILLPENDING 待填写
	// <br/>FILLACCEPT 填写完成	
	// <br/>FILLREJECT 拒绝填写	
	// <br/>WAITPICKUP 待领取	
	// <br/>ACCEPT 已签署	
	// <br/>REJECT 拒签 
	// <br/>DEADLINE 过期没人处理 
	// <br/>CANCEL 流程已撤回	
	// <br/>FORWARD 已经转他人处理
	// <br/>STOP 流程已终止	
	// <br/>RELIEVED 解除协议（已解除）
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveStatus *string `json:"ApproveStatus,omitnil" name:"ApproveStatus"`

	// 签署人自定义信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveMessage *string `json:"ApproveMessage,omitnil" name:"ApproveMessage"`

	// 签署人签署时间戳，单位秒
	ApproveTime *int64 `json:"ApproveTime,omitnil" name:"ApproveTime"`

	// 参与者类型 
	// <br/>ORGANIZATION：企业签署人
	// <br/>PERSON：个人签署人
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveType *string `json:"ApproveType,omitnil" name:"ApproveType"`

	// 自定义签署人角色
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverRoleName *string `json:"ApproverRoleName,omitnil" name:"ApproverRoleName"`
}

type FlowApproverInfo struct {
	// 签署人姓名，最大长度50个字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 签署人的证件类型
	// 1.ID_CARD 居民身份证
	// 2.HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证
	// 3.HONGKONG_AND_MACAO 港澳居民来往内地通行证
	// 4.OTHER_CARD_TYPE 其他（需要使用该类型请先联系运营经理）
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 签署人证件号（长度不超过18位）
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 签署人手机号，脱敏显示。大陆手机号为11位，暂不支持海外手机号。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 企业签署方工商营业执照上的企业名称，签署方为非发起方企业场景下必传，最大长度64个字符；
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 指定签署人非第三方平台子客企业下员工，在ApproverType为ORGANIZATION时指定。
	// 默认为false，即签署人位于同一个第三方平台应用号下；默认为false，即签署人位于同一个第三方应用号下；
	NotChannelOrganization *bool `json:"NotChannelOrganization,omitnil" name:"NotChannelOrganization"`

	// 用户侧第三方id，最大长度64个字符
	// 当签署方为同一第三方平台下的员工时，该字段若不指定，则发起【待领取】的流程
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 企业签署方在同一第三方平台应用下的其他合作企业OpenId，签署方为非发起方企业场景下必传，最大长度64个字符；
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 签署人类型
	// PERSON-个人/自然人；
	// PERSON_AUTO_SIGN-个人自动签署，适用于个人自动签场景
	// 注: 个人自动签场景为白名单功能, 使用前请联系对接的客户经理沟通。
	// ORGANIZATION-企业（企业签署方或模板发起时的企业静默签）；
	// ENTERPRISESERVER-企业自动签（他方企业自动签署或文件发起时的本方企业自动签）
	// 
	// 若要实现他方企业（同一应用下）自动签，需要满足3个条件：
	// 条件1：ApproverType 设置为ENTERPRISESERVER
	// 条件2：子客之间完成授权
	// 条件3：联系对接的客户经理沟通
	ApproverType *string `json:"ApproverType,omitnil" name:"ApproverType"`

	// 签署流程签署人在模板中对应的签署人Id；在非单方签署、以及非B2C签署的场景下必传，用于指定当前签署方在签署流程中的位置；
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 签署截止时间戳，默认一年
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

	// 合同的强制预览时间：3~300s，未指定则按合同页数计算
	PreReadTime *int64 `json:"PreReadTime,omitnil" name:"PreReadTime"`

	// 签署完前端跳转的url，此字段的用法场景请联系客户经理确认
	JumpUrl *string `json:"JumpUrl,omitnil" name:"JumpUrl"`

	// 签署人个性化能力值
	ApproverOption *ApproverOption `json:"ApproverOption,omitnil" name:"ApproverOption"`

	// 当前签署方进行签署操作是否需要企业内部审批，true 则为需要
	ApproverNeedSignReview *bool `json:"ApproverNeedSignReview,omitnil" name:"ApproverNeedSignReview"`

	// 签署人查看合同时认证方式, 1-实名查看 2-短信验证码查看(企业签署方不支持该方式) 如果不传默认为1
	// 查看合同的认证方式 Flow层级的优先于approver层级的
	// （当手写签名方式为OCR_ESIGN时，合同认证方式2无效，因为这种签名方式依赖实名认证）
	ApproverVerifyTypes []*int64 `json:"ApproverVerifyTypes,omitnil" name:"ApproverVerifyTypes"`

	// 签署人签署合同时的认证方式
	// 1-人脸认证 2-签署密码 3-运营商三要素(默认为1,2)
	ApproverSignTypes []*int64 `json:"ApproverSignTypes,omitnil" name:"ApproverSignTypes"`

	// 签署ID
	// - 发起流程时系统自动补充
	// - 创建签署链接时，可以通过查询详情接口获得签署人的SignId，然后可传入此值为该签署人创建签署链接，无需再传姓名、手机号、证件号等其他信息
	SignId *string `json:"SignId,omitnil" name:"SignId"`

	// SMS: 短信(需确保“电子签短信通知签署方”功能是开启状态才能生效); NONE: 不发信息
	// 默认为SMS(签署方为子客时该字段不生效)
	NotifyType *string `json:"NotifyType,omitnil" name:"NotifyType"`

	// [通过文件创建签署流程](https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowByFiles)时,如果设置了外层参数SignBeanTag=1(允许签署过程中添加签署控件),则可通过此参数明确规定合同所使用的签署控件类型（骑缝章、普通章法人章等）和具体的印章（印章ID）或签名方式。
	// 
	// 注：`限制印章控件或骑缝章控件情况下,仅本企业签署方可以指定具体印章（通过传递ComponentValue,支持多个），他方企业或个人只支持限制控件类型。`
	AddSignComponentsLimits []*ComponentLimit `json:"AddSignComponentsLimits,omitnil" name:"AddSignComponentsLimits"`

	// 自定义签署方角色名称
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
	// 签署短链接，不支持小程序嵌入，只支持移动端浏览器打开。注意该链接有效期为30分钟，同时需要注意保密，不要外泄给无关用户。
	SignUrl *string `json:"SignUrl,omitnil" name:"SignUrl"`

	// 签署人类型 PERSON-个人
	ApproverType *string `json:"ApproverType,omitnil" name:"ApproverType"`

	// 签署人姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 签署人手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 签署长链接，支持小程序嵌入。注意该链接有效期为30分钟，同时需要注意保密，不要外泄给无关用户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LongUrl *string `json:"LongUrl,omitnil" name:"LongUrl"`
}

type FlowDetailInfo struct {
	// 合同(流程)的Id
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 合同(流程)的名字
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 合同(流程)的类型
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 合同(流程)的状态, 状态如下
	// 
	// INIT 合同创建
	// PART 合同签署中
	// REJECT 合同拒签
	// ALL 合同签署完成
	// DEADLINE 合同流签(合同过期)
	// CANCEL 合同撤回
	// RELIEVED 解除协议（已解除）
	//  
	FlowStatus *string `json:"FlowStatus,omitnil" name:"FlowStatus"`

	// 合同(流程)的信息
	FlowMessage *string `json:"FlowMessage,omitnil" name:"FlowMessage"`

	// 合同(流程)的创建时间戳，单位秒
	CreateOn *int64 `json:"CreateOn,omitnil" name:"CreateOn"`

	// 合同(流程)的签署截止时间戳，单位秒
	DeadLine *int64 `json:"DeadLine,omitnil" name:"DeadLine"`

	// 用户自定义数据
	CustomData *string `json:"CustomData,omitnil" name:"CustomData"`

	// 合同(流程)的签署人数组
	FlowApproverInfos []*FlowApproverDetail `json:"FlowApproverInfos,omitnil" name:"FlowApproverInfos"`

	// 合同(流程)关注方信息列表
	CcInfos []*FlowApproverDetail `json:"CcInfos,omitnil" name:"CcInfos"`

	// 是否需要发起前审批，当NeedCreateReview为true，表明当前流程是需要发起前审核的合同，可能无法进行查看，签署操作，需要等审核完成后，才可以继续后续流程
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
	// 发起方企业经办人（即签署人为发起方企业员工）是否需要对子合同进行独立的意愿确认：
	// fasle：发起方企业经办人签署时对所有子合同进行统一的意愿确认
	// true：发起方企业经办人签署时需要对子合同进行独立的意愿确认
	// 默认为fasle。
	SelfOrganizationApproverSignEach *bool `json:"SelfOrganizationApproverSignEach,omitnil" name:"SelfOrganizationApproverSignEach"`

	// 非发起方企业经办人（即：签署人为个人或者不为发起方企业的员工）是否需要对子合同进行独立的意愿确认：
	// fasle：非发起方企业经办人签署时对所有子合同进行统一的意愿确认
	// true：非发起方企业经办人签署时需要对子合同进行独立的意愿确认
	// 默认为false。
	OtherApproverSignEach *bool `json:"OtherApproverSignEach,omitnil" name:"OtherApproverSignEach"`
}

type FlowInfo struct {
	// 合同名字，最大长度200个字符
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 签署截止时间戳，超过有效签署时间则该签署流程失败，默认一年
	Deadline *int64 `json:"Deadline,omitnil" name:"Deadline"`

	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 多个签署人信息，最大支持50个签署方
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitnil" name:"FlowApprovers"`

	// 表单K-V对列表
	FormFields []*FormField `json:"FormFields,omitnil" name:"FormFields"`

	// 回调地址，最大长度1000个字符
	CallbackUrl *string `json:"CallbackUrl,omitnil" name:"CallbackUrl"`

	// 合同类型，如：1. “劳务”；2. “销售”；3. “租赁”；4. “其他”，最大长度200个字符
	FlowType *string `json:"FlowType,omitnil" name:"FlowType"`

	// 合同描述，最大长度1000个字符
	FlowDescription *string `json:"FlowDescription,omitnil" name:"FlowDescription"`

	//  第三方应用平台的业务信息，最大长度1000个字符。
	CustomerData *string `json:"CustomerData,omitnil" name:"CustomerData"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitnil" name:"CustomShowMap"`

	// 被抄送人的信息列表，抄送功能暂不开放
	CcInfos []*CcInfo `json:"CcInfos,omitnil" name:"CcInfos"`

	// 发起方企业的签署人进行签署操作是否需要企业内部审批。
	// 若设置为true,审核结果需通过接口 ChannelCreateFlowSignReview 通知电子签，审核通过后，发起方企业签署人方可进行签署操作，否则会阻塞其签署操作。
	// 
	// 注：企业可以通过此功能与企业内部的审批流程进行关联，支持手动、静默签署合同。
	NeedSignReview *bool `json:"NeedSignReview,omitnil" name:"NeedSignReview"`

	// 给关注人发送短信通知的类型，0-合同发起时通知 1-签署完成后通知
	CcNotifyType *int64 `json:"CcNotifyType,omitnil" name:"CcNotifyType"`

	// 个人自动签场景。发起自动签署时，需设置对应自动签署场景，目前仅支持场景：处方单-E_PRESCRIPTION_AUTO_SIGN
	AutoSignScene *string `json:"AutoSignScene,omitnil" name:"AutoSignScene"`
}

type FlowResourceUrlInfo struct {
	// 流程对应Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 流程对应资源链接信息数组
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 文件夹数组，签署流程总数不能超过50个，一个文件夹下，不能超过20个签署流程
	DownLoadFlows []*DownloadFlowInfo `json:"DownLoadFlows,omitnil" name:"DownLoadFlows"`

	// 操作者的信息，不用传
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type GetDownloadFlowUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 文件夹数组，签署流程总数不能超过50个，一个文件夹下，不能超过20个签署流程
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
	// 合同（流程）下载地址
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	// 
	// 注: 此接口 参数Agent. ProxyOperator.OpenId 需要传递超管或者法人的OpenId
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	//   扩展服务类型
	//   AUTO_SIGN             企业静默签（自动签署）
	//   OVERSEA_SIGN          企业与港澳台居民*签署合同
	//   MOBILE_CHECK_APPROVER 使用手机号验证签署方身份
	//   PAGING_SEAL           骑缝章
	//   DOWNLOAD_FLOW         授权渠道下载合同 
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 操作类型 
	// OPEN:开通 
	// CLOSE:关闭
	Operate *string `json:"Operate,omitnil" name:"Operate"`

	// 链接跳转类型，支持以下类型
	// <ul><li>WEIXINAPP : 短链直接跳转到电子签小程序  (默认值)</li>
	// <li>APP : 第三方APP或小程序跳转电子签小程序</li></ul>
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`
}

type ModifyExtendedServiceRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	// 
	// 注: 此接口 参数Agent. ProxyOperator.OpenId 需要传递超管或者法人的OpenId
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	//   扩展服务类型
	//   AUTO_SIGN             企业静默签（自动签署）
	//   OVERSEA_SIGN          企业与港澳台居民*签署合同
	//   MOBILE_CHECK_APPROVER 使用手机号验证签署方身份
	//   PAGING_SEAL           骑缝章
	//   DOWNLOAD_FLOW         授权渠道下载合同 
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 操作类型 
	// OPEN:开通 
	// CLOSE:关闭
	Operate *string `json:"Operate,omitnil" name:"Operate"`

	// 链接跳转类型，支持以下类型
	// <ul><li>WEIXINAPP : 短链直接跳转到电子签小程序  (默认值)</li>
	// <li>APP : 第三方APP或小程序跳转电子签小程序</li></ul>
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
	// 应用相关信息。 
	// 此接口Agent.AppId必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作类型，
	// 查询:"SELECT"，
	// 删除:"DELETE"，
	// 更新:"UPDATE"
	OperateType *string `json:"OperateType,omitnil" name:"OperateType"`

	// 第三方应用平台模板库模板唯一标识
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 合作企业方第三方机构唯一标识数据.
	// 支持多个， 用","进行分隔
	ProxyOrganizationOpenIds *string `json:"ProxyOrganizationOpenIds,omitnil" name:"ProxyOrganizationOpenIds"`

	// 模板可见性, 
	// 全部可见-"all",
	//  部分可见-"part"
	AuthTag *string `json:"AuthTag,omitnil" name:"AuthTag"`

	// 当OperateType=UPDATE时，可以通过设置此字段对模板启停用状态进行操作。
	// 若此字段值为0，则不会修改模板Available，
	// 1为启用模板，
	// 2为停用模板。
	// 启用后模板可以正常领取。停用后，推送方式为【自动推送】的模板则无法被子客使用，推送方式为【手动领取】的模板则无法出现被模板库被子客领用。如果Available更新失败，会直接返回错误。
	Available *int64 `json:"Available,omitnil" name:"Available"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type OperateChannelTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 
	// 此接口Agent.AppId必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作类型，
	// 查询:"SELECT"，
	// 删除:"DELETE"，
	// 更新:"UPDATE"
	OperateType *string `json:"OperateType,omitnil" name:"OperateType"`

	// 第三方应用平台模板库模板唯一标识
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 合作企业方第三方机构唯一标识数据.
	// 支持多个， 用","进行分隔
	ProxyOrganizationOpenIds *string `json:"ProxyOrganizationOpenIds,omitnil" name:"ProxyOrganizationOpenIds"`

	// 模板可见性, 
	// 全部可见-"all",
	//  部分可见-"part"
	AuthTag *string `json:"AuthTag,omitnil" name:"AuthTag"`

	// 当OperateType=UPDATE时，可以通过设置此字段对模板启停用状态进行操作。
	// 若此字段值为0，则不会修改模板Available，
	// 1为启用模板，
	// 2为停用模板。
	// 启用后模板可以正常领取。停用后，推送方式为【自动推送】的模板则无法被子客使用，推送方式为【手动领取】的模板则无法出现被模板库被子客领用。如果Available更新失败，会直接返回错误。
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
	// 腾讯电子签颁发给第三方应用平台的应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 第三方应用平台模板库模板唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 描述模板可见性更改的结果，和参数中Available无关。
	// 全部成功-"all-success",
	// 部分成功-"part-success", 
	// 全部失败-"fail"，失败的会在FailMessageList中展示。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateResult *string `json:"OperateResult,omitnil" name:"OperateResult"`

	// 模板可见性, 
	// 全部可见-"all", 
	// 部分可见-"part"
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthTag *string `json:"AuthTag,omitnil" name:"AuthTag"`

	// 合作企业方第三方机构唯一标识数据
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
	// 验签结果。0-签名域未签名；1-验签成功； 3-验签失败；4-未找到签名域：文件内没有签名域；5-签名值格式不正确。
	VerifyResult *int64 `json:"VerifyResult,omitnil" name:"VerifyResult"`

	// 签署平台，如果文件是在腾讯电子签平台签署，则返回腾讯电子签，如果文件不在腾讯电子签平台签署，则返回其他平台。
	SignPlatform *string `json:"SignPlatform,omitnil" name:"SignPlatform"`

	// 签署人名称
	SignerName *string `json:"SignerName,omitnil" name:"SignerName"`

	// 签署时间戳，单位秒
	SignTime *int64 `json:"SignTime,omitnil" name:"SignTime"`

	// 签名算法
	SignAlgorithm *string `json:"SignAlgorithm,omitnil" name:"SignAlgorithm"`

	// 签名证书序列号
	CertSn *string `json:"CertSn,omitnil" name:"CertSn"`

	// 证书起始时间戳，单位秒
	CertNotBefore *int64 `json:"CertNotBefore,omitnil" name:"CertNotBefore"`

	// 证书过期时间戳，单位秒
	CertNotAfter *int64 `json:"CertNotAfter,omitnil" name:"CertNotAfter"`

	// 签名类型
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
	// 对应Agent-ProxyOperator-OpenId。第三方应用平台自定义，对子客企业员的唯一标识。一个OpenId在一个子客企业内唯一对应一个真实员工，不可在其他子客企业内重复使用。（例如，可以使用经办人企业名+员工身份证的hash值，需要第三方应用平台保存），最大64位字符串
	Id *string `json:"Id,omitnil" name:"Id"`

	// 经办人姓名，最大长度50个字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 经办人身份证件类型
	// 1.ID_CARD 居民身份证
	// 2.HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证
	// 3.HONGKONG_AND_MACAO 港澳居民来往内地通行证
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 经办人证件号
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 经办人手机号，大陆手机号输入11位，暂不支持海外手机号。
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 默认角色，值为以下三个对应的英文：
	// 业务管理员：admin
	// 经办人：channel-normal-operator
	// 业务员：channel-sales-man
	DefaultRole *string `json:"DefaultRole,omitnil" name:"DefaultRole"`
}

type Recipient struct {
	// 签署人唯一标识，在通过模板发起合同的时候对应签署方ID
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 参与者类型，默认为空。
	// ENTERPRISE-企业；
	// INDIVIDUAL-个人；
	// PROMOTER-发起方
	RecipientType *string `json:"RecipientType,omitnil" name:"RecipientType"`

	// 描述信息	
	Description *string `json:"Description,omitnil" name:"Description"`

	// 角色名称	
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
}

type RecipientComponentInfo struct {
	// 参与方Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"RecipientId,omitnil" name:"RecipientId"`

	// 参与方填写状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientFillStatus *string `json:"RecipientFillStatus,omitnil" name:"RecipientFillStatus"`

	// 是否发起方
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPromoter *bool `json:"IsPromoter,omitnil" name:"IsPromoter"`

	// 填写控件内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components []*FilledComponent `json:"Components,omitnil" name:"Components"`
}

type ReleasedApprover struct {
	// 企业签署方工商营业执照上的企业名称，签署方为非发起方企业场景下必传，最大长度64个字符
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// 签署人在原流程中的签署人列表中的顺序序号（从0开始，按顺序依次递增），如果不清楚原流程中的签署人列表，可以通过DescribeFlows接口查看
	ApproverNumber *uint64 `json:"ApproverNumber,omitnil" name:"ApproverNumber"`

	// 签署人类型，目前仅支持
	// ORGANIZATION-企业
	// ENTERPRISESERVER-企业静默签
	ApproverType *string `json:"ApproverType,omitnil" name:"ApproverType"`

	// 签署人姓名，最大长度50个字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 签署人身份证件类型
	// 1.ID_CARD 居民身份证
	// 2.HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证
	// 3.HONGKONG_AND_MACAO 港澳居民来往内地通行证
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 签署人证件号
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 签署人手机号，脱敏显示。大陆手机号为11位，暂不支持海外手机号
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// 企业签署方在同一第三方应用下的其他合作企业OpenId，签署方为非发起方企业场景下必传，最大长度64个字符
	OrganizationOpenId *string `json:"OrganizationOpenId,omitnil" name:"OrganizationOpenId"`

	// 用户侧第三方id，最大长度64个字符
	// 当签署方为同一第三方应用下的员工时，该字必传
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 签署控件类型，支持自定义企业签署方的签署控件为“印章”或“签名”
	// - SIGN_SEAL-默认为印章控件类型
	// - SIGN_SIGNATURE-手写签名控件类型
	ApproverSignComponentType *string `json:"ApproverSignComponentType,omitnil" name:"ApproverSignComponentType"`

	// 签署方自定义控件别名，最大长度20个字符
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
	// 是否能够催办，true-是，false-否
	CanRemind *bool `json:"CanRemind,omitnil" name:"CanRemind"`

	// 合同id
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 催办详情信息
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
	// 二维码id
	QrCodeId *string `json:"QrCodeId,omitnil" name:"QrCodeId"`

	// 二维码url
	QrCodeUrl *string `json:"QrCodeUrl,omitnil" name:"QrCodeUrl"`

	// 二维码过期时间
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`
}

type SignUrl struct {
	// 小程序签署链接
	AppSignUrl *string `json:"AppSignUrl,omitnil" name:"AppSignUrl"`

	// 签署链接有效时间
	EffectiveTime *string `json:"EffectiveTime,omitnil" name:"EffectiveTime"`

	// 移动端签署链接
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

	// 参与者类型:
	// ORGANIZATION 企业经办人
	// PERSON 自然人
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
	// 对应Agent-ProxyOperator-OpenId。第三方应用平台自定义，对子客企业员的唯一标识。一个OpenId在一个子客企业内唯一对应一个真实员工，不可在其他子客企业内重复使用。（例如，可以使用经办人企业名+员工身份证的hash值，需要第三方应用平台保存），最大64位字符串
	Id *string `json:"Id,omitnil" name:"Id"`

	// 失败原因
	// 例如：Id不符合规范、证件号码不合法等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`
}

// Predefined struct for user
type SyncProxyOrganizationOperatorsRequestParams struct {
	// 应用相关信息。 此接口Agent.AppId 和 Agent.ProxyOrganizationOpenId必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作类型，新增:"CREATE"，修改:"UPDATE"，离职:"RESIGN"
	OperatorType *string `json:"OperatorType,omitnil" name:"OperatorType"`

	// 经办人信息列表，最大长度200
	ProxyOrganizationOperators []*ProxyOrganizationOperator `json:"ProxyOrganizationOperators,omitnil" name:"ProxyOrganizationOperators"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type SyncProxyOrganizationOperatorsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.AppId 和 Agent.ProxyOrganizationOpenId必填。
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 操作类型，新增:"CREATE"，修改:"UPDATE"，离职:"RESIGN"
	OperatorType *string `json:"OperatorType,omitnil" name:"OperatorType"`

	// 经办人信息列表，最大长度200
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
	// Status 同步状态,全部同步失败接口会直接报错
	// 1-成功 
	// 2-部分成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 同步失败经办人及其失败原因
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
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 第三方平台子客企业名称，最大长度64个字符
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil" name:"ProxyOrganizationName"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitnil" name:"BusinessLicense"`

	// 第三方平台子客企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil" name:"UniformSocialCreditCode"`

	// 第三方平台子客企业法人/负责人姓名
	ProxyLegalName *string `json:"ProxyLegalName,omitnil" name:"ProxyLegalName"`

	// 暂未开放
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 第三方平台子客企业法人/负责人证件类型，默认居民身份证（ID_CARD）类型，暂不支持其他类型
	ProxyLegalIdCardType *string `json:"ProxyLegalIdCardType,omitnil" name:"ProxyLegalIdCardType"`

	// 第三方平台子客企业法人/负责人证件号
	ProxyLegalIdCardNumber *string `json:"ProxyLegalIdCardNumber,omitnil" name:"ProxyLegalIdCardNumber"`
}

type SyncProxyOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 第三方平台子客企业名称，最大长度64个字符
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil" name:"ProxyOrganizationName"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitnil" name:"BusinessLicense"`

	// 第三方平台子客企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil" name:"UniformSocialCreditCode"`

	// 第三方平台子客企业法人/负责人姓名
	ProxyLegalName *string `json:"ProxyLegalName,omitnil" name:"ProxyLegalName"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`

	// 第三方平台子客企业法人/负责人证件类型，默认居民身份证（ID_CARD）类型，暂不支持其他类型
	ProxyLegalIdCardType *string `json:"ProxyLegalIdCardType,omitnil" name:"ProxyLegalIdCardType"`

	// 第三方平台子客企业法人/负责人证件号
	ProxyLegalIdCardNumber *string `json:"ProxyLegalIdCardNumber,omitnil" name:"ProxyLegalIdCardNumber"`
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
	// 应用相关信息，若是第三方应用集成调用 若是第三方应用集成调用,Agent.AppId 和 Agent.ProxyOrganizationOpenId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 文件对应业务类型
	// 1. TEMPLATE - 模板； 文件类型：.pdf/.doc/.docx/.html
	// 2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html
	BusinessType *string `json:"BusinessType,omitnil" name:"BusinessType"`

	// 上传文件内容数组，最多支持20个文件
	FileInfos []*UploadFile `json:"FileInfos,omitnil" name:"FileInfos"`

	// 操作者的信息
	//
	// Deprecated: Operator is deprecated.
	Operator *UserInfo `json:"Operator,omitnil" name:"Operator"`
}

type UploadFilesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息，若是第三方应用集成调用 若是第三方应用集成调用,Agent.AppId 和 Agent.ProxyOrganizationOpenId 必填
	Agent *Agent `json:"Agent,omitnil" name:"Agent"`

	// 文件对应业务类型
	// 1. TEMPLATE - 模板； 文件类型：.pdf/.doc/.docx/.html
	// 2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html
	BusinessType *string `json:"BusinessType,omitnil" name:"BusinessType"`

	// 上传文件内容数组，最多支持20个文件
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 文件id数组，有效期一个小时；有效期内此文件id可以反复使用
	FileIds []*string `json:"FileIds,omitnil" name:"FileIds"`

	// 文件Url
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
	// 子客企业唯一标识
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitnil" name:"ProxyOrganizationOpenId"`

	// 子客企业名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitnil" name:"ProxyOrganizationName"`

	// 日期，当需要汇总数据时日期为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitnil" name:"Date"`

	// 消耗数量
	Usage *uint64 `json:"Usage,omitnil" name:"Usage"`

	// 撤回数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cancel *uint64 `json:"Cancel,omitnil" name:"Cancel"`

	// 消耗渠道
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowChannel *string `json:"FlowChannel,omitnil" name:"FlowChannel"`
}

type UserInfo struct {
	// 第三方应用平台自定义，对应第三方平台子客企业员的唯一标识。一个OpenId在一个子客企业内唯一对应一个真实员工，不可在其他子客企业内重复使用。（例如，可以使用经办人企业名+员工身份证的hash值，需要第三方应用平台保存），最大64位字符串
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
	// 姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 证件类型: 
	// ID_CARD 身份证
	// HONGKONG_AND_MACAO 港澳居民来往内地通行证
	// HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证(格式同居民身份证)
	IdCardType *string `json:"IdCardType,omitnil" name:"IdCardType"`

	// 证件号，如果有 X 请大写
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