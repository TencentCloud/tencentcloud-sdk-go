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

package v20180409

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateProbeTasksRequestParams struct {
	// 批量任务名-地址
	BatchTasks []*ProbeTaskBasicConfiguration `json:"BatchTasks,omitnil,omitempty" name:"BatchTasks"`

	// 任务类型，如1、2、3、4、5、6、7；1-页面性能、2-文件上传、3-文件下载、4-端口性能、5-网络质量、6-音视频体验、7-域名whois
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 拨测节点，如10001，具体拨测地域运营商对应的拨测点编号可联系云拨测确认。
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 拨测间隔
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 拨测参数，详细可参考云拨测官方文档,链接:https://cloud.tencent.com/document/product/248/87308#createprobetasks。
	Parameters *string `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// 任务分类
	// <li>1 = PC</li>
	// <li> 2 = Mobile </li>
	TaskCategory *int64 `json:"TaskCategory,omitnil,omitempty" name:"TaskCategory"`

	// 定时任务cron表达式
	Cron *string `json:"Cron,omitnil,omitempty" name:"Cron"`

	// 资源标签值
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 测试类型，包含定时测试与即时测试。0-定时拨测，其它表示即时拨测。
	ProbeType *uint64 `json:"ProbeType,omitnil,omitempty" name:"ProbeType"`

	// 插件类型，如CDN，详情参考云拨测官方文档。
	PluginSource *string `json:"PluginSource,omitnil,omitempty" name:"PluginSource"`

	// 客户端ID
	ClientNum *string `json:"ClientNum,omitnil,omitempty" name:"ClientNum"`

	// 拨测点IP类型：0-不限制IP类型，1-IPv4，2-IPv6
	NodeIpType *int64 `json:"NodeIpType,omitnil,omitempty" name:"NodeIpType"`

	// 供应商子账户同步标志
	SubSyncFlag *int64 `json:"SubSyncFlag,omitnil,omitempty" name:"SubSyncFlag"`
}

type CreateProbeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 批量任务名-地址
	BatchTasks []*ProbeTaskBasicConfiguration `json:"BatchTasks,omitnil,omitempty" name:"BatchTasks"`

	// 任务类型，如1、2、3、4、5、6、7；1-页面性能、2-文件上传、3-文件下载、4-端口性能、5-网络质量、6-音视频体验、7-域名whois
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 拨测节点，如10001，具体拨测地域运营商对应的拨测点编号可联系云拨测确认。
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 拨测间隔
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 拨测参数，详细可参考云拨测官方文档,链接:https://cloud.tencent.com/document/product/248/87308#createprobetasks。
	Parameters *string `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// 任务分类
	// <li>1 = PC</li>
	// <li> 2 = Mobile </li>
	TaskCategory *int64 `json:"TaskCategory,omitnil,omitempty" name:"TaskCategory"`

	// 定时任务cron表达式
	Cron *string `json:"Cron,omitnil,omitempty" name:"Cron"`

	// 资源标签值
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 测试类型，包含定时测试与即时测试。0-定时拨测，其它表示即时拨测。
	ProbeType *uint64 `json:"ProbeType,omitnil,omitempty" name:"ProbeType"`

	// 插件类型，如CDN，详情参考云拨测官方文档。
	PluginSource *string `json:"PluginSource,omitnil,omitempty" name:"PluginSource"`

	// 客户端ID
	ClientNum *string `json:"ClientNum,omitnil,omitempty" name:"ClientNum"`

	// 拨测点IP类型：0-不限制IP类型，1-IPv4，2-IPv6
	NodeIpType *int64 `json:"NodeIpType,omitnil,omitempty" name:"NodeIpType"`

	// 供应商子账户同步标志
	SubSyncFlag *int64 `json:"SubSyncFlag,omitnil,omitempty" name:"SubSyncFlag"`
}

func (r *CreateProbeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProbeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchTasks")
	delete(f, "TaskType")
	delete(f, "Nodes")
	delete(f, "Interval")
	delete(f, "Parameters")
	delete(f, "TaskCategory")
	delete(f, "Cron")
	delete(f, "Tag")
	delete(f, "ProbeType")
	delete(f, "PluginSource")
	delete(f, "ClientNum")
	delete(f, "NodeIpType")
	delete(f, "SubSyncFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProbeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProbeTasksResponseParams struct {
	// 任务ID列表
	TaskIDs []*string `json:"TaskIDs,omitnil,omitempty" name:"TaskIDs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProbeTasksResponse struct {
	*tchttp.BaseResponse
	Response *CreateProbeTasksResponseParams `json:"Response"`
}

func (r *CreateProbeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProbeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProbeTaskRequestParams struct {
	// 任务 ID
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

type DeleteProbeTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

func (r *DeleteProbeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProbeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProbeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProbeTaskResponseParams struct {
	// 任务总量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 任务成功量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 任务执行结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*TaskResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProbeTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProbeTaskResponseParams `json:"Response"`
}

func (r *DeleteProbeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProbeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDetailedSingleProbeDataRequestParams struct {
	// 开始时间戳（毫秒级）
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间戳（毫秒级）
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务类型
	// AnalyzeTaskType_Network：网络质量
	// AnalyzeTaskType_Browse：页面性能
	// AnalyzeTaskType_UploadDownload：文件传输（含文件上传、文件下载）
	// AnalyzeTaskType_Transport：端口性能
	// AnalyzeTaskType_MediaStream：音视频体验
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 待排序字段
	// 可以填写 ProbeTime 拨测时间排序
	// 也可填写SelectedFields 中的选中字段
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// true表示升序
	Ascending *bool `json:"Ascending,omitnil,omitempty" name:"Ascending"`

	// 选中字段，如ProbeTime、TransferTime、TransferSize等。
	SelectedFields []*string `json:"SelectedFields,omitnil,omitempty" name:"SelectedFields"`

	// 起始取数位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 取数数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 任务ID
	TaskID []*string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 拨测点运营商
	// 	
	// 这里实际按拨测结果中的运营商来填写即可
	// 
	// 电信：中国电信
	// 移动：中国移动
	// 联通：中国联通
	Operators []*string `json:"Operators,omitnil,omitempty" name:"Operators"`

	// 拨测点地区
	// 	
	// 这里实际按拨测结果中的地区来填写即可
	// 
	// 国内一般是省级单位，如广东、广西、中国香港；直辖市则填北京、上海
	// 
	// 境外一般是国家名，如澳大利亚、新加坡
	Districts []*string `json:"Districts,omitnil,omitempty" name:"Districts"`

	// 错误类型
	ErrorTypes []*string `json:"ErrorTypes,omitnil,omitempty" name:"ErrorTypes"`

	// 城市
	// 这里实际按拨测结果中的城市来填写即可
	// 
	// 示例：
	// 
	// 深圳市
	// 武汉市
	// 首尔
	// 多伦多
	City []*string `json:"City,omitnil,omitempty" name:"City"`

	// es scroll查询id
	ScrollID *string `json:"ScrollID,omitnil,omitempty" name:"ScrollID"`

	// 详情数据下载
	QueryFlag *string `json:"QueryFlag,omitnil,omitempty" name:"QueryFlag"`
}

type DescribeDetailedSingleProbeDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间戳（毫秒级）
	BeginTime *uint64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间戳（毫秒级）
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务类型
	// AnalyzeTaskType_Network：网络质量
	// AnalyzeTaskType_Browse：页面性能
	// AnalyzeTaskType_UploadDownload：文件传输（含文件上传、文件下载）
	// AnalyzeTaskType_Transport：端口性能
	// AnalyzeTaskType_MediaStream：音视频体验
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 待排序字段
	// 可以填写 ProbeTime 拨测时间排序
	// 也可填写SelectedFields 中的选中字段
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// true表示升序
	Ascending *bool `json:"Ascending,omitnil,omitempty" name:"Ascending"`

	// 选中字段，如ProbeTime、TransferTime、TransferSize等。
	SelectedFields []*string `json:"SelectedFields,omitnil,omitempty" name:"SelectedFields"`

	// 起始取数位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 取数数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 任务ID
	TaskID []*string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 拨测点运营商
	// 	
	// 这里实际按拨测结果中的运营商来填写即可
	// 
	// 电信：中国电信
	// 移动：中国移动
	// 联通：中国联通
	Operators []*string `json:"Operators,omitnil,omitempty" name:"Operators"`

	// 拨测点地区
	// 	
	// 这里实际按拨测结果中的地区来填写即可
	// 
	// 国内一般是省级单位，如广东、广西、中国香港；直辖市则填北京、上海
	// 
	// 境外一般是国家名，如澳大利亚、新加坡
	Districts []*string `json:"Districts,omitnil,omitempty" name:"Districts"`

	// 错误类型
	ErrorTypes []*string `json:"ErrorTypes,omitnil,omitempty" name:"ErrorTypes"`

	// 城市
	// 这里实际按拨测结果中的城市来填写即可
	// 
	// 示例：
	// 
	// 深圳市
	// 武汉市
	// 首尔
	// 多伦多
	City []*string `json:"City,omitnil,omitempty" name:"City"`

	// es scroll查询id
	ScrollID *string `json:"ScrollID,omitnil,omitempty" name:"ScrollID"`

	// 详情数据下载
	QueryFlag *string `json:"QueryFlag,omitnil,omitempty" name:"QueryFlag"`
}

func (r *DescribeDetailedSingleProbeDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetailedSingleProbeDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "TaskType")
	delete(f, "SortField")
	delete(f, "Ascending")
	delete(f, "SelectedFields")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TaskID")
	delete(f, "Operators")
	delete(f, "Districts")
	delete(f, "ErrorTypes")
	delete(f, "City")
	delete(f, "ScrollID")
	delete(f, "QueryFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDetailedSingleProbeDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDetailedSingleProbeDataResponseParams struct {
	// 单次详情数据
	DataSet []*DetailedSingleDataDefine `json:"DataSet,omitnil,omitempty" name:"DataSet"`

	// 符合条件的数据总数
	TotalNumber *int64 `json:"TotalNumber,omitnil,omitempty" name:"TotalNumber"`

	// es scroll查询的id
	ScrollID *string `json:"ScrollID,omitnil,omitempty" name:"ScrollID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDetailedSingleProbeDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDetailedSingleProbeDataResponseParams `json:"Response"`
}

func (r *DescribeDetailedSingleProbeDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetailedSingleProbeDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstantTasksRequestParams struct {
	// 数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 起始位置
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeInstantTasksRequest struct {
	*tchttp.BaseRequest
	
	// 数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 起始位置
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeInstantTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstantTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstantTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstantTasksResponseParams struct {
	// 任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*SingleInstantTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstantTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstantTasksResponseParams `json:"Response"`
}

func (r *DescribeInstantTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstantTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodesRequestParams struct {
	// 节点类型
	// <li> 1 = IDC </li>
	// <li> 2 = LastMile </li>
	// <li> 3 = Mobile </li>
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点区域
	// <li> 1 = 中国大陆 </li>
	// <li> 2 = 港澳台 </li>
	// <li> 3 = 境外</li>
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`

	// 是否IPv6
	IsIPv6 *bool `json:"IsIPv6,omitnil,omitempty" name:"IsIPv6"`

	// 名字模糊搜索
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 付费模式
	// <li>1 = 试用版本</li>
	// <li> 2 = 付费版本 </li>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 任务类型
	// <li>1 = 页面性能</li>
	// <li>2 = 文件上传</li>
	// <li>3 = 文件下载</li>
	// <li>4 = 端口性能</li>
	// <li>5 = 网络质量</li>
	// <li>6 = 音视频体验</li>
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type DescribeNodesRequest struct {
	*tchttp.BaseRequest
	
	// 节点类型
	// <li> 1 = IDC </li>
	// <li> 2 = LastMile </li>
	// <li> 3 = Mobile </li>
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点区域
	// <li> 1 = 中国大陆 </li>
	// <li> 2 = 港澳台 </li>
	// <li> 3 = 境外</li>
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`

	// 是否IPv6
	IsIPv6 *bool `json:"IsIPv6,omitnil,omitempty" name:"IsIPv6"`

	// 名字模糊搜索
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 付费模式
	// <li>1 = 试用版本</li>
	// <li> 2 = 付费版本 </li>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 任务类型
	// <li>1 = 页面性能</li>
	// <li>2 = 文件上传</li>
	// <li>3 = 文件下载</li>
	// <li>4 = 端口性能</li>
	// <li>5 = 网络质量</li>
	// <li>6 = 音视频体验</li>
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

func (r *DescribeNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeType")
	delete(f, "Location")
	delete(f, "IsIPv6")
	delete(f, "NodeName")
	delete(f, "PayMode")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodesResponseParams struct {
	// 节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeSet []*NodeDefineExt `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodesResponseParams `json:"Response"`
}

func (r *DescribeNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProbeMetricDataRequestParams struct {
	// 分析任务类型，支持以下几种类型：
	// AnalyzeTaskType_Network：网络质量
	// AnalyzeTaskType_Browse：页面性能
	// AnalyzeTaskType_Transport：端口性能
	// AnalyzeTaskType_UploadDownload：文件传输
	// AnalyzeTaskType_MediaStream：音视频体验
	AnalyzeTaskType *string `json:"AnalyzeTaskType,omitnil,omitempty" name:"AnalyzeTaskType"`

	// 指标类型（counter、gauge以及histogram），指标查询默认传gauge
	MetricType *string `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// 指标详细字段，可以传递传具体的指标也可以对指标进行聚合查询例如："avg(ping_time)"代表整体时延(ms)；不同的任务类型支持不同的field查询，以及聚合规则，详情可见https://cloud.tencent.com/document/product/248/87584。
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 过滤条件可以传单个过滤条件也可以拼接多个参数
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 聚合时间, 1m、1d、30d 等等
	GroupBy *string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// 多条件过滤，支持多个过滤条件组合查询
	// 例如：[""host" = 'www.test.com'", "time >= now()-1h"]
	Filters []*string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeProbeMetricDataRequest struct {
	*tchttp.BaseRequest
	
	// 分析任务类型，支持以下几种类型：
	// AnalyzeTaskType_Network：网络质量
	// AnalyzeTaskType_Browse：页面性能
	// AnalyzeTaskType_Transport：端口性能
	// AnalyzeTaskType_UploadDownload：文件传输
	// AnalyzeTaskType_MediaStream：音视频体验
	AnalyzeTaskType *string `json:"AnalyzeTaskType,omitnil,omitempty" name:"AnalyzeTaskType"`

	// 指标类型（counter、gauge以及histogram），指标查询默认传gauge
	MetricType *string `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// 指标详细字段，可以传递传具体的指标也可以对指标进行聚合查询例如："avg(ping_time)"代表整体时延(ms)；不同的任务类型支持不同的field查询，以及聚合规则，详情可见https://cloud.tencent.com/document/product/248/87584。
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 过滤条件可以传单个过滤条件也可以拼接多个参数
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 聚合时间, 1m、1d、30d 等等
	GroupBy *string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// 多条件过滤，支持多个过滤条件组合查询
	// 例如：[""host" = 'www.test.com'", "time >= now()-1h"]
	Filters []*string `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeProbeMetricDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProbeMetricDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AnalyzeTaskType")
	delete(f, "MetricType")
	delete(f, "Field")
	delete(f, "Filter")
	delete(f, "GroupBy")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProbeMetricDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProbeMetricDataResponseParams struct {
	//  返回指标 JSON 序列化后的字符串，具体如下所示："[{\"name\":\"task_navigate_request_gauge\",\"columns\":[\"time\",\"avg(first_screen_time) / 1000\"],\"values\":[[1641571200,6.756600000000001]],\"tags\":null}]"
	MetricSet *string `json:"MetricSet,omitnil,omitempty" name:"MetricSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProbeMetricDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProbeMetricDataResponseParams `json:"Response"`
}

func (r *DescribeProbeMetricDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProbeMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProbeNodesRequestParams struct {
	// 节点类型
	// <li> 1 = IDC </li>
	// <li> 2 = LastMile </li>
	// <li> 3 = Mobile </li>
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点区域
	// <li> 1 = 中国大陆 </li>
	// <li> 2 = 港澳台 </li>
	// <li> 3 = 海外 </li>
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`

	// 是否IPv6
	IsIPv6 *bool `json:"IsIPv6,omitnil,omitempty" name:"IsIPv6"`

	// 名字模糊搜索
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 付费模式
	// <li>1 = 试用版本</li>
	// <li> 2 = 付费版本 </li>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type DescribeProbeNodesRequest struct {
	*tchttp.BaseRequest
	
	// 节点类型
	// <li> 1 = IDC </li>
	// <li> 2 = LastMile </li>
	// <li> 3 = Mobile </li>
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点区域
	// <li> 1 = 中国大陆 </li>
	// <li> 2 = 港澳台 </li>
	// <li> 3 = 海外 </li>
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`

	// 是否IPv6
	IsIPv6 *bool `json:"IsIPv6,omitnil,omitempty" name:"IsIPv6"`

	// 名字模糊搜索
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 付费模式
	// <li>1 = 试用版本</li>
	// <li> 2 = 付费版本 </li>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

func (r *DescribeProbeNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProbeNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeType")
	delete(f, "Location")
	delete(f, "IsIPv6")
	delete(f, "NodeName")
	delete(f, "PayMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProbeNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProbeNodesResponseParams struct {
	// 节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeSet []*NodeDefine `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProbeNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProbeNodesResponseParams `json:"Response"`
}

func (r *DescribeProbeNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProbeNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProbeTasksRequestParams struct {
	// 任务 ID  列表
	TaskIDs []*string `json:"TaskIDs,omitnil,omitempty" name:"TaskIDs"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 拨测目标
	TargetAddress *string `json:"TargetAddress,omitnil,omitempty" name:"TargetAddress"`

	// 任务状态列表
	// <li>1 = 创建中</li>
	// <li> 2 = 运行中 </li>
	// <li> 3 = 运行异常 </li>
	// <li> 4 = 暂停中 </li>
	// <li> 5 = 暂停异常 </li>
	// <li> 6 = 任务暂停 </li>
	// <li> 7 = 任务删除中 </li>
	// <li> 8 = 任务删除异常 </li>
	// <li> 9 = 任务删除</li>
	// <li> 10 = 定时任务暂停中 </li>
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 付费模式
	// <li>1 = 试用版本</li>
	// <li> 2 = 付费版本 </li>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 订单状态
	// <li>1 = 正常</li>
	// <li> 2 = 欠费 </li>
	OrderState *int64 `json:"OrderState,omitnil,omitempty" name:"OrderState"`

	// 拨测类型
	// <li>1 = 页面浏览</li>
	// <li> 2 =文件上传 </li>
	// <li> 3 = 文件下载</li>
	// <li> 4 = 端口性能 </li>
	// <li> 5 = 网络质量 </li>
	// <li> 6 =流媒体 </li>
	// 
	// 即使拨测只支持页面浏览，网络质量，文件下载
	TaskType []*int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 节点类型
	TaskCategory []*int64 `json:"TaskCategory,omitnil,omitempty" name:"TaskCategory"`

	// 排序的列
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 是否正序
	Ascend *bool `json:"Ascend,omitnil,omitempty" name:"Ascend"`

	// 资源标签值
	TagFilters []*KeyValuePair `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeProbeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID  列表
	TaskIDs []*string `json:"TaskIDs,omitnil,omitempty" name:"TaskIDs"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 拨测目标
	TargetAddress *string `json:"TargetAddress,omitnil,omitempty" name:"TargetAddress"`

	// 任务状态列表
	// <li>1 = 创建中</li>
	// <li> 2 = 运行中 </li>
	// <li> 3 = 运行异常 </li>
	// <li> 4 = 暂停中 </li>
	// <li> 5 = 暂停异常 </li>
	// <li> 6 = 任务暂停 </li>
	// <li> 7 = 任务删除中 </li>
	// <li> 8 = 任务删除异常 </li>
	// <li> 9 = 任务删除</li>
	// <li> 10 = 定时任务暂停中 </li>
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 付费模式
	// <li>1 = 试用版本</li>
	// <li> 2 = 付费版本 </li>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 订单状态
	// <li>1 = 正常</li>
	// <li> 2 = 欠费 </li>
	OrderState *int64 `json:"OrderState,omitnil,omitempty" name:"OrderState"`

	// 拨测类型
	// <li>1 = 页面浏览</li>
	// <li> 2 =文件上传 </li>
	// <li> 3 = 文件下载</li>
	// <li> 4 = 端口性能 </li>
	// <li> 5 = 网络质量 </li>
	// <li> 6 =流媒体 </li>
	// 
	// 即使拨测只支持页面浏览，网络质量，文件下载
	TaskType []*int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 节点类型
	TaskCategory []*int64 `json:"TaskCategory,omitnil,omitempty" name:"TaskCategory"`

	// 排序的列
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 是否正序
	Ascend *bool `json:"Ascend,omitnil,omitempty" name:"Ascend"`

	// 资源标签值
	TagFilters []*KeyValuePair `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

func (r *DescribeProbeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProbeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIDs")
	delete(f, "TaskName")
	delete(f, "TargetAddress")
	delete(f, "TaskStatus")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PayMode")
	delete(f, "OrderState")
	delete(f, "TaskType")
	delete(f, "TaskCategory")
	delete(f, "OrderBy")
	delete(f, "Ascend")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProbeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProbeTasksResponseParams struct {
	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskSet []*ProbeTask `json:"TaskSet,omitnil,omitempty" name:"TaskSet"`

	// 任务总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProbeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProbeTasksResponseParams `json:"Response"`
}

func (r *DescribeProbeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProbeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailedSingleDataDefine struct {
	// 拨测时间戳
	ProbeTime *uint64 `json:"ProbeTime,omitnil,omitempty" name:"ProbeTime"`

	// 储存所有string类型字段
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 储存所有float类型字段
	Fields []*Field `json:"Fields,omitnil,omitempty" name:"Fields"`
}

type Field struct {
	// 自定义字段编号
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义字段名称/说明
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段值
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type KeyValuePair struct {
	// 健
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Label struct {
	// 自定义字段编号
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 自定义字段名称/说明
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type NodeDefine struct {
	// 节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 节点代码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 节点类型
	// <li> 1 = IDC </li>
	// <li> 2 = LastMile </li>
	// <li> 3 = Mobile </li>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 网络服务商
	NetService *string `json:"NetService,omitnil,omitempty" name:"NetService"`

	// 区域
	District *string `json:"District,omitnil,omitempty" name:"District"`

	// 城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// IP 类型
	// <li> 1 = IPv4 </li>
	// <li> 2 = IPv6 </li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPType *int64 `json:"IPType,omitnil,omitempty" name:"IPType"`

	// 区域
	// <li> 1 = 中国大陆 </li>
	// <li> 2 = 港澳台 </li>
	// <li> 3 = 国外 </li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`

	// 节点类型  如果为base 则为可用性拨测点，为空则为高级拨测点
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeType *string `json:"CodeType,omitnil,omitempty" name:"CodeType"`

	// 节点状态：1-运行,2-下线
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeDefineStatus *uint64 `json:"NodeDefineStatus,omitnil,omitempty" name:"NodeDefineStatus"`
}

type NodeDefineExt struct {
	// 节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 节点代码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 节点类型
	// <li> 1 = IDC </li>
	// <li> 2 = LastMile </li>
	// <li> 3 = Mobile </li>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 网络服务商
	NetService *string `json:"NetService,omitnil,omitempty" name:"NetService"`

	// 区域
	District *string `json:"District,omitnil,omitempty" name:"District"`

	// 城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// IP 类型
	// <li> 1 = IPv4 </li>
	// <li> 2 = IPv6 </li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPType *int64 `json:"IPType,omitnil,omitempty" name:"IPType"`

	// 区域
	// <li> 1 = 中国大陆 </li>
	// <li> 2 = 港澳台 </li>
	// <li> 3 = 境外 </li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`

	// 节点类型  如果为base 则为可用性拨测点，为空则为高级拨测点
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeType *string `json:"CodeType,omitnil,omitempty" name:"CodeType"`

	// 节点支持的任务类型。1: 页面性能 2: 文件上传 3: 文件下载 4: 端口性能 5: 网络质量 6: 音视频体验
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypes []*int64 `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`
}

type ProbeTask struct {
	// 任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 拨测类型
	// <li>1 = 页面浏览</li>
	// <li> 2 =文件上传 </li>
	// <li> 3 = 文件下载</li>
	// <li> 4 = 端口性能 </li>
	// <li> 5 = 网络质量 </li>
	// <li> 6 =流媒体 </li>
	// 
	// 即时拨测只支持页面浏览，网络质量，文件下载
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 拨测节点列表
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 拨测任务所选的拨测点IP类型，0-不限，1-IPv4，2-IPv6
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeIpType *int64 `json:"NodeIpType,omitnil,omitempty" name:"NodeIpType"`

	// 拨测间隔
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 拨测参数
	Parameters *string `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// 任务状态
	// <li>1 = 创建中</li>
	// <li> 2 = 运行中 </li>
	// <li> 3 = 运行异常 </li>
	// <li> 4 = 暂停中 </li>
	// <li> 5 = 暂停异常 </li>
	// <li> 6 = 任务暂停 </li>
	// <li> 7 = 任务删除中 </li>
	// <li> 8 = 任务删除异常 </li>
	// <li> 9 = 任务删除</li>
	// <li> 10 = 定时任务暂停中 </li>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 目标地址
	TargetAddress *string `json:"TargetAddress,omitnil,omitempty" name:"TargetAddress"`

	// 付费模式
	// <li>1 = 试用版本</li>
	// <li> 2 = 付费版本 </li>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 订单状态
	// <li>1 = 正常</li>
	// <li> 2 = 欠费 </li>
	OrderState *int64 `json:"OrderState,omitnil,omitempty" name:"OrderState"`

	// 任务分类
	// <li>1 = PC</li>
	// <li> 2 = Mobile </li>
	TaskCategory *int64 `json:"TaskCategory,omitnil,omitempty" name:"TaskCategory"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 定时任务cron表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cron *string `json:"Cron,omitnil,omitempty" name:"Cron"`

	// 定时任务启动状态
	// <li>1 = 定时任务表达式生效</li>
	// <li> 2 = 定时任务表达式未生效（一般为任务手动暂停）</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CronState *int64 `json:"CronState,omitnil,omitempty" name:"CronState"`

	// 任务当前绑定的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagInfoList []*KeyValuePair `json:"TagInfoList,omitnil,omitempty" name:"TagInfoList"`

	// 是否为同步账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubSyncFlag *int64 `json:"SubSyncFlag,omitnil,omitempty" name:"SubSyncFlag"`
}

type ProbeTaskBasicConfiguration struct {
	// 拨测任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 拨测目标地址
	TargetAddress *string `json:"TargetAddress,omitnil,omitempty" name:"TargetAddress"`
}

// Predefined struct for user
type ResumeProbeTaskRequestParams struct {
	// 任务 ID
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

type ResumeProbeTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

func (r *ResumeProbeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeProbeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeProbeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeProbeTaskResponseParams struct {
	// 任务总量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 任务成功量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 任务执行详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*TaskResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeProbeTaskResponse struct {
	*tchttp.BaseResponse
	Response *ResumeProbeTaskResponseParams `json:"Response"`
}

func (r *ResumeProbeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeProbeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SingleInstantTask struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务地址
	TargetAddress *string `json:"TargetAddress,omitnil,omitempty" name:"TargetAddress"`

	// 任务类型
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 测试时间
	ProbeTime *uint64 `json:"ProbeTime,omitnil,omitempty" name:"ProbeTime"`

	// 任务状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 成功率
	SuccessRate *float64 `json:"SuccessRate,omitnil,omitempty" name:"SuccessRate"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 节点类型
	TaskCategory *uint64 `json:"TaskCategory,omitnil,omitempty" name:"TaskCategory"`
}

// Predefined struct for user
type SuspendProbeTaskRequestParams struct {
	// 任务 ID
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

type SuspendProbeTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

func (r *SuspendProbeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SuspendProbeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SuspendProbeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SuspendProbeTaskResponseParams struct {
	// 任务总量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 任务成功量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 任务执行结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*TaskResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SuspendProbeTaskResponse struct {
	*tchttp.BaseResponse
	Response *SuspendProbeTaskResponseParams `json:"Response"`
}

func (r *SuspendProbeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SuspendProbeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskResult struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

// Predefined struct for user
type UpdateProbeTaskAttributesRequestParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名，该参数为空时不作任何修改。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type UpdateProbeTaskAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名，该参数为空时不作任何修改。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *UpdateProbeTaskAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProbeTaskAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateProbeTaskAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProbeTaskAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateProbeTaskAttributesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateProbeTaskAttributesResponseParams `json:"Response"`
}

func (r *UpdateProbeTaskAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProbeTaskAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProbeTaskConfigurationListRequestParams struct {
	// 任务 ID，如task-n1wchki8
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 拨测节点，如10001，详细地区运营商拨测编号请联系云拨测。
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 拨测间隔，如30，单位为分钟。
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 拨测参数，详细参数配置可参考云拨测官网文档。
	Parameters *string `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// 定时任务cron表达式
	Cron *string `json:"Cron,omitnil,omitempty" name:"Cron"`

	// 预付费套餐id
	// 需要与taskId对应
	ResourceIDs []*string `json:"ResourceIDs,omitnil,omitempty" name:"ResourceIDs"`

	// 拨测节点的IP类型，0-不限，1-IPv4，2-IPv6
	NodeIpType *int64 `json:"NodeIpType,omitnil,omitempty" name:"NodeIpType"`
}

type UpdateProbeTaskConfigurationListRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID，如task-n1wchki8
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 拨测节点，如10001，详细地区运营商拨测编号请联系云拨测。
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 拨测间隔，如30，单位为分钟。
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 拨测参数，详细参数配置可参考云拨测官网文档。
	Parameters *string `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// 定时任务cron表达式
	Cron *string `json:"Cron,omitnil,omitempty" name:"Cron"`

	// 预付费套餐id
	// 需要与taskId对应
	ResourceIDs []*string `json:"ResourceIDs,omitnil,omitempty" name:"ResourceIDs"`

	// 拨测节点的IP类型，0-不限，1-IPv4，2-IPv6
	NodeIpType *int64 `json:"NodeIpType,omitnil,omitempty" name:"NodeIpType"`
}

func (r *UpdateProbeTaskConfigurationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProbeTaskConfigurationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "Nodes")
	delete(f, "Interval")
	delete(f, "Parameters")
	delete(f, "Cron")
	delete(f, "ResourceIDs")
	delete(f, "NodeIpType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateProbeTaskConfigurationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProbeTaskConfigurationListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateProbeTaskConfigurationListResponse struct {
	*tchttp.BaseResponse
	Response *UpdateProbeTaskConfigurationListResponseParams `json:"Response"`
}

func (r *UpdateProbeTaskConfigurationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProbeTaskConfigurationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}