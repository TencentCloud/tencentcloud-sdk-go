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

package v20200507

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CopyCryptoColumnPolicyRequestParams struct {
	// 实例Id
	CasbId *string `json:"CasbId,omitempty" name:"CasbId"`

	// 元数据id
	MetaDataId *string `json:"MetaDataId,omitempty" name:"MetaDataId"`

	// 目标实例Id 如果和实例Id相同则为同CasbId下的策略复制
	DstCasbId *string `json:"DstCasbId,omitempty" name:"DstCasbId"`

	// 目标实例Id 如果和[元数据id]相同则为同元数据下的策略复制
	DstMetaDataId *string `json:"DstMetaDataId,omitempty" name:"DstMetaDataId"`

	// 筛选来源数据库的表
	SrcTableFilter []*CryptoCopyColumnPolicyTableFilter `json:"SrcTableFilter,omitempty" name:"SrcTableFilter"`

	// 复制同元数据下的策略，需要填写目标数据库名
	DstDatabaseName *string `json:"DstDatabaseName,omitempty" name:"DstDatabaseName"`
}

type CopyCryptoColumnPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	CasbId *string `json:"CasbId,omitempty" name:"CasbId"`

	// 元数据id
	MetaDataId *string `json:"MetaDataId,omitempty" name:"MetaDataId"`

	// 目标实例Id 如果和实例Id相同则为同CasbId下的策略复制
	DstCasbId *string `json:"DstCasbId,omitempty" name:"DstCasbId"`

	// 目标实例Id 如果和[元数据id]相同则为同元数据下的策略复制
	DstMetaDataId *string `json:"DstMetaDataId,omitempty" name:"DstMetaDataId"`

	// 筛选来源数据库的表
	SrcTableFilter []*CryptoCopyColumnPolicyTableFilter `json:"SrcTableFilter,omitempty" name:"SrcTableFilter"`

	// 复制同元数据下的策略，需要填写目标数据库名
	DstDatabaseName *string `json:"DstDatabaseName,omitempty" name:"DstDatabaseName"`
}

func (r *CopyCryptoColumnPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyCryptoColumnPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CasbId")
	delete(f, "MetaDataId")
	delete(f, "DstCasbId")
	delete(f, "DstMetaDataId")
	delete(f, "SrcTableFilter")
	delete(f, "DstDatabaseName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyCryptoColumnPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyCryptoColumnPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CopyCryptoColumnPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CopyCryptoColumnPolicyResponseParams `json:"Response"`
}

func (r *CopyCryptoColumnPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyCryptoColumnPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CryptoCopyColumnPolicyTableFilter struct {
	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 表名称
	TableNameSet []*string `json:"TableNameSet,omitempty" name:"TableNameSet"`
}