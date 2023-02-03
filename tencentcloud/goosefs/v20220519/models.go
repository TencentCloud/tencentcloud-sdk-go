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
	// 数据流通任务类型, FS_TO_COS(文件系统到COS Bucket),或者COS_TO_FS(COS Bucket到文件系统)
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// COS存储桶名
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 对于FS_TO_COS, TaskPath是Bucket映射目录的相对路径, 对于COS_TO_FS是COS上的路径。如果置为空, 则表示全部数据
	TaskPath *string `json:"TaskPath,omitempty" name:"TaskPath"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 数据流通方式 MSP_AFM 手动加载  RAW_AFM 按需加载
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 文件列表下载地址，以http开头
	TextLocation *string `json:"TextLocation,omitempty" name:"TextLocation"`
}

type CreateDataRepositoryTaskRequest struct {
	*tchttp.BaseRequest
	
	// 数据流通任务类型, FS_TO_COS(文件系统到COS Bucket),或者COS_TO_FS(COS Bucket到文件系统)
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// COS存储桶名
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 对于FS_TO_COS, TaskPath是Bucket映射目录的相对路径, 对于COS_TO_FS是COS上的路径。如果置为空, 则表示全部数据
	TaskPath *string `json:"TaskPath,omitempty" name:"TaskPath"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 数据流通方式 MSP_AFM 手动加载  RAW_AFM 按需加载
	RepositoryType *string `json:"RepositoryType,omitempty" name:"RepositoryType"`

	// 文件列表下载地址，以http开头
	TextLocation *string `json:"TextLocation,omitempty" name:"TextLocation"`
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
	delete(f, "RepositoryType")
	delete(f, "TextLocation")
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

// Predefined struct for user
type DescribeDataRepositoryTaskStatusRequestParams struct {
	// task id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// file system id
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DescribeDataRepositoryTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// task id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// file system id
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeDataRepositoryTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataRepositoryTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataRepositoryTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataRepositoryTaskStatusResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态 0(初始化中), 1(运行中), 2(已完成), 3(任务失败)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataRepositoryTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataRepositoryTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeDataRepositoryTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataRepositoryTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}