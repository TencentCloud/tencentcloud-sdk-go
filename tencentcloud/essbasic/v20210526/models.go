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
	// 应用的唯一标识。不同的业务系统可以采用不同的AppId，不同AppId下的数据是隔离的。可以由控制台开发者中心-应用集成自主生成。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 第三方应用平台自定义，对应第三方平台子客企业的唯一标识。一个第三方平台子客企业主体与子客企业ProxyOrganizationOpenId是一一对应的，不可更改，不可重复使用。（例如，可以使用企业名称的hash值，或者社会统一信用代码的hash值，或者随机hash值，需要第三方应用平台保存），最大64位字符串
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitempty" name:"ProxyOrganizationOpenId"`

	// 第三方平台子客企业中的员工/经办人，通过第三方应用平台进入电子签完成实名、且被赋予相关权限后，可以参与到企业资源的管理或签署流程中。
	ProxyOperator *UserInfo `json:"ProxyOperator,omitempty" name:"ProxyOperator"`

	// 在第三方平台子客企业开通电子签后，会生成唯一的子客应用Id（ProxyAppId）用于代理调用时的鉴权，在子客开通的回调中获取。
	ProxyAppId *string `json:"ProxyAppId,omitempty" name:"ProxyAppId"`

	// 内部参数，暂未开放使用
	ProxyOrganizationId *string `json:"ProxyOrganizationId,omitempty" name:"ProxyOrganizationId"`
}

type ApproverOption struct {
	// 是否隐藏一键签署 false-不隐藏,默认 true-隐藏
	HideOneKeySign *bool `json:"HideOneKeySign,omitempty" name:"HideOneKeySign"`
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

type AuthFailMessage struct {
	// 合作企业Id
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitempty" name:"ProxyOrganizationOpenId"`

	// 出错信息
	Message *string `json:"Message,omitempty" name:"Message"`
}

type AuthorizedUser struct {
	// 用户openid
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
}

type BaseFlowInfo struct {
	// 合同流程名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 合同流程类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 合同流程描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 合同流程截止时间，unix时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 是否顺序签署(true:无序签,false:顺序签)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unordered *bool `json:"Unordered,omitempty" name:"Unordered"`

	// 打开智能添加填写区(默认开启，打开:"OPEN" 关闭："CLOSE")
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntelligentStatus *string `json:"IntelligentStatus,omitempty" name:"IntelligentStatus"`

	// 填写控件内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormFields []*FormField `json:"FormFields,omitempty" name:"FormFields"`

	// 本企业(发起方企业)是否需要签署审批，true：开启本企业签署审批
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeedSignReview *bool `json:"NeedSignReview,omitempty" name:"NeedSignReview"`
}

type CcInfo struct {
	// 被抄送人手机号，大陆11位手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 被抄送人姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 被抄送人类型
	// 0--个人. 1--员工
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcType *int64 `json:"CcType,omitempty" name:"CcType"`

	// 被抄送人权限
	// 0--可查看
	// 1--可查看也可下载
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcPermission *int64 `json:"CcPermission,omitempty" name:"CcPermission"`
}

// Predefined struct for user
type ChannelBatchCancelFlowsRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程Id数组，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 撤销理由
	CancelMessage *string `json:"CancelMessage,omitempty" name:"CancelMessage"`

	// 撤销理由自定义格式；选项：
	// 0 默认格式
	// 1 只保留身份信息：展示为【发起方】
	// 2 保留身份信息+企业名称：展示为【发起方xxx公司】
	// 3 保留身份信息+企业名称+经办人名称：展示为【发起方xxxx公司-经办人姓名】
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitempty" name:"CancelMessageFormat"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelBatchCancelFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程Id数组，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 撤销理由
	CancelMessage *string `json:"CancelMessage,omitempty" name:"CancelMessage"`

	// 撤销理由自定义格式；选项：
	// 0 默认格式
	// 1 只保留身份信息：展示为【发起方】
	// 2 保留身份信息+企业名称：展示为【发起方xxx公司】
	// 3 保留身份信息+企业名称+经办人名称：展示为【发起方xxxx公司-经办人姓名】
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitempty" name:"CancelMessageFormat"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	FailMessages []*string `json:"FailMessages,omitempty" name:"FailMessages"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 撤回原因，最大不超过200字符
	CancelMessage *string `json:"CancelMessage,omitempty" name:"CancelMessage"`

	// 撤销理由自定义格式；选项：
	// 0 默认格式
	// 1 只保留身份信息：展示为【发起方】
	// 2 保留身份信息+企业名称：展示为【发起方xxx公司】
	// 3 保留身份信息+企业名称+经办人名称：展示为【发起方xxxx公司-经办人姓名】
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitempty" name:"CancelMessageFormat"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelCancelFlowRequest struct {
	*tchttp.BaseRequest
	
	// 签署流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 撤回原因，最大不超过200字符
	CancelMessage *string `json:"CancelMessage,omitempty" name:"CancelMessage"`

	// 撤销理由自定义格式；选项：
	// 0 默认格式
	// 1 只保留身份信息：展示为【发起方】
	// 2 保留身份信息+企业名称：展示为【发起方xxx公司】
	// 3 保留身份信息+企业名称+经办人名称：展示为【发起方xxxx公司-经办人姓名】
	CancelMessageFormat *int64 `json:"CancelMessageFormat,omitempty" name:"CancelMessageFormat"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 二维码id
	QrCodeId *string `json:"QrCodeId,omitempty" name:"QrCodeId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelCancelMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 二维码id
	QrCodeId *string `json:"QrCodeId,omitempty" name:"QrCodeId"`

	// 暂未开放
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
type ChannelCreateBatchCancelFlowUrlRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程Id数组
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelCreateBatchCancelFlowUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程Id数组
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	BatchCancelFlowUrl *string `json:"BatchCancelFlowUrl,omitempty" name:"BatchCancelFlowUrl"`

	// 签署流程批量撤销失败原因
	FailMessages []*string `json:"FailMessages,omitempty" name:"FailMessages"`

	// 签署撤销url过期时间-年月日-时分秒
	UrlExpireOn *string `json:"UrlExpireOn,omitempty" name:"UrlExpireOn"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ChannelCreateBoundFlowsRequestParams struct {
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 领取的合同id列表
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelCreateBoundFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 领取的合同id列表
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 资源类型 取值范围doc,docx,html,xls,xlsx之一
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，长度限制为256字符
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源Id，通过UploadFiles获取
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 调用方用户信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`
}

type ChannelCreateConvertTaskApiRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 资源类型 取值范围doc,docx,html,xls,xlsx之一
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，长度限制为256字符
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源Id，通过UploadFiles获取
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 调用方用户信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// WEB嵌入资源类型，取值范围：CREATE_SEAL创建印章，CREATE_TEMPLATE创建模板，MODIFY_TEMPLATE修改模板，PREVIEW_TEMPLATE预览模板，PREVIEW_FLOW预览流程
	EmbedType *string `json:"EmbedType,omitempty" name:"EmbedType"`

	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 渠道操作者信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// WEB嵌入的业务资源ID，EmbedType取值MODIFY_TEMPLATE或PREVIEW_TEMPLATE或 PREVIEW_FLOW时BusinessId必填
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 是否隐藏控件，只有预览模板时生效
	HiddenComponents *bool `json:"HiddenComponents,omitempty" name:"HiddenComponents"`
}

