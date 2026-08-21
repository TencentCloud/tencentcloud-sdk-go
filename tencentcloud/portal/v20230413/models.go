// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20230413

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type SearchDocumentItem struct {
	// <p>文档URL</p>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>文档标题</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>产品名称</p>
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// <p>文档片段</p>
	Snippet *string `json:"Snippet,omitnil,omitempty" name:"Snippet"`
}

// Predefined struct for user
type SearchDocumentsRequestParams struct {
	// <p>搜索关键词</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>页码</p><p>取值范围：[1, 99]</p>
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>每页条数</p><p>取值范围：[1, 20]</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>产品名称</p>
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`
}

type SearchDocumentsRequest struct {
	*tchttp.BaseRequest
	
	// <p>搜索关键词</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>页码</p><p>取值范围：[1, 99]</p>
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// <p>每页条数</p><p>取值范围：[1, 20]</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>产品名称</p>
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`
}

func (r *SearchDocumentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchDocumentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "ProductName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchDocumentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchDocumentsResponseParams struct {
	// <p>总数</p>
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>文档列表</p>
	Documents []*SearchDocumentItem `json:"Documents,omitnil,omitempty" name:"Documents"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchDocumentsResponse struct {
	*tchttp.BaseResponse
	Response *SearchDocumentsResponseParams `json:"Response"`
}

func (r *SearchDocumentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchDocumentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}