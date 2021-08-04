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

package v20210622

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateProjectRequest struct {
	*tchttp.BaseRequest

	// 创建的项目名(不为空且最长为 200)
	Name *string `json:"Name,omitempty" name:"Name"`

	// 项目对应实例 id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 项目抽样率(大于等于 0)
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// 是否开启聚类
	EnableURLGroup *uint64 `json:"EnableURLGroup,omitempty" name:"EnableURLGroup"`

	// 项目类型("web", "mp", "android", "ios", "node", "hippy", "weex", "viola", "rn")
	Type *string `json:"Type,omitempty" name:"Type"`

	// 项目对应仓库地址(可选，最长为 256)
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// 项目对应网页地址(可选，最长为 256)
	URL *string `json:"URL,omitempty" name:"URL"`

	// 创建的项目描述(可选，最长为 1000)
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "InstanceID")
	delete(f, "Rate")
	delete(f, "EnableURLGroup")
	delete(f, "Type")
	delete(f, "Repo")
	delete(f, "URL")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目 id
		ID *uint64 `json:"ID,omitempty" name:"ID"`

		// 项目唯一key
		Key *string `json:"Key,omitempty" name:"Key"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