type ChannelCreateEmbedWebUrlRequest struct {
	*tchttp.BaseRequest
	
	// WEB嵌入资源类型，取值范围：CREATE_SEAL创建印章，CREATE_TEMPLATE创建模板，MODIFY_TEMPLATE修改模板，PREVIEW_TEMPLATE预览模板，PREVIEW_FLOW预览流程
	EmbedType *string `json:"EmbedType,omitempty" name:"EmbedType"`

	// 渠道应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 渠道操作者信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// WEB嵌入的业务资源ID，EmbedType取值MODIFY_TEMPLATE或PREVIEW_TEMPLATE或 PREVIEW_FLOW时BusinessId必填
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 是否隐藏控件，只有预览模板时生效
	HiddenComponents *bool `json:"HiddenComponents,omitempty" name:"HiddenComponents"`
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
	delete(f, "EmbedType")
	delete(f, "Agent")
	delete(f, "Operator")
	delete(f, "BusinessId")
	delete(f, "HiddenComponents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateEmbedWebUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateEmbedWebUrlResponseParams struct {
	// 嵌入的web链接
	WebUrl *string `json:"WebUrl,omitempty" name:"WebUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程名称，长度不超过200个字符
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署流程签约方列表，最多不超过50个参与方
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitempty" name:"FlowApprovers"`

	// 签署文件资源Id列表，目前仅支持单个文件
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 签署文件中的发起方的填写控件，需要在发起的时候进行填充
	Components []*Component `json:"Components,omitempty" name:"Components"`

	// 签署流程截止时间，十位数时间戳，最大值为33162419560，即3020年
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 签署流程回调地址，长度不超过255个字符
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 合同签署顺序类型(无序签,顺序签)，默认为false，即有序签署。有序签署时以传入FlowApprovers数组的顺序作为签署顺序
	Unordered *bool `json:"Unordered,omitempty" name:"Unordered"`

	// 签署流程的类型，长度不超过255个字符
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 签署流程的描述，长度不超过1000个字符
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitempty" name:"CustomShowMap"`

	// 业务信息，最大长度1000个字符。发起自动签署时，需设置对应自动签署场景，目前仅支持场景：处方单-E_PRESCRIPTION_AUTO_SIGN
	CustomerData *string `json:"CustomerData,omitempty" name:"CustomerData"`

	// 发起方企业的签署人进行签署操作是否需要企业内部审批。 若设置为true,审核结果需通过接口 ChannelCreateFlowSignReview 通知电子签，审核通过后，发起方企业签署人方可进行签署操作，否则会阻塞其签署操作。  注：企业可以通过此功能与企业内部的审批流程进行关联，支持手动、静默签署合同。
	NeedSignReview *bool `json:"NeedSignReview,omitempty" name:"NeedSignReview"`

	// 签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证
	// 参数说明：可选人脸识别或手机号验证两种方式，若选择后者，未实名个人签署方在签署合同时，无需经过实名认证和意愿确认两次人脸识别，该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitempty" name:"ApproverVerifyType"`

	// 标识是否允许发起后添加控件。0为不允许1为允许。如果为1，创建的时候不能有签署控件，只能创建后添加。注意发起后添加控件功能不支持添加骑缝章和签批控件
	SignBeanTag *int64 `json:"SignBeanTag,omitempty" name:"SignBeanTag"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 被抄送人信息列表
	CcInfos []*CcInfo `json:"CcInfos,omitempty" name:"CcInfos"`

	// 给关注人发送短信通知的类型，0-合同发起时通知 1-签署完成后通知
	CcNotifyType *int64 `json:"CcNotifyType,omitempty" name:"CcNotifyType"`
}

type ChannelCreateFlowByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程名称，长度不超过200个字符
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署流程签约方列表，最多不超过50个参与方
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitempty" name:"FlowApprovers"`

	// 签署文件资源Id列表，目前仅支持单个文件
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 签署文件中的发起方的填写控件，需要在发起的时候进行填充
	Components []*Component `json:"Components,omitempty" name:"Components"`

	// 签署流程截止时间，十位数时间戳，最大值为33162419560，即3020年
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 签署流程回调地址，长度不超过255个字符
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 合同签署顺序类型(无序签,顺序签)，默认为false，即有序签署。有序签署时以传入FlowApprovers数组的顺序作为签署顺序
	Unordered *bool `json:"Unordered,omitempty" name:"Unordered"`

	// 签署流程的类型，长度不超过255个字符
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 签署流程的描述，长度不超过1000个字符
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitempty" name:"CustomShowMap"`

	// 业务信息，最大长度1000个字符。发起自动签署时，需设置对应自动签署场景，目前仅支持场景：处方单-E_PRESCRIPTION_AUTO_SIGN
	CustomerData *string `json:"CustomerData,omitempty" name:"CustomerData"`

	// 发起方企业的签署人进行签署操作是否需要企业内部审批。 若设置为true,审核结果需通过接口 ChannelCreateFlowSignReview 通知电子签，审核通过后，发起方企业签署人方可进行签署操作，否则会阻塞其签署操作。  注：企业可以通过此功能与企业内部的审批流程进行关联，支持手动、静默签署合同。
	NeedSignReview *bool `json:"NeedSignReview,omitempty" name:"NeedSignReview"`

	// 签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证
	// 参数说明：可选人脸识别或手机号验证两种方式，若选择后者，未实名个人签署方在签署合同时，无需经过实名认证和意愿确认两次人脸识别，该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitempty" name:"ApproverVerifyType"`

	// 标识是否允许发起后添加控件。0为不允许1为允许。如果为1，创建的时候不能有签署控件，只能创建后添加。注意发起后添加控件功能不支持添加骑缝章和签批控件
	SignBeanTag *int64 `json:"SignBeanTag,omitempty" name:"SignBeanTag"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 被抄送人信息列表
	CcInfos []*CcInfo `json:"CcInfos,omitempty" name:"CcInfos"`

	// 给关注人发送短信通知的类型，0-合同发起时通知 1-签署完成后通知
	CcNotifyType *int64 `json:"CcNotifyType,omitempty" name:"CcNotifyType"`
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
	delete(f, "CustomerData")
	delete(f, "NeedSignReview")
	delete(f, "ApproverVerifyType")
	delete(f, "SignBeanTag")
	delete(f, "Operator")
	delete(f, "CcInfos")
	delete(f, "CcNotifyType")
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
type ChannelCreateFlowGroupByFilesRequestParams struct {
	// 每个子合同的发起所需的信息，数量限制2-100
	FlowFileInfos []*FlowFileInfo `json:"FlowFileInfos,omitempty" name:"FlowFileInfos"`

	// 合同组名称，长度不超过200个字符
	FlowGroupName *string `json:"FlowGroupName,omitempty" name:"FlowGroupName"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证
	// 参数说明：若选择后者，未实名的个人签署方查看合同时，无需进行人脸识别实名认证（但签署合同时仍然需要人脸实名），该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitempty" name:"ApproverVerifyType"`

	// 操作者的信息，此参数不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelCreateFlowGroupByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 每个子合同的发起所需的信息，数量限制2-100
	FlowFileInfos []*FlowFileInfo `json:"FlowFileInfos,omitempty" name:"FlowFileInfos"`

	// 合同组名称，长度不超过200个字符
	FlowGroupName *string `json:"FlowGroupName,omitempty" name:"FlowGroupName"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署人校验方式
	// VerifyCheck: 人脸识别（默认）
	// MobileCheck：手机号验证
	// 参数说明：若选择后者，未实名的个人签署方查看合同时，无需进行人脸识别实名认证（但签署合同时仍然需要人脸实名），该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"ApproverVerifyType,omitempty" name:"ApproverVerifyType"`

	// 操作者的信息，此参数不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	FlowGroupId *string `json:"FlowGroupId,omitempty" name:"FlowGroupId"`

	// 子合同ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ChannelCreateFlowRemindsRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程Id数组，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`
}

type ChannelCreateFlowRemindsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程Id数组，最多100个，超过100不处理
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`
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
	RemindFlowRecords []*RemindFlowRecords `json:"RemindFlowRecords,omitempty" name:"RemindFlowRecords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 企业内部审核结果
	// PASS: 通过
	// REJECT: 拒绝
	// SIGN_REJECT:拒签(流程结束)
	ReviewType *string `json:"ReviewType,omitempty" name:"ReviewType"`

	// 审核原因 
	// 当ReviewType 是REJECT 时此字段必填,字符串长度不超过200
	ReviewMessage *string `json:"ReviewMessage,omitempty" name:"ReviewMessage"`

	// 签署节点审核时需要指定
	RecipientId *string `json:"RecipientId,omitempty" name:"RecipientId"`
}

type ChannelCreateFlowSignReviewRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 企业内部审核结果
	// PASS: 通过
	// REJECT: 拒绝
	// SIGN_REJECT:拒签(流程结束)
	ReviewType *string `json:"ReviewType,omitempty" name:"ReviewType"`

	// 审核原因 
	// 当ReviewType 是REJECT 时此字段必填,字符串长度不超过200
	ReviewMessage *string `json:"ReviewMessage,omitempty" name:"ReviewMessage"`

	// 签署节点审核时需要指定
	RecipientId *string `json:"RecipientId,omitempty" name:"RecipientId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowSignReviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowSignReviewResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 流程签署人，其中Name和Mobile必传，其他可不传，ApproverType目前只支持PERSON类型的签署人，如果不传默认为该值。还需注意签署人只能有手写签名和时间类型的签署控件，其他类型的填写控件和签署控件暂时都未支持。
	FlowApproverInfos []*FlowApproverInfo `json:"FlowApproverInfos,omitempty" name:"FlowApproverInfos"`

	// 用户信息，暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 机构信息，暂未开放
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`
}

type ChannelCreateFlowSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 流程签署人，其中Name和Mobile必传，其他可不传，ApproverType目前只支持PERSON类型的签署人，如果不传默认为该值。还需注意签署人只能有手写签名和时间类型的签署控件，其他类型的填写控件和签署控件暂时都未支持。
	FlowApproverInfos []*FlowApproverInfo `json:"FlowApproverInfos,omitempty" name:"FlowApproverInfos"`

	// 用户信息，暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 机构信息，暂未开放
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateFlowSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateFlowSignUrlResponseParams struct {
	// 签署人签署链接信息
	FlowApproverUrlInfos []*FlowApproverUrlInfo `json:"FlowApproverUrlInfos,omitempty" name:"FlowApproverUrlInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

	// 限制二维码用户条件
	Restrictions []*ApproverRestriction `json:"Restrictions,omitempty" name:"Restrictions"`

	// 回调地址，最大长度1000个字符
	// 不传默认使用第三方应用号配置的回调地址
	// 回调时机:用户通过签署二维码发起合同时，企业额度不足导致失败
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 限制二维码用户条件（已弃用）
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitempty" name:"ApproverRestrictions"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelCreateMultiFlowSignQRCodeRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。
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

	// 限制二维码用户条件
	Restrictions []*ApproverRestriction `json:"Restrictions,omitempty" name:"Restrictions"`

	// 回调地址，最大长度1000个字符
	// 不传默认使用第三方应用号配置的回调地址
	// 回调时机:用户通过签署二维码发起合同时，企业额度不足导致失败
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 限制二维码用户条件（已弃用）
	ApproverRestrictions *ApproverRestriction `json:"ApproverRestrictions,omitempty" name:"ApproverRestrictions"`

	// 暂未开放
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
	delete(f, "Restrictions")
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
	// 签署二维码对象
	QrCode *SignQrCode `json:"QrCode,omitempty" name:"QrCode"`

	// 签署链接对象
	SignUrls *SignUrl `json:"SignUrls,omitempty" name:"SignUrls"`

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

// Predefined struct for user
type ChannelCreatePrepareFlowRequestParams struct {
	// 资源id，与ResourceType对应
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源类型，1：模板，目前仅支持模板，与ResourceId对应
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 合同流程基础信息
	FlowInfo *BaseFlowInfo `json:"FlowInfo,omitempty" name:"FlowInfo"`

	// 合同签署人信息
	FlowApproverList []*CommonFlowApprover `json:"FlowApproverList,omitempty" name:"FlowApproverList"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 合同流程配置信息
	FlowOption *CreateFlowOption `json:"FlowOption,omitempty" name:"FlowOption"`

	// 该参数不可用，请通过获取 web 可嵌入接口获取合同流程预览 URL
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 该参数不可用，请通过获取 web 可嵌入接口获取合同流程预览 URL
	NeedPreview *bool `json:"NeedPreview,omitempty" name:"NeedPreview"`

	// 企业机构信息，不用传
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`

	// 操作人（用户）信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelCreatePrepareFlowRequest struct {
	*tchttp.BaseRequest
	
	// 资源id，与ResourceType对应
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源类型，1：模板，目前仅支持模板，与ResourceId对应
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 合同流程基础信息
	FlowInfo *BaseFlowInfo `json:"FlowInfo,omitempty" name:"FlowInfo"`

	// 合同签署人信息
	FlowApproverList []*CommonFlowApprover `json:"FlowApproverList,omitempty" name:"FlowApproverList"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 合同流程配置信息
	FlowOption *CreateFlowOption `json:"FlowOption,omitempty" name:"FlowOption"`

	// 该参数不可用，请通过获取 web 可嵌入接口获取合同流程预览 URL
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 该参数不可用，请通过获取 web 可嵌入接口获取合同流程预览 URL
	NeedPreview *bool `json:"NeedPreview,omitempty" name:"NeedPreview"`

	// 企业机构信息，不用传
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`

	// 操作人（用户）信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	delete(f, "FlowApproverList")
	delete(f, "Agent")
	delete(f, "FlowOption")
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
	// 预发起的合同链接
	PrepareFlowUrl *string `json:"PrepareFlowUrl,omitempty" name:"PrepareFlowUrl"`

	// 合同发起后预览链接
	PreviewFlowUrl *string `json:"PreviewFlowUrl,omitempty" name:"PreviewFlowUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ChannelCreateReleaseFlowRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 待解除的流程编号（即原流程的编号）
	NeedRelievedFlowId *string `json:"NeedRelievedFlowId,omitempty" name:"NeedRelievedFlowId"`

	// 解除协议内容
	ReliveInfo *RelieveInfo `json:"ReliveInfo,omitempty" name:"ReliveInfo"`

	// 非必须，解除协议的本企业签署人列表，默认使用原流程的签署人列表；当解除协议的签署人与原流程的签署人不能相同时（例如原流程签署人离职了），需要指定本企业的其他签署人来替换原流程中的原签署人，注意需要指明ApproverNumber来代表需要替换哪一个签署人，解除协议的签署人数量不能多于原流程的签署人数量
	ReleasedApprovers []*ReleasedApprover `json:"ReleasedApprovers,omitempty" name:"ReleasedApprovers"`

	// 签署完回调url，最大长度1000个字符
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelCreateReleaseFlowRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 待解除的流程编号（即原流程的编号）
	NeedRelievedFlowId *string `json:"NeedRelievedFlowId,omitempty" name:"NeedRelievedFlowId"`

	// 解除协议内容
	ReliveInfo *RelieveInfo `json:"ReliveInfo,omitempty" name:"ReliveInfo"`

	// 非必须，解除协议的本企业签署人列表，默认使用原流程的签署人列表；当解除协议的签署人与原流程的签署人不能相同时（例如原流程签署人离职了），需要指定本企业的其他签署人来替换原流程中的原签署人，注意需要指明ApproverNumber来代表需要替换哪一个签署人，解除协议的签署人数量不能多于原流程的签署人数量
	ReleasedApprovers []*ReleasedApprover `json:"ReleasedApprovers,omitempty" name:"ReleasedApprovers"`

	// 签署完回调url，最大长度1000个字符
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateReleaseFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateReleaseFlowResponseParams struct {
	// 解除协议流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ChannelCreateSealPolicyRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 指定印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 指定待授权的用户ID数组
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// 企业机构信息，不用传
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`

	// 操作人（用户）信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelCreateSealPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 指定印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 指定待授权的用户ID数组
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// 企业机构信息，不用传
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`

	// 操作人（用户）信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	delete(f, "Organization")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelCreateSealPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelCreateSealPolicyResponseParams struct {
	// 最终授权成功的用户ID数组。其他的跳过的是已经授权了的
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ChannelDeleteSealPoliciesRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 指定印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 指定用户ID数组
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// 组织机构信息，不用传
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`

	// 操作人（用户）信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelDeleteSealPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 指定印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 指定用户ID数组
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// 组织机构信息，不用传
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`

	// 操作人（用户）信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 返回最大数量，最大为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 查询过滤实名用户，Key为Status，Values为["IsVerified"]
	// 根据第三方系统openId过滤查询员工时,Key为StaffOpenId,Values为["OpenId","OpenId",...]
	// 查询离职员工时，Key为Status，Values为["QuiteJob"]
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0，最大为20000
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelDescribeEmployeesRequest struct {
	*tchttp.BaseRequest
	
	// 返回最大数量，最大为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 查询过滤实名用户，Key为Status，Values为["IsVerified"]
	// 根据第三方系统openId过滤查询员工时,Key为StaffOpenId,Values为["OpenId","OpenId",...]
	// 查询离职员工时，Key为Status，Values为["QuiteJob"]
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0，最大为20000
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	delete(f, "Limit")
	delete(f, "Agent")
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
	Employees []*Staff `json:"Employees,omitempty" name:"Employees"`

	// 偏移量，默认为0，最大为20000
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，最大为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 符合条件的员工数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ChannelDescribeOrganizationSealsRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 返回最大数量，最大为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0，最大为20000
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询信息类型，为1时返回授权用户，为其他值时不返回
	InfoType *int64 `json:"InfoType,omitempty" name:"InfoType"`

	// 印章id（没有输入返回所有）
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 印章类型列表（都是组织机构印章）。
	// 为空时查询所有类型的印章。
	// 目前支持以下类型：
	// OFFICIAL：企业公章；
	// CONTRACT：合同专用章；
	// ORGANIZATION_SEAL：企业印章(图片上传创建)；
	// LEGAL_PERSON_SEAL：法定代表人章
	SealTypes []*string `json:"SealTypes,omitempty" name:"SealTypes"`
}

type ChannelDescribeOrganizationSealsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 返回最大数量，最大为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0，最大为20000
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询信息类型，为1时返回授权用户，为其他值时不返回
	InfoType *int64 `json:"InfoType,omitempty" name:"InfoType"`

	// 印章id（没有输入返回所有）
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 印章类型列表（都是组织机构印章）。
	// 为空时查询所有类型的印章。
	// 目前支持以下类型：
	// OFFICIAL：企业公章；
	// CONTRACT：合同专用章；
	// ORGANIZATION_SEAL：企业印章(图片上传创建)；
	// LEGAL_PERSON_SEAL：法定代表人章
	SealTypes []*string `json:"SealTypes,omitempty" name:"SealTypes"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 查询到的印章结果数组
	Seals []*OccupiedSeal `json:"Seals,omitempty" name:"Seals"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ChannelGetTaskResultApiRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 任务Id，通过ChannelCreateConvertTaskApi接口获得
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`
}

type ChannelGetTaskResultApiRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 任务Id，通过ChannelCreateConvertTaskApi接口获得
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 暂未开放
	Organization *OrganizationInfo `json:"Organization,omitempty" name:"Organization"`
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

	// 预览文件Url，有效期30分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewUrl *string `json:"PreviewUrl,omitempty" name:"PreviewUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ChannelUpdateSealStatusRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 操作的印章状态，DISABLE-停用印章
	Status *string `json:"Status,omitempty" name:"Status"`

	// 印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 更新印章状态原因说明
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type ChannelUpdateSealStatusRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 操作的印章状态，DISABLE-停用印章
	Status *string `json:"Status,omitempty" name:"Status"`

	// 印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 更新印章状态原因说明
	Reason *string `json:"Reason,omitempty" name:"Reason"`
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
	delete(f, "Operator")
	delete(f, "Reason")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChannelUpdateSealStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChannelUpdateSealStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 合同Id，流程Id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type ChannelVerifyPdfRequest struct {
	*tchttp.BaseRequest
	
	// 合同Id，流程Id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	VerifyResult *int64 `json:"VerifyResult,omitempty" name:"VerifyResult"`

	// 验签结果详情,内部状态1-验签成功，在电子签签署；2-验签成功，在其他平台签署；3-验签失败；4-pdf文件没有签名域
	// ；5-文件签名格式错误
	PdfVerifyResults []*PdfVerifyResult `json:"PdfVerifyResults,omitempty" name:"PdfVerifyResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanEditApprover *bool `json:"CanEditApprover,omitempty" name:"CanEditApprover"`
}

type CommonFlowApprover struct {
	// 指定当前签署人为第三方应用集成子客，默认false：当前签署人为第三方应用集成子客，true：当前签署人为saas企业用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotChannelOrganization *bool `json:"NotChannelOrganization,omitempty" name:"NotChannelOrganization"`

	// 签署人类型,目前支持：0-企业签署人，1-个人签署人，3-企业静默签署人
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverType *int64 `json:"ApproverType,omitempty" name:"ApproverType"`

	// 企业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"OrganizationId,omitempty" name:"OrganizationId"`

	// 企业OpenId，第三方应用集成非静默签子客企业签署人发起合同毕传
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationOpenId *string `json:"OrganizationOpenId,omitempty" name:"OrganizationOpenId"`

	// 企业名称，第三方应用集成非静默签子客企业签署人必传，saas企业签署人必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户openId，第三方应用集成非静默签子客企业签署人必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 签署人名称，saas企业签署人，个人签署人必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverName *string `json:"ApproverName,omitempty" name:"ApproverName"`

	// 签署人手机号，saas企业签署人，个人签署人必传
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverMobile *string `json:"ApproverMobile,omitempty" name:"ApproverMobile"`

	// 签署人Id，使用模板发起是，对应模板配置中的签署人RecipientId
	// 注意：模板发起时该字段必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"RecipientId,omitempty" name:"RecipientId"`

	// 签署前置条件：阅读时长限制，不传默认10s,最大300s，最小3s
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreReadTime *int64 `json:"PreReadTime,omitempty" name:"PreReadTime"`

	// 签署前置条件：阅读全文限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFullText *bool `json:"IsFullText,omitempty" name:"IsFullText"`

	// 通知类型：SMS（短信） NONE（不做通知）, 不传 默认SMS
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyType *string `json:"NotifyType,omitempty" name:"NotifyType"`

	// 签署人配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverOption *CommonApproverOption `json:"ApproverOption,omitempty" name:"ApproverOption"`
}

type Component struct {
	// 控件编号
	// 
	// CreateFlowByTemplates发起合同时优先以ComponentId（不为空）填充；否则以ComponentName填充
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
	// TEXT - 普通文本控件，输入文本字符串；
	// MULTI_LINE_TEXT - 多行文本控件，输入文本字符串；
	// CHECK_BOX - 勾选框控件，若选中填写ComponentValue 填写 true或者 false 字符串；
	// FILL_IMAGE - 图片控件，ComponentValue 填写图片的资源 ID；
	// DYNAMIC_TABLE - 动态表格控件；
	// ATTACHMENT - 附件控件,ComponentValue 填写福建图片的资源 ID列表，以逗号分割；
	// SELECTOR - 选择器控件，ComponentValue填写选择的字符串内容；
	// DATE - 日期控件；默认是格式化为xxxx年xx月xx日字符串；
	// DISTRICT - 省市区行政区划控件，ComponentValue填写省市区行政区划字符串内容；
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
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// 控件简称，不能超过30个字符
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// 定义控件是否为必填项，默认为false
	ComponentRequired *bool `json:"ComponentRequired,omitempty" name:"ComponentRequired"`

	// 控件关联的签署方id
	ComponentRecipientId *string `json:"ComponentRecipientId,omitempty" name:"ComponentRecipientId"`

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
	// 
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
	// BORDERLESS_ESIGN – 自动生成无边框腾讯体
	// OCR_ESIGN -- AI智能识别手写签名
	// ESIGN -- 个人印章类型
	// 如：{“ComponentTypeLimit”: [“BORDERLESS_ESIGN”]}
	ComponentExtra *string `json:"ComponentExtra,omitempty" name:"ComponentExtra"`

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

	// 平台企业控件ID。
	// 如果不为空，属于平台企业预设控件；
	ChannelComponentId *string `json:"ChannelComponentId,omitempty" name:"ChannelComponentId"`

	// 指定关键字排序规则，Positive-正序，Reverse-倒序。传入Positive时会根据关键字在PDF文件内的顺序进行排列。在指定KeywordIndexes时，0代表在PDF内查找内容时，查找到的第一个关键字。
	// 传入Reverse时会根据关键字在PDF文件内的反序进行排列。在指定KeywordIndexes时，0代表在PDF内查找内容时，查找到的最后一个关键字。
	KeywordOrder *string `json:"KeywordOrder,omitempty" name:"KeywordOrder"`

	// 指定关键字页码，可选参数，指定页码后，将只在指定的页码内查找关键字，非该页码的关键字将不会查询出来
	KeywordPage *int64 `json:"KeywordPage,omitempty" name:"KeywordPage"`

	// 关键字位置模式，Middle-居中，Below-正下方，Right-正右方，LowerRight-右上角，UpperRight-右下角。示例：如果设置Middle的关键字盖章，则印章的中心会和关键字的中心重合，如果设置Below，则印章在关键字的正下方
	RelativeLocation *string `json:"RelativeLocation,omitempty" name:"RelativeLocation"`

	// 关键字索引，可选参数，如果一个关键字在PDF文件中存在多个，可以通过关键字索引指定使用第几个关键字作为最后的结果，可指定多个索引。示例[0,2]，说明使用PDF文件内第1个和第3个关键字位置。
	KeywordIndexes []*int64 `json:"KeywordIndexes,omitempty" name:"KeywordIndexes"`
}

// Predefined struct for user
type CreateChannelFlowEvidenceReportRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type CreateChannelFlowEvidenceReportRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	// 出证报告 ID，用于查询出证报告接口DescribeChannelFlowEvidenceReport时用到
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportId *string `json:"ReportId,omitempty" name:"ReportId"`

	// 执行中：EvidenceStatusExecuting
	// 成功：EvidenceStatusSuccess
	// 失败：EvidenceStatusFailed
	Status *string `json:"Status,omitempty" name:"Status"`

	// 废除，字段无效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportUrl *string `json:"ReportUrl,omitempty" name:"ReportUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateConsoleLoginUrlRequestParams struct {
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 子客企业名称，最大长度64个字符
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitempty" name:"ProxyOrganizationName"`

	// 子客企业经办人的姓名，最大长度50个字符
	ProxyOperatorName *string `json:"ProxyOperatorName,omitempty" name:"ProxyOperatorName"`

	// PC控制台指定模块，文件/合同管理:"DOCUMENT"，模板管理:"TEMPLATE"，印章管理:"SEAL"，组织架构/人员:"OPERATOR"，空字符串："账号信息"。 EndPoint为"CHANNEL"/"APP"只支持"SEAL"-印章管理
	Module *string `json:"Module,omitempty" name:"Module"`

	// 控制台指定模块Id
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 子客企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`

	// 是否展示左侧菜单栏 是：ENABLE（默认） 否：DISABLE
	MenuStatus *string `json:"MenuStatus,omitempty" name:"MenuStatus"`

	// 链接跳转类型："PC"-PC控制台，“CHANNEL”-H5跳转到电子签小程序；“APP”-第三方APP或小程序跳转电子签小程序，默认为PC控制台
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 触发自动跳转事件，仅对App类型有效，"VERIFIED":企业认证完成/员工认证完成后跳回原App/小程序
	AutoJumpBackEvent *string `json:"AutoJumpBackEvent,omitempty" name:"AutoJumpBackEvent"`

	// 支持的授权方式,授权方式: "1" - 上传授权书认证  "2" - 法定代表人认证
	AuthorizationTypes []*int64 `json:"AuthorizationTypes,omitempty" name:"AuthorizationTypes"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type CreateConsoleLoginUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 子客企业名称，最大长度64个字符
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitempty" name:"ProxyOrganizationName"`

	// 子客企业经办人的姓名，最大长度50个字符
	ProxyOperatorName *string `json:"ProxyOperatorName,omitempty" name:"ProxyOperatorName"`

	// PC控制台指定模块，文件/合同管理:"DOCUMENT"，模板管理:"TEMPLATE"，印章管理:"SEAL"，组织架构/人员:"OPERATOR"，空字符串："账号信息"。 EndPoint为"CHANNEL"/"APP"只支持"SEAL"-印章管理
	Module *string `json:"Module,omitempty" name:"Module"`

	// 控制台指定模块Id
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`

	// 子客企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`

	// 是否展示左侧菜单栏 是：ENABLE（默认） 否：DISABLE
	MenuStatus *string `json:"MenuStatus,omitempty" name:"MenuStatus"`

	// 链接跳转类型："PC"-PC控制台，“CHANNEL”-H5跳转到电子签小程序；“APP”-第三方APP或小程序跳转电子签小程序，默认为PC控制台
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 触发自动跳转事件，仅对App类型有效，"VERIFIED":企业认证完成/员工认证完成后跳回原App/小程序
	AutoJumpBackEvent *string `json:"AutoJumpBackEvent,omitempty" name:"AutoJumpBackEvent"`

	// 支持的授权方式,授权方式: "1" - 上传授权书认证  "2" - 法定代表人认证
	AuthorizationTypes []*int64 `json:"AuthorizationTypes,omitempty" name:"AuthorizationTypes"`

	// 暂未开放
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
	// 2. 若企业认证完成且员工认证完成后，重新获取pc端的链接5分钟之内有效，且只能访问一次
	// 3. 若企业认证完成且员工认证完成后，重新获取H5/APP的链接只要在有效期内（一年）都可以访问
	// 4. 此链接仅单次有效，使用后需要再次创建新的链接（部分聊天软件，如企业微信默认会对链接进行解析，此时需要使用类似“代码片段”的方式或者放到txt文件里发送链接）
	// 5. 创建的链接应避免被转义，如：&被转义为\u0026；如使用Postman请求后，请选择响应类型为 JSON，否则链接将被转义
	ConsoleUrl *string `json:"ConsoleUrl,omitempty" name:"ConsoleUrl"`

	// 子客企业是否已开通腾讯电子签
	IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`

	// 当前经办人是否已认证（false:未认证 true:已认证）
	ProxyOperatorIsVerified *bool `json:"ProxyOperatorIsVerified,omitempty" name:"ProxyOperatorIsVerified"`

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

type CreateFlowOption struct {
	// 是否允许修改合同信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanEditFlow *bool `json:"CanEditFlow,omitempty" name:"CanEditFlow"`
}

// Predefined struct for user
type CreateFlowsByTemplatesRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 多个合同（签署流程）信息，最多支持20个
	FlowInfos []*FlowInfo `json:"FlowInfos,omitempty" name:"FlowInfos"`

	// 是否为预览模式；默认为false，即非预览模式，此时发起合同并返回FlowIds；若为预览模式，不会发起合同，会返回PreviewUrls；
	// 预览链接有效期300秒；
	// 同时，如果预览的文件中指定了动态表格控件，需要进行异步合成；此时此接口返回的是合成前的文档预览链接，而合成完成后的文档预览链接会通过：回调通知的方式、或使用返回的TaskInfo中的TaskId通过ChannelGetTaskResultApi接口查询；
	NeedPreview *bool `json:"NeedPreview,omitempty" name:"NeedPreview"`

	// 预览链接类型 默认:0-文件流, 1- H5链接 注意:此参数在NeedPreview 为true 时有效,
	PreviewType *int64 `json:"PreviewType,omitempty" name:"PreviewType"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type CreateFlowsByTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 多个合同（签署流程）信息，最多支持20个
	FlowInfos []*FlowInfo `json:"FlowInfos,omitempty" name:"FlowInfos"`

	// 是否为预览模式；默认为false，即非预览模式，此时发起合同并返回FlowIds；若为预览模式，不会发起合同，会返回PreviewUrls；
	// 预览链接有效期300秒；
	// 同时，如果预览的文件中指定了动态表格控件，需要进行异步合成；此时此接口返回的是合成前的文档预览链接，而合成完成后的文档预览链接会通过：回调通知的方式、或使用返回的TaskInfo中的TaskId通过ChannelGetTaskResultApi接口查询；
	NeedPreview *bool `json:"NeedPreview,omitempty" name:"NeedPreview"`

	// 预览链接类型 默认:0-文件流, 1- H5链接 注意:此参数在NeedPreview 为true 时有效,
	PreviewType *int64 `json:"PreviewType,omitempty" name:"PreviewType"`

	// 操作者的信息，不用传
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
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 业务信息，限制1024字符
	CustomerData []*string `json:"CustomerData,omitempty" name:"CustomerData"`

	// 创建消息，对应多个合同ID，
	// 成功为“”,创建失败则对应失败消息
	ErrorMessages []*string `json:"ErrorMessages,omitempty" name:"ErrorMessages"`

	// 预览模式下返回的预览文件url数组
	PreviewUrls []*string `json:"PreviewUrls,omitempty" name:"PreviewUrls"`

	// 复杂文档合成任务（如，包含动态表格的预览任务）的任务信息数组；
	// 如果文档需要异步合成，此字段会返回该异步任务的任务信息，后续可以通过ChannelGetTaskResultApi接口查询任务详情；
	TaskInfos []*TaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos"`

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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 印章名称，最大长度不超过50字符
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 印章图片base64，大小不超过10M（原始图片不超过7.6M）
	SealImage *string `json:"SealImage,omitempty" name:"SealImage"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type CreateSealByImageRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 印章名称，最大长度不超过50字符
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 印章图片base64，大小不超过10M（原始图片不超过7.6M）
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程编号数组，最多支持100个。(备注：该参数和合同组编号必须二选一)
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 合同组编号(备注：该参数和合同(流程)编号数组必须二选一)
	FlowGroupId *string `json:"FlowGroupId,omitempty" name:"FlowGroupId"`

	// 签署链接类型：“WEIXINAPP”-短链直接跳小程序；“CHANNEL”-跳转H5页面；“APP”-第三方APP或小程序跳转电子签小程序；"LONGURL2WEIXINAPP"-长链接跳转小程序；默认“WEIXINAPP”类型，即跳转至小程序；
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 签署链接生成类型，默认是 "ALL"；
	// "ALL"：全部签署方签署链接，此时不会给自动签署的签署方创建签署链接；
	// "CHANNEL"：第三方平台子客企业企业；
	// "NOT_CHANNEL"：非第三方平台子客企业企业；
	// "PERSON"：个人；
	// "FOLLOWER"：关注方，目前是合同抄送方；
	GenerateType *string `json:"GenerateType,omitempty" name:"GenerateType"`

	// 非第三方平台子客企业参与方的企业名称，GenerateType为"NOT_CHANNEL"时必填
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 参与人姓名，GenerateType为"PERSON"时必填
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参与人手机号；
	// GenerateType为"PERSON"或"FOLLOWER"时必填
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 第三方平台子客企业的企业OpenId，GenerateType为"CHANNEL"时必填
	OrganizationOpenId *string `json:"OrganizationOpenId,omitempty" name:"OrganizationOpenId"`

	// 第三方平台子客企业参与人OpenId，GenerateType为"CHANNEL"时可用，指定到具体参与人, 仅展示已经实名的经办人信息
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// Endpoint为"APP" 类型的签署链接，可以设置此值；支持调用方小程序打开签署链接，在电子签小程序完成签署后自动回跳至调用方小程序
	AutoJumpBack *bool `json:"AutoJumpBack,omitempty" name:"AutoJumpBack"`

	// 签署完之后的H5页面的跳转链接，针对Endpoint为CHANNEL时有效，最大长度1000个字符。
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type CreateSignUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 签署流程编号数组，最多支持100个。(备注：该参数和合同组编号必须二选一)
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 合同组编号(备注：该参数和合同(流程)编号数组必须二选一)
	FlowGroupId *string `json:"FlowGroupId,omitempty" name:"FlowGroupId"`

	// 签署链接类型：“WEIXINAPP”-短链直接跳小程序；“CHANNEL”-跳转H5页面；“APP”-第三方APP或小程序跳转电子签小程序；"LONGURL2WEIXINAPP"-长链接跳转小程序；默认“WEIXINAPP”类型，即跳转至小程序；
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 签署链接生成类型，默认是 "ALL"；
	// "ALL"：全部签署方签署链接，此时不会给自动签署的签署方创建签署链接；
	// "CHANNEL"：第三方平台子客企业企业；
	// "NOT_CHANNEL"：非第三方平台子客企业企业；
	// "PERSON"：个人；
	// "FOLLOWER"：关注方，目前是合同抄送方；
	GenerateType *string `json:"GenerateType,omitempty" name:"GenerateType"`

	// 非第三方平台子客企业参与方的企业名称，GenerateType为"NOT_CHANNEL"时必填
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 参与人姓名，GenerateType为"PERSON"时必填
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参与人手机号；
	// GenerateType为"PERSON"或"FOLLOWER"时必填
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 第三方平台子客企业的企业OpenId，GenerateType为"CHANNEL"时必填
	OrganizationOpenId *string `json:"OrganizationOpenId,omitempty" name:"OrganizationOpenId"`

	// 第三方平台子客企业参与人OpenId，GenerateType为"CHANNEL"时可用，指定到具体参与人, 仅展示已经实名的经办人信息
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// Endpoint为"APP" 类型的签署链接，可以设置此值；支持调用方小程序打开签署链接，在电子签小程序完成签署后自动回跳至调用方小程序
	AutoJumpBack *bool `json:"AutoJumpBack,omitempty" name:"AutoJumpBack"`

	// 签署完之后的H5页面的跳转链接，针对Endpoint为CHANNEL时有效，最大长度1000个字符。
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 暂未开放
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

type Department struct {
	// 部门id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`

	// 部门名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentName *string `json:"DepartmentName,omitempty" name:"DepartmentName"`
}

// Predefined struct for user
type DescribeChannelFlowEvidenceReportRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 出证报告编号
	ReportId *string `json:"ReportId,omitempty" name:"ReportId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type DescribeChannelFlowEvidenceReportRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 出证报告编号
	ReportId *string `json:"ReportId,omitempty" name:"ReportId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
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
	// 出证报告 URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportUrl *string `json:"ReportUrl,omitempty" name:"ReportUrl"`

	// 执行中：EvidenceStatusExecuting
	// 成功：EvidenceStatusSuccess
	// 失败：EvidenceStatusFailed
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填
	// 
	// 注: 此接口 参数Agent. ProxyOperator.OpenId 需要传递超管或者法人的OpenId
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
}

type DescribeExtendedServiceAuthInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填
	// 
	// 注: 此接口 参数Agent. ProxyOperator.OpenId 需要传递超管或者法人的OpenId
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`
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
	AuthInfo []*ExtentServiceAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 合同(流程)编号数组，最多支持100个。
	// （备注：该参数和合同组编号必须二选一）
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 合同组编号（备注：该参数和合同(流程)编号数组必须二选一）
	FlowGroupId *string `json:"FlowGroupId,omitempty" name:"FlowGroupId"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type DescribeFlowDetailInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 合同(流程)编号数组，最多支持100个。
	// （备注：该参数和合同组编号必须二选一）
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 合同组编号（备注：该参数和合同(流程)编号数组必须二选一）
	FlowGroupId *string `json:"FlowGroupId,omitempty" name:"FlowGroupId"`

	// 暂未开放
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
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 第三方平台子客企业OpenId
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitempty" name:"ProxyOrganizationOpenId"`

	// 合同(签署流程)的具体详细描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowInfo []*FlowDetailInfo `json:"FlowInfo,omitempty" name:"FlowInfo"`

	// 合同组编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowGroupId *string `json:"FlowGroupId,omitempty" name:"FlowGroupId"`

	// 合同组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowGroupName *string `json:"FlowGroupName,omitempty" name:"FlowGroupName"`

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
	// 应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 查询资源所对应的签署流程Id，最多支持50个
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type DescribeResourceUrlsByFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 查询资源所对应的签署流程Id，最多支持50个
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 操作者的信息，不用传
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
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 模板唯一标识，查询单个模板时使用
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 查询内容：0-模板列表及详情（默认），1-仅模板列表
	ContentType *int64 `json:"ContentType,omitempty" name:"ContentType"`

	// 查询个数，默认20，最大100；在查询列表的时候有效
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移位置，默认0；在查询列表的时候有效
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否返回所有组件信息。默认false，只返回发起方控件；true，返回所有签署方控件
	QueryAllComponents *bool `json:"QueryAllComponents,omitempty" name:"QueryAllComponents"`

	// 模糊搜索模板名称，最大长度200
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 是否获取模板预览链接
	WithPreviewUrl *bool `json:"WithPreviewUrl,omitempty" name:"WithPreviewUrl"`

	// 是否获取模板的PDF文件链接- 第三方应用集成需要开启白名单时才能使用。
	WithPdfUrl *bool `json:"WithPdfUrl,omitempty" name:"WithPdfUrl"`

	// 模板ID
	ChannelTemplateId *string `json:"ChannelTemplateId,omitempty" name:"ChannelTemplateId"`
}

type DescribeTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 模板唯一标识，查询单个模板时使用
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 查询内容：0-模板列表及详情（默认），1-仅模板列表
	ContentType *int64 `json:"ContentType,omitempty" name:"ContentType"`

	// 查询个数，默认20，最大100；在查询列表的时候有效
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移位置，默认0；在查询列表的时候有效
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否返回所有组件信息。默认false，只返回发起方控件；true，返回所有签署方控件
	QueryAllComponents *bool `json:"QueryAllComponents,omitempty" name:"QueryAllComponents"`

	// 模糊搜索模板名称，最大长度200
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 是否获取模板预览链接
	WithPreviewUrl *bool `json:"WithPreviewUrl,omitempty" name:"WithPreviewUrl"`

	// 是否获取模板的PDF文件链接- 第三方应用集成需要开启白名单时才能使用。
	WithPdfUrl *bool `json:"WithPdfUrl,omitempty" name:"WithPdfUrl"`

	// 模板ID
	ChannelTemplateId *string `json:"ChannelTemplateId,omitempty" name:"ChannelTemplateId"`
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
	delete(f, "Operator")
	delete(f, "WithPreviewUrl")
	delete(f, "WithPdfUrl")
	delete(f, "ChannelTemplateId")
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
	// 应用信息，此接口Agent.AppId必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 开始时间，例如：2021-03-21
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束时间，例如：2021-06-21；
	// 开始时间到结束时间的区间长度小于等于90天。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 是否汇总数据，默认不汇总。
	// 不汇总：返回在统计区间内第三方平台下所有企业的每日明细，即每个企业N条数据，N为统计天数；
	// 汇总：返回在统计区间内第三方平台下所有企业的汇总后数据，即每个企业一条数据；
	NeedAggregate *bool `json:"NeedAggregate,omitempty" name:"NeedAggregate"`

	// 单次返回的最多条目数量。默认为1000，且不能超过1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认是0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type DescribeUsageRequest struct {
	*tchttp.BaseRequest
	
	// 应用信息，此接口Agent.AppId必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 开始时间，例如：2021-03-21
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束时间，例如：2021-06-21；
	// 开始时间到结束时间的区间长度小于等于90天。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 是否汇总数据，默认不汇总。
	// 不汇总：返回在统计区间内第三方平台下所有企业的每日明细，即每个企业N条数据，N为统计天数；
	// 汇总：返回在统计区间内第三方平台下所有企业的汇总后数据，即每个企业一条数据；
	NeedAggregate *bool `json:"NeedAggregate,omitempty" name:"NeedAggregate"`

	// 单次返回的最多条目数量。默认为1000，且不能超过1000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认是0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 暂未开放
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

type ExtentServiceAuthInfo struct {
	// 扩展服务类型
	//   AUTO_SIGN             企业静默签（自动签署）
	//   OVERSEA_SIGN          企业与港澳台居民*签署合同
	//   MOBILE_CHECK_APPROVER 使用手机号验证签署方身份
	//   PAGING_SEAL           骑缝章
	//   DOWNLOAD_FLOW         授权渠道下载合同 
	Type *string `json:"Type,omitempty" name:"Type"`

	// 扩展服务名称 
	Name *string `json:"Name,omitempty" name:"Name"`

	// 服务状态 
	// ENABLE 开启 
	// DISABLE 关闭
	Status *string `json:"Status,omitempty" name:"Status"`

	// 最近操作人openid（经办人openid）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorOpenId *string `json:"OperatorOpenId,omitempty" name:"OperatorOpenId"`

	// 最近操作时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateOn *int64 `json:"OperateOn,omitempty" name:"OperateOn"`
}

type Filter struct {
	// 查询过滤条件的Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 查询过滤条件的Value列表
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FlowApproverDetail struct {
	// 模板配置时候的签署人id,与控件绑定
	ReceiptId *string `json:"ReceiptId,omitempty" name:"ReceiptId"`

	// 平台企业的第三方id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitempty" name:"ProxyOrganizationOpenId"`

	// 平台企业操作人的第三方id
	ProxyOperatorOpenId *string `json:"ProxyOperatorOpenId,omitempty" name:"ProxyOperatorOpenId"`

	// 平台企业名称
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

	// 签署人身份证件类型
	// 1.ID_CARD 居民身份证
	// 2.HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证
	// 3.HONGKONG_AND_MACAO 港澳居民来往内地通行证
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 签署人证件号
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 签署人手机号，脱敏显示。大陆手机号为11位，暂不支持海外手机号。
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 企业签署方工商营业执照上的企业名称，签署方为非发起方企业场景下必传，最大长度64个字符；
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 指定签署人非第三方平台子客企业下员工，在ApproverType为ORGANIZATION时指定。
	// 默认为false，即签署人位于同一个第三方平台应用号下；默认为false，即签署人位于同一个第三方应用号下；
	NotChannelOrganization *bool `json:"NotChannelOrganization,omitempty" name:"NotChannelOrganization"`

	// 用户侧第三方id，最大长度64个字符
	// 当签署方为同一第三方平台下的员工时，该字段若不指定，则发起【待领取】的流程
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 企业签署方在同一第三方平台应用下的其他合作企业OpenId，签署方为非发起方企业场景下必传，最大长度64个字符；
	OrganizationOpenId *string `json:"OrganizationOpenId,omitempty" name:"OrganizationOpenId"`

	// 签署人类型
	// PERSON-个人/自然人；
	// PERSON_AUTO_SIGN-个人自动签（定制化场景下使用）；
	// ORGANIZATION-企业（企业签署方或模板发起时的企业静默签）；
	// ENTERPRISESERVER-企业静默签（文件发起时的企业静默签字）。
	ApproverType *string `json:"ApproverType,omitempty" name:"ApproverType"`

	// 签署流程签署人在模板中对应的签署人Id；在非单方签署、以及非B2C签署的场景下必传，用于指定当前签署方在签署流程中的位置；
	RecipientId *string `json:"RecipientId,omitempty" name:"RecipientId"`

	// 签署截止时间，默认一年
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 签署完回调url，最大长度1000个字符
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 使用PDF文件直接发起合同时，签署人指定的签署控件
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 个人签署方指定签署控件类型，目前支持：OCR_ESIGN -AI智慧手写签名
	// HANDWRITE -手写签名
	ComponentLimitType []*string `json:"ComponentLimitType,omitempty" name:"ComponentLimitType"`

	// 合同的强制预览时间：3~300s，未指定则按合同页数计算
	PreReadTime *int64 `json:"PreReadTime,omitempty" name:"PreReadTime"`

	// 签署完前端跳转的url，暂未使用
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 签署人个性化能力值
	ApproverOption *ApproverOption `json:"ApproverOption,omitempty" name:"ApproverOption"`

	// 当前签署方进行签署操作是否需要企业内部审批，true 则为需要
	ApproverNeedSignReview *bool `json:"ApproverNeedSignReview,omitempty" name:"ApproverNeedSignReview"`
}

type FlowApproverUrlInfo struct {
	// 签署链接，注意该链接有效期为30分钟，同时需要注意保密，不要外泄给无关用户。
	SignUrl *string `json:"SignUrl,omitempty" name:"SignUrl"`

	// 签署人手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 签署人姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 签署人类型 PERSON-个人
	ApproverType *string `json:"ApproverType,omitempty" name:"ApproverType"`
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

type FlowFileInfo struct {
	// 签署文件资源Id列表，目前仅支持单个文件
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 签署流程名称，长度不超过200个字符
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署流程签约方列表，最多不超过5个参与方
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitempty" name:"FlowApprovers"`

	// 签署流程截止时间，十位数时间戳，最大值为33162419560，即3020年
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 签署流程的描述，长度不超过1000个字符
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 签署流程的类型，长度不超过255个字符
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 签署流程回调地址，长度不超过255个字符
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 第三方应用的业务信息，最大长度1000个字符。发起自动签署时，需设置对应自动签署场景，目前仅支持场景：处方单-E_PRESCRIPTION_AUTO_SIGN
	CustomerData *string `json:"CustomerData,omitempty" name:"CustomerData"`

	// 合同签署顺序类型(无序签,顺序签)，默认为false，即有序签署
	Unordered *bool `json:"Unordered,omitempty" name:"Unordered"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitempty" name:"CustomShowMap"`

	// 本企业(发起方企业)是否需要签署审批
	NeedSignReview *bool `json:"NeedSignReview,omitempty" name:"NeedSignReview"`
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

	//  第三方应用平台的业务信息，最大长度1000个字符。发起自动签署时，需设置对应自动签署场景，目前仅支持场景：处方单-E_PRESCRIPTION_AUTO_SIGN
	CustomerData *string `json:"CustomerData,omitempty" name:"CustomerData"`

	// 合同显示的页卡模板，说明：只支持{合同名称}, {发起方企业}, {发起方姓名}, {签署方N企业}, {签署方N姓名}，且N不能超过签署人的数量，N从1开始
	CustomShowMap *string `json:"CustomShowMap,omitempty" name:"CustomShowMap"`

	// 被抄送人的信息列表，抄送功能暂不开放
	CcInfos []*CcInfo `json:"CcInfos,omitempty" name:"CcInfos"`

	// 发起方企业的签署人进行签署操作是否需要企业内部审批。
	// 若设置为true,审核结果需通过接口 ChannelCreateFlowSignReview 通知电子签，审核通过后，发起方企业签署人方可进行签署操作，否则会阻塞其签署操作。
	// 
	// 注：企业可以通过此功能与企业内部的审批流程进行关联，支持手动、静默签署合同。
	NeedSignReview *bool `json:"NeedSignReview,omitempty" name:"NeedSignReview"`

	// 给关注人发送短信通知的类型，0-合同发起时通知 1-签署完成后通知
	CcNotifyType *int64 `json:"CcNotifyType,omitempty" name:"CcNotifyType"`
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
	// 控件填充vaule，ComponentType和传入值类型对应关系：
	// TEXT - 文本内容
	// MULTI_LINE_TEXT - 文本内容
	// CHECK_BOX - true/false
	// FILL_IMAGE、ATTACHMENT - 附件的FileId，需要通过UploadFiles接口上传获取
	// SELECTOR - 选项值
	// DYNAMIC_TABLE - 传入json格式的表格内容，具体见数据结构FlowInfo：https://cloud.tencent.com/document/api/1420/61525#FlowInfo
	ComponentValue *string `json:"ComponentValue,omitempty" name:"ComponentValue"`

	// 表单域或控件的ID，跟ComponentName二选一，不能全为空；
	// CreateFlowsByTemplates 接口不使用此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`

	// 控件的名字，跟ComponentId二选一，不能全为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
}

// Predefined struct for user
type GetDownloadFlowUrlRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 文件夹数组，签署流程总数不能超过50个，一个文件夹下，不能超过20个签署流程
	DownLoadFlows []*DownloadFlowInfo `json:"DownLoadFlows,omitempty" name:"DownLoadFlows"`

	// 操作者的信息，不用传
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type GetDownloadFlowUrlRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 文件夹数组，签署流程总数不能超过50个，一个文件夹下，不能超过20个签署流程
	DownLoadFlows []*DownloadFlowInfo `json:"DownLoadFlows,omitempty" name:"DownLoadFlows"`

	// 操作者的信息，不用传
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
type ModifyExtendedServiceRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	// 
	// 注: 此接口 参数Agent. ProxyOperator.OpenId 需要传递超管或者法人的OpenId
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	//   扩展服务类型
	//   AUTO_SIGN             企业静默签（自动签署）
	//   OVERSEA_SIGN          企业与港澳台居民*签署合同
	//   MOBILE_CHECK_APPROVER 使用手机号验证签署方身份
	//   PAGING_SEAL           骑缝章
	//   DOWNLOAD_FLOW         授权渠道下载合同 
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 操作类型 
	// OPEN:开通 
	// CLOSE:关闭
	Operate *string `json:"Operate,omitempty" name:"Operate"`
}

type ModifyExtendedServiceRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	// 
	// 注: 此接口 参数Agent. ProxyOperator.OpenId 需要传递超管或者法人的OpenId
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	//   扩展服务类型
	//   AUTO_SIGN             企业静默签（自动签署）
	//   OVERSEA_SIGN          企业与港澳台居民*签署合同
	//   MOBILE_CHECK_APPROVER 使用手机号验证签署方身份
	//   PAGING_SEAL           骑缝章
	//   DOWNLOAD_FLOW         授权渠道下载合同 
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 操作类型 
	// OPEN:开通 
	// CLOSE:关闭
	Operate *string `json:"Operate,omitempty" name:"Operate"`
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
	OperateUrl *string `json:"OperateUrl,omitempty" name:"OperateUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 电子印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 电子印章授权时间戳
	CreateOn *int64 `json:"CreateOn,omitempty" name:"CreateOn"`

	// 电子印章授权人
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 电子印章策略Id
	SealPolicyId *string `json:"SealPolicyId,omitempty" name:"SealPolicyId"`

	// 印章状态，有以下六种：CHECKING（审核中）SUCCESS（已启用）FAIL（审核拒绝）CHECKING-SADM（待超管审核）DISABLE（已停用）STOPPED（已终止）
	SealStatus *string `json:"SealStatus,omitempty" name:"SealStatus"`

	// 审核失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`

	// 印章图片url，5分钟内有效
	Url *string `json:"Url,omitempty" name:"Url"`

	// 印章类型
	SealType *string `json:"SealType,omitempty" name:"SealType"`

	// 用印申请是否为永久授权
	IsAllTime *bool `json:"IsAllTime,omitempty" name:"IsAllTime"`

	// 授权人列表
	AuthorizedUsers []*AuthorizedUser `json:"AuthorizedUsers,omitempty" name:"AuthorizedUsers"`
}

// Predefined struct for user
type OperateChannelTemplateRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 操作类型，查询:"SELECT"，删除:"DELETE"，更新:"UPDATE"
	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`

	// 第三方应用平台模板库模板唯一标识
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 合作企业方第三方机构唯一标识数据，支持多个， 用","进行分隔
	ProxyOrganizationOpenIds *string `json:"ProxyOrganizationOpenIds,omitempty" name:"ProxyOrganizationOpenIds"`

	// 模板可见性, 全部可见-"all", 部分可见-"part"
	AuthTag *string `json:"AuthTag,omitempty" name:"AuthTag"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type OperateChannelTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 操作类型，查询:"SELECT"，删除:"DELETE"，更新:"UPDATE"
	OperateType *string `json:"OperateType,omitempty" name:"OperateType"`

	// 第三方应用平台模板库模板唯一标识
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 合作企业方第三方机构唯一标识数据，支持多个， 用","进行分隔
	ProxyOrganizationOpenIds *string `json:"ProxyOrganizationOpenIds,omitempty" name:"ProxyOrganizationOpenIds"`

	// 模板可见性, 全部可见-"all", 部分可见-"part"
	AuthTag *string `json:"AuthTag,omitempty" name:"AuthTag"`

	// 暂未开放
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
	// 腾讯电子签颁发给第三方应用平台的应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 第三方应用平台模板库模板唯一标识
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

type OrganizationInfo struct {
	// 用户在渠道的机构编号
	OrganizationOpenId *string `json:"OrganizationOpenId,omitempty" name:"OrganizationOpenId"`

	// 用户真实的IP
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 机构的代理IP
	ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`

	// 机构在平台的编号
	OrganizationId *string `json:"OrganizationId,omitempty" name:"OrganizationId"`

	// 用户渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`
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

	// 签名类型
	SignType *int64 `json:"SignType,omitempty" name:"SignType"`

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

// Predefined struct for user
type PrepareFlowsRequestParams struct {
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 多个合同（签署流程）信息，最大支持20个签署流程。
	FlowInfos []*FlowInfo `json:"FlowInfos,omitempty" name:"FlowInfos"`

	// 操作完成后的跳转地址，最大长度200
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type PrepareFlowsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 多个合同（签署流程）信息，最大支持20个签署流程。
	FlowInfos []*FlowInfo `json:"FlowInfos,omitempty" name:"FlowInfos"`

	// 操作完成后的跳转地址，最大长度200
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 暂未开放
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
	// 对应Agent-ProxyOperator-OpenId。第三方应用平台自定义，对子客企业员的唯一标识。一个OpenId在一个子客企业内唯一对应一个真实员工，不可在其他子客企业内重复使用。（例如，可以使用经办人企业名+员工身份证的hash值，需要第三方应用平台保存），最大64位字符串
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

	// 默认角色，值为以下三个对应的英文：
	// 业务管理员：admin
	// 经办人：channel-normal-operator
	// 业务员：channel-sales-man
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultRole *string `json:"DefaultRole,omitempty" name:"DefaultRole"`
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

type ReleasedApprover struct {
	// 企业签署方工商营业执照上的企业名称，签署方为非发起方企业场景下必传，最大长度64个字符
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 签署人在原流程中的签署人列表中的顺序序号（从0开始，按顺序依次递增），如果不清楚原流程中的签署人列表，可以通过DescribeFlows接口查看
	ApproverNumber *uint64 `json:"ApproverNumber,omitempty" name:"ApproverNumber"`

	// 签署人类型，目前仅支持
	// ORGANIZATION-企业
	ApproverType *string `json:"ApproverType,omitempty" name:"ApproverType"`

	// 签署人姓名，最大长度50个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 签署人身份证件类型
	// 1.ID_CARD 居民身份证
	// 2.HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证
	// 3.HONGKONG_AND_MACAO 港澳居民来往内地通行证
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 签署人证件号
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 签署人手机号，脱敏显示。大陆手机号为11位，暂不支持海外手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 企业签署方在同一第三方应用下的其他合作企业OpenId，签署方为非发起方企业场景下必传，最大长度64个字符
	OrganizationOpenId *string `json:"OrganizationOpenId,omitempty" name:"OrganizationOpenId"`

	// 用户侧第三方id，最大长度64个字符
	// 当签署方为同一第三方应用下的员工时，该字必传
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
}

type RelieveInfo struct {
	// 解除理由，最大支持200个字
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 解除后仍然有效的条款，保留条款，最大支持200个字
	RemainInForceItem *string `json:"RemainInForceItem,omitempty" name:"RemainInForceItem"`

	// 原合同事项处理-费用结算，最大支持200个字
	OriginalExpenseSettlement *string `json:"OriginalExpenseSettlement,omitempty" name:"OriginalExpenseSettlement"`

	// 原合同事项处理-其他事项，最大支持200个字
	OriginalOtherSettlement *string `json:"OriginalOtherSettlement,omitempty" name:"OriginalOtherSettlement"`

	// 其他约定，最大支持200个字
	OtherDeals *string `json:"OtherDeals,omitempty" name:"OtherDeals"`
}

type RemindFlowRecords struct {
	// 是否能够催办
	CanRemind *bool `json:"CanRemind,omitempty" name:"CanRemind"`

	// 合同id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 催办详情
	RemindMessage *string `json:"RemindMessage,omitempty" name:"RemindMessage"`
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

type SignUrl struct {
	// 小程序签署链接
	AppSignUrl *string `json:"AppSignUrl,omitempty" name:"AppSignUrl"`

	// 签署链接有效时间
	EffectiveTime *string `json:"EffectiveTime,omitempty" name:"EffectiveTime"`

	// 移动端签署链接
	HttpSignUrl *string `json:"HttpSignUrl,omitempty" name:"HttpSignUrl"`
}

type SignUrlInfo struct {
	// 签署链接，过期时间为30天
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignUrl *string `json:"SignUrl,omitempty" name:"SignUrl"`

	// 合同过期时间
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

	// 合同组签署链接对应的合同组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowGroupId *string `json:"FlowGroupId,omitempty" name:"FlowGroupId"`
}

type Staff struct {
	// 员工在电子签平台的id
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 显示的员工名
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 员工手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 员工邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitempty" name:"Email"`

	// 员工在第三方平台id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 员工角色
	// 注意：此字段可能返回 null，表示取不到有效值。
	Roles []*StaffRole `json:"Roles,omitempty" name:"Roles"`

	// 员工部门
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitempty" name:"Department"`

	// 员工是否实名
	Verified *bool `json:"Verified,omitempty" name:"Verified"`

	// 员工创建时间戳
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 员工实名时间戳
	VerifiedOn *int64 `json:"VerifiedOn,omitempty" name:"VerifiedOn"`

	// 员工是否离职：0-未离职，1-离职
	QuiteJob *int64 `json:"QuiteJob,omitempty" name:"QuiteJob"`
}

type StaffRole struct {
	// 角色id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
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
	// 应用相关信息。 此接口Agent.AppId 和 Agent.ProxyOrganizationOpenId必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 操作类型，新增:"CREATE"，修改:"UPDATE"，离职:"RESIGN"
	OperatorType *string `json:"OperatorType,omitempty" name:"OperatorType"`

	// 经办人信息列表，最大长度200
	ProxyOrganizationOperators []*ProxyOrganizationOperator `json:"ProxyOrganizationOperators,omitempty" name:"ProxyOrganizationOperators"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type SyncProxyOrganizationOperatorsRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息。 此接口Agent.AppId 和 Agent.ProxyOrganizationOpenId必填。
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 操作类型，新增:"CREATE"，修改:"UPDATE"，离职:"RESIGN"
	OperatorType *string `json:"OperatorType,omitempty" name:"OperatorType"`

	// 经办人信息列表，最大长度200
	ProxyOrganizationOperators []*ProxyOrganizationOperator `json:"ProxyOrganizationOperators,omitempty" name:"ProxyOrganizationOperators"`

	// 暂未开放
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

	// 第三方平台子客企业名称，最大长度64个字符
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitempty" name:"ProxyOrganizationName"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitempty" name:"BusinessLicense"`

	// 第三方平台子客企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`

	// 第三方平台子客企业法人/负责人姓名
	ProxyLegalName *string `json:"ProxyLegalName,omitempty" name:"ProxyLegalName"`

	// 暂未开放
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type SyncProxyOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 第三方平台子客企业名称，最大长度64个字符
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitempty" name:"ProxyOrganizationName"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitempty" name:"BusinessLicense"`

	// 第三方平台子客企业统一社会信用代码，最大长度200个字符
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`

	// 第三方平台子客企业法人/负责人姓名
	ProxyLegalName *string `json:"ProxyLegalName,omitempty" name:"ProxyLegalName"`

	// 暂未开放
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
	delete(f, "ProxyLegalName")
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

type TaskInfo struct {
	// 合成任务Id，可以通过 ChannelGetTaskResultApi 接口获取任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态：READY - 任务已完成；NOTREADY - 任务未完成；
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
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

	// 模板中的流程参与人信息
	Recipients []*Recipient `json:"Recipients,omitempty" name:"Recipients"`

	// 签署区模板信息结构
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 模板类型：1-静默签；3-普通模板
	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`

	// 是否是发起人 ,已弃用
	IsPromoter *bool `json:"IsPromoter,omitempty" name:"IsPromoter"`

	// 模板的创建者信息
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 模板创建的时间戳（精确到秒）
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 模板的H5预览链接,可以通过浏览器打开此链接预览模板，或者嵌入到iframe中预览模板。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewUrl *string `json:"PreviewUrl,omitempty" name:"PreviewUrl"`

	// 第三方应用集成-模板PDF文件链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	PdfUrl *string `json:"PdfUrl,omitempty" name:"PdfUrl"`

	// 关联的平台企业模板ID
	ChannelTemplateId *string `json:"ChannelTemplateId,omitempty" name:"ChannelTemplateId"`

	// 关联的平台企业模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelTemplateName *string `json:"ChannelTemplateName,omitempty" name:"ChannelTemplateName"`

	// 0-需要子客企业手动领取平台企业的模板(默认); 1-平台自动设置子客模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelAutoSave *int64 `json:"ChannelAutoSave,omitempty" name:"ChannelAutoSave"`

	// 模板版本，全数字字符。默认为空，初始版本为yyyyMMdd001。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateVersion *string `json:"TemplateVersion,omitempty" name:"TemplateVersion"`
}

type UploadFile struct {
	// Base64编码后的文件内容
	FileBody *string `json:"FileBody,omitempty" name:"FileBody"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

// Predefined struct for user
type UploadFilesRequestParams struct {
	// 应用相关信息，若是第三方应用集成调用 appid 和proxyappid 必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 文件对应业务类型
	// 1. TEMPLATE - 模板； 文件类型：.pdf/.doc/.docx/.html
	// 2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 上传文件内容数组，最多支持20个文件
	FileInfos []*UploadFile `json:"FileInfos,omitempty" name:"FileInfos"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`
}

type UploadFilesRequest struct {
	*tchttp.BaseRequest
	
	// 应用相关信息，若是第三方应用集成调用 appid 和proxyappid 必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 文件对应业务类型
	// 1. TEMPLATE - 模板； 文件类型：.pdf/.doc/.docx/.html
	// 2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html
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
	// 文件id数组，有效期一个小时；有效期内此文件id可以反复使用
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
	// 子客企业唯一标识
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitempty" name:"ProxyOrganizationOpenId"`

	// 子客企业名
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

	// 消耗渠道
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowChannel *string `json:"FlowChannel,omitempty" name:"FlowChannel"`
}

type UserInfo struct {
	// 第三方应用平台自定义，对应第三方平台子客企业员的唯一标识。一个OpenId在一个子客企业内唯一对应一个真实员工，不可在其他子客企业内重复使用。（例如，可以使用经办人企业名+员工身份证的hash值，需要第三方应用平台保存），最大64位字符串
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 内部参数，暂未开放使用
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 内部参数，暂未开放使用
	CustomUserId *string `json:"CustomUserId,omitempty" name:"CustomUserId"`

	// 内部参数，暂未开放使用
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 内部参数，暂未开放使用
	ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
}