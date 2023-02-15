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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateBPFakeAPPListRequestParams struct {
	// 仿冒应用下载链接。请严格按照模版进行填写：https://bma-privacy-detection-1251316161.cosgz.myqcloud.com/20221206/f8c7521fbd84f4c4e7c2a25ac233857e/批量仿冒应用举报模板.xlsx
	FakeAPPs *string `json:"FakeAPPs,omitempty" name:"FakeAPPs"`
}

type CreateBPFakeAPPListRequest struct {
	*tchttp.BaseRequest
	
	// 仿冒应用下载链接。请严格按照模版进行填写：https://bma-privacy-detection-1251316161.cosgz.myqcloud.com/20221206/f8c7521fbd84f4c4e7c2a25ac233857e/批量仿冒应用举报模板.xlsx
	FakeAPPs *string `json:"FakeAPPs,omitempty" name:"FakeAPPs"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CompanyId *int64 `json:"CompanyId,omitempty" name:"CompanyId"`

	// 仿冒应用名称
	FakeAPPName *string `json:"FakeAPPName,omitempty" name:"FakeAPPName"`

	// 仿冒来源
	APPChan *string `json:"APPChan,omitempty" name:"APPChan"`

	// 仿冒应用包名
	FakeAPPPackageName *string `json:"FakeAPPPackageName,omitempty" name:"FakeAPPPackageName"`

	// 仿冒应用证书
	FakeAPPCert *string `json:"FakeAPPCert,omitempty" name:"FakeAPPCert"`

	// 仿冒应用大小
	FakeAPPSize *string `json:"FakeAPPSize,omitempty" name:"FakeAPPSize"`

	// 仿冒截图
	FakeAPPSnapshots []*string `json:"FakeAPPSnapshots,omitempty" name:"FakeAPPSnapshots"`

	// 备注
	Note *string `json:"Note,omitempty" name:"Note"`
}

type CreateBPFakeAPPRequest struct {
	*tchttp.BaseRequest
	
	// 企业id
	CompanyId *int64 `json:"CompanyId,omitempty" name:"CompanyId"`

	// 仿冒应用名称
	FakeAPPName *string `json:"FakeAPPName,omitempty" name:"FakeAPPName"`

	// 仿冒来源
	APPChan *string `json:"APPChan,omitempty" name:"APPChan"`

	// 仿冒应用包名
	FakeAPPPackageName *string `json:"FakeAPPPackageName,omitempty" name:"FakeAPPPackageName"`

	// 仿冒应用证书
	FakeAPPCert *string `json:"FakeAPPCert,omitempty" name:"FakeAPPCert"`

	// 仿冒应用大小
	FakeAPPSize *string `json:"FakeAPPSize,omitempty" name:"FakeAPPSize"`

	// 仿冒截图
	FakeAPPSnapshots []*string `json:"FakeAPPSnapshots,omitempty" name:"FakeAPPSnapshots"`

	// 备注
	Note *string `json:"Note,omitempty" name:"Note"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CompanyId *int64 `json:"CompanyId,omitempty" name:"CompanyId"`

	// 仿冒网址
	FakeURL *string `json:"FakeURL,omitempty" name:"FakeURL"`

	// 仿冒网址截图
	FakeURLSnapshots []*string `json:"FakeURLSnapshots,omitempty" name:"FakeURLSnapshots"`

	// 备注
	Note *string `json:"Note,omitempty" name:"Note"`
}

type CreateBPFakeURLRequest struct {
	*tchttp.BaseRequest
	
	// 企业id
	CompanyId *int64 `json:"CompanyId,omitempty" name:"CompanyId"`

	// 仿冒网址
	FakeURL *string `json:"FakeURL,omitempty" name:"FakeURL"`

	// 仿冒网址截图
	FakeURLSnapshots []*string `json:"FakeURLSnapshots,omitempty" name:"FakeURLSnapshots"`

	// 备注
	Note *string `json:"Note,omitempty" name:"Note"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 仿冒网址下载链接：请严格按照模版要求填写，https://bma-privacy-detection-1251316161.cosgz.myqcloud.com/20221124/ff3273b24104d03fa3a8d0629a7f71a9/批量仿冒网址举报模板.xlsx
	FakeURLs *string `json:"FakeURLs,omitempty" name:"FakeURLs"`
}

type CreateBPFakeURLsRequest struct {
	*tchttp.BaseRequest
	
	// 仿冒网址下载链接：请严格按照模版要求填写，https://bma-privacy-detection-1251316161.cosgz.myqcloud.com/20221124/ff3273b24104d03fa3a8d0629a7f71a9/批量仿冒网址举报模板.xlsx
	FakeURLs *string `json:"FakeURLs,omitempty" name:"FakeURLs"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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