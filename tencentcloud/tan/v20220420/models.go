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

package v20220420

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateBlockNodeRecordsRequestParams struct {
	// 盘查组id，可在“盘查组概览”功能中获取。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 节点id，可在“数据接入管理”中获取。
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// 节点数据json，具体demo请参考输入示例，其中key为数据接入管理中节点内创建的属性变量名，value为期望的推送值。
	Records *string `json:"Records,omitempty" name:"Records"`
}

type CreateBlockNodeRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 盘查组id，可在“盘查组概览”功能中获取。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 节点id，可在“数据接入管理”中获取。
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// 节点数据json，具体demo请参考输入示例，其中key为数据接入管理中节点内创建的属性变量名，value为期望的推送值。
	Records *string `json:"Records,omitempty" name:"Records"`
}

func (r *CreateBlockNodeRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlockNodeRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "NodeId")
	delete(f, "Records")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBlockNodeRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBlockNodeRecordsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBlockNodeRecordsResponse struct {
	*tchttp.BaseResponse
	Response *CreateBlockNodeRecordsResponseParams `json:"Response"`
}

func (r *CreateBlockNodeRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlockNodeRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}