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

package v20180808

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ApiEnvironmentStrategy struct {

	// API唯一ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 用户自定义API名称。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API的路径。如/path。
	Path *string `json:"Path,omitempty" name:"Path"`

	// API的方法。如GET。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 环境的限流信息。
	EnvironmentStrategySet []*EnvironmentStrategy `json:"EnvironmentStrategySet,omitempty" name:"EnvironmentStrategySet" list`
}

type ApiEnvironmentStrategyStataus struct {

	// API绑定的限流策略数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API绑定的限流策略列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiEnvironmentStrategySet []*ApiEnvironmentStrategy `json:"ApiEnvironmentStrategySet,omitempty" name:"ApiEnvironmentStrategySet" list`
}

type ApiIdStatus struct {

	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API唯一ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API描述
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// API PATH。
	Path *string `json:"Path,omitempty" name:"Path"`

	// API METHOD。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 服务创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 服务修改时间。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// API名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// VPC唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// API类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API协议。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 是否买后调试。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// 授权类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// API业务类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// 关联授权API唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// 授权API关联的业务API列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds" list`

	// oauth配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// oauth2.0API请求，token存放位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`
}

type ApiInfo struct {

	// API 所在的服务唯一 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 所在的服务的名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// API 所在的服务的描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// API 接口唯一 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API 接口的描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// 创建时间，按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 最后修改时间，按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// API 接口的名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API 类型。可取值为NORMAL（普通API）、TSF（微服务API）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API 的前端请求类型，如 HTTP 或 HTTPS 或者 HTTP 和 HTTPS。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// API 鉴权类型。可取值为 SECRET（密钥对鉴权）、NONE（免鉴权）、OAUTH。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// OAUTH API的类型。可取值为NORMAL（业务API）、OAUTH（授权API）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// OAUTH 业务API 关联的授权API 唯一 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// OAUTH配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// 是否购买后调试（云市场预留参数）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// 请求的前端配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// 返回类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// 自定义响应配置成功响应示例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// 自定义响应配置失败响应示例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// 用户自定义错误码配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseErrorCodes []*ErrorCodes `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes" list`

	// 前端请求参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters" list`

	// API 的后端服务超时时间，单位是秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// API 的后端服务类型。可取值为 HTTP、MOCK、TSF、CLB、SCF、WEBSOCKET、TARGET（内测）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// API 的后端服务配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// API的后端服务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters" list`

	// 常量参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters" list`

	// API 的后端 Mock 返回信息。如果 ServiceType 是 Mock，则此参数必传。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// scf 函数名称。当后端类型是SCF时生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// scf 函数命名空间。当后端类型是SCF时生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// scf函数版本。当后端类型是SCF时生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// 是否开启集成响应。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// scf websocket注册函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// scf websocket注册函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// scf websocket传输函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// scf websocket清理函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// scf websocket清理函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// scf websocket清理函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// WEBSOCKET 回推地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternalDomain *string `json:"InternalDomain,omitempty" name:"InternalDomain"`

	// scf websocket传输函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// scf websocket传输函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// scf websocket传输函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// API绑定微服务服务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroServices []*MicroService `json:"MicroServices,omitempty" name:"MicroServices" list`

	// 微服务信息详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroServicesInfo []*int64 `json:"MicroServicesInfo,omitempty" name:"MicroServicesInfo" list`

	// 微服务的负载均衡配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// 微服务的健康检查配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// 是否开启跨域。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// API绑定的tag信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// API已发布的环境信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environments []*string `json:"Environments,omitempty" name:"Environments" list`
}

type ApiKey struct {

	// 创建的 API 密钥 ID 。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// 创建的 API 密钥 Key。
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`

	// 密钥类型，auto 或者 manual。
	AccessKeyType *string `json:"AccessKeyType,omitempty" name:"AccessKeyType"`

	// 用户自定义密钥名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 最后一次修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 密钥状态。0表示禁用，1表示启用。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ApiKeysStatus struct {

	// 符合条件的 API 密钥数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API 密钥列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiKeySet []*ApiKey `json:"ApiKeySet,omitempty" name:"ApiKeySet" list`
}

type ApiRequestConfig struct {

	// path
	Path *string `json:"Path,omitempty" name:"Path"`

	// 方法
	Method *string `json:"Method,omitempty" name:"Method"`
}

type ApiUsagePlan struct {

	// 服务唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 唯一 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API 名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API 路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// API 方法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 使用计划的唯一 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 使用计划的名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 使用计划的描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 使用计划绑定的服务环境。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 已经使用的配额。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InUseRequestNum *int64 `json:"InUseRequestNum,omitempty" name:"InUseRequestNum"`

	// 请求配额总量，-1表示没有限制。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 请求 QPS 上限，-1 表示没有限制。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// 使用计划创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 使用计划最后修改时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 服务名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type ApiUsagePlanSet struct {

	// API 绑定的使用计划总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API 绑定使用计划列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiUsagePlanList []*ApiUsagePlan `json:"ApiUsagePlanList,omitempty" name:"ApiUsagePlanList" list`
}

type ApisStatus struct {

	// 符合条件的 API 接口数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API 接口列表。
	ApiIdStatusSet []*DesApisStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet" list`
}

type BindEnvironmentRequest struct {
	*tchttp.BaseRequest

	// 待绑定的使用计划唯一 ID 列表。
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds" list`

	// 绑定类型，取值为API、SERVICE，默认值为SERVICE。
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// 待绑定的环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 待绑定的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API唯一ID数组，当bindType=API时，需要传入此参数。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *BindEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindEnvironmentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 绑定操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindEnvironmentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 待绑定的IP策略所属的服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待绑定的IP策略唯一ID。
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// IP策略待绑定的环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// IP策略待绑定的API列表。
	BindApiIds []*string `json:"BindApiIds,omitempty" name:"BindApiIds" list`
}

func (r *BindIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 绑定操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSecretIdsRequest struct {
	*tchttp.BaseRequest

	// 待绑定的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 待绑定的密钥 ID 数组。
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds" list`
}

func (r *BindSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSecretIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 绑定操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSecretIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSubDomainRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待绑定的自定义的域名。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 服务支持协议，可选值为http、https、http&https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 网络类型，可选值为OUTER、INNER。
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 是否使用默认路径映射，默认为 true。为 false 时，表示自定义路径映射，此时 PathMappingSet 必填。
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// 默认域名。
	NetSubDomain *string `json:"NetSubDomain,omitempty" name:"NetSubDomain"`

	// 待绑定自定义域名的证书唯一 ID。针对Protocol 为https或http&https可以选择上传。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 自定义域名路径映射，最多输入三个Environment，并且只能分别取值“test”、 ”prepub“、”release“。
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet" list`
}

func (r *BindSubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSubDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BindSubDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindSubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindSubDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConstantParameter struct {

	// 常量参数名称。只有 ServiceType 是 HTTP 才会用到此参数。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 常量参数描述。只有 ServiceType 是 HTTP 才会用到此参数。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 常量参数位置。只有 ServiceType 是 HTTP 才会用到此参数。
	Position *string `json:"Position,omitempty" name:"Position"`

	// 常量参数默认值。只有 ServiceType 是 HTTP 才会用到此参数。
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
}

type CreateApiKeyRequest struct {
	*tchttp.BaseRequest

	// 用户自定义密钥名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 密钥类型，支持 auto 和 manual（自定义密钥），默认为 auto。
	AccessKeyType *string `json:"AccessKeyType,omitempty" name:"AccessKeyType"`

	// 用户自定义密钥 ID，AccessKeyType 为 manual 时必传。长度为5 - 50字符，由字母、数字、英文下划线组成。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// 用户自定义密钥 Key，AccessKeyType 为 manual 时必传。长度为10 - 50字符，由字母、数字、英文下划线。
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
}

func (r *CreateApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新增的密钥详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ApiKey `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiRequest struct {
	*tchttp.BaseRequest

	// API 所在的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 的后端服务类型。支持HTTP、MOCK、TSF、SCF、WEBSOCKET、TARGET（内测）。
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// API 的后端服务超时时间，单位是秒。
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// API 的前端请求协议，支持HTTP和WEBSOCKET。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 请求的前端配置。
	RequestConfig *ApiRequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// 用户自定义的 API 名称。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// 用户自定义的 API 接口描述。
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// API 类型，支持NORMAL（普通API）和TSF（微服务API），默认为NORMAL。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API 鉴权类型。支持SECRET（密钥对鉴权）、NONE（免鉴权）、OAUTH。默认为NONE。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 是否开启跨域。
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// 常量参数。
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters" list`

	// 前端请求参数。
	RequestParameters []*RequestParameter `json:"RequestParameters,omitempty" name:"RequestParameters" list`

	// 当AuthType 为 OAUTH时，该字段有效， NORMAL：业务api OAUTH：授权API。
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// API 的后端 Mock 返回信息。如果 ServiceType 是 Mock，则此参数必传。
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// API绑定微服务服务列表。
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices" list`

	// 微服务的负载均衡配置。
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// 微服务的健康检查配置。
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// target类型后端资源信息。（内测阶段）
	TargetServices []*TargetServicesReq `json:"TargetServices,omitempty" name:"TargetServices" list`

	// target类型负载均衡配置。（内测阶段）
	TargetServicesLoadBalanceConf *int64 `json:"TargetServicesLoadBalanceConf,omitempty" name:"TargetServicesLoadBalanceConf"`

	// target健康检查配置。（内测阶段）
	TargetServicesHealthCheckConf *HealthCheckConf `json:"TargetServicesHealthCheckConf,omitempty" name:"TargetServicesHealthCheckConf"`

	// scf 函数名称。当后端类型是SCF时生效。
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// scf websocket注册函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// scf websocket清理函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// scf websocket传输函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// scf 函数命名空间。当后端类型是SCF时生效。
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// scf函数版本。当后端类型是SCF时生效。
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// scf websocket注册函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// scf websocket传输函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// scf websocket传输函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// scf websocket传输函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// scf websocket清理函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// scf websocket清理函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// 是否开启响应集成。当后端类型是SCF时生效。
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// 开始调试后计费。（云市场预留字段）
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// 是否删除自定义响应配置错误码，如果不传或者传 False，不删除，当传 True 时，则删除此 API 所有自定义响应配置错误码。
	IsDeleteResponseErrorCodes *bool `json:"IsDeleteResponseErrorCodes,omitempty" name:"IsDeleteResponseErrorCodes"`

	// 返回类型。
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// 自定义响应配置成功响应示例。
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// 自定义响应配置失败响应示例。
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// API 的后端服务配置。
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// 关联的授权API 唯一 ID，当AuthType为OAUTH且ApiBusinessType为NORMAL时生效。标示业务API绑定的oauth2.0授权API唯一ID。
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// API的后端服务参数。
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters" list`

	// oauth配置。当AuthType是OAUTH时生效。
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// 用户自定义错误码配置。
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes" list`

	// tsf serverless 命名空间ID。（内测中）
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" name:"TargetNamespaceId"`

	// 用户类型。
	UserType *string `json:"UserType,omitempty" name:"UserType"`
}

func (r *CreateApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// api信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *CreateApiRsp `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApiRsp struct {

	// api id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// path
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// method
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type CreateIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务的唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 用户自定义的策略名称。
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 策略类型。支持WHITE（白名单）和BLACK（黑名单）。
	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`

	// 策略详情。
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

func (r *CreateIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新建的IP策略详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *IPStrategy `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceRequest struct {
	*tchttp.BaseRequest

	// 用户自定义的服务名称。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 服务的前端请求类型。如 http、https、http&https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 用户自定义的服务描述。
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// 独立集群名称，用于指定创建服务所在的独立集群。
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// 网络类型列表，用于指定支持的访问类型，INNER为内网访问，OUTER为外网访问。默认为OUTER。
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`

	// IP版本号，支持IPv4和IPv6，默认为IPv4。
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// 集群名称。保留字段，tsf serverlss类型使用。
	SetServerName *string `json:"SetServerName,omitempty" name:"SetServerName"`

	// 用户类型。保留类型，serverless用户使用。
	AppIdType *string `json:"AppIdType,omitempty" name:"AppIdType"`

	// 标签。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

func (r *CreateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务唯一ID。
		ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

		// 用户自定义服务名称。
		ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

		// 用户自定义服务描述。
		ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

		// 外网默认域名。
		OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

		// vpc内网默认域名。
		InnerSubDomain *string `json:"InnerSubDomain,omitempty" name:"InnerSubDomain"`

		// 服务创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

		// 网络类型列表，INNER为内网访问，OUTER为外网访问。
		NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`

		// IP版本号。
	// 注意：此字段可能返回 null，表示取不到有效值。
		IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 用户自定义的使用计划名称。
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 用户自定义的使用计划描述。
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 请求配额总数，取值范围为-1或者[1, 99999999]，默认为-1，表示不开启。
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 每秒请求限制数，取值范围为-1或者[1, 2000]，默认-1，表示不开启。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
}

func (r *CreateUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 使用计划详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiKeyRequest struct {
	*tchttp.BaseRequest

	// 待删除的密钥 ID。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DeleteApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiRequest struct {
	*tchttp.BaseRequest

	// API 所在的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 接口唯一 ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DeleteApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 待删除的IP策略所属的服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待删除的IP策略唯一ID。
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DeleteIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceRequest struct {
	*tchttp.BaseRequest

	// 待删除服务的唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DeleteServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceSubDomainMappingRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务绑定的自定义域名。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 待删除映射的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *DeleteServiceSubDomainMappingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceSubDomainMappingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceSubDomainMappingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除自定义域名的路径映射操作是否成功。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceSubDomainMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceSubDomainMappingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 待删除的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
}

func (r *DeleteUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DemoteServiceUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 使用计划ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 待降级的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名称。
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *DemoteServiceUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DemoteServiceUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DemoteServiceUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 降级操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DemoteServiceUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DemoteServiceUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DesApisStatus struct {

	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API唯一ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 用户自定义的 API 接口描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 最后修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// API 接口的名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// VPCID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// VPC唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// API类型。取值为NORMAL（普通API）和TSF（微服务API）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API协议。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 是否买后调试。（云市场预留字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// API 鉴权类型。取值为SECRET（密钥对鉴权）、NONE（免鉴权）、OAUTH。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// OAUTH API的类型。当AuthType 为 OAUTH时该字段有效， 取值为NORMAL（业务API）和 OAUTH（授权API）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// 关联的授权API 唯一 ID，当AuthType为OAUTH且ApiBusinessType为NORMAL时生效。标示业务API绑定的oauth2.0授权API唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// OAUTH 配置信息。当AuthType是OAUTH时生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// 授权API关联的业务API列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds" list`

	// API关联的标签信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// API 的路径，如 /path。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// API 的请求方法，如 GET。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`
}

type DescribeApiEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// API所属服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境列表。
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames" list`

	// API唯一ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeApiEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiEnvironmentStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// api绑定策略详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ApiEnvironmentStrategyStataus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiEnvironmentStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeyRequest struct {
	*tchttp.BaseRequest

	// API 密钥 ID。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DescribeApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 密钥详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ApiKey `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeysStatusRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持AccessKeyId、AccessKeySecret、SecretName、NotUsagePlanId、Status、KeyWord（ 可以匹配name或者path）。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeApiKeysStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiKeysStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiKeysStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 密钥列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ApiKeysStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiKeysStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiKeysStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiRequest struct {
	*tchttp.BaseRequest

	// API 所在的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 接口唯一 ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DescribeApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API 详情。
		Result *ApiInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeApiUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApiUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// api绑定使用计划列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ApiUsagePlanSet `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApiUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApiUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApisStatusRequest struct {
	*tchttp.BaseRequest

	// API 所在的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// API过滤条件。支持ApiId、ApiName、ApiPath、ApiType、AuthRelationApiId、AuthType、ApiBuniessType、NotUsagePlanId、Environment、Tags (values为 $tag_key:tag_value的列表)、TagKeys （values 为 tag key的列表）。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeApisStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApisStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApisStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// API 详情列表。
		Result *ApisStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApisStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApisStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategyApisStatusRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 策略唯一ID。
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 策略所在环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持 ApiPath、ApiName、KeyWord（模糊查询Path 和Name）。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeIPStrategyApisStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIPStrategyApisStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategyApisStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境绑定API列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *IPStrategyApiStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPStrategyApisStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIPStrategyApisStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// IP 策略唯一ID。
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 策略关联的环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。预留字段，目前不支持过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// IP策略详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *IPStrategy `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategysStatusRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 过滤条件。支持StrategyName。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeIPStrategysStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIPStrategysStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIPStrategysStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的策略列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *IPStrategysStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIPStrategysStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIPStrategysStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSearchRequest struct {
	*tchttp.BaseRequest

	// 日志开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 保留字段
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 根据上次返回的ConText，获取后续的内容，最多可获取10000条
	ConText *string `json:"ConText,omitempty" name:"ConText"`

	// 按时间排序 asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 保留字段
	Query *string `json:"Query,omitempty" name:"Query"`

	// 检索条件,支持的检索条件如下：
	// req_id：“=”
	// api_id：“=”
	// cip：“=”
	// uip：“:”
	// err_msg：“:”
	// rsp_st：“=” 、“!=” 、 “:” 、 “>” 、 “<”
	// req_t：”>=“ 、 ”<=“
	// 
	// 说明：
	// “:”表示包含，“!=”表示不等于，字段含义见输出参数的LogSet说明
	LogQuerys []*LogQuery `json:"LogQuerys,omitempty" name:"LogQuerys" list`
}

func (r *DescribeLogSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogSearchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSearchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 获取更多检索结果的游标，值为""表示无后续结果
		ConText *string `json:"ConText,omitempty" name:"ConText"`

		// 由0或多条日志组成，每条日志格式如下：
	// '[$app_id][$env_name][$service_id][$http_host][$api_id][$uri][$scheme][rsp_st:$status][ups_st:$upstream_status]'
	// '[cip:$remote_addr][uip:$upstream_addr][vip:$server_addr][rsp_len:$bytes_sent][req_len:$request_length]'
	// '[req_t:$request_time][ups_rsp_t:$upstream_response_time][ups_conn_t:$upstream_connect_time][ups_head_t:$upstream_header_time]’
	// '[err_msg:$err_msg][tcp_rtt:$tcpinfo_rtt][$pid][$time_local][req_id:$request_id]';
	// 
	// 说明：
	// app_id： 用户 ID。
	// env_name：环境名称。
	// service_id： 服务 ID。
	// http_host： 域名。
	// api_id： API 的 ID。
	// uri：请求的路径。
	// scheme： HTTP/HTTPS 协议。
	// rsp_st： 请求响应状态码。
	// ups_st： 后端业务服务器的响应状态码（如果请求透传到后端，改变量不为空。如果请求在 APIGW 就被拦截了，那么该变量显示为 -）。
	// cip： 客户端 IP。
	// uip： 后端业务服务（upstream）的 IP。
	// vip： 请求访问的 VIP。
	// rsp_len： 响应长度。
	// req_len： 请求长度。
	// req_t： 请求响应的总时间。
	// ups_rsp_t： 后端响应的总时间（apigw 建立连接到接收到后端响应的时间）。
	// ups_conn_t: 与后端业务服务器连接建立成功时间。
	// ups_head_t：后端响应的头部到达时间。
	// err_msg： 错误信息。
	// tcp_rtt： 客户端 TCP 连接信息，RTT（Round Trip Time）由三部分组成：链路的传播时间（propagation delay）、末端系统的处理时间、路由器缓存中的排队和处理时间（queuing delay）。
	// req_id：请求id。
		LogSet []*string `json:"LogSet,omitempty" name:"LogSet" list`

		// 单次搜索返回的日志条数，TotalCount <= Limit
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLogSearchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentListRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务绑定环境详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ServiceEnvironmentSet `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentReleaseHistoryRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名称。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentReleaseHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentReleaseHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentReleaseHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务发布历史。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ServiceReleaseHistory `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentReleaseHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentReleaseHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 限流策略列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ServiceEnvironmentStrategyStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceEnvironmentStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceReleaseVersionRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceReleaseVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceReleaseVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceReleaseVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务发布版本列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ServiceReleaseVersion `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceReleaseVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceReleaseVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *DescribeServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务唯一ID。
		ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

		// 服务 环境列表。
		AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments" list`

		// 服务名称。
		ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

		// 服务描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

		// 服务支持协议，可选值为http、https、http&https。
		Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

		// 服务创建时间。
		CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

		// 服务修改时间。
		ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

		// 独立集群名称。
		ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

		// 网络类型列表，INNER为内网访问，OUTER为外网访问。
		NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`

		// 内网访问子域名。
		InternalSubDomain *string `json:"InternalSubDomain,omitempty" name:"InternalSubDomain"`

		// 外网访问子域名。
		OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

		// 内网访问http服务端口号。
		InnerHttpPort *int64 `json:"InnerHttpPort,omitempty" name:"InnerHttpPort"`

		// 内网访问https端口号。
		InnerHttpsPort *int64 `json:"InnerHttpsPort,omitempty" name:"InnerHttpsPort"`

		// API总数。
		ApiTotalCount *int64 `json:"ApiTotalCount,omitempty" name:"ApiTotalCount"`

		// API列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApiIdStatusSet []*ApiIdStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet" list`

		// 使用计划总数量。
		UsagePlanTotalCount *int64 `json:"UsagePlanTotalCount,omitempty" name:"UsagePlanTotalCount"`

		// 使用计划数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UsagePlanList []*UsagePlan `json:"UsagePlanList,omitempty" name:"UsagePlanList" list`

		// IP版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
		IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

		// 此服务的用户类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserType *string `json:"UserType,omitempty" name:"UserType"`

		// 预留字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
		SetId *int64 `json:"SetId,omitempty" name:"SetId"`

		// 服务绑定的标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainMappingsRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务绑定的自定义域名。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

func (r *DescribeServiceSubDomainMappingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceSubDomainMappingsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainMappingsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 自定义路径映射列表。
		Result *ServiceSubDomainMappings `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceSubDomainMappingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceSubDomainMappingsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainsRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceSubDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceSubDomainsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceSubDomainsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询服务自定义域名列表。
		Result *DomainSets `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceSubDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceSubDomainsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务绑定使用计划列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ServiceUsagePlanSet `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServiceUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServicesStatusRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持ServiceId、ServiceName、NotUsagePlanId、Environment、IpVersion。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeServicesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServicesStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServicesStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务列表查询结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ServicesStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServicesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServicesStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanEnvironmentsRequest struct {
	*tchttp.BaseRequest

	// 待查询的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 定类型，取值为 API、SERVICE，默认值为 SERVICE。
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUsagePlanEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanEnvironmentsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 使用计划绑定详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *UsagePlanEnvironmentStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlanEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanEnvironmentsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 待查询的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
}

func (r *DescribeUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 使用计划详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanSecretIdsRequest struct {
	*tchttp.BaseRequest

	// 绑定的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeUsagePlanSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanSecretIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlanSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 使用计划绑定的密钥列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *UsagePlanBindSecretStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlanSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlanSecretIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlansStatusRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 使用计划过滤条件。支持UsagePlanId、UsagePlanName、NotServiceId、NotApiId、Environment。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeUsagePlansStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlansStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUsagePlansStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 使用计划列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *UsagePlansStatus `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUsagePlansStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUsagePlansStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableApiKeyRequest struct {
	*tchttp.BaseRequest

	// 待禁用的密钥 ID。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *DisableApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 禁用密钥操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DocumentSDK struct {

	// 生成的 document 会存放到 COS 中，此出参返回产生文件的下载链接。
	DocumentURL *string `json:"DocumentURL,omitempty" name:"DocumentURL"`

	// 生成的 SDK 会存放到 COS 中，此出参返回产生 SDK 文件的下载链接。
	SdkURL *string `json:"SdkURL,omitempty" name:"SdkURL"`
}

type DomainSetList struct {

	// 域名名称。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 域名解析状态。True 表示正常解析，False 表示解析失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 证书ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 是否使用默认路径映射。
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// 自定义域名协议类型。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 网络类型（'INNER' 或 'OUTER'）。
	NetType *string `json:"NetType,omitempty" name:"NetType"`
}

type DomainSets struct {

	// 服务下的自定义域名数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 自定义服务域名列表。
	DomainSet []*DomainSetList `json:"DomainSet,omitempty" name:"DomainSet" list`
}

type EnableApiKeyRequest struct {
	*tchttp.BaseRequest

	// 待启用的密钥 ID。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
}

func (r *EnableApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 启动密钥操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Environment struct {

	// 环境名称。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 访问路径。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 发布状态，1 表示已发布，0 表示未发布。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 运行版本。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`
}

type EnvironmentStrategy struct {

	// 环境名
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 限流值
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`
}

type ErrorCodes struct {

	// 自定义响应配置错误码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 自定义响应配置错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 自定义响应配置错误码备注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 自定义错误码转换。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConvertedCode *int64 `json:"ConvertedCode,omitempty" name:"ConvertedCode"`

	// 是否需要开启错误码转换。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeedConvert *bool `json:"NeedConvert,omitempty" name:"NeedConvert"`
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type GenerateApiDocumentRequest struct {
	*tchttp.BaseRequest

	// 待创建文档的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待创建 SDK 的服务所在环境。
	GenEnvironment *string `json:"GenEnvironment,omitempty" name:"GenEnvironment"`

	// 待创建 SDK 的语言。当前只支持 Python 和 JavaScript。
	GenLanguage *string `json:"GenLanguage,omitempty" name:"GenLanguage"`
}

func (r *GenerateApiDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GenerateApiDocumentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GenerateApiDocumentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// api文档&sdk链接。
		Result *DocumentSDK `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateApiDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GenerateApiDocumentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HealthCheckConf struct {

	// 是否开启健康检查。
	IsHealthCheck *bool `json:"IsHealthCheck,omitempty" name:"IsHealthCheck"`

	// 健康检查阈值。
	RequestVolumeThreshold *int64 `json:"RequestVolumeThreshold,omitempty" name:"RequestVolumeThreshold"`

	// 窗口大小。
	SleepWindowInMilliseconds *int64 `json:"SleepWindowInMilliseconds,omitempty" name:"SleepWindowInMilliseconds"`

	// 阈值百分比。
	ErrorThresholdPercentage *int64 `json:"ErrorThresholdPercentage,omitempty" name:"ErrorThresholdPercentage"`
}

type IPStrategy struct {

	// 策略唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 用户自定义策略名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 策略类型。支持WHITE（白名单）和BLACK（黑名单）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`

	// IP列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 修改时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 最后修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 策略绑定的API数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindApiTotalCount *int64 `json:"BindApiTotalCount,omitempty" name:"BindApiTotalCount"`

	// 绑定的API详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindApis []*DesApisStatus `json:"BindApis,omitempty" name:"BindApis" list`
}

type IPStrategyApi struct {

	// API 唯一 ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 用户自定义的 API 名称。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API 类型。取值为NORMAL（普通API）和TSF （微服务API）。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API 的路径。如 /path。
	Path *string `json:"Path,omitempty" name:"Path"`

	// API 的请求方法。如 GET。
	Method *string `json:"Method,omitempty" name:"Method"`

	// API 已经绑定的其他策略唯一ID。
	OtherIPStrategyId *string `json:"OtherIPStrategyId,omitempty" name:"OtherIPStrategyId"`

	// API 已经绑定的环境。
	OtherEnvironmentName *string `json:"OtherEnvironmentName,omitempty" name:"OtherEnvironmentName"`
}

type IPStrategyApiStatus struct {

	// 环境绑定API数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 环境绑定API详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiIdStatusSet []*IPStrategyApi `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet" list`
}

type IPStrategysStatus struct {

	// 策略数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 策略列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategySet []*IPStrategy `json:"StrategySet,omitempty" name:"StrategySet" list`
}

type LogQuery struct {

	// 检索字段
	Name *string `json:"Name,omitempty" name:"Name"`

	// 操作符
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 检索值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type MicroService struct {

	// 微服务集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 微服务命名空间ID。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务名称。
	MicroServiceName *string `json:"MicroServiceName,omitempty" name:"MicroServiceName"`
}

type MicroServiceReq struct {

	// 微服务集群。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 微服务命名空间。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务名称。
	MicroServiceName *string `json:"MicroServiceName,omitempty" name:"MicroServiceName"`
}

type ModifyApiEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 限流值。
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 环境名。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// API列表。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *ModifyApiEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiEnvironmentStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiEnvironmentStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiIncrementRequest struct {
	*tchttp.BaseRequest

	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 接口ID
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 需要修改的API auth类型(可选择OAUTH-授权API)
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// oauth接口需要修改的公钥值
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// oauth接口重定向地址
	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

func (r *ModifyApiIncrementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiIncrementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiIncrementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiIncrementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiIncrementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiRequest struct {
	*tchttp.BaseRequest

	// API 所在的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 的后端服务类型。支持HTTP、MOCK、TSF、CLB、SCF、WEBSOCKET、TARGET（内测）。
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 请求的前端配置。
	RequestConfig *RequestConfig `json:"RequestConfig,omitempty" name:"RequestConfig"`

	// API 接口唯一 ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 用户自定义的 API 名称。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// 用户自定义的 API 接口描述。
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// API 类型，支持NORMAL和TSF，默认为NORMAL。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API 鉴权类型。支持SECRET、NONE、OAUTH。默认为NONE。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 是否需要签名认证，True 表示需要，False 表示不需要。待废弃。
	AuthRequired *bool `json:"AuthRequired,omitempty" name:"AuthRequired"`

	// API 的后端服务超时时间，单位是秒。
	ServiceTimeout *int64 `json:"ServiceTimeout,omitempty" name:"ServiceTimeout"`

	// API 的前端请求类型，如 HTTP 或 HTTPS 或者 HTTP 和 HTTPS。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 是否需要开启跨域，Ture 表示需要，False 表示不需要。
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// 常量参数。
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters" list`

	// 前端请求参数。
	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters" list`

	// 当AuthType 为 OAUTH时，该字段有效， NORMAL：业务api   OAUTH：授权API。
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// API 的后端 Mock 返回信息。如果 ServiceType 是 Mock，则此参数必传。
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// API绑定微服务服务列表。
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices" list`

	// 微服务的负载均衡配置。
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// 微服务的健康检查配置。
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// target类型负载均衡配置。（内测阶段）
	TargetServicesLoadBalanceConf *int64 `json:"TargetServicesLoadBalanceConf,omitempty" name:"TargetServicesLoadBalanceConf"`

	// target健康检查配置。（内测阶段）
	TargetServicesHealthCheckConf *HealthCheckConf `json:"TargetServicesHealthCheckConf,omitempty" name:"TargetServicesHealthCheckConf"`

	// scf 函数名称。当后端类型是SCF时生效。
	ServiceScfFunctionName *string `json:"ServiceScfFunctionName,omitempty" name:"ServiceScfFunctionName"`

	// scf websocket注册函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketRegisterFunctionName *string `json:"ServiceWebsocketRegisterFunctionName,omitempty" name:"ServiceWebsocketRegisterFunctionName"`

	// scf websocket清理函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketCleanupFunctionName *string `json:"ServiceWebsocketCleanupFunctionName,omitempty" name:"ServiceWebsocketCleanupFunctionName"`

	// scf websocket传输函数。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketTransportFunctionName *string `json:"ServiceWebsocketTransportFunctionName,omitempty" name:"ServiceWebsocketTransportFunctionName"`

	// scf 函数命名空间。当后端类型是SCF时生效。
	ServiceScfFunctionNamespace *string `json:"ServiceScfFunctionNamespace,omitempty" name:"ServiceScfFunctionNamespace"`

	// scf函数版本。当后端类型是SCF时生效。
	ServiceScfFunctionQualifier *string `json:"ServiceScfFunctionQualifier,omitempty" name:"ServiceScfFunctionQualifier"`

	// scf websocket注册函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketRegisterFunctionNamespace *string `json:"ServiceWebsocketRegisterFunctionNamespace,omitempty" name:"ServiceWebsocketRegisterFunctionNamespace"`

	// scf websocket传输函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketRegisterFunctionQualifier *string `json:"ServiceWebsocketRegisterFunctionQualifier,omitempty" name:"ServiceWebsocketRegisterFunctionQualifier"`

	// scf websocket传输函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketTransportFunctionNamespace *string `json:"ServiceWebsocketTransportFunctionNamespace,omitempty" name:"ServiceWebsocketTransportFunctionNamespace"`

	// scf websocket传输函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketTransportFunctionQualifier *string `json:"ServiceWebsocketTransportFunctionQualifier,omitempty" name:"ServiceWebsocketTransportFunctionQualifier"`

	// scf websocket清理函数命名空间。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketCleanupFunctionNamespace *string `json:"ServiceWebsocketCleanupFunctionNamespace,omitempty" name:"ServiceWebsocketCleanupFunctionNamespace"`

	// scf websocket清理函数版本。当前端类型是WEBSOCKET且后端类型是SCF时生效。
	ServiceWebsocketCleanupFunctionQualifier *string `json:"ServiceWebsocketCleanupFunctionQualifier,omitempty" name:"ServiceWebsocketCleanupFunctionQualifier"`

	// 是否开启响应集成。当后端类型是SCF时生效。
	ServiceScfIsIntegratedResponse *bool `json:"ServiceScfIsIntegratedResponse,omitempty" name:"ServiceScfIsIntegratedResponse"`

	// 开始调试后计费。（云市场预留字段）
	IsDebugAfterCharge *bool `json:"IsDebugAfterCharge,omitempty" name:"IsDebugAfterCharge"`

	// 标签。
	TagSpecifications *Tag `json:"TagSpecifications,omitempty" name:"TagSpecifications"`

	// 是否删除自定义响应配置错误码，如果不传或者传 False，不删除，当传 True 时，则删除此 API 所有自定义响应配置错误码。
	IsDeleteResponseErrorCodes *bool `json:"IsDeleteResponseErrorCodes,omitempty" name:"IsDeleteResponseErrorCodes"`

	// 返回类型。
	ResponseType *string `json:"ResponseType,omitempty" name:"ResponseType"`

	// 自定义响应配置成功响应示例。
	ResponseSuccessExample *string `json:"ResponseSuccessExample,omitempty" name:"ResponseSuccessExample"`

	// 自定义响应配置失败响应示例。
	ResponseFailExample *string `json:"ResponseFailExample,omitempty" name:"ResponseFailExample"`

	// API 的后端服务配置。
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" name:"ServiceConfig"`

	// 关联的授权API 唯一 ID，当AuthType为OAUTH且ApiBusinessType为NORMAL时生效。标示业务API绑定的oauth2.0授权API唯一ID。
	AuthRelationApiId *string `json:"AuthRelationApiId,omitempty" name:"AuthRelationApiId"`

	// API的后端服务参数。
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters" list`

	// oauth配置。当AuthType是OAUTH时生效。
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// 用户自定义错误码配置。
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes" list`
}

func (r *ModifyApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyApiResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyApiResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 待修改的策略所属服务的唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待修改的策略唯一ID。
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 待修改的策略详情。
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

func (r *ModifyIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest

	// 服务的唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 限流值。
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 环境列表。
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames" list`
}

func (r *ModifyServiceEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceEnvironmentStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceEnvironmentStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceRequest struct {
	*tchttp.BaseRequest

	// 待修改服务的唯一 Id。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 修改后的服务名称。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 修改后的服务描述。
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// 修改后的服务前端请求类型，如 http、https和 http&https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 网络类型列表，用于指定支持的访问类型，INNER为内网访问，OUTER为外网访问。默认为OUTER。
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`
}

func (r *ModifyServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubDomainRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待修改路径映射的自定义的域名。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 是否修改为使用默认路径映射。为 true，表示使用默认路径映射，为 false，表示使用自定义路径映射。
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// 证书 ID，协议包含 HTTPS 的时候要传该字段。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 修改后的自定义域名协议类型。（http，https 或 http&https)
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 修改后的路径映射列表。
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet" list`

	// 网络类型 （'INNER' 或 'OUTER'）
	NetType *string `json:"NetType,omitempty" name:"NetType"`
}

func (r *ModifySubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改自定义域名操作是否成功。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUsagePlanRequest struct {
	*tchttp.BaseRequest

	// 使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 修改后的用户自定义的使用计划名称。
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 修改后的用户自定义的使用计划描述。
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 请求配额总数，取值范围为-1或者[1, 99999999]，默认为-1，表示不开启。
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 每秒请求限制数，取值范围为-1或者[1, 2000]，默认-1，表示不开启。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
}

func (r *ModifyUsagePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUsagePlanRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 使用计划详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUsagePlanResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OauthConfig struct {

	// 公钥，用于验证用户token。
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// token传递位置。
	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`

	// 重定向地址，用于引导用户登录操作。
	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

type PathMapping struct {

	// 路径。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 发布环境，可选值为“test”、 ”prepub“、”release“。
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type ReleaseService struct {

	// 发布时的备注信息填写。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// 发布的版本id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" name:"ReleaseVersion"`
}

type ReleaseServiceRequest struct {
	*tchttp.BaseRequest

	// 待发布服务的唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待发布的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 本次的发布描述。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// apiId列表，预留字段，默认全量api发布。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *ReleaseServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发布信息。
		Result *ReleaseService `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReqParameter struct {

	// API 的前端参数名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// API 的前端参数位置，如 header。目前支持 header、query、path。
	Position *string `json:"Position,omitempty" name:"Position"`

	// API 的前端参数类型，如 String、int。
	Type *string `json:"Type,omitempty" name:"Type"`

	// API 的前端参数默认值。
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// API 的前端参数是否必填，True：表示必填，False：表示可选。
	Required *bool `json:"Required,omitempty" name:"Required"`

	// API 的前端参数备注。
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type RequestConfig struct {

	// API 的路径，如 /path。
	Path *string `json:"Path,omitempty" name:"Path"`

	// API 的请求方法，如 GET。
	Method *string `json:"Method,omitempty" name:"Method"`
}

type RequestParameter struct {

	// 请求参数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 参数位置
	Position *string `json:"Position,omitempty" name:"Position"`

	// 参数类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 是否必须
	Required *bool `json:"Required,omitempty" name:"Required"`
}

type ResponseErrorCodeReq struct {

	// 自定义响应配置错误码。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 自定义响应配置错误信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 自定义响应配置错误码备注。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 自定义错误码转换。
	ConvertedCode *int64 `json:"ConvertedCode,omitempty" name:"ConvertedCode"`

	// 是否需要开启错误码转换。
	NeedConvert *bool `json:"NeedConvert,omitempty" name:"NeedConvert"`
}

type Service struct {

	// 内网访问https端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpsPort *int64 `json:"InnerHttpsPort,omitempty" name:"InnerHttpsPort"`

	// 用户自定义的服务描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// 服务的前端请求类型。如http、https 或者 http&https。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 最后修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 服务支持的网络类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes" list`

	// 独占集群名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// 服务唯一 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// IP版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// 已经发布的环境列表。如test、prepub、release。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments" list`

	// 用户自定义的服务名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 系统为该服务分配的外网域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 内网访问http端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpPort *uint64 `json:"InnerHttpPort,omitempty" name:"InnerHttpPort"`

	// 系统为该服务自动分配的内网域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerSubDomain *string `json:"InnerSubDomain,omitempty" name:"InnerSubDomain"`

	// 服务的计费状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeIsolateStatus *int64 `json:"TradeIsolateStatus,omitempty" name:"TradeIsolateStatus"`

	// 服务绑定的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

type ServiceConfig struct {

	// 后端类型。启用vpc时生效，目前支持的类型为clb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// vpc 的唯一ID。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// API 的后端服务url。如果ServiceType是HTTP，则此参数必传。
	Url *string `json:"Url,omitempty" name:"Url"`

	// API 的后端服务路径，如 /path。如果 ServiceType 是 HTTP，则此参数必传。前后端路径可不同。
	Path *string `json:"Path,omitempty" name:"Path"`

	// API的后端服务请求方法，如 GET。如果 ServiceType 是 HTTP，则此参数必传。前后端方法可不同。
	Method *string `json:"Method,omitempty" name:"Method"`
}

type ServiceEnvironmentSet struct {

	// 服务绑定环境总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 服务绑定环境列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentList []*Environment `json:"EnvironmentList,omitempty" name:"EnvironmentList" list`
}

type ServiceEnvironmentStrategy struct {

	// 环境名。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 访问服务对应环境的url。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 发布状态。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 发布的版本号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 限流值。
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`
}

type ServiceEnvironmentStrategyStatus struct {

	// 限流策略数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 限流策略列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentList []*ServiceEnvironmentStrategy `json:"EnvironmentList,omitempty" name:"EnvironmentList" list`
}

type ServiceParameter struct {

	// API的后端服务参数名称。只有ServiceType是HTTP才会用到此参数。前后端参数名称可不同。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// API 的后端服务参数位置，如 head。只有 ServiceType 是 HTTP 才会用到此参数。前后端参数位置可配置不同。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *string `json:"Position,omitempty" name:"Position"`

	// API 的后端服务参数对应的前端参数位置，如 head。只有 ServiceType 是 HTTP 才会用到此参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelevantRequestParameterPosition *string `json:"RelevantRequestParameterPosition,omitempty" name:"RelevantRequestParameterPosition"`

	// API 的后端服务参数对应的前端参数名称。只有 ServiceType 是 HTTP 才会用到此参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelevantRequestParameterName *string `json:"RelevantRequestParameterName,omitempty" name:"RelevantRequestParameterName"`

	// API 的后端服务参数默认值。只有 ServiceType 是 HTTP 才会用到此参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// API 的后端服务参数备注。只有 ServiceType 是 HTTP 才会用到此参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelevantRequestParameterDesc *string `json:"RelevantRequestParameterDesc,omitempty" name:"RelevantRequestParameterDesc"`

	// API 的后端服务参数类型。只有 ServiceType 是 HTTP 才会用到此参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelevantRequestParameterType *string `json:"RelevantRequestParameterType,omitempty" name:"RelevantRequestParameterType"`
}

type ServiceReleaseHistory struct {

	// 发布版本总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 历史版本列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionList []*ServiceReleaseHistoryInfo `json:"VersionList,omitempty" name:"VersionList" list`
}

type ServiceReleaseHistoryInfo struct {

	// 版本号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 版本描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`

	// 版本发布时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`
}

type ServiceReleaseVersion struct {

	// 发布版本总数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 发布版本列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionList []*ServiceReleaseHistoryInfo `json:"VersionList,omitempty" name:"VersionList" list`
}

type ServiceSubDomainMappings struct {

	// 是否使用默认路径映射，为 True 表示使用默认路径映射；为 False 的话，表示使用自定义路径映射，此时 PathMappingSet 不为空。
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// 自定义路径映射列表。
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet" list`
}

type ServiceUsagePlanSet struct {

	// 服务上绑定的使用计划总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 服务上绑定的使用计划列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceUsagePlanList []*ApiUsagePlan `json:"ServiceUsagePlanList,omitempty" name:"ServiceUsagePlanList" list`
}

type ServicesStatus struct {

	// 服务列表总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 服务列表详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceSet []*Service `json:"ServiceSet,omitempty" name:"ServiceSet" list`
}

type Tag struct {

	// 标签的 key。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 便签的 value。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TargetServicesReq struct {

	// vm ip
	VmIp *string `json:"VmIp,omitempty" name:"VmIp"`

	// vpc id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vm port
	VmPort *int64 `json:"VmPort,omitempty" name:"VmPort"`

	// cvm所在宿主机ip
	HostIp *string `json:"HostIp,omitempty" name:"HostIp"`

	// docker ip
	DockerIp *string `json:"DockerIp,omitempty" name:"DockerIp"`
}

type TsfLoadBalanceConfResp struct {

	// 是否开启负载均衡。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLoadBalance *bool `json:"IsLoadBalance,omitempty" name:"IsLoadBalance"`

	// 负载均衡方式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 是否开启会话保持。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionStickRequired *bool `json:"SessionStickRequired,omitempty" name:"SessionStickRequired"`

	// 会话保持超时时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionStickTimeout *int64 `json:"SessionStickTimeout,omitempty" name:"SessionStickTimeout"`
}

type UnBindEnvironmentRequest struct {
	*tchttp.BaseRequest

	// 绑定类型，取值为 API、SERVICE，默认值为 SERVICE。
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// 待绑定的使用计划唯一 ID 列表。
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds" list`

	// 待解绑的服务环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 待解绑的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 唯一 ID 数组，当 BindType=API 时，需要传入此参数。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *UnBindEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindEnvironmentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 解绑操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindEnvironmentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindIPStrategyRequest struct {
	*tchttp.BaseRequest

	// 待解绑的服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待解绑的IP策略唯一ID。
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 待解绑的环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 待解绑的 API 列表。
	UnBindApiIds []*string `json:"UnBindApiIds,omitempty" name:"UnBindApiIds" list`
}

func (r *UnBindIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindIPStrategyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 解绑操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindIPStrategyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindSecretIdsRequest struct {
	*tchttp.BaseRequest

	// 待解绑的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 待解绑的密钥 ID 数组。
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds" list`
}

func (r *UnBindSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindSecretIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 解绑操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindSecretIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindSubDomainRequest struct {
	*tchttp.BaseRequest

	// 服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待解绑的自定义的域名。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

func (r *UnBindSubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindSubDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnBindSubDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 解绑自定义域名操作是否成功。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindSubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindSubDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnReleaseServiceRequest struct {
	*tchttp.BaseRequest

	// 待下线服务的唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待下线的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 保留字段，待下线的API列表。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds" list`
}

func (r *UnReleaseServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnReleaseServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnReleaseServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 下线操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnReleaseServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnReleaseServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateApiKeyRequest struct {
	*tchttp.BaseRequest

	// 待更换的密钥 ID。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// 待更换的密钥 Key，更新自定义密钥时，该字段为必传。长度10 - 50字符，包括字母、数字、英文下划线。
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
}

func (r *UpdateApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 更换后的密钥详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ApiKey `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceRequest struct {
	*tchttp.BaseRequest

	// 待切换服务的唯一 Id。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待切换的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 切换的版本号。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 本次的切换描述。
	UpdateDesc *string `json:"UpdateDesc,omitempty" name:"UpdateDesc"`
}

func (r *UpdateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateServiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 切换版本操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateServiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UsagePlan struct {

	// 环境名称。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 使用计划唯一ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 使用计划名称。
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 使用计划描述。
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 使用计划qps，-1表示没有限制。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// 使用计划时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 使用计划修改时间。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type UsagePlanBindEnvironment struct {

	// 环境名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 服务唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

type UsagePlanBindSecret struct {

	// 密钥ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// 密钥名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 密钥状态，0表示已禁用，1表示启用中。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type UsagePlanBindSecretStatus struct {

	// 使用计划绑定密钥的数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 密钥详情列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKeyList []*UsagePlanBindSecret `json:"AccessKeyList,omitempty" name:"AccessKeyList" list`
}

type UsagePlanEnvironment struct {

	// 绑定的服务唯一 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 的唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API 的名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API 的路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// API 的方法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 已经绑定的环境名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 已经使用的配额。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InUseRequestNum *int64 `json:"InUseRequestNum,omitempty" name:"InUseRequestNum"`

	// 最大请求量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 每秒最大请求次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 最后修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 服务名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type UsagePlanEnvironmentStatus struct {

	// 使用计划绑定的服务的环境数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 使用计划已经绑定的各个服务的环境状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentList []*UsagePlanEnvironment `json:"EnvironmentList,omitempty" name:"EnvironmentList" list`
}

type UsagePlanInfo struct {

	// 使用计划唯一 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 使用计划名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 使用计划描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 初始化调用次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitQuota *int64 `json:"InitQuota,omitempty" name:"InitQuota"`

	// 每秒请求限制数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// 最大调用次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 是否隐藏。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsHide *int64 `json:"IsHide,omitempty" name:"IsHide"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 最后修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 绑定密钥的数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindSecretIdTotalCount *int64 `json:"BindSecretIdTotalCount,omitempty" name:"BindSecretIdTotalCount"`

	// 绑定密钥的详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindSecretIds []*string `json:"BindSecretIds,omitempty" name:"BindSecretIds" list`

	// 绑定环境数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindEnvironmentTotalCount *int64 `json:"BindEnvironmentTotalCount,omitempty" name:"BindEnvironmentTotalCount"`

	// 绑定环境详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindEnvironments []*UsagePlanBindEnvironment `json:"BindEnvironments,omitempty" name:"BindEnvironments" list`
}

type UsagePlanStatusInfo struct {

	// 使用计划唯一 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 用户自定义的使用计划名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 用户自定义的使用计划描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 每秒最大请求次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`

	// 请求配额总量，-1表示没有限制。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 最后修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type UsagePlansStatus struct {

	// 符合条件的使用计划数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 使用计划列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanStatusSet []*UsagePlanStatusInfo `json:"UsagePlanStatusSet,omitempty" name:"UsagePlanStatusSet" list`
}
