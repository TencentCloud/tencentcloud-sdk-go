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

package v20180625

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BindL4Backend struct {
	// 待绑定的主机端口，可选值1~65535。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 待绑定的黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 待绑定的主机权重，可选值0~100。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 自定义探测的主机端口，可选值1~65535。（需要监听器开启自定义健康检查）
	ProbePort *int64 `json:"ProbePort,omitempty" name:"ProbePort"`
}

// Predefined struct for user
type BindL4BackendsRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 四层监听器实例ID，可通过接口DescribeL4Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待绑定的主机信息。可以绑定多个主机端口。目前一个四层监听器下面最多允许绑定255个主机端口。
	BackendSet []*BindL4Backend `json:"BackendSet,omitempty" name:"BackendSet"`

	// 绑定类型。0：物理机 1：虚拟机 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

type BindL4BackendsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 四层监听器实例ID，可通过接口DescribeL4Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待绑定的主机信息。可以绑定多个主机端口。目前一个四层监听器下面最多允许绑定255个主机端口。
	BackendSet []*BindL4Backend `json:"BackendSet,omitempty" name:"BackendSet"`

	// 绑定类型。0：物理机 1：虚拟机 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

func (r *BindL4BackendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindL4BackendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "BackendSet")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindL4BackendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindL4BackendsResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindL4BackendsResponse struct {
	*tchttp.BaseResponse
	Response *BindL4BackendsResponseParams `json:"Response"`
}

func (r *BindL4BackendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindL4BackendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindL7Backend struct {
	// 待绑定的主机端口，可选值1~65535。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 待绑定的主机权重，可选值0~100。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

// Predefined struct for user
type BindL7BackendsRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID，可通过接口DescribeL7Rules查询。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 待绑定的主机信息。可以绑定多个主机端口。目前一个七层转发路径下面最多允许绑定255个主机端口。
	BackendSet []*BindL7Backend `json:"BackendSet,omitempty" name:"BackendSet"`

	// 绑定类型。0：物理机，1：虚拟机 2：半托管机器。
	BindType *uint64 `json:"BindType,omitempty" name:"BindType"`
}

type BindL7BackendsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID，可通过接口DescribeL7Rules查询。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 待绑定的主机信息。可以绑定多个主机端口。目前一个七层转发路径下面最多允许绑定255个主机端口。
	BackendSet []*BindL7Backend `json:"BackendSet,omitempty" name:"BackendSet"`

	// 绑定类型。0：物理机，1：虚拟机 2：半托管机器。
	BindType *uint64 `json:"BindType,omitempty" name:"BindType"`
}

func (r *BindL7BackendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindL7BackendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "DomainId")
	delete(f, "LocationId")
	delete(f, "BackendSet")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindL7BackendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindL7BackendsResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindL7BackendsResponse struct {
	*tchttp.BaseResponse
	Response *BindL7BackendsResponseParams `json:"Response"`
}

func (r *BindL7BackendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindL7BackendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindTrafficMirrorListenersRequestParams struct {
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 七层监听器实例ID数组，可通过接口DescribeL7Listeners查询。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

type BindTrafficMirrorListenersRequest struct {
	*tchttp.BaseRequest
	
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 七层监听器实例ID数组，可通过接口DescribeL7Listeners查询。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

func (r *BindTrafficMirrorListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindTrafficMirrorListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficMirrorId")
	delete(f, "ListenerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindTrafficMirrorListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindTrafficMirrorListenersResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindTrafficMirrorListenersResponse struct {
	*tchttp.BaseResponse
	Response *BindTrafficMirrorListenersResponseParams `json:"Response"`
}

func (r *BindTrafficMirrorListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindTrafficMirrorListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindTrafficMirrorReceiver struct {
	// 待绑定的主机端口，可选值1~65535。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 待绑定的主机实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 待绑定的主机权重，可选值0~100。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

// Predefined struct for user
type BindTrafficMirrorReceiversRequestParams struct {
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 待绑定的黑石物理机信息数组。
	ReceiverSet []*BindTrafficMirrorReceiver `json:"ReceiverSet,omitempty" name:"ReceiverSet"`
}

type BindTrafficMirrorReceiversRequest struct {
	*tchttp.BaseRequest
	
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 待绑定的黑石物理机信息数组。
	ReceiverSet []*BindTrafficMirrorReceiver `json:"ReceiverSet,omitempty" name:"ReceiverSet"`
}

func (r *BindTrafficMirrorReceiversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindTrafficMirrorReceiversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficMirrorId")
	delete(f, "ReceiverSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindTrafficMirrorReceiversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindTrafficMirrorReceiversResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindTrafficMirrorReceiversResponse struct {
	*tchttp.BaseResponse
	Response *BindTrafficMirrorReceiversResponseParams `json:"Response"`
}

func (r *BindTrafficMirrorReceiversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindTrafficMirrorReceiversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CertDetailLoadBalancer struct {
	// 黑石负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 黑石负载均衡实例名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 该黑石负载均衡所在的VpcId。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 该黑石负载均衡所在的regionId。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

type CreateL4Listener struct {
	// 监听器监听端口，可选值1~65535。
	LoadBalancerPort *int64 `json:"LoadBalancerPort,omitempty" name:"LoadBalancerPort"`

	// 监听器协议类型，可选值tcp，udp。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器的会话保持时间，单位：秒。可选值：900~3600,不传表示不开启会话保持。
	SessionExpire *int64 `json:"SessionExpire,omitempty" name:"SessionExpire"`

	// 是否开启健康检查：1（开启）、0（关闭）。默认值0，表示关闭。
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 健康检查的响应超时时间，可选值：2-60，默认值：2，单位:秒。<br><font color="red">响应超时时间要小于检查间隔时间。</font>
	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 健康检查检查间隔时间，默认值：5，可选值：5-300，单位：秒。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 健康阈值，默认值：3，表示当连续探测三次健康则表示该转发正常，可选值：2-10，单位：次。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 不健康阈值，默认值：3，表示当连续探测三次不健康则表示该转发不正常，可选值：2-10，单位：次。
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 监听器最大带宽值，用于计费模式为固定带宽计费，可选值：0-1000，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 是否开启自定义健康检查：1（开启）、0（关闭）。默认值0，表示关闭。（该字段在健康检查开启的情况下才生效）
	CustomHealthSwitch *int64 `json:"CustomHealthSwitch,omitempty" name:"CustomHealthSwitch"`

	// 自定义健康探测内容类型，可选值：text（文本）、hexadecimal（十六进制）。
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 探测内容类型为文本方式时，针对请求文本中换行替换方式。可选值：1（替换为LF）、2（替换为CR）、3（替换为LF+CR）。
	LineSeparatorType *int64 `json:"LineSeparatorType,omitempty" name:"LineSeparatorType"`

	// 自定义探测请求内容。
	HealthRequest *string `json:"HealthRequest,omitempty" name:"HealthRequest"`

	// 自定义探测返回内容。
	HealthResponse *string `json:"HealthResponse,omitempty" name:"HealthResponse"`

	// 是否开启toa。可选值：0（关闭）、1（开启），默认关闭。（该字段在负载均衡为fullnat类型下才生效）
	ToaFlag *int64 `json:"ToaFlag,omitempty" name:"ToaFlag"`
}

// Predefined struct for user
type CreateL4ListenersRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 监听器信息数组，可以创建多个监听器。目前一个负载均衡下面最多允许创建50个监听器
	ListenerSet []*CreateL4Listener `json:"ListenerSet,omitempty" name:"ListenerSet"`
}

type CreateL4ListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 监听器信息数组，可以创建多个监听器。目前一个负载均衡下面最多允许创建50个监听器
	ListenerSet []*CreateL4Listener `json:"ListenerSet,omitempty" name:"ListenerSet"`
}

func (r *CreateL4ListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4ListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL4ListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL4ListenersResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateL4ListenersResponse struct {
	*tchttp.BaseResponse
	Response *CreateL4ListenersResponseParams `json:"Response"`
}

func (r *CreateL4ListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL4ListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7Listener struct {
	// 七层监听器端口，可选值1~65535。
	LoadBalancerPort *int64 `json:"LoadBalancerPort,omitempty" name:"LoadBalancerPort"`

	// 七层监听器协议类型，可选值：http,https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 七层监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 认证方式：0（不认证，用于http），1（单向认证，用于https），2（双向认证，用于https）。当创建的是https类型的监听器时，此值必选。
	SslMode *int64 `json:"SslMode,omitempty" name:"SslMode"`

	// 服务端证书ID。当创建的是https类型的监听器时，此值必选。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 服务端证书名称。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 服务端证书内容。
	CertContent *string `json:"CertContent,omitempty" name:"CertContent"`

	// 服务端证书密钥。
	CertKey *string `json:"CertKey,omitempty" name:"CertKey"`

	// 客户端证书ID。
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// 客户端证书名称。
	CertCaName *string `json:"CertCaName,omitempty" name:"CertCaName"`

	// 客户端证书内容。
	CertCaContent *string `json:"CertCaContent,omitempty" name:"CertCaContent"`

	// 用于计费模式为固定带宽计费，指定监听器最大带宽值，可选值：0-1000，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 转发协议。当Protocol为https时并且SslMode为1或2时，有意义。可选的值为0：https，1：spdy，2：http2，3：spdy+http2。
	ForwardProtocol *int64 `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`
}

// Predefined struct for user
type CreateL7ListenersRequestParams struct {
	// 负载均衡实例ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器信息数组，可以创建多个七层监听器。目前一个负载均衡下面最多允许创建50个七层监听器。
	ListenerSet []*CreateL7Listener `json:"ListenerSet,omitempty" name:"ListenerSet"`
}

type CreateL7ListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器信息数组，可以创建多个七层监听器。目前一个负载均衡下面最多允许创建50个七层监听器。
	ListenerSet []*CreateL7Listener `json:"ListenerSet,omitempty" name:"ListenerSet"`
}

func (r *CreateL7ListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7ListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL7ListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7ListenersResponseParams struct {
	// 新建的负载均衡七层监听器的唯一ID列表。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateL7ListenersResponse struct {
	*tchttp.BaseResponse
	Response *CreateL7ListenersResponseParams `json:"Response"`
}

func (r *CreateL7ListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7ListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateL7Rule struct {
	// 七层转发规则的转发域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 七层转发规则的转发路径。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 会话保持时间，单位：秒。可选值：30~3600。默认值0，表示不开启会话保持。
	SessionExpire *int64 `json:"SessionExpire,omitempty" name:"SessionExpire"`

	// 健康检查开关：1（开启）、0（关闭）。默认值0，表示关闭。
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 健康检查检查间隔时间，默认值：5，可选值：5-300，单位：秒。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 健康检查健康阈值，默认值：3，表示当连续探测三次健康则表示该转发正常，可选值：2-10，单位：次。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 健康检查不健康阈值，默认值：5，表示当连续探测五次不健康则表示该转发不正常，可选值：2-10，单位：次。
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 健康检查中认为健康的HTTP返回码的组合。可选值为1~5的集合，1表示HTTP返回码为1xx认为健康。2表示HTTP返回码为2xx认为健康。3表示HTTP返回码为3xx认为健康。4表示HTTP返回码为4xx认为健康。5表示HTTP返回码为5xx认为健康。
	HttpCodes []*uint64 `json:"HttpCodes,omitempty" name:"HttpCodes"`

	// 健康检查检查路径。
	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`

	// 健康检查检查域名。如果创建规则的域名使用通配符或正则表达式，则健康检查检查域名可自定义，否则必须跟健康检查检查域名一样。
	HttpCheckDomain *string `json:"HttpCheckDomain,omitempty" name:"HttpCheckDomain"`

	// 均衡方式：ip_hash、wrr。默认值wrr。
	BalanceMode *string `json:"BalanceMode,omitempty" name:"BalanceMode"`
}

// Predefined struct for user
type CreateL7RulesRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 七层转发规则信息数组，可以创建多个七层转发规则。目前一个七层监听器下面最多允许创建50个七层转发域名，而每一个转发域名下最多可以创建100个转发规则。目前只能单条创建，不能批量创建。
	RuleSet []*CreateL7Rule `json:"RuleSet,omitempty" name:"RuleSet"`
}

type CreateL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 七层转发规则信息数组，可以创建多个七层转发规则。目前一个七层监听器下面最多允许创建50个七层转发域名，而每一个转发域名下最多可以创建100个转发规则。目前只能单条创建，不能批量创建。
	RuleSet []*CreateL7Rule `json:"RuleSet,omitempty" name:"RuleSet"`
}

func (r *CreateL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "RuleSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7RulesResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateL7RulesResponseParams `json:"Response"`
}

func (r *CreateL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLoadBalancerBzConf struct {
	// 按月/按小时计费。
	BzPayMode *string `json:"BzPayMode,omitempty" name:"BzPayMode"`

	// 四层可选按带宽，连接数衡量。
	BzL4Metrics *string `json:"BzL4Metrics,omitempty" name:"BzL4Metrics"`

	// 七层可选按qps衡量。
	BzL7Metrics *string `json:"BzL7Metrics,omitempty" name:"BzL7Metrics"`
}

// Predefined struct for user
type CreateLoadBalancersRequestParams struct {
	// 黑石负载均衡实例所属的私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 负载均衡的类型，取值为open或internal。open表示公网(有日租)，internal表示内网。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 在私有网络内购买内网负载均衡实例的时候需要指定子网ID，内网负载均衡实例的VIP将从这个子网中产生。其他情况不用填写该字段。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 负载均衡所属项目ID。不填则属于默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 购买黑石负载均衡实例的数量。默认值为1, 最大值为20。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 黑石负载均衡的计费模式，取值为flow和bandwidth，其中flow模式表示流量模式，bandwidth表示带宽模式。默认值为flow。
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 负载均衡对应的TGW集群类别，取值为tunnel、fullnat或dnat。tunnel表示隧道集群，fullnat表示FULLNAT集群（普通外网负载均衡），dnat表示DNAT集群（增强型外网负载均衡）。默认值为fullnat。如需获取client IP，可以选择 tunnel 模式，fullnat 模式（tcp 通过toa 获取），dnat 模式。
	TgwSetType *string `json:"TgwSetType,omitempty" name:"TgwSetType"`

	// 负载均衡的独占类别，取值为0表示非独占，1表示四层独占，2表示七层独占，3表示四层和七层独占，4表示共享容灾。
	Exclusive *int64 `json:"Exclusive,omitempty" name:"Exclusive"`

	// 指定的VIP，如果指定，则数量必须与goodsNum一致。如果不指定，则由后台分配随机VIP。
	SpecifiedVips []*string `json:"SpecifiedVips,omitempty" name:"SpecifiedVips"`

	// （未全地域开放）保障型负载均衡设定参数，如果类别选择保障型则需传入此参数。
	BzConf *CreateLoadBalancerBzConf `json:"BzConf,omitempty" name:"BzConf"`

	// IP协议类型。可取的值为“ipv4”或“ipv6”。
	IpProtocolType *string `json:"IpProtocolType,omitempty" name:"IpProtocolType"`
}

type CreateLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 黑石负载均衡实例所属的私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 负载均衡的类型，取值为open或internal。open表示公网(有日租)，internal表示内网。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 在私有网络内购买内网负载均衡实例的时候需要指定子网ID，内网负载均衡实例的VIP将从这个子网中产生。其他情况不用填写该字段。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 负载均衡所属项目ID。不填则属于默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 购买黑石负载均衡实例的数量。默认值为1, 最大值为20。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 黑石负载均衡的计费模式，取值为flow和bandwidth，其中flow模式表示流量模式，bandwidth表示带宽模式。默认值为flow。
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 负载均衡对应的TGW集群类别，取值为tunnel、fullnat或dnat。tunnel表示隧道集群，fullnat表示FULLNAT集群（普通外网负载均衡），dnat表示DNAT集群（增强型外网负载均衡）。默认值为fullnat。如需获取client IP，可以选择 tunnel 模式，fullnat 模式（tcp 通过toa 获取），dnat 模式。
	TgwSetType *string `json:"TgwSetType,omitempty" name:"TgwSetType"`

	// 负载均衡的独占类别，取值为0表示非独占，1表示四层独占，2表示七层独占，3表示四层和七层独占，4表示共享容灾。
	Exclusive *int64 `json:"Exclusive,omitempty" name:"Exclusive"`

	// 指定的VIP，如果指定，则数量必须与goodsNum一致。如果不指定，则由后台分配随机VIP。
	SpecifiedVips []*string `json:"SpecifiedVips,omitempty" name:"SpecifiedVips"`

	// （未全地域开放）保障型负载均衡设定参数，如果类别选择保障型则需传入此参数。
	BzConf *CreateLoadBalancerBzConf `json:"BzConf,omitempty" name:"BzConf"`

	// IP协议类型。可取的值为“ipv4”或“ipv6”。
	IpProtocolType *string `json:"IpProtocolType,omitempty" name:"IpProtocolType"`
}

func (r *CreateLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "LoadBalancerType")
	delete(f, "SubnetId")
	delete(f, "ProjectId")
	delete(f, "GoodsNum")
	delete(f, "PayMode")
	delete(f, "TgwSetType")
	delete(f, "Exclusive")
	delete(f, "SpecifiedVips")
	delete(f, "BzConf")
	delete(f, "IpProtocolType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancersResponseParams struct {
	// 创建的黑石负载均衡实例ID。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 创建负载均衡的异步任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoadBalancersResponseParams `json:"Response"`
}

func (r *CreateLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrafficMirrorRequestParams struct {
	// 流量镜像实例别名。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 流量镜像实例所属的私有网络ID，形如：vpc-xxx。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type CreateTrafficMirrorRequest struct {
	*tchttp.BaseRequest
	
	// 流量镜像实例别名。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 流量镜像实例所属的私有网络ID，形如：vpc-xxx。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *CreateTrafficMirrorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrafficMirrorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Alias")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTrafficMirrorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrafficMirrorResponseParams struct {
	// 流量镜像实例ID
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTrafficMirrorResponse struct {
	*tchttp.BaseResponse
	Response *CreateTrafficMirrorResponseParams `json:"Response"`
}

func (r *CreateTrafficMirrorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrafficMirrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL7DomainsRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID列表，可通过接口DescribeL7Rules查询。
	DomainIds []*string `json:"DomainIds,omitempty" name:"DomainIds"`
}

type DeleteL7DomainsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID列表，可通过接口DescribeL7Rules查询。
	DomainIds []*string `json:"DomainIds,omitempty" name:"DomainIds"`
}

func (r *DeleteL7DomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL7DomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "DomainIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteL7DomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL7DomainsResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteL7DomainsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteL7DomainsResponseParams `json:"Response"`
}

func (r *DeleteL7DomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL7DomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL7RulesRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID列表，可通过接口DescribeL7Rules查询。
	LocationIds []*string `json:"LocationIds,omitempty" name:"LocationIds"`
}

type DeleteL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID列表，可通过接口DescribeL7Rules查询。
	LocationIds []*string `json:"LocationIds,omitempty" name:"LocationIds"`
}

func (r *DeleteL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "DomainId")
	delete(f, "LocationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL7RulesResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteL7RulesResponseParams `json:"Response"`
}

func (r *DeleteL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenersRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 待删除的负载均衡四层和七层监听器ID列表，可通过接口DescribeL4Listeners和DescribeL7Listeners查询。目前同时只能删除一种类型的监听器，并且删除七层监听器的数量上限为一个。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

type DeleteListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 待删除的负载均衡四层和七层监听器ID列表，可通过接口DescribeL4Listeners和DescribeL7Listeners查询。目前同时只能删除一种类型的监听器，并且删除七层监听器的数量上限为一个。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

func (r *DeleteListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenersResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteListenersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteListenersResponseParams `json:"Response"`
}

func (r *DeleteListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
}

type DeleteLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
}

func (r *DeleteLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancerResponseParams `json:"Response"`
}

func (r *DeleteLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrafficMirrorRequestParams struct {
	// 流量镜像实例ID数组，可以批量删除，每次删除上限为20
	TrafficMirrorIds []*string `json:"TrafficMirrorIds,omitempty" name:"TrafficMirrorIds"`
}

type DeleteTrafficMirrorRequest struct {
	*tchttp.BaseRequest
	
	// 流量镜像实例ID数组，可以批量删除，每次删除上限为20
	TrafficMirrorIds []*string `json:"TrafficMirrorIds,omitempty" name:"TrafficMirrorIds"`
}

func (r *DeleteTrafficMirrorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrafficMirrorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficMirrorIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTrafficMirrorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrafficMirrorResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTrafficMirrorResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTrafficMirrorResponseParams `json:"Response"`
}

func (r *DeleteTrafficMirrorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrafficMirrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertDetailRequestParams struct {
	// 证书ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`
}

type DescribeCertDetailRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`
}

func (r *DescribeCertDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertDetailResponseParams struct {
	// 证书ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 证书名称。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 证书类型（SVR=服务器证书，CA=客户端证书）。
	CertType *string `json:"CertType,omitempty" name:"CertType"`

	// 证书内容。
	CertContent *string `json:"CertContent,omitempty" name:"CertContent"`

	// 证书主域名。
	CertDomain *string `json:"CertDomain,omitempty" name:"CertDomain"`

	// 证书子域名列表。
	CertSubjectDomain []*string `json:"CertSubjectDomain,omitempty" name:"CertSubjectDomain"`

	// 证书上传时间。
	CertUploadTime *string `json:"CertUploadTime,omitempty" name:"CertUploadTime"`

	// 证书生效时间。
	CertBeginTime *string `json:"CertBeginTime,omitempty" name:"CertBeginTime"`

	// 证书失效时间。
	CertEndTime *string `json:"CertEndTime,omitempty" name:"CertEndTime"`

	// 该证书关联的黑石负载均衡对象列表。
	CertLoadBalancerSet []*CertDetailLoadBalancer `json:"CertLoadBalancerSet,omitempty" name:"CertLoadBalancerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCertDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertDetailResponseParams `json:"Response"`
}

func (r *DescribeCertDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesBindInfoRequestParams struct {
	// 黑石私有网络唯一ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 主机ID或虚机IP列表，可用于获取绑定了该主机的负载均衡列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeDevicesBindInfoRequest struct {
	*tchttp.BaseRequest
	
	// 黑石私有网络唯一ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 主机ID或虚机IP列表，可用于获取绑定了该主机的负载均衡列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeDevicesBindInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesBindInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicesBindInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesBindInfoResponseParams struct {
	// 返回的负载均衡绑定信息。
	LoadBalancerSet []*DevicesBindInfoLoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDevicesBindInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicesBindInfoResponseParams `json:"Response"`
}

func (r *DescribeDevicesBindInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesBindInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeL4Backend struct {
	// 待绑定的主机端口，可选值1~65535。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 黑石物理机的主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type DescribeL4BackendsRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡四层监听器ID，可通过接口DescribeL4Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待查询的主机信息。
	BackendSet []*DescribeL4Backend `json:"BackendSet,omitempty" name:"BackendSet"`
}

type DescribeL4BackendsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡四层监听器ID，可通过接口DescribeL4Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待查询的主机信息。
	BackendSet []*DescribeL4Backend `json:"BackendSet,omitempty" name:"BackendSet"`
}

func (r *DescribeL4BackendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4BackendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "BackendSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL4BackendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4BackendsResponseParams struct {
	// 返回的绑定关系列表。
	BackendSet []*L4Backend `json:"BackendSet,omitempty" name:"BackendSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeL4BackendsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL4BackendsResponseParams `json:"Response"`
}

func (r *DescribeL4BackendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4BackendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4ListenerInfoRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 查找的键值，可用于模糊查找该名称的监听器。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 主机ID或虚机IP列表，可用于获取绑定了该主机的监听器。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeL4ListenerInfoRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 查找的键值，可用于模糊查找该名称的监听器。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 主机ID或虚机IP列表，可用于获取绑定了该主机的监听器。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeL4ListenerInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4ListenerInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SearchKey")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL4ListenerInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4ListenerInfoResponseParams struct {
	// 返回的四层监听器列表。
	ListenerSet []*L4ListenerInfo `json:"ListenerSet,omitempty" name:"ListenerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeL4ListenerInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL4ListenerInfoResponseParams `json:"Response"`
}

func (r *DescribeL4ListenerInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4ListenerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4ListenersRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 四层监听器实例ID数组，可通过接口DescribeL4Listeners查询。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

type DescribeL4ListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 四层监听器实例ID数组，可通过接口DescribeL4Listeners查询。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

func (r *DescribeL4ListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4ListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL4ListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL4ListenersResponseParams struct {
	// 监听器信息数组。
	ListenerSet []*L4Listener `json:"ListenerSet,omitempty" name:"ListenerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeL4ListenersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL4ListenersResponseParams `json:"Response"`
}

func (r *DescribeL4ListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL4ListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7BackendsRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID，可通过接口DescribeL7Rules查询。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 查询条件，传'all'则查询所有与规则绑定的主机信息。如果为all时，DomainId和LocationId参数没有意义不必传入，否则DomainId和LocationId参数必须传入。
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
}

type DescribeL7BackendsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID，可通过接口DescribeL7Rules查询。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 查询条件，传'all'则查询所有与规则绑定的主机信息。如果为all时，DomainId和LocationId参数没有意义不必传入，否则DomainId和LocationId参数必须传入。
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`
}

func (r *DescribeL7BackendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7BackendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "DomainId")
	delete(f, "LocationId")
	delete(f, "QueryType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL7BackendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7BackendsResponseParams struct {
	// 返回的绑定关系列表。
	BackendSet []*L7Backend `json:"BackendSet,omitempty" name:"BackendSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeL7BackendsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL7BackendsResponseParams `json:"Response"`
}

func (r *DescribeL7BackendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7BackendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7ListenerInfoRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 查找的键值，可用于模糊查找有该转发域名的监听器。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 主机ID或虚机IP列表，可用于获取绑定了该主机的监听器。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 是否获取转发规则下的主机信息。默认为0，不获取。
	IfGetBackendInfo *int64 `json:"IfGetBackendInfo,omitempty" name:"IfGetBackendInfo"`
}

type DescribeL7ListenerInfoRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 查找的键值，可用于模糊查找有该转发域名的监听器。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 主机ID或虚机IP列表，可用于获取绑定了该主机的监听器。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 是否获取转发规则下的主机信息。默认为0，不获取。
	IfGetBackendInfo *int64 `json:"IfGetBackendInfo,omitempty" name:"IfGetBackendInfo"`
}

func (r *DescribeL7ListenerInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7ListenerInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SearchKey")
	delete(f, "InstanceIds")
	delete(f, "IfGetBackendInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL7ListenerInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7ListenerInfoResponseParams struct {
	// 返回的七层监听器列表。
	ListenerSet []*L7ListenerInfo `json:"ListenerSet,omitempty" name:"ListenerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeL7ListenerInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL7ListenerInfoResponseParams `json:"Response"`
}

func (r *DescribeL7ListenerInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7ListenerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7ListenersExRequestParams struct {
	// 返回的监听器中标识是否绑定在此流量镜像中。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 待获取监听器所在的VPC的ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 此VPC中获取负载均衡的偏移。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 此VPC中获取负载均衡的数量。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// LoadBalancerId - String - （过滤条件）负载均衡ID。
	// LoadBalancerName - String - （过滤条件）负载均衡名称。
	// Vip - String - （过滤条件）VIP。
	// ListenerId - String - （过滤条件）监听器ID。
	// ListenerName -  String - （过滤条件）监听器名称。
	// Protocol -  String - （过滤条件）七层协议。
	// LoadBalancerPort -  String - （过滤条件）监听器端口。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeL7ListenersExRequest struct {
	*tchttp.BaseRequest
	
	// 返回的监听器中标识是否绑定在此流量镜像中。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 待获取监听器所在的VPC的ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 此VPC中获取负载均衡的偏移。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 此VPC中获取负载均衡的数量。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件。
	// LoadBalancerId - String - （过滤条件）负载均衡ID。
	// LoadBalancerName - String - （过滤条件）负载均衡名称。
	// Vip - String - （过滤条件）VIP。
	// ListenerId - String - （过滤条件）监听器ID。
	// ListenerName -  String - （过滤条件）监听器名称。
	// Protocol -  String - （过滤条件）七层协议。
	// LoadBalancerPort -  String - （过滤条件）监听器端口。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeL7ListenersExRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7ListenersExRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficMirrorId")
	delete(f, "VpcId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL7ListenersExRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7ListenersExResponseParams struct {
	// 此指定VPC中负载均衡的总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 符合条件的监听器。
	ListenerSet []*L7ExListener `json:"ListenerSet,omitempty" name:"ListenerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeL7ListenersExResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL7ListenersExResponseParams `json:"Response"`
}

func (r *DescribeL7ListenersExResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7ListenersExResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7ListenersRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID列表，可通过接口DescribeL7Listeners查询。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

type DescribeL7ListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID列表，可通过接口DescribeL7Listeners查询。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

func (r *DescribeL7ListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7ListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL7ListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7ListenersResponseParams struct {
	// 返回的七层监听器列表。
	ListenerSet []*L7Listener `json:"ListenerSet,omitempty" name:"ListenerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeL7ListenersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL7ListenersResponseParams `json:"Response"`
}

func (r *DescribeL7ListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7ListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7RulesRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名ID列表，可通过接口DescribeL7Rules查询。
	DomainIds []*string `json:"DomainIds,omitempty" name:"DomainIds"`
}

type DescribeL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名ID列表，可通过接口DescribeL7Rules查询。
	DomainIds []*string `json:"DomainIds,omitempty" name:"DomainIds"`
}

func (r *DescribeL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "DomainIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7RulesResponseParams struct {
	// 返回的转发规则列表。
	RuleSet []*L7Rule `json:"RuleSet,omitempty" name:"RuleSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL7RulesResponseParams `json:"Response"`
}

func (r *DescribeL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerPortInfoRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
}

type DescribeLoadBalancerPortInfoRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
}

func (r *DescribeLoadBalancerPortInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerPortInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancerPortInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerPortInfoResponseParams struct {
	// 返回的监听器列表（四层和七层）。
	ListenerSet []*LoadBalancerPortInfoListener `json:"ListenerSet,omitempty" name:"ListenerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoadBalancerPortInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancerPortInfoResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancerPortInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerPortInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerTaskResultRequestParams struct {
	// 任务ID。由具体的异步操作接口提供。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeLoadBalancerTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。由具体的异步操作接口提供。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeLoadBalancerTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancerTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerTaskResultResponseParams struct {
	// 任务当前状态。0：成功，1：失败，2：进行中。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoadBalancerTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancerTaskResultResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancerTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersRequestParams struct {
	// 负载均衡器ID数组
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 负载均衡的类型 : open表示公网LB类型，internal表示内网LB类型
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 负载均衡器名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡域名。规则：1-60个小写英文字母、数字、点号“.”或连接线“-”。内网类型的负载均衡不能配置该字段
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 负载均衡获得的公网IP地址,支持多个
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// 数据偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数据长度，默认为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模糊查找名称、域名、VIP
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 排序字段，支持：loadBalancerName,createTime,domain,loadBalancerType
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 1倒序，0顺序，默认顺序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 是否筛选独占集群，0表示非独占集群，1表示四层独占集群，2表示七层独占集群，3表示四层和七层独占集群，4表示共享容灾
	Exclusive *uint64 `json:"Exclusive,omitempty" name:"Exclusive"`

	// 该负载均衡对应的tgw集群（fullnat,tunnel,dnat）
	TgwSetType *string `json:"TgwSetType,omitempty" name:"TgwSetType"`

	// 该负载均衡对应的所在的私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 'CONFLIST' 查询带confId的LB列表，'CONFID' 查询某个confId绑定的LB列表
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// 个性化配置ID
	ConfId *string `json:"ConfId,omitempty" name:"ConfId"`
}

type DescribeLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡器ID数组
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 负载均衡的类型 : open表示公网LB类型，internal表示内网LB类型
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 负载均衡器名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡域名。规则：1-60个小写英文字母、数字、点号“.”或连接线“-”。内网类型的负载均衡不能配置该字段
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 负载均衡获得的公网IP地址,支持多个
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// 数据偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数据长度，默认为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 模糊查找名称、域名、VIP
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 排序字段，支持：loadBalancerName,createTime,domain,loadBalancerType
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 1倒序，0顺序，默认顺序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 是否筛选独占集群，0表示非独占集群，1表示四层独占集群，2表示七层独占集群，3表示四层和七层独占集群，4表示共享容灾
	Exclusive *uint64 `json:"Exclusive,omitempty" name:"Exclusive"`

	// 该负载均衡对应的tgw集群（fullnat,tunnel,dnat）
	TgwSetType *string `json:"TgwSetType,omitempty" name:"TgwSetType"`

	// 该负载均衡对应的所在的私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 'CONFLIST' 查询带confId的LB列表，'CONFID' 查询某个confId绑定的LB列表
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// 个性化配置ID
	ConfId *string `json:"ConfId,omitempty" name:"ConfId"`
}

func (r *DescribeLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "LoadBalancerType")
	delete(f, "LoadBalancerName")
	delete(f, "Domain")
	delete(f, "LoadBalancerVips")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "ProjectId")
	delete(f, "Exclusive")
	delete(f, "TgwSetType")
	delete(f, "VpcId")
	delete(f, "QueryType")
	delete(f, "ConfId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersResponseParams struct {
	// 返回负载均衡信息列表。
	LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`

	// 符合条件的负载均衡总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancersResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficMirrorListenersRequestParams struct {
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 分页的偏移量，也即从第几条记录开始查询
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次查询返回的条目数，默认值：500。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 待搜索的负载均衡Id。
	SearchLoadBalancerIds []*string `json:"SearchLoadBalancerIds,omitempty" name:"SearchLoadBalancerIds"`

	// 待搜索的负载均衡名称。
	SearchLoadBalancerNames []*string `json:"SearchLoadBalancerNames,omitempty" name:"SearchLoadBalancerNames"`

	// 待搜索的Vip。
	SearchVips []*string `json:"SearchVips,omitempty" name:"SearchVips"`

	// 待搜索的监听器ID。
	SearchListenerIds []*string `json:"SearchListenerIds,omitempty" name:"SearchListenerIds"`

	// 待搜索的监听器名称。
	SearchListenerNames []*string `json:"SearchListenerNames,omitempty" name:"SearchListenerNames"`

	// 待搜索的协议名称。
	SearchProtocols []*string `json:"SearchProtocols,omitempty" name:"SearchProtocols"`

	// 待搜索的端口。
	SearchLoadBalancerPorts []*uint64 `json:"SearchLoadBalancerPorts,omitempty" name:"SearchLoadBalancerPorts"`
}

type DescribeTrafficMirrorListenersRequest struct {
	*tchttp.BaseRequest
	
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 分页的偏移量，也即从第几条记录开始查询
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次查询返回的条目数，默认值：500。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 待搜索的负载均衡Id。
	SearchLoadBalancerIds []*string `json:"SearchLoadBalancerIds,omitempty" name:"SearchLoadBalancerIds"`

	// 待搜索的负载均衡名称。
	SearchLoadBalancerNames []*string `json:"SearchLoadBalancerNames,omitempty" name:"SearchLoadBalancerNames"`

	// 待搜索的Vip。
	SearchVips []*string `json:"SearchVips,omitempty" name:"SearchVips"`

	// 待搜索的监听器ID。
	SearchListenerIds []*string `json:"SearchListenerIds,omitempty" name:"SearchListenerIds"`

	// 待搜索的监听器名称。
	SearchListenerNames []*string `json:"SearchListenerNames,omitempty" name:"SearchListenerNames"`

	// 待搜索的协议名称。
	SearchProtocols []*string `json:"SearchProtocols,omitempty" name:"SearchProtocols"`

	// 待搜索的端口。
	SearchLoadBalancerPorts []*uint64 `json:"SearchLoadBalancerPorts,omitempty" name:"SearchLoadBalancerPorts"`
}

func (r *DescribeTrafficMirrorListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficMirrorListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficMirrorId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchLoadBalancerIds")
	delete(f, "SearchLoadBalancerNames")
	delete(f, "SearchVips")
	delete(f, "SearchListenerIds")
	delete(f, "SearchListenerNames")
	delete(f, "SearchProtocols")
	delete(f, "SearchLoadBalancerPorts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrafficMirrorListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficMirrorListenersResponseParams struct {
	// 监听器列表。
	ListenerSet []*TrafficMirrorListener `json:"ListenerSet,omitempty" name:"ListenerSet"`

	// 监听器总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrafficMirrorListenersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrafficMirrorListenersResponseParams `json:"Response"`
}

func (r *DescribeTrafficMirrorListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficMirrorListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTrafficMirrorReceiver struct {
	// 物理机实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 物理机绑定的端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

// Predefined struct for user
type DescribeTrafficMirrorReceiverHealthStatusRequestParams struct {
	// 查询所在的流量镜像ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 流量镜像接收机实例ID和端口数组。
	ReceiverSet []*DescribeTrafficMirrorReceiver `json:"ReceiverSet,omitempty" name:"ReceiverSet"`
}

type DescribeTrafficMirrorReceiverHealthStatusRequest struct {
	*tchttp.BaseRequest
	
	// 查询所在的流量镜像ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 流量镜像接收机实例ID和端口数组。
	ReceiverSet []*DescribeTrafficMirrorReceiver `json:"ReceiverSet,omitempty" name:"ReceiverSet"`
}

func (r *DescribeTrafficMirrorReceiverHealthStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficMirrorReceiverHealthStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficMirrorId")
	delete(f, "ReceiverSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrafficMirrorReceiverHealthStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficMirrorReceiverHealthStatusResponseParams struct {
	// 内网IP和端口对应的状态。
	ReceiversStatusSet []*TrafficMirrorReciversStatus `json:"ReceiversStatusSet,omitempty" name:"ReceiversStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrafficMirrorReceiverHealthStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrafficMirrorReceiverHealthStatusResponseParams `json:"Response"`
}

func (r *DescribeTrafficMirrorReceiverHealthStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficMirrorReceiverHealthStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficMirrorReceiversRequestParams struct {
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 接收机黑石物理机实例ID数组。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 接收机接收端口数组。
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// 接收机实例权重数组。
	Weights []*int64 `json:"Weights,omitempty" name:"Weights"`

	// 分页的偏移量，也即从第几条记录开始查询
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次查询返回的条目数，默认值：500。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索instance或者alias
	VagueStr *string `json:"VagueStr,omitempty" name:"VagueStr"`

	// 搜索IP
	VagueIp *string `json:"VagueIp,omitempty" name:"VagueIp"`
}

type DescribeTrafficMirrorReceiversRequest struct {
	*tchttp.BaseRequest
	
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 接收机黑石物理机实例ID数组。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 接收机接收端口数组。
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// 接收机实例权重数组。
	Weights []*int64 `json:"Weights,omitempty" name:"Weights"`

	// 分页的偏移量，也即从第几条记录开始查询
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次查询返回的条目数，默认值：500。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索instance或者alias
	VagueStr *string `json:"VagueStr,omitempty" name:"VagueStr"`

	// 搜索IP
	VagueIp *string `json:"VagueIp,omitempty" name:"VagueIp"`
}

func (r *DescribeTrafficMirrorReceiversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficMirrorReceiversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficMirrorId")
	delete(f, "InstanceIds")
	delete(f, "Ports")
	delete(f, "Weights")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "VagueStr")
	delete(f, "VagueIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrafficMirrorReceiversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficMirrorReceiversResponseParams struct {
	// 接收机列表，具体结构描述如data结构所示。
	ReceiverSet []*TrafficMirrorReceiver `json:"ReceiverSet,omitempty" name:"ReceiverSet"`

	// 接收机总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrafficMirrorReceiversResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrafficMirrorReceiversResponseParams `json:"Response"`
}

func (r *DescribeTrafficMirrorReceiversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficMirrorReceiversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficMirrorsRequestParams struct {
	// 流量镜像实例ID的数组，支持批量查询
	TrafficMirrorIds []*string `json:"TrafficMirrorIds,omitempty" name:"TrafficMirrorIds"`

	// 流量镜像实例别名数组。
	Aliases []*string `json:"Aliases,omitempty" name:"Aliases"`

	// 流量镜像实例所属的私有网络ID数组，形如：vpc-xxx。
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// 分页的偏移量，也即从第几条记录开始查询
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次查询返回的条目数，默认值：500。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段。trafficMirrorId或者createTime。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，取值：0:增序(默认)，1:降序
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 模糊匹配trafficMirrorId或者alias字段。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

type DescribeTrafficMirrorsRequest struct {
	*tchttp.BaseRequest
	
	// 流量镜像实例ID的数组，支持批量查询
	TrafficMirrorIds []*string `json:"TrafficMirrorIds,omitempty" name:"TrafficMirrorIds"`

	// 流量镜像实例别名数组。
	Aliases []*string `json:"Aliases,omitempty" name:"Aliases"`

	// 流量镜像实例所属的私有网络ID数组，形如：vpc-xxx。
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// 分页的偏移量，也即从第几条记录开始查询
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 单次查询返回的条目数，默认值：500。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段。trafficMirrorId或者createTime。
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，取值：0:增序(默认)，1:降序
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 模糊匹配trafficMirrorId或者alias字段。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeTrafficMirrorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficMirrorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficMirrorIds")
	delete(f, "Aliases")
	delete(f, "VpcIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrafficMirrorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficMirrorsResponseParams struct {
	// 流量镜像总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 对象数组。数组元素为流量镜像信息，具体结构描述如list结构所示。
	TrafficMirrorSet []*TrafficMirror `json:"TrafficMirrorSet,omitempty" name:"TrafficMirrorSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrafficMirrorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrafficMirrorsResponseParams `json:"Response"`
}

func (r *DescribeTrafficMirrorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficMirrorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DevicesBindInfoBackend struct {
	// 黑石物理机的主机ID、托管主机ID或虚拟机IP。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type DevicesBindInfoL4Listener struct {
	// 七层监听器实例ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 七层监听器协议类型，可选值：http,https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 七层监听器的监听端口。
	LoadBalancerPort *int64 `json:"LoadBalancerPort,omitempty" name:"LoadBalancerPort"`

	// 该转发路径所绑定的主机列表。
	BackendSet []*DevicesBindInfoBackend `json:"BackendSet,omitempty" name:"BackendSet"`
}

type DevicesBindInfoL7Listener struct {
	// 七层监听器实例ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 七层监听器协议类型，可选值：http,https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 七层监听器的监听端口。
	LoadBalancerPort *int64 `json:"LoadBalancerPort,omitempty" name:"LoadBalancerPort"`

	// 返回的转发规则列表。
	RuleSet []*DevicesBindInfoRule `json:"RuleSet,omitempty" name:"RuleSet"`
}

type DevicesBindInfoLoadBalancer struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 开发商AppId。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 负载均衡所属的项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 黑石私有网络唯一ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 负载均衡的IP地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 负载均衡对应的TGW集群类别，取值为tunnel或fullnat。tunnel表示隧道集群，fullnat表示FULLNAT集群。
	TgwSetType *string `json:"TgwSetType,omitempty" name:"TgwSetType"`

	// 是否独占TGW集群。
	Exclusive *int64 `json:"Exclusive,omitempty" name:"Exclusive"`

	// 具有该绑定关系的四层监听器列表。
	L4ListenerSet []*DevicesBindInfoL4Listener `json:"L4ListenerSet,omitempty" name:"L4ListenerSet"`

	// 具有该绑定关系的七层监听器列表。
	L7ListenerSet []*DevicesBindInfoL7Listener `json:"L7ListenerSet,omitempty" name:"L7ListenerSet"`
}

type DevicesBindInfoLocation struct {
	// 转发路径。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 转发路径实例ID。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 该转发路径所绑定的主机列表。
	BackendSet []*DevicesBindInfoBackend `json:"BackendSet,omitempty" name:"BackendSet"`
}

type DevicesBindInfoRule struct {
	// 转发域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发域名ID。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径列表。
	LocationSet []*DevicesBindInfoLocation `json:"LocationSet,omitempty" name:"LocationSet"`
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type L4Backend struct {
	// 绑定类别（0代表黑石物理机，1代表虚拟机IP）。
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`

	// 主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 权重。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 当前绑定关系的健康检查状态（Dead代表不健康，Alive代表健康）。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 黑石物理机的主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 黑石物理机的别名。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 主机IP。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 黑石物理机当前可以执行的操作。
	Operates []*string `json:"Operates,omitempty" name:"Operates"`

	// 主机探测端口。
	ProbePort *int64 `json:"ProbePort,omitempty" name:"ProbePort"`
}

type L4Listener struct {
	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 用户自定义的监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 负载均衡实例监听器协议类型，可选值tcp，udp。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 负载均衡监听器的监听接口，可选值1~65535。
	LoadBalancerPort *int64 `json:"LoadBalancerPort,omitempty" name:"LoadBalancerPort"`

	// 用于计费模式为固定带宽计费，指定监听器最大带宽值，可选值：0-1000，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 监听器的类别：L4Listener（四层监听器），L7Listener（七层监听器）。
	ListenerType *string `json:"ListenerType,omitempty" name:"ListenerType"`

	// 会话保持时间。单位：秒
	SessionExpire *int64 `json:"SessionExpire,omitempty" name:"SessionExpire"`

	// 是否开启了检查：1（开启）、0（关闭）。
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 响应超时时间，单位：秒。
	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 检查间隔，单位：秒。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 负载均衡监听器健康阈值，默认值：3，表示当连续探测三次健康则表示该转发正常，可选值：2-10，单位：次。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 负载均衡监听器不健康阈值，默认值：3，表示当连续探测三次不健康则表示该转发不正常，可选值：2-10，单位：次。
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 是否开启自定义健康检查：1（开启）、0（关闭）。默认值0，表示关闭。（该字段在健康检查开启的情况下才生效）
	CustomHealthSwitch *int64 `json:"CustomHealthSwitch,omitempty" name:"CustomHealthSwitch"`

	// 自定义健康探测内容类型，可选值：text（文本）、hexadecimal（十六进制）。
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 探测内容类型为文本方式时，针对请求文本中换行替换方式。可选值：1（替换为LF）、2（替换为CR）、3（替换为LF+CR）。
	LineSeparatorType *int64 `json:"LineSeparatorType,omitempty" name:"LineSeparatorType"`

	// 自定义探测请求内容。
	HealthRequest *string `json:"HealthRequest,omitempty" name:"HealthRequest"`

	// 自定义探测返回内容。
	HealthResponse *string `json:"HealthResponse,omitempty" name:"HealthResponse"`

	// 是否开启toa：1（开启）、0（关闭）。
	ToaFlag *int64 `json:"ToaFlag,omitempty" name:"ToaFlag"`

	// 监听器当前状态（0代表创建中，1代表正常运行，2代表创建失败，3代表删除中，4代表删除失败）。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间戳。
	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`

	// 转发后端服务器调度类型。
	BalanceMode *string `json:"BalanceMode,omitempty" name:"BalanceMode"`
}

type L4ListenerInfo struct {
	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 用户自定义的监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 负载均衡实例监听器协议类型，可选值tcp，udp。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 负载均衡监听器的监听接口，可选值1~65535。
	LoadBalancerPort *int64 `json:"LoadBalancerPort,omitempty" name:"LoadBalancerPort"`

	// 用于计费模式为固定带宽计费，指定监听器最大带宽值，可选值：0-1000，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 监听器的类别：L4Listener（四层监听器），L7Listener（七层监听器）。
	ListenerType *string `json:"ListenerType,omitempty" name:"ListenerType"`

	// 会话保持时间。单位：秒
	SessionExpire *int64 `json:"SessionExpire,omitempty" name:"SessionExpire"`

	// 是否开启了检查：1（开启）、0（关闭）。
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 响应超时时间，单位：秒。
	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 检查间隔，单位：秒。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 负载均衡监听器健康阈值，默认值：3，表示当连续探测三次健康则表示该转发正常，可选值：2-10，单位：次。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 负载均衡监听器不健康阈值，默认值：3，表示当连续探测三次不健康则表示该转发不正常，可选值：2-10，单位：次。
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 监听器当前状态（0代表创建中，1代表正常运行，2代表创建失败，3代表删除中，4代表删除失败）。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间戳。
	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`

	// 是否开启自定义健康检查：1（开启）、0（关闭）。默认值0，表示关闭。（该字段在健康检查开启的情况下才生效）
	CustomHealthSwitch *int64 `json:"CustomHealthSwitch,omitempty" name:"CustomHealthSwitch"`

	// 自定义健康探测内容类型，可选值：text（文本）、hexadecimal（十六进制）。
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 探测内容类型为文本方式时，针对请求文本中换行替换方式。可选值：1（替换为LF）、2（替换为CR）、3（替换为LF+CR）。
	LineSeparatorType *int64 `json:"LineSeparatorType,omitempty" name:"LineSeparatorType"`

	// 自定义探测请求内容。
	HealthRequest *string `json:"HealthRequest,omitempty" name:"HealthRequest"`

	// 自定义探测返回内容。
	HealthResponse *string `json:"HealthResponse,omitempty" name:"HealthResponse"`

	// 是否开启toa：1（开启）、0（关闭）。
	ToaFlag *int64 `json:"ToaFlag,omitempty" name:"ToaFlag"`

	// 转发后端服务器调度类型。
	BalanceMode *string `json:"BalanceMode,omitempty" name:"BalanceMode"`
}

type L7Backend struct {
	// 绑定类别（0代表黑石物理机，1代表虚拟机IP）。
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`

	// 主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 权重。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 当前绑定关系的健康检查状态（Dead代表不健康，Alive代表健康）。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 黑石物理机的主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 黑石物理机的别名。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 主机IP。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 黑石物理机的管理IP。
	MgtIp *string `json:"MgtIp,omitempty" name:"MgtIp"`

	// 黑石物理机当前可以执行的操作。
	Operates []*string `json:"Operates,omitempty" name:"Operates"`
}

type L7ExListener struct {
	// 绑定的监听器唯一ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 七层监听器协议类型，可选值：http,https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器的监听端口。
	LoadBalancerPort *uint64 `json:"LoadBalancerPort,omitempty" name:"LoadBalancerPort"`

	// 当前带宽。
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 带宽上限。
	MaxBandwidth *uint64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 监听器类型。
	ListenerType *string `json:"ListenerType,omitempty" name:"ListenerType"`

	// 认证方式：0（不认证，用于http），1（单向认证，用于https），2（双向认证，用于https）。
	SslMode *int64 `json:"SslMode,omitempty" name:"SslMode"`

	// 服务端证书ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 客户端证书ID。
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// 添加时间。
	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`

	// 负载均衡名ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 私有网络名称。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 私有网络Cidr。
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 负载均衡的VIP。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// 负载均衡名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡IPV6的VIP。
	LoadBalancerVipv6s []*string `json:"LoadBalancerVipv6s,omitempty" name:"LoadBalancerVipv6s"`

	// 支持的IP协议类型。ipv4或者是ipv6。
	IpProtocolType *string `json:"IpProtocolType,omitempty" name:"IpProtocolType"`

	// 是否绑定在入参指定的流量镜像中。
	BindTrafficMirror *bool `json:"BindTrafficMirror,omitempty" name:"BindTrafficMirror"`
}

type L7Listener struct {
	// 七层监听器实例ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 七层监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 七层监听器协议类型，可选值：http,https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 七层监听器的监听端口。
	LoadBalancerPort *int64 `json:"LoadBalancerPort,omitempty" name:"LoadBalancerPort"`

	// 计费模式为按固定带宽方式时监听器的限速值，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 监听器的类别：L4Listener（四层监听器），L7Listener（七层监听器）。
	ListenerType *string `json:"ListenerType,omitempty" name:"ListenerType"`

	// 七层监听器的认证方式：0（不认证，用于http），1（单向认证，用于https），2（双向认证，用于https）。
	SslMode *int64 `json:"SslMode,omitempty" name:"SslMode"`

	// 七层监听器关联的服务端证书ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 七层监听器关联的客户端证书ID。
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// 监听器当前状态（0代表创建中，1代表正常运行，2代表创建失败，3代表删除中，4代表删除失败）。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间戳。
	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`

	// https转发类型。0：https。1：spdy。2：http2。3：spdy+http2。
	ForwardProtocol *int64 `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`
}

type L7ListenerInfo struct {
	// 七层监听器实例ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 七层监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 七层监听器协议类型，可选值：http,https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 七层监听器的监听端口。
	LoadBalancerPort *int64 `json:"LoadBalancerPort,omitempty" name:"LoadBalancerPort"`

	// 计费模式为按固定带宽方式时监听器的限速值，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 监听器的类别：L4Listener（四层监听器），L7Listener（七层监听器）。
	ListenerType *string `json:"ListenerType,omitempty" name:"ListenerType"`

	// 七层监听器的认证方式：0（不认证，用于http），1（单向认证，用于https），2（双向认证，用于https）。
	SslMode *int64 `json:"SslMode,omitempty" name:"SslMode"`

	// 七层监听器关联的服务端证书ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 七层监听器关联的客户端证书ID。
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// 当前绑定关系的健康检查状态（Dead代表不健康，Alive代表健康）。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间戳。
	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`

	// 返回的转发规则列表。
	RuleSet []*L7ListenerInfoRule `json:"RuleSet,omitempty" name:"RuleSet"`

	// https转发类型。0：https。1：spdy。2：http2。3：spdy+http2。
	ForwardProtocol *int64 `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`
}

type L7ListenerInfoBackend struct {
	// 绑定类别（0代表黑石物理机，1代表虚拟机IP）。
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`

	// 主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 权重。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 当前绑定关系的健康检查状态（Dead代表不健康，Alive代表健康）。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 黑石物理机的主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 黑石物理机的别名。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 主机IP。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
}

type L7ListenerInfoLocation struct {
	// 转发路径。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 转发路径实例ID。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 会话保持时间。
	SessionExpire *int64 `json:"SessionExpire,omitempty" name:"SessionExpire"`

	// 是否开启健康检查。
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 健康检查检查路径。
	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`

	// 健康检查检查域名。
	HttpCheckDomain *string `json:"HttpCheckDomain,omitempty" name:"HttpCheckDomain"`

	// 健康检查检查间隔时间。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 健康检查健康阈值。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 健康检查不健康阈值。
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 健康检查中认为健康的HTTP返回码的组合。可选值为1~5的集合，1表示HTTP返回码为1xx认为健康。2表示HTTP返回码为2xx认为健康。3表示HTTP返回码为3xx认为健康。4表示HTTP返回码为4xx认为健康。5表示HTTP返回码为5xx认为健康。
	HttpCodes []*uint64 `json:"HttpCodes,omitempty" name:"HttpCodes"`

	// 均衡方式。
	BalanceMode *string `json:"BalanceMode,omitempty" name:"BalanceMode"`

	// 当前绑定关系的健康检查状态（Dead代表不健康，Alive代表健康）。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间戳。
	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`

	// 该转发路径所绑定的主机列表。
	BackendSet []*L7ListenerInfoBackend `json:"BackendSet,omitempty" name:"BackendSet"`
}

type L7ListenerInfoRule struct {
	// 转发域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发域名实例ID。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 当前绑定关系的健康检查状态（Dead代表不健康，Alive代表健康）。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间戳。
	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`

	// 该转发域名下面的转发路径列表。
	LocationSet []*L7ListenerInfoLocation `json:"LocationSet,omitempty" name:"LocationSet"`
}

type L7Rule struct {
	// 转发域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发域名实例ID。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径当前状态（0代表创建中，1代表正常运行，2代表创建失败，3代表删除中，4代表删除失败）。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间戳。
	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`

	// 该转发域名下面的转发路径列表。
	LocationSet []*L7RulesLocation `json:"LocationSet,omitempty" name:"LocationSet"`
}

type L7RulesLocation struct {
	// 转发路径。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 转发路径实例ID。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 会话保持时间。
	SessionExpire *int64 `json:"SessionExpire,omitempty" name:"SessionExpire"`

	// 是否开启健康检查。
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 健康检查检查路径。
	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`

	// 健康检查检查域名。
	HttpCheckDomain *string `json:"HttpCheckDomain,omitempty" name:"HttpCheckDomain"`

	// 健康检查检查间隔时间。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 健康检查健康阈值。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 健康检查不健康阈值。
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 健康检查中认为健康的HTTP返回码的组合。可选值为1~5的集合，1表示HTTP返回码为1xx认为健康。2表示HTTP返回码为2xx认为健康。3表示HTTP返回码为3xx认为健康。4表示HTTP返回码为4xx认为健康。5表示HTTP返回码为5xx认为健康。
	HttpCodes []*uint64 `json:"HttpCodes,omitempty" name:"HttpCodes"`

	// 均衡方式。
	BalanceMode *string `json:"BalanceMode,omitempty" name:"BalanceMode"`

	// 转发路径当前状态（0代表创建中，1代表正常运行，2代表创建失败，3代表删除中，4代表删除失败）。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间戳。
	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`
}

type LoadBalancer struct {
	// 负载均衡器ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 项目ID，通过v2/DescribeProject 接口获得
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 负载均衡器名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡的类型 : open表示公网负载均衡类型，internal表示内网负载均衡类型
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 是否筛选独占集群，0表示非独占集群，1表示四层独占集群，2表示七层独占集群，3表示四层和七层独占集群，4表示共享容灾
	Exclusive *uint64 `json:"Exclusive,omitempty" name:"Exclusive"`

	// 该负载均衡对应的tgw集群（fullnat,tunnel,dnat）
	TgwSetType *string `json:"TgwSetType,omitempty" name:"TgwSetType"`

	// 负载均衡域名。规则：1-60个小写英文字母、数字、点号“.”或连接线“-”。内网类型的负载均衡不能配置该字段
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 该负载均衡对应的所在的VpcId
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 该负载均衡对应的所在的SubnetId
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 无
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 无
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 无
	LatestPayMode *string `json:"LatestPayMode,omitempty" name:"LatestPayMode"`

	// 无
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 无
	StatusTime *string `json:"StatusTime,omitempty" name:"StatusTime"`

	// 私有网络名称。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 私有网络Cidr。
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 负载均衡的IPV4的VIP。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// 无
	SupportListenerTypes []*string `json:"SupportListenerTypes,omitempty" name:"SupportListenerTypes"`

	// 无
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 负载均衡个性化配置ID
	ConfId *string `json:"ConfId,omitempty" name:"ConfId"`

	// 无
	ConfName *string `json:"ConfName,omitempty" name:"ConfName"`

	// 负载均衡的IPV6的VIP。
	LoadBalancerVipv6s []*string `json:"LoadBalancerVipv6s,omitempty" name:"LoadBalancerVipv6s"`

	// 负载均衡IP协议类型。ipv4或者ipv6。
	IpProtocolType *string `json:"IpProtocolType,omitempty" name:"IpProtocolType"`

	// 保障型网关计费形式
	BzPayMode *string `json:"BzPayMode,omitempty" name:"BzPayMode"`

	// 保障型网关四层计费指标
	BzL4Metrics *string `json:"BzL4Metrics,omitempty" name:"BzL4Metrics"`

	// 保障型网关七层计费指标
	BzL7Metrics *string `json:"BzL7Metrics,omitempty" name:"BzL7Metrics"`

	// 该负载均衡对应的所在的整形类型的VpcId
	IntVpcId *uint64 `json:"IntVpcId,omitempty" name:"IntVpcId"`

	// 负载均衡的IPV6或者IPV4的VIP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurVips []*string `json:"CurVips,omitempty" name:"CurVips"`
}

type LoadBalancerPortInfoListener struct {
	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器协议类型，可选值：http，https，tcp，udp。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器的监听端口。
	LoadBalancerPort *int64 `json:"LoadBalancerPort,omitempty" name:"LoadBalancerPort"`

	// 计费模式为按固定带宽方式时监听器的限速值，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 监听器当前状态（0代表创建中，1代表正常运行，2代表创建失败，3代表删除中，4代表删除失败）。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 与监听器绑定的主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

// Predefined struct for user
type ModifyL4BackendPortRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡四层监听器ID，可通过接口DescribeL4Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 已绑定的主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 新的主机端口，可选值1~65535。
	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`

	// 绑定类型。0：物理机  1：虚拟机 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

type ModifyL4BackendPortRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡四层监听器ID，可通过接口DescribeL4Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 已绑定的主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 新的主机端口，可选值1~65535。
	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`

	// 绑定类型。0：物理机  1：虚拟机 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

func (r *ModifyL4BackendPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4BackendPortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "InstanceId")
	delete(f, "Port")
	delete(f, "NewPort")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4BackendPortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4BackendPortResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyL4BackendPortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4BackendPortResponseParams `json:"Response"`
}

func (r *ModifyL4BackendPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4BackendPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4BackendProbePortRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡四层监听器ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 已绑定的主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 新的探测端口，可选值1~65535。
	ProbePort *int64 `json:"ProbePort,omitempty" name:"ProbePort"`

	// 绑定类型。0：物理机 1：虚拟机IP 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

type ModifyL4BackendProbePortRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡四层监听器ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 已绑定的主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 新的探测端口，可选值1~65535。
	ProbePort *int64 `json:"ProbePort,omitempty" name:"ProbePort"`

	// 绑定类型。0：物理机 1：虚拟机IP 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

func (r *ModifyL4BackendProbePortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4BackendProbePortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "InstanceId")
	delete(f, "Port")
	delete(f, "ProbePort")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4BackendProbePortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4BackendProbePortResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyL4BackendProbePortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4BackendProbePortResponseParams `json:"Response"`
}

func (r *ModifyL4BackendProbePortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4BackendProbePortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4BackendWeightRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡四层监听器ID，可通过接口DescribeL4Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 权重信息，可选值0~100。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 已绑定的主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 绑定类型。0：物理机 1：虚拟机 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

type ModifyL4BackendWeightRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡四层监听器ID，可通过接口DescribeL4Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 权重信息，可选值0~100。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 已绑定的主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 绑定类型。0：物理机 1：虚拟机 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

func (r *ModifyL4BackendWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4BackendWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "InstanceId")
	delete(f, "Weight")
	delete(f, "Port")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4BackendWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4BackendWeightResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyL4BackendWeightResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4BackendWeightResponseParams `json:"Response"`
}

func (r *ModifyL4BackendWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4BackendWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ListenerRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 四层监听器ID。可通过接口DescribeL4Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 四层监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 会话保持时间，单位：秒。可选值：900~3600。
	SessionExpire *int64 `json:"SessionExpire,omitempty" name:"SessionExpire"`

	// 是否开启健康检查：1（开启）、0（关闭）。默认值0，表示关闭。
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 健康检查的响应超时时间，可选值：2-60，默认值：2，单位:秒。<br><font color="red">响应超时时间要小于检查间隔时间。</font>
	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 健康检查间隔，默认值：5，可选值：5-300，单位：秒。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 健康阈值，默认值：3，表示当连续探测三次健康则表示该转发正常，可选值：2-10，单位：次。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 不健康阈值，默认值：3，表示当连续探测三次不健康则表示该转发不正常，可选值：2-10，单位：次。
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 监听器最大带宽值，用于计费模式为固定带宽计费。可选值：0-1000，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 是否开启自定义健康检查：1（开启）、0（关闭）。默认值0，表示关闭。（该字段在健康检查开启的情况下才生效）
	CustomHealthSwitch *int64 `json:"CustomHealthSwitch,omitempty" name:"CustomHealthSwitch"`

	// 自定义健康探测内容类型，可选值：text（文本）、hexadecimal（十六进制）。
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 探测内容类型为文本方式时，针对请求文本中换行替换方式。可选值：1（替换为LF）、2（替换为CR）、3（替换为LF+CR）。
	LineSeparatorType *int64 `json:"LineSeparatorType,omitempty" name:"LineSeparatorType"`

	// 自定义探测请求内容。
	HealthRequest *string `json:"HealthRequest,omitempty" name:"HealthRequest"`

	// 自定义探测返回内容。
	HealthResponse *string `json:"HealthResponse,omitempty" name:"HealthResponse"`

	// 是否开启toa。可选值：0（关闭）、1（开启），默认关闭。（该字段在负载均衡为fullnat类型下才生效）
	ToaFlag *int64 `json:"ToaFlag,omitempty" name:"ToaFlag"`

	// 四层调度方式。wrr，wlc。
	BalanceMode *string `json:"BalanceMode,omitempty" name:"BalanceMode"`
}

type ModifyL4ListenerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 四层监听器ID。可通过接口DescribeL4Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 四层监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 会话保持时间，单位：秒。可选值：900~3600。
	SessionExpire *int64 `json:"SessionExpire,omitempty" name:"SessionExpire"`

	// 是否开启健康检查：1（开启）、0（关闭）。默认值0，表示关闭。
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 健康检查的响应超时时间，可选值：2-60，默认值：2，单位:秒。<br><font color="red">响应超时时间要小于检查间隔时间。</font>
	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 健康检查间隔，默认值：5，可选值：5-300，单位：秒。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 健康阈值，默认值：3，表示当连续探测三次健康则表示该转发正常，可选值：2-10，单位：次。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 不健康阈值，默认值：3，表示当连续探测三次不健康则表示该转发不正常，可选值：2-10，单位：次。
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 监听器最大带宽值，用于计费模式为固定带宽计费。可选值：0-1000，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 是否开启自定义健康检查：1（开启）、0（关闭）。默认值0，表示关闭。（该字段在健康检查开启的情况下才生效）
	CustomHealthSwitch *int64 `json:"CustomHealthSwitch,omitempty" name:"CustomHealthSwitch"`

	// 自定义健康探测内容类型，可选值：text（文本）、hexadecimal（十六进制）。
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 探测内容类型为文本方式时，针对请求文本中换行替换方式。可选值：1（替换为LF）、2（替换为CR）、3（替换为LF+CR）。
	LineSeparatorType *int64 `json:"LineSeparatorType,omitempty" name:"LineSeparatorType"`

	// 自定义探测请求内容。
	HealthRequest *string `json:"HealthRequest,omitempty" name:"HealthRequest"`

	// 自定义探测返回内容。
	HealthResponse *string `json:"HealthResponse,omitempty" name:"HealthResponse"`

	// 是否开启toa。可选值：0（关闭）、1（开启），默认关闭。（该字段在负载均衡为fullnat类型下才生效）
	ToaFlag *int64 `json:"ToaFlag,omitempty" name:"ToaFlag"`

	// 四层调度方式。wrr，wlc。
	BalanceMode *string `json:"BalanceMode,omitempty" name:"BalanceMode"`
}

func (r *ModifyL4ListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "SessionExpire")
	delete(f, "HealthSwitch")
	delete(f, "TimeOut")
	delete(f, "IntervalTime")
	delete(f, "HealthNum")
	delete(f, "UnhealthNum")
	delete(f, "Bandwidth")
	delete(f, "CustomHealthSwitch")
	delete(f, "InputType")
	delete(f, "LineSeparatorType")
	delete(f, "HealthRequest")
	delete(f, "HealthResponse")
	delete(f, "ToaFlag")
	delete(f, "BalanceMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL4ListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL4ListenerResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyL4ListenerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL4ListenerResponseParams `json:"Response"`
}

func (r *ModifyL4ListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL4ListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7BackendPortRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID，可通过接口DescribeL7Rules查询。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 已绑定的主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 新的主机端口，可选值1~65535。
	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`

	// 绑定类型。0：物理机 1：虚拟机 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

type ModifyL7BackendPortRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID，可通过接口DescribeL7Rules查询。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 已绑定的主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 新的主机端口，可选值1~65535。
	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`

	// 绑定类型。0：物理机 1：虚拟机 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

func (r *ModifyL7BackendPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7BackendPortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "DomainId")
	delete(f, "LocationId")
	delete(f, "InstanceId")
	delete(f, "Port")
	delete(f, "NewPort")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL7BackendPortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7BackendPortResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyL7BackendPortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL7BackendPortResponseParams `json:"Response"`
}

func (r *ModifyL7BackendPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7BackendPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7BackendWeightRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID，可通过接口DescribeL7Rules查询。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 权重信息，可选值0~100。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 已绑定的主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 绑定类型。0：物理机 1：虚拟机 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

type ModifyL7BackendWeightRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID，可通过接口DescribeL7Rules查询。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 权重信息，可选值0~100。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 已绑定的主机端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 绑定类型。0：物理机 1：虚拟机 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

func (r *ModifyL7BackendWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7BackendWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "DomainId")
	delete(f, "LocationId")
	delete(f, "InstanceId")
	delete(f, "Weight")
	delete(f, "Port")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL7BackendWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7BackendWeightResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyL7BackendWeightResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL7BackendWeightResponseParams `json:"Response"`
}

func (r *ModifyL7BackendWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7BackendWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7ListenerRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 七层监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 认证方式：0（不认证，用于http），1（单向认证，用于https），2（双向认证，用于https）。
	SslMode *int64 `json:"SslMode,omitempty" name:"SslMode"`

	// 服务端证书ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 服务端证书名称。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 服务端证书内容。
	CertContent *string `json:"CertContent,omitempty" name:"CertContent"`

	// 服务端证书密钥。
	CertKey *string `json:"CertKey,omitempty" name:"CertKey"`

	// 客户端证书ID。
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// 客户端证书名称。
	CertCaName *string `json:"CertCaName,omitempty" name:"CertCaName"`

	// 客户端证书内容。
	CertCaContent *string `json:"CertCaContent,omitempty" name:"CertCaContent"`

	// 计费模式为按固定带宽方式时监听器的限速值，可选值：0-1000，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 转发协议。当监听器Protocol为https时并且SslMode为1或2时，有意义。可选的值为0：https，1：spdy，2：http2，3：spdy+http2。
	ForwardProtocol *int64 `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`
}

type ModifyL7ListenerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 七层监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 认证方式：0（不认证，用于http），1（单向认证，用于https），2（双向认证，用于https）。
	SslMode *int64 `json:"SslMode,omitempty" name:"SslMode"`

	// 服务端证书ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 服务端证书名称。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 服务端证书内容。
	CertContent *string `json:"CertContent,omitempty" name:"CertContent"`

	// 服务端证书密钥。
	CertKey *string `json:"CertKey,omitempty" name:"CertKey"`

	// 客户端证书ID。
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// 客户端证书名称。
	CertCaName *string `json:"CertCaName,omitempty" name:"CertCaName"`

	// 客户端证书内容。
	CertCaContent *string `json:"CertCaContent,omitempty" name:"CertCaContent"`

	// 计费模式为按固定带宽方式时监听器的限速值，可选值：0-1000，单位：Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 转发协议。当监听器Protocol为https时并且SslMode为1或2时，有意义。可选的值为0：https，1：spdy，2：http2，3：spdy+http2。
	ForwardProtocol *int64 `json:"ForwardProtocol,omitempty" name:"ForwardProtocol"`
}

func (r *ModifyL7ListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7ListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "SslMode")
	delete(f, "CertId")
	delete(f, "CertName")
	delete(f, "CertContent")
	delete(f, "CertKey")
	delete(f, "CertCaId")
	delete(f, "CertCaName")
	delete(f, "CertCaContent")
	delete(f, "Bandwidth")
	delete(f, "ForwardProtocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL7ListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7ListenerResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用[DescribeLoadBalancerTaskResult](/document/product/386/9308)接口来查询任务操作结果
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyL7ListenerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL7ListenerResponseParams `json:"Response"`
}

func (r *ModifyL7ListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7ListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyL7LocationRule struct {
	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID，可通过接口DescribeL7Rules查询。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 转发路径。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 会话保持时间，单位：秒。可选值：30~3600。默认值0，表示不开启会话保持。
	SessionExpire *int64 `json:"SessionExpire,omitempty" name:"SessionExpire"`

	// 健康检查开关：1（开启）、0（关闭）。默认值0，表示关闭。
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 健康检查检查间隔时间，默认值：5，可选值：5-300，单位：秒。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 健康检查健康阈值，默认值：3，表示当连续探测三次健康则表示该转发正常，可选值：2-10，单位：次。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 健康检查不健康阈值，默认值：5，表示当连续探测五次不健康则表示该转发不正常，可选值：2-10，单位：次。
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 健康检查中认为健康的HTTP返回码的组合。可选值为1~5的集合，1表示HTTP返回码为1xx认为健康。2表示HTTP返回码为2xx认为健康。3表示HTTP返回码为3xx认为健康。4表示HTTP返回码为4xx认为健康。5表示HTTP返回码为5xx认为健康。
	HttpCodes []*uint64 `json:"HttpCodes,omitempty" name:"HttpCodes"`

	// 健康检查检查路径。
	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`

	// 健康检查检查域名。如果规则的域名使用通配符或正则表达式，则健康检查检查域名可自定义，否则必须跟健康检查检查域名一样。不填表示不修改。
	HttpCheckDomain *string `json:"HttpCheckDomain,omitempty" name:"HttpCheckDomain"`

	// 均衡方式：ip_hash、wrr。默认值wrr。
	BalanceMode *string `json:"BalanceMode,omitempty" name:"BalanceMode"`

	// 转发域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

// Predefined struct for user
type ModifyL7LocationsRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待更新的七层转发规则信息数组。
	RuleSet []*ModifyL7LocationRule `json:"RuleSet,omitempty" name:"RuleSet"`
}

type ModifyL7LocationsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待更新的七层转发规则信息数组。
	RuleSet []*ModifyL7LocationRule `json:"RuleSet,omitempty" name:"RuleSet"`
}

func (r *ModifyL7LocationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7LocationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "RuleSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyL7LocationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyL7LocationsResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyL7LocationsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyL7LocationsResponseParams `json:"Response"`
}

func (r *ModifyL7LocationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyL7LocationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancerChargeModeListener struct {
	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 协议类型。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 带宽。
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
}

// Predefined struct for user
type ModifyLoadBalancerChargeModeRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 计费方式。flow或bandwidth。
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 监听器信息，当计费方式选为 bandwidth 且此负载均衡实例下存在监听器时需填入此字段，可以自定义每个监听器带宽上限。
	ListenerSet []*ModifyLoadBalancerChargeModeListener `json:"ListenerSet,omitempty" name:"ListenerSet"`
}

type ModifyLoadBalancerChargeModeRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 计费方式。flow或bandwidth。
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 监听器信息，当计费方式选为 bandwidth 且此负载均衡实例下存在监听器时需填入此字段，可以自定义每个监听器带宽上限。
	ListenerSet []*ModifyLoadBalancerChargeModeListener `json:"ListenerSet,omitempty" name:"ListenerSet"`
}

func (r *ModifyLoadBalancerChargeModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerChargeModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "PayMode")
	delete(f, "ListenerSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerChargeModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerChargeModeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerChargeModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerChargeModeResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancerChargeModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerChargeModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡器名称，规则：1-20个英文、汉字、数字、连接线“-”或下划线“_”。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 域名前缀，负载均衡的域名由用户输入的域名前缀与配置文件中的域名后缀一起组合而成，保证是唯一的域名。规则：1-20个小写英文字母、数字或连接线“-”。内网类型的负载均衡不能配置该字段。
	DomainPrefix *string `json:"DomainPrefix,omitempty" name:"DomainPrefix"`
}

type ModifyLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡器名称，规则：1-20个英文、汉字、数字、连接线“-”或下划线“_”。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 域名前缀，负载均衡的域名由用户输入的域名前缀与配置文件中的域名后缀一起组合而成，保证是唯一的域名。规则：1-20个小写英文字母、数字或连接线“-”。内网类型的负载均衡不能配置该字段。
	DomainPrefix *string `json:"DomainPrefix,omitempty" name:"DomainPrefix"`
}

func (r *ModifyLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "LoadBalancerName")
	delete(f, "DomainPrefix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceCertRequestParams struct {
	// 要被替换的证书ID
	OldCertId *string `json:"OldCertId,omitempty" name:"OldCertId"`

	// 证书内容
	NewCert *string `json:"NewCert,omitempty" name:"NewCert"`

	// 证书名称
	NewAlias *string `json:"NewAlias,omitempty" name:"NewAlias"`

	// 私钥内容，证书类型为SVR时不需要传递
	NewKey *string `json:"NewKey,omitempty" name:"NewKey"`

	// 是否删除旧证书，0 表示不删除，1 表示删除
	DeleteOld *uint64 `json:"DeleteOld,omitempty" name:"DeleteOld"`
}

type ReplaceCertRequest struct {
	*tchttp.BaseRequest
	
	// 要被替换的证书ID
	OldCertId *string `json:"OldCertId,omitempty" name:"OldCertId"`

	// 证书内容
	NewCert *string `json:"NewCert,omitempty" name:"NewCert"`

	// 证书名称
	NewAlias *string `json:"NewAlias,omitempty" name:"NewAlias"`

	// 私钥内容，证书类型为SVR时不需要传递
	NewKey *string `json:"NewKey,omitempty" name:"NewKey"`

	// 是否删除旧证书，0 表示不删除，1 表示删除
	DeleteOld *uint64 `json:"DeleteOld,omitempty" name:"DeleteOld"`
}

func (r *ReplaceCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OldCertId")
	delete(f, "NewCert")
	delete(f, "NewAlias")
	delete(f, "NewKey")
	delete(f, "DeleteOld")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceCertResponseParams struct {
	// 新证书ID。
	NewCertId *string `json:"NewCertId,omitempty" name:"NewCertId"`

	// 旧证书ID。
	OldCertId *string `json:"OldCertId,omitempty" name:"OldCertId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReplaceCertResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceCertResponseParams `json:"Response"`
}

func (r *ReplaceCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTrafficMirrorAliasRequestParams struct {
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 流量镜像实例别名。
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type SetTrafficMirrorAliasRequest struct {
	*tchttp.BaseRequest
	
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 流量镜像实例别名。
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *SetTrafficMirrorAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTrafficMirrorAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficMirrorId")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTrafficMirrorAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTrafficMirrorAliasResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetTrafficMirrorAliasResponse struct {
	*tchttp.BaseResponse
	Response *SetTrafficMirrorAliasResponseParams `json:"Response"`
}

func (r *SetTrafficMirrorAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTrafficMirrorAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTrafficMirrorHealthSwitchRequestParams struct {
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 健康检查开关，0：关闭，1：打开
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 健康检查判断健康的次数，最小值2，最大值10。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 健康检查判断不健康的次数，最小值2，最大值10。
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 健康检查间隔，单位：秒，最小值5，最大值300。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 检查的域名配置。
	HttpCheckDomain *string `json:"HttpCheckDomain,omitempty" name:"HttpCheckDomain"`

	// 检查的路径配置。
	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`

	// 健康检查中认为健康的HTTP返回码的组合。可选值为1~5的集合，1表示HTTP返回码为1xx认为健康。2表示HTTP返回码为2xx认为健康。3表示HTTP返回码为3xx认为健康。4表示HTTP返回码为4xx认为健康。5表示HTTP返回码为5xx认为健康。
	HttpCodes []*int64 `json:"HttpCodes,omitempty" name:"HttpCodes"`
}

type SetTrafficMirrorHealthSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 健康检查开关，0：关闭，1：打开
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 健康检查判断健康的次数，最小值2，最大值10。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 健康检查判断不健康的次数，最小值2，最大值10。
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 健康检查间隔，单位：秒，最小值5，最大值300。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 检查的域名配置。
	HttpCheckDomain *string `json:"HttpCheckDomain,omitempty" name:"HttpCheckDomain"`

	// 检查的路径配置。
	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`

	// 健康检查中认为健康的HTTP返回码的组合。可选值为1~5的集合，1表示HTTP返回码为1xx认为健康。2表示HTTP返回码为2xx认为健康。3表示HTTP返回码为3xx认为健康。4表示HTTP返回码为4xx认为健康。5表示HTTP返回码为5xx认为健康。
	HttpCodes []*int64 `json:"HttpCodes,omitempty" name:"HttpCodes"`
}

func (r *SetTrafficMirrorHealthSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTrafficMirrorHealthSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficMirrorId")
	delete(f, "HealthSwitch")
	delete(f, "HealthNum")
	delete(f, "UnhealthNum")
	delete(f, "IntervalTime")
	delete(f, "HttpCheckDomain")
	delete(f, "HttpCheckPath")
	delete(f, "HttpCodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTrafficMirrorHealthSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTrafficMirrorHealthSwitchResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetTrafficMirrorHealthSwitchResponse struct {
	*tchttp.BaseResponse
	Response *SetTrafficMirrorHealthSwitchResponseParams `json:"Response"`
}

func (r *SetTrafficMirrorHealthSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTrafficMirrorHealthSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrafficMirror struct {
	// 流量镜像ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 流量镜像名称。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 流量镜像所在的私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 接收机负载均衡方式。wrr，ip_hash，wlc。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 是否开始对接收机的健康检查。0：关闭，非0：开启。
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 健康阈值。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 不健康阈值。
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 检查间隔。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 检查域名。
	HttpCheckDomain *string `json:"HttpCheckDomain,omitempty" name:"HttpCheckDomain"`

	// 检查目录。
	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`

	// 健康检查返回码。 1：1xx，2：2xx，3：3xx，4：4xx，5：5xx。
	HttpCodes []*int64 `json:"HttpCodes,omitempty" name:"HttpCodes"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 流量镜像所在私有网络的Cidr。
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 流量镜像所在私有网络的名称。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type TrafficMirrorListener struct {
	// 监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 七层监听器协议类型，可选值：http,https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器的监听端口。
	LoadBalancerPort *uint64 `json:"LoadBalancerPort,omitempty" name:"LoadBalancerPort"`

	// 当前带宽。
	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 带宽上限。
	MaxBandwidth *uint64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// 监听器类型。
	ListenerType *string `json:"ListenerType,omitempty" name:"ListenerType"`

	// 认证方式：0（不认证，用于http），1（单向认证，用于https），2（双向认证，用于https）。
	SslMode *int64 `json:"SslMode,omitempty" name:"SslMode"`

	// 服务端证书ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 客户端证书ID。
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// 添加时间。
	AddTimestamp *string `json:"AddTimestamp,omitempty" name:"AddTimestamp"`

	// 负载均衡ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 私有网络名称。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 私有网络Cidr。
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 负载均衡的VIP。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// 负载均衡名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡的IPV6的VIP。
	LoadBalancerVipv6s []*string `json:"LoadBalancerVipv6s,omitempty" name:"LoadBalancerVipv6s"`

	// 支持的IP协议类型。ipv4或者是ipv6。
	IpProtocolType *string `json:"IpProtocolType,omitempty" name:"IpProtocolType"`
}

type TrafficMirrorPortStatus struct {
	// 接收机端口。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 状态。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type TrafficMirrorReceiver struct {
	// 接收机实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 接收机接收端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 接收机权重。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 流量镜像ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 接收机别名。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 接收机内网IP地址。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 接收机所在的子网的ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 接收机所在的子网的名称。
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 接收机所在的子网的Cidr。
	SubnetCidrBlock *string `json:"SubnetCidrBlock,omitempty" name:"SubnetCidrBlock"`

	// 接收机所在的私有网络的ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 接收机所在的私有网络的名称。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 接收机所在的私有网络的Cidr。
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 接收机的健康状态。
	HealthStatus *string `json:"HealthStatus,omitempty" name:"HealthStatus"`

	// 接收机的可以执行的操作集合。
	Operates []*string `json:"Operates,omitempty" name:"Operates"`
}

type TrafficMirrorReciversStatus struct {
	// 内网IP。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 端口及对应的状态。
	ReceiversPortStatusSet []*TrafficMirrorPortStatus `json:"ReceiversPortStatusSet,omitempty" name:"ReceiversPortStatusSet"`
}

type UnbindL4Backend struct {
	// 待解绑的主机端口，可选值1~65535。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type UnbindL4BackendsRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡四层监听器ID，可通过接口DescribeL4Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待解绑的主机信息。可以绑定多个主机端口。目前一个四层监听器下面最多允许绑定255个主机端口。
	BackendSet []*UnbindL4Backend `json:"BackendSet,omitempty" name:"BackendSet"`

	// 绑定类型。0：物理机 1：虚拟机 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

type UnbindL4BackendsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡四层监听器ID，可通过接口DescribeL4Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待解绑的主机信息。可以绑定多个主机端口。目前一个四层监听器下面最多允许绑定255个主机端口。
	BackendSet []*UnbindL4Backend `json:"BackendSet,omitempty" name:"BackendSet"`

	// 绑定类型。0：物理机 1：虚拟机 2：半托管机器
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`
}

func (r *UnbindL4BackendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindL4BackendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "BackendSet")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindL4BackendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindL4BackendsResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindL4BackendsResponse struct {
	*tchttp.BaseResponse
	Response *UnbindL4BackendsResponseParams `json:"Response"`
}

func (r *UnbindL4BackendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindL4BackendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindL7Backend struct {
	// 待解绑的主机端口，可选值1~65535。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 黑石物理机主机ID、虚拟机IP或者是半托管主机ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type UnbindL7BackendsRequestParams struct {
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID，可通过接口DescribeL7Rules查询。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 待绑定的主机信息。
	BackendSet []*UnbindL7Backend `json:"BackendSet,omitempty" name:"BackendSet"`

	// 绑定类型。0：物理机  1：虚拟机 2：半托管机器
	BindType *uint64 `json:"BindType,omitempty" name:"BindType"`
}

type UnbindL7BackendsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可通过接口DescribeLoadBalancers查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 七层监听器实例ID，可通过接口DescribeL7Listeners查询。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发域名实例ID，可通过接口DescribeL7Rules查询。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 转发路径实例ID，可通过接口DescribeL7Rules查询。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 待绑定的主机信息。
	BackendSet []*UnbindL7Backend `json:"BackendSet,omitempty" name:"BackendSet"`

	// 绑定类型。0：物理机  1：虚拟机 2：半托管机器
	BindType *uint64 `json:"BindType,omitempty" name:"BindType"`
}

func (r *UnbindL7BackendsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindL7BackendsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "DomainId")
	delete(f, "LocationId")
	delete(f, "BackendSet")
	delete(f, "BindType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindL7BackendsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindL7BackendsResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindL7BackendsResponse struct {
	*tchttp.BaseResponse
	Response *UnbindL7BackendsResponseParams `json:"Response"`
}

func (r *UnbindL7BackendsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindL7BackendsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindTrafficMirrorListenersRequestParams struct {
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 七层监听器实例ID数组，可通过接口DescribeL7Listeners查询。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

type UnbindTrafficMirrorListenersRequest struct {
	*tchttp.BaseRequest
	
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 七层监听器实例ID数组，可通过接口DescribeL7Listeners查询。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

func (r *UnbindTrafficMirrorListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindTrafficMirrorListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficMirrorId")
	delete(f, "ListenerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindTrafficMirrorListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindTrafficMirrorListenersResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindTrafficMirrorListenersResponse struct {
	*tchttp.BaseResponse
	Response *UnbindTrafficMirrorListenersResponseParams `json:"Response"`
}

func (r *UnbindTrafficMirrorListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindTrafficMirrorListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindTrafficMirrorReceiver struct {
	// 待解绑的主机端口，可选值1~65535。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 待解绑的主机实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type UnbindTrafficMirrorReceiversRequestParams struct {
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 待绑定的主机实例ID和端口数组。
	ReceiverSet []*UnbindTrafficMirrorReceiver `json:"ReceiverSet,omitempty" name:"ReceiverSet"`
}

type UnbindTrafficMirrorReceiversRequest struct {
	*tchttp.BaseRequest
	
	// 流量镜像实例ID。
	TrafficMirrorId *string `json:"TrafficMirrorId,omitempty" name:"TrafficMirrorId"`

	// 待绑定的主机实例ID和端口数组。
	ReceiverSet []*UnbindTrafficMirrorReceiver `json:"ReceiverSet,omitempty" name:"ReceiverSet"`
}

func (r *UnbindTrafficMirrorReceiversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindTrafficMirrorReceiversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficMirrorId")
	delete(f, "ReceiverSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindTrafficMirrorReceiversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindTrafficMirrorReceiversResponseParams struct {
	// 任务ID。该接口为异步任务，可根据本参数调用DescribeLoadBalancerTaskResult接口来查询任务操作结果。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindTrafficMirrorReceiversResponse struct {
	*tchttp.BaseResponse
	Response *UnbindTrafficMirrorReceiversResponseParams `json:"Response"`
}

func (r *UnbindTrafficMirrorReceiversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindTrafficMirrorReceiversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadCertRequestParams struct {
	// 证书类型，可选值：CA，SVR。
	CertType *string `json:"CertType,omitempty" name:"CertType"`

	// 证书内容。
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// 证书别名。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 私钥内容，证书类型为SVR时不需要传递。
	Key *string `json:"Key,omitempty" name:"Key"`
}

type UploadCertRequest struct {
	*tchttp.BaseRequest
	
	// 证书类型，可选值：CA，SVR。
	CertType *string `json:"CertType,omitempty" name:"CertType"`

	// 证书内容。
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// 证书别名。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 私钥内容，证书类型为SVR时不需要传递。
	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *UploadCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertType")
	delete(f, "Cert")
	delete(f, "Alias")
	delete(f, "Key")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadCertResponseParams struct {
	// 新建的证书ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadCertResponse struct {
	*tchttp.BaseResponse
	Response *UploadCertResponseParams `json:"Response"`
}

func (r *UploadCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}