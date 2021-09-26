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

	// 腾讯电子签颁发给渠道侧合作企业的企业ID
	ProxyOrganizationId *string `json:"ProxyOrganizationId,omitempty" name:"ProxyOrganizationId"`

	// 腾讯电子签颁发给渠道侧合作企业的应用ID
	ProxyAppId *string `json:"ProxyAppId,omitempty" name:"ProxyAppId"`

	// 渠道/平台合作企业经办人（操作员）
	ProxyOperator *UserInfo `json:"ProxyOperator,omitempty" name:"ProxyOperator"`

	// 渠道/平台合作企业的企业ID
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitempty" name:"ProxyOrganizationOpenId"`
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

	// 控件简称
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// 定义控件是否为必填项，默认为false
	ComponentRequired *bool `json:"ComponentRequired,omitempty" name:"ComponentRequired"`

	// 控件所属文件的序号 (文档中文件的排列序号)
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

	// 参数控件所在页码
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

	// 控件描述
	ComponentDescription *string `json:"ComponentDescription,omitempty" name:"ComponentDescription"`
}

type CreateConsoleLoginUrlRequest struct {
	*tchttp.BaseRequest

	// 应用信息
	// 此接口Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 渠道侧合作企业名称
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitempty" name:"ProxyOrganizationName"`

	// 渠道侧合作企业统一社会信用代码
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`

	// 渠道侧合作企业经办人的姓名
	ProxyOperatorName *string `json:"ProxyOperatorName,omitempty" name:"ProxyOperatorName"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 控制台指定模块，文件/合同管理:"DOCUMENT"，模版管理:"TEMPLATE"，印章管理:"SEAL"，组织架构/人员:"OPERATOR"，空字符串："账号信息"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 控制台指定模块Id
	ModuleId *string `json:"ModuleId,omitempty" name:"ModuleId"`
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
	delete(f, "Operator")
	delete(f, "Module")
	delete(f, "ModuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsoleLoginUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateConsoleLoginUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 控制台url
		ConsoleUrl *string `json:"ConsoleUrl,omitempty" name:"ConsoleUrl"`

		// 渠道合作企业是否认证开通腾讯电子签。
	// 当渠道合作企业未完成认证开通腾讯电子签,建议先调用同步企业信息(SyncProxyOrganization)和同步经办人信息(SyncProxyOrganizationOperators)接口成功后再跳转到登录页面。
		IsActivated *bool `json:"IsActivated,omitempty" name:"IsActivated"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateFlowsByTemplatesRequest struct {
	*tchttp.BaseRequest

	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 多个合同（流程）信息
	FlowInfos []*FlowInfo `json:"FlowInfos,omitempty" name:"FlowInfos"`

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
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowsByTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlowsByTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 多个合同ID
		FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

		// 渠道的业务信息，限制1024字符
		CustomerData []*string `json:"CustomerData,omitempty" name:"CustomerData"`

		// 创建消息，对应多个合同ID，
	// 成功为“”,创建失败则对应失败消息
		ErrorMessages []*string `json:"ErrorMessages,omitempty" name:"ErrorMessages"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateSignUrlsRequest struct {
	*tchttp.BaseRequest

	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 所签署合同ID数组
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 签署链接类型，默认：“WEIXINAPP”-直接跳小程序; “CHANNEL”-跳转H5页面; “APP”-第三方APP或小程序跳转电子签小程序;
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 签署完成后H5引导页跳转URL
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`
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
	delete(f, "Operator")
	delete(f, "Endpoint")
	delete(f, "JumpUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSignUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSignUrlsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 签署参与者签署H5链接信息数组
		SignUrlInfos []*SignUrlInfo `json:"SignUrlInfos,omitempty" name:"SignUrlInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeTemplatesRequest struct {
	*tchttp.BaseRequest

	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 模版唯一标识
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
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
	delete(f, "Operator")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板详情
		Templates []*TemplateInfo `json:"Templates,omitempty" name:"Templates"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeUsageRequest struct {
	*tchttp.BaseRequest

	// 应用信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 开始时间 eg:2021-03-21
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束时间 eg:2021-06-21 
	// 开始时间到结束时间的区间长度小于等于90天
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 操作者的信息
	Operator *UserInfo `json:"Operator,omitempty" name:"Operator"`

	// 是否汇总数据，默认不汇总
	// 不汇总:返回在统计区间内渠道下所有企业的每日明细，即每个企业N条数据，N为统计天数
	// 汇总:返回在统计区间内渠道下所有企业的汇总后数据，即每个企业一条数据
	NeedAggregate *bool `json:"NeedAggregate,omitempty" name:"NeedAggregate"`

	// 单次返回的最多条目数量,默认为1000,且不能超过1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量,默认是0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
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
	delete(f, "Operator")
	delete(f, "NeedAggregate")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用量明细条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 用量明细
	// 注意：此字段可能返回 null，表示取不到有效值。
		Details []*UsageDetail `json:"Details,omitempty" name:"Details"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type FlowApproverInfo struct {

	// 签署人姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 签署人手机号，脱敏显示
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 经办人身份证号
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 签署完前端跳转的url，暂未使用
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 签署截止时间
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 签署完回调url
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 签署人类型，PERSON和ORGANIZATION
	ApproverType *string `json:"ApproverType,omitempty" name:"ApproverType"`

	// 用户侧第三方id
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
}

type FlowInfo struct {

	// 合同名字
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署截止时间戳，超过有效签署时间则该签署流程失败
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 模版ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 合同类型：
	// 1. “劳务”
	// 2. “销售”
	// 3. “租赁”
	// 4. “其他”
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 多个签署人信息
	FlowApprovers []*FlowApproverInfo `json:"FlowApprovers,omitempty" name:"FlowApprovers"`

	// 表单K-V对列表
	FormFields []*FormField `json:"FormFields,omitempty" name:"FormFields"`

	// 合同描述
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 渠道的业务信息，限制1024字符
	CustomerData *string `json:"CustomerData,omitempty" name:"CustomerData"`
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

type PrepareFlowsRequest struct {
	*tchttp.BaseRequest

	// 渠道应用相关信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 多个合同（流程）信息
	FlowInfos []*FlowInfo `json:"FlowInfos,omitempty" name:"FlowInfos"`

	// 操作完成后的跳转地址
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

type PrepareFlowsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 待发起文件确认页
		ConfirmUrl *string `json:"ConfirmUrl,omitempty" name:"ConfirmUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 经办人ID（渠道颁发）
	Id *string `json:"Id,omitempty" name:"Id"`

	// 经办人姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 经办人身份证件类型
	// 用户证件类型：默认ID_CARD
	// 1. ID_CARD - 居民身份证
	// 2. HOUSEHOLD_REGISTER - 户口本
	// 3. TEMP_ID_CARD - 临时居民身份证
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 经办人身份证号
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 经办人手机号
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
}

type SignUrlInfo struct {

	// 签署链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignUrl *string `json:"SignUrl,omitempty" name:"SignUrl"`

	// 链接失效时间
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
}

type SyncFailReason struct {

	// 经办人Id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 失败原因
	// 例如：Id不符合规范、证件号码不合法等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type SyncProxyOrganizationOperatorsRequest struct {
	*tchttp.BaseRequest

	// 操作类型，新增:"CREATE"，修改:"UPDATE"，离职:"RESIGN"
	OperatorType *string `json:"OperatorType,omitempty" name:"OperatorType"`

	// 应用信息
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 经办人信息列表
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
	delete(f, "OperatorType")
	delete(f, "Agent")
	delete(f, "ProxyOrganizationOperators")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncProxyOrganizationOperatorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SyncProxyOrganizationOperatorsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type SyncProxyOrganizationRequest struct {
	*tchttp.BaseRequest

	// 应用信息
	// 此接口Agent.ProxyOrganizationOpenId必填
	Agent *Agent `json:"Agent,omitempty" name:"Agent"`

	// 渠道侧合作企业名称
	ProxyOrganizationName *string `json:"ProxyOrganizationName,omitempty" name:"ProxyOrganizationName"`

	// 渠道侧合作企业统一社会信用代码
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitempty" name:"UniformSocialCreditCode"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"BusinessLicense,omitempty" name:"BusinessLicense"`

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
	delete(f, "UniformSocialCreditCode")
	delete(f, "BusinessLicense")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncProxyOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SyncProxyOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 模板的创建者信息
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 模板创建的时间戳（精确到秒）
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 模板类型：1-静默签；2-静默签授权；3-普通模版
	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`

	// 模板中的流程参与人信息
	Recipients []*Recipient `json:"Recipients,omitempty" name:"Recipients"`
}

type UsageDetail struct {

	// 渠道侧合作企业唯一标识
	ProxyOrganizationOpenId *string `json:"ProxyOrganizationOpenId,omitempty" name:"ProxyOrganizationOpenId"`

	// 消耗量
	Usage *uint64 `json:"Usage,omitempty" name:"Usage"`

	// 日期，当需要汇总数据时日期为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitempty" name:"Date"`
}

type UserInfo struct {

	// 自定义用户编号
	CustomUserId *string `json:"CustomUserId,omitempty" name:"CustomUserId"`

	// 用户的来源渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 用户在渠道的编号
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 用户真实IP
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 用户代理IP
	ProxyIp *string `json:"ProxyIp,omitempty" name:"ProxyIp"`
}
