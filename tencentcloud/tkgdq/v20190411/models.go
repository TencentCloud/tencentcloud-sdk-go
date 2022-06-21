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

package v20190411

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeEntityRequestParams struct {
	// 实体名称
	EntityName *string `json:"EntityName,omitempty" name:"EntityName"`
}

type DescribeEntityRequest struct {
	*tchttp.BaseRequest
	
	// 实体名称
	EntityName *string `json:"EntityName,omitempty" name:"EntityName"`
}

func (r *DescribeEntityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEntityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EntityName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEntityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEntityResponseParams struct {
	// 返回查询实体相关信息
	Content *string `json:"Content,omitempty" name:"Content"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEntityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEntityResponseParams `json:"Response"`
}

func (r *DescribeEntityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEntityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelationRequestParams struct {
	// 输入第一个实体
	LeftEntityName *string `json:"LeftEntityName,omitempty" name:"LeftEntityName"`

	// 输入第二个实体
	RightEntityName *string `json:"RightEntityName,omitempty" name:"RightEntityName"`
}

type DescribeRelationRequest struct {
	*tchttp.BaseRequest
	
	// 输入第一个实体
	LeftEntityName *string `json:"LeftEntityName,omitempty" name:"LeftEntityName"`

	// 输入第二个实体
	RightEntityName *string `json:"RightEntityName,omitempty" name:"RightEntityName"`
}

func (r *DescribeRelationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LeftEntityName")
	delete(f, "RightEntityName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRelationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelationResponseParams struct {
	// 返回查询实体间的关系
	Content []*EntityRelationContent `json:"Content,omitempty" name:"Content"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRelationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRelationResponseParams `json:"Response"`
}

func (r *DescribeRelationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTripleRequestParams struct {
	// 三元组查询条件
	TripleCondition *string `json:"TripleCondition,omitempty" name:"TripleCondition"`
}

type DescribeTripleRequest struct {
	*tchttp.BaseRequest
	
	// 三元组查询条件
	TripleCondition *string `json:"TripleCondition,omitempty" name:"TripleCondition"`
}

func (r *DescribeTripleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTripleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TripleCondition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTripleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTripleResponseParams struct {
	// 返回三元组信息
	Content []*TripleContent `json:"Content,omitempty" name:"Content"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTripleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTripleResponseParams `json:"Response"`
}

func (r *DescribeTripleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTripleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EntityRelationContent struct {
	// 实体关系查询返回关系的object
	Object []*EntityRelationObject `json:"Object,omitempty" name:"Object"`

	// 实体关系查询返回关系的subject
	Subject []*EntityRelationSubject `json:"Subject,omitempty" name:"Subject"`

	// 实体关系查询返回的关系名称
	Relation *string `json:"Relation,omitempty" name:"Relation"`
}

type EntityRelationObject struct {
	// object对应id
	Id []*string `json:"Id,omitempty" name:"Id"`

	// object对应name
	Name []*string `json:"Name,omitempty" name:"Name"`

	// object对应popular值
	Popular []*int64 `json:"Popular,omitempty" name:"Popular"`
}

type EntityRelationSubject struct {
	// Subject对应id
	Id []*string `json:"Id,omitempty" name:"Id"`

	// Subject对应name
	Name []*string `json:"Name,omitempty" name:"Name"`

	// Subject对应popular
	Popular []*int64 `json:"Popular,omitempty" name:"Popular"`
}

type TripleContent struct {
	// 实体id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 实体名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实体order
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 实体流行度
	Popular *int64 `json:"Popular,omitempty" name:"Popular"`
}