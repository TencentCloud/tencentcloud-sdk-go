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

type ApplyBlackListRequest struct {
	*tchttp.BaseRequest

	// 模块
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 黑名单列表
	BlackList []*SingleBlackApply `json:"BlackList,omitempty" name:"BlackList" list`

	// 实例ID，不传默认为系统分配的初始实例
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

func (r *ApplyBlackListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyBlackListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApplyBlackListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyBlackListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyBlackListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApplyCreditAuditRequest struct {
	*tchttp.BaseRequest

	// 模块
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 实例ID
	InstId *string `json:"InstId,omitempty" name:"InstId"`

	// 产品ID，形如P******。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 信审任务ID，同一天内，同一InstId下，同一CaseId只能调用一次。
	CaseId *string `json:"CaseId,omitempty" name:"CaseId"`

	// 回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// JSON格式的业务字段。
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *ApplyCreditAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyCreditAuditRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApplyCreditAuditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求日期
		RequestDate *string `json:"RequestDate,omitempty" name:"RequestDate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyCreditAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyCreditAuditResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCreditResultRequest struct {
	*tchttp.BaseRequest

	// 模块
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 实例ID
	InstId *string `json:"InstId,omitempty" name:"InstId"`

	// 产品ID，形如P******。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 信审任务ID
	CaseId *string `json:"CaseId,omitempty" name:"CaseId"`

	// 请求日期
	RequestDate *string `json:"RequestDate,omitempty" name:"RequestDate"`
}

func (r *DescribeCreditResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCreditResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCreditResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// <p>呼叫结果，取值范围：</p><ul style="margin-bottom:0px;"><li>NON：接通</li><li>DBU：号码忙</li><li>DRF：不在服务区</li><li>ANA：欠费未接听</li><li>REJ：拒接</li><li>SHU：关机</li><li>NAN：空号</li><li>HAL：停机</li><li>DAD：未接听</li><li>EXE：其他异常</li></ul>
		ResultCode *string `json:"ResultCode,omitempty" name:"ResultCode"`

		// 客户标识代码，多个标识码以英文逗号分隔，ResultCode为NON时才有。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClientCode *string `json:"ClientCode,omitempty" name:"ClientCode"`

		// 开始振铃时间，ResultCode为NON或DAD时才有此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
		RingStartTime *string `json:"RingStartTime,omitempty" name:"RingStartTime"`

		// 振铃时长
		RingDuration *int64 `json:"RingDuration,omitempty" name:"RingDuration"`

		// 接通时长
		AnswerDuration *int64 `json:"AnswerDuration,omitempty" name:"AnswerDuration"`

		// JSON格式的扩展信息字段，ResultCode为NON时才有。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ContextValue *string `json:"ContextValue,omitempty" name:"ContextValue"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCreditResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCreditResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordsRequest struct {
	*tchttp.BaseRequest

	// 模块
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 案件编号
	AccountNum *string `json:"AccountNum,omitempty" name:"AccountNum"`

	// 被叫号码
	CalledPhone *string `json:"CalledPhone,omitempty" name:"CalledPhone"`

	// 查询起始日期
	StartBizDate *string `json:"StartBizDate,omitempty" name:"StartBizDate"`

	// 查询结束日期
	EndBizDate *string `json:"EndBizDate,omitempty" name:"EndBizDate"`

	// 分页参数，索引，从0开始
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 分页参数，页长
	Limit *string `json:"Limit,omitempty" name:"Limit"`

	// 实例ID，不传默认为系统分配的初始实例
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

func (r *DescribeRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRecordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 录音列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		RecordList []*SingleRecord `json:"RecordList,omitempty" name:"RecordList" list`

		// 录音总量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRecordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务ID，形如abc-a0b1c2xyz
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
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

		// 任务结果，例如上传成功时返回“File Uploading Task Success.”
		TaskResult *string `json:"TaskResult,omitempty" name:"TaskResult"`

		// <p>任务类型：</p><ul style="margin-bottom:0px;"><li>报告下载：001</li><li>催收数据上传：002</li><li>还款数据上传：003</li><li>回访数据上传：004</li><li>停拨数据上传：005</li></ul>
		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

		// 过滤文件下载链接，有过滤数据时才存在。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskFileUrl *string `json:"TaskFileUrl,omitempty" name:"TaskFileUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 报告日期
	ReportDate *string `json:"ReportDate,omitempty" name:"ReportDate"`

	// 实例ID，不传默认为系统分配的初始实例。
	InstId *string `json:"InstId,omitempty" name:"InstId"`
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

		// 催收日报下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DailyReportUrl *string `json:"DailyReportUrl,omitempty" name:"DailyReportUrl"`

		// 催收结果下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		ResultReportUrl *string `json:"ResultReportUrl,omitempty" name:"ResultReportUrl"`

		// 催收明细下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		DetailReportUrl *string `json:"DetailReportUrl,omitempty" name:"DetailReportUrl"`

		// 回访日报下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		CallbackDailyReportUrl *string `json:"CallbackDailyReportUrl,omitempty" name:"CallbackDailyReportUrl"`

		// 回访结果下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		CallbackResultReportUrl *string `json:"CallbackResultReportUrl,omitempty" name:"CallbackResultReportUrl"`

		// 回访明细下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		CallbackDetailReportUrl *string `json:"CallbackDetailReportUrl,omitempty" name:"CallbackDetailReportUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadReportResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SingleBlackApply struct {

	// 黑名单类型，01代表手机号码。
	BlackType *string `json:"BlackType,omitempty" name:"BlackType"`

	// 操作类型，A为新增，D为删除。
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 黑名单值，BlackType为01时，填写11位手机号码。
	BlackValue *string `json:"BlackValue,omitempty" name:"BlackValue"`

	// 备注。
	BlackDescription *string `json:"BlackDescription,omitempty" name:"BlackDescription"`

	// 黑名单生效截止日期，格式为YYYY-MM-DD，不填默认为永久。
	BlackValidDate *string `json:"BlackValidDate,omitempty" name:"BlackValidDate"`
}

type SingleRecord struct {

	// 案件编号。
	AccountNum *string `json:"AccountNum,omitempty" name:"AccountNum"`

	// 外呼日期。
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`

	// 开始呼叫时间。
	CallStartTime *string `json:"CallStartTime,omitempty" name:"CallStartTime"`

	// 主叫号码。
	CallerPhone *string `json:"CallerPhone,omitempty" name:"CallerPhone"`

	// 呼叫方向，O为呼出，I为呼入。
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// 通话时长。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 产品ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 录音下载链接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordCosUrl *string `json:"RecordCosUrl,omitempty" name:"RecordCosUrl"`
}

type UploadDataFileRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// <p>上传类型，不填默认催收文件，取值范围：</p><ul style="margin-bottom:0px;"><li>data：催收文件</li><li>repay：还款文件</li><li>callback：回访文件</li></ul>
	UploadModel *string `json:"UploadModel,omitempty" name:"UploadModel"`

	// 文件，文件与文件地址上传只可选用一种，必须使用multipart/form-data协议来上传二进制流文件，建议使用xlsx格式，大小不超过5MB。
	File *binary `json:"File,omitempty" name:"File"`

	// 文件上传地址，文件与文件地址上传只可选用一种，大小不超过50MB。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 实例ID，不传默认为系统分配的初始实例。
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

func (r *UploadDataFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadDataFileRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadDataFileResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据ID
		DataResId *string `json:"DataResId,omitempty" name:"DataResId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadDataFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadDataFileResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadFileRequest struct {
	*tchttp.BaseRequest

	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 文件上传地址，要求地址协议为HTTPS，且URL端口必须为443
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件日期
	FileDate *string `json:"FileDate,omitempty" name:"FileDate"`
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
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadFileResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
