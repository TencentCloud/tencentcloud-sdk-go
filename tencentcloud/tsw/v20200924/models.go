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

package v20200924

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AgentShell struct {
	// 鉴权token
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 数据接收Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	EtlIp *string `json:"EtlIp,omitempty" name:"EtlIp"`

	// 数据接收port
	// 注意：此字段可能返回 null，表示取不到有效值。
	EtlPort *string `json:"EtlPort,omitempty" name:"EtlPort"`

	// 手动接入脚本串
	// 注意：此字段可能返回 null，表示取不到有效值。
	ByHandAccess *string `json:"ByHandAccess,omitempty" name:"ByHandAccess"`

	// 自动接入脚本串
	// 注意：此字段可能返回 null，表示取不到有效值。
	ByShellAccess *string `json:"ByShellAccess,omitempty" name:"ByShellAccess"`

	// SkyWalking数据接收port
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkyWalkingPort *string `json:"SkyWalkingPort,omitempty" name:"SkyWalkingPort"`

	// Zipkin数据接收port
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZipkinPort *string `json:"ZipkinPort,omitempty" name:"ZipkinPort"`

	// Jaeger数据接收port
	// 注意：此字段可能返回 null，表示取不到有效值。
	JaegerPort *string `json:"JaegerPort,omitempty" name:"JaegerPort"`
}

// Predefined struct for user
type DescribeAgentShellRequestParams struct {

}

type DescribeAgentShellRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAgentShellRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentShellRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentShellRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentShellResponseParams struct {
	// 接入信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *AgentShell `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentShellResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentShellResponseParams `json:"Response"`
}

func (r *DescribeAgentShellResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentShellResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}