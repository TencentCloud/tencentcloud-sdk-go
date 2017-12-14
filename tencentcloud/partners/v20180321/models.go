// Copyright 1999-2018 Tencent Ltd.
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

package cvm

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AgentBillElem struct {
	// 代理商账号ID
	Uin *string `json:"Uin" name:"Uin"`
	// 订单号，仅对预付费账单有意义
	OrderId *string `json:"OrderId" name:"OrderId"`
	// 代客账号ID
	ClientUin *string `json:"ClientUin" name:"ClientUin"`
	// 代客备注名称
	ClientRemark *string `json:"ClientRemark" name:"ClientRemark"`
	// 支付时间
	PayTime *string `json:"PayTime" name:"PayTime"`
	// 云产品名称
	GoodsType *string `json:"GoodsType" name:"GoodsType"`
	// 预付费/后付费
	PayMode *string `json:"PayMode" name:"PayMode"`
	// 支付月份
	SettleMonth *string `json:"SettleMonth" name:"SettleMonth"`
	// 支付金额，单位分
	Amt *uint64 `json:"Amt" name:"Amt"`
}

type AgentClientElem struct {
	// 代理商账号ID
	Uin *string `json:"Uin" name:"Uin"`
	// 代客账号ID
	ClientUin *string `json:"ClientUin" name:"ClientUin"`
	// 代客申请时间戳
	ApplyTime *uint64 `json:"ApplyTime" name:"ApplyTime"`
	// 代客类型，可能值为a/b
	ClientFlag *string `json:"ClientFlag" name:"ClientFlag"`
	// 代客邮箱，打码显示
	Mail *string `json:"Mail" name:"Mail"`
	// 代客手机，打码显示
	Phone *string `json:"Phone" name:"Phone"`
	// 0表示不欠费，1表示欠费
	HasOverdueBill *uint64 `json:"HasOverdueBill" name:"HasOverdueBill"`
}

type AuditApplyClientRequest struct {
	*tchttp.BaseRequest
	// 待审核客户账号ID
	ClientUin *string `json:"ClientUin" name:"ClientUin"`
	// 审核结果，可能的取值：accept/reject
	AuditResult *string `json:"AuditResult" name:"AuditResult"`
	// 申请理由，B类客户审核通过时必须填写申请理由
	Note *string `json:"Note" name:"Note"`
}

func (r *AuditApplyClientRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AuditApplyClientRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AuditApplyClientResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 代理商账号ID
		Uin *string `json:"Uin" name:"Uin"`
		// 客户账号ID
		ClientUin *string `json:"ClientUin" name:"ClientUin"`
		// 审核结果，包括accept/reject/qcloudaudit（腾讯云审核）
		AuditResult *string `json:"AuditResult" name:"AuditResult"`
		// 关联时间对应的时间戳
		AgentTime *uint64 `json:"AgentTime" name:"AgentTime"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuditApplyClientResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AuditApplyClientResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentBillsRequest struct {
	*tchttp.BaseRequest
	// 支付月份，如2018-02
	SettleMonth *string `json:"SettleMonth" name:"SettleMonth"`
	// 客户账号ID
	ClientUin *string `json:"ClientUin" name:"ClientUin"`
	// 支付方式，prepay/postpay
	PayMode *string `json:"PayMode" name:"PayMode"`
	// 预付费订单号
	OrderId *string `json:"OrderId" name:"OrderId"`
	// 客户备注名称
	ClientRemark *string `json:"ClientRemark" name:"ClientRemark"`
	// 偏移量
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 限制数目
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribeAgentBillsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentBillsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentBillsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合查询条件列表总数量
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 业务明细列表
		AgentBillSet []*AgentBillElem `json:"AgentBillSet" name:"AgentBillSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentBillsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentBillsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentClientsRequest struct {
	*tchttp.BaseRequest
	// 客户账号ID
	ClientUin *string `json:"ClientUin" name:"ClientUin"`
	// 客户名称。由于涉及隐私，名称打码显示，故名称仅支持打码后的模糊搜索
	ClientName *string `json:"ClientName" name:"ClientName"`
	// 客户类型，a/b，类型定义参考代理商相关政策文档
	ClientFlag *string `json:"ClientFlag" name:"ClientFlag"`
	// ASC/DESC， 不区分大小写，按申请时间排序
	OrderDirection *string `json:"OrderDirection" name:"OrderDirection"`
	// 偏移量
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 限制数目
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribeAgentClientsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentClientsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentClientsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 代客列表
		AgentClientSet []*AgentClientElem `json:"AgentClientSet" name:"AgentClientSet" list`
		// 符合条件的代客总数
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentClientsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentClientsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRebateInfosRequest struct {
	*tchttp.BaseRequest
	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth" name:"RebateMonth"`
	// 偏移量
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 限制数目
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribeRebateInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRebateInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRebateInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 返佣信息列表
		RebateInfoSet []*RebateInfoElem `json:"RebateInfoSet" name:"RebateInfoSet" list`
		// 符合查询条件返佣信息数目
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRebateInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRebateInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClientRemarkRequest struct {
	*tchttp.BaseRequest
	// 客户备注名称
	ClientRemark *string `json:"ClientRemark" name:"ClientRemark"`
	// 客户账号ID
	ClientUin *string `json:"ClientUin" name:"ClientUin"`
}

func (r *ModifyClientRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClientRemarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClientRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClientRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClientRemarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RebateInfoElem struct {
	// 代理商账号ID
	Uin *string `json:"Uin" name:"Uin"`
	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth" name:"RebateMonth"`
	// 返佣金额，单位分
	Amt *uint64 `json:"Amt" name:"Amt"`
	// 月度业绩，单位分
	MonthSales *uint64 `json:"MonthSales" name:"MonthSales"`
	// 季度业绩，单位分
	QuarterSales *uint64 `json:"QuarterSales" name:"QuarterSales"`
	// NORMAL(正常)/HAS_OVERDUE_BILL(欠费)/NO_CONTRACT(缺合同)
	ExceptionFlag *string `json:"ExceptionFlag" name:"ExceptionFlag"`
}
