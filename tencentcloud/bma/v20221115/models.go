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

package v20221115

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type BrandData struct {
	// 品牌Id
	CompanyId *int64 `json:"CompanyId,omitnil" name:"CompanyId"`

	// 企业名称
	CompanyName *string `json:"CompanyName,omitnil" name:"CompanyName"`

	// 品牌名称
	BrandName *string `json:"BrandName,omitnil" name:"BrandName"`

	// 联系电话
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 营业执照
	License *string `json:"License,omitnil" name:"License"`

	// 营业执照审核状态
	LicenseStatus *int64 `json:"LicenseStatus,omitnil" name:"LicenseStatus"`

	// 营业执照审核状态说明
	LicenseNote *string `json:"LicenseNote,omitnil" name:"LicenseNote"`

	// 授权书
	Authorization *string `json:"Authorization,omitnil" name:"Authorization"`

	// 授权书审核状态
	AuthorizationStatus *int64 `json:"AuthorizationStatus,omitnil" name:"AuthorizationStatus"`

	// 授权书审核状态说明
	AuthorizationNote *string `json:"AuthorizationNote,omitnil" name:"AuthorizationNote"`

	// 商标信息
	Trademarks []*TrademarkData `json:"Trademarks,omitnil" name:"Trademarks"`

	// 新增时间
	InsertTime *string `json:"InsertTime,omitnil" name:"InsertTime"`

	// 服务信息
	Services *ServiceData `json:"Services,omitnil" name:"Services"`

	// 账号id
	Uin *string `json:"Uin,omitnil" name:"Uin"`
}

// Predefined struct for user
type CreateBPBrandRequestParams struct {
	// 品牌名称
	BrandName *string `json:"BrandName,omitnil" name:"BrandName"`

	// 企业名称
	CompanyName *string `json:"CompanyName,omitnil" name:"CompanyName"`

	// 品牌logo
	BrandLogo *string `json:"BrandLogo,omitnil" name:"BrandLogo"`

	// 联系电话
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 营业执照
	License *string `json:"License,omitnil" name:"License"`

	// 授权书
	Authorization *string `json:"Authorization,omitnil" name:"Authorization"`

	// 商标名称
	TrademarkNames []*string `json:"TrademarkNames,omitnil" name:"TrademarkNames"`

	// 商标证明
	Trademarks []*string `json:"Trademarks,omitnil" name:"Trademarks"`

	// 是否涉及转让: 0-不转让 1-转让
	IsTransfers []*string `json:"IsTransfers,omitnil" name:"IsTransfers"`

	// 转让证明
	Transfers []*string `json:"Transfers,omitnil" name:"Transfers"`

	// 保护网址
	ProtectURLs []*string `json:"ProtectURLs,omitnil" name:"ProtectURLs"`

	// 保护应用
	ProtectAPPs []*string `json:"ProtectAPPs,omitnil" name:"ProtectAPPs"`

	// 保护公众号
	ProtectOfficialAccounts []*string `json:"ProtectOfficialAccounts,omitnil" name:"ProtectOfficialAccounts"`

	// 保护小程序
	ProtectMiniPrograms []*string `json:"ProtectMiniPrograms,omitnil" name:"ProtectMiniPrograms"`

	// 请求来源：0-反钓鱼 2-反假冒
	APISource *int64 `json:"APISource,omitnil" name:"APISource"`
}

type CreateBPBrandRequest struct {
	*tchttp.BaseRequest
	
	// 品牌名称
	BrandName *string `json:"BrandName,omitnil" name:"BrandName"`

	// 企业名称
	CompanyName *string `json:"CompanyName,omitnil" name:"CompanyName"`

	// 品牌logo
	BrandLogo *string `json:"BrandLogo,omitnil" name:"BrandLogo"`

	// 联系电话
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 营业执照
	License *string `json:"License,omitnil" name:"License"`

	// 授权书
	Authorization *string `json:"Authorization,omitnil" name:"Authorization"`

	// 商标名称
	TrademarkNames []*string `json:"TrademarkNames,omitnil" name:"TrademarkNames"`

	// 商标证明
	Trademarks []*string `json:"Trademarks,omitnil" name:"Trademarks"`

	// 是否涉及转让: 0-不转让 1-转让
	IsTransfers []*string `json:"IsTransfers,omitnil" name:"IsTransfers"`

	// 转让证明
	Transfers []*string `json:"Transfers,omitnil" name:"Transfers"`

	// 保护网址
	ProtectURLs []*string `json:"ProtectURLs,omitnil" name:"ProtectURLs"`

	// 保护应用
	ProtectAPPs []*string `json:"ProtectAPPs,omitnil" name:"ProtectAPPs"`

	// 保护公众号
	ProtectOfficialAccounts []*string `json:"ProtectOfficialAccounts,omitnil" name:"ProtectOfficialAccounts"`

	// 保护小程序
	ProtectMiniPrograms []*string `json:"ProtectMiniPrograms,omitnil" name:"ProtectMiniPrograms"`

	// 请求来源：0-反钓鱼 2-反假冒
	APISource *int64 `json:"APISource,omitnil" name:"APISource"`
}

func (r *CreateBPBrandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPBrandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BrandName")
	delete(f, "CompanyName")
	delete(f, "BrandLogo")
	delete(f, "Phone")
	delete(f, "License")
	delete(f, "Authorization")
	delete(f, "TrademarkNames")
	delete(f, "Trademarks")
	delete(f, "IsTransfers")
	delete(f, "Transfers")
	delete(f, "ProtectURLs")
	delete(f, "ProtectAPPs")
	delete(f, "ProtectOfficialAccounts")
	delete(f, "ProtectMiniPrograms")
	delete(f, "APISource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPBrandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPBrandResponseParams struct {
	// 企业id
	CompanyId *int64 `json:"CompanyId,omitnil" name:"CompanyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBPBrandResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPBrandResponseParams `json:"Response"`
}

func (r *CreateBPBrandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPBrandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFakeAPPListRequestParams struct {
	// 批量模板
	FakeAPPs *string `json:"FakeAPPs,omitnil" name:"FakeAPPs"`
}

type CreateBPFakeAPPListRequest struct {
	*tchttp.BaseRequest
	
	// 批量模板
	FakeAPPs *string `json:"FakeAPPs,omitnil" name:"FakeAPPs"`
}

func (r *CreateBPFakeAPPListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFakeAPPListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FakeAPPs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPFakeAPPListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFakeAPPListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBPFakeAPPListResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPFakeAPPListResponseParams `json:"Response"`
}

func (r *CreateBPFakeAPPListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFakeAPPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFakeAPPRequestParams struct {
	// 企业id
	CompanyId *int64 `json:"CompanyId,omitnil" name:"CompanyId"`

	// 仿冒应用名称
	FakeAPPName *string `json:"FakeAPPName,omitnil" name:"FakeAPPName"`

	// 仿冒来源
	APPChan *string `json:"APPChan,omitnil" name:"APPChan"`

	// 仿冒应用包名
	FakeAPPPackageName *string `json:"FakeAPPPackageName,omitnil" name:"FakeAPPPackageName"`

	// 仿冒应用证书
	FakeAPPCert *string `json:"FakeAPPCert,omitnil" name:"FakeAPPCert"`

	// 仿冒应用大小
	FakeAPPSize *string `json:"FakeAPPSize,omitnil" name:"FakeAPPSize"`

	// 仿冒截图
	FakeAPPSnapshots []*string `json:"FakeAPPSnapshots,omitnil" name:"FakeAPPSnapshots"`

	// 备注
	Note *string `json:"Note,omitnil" name:"Note"`
}

type CreateBPFakeAPPRequest struct {
	*tchttp.BaseRequest
	
	// 企业id
	CompanyId *int64 `json:"CompanyId,omitnil" name:"CompanyId"`

	// 仿冒应用名称
	FakeAPPName *string `json:"FakeAPPName,omitnil" name:"FakeAPPName"`

	// 仿冒来源
	APPChan *string `json:"APPChan,omitnil" name:"APPChan"`

	// 仿冒应用包名
	FakeAPPPackageName *string `json:"FakeAPPPackageName,omitnil" name:"FakeAPPPackageName"`

	// 仿冒应用证书
	FakeAPPCert *string `json:"FakeAPPCert,omitnil" name:"FakeAPPCert"`

	// 仿冒应用大小
	FakeAPPSize *string `json:"FakeAPPSize,omitnil" name:"FakeAPPSize"`

	// 仿冒截图
	FakeAPPSnapshots []*string `json:"FakeAPPSnapshots,omitnil" name:"FakeAPPSnapshots"`

	// 备注
	Note *string `json:"Note,omitnil" name:"Note"`
}

func (r *CreateBPFakeAPPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFakeAPPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "FakeAPPName")
	delete(f, "APPChan")
	delete(f, "FakeAPPPackageName")
	delete(f, "FakeAPPCert")
	delete(f, "FakeAPPSize")
	delete(f, "FakeAPPSnapshots")
	delete(f, "Note")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPFakeAPPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFakeAPPResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBPFakeAPPResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPFakeAPPResponseParams `json:"Response"`
}

func (r *CreateBPFakeAPPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFakeAPPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFakeURLRequestParams struct {
	// 企业id
	CompanyId *int64 `json:"CompanyId,omitnil" name:"CompanyId"`

	// 仿冒网址
	FakeURL *string `json:"FakeURL,omitnil" name:"FakeURL"`

	// 仿冒网址截图
	FakeURLSnapshots []*string `json:"FakeURLSnapshots,omitnil" name:"FakeURLSnapshots"`

	// 备注
	Note *string `json:"Note,omitnil" name:"Note"`
}

type CreateBPFakeURLRequest struct {
	*tchttp.BaseRequest
	
	// 企业id
	CompanyId *int64 `json:"CompanyId,omitnil" name:"CompanyId"`

	// 仿冒网址
	FakeURL *string `json:"FakeURL,omitnil" name:"FakeURL"`

	// 仿冒网址截图
	FakeURLSnapshots []*string `json:"FakeURLSnapshots,omitnil" name:"FakeURLSnapshots"`

	// 备注
	Note *string `json:"Note,omitnil" name:"Note"`
}

func (r *CreateBPFakeURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFakeURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "FakeURL")
	delete(f, "FakeURLSnapshots")
	delete(f, "Note")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPFakeURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFakeURLResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBPFakeURLResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPFakeURLResponseParams `json:"Response"`
}

func (r *CreateBPFakeURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFakeURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFakeURLsRequestParams struct {
	// 批量模板
	FakeURLs *string `json:"FakeURLs,omitnil" name:"FakeURLs"`
}

type CreateBPFakeURLsRequest struct {
	*tchttp.BaseRequest
	
	// 批量模板
	FakeURLs *string `json:"FakeURLs,omitnil" name:"FakeURLs"`
}

func (r *CreateBPFakeURLsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFakeURLsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FakeURLs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPFakeURLsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPFakeURLsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBPFakeURLsResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPFakeURLsResponseParams `json:"Response"`
}

func (r *CreateBPFakeURLsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPFakeURLsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPWhiteListRequestParams struct {
	// 企业Id
	CompanyId *int64 `json:"CompanyId,omitnil" name:"CompanyId"`

	// 白名单类型：0-网站 1-应用 2-公众号 3-小程
	WhiteListType *int64 `json:"WhiteListType,omitnil" name:"WhiteListType"`

	// 白名单名称
	WhiteLists []*string `json:"WhiteLists,omitnil" name:"WhiteLists"`

	// 白名单备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type CreateBPWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 企业Id
	CompanyId *int64 `json:"CompanyId,omitnil" name:"CompanyId"`

	// 白名单类型：0-网站 1-应用 2-公众号 3-小程
	WhiteListType *int64 `json:"WhiteListType,omitnil" name:"WhiteListType"`

	// 白名单名称
	WhiteLists []*string `json:"WhiteLists,omitnil" name:"WhiteLists"`

	// 白名单备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *CreateBPWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "WhiteListType")
	delete(f, "WhiteLists")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBPWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBPWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBPWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *CreateBPWhiteListResponseParams `json:"Response"`
}

func (r *CreateBPWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBPWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBPWhiteListRequestParams struct {
	// 白名单id
	WhiteListId *int64 `json:"WhiteListId,omitnil" name:"WhiteListId"`
}

type DeleteBPWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 白名单id
	WhiteListId *int64 `json:"WhiteListId,omitnil" name:"WhiteListId"`
}

func (r *DeleteBPWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBPWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WhiteListId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBPWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBPWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBPWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBPWhiteListResponseParams `json:"Response"`
}

func (r *DeleteBPWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBPWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPBrandsRequestParams struct {

}

type DescribeBPBrandsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBPBrandsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPBrandsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBPBrandsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPBrandsResponseParams struct {
	// 品牌信息
	Brands []*BrandData `json:"Brands,omitnil" name:"Brands"`

	// 品牌审核通知栏状态：0 不显示 1 显示
	NoticeStatus *int64 `json:"NoticeStatus,omitnil" name:"NoticeStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBPBrandsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBPBrandsResponseParams `json:"Response"`
}

func (r *DescribeBPBrandsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPBrandsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPFakeAPPListRequestParams struct {
	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

type DescribeBPFakeAPPListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

func (r *DescribeBPFakeAPPListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPFakeAPPListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBPFakeAPPListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPFakeAPPListResponseParams struct {
	// 仿冒应用列表
	FakeAPPList []*FakeAPPData `json:"FakeAPPList,omitnil" name:"FakeAPPList"`

	// 仿冒应用总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBPFakeAPPListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBPFakeAPPListResponseParams `json:"Response"`
}

func (r *DescribeBPFakeAPPListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPFakeAPPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPFakeURLsRequestParams struct {
	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

type DescribeBPFakeURLsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

func (r *DescribeBPFakeURLsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPFakeURLsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBPFakeURLsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPFakeURLsResponseParams struct {
	// 仿冒网址列表
	FakeURLs []*FakeURLData `json:"FakeURLs,omitnil" name:"FakeURLs"`

	// 仿冒网址总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBPFakeURLsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBPFakeURLsResponseParams `json:"Response"`
}

func (r *DescribeBPFakeURLsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPFakeURLsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPWhiteListsRequestParams struct {
	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

type DescribeBPWhiteListsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 页数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`
}

func (r *DescribeBPWhiteListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPWhiteListsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBPWhiteListsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBPWhiteListsResponseParams struct {
	// 白名单列表
	WhiteLists []*WhiteListData `json:"WhiteLists,omitnil" name:"WhiteLists"`

	// 白名单总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBPWhiteListsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBPWhiteListsResponseParams `json:"Response"`
}

func (r *DescribeBPWhiteListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBPWhiteListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FakeAPPData struct {
	// 仿冒应用id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FakeAPPId *int64 `json:"FakeAPPId,omitnil" name:"FakeAPPId"`

	// 品牌名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrandName *string `json:"BrandName,omitnil" name:"BrandName"`

	// 仿冒来源：0-系统检测 1-人工举报
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *int64 `json:"Origin,omitnil" name:"Origin"`

	// 仿冒应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FakeAPPName *string `json:"FakeAPPName,omitnil" name:"FakeAPPName"`

	// 仿冒应用包名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FakeAPPPackageName *string `json:"FakeAPPPackageName,omitnil" name:"FakeAPPPackageName"`

	// 仿冒应用证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	FakeAPPCert *string `json:"FakeAPPCert,omitnil" name:"FakeAPPCert"`

	// 仿冒应用大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FakeAPPSize *string `json:"FakeAPPSize,omitnil" name:"FakeAPPSize"`

	// 热度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Heat *int64 `json:"Heat,omitnil" name:"Heat"`

	// 协助处置状态：0-未处置 1-处置中 2-处置成功 3-处置失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockStatus *int64 `json:"BlockStatus,omitnil" name:"BlockStatus"`

	// 协助处置状态说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockNote *string `json:"BlockNote,omitnil" name:"BlockNote"`

	// 关停状态：0-未关停 1-关停中 2-关停成功 3-关停失败 4-重复上架
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineStatus *int64 `json:"OfflineStatus,omitnil" name:"OfflineStatus"`

	// 关停状态说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineNote *string `json:"OfflineNote,omitnil" name:"OfflineNote"`

	// app来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadWay *string `json:"DownloadWay,omitnil" name:"DownloadWay"`

	// 新增时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitnil" name:"InsertTime"`

	// App下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadCosURL *string `json:"DownloadCosURL,omitnil" name:"DownloadCosURL"`

	// 资质证明状态:0-不可用 1-可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificationStatus *int64 `json:"CertificationStatus,omitnil" name:"CertificationStatus"`
}

type FakeURLData struct {
	// 仿冒网址id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FakeURLId *int64 `json:"FakeURLId,omitnil" name:"FakeURLId"`

	// 品牌名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrandName *string `json:"BrandName,omitnil" name:"BrandName"`

	// 仿冒来源：0-系统检测 1-人工举报
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *int64 `json:"Origin,omitnil" name:"Origin"`

	// 仿冒网址
	// 注意：此字段可能返回 null，表示取不到有效值。
	FakeURL *string `json:"FakeURL,omitnil" name:"FakeURL"`

	// 仿冒域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FakeDomain *string `json:"FakeDomain,omitnil" name:"FakeDomain"`

	// 热度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Heat *int64 `json:"Heat,omitnil" name:"Heat"`

	// 拦截处置状态：0-未处置 1-处置中 2-处置成功 3-处置失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockStatus *int64 `json:"BlockStatus,omitnil" name:"BlockStatus"`

	// 拦截处置状态说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockNote *string `json:"BlockNote,omitnil" name:"BlockNote"`

	// 关停状态：0-未关停 1-关停中 2-关停成功 3-关停失败 4-重复上架
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineStatus *int64 `json:"OfflineStatus,omitnil" name:"OfflineStatus"`

	// 关停状态说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineNote *string `json:"OfflineNote,omitnil" name:"OfflineNote"`

	// ip地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitnil" name:"IP"`

	// ip地理位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPLocation *string `json:"IPLocation,omitnil" name:"IPLocation"`

	// 网站所属单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebCompany *string `json:"WebCompany,omitnil" name:"WebCompany"`

	// 网站性质
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebAttribute *string `json:"WebAttribute,omitnil" name:"WebAttribute"`

	// 网站名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebName *string `json:"WebName,omitnil" name:"WebName"`

	// 备案号
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebICP *string `json:"WebICP,omitnil" name:"WebICP"`

	// 网站创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebCreateTime *string `json:"WebCreateTime,omitnil" name:"WebCreateTime"`

	// 网站过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebExpireTime *string `json:"WebExpireTime,omitnil" name:"WebExpireTime"`

	// 新增时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitnil" name:"InsertTime"`

	// 资质证明状态：0-不可用 1-可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificationStatus *int64 `json:"CertificationStatus,omitnil" name:"CertificationStatus"`

	// 网址截图
	// 注意：此字段可能返回 null，表示取不到有效值。
	Snapshot *string `json:"Snapshot,omitnil" name:"Snapshot"`

	// 账户资源状态：0-不可用 1-可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountStatus *int64 `json:"AccountStatus,omitnil" name:"AccountStatus"`

	// 审核状态：0-未审核 1-审核中 2-审核成功 3-审核失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuditStatus *int64 `json:"AuditStatus,omitnil" name:"AuditStatus"`

	// 下线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineTime *string `json:"OfflineTime,omitnil" name:"OfflineTime"`
}

type Filter struct {
	// 过滤参数键
	Name *string `json:"Name,omitnil" name:"Name"`

	// 过滤参数值
	Value *string `json:"Value,omitnil" name:"Value"`
}

type ServiceData struct {
	// 网站保护关联资产数
	ProtectURLCount *int64 `json:"ProtectURLCount,omitnil" name:"ProtectURLCount"`

	// 网站保护服务到期时间
	ProtectURLExpireTime *string `json:"ProtectURLExpireTime,omitnil" name:"ProtectURLExpireTime"`

	// 应用保护关联资产数
	ProtectAPPCount *int64 `json:"ProtectAPPCount,omitnil" name:"ProtectAPPCount"`

	// 应用保护服务到期时间
	ProtectAPPExpireTime *string `json:"ProtectAPPExpireTime,omitnil" name:"ProtectAPPExpireTime"`

	// 公众号保护关联资产数
	ProtectOfficialAccountCount *int64 `json:"ProtectOfficialAccountCount,omitnil" name:"ProtectOfficialAccountCount"`

	// 公众号保护服务到期时间
	ProtectOfficialAccountExpireTime *string `json:"ProtectOfficialAccountExpireTime,omitnil" name:"ProtectOfficialAccountExpireTime"`

	// 小程序保护关联资产数
	ProtectMiniProgramCount *int64 `json:"ProtectMiniProgramCount,omitnil" name:"ProtectMiniProgramCount"`

	// 小程序保护服务到期时间
	ProtectMiniProgramExpireTime *string `json:"ProtectMiniProgramExpireTime,omitnil" name:"ProtectMiniProgramExpireTime"`

	// 关停下架使用次数
	OfflineCount *int64 `json:"OfflineCount,omitnil" name:"OfflineCount"`
}

type TrademarkData struct {
	// 商标证明
	Trademark *string `json:"Trademark,omitnil" name:"Trademark"`

	// 商标审核状态
	TrademarkStatus *int64 `json:"TrademarkStatus,omitnil" name:"TrademarkStatus"`

	// 商标审核状态说明
	TrademarkNote *string `json:"TrademarkNote,omitnil" name:"TrademarkNote"`

	// 商标id
	TrademarkId *int64 `json:"TrademarkId,omitnil" name:"TrademarkId"`

	// 商标转让书
	Transfer *string `json:"Transfer,omitnil" name:"Transfer"`

	// 商标转让书审核状态
	TransferStatus *int64 `json:"TransferStatus,omitnil" name:"TransferStatus"`

	// 商标转让书审核状态说明
	TransferNote *string `json:"TransferNote,omitnil" name:"TransferNote"`

	// 商标名称
	TrademarkName *string `json:"TrademarkName,omitnil" name:"TrademarkName"`
}

type WhiteListData struct {
	// 白名单id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WhiteListId *int64 `json:"WhiteListId,omitnil" name:"WhiteListId"`

	// 企业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompanyId *int64 `json:"CompanyId,omitnil" name:"CompanyId"`

	// 品牌名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrandName *string `json:"BrandName,omitnil" name:"BrandName"`

	// 资产类型：0-网站 1-app 2-公众号 3-小程序
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetsType *int64 `json:"AssetsType,omitnil" name:"AssetsType"`

	// 白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	WhiteList *string `json:"WhiteList,omitnil" name:"WhiteList"`

	// 白名单备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 新增时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitnil" name:"InsertTime"`
}