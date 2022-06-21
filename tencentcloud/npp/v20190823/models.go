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

package v20190823

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CallBackCdr struct {
	// 呼叫通话 ID
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 主叫号码
	Src *string `json:"Src,omitempty" name:"Src"`

	// 被叫号码
	Dst *string `json:"Dst,omitempty" name:"Dst"`

	// 主叫呼叫开始时间
	StartSrcCallTime *string `json:"StartSrcCallTime,omitempty" name:"StartSrcCallTime"`

	// 主叫响铃开始时间
	StartSrcRingTime *string `json:"StartSrcRingTime,omitempty" name:"StartSrcRingTime"`

	// 主叫接听时间
	SrcAcceptTime *string `json:"SrcAcceptTime,omitempty" name:"SrcAcceptTime"`

	// 被叫呼叫开始时间
	StartDstCallTime *string `json:"StartDstCallTime,omitempty" name:"StartDstCallTime"`

	// 被叫响铃开始时间
	StartDstRingTime *string `json:"StartDstRingTime,omitempty" name:"StartDstRingTime"`

	// 被叫接听时间
	DstAcceptTime *string `json:"DstAcceptTime,omitempty" name:"DstAcceptTime"`

	// 用户挂机通话结束时间
	EndCallTime *string `json:"EndCallTime,omitempty" name:"EndCallTime"`

	// 通话最后状态：0：未知状态 1：正常通话 2：主叫未接 3：主叫接听，被叫未接 4：主叫未接通 5：被叫未接通
	CallEndStatus *string `json:"CallEndStatus,omitempty" name:"CallEndStatus"`

	// 通话计费时间
	Duration *string `json:"Duration,omitempty" name:"Duration"`

	// 录音 URL，如果不录音或录音失败，该值为空
	RecordUrl *string `json:"RecordUrl,omitempty" name:"RecordUrl"`

	// 通话类型(1: VOIP 2:IP TO PSTN 3: PSTN TO PSTN)，如果话单中没有该字段，默认值为回拨 3 (PSTN TO PSTN)
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallType *string `json:"CallType,omitempty" name:"CallType"`

	// 同回拨请求中的 bizId，如果回拨请求中带 bizId 会有该字段返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizId *string `json:"BizId,omitempty" name:"BizId"`

	// 订单 ID,最大长度不超过 64 个字节，对于一些有订单状态 App 相关应用（如达人帮接入 App 应用)，该字段只在帐单中带上，其它回调不附带该字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

type CallBackPhoneCode struct {
	// 国家码，统一以 00 开头
	Nation *string `json:"Nation,omitempty" name:"Nation"`

	// 号码（固话区号前加 0，如075586013388）
	Phone *string `json:"Phone,omitempty" name:"Phone"`
}

// Predefined struct for user
type CreateCallBackRequestParams struct {
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 主叫号码(必须为 11 位手机号，号码前加 0086，如 008613631686024)
	Src *string `json:"Src,omitempty" name:"Src"`

	// 被叫号码(必须为 11 位手机或固话号码,号码前加 0086，如008613631686024，固话如：0086075586013388)
	Dst *string `json:"Dst,omitempty" name:"Dst"`

	// 主叫显示系统分配的固话号码，如不填显示随机分配号码
	SrcDisplayNum *string `json:"SrcDisplayNum,omitempty" name:"SrcDisplayNum"`

	// 被叫显示系统分配的固话号码，如不填显示随机分配号码
	DstDisplayNum *string `json:"DstDisplayNum,omitempty" name:"DstDisplayNum"`

	// 是否录音，0 表示不录音，1 表示录音。默认为不录音
	Record *string `json:"Record,omitempty" name:"Record"`

	// 允许最大通话时间，不填默认为 30 分钟（单位：分钟）
	MaxAllowTime *string `json:"MaxAllowTime,omitempty" name:"MaxAllowTime"`

	// 主叫发起呼叫状态：1 被叫发起呼叫状态：256 主叫响铃状态：2 被叫响铃状态：512 主叫接听状态：4 被叫接听状态：1024 主叫拒绝接听状态：8 被叫拒绝接听状态：2048 主叫正常挂机状态：16 被叫正常挂机状态：4096 主叫呼叫异常：32 被叫呼叫异常：8192
	// 例如：当值为 0：表示所有状态不需要推送；当值为 4：表示只要推送主叫接听状态；当值为 16191：表示所有状态都需要推送(上面所有值和)
	StatusFlag *string `json:"StatusFlag,omitempty" name:"StatusFlag"`

	// 状态回调通知地址，正式环境可以配置默认推送地址
	StatusUrl *string `json:"StatusUrl,omitempty" name:"StatusUrl"`

	// 话单回调通知地址，正式环境可以配置默认推送地址
	HangupUrl *string `json:"HangupUrl,omitempty" name:"HangupUrl"`

	// 录单 URL 回调通知地址，正式环境可以配置默认推送地址
	RecordUrl *string `json:"RecordUrl,omitempty" name:"RecordUrl"`

	// 业务应用 key，业务用该 key 可以区分内部业务或客户产品等，该 key 需保证在该 appId 下全局唯一，最大长度不超过 64 个字节，bizId 只能包含数字、字母。
	BizId *string `json:"BizId,omitempty" name:"BizId"`

	// 最后一次呼叫 callId，带上该字段以后，平台会参考该 callId 分配线路，优先不分配该 callId 通话线路（注：谨慎使用，这个会影响线路调度）
	LastCallId *string `json:"LastCallId,omitempty" name:"LastCallId"`

	// 结构体，主叫呼叫预处理操作，根据不同操作确认是否呼通被叫。如需使用，本结构体需要与 keyList 结构体配合使用，此时这两个参数都为必填项
	PreCallerHandle *RreCallerHandle `json:"PreCallerHandle,omitempty" name:"PreCallerHandle"`

	// 订单 ID，最大长度不超过64个字节，对于一些有订单状态 App 相关应用使用（如达人帮接入 App 应用)，该字段只在帐单中带上，其它回调不附带该字段
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

type CreateCallBackRequest struct {
	*tchttp.BaseRequest
	
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 主叫号码(必须为 11 位手机号，号码前加 0086，如 008613631686024)
	Src *string `json:"Src,omitempty" name:"Src"`

	// 被叫号码(必须为 11 位手机或固话号码,号码前加 0086，如008613631686024，固话如：0086075586013388)
	Dst *string `json:"Dst,omitempty" name:"Dst"`

	// 主叫显示系统分配的固话号码，如不填显示随机分配号码
	SrcDisplayNum *string `json:"SrcDisplayNum,omitempty" name:"SrcDisplayNum"`

	// 被叫显示系统分配的固话号码，如不填显示随机分配号码
	DstDisplayNum *string `json:"DstDisplayNum,omitempty" name:"DstDisplayNum"`

	// 是否录音，0 表示不录音，1 表示录音。默认为不录音
	Record *string `json:"Record,omitempty" name:"Record"`

	// 允许最大通话时间，不填默认为 30 分钟（单位：分钟）
	MaxAllowTime *string `json:"MaxAllowTime,omitempty" name:"MaxAllowTime"`

	// 主叫发起呼叫状态：1 被叫发起呼叫状态：256 主叫响铃状态：2 被叫响铃状态：512 主叫接听状态：4 被叫接听状态：1024 主叫拒绝接听状态：8 被叫拒绝接听状态：2048 主叫正常挂机状态：16 被叫正常挂机状态：4096 主叫呼叫异常：32 被叫呼叫异常：8192
	// 例如：当值为 0：表示所有状态不需要推送；当值为 4：表示只要推送主叫接听状态；当值为 16191：表示所有状态都需要推送(上面所有值和)
	StatusFlag *string `json:"StatusFlag,omitempty" name:"StatusFlag"`

	// 状态回调通知地址，正式环境可以配置默认推送地址
	StatusUrl *string `json:"StatusUrl,omitempty" name:"StatusUrl"`

	// 话单回调通知地址，正式环境可以配置默认推送地址
	HangupUrl *string `json:"HangupUrl,omitempty" name:"HangupUrl"`

	// 录单 URL 回调通知地址，正式环境可以配置默认推送地址
	RecordUrl *string `json:"RecordUrl,omitempty" name:"RecordUrl"`

	// 业务应用 key，业务用该 key 可以区分内部业务或客户产品等，该 key 需保证在该 appId 下全局唯一，最大长度不超过 64 个字节，bizId 只能包含数字、字母。
	BizId *string `json:"BizId,omitempty" name:"BizId"`

	// 最后一次呼叫 callId，带上该字段以后，平台会参考该 callId 分配线路，优先不分配该 callId 通话线路（注：谨慎使用，这个会影响线路调度）
	LastCallId *string `json:"LastCallId,omitempty" name:"LastCallId"`

	// 结构体，主叫呼叫预处理操作，根据不同操作确认是否呼通被叫。如需使用，本结构体需要与 keyList 结构体配合使用，此时这两个参数都为必填项
	PreCallerHandle *RreCallerHandle `json:"PreCallerHandle,omitempty" name:"PreCallerHandle"`

	// 订单 ID，最大长度不超过64个字节，对于一些有订单状态 App 相关应用使用（如达人帮接入 App 应用)，该字段只在帐单中带上，其它回调不附带该字段
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

func (r *CreateCallBackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCallBackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizAppId")
	delete(f, "Src")
	delete(f, "Dst")
	delete(f, "SrcDisplayNum")
	delete(f, "DstDisplayNum")
	delete(f, "Record")
	delete(f, "MaxAllowTime")
	delete(f, "StatusFlag")
	delete(f, "StatusUrl")
	delete(f, "HangupUrl")
	delete(f, "RecordUrl")
	delete(f, "BizId")
	delete(f, "LastCallId")
	delete(f, "PreCallerHandle")
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCallBackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCallBackResponseParams struct {
	// 话单id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 主叫显示系统分配的固话号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcDisplayNum *string `json:"SrcDisplayNum,omitempty" name:"SrcDisplayNum"`

	// 被叫显示系统分配的固话号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstDisplayNum *string `json:"DstDisplayNum,omitempty" name:"DstDisplayNum"`

	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 错误原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCallBackResponse struct {
	*tchttp.BaseResponse
	Response *CreateCallBackResponseParams `json:"Response"`
}

func (r *CreateCallBackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCallBackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DelVirtualNumRequestParams struct {
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 双方号码 + 中间号绑定 ID，该 ID 全局唯一
	BindId *string `json:"BindId,omitempty" name:"BindId"`

	// 应用二级业务 ID，bizId 需保证在该 appId 下全局唯一，最大长度不超过 16 个字节。
	BizId *string `json:"BizId,omitempty" name:"BizId"`
}

type DelVirtualNumRequest struct {
	*tchttp.BaseRequest
	
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 双方号码 + 中间号绑定 ID，该 ID 全局唯一
	BindId *string `json:"BindId,omitempty" name:"BindId"`

	// 应用二级业务 ID，bizId 需保证在该 appId 下全局唯一，最大长度不超过 16 个字节。
	BizId *string `json:"BizId,omitempty" name:"BizId"`
}

func (r *DelVirtualNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DelVirtualNumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizAppId")
	delete(f, "BindId")
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DelVirtualNumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DelVirtualNumResponseParams struct {
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 绑定 ID，该 ID 全局唯一
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindId *string `json:"BindId,omitempty" name:"BindId"`

	// 中间号还剩引用计数，如果计数为 0 会解绑
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefLeftNum *string `json:"RefLeftNum,omitempty" name:"RefLeftNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DelVirtualNumResponse struct {
	*tchttp.BaseResponse
	Response *DelVirtualNumResponseParams `json:"Response"`
}

func (r *DelVirtualNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DelVirtualNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCallBackRequestParams struct {
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 回拨请求响应中返回的 callId
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 0：不管通话状态直接拆线（默认) 1：主叫响铃以后状态不拆线 2：主叫接听以后状态不拆线 3：被叫响铃以后状态不拆线 4：被叫接听以后状态不拆线
	CancelFlag *string `json:"CancelFlag,omitempty" name:"CancelFlag"`
}

type DeleteCallBackRequest struct {
	*tchttp.BaseRequest
	
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 回拨请求响应中返回的 callId
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 0：不管通话状态直接拆线（默认) 1：主叫响铃以后状态不拆线 2：主叫接听以后状态不拆线 3：被叫响铃以后状态不拆线 4：被叫接听以后状态不拆线
	CancelFlag *string `json:"CancelFlag,omitempty" name:"CancelFlag"`
}

func (r *DeleteCallBackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCallBackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizAppId")
	delete(f, "CallId")
	delete(f, "CancelFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCallBackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCallBackResponseParams struct {
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 错误原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 话单id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCallBackResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCallBackResponseParams `json:"Response"`
}

func (r *DeleteCallBackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCallBackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallBackCdrRequestParams struct {
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 回拨请求响应中返回的 callId，按 callId 查询该话单详细信息
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 查询主叫用户产生的呼叫话单，如填空表示拉取这个时间段所有话单
	Src *string `json:"Src,omitempty" name:"Src"`

	// 话单开始时间戳
	StartTimeStamp *string `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 话单结束时间戳
	EndTimeStamp *string `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`
}

type DescribeCallBackCdrRequest struct {
	*tchttp.BaseRequest
	
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 回拨请求响应中返回的 callId，按 callId 查询该话单详细信息
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 查询主叫用户产生的呼叫话单，如填空表示拉取这个时间段所有话单
	Src *string `json:"Src,omitempty" name:"Src"`

	// 话单开始时间戳
	StartTimeStamp *string `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 话单结束时间戳
	EndTimeStamp *string `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`
}

func (r *DescribeCallBackCdrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallBackCdrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizAppId")
	delete(f, "CallId")
	delete(f, "Src")
	delete(f, "StartTimeStamp")
	delete(f, "EndTimeStamp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallBackCdrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallBackCdrResponseParams struct {
	// 话单列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cdr []*CallBackCdr `json:"Cdr,omitempty" name:"Cdr"`

	// 偏移
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 错误原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCallBackCdrResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallBackCdrResponseParams `json:"Response"`
}

func (r *DescribeCallBackCdrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallBackCdrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallBackStatusRequestParams struct {
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 回拨请求响应中返回的 callId
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 主叫号码
	Src *string `json:"Src,omitempty" name:"Src"`

	// 被叫号码
	Dst *string `json:"Dst,omitempty" name:"Dst"`

	// 通话最后状态：0：未知状态 1：主叫响铃中 2：主叫接听 3：被叫响铃中 4：正常通话中 5：通话结束
	CallStatus *string `json:"CallStatus,omitempty" name:"CallStatus"`
}

type DescribeCallBackStatusRequest struct {
	*tchttp.BaseRequest
	
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 回拨请求响应中返回的 callId
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 主叫号码
	Src *string `json:"Src,omitempty" name:"Src"`

	// 被叫号码
	Dst *string `json:"Dst,omitempty" name:"Dst"`

	// 通话最后状态：0：未知状态 1：主叫响铃中 2：主叫接听 3：被叫响铃中 4：正常通话中 5：通话结束
	CallStatus *string `json:"CallStatus,omitempty" name:"CallStatus"`
}

func (r *DescribeCallBackStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallBackStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizAppId")
	delete(f, "CallId")
	delete(f, "Src")
	delete(f, "Dst")
	delete(f, "CallStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallBackStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallBackStatusResponseParams struct {
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 错误信息
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 业务appid
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 回拨请求响应中返回的 callId
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 主叫号码
	Src *string `json:"Src,omitempty" name:"Src"`

	// 被叫号码
	Dst *string `json:"Dst,omitempty" name:"Dst"`

	// 通话最后状态：0：未知状态 1：主叫响铃中 2：主叫接听 3：被叫响铃中 4：正常通话中 5：通话结束
	CallStatus *string `json:"CallStatus,omitempty" name:"CallStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCallBackStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallBackStatusResponseParams `json:"Response"`
}

func (r *DescribeCallBackStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallBackStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallerDisplayListRequestParams struct {
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`
}

type DescribeCallerDisplayListRequest struct {
	*tchttp.BaseRequest
	
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`
}

func (r *DescribeCallerDisplayListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallerDisplayListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallerDisplayListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallerDisplayListResponseParams struct {
	// appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 主叫显号号码集合，codeList[0...*] 结构体数组，如果业务是主被叫互显，该字段为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeList []*CallBackPhoneCode `json:"CodeList,omitempty" name:"CodeList"`

	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 错误原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCallerDisplayListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallerDisplayListResponseParams `json:"Response"`
}

func (r *DescribeCallerDisplayListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallerDisplayListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type Get400CdrRequestParams struct {
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 通话唯一标识 callId，即直拨呼叫响应中返回的 callId
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 查询主叫用户产生的呼叫话单（0086开头），设置为空表示拉取该时间段的所有话单
	Src *string `json:"Src,omitempty" name:"Src"`

	// 话单开始时间戳
	StartTimeStamp *string `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 话单结束时间戳
	EndTimeStamp *string `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`
}

type Get400CdrRequest struct {
	*tchttp.BaseRequest
	
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 通话唯一标识 callId，即直拨呼叫响应中返回的 callId
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 查询主叫用户产生的呼叫话单（0086开头），设置为空表示拉取该时间段的所有话单
	Src *string `json:"Src,omitempty" name:"Src"`

	// 话单开始时间戳
	StartTimeStamp *string `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 话单结束时间戳
	EndTimeStamp *string `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`
}

func (r *Get400CdrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *Get400CdrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizAppId")
	delete(f, "CallId")
	delete(f, "Src")
	delete(f, "StartTimeStamp")
	delete(f, "EndTimeStamp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "Get400CdrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type Get400CdrResponseParams struct {
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 错误原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 偏移
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 话单列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cdr []*VirturalNumCdr `json:"Cdr,omitempty" name:"Cdr"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type Get400CdrResponse struct {
	*tchttp.BaseResponse
	Response *Get400CdrResponseParams `json:"Response"`
}

func (r *Get400CdrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *Get400CdrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetVirtualNumRequestParams struct {
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 被叫号码(号码前加 0086，如 008613631686024)
	Dst *string `json:"Dst,omitempty" name:"Dst"`

	// 主叫号码(号码前加 0086，如 008613631686024)，xb 模式下是不用填写，axb 模式下是必选
	Src *string `json:"Src,omitempty" name:"Src"`

	// {“accreditList”:[“008613631686024”,”008612345678910”]}，主要用于 N-1 场景，号码绑定非共享是独占型，指定了 dst 独占中间号绑定，accreditList 表示这个列表成员可以拨打 dst 绑 定的中间号，默认值为空，表示所有号码都可以拨打独占型中间号绑定，最大集合不允许超过 30 个，仅适用于xb模式
	AccreditList []*string `json:"AccreditList,omitempty" name:"AccreditList"`

	// 指定中间号（格式：008617013541251），如果该中间号已被使用则返回绑定失败。如果不带该字段则由腾讯侧从号码池里自动分配
	AssignVirtualNum *string `json:"AssignVirtualNum,omitempty" name:"AssignVirtualNum"`

	// 是否录音，0表示不录音，1表示录音。默认为不录音，注意如果需要录音回调，通话完成后需要等待一段时间，收到录音回调之后，再解绑中间号。
	Record *string `json:"Record,omitempty" name:"Record"`

	// 主被叫显号号码归属地，指定该参数说明显号归属该城市，如果没有该城市号码会随机选取一个城市或者后台配置返回107，返回码详见 《腾讯-中间号-城市id.xlsx》
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// 应用二级业务 ID，bizId 需保证在该 appId 下全局唯一，最大长度不超过 16 个字节。
	BizId *string `json:"BizId,omitempty" name:"BizId"`

	// 号码最大绑定时间，不填默认为 24 小时，最长绑定时间是168小时，单位秒
	MaxAssignTime *string `json:"MaxAssignTime,omitempty" name:"MaxAssignTime"`

	// 主叫发起呼叫状态：1
	// 被叫发起呼叫状态：256
	// 主叫响铃状态：2
	// 被叫响铃状态：512
	// 主叫接听状态：4
	// 被叫接听状态：1024
	// 主叫拒绝接听状态：8
	// 被叫拒绝接听状态：2048
	// 主叫正常挂机状态：16
	// 被叫正常挂机状态：4096
	// 主叫呼叫异常：32
	// 被叫呼叫异常：8192
	// 
	// 例如：
	// 值为 0：表示所有状态不需要推送
	// 值为 4：表示只要推送主叫接听状态
	// 值为 16191：表示所有状态都需要推送（上面所有值和）
	StatusFlag *string `json:"StatusFlag,omitempty" name:"StatusFlag"`

	// 请填写statusFlag并设置值，状态回调通知地址，正式环境可以配置默认推送地址
	StatusUrl *string `json:"StatusUrl,omitempty" name:"StatusUrl"`

	// 话单回调通知地址，正式环境可以配置默认推送地址
	HangupUrl *string `json:"HangupUrl,omitempty" name:"HangupUrl"`

	// 录单 URL 回调通知地址，正式环境可以配置默认推送地址
	RecordUrl *string `json:"RecordUrl,omitempty" name:"RecordUrl"`
}

type GetVirtualNumRequest struct {
	*tchttp.BaseRequest
	
	// 业务appid
	BizAppId *string `json:"BizAppId,omitempty" name:"BizAppId"`

	// 被叫号码(号码前加 0086，如 008613631686024)
	Dst *string `json:"Dst,omitempty" name:"Dst"`

	// 主叫号码(号码前加 0086，如 008613631686024)，xb 模式下是不用填写，axb 模式下是必选
	Src *string `json:"Src,omitempty" name:"Src"`

	// {“accreditList”:[“008613631686024”,”008612345678910”]}，主要用于 N-1 场景，号码绑定非共享是独占型，指定了 dst 独占中间号绑定，accreditList 表示这个列表成员可以拨打 dst 绑 定的中间号，默认值为空，表示所有号码都可以拨打独占型中间号绑定，最大集合不允许超过 30 个，仅适用于xb模式
	AccreditList []*string `json:"AccreditList,omitempty" name:"AccreditList"`

	// 指定中间号（格式：008617013541251），如果该中间号已被使用则返回绑定失败。如果不带该字段则由腾讯侧从号码池里自动分配
	AssignVirtualNum *string `json:"AssignVirtualNum,omitempty" name:"AssignVirtualNum"`

	// 是否录音，0表示不录音，1表示录音。默认为不录音，注意如果需要录音回调，通话完成后需要等待一段时间，收到录音回调之后，再解绑中间号。
	Record *string `json:"Record,omitempty" name:"Record"`

	// 主被叫显号号码归属地，指定该参数说明显号归属该城市，如果没有该城市号码会随机选取一个城市或者后台配置返回107，返回码详见 《腾讯-中间号-城市id.xlsx》
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// 应用二级业务 ID，bizId 需保证在该 appId 下全局唯一，最大长度不超过 16 个字节。
	BizId *string `json:"BizId,omitempty" name:"BizId"`

	// 号码最大绑定时间，不填默认为 24 小时，最长绑定时间是168小时，单位秒
	MaxAssignTime *string `json:"MaxAssignTime,omitempty" name:"MaxAssignTime"`

	// 主叫发起呼叫状态：1
	// 被叫发起呼叫状态：256
	// 主叫响铃状态：2
	// 被叫响铃状态：512
	// 主叫接听状态：4
	// 被叫接听状态：1024
	// 主叫拒绝接听状态：8
	// 被叫拒绝接听状态：2048
	// 主叫正常挂机状态：16
	// 被叫正常挂机状态：4096
	// 主叫呼叫异常：32
	// 被叫呼叫异常：8192
	// 
	// 例如：
	// 值为 0：表示所有状态不需要推送
	// 值为 4：表示只要推送主叫接听状态
	// 值为 16191：表示所有状态都需要推送（上面所有值和）
	StatusFlag *string `json:"StatusFlag,omitempty" name:"StatusFlag"`

	// 请填写statusFlag并设置值，状态回调通知地址，正式环境可以配置默认推送地址
	StatusUrl *string `json:"StatusUrl,omitempty" name:"StatusUrl"`

	// 话单回调通知地址，正式环境可以配置默认推送地址
	HangupUrl *string `json:"HangupUrl,omitempty" name:"HangupUrl"`

	// 录单 URL 回调通知地址，正式环境可以配置默认推送地址
	RecordUrl *string `json:"RecordUrl,omitempty" name:"RecordUrl"`
}

func (r *GetVirtualNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVirtualNumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizAppId")
	delete(f, "Dst")
	delete(f, "Src")
	delete(f, "AccreditList")
	delete(f, "AssignVirtualNum")
	delete(f, "Record")
	delete(f, "CityId")
	delete(f, "BizId")
	delete(f, "MaxAssignTime")
	delete(f, "StatusFlag")
	delete(f, "StatusUrl")
	delete(f, "HangupUrl")
	delete(f, "RecordUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetVirtualNumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetVirtualNumResponseParams struct {
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 绑定 ID，该 ID 全局唯一
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindId *string `json:"BindId,omitempty" name:"BindId"`

	// 中间号还剩引用计数，如果计数为 0 会解绑
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefNum *string `json:"RefNum,omitempty" name:"RefNum"`

	// 中间号
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualNum *string `json:"VirtualNum,omitempty" name:"VirtualNum"`

	// 错误原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetVirtualNumResponse struct {
	*tchttp.BaseResponse
	Response *GetVirtualNumResponseParams `json:"Response"`
}

func (r *GetVirtualNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVirtualNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyList struct {
	// 用户按键（0-9、*、#、A-D)
	Key *string `json:"Key,omitempty" name:"Key"`

	// 1: 呼通被叫 2：interruptPrompt 重播提示 3：拆线
	Operate *string `json:"Operate,omitempty" name:"Operate"`
}

type RreCallerHandle struct {
	// 呼叫主叫以后，给主叫用户的语音提示，播放该提示时用户所有按键无效
	ReadPrompt *string `json:"ReadPrompt,omitempty" name:"ReadPrompt"`

	// 可中断提示，播放该提示时，用户可以按键
	InterruptPrompt *string `json:"InterruptPrompt,omitempty" name:"InterruptPrompt"`

	// 对应按键操作,如果没有结构体里定义按键操作用户按键以后都从 interruptPrompt 重新播放
	KeyList []*KeyList `json:"KeyList,omitempty" name:"KeyList"`

	// 最多重复播放次数，超过该次数拆线
	RepeatTimes *string `json:"RepeatTimes,omitempty" name:"RepeatTimes"`

	// 用户按键回调通知地址，如果为空不回调
	KeyPressUrl *string `json:"KeyPressUrl,omitempty" name:"KeyPressUrl"`

	// 提示音男声女声：1女声，2男声。默认女声
	PromptGender *string `json:"PromptGender,omitempty" name:"PromptGender"`
}

type VirturalNumCdr struct {
	// 呼叫通话 ID
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 双方号码 + 中间号绑定 ID，该 ID 全局唯一
	BindId *string `json:"BindId,omitempty" name:"BindId"`

	// 主叫号码
	Src *string `json:"Src,omitempty" name:"Src"`

	// 被叫号码
	Dst *string `json:"Dst,omitempty" name:"Dst"`

	// 主叫通讯录直拨虚拟保护号码
	DstVirtualNum *string `json:"DstVirtualNum,omitempty" name:"DstVirtualNum"`

	// 虚拟保护号码平台收到呼叫时间
	CallCenterAcceptTime *string `json:"CallCenterAcceptTime,omitempty" name:"CallCenterAcceptTime"`

	// 被叫呼叫开始时间
	StartDstCallTime *string `json:"StartDstCallTime,omitempty" name:"StartDstCallTime"`

	// 被叫响铃开始时间
	StartDstRingTime *string `json:"StartDstRingTime,omitempty" name:"StartDstRingTime"`

	// 被叫接听时间
	DstAcceptTime *string `json:"DstAcceptTime,omitempty" name:"DstAcceptTime"`

	// 用户挂机通话结束时间
	EndCallTime *string `json:"EndCallTime,omitempty" name:"EndCallTime"`

	// 通话最后状态：0：未知状态 1：正常通话 2：查询呼叫转移被叫号异常 3：未接通 4：未接听 5：拒接挂断 6：关机 7：空号 8：通话中 9：欠费 10：运营商线路或平台异常
	CallEndStatus *string `json:"CallEndStatus,omitempty" name:"CallEndStatus"`

	// 主叫接通虚拟保护号码到通话结束通话时间
	SrcDuration *string `json:"SrcDuration,omitempty" name:"SrcDuration"`

	// 呼叫转接被叫接通到通话结束通话时间
	DstDuration *string `json:"DstDuration,omitempty" name:"DstDuration"`

	// 录音 URL，如果不录音或录音失败，该值为空
	RecordUrl *string `json:"RecordUrl,omitempty" name:"RecordUrl"`
}