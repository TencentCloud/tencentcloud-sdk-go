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

package v20220519

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateDataRepositoryTaskRequestParams struct {
	// 数据流通任务类型, FS_TO_COS(文件系统到COS Bucket),或者Bucket到文件系统(COS_TO_FS)
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// bucket名
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 对于FS_TO_COS, TaskPath是Bucket映射目录的相对路径, 对于COS_TO_FS是COS上的路径。如果置位空, 则表示全部数据
	TaskPath *string `json:"TaskPath,omitempty" name:"TaskPath"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

type CreateDataRepositoryTaskRequest struct {
	*tchttp.BaseRequest
	
	// 数据流通任务类型, FS_TO_COS(文件系统到COS Bucket),或者Bucket到文件系统(COS_TO_FS)
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// bucket名
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 对于FS_TO_COS, TaskPath是Bucket映射目录的相对路径, 对于COS_TO_FS是COS上的路径。如果置位空, 则表示全部数据
	TaskPath *string `json:"TaskPath,omitempty" name:"TaskPath"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

func (r *CreateDataRepositoryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataRepositoryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "Bucket")
	delete(f, "FileSystemId")
	delete(f, "TaskPath")
	delete(f, "TaskName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataRepositoryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataRepositoryTaskResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDataRepositoryTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataRepositoryTaskResponseParams `json:"Response"`
}

func (r *CreateDataRepositoryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataRepositoryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}