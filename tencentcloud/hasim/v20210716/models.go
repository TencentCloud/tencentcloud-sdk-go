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

package v20210716

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateRuleRequestParams struct {
	// 自动化规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则类型：用量类(101 当月|102有效期内)、位置类(201行政区|202移动距离)、网络质量类(301网络盲点)
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 是否激活
	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`

	// 触发动作：1 邮件 2 API请求 3 微信 4 停卡 5 地图标识为盲点
	Notice *int64 `json:"Notice,omitempty" name:"Notice"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 推送的API地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 用量阈值
	DataThreshold *int64 `json:"DataThreshold,omitempty" name:"DataThreshold"`

	// 行政区类型：1. 省份 2. 城市 3. 区
	District *int64 `json:"District,omitempty" name:"District"`

	// 心跳移动距离阈值
	Distance *int64 `json:"Distance,omitempty" name:"Distance"`

	// 信号强度阈值
	SignalStrength *int64 `json:"SignalStrength,omitempty" name:"SignalStrength"`

	// 盲点时间阈值，天
	LostDay *int64 `json:"LostDay,omitempty" name:"LostDay"`

	// 标签ID集合
	TagIDs []*int64 `json:"TagIDs,omitempty" name:"TagIDs"`

	// 资费计划
	SalePlan *string `json:"SalePlan,omitempty" name:"SalePlan"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 自动化规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则类型：用量类(101 当月|102有效期内)、位置类(201行政区|202移动距离)、网络质量类(301网络盲点)
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 是否激活
	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`

	// 触发动作：1 邮件 2 API请求 3 微信 4 停卡 5 地图标识为盲点
	Notice *int64 `json:"Notice,omitempty" name:"Notice"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 推送的API地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 用量阈值
	DataThreshold *int64 `json:"DataThreshold,omitempty" name:"DataThreshold"`

	// 行政区类型：1. 省份 2. 城市 3. 区
	District *int64 `json:"District,omitempty" name:"District"`

	// 心跳移动距离阈值
	Distance *int64 `json:"Distance,omitempty" name:"Distance"`

	// 信号强度阈值
	SignalStrength *int64 `json:"SignalStrength,omitempty" name:"SignalStrength"`

	// 盲点时间阈值，天
	LostDay *int64 `json:"LostDay,omitempty" name:"LostDay"`

	// 标签ID集合
	TagIDs []*int64 `json:"TagIDs,omitempty" name:"TagIDs"`

	// 资费计划
	SalePlan *string `json:"SalePlan,omitempty" name:"SalePlan"`
}

func (r *CreateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "IsActive")
	delete(f, "Notice")
	delete(f, "Email")
	delete(f, "Url")
	delete(f, "DataThreshold")
	delete(f, "District")
	delete(f, "Distance")
	delete(f, "SignalStrength")
	delete(f, "LostDay")
	delete(f, "TagIDs")
	delete(f, "SalePlan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRuleResponseParams `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTacticRequestParams struct {
	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否自动执行
	IsAuto *int64 `json:"IsAuto,omitempty" name:"IsAuto"`

	// 心跳上报间隔(s)
	PingInterval *int64 `json:"PingInterval,omitempty" name:"PingInterval"`

	// 是否开启弱信号检测
	IsWeak *int64 `json:"IsWeak,omitempty" name:"IsWeak"`

	// 弱信号阈值（-dbm）
	WeakThreshold *int64 `json:"WeakThreshold,omitempty" name:"WeakThreshold"`

	// 是否开启时延切换
	IsDelay *int64 `json:"IsDelay,omitempty" name:"IsDelay"`

	// 网络时延阈值（ms）
	DelayThreshold *int64 `json:"DelayThreshold,omitempty" name:"DelayThreshold"`

	// 是否开启假信号检测
	IsFake *int64 `json:"IsFake,omitempty" name:"IsFake"`

	// 假信号检测IP字符串，用逗号分隔
	FakeIP *string `json:"FakeIP,omitempty" name:"FakeIP"`

	// 假信号检测间隔（s）
	FakeInterval *int64 `json:"FakeInterval,omitempty" name:"FakeInterval"`

	// 是否开启网络制式检测
	IsNet *int64 `json:"IsNet,omitempty" name:"IsNet"`

	// 网络回落制式 1 2G、 2 3G 、 3 2/3G
	Network *int64 `json:"Network,omitempty" name:"Network"`

	// 是否开启移动检测
	IsMove *int64 `json:"IsMove,omitempty" name:"IsMove"`

	// 是否开启最优先运营商
	IsPriorityTele *int64 `json:"IsPriorityTele,omitempty" name:"IsPriorityTele"`

	// 最优先运营商 1 移动、 2 联通、 3 电信 4 上次在线运营商
	PriorityTele *int64 `json:"PriorityTele,omitempty" name:"PriorityTele"`

	// 是否开启最不优先运营商
	IsBottomTele *int64 `json:"IsBottomTele,omitempty" name:"IsBottomTele"`

	// 最不优先运营商 1 移动、 2 联通、 3 电信
	BottomTele *int64 `json:"BottomTele,omitempty" name:"BottomTele"`

	// 最优先信号选取策略
	IsBestSignal *int64 `json:"IsBestSignal,omitempty" name:"IsBestSignal"`
}

type CreateTacticRequest struct {
	*tchttp.BaseRequest
	
	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否自动执行
	IsAuto *int64 `json:"IsAuto,omitempty" name:"IsAuto"`

	// 心跳上报间隔(s)
	PingInterval *int64 `json:"PingInterval,omitempty" name:"PingInterval"`

	// 是否开启弱信号检测
	IsWeak *int64 `json:"IsWeak,omitempty" name:"IsWeak"`

	// 弱信号阈值（-dbm）
	WeakThreshold *int64 `json:"WeakThreshold,omitempty" name:"WeakThreshold"`

	// 是否开启时延切换
	IsDelay *int64 `json:"IsDelay,omitempty" name:"IsDelay"`

	// 网络时延阈值（ms）
	DelayThreshold *int64 `json:"DelayThreshold,omitempty" name:"DelayThreshold"`

	// 是否开启假信号检测
	IsFake *int64 `json:"IsFake,omitempty" name:"IsFake"`

	// 假信号检测IP字符串，用逗号分隔
	FakeIP *string `json:"FakeIP,omitempty" name:"FakeIP"`

	// 假信号检测间隔（s）
	FakeInterval *int64 `json:"FakeInterval,omitempty" name:"FakeInterval"`

	// 是否开启网络制式检测
	IsNet *int64 `json:"IsNet,omitempty" name:"IsNet"`

	// 网络回落制式 1 2G、 2 3G 、 3 2/3G
	Network *int64 `json:"Network,omitempty" name:"Network"`

	// 是否开启移动检测
	IsMove *int64 `json:"IsMove,omitempty" name:"IsMove"`

	// 是否开启最优先运营商
	IsPriorityTele *int64 `json:"IsPriorityTele,omitempty" name:"IsPriorityTele"`

	// 最优先运营商 1 移动、 2 联通、 3 电信 4 上次在线运营商
	PriorityTele *int64 `json:"PriorityTele,omitempty" name:"PriorityTele"`

	// 是否开启最不优先运营商
	IsBottomTele *int64 `json:"IsBottomTele,omitempty" name:"IsBottomTele"`

	// 最不优先运营商 1 移动、 2 联通、 3 电信
	BottomTele *int64 `json:"BottomTele,omitempty" name:"BottomTele"`

	// 最优先信号选取策略
	IsBestSignal *int64 `json:"IsBestSignal,omitempty" name:"IsBestSignal"`
}

func (r *CreateTacticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTacticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "IsAuto")
	delete(f, "PingInterval")
	delete(f, "IsWeak")
	delete(f, "WeakThreshold")
	delete(f, "IsDelay")
	delete(f, "DelayThreshold")
	delete(f, "IsFake")
	delete(f, "FakeIP")
	delete(f, "FakeInterval")
	delete(f, "IsNet")
	delete(f, "Network")
	delete(f, "IsMove")
	delete(f, "IsPriorityTele")
	delete(f, "PriorityTele")
	delete(f, "IsBottomTele")
	delete(f, "BottomTele")
	delete(f, "IsBestSignal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTacticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTacticResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTacticResponse struct {
	*tchttp.BaseResponse
	Response *CreateTacticResponseParams `json:"Response"`
}

func (r *CreateTacticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTacticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRequestParams struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type CreateTagRequest struct {
	*tchttp.BaseRequest
	
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *CreateTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTagResponse struct {
	*tchttp.BaseResponse
	Response *CreateTagResponseParams `json:"Response"`
}

func (r *CreateTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleRequestParams struct {
	// 自动化规则ID
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 自动化规则ID
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRuleResponseParams `json:"Response"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTacticRequestParams struct {
	// 策略ID
	TacticID *int64 `json:"TacticID,omitempty" name:"TacticID"`
}

type DeleteTacticRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID
	TacticID *int64 `json:"TacticID,omitempty" name:"TacticID"`
}

func (r *DeleteTacticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTacticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TacticID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTacticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTacticResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTacticResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTacticResponseParams `json:"Response"`
}

func (r *DeleteTacticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTacticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagRequestParams struct {
	// 标签ID
	TagID *int64 `json:"TagID,omitempty" name:"TagID"`
}

type DeleteTagRequest struct {
	*tchttp.BaseRequest
	
	// 标签ID
	TagID *int64 `json:"TagID,omitempty" name:"TagID"`
}

func (r *DeleteTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTagResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTagResponseParams `json:"Response"`
}

func (r *DeleteTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLinkRequestParams struct {
	// 云兔卡ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 具体的账号
	UinAccount *string `json:"UinAccount,omitempty" name:"UinAccount"`
}

type DescribeLinkRequest struct {
	*tchttp.BaseRequest
	
	// 云兔卡ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 具体的账号
	UinAccount *string `json:"UinAccount,omitempty" name:"UinAccount"`
}

func (r *DescribeLinkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLinkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LinkID")
	delete(f, "UinAccount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLinkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLinkResponseParams struct {
	// 云兔连接详细信息
	Data *LinkDetailInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLinkResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLinkResponseParams `json:"Response"`
}

func (r *DescribeLinkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLinksRequestParams struct {
	// 云兔卡ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 运营商ICCID
	ICCID *string `json:"ICCID,omitempty" name:"ICCID"`

	// 设备码
	IMEI *string `json:"IMEI,omitempty" name:"IMEI"`

	// 卡片状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 运营商 1移动 2联通 3电信
	TeleOperator *uint64 `json:"TeleOperator,omitempty" name:"TeleOperator"`

	// 标签ID
	TagID *uint64 `json:"TagID,omitempty" name:"TagID"`

	// 策略ID
	TacticID *uint64 `json:"TacticID,omitempty" name:"TacticID"`

	// 设备在线状态 0 未激活 1 在线 2 离线
	LinkedState *int64 `json:"LinkedState,omitempty" name:"LinkedState"`

	// 标签ID 集合
	TagIDs []*int64 `json:"TagIDs,omitempty" name:"TagIDs"`

	// 翻页大小, 默认翻页大小为10，最大数量为500
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeLinksRequest struct {
	*tchttp.BaseRequest
	
	// 云兔卡ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 运营商ICCID
	ICCID *string `json:"ICCID,omitempty" name:"ICCID"`

	// 设备码
	IMEI *string `json:"IMEI,omitempty" name:"IMEI"`

	// 卡片状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 运营商 1移动 2联通 3电信
	TeleOperator *uint64 `json:"TeleOperator,omitempty" name:"TeleOperator"`

	// 标签ID
	TagID *uint64 `json:"TagID,omitempty" name:"TagID"`

	// 策略ID
	TacticID *uint64 `json:"TacticID,omitempty" name:"TacticID"`

	// 设备在线状态 0 未激活 1 在线 2 离线
	LinkedState *int64 `json:"LinkedState,omitempty" name:"LinkedState"`

	// 标签ID 集合
	TagIDs []*int64 `json:"TagIDs,omitempty" name:"TagIDs"`

	// 翻页大小, 默认翻页大小为10，最大数量为500
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页起始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLinksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLinksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LinkID")
	delete(f, "ICCID")
	delete(f, "IMEI")
	delete(f, "Status")
	delete(f, "TeleOperator")
	delete(f, "TagID")
	delete(f, "TacticID")
	delete(f, "LinkedState")
	delete(f, "TagIDs")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLinksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLinksResponseParams struct {
	// 云兔连接响应信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *LinkInfos `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLinksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLinksResponseParams `json:"Response"`
}

func (r *DescribeLinksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLinksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrdersRequestParams struct {
	// 子订单ID
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 审批状态 0全部 1通过 2驳回 3待审核
	AuditStatus *int64 `json:"AuditStatus,omitempty" name:"AuditStatus"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 开始时间,例如2022-06-30 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间,例如2022-06-30 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest
	
	// 子订单ID
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 审批状态 0全部 1通过 2驳回 3待审核
	AuditStatus *int64 `json:"AuditStatus,omitempty" name:"AuditStatus"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 开始时间,例如2022-06-30 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间,例如2022-06-30 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealName")
	delete(f, "AuditStatus")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrdersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrdersResponseParams struct {
	// 订单列表
	Data *Orders `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrdersResponseParams `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleRequestParams struct {
	// 自动化规则ID
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
}

type DescribeRuleRequest struct {
	*tchttp.BaseRequest
	
	// 自动化规则ID
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
}

func (r *DescribeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleResponseParams struct {
	// 策略信息
	Data *RuleDetail `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleResponseParams `json:"Response"`
}

func (r *DescribeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesRequestParams struct {
	// 自动化规则ID
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 自动化规则ID
	RuleIDs []*int64 `json:"RuleIDs,omitempty" name:"RuleIDs"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 类型
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 是否激活
	IsActive *int64 `json:"IsActive,omitempty" name:"IsActive"`

	// 翻页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest
	
	// 自动化规则ID
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 自动化规则ID
	RuleIDs []*int64 `json:"RuleIDs,omitempty" name:"RuleIDs"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 类型
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 是否激活
	IsActive *int64 `json:"IsActive,omitempty" name:"IsActive"`

	// 翻页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleID")
	delete(f, "RuleIDs")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "IsActive")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesResponseParams struct {
	// 自动化规则列表集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RuleInfos `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRulesResponseParams `json:"Response"`
}

func (r *DescribeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTacticRequestParams struct {
	// 策略ID
	TacticID *int64 `json:"TacticID,omitempty" name:"TacticID"`
}

type DescribeTacticRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID
	TacticID *int64 `json:"TacticID,omitempty" name:"TacticID"`
}

func (r *DescribeTacticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTacticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TacticID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTacticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTacticResponseParams struct {
	// 策略信息
	Data *Tactic `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTacticResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTacticResponseParams `json:"Response"`
}

func (r *DescribeTacticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTacticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTacticsRequestParams struct {
	// 策略ID
	TacticID *int64 `json:"TacticID,omitempty" name:"TacticID"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeTacticsRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID
	TacticID *int64 `json:"TacticID,omitempty" name:"TacticID"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeTacticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTacticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TacticID")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTacticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTacticsResponseParams struct {
	// 策略集合信息
	Data *TacticInfos `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTacticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTacticsResponseParams `json:"Response"`
}

func (r *DescribeTacticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTacticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsRequestParams struct {
	// 标签名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeTagsRequest struct {
	*tchttp.BaseRequest
	
	// 标签名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagsResponseParams struct {
	// 列表
	Data *TagInfos `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagsResponseParams `json:"Response"`
}

func (r *DescribeTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceReport struct {
	// 移动设备ID
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// 经度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lng *string `json:"Lng,omitempty" name:"Lng"`

	// 维度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lat *string `json:"Lat,omitempty" name:"Lat"`

	// 运营商基站ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lac *string `json:"Lac,omitempty" name:"Lac"`

	// 小区CellID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cell *string `json:"Cell,omitempty" name:"Cell"`

	// 当前上报运营商ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 信号强度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rss *int64 `json:"Rss,omitempty" name:"Rss"`

	// 运营商: 1 移动 2 联通 3 电信
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tele *int64 `json:"Tele,omitempty" name:"Tele"`

	// 当前设备策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tid *int64 `json:"Tid,omitempty" name:"Tid"`

	// 心跳间隔,单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ping *int64 `json:"Ping,omitempty" name:"Ping"`

	// 网络延迟,单位毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Delay *int64 `json:"Delay,omitempty" name:"Delay"`

	// 高级日志启停状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Log *int64 `json:"Log,omitempty" name:"Log"`

	// 设备型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevType *string `json:"DevType,omitempty" name:"DevType"`

	// 设备型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevModel *string `json:"DevModel,omitempty" name:"DevModel"`

	// 设备版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 设备刷新时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadTime *string `json:"UploadTime,omitempty" name:"UploadTime"`

	// 网络环境: 0 正常 1 弱网
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 每月第一次上报心跳时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonthFirstTime *string `json:"MonthFirstTime,omitempty" name:"MonthFirstTime"`
}

type LinkDetailInfo struct {
	// 云兔连接ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 卡片状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveTime *string `json:"ActiveTime,omitempty" name:"ActiveTime"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 数据用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataUse *float64 `json:"DataUse,omitempty" name:"DataUse"`

	// 语音用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioUse *int64 `json:"AudioUse,omitempty" name:"AudioUse"`

	// 短信用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmsUse *int64 `json:"SmsUse,omitempty" name:"SmsUse"`

	// 在线状态 0 未激活 1 在线 2 离线
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkedState *int64 `json:"LinkedState,omitempty" name:"LinkedState"`

	// 预期策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TacticID *int64 `json:"TacticID,omitempty" name:"TacticID"`

	// 策略下发状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TacticStatus *int64 `json:"TacticStatus,omitempty" name:"TacticStatus"`

	// 策略下发成功过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TacticExpireTime *string `json:"TacticExpireTime,omitempty" name:"TacticExpireTime"`

	// 高级日志预期状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsActiveLog *bool `json:"IsActiveLog,omitempty" name:"IsActiveLog"`

	// 运营商 1移动 2联通 3电信
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeleOperator *int64 `json:"TeleOperator,omitempty" name:"TeleOperator"`

	// 设备最新上报信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Report *DeviceReport `json:"Report,omitempty" name:"Report"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 运营商ICCID信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cards []*TeleOperatorCard `json:"Cards,omitempty" name:"Cards"`

	// 云兔实际卡片ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CardID *string `json:"CardID,omitempty" name:"CardID"`
}

type LinkInfo struct {
	// 云兔连接ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 卡片状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveTime *string `json:"ActiveTime,omitempty" name:"ActiveTime"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 数据用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataUse *float64 `json:"DataUse,omitempty" name:"DataUse"`

	// 语音用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioUse *int64 `json:"AudioUse,omitempty" name:"AudioUse"`

	// 短信用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmsUse *int64 `json:"SmsUse,omitempty" name:"SmsUse"`

	// 在线状态 0 未激活 1 在线 2 离线
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkedState *int64 `json:"LinkedState,omitempty" name:"LinkedState"`

	// 预期策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TacticID *int64 `json:"TacticID,omitempty" name:"TacticID"`

	// 策略下发状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TacticStatus *int64 `json:"TacticStatus,omitempty" name:"TacticStatus"`

	// 策略下发成功过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TacticExpireTime *string `json:"TacticExpireTime,omitempty" name:"TacticExpireTime"`

	// 高级日志预期状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsActiveLog *bool `json:"IsActiveLog,omitempty" name:"IsActiveLog"`

	// 运营商 1移动 2联通 3电信
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeleOperator *int64 `json:"TeleOperator,omitempty" name:"TeleOperator"`

	// 设备最新上报信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Report *DeviceReport `json:"Report,omitempty" name:"Report"`
}

type LinkInfos struct {
	// 总量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 云兔连接列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*LinkInfo `json:"List,omitempty" name:"List"`
}

// Predefined struct for user
type ModifyLinkAdvancedLogRequestParams struct {
	// 云兔ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 是否激活高级日志 0 关闭 1激活
	IsAdLog *int64 `json:"IsAdLog,omitempty" name:"IsAdLog"`
}

type ModifyLinkAdvancedLogRequest struct {
	*tchttp.BaseRequest
	
	// 云兔ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 是否激活高级日志 0 关闭 1激活
	IsAdLog *int64 `json:"IsAdLog,omitempty" name:"IsAdLog"`
}

func (r *ModifyLinkAdvancedLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLinkAdvancedLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LinkID")
	delete(f, "IsAdLog")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLinkAdvancedLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLinkAdvancedLogResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLinkAdvancedLogResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLinkAdvancedLogResponseParams `json:"Response"`
}

func (r *ModifyLinkAdvancedLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLinkAdvancedLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLinkTacticRequestParams struct {
	// 云兔ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 策略ID
	TacticID *int64 `json:"TacticID,omitempty" name:"TacticID"`
}

type ModifyLinkTacticRequest struct {
	*tchttp.BaseRequest
	
	// 云兔ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 策略ID
	TacticID *int64 `json:"TacticID,omitempty" name:"TacticID"`
}

func (r *ModifyLinkTacticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLinkTacticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LinkID")
	delete(f, "TacticID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLinkTacticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLinkTacticResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLinkTacticResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLinkTacticResponseParams `json:"Response"`
}

func (r *ModifyLinkTacticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLinkTacticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLinkTeleRequestParams struct {
	// 云兔ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 运营商 1 移动 2 联通 3 电信
	TeleOperator *int64 `json:"TeleOperator,omitempty" name:"TeleOperator"`
}

type ModifyLinkTeleRequest struct {
	*tchttp.BaseRequest
	
	// 云兔ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 运营商 1 移动 2 联通 3 电信
	TeleOperator *int64 `json:"TeleOperator,omitempty" name:"TeleOperator"`
}

func (r *ModifyLinkTeleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLinkTeleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LinkID")
	delete(f, "TeleOperator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLinkTeleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLinkTeleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLinkTeleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLinkTeleResponseParams `json:"Response"`
}

func (r *ModifyLinkTeleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLinkTeleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleRequestParams struct {
	// 自动化规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则类型：用量类(101 当月|102有效期内)、位置类(201行政区|202移动距离)、网络质量类(301网络盲点)
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 是否激活
	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`

	// 触发动作：1 邮件 2 API请求 3 微信 4 停卡 5 地图标识为盲点
	Notice *int64 `json:"Notice,omitempty" name:"Notice"`

	// 自动化规则ID
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 推送的API地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 用量阈值
	DataThreshold *int64 `json:"DataThreshold,omitempty" name:"DataThreshold"`

	// 行政区类型：1. 省份 2. 城市 3. 区
	District *int64 `json:"District,omitempty" name:"District"`

	// 心跳移动距离阈值
	Distance *int64 `json:"Distance,omitempty" name:"Distance"`

	// 信号强度阈值
	SignalStrength *int64 `json:"SignalStrength,omitempty" name:"SignalStrength"`

	// 标签ID集合
	TagIDs []*int64 `json:"TagIDs,omitempty" name:"TagIDs"`

	// 资费计划
	SalePlan *string `json:"SalePlan,omitempty" name:"SalePlan"`

	// 具体的账号
	UinAccount *string `json:"UinAccount,omitempty" name:"UinAccount"`
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 自动化规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则类型：用量类(101 当月|102有效期内)、位置类(201行政区|202移动距离)、网络质量类(301网络盲点)
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 是否激活
	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`

	// 触发动作：1 邮件 2 API请求 3 微信 4 停卡 5 地图标识为盲点
	Notice *int64 `json:"Notice,omitempty" name:"Notice"`

	// 自动化规则ID
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 推送的API地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 用量阈值
	DataThreshold *int64 `json:"DataThreshold,omitempty" name:"DataThreshold"`

	// 行政区类型：1. 省份 2. 城市 3. 区
	District *int64 `json:"District,omitempty" name:"District"`

	// 心跳移动距离阈值
	Distance *int64 `json:"Distance,omitempty" name:"Distance"`

	// 信号强度阈值
	SignalStrength *int64 `json:"SignalStrength,omitempty" name:"SignalStrength"`

	// 标签ID集合
	TagIDs []*int64 `json:"TagIDs,omitempty" name:"TagIDs"`

	// 资费计划
	SalePlan *string `json:"SalePlan,omitempty" name:"SalePlan"`

	// 具体的账号
	UinAccount *string `json:"UinAccount,omitempty" name:"UinAccount"`
}

func (r *ModifyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "IsActive")
	delete(f, "Notice")
	delete(f, "RuleID")
	delete(f, "Email")
	delete(f, "Url")
	delete(f, "DataThreshold")
	delete(f, "District")
	delete(f, "Distance")
	delete(f, "SignalStrength")
	delete(f, "TagIDs")
	delete(f, "SalePlan")
	delete(f, "UinAccount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRuleResponseParams `json:"Response"`
}

func (r *ModifyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleStatusRequestParams struct {
	// 自动化规则ID
	RuleID *uint64 `json:"RuleID,omitempty" name:"RuleID"`

	// 是否激活
	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`
}

type ModifyRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 自动化规则ID
	RuleID *uint64 `json:"RuleID,omitempty" name:"RuleID"`

	// 是否激活
	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`
}

func (r *ModifyRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleID")
	delete(f, "IsActive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTacticRequestParams struct {
	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否自动执行
	IsAuto *int64 `json:"IsAuto,omitempty" name:"IsAuto"`

	// 心跳上报间隔(s)
	PingInterval *int64 `json:"PingInterval,omitempty" name:"PingInterval"`

	// 是否开启弱信号检测
	IsWeak *int64 `json:"IsWeak,omitempty" name:"IsWeak"`

	// 弱信号阈值（-dbm）
	WeakThreshold *int64 `json:"WeakThreshold,omitempty" name:"WeakThreshold"`

	// 是否开启时延切换
	IsDelay *int64 `json:"IsDelay,omitempty" name:"IsDelay"`

	// 网络时延阈值（ms）
	DelayThreshold *int64 `json:"DelayThreshold,omitempty" name:"DelayThreshold"`

	// 是否开启假信号检测
	IsFake *int64 `json:"IsFake,omitempty" name:"IsFake"`

	// 假信号检测间隔（s）
	FakeInterval *int64 `json:"FakeInterval,omitempty" name:"FakeInterval"`

	// 是否开启网络制式检测
	IsNet *int64 `json:"IsNet,omitempty" name:"IsNet"`

	// 网络回落制式 1 2G、 2 3G 、 3 2/3G
	Network *int64 `json:"Network,omitempty" name:"Network"`

	// 是否开启移动检测
	IsMove *int64 `json:"IsMove,omitempty" name:"IsMove"`

	// 策略ID
	TacticID *int64 `json:"TacticID,omitempty" name:"TacticID"`

	// 是否开启最优先运营商
	IsPriorityTele *int64 `json:"IsPriorityTele,omitempty" name:"IsPriorityTele"`

	// 最优先运营商 1 移动、 2 联通、 3 电信 4 上次在线运营商
	PriorityTele *int64 `json:"PriorityTele,omitempty" name:"PriorityTele"`

	// 是否开启最不优先运营商
	IsBottomTele *int64 `json:"IsBottomTele,omitempty" name:"IsBottomTele"`

	// 最不优先运营商 1 移动、 2 联通、 3 电信
	BottomTele *int64 `json:"BottomTele,omitempty" name:"BottomTele"`

	// 是否最优先信号选取策略
	IsBestSignal *int64 `json:"IsBestSignal,omitempty" name:"IsBestSignal"`

	// 假信号检测IP字符串，用逗号分隔
	FakeIP *string `json:"FakeIP,omitempty" name:"FakeIP"`
}

type ModifyTacticRequest struct {
	*tchttp.BaseRequest
	
	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否自动执行
	IsAuto *int64 `json:"IsAuto,omitempty" name:"IsAuto"`

	// 心跳上报间隔(s)
	PingInterval *int64 `json:"PingInterval,omitempty" name:"PingInterval"`

	// 是否开启弱信号检测
	IsWeak *int64 `json:"IsWeak,omitempty" name:"IsWeak"`

	// 弱信号阈值（-dbm）
	WeakThreshold *int64 `json:"WeakThreshold,omitempty" name:"WeakThreshold"`

	// 是否开启时延切换
	IsDelay *int64 `json:"IsDelay,omitempty" name:"IsDelay"`

	// 网络时延阈值（ms）
	DelayThreshold *int64 `json:"DelayThreshold,omitempty" name:"DelayThreshold"`

	// 是否开启假信号检测
	IsFake *int64 `json:"IsFake,omitempty" name:"IsFake"`

	// 假信号检测间隔（s）
	FakeInterval *int64 `json:"FakeInterval,omitempty" name:"FakeInterval"`

	// 是否开启网络制式检测
	IsNet *int64 `json:"IsNet,omitempty" name:"IsNet"`

	// 网络回落制式 1 2G、 2 3G 、 3 2/3G
	Network *int64 `json:"Network,omitempty" name:"Network"`

	// 是否开启移动检测
	IsMove *int64 `json:"IsMove,omitempty" name:"IsMove"`

	// 策略ID
	TacticID *int64 `json:"TacticID,omitempty" name:"TacticID"`

	// 是否开启最优先运营商
	IsPriorityTele *int64 `json:"IsPriorityTele,omitempty" name:"IsPriorityTele"`

	// 最优先运营商 1 移动、 2 联通、 3 电信 4 上次在线运营商
	PriorityTele *int64 `json:"PriorityTele,omitempty" name:"PriorityTele"`

	// 是否开启最不优先运营商
	IsBottomTele *int64 `json:"IsBottomTele,omitempty" name:"IsBottomTele"`

	// 最不优先运营商 1 移动、 2 联通、 3 电信
	BottomTele *int64 `json:"BottomTele,omitempty" name:"BottomTele"`

	// 是否最优先信号选取策略
	IsBestSignal *int64 `json:"IsBestSignal,omitempty" name:"IsBestSignal"`

	// 假信号检测IP字符串，用逗号分隔
	FakeIP *string `json:"FakeIP,omitempty" name:"FakeIP"`
}

func (r *ModifyTacticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTacticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "IsAuto")
	delete(f, "PingInterval")
	delete(f, "IsWeak")
	delete(f, "WeakThreshold")
	delete(f, "IsDelay")
	delete(f, "DelayThreshold")
	delete(f, "IsFake")
	delete(f, "FakeInterval")
	delete(f, "IsNet")
	delete(f, "Network")
	delete(f, "IsMove")
	delete(f, "TacticID")
	delete(f, "IsPriorityTele")
	delete(f, "PriorityTele")
	delete(f, "IsBottomTele")
	delete(f, "BottomTele")
	delete(f, "IsBestSignal")
	delete(f, "FakeIP")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTacticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTacticResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTacticResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTacticResponseParams `json:"Response"`
}

func (r *ModifyTacticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTacticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTagRequestParams struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标签ID
	TagID *int64 `json:"TagID,omitempty" name:"TagID"`

	// 备注
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type ModifyTagRequest struct {
	*tchttp.BaseRequest
	
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标签ID
	TagID *int64 `json:"TagID,omitempty" name:"TagID"`

	// 备注
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *ModifyTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "TagID")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTagResponseParams `json:"Response"`
}

func (r *ModifyTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrderInfo struct {
	// 子订单ID
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 订单账户
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 购买数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuyNum *int64 `json:"BuyNum,omitempty" name:"BuyNum"`

	// 行业代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndustryCode *string `json:"IndustryCode,omitempty" name:"IndustryCode"`

	// 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 联系人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Contact *string `json:"Contact,omitempty" name:"Contact"`

	// 电话号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msisdn *string `json:"Msisdn,omitempty" name:"Msisdn"`

	// 卡片规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Specification *string `json:"Specification,omitempty" name:"Specification"`

	// 用户订单备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`

	// 审批状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuditStatus *string `json:"AuditStatus,omitempty" name:"AuditStatus"`

	// 发货状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowStatus *string `json:"FlowStatus,omitempty" name:"FlowStatus"`

	// 审批备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 退费订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundBigDealId *string `json:"RefundBigDealId,omitempty" name:"RefundBigDealId"`
}

type Orders struct {
	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 订单集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*OrderInfo `json:"List,omitempty" name:"List"`
}

// Predefined struct for user
type RenewLinkInfoRequestParams struct {
	// 云兔ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 具体的账号
	UinAccount *string `json:"UinAccount,omitempty" name:"UinAccount"`
}

type RenewLinkInfoRequest struct {
	*tchttp.BaseRequest
	
	// 云兔ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 具体的账号
	UinAccount *string `json:"UinAccount,omitempty" name:"UinAccount"`
}

func (r *RenewLinkInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewLinkInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LinkID")
	delete(f, "UinAccount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewLinkInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewLinkInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RenewLinkInfoResponse struct {
	*tchttp.BaseResponse
	Response *RenewLinkInfoResponseParams `json:"Response"`
}

func (r *RenewLinkInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewLinkInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Rule struct {
	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 删除时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeletedAt *string `json:"DeletedAt,omitempty" name:"DeletedAt"`

	// 规则类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 是否激活
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`

	// 触发动作：1 邮件 2 API请求 5 停卡 6 地图标识为盲点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Notice *int64 `json:"Notice,omitempty" name:"Notice"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitempty" name:"Email"`

	// 回调API地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 用量类：用量阈值,单位MB
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataThreshold *int64 `json:"DataThreshold,omitempty" name:"DataThreshold"`

	// 行政区类型：1. 省份 2. 城市 3. 区
	// 注意：此字段可能返回 null，表示取不到有效值。
	District *int64 `json:"District,omitempty" name:"District"`

	// 移动距离阈值，单位KM
	// 注意：此字段可能返回 null，表示取不到有效值。
	Distance *int64 `json:"Distance,omitempty" name:"Distance"`

	// 信号强度阈值(-dbm）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignalStrength *int64 `json:"SignalStrength,omitempty" name:"SignalStrength"`

	// 盲点阈值天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LostDay *int64 `json:"LostDay,omitempty" name:"LostDay"`

	// 绑定的标签ID集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagIDs []*uint64 `json:"TagIDs,omitempty" name:"TagIDs"`

	// 绑定的资费计划
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalePlan *string `json:"SalePlan,omitempty" name:"SalePlan"`
}

type RuleDetail struct {
	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 删除时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeletedAt *string `json:"DeletedAt,omitempty" name:"DeletedAt"`

	// 规则类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 是否激活
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsActive *bool `json:"IsActive,omitempty" name:"IsActive"`

	// 触发动作：1 邮件 2 API请求 5 停卡 6 地图标识为盲点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Notice *int64 `json:"Notice,omitempty" name:"Notice"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitempty" name:"Email"`

	// 回调API地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 用量类：用量阈值,单位MB
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataThreshold *int64 `json:"DataThreshold,omitempty" name:"DataThreshold"`

	// 行政区类型：1. 省份 2. 城市 3. 区
	// 注意：此字段可能返回 null，表示取不到有效值。
	District *int64 `json:"District,omitempty" name:"District"`

	// 移动距离阈值，单位KM
	// 注意：此字段可能返回 null，表示取不到有效值。
	Distance *int64 `json:"Distance,omitempty" name:"Distance"`

	// 信号强度阈值(-dbm）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignalStrength *int64 `json:"SignalStrength,omitempty" name:"SignalStrength"`

	// 盲点阈值天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LostDay *int64 `json:"LostDay,omitempty" name:"LostDay"`

	// 标签ID集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagIDs []*int64 `json:"TagIDs,omitempty" name:"TagIDs"`

	// 资费信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalePlan *string `json:"SalePlan,omitempty" name:"SalePlan"`
}

type RuleInfos struct {
	// 总量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*Rule `json:"List,omitempty" name:"List"`
}

type Tactic struct {
	// 策略ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 是否自动执行策略
	IsAuto *int64 `json:"IsAuto,omitempty" name:"IsAuto"`

	// 设备上报信息间隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PingInterval *int64 `json:"PingInterval,omitempty" name:"PingInterval"`

	// 是否开启弱信号检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWeak *int64 `json:"IsWeak,omitempty" name:"IsWeak"`

	// 弱信号阈值（-dbm）
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeakThreshold *int64 `json:"WeakThreshold,omitempty" name:"WeakThreshold"`

	// 忘了时延切换
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDelay *int64 `json:"IsDelay,omitempty" name:"IsDelay"`

	// 时延阈值（ms）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayThreshold *int64 `json:"DelayThreshold,omitempty" name:"DelayThreshold"`

	// 是否开启假信号检测
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFake *int64 `json:"IsFake,omitempty" name:"IsFake"`

	// 假信号检测IP字符串，用逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	FakeIP *string `json:"FakeIP,omitempty" name:"FakeIP"`

	// 假信号检测间隔（s）
	// 注意：此字段可能返回 null，表示取不到有效值。
	FakeInterval *int64 `json:"FakeInterval,omitempty" name:"FakeInterval"`

	// 是否开启网络制式检测
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNet *int64 `json:"IsNet,omitempty" name:"IsNet"`

	// 网络回落制式 1: 2G、 2: 3G 、 3: 2/3G
	// 注意：此字段可能返回 null，表示取不到有效值。
	Network *int64 `json:"Network,omitempty" name:"Network"`

	// 是否开启移动检测
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMove *int64 `json:"IsMove,omitempty" name:"IsMove"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否开启最优先运营商
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPriorityTele *int64 `json:"IsPriorityTele,omitempty" name:"IsPriorityTele"`

	// 最优先运营商 1 移动、 2 联通、 3 电信 4 上次在线运营商
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriorityTele *int64 `json:"PriorityTele,omitempty" name:"PriorityTele"`

	// 是否开启最不优先运营商
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsBottomTele *int64 `json:"IsBottomTele,omitempty" name:"IsBottomTele"`

	// 最不优先运营商 1 移动、 2 联通、 3 电信
	// 注意：此字段可能返回 null，表示取不到有效值。
	BottomTele *int64 `json:"BottomTele,omitempty" name:"BottomTele"`

	// 是否开启最优先信号选取策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsBestSignal *int64 `json:"IsBestSignal,omitempty" name:"IsBestSignal"`
}

type TacticInfos struct {
	// 总量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 策略列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*Tactic `json:"List,omitempty" name:"List"`
}

type Tag struct {
	// 标签名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 标签ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type TagInfos struct {
	// 总量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*Tag `json:"List,omitempty" name:"List"`
}

type TeleOperatorCard struct {
	// 开户时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountTime *string `json:"AccountTime,omitempty" name:"AccountTime"`

	// 激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveTime *string `json:"ActiveTime,omitempty" name:"ActiveTime"`

	// 运营商ICCID
	ICCID *string `json:"ICCID,omitempty" name:"ICCID"`

	// 云兔卡ID
	LinkID *int64 `json:"LinkID,omitempty" name:"LinkID"`

	// 电话号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msisdn *string `json:"Msisdn,omitempty" name:"Msisdn"`

	// 移动用户识别码
	// 注意：此字段可能返回 null，表示取不到有效值。
	IMSI *string `json:"IMSI,omitempty" name:"IMSI"`

	// 运营商: 1 移动 2 联通 3 电信
	TeleOperator *int64 `json:"TeleOperator,omitempty" name:"TeleOperator"`
}