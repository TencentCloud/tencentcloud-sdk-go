// Copyright 1999-2018 Tencent Ltd.
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

package v20180416

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type InvokeRequest struct {
	*tchttp.BaseRequest
	// 函数名称。
	FunctionName *string `json:"FunctionName" name:"FunctionName"`
	// RequestResponse(同步) 和 Event(异步)，默认为同步。
	InvocationType *string `json:"InvocationType" name:"InvocationType"`
	// 触发函数的版本号。
	Qualifier *string `json:"Qualifier" name:"Qualifier"`
	// 运行函数时的参数，以json格式传入，最大支持的参数长度是 1M。
	ClientContext *string `json:"ClientContext" name:"ClientContext"`
	// 同步调用时指定该字段，返回值会包含4K的日志，可选值为None和Tail，默认值为None。当该值为Tail时，返回参数中的logMsg字段会包含对应的函数执行日志。
	LogType *string `json:"LogType" name:"LogType"`
}

func (r *InvokeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InvokeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InvokeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 函数执行结果
		Result *Result `json:"Result" name:"Result"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *InvokeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InvokeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Result struct {
	// 表示执行过程中的日志输出，异步调用返回为空
	Log *string `json:"Log" name:"Log"`
	// 表示执行函数的返回，异步调用返回为空
	RetMsg *string `json:"RetMsg" name:"RetMsg"`
	// 表示执行函数的错误返回信息，异步调用返回为空
	ErrMsg *string `json:"ErrMsg" name:"ErrMsg"`
	// 执行函数时的内存大小，单位为Byte，异步调用返回为空
	MemUsage *int64 `json:"MemUsage" name:"MemUsage"`
	// 表示执行函数的耗时，单位是毫秒，异步调用返回为空
	Duration *float64 `json:"Duration" name:"Duration"`
	// 表示函数的计费耗时，单位是毫秒，异步调用返回为空
	BillDuration *int64 `json:"BillDuration" name:"BillDuration"`
	// 此次函数执行的Id
	FunctionRequestId *string `json:"FunctionRequestId" name:"FunctionRequestId"`
	// 0为正确，异步调用返回为空
	InvokeResult *int64 `json:"InvokeResult" name:"InvokeResult"`
}
