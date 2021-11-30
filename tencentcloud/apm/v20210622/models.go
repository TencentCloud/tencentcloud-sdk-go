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

type ApmAgentInfo struct {

	// Agent下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentDownloadURL *string `json:"AgentDownloadURL,omitempty" name:"AgentDownloadURL"`

	// Collector上报地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectorURL *string `json:"CollectorURL,omitempty" name:"CollectorURL"`

	// Token信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 外网上报地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicCollectorURL *string `json:"PublicCollectorURL,omitempty" name:"PublicCollectorURL"`

	// 自研VPC上报地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerCollectorURL *string `json:"InnerCollectorURL,omitempty" name:"InnerCollectorURL"`

	// 内网上报地址(Private Link上报地址)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateLinkCollectorURL *string `json:"PrivateLinkCollectorURL,omitempty" name:"PrivateLinkCollectorURL"`
}

type DescribeApmAgentRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 接入方式
	AgentType *string `json:"AgentType,omitempty" name:"AgentType"`

	// 环境
	NetworkMode *string `json:"NetworkMode,omitempty" name:"NetworkMode"`

	// 语言
	LanguageEnvironment *string `json:"LanguageEnvironment,omitempty" name:"LanguageEnvironment"`
}

func (r *DescribeApmAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentType")
	delete(f, "NetworkMode")
	delete(f, "LanguageEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmAgentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Agent信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApmAgent *ApmAgentInfo `json:"ApmAgent,omitempty" name:"ApmAgent"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
