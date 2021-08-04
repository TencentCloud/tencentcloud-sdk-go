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

package v20210701

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DeployServiceBatchDetail struct {

	// 旧实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldPodList *DeployServicePodDetail `json:"OldPodList,omitempty" name:"OldPodList"`

	// 新实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewPodList *DeployServicePodDetail `json:"NewPodList,omitempty" name:"NewPodList"`

	// 当前批次状态："WaitForTimeExceed", "WaitForResume", "Deploying", "Finish", "NotStart"
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchStatus *string `json:"BatchStatus,omitempty" name:"BatchStatus"`

	// 该批次预计旧实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodNum *int64 `json:"PodNum,omitempty" name:"PodNum"`

	// 批次id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchIndex *int64 `json:"BatchIndex,omitempty" name:"BatchIndex"`
}

type DeployServicePodDetail struct {

	// pod Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodId *string `json:"PodId,omitempty" name:"PodId"`

	// pod状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodStatus []*string `json:"PodStatus,omitempty" name:"PodStatus"`

	// pod版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodVersion *string `json:"PodVersion,omitempty" name:"PodVersion"`

	// pod创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// pod所在可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DeployStrategyConf struct {

	// 总分批数
	TotalBatchCount *int64 `json:"TotalBatchCount,omitempty" name:"TotalBatchCount"`

	// beta分批实例数
	BetaBatchNum *int64 `json:"BetaBatchNum,omitempty" name:"BetaBatchNum"`

	// 分批策略：0-全自动，1-全手动，2-beta分批，beta批一定是手动的
	DeployStrategyType *int64 `json:"DeployStrategyType,omitempty" name:"DeployStrategyType"`

	// 每批暂停间隔
	BatchInterval *int64 `json:"BatchInterval,omitempty" name:"BatchInterval"`
}

type DescribeDeployApplicationDetailRequest struct {
	*tchttp.BaseRequest

	// 服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *DescribeDeployApplicationDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployApplicationDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeployApplicationDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeployApplicationDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分批发布结果详情
		Result *TemDeployApplicationDetailInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeployApplicationDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployApplicationDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRunPodPage struct {

	// 分页下标
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 请求id
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

	// 条目
	PodList []*RunVersionPod `json:"PodList,omitempty" name:"PodList"`
}

type ResumeDeployApplicationRequest struct {
	*tchttp.BaseRequest

	// 需要开始下一批次的服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *ResumeDeployApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeDeployApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeDeployApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResumeDeployApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeDeployApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeDeployApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevertDeployApplicationRequest struct {
	*tchttp.BaseRequest

	// 需要回滚的服务id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 需要回滚的服务所在环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *RevertDeployApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevertDeployApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "EnvironmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevertDeployApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RevertDeployApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevertDeployApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevertDeployApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunVersionPod struct {

	// shell地址
	Webshell *string `json:"Webshell,omitempty" name:"Webshell"`

	// pod的id
	PodId *string `json:"PodId,omitempty" name:"PodId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例的ip
	PodIp *string `json:"PodIp,omitempty" name:"PodIp"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 部署版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployVersion *string `json:"DeployVersion,omitempty" name:"DeployVersion"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`
}

type TemDeployApplicationDetailInfo struct {

	// 分批发布策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployStrategyConf *DeployStrategyConf `json:"DeployStrategyConf,omitempty" name:"DeployStrategyConf"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 当前状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// beta分批详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BetaBatchDetail *DeployServiceBatchDetail `json:"BetaBatchDetail,omitempty" name:"BetaBatchDetail"`

	// 其他分批详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherBatchDetail []*DeployServiceBatchDetail `json:"OtherBatchDetail,omitempty" name:"OtherBatchDetail"`

	// 老版本pod列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldVersionPodList *DescribeRunPodPage `json:"OldVersionPodList,omitempty" name:"OldVersionPodList"`

	// 当前批次id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentBatchIndex *int64 `json:"CurrentBatchIndex,omitempty" name:"CurrentBatchIndex"`

	// 错误原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 当前批次状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentBatchStatus *string `json:"CurrentBatchStatus,omitempty" name:"CurrentBatchStatus"`
}
