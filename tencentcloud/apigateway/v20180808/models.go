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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type APIDoc struct {
	// API文档ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`

	// API文档名称
	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`

	// API文档构建状态
	ApiDocStatus *string `json:"ApiDocStatus,omitempty" name:"ApiDocStatus"`
}

type APIDocInfo struct {
	// API文档ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`

	// API文档名称
	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`

	// API文档构建状态
	ApiDocStatus *string `json:"ApiDocStatus,omitempty" name:"ApiDocStatus"`

	// API文档API数量
	ApiCount *int64 `json:"ApiCount,omitempty" name:"ApiCount"`

	// API文档查看次数
	ViewCount *int64 `json:"ViewCount,omitempty" name:"ViewCount"`

	// API文档发布次数
	ReleaseCount *int64 `json:"ReleaseCount,omitempty" name:"ReleaseCount"`

	// API文档访问URI
	ApiDocUri *string `json:"ApiDocUri,omitempty" name:"ApiDocUri"`

	// API文档分享密码
	SharePassword *string `json:"SharePassword,omitempty" name:"SharePassword"`

	// API文档更新时间
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境信息
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 生成API文档的API ID
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`

	// 服务名称
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 生成API文档的API名称
	ApiNames []*string `json:"ApiNames,omitempty" name:"ApiNames"`
}

type APIDocs struct {
	// API文档数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API文档基本信息
	APIDocSet []*APIDoc `json:"APIDocSet,omitempty" name:"APIDocSet"`
}

type ApiAppApiInfo struct {
	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`

	// 应用ID
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// Api的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Api名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 授权绑定时间，按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizedTime *string `json:"AuthorizedTime,omitempty" name:"AuthorizedTime"`

	// Api所属地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`

	// 授权绑定的环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`
}

type ApiAppApiInfos struct {
	// 数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 应用绑定的Api信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiAppApiSet []*ApiAppApiInfo `json:"ApiAppApiSet,omitempty" name:"ApiAppApiSet"`
}

type ApiAppInfo struct {
	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`

	// 应用ID
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// 应用SECRET
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiAppSecret *string `json:"ApiAppSecret,omitempty" name:"ApiAppSecret"`

	// 应用描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`

	// 创建时间，按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 修改时间，按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 应用KEY
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiAppKey *string `json:"ApiAppKey,omitempty" name:"ApiAppKey"`
}

type ApiAppInfos struct {
	// 应用数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 应用信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiAppSet []*ApiAppInfo `json:"ApiAppSet,omitempty" name:"ApiAppSet"`
}

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
	EnvironmentStrategySet []*EnvironmentStrategy `json:"EnvironmentStrategySet,omitempty" name:"EnvironmentStrategySet"`
}

type ApiEnvironmentStrategyStataus struct {
	// API绑定的限流策略数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API绑定的限流策略列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiEnvironmentStrategySet []*ApiEnvironmentStrategy `json:"ApiEnvironmentStrategySet,omitempty" name:"ApiEnvironmentStrategySet"`
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
	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds"`

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
	ResponseErrorCodes []*ErrorCodes `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`

	// 前端请求参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`

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
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`

	// 常量参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`

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
	MicroServices []*MicroService `json:"MicroServices,omitempty" name:"MicroServices"`

	// 微服务信息详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroServicesInfo []*int64 `json:"MicroServicesInfo,omitempty" name:"MicroServicesInfo"`

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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// API已发布的环境信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environments []*string `json:"Environments,omitempty" name:"Environments"`

	// 是否开启Base64编码，只有后端为scf时才会生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`

	// 是否开启Base64编码的header触发，只有后端为scf时才会生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsBase64Trigger *bool `json:"IsBase64Trigger,omitempty" name:"IsBase64Trigger"`

	// Header触发规则，总规则数量不超过10。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Base64EncodedTriggerRules []*Base64EncodedTriggerRule `json:"Base64EncodedTriggerRules,omitempty" name:"Base64EncodedTriggerRules"`
}

type ApiInfoSummary struct {
	// 插件相关的API总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 插件相关的API信息。
	ApiSet []*AvailableApiInfo `json:"ApiSet,omitempty" name:"ApiSet"`
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
	ApiKeySet []*ApiKey `json:"ApiKeySet,omitempty" name:"ApiKeySet"`
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
	ApiUsagePlanList []*ApiUsagePlan `json:"ApiUsagePlanList,omitempty" name:"ApiUsagePlanList"`
}

type ApisStatus struct {
	// 符合条件的 API 接口数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// API 接口列表。
	ApiIdStatusSet []*DesApisStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet"`
}

// Predefined struct for user
type AttachPluginRequestParams struct {
	// 绑定的API网关插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 要操作的服务ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 要操作API的环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 要绑定的API列表。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type AttachPluginRequest struct {
	*tchttp.BaseRequest
	
	// 绑定的API网关插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 要操作的服务ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 要操作API的环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 要绑定的API列表。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *AttachPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachPluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachPluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachPluginResponseParams struct {
	// 绑定操作是否成功。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AttachPluginResponse struct {
	*tchttp.BaseResponse
	Response *AttachPluginResponseParams `json:"Response"`
}

func (r *AttachPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachedApiInfo struct {
	// API所在服务ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API所在服务名称。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// API所在服务描述信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// API ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API名称。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiDesc *string `json:"ApiDesc,omitempty" name:"ApiDesc"`

	// 插件绑定API的环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 插件和API绑定时间。
	AttachedTime *string `json:"AttachedTime,omitempty" name:"AttachedTime"`
}

type AttachedApiSummary struct {
	// 插件绑定的API数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 插件绑定的API信息。
	AttachedApis []*AttachedApiInfo `json:"AttachedApis,omitempty" name:"AttachedApis"`
}

type AttachedPluginInfo struct {
	// 插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 环境信息。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 绑定时间。
	AttachedTime *string `json:"AttachedTime,omitempty" name:"AttachedTime"`

	// 插件名称。
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 插件类型。
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// 插件描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 插件定义语句。
	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`
}

type AttachedPluginSummary struct {
	// 已绑定的插件总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 已绑定的插件信息。
	PluginSummary []*AttachedPluginInfo `json:"PluginSummary,omitempty" name:"PluginSummary"`
}

type AvailableApiInfo struct {
	// API ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// API名称。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// API类型。
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// API路径。
	Path *string `json:"Path,omitempty" name:"Path"`

	// API方法。
	Method *string `json:"Method,omitempty" name:"Method"`

	// API是否绑定其他插件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachedOtherPlugin *bool `json:"AttachedOtherPlugin,omitempty" name:"AttachedOtherPlugin"`

	// API是否绑定当前插件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAttached *bool `json:"IsAttached,omitempty" name:"IsAttached"`
}

type Base64EncodedTriggerRule struct {
	// 进行编码触发的header，可选值 "Accept"和"Content_Type" 对应实际数据流请求header中的Accept和 Content-Type。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 进行编码触发的header的可选值数组, 数组元素的字符串最大长度为40，元素可以包括数字，英文字母以及特殊字符，特殊字符的可选值为： `.`    `+`    `*`   `-`   `/`  `_` 
	// 
	// 例如 [
	//     "application/x-vpeg005",
	//     "application/xhtml+xml",
	//     "application/vnd.ms-project",
	//     "application/vnd.rn-rn_music_package"
	// ] 等都是合法的。
	Value []*string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type BindApiAppRequestParams struct {
	// 待绑定的应用唯一 ID 。
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// 待绑定的环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 待绑定的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待绑定的API唯一ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

type BindApiAppRequest struct {
	*tchttp.BaseRequest
	
	// 待绑定的应用唯一 ID 。
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// 待绑定的环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 待绑定的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待绑定的API唯一ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *BindApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	delete(f, "Environment")
	delete(f, "ServiceId")
	delete(f, "ApiId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindApiAppResponseParams struct {
	// 绑定操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindApiAppResponse struct {
	*tchttp.BaseResponse
	Response *BindApiAppResponseParams `json:"Response"`
}

func (r *BindApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindApiInfo struct {
	// api唯一id
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Service唯一id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// api名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`

	// 服务名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 绑定时间
	BindTime *string `json:"BindTime,omitempty" name:"BindTime"`
}

// Predefined struct for user
type BindEnvironmentRequestParams struct {
	// 待绑定的使用计划唯一 ID 列表。
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds"`

	// 绑定类型，取值为API、SERVICE，默认值为SERVICE。
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// 待绑定的环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 待绑定的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API唯一ID数组，当bindType=API时，需要传入此参数。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type BindEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 待绑定的使用计划唯一 ID 列表。
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds"`

	// 绑定类型，取值为API、SERVICE，默认值为SERVICE。
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// 待绑定的环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 待绑定的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API唯一ID数组，当bindType=API时，需要传入此参数。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *BindEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanIds")
	delete(f, "BindType")
	delete(f, "Environment")
	delete(f, "ServiceId")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindEnvironmentResponseParams struct {
	// 绑定操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *BindEnvironmentResponseParams `json:"Response"`
}

func (r *BindEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindIPStrategyRequestParams struct {
	// 待绑定的IP策略所属的服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待绑定的IP策略唯一ID。
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// IP策略待绑定的环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// IP策略待绑定的API列表。
	BindApiIds []*string `json:"BindApiIds,omitempty" name:"BindApiIds"`
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
	BindApiIds []*string `json:"BindApiIds,omitempty" name:"BindApiIds"`
}

func (r *BindIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindIPStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyId")
	delete(f, "EnvironmentName")
	delete(f, "BindApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindIPStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindIPStrategyResponseParams struct {
	// 绑定操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *BindIPStrategyResponseParams `json:"Response"`
}

func (r *BindIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSecretIdsRequestParams struct {
	// 待绑定的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 待绑定的密钥 ID 数组。
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds"`
}

type BindSecretIdsRequest struct {
	*tchttp.BaseRequest
	
	// 待绑定的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 待绑定的密钥 ID 数组。
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds"`
}

func (r *BindSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSecretIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	delete(f, "AccessKeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindSecretIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSecretIdsResponseParams struct {
	// 绑定操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *BindSecretIdsResponseParams `json:"Response"`
}

func (r *BindSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSecretIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSubDomainRequestParams struct {
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
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`

	// 是否将HTTP请求强制跳转 HTTPS，默认为false。参数为 true时，API网关会将所有使用该自定义域名的 HTTP 协议的请求重定向至 HTTPS 协议进行转发。
	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`
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
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`

	// 是否将HTTP请求强制跳转 HTTPS，默认为false。参数为 true时，API网关会将所有使用该自定义域名的 HTTP 协议的请求重定向至 HTTPS 协议进行转发。
	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`
}

func (r *BindSubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSubDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "SubDomain")
	delete(f, "Protocol")
	delete(f, "NetType")
	delete(f, "IsDefaultMapping")
	delete(f, "NetSubDomain")
	delete(f, "CertificateId")
	delete(f, "PathMappingSet")
	delete(f, "IsForcedHttps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindSubDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindSubDomainResponseParams struct {
	// 绑定操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindSubDomainResponse struct {
	*tchttp.BaseResponse
	Response *BindSubDomainResponseParams `json:"Response"`
}

func (r *BindSubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindSubDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BuildAPIDocRequestParams struct {
	// API文档ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

type BuildAPIDocRequest struct {
	*tchttp.BaseRequest
	
	// API文档ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *BuildAPIDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BuildAPIDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiDocId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BuildAPIDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BuildAPIDocResponseParams struct {
	// 操作是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BuildAPIDocResponse struct {
	*tchttp.BaseResponse
	Response *BuildAPIDocResponseParams `json:"Response"`
}

func (r *BuildAPIDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BuildAPIDocResponse) FromJsonString(s string) error {
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

type CosConfig struct {
	// API调用后端COS的方式，前端请求方法与Action的可选值为：
	// GET：GetObject
	// PUT：PutObject
	// POST：PostObject、AppendObject
	// HEAD： HeadObject
	// DELETE： DeleteObject。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitempty" name:"Action"`

	// API后端COS的存储桶名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// API调用后端COS的签名开关，默认为false。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Authorization *bool `json:"Authorization,omitempty" name:"Authorization"`

	// API后端COS的路径匹配模式，可选值：
	// BackEndPath ： 后端路径匹配
	// FullPath ： 全路径匹配
	// 
	// 默认值为：BackEndPath
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathMatchMode *string `json:"PathMatchMode,omitempty" name:"PathMatchMode"`
}

// Predefined struct for user
type CreateAPIDocRequestParams struct {
	// API文档名称
	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`

	// 服务名称
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名称
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 生成文档的API列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type CreateAPIDocRequest struct {
	*tchttp.BaseRequest
	
	// API文档名称
	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`

	// 服务名称
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名称
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 生成文档的API列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *CreateAPIDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAPIDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiDocName")
	delete(f, "ServiceId")
	delete(f, "Environment")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAPIDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAPIDocResponseParams struct {
	// API文档基本信息
	Result *APIDoc `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAPIDocResponse struct {
	*tchttp.BaseResponse
	Response *CreateAPIDocResponseParams `json:"Response"`
}

func (r *CreateAPIDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAPIDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiAppRequestParams struct {
	// 用户自定义应用名称。
	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`

	// 应用描述
	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`
}

type CreateApiAppRequest struct {
	*tchttp.BaseRequest
	
	// 用户自定义应用名称。
	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`

	// 应用描述
	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`
}

func (r *CreateApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppName")
	delete(f, "ApiAppDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiAppResponseParams struct {
	// 新增的应用详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApiAppInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApiAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateApiAppResponseParams `json:"Response"`
}

func (r *CreateApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiKeyRequestParams struct {
	// 用户自定义密钥名称。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 密钥类型，支持 auto 和 manual（自定义密钥），默认为 auto。
	AccessKeyType *string `json:"AccessKeyType,omitempty" name:"AccessKeyType"`

	// 用户自定义密钥 ID，AccessKeyType 为 manual 时必传。长度为5 - 50字符，由字母、数字、英文下划线组成。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// 用户自定义密钥 Key，AccessKeyType 为 manual 时必传。长度为10 - 50字符，由字母、数字、英文下划线。
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretName")
	delete(f, "AccessKeyType")
	delete(f, "AccessKeyId")
	delete(f, "AccessKeySecret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiKeyResponseParams struct {
	// 新增的密钥详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApiKey `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateApiKeyResponseParams `json:"Response"`
}

func (r *CreateApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiRequestParams struct {
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

	// API 鉴权类型。支持SECRET（密钥对鉴权）、NONE（免鉴权）、OAUTH、APP（应用认证）。默认为NONE。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 是否开启跨域。
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// 常量参数。
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`

	// 前端请求参数。
	RequestParameters []*RequestParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`

	// 当AuthType 为 OAUTH时，该字段有效， NORMAL：业务api OAUTH：授权API。
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// API 的后端 Mock 返回信息。如果 ServiceType 是 Mock，则此参数必传。
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// API绑定微服务服务列表。
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices"`

	// 微服务的负载均衡配置。
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// 微服务的健康检查配置。
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// target类型后端资源信息。（内测阶段）
	TargetServices []*TargetServicesReq `json:"TargetServices,omitempty" name:"TargetServices"`

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
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`

	// oauth配置。当AuthType是OAUTH时生效。
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// 用户自定义错误码配置。
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`

	// tsf serverless 命名空间ID。（内测中）
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" name:"TargetNamespaceId"`

	// 用户类型。
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// 是否打开Base64编码，只有后端是scf时才会生效。
	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`

	// 事件总线ID。
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// scf函数类型。当后端类型是SCF时生效。支持事件触发(EVENT)，http直通云函数(HTTP)。
	ServiceScfFunctionType *string `json:"ServiceScfFunctionType,omitempty" name:"ServiceScfFunctionType"`

	// EIAM应用类型。
	EIAMAppType *string `json:"EIAMAppType,omitempty" name:"EIAMAppType"`

	// EIAM应用认证类型，支持仅认证（AuthenticationOnly）、认证和鉴权（Authorization）。
	EIAMAuthType *string `json:"EIAMAuthType,omitempty" name:"EIAMAuthType"`

	// EIAM应用Token 有效时间，单位为秒，默认为7200秒。
	TokenTimeout *int64 `json:"TokenTimeout,omitempty" name:"TokenTimeout"`

	// EIAM应用ID。
	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`

	// 资源的Owner
	Owner *string `json:"Owner,omitempty" name:"Owner"`
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

	// API 鉴权类型。支持SECRET（密钥对鉴权）、NONE（免鉴权）、OAUTH、APP（应用认证）。默认为NONE。
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 是否开启跨域。
	EnableCORS *bool `json:"EnableCORS,omitempty" name:"EnableCORS"`

	// 常量参数。
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`

	// 前端请求参数。
	RequestParameters []*RequestParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`

	// 当AuthType 为 OAUTH时，该字段有效， NORMAL：业务api OAUTH：授权API。
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// API 的后端 Mock 返回信息。如果 ServiceType 是 Mock，则此参数必传。
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// API绑定微服务服务列表。
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices"`

	// 微服务的负载均衡配置。
	ServiceTsfLoadBalanceConf *TsfLoadBalanceConfResp `json:"ServiceTsfLoadBalanceConf,omitempty" name:"ServiceTsfLoadBalanceConf"`

	// 微服务的健康检查配置。
	ServiceTsfHealthCheckConf *HealthCheckConf `json:"ServiceTsfHealthCheckConf,omitempty" name:"ServiceTsfHealthCheckConf"`

	// target类型后端资源信息。（内测阶段）
	TargetServices []*TargetServicesReq `json:"TargetServices,omitempty" name:"TargetServices"`

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
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`

	// oauth配置。当AuthType是OAUTH时生效。
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// 用户自定义错误码配置。
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`

	// tsf serverless 命名空间ID。（内测中）
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" name:"TargetNamespaceId"`

	// 用户类型。
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// 是否打开Base64编码，只有后端是scf时才会生效。
	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`

	// 事件总线ID。
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// scf函数类型。当后端类型是SCF时生效。支持事件触发(EVENT)，http直通云函数(HTTP)。
	ServiceScfFunctionType *string `json:"ServiceScfFunctionType,omitempty" name:"ServiceScfFunctionType"`

	// EIAM应用类型。
	EIAMAppType *string `json:"EIAMAppType,omitempty" name:"EIAMAppType"`

	// EIAM应用认证类型，支持仅认证（AuthenticationOnly）、认证和鉴权（Authorization）。
	EIAMAuthType *string `json:"EIAMAuthType,omitempty" name:"EIAMAuthType"`

	// EIAM应用Token 有效时间，单位为秒，默认为7200秒。
	TokenTimeout *int64 `json:"TokenTimeout,omitempty" name:"TokenTimeout"`

	// EIAM应用ID。
	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`

	// 资源的Owner
	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

func (r *CreateApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ServiceType")
	delete(f, "ServiceTimeout")
	delete(f, "Protocol")
	delete(f, "RequestConfig")
	delete(f, "ApiName")
	delete(f, "ApiDesc")
	delete(f, "ApiType")
	delete(f, "AuthType")
	delete(f, "EnableCORS")
	delete(f, "ConstantParameters")
	delete(f, "RequestParameters")
	delete(f, "ApiBusinessType")
	delete(f, "ServiceMockReturnMessage")
	delete(f, "MicroServices")
	delete(f, "ServiceTsfLoadBalanceConf")
	delete(f, "ServiceTsfHealthCheckConf")
	delete(f, "TargetServices")
	delete(f, "TargetServicesLoadBalanceConf")
	delete(f, "TargetServicesHealthCheckConf")
	delete(f, "ServiceScfFunctionName")
	delete(f, "ServiceWebsocketRegisterFunctionName")
	delete(f, "ServiceWebsocketCleanupFunctionName")
	delete(f, "ServiceWebsocketTransportFunctionName")
	delete(f, "ServiceScfFunctionNamespace")
	delete(f, "ServiceScfFunctionQualifier")
	delete(f, "ServiceWebsocketRegisterFunctionNamespace")
	delete(f, "ServiceWebsocketRegisterFunctionQualifier")
	delete(f, "ServiceWebsocketTransportFunctionNamespace")
	delete(f, "ServiceWebsocketTransportFunctionQualifier")
	delete(f, "ServiceWebsocketCleanupFunctionNamespace")
	delete(f, "ServiceWebsocketCleanupFunctionQualifier")
	delete(f, "ServiceScfIsIntegratedResponse")
	delete(f, "IsDebugAfterCharge")
	delete(f, "IsDeleteResponseErrorCodes")
	delete(f, "ResponseType")
	delete(f, "ResponseSuccessExample")
	delete(f, "ResponseFailExample")
	delete(f, "ServiceConfig")
	delete(f, "AuthRelationApiId")
	delete(f, "ServiceParameters")
	delete(f, "OauthConfig")
	delete(f, "ResponseErrorCodes")
	delete(f, "TargetNamespaceId")
	delete(f, "UserType")
	delete(f, "IsBase64Encoded")
	delete(f, "EventBusId")
	delete(f, "ServiceScfFunctionType")
	delete(f, "EIAMAppType")
	delete(f, "EIAMAuthType")
	delete(f, "TokenTimeout")
	delete(f, "EIAMAppId")
	delete(f, "Owner")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiResponseParams struct {
	// api信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateApiRsp `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApiResponse struct {
	*tchttp.BaseResponse
	Response *CreateApiResponseParams `json:"Response"`
}

func (r *CreateApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApiRsp struct {
	// api id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 请求方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 导入状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// api name
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiName *string `json:"ApiName,omitempty" name:"ApiName"`
}

type CreateApiRspSet struct {
	// 个数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的数组
	ApiSet []*CreateApiRsp `json:"ApiSet,omitempty" name:"ApiSet"`
}

// Predefined struct for user
type CreateIPStrategyRequestParams struct {
	// 服务的唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 用户自定义的策略名称。
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 策略类型。支持WHITE（白名单）和BLACK（黑名单）。
	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`

	// 策略详情，多个ip 使用\n 分隔符分开。
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

type CreateIPStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 服务的唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 用户自定义的策略名称。
	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`

	// 策略类型。支持WHITE（白名单）和BLACK（黑名单）。
	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`

	// 策略详情，多个ip 使用\n 分隔符分开。
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
}

func (r *CreateIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIPStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyName")
	delete(f, "StrategyType")
	delete(f, "StrategyData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIPStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIPStrategyResponseParams struct {
	// 新建的IP策略详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *IPStrategy `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *CreateIPStrategyResponseParams `json:"Response"`
}

func (r *CreateIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePluginRequestParams struct {
	// 用户自定义的插件名称。最长50个字符，最短2个字符，支持 a-z,A-Z,0-9,_, 必须字母开头，字母或者数字结尾。
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 插件类型。目前支持IPControl, TrafficControl, Cors, CustomReq, CustomAuth，Routing，TrafficControlByParameter, CircuitBreaker, ProxyCache。
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// 插件定义语句，支持json。
	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`

	// 插件描述，限定200字以内。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreatePluginRequest struct {
	*tchttp.BaseRequest
	
	// 用户自定义的插件名称。最长50个字符，最短2个字符，支持 a-z,A-Z,0-9,_, 必须字母开头，字母或者数字结尾。
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 插件类型。目前支持IPControl, TrafficControl, Cors, CustomReq, CustomAuth，Routing，TrafficControlByParameter, CircuitBreaker, ProxyCache。
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// 插件定义语句，支持json。
	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`

	// 插件描述，限定200字以内。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreatePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginName")
	delete(f, "PluginType")
	delete(f, "PluginData")
	delete(f, "Description")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePluginResponseParams struct {
	// 新建的插件详情。
	Result *Plugin `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePluginResponse struct {
	*tchttp.BaseResponse
	Response *CreatePluginResponseParams `json:"Response"`
}

func (r *CreatePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceRequestParams struct {
	// 用户自定义的服务名称。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 服务的前端请求类型。如 http、https、http&https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 用户自定义的服务描述。
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// 独立集群名称，用于指定创建服务所在的独立集群。
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// 网络类型列表，用于指定支持的访问类型，INNER为内网访问，OUTER为外网访问。默认为OUTER。
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`

	// IP版本号，支持IPv4和IPv6，默认为IPv4。
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// 集群名称。保留字段，tsf serverlss类型使用。
	SetServerName *string `json:"SetServerName,omitempty" name:"SetServerName"`

	// 用户类型。保留类型，serverless用户使用。
	AppIdType *string `json:"AppIdType,omitempty" name:"AppIdType"`

	// 标签。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 独享实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// vpc属性
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
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
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`

	// IP版本号，支持IPv4和IPv6，默认为IPv4。
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// 集群名称。保留字段，tsf serverlss类型使用。
	SetServerName *string `json:"SetServerName,omitempty" name:"SetServerName"`

	// 用户类型。保留类型，serverless用户使用。
	AppIdType *string `json:"AppIdType,omitempty" name:"AppIdType"`

	// 标签。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 独享实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// vpc属性
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
}

func (r *CreateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceName")
	delete(f, "Protocol")
	delete(f, "ServiceDesc")
	delete(f, "ExclusiveSetName")
	delete(f, "NetTypes")
	delete(f, "IpVersion")
	delete(f, "SetServerName")
	delete(f, "AppIdType")
	delete(f, "Tags")
	delete(f, "InstanceId")
	delete(f, "UniqVpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceResponseParams struct {
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
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`

	// IP版本号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceResponseParams `json:"Response"`
}

func (r *CreateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUpstreamRequestParams struct {
	// 后端协议，取值范围：HTTP, HTTPS
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 负载均衡算法，取值范围：ROUND-ROBIN
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// VPC唯一ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 后端通道名字
	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`

	// 后端通道描述
	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`

	// 后端访问类型，取值范围：IP_PORT, K8S
	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// 请求重试次数，默认3次
	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`

	// 网关转发到后端的Host请求头
	UpstreamHost *string `json:"UpstreamHost,omitempty" name:"UpstreamHost"`

	// 后端节点
	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`

	// 标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 健康检查配置，目前只支持VPC通道
	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`

	// K8S容器服务的配置
	K8sService []*K8sService `json:"K8sService,omitempty" name:"K8sService"`
}

type CreateUpstreamRequest struct {
	*tchttp.BaseRequest
	
	// 后端协议，取值范围：HTTP, HTTPS
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 负载均衡算法，取值范围：ROUND-ROBIN
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// VPC唯一ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 后端通道名字
	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`

	// 后端通道描述
	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`

	// 后端访问类型，取值范围：IP_PORT, K8S
	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// 请求重试次数，默认3次
	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`

	// 网关转发到后端的Host请求头
	UpstreamHost *string `json:"UpstreamHost,omitempty" name:"UpstreamHost"`

	// 后端节点
	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`

	// 标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 健康检查配置，目前只支持VPC通道
	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`

	// K8S容器服务的配置
	K8sService []*K8sService `json:"K8sService,omitempty" name:"K8sService"`
}

func (r *CreateUpstreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUpstreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Scheme")
	delete(f, "Algorithm")
	delete(f, "UniqVpcId")
	delete(f, "UpstreamName")
	delete(f, "UpstreamDescription")
	delete(f, "UpstreamType")
	delete(f, "Retries")
	delete(f, "UpstreamHost")
	delete(f, "Nodes")
	delete(f, "Tags")
	delete(f, "HealthChecker")
	delete(f, "K8sService")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUpstreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUpstreamResponseParams struct {
	// 创建返回的唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUpstreamResponse struct {
	*tchttp.BaseResponse
	Response *CreateUpstreamResponseParams `json:"Response"`
}

func (r *CreateUpstreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUpstreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUsagePlanRequestParams struct {
	// 用户自定义的使用计划名称。
	UsagePlanName *string `json:"UsagePlanName,omitempty" name:"UsagePlanName"`

	// 用户自定义的使用计划描述。
	UsagePlanDesc *string `json:"UsagePlanDesc,omitempty" name:"UsagePlanDesc"`

	// 请求配额总数，取值范围为-1或者[1, 99999999]，默认为-1，表示不开启。
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 每秒请求限制数，取值范围为-1或者[1, 2000]，默认-1，表示不开启。
	MaxRequestNumPreSec *int64 `json:"MaxRequestNumPreSec,omitempty" name:"MaxRequestNumPreSec"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanName")
	delete(f, "UsagePlanDesc")
	delete(f, "MaxRequestNum")
	delete(f, "MaxRequestNumPreSec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUsagePlanResponseParams struct {
	// 使用计划详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateUsagePlanResponseParams `json:"Response"`
}

func (r *CreateUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAPIDocRequestParams struct {
	// API文档ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

type DeleteAPIDocRequest struct {
	*tchttp.BaseRequest
	
	// API文档ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *DeleteAPIDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAPIDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiDocId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAPIDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAPIDocResponseParams struct {
	// 操作是否成功
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAPIDocResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAPIDocResponseParams `json:"Response"`
}

func (r *DeleteAPIDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAPIDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiAppRequestParams struct {
	// 应用唯一 ID。
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
}

type DeleteApiAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用唯一 ID。
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
}

func (r *DeleteApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiAppResponseParams struct {
	// 删除操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApiAppResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApiAppResponseParams `json:"Response"`
}

func (r *DeleteApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiKeyRequestParams struct {
	// 待删除的密钥 ID。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiKeyResponseParams struct {
	// 删除操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApiKeyResponseParams `json:"Response"`
}

func (r *DeleteApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiRequestParams struct {
	// API 所在的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 接口唯一 ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ApiId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiResponseParams struct {
	// 删除操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApiResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApiResponseParams `json:"Response"`
}

func (r *DeleteApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIPStrategyRequestParams struct {
	// 待删除的IP策略所属的服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待删除的IP策略唯一ID。
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIPStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIPStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIPStrategyResponseParams struct {
	// 删除操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIPStrategyResponseParams `json:"Response"`
}

func (r *DeleteIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePluginRequestParams struct {
	// 要删除的API网关插件的ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
}

type DeletePluginRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的API网关插件的ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
}

func (r *DeletePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePluginResponseParams struct {
	// 删除操作是否成功。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePluginResponse struct {
	*tchttp.BaseResponse
	Response *DeletePluginResponseParams `json:"Response"`
}

func (r *DeletePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceRequestParams struct {
	// 待删除服务的唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 跳过删除前置条件校验（仅支持独享实例上的服务）
	SkipVerification *int64 `json:"SkipVerification,omitempty" name:"SkipVerification"`
}

type DeleteServiceRequest struct {
	*tchttp.BaseRequest
	
	// 待删除服务的唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 跳过删除前置条件校验（仅支持独享实例上的服务）
	SkipVerification *int64 `json:"SkipVerification,omitempty" name:"SkipVerification"`
}

func (r *DeleteServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "SkipVerification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceResponseParams struct {
	// 删除操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceResponseParams `json:"Response"`
}

func (r *DeleteServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceSubDomainMappingRequestParams struct {
	// 服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务绑定的自定义域名。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 待删除映射的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。
	Environment *string `json:"Environment,omitempty" name:"Environment"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceSubDomainMappingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "SubDomain")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceSubDomainMappingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceSubDomainMappingResponseParams struct {
	// 删除自定义域名的路径映射操作是否成功。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteServiceSubDomainMappingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceSubDomainMappingResponseParams `json:"Response"`
}

func (r *DeleteServiceSubDomainMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceSubDomainMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUpstreamRequestParams struct {
	// 待删除的后端通道ID
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`
}

type DeleteUpstreamRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的后端通道ID
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`
}

func (r *DeleteUpstreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUpstreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UpstreamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUpstreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUpstreamResponseParams struct {
	// 成功删除的后端通道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUpstreamResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUpstreamResponseParams `json:"Response"`
}

func (r *DeleteUpstreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUpstreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsagePlanRequestParams struct {
	// 待删除的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsagePlanResponseParams struct {
	// 删除操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUsagePlanResponseParams `json:"Response"`
}

func (r *DeleteUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DemoteServiceUsagePlanRequestParams struct {
	// 使用计划ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 待降级的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名称。
	Environment *string `json:"Environment,omitempty" name:"Environment"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DemoteServiceUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	delete(f, "ServiceId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DemoteServiceUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DemoteServiceUsagePlanResponseParams struct {
	// 降级操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DemoteServiceUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *DemoteServiceUsagePlanResponseParams `json:"Response"`
}

func (r *DemoteServiceUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// API 鉴权类型。取值为SECRET（密钥对鉴权）、NONE（免鉴权）、OAUTH、EIAM。
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
	RelationBuniessApiIds []*string `json:"RelationBuniessApiIds,omitempty" name:"RelationBuniessApiIds"`

	// API关联的标签信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// API 的路径，如 /path。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// API 的请求方法，如 GET。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`
}

// Predefined struct for user
type DescribeAPIDocDetailRequestParams struct {
	// API文档ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

type DescribeAPIDocDetailRequest struct {
	*tchttp.BaseRequest
	
	// API文档ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *DescribeAPIDocDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIDocDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiDocId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAPIDocDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPIDocDetailResponseParams struct {
	// API文档详细信息
	Result *APIDocInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAPIDocDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAPIDocDetailResponseParams `json:"Response"`
}

func (r *DescribeAPIDocDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIDocDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPIDocsRequestParams struct {
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeAPIDocsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAPIDocsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIDocsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAPIDocsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPIDocsResponseParams struct {
	// API文档列表信息
	Result *APIDocs `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAPIDocsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAPIDocsResponseParams `json:"Response"`
}

func (r *DescribeAPIDocsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIDocsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllPluginApisRequestParams struct {
	// 要查询的服务ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 要查询的插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 环境信息。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeAllPluginApisRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的服务ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 要查询的插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 环境信息。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAllPluginApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllPluginApisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "PluginId")
	delete(f, "EnvironmentName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllPluginApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllPluginApisResponseParams struct {
	// 插件相关API的列表。
	Result *ApiInfoSummary `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAllPluginApisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllPluginApisResponseParams `json:"Response"`
}

func (r *DescribeAllPluginApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllPluginApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiAppBindApisStatusRequestParams struct {
	// 应用ID
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持ApiId、ApiName、ServiceId、Environment 、KeyWord（ 可以匹配name或者ID）。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeApiAppBindApisStatusRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持ApiId、ApiName、ServiceId、Environment 、KeyWord（ 可以匹配name或者ID）。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApiAppBindApisStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiAppBindApisStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiAppBindApisStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiAppBindApisStatusResponseParams struct {
	// 应用绑定的Api列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApiAppApiInfos `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiAppBindApisStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiAppBindApisStatusResponseParams `json:"Response"`
}

func (r *DescribeApiAppBindApisStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiAppBindApisStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiAppRequestParams struct {
	// 应用ID。
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
}

type DescribeApiAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID。
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`
}

func (r *DescribeApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiAppResponseParams struct {
	// 应用详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApiAppInfos `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiAppResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiAppResponseParams `json:"Response"`
}

func (r *DescribeApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiAppsStatusRequestParams struct {
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持ApiAppId、ApiAppName、KeyWord（ 可以匹配name或者ID）。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeApiAppsStatusRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持ApiAppId、ApiAppName、KeyWord（ 可以匹配name或者ID）。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApiAppsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiAppsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiAppsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiAppsStatusResponseParams struct {
	// 应用列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApiAppInfos `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiAppsStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiAppsStatusResponseParams `json:"Response"`
}

func (r *DescribeApiAppsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiAppsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiBindApiAppsStatusRequestParams struct {
	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Api的ID的数组
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持ApiAppId、Environment、KeyWord（ 可以匹配name或者ID）。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeApiBindApiAppsStatusRequest struct {
	*tchttp.BaseRequest
	
	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// Api的ID的数组
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持ApiAppId、Environment、KeyWord（ 可以匹配name或者ID）。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApiBindApiAppsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiBindApiAppsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ApiIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiBindApiAppsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiBindApiAppsStatusResponseParams struct {
	// 应用绑定的Api列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApiAppApiInfos `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiBindApiAppsStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiBindApiAppsStatusResponseParams `json:"Response"`
}

func (r *DescribeApiBindApiAppsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiBindApiAppsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiEnvironmentStrategyRequestParams struct {
	// API所属服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境列表。
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames"`

	// API唯一ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeApiEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest
	
	// API所属服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境列表。
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames"`

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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiEnvironmentStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "EnvironmentNames")
	delete(f, "ApiId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiEnvironmentStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiEnvironmentStrategyResponseParams struct {
	// api绑定策略详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApiEnvironmentStrategyStataus `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiEnvironmentStrategyResponseParams `json:"Response"`
}

func (r *DescribeApiEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiEnvironmentStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiForApiAppRequestParams struct {
	// API 所在的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 接口唯一 ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Api所属地域
	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`
}

type DescribeApiForApiAppRequest struct {
	*tchttp.BaseRequest
	
	// API 所在的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 接口唯一 ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// Api所属地域
	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`
}

func (r *DescribeApiForApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiForApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ApiId")
	delete(f, "ApiRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiForApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiForApiAppResponseParams struct {
	// API 详情。
	Result *ApiInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiForApiAppResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiForApiAppResponseParams `json:"Response"`
}

func (r *DescribeApiForApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiForApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeyRequestParams struct {
	// API 密钥 ID。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeyResponseParams struct {
	// 密钥详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApiKey `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiKeyResponseParams `json:"Response"`
}

func (r *DescribeApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeysStatusRequestParams struct {
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持AccessKeyId、AccessKeySecret、SecretName、NotUsagePlanId、Status、KeyWord（ 可以匹配name或者path）。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeApiKeysStatusRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持AccessKeyId、AccessKeySecret、SecretName、NotUsagePlanId、Status、KeyWord（ 可以匹配name或者path）。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApiKeysStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiKeysStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiKeysStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeysStatusResponseParams struct {
	// 密钥列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApiKeysStatus `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiKeysStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiKeysStatusResponseParams `json:"Response"`
}

func (r *DescribeApiKeysStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiKeysStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiRequestParams struct {
	// API 所在的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 接口唯一 ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ApiId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiResponseParams struct {
	// API 详情。
	Result *ApiInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiResponseParams `json:"Response"`
}

func (r *DescribeApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiUsagePlanRequestParams struct {
	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiUsagePlanResponseParams struct {
	// api绑定使用计划列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApiUsagePlanSet `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApiUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiUsagePlanResponseParams `json:"Response"`
}

func (r *DescribeApiUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApisStatusRequestParams struct {
	// API 所在的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// API过滤条件。支持ApiId、ApiName、ApiPath、ApiType、AuthRelationApiId、AuthType、ApiBuniessType、NotUsagePlanId、 Environment、Tags (values为 $tag_key:tag_value的列表)、TagKeys （values 为 tag key的列表），其中NotUsagePlanId和Environment必须同时使用，不能单独使用一个。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeApisStatusRequest struct {
	*tchttp.BaseRequest
	
	// API 所在的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// API过滤条件。支持ApiId、ApiName、ApiPath、ApiType、AuthRelationApiId、AuthType、ApiBuniessType、NotUsagePlanId、 Environment、Tags (values为 $tag_key:tag_value的列表)、TagKeys （values 为 tag key的列表），其中NotUsagePlanId和Environment必须同时使用，不能单独使用一个。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeApisStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApisStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApisStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApisStatusResponseParams struct {
	// API 详情列表。
	Result *ApisStatus `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApisStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApisStatusResponseParams `json:"Response"`
}

func (r *DescribeApisStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApisStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExclusiveInstanceDetailRequestParams struct {
	// 独享实例唯一id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeExclusiveInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// 独享实例唯一id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeExclusiveInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExclusiveInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExclusiveInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExclusiveInstanceDetailResponseParams struct {
	// 独享实例详情
	Result *InstanceDetail `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExclusiveInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExclusiveInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeExclusiveInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExclusiveInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExclusiveInstancesRequestParams struct {
	// 分页查询，limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页查询，offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeExclusiveInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询，limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页查询，offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeExclusiveInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExclusiveInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExclusiveInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExclusiveInstancesResponseParams struct {
	// 独享实例列表查询结果
	Result *InstanceInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExclusiveInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExclusiveInstancesResponseParams `json:"Response"`
}

func (r *DescribeExclusiveInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExclusiveInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExclusiveInstancesStatusRequestParams struct {
	// 分页查询，limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页查询，offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeExclusiveInstancesStatusRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询，limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页查询，offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeExclusiveInstancesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExclusiveInstancesStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExclusiveInstancesStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExclusiveInstancesStatusResponseParams struct {
	// 独享实例列表查询结果
	Result *InstanceSummary `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExclusiveInstancesStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExclusiveInstancesStatusResponseParams `json:"Response"`
}

func (r *DescribeExclusiveInstancesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExclusiveInstancesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStrategyApisStatusRequestParams struct {
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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIPStrategyApisStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStrategyApisStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyId")
	delete(f, "EnvironmentName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPStrategyApisStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStrategyApisStatusResponseParams struct {
	// 环境绑定API列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *IPStrategyApiStatus `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIPStrategyApisStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPStrategyApisStatusResponseParams `json:"Response"`
}

func (r *DescribeIPStrategyApisStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStrategyApisStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStrategyRequestParams struct {
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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyId")
	delete(f, "EnvironmentName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStrategyResponseParams struct {
	// IP策略详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *IPStrategy `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPStrategyResponseParams `json:"Response"`
}

func (r *DescribeIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStrategysStatusRequestParams struct {
	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 过滤条件。支持StrategyName。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeIPStrategysStatusRequest struct {
	*tchttp.BaseRequest
	
	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 过滤条件。支持StrategyName。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeIPStrategysStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStrategysStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPStrategysStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStrategysStatusResponseParams struct {
	// 符合条件的策略列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *IPStrategysStatus `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIPStrategysStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPStrategysStatusResponseParams `json:"Response"`
}

func (r *DescribeIPStrategysStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStrategysStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogSearchRequestParams struct {
	// 日志开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务id
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 保留字段
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

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
	LogQuerys []*LogQuery `json:"LogQuerys,omitempty" name:"LogQuerys"`
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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

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
	LogQuerys []*LogQuery `json:"LogQuerys,omitempty" name:"LogQuerys"`
}

func (r *DescribeLogSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogSearchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ServiceId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "ConText")
	delete(f, "Sort")
	delete(f, "Query")
	delete(f, "LogQuerys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogSearchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogSearchResponseParams struct {
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
	LogSet []*string `json:"LogSet,omitempty" name:"LogSet"`

	// 单次搜索返回的日志条数，TotalCount <= Limit
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogSearchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogSearchResponseParams `json:"Response"`
}

func (r *DescribeLogSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogSearchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginApisRequestParams struct {
	// 查询的插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribePluginApisRequest struct {
	*tchttp.BaseRequest
	
	// 查询的插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePluginApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginApisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePluginApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginApisResponseParams struct {
	// 插件绑定的API列表信息。
	Result *AttachedApiSummary `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePluginApisResponse struct {
	*tchttp.BaseResponse
	Response *DescribePluginApisResponseParams `json:"Response"`
}

func (r *DescribePluginApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginRequestParams struct {
	// 要查询的插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribePluginRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginResponseParams struct {
	// 插件详情。
	Result *Plugin `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePluginResponse struct {
	*tchttp.BaseResponse
	Response *DescribePluginResponseParams `json:"Response"`
}

func (r *DescribePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginsByApiRequestParams struct {
	// 要查询的API ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 要查询的服务ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境信息。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribePluginsByApiRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的API ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`

	// 要查询的服务ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境信息。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePluginsByApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginsByApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiId")
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePluginsByApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginsByApiResponseParams struct {
	// 插件可绑定的API列表信息。
	Result *AttachedPluginSummary `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePluginsByApiResponse struct {
	*tchttp.BaseResponse
	Response *DescribePluginsByApiResponseParams `json:"Response"`
}

func (r *DescribePluginsByApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginsByApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginsRequestParams struct {
	// 要查询的插件列表。
	PluginIds []*string `json:"PluginIds,omitempty" name:"PluginIds"`

	// 要查询的插件名称。
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 要查询的插件类型。
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。预留字段，目前不支持过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePluginsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的插件列表。
	PluginIds []*string `json:"PluginIds,omitempty" name:"PluginIds"`

	// 要查询的插件名称。
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 要查询的插件类型。
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。预留字段，目前不支持过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePluginsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginIds")
	delete(f, "PluginName")
	delete(f, "PluginType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePluginsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginsResponseParams struct {
	// 插件详情。
	Result *PluginSummary `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePluginsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePluginsResponseParams `json:"Response"`
}

func (r *DescribePluginsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceEnvironmentListRequestParams struct {
	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceEnvironmentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceEnvironmentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceEnvironmentListResponseParams struct {
	// 服务绑定环境详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ServiceEnvironmentSet `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceEnvironmentListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceEnvironmentListResponseParams `json:"Response"`
}

func (r *DescribeServiceEnvironmentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceEnvironmentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceEnvironmentReleaseHistoryRequestParams struct {
	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名称。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceEnvironmentReleaseHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceEnvironmentReleaseHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceEnvironmentReleaseHistoryResponseParams struct {
	// 服务发布历史。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ServiceReleaseHistory `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceEnvironmentReleaseHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceEnvironmentReleaseHistoryResponseParams `json:"Response"`
}

func (r *DescribeServiceEnvironmentReleaseHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceEnvironmentReleaseHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceEnvironmentStrategyRequestParams struct {
	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceEnvironmentStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceEnvironmentStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceEnvironmentStrategyResponseParams struct {
	// 限流策略列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ServiceEnvironmentStrategyStatus `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceEnvironmentStrategyResponseParams `json:"Response"`
}

func (r *DescribeServiceEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceEnvironmentStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceForApiAppRequestParams struct {
	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务所属的地域
	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`
}

type DescribeServiceForApiAppRequest struct {
	*tchttp.BaseRequest
	
	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务所属的地域
	ApiRegion *string `json:"ApiRegion,omitempty" name:"ApiRegion"`
}

func (r *DescribeServiceForApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceForApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ApiRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceForApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceForApiAppResponseParams struct {
	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务 环境列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments"`

	// 服务名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 服务描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// 服务支持协议，可选值为http、https、http&https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 服务创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 服务修改时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 独立集群名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExclusiveSetName *string `json:"ExclusiveSetName,omitempty" name:"ExclusiveSetName"`

	// 网络类型列表，INNER为内网访问，OUTER为外网访问。
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`

	// 内网访问子域名。
	InternalSubDomain *string `json:"InternalSubDomain,omitempty" name:"InternalSubDomain"`

	// 外网访问子域名。
	OuterSubDomain *string `json:"OuterSubDomain,omitempty" name:"OuterSubDomain"`

	// 内网访问http服务端口号。
	InnerHttpPort *int64 `json:"InnerHttpPort,omitempty" name:"InnerHttpPort"`

	// 内网访问https端口号。
	InnerHttpsPort *int64 `json:"InnerHttpsPort,omitempty" name:"InnerHttpsPort"`

	// API总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiTotalCount *int64 `json:"ApiTotalCount,omitempty" name:"ApiTotalCount"`

	// API列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiIdStatusSet []*ApiIdStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet"`

	// 使用计划总数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanTotalCount *int64 `json:"UsagePlanTotalCount,omitempty" name:"UsagePlanTotalCount"`

	// 使用计划数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanList []*UsagePlan `json:"UsagePlanList,omitempty" name:"UsagePlanList"`

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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceForApiAppResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceForApiAppResponseParams `json:"Response"`
}

func (r *DescribeServiceForApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceForApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceReleaseVersionRequestParams struct {
	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceReleaseVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceReleaseVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceReleaseVersionResponseParams struct {
	// 服务发布版本列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ServiceReleaseVersion `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceReleaseVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceReleaseVersionResponseParams `json:"Response"`
}

func (r *DescribeServiceReleaseVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceReleaseVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceRequestParams struct {
	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceResponseParams struct {
	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务 环境列表。
	AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments"`

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
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`

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
	ApiIdStatusSet []*ApiIdStatus `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet"`

	// 使用计划总数量。
	UsagePlanTotalCount *int64 `json:"UsagePlanTotalCount,omitempty" name:"UsagePlanTotalCount"`

	// 使用计划数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsagePlanList []*UsagePlan `json:"UsagePlanList,omitempty" name:"UsagePlanList"`

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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 独享实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 独享实例name
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetType *string `json:"SetType,omitempty" name:"SetType"`

	// 服务部署的集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeploymentType *string `json:"DeploymentType,omitempty" name:"DeploymentType"`

	// 特殊用途, NULL和DEFAULT表示无特殊用途，其他用途如HTTP_DNS等
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialUse *string `json:"SpecialUse,omitempty" name:"SpecialUse"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceResponseParams `json:"Response"`
}

func (r *DescribeServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceSubDomainMappingsRequestParams struct {
	// 服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 服务绑定的自定义域名。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceSubDomainMappingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "SubDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceSubDomainMappingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceSubDomainMappingsResponseParams struct {
	// 自定义路径映射列表。
	Result *ServiceSubDomainMappings `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceSubDomainMappingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceSubDomainMappingsResponseParams `json:"Response"`
}

func (r *DescribeServiceSubDomainMappingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceSubDomainMappingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceSubDomainsRequestParams struct {
	// 服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceSubDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceSubDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceSubDomainsResponseParams struct {
	// 查询服务自定义域名列表。
	Result *DomainSets `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceSubDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceSubDomainsResponseParams `json:"Response"`
}

func (r *DescribeServiceSubDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceSubDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceUsagePlanRequestParams struct {
	// 待查询的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceUsagePlanResponseParams struct {
	// 服务绑定使用计划列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ServiceUsagePlanSet `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceUsagePlanResponseParams `json:"Response"`
}

func (r *DescribeServiceUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServicesStatusRequestParams struct {
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持ServiceId、ServiceName、NotUsagePlanId、Environment、IpVersion、InstanceId、NetType、EIAMAppId。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeServicesStatusRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件。支持ServiceId、ServiceName、NotUsagePlanId、Environment、IpVersion、InstanceId、NetType、EIAMAppId。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeServicesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServicesStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServicesStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServicesStatusResponseParams struct {
	// 服务列表查询结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ServicesStatus `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServicesStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServicesStatusResponseParams `json:"Response"`
}

func (r *DescribeServicesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServicesStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpstreamBindApis struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 绑定的API信息
	BindApiSet []*BindApiInfo `json:"BindApiSet,omitempty" name:"BindApiSet"`
}

// Predefined struct for user
type DescribeUpstreamBindApisRequestParams struct {
	// 分页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 后端通道ID
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// ServiceId和ApiId过滤查询
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeUpstreamBindApisRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 后端通道ID
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// ServiceId和ApiId过滤查询
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUpstreamBindApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpstreamBindApisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "UpstreamId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUpstreamBindApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpstreamBindApisResponseParams struct {
	// 查询结果
	Result *DescribeUpstreamBindApis `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUpstreamBindApisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUpstreamBindApisResponseParams `json:"Response"`
}

func (r *DescribeUpstreamBindApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpstreamBindApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpstreamInfo struct {
	// 查询总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 查询列表
	UpstreamSet []*UpstreamInfo `json:"UpstreamSet,omitempty" name:"UpstreamSet"`
}

// Predefined struct for user
type DescribeUpstreamsRequestParams struct {
	// 分页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，支持后端通道ID（UpstreamId）、后端通道名字（UpstreamName）过滤查询
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeUpstreamsRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，支持后端通道ID（UpstreamId）、后端通道名字（UpstreamName）过滤查询
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUpstreamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpstreamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUpstreamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpstreamsResponseParams struct {
	// 查询结果
	Result *DescribeUpstreamInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUpstreamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUpstreamsResponseParams `json:"Response"`
}

func (r *DescribeUpstreamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpstreamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlanEnvironmentsRequestParams struct {
	// 待查询的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 定类型，取值为 API、SERVICE，默认值为 SERVICE。
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlanEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	delete(f, "BindType")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsagePlanEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlanEnvironmentsResponseParams struct {
	// 使用计划绑定详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *UsagePlanEnvironmentStatus `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsagePlanEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsagePlanEnvironmentsResponseParams `json:"Response"`
}

func (r *DescribeUsagePlanEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlanEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlanRequestParams struct {
	// 待查询的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlanResponseParams struct {
	// 使用计划详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsagePlanResponseParams `json:"Response"`
}

func (r *DescribeUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlanSecretIdsRequestParams struct {
	// 绑定的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlanSecretIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsagePlanSecretIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlanSecretIdsResponseParams struct {
	// 使用计划绑定的密钥列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *UsagePlanBindSecretStatus `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsagePlanSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsagePlanSecretIdsResponseParams `json:"Response"`
}

func (r *DescribeUsagePlanSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlanSecretIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlansStatusRequestParams struct {
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 使用计划过滤条件。支持UsagePlanId、UsagePlanName、NotServiceId、NotApiId、Environment。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeUsagePlansStatusRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 使用计划过滤条件。支持UsagePlanId、UsagePlanName、NotServiceId、NotApiId、Environment。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeUsagePlansStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlansStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsagePlansStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsagePlansStatusResponseParams struct {
	// 使用计划列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *UsagePlansStatus `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsagePlansStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsagePlansStatusResponseParams `json:"Response"`
}

func (r *DescribeUsagePlansStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsagePlansStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachPluginRequestParams struct {
	// 要解绑的API网关插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 要操作的服务ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 要操作API的环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 要解绑的API ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

type DetachPluginRequest struct {
	*tchttp.BaseRequest
	
	// 要解绑的API网关插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 要操作的服务ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 要操作API的环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 要解绑的API ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *DetachPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachPluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "ApiId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachPluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachPluginResponseParams struct {
	// 解绑操作是否成功。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetachPluginResponse struct {
	*tchttp.BaseResponse
	Response *DetachPluginResponseParams `json:"Response"`
}

func (r *DetachPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableApiKeyRequestParams struct {
	// 待禁用的密钥 ID。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableApiKeyResponseParams struct {
	// 禁用密钥操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *DisableApiKeyResponseParams `json:"Response"`
}

func (r *DisableApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 域名解析状态。1 表示正常解析，0 表示解析失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 证书ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 是否使用默认路径映射。
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// 自定义域名协议类型。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 网络类型（'INNER' 或 'OUTER'）。
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 是否将HTTP请求强制跳转 HTTPS，默认为false。参数为 true时，API网关会将所有使用该自定义域名的 HTTP 协议的请求重定向至 HTTPS 协议进行转发。
	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`

	// 域名备案注册状态
	RegistrationStatus *bool `json:"RegistrationStatus,omitempty" name:"RegistrationStatus"`
}

type DomainSets struct {
	// 服务下的自定义域名数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 自定义服务域名列表。
	DomainSet []*DomainSetList `json:"DomainSet,omitempty" name:"DomainSet"`
}

// Predefined struct for user
type EnableApiKeyRequestParams struct {
	// 待启用的密钥 ID。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableApiKeyResponseParams struct {
	// 启动密钥操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *EnableApiKeyResponseParams `json:"Response"`
}

func (r *EnableApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

	// 限流最大值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxQuota *int64 `json:"MaxQuota,omitempty" name:"MaxQuota"`
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
	Values []*string `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type GenerateApiDocumentRequestParams struct {
	// 待创建文档的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待创建 SDK 的服务所在环境。
	GenEnvironment *string `json:"GenEnvironment,omitempty" name:"GenEnvironment"`

	// 待创建 SDK 的语言。当前只支持 Python 和 JavaScript。
	GenLanguage *string `json:"GenLanguage,omitempty" name:"GenLanguage"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateApiDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "GenEnvironment")
	delete(f, "GenLanguage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateApiDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateApiDocumentResponseParams struct {
	// api文档&sdk链接。
	Result *DocumentSDK `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateApiDocumentResponse struct {
	*tchttp.BaseResponse
	Response *GenerateApiDocumentResponseParams `json:"Response"`
}

func (r *GenerateApiDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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
	BindApis []*DesApisStatus `json:"BindApis,omitempty" name:"BindApis"`
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
	ApiIdStatusSet []*IPStrategyApi `json:"ApiIdStatusSet,omitempty" name:"ApiIdStatusSet"`
}

type IPStrategysStatus struct {
	// 策略数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 策略列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategySet []*IPStrategy `json:"StrategySet,omitempty" name:"StrategySet"`
}

// Predefined struct for user
type ImportOpenApiRequestParams struct {
	// API所在的服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// openAPI正文内容。
	Content *string `json:"Content,omitempty" name:"Content"`

	// Content格式，只能是YAML或者JSON，默认是YAML。
	EncodeType *string `json:"EncodeType,omitempty" name:"EncodeType"`

	// Content版本，默认是openAPI，目前只支持openAPI。
	ContentVersion *string `json:"ContentVersion,omitempty" name:"ContentVersion"`
}

type ImportOpenApiRequest struct {
	*tchttp.BaseRequest
	
	// API所在的服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// openAPI正文内容。
	Content *string `json:"Content,omitempty" name:"Content"`

	// Content格式，只能是YAML或者JSON，默认是YAML。
	EncodeType *string `json:"EncodeType,omitempty" name:"EncodeType"`

	// Content版本，默认是openAPI，目前只支持openAPI。
	ContentVersion *string `json:"ContentVersion,omitempty" name:"ContentVersion"`
}

func (r *ImportOpenApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportOpenApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Content")
	delete(f, "EncodeType")
	delete(f, "ContentVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportOpenApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportOpenApiResponseParams struct {
	// 导入OpenApi返回参数。
	Result *CreateApiRspSet `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportOpenApiResponse struct {
	*tchttp.BaseResponse
	Response *ImportOpenApiResponseParams `json:"Response"`
}

func (r *ImportOpenApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportOpenApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {
	// 自动续费标示
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 预付费到期时间
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
}

type InstanceDetail struct {
	// 独享实例唯一id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 独享实例名字
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 独享实例描述
	InstanceDescription *string `json:"InstanceDescription,omitempty" name:"InstanceDescription"`

	// 独享实例计费类型
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 独享实例状态
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// 独享实例预付费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 独享实例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 独享实例网络类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConfig *NetworkConfig `json:"NetworkConfig,omitempty" name:"NetworkConfig"`

	// 独享实例vpc配置
	VpcConfig *VpcConfig `json:"VpcConfig,omitempty" name:"VpcConfig"`

	// 独享实例参数配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parameters []*ParameterInfo `json:"Parameters,omitempty" name:"Parameters"`

	// 独享实例隔离时间
	IsolationStartedTime *string `json:"IsolationStartedTime,omitempty" name:"IsolationStartedTime"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 可用区列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zones *string `json:"Zones,omitempty" name:"Zones"`
}

type InstanceInfo struct {
	// 独享实例唯一id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 独享实例name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 独享实例描述
	InstanceDescription *string `json:"InstanceDescription,omitempty" name:"InstanceDescription"`

	// 独享实例计费类型
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 独享实例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 独享实例状态
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// 独享实例创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 资源ID同唯一id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type InstanceParameterInput struct {
	// ServiceRequestNumPreSec，ApiRequestNumPreSec
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type InstanceSummary struct {
	// 专享实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 专享实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSet []*InstanceInfo `json:"InstanceSet,omitempty" name:"InstanceSet"`
}

type K8sLabel struct {
	// Label的Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Label的Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type K8sService struct {
	// 权重
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// k8s集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 容器命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 容器服务的名字
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 服务的端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 额外选择的Pod的Label
	ExtraLabels []*K8sLabel `json:"ExtraLabels,omitempty" name:"ExtraLabels"`

	// 自定义的服务名字，可选
	Name *string `json:"Name,omitempty" name:"Name"`
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

// Predefined struct for user
type ModifyAPIDocRequestParams struct {
	// API文档ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`

	// API文档名称
	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`

	// 服务名称
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名称
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 生成文档的API列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type ModifyAPIDocRequest struct {
	*tchttp.BaseRequest
	
	// API文档ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`

	// API文档名称
	ApiDocName *string `json:"ApiDocName,omitempty" name:"ApiDocName"`

	// 服务名称
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 环境名称
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 生成文档的API列表
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *ModifyAPIDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAPIDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiDocId")
	delete(f, "ApiDocName")
	delete(f, "ServiceId")
	delete(f, "Environment")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAPIDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAPIDocResponseParams struct {
	// API文档基本信息
	Result *APIDoc `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAPIDocResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAPIDocResponseParams `json:"Response"`
}

func (r *ModifyAPIDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAPIDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiAppRequestParams struct {
	// 应用唯一 ID。
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// 修改的应用名称
	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`

	// 修改的应用描述
	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`
}

type ModifyApiAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用唯一 ID。
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// 修改的应用名称
	ApiAppName *string `json:"ApiAppName,omitempty" name:"ApiAppName"`

	// 修改的应用描述
	ApiAppDesc *string `json:"ApiAppDesc,omitempty" name:"ApiAppDesc"`
}

func (r *ModifyApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	delete(f, "ApiAppName")
	delete(f, "ApiAppDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiAppResponseParams struct {
	// 修改操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApiAppResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApiAppResponseParams `json:"Response"`
}

func (r *ModifyApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiEnvironmentStrategyRequestParams struct {
	// 服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 限流值。
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 环境名。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// API列表。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
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
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *ModifyApiEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiEnvironmentStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Strategy")
	delete(f, "EnvironmentName")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApiEnvironmentStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiEnvironmentStrategyResponseParams struct {
	// 修改操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApiEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApiEnvironmentStrategyResponseParams `json:"Response"`
}

func (r *ModifyApiEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiEnvironmentStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiIncrementRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiIncrementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ApiId")
	delete(f, "BusinessType")
	delete(f, "PublicKey")
	delete(f, "LoginRedirectUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApiIncrementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiIncrementResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApiIncrementResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApiIncrementResponseParams `json:"Response"`
}

func (r *ModifyApiIncrementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiIncrementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiRequestParams struct {
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

	// API 鉴权类型。支持SECRET、NONE、OAUTH、APP。默认为NONE。
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
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`

	// 前端请求参数。
	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`

	// 当AuthType 为 OAUTH时，该字段有效， NORMAL：业务api   OAUTH：授权API。
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// API 的后端 Mock 返回信息。如果 ServiceType 是 Mock，则此参数必传。
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// API绑定微服务服务列表。
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices"`

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
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`

	// oauth配置。当AuthType是OAUTH时生效。
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// 用户自定义错误码配置。
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`

	// 是否开启Base64编码，只有后端为scf时才会生效。
	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`

	// 是否开启Base64编码的header触发，只有后端为scf时才会生效。
	IsBase64Trigger *bool `json:"IsBase64Trigger,omitempty" name:"IsBase64Trigger"`

	// Header触发规则，总规则数不能超过10。
	Base64EncodedTriggerRules []*Base64EncodedTriggerRule `json:"Base64EncodedTriggerRules,omitempty" name:"Base64EncodedTriggerRules"`

	// 事件总线ID。
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// scf函数类型。当后端类型是SCF时生效。支持事件触发(EVENT)，http直通云函数(HTTP)。
	ServiceScfFunctionType *string `json:"ServiceScfFunctionType,omitempty" name:"ServiceScfFunctionType"`

	// EIAM应用类型。
	EIAMAppType *string `json:"EIAMAppType,omitempty" name:"EIAMAppType"`

	// EIAM应用认证类型，支持仅认证（AuthenticationOnly）、认证和鉴权（Authorization）。
	EIAMAuthType *string `json:"EIAMAuthType,omitempty" name:"EIAMAuthType"`

	// EIAM应用Token 有效时间，单位为秒，默认为7200秒。
	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`

	// EIAM应用ID。
	TokenTimeout *int64 `json:"TokenTimeout,omitempty" name:"TokenTimeout"`
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

	// API 鉴权类型。支持SECRET、NONE、OAUTH、APP。默认为NONE。
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
	ConstantParameters []*ConstantParameter `json:"ConstantParameters,omitempty" name:"ConstantParameters"`

	// 前端请求参数。
	RequestParameters []*ReqParameter `json:"RequestParameters,omitempty" name:"RequestParameters"`

	// 当AuthType 为 OAUTH时，该字段有效， NORMAL：业务api   OAUTH：授权API。
	ApiBusinessType *string `json:"ApiBusinessType,omitempty" name:"ApiBusinessType"`

	// API 的后端 Mock 返回信息。如果 ServiceType 是 Mock，则此参数必传。
	ServiceMockReturnMessage *string `json:"ServiceMockReturnMessage,omitempty" name:"ServiceMockReturnMessage"`

	// API绑定微服务服务列表。
	MicroServices []*MicroServiceReq `json:"MicroServices,omitempty" name:"MicroServices"`

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
	ServiceParameters []*ServiceParameter `json:"ServiceParameters,omitempty" name:"ServiceParameters"`

	// oauth配置。当AuthType是OAUTH时生效。
	OauthConfig *OauthConfig `json:"OauthConfig,omitempty" name:"OauthConfig"`

	// 用户自定义错误码配置。
	ResponseErrorCodes []*ResponseErrorCodeReq `json:"ResponseErrorCodes,omitempty" name:"ResponseErrorCodes"`

	// 是否开启Base64编码，只有后端为scf时才会生效。
	IsBase64Encoded *bool `json:"IsBase64Encoded,omitempty" name:"IsBase64Encoded"`

	// 是否开启Base64编码的header触发，只有后端为scf时才会生效。
	IsBase64Trigger *bool `json:"IsBase64Trigger,omitempty" name:"IsBase64Trigger"`

	// Header触发规则，总规则数不能超过10。
	Base64EncodedTriggerRules []*Base64EncodedTriggerRule `json:"Base64EncodedTriggerRules,omitempty" name:"Base64EncodedTriggerRules"`

	// 事件总线ID。
	EventBusId *string `json:"EventBusId,omitempty" name:"EventBusId"`

	// scf函数类型。当后端类型是SCF时生效。支持事件触发(EVENT)，http直通云函数(HTTP)。
	ServiceScfFunctionType *string `json:"ServiceScfFunctionType,omitempty" name:"ServiceScfFunctionType"`

	// EIAM应用类型。
	EIAMAppType *string `json:"EIAMAppType,omitempty" name:"EIAMAppType"`

	// EIAM应用认证类型，支持仅认证（AuthenticationOnly）、认证和鉴权（Authorization）。
	EIAMAuthType *string `json:"EIAMAuthType,omitempty" name:"EIAMAuthType"`

	// EIAM应用Token 有效时间，单位为秒，默认为7200秒。
	EIAMAppId *string `json:"EIAMAppId,omitempty" name:"EIAMAppId"`

	// EIAM应用ID。
	TokenTimeout *int64 `json:"TokenTimeout,omitempty" name:"TokenTimeout"`
}

func (r *ModifyApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ServiceType")
	delete(f, "RequestConfig")
	delete(f, "ApiId")
	delete(f, "ApiName")
	delete(f, "ApiDesc")
	delete(f, "ApiType")
	delete(f, "AuthType")
	delete(f, "AuthRequired")
	delete(f, "ServiceTimeout")
	delete(f, "Protocol")
	delete(f, "EnableCORS")
	delete(f, "ConstantParameters")
	delete(f, "RequestParameters")
	delete(f, "ApiBusinessType")
	delete(f, "ServiceMockReturnMessage")
	delete(f, "MicroServices")
	delete(f, "ServiceTsfLoadBalanceConf")
	delete(f, "ServiceTsfHealthCheckConf")
	delete(f, "TargetServicesLoadBalanceConf")
	delete(f, "TargetServicesHealthCheckConf")
	delete(f, "ServiceScfFunctionName")
	delete(f, "ServiceWebsocketRegisterFunctionName")
	delete(f, "ServiceWebsocketCleanupFunctionName")
	delete(f, "ServiceWebsocketTransportFunctionName")
	delete(f, "ServiceScfFunctionNamespace")
	delete(f, "ServiceScfFunctionQualifier")
	delete(f, "ServiceWebsocketRegisterFunctionNamespace")
	delete(f, "ServiceWebsocketRegisterFunctionQualifier")
	delete(f, "ServiceWebsocketTransportFunctionNamespace")
	delete(f, "ServiceWebsocketTransportFunctionQualifier")
	delete(f, "ServiceWebsocketCleanupFunctionNamespace")
	delete(f, "ServiceWebsocketCleanupFunctionQualifier")
	delete(f, "ServiceScfIsIntegratedResponse")
	delete(f, "IsDebugAfterCharge")
	delete(f, "TagSpecifications")
	delete(f, "IsDeleteResponseErrorCodes")
	delete(f, "ResponseType")
	delete(f, "ResponseSuccessExample")
	delete(f, "ResponseFailExample")
	delete(f, "ServiceConfig")
	delete(f, "AuthRelationApiId")
	delete(f, "ServiceParameters")
	delete(f, "OauthConfig")
	delete(f, "ResponseErrorCodes")
	delete(f, "IsBase64Encoded")
	delete(f, "IsBase64Trigger")
	delete(f, "Base64EncodedTriggerRules")
	delete(f, "EventBusId")
	delete(f, "ServiceScfFunctionType")
	delete(f, "EIAMAppType")
	delete(f, "EIAMAuthType")
	delete(f, "EIAMAppId")
	delete(f, "TokenTimeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApiResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApiResponseParams `json:"Response"`
}

func (r *ModifyApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyExclusiveInstanceRequestParams struct {
	// 独享实例唯一id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 独享实例name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 独享实例描述
	InstanceDescription *string `json:"InstanceDescription,omitempty" name:"InstanceDescription"`

	// 独享实例参数配置
	Parameters []*InstanceParameterInput `json:"Parameters,omitempty" name:"Parameters"`
}

type ModifyExclusiveInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 独享实例唯一id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 独享实例name
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 独享实例描述
	InstanceDescription *string `json:"InstanceDescription,omitempty" name:"InstanceDescription"`

	// 独享实例参数配置
	Parameters []*InstanceParameterInput `json:"Parameters,omitempty" name:"Parameters"`
}

func (r *ModifyExclusiveInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyExclusiveInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "InstanceDescription")
	delete(f, "Parameters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyExclusiveInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyExclusiveInstanceResponseParams struct {
	// 独享实例详情信息
	Result *InstanceDetail `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyExclusiveInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyExclusiveInstanceResponseParams `json:"Response"`
}

func (r *ModifyExclusiveInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyExclusiveInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIPStrategyRequestParams struct {
	// 待修改的策略所属服务的唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待修改的策略唯一ID。
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 待修改的策略详情。
	StrategyData *string `json:"StrategyData,omitempty" name:"StrategyData"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIPStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyId")
	delete(f, "StrategyData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIPStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIPStrategyResponseParams struct {
	// 修改操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIPStrategyResponseParams `json:"Response"`
}

func (r *ModifyIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPluginRequestParams struct {
	// 要修改的插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 要修改的API网关插件名称。最长50个字符，支持 a-z,A-Z,0-9,_, 必须字母开头，字母或者数字结尾。
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 要修改的插件描述，限定200字以内。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 要修改的插件定义语句，支持json。
	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`
}

type ModifyPluginRequest struct {
	*tchttp.BaseRequest
	
	// 要修改的插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 要修改的API网关插件名称。最长50个字符，支持 a-z,A-Z,0-9,_, 必须字母开头，字母或者数字结尾。
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 要修改的插件描述，限定200字以内。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 要修改的插件定义语句，支持json。
	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`
}

func (r *ModifyPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "PluginName")
	delete(f, "Description")
	delete(f, "PluginData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPluginResponseParams struct {
	// 修改操作是否成功。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPluginResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPluginResponseParams `json:"Response"`
}

func (r *ModifyPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceEnvironmentStrategyRequestParams struct {
	// 服务的唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 限流值。
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 环境列表。
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames"`
}

type ModifyServiceEnvironmentStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 服务的唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 限流值。
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 环境列表。
	EnvironmentNames []*string `json:"EnvironmentNames,omitempty" name:"EnvironmentNames"`
}

func (r *ModifyServiceEnvironmentStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceEnvironmentStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "Strategy")
	delete(f, "EnvironmentNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceEnvironmentStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceEnvironmentStrategyResponseParams struct {
	// 修改操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyServiceEnvironmentStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceEnvironmentStrategyResponseParams `json:"Response"`
}

func (r *ModifyServiceEnvironmentStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceEnvironmentStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceRequestParams struct {
	// 待修改服务的唯一 Id。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 修改后的服务名称。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 修改后的服务描述。
	ServiceDesc *string `json:"ServiceDesc,omitempty" name:"ServiceDesc"`

	// 修改后的服务前端请求类型，如 http、https和 http&https。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 网络类型列表，用于指定支持的访问类型，INNER为内网访问，OUTER为外网访问。默认为OUTER。
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`
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
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`
}

func (r *ModifyServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ServiceName")
	delete(f, "ServiceDesc")
	delete(f, "Protocol")
	delete(f, "NetTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceResponseParams `json:"Response"`
}

func (r *ModifyServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubDomainRequestParams struct {
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
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`

	// 网络类型 （'INNER' 或 'OUTER'）
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 是否将HTTP请求强制跳转 HTTPS，默认为false。参数为 true时，API网关会将所有使用该自定义域名的 HTTP 协议的请求重定向至 HTTPS 协议进行转发。
	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`
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
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`

	// 网络类型 （'INNER' 或 'OUTER'）
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 是否将HTTP请求强制跳转 HTTPS，默认为false。参数为 true时，API网关会将所有使用该自定义域名的 HTTP 协议的请求重定向至 HTTPS 协议进行转发。
	IsForcedHttps *bool `json:"IsForcedHttps,omitempty" name:"IsForcedHttps"`
}

func (r *ModifySubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "SubDomain")
	delete(f, "IsDefaultMapping")
	delete(f, "CertificateId")
	delete(f, "Protocol")
	delete(f, "PathMappingSet")
	delete(f, "NetType")
	delete(f, "IsForcedHttps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubDomainResponseParams struct {
	// 修改自定义域名操作是否成功。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubDomainResponseParams `json:"Response"`
}

func (r *ModifySubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUpstreamRequestParams struct {
	// 后端通道唯一ID
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// 后端通道名字
	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`

	// 后端通道描述
	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`

	// 后端协议，取值范围：HTTP, HTTPS
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 后端访问类型，取值范围：IP_PORT, K8S
	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// 负载均衡算法，取值范围：ROUND_ROBIN
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// VPC唯一ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 请求重试次数，默认3次
	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`

	// 网关转发到后端的 Host 请求头
	UpstreamHost *string `json:"UpstreamHost,omitempty" name:"UpstreamHost"`

	// 后端节点列表
	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`

	// 健康检查配置，目前只支持VPC通道
	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`

	// 容器服务配置
	K8sService []*K8sService `json:"K8sService,omitempty" name:"K8sService"`
}

type ModifyUpstreamRequest struct {
	*tchttp.BaseRequest
	
	// 后端通道唯一ID
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// 后端通道名字
	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`

	// 后端通道描述
	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`

	// 后端协议，取值范围：HTTP, HTTPS
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 后端访问类型，取值范围：IP_PORT, K8S
	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// 负载均衡算法，取值范围：ROUND_ROBIN
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// VPC唯一ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 请求重试次数，默认3次
	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`

	// 网关转发到后端的 Host 请求头
	UpstreamHost *string `json:"UpstreamHost,omitempty" name:"UpstreamHost"`

	// 后端节点列表
	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`

	// 健康检查配置，目前只支持VPC通道
	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`

	// 容器服务配置
	K8sService []*K8sService `json:"K8sService,omitempty" name:"K8sService"`
}

func (r *ModifyUpstreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUpstreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UpstreamId")
	delete(f, "UpstreamName")
	delete(f, "UpstreamDescription")
	delete(f, "Scheme")
	delete(f, "UpstreamType")
	delete(f, "Algorithm")
	delete(f, "UniqVpcId")
	delete(f, "Retries")
	delete(f, "UpstreamHost")
	delete(f, "Nodes")
	delete(f, "HealthChecker")
	delete(f, "K8sService")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUpstreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUpstreamResponseParams struct {
	// 返回修改后的后端通道信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *UpstreamInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUpstreamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUpstreamResponseParams `json:"Response"`
}

func (r *ModifyUpstreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUpstreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUsagePlanRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUsagePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	delete(f, "UsagePlanName")
	delete(f, "UsagePlanDesc")
	delete(f, "MaxRequestNum")
	delete(f, "MaxRequestNumPreSec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUsagePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUsagePlanResponseParams struct {
	// 使用计划详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *UsagePlanInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUsagePlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUsagePlanResponseParams `json:"Response"`
}

func (r *ModifyUsagePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUsagePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkConfig struct {
	// 最大出带宽
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// EnableInternetInbound信息
	EnableInternetInbound *bool `json:"EnableInternetInbound,omitempty" name:"EnableInternetInbound"`

	// EnableInternetOutbound信息
	EnableInternetOutbound *bool `json:"EnableInternetOutbound,omitempty" name:"EnableInternetOutbound"`

	// InboundIpAddresses信息
	InboundIpAddresses []*string `json:"InboundIpAddresses,omitempty" name:"InboundIpAddresses"`

	// OutboundIpAddresses信息
	OutboundIpAddresses []*string `json:"OutboundIpAddresses,omitempty" name:"OutboundIpAddresses"`
}

type OauthConfig struct {
	// 公钥，用于验证用户token。
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// token传递位置。
	TokenLocation *string `json:"TokenLocation,omitempty" name:"TokenLocation"`

	// 重定向地址，用于引导用户登录操作。
	LoginRedirectUrl *string `json:"LoginRedirectUrl,omitempty" name:"LoginRedirectUrl"`
}

type ParameterInfo struct {
	// 名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 当前值
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// 默认值
	Default *int64 `json:"Default,omitempty" name:"Default"`

	// 单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 类型, integer|string
	Type *string `json:"Type,omitempty" name:"Type"`

	// 最小
	Minimum *int64 `json:"Minimum,omitempty" name:"Minimum"`

	// 最大
	Maximum *int64 `json:"Maximum,omitempty" name:"Maximum"`

	// 修改时间
	ModifedTime *string `json:"ModifedTime,omitempty" name:"ModifedTime"`

	// 字符类型的值，当Type为string时才有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueString *string `json:"ValueString,omitempty" name:"ValueString"`

	// 字符类型的默认值，当Type为string时才有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValueString *string `json:"DefaultValueString,omitempty" name:"DefaultValueString"`

	// 可调整范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Range *string `json:"Range,omitempty" name:"Range"`
}

type PathMapping struct {
	// 路径。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 发布环境，可选值为“test”、 ”prepub“、”release“。
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type Plugin struct {
	// 插件ID。
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// 插件名称。
	PluginName *string `json:"PluginName,omitempty" name:"PluginName"`

	// 插件类型。
	PluginType *string `json:"PluginType,omitempty" name:"PluginType"`

	// 插件定义语句。
	PluginData *string `json:"PluginData,omitempty" name:"PluginData"`

	// 插件描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 插件创建时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 插件修改时间。按照 ISO8601 标准表示，并且使用 UTC 时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 插件绑定的API总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachedApiTotalCount *int64 `json:"AttachedApiTotalCount,omitempty" name:"AttachedApiTotalCount"`

	// 插件绑定的API信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachedApis []*AttachedApiInfo `json:"AttachedApis,omitempty" name:"AttachedApis"`
}

type PluginSummary struct {
	// 插件个数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 插件详情。
	PluginSet []*Plugin `json:"PluginSet,omitempty" name:"PluginSet"`
}

type ReleaseService struct {
	// 发布时的备注信息填写。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// 发布的版本id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" name:"ReleaseVersion"`
}

// Predefined struct for user
type ReleaseServiceRequestParams struct {
	// 待发布服务的唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待发布的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 本次的发布描述。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// apiId列表，预留字段，默认全量api发布。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
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
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *ReleaseServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "ReleaseDesc")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseServiceResponseParams struct {
	// 发布信息。
	Result *ReleaseService `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReleaseServiceResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseServiceResponseParams `json:"Response"`
}

func (r *ReleaseServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
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

// Predefined struct for user
type ResetAPIDocPasswordRequestParams struct {
	// API文档ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

type ResetAPIDocPasswordRequest struct {
	*tchttp.BaseRequest
	
	// API文档ID
	ApiDocId *string `json:"ApiDocId,omitempty" name:"ApiDocId"`
}

func (r *ResetAPIDocPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAPIDocPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiDocId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAPIDocPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAPIDocPasswordResponseParams struct {
	// API文档基本信息
	Result *APIDoc `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetAPIDocPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetAPIDocPasswordResponseParams `json:"Response"`
}

func (r *ResetAPIDocPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAPIDocPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	NetTypes []*string `json:"NetTypes,omitempty" name:"NetTypes"`

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
	AvailableEnvironments []*string `json:"AvailableEnvironments,omitempty" name:"AvailableEnvironments"`

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
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 独享实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetType *string `json:"SetType,omitempty" name:"SetType"`

	// 服务部署的集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeploymentType *string `json:"DeploymentType,omitempty" name:"DeploymentType"`
}

type ServiceConfig struct {
	// 后端类型。启用vpc时生效，目前支持的类型为clb, cvm和upstream
	Product *string `json:"Product,omitempty" name:"Product"`

	// vpc 的唯一ID。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// API 的后端服务url。如果ServiceType是HTTP，则此参数必传。
	Url *string `json:"Url,omitempty" name:"Url"`

	// API 的后端服务路径，如 /path。如果 ServiceType 是 HTTP，则此参数必传。前后端路径可不同。
	Path *string `json:"Path,omitempty" name:"Path"`

	// API的后端服务请求方法，如 GET。如果 ServiceType 是 HTTP，则此参数必传。前后端方法可不同。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 当绑定vpc通道才需要
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// API后端COS配置。如果 ServiceType 是 COS，则此参数必传。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosConfig *CosConfig `json:"CosConfig,omitempty" name:"CosConfig"`
}

type ServiceEnvironmentSet struct {
	// 服务绑定环境总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 服务绑定环境列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentList []*Environment `json:"EnvironmentList,omitempty" name:"EnvironmentList"`
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

	// 最大限流值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxStrategy *int64 `json:"MaxStrategy,omitempty" name:"MaxStrategy"`
}

type ServiceEnvironmentStrategyStatus struct {
	// 限流策略数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 限流策略列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentList []*ServiceEnvironmentStrategy `json:"EnvironmentList,omitempty" name:"EnvironmentList"`
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
	VersionList []*ServiceReleaseHistoryInfo `json:"VersionList,omitempty" name:"VersionList"`
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
	VersionList []*ServiceReleaseHistoryInfo `json:"VersionList,omitempty" name:"VersionList"`
}

type ServiceSubDomainMappings struct {
	// 是否使用默认路径映射，为 True 表示使用默认路径映射；为 False 的话，表示使用自定义路径映射，此时 PathMappingSet 不为空。
	IsDefaultMapping *bool `json:"IsDefaultMapping,omitempty" name:"IsDefaultMapping"`

	// 自定义路径映射列表。
	PathMappingSet []*PathMapping `json:"PathMappingSet,omitempty" name:"PathMappingSet"`
}

type ServiceUsagePlanSet struct {
	// 服务上绑定的使用计划总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 服务上绑定的使用计划列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceUsagePlanList []*ApiUsagePlan `json:"ServiceUsagePlanList,omitempty" name:"ServiceUsagePlanList"`
}

type ServicesStatus struct {
	// 服务列表总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 服务列表详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceSet []*Service `json:"ServiceSet,omitempty" name:"ServiceSet"`
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

// Predefined struct for user
type UnBindEnvironmentRequestParams struct {
	// 绑定类型，取值为 API、SERVICE，默认值为 SERVICE。
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// 待绑定的使用计划唯一 ID 列表。
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds"`

	// 待解绑的服务环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 待解绑的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 唯一 ID 数组，当 BindType=API 时，需要传入此参数。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type UnBindEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 绑定类型，取值为 API、SERVICE，默认值为 SERVICE。
	BindType *string `json:"BindType,omitempty" name:"BindType"`

	// 待绑定的使用计划唯一 ID 列表。
	UsagePlanIds []*string `json:"UsagePlanIds,omitempty" name:"UsagePlanIds"`

	// 待解绑的服务环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 待解绑的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// API 唯一 ID 数组，当 BindType=API 时，需要传入此参数。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *UnBindEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BindType")
	delete(f, "UsagePlanIds")
	delete(f, "Environment")
	delete(f, "ServiceId")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindEnvironmentResponseParams struct {
	// 解绑操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnBindEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *UnBindEnvironmentResponseParams `json:"Response"`
}

func (r *UnBindEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindIPStrategyRequestParams struct {
	// 待解绑的服务唯一ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待解绑的IP策略唯一ID。
	StrategyId *string `json:"StrategyId,omitempty" name:"StrategyId"`

	// 待解绑的环境。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 待解绑的 API 列表。
	UnBindApiIds []*string `json:"UnBindApiIds,omitempty" name:"UnBindApiIds"`
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
	UnBindApiIds []*string `json:"UnBindApiIds,omitempty" name:"UnBindApiIds"`
}

func (r *UnBindIPStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindIPStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "StrategyId")
	delete(f, "EnvironmentName")
	delete(f, "UnBindApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindIPStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindIPStrategyResponseParams struct {
	// 解绑操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnBindIPStrategyResponse struct {
	*tchttp.BaseResponse
	Response *UnBindIPStrategyResponseParams `json:"Response"`
}

func (r *UnBindIPStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindIPStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindSecretIdsRequestParams struct {
	// 待解绑的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 待解绑的密钥 ID 数组。
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds"`
}

type UnBindSecretIdsRequest struct {
	*tchttp.BaseRequest
	
	// 待解绑的使用计划唯一 ID。
	UsagePlanId *string `json:"UsagePlanId,omitempty" name:"UsagePlanId"`

	// 待解绑的密钥 ID 数组。
	AccessKeyIds []*string `json:"AccessKeyIds,omitempty" name:"AccessKeyIds"`
}

func (r *UnBindSecretIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindSecretIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UsagePlanId")
	delete(f, "AccessKeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindSecretIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindSecretIdsResponseParams struct {
	// 解绑操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnBindSecretIdsResponse struct {
	*tchttp.BaseResponse
	Response *UnBindSecretIdsResponseParams `json:"Response"`
}

func (r *UnBindSecretIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindSecretIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindSubDomainRequestParams struct {
	// 服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待解绑的自定义的域名。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindSubDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "SubDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindSubDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindSubDomainResponseParams struct {
	// 解绑自定义域名操作是否成功。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnBindSubDomainResponse struct {
	*tchttp.BaseResponse
	Response *UnBindSubDomainResponseParams `json:"Response"`
}

func (r *UnBindSubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindSubDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnReleaseServiceRequestParams struct {
	// 待下线服务的唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待下线的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 保留字段，待下线的API列表。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

type UnReleaseServiceRequest struct {
	*tchttp.BaseRequest
	
	// 待下线服务的唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待下线的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 保留字段，待下线的API列表。
	ApiIds []*string `json:"ApiIds,omitempty" name:"ApiIds"`
}

func (r *UnReleaseServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnReleaseServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "ApiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnReleaseServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnReleaseServiceResponseParams struct {
	// 下线操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnReleaseServiceResponse struct {
	*tchttp.BaseResponse
	Response *UnReleaseServiceResponseParams `json:"Response"`
}

func (r *UnReleaseServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnReleaseServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindApiAppRequestParams struct {
	// 待绑定的应用唯一 ID 。
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// 待绑定的环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 待绑定的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待绑定的API唯一ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

type UnbindApiAppRequest struct {
	*tchttp.BaseRequest
	
	// 待绑定的应用唯一 ID 。
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// 待绑定的环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 待绑定的服务唯一 ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待绑定的API唯一ID。
	ApiId *string `json:"ApiId,omitempty" name:"ApiId"`
}

func (r *UnbindApiAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindApiAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	delete(f, "Environment")
	delete(f, "ServiceId")
	delete(f, "ApiId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindApiAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindApiAppResponseParams struct {
	// 解除绑定操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindApiAppResponse struct {
	*tchttp.BaseResponse
	Response *UnbindApiAppResponseParams `json:"Response"`
}

func (r *UnbindApiAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindApiAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiAppKeyRequestParams struct {
	// 应用唯一 ID。
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// 应用的Key。
	ApiAppKey *string `json:"ApiAppKey,omitempty" name:"ApiAppKey"`

	// 应用的Secret。
	ApiAppSecret *string `json:"ApiAppSecret,omitempty" name:"ApiAppSecret"`
}

type UpdateApiAppKeyRequest struct {
	*tchttp.BaseRequest
	
	// 应用唯一 ID。
	ApiAppId *string `json:"ApiAppId,omitempty" name:"ApiAppId"`

	// 应用的Key。
	ApiAppKey *string `json:"ApiAppKey,omitempty" name:"ApiAppKey"`

	// 应用的Secret。
	ApiAppSecret *string `json:"ApiAppSecret,omitempty" name:"ApiAppSecret"`
}

func (r *UpdateApiAppKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiAppKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiAppId")
	delete(f, "ApiAppKey")
	delete(f, "ApiAppSecret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateApiAppKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiAppKeyResponseParams struct {
	// 更新操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateApiAppKeyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateApiAppKeyResponseParams `json:"Response"`
}

func (r *UpdateApiAppKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiAppKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiKeyRequestParams struct {
	// 待更换的密钥 ID。
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// 待更换的密钥 Key，更新自定义密钥时，该字段为必传。长度10 - 50字符，包括字母、数字、英文下划线。
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" name:"AccessKeySecret"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	delete(f, "AccessKeySecret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApiKeyResponseParams struct {
	// 更换后的密钥详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApiKey `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateApiKeyResponseParams `json:"Response"`
}

func (r *UpdateApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateServiceRequestParams struct {
	// 待切换服务的唯一 Id。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 待切换的环境名称，当前支持三个环境，test（测试环境）、prepub（预发布环境）和 release（发布环境）。
	EnvironmentName *string `json:"EnvironmentName,omitempty" name:"EnvironmentName"`

	// 切换的版本号。
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 本次的切换描述。
	UpdateDesc *string `json:"UpdateDesc,omitempty" name:"UpdateDesc"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "EnvironmentName")
	delete(f, "VersionName")
	delete(f, "UpdateDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateServiceResponseParams struct {
	// 切换版本操作是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateServiceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateServiceResponseParams `json:"Response"`
}

func (r *UpdateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpstreamHealthChecker struct {
	// 标识是否开启主动健康检查。
	EnableActiveCheck *bool `json:"EnableActiveCheck,omitempty" name:"EnableActiveCheck"`

	// 标识是否开启被动健康检查。
	EnablePassiveCheck *bool `json:"EnablePassiveCheck,omitempty" name:"EnablePassiveCheck"`

	// 健康检查时，判断为成功请求的 HTTP 状态码。
	HealthyHttpStatus *string `json:"HealthyHttpStatus,omitempty" name:"HealthyHttpStatus"`

	// 健康检查时，判断为失败请求的 HTTP 状态码。
	UnhealthyHttpStatus *string `json:"UnhealthyHttpStatus,omitempty" name:"UnhealthyHttpStatus"`

	// TCP连续错误阈值。0 表示禁用 TCP 检查。取值范围：[0, 254]。
	TcpFailureThreshold *uint64 `json:"TcpFailureThreshold,omitempty" name:"TcpFailureThreshold"`

	// 连续超时阈值。0 表示禁用超时检查。取值范围：[0, 254]。
	TimeoutThreshold *uint64 `json:"TimeoutThreshold,omitempty" name:"TimeoutThreshold"`

	// HTTP连续错误阈值。0 表示禁用HTTP检查。取值范围：[0, 254]。
	HttpFailureThreshold *uint64 `json:"HttpFailureThreshold,omitempty" name:"HttpFailureThreshold"`

	// 主动健康检查时探测请求的路径。默认为"/"。
	ActiveCheckHttpPath *string `json:"ActiveCheckHttpPath,omitempty" name:"ActiveCheckHttpPath"`

	// 主动健康检查的探测请求超时，单位秒。默认为5秒。
	ActiveCheckTimeout *uint64 `json:"ActiveCheckTimeout,omitempty" name:"ActiveCheckTimeout"`

	// 主动健康检查的时间间隔，默认5秒。
	ActiveCheckInterval *uint64 `json:"ActiveCheckInterval,omitempty" name:"ActiveCheckInterval"`

	// 主动健康检查时探测请求的的请求头。
	ActiveRequestHeader []*UpstreamHealthCheckerReqHeaders `json:"ActiveRequestHeader,omitempty" name:"ActiveRequestHeader"`

	// 异常节点的状态自动恢复时间，单位秒。当只开启被动检查的话，必须设置为 > 0 的值，否则被动异常节点将无法恢复。默认30秒。
	UnhealthyTimeout *uint64 `json:"UnhealthyTimeout,omitempty" name:"UnhealthyTimeout"`
}

type UpstreamHealthCheckerReqHeaders struct {

}

type UpstreamInfo struct {
	// 后端通道唯一ID
	UpstreamId *string `json:"UpstreamId,omitempty" name:"UpstreamId"`

	// 后端通道名字
	UpstreamName *string `json:"UpstreamName,omitempty" name:"UpstreamName"`

	// 后端通道描述
	UpstreamDescription *string `json:"UpstreamDescription,omitempty" name:"UpstreamDescription"`

	// 后端协议，取值范围：HTTP, HTTPS
	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`

	// 负载均衡算法，取值范围：ROUND_ROBIN
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`

	// VPC唯一ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 请求重试次数
	Retries *uint64 `json:"Retries,omitempty" name:"Retries"`

	// 后端节点
	Nodes []*UpstreamNode `json:"Nodes,omitempty" name:"Nodes"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 健康检查配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthChecker *UpstreamHealthChecker `json:"HealthChecker,omitempty" name:"HealthChecker"`

	// 后端的类型，取值范围：IP_PORT, K8S
	UpstreamType *string `json:"UpstreamType,omitempty" name:"UpstreamType"`

	// K8S容器服务配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	K8sServices []*K8sService `json:"K8sServices,omitempty" name:"K8sServices"`

	// 网关转发给后端的Host请求头
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamHost *string `json:"UpstreamHost,omitempty" name:"UpstreamHost"`
}

type UpstreamNode struct {
	// IP或域名
	Host *string `json:"Host,omitempty" name:"Host"`

	// 端口[0, 65535]
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 权重[0, 100], 0为禁用
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// CVM实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VmInstanceId *string `json:"VmInstanceId,omitempty" name:"VmInstanceId"`

	// 染色标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 节点健康状态，创建、编辑时不需要传该参数。OFF：关闭，HEALTHY：健康，UNHEALTHY：异常，NO_DATA：数据未上报。目前只支持VPC通道。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Healthy *string `json:"Healthy,omitempty" name:"Healthy"`

	// K8S容器服务名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// K8S命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameSpace *string `json:"NameSpace,omitempty" name:"NameSpace"`

	// TKE集群的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Node的来源，取值范围：K8S
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitempty" name:"Source"`

	// API网关内部记录唯一的服务名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqueServiceName *string `json:"UniqueServiceName,omitempty" name:"UniqueServiceName"`
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
	AccessKeyList []*UsagePlanBindSecret `json:"AccessKeyList,omitempty" name:"AccessKeyList"`
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
	EnvironmentList []*UsagePlanEnvironment `json:"EnvironmentList,omitempty" name:"EnvironmentList"`
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
	BindSecretIds []*string `json:"BindSecretIds,omitempty" name:"BindSecretIds"`

	// 绑定环境数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindEnvironmentTotalCount *int64 `json:"BindEnvironmentTotalCount,omitempty" name:"BindEnvironmentTotalCount"`

	// 绑定环境详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindEnvironments []*UsagePlanBindEnvironment `json:"BindEnvironments,omitempty" name:"BindEnvironments"`
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
	UsagePlanStatusSet []*UsagePlanStatusInfo `json:"UsagePlanStatusSet,omitempty" name:"UsagePlanStatusSet"`
}

type VpcConfig struct {
	// vpcid
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// subnetid
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
}