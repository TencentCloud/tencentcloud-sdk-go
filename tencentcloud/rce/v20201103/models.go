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

package v20201103

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccountInfo struct {
	// 用户账号类型；默认开通QQOpenId、手机号MD5权限；如果需要使用微信OpenId入参，则需要"提交工单"或联系对接人进行资格审核，审核通过后方可正常使用微信开放账号。
	// 1：QQ开放账号
	// 2：微信开放账号
	// 10004：手机号MD5，中国大陆11位手机号进行MD5加密，取32位小写值
	// 10005：手机号SHA256，中国大陆11位手机号进行SHA256加密，取64位小写值
	AccountType *uint64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// QQ账号信息，AccountType是"1"时，该字段必填。
	QQAccount *QQAccountInfo `json:"QQAccount,omitnil,omitempty" name:"QQAccount"`

	// 微信账号信息，AccountType是"2"时，该字段必填。
	WeChatAccount *WeChatAccountInfo `json:"WeChatAccount,omitnil,omitempty" name:"WeChatAccount"`

	// 其它账号信息，AccountType是10004或10005时，该字段必填。
	OtherAccount *OtherAccountInfo `json:"OtherAccount,omitnil,omitempty" name:"OtherAccount"`
}

// Predefined struct for user
type CreateNameListRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputCreateNameListFront `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type CreateNameListRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputCreateNameListFront `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *CreateNameListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNameListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNameListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNameListResponseParams struct {
	// 业务出参
	Data *OutputCreateNameListFront `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNameListResponse struct {
	*tchttp.BaseResponse
	Response *CreateNameListResponseParams `json:"Response"`
}

func (r *CreateNameListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataAuthorizationInfo struct {
	// 数据委托方、需求方：客户主体名称。
	// 
	// 示例值：某某有限公司。
	DataProviderName *string `json:"DataProviderName,omitnil,omitempty" name:"DataProviderName"`

	// 数据受托方、提供方：腾讯云主体名称。
	// 
	// 固定填：腾讯云计算（北京）有限责任公司
	// 
	// 示例值：腾讯云计算（北京）有限责任公司
	DataRecipientName *string `json:"DataRecipientName,omitnil,omitempty" name:"DataRecipientName"`

	// 客户请求RCE所提供的用户数据类型，支持多选。实际以接口请求传参为准。
	// 
	// 1-手机号；
	// 
	// 2-微信开放账号；
	// 
	// 3-QQ开放账号；
	// 
	// 4-IP地址；
	// 
	// 999-其它；
	// 
	// 示例值：[1, 4]
	UserDataType []*uint64 `json:"UserDataType,omitnil,omitempty" name:"UserDataType"`

	// 客户是否已按[合规指南](https://rule.tencent.com/rule/202409130001)要求获取用户授权，同意客户委托腾讯云处理入参信息
	// 1-已授权；其它值为未授权。
	// 示例值：1
	IsAuthorize *uint64 `json:"IsAuthorize,omitnil,omitempty" name:"IsAuthorize"`

	// 客户是否已按[合规指南](https://rule.tencent.com/rule/202409130001)要求获取用户授权，同意腾讯云结合客户提供的信息，对已合法收集的用户数据进行必要处理得出服务结果，并返回给客户。
	// 1-已授权；其它值为未授权。
	// 示例值：1
	IsOrderHandling *uint64 `json:"IsOrderHandling,omitnil,omitempty" name:"IsOrderHandling"`

	// 客户获得的用户授权期限时间戳（单位秒）。
	// 
	// 不填默认无固定期限。
	// 
	// 示例值：1719805604
	AuthorizationTerm *uint64 `json:"AuthorizationTerm,omitnil,omitempty" name:"AuthorizationTerm"`

	// 	
	// 客户获得用户授权所依赖的协议地址。
	// 
	// 示例值：https://www.*****.com/*
	PrivacyPolicyLink *string `json:"PrivacyPolicyLink,omitnil,omitempty" name:"PrivacyPolicyLink"`

	// 是否是用户个人敏感数据。
	// 
	// 固定填：1。
	// 
	// 示例值：1
	IsPersonalData *uint64 `json:"IsPersonalData,omitnil,omitempty" name:"IsPersonalData"`
}

type DataContentInfo struct {
	// 名单数据内容
	DataContent *string `json:"DataContent,omitnil,omitempty" name:"DataContent"`

	// 名单数据描述
	DataRemark *string `json:"DataRemark,omitnil,omitempty" name:"DataRemark"`

	// 名单数据开始时间，时间格式示例"2024-05-05 12:10:15"
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 名单数据结束时间，时间格式示例"2024-05-05 12:10:15"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

// Predefined struct for user
type DeleteNameListDataRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputDeleteNameListData `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type DeleteNameListDataRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputDeleteNameListData `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *DeleteNameListDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNameListDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNameListDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNameListDataResponseParams struct {
	// 业务出参
	Data *OutputDeleteNameListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNameListDataResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNameListDataResponseParams `json:"Response"`
}

func (r *DeleteNameListDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNameListDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNameListRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputDeleteNameListFront `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type DeleteNameListRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputDeleteNameListFront `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *DeleteNameListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNameListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNameListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNameListResponseParams struct {
	// 业务出参
	Data *OutputDeleteNameListFront `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNameListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNameListResponseParams `json:"Response"`
}

func (r *DeleteNameListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNameListDataListRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputDescribeDataListFront `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type DescribeNameListDataListRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputDescribeDataListFront `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *DescribeNameListDataListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNameListDataListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNameListDataListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNameListDataListResponseParams struct {
	// 业务出参
	Data *OutputDescribeDataListFrontData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNameListDataListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNameListDataListResponseParams `json:"Response"`
}

func (r *DescribeNameListDataListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNameListDataListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNameListDetailRequestParams struct {
	// 业务入参	
	BusinessSecurityData *InputDescribeNameListDetail `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type DescribeNameListDetailRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参	
	BusinessSecurityData *InputDescribeNameListDetail `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *DescribeNameListDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNameListDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNameListDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNameListDetailResponseParams struct {
	// 黑白名单列表详情业务出参
	Data *OutputDescribeNameListDetailFront `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNameListDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNameListDetailResponseParams `json:"Response"`
}

func (r *DescribeNameListDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNameListDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNameListRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputDescribeNameListFront `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type DescribeNameListRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputDescribeNameListFront `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *DescribeNameListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNameListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNameListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNameListResponseParams struct {
	// 业务出参
	Data *OutputDescribeNameListFrontFixListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNameListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNameListResponseParams `json:"Response"`
}

func (r *DescribeNameListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserUsageCntRequestParams struct {

}

type DescribeUserUsageCntRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserUsageCntRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserUsageCntRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserUsageCntRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserUsageCntResponseParams struct {
	// 业务出参
	Data *OutputDescribeUserUsageCntData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserUsageCntResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserUsageCntResponseParams `json:"Response"`
}

func (r *DescribeUserUsageCntResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserUsageCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportNameListDataRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputImportNameListDataFront `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type ImportNameListDataRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputImportNameListDataFront `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *ImportNameListDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportNameListDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportNameListDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportNameListDataResponseParams struct {
	// 业务出参
	Data *OutputImportNameListDataFront `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportNameListDataResponse struct {
	*tchttp.BaseResponse
	Response *ImportNameListDataResponseParams `json:"Response"`
}

func (r *ImportNameListDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportNameListDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InputCreateNameListFront struct {
	// 名单名称
	ListName *string `json:"ListName,omitnil,omitempty" name:"ListName"`

	// 名单类型 [1 黑名单 2白名单]
	ListType *int64 `json:"ListType,omitnil,omitempty" name:"ListType"`

	// 数据类型[1 手机号 2 qqOpenId 3 2echatOpenId 4 ip 6 idfa 7 imei]
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// 描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 加密类型[0 无需加密 1 MD5加密 2 SHA256加密]
	EncryptionType *int64 `json:"EncryptionType,omitnil,omitempty" name:"EncryptionType"`

	// 场景Code，all_scene代表全部场景
	SceneCode *string `json:"SceneCode,omitnil,omitempty" name:"SceneCode"`
}

type InputCryptoManageMarketingRisk struct {
	// 是否授权：1已授权，否则未授权。
	//  调用全栈式风控引擎接口服务时，客户需先明确授权
	// 
	IsAuthorized *string `json:"IsAuthorized,omitnil,omitempty" name:"IsAuthorized"`

	// 加密类型：1AES加密
	CryptoType *string `json:"CryptoType,omitnil,omitempty" name:"CryptoType"`

	// 加密内容，非空时接口采用加密模式。
	CryptoContent *string `json:"CryptoContent,omitnil,omitempty" name:"CryptoContent"`
}

type InputDeleteNameListData struct {
	// 黑白名单数据ID集合
	NameListDataIdList []*int64 `json:"NameListDataIdList,omitnil,omitempty" name:"NameListDataIdList"`
}

type InputDeleteNameListFront struct {
	// 名单ID
	NameListId *int64 `json:"NameListId,omitnil,omitempty" name:"NameListId"`
}

type InputDescribeDataListFront struct {
	// 名单ID
	NameListId *int64 `json:"NameListId,omitnil,omitempty" name:"NameListId"`

	// 当前页数
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示条数	
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 搜索关键字，按照名单数据名称或加密名单数据名称搜索
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// 黑白名单列表状态[1 启用 2 停用]
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type InputDescribeNameListDetail struct {
	// 名单ID
	NameListId *int64 `json:"NameListId,omitnil,omitempty" name:"NameListId"`
}

type InputDescribeNameListFront struct {
	// 当前页数
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 名单类型 [1 黑名单 2 白名单]
	ListType *int64 `json:"ListType,omitnil,omitempty" name:"ListType"`

	// 数据类型[1 手机号 2 qqOpenId 3 wechatOpenId 4 ip 6 idfa 7 imei]
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// 关键字，按照名单名称搜索
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// 记录状态[1 启用 2 停用]
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type InputDetails struct {
	// 字段名称
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 字段值
	FieldValue *string `json:"FieldValue,omitnil,omitempty" name:"FieldValue"`
}

type InputImportNameListDataFront struct {
	// 名单ID
	NameListId *int64 `json:"NameListId,omitnil,omitempty" name:"NameListId"`

	// 数据来源，固定传2（手工录入）
	DataSource *int64 `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// 黑白名单数据内容
	DataContentInfo []*DataContentInfo `json:"DataContentInfo,omitnil,omitempty" name:"DataContentInfo"`
}

type InputManageMarketingRisk struct {
	// 用户账号类型；默认开通QQOpenId、手机号MD5权限；如果需要使用微信OpenId入参，则需要"提交工单"或联系对接人进行资格审核，审核通过后方可正常使用微信开放账号。
	// 1：QQ开放账号
	// 2：微信开放账号
	// 10004：手机号MD5，中国大陆11位手机号进行MD5加密，取32位小写值
	// 10005：手机号SHA256，中国大陆11位手机号进行SHA256加密，取64位小写值
	Account *AccountInfo `json:"Account,omitnil,omitempty" name:"Account"`

	// 场景码，用于识别和区分不同的业务场景，可在控制台上新建和管理
	// 控制台链接：https://console.cloud.tencent.com/rce/risk/strategy/scene-root
	// 活动防刷默认场景码：e_activity_antirush 
	// 登录保护默认场景码：e_login_protection
	// 注册保护默认场景码：e_register_protection
	SceneCode *string `json:"SceneCode,omitnil,omitempty" name:"SceneCode"`

	// 用户外网ip（传入用户非外网ip会影响判断结果）。
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 用户操作时间戳，精确到秒。
	PostTime *uint64 `json:"PostTime,omitnil,omitempty" name:"PostTime"`

	// 业务平台用户唯一标识，支持自定义。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 设备指纹DeviceToken值，集成设备指纹后获取；如果集成了相应的设备指纹，该字段必填。
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// 设备指纹 BusinessId。
	DeviceBusinessId *int64 `json:"DeviceBusinessId,omitnil,omitempty" name:"DeviceBusinessId"`

	// 业务ID。网站或应用在多个业务中使用此服务，通过此ID区分统计数据。
	BusinessId *uint64 `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 昵称，UTF-8 编码。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 用户邮箱地址。
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// 是否识别设备异常：
	// 0：不识别。
	// 1：识别。
	CheckDevice *int64 `json:"CheckDevice,omitnil,omitempty" name:"CheckDevice"`

	// 用户HTTP请求中的Cookie进行2次hash的值，只要保证相同Cookie的hash值一致即可。
	CookieHash *string `json:"CookieHash,omitnil,omitempty" name:"CookieHash"`

	// 用户HTTP请求的Referer值。
	Referer *string `json:"Referer,omitnil,omitempty" name:"Referer"`

	// 用户HTTP请求的User-Agent值。
	UserAgent *string `json:"UserAgent,omitnil,omitempty" name:"UserAgent"`

	// 用户HTTP请求的X-Forwarded-For值。
	XForwardedFor *string `json:"XForwardedFor,omitnil,omitempty" name:"XForwardedFor"`

	// MAC地址或设备唯一标识。
	MacAddress *string `json:"MacAddress,omitnil,omitempty" name:"MacAddress"`

	// 手机制造商ID，如果手机注册，请带上此信息。
	VendorId *string `json:"VendorId,omitnil,omitempty" name:"VendorId"`

	// 设备类型(已不推荐使用)。
	DeviceType *int64 `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 扩展字段。
	Details []*InputDetails `json:"Details,omitnil,omitempty" name:"Details"`

	// 邀请助力场景相关信息。
	Sponsor *SponsorInfo `json:"Sponsor,omitnil,omitempty" name:"Sponsor"`

	// 详情请跳转至OnlineScamInfo查看。
	OnlineScam *OnlineScamInfo `json:"OnlineScam,omitnil,omitempty" name:"OnlineScam"`

	// 1：Android
	// 2：iOS
	// 3：H5
	// 4：小程序
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 数据授权信息。
	DataAuthorization *DataAuthorizationInfo `json:"DataAuthorization,omitnil,omitempty" name:"DataAuthorization"`
}

type InputModifyNameFront struct {
	// 名单ID
	NameListId *int64 `json:"NameListId,omitnil,omitempty" name:"NameListId"`

	// 名单名称
	ListName *string `json:"ListName,omitnil,omitempty" name:"ListName"`

	// 名单状态 [1 启用 2 停用]
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type InputModifyNameListDataFront struct {
	// 名单数据ID
	NameListDataId *int64 `json:"NameListDataId,omitnil,omitempty" name:"NameListDataId"`

	// 名单数据内容
	DataContent *string `json:"DataContent,omitnil,omitempty" name:"DataContent"`

	// 名单数据开始时间，时间格式示例"2024-05-05 12:10:15"
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 名单数据结束时间，时间格式示例"2024-05-05 12:10:15"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 记录状态 [1 启用 2 停用]
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 名单数据描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type InputModifyNameListDataFrontListData struct {
	// 名单数据集合
	DataList []*InputModifyNameListDataFront `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type ManageMarketingRiskRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputManageMarketingRisk `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`

	// 业务入参
	BusinessCryptoData *InputCryptoManageMarketingRisk `json:"BusinessCryptoData,omitnil,omitempty" name:"BusinessCryptoData"`
}

type ManageMarketingRiskRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputManageMarketingRisk `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`

	// 业务入参
	BusinessCryptoData *InputCryptoManageMarketingRisk `json:"BusinessCryptoData,omitnil,omitempty" name:"BusinessCryptoData"`
}

func (r *ManageMarketingRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageMarketingRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	delete(f, "BusinessCryptoData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageMarketingRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageMarketingRiskResponseParams struct {
	// 业务出参
	Data *OutputManageMarketingRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ManageMarketingRiskResponse struct {
	*tchttp.BaseResponse
	Response *ManageMarketingRiskResponseParams `json:"Response"`
}

func (r *ManageMarketingRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageMarketingRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNameListDataRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputModifyNameListDataFrontListData `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type ModifyNameListDataRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputModifyNameListDataFrontListData `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *ModifyNameListDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNameListDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNameListDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNameListDataResponseParams struct {
	// 业务出参
	Data *OutputModifyNameListFront `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNameListDataResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNameListDataResponseParams `json:"Response"`
}

func (r *ModifyNameListDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNameListDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNameListRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputModifyNameFront `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type ModifyNameListRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputModifyNameFront `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *ModifyNameListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNameListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNameListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNameListResponseParams struct {
	// 业务出参
	Data *OutputModifyNameFront `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNameListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNameListResponseParams `json:"Response"`
}

func (r *ModifyNameListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OnlineScamInfo struct {
	// 内容标签。
	ContentLabel *string `json:"ContentLabel,omitnil,omitempty" name:"ContentLabel"`

	// 内容风险等级：
	// 0：正常。
	// 1：可疑。
	ContentRiskLevel *int64 `json:"ContentRiskLevel,omitnil,omitempty" name:"ContentRiskLevel"`

	// 内容产生形式：
	// 0：对话。
	// 1：广播。
	ContentType *int64 `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 类型
	FraudType *int64 `json:"FraudType,omitnil,omitempty" name:"FraudType"`

	// 账号
	FraudAccount *string `json:"FraudAccount,omitnil,omitempty" name:"FraudAccount"`
}

type OtherAccountInfo struct {
	// 其他账号信息；
	// AccountType是10004时，填入中国大陆标准11位手机号的MD5值
	// AccountType是10005时，填入中国大陆标准11位手机号的SHA256值
	// 注释：
	// MD5手机号加密方式，使用中国大陆11位手机号进行MD5加密，加密后取32位小写值。
	// SHA256手机号加密方式，使用中国大陆11位手机号进行SHA256加密，加密后取64位小写值。
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 账号绑定的MD5或SHA256加密的手机号（该字段已不推荐使用）。
	// 注释：支持标准中国大陆11位手机号MD5加密后位的32位小写字符串；
	//      支持标准中国大陆11位手机号SHA256加密后位的64位小写字符串。
	MobilePhone *string `json:"MobilePhone,omitnil,omitempty" name:"MobilePhone"`

	// 用户设备号（该字段已不推荐使用）。
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type OuntputDescribeDataListInfo struct {
	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*OutputDescribeDataListFront `json:"List,omitnil,omitempty" name:"List"`
}

type OutputCreateNameListFront struct {
	// 错误码，0 表示成功，非0表示失败错误码。 0：成功 1002：参数错误 4300：未开通服务 6000：系统内部错误
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 空数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputDeleteNameListData struct {
	// 错误码，0 表示成功，非0表示失败错误码。 0：成功 1002：参数错误 4300：未开通服务 6000：系统内部错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 空数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputDeleteNameListFront struct {
	// 错误码，0 表示成功，非0表示失败错误码。 0：成功 1002：参数错误 4300：未开通服务 6000：系统内部错误
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 空数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputDescribeDataListFront struct {
	// 名单数据ID
	NameListDataId *int64 `json:"NameListDataId,omitnil,omitempty" name:"NameListDataId"`

	// 名单ID
	NameListId *int64 `json:"NameListId,omitnil,omitempty" name:"NameListId"`

	// 名单数据内容
	DataContent *string `json:"DataContent,omitnil,omitempty" name:"DataContent"`

	// 数据来源，固定传2（手工录入）
	DataSource *int64 `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// 名单数据开始时间，时间格式示例"2024-05-05 12:10:15"
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 名单数据结束时间，时间格式示例"2024-05-05 12:10:15"
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 名单数据状态 [1 启用 2 停用]
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 名单数据描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 名单数据创建时间，时间格式示例"2024-05-05 12:10:15"
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 名单数据更新时间，时间格式示例"2024-05-05 12:10:15"
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 加密名单数据内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptDataContent *string `json:"EncryptDataContent,omitnil,omitempty" name:"EncryptDataContent"`
}

type OutputDescribeDataListFrontData struct {
	// 错误码，0 表示成功，非0表示失败错误码。 0：成功 1002：参数错误 4300：未开通服务 6000：系统内部错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 黑白名单数据信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *OuntputDescribeDataListInfo `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputDescribeNameListDetail struct {
	// 名单ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameListId *int64 `json:"NameListId,omitnil,omitempty" name:"NameListId"`

	// 名单名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListName *string `json:"ListName,omitnil,omitempty" name:"ListName"`

	// 名单类型 [1 黑名单 2 白名单]
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListType *int64 `json:"ListType,omitnil,omitempty" name:"ListType"`

	// 数据类型[1 手机号 2 qqOpenId 3 2echatOpenId 4 ip 6 idfa 7 imei]
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// 场景Code
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneCode *string `json:"SceneCode,omitnil,omitempty" name:"SceneCode"`

	// 名单列表状态 [1 启用 2 停用]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间，时间格式示例"2024-05-05 12:10:15"
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，时间格式示例"2024-05-05 12:10:15"
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 加密类型 [0 无需加密，1 MD5加密，2 SHA256加密]
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptionType *int64 `json:"EncryptionType,omitnil,omitempty" name:"EncryptionType"`
}

type OutputDescribeNameListDetailFront struct {
	// 错误码，0 表示成功，非0表示失败错误码。 0：成功 1002：参数错误 4300：未开通服务 6000：系统内部错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 列表详情信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *OutputDescribeNameListDetail `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputDescribeNameListFrontFix struct {
	// 名单ID
	NameListId *int64 `json:"NameListId,omitnil,omitempty" name:"NameListId"`

	// 名单名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListName *string `json:"ListName,omitnil,omitempty" name:"ListName"`

	// 名单类型 [1 黑名单 2 白名单]
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListType *int64 `json:"ListType,omitnil,omitempty" name:"ListType"`

	// 数据类型[1 手机号 2 qqOpenId 3 2echatOpenId 4 ip 6 idfa 7 imei]
	DataType *int64 `json:"DataType,omitnil,omitempty" name:"DataType"`

	// 记录状态 [1 启用 2 停用]
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间，时间格式示例"2024-05-05 12:10:15"
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，时间格式示例"2024-05-05 12:10:15"
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 有效数据/数据总数
	EffectCount *string `json:"EffectCount,omitnil,omitempty" name:"EffectCount"`

	// 加密类型[0 无需加密 1 MD5加密 2 SHA256加密]
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptionType *int64 `json:"EncryptionType,omitnil,omitempty" name:"EncryptionType"`

	// 场景Code，all_scene代表全部场景
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneCode *string `json:"SceneCode,omitnil,omitempty" name:"SceneCode"`
}

type OutputDescribeNameListFrontFixListData struct {
	// 错误码，0 表示成功，非0表示失败错误码。 0：成功 1002：参数错误 4300：未开通服务 6000：系统内部错误
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 黑白名单列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *OutputDescribeNameListInfo `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputDescribeNameListInfo struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*OutputDescribeNameListFrontFix `json:"List,omitnil,omitempty" name:"List"`
}

type OutputDescribeUserUsageCnt struct {
	// 当前付费模式，0 后付费 1 预付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 后付费本月使用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AfterPayModeThisMonthUsedCnt *int64 `json:"AfterPayModeThisMonthUsedCnt,omitnil,omitempty" name:"AfterPayModeThisMonthUsedCnt"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 超出时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 后付费上月使用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AfterPayModeLastMonthUsedCnt *int64 `json:"AfterPayModeLastMonthUsedCnt,omitnil,omitempty" name:"AfterPayModeLastMonthUsedCnt"`

	// 预付费总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeforePayModeTotalUsedCnt *int64 `json:"BeforePayModeTotalUsedCnt,omitnil,omitempty" name:"BeforePayModeTotalUsedCnt"`

	// 预付费剩余用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeforePayModeRemainUsedCnt *int64 `json:"BeforePayModeRemainUsedCnt,omitnil,omitempty" name:"BeforePayModeRemainUsedCnt"`
}

type OutputDescribeUserUsageCntData struct {
	// 错误码，0 表示成功，非0表示失败错误码。 0：成功 4300：未开通服务
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 业务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *OutputDescribeUserUsageCnt `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputImportNameListDataFront struct {
	// 错误码，0 表示成功，非0表示失败错误码。 0：成功 1002：参数错误 4300：未开通服务 6000：系统内部错误
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 空数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputManageMarketingRisk struct {
	// 错误码，0 表示成功，非0表示失败错误码。
	// 0：成功
	// 1：错误
	// 1002：参数错误
	// 4300：未开通服务
	// 4301：后端未创建对应产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// UTF-8编码，出错消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 业务详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *OutputManageMarketingRiskValue `json:"Value,omitnil,omitempty" name:"Value"`

	// 控制台显示的req_id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UUid *string `json:"UUid,omitnil,omitempty" name:"UUid"`
}

type OutputManageMarketingRiskValue struct {
	// 账号ID：对应输入参数。
	// 当AccountType为1时，对应QQ的OpenId；
	// 当AccountType为2时，对应微信的OpenId/UnionId；
	// 当AccountType为10004时，对应手机号的MD5值；
	// 当AccountType为10005时，对应手机号的SHA256值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 操作时间戳，单位秒（对应输入参数）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostTime *uint64 `json:"PostTime,omitnil,omitempty" name:"PostTime"`

	// 业务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociateAccount *string `json:"AssociateAccount,omitnil,omitempty" name:"AssociateAccount"`

	// 操作来源的外网IP（对应输入参数）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 风险等级
	// pass：无恶意
	// review：低风险，需要人工审核
	// reject：高风险，建议拦截
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 风险类型，可能同时命中多个风险类型
	// 1: 账号信用低，账号近期存在因恶意被处罚历史，网络低活跃，被举报等因素。
	// 11: 疑似低活跃账号，账号活跃度与正常用户有差异。
	// 2: 垃圾账号，疑似批量注册小号，近期存在严重违规或大量举报。
	// 21: 疑似小号，账号有疑似线上养号，小号等行为。
	// 22: 疑似违规账号，账号曾有违规行为、曾被举报过、曾因违规被处罚过等。
	// 3: 无效账号，送检账号参数无法成功解析，请检查微信 OpenId 是否有误/AppId 与 QQ OpenId 无法关联/微信 OpenId 权限是否开通/手机号是否为中国大陆手机号；
	// 4: 黑名单，该账号在业务侧有过拉黑记录。
	// 5: 白名单，业务自行有添加过白名单记录。
	// 101: 批量操作，存在 IP/设备/环境等因素的聚集性异常。
	// 1011: 疑似 IP 属性聚集，出现 IP 聚集。
	// 1012: 疑似设备属性聚集，出现设备聚集。
	// 102: 自动机，疑似自动机批量请求。
	// 103: 恶意行为-网赚，疑似网赚。
	// 104: 微信登录态无效，检查 WeChatAccessToken 参数，是否已经失效。
	// 201: 环境风险，环境异常操作 IP/设备/环境存在异常。当前 IP 为非常用 IP 或恶意 IP 段。
	// 2011: 疑似非常用IP，请求当前请求 IP 非该账号常用 IP。
	// 2012: 疑似 IP 异常，使用 IDC 机房 IP 或使用代理 IP 或使用恶意 IP 等。
	// 205: 非公网有效 IP，传进来的 IP 地址为内网 IP 地址或者 IP 保留地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskType []*int64 `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// 设备指纹ID，如果集成了设备指纹，并传入了正确的DeviceToken和Platform，该字段正常输出；如果DeviceToken异常（校验不通过），则会在RiskType中返回"-1"标签，ConstId字段为空；如果没有集成设备指纹ConstId字段默认为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConstId *string `json:"ConstId,omitnil,omitempty" name:"ConstId"`

	// 风险扩展数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskInformation *string `json:"RiskInformation,omitnil,omitempty" name:"RiskInformation"`
}

type OutputModifyNameFront struct {
	// 错误码，0 表示成功，非0表示失败错误码。 0：成功 1002：参数错误 4300：未开通服务 6000：系统内部错误
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 空数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputModifyNameListFront struct {
	// 错误码，0 表示成功，非0表示失败错误码。 0：成功 1002：参数错误 4300：未开通服务 6000：系统内部错误
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 空数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type QQAccountInfo struct {
	// QQ的OpenId。
	QQOpenId *string `json:"QQOpenId,omitnil,omitempty" name:"QQOpenId"`

	// QQ分配给网站或应用的AppId，用来唯一标识网站或应用。
	AppIdUser *string `json:"AppIdUser,omitnil,omitempty" name:"AppIdUser"`

	// 用于标识QQ用户登录后所关联业务自身的账号ID。
	AssociateAccount *string `json:"AssociateAccount,omitnil,omitempty" name:"AssociateAccount"`

	// 账号绑定的MD5或SHA256加密的手机号。
	// 注释：支持标准中国大陆11位手机号MD5加密后位的32位小写字符串；
	//      支持标准中国大陆11位手机号SHA256加密后位的64位小写字符串。
	MobilePhone *string `json:"MobilePhone,omitnil,omitempty" name:"MobilePhone"`

	// 用户设备号（已不推荐使用）。
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type SponsorInfo struct {
	// 助力场景建议填写：活动发起人微信OpenId。
	SponsorOpenId *string `json:"SponsorOpenId,omitnil,omitempty" name:"SponsorOpenId"`

	// 助力场景建议填写：发起人设备号
	SponsorDeviceNumber *string `json:"SponsorDeviceNumber,omitnil,omitempty" name:"SponsorDeviceNumber"`

	// 助力场景建议填写：发起人的MD5手机号
	SponsorPhone *string `json:"SponsorPhone,omitnil,omitempty" name:"SponsorPhone"`

	// 助力场景建议填写：发起人IP
	SponsorIp *string `json:"SponsorIp,omitnil,omitempty" name:"SponsorIp"`

	// 助力场景建议填写：活动链接
	CampaignUrl *string `json:"CampaignUrl,omitnil,omitempty" name:"CampaignUrl"`
}

type WeChatAccountInfo struct {
	// 微信的OpenId/UnionId。
	WeChatOpenId *string `json:"WeChatOpenId,omitnil,omitempty" name:"WeChatOpenId"`

	// 微信开放账号类型：
	// 1：微信公众号/微信第三方登录。
	// 2：微信小程序。
	WeChatSubType *uint64 `json:"WeChatSubType,omitnil,omitempty" name:"WeChatSubType"`

	// 随机串。如果WeChatSubType是2，该字段必填。Token签名随机数，建议16个字符。
	RandStr *string `json:"RandStr,omitnil,omitempty" name:"RandStr"`

	// 如果WeChatSubType 是1，填入授权的 access_token（注意：不是普通 access_token，详情请参阅官方说明文档。获取网页版本的 access_token 时，scope 字段必需填写snsapi_userinfo
	// 如果WeChatSubType是2，填入以session_key 为密钥签名随机数RandStr（hmac_sha256签名算法）得到的字符串。
	WeChatAccessToken *string `json:"WeChatAccessToken,omitnil,omitempty" name:"WeChatAccessToken"`

	// 用于标识微信用户登录后所关联业务自身的账号ID。
	AssociateAccount *string `json:"AssociateAccount,omitnil,omitempty" name:"AssociateAccount"`

	// 账号绑定的MD5或SHA256加密的手机号。
	// 注释：支持标准中国大陆11位手机号MD5加密后位的32位小写字符串；
	//      支持标准中国大陆11位手机号SHA256加密后位的64位小写字符串。
	MobilePhone *string `json:"MobilePhone,omitnil,omitempty" name:"MobilePhone"`

	// 用户设备号（已不推荐使用）。
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}