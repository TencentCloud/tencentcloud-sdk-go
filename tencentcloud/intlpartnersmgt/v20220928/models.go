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

package v20220928

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AllocateCustomerCreditRequestParams struct {
	// 分配客户信用的具体值
	AddedCredit *float64 `json:"AddedCredit,omitempty" name:"AddedCredit"`

	// 客户uin
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`
}

type AllocateCustomerCreditRequest struct {
	*tchttp.BaseRequest
	
	// 分配客户信用的具体值
	AddedCredit *float64 `json:"AddedCredit,omitempty" name:"AddedCredit"`

	// 客户uin
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`
}

func (r *AllocateCustomerCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateCustomerCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddedCredit")
	delete(f, "ClientUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateCustomerCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateCustomerCreditResponseParams struct {
	// 更新后的信用总额
	TotalCredit *float64 `json:"TotalCredit,omitempty" name:"TotalCredit"`

	// 更新后的信用余额
	RemainingCredit *float64 `json:"RemainingCredit,omitempty" name:"RemainingCredit"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AllocateCustomerCreditResponse struct {
	*tchttp.BaseResponse
	Response *AllocateCustomerCreditResponseParams `json:"Response"`
}

func (r *AllocateCustomerCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateCustomerCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CountryCodeItem struct {
	// 国家英文名
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// 国家中文名
	Name *string `json:"Name,omitempty" name:"Name"`

	// ISO2标准国家/地区代码
	IOS2 *string `json:"IOS2,omitempty" name:"IOS2"`

	// ISO3标准国家/地区代码
	IOS3 *string `json:"IOS3,omitempty" name:"IOS3"`

	// 电话代码
	Code *string `json:"Code,omitempty" name:"Code"`
}

// Predefined struct for user
type CreateAccountRequestParams struct {
	// 新建客户的账户类型标识。本接口取值为：personal或company
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// 注册邮件地址。需要调用方保证邮件地址的有效性和正确性。
	// 需要满足邮件的格式。如：account@qq.com
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 账户密码。
	// 长度限制：[8,20]。
	// 需同时包含数字、字母以及特殊符号（!@#$%^&*()等非空格）
	Password *string `json:"Password,omitempty" name:"Password"`

	// 二次确认密码。必须和Password值相同
	ConfirmPassword *string `json:"ConfirmPassword,omitempty" name:"ConfirmPassword"`

	// 客户手机号码。需要调用方保证手机号码的有效性和正确性。
	// 长度限制：[1,32]。支持全球手机号。如18888888888
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 客户的国家/地区代码。取值参考获取国家/地区码接口GetCountryCodes。如852
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// 客户的IOS2标准国家/地区代码。参考获取国家/地区码接口GetCountryCodes。需要与CountryCode值对应。如HK
	Area *string `json:"Area,omitempty" name:"Area"`

	// 拓展字段，默认为空
	Extended *string `json:"Extended,omitempty" name:"Extended"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest
	
	// 新建客户的账户类型标识。本接口取值为：personal或company
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// 注册邮件地址。需要调用方保证邮件地址的有效性和正确性。
	// 需要满足邮件的格式。如：account@qq.com
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 账户密码。
	// 长度限制：[8,20]。
	// 需同时包含数字、字母以及特殊符号（!@#$%^&*()等非空格）
	Password *string `json:"Password,omitempty" name:"Password"`

	// 二次确认密码。必须和Password值相同
	ConfirmPassword *string `json:"ConfirmPassword,omitempty" name:"ConfirmPassword"`

	// 客户手机号码。需要调用方保证手机号码的有效性和正确性。
	// 长度限制：[1,32]。支持全球手机号。如18888888888
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 客户的国家/地区代码。取值参考获取国家/地区码接口GetCountryCodes。如852
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// 客户的IOS2标准国家/地区代码。参考获取国家/地区码接口GetCountryCodes。需要与CountryCode值对应。如HK
	Area *string `json:"Area,omitempty" name:"Area"`

	// 拓展字段，默认为空
	Extended *string `json:"Extended,omitempty" name:"Extended"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountType")
	delete(f, "Mail")
	delete(f, "Password")
	delete(f, "ConfirmPassword")
	delete(f, "PhoneNum")
	delete(f, "CountryCode")
	delete(f, "Area")
	delete(f, "Extended")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountResponseParams struct {
	// 账号的uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountResponseParams `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCountryCodesRequestParams struct {

}

type GetCountryCodesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetCountryCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCountryCodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCountryCodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCountryCodesResponseParams struct {
	// 国家地区代码列表
	Data []*CountryCodeItem `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetCountryCodesResponse struct {
	*tchttp.BaseResponse
	Response *GetCountryCodesResponseParams `json:"Response"`
}

func (r *GetCountryCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCountryCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCreditAllocationHistoryData struct {
	// 分配时间
	AllocatedTime *string `json:"AllocatedTime,omitempty" name:"AllocatedTime"`

	// 操作员
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 分配的信用值
	Credit *float64 `json:"Credit,omitempty" name:"Credit"`

	// 分配的总金额
	AllocatedCredit *float64 `json:"AllocatedCredit,omitempty" name:"AllocatedCredit"`
}

// Predefined struct for user
type QueryCreditAllocationHistoryRequestParams struct {
	// 客户uin
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`

	// 翻页参数，所在页数
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 翻页参数，每页数据量
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

type QueryCreditAllocationHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 客户uin
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`

	// 翻页参数，所在页数
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 翻页参数，每页数据量
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *QueryCreditAllocationHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditAllocationHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	delete(f, "Page")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCreditAllocationHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCreditAllocationHistoryResponseParams struct {
	// 历史信息总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 历史信息详细列表数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	History []*QueryCreditAllocationHistoryData `json:"History,omitempty" name:"History"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryCreditAllocationHistoryResponse struct {
	*tchttp.BaseResponse
	Response *QueryCreditAllocationHistoryResponseParams `json:"Response"`
}

func (r *QueryCreditAllocationHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditAllocationHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCreditByUinListRequestParams struct {
	// 用户列表
	UinList []*uint64 `json:"UinList,omitempty" name:"UinList"`
}

type QueryCreditByUinListRequest struct {
	*tchttp.BaseRequest
	
	// 用户列表
	UinList []*uint64 `json:"UinList,omitempty" name:"UinList"`
}

func (r *QueryCreditByUinListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditByUinListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UinList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCreditByUinListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCreditByUinListResponseParams struct {
	// 用户信息列表
	Data []*QueryDirectCustomersCreditData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryCreditByUinListResponse struct {
	*tchttp.BaseResponse
	Response *QueryCreditByUinListResponseParams `json:"Response"`
}

func (r *QueryCreditByUinListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditByUinListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomersCreditData struct {
	// 名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 电话
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 欠费标记
	Arrears *string `json:"Arrears,omitempty" name:"Arrears"`

	// 绑定时间
	AssociationTime *string `json:"AssociationTime,omitempty" name:"AssociationTime"`

	// 最近到期时间
	RecentExpiry *string `json:"RecentExpiry,omitempty" name:"RecentExpiry"`

	// 子客uin
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`

	// 子客授信额度
	Credit *float64 `json:"Credit,omitempty" name:"Credit"`

	// 子客剩余额度
	RemainingCredit *float64 `json:"RemainingCredit,omitempty" name:"RemainingCredit"`

	// 0：未实名 1: 个人实名 2：企业实名
	IdentifyType *uint64 `json:"IdentifyType,omitempty" name:"IdentifyType"`

	// 客户备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 强制状态
	Force *int64 `json:"Force,omitempty" name:"Force"`
}

// Predefined struct for user
type QueryCustomersCreditRequestParams struct {
	// 搜索条件类型，只能是：ClientUin|Name|Remark|Email四选一
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// 查询条件值
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 分页参数：当前页数，从1开始
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// 分页参数：每页条数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 排序参数，根据AssociationTime按照空或者desc：逆序，asc：顺序的方式进行排序
	Order *string `json:"Order,omitempty" name:"Order"`
}

type QueryCustomersCreditRequest struct {
	*tchttp.BaseRequest
	
	// 搜索条件类型，只能是：ClientUin|Name|Remark|Email四选一
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// 查询条件值
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 分页参数：当前页数，从1开始
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// 分页参数：每页条数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 排序参数，根据AssociationTime按照空或者desc：逆序，asc：顺序的方式进行排序
	Order *string `json:"Order,omitempty" name:"Order"`
}

func (r *QueryCustomersCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCustomersCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FilterType")
	delete(f, "Filter")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCustomersCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCustomersCreditResponseParams struct {
	// 查询客户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*QueryCustomersCreditData `json:"Data,omitempty" name:"Data"`

	// 客户数量
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryCustomersCreditResponse struct {
	*tchttp.BaseResponse
	Response *QueryCustomersCreditResponseParams `json:"Response"`
}

func (r *QueryCustomersCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCustomersCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDirectCustomersCreditData struct {
	// 用户Uin
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 总信用值
	TotalCredit *float64 `json:"TotalCredit,omitempty" name:"TotalCredit"`

	// 剩余信用值
	RemainingCredit *float64 `json:"RemainingCredit,omitempty" name:"RemainingCredit"`
}

// Predefined struct for user
type QueryDirectCustomersCreditRequestParams struct {

}

type QueryDirectCustomersCreditRequest struct {
	*tchttp.BaseRequest
	
}

func (r *QueryDirectCustomersCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDirectCustomersCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryDirectCustomersCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryDirectCustomersCreditResponseParams struct {
	// 直接子客信息列表
	Data []*QueryDirectCustomersCreditData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryDirectCustomersCreditResponse struct {
	*tchttp.BaseResponse
	Response *QueryDirectCustomersCreditResponseParams `json:"Response"`
}

func (r *QueryDirectCustomersCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDirectCustomersCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryPartnerCreditRequestParams struct {

}

type QueryPartnerCreditRequest struct {
	*tchttp.BaseRequest
	
}

func (r *QueryPartnerCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryPartnerCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryPartnerCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryPartnerCreditResponseParams struct {
	// 已分配额度
	AllocatedCredit *float64 `json:"AllocatedCredit,omitempty" name:"AllocatedCredit"`

	// 额度总数
	TotalCredit *float64 `json:"TotalCredit,omitempty" name:"TotalCredit"`

	// 剩余额度
	RemainingCredit *float64 `json:"RemainingCredit,omitempty" name:"RemainingCredit"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryPartnerCreditResponse struct {
	*tchttp.BaseResponse
	Response *QueryPartnerCreditResponseParams `json:"Response"`
}

func (r *QueryPartnerCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryPartnerCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryVoucherAmountByUinItem struct {
	// 子客uin
	ClientUin *int64 `json:"ClientUin,omitempty" name:"ClientUin"`

	// 代金券额度
	TotalAmount *float64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

	// 代金券余额
	RemainAmount *float64 `json:"RemainAmount,omitempty" name:"RemainAmount"`
}

// Predefined struct for user
type QueryVoucherAmountByUinRequestParams struct {
	// 子客uin列表
	ClientUins []*uint64 `json:"ClientUins,omitempty" name:"ClientUins"`
}

type QueryVoucherAmountByUinRequest struct {
	*tchttp.BaseRequest
	
	// 子客uin列表
	ClientUins []*uint64 `json:"ClientUins,omitempty" name:"ClientUins"`
}

func (r *QueryVoucherAmountByUinRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherAmountByUinRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryVoucherAmountByUinRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryVoucherAmountByUinResponseParams struct {
	// 子客代金券额度数据
	Data []*QueryVoucherAmountByUinItem `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryVoucherAmountByUinResponse struct {
	*tchttp.BaseResponse
	Response *QueryVoucherAmountByUinResponseParams `json:"Response"`
}

func (r *QueryVoucherAmountByUinResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherAmountByUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryVoucherListByUinItem struct {
	// 子客uin
	ClientUin *int64 `json:"ClientUin,omitempty" name:"ClientUin"`

	// 券总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 券详情
	Data []*QueryVoucherListByUinVoucherItem `json:"Data,omitempty" name:"Data"`
}

// Predefined struct for user
type QueryVoucherListByUinRequestParams struct {
	// 子客uin列表
	ClientUins []*uint64 `json:"ClientUins,omitempty" name:"ClientUins"`

	// 状态，不传时默认查询所有状态。Unused,Used,Expired
	Status *string `json:"Status,omitempty" name:"Status"`
}

type QueryVoucherListByUinRequest struct {
	*tchttp.BaseRequest
	
	// 子客uin列表
	ClientUins []*uint64 `json:"ClientUins,omitempty" name:"ClientUins"`

	// 状态，不传时默认查询所有状态。Unused,Used,Expired
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *QueryVoucherListByUinRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherListByUinRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUins")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryVoucherListByUinRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryVoucherListByUinResponseParams struct {
	// 子客代金券数据
	Data []*QueryVoucherListByUinItem `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryVoucherListByUinResponse struct {
	*tchttp.BaseResponse
	Response *QueryVoucherListByUinResponseParams `json:"Response"`
}

func (r *QueryVoucherListByUinResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherListByUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryVoucherListByUinVoucherItem struct {
	// 券号
	VoucherId *string `json:"VoucherId,omitempty" name:"VoucherId"`

	// 状态
	VoucherStatus *string `json:"VoucherStatus,omitempty" name:"VoucherStatus"`

	// 面额
	TotalAmount *float64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

	// 余额
	RemainAmount *float64 `json:"RemainAmount,omitempty" name:"RemainAmount"`
}

// Predefined struct for user
type QueryVoucherPoolRequestParams struct {

}

type QueryVoucherPoolRequest struct {
	*tchttp.BaseRequest
	
}

func (r *QueryVoucherPoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherPoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryVoucherPoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryVoucherPoolResponseParams struct {
	// 经销商姓名
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 经销商角色类型：1:经销商 2:总经销商 3:二级经销商
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

	// 总额度
	TotalQuota *float64 `json:"TotalQuota,omitempty" name:"TotalQuota"`

	// 剩余额度
	RemainingQuota *float64 `json:"RemainingQuota,omitempty" name:"RemainingQuota"`

	// 已发放的代金券数量
	IssuedNum *int64 `json:"IssuedNum,omitempty" name:"IssuedNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryVoucherPoolResponse struct {
	*tchttp.BaseResponse
	Response *QueryVoucherPoolResponseParams `json:"Response"`
}

func (r *QueryVoucherPoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}