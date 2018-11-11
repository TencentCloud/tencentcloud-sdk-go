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

package v20180321

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	// 模块名
	Module *string `json:"Module" name:"Module"`
	// 操作名
	Operation *string `json:"Operation" name:"Operation"`
	// 任务ID
	TaskId *int64 `json:"TaskId" name:"TaskId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 任务结果
		TaskResult *string `json:"TaskResult" name:"TaskResult"`
		// 任务类型，010代表上传任务
		TaskType *string `json:"TaskType" name:"TaskType"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadReportRequest struct {
	*tchttp.BaseRequest
	// 模块名
	Module *string `json:"Module" name:"Module"`
	// 操作名
	Operation *string `json:"Operation" name:"Operation"`
	// 报告日期
	ReportDate *string `json:"ReportDate" name:"ReportDate"`
}

func (r *DownloadReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadReportRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadReportResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 日报下载地址
		DailyReportUrl *string `json:"DailyReportUrl" name:"DailyReportUrl"`
		// 结果下载地址
		ResultReportUrl *string `json:"ResultReportUrl" name:"ResultReportUrl"`
		// 明细下载地址
		DetailReportUrl *string `json:"DetailReportUrl" name:"DetailReportUrl"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadReportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadFileRequest struct {
	*tchttp.BaseRequest
	// 模块名
	Module *string `json:"Module" name:"Module"`
	// 操作名
	Operation *string `json:"Operation" name:"Operation"`
	// 文件上传地址，要求地址协议为HTTPS，且URL端口必须为443
	FileUrl *string `json:"FileUrl" name:"FileUrl"`
	// 文件名
	FileName *string `json:"FileName" name:"FileName"`
	// 文件日期
	FileDate *string `json:"FileDate" name:"FileDate"`
}

func (r *UploadFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadFileRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadFileResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 任务ID
		TaskId *int64 `json:"TaskId" name:"TaskId"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadFileResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
