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

package v20201127

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CallDetailItem struct {
	// 数据类型 0 imei 1 qimei 2 qq 3 phone 7:IDFA 8:MD5(imei)
	DataType *uint64 `json:"DataType,omitempty" name:"DataType"`

	// 有效数据量
	ValidAmount *uint64 `json:"ValidAmount,omitempty" name:"ValidAmount"`

	// 调用时间
	Date *string `json:"Date,omitempty" name:"Date"`
}

type CallDetails struct {
	// 符合条件的总条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 调用明细数组
	CallDetailSet []*CallDetailItem `json:"CallDetailSet,omitempty" name:"CallDetailSet"`
}

type CallStatItem struct {
	// 当前统计量的时间段
	Date *string `json:"Date,omitempty" name:"Date"`

	// 当前时间段的调用量
	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`
}

type GeneralStat struct {
	// 今日调用量
	TodayAmount *uint64 `json:"TodayAmount,omitempty" name:"TodayAmount"`

	// 本周调用量
	WeekAmount *uint64 `json:"WeekAmount,omitempty" name:"WeekAmount"`

	// 本月调用量
	MonthAmount *uint64 `json:"MonthAmount,omitempty" name:"MonthAmount"`

	// 总调用量
	TotalAmount *uint64 `json:"TotalAmount,omitempty" name:"TotalAmount"`
}

// Predefined struct for user
type GetTaskDetailRequestParams struct {
	// 任务ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type GetTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *GetTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskDetailResponseParams struct {
	// 画像洞察任务TAG详细数据列表
	TaskDetailDataList []*TaskDetailData `json:"TaskDetailDataList,omitempty" name:"TaskDetailDataList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskDetailResponseParams `json:"Response"`
}

func (r *GetTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskListRequestParams struct {
	// 查询分页页码
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 查询分页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 查询起始时间（13位数字的UNIX时间戳，单位毫秒 ）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间（13位数字的UNIX时间戳，单位毫秒 ）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 查询任务状态 0:默认状态 1:任务正在运行 2:任务运行成功 3:任务运行失败
	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type GetTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 查询分页页码
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 查询分页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 查询起始时间（13位数字的UNIX时间戳，单位毫秒 ）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间（13位数字的UNIX时间戳，单位毫秒 ）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 查询任务状态 0:默认状态 1:任务正在运行 2:任务运行成功 3:任务运行失败
	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

func (r *GetTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskName")
	delete(f, "TaskStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskListResponseParams struct {
	// 任务列表对象
	TaskListData *TaskListData `json:"TaskListData,omitempty" name:"TaskListData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTaskListResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskListResponseParams `json:"Response"`
}

func (r *GetTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LabelDetailData struct {
	// 标签数据对象
	Value *LabelValue `json:"Value,omitempty" name:"Value"`

	// 标签表述，如"汽车资讯"、"游戏#手游"等
	Label *string `json:"Label,omitempty" name:"Label"`
}

type LabelValue struct {
	// 标签覆盖率占比（在整个上传的ID列表中的覆盖率）
	Proportion *float64 `json:"Proportion,omitempty" name:"Proportion"`

	// 标签大盘覆盖率占比
	Market *float64 `json:"Market,omitempty" name:"Market"`

	// TGI指数，由Proportion除以Market得到
	Tgi *float64 `json:"Tgi,omitempty" name:"Tgi"`
}

type ListModel struct {
	// 任务ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务起始时间（13位数字的UNIX 时间戳，单位毫秒 ）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 任务状态 0:默认状态 1:任务正在运行 2:任务运行成功 3:任务运行失败
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 画像覆盖人数
	Available *uint64 `json:"Available,omitempty" name:"Available"`

	// 任务失败描述信息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

// Predefined struct for user
type PredictRatingRequestParams struct {
	// ID标志的类型，0:IMEI 7:IDFA 8:MD5(imei) 100: 手机号明文 101: 手机号md5加密
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 请求唯一标志ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type PredictRatingRequest struct {
	*tchttp.BaseRequest
	
	// ID标志的类型，0:IMEI 7:IDFA 8:MD5(imei) 100: 手机号明文 101: 手机号md5加密
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 请求唯一标志ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *PredictRatingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PredictRatingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PredictRatingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PredictRatingResponseParams struct {
	// 意向评级
	RatingData *RatingData `json:"RatingData,omitempty" name:"RatingData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PredictRatingResponse struct {
	*tchttp.BaseResponse
	Response *PredictRatingResponseParams `json:"Response"`
}

func (r *PredictRatingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PredictRatingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCallDetailsRequestParams struct {
	// 请求类型 1:人群特征洞察统计 2:购车意向预测统计
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 开始时间戳（毫秒）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳（毫秒）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页个数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

type QueryCallDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 请求类型 1:人群特征洞察统计 2:购车意向预测统计
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 开始时间戳（毫秒）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳（毫秒）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 页数
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页个数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *QueryCallDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCallDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCallDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCallDetailsResponseParams struct {
	// 调用明细
	CallDetails *CallDetails `json:"CallDetails,omitempty" name:"CallDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryCallDetailsResponse struct {
	*tchttp.BaseResponse
	Response *QueryCallDetailsResponseParams `json:"Response"`
}

func (r *QueryCallDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCallDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCallStatRequestParams struct {
	// 请求类型 1:人群特征洞察统计 2:购车意向预测统计
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 开始时间戳（毫秒）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳（毫秒）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

type QueryCallStatRequest struct {
	*tchttp.BaseRequest
	
	// 请求类型 1:人群特征洞察统计 2:购车意向预测统计
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 开始时间戳（毫秒）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳（毫秒）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *QueryCallStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCallStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCallStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCallStatResponseParams struct {
	// 调用量数组
	CallSet []*CallStatItem `json:"CallSet,omitempty" name:"CallSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryCallStatResponse struct {
	*tchttp.BaseResponse
	Response *QueryCallStatResponseParams `json:"Response"`
}

func (r *QueryCallStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCallStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryGeneralStatRequestParams struct {
	// 请求类型:1,人群特征洞察统计 2购车意向预测统计
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type QueryGeneralStatRequest struct {
	*tchttp.BaseRequest
	
	// 请求类型:1,人群特征洞察统计 2购车意向预测统计
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *QueryGeneralStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryGeneralStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryGeneralStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryGeneralStatResponseParams struct {
	// 调用量信息
	GeneralStat *GeneralStat `json:"GeneralStat,omitempty" name:"GeneralStat"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryGeneralStatResponse struct {
	*tchttp.BaseResponse
	Response *QueryGeneralStatResponseParams `json:"Response"`
}

func (r *QueryGeneralStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryGeneralStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RatingData struct {
	// 线索评级（取值：0、1、2、3分别代表无、低、中、高意愿）
	Rank *int64 `json:"Rank,omitempty" name:"Rank"`
}

type TaskData struct {
	// 画像洞察任务ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type TaskDetailData struct {
	// 画像TAG ID
	TagId *uint64 `json:"TagId,omitempty" name:"TagId"`

	// 画像TAG描述（如“省份分布”）
	TagDesc *string `json:"TagDesc,omitempty" name:"TagDesc"`

	// 画像Label对象列表（一个TAG对于N个Label，例如“省份分布”TAG对应“广东省”、“浙江省”等多个Label）
	LabelDetailDataList []*LabelDetailData `json:"LabelDetailDataList,omitempty" name:"LabelDetailDataList"`
}

type TaskListData struct {
	// 查询分页页码
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 查询分页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 任务列表总记录数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务列表
	TaskList []*ListModel `json:"TaskList,omitempty" name:"TaskList"`
}

// Predefined struct for user
type UploadIdRequestParams struct {
	// id标志的类型: 0:imei 7:IDFA 8:MD5(imei)
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// ID列表（ID间使用换行符分割、然后使用Base64编码）
	IdListBase64 *string `json:"IdListBase64,omitempty" name:"IdListBase64"`
}

type UploadIdRequest struct {
	*tchttp.BaseRequest
	
	// id标志的类型: 0:imei 7:IDFA 8:MD5(imei)
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// ID列表（ID间使用换行符分割、然后使用Base64编码）
	IdListBase64 *string `json:"IdListBase64,omitempty" name:"IdListBase64"`
}

func (r *UploadIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "TaskName")
	delete(f, "IdListBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadIdResponseParams struct {
	// 画像洞察任务ID等信息
	TaskData *TaskData `json:"TaskData,omitempty" name:"TaskData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadIdResponse struct {
	*tchttp.BaseResponse
	Response *UploadIdResponseParams `json:"Response"`
}

func (r *UploadIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}