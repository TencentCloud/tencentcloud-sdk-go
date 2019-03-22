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

package v20180416

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Code struct {

	// 对象存储桶名称
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// 对象存储对象路径
	CosObjectName *string `json:"CosObjectName,omitempty" name:"CosObjectName"`

	// 包含函数代码文件及其依赖项的 zip 格式文件，使用该接口时要求将 zip 文件的内容转成 base64 编码，最大支持20M
	ZipFile *string `json:"ZipFile,omitempty" name:"ZipFile"`

	// 对象存储的地域，地域为北京时需要传入ap-beijing,北京一区时需要传递ap-beijing-1，其他的地域不需要传递。
	CosBucketRegion *string `json:"CosBucketRegion,omitempty" name:"CosBucketRegion"`

	// 如果是通过Demo创建的话，需要传入DemoId
	DemoId *string `json:"DemoId,omitempty" name:"DemoId"`

	// 如果是从TempCos创建的话，需要传入TempCosObjectName
	TempCosObjectName *string `json:"TempCosObjectName,omitempty" name:"TempCosObjectName"`
}

type CopyFunctionRequest struct {
	*tchttp.BaseRequest

	// 函数名
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// 新函数的名称
	NewFunctionName *string `json:"NewFunctionName,omitempty" name:"NewFunctionName"`

	// 命名空间，默认为default
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 将函数复制到的命名空间，默认为default
	TargetNamespace *string `json:"TargetNamespace,omitempty" name:"TargetNamespace"`

	// 函数描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 要将函数复制到的地域，不填则默认为当前地域
	TargetRegion *string `json:"TargetRegion,omitempty" name:"TargetRegion"`
}

func (r *CopyFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyFunctionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CopyFunctionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CopyFunctionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFunctionRequest struct {
	*tchttp.BaseRequest

	// 创建的函数名称，函数名称支持26个英文字母大小写、数字、连接符和下划线，第一个字符只能以字母开头，最后一个字符不能为连接符或者下划线，名称长度2-60
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// 函数的代码. 注意：不能同时指定Cos与ZipFile
	Code *Code `json:"Code,omitempty" name:"Code"`

	// 函数处理方法名称，名称格式支持 "文件名称.方法名称" 形式，文件名称和函数名称之间以"."隔开，文件名称和函数名称要求以字母开始和结尾，中间允许插入字母、数字、下划线和连接符，文件名称和函数名字的长度要求是 2-60 个字符
	Handler *string `json:"Handler,omitempty" name:"Handler"`

	// 函数描述,最大支持 1000 个英文字母、数字、空格、逗号、换行符和英文句号，支持中文
	Description *string `json:"Description,omitempty" name:"Description"`

	// 函数运行时内存大小，默认为 128M，可选范围 128MB-1536MB，并且以 128MB 为阶梯
	MemorySize *int64 `json:"MemorySize,omitempty" name:"MemorySize"`

	// 函数最长执行时间，单位为秒，可选值范围 1-300 秒，默认为 3 秒
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// 函数的环境变量
	Environment *Environment `json:"Environment,omitempty" name:"Environment"`

	// 函数运行环境，目前仅支持 Python2.7，Python3.6，Nodejs6.10， PHP5， PHP7，Golang1 和 Java8，默认Python2.7
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 函数的私有网络配置
	VpcConfig *VpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`

	// 函数日志投递到的CLS LogsetID
	ClsLogsetId *string `json:"ClsLogsetId,omitempty" name:"ClsLogsetId"`

	// 函数日志投递到的CLS TopicID
	ClsTopicId *string `json:"ClsTopicId,omitempty" name:"ClsTopicId"`
}

func (r *CreateFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFunctionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFunctionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFunctionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTriggerRequest struct {
	*tchttp.BaseRequest

	// 新建触发器绑定的函数名称
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// 新建触发器名称。如果是定时触发器，名称支持英文字母、数字、连接符和下划线，最长100个字符；如果是其他触发器，见具体触发器绑定参数的说明
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 触发器类型，目前支持 cos 、cmq、 timer、 ckafka类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 触发器对应的参数，如果是 timer 类型的触发器其内容是 Linux cron 表达式，如果是其他触发器，见具体触发器说明
	TriggerDesc *string `json:"TriggerDesc,omitempty" name:"TriggerDesc"`

	// 函数的版本
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`

	// 触发器的初始是能状态 OPEN表示开启 CLOSE表示关闭
	Enable *string `json:"Enable,omitempty" name:"Enable"`
}

func (r *CreateTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTriggerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTriggerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 触发器信息
		TriggerInfo *Trigger `json:"TriggerInfo,omitempty" name:"TriggerInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTriggerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFunctionRequest struct {
	*tchttp.BaseRequest

	// 要删除的函数名称
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`
}

func (r *DeleteFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFunctionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFunctionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFunctionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTriggerRequest struct {
	*tchttp.BaseRequest

	// 函数的名称
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// 要删除的触发器名称
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 要删除的触发器类型，目前支持 cos 、cmq、 timer、ckafka 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 如果删除的触发器类型为 COS 触发器，该字段为必填值，存放 JSON 格式的数据 {"event":"cos:ObjectCreated:*"}，数据内容和 SetTrigger 接口中该字段的格式相同；如果删除的触发器类型为定时触发器或 CMQ 触发器，可以不指定该字段
	TriggerDesc *string `json:"TriggerDesc,omitempty" name:"TriggerDesc"`

	// 函数的版本信息
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`
}

func (r *DeleteTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTriggerRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTriggerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTriggerResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Environment struct {

	// 环境变量数组
	Variables []*Variable `json:"Variables,omitempty" name:"Variables" list`
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type Function struct {

	// 修改时间
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

	// 创建时间
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 运行时
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 函数名称
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// 函数ID
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 函数状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 函数状态详情
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 函数描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 函数标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

type FunctionLog struct {

	// 函数的名称
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// 函数执行完成后的返回值
	RetMsg *string `json:"RetMsg,omitempty" name:"RetMsg"`

	// 执行该函数对应的requestId
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

	// 函数开始执行时的时间点
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 函数执行结果，如果是 0 表示执行成功，其他值表示失败
	RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`

	// 函数调用是否结束，如果是 1 表示执行结束，其他值表示调用异常
	InvokeFinished *int64 `json:"InvokeFinished,omitempty" name:"InvokeFinished"`

	// 函数执行耗时，单位为 ms
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 函数计费时间，根据 duration 向上取最近的 100ms，单位为ms
	BillDuration *int64 `json:"BillDuration,omitempty" name:"BillDuration"`

	// 函数执行时消耗实际内存大小，单位为 Byte
	MemUsage *int64 `json:"MemUsage,omitempty" name:"MemUsage"`

	// 函数执行过程中的日志输出
	Log *string `json:"Log,omitempty" name:"Log"`
}

type GetFunctionLogsRequest struct {
	*tchttp.BaseRequest

	// 函数的名称
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// 数据的偏移量，Offset+Limit不能大于10000
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数据的长度，Offset+Limit不能大于10000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 以升序还是降序的方式对日志进行排序，可选值 desc和 asc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 根据某个字段排序日志,支持以下字段：function_name, duration, mem_usage, start_time
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 日志过滤条件。可用来区分正确和错误日志，filter.retCode=not0 表示只返回错误日志，filter.retCode=is0 表示只返回正确日志，不传，则返回所有日志
	Filter *LogFilter `json:"Filter,omitempty" name:"Filter"`

	// 函数的版本
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`

	// 执行该函数对应的requestId
	FunctionRequestId *string `json:"FunctionRequestId,omitempty" name:"FunctionRequestId"`

	// 查询的具体日期，例如：2017-05-16 20:00:00，只能与endtime相差一天之内
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询的具体日期，例如：2017-05-16 20:59:59，只能与startTime相差一天之内
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetFunctionLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFunctionLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFunctionLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 函数日志的总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 函数日志信息
		Data []*FunctionLog `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFunctionLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFunctionLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFunctionRequest struct {
	*tchttp.BaseRequest

	// 需要获取详情的函数名称
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// 函数的版本号
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`

	// 是否显示代码, TRUE表示显示代码，FALSE表示不显示代码,大于1M的入口文件不会显示
	ShowCode *string `json:"ShowCode,omitempty" name:"ShowCode"`
}

func (r *GetFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFunctionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFunctionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 函数的最后修改时间
		ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

		// 函数的代码
		CodeInfo *string `json:"CodeInfo,omitempty" name:"CodeInfo"`

		// 函数的描述信息
		Description *string `json:"Description,omitempty" name:"Description"`

		// 函数的触发器列表
		Triggers []*Trigger `json:"Triggers,omitempty" name:"Triggers" list`

		// 函数的入口
		Handler *string `json:"Handler,omitempty" name:"Handler"`

		// 函数代码大小
		CodeSize *int64 `json:"CodeSize,omitempty" name:"CodeSize"`

		// 函数的超时时间
		Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

		// 函数的版本
		FunctionVersion *string `json:"FunctionVersion,omitempty" name:"FunctionVersion"`

		// 函数的最大可用内存
		MemorySize *int64 `json:"MemorySize,omitempty" name:"MemorySize"`

		// 函数的运行环境
		Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

		// 函数的名称
		FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

		// 函数的私有网络
		VpcConfig *VpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`

		// 是否使用GPU
		UseGpu *string `json:"UseGpu,omitempty" name:"UseGpu"`

		// 函数的环境变量
		Environment *Environment `json:"Environment,omitempty" name:"Environment"`

		// 代码是否正确
		CodeResult *string `json:"CodeResult,omitempty" name:"CodeResult"`

		// 代码错误信息
		CodeError *string `json:"CodeError,omitempty" name:"CodeError"`

		// 代码错误码
		ErrNo *int64 `json:"ErrNo,omitempty" name:"ErrNo"`

		// 函数的命名空间
		Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

		// 函数绑定的角色
		Role *string `json:"Role,omitempty" name:"Role"`

		// 是否自动安装依赖
		InstallDependency *string `json:"InstallDependency,omitempty" name:"InstallDependency"`

		// 函数状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 状态描述
		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

		// 日志投递到的Cls日志集
		ClsLogsetId *string `json:"ClsLogsetId,omitempty" name:"ClsLogsetId"`

		// 日志投递到的Cls Topic
		ClsTopicId *string `json:"ClsTopicId,omitempty" name:"ClsTopicId"`

		// 函数ID
		FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

		// 函数的标签列表
		Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFunctionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InvokeRequest struct {
	*tchttp.BaseRequest

	// 函数名称
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// RequestResponse(同步) 和 Event(异步)，默认为同步
	InvocationType *string `json:"InvocationType,omitempty" name:"InvocationType"`

	// 触发函数的版本号
	Qualifier *string `json:"Qualifier,omitempty" name:"Qualifier"`

	// 运行函数时的参数，以json格式传入，最大支持的参数长度是 1M
	ClientContext *string `json:"ClientContext,omitempty" name:"ClientContext"`

	// 同步调用时指定该字段，返回值会包含4K的日志，可选值为None和Tail，默认值为None。当该值为Tail时，返回参数中的logMsg字段会包含对应的函数执行日志
	LogType *string `json:"LogType,omitempty" name:"LogType"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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
		Result *Result `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InvokeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InvokeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListFunctionsRequest struct {
	*tchttp.BaseRequest

	// 以升序还是降序的方式返回结果，可选值 ASC 和 DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// 根据哪个字段进行返回结果排序,支持以下字段：AddTime, ModTime, FunctionName
	Orderby *string `json:"Orderby,omitempty" name:"Orderby"`

	// 数据偏移量，默认值为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数据长度，默认值为 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 支持FunctionName模糊匹配
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 函数描述，支持模糊搜索
	Description *string `json:"Description,omitempty" name:"Description"`

	// 过滤条件。
	// - tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *ListFunctionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListFunctionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListFunctionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 函数列表
		Functions []*Function `json:"Functions,omitempty" name:"Functions" list`

		// 总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListFunctionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListFunctionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LogFilter struct {

	// filter.RetCode=not0 表示只返回错误日志，filter.RetCode=is0 表示只返回正确日志，无输入则返回所有日志。
	RetCode *string `json:"RetCode,omitempty" name:"RetCode"`
}

type Result struct {

	// 表示执行过程中的日志输出，异步调用返回为空
	Log *string `json:"Log,omitempty" name:"Log"`

	// 表示执行函数的返回，异步调用返回为空
	RetMsg *string `json:"RetMsg,omitempty" name:"RetMsg"`

	// 表示执行函数的错误返回信息，异步调用返回为空
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 执行函数时的内存大小，单位为Byte，异步调用返回为空
	MemUsage *int64 `json:"MemUsage,omitempty" name:"MemUsage"`

	// 表示执行函数的耗时，单位是毫秒，异步调用返回为空
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 表示函数的计费耗时，单位是毫秒，异步调用返回为空
	BillDuration *int64 `json:"BillDuration,omitempty" name:"BillDuration"`

	// 此次函数执行的Id
	FunctionRequestId *string `json:"FunctionRequestId,omitempty" name:"FunctionRequestId"`

	// 0为正确，异步调用返回为空
	InvokeResult *int64 `json:"InvokeResult,omitempty" name:"InvokeResult"`
}

type Tag struct {

	// 标签的key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签的value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Trigger struct {

	// 触发器最后修改时间
	ModTime *string `json:"ModTime,omitempty" name:"ModTime"`

	// 触发器类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 触发器详细配置
	TriggerDesc *string `json:"TriggerDesc,omitempty" name:"TriggerDesc"`

	// 触发器名称
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 触发器创建时间
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 使能开关
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 客户自定义参数
	CustomArgument *string `json:"CustomArgument,omitempty" name:"CustomArgument"`
}

type UpdateFunctionCodeRequest struct {
	*tchttp.BaseRequest

	// 函数处理方法名称。名称格式支持“文件名称.函数名称”形式，文件名称和函数名称之间以"."隔开，文件名称和函数名称要求以字母开始和结尾，中间允许插入字母、数字、下划线和连接符，文件名称和函数名字的长度要求 2-60 个字符
	Handler *string `json:"Handler,omitempty" name:"Handler"`

	// 要修改的函数名称
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// 对象存储桶名称
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// 对象存储对象路径
	CosObjectName *string `json:"CosObjectName,omitempty" name:"CosObjectName"`

	// 包含函数代码文件及其依赖项的 zip 格式文件，使用该接口时要求将 zip 文件的内容转成 base64 编码，最大支持20M
	ZipFile *string `json:"ZipFile,omitempty" name:"ZipFile"`

	// 对象存储的地域，注：北京分为ap-beijing和ap-beijing-1
	CosBucketRegion *string `json:"CosBucketRegion,omitempty" name:"CosBucketRegion"`
}

func (r *UpdateFunctionCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFunctionCodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFunctionCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFunctionCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFunctionCodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFunctionConfigurationRequest struct {
	*tchttp.BaseRequest

	// 要修改的函数名称
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// 函数描述。最大支持 1000 个英文字母、数字、空格、逗号和英文句号，支持中文
	Description *string `json:"Description,omitempty" name:"Description"`

	// 函数运行时内存大小，默认为 128 M，可选范 128 M-1536 M
	MemorySize *int64 `json:"MemorySize,omitempty" name:"MemorySize"`

	// 函数最长执行时间，单位为秒，可选值范 1-300 秒，默认为 3 秒
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// 函数运行环境，目前仅支持 Python2.7，Python3.6，Nodejs6.10，PHP5， PHP7，Golang1 和 Java8
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 函数的环境变量
	Environment *Environment `json:"Environment,omitempty" name:"Environment"`

	// 函数的私有网络配置
	VpcConfig *VpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`
}

func (r *UpdateFunctionConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFunctionConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFunctionConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFunctionConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFunctionConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Variable struct {

	// 变量的名称
	Key *string `json:"Key,omitempty" name:"Key"`

	// 变量的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type VpcConfig struct {

	// 私有网络 的 id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网的 id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
