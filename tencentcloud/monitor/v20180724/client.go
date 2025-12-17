// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20180724

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-24"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewBindPrometheusManagedGrafanaRequest() (request *BindPrometheusManagedGrafanaRequest) {
    request = &BindPrometheusManagedGrafanaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "BindPrometheusManagedGrafana")
    
    
    return
}

func NewBindPrometheusManagedGrafanaResponse() (response *BindPrometheusManagedGrafanaResponse) {
    response = &BindPrometheusManagedGrafanaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindPrometheusManagedGrafana
// 绑定 Grafana 可视化服务实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) BindPrometheusManagedGrafana(request *BindPrometheusManagedGrafanaRequest) (response *BindPrometheusManagedGrafanaResponse, err error) {
    return c.BindPrometheusManagedGrafanaWithContext(context.Background(), request)
}

// BindPrometheusManagedGrafana
// 绑定 Grafana 可视化服务实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) BindPrometheusManagedGrafanaWithContext(ctx context.Context, request *BindPrometheusManagedGrafanaRequest) (response *BindPrometheusManagedGrafanaResponse, err error) {
    if request == nil {
        request = NewBindPrometheusManagedGrafanaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "BindPrometheusManagedGrafana")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindPrometheusManagedGrafana require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindPrometheusManagedGrafanaResponse()
    err = c.Send(request, response)
    return
}

func NewBindingPolicyObjectRequest() (request *BindingPolicyObjectRequest) {
    request = &BindingPolicyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "BindingPolicyObject")
    
    
    return
}

func NewBindingPolicyObjectResponse() (response *BindingPolicyObjectResponse) {
    response = &BindingPolicyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindingPolicyObject
// 将告警策略绑定到特定对象
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindingPolicyObject(request *BindingPolicyObjectRequest) (response *BindingPolicyObjectResponse, err error) {
    return c.BindingPolicyObjectWithContext(context.Background(), request)
}

// BindingPolicyObject
// 将告警策略绑定到特定对象
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindingPolicyObjectWithContext(ctx context.Context, request *BindingPolicyObjectRequest) (response *BindingPolicyObjectResponse, err error) {
    if request == nil {
        request = NewBindingPolicyObjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "BindingPolicyObject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindingPolicyObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindingPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewBindingPolicyTagRequest() (request *BindingPolicyTagRequest) {
    request = &BindingPolicyTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "BindingPolicyTag")
    
    
    return
}

func NewBindingPolicyTagResponse() (response *BindingPolicyTagResponse) {
    response = &BindingPolicyTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindingPolicyTag
// 策略绑定标签
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) BindingPolicyTag(request *BindingPolicyTagRequest) (response *BindingPolicyTagResponse, err error) {
    return c.BindingPolicyTagWithContext(context.Background(), request)
}

// BindingPolicyTag
// 策略绑定标签
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) BindingPolicyTagWithContext(ctx context.Context, request *BindingPolicyTagRequest) (response *BindingPolicyTagResponse, err error) {
    if request == nil {
        request = NewBindingPolicyTagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "BindingPolicyTag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindingPolicyTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindingPolicyTagResponse()
    err = c.Send(request, response)
    return
}

func NewCleanGrafanaInstanceRequest() (request *CleanGrafanaInstanceRequest) {
    request = &CleanGrafanaInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CleanGrafanaInstance")
    
    
    return
}

func NewCleanGrafanaInstanceResponse() (response *CleanGrafanaInstanceResponse) {
    response = &CleanGrafanaInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CleanGrafanaInstance
// 强制销毁 Grafana 实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CleanGrafanaInstance(request *CleanGrafanaInstanceRequest) (response *CleanGrafanaInstanceResponse, err error) {
    return c.CleanGrafanaInstanceWithContext(context.Background(), request)
}

// CleanGrafanaInstance
// 强制销毁 Grafana 实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CleanGrafanaInstanceWithContext(ctx context.Context, request *CleanGrafanaInstanceRequest) (response *CleanGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewCleanGrafanaInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CleanGrafanaInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CleanGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCleanGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmNoticeRequest() (request *CreateAlarmNoticeRequest) {
    request = &CreateAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAlarmNotice")
    
    
    return
}

func NewCreateAlarmNoticeResponse() (response *CreateAlarmNoticeResponse) {
    response = &CreateAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlarmNotice
// 创建通知模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmNotice(request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
    return c.CreateAlarmNoticeWithContext(context.Background(), request)
}

// CreateAlarmNotice
// 创建通知模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmNoticeWithContext(ctx context.Context, request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewCreateAlarmNoticeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateAlarmNotice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmPolicyRequest() (request *CreateAlarmPolicyRequest) {
    request = &CreateAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAlarmPolicy")
    
    
    return
}

func NewCreateAlarmPolicyResponse() (response *CreateAlarmPolicyResponse) {
    response = &CreateAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlarmPolicy
// 创建告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmPolicy(request *CreateAlarmPolicyRequest) (response *CreateAlarmPolicyResponse, err error) {
    return c.CreateAlarmPolicyWithContext(context.Background(), request)
}

// CreateAlarmPolicy
// 创建告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmPolicyWithContext(ctx context.Context, request *CreateAlarmPolicyRequest) (response *CreateAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAlarmPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateAlarmPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarmPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmShieldRequest() (request *CreateAlarmShieldRequest) {
    request = &CreateAlarmShieldRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAlarmShield")
    
    
    return
}

func NewCreateAlarmShieldResponse() (response *CreateAlarmShieldResponse) {
    response = &CreateAlarmShieldResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlarmShield
// 创建告警屏蔽规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmShield(request *CreateAlarmShieldRequest) (response *CreateAlarmShieldResponse, err error) {
    return c.CreateAlarmShieldWithContext(context.Background(), request)
}

// CreateAlarmShield
// 创建告警屏蔽规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmShieldWithContext(ctx context.Context, request *CreateAlarmShieldRequest) (response *CreateAlarmShieldResponse, err error) {
    if request == nil {
        request = NewCreateAlarmShieldRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateAlarmShield")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarmShield require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmShieldResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlertRuleRequest() (request *CreateAlertRuleRequest) {
    request = &CreateAlertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAlertRule")
    
    
    return
}

func NewCreateAlertRuleResponse() (response *CreateAlertRuleResponse) {
    response = &CreateAlertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlertRule
// 创建 Prometheus 告警规则。
//
// 
//
// 请注意，**告警对象和告警消息是 Prometheus Rule Annotations 的特殊字段，需要通过 annotations 来传递，对应的 Key 分别为summary/description**，，请参考 [Prometheus Rule更多配置请参考](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/)。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateAlertRule(request *CreateAlertRuleRequest) (response *CreateAlertRuleResponse, err error) {
    return c.CreateAlertRuleWithContext(context.Background(), request)
}

// CreateAlertRule
// 创建 Prometheus 告警规则。
//
// 
//
// 请注意，**告警对象和告警消息是 Prometheus Rule Annotations 的特殊字段，需要通过 annotations 来传递，对应的 Key 分别为summary/description**，，请参考 [Prometheus Rule更多配置请参考](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/)。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateAlertRuleWithContext(ctx context.Context, request *CreateAlertRuleRequest) (response *CreateAlertRuleResponse, err error) {
    if request == nil {
        request = NewCreateAlertRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateAlertRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlertRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConditionsTemplateRequest() (request *CreateConditionsTemplateRequest) {
    request = &CreateConditionsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateConditionsTemplate")
    
    
    return
}

func NewCreateConditionsTemplateResponse() (response *CreateConditionsTemplateResponse) {
    response = &CreateConditionsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConditionsTemplate
// 创建告警条件模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateConditionsTemplate(request *CreateConditionsTemplateRequest) (response *CreateConditionsTemplateResponse, err error) {
    return c.CreateConditionsTemplateWithContext(context.Background(), request)
}

// CreateConditionsTemplate
// 创建告警条件模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateConditionsTemplateWithContext(ctx context.Context, request *CreateConditionsTemplateRequest) (response *CreateConditionsTemplateResponse, err error) {
    if request == nil {
        request = NewCreateConditionsTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateConditionsTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConditionsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConditionsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExporterIntegrationRequest() (request *CreateExporterIntegrationRequest) {
    request = &CreateExporterIntegrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateExporterIntegration")
    
    
    return
}

func NewCreateExporterIntegrationResponse() (response *CreateExporterIntegrationResponse) {
    response = &CreateExporterIntegrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateExporterIntegration
// 创建集成中心 exporter 集成，因集成较多，建议控制台创建集成。(前提：已授权创建托管 EKS 集群，验证方式：1. 控制台界面确认，未提示授权则表示已授权创建；2. 通过 DescribePrometheusInstanceInitStatus 接口查询集群状态，如果托管集群不存在，可通过 RunPrometheusInstance 接口创建)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateExporterIntegration(request *CreateExporterIntegrationRequest) (response *CreateExporterIntegrationResponse, err error) {
    return c.CreateExporterIntegrationWithContext(context.Background(), request)
}

// CreateExporterIntegration
// 创建集成中心 exporter 集成，因集成较多，建议控制台创建集成。(前提：已授权创建托管 EKS 集群，验证方式：1. 控制台界面确认，未提示授权则表示已授权创建；2. 通过 DescribePrometheusInstanceInitStatus 接口查询集群状态，如果托管集群不存在，可通过 RunPrometheusInstance 接口创建)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateExporterIntegrationWithContext(ctx context.Context, request *CreateExporterIntegrationRequest) (response *CreateExporterIntegrationResponse, err error) {
    if request == nil {
        request = NewCreateExporterIntegrationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateExporterIntegration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExporterIntegration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExporterIntegrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExternalClusterRequest() (request *CreateExternalClusterRequest) {
    request = &CreateExternalClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateExternalCluster")
    
    
    return
}

func NewCreateExternalClusterResponse() (response *CreateExternalClusterResponse) {
    response = &CreateExternalClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateExternalCluster
// 注册外部集群到云上 TMP 实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCECONFLICT = "FailedOperation.ResourceConflict"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateExternalCluster(request *CreateExternalClusterRequest) (response *CreateExternalClusterResponse, err error) {
    return c.CreateExternalClusterWithContext(context.Background(), request)
}

// CreateExternalCluster
// 注册外部集群到云上 TMP 实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCECONFLICT = "FailedOperation.ResourceConflict"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateExternalClusterWithContext(ctx context.Context, request *CreateExternalClusterRequest) (response *CreateExternalClusterResponse, err error) {
    if request == nil {
        request = NewCreateExternalClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateExternalCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExternalCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExternalClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGrafanaInstanceRequest() (request *CreateGrafanaInstanceRequest) {
    request = &CreateGrafanaInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateGrafanaInstance")
    
    
    return
}

func NewCreateGrafanaInstanceResponse() (response *CreateGrafanaInstanceResponse) {
    response = &CreateGrafanaInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGrafanaInstance
// 本接口（CreateGrafanaInstance）用于创建 Grafana 包年包月实例，默认基础版、到期自动续费。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_CREATEINSTANCELIMITED = "FailedOperation.CreateInstanceLimited"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_REGIONUNAVAILABLE = "FailedOperation.RegionUnavailable"
//  FAILEDOPERATION_ZONEUNAVAILABLE = "FailedOperation.ZoneUnavailable"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) CreateGrafanaInstance(request *CreateGrafanaInstanceRequest) (response *CreateGrafanaInstanceResponse, err error) {
    return c.CreateGrafanaInstanceWithContext(context.Background(), request)
}

// CreateGrafanaInstance
// 本接口（CreateGrafanaInstance）用于创建 Grafana 包年包月实例，默认基础版、到期自动续费。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_CREATEINSTANCELIMITED = "FailedOperation.CreateInstanceLimited"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_REGIONUNAVAILABLE = "FailedOperation.RegionUnavailable"
//  FAILEDOPERATION_ZONEUNAVAILABLE = "FailedOperation.ZoneUnavailable"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) CreateGrafanaInstanceWithContext(ctx context.Context, request *CreateGrafanaInstanceRequest) (response *CreateGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewCreateGrafanaInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateGrafanaInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGrafanaIntegrationRequest() (request *CreateGrafanaIntegrationRequest) {
    request = &CreateGrafanaIntegrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateGrafanaIntegration")
    
    
    return
}

func NewCreateGrafanaIntegrationResponse() (response *CreateGrafanaIntegrationResponse) {
    response = &CreateGrafanaIntegrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGrafanaIntegration
// 创建 Grafana 集成配置，其中 Prometheus 集成不通过该接口创建，可参考 BindPrometheusManagedGrafana 接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateGrafanaIntegration(request *CreateGrafanaIntegrationRequest) (response *CreateGrafanaIntegrationResponse, err error) {
    return c.CreateGrafanaIntegrationWithContext(context.Background(), request)
}

// CreateGrafanaIntegration
// 创建 Grafana 集成配置，其中 Prometheus 集成不通过该接口创建，可参考 BindPrometheusManagedGrafana 接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateGrafanaIntegrationWithContext(ctx context.Context, request *CreateGrafanaIntegrationRequest) (response *CreateGrafanaIntegrationResponse, err error) {
    if request == nil {
        request = NewCreateGrafanaIntegrationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateGrafanaIntegration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGrafanaIntegration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGrafanaIntegrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGrafanaNotificationChannelRequest() (request *CreateGrafanaNotificationChannelRequest) {
    request = &CreateGrafanaNotificationChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateGrafanaNotificationChannel")
    
    
    return
}

func NewCreateGrafanaNotificationChannelResponse() (response *CreateGrafanaNotificationChannelResponse) {
    response = &CreateGrafanaNotificationChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGrafanaNotificationChannel
// 创建 Grafana 告警通道
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateGrafanaNotificationChannel(request *CreateGrafanaNotificationChannelRequest) (response *CreateGrafanaNotificationChannelResponse, err error) {
    return c.CreateGrafanaNotificationChannelWithContext(context.Background(), request)
}

// CreateGrafanaNotificationChannel
// 创建 Grafana 告警通道
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateGrafanaNotificationChannelWithContext(ctx context.Context, request *CreateGrafanaNotificationChannelRequest) (response *CreateGrafanaNotificationChannelResponse, err error) {
    if request == nil {
        request = NewCreateGrafanaNotificationChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateGrafanaNotificationChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGrafanaNotificationChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGrafanaNotificationChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePolicyGroupRequest() (request *CreatePolicyGroupRequest) {
    request = &CreatePolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePolicyGroup")
    
    
    return
}

func NewCreatePolicyGroupResponse() (response *CreatePolicyGroupResponse) {
    response = &CreatePolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePolicyGroup
// 增加策略组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePolicyGroup(request *CreatePolicyGroupRequest) (response *CreatePolicyGroupResponse, err error) {
    return c.CreatePolicyGroupWithContext(context.Background(), request)
}

// CreatePolicyGroup
// 增加策略组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePolicyGroupWithContext(ctx context.Context, request *CreatePolicyGroupRequest) (response *CreatePolicyGroupResponse, err error) {
    if request == nil {
        request = NewCreatePolicyGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePolicyGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePolicyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusAgentRequest() (request *CreatePrometheusAgentRequest) {
    request = &CreatePrometheusAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusAgent")
    
    
    return
}

func NewCreatePrometheusAgentResponse() (response *CreatePrometheusAgentResponse) {
    response = &CreatePrometheusAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusAgent
// 创建 Prometheus CVM Agent
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAgent(request *CreatePrometheusAgentRequest) (response *CreatePrometheusAgentResponse, err error) {
    return c.CreatePrometheusAgentWithContext(context.Background(), request)
}

// CreatePrometheusAgent
// 创建 Prometheus CVM Agent
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAgentWithContext(ctx context.Context, request *CreatePrometheusAgentRequest) (response *CreatePrometheusAgentResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusAgentResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusAlertGroupRequest() (request *CreatePrometheusAlertGroupRequest) {
    request = &CreatePrometheusAlertGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusAlertGroup")
    
    
    return
}

func NewCreatePrometheusAlertGroupResponse() (response *CreatePrometheusAlertGroupResponse) {
    response = &CreatePrometheusAlertGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusAlertGroup
// 创建Prometheus告警规则分组
//
// 
//
// 告警分组中可包含多条告警规则，分组内告警消息通过告警分组的通知模板发送。
//
// 支持单个告警分组下分别创建启用/禁用的告警规则。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreatePrometheusAlertGroup(request *CreatePrometheusAlertGroupRequest) (response *CreatePrometheusAlertGroupResponse, err error) {
    return c.CreatePrometheusAlertGroupWithContext(context.Background(), request)
}

// CreatePrometheusAlertGroup
// 创建Prometheus告警规则分组
//
// 
//
// 告警分组中可包含多条告警规则，分组内告警消息通过告警分组的通知模板发送。
//
// 支持单个告警分组下分别创建启用/禁用的告警规则。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreatePrometheusAlertGroupWithContext(ctx context.Context, request *CreatePrometheusAlertGroupRequest) (response *CreatePrometheusAlertGroupResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusAlertGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusAlertGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusAlertGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusAlertGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusAlertPolicyRequest() (request *CreatePrometheusAlertPolicyRequest) {
    request = &CreatePrometheusAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusAlertPolicy")
    
    
    return
}

func NewCreatePrometheusAlertPolicyResponse() (response *CreatePrometheusAlertPolicyResponse) {
    response = &CreatePrometheusAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusAlertPolicy
// 创建 Prometheus 告警策略(将逐步废弃，建议使用 CreatePrometheusAlertGroup 创建告警策略)
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAlertPolicy(request *CreatePrometheusAlertPolicyRequest) (response *CreatePrometheusAlertPolicyResponse, err error) {
    return c.CreatePrometheusAlertPolicyWithContext(context.Background(), request)
}

// CreatePrometheusAlertPolicy
// 创建 Prometheus 告警策略(将逐步废弃，建议使用 CreatePrometheusAlertGroup 创建告警策略)
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAlertPolicyWithContext(ctx context.Context, request *CreatePrometheusAlertPolicyRequest) (response *CreatePrometheusAlertPolicyResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusAlertPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusAlertPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusAlertPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusClusterAgentRequest() (request *CreatePrometheusClusterAgentRequest) {
    request = &CreatePrometheusClusterAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusClusterAgent")
    
    
    return
}

func NewCreatePrometheusClusterAgentResponse() (response *CreatePrometheusClusterAgentResponse) {
    response = &CreatePrometheusClusterAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusClusterAgent
// 与腾讯云可观测融合的2.0实例关联集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusClusterAgent(request *CreatePrometheusClusterAgentRequest) (response *CreatePrometheusClusterAgentResponse, err error) {
    return c.CreatePrometheusClusterAgentWithContext(context.Background(), request)
}

// CreatePrometheusClusterAgent
// 与腾讯云可观测融合的2.0实例关联集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusClusterAgentWithContext(ctx context.Context, request *CreatePrometheusClusterAgentRequest) (response *CreatePrometheusClusterAgentResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusClusterAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusClusterAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusClusterAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusClusterAgentResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusConfigRequest() (request *CreatePrometheusConfigRequest) {
    request = &CreatePrometheusConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusConfig")
    
    
    return
}

func NewCreatePrometheusConfigResponse() (response *CreatePrometheusConfigResponse) {
    response = &CreatePrometheusConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusConfig
// 创建prometheus配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  RESOURCEINUSE_RESOURCEEXISTALREADY = "ResourceInUse.ResourceExistAlready"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusConfig(request *CreatePrometheusConfigRequest) (response *CreatePrometheusConfigResponse, err error) {
    return c.CreatePrometheusConfigWithContext(context.Background(), request)
}

// CreatePrometheusConfig
// 创建prometheus配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  RESOURCEINUSE_RESOURCEEXISTALREADY = "ResourceInUse.ResourceExistAlready"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusConfigWithContext(ctx context.Context, request *CreatePrometheusConfigRequest) (response *CreatePrometheusConfigResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusGlobalNotificationRequest() (request *CreatePrometheusGlobalNotificationRequest) {
    request = &CreatePrometheusGlobalNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusGlobalNotification")
    
    
    return
}

func NewCreatePrometheusGlobalNotificationResponse() (response *CreatePrometheusGlobalNotificationResponse) {
    response = &CreatePrometheusGlobalNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusGlobalNotification
// 创建全局告警通知渠道。集群内创建的告警规则如果未配置告警通知渠道，默认走全局告警通知渠道（建议在控制台创建告警，集群内创建告警不易维护）
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusGlobalNotification(request *CreatePrometheusGlobalNotificationRequest) (response *CreatePrometheusGlobalNotificationResponse, err error) {
    return c.CreatePrometheusGlobalNotificationWithContext(context.Background(), request)
}

// CreatePrometheusGlobalNotification
// 创建全局告警通知渠道。集群内创建的告警规则如果未配置告警通知渠道，默认走全局告警通知渠道（建议在控制台创建告警，集群内创建告警不易维护）
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusGlobalNotificationWithContext(ctx context.Context, request *CreatePrometheusGlobalNotificationRequest) (response *CreatePrometheusGlobalNotificationResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusGlobalNotificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusGlobalNotification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusGlobalNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusGlobalNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusMultiTenantInstancePostPayModeRequest() (request *CreatePrometheusMultiTenantInstancePostPayModeRequest) {
    request = &CreatePrometheusMultiTenantInstancePostPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusMultiTenantInstancePostPayMode")
    
    
    return
}

func NewCreatePrometheusMultiTenantInstancePostPayModeResponse() (response *CreatePrometheusMultiTenantInstancePostPayModeResponse) {
    response = &CreatePrometheusMultiTenantInstancePostPayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusMultiTenantInstancePostPayMode
// 创建按量 Prometheus 实例，根据用量收费实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_CREATEINSTANCE = "FailedOperation.CreateInstance"
//  FAILEDOPERATION_CREATEINSTANCELIMITED = "FailedOperation.CreateInstanceLimited"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreatePrometheusMultiTenantInstancePostPayMode(request *CreatePrometheusMultiTenantInstancePostPayModeRequest) (response *CreatePrometheusMultiTenantInstancePostPayModeResponse, err error) {
    return c.CreatePrometheusMultiTenantInstancePostPayModeWithContext(context.Background(), request)
}

// CreatePrometheusMultiTenantInstancePostPayMode
// 创建按量 Prometheus 实例，根据用量收费实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_CREATEINSTANCE = "FailedOperation.CreateInstance"
//  FAILEDOPERATION_CREATEINSTANCELIMITED = "FailedOperation.CreateInstanceLimited"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreatePrometheusMultiTenantInstancePostPayModeWithContext(ctx context.Context, request *CreatePrometheusMultiTenantInstancePostPayModeRequest) (response *CreatePrometheusMultiTenantInstancePostPayModeResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusMultiTenantInstancePostPayModeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusMultiTenantInstancePostPayMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusMultiTenantInstancePostPayMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusMultiTenantInstancePostPayModeResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusRecordRuleYamlRequest() (request *CreatePrometheusRecordRuleYamlRequest) {
    request = &CreatePrometheusRecordRuleYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusRecordRuleYaml")
    
    
    return
}

func NewCreatePrometheusRecordRuleYamlResponse() (response *CreatePrometheusRecordRuleYamlResponse) {
    response = &CreatePrometheusRecordRuleYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusRecordRuleYaml
// 以Yaml的方式创建聚合规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreatePrometheusRecordRuleYaml(request *CreatePrometheusRecordRuleYamlRequest) (response *CreatePrometheusRecordRuleYamlResponse, err error) {
    return c.CreatePrometheusRecordRuleYamlWithContext(context.Background(), request)
}

// CreatePrometheusRecordRuleYaml
// 以Yaml的方式创建聚合规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreatePrometheusRecordRuleYamlWithContext(ctx context.Context, request *CreatePrometheusRecordRuleYamlRequest) (response *CreatePrometheusRecordRuleYamlResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusRecordRuleYamlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusRecordRuleYaml")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusRecordRuleYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusRecordRuleYamlResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusScrapeJobRequest() (request *CreatePrometheusScrapeJobRequest) {
    request = &CreatePrometheusScrapeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusScrapeJob")
    
    
    return
}

func NewCreatePrometheusScrapeJobResponse() (response *CreatePrometheusScrapeJobResponse) {
    response = &CreatePrometheusScrapeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusScrapeJob
// 创建 Prometheus Agent 抓取任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusScrapeJob(request *CreatePrometheusScrapeJobRequest) (response *CreatePrometheusScrapeJobResponse, err error) {
    return c.CreatePrometheusScrapeJobWithContext(context.Background(), request)
}

// CreatePrometheusScrapeJob
// 创建 Prometheus Agent 抓取任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusScrapeJobWithContext(ctx context.Context, request *CreatePrometheusScrapeJobRequest) (response *CreatePrometheusScrapeJobResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusScrapeJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusScrapeJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusScrapeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusScrapeJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusTempRequest() (request *CreatePrometheusTempRequest) {
    request = &CreatePrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusTemp")
    
    
    return
}

func NewCreatePrometheusTempResponse() (response *CreatePrometheusTempResponse) {
    response = &CreatePrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusTemp
// 创建一个云原生Prometheus模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CreatePrometheusTemp(request *CreatePrometheusTempRequest) (response *CreatePrometheusTempResponse, err error) {
    return c.CreatePrometheusTempWithContext(context.Background(), request)
}

// CreatePrometheusTemp
// 创建一个云原生Prometheus模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CreatePrometheusTempWithContext(ctx context.Context, request *CreatePrometheusTempRequest) (response *CreatePrometheusTempResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusTempRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusTemp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordingRuleRequest() (request *CreateRecordingRuleRequest) {
    request = &CreateRecordingRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateRecordingRule")
    
    
    return
}

func NewCreateRecordingRuleResponse() (response *CreateRecordingRuleResponse) {
    response = &CreateRecordingRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRecordingRule
// 创建 Prometheus 的预聚合规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRecordingRule(request *CreateRecordingRuleRequest) (response *CreateRecordingRuleResponse, err error) {
    return c.CreateRecordingRuleWithContext(context.Background(), request)
}

// CreateRecordingRule
// 创建 Prometheus 的预聚合规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRecordingRuleWithContext(ctx context.Context, request *CreateRecordingRuleRequest) (response *CreateRecordingRuleResponse, err error) {
    if request == nil {
        request = NewCreateRecordingRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateRecordingRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRecordingRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRecordingRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSSOAccountRequest() (request *CreateSSOAccountRequest) {
    request = &CreateSSOAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateSSOAccount")
    
    
    return
}

func NewCreateSSOAccountResponse() (response *CreateSSOAccountResponse) {
    response = &CreateSSOAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSSOAccount
// Grafana实例授权其他腾讯云用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSSOAccount(request *CreateSSOAccountRequest) (response *CreateSSOAccountResponse, err error) {
    return c.CreateSSOAccountWithContext(context.Background(), request)
}

// CreateSSOAccount
// Grafana实例授权其他腾讯云用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSSOAccountWithContext(ctx context.Context, request *CreateSSOAccountRequest) (response *CreateSSOAccountResponse, err error) {
    if request == nil {
        request = NewCreateSSOAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateSSOAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSSOAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSSOAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceDiscoveryRequest() (request *CreateServiceDiscoveryRequest) {
    request = &CreateServiceDiscoveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateServiceDiscovery")
    
    
    return
}

func NewCreateServiceDiscoveryResponse() (response *CreateServiceDiscoveryResponse) {
    response = &CreateServiceDiscoveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateServiceDiscovery
// 在腾讯云容器服务下创建 Prometheus 服务发现。
//
// <p>注意：前提条件，已经通过 Prometheus 控制台集成了对应的腾讯云容器服务，具体请参考
//
// <a href="https://cloud.tencent.com/document/product/248/48859" target="_blank">Agent 安装</a>。</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKECLIENTAUTHFAIL = "FailedOperation.TKEClientAuthFail"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateServiceDiscovery(request *CreateServiceDiscoveryRequest) (response *CreateServiceDiscoveryResponse, err error) {
    return c.CreateServiceDiscoveryWithContext(context.Background(), request)
}

// CreateServiceDiscovery
// 在腾讯云容器服务下创建 Prometheus 服务发现。
//
// <p>注意：前提条件，已经通过 Prometheus 控制台集成了对应的腾讯云容器服务，具体请参考
//
// <a href="https://cloud.tencent.com/document/product/248/48859" target="_blank">Agent 安装</a>。</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKECLIENTAUTHFAIL = "FailedOperation.TKEClientAuthFail"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateServiceDiscoveryWithContext(ctx context.Context, request *CreateServiceDiscoveryRequest) (response *CreateServiceDiscoveryResponse, err error) {
    if request == nil {
        request = NewCreateServiceDiscoveryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateServiceDiscovery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateServiceDiscovery require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServiceDiscoveryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmNoticesRequest() (request *DeleteAlarmNoticesRequest) {
    request = &DeleteAlarmNoticesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlarmNotices")
    
    
    return
}

func NewDeleteAlarmNoticesResponse() (response *DeleteAlarmNoticesResponse) {
    response = &DeleteAlarmNoticesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlarmNotices
// 删除告警通知模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmNotices(request *DeleteAlarmNoticesRequest) (response *DeleteAlarmNoticesResponse, err error) {
    return c.DeleteAlarmNoticesWithContext(context.Background(), request)
}

// DeleteAlarmNotices
// 删除告警通知模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmNoticesWithContext(ctx context.Context, request *DeleteAlarmNoticesRequest) (response *DeleteAlarmNoticesResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmNoticesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteAlarmNotices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarmNotices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmNoticesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmPolicyRequest() (request *DeleteAlarmPolicyRequest) {
    request = &DeleteAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlarmPolicy")
    
    
    return
}

func NewDeleteAlarmPolicyResponse() (response *DeleteAlarmPolicyResponse) {
    response = &DeleteAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlarmPolicy
// 删除告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmPolicy(request *DeleteAlarmPolicyRequest) (response *DeleteAlarmPolicyResponse, err error) {
    return c.DeleteAlarmPolicyWithContext(context.Background(), request)
}

// DeleteAlarmPolicy
// 删除告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmPolicyWithContext(ctx context.Context, request *DeleteAlarmPolicyRequest) (response *DeleteAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteAlarmPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarmPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmShieldsRequest() (request *DeleteAlarmShieldsRequest) {
    request = &DeleteAlarmShieldsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlarmShields")
    
    
    return
}

func NewDeleteAlarmShieldsResponse() (response *DeleteAlarmShieldsResponse) {
    response = &DeleteAlarmShieldsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlarmShields
// 删除告警屏蔽规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmShields(request *DeleteAlarmShieldsRequest) (response *DeleteAlarmShieldsResponse, err error) {
    return c.DeleteAlarmShieldsWithContext(context.Background(), request)
}

// DeleteAlarmShields
// 删除告警屏蔽规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmShieldsWithContext(ctx context.Context, request *DeleteAlarmShieldsRequest) (response *DeleteAlarmShieldsResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmShieldsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteAlarmShields")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarmShields require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmShieldsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlertRulesRequest() (request *DeleteAlertRulesRequest) {
    request = &DeleteAlertRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlertRules")
    
    
    return
}

func NewDeleteAlertRulesResponse() (response *DeleteAlertRulesResponse) {
    response = &DeleteAlertRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlertRules
// 批量删除 Prometheus 报警规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAlertRules(request *DeleteAlertRulesRequest) (response *DeleteAlertRulesResponse, err error) {
    return c.DeleteAlertRulesWithContext(context.Background(), request)
}

// DeleteAlertRules
// 批量删除 Prometheus 报警规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAlertRulesWithContext(ctx context.Context, request *DeleteAlertRulesRequest) (response *DeleteAlertRulesResponse, err error) {
    if request == nil {
        request = NewDeleteAlertRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteAlertRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlertRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlertRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteExporterIntegrationRequest() (request *DeleteExporterIntegrationRequest) {
    request = &DeleteExporterIntegrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteExporterIntegration")
    
    
    return
}

func NewDeleteExporterIntegrationResponse() (response *DeleteExporterIntegrationResponse) {
    response = &DeleteExporterIntegrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteExporterIntegration
// 删除集成中心 exporter 集成
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteExporterIntegration(request *DeleteExporterIntegrationRequest) (response *DeleteExporterIntegrationResponse, err error) {
    return c.DeleteExporterIntegrationWithContext(context.Background(), request)
}

// DeleteExporterIntegration
// 删除集成中心 exporter 集成
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteExporterIntegrationWithContext(ctx context.Context, request *DeleteExporterIntegrationRequest) (response *DeleteExporterIntegrationResponse, err error) {
    if request == nil {
        request = NewDeleteExporterIntegrationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteExporterIntegration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteExporterIntegration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteExporterIntegrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGrafanaInstanceRequest() (request *DeleteGrafanaInstanceRequest) {
    request = &DeleteGrafanaInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteGrafanaInstance")
    
    
    return
}

func NewDeleteGrafanaInstanceResponse() (response *DeleteGrafanaInstanceResponse) {
    response = &DeleteGrafanaInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGrafanaInstance
// 本接口（DeleteGrafanaInstance）用于 Grafana 包年包月实例的退费，调用后实例处于停服状态，不可使用，7天后自动销毁。
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaInstance(request *DeleteGrafanaInstanceRequest) (response *DeleteGrafanaInstanceResponse, err error) {
    return c.DeleteGrafanaInstanceWithContext(context.Background(), request)
}

// DeleteGrafanaInstance
// 本接口（DeleteGrafanaInstance）用于 Grafana 包年包月实例的退费，调用后实例处于停服状态，不可使用，7天后自动销毁。
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaInstanceWithContext(ctx context.Context, request *DeleteGrafanaInstanceRequest) (response *DeleteGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteGrafanaInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteGrafanaInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGrafanaIntegrationRequest() (request *DeleteGrafanaIntegrationRequest) {
    request = &DeleteGrafanaIntegrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteGrafanaIntegration")
    
    
    return
}

func NewDeleteGrafanaIntegrationResponse() (response *DeleteGrafanaIntegrationResponse) {
    response = &DeleteGrafanaIntegrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGrafanaIntegration
// 删除 Grafana 集成配置
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaIntegration(request *DeleteGrafanaIntegrationRequest) (response *DeleteGrafanaIntegrationResponse, err error) {
    return c.DeleteGrafanaIntegrationWithContext(context.Background(), request)
}

// DeleteGrafanaIntegration
// 删除 Grafana 集成配置
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaIntegrationWithContext(ctx context.Context, request *DeleteGrafanaIntegrationRequest) (response *DeleteGrafanaIntegrationResponse, err error) {
    if request == nil {
        request = NewDeleteGrafanaIntegrationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteGrafanaIntegration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGrafanaIntegration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGrafanaIntegrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGrafanaNotificationChannelRequest() (request *DeleteGrafanaNotificationChannelRequest) {
    request = &DeleteGrafanaNotificationChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteGrafanaNotificationChannel")
    
    
    return
}

func NewDeleteGrafanaNotificationChannelResponse() (response *DeleteGrafanaNotificationChannelResponse) {
    response = &DeleteGrafanaNotificationChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGrafanaNotificationChannel
// 删除 Grafana 告警通道
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaNotificationChannel(request *DeleteGrafanaNotificationChannelRequest) (response *DeleteGrafanaNotificationChannelResponse, err error) {
    return c.DeleteGrafanaNotificationChannelWithContext(context.Background(), request)
}

// DeleteGrafanaNotificationChannel
// 删除 Grafana 告警通道
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaNotificationChannelWithContext(ctx context.Context, request *DeleteGrafanaNotificationChannelRequest) (response *DeleteGrafanaNotificationChannelResponse, err error) {
    if request == nil {
        request = NewDeleteGrafanaNotificationChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteGrafanaNotificationChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGrafanaNotificationChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGrafanaNotificationChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePolicyGroupRequest() (request *DeletePolicyGroupRequest) {
    request = &DeletePolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePolicyGroup")
    
    
    return
}

func NewDeletePolicyGroupResponse() (response *DeletePolicyGroupResponse) {
    response = &DeletePolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePolicyGroup
// 删除告警策略组
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeletePolicyGroup(request *DeletePolicyGroupRequest) (response *DeletePolicyGroupResponse, err error) {
    return c.DeletePolicyGroupWithContext(context.Background(), request)
}

// DeletePolicyGroup
// 删除告警策略组
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeletePolicyGroupWithContext(ctx context.Context, request *DeletePolicyGroupRequest) (response *DeletePolicyGroupResponse, err error) {
    if request == nil {
        request = NewDeletePolicyGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePolicyGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePolicyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusAlertGroupsRequest() (request *DeletePrometheusAlertGroupsRequest) {
    request = &DeletePrometheusAlertGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusAlertGroups")
    
    
    return
}

func NewDeletePrometheusAlertGroupsResponse() (response *DeletePrometheusAlertGroupsResponse) {
    response = &DeletePrometheusAlertGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusAlertGroups
// 删除Prometheus告警规则分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeletePrometheusAlertGroups(request *DeletePrometheusAlertGroupsRequest) (response *DeletePrometheusAlertGroupsResponse, err error) {
    return c.DeletePrometheusAlertGroupsWithContext(context.Background(), request)
}

// DeletePrometheusAlertGroups
// 删除Prometheus告警规则分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeletePrometheusAlertGroupsWithContext(ctx context.Context, request *DeletePrometheusAlertGroupsRequest) (response *DeletePrometheusAlertGroupsResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusAlertGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusAlertGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusAlertGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusAlertGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusAlertPolicyRequest() (request *DeletePrometheusAlertPolicyRequest) {
    request = &DeletePrometheusAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusAlertPolicy")
    
    
    return
}

func NewDeletePrometheusAlertPolicyResponse() (response *DeletePrometheusAlertPolicyResponse) {
    response = &DeletePrometheusAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusAlertPolicy
// 删除2.0实例告警策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusAlertPolicy(request *DeletePrometheusAlertPolicyRequest) (response *DeletePrometheusAlertPolicyResponse, err error) {
    return c.DeletePrometheusAlertPolicyWithContext(context.Background(), request)
}

// DeletePrometheusAlertPolicy
// 删除2.0实例告警策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusAlertPolicyWithContext(ctx context.Context, request *DeletePrometheusAlertPolicyRequest) (response *DeletePrometheusAlertPolicyResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusAlertPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusAlertPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusAlertPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusClusterAgentRequest() (request *DeletePrometheusClusterAgentRequest) {
    request = &DeletePrometheusClusterAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusClusterAgent")
    
    
    return
}

func NewDeletePrometheusClusterAgentResponse() (response *DeletePrometheusClusterAgentResponse) {
    response = &DeletePrometheusClusterAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusClusterAgent
// 解除TMP实例的集群关联
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusClusterAgent(request *DeletePrometheusClusterAgentRequest) (response *DeletePrometheusClusterAgentResponse, err error) {
    return c.DeletePrometheusClusterAgentWithContext(context.Background(), request)
}

// DeletePrometheusClusterAgent
// 解除TMP实例的集群关联
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusClusterAgentWithContext(ctx context.Context, request *DeletePrometheusClusterAgentRequest) (response *DeletePrometheusClusterAgentResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusClusterAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusClusterAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusClusterAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusClusterAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusConfigRequest() (request *DeletePrometheusConfigRequest) {
    request = &DeletePrometheusConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusConfig")
    
    
    return
}

func NewDeletePrometheusConfigResponse() (response *DeletePrometheusConfigResponse) {
    response = &DeletePrometheusConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusConfig
// 删除Prometheus配置，如果目标不存在，将返回成功
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) DeletePrometheusConfig(request *DeletePrometheusConfigRequest) (response *DeletePrometheusConfigResponse, err error) {
    return c.DeletePrometheusConfigWithContext(context.Background(), request)
}

// DeletePrometheusConfig
// 删除Prometheus配置，如果目标不存在，将返回成功
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) DeletePrometheusConfigWithContext(ctx context.Context, request *DeletePrometheusConfigRequest) (response *DeletePrometheusConfigResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusRecordRuleYamlRequest() (request *DeletePrometheusRecordRuleYamlRequest) {
    request = &DeletePrometheusRecordRuleYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusRecordRuleYaml")
    
    
    return
}

func NewDeletePrometheusRecordRuleYamlResponse() (response *DeletePrometheusRecordRuleYamlResponse) {
    response = &DeletePrometheusRecordRuleYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusRecordRuleYaml
// 删除聚合实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusRecordRuleYaml(request *DeletePrometheusRecordRuleYamlRequest) (response *DeletePrometheusRecordRuleYamlResponse, err error) {
    return c.DeletePrometheusRecordRuleYamlWithContext(context.Background(), request)
}

// DeletePrometheusRecordRuleYaml
// 删除聚合实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusRecordRuleYamlWithContext(ctx context.Context, request *DeletePrometheusRecordRuleYamlRequest) (response *DeletePrometheusRecordRuleYamlResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusRecordRuleYamlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusRecordRuleYaml")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusRecordRuleYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusRecordRuleYamlResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusScrapeJobsRequest() (request *DeletePrometheusScrapeJobsRequest) {
    request = &DeletePrometheusScrapeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusScrapeJobs")
    
    
    return
}

func NewDeletePrometheusScrapeJobsResponse() (response *DeletePrometheusScrapeJobsResponse) {
    response = &DeletePrometheusScrapeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusScrapeJobs
// 删除 Prometheus Agent 抓取任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePrometheusScrapeJobs(request *DeletePrometheusScrapeJobsRequest) (response *DeletePrometheusScrapeJobsResponse, err error) {
    return c.DeletePrometheusScrapeJobsWithContext(context.Background(), request)
}

// DeletePrometheusScrapeJobs
// 删除 Prometheus Agent 抓取任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePrometheusScrapeJobsWithContext(ctx context.Context, request *DeletePrometheusScrapeJobsRequest) (response *DeletePrometheusScrapeJobsResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusScrapeJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusScrapeJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusScrapeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusScrapeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusTempRequest() (request *DeletePrometheusTempRequest) {
    request = &DeletePrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusTemp")
    
    
    return
}

func NewDeletePrometheusTempResponse() (response *DeletePrometheusTempResponse) {
    response = &DeletePrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusTemp
// 删除一个云原生Prometheus配置模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePrometheusTemp(request *DeletePrometheusTempRequest) (response *DeletePrometheusTempResponse, err error) {
    return c.DeletePrometheusTempWithContext(context.Background(), request)
}

// DeletePrometheusTemp
// 删除一个云原生Prometheus配置模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePrometheusTempWithContext(ctx context.Context, request *DeletePrometheusTempRequest) (response *DeletePrometheusTempResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusTempRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusTemp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusTempSyncRequest() (request *DeletePrometheusTempSyncRequest) {
    request = &DeletePrometheusTempSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusTempSync")
    
    
    return
}

func NewDeletePrometheusTempSyncResponse() (response *DeletePrometheusTempSyncResponse) {
    response = &DeletePrometheusTempSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusTempSync
// 解除模板同步，这将会删除目标中该模板所生产的配置，针对V2版本实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeletePrometheusTempSync(request *DeletePrometheusTempSyncRequest) (response *DeletePrometheusTempSyncResponse, err error) {
    return c.DeletePrometheusTempSyncWithContext(context.Background(), request)
}

// DeletePrometheusTempSync
// 解除模板同步，这将会删除目标中该模板所生产的配置，针对V2版本实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeletePrometheusTempSyncWithContext(ctx context.Context, request *DeletePrometheusTempSyncRequest) (response *DeletePrometheusTempSyncResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusTempSyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusTempSync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusTempSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusTempSyncResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordingRulesRequest() (request *DeleteRecordingRulesRequest) {
    request = &DeleteRecordingRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteRecordingRules")
    
    
    return
}

func NewDeleteRecordingRulesResponse() (response *DeleteRecordingRulesResponse) {
    response = &DeleteRecordingRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecordingRules
// 批量删除 Prometheus 预聚合规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRecordingRules(request *DeleteRecordingRulesRequest) (response *DeleteRecordingRulesResponse, err error) {
    return c.DeleteRecordingRulesWithContext(context.Background(), request)
}

// DeleteRecordingRules
// 批量删除 Prometheus 预聚合规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRecordingRulesWithContext(ctx context.Context, request *DeleteRecordingRulesRequest) (response *DeleteRecordingRulesResponse, err error) {
    if request == nil {
        request = NewDeleteRecordingRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteRecordingRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordingRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordingRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSSOAccountRequest() (request *DeleteSSOAccountRequest) {
    request = &DeleteSSOAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteSSOAccount")
    
    
    return
}

func NewDeleteSSOAccountResponse() (response *DeleteSSOAccountResponse) {
    response = &DeleteSSOAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSSOAccount
// Grafana可视化服务 删除授权用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteSSOAccount(request *DeleteSSOAccountRequest) (response *DeleteSSOAccountResponse, err error) {
    return c.DeleteSSOAccountWithContext(context.Background(), request)
}

// DeleteSSOAccount
// Grafana可视化服务 删除授权用户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteSSOAccountWithContext(ctx context.Context, request *DeleteSSOAccountRequest) (response *DeleteSSOAccountResponse, err error) {
    if request == nil {
        request = NewDeleteSSOAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteSSOAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSSOAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSSOAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceDiscoveryRequest() (request *DeleteServiceDiscoveryRequest) {
    request = &DeleteServiceDiscoveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteServiceDiscovery")
    
    
    return
}

func NewDeleteServiceDiscoveryResponse() (response *DeleteServiceDiscoveryResponse) {
    response = &DeleteServiceDiscoveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteServiceDiscovery
// 删除在腾讯云容器服务下创建的 Prometheus 服务发现。
//
// <p>注意：前提条件，已经通过 Prometheus 控制台集成了对应的腾讯云容器服务，具体请参考
//
// <a href="https://cloud.tencent.com/document/product/248/48859" target="_blank">Agent 安装</a>。</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteServiceDiscovery(request *DeleteServiceDiscoveryRequest) (response *DeleteServiceDiscoveryResponse, err error) {
    return c.DeleteServiceDiscoveryWithContext(context.Background(), request)
}

// DeleteServiceDiscovery
// 删除在腾讯云容器服务下创建的 Prometheus 服务发现。
//
// <p>注意：前提条件，已经通过 Prometheus 控制台集成了对应的腾讯云容器服务，具体请参考
//
// <a href="https://cloud.tencent.com/document/product/248/48859" target="_blank">Agent 安装</a>。</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteServiceDiscoveryWithContext(ctx context.Context, request *DeleteServiceDiscoveryRequest) (response *DeleteServiceDiscoveryResponse, err error) {
    if request == nil {
        request = NewDeleteServiceDiscoveryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteServiceDiscovery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServiceDiscovery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServiceDiscoveryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccidentEventListRequest() (request *DescribeAccidentEventListRequest) {
    request = &DescribeAccidentEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAccidentEventList")
    
    
    return
}

func NewDescribeAccidentEventListResponse() (response *DescribeAccidentEventListResponse) {
    response = &DescribeAccidentEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccidentEventList
// 获取平台事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccidentEventList(request *DescribeAccidentEventListRequest) (response *DescribeAccidentEventListResponse, err error) {
    return c.DescribeAccidentEventListWithContext(context.Background(), request)
}

// DescribeAccidentEventList
// 获取平台事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccidentEventListWithContext(ctx context.Context, request *DescribeAccidentEventListRequest) (response *DescribeAccidentEventListResponse, err error) {
    if request == nil {
        request = NewDescribeAccidentEventListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAccidentEventList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccidentEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccidentEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmEventsRequest() (request *DescribeAlarmEventsRequest) {
    request = &DescribeAlarmEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmEvents")
    
    
    return
}

func NewDescribeAlarmEventsResponse() (response *DescribeAlarmEventsResponse) {
    response = &DescribeAlarmEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmEvents
// 查询告警事件列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmEvents(request *DescribeAlarmEventsRequest) (response *DescribeAlarmEventsResponse, err error) {
    return c.DescribeAlarmEventsWithContext(context.Background(), request)
}

// DescribeAlarmEvents
// 查询告警事件列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmEventsWithContext(ctx context.Context, request *DescribeAlarmEventsRequest) (response *DescribeAlarmEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmEventsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmEvents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmHistoriesRequest() (request *DescribeAlarmHistoriesRequest) {
    request = &DescribeAlarmHistoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmHistories")
    
    
    return
}

func NewDescribeAlarmHistoriesResponse() (response *DescribeAlarmHistoriesResponse) {
    response = &DescribeAlarmHistoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmHistories
// 查询告警历史
//
// 
//
// 请注意，**如果使用子用户进行告警历史的查询，只能查询到被授权项目下的告警历史**，或不区分项目的产品的告警历史。如何对子账户授予项目的权限，请参考 [访问管理-项目与标签](https://cloud.tencent.com/document/product/598/32738)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmHistories(request *DescribeAlarmHistoriesRequest) (response *DescribeAlarmHistoriesResponse, err error) {
    return c.DescribeAlarmHistoriesWithContext(context.Background(), request)
}

// DescribeAlarmHistories
// 查询告警历史
//
// 
//
// 请注意，**如果使用子用户进行告警历史的查询，只能查询到被授权项目下的告警历史**，或不区分项目的产品的告警历史。如何对子账户授予项目的权限，请参考 [访问管理-项目与标签](https://cloud.tencent.com/document/product/598/32738)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmHistoriesWithContext(ctx context.Context, request *DescribeAlarmHistoriesRequest) (response *DescribeAlarmHistoriesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmHistoriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmHistories")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmHistories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmHistoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmMetricsRequest() (request *DescribeAlarmMetricsRequest) {
    request = &DescribeAlarmMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmMetrics")
    
    
    return
}

func NewDescribeAlarmMetricsResponse() (response *DescribeAlarmMetricsResponse) {
    response = &DescribeAlarmMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmMetrics
// 查询告警指标列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILEDWITHCORRUPTIONDATA = "FailedOperation.DoTRPCTransferFailedWithCorruptionData"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmMetrics(request *DescribeAlarmMetricsRequest) (response *DescribeAlarmMetricsResponse, err error) {
    return c.DescribeAlarmMetricsWithContext(context.Background(), request)
}

// DescribeAlarmMetrics
// 查询告警指标列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILEDWITHCORRUPTIONDATA = "FailedOperation.DoTRPCTransferFailedWithCorruptionData"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmMetricsWithContext(ctx context.Context, request *DescribeAlarmMetricsRequest) (response *DescribeAlarmMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmMetricsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmMetrics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmNoticeRequest() (request *DescribeAlarmNoticeRequest) {
    request = &DescribeAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmNotice")
    
    
    return
}

func NewDescribeAlarmNoticeResponse() (response *DescribeAlarmNoticeResponse) {
    response = &DescribeAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmNotice
// 查询单个通知模板的详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNotice(request *DescribeAlarmNoticeRequest) (response *DescribeAlarmNoticeResponse, err error) {
    return c.DescribeAlarmNoticeWithContext(context.Background(), request)
}

// DescribeAlarmNotice
// 查询单个通知模板的详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticeWithContext(ctx context.Context, request *DescribeAlarmNoticeRequest) (response *DescribeAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmNotice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmNoticeCallbacksRequest() (request *DescribeAlarmNoticeCallbacksRequest) {
    request = &DescribeAlarmNoticeCallbacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmNoticeCallbacks")
    
    
    return
}

func NewDescribeAlarmNoticeCallbacksResponse() (response *DescribeAlarmNoticeCallbacksResponse) {
    response = &DescribeAlarmNoticeCallbacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmNoticeCallbacks
// 获取告警通知模板所有回调URL
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticeCallbacks(request *DescribeAlarmNoticeCallbacksRequest) (response *DescribeAlarmNoticeCallbacksResponse, err error) {
    return c.DescribeAlarmNoticeCallbacksWithContext(context.Background(), request)
}

// DescribeAlarmNoticeCallbacks
// 获取告警通知模板所有回调URL
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticeCallbacksWithContext(ctx context.Context, request *DescribeAlarmNoticeCallbacksRequest) (response *DescribeAlarmNoticeCallbacksResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticeCallbacksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmNoticeCallbacks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmNoticeCallbacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmNoticeCallbacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmNoticesRequest() (request *DescribeAlarmNoticesRequest) {
    request = &DescribeAlarmNoticesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmNotices")
    
    
    return
}

func NewDescribeAlarmNoticesResponse() (response *DescribeAlarmNoticesResponse) {
    response = &DescribeAlarmNoticesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmNotices
// 查询通知模板列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNotices(request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
    return c.DescribeAlarmNoticesWithContext(context.Background(), request)
}

// DescribeAlarmNotices
// 查询通知模板列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticesWithContext(ctx context.Context, request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmNotices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmNotices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmNoticesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmPoliciesRequest() (request *DescribeAlarmPoliciesRequest) {
    request = &DescribeAlarmPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmPolicies")
    
    
    return
}

func NewDescribeAlarmPoliciesResponse() (response *DescribeAlarmPoliciesResponse) {
    response = &DescribeAlarmPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmPolicies
// 查询告警策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPolicies(request *DescribeAlarmPoliciesRequest) (response *DescribeAlarmPoliciesResponse, err error) {
    return c.DescribeAlarmPoliciesWithContext(context.Background(), request)
}

// DescribeAlarmPolicies
// 查询告警策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPoliciesWithContext(ctx context.Context, request *DescribeAlarmPoliciesRequest) (response *DescribeAlarmPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmPolicyRequest() (request *DescribeAlarmPolicyRequest) {
    request = &DescribeAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmPolicy")
    
    
    return
}

func NewDescribeAlarmPolicyResponse() (response *DescribeAlarmPolicyResponse) {
    response = &DescribeAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmPolicy
// 获取单个告警策略详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPolicy(request *DescribeAlarmPolicyRequest) (response *DescribeAlarmPolicyResponse, err error) {
    return c.DescribeAlarmPolicyWithContext(context.Background(), request)
}

// DescribeAlarmPolicy
// 获取单个告警策略详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPolicyWithContext(ctx context.Context, request *DescribeAlarmPolicyRequest) (response *DescribeAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmSmsQuotaRequest() (request *DescribeAlarmSmsQuotaRequest) {
    request = &DescribeAlarmSmsQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmSmsQuota")
    
    
    return
}

func NewDescribeAlarmSmsQuotaResponse() (response *DescribeAlarmSmsQuotaResponse) {
    response = &DescribeAlarmSmsQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmSmsQuota
// 获取告警短信配额
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DRUIDTABLENOTFOUND = "FailedOperation.DruidTableNotFound"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmSmsQuota(request *DescribeAlarmSmsQuotaRequest) (response *DescribeAlarmSmsQuotaResponse, err error) {
    return c.DescribeAlarmSmsQuotaWithContext(context.Background(), request)
}

// DescribeAlarmSmsQuota
// 获取告警短信配额
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DRUIDTABLENOTFOUND = "FailedOperation.DruidTableNotFound"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmSmsQuotaWithContext(ctx context.Context, request *DescribeAlarmSmsQuotaRequest) (response *DescribeAlarmSmsQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmSmsQuotaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmSmsQuota")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmSmsQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmSmsQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertRulesRequest() (request *DescribeAlertRulesRequest) {
    request = &DescribeAlertRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlertRules")
    
    
    return
}

func NewDescribeAlertRulesResponse() (response *DescribeAlertRulesResponse) {
    response = &DescribeAlertRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlertRules
// Prometheus 报警规则查询接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAlertRules(request *DescribeAlertRulesRequest) (response *DescribeAlertRulesResponse, err error) {
    return c.DescribeAlertRulesWithContext(context.Background(), request)
}

// DescribeAlertRules
// Prometheus 报警规则查询接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAlertRulesWithContext(ctx context.Context, request *DescribeAlertRulesRequest) (response *DescribeAlertRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAlertRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlertRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlertRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlertRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllNamespacesRequest() (request *DescribeAllNamespacesRequest) {
    request = &DescribeAllNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAllNamespaces")
    
    
    return
}

func NewDescribeAllNamespacesResponse() (response *DescribeAllNamespacesResponse) {
    response = &DescribeAllNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllNamespaces
// 查询所有名字空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAllNamespaces(request *DescribeAllNamespacesRequest) (response *DescribeAllNamespacesResponse, err error) {
    return c.DescribeAllNamespacesWithContext(context.Background(), request)
}

// DescribeAllNamespaces
// 查询所有名字空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAllNamespacesWithContext(ctx context.Context, request *DescribeAllNamespacesRequest) (response *DescribeAllNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeAllNamespacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAllNamespaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaseMetricsRequest() (request *DescribeBaseMetricsRequest) {
    request = &DescribeBaseMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBaseMetrics")
    
    
    return
}

func NewDescribeBaseMetricsResponse() (response *DescribeBaseMetricsResponse) {
    response = &DescribeBaseMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBaseMetrics
// 获取基础指标属性
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBaseMetrics(request *DescribeBaseMetricsRequest) (response *DescribeBaseMetricsResponse, err error) {
    return c.DescribeBaseMetricsWithContext(context.Background(), request)
}

// DescribeBaseMetrics
// 获取基础指标属性
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBaseMetricsWithContext(ctx context.Context, request *DescribeBaseMetricsRequest) (response *DescribeBaseMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeBaseMetricsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeBaseMetrics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaseMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaseMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicAlarmListRequest() (request *DescribeBasicAlarmListRequest) {
    request = &DescribeBasicAlarmListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBasicAlarmList")
    
    
    return
}

func NewDescribeBasicAlarmListResponse() (response *DescribeBasicAlarmListResponse) {
    response = &DescribeBasicAlarmListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBasicAlarmList
// 获取基础告警列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBasicAlarmList(request *DescribeBasicAlarmListRequest) (response *DescribeBasicAlarmListResponse, err error) {
    return c.DescribeBasicAlarmListWithContext(context.Background(), request)
}

// DescribeBasicAlarmList
// 获取基础告警列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBasicAlarmListWithContext(ctx context.Context, request *DescribeBasicAlarmListRequest) (response *DescribeBasicAlarmListResponse, err error) {
    if request == nil {
        request = NewDescribeBasicAlarmListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeBasicAlarmList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBasicAlarmList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBasicAlarmListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindingPolicyObjectListRequest() (request *DescribeBindingPolicyObjectListRequest) {
    request = &DescribeBindingPolicyObjectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBindingPolicyObjectList")
    
    
    return
}

func NewDescribeBindingPolicyObjectListResponse() (response *DescribeBindingPolicyObjectListResponse) {
    response = &DescribeBindingPolicyObjectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBindingPolicyObjectList
// 获取已绑定对象列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DRUIDTABLENOTFOUND = "FailedOperation.DruidTableNotFound"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindingPolicyObjectList(request *DescribeBindingPolicyObjectListRequest) (response *DescribeBindingPolicyObjectListResponse, err error) {
    return c.DescribeBindingPolicyObjectListWithContext(context.Background(), request)
}

// DescribeBindingPolicyObjectList
// 获取已绑定对象列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DRUIDTABLENOTFOUND = "FailedOperation.DruidTableNotFound"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindingPolicyObjectListWithContext(ctx context.Context, request *DescribeBindingPolicyObjectListRequest) (response *DescribeBindingPolicyObjectListResponse, err error) {
    if request == nil {
        request = NewDescribeBindingPolicyObjectListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeBindingPolicyObjectList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBindingPolicyObjectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBindingPolicyObjectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAgentCreatingProgressRequest() (request *DescribeClusterAgentCreatingProgressRequest) {
    request = &DescribeClusterAgentCreatingProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeClusterAgentCreatingProgress")
    
    
    return
}

func NewDescribeClusterAgentCreatingProgressResponse() (response *DescribeClusterAgentCreatingProgressResponse) {
    response = &DescribeClusterAgentCreatingProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterAgentCreatingProgress
// 获取prom实例中集群详细的关联状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeClusterAgentCreatingProgress(request *DescribeClusterAgentCreatingProgressRequest) (response *DescribeClusterAgentCreatingProgressResponse, err error) {
    return c.DescribeClusterAgentCreatingProgressWithContext(context.Background(), request)
}

// DescribeClusterAgentCreatingProgress
// 获取prom实例中集群详细的关联状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeClusterAgentCreatingProgressWithContext(ctx context.Context, request *DescribeClusterAgentCreatingProgressRequest) (response *DescribeClusterAgentCreatingProgressResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAgentCreatingProgressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeClusterAgentCreatingProgress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterAgentCreatingProgress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterAgentCreatingProgressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConditionsTemplateListRequest() (request *DescribeConditionsTemplateListRequest) {
    request = &DescribeConditionsTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeConditionsTemplateList")
    
    
    return
}

func NewDescribeConditionsTemplateListResponse() (response *DescribeConditionsTemplateListResponse) {
    response = &DescribeConditionsTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConditionsTemplateList
// 获取条件模板列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConditionsTemplateList(request *DescribeConditionsTemplateListRequest) (response *DescribeConditionsTemplateListResponse, err error) {
    return c.DescribeConditionsTemplateListWithContext(context.Background(), request)
}

// DescribeConditionsTemplateList
// 获取条件模板列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConditionsTemplateListWithContext(ctx context.Context, request *DescribeConditionsTemplateListRequest) (response *DescribeConditionsTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeConditionsTemplateListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeConditionsTemplateList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConditionsTemplateList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConditionsTemplateListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDNSConfigRequest() (request *DescribeDNSConfigRequest) {
    request = &DescribeDNSConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeDNSConfig")
    
    
    return
}

func NewDescribeDNSConfigResponse() (response *DescribeDNSConfigResponse) {
    response = &DescribeDNSConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDNSConfig
// 列出 Grafana DNS 配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDNSConfig(request *DescribeDNSConfigRequest) (response *DescribeDNSConfigResponse, err error) {
    return c.DescribeDNSConfigWithContext(context.Background(), request)
}

// DescribeDNSConfig
// 列出 Grafana DNS 配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDNSConfigWithContext(ctx context.Context, request *DescribeDNSConfigRequest) (response *DescribeDNSConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDNSConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeDNSConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDNSConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDNSConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExporterIntegrationsRequest() (request *DescribeExporterIntegrationsRequest) {
    request = &DescribeExporterIntegrationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeExporterIntegrations")
    
    
    return
}

func NewDescribeExporterIntegrationsResponse() (response *DescribeExporterIntegrationsResponse) {
    response = &DescribeExporterIntegrationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExporterIntegrations
// 查询集成中心 exporter 集成列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTNOTALLOWED = "FailedOperation.AgentNotAllowed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeExporterIntegrations(request *DescribeExporterIntegrationsRequest) (response *DescribeExporterIntegrationsResponse, err error) {
    return c.DescribeExporterIntegrationsWithContext(context.Background(), request)
}

// DescribeExporterIntegrations
// 查询集成中心 exporter 集成列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTNOTALLOWED = "FailedOperation.AgentNotAllowed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeExporterIntegrationsWithContext(ctx context.Context, request *DescribeExporterIntegrationsRequest) (response *DescribeExporterIntegrationsResponse, err error) {
    if request == nil {
        request = NewDescribeExporterIntegrationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeExporterIntegrations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExporterIntegrations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExporterIntegrationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExternalClusterRegisterCommandRequest() (request *DescribeExternalClusterRegisterCommandRequest) {
    request = &DescribeExternalClusterRegisterCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeExternalClusterRegisterCommand")
    
    
    return
}

func NewDescribeExternalClusterRegisterCommandResponse() (response *DescribeExternalClusterRegisterCommandResponse) {
    response = &DescribeExternalClusterRegisterCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExternalClusterRegisterCommand
// 查看外部集群注册命令
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeExternalClusterRegisterCommand(request *DescribeExternalClusterRegisterCommandRequest) (response *DescribeExternalClusterRegisterCommandResponse, err error) {
    return c.DescribeExternalClusterRegisterCommandWithContext(context.Background(), request)
}

// DescribeExternalClusterRegisterCommand
// 查看外部集群注册命令
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeExternalClusterRegisterCommandWithContext(ctx context.Context, request *DescribeExternalClusterRegisterCommandRequest) (response *DescribeExternalClusterRegisterCommandResponse, err error) {
    if request == nil {
        request = NewDescribeExternalClusterRegisterCommandRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeExternalClusterRegisterCommand")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExternalClusterRegisterCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExternalClusterRegisterCommandResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExternalClusterUninstallCommandRequest() (request *DescribeExternalClusterUninstallCommandRequest) {
    request = &DescribeExternalClusterUninstallCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeExternalClusterUninstallCommand")
    
    
    return
}

func NewDescribeExternalClusterUninstallCommandResponse() (response *DescribeExternalClusterUninstallCommandResponse) {
    response = &DescribeExternalClusterUninstallCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExternalClusterUninstallCommand
// 查看外部集群 Agent 卸载命令
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeExternalClusterUninstallCommand(request *DescribeExternalClusterUninstallCommandRequest) (response *DescribeExternalClusterUninstallCommandResponse, err error) {
    return c.DescribeExternalClusterUninstallCommandWithContext(context.Background(), request)
}

// DescribeExternalClusterUninstallCommand
// 查看外部集群 Agent 卸载命令
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeExternalClusterUninstallCommandWithContext(ctx context.Context, request *DescribeExternalClusterUninstallCommandRequest) (response *DescribeExternalClusterUninstallCommandResponse, err error) {
    if request == nil {
        request = NewDescribeExternalClusterUninstallCommandRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeExternalClusterUninstallCommand")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExternalClusterUninstallCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExternalClusterUninstallCommandResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaChannelsRequest() (request *DescribeGrafanaChannelsRequest) {
    request = &DescribeGrafanaChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaChannels")
    
    
    return
}

func NewDescribeGrafanaChannelsResponse() (response *DescribeGrafanaChannelsResponse) {
    response = &DescribeGrafanaChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaChannels
// 列出 Grafana 所有告警通道
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaChannels(request *DescribeGrafanaChannelsRequest) (response *DescribeGrafanaChannelsResponse, err error) {
    return c.DescribeGrafanaChannelsWithContext(context.Background(), request)
}

// DescribeGrafanaChannels
// 列出 Grafana 所有告警通道
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaChannelsWithContext(ctx context.Context, request *DescribeGrafanaChannelsRequest) (response *DescribeGrafanaChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaChannelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaConfigRequest() (request *DescribeGrafanaConfigRequest) {
    request = &DescribeGrafanaConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaConfig")
    
    
    return
}

func NewDescribeGrafanaConfigResponse() (response *DescribeGrafanaConfigResponse) {
    response = &DescribeGrafanaConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaConfig
// 列出 Grafana 的设置，即 grafana.ini 文件内容
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaConfig(request *DescribeGrafanaConfigRequest) (response *DescribeGrafanaConfigResponse, err error) {
    return c.DescribeGrafanaConfigWithContext(context.Background(), request)
}

// DescribeGrafanaConfig
// 列出 Grafana 的设置，即 grafana.ini 文件内容
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaConfigWithContext(ctx context.Context, request *DescribeGrafanaConfigRequest) (response *DescribeGrafanaConfigResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaEnvironmentsRequest() (request *DescribeGrafanaEnvironmentsRequest) {
    request = &DescribeGrafanaEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaEnvironments")
    
    
    return
}

func NewDescribeGrafanaEnvironmentsResponse() (response *DescribeGrafanaEnvironmentsResponse) {
    response = &DescribeGrafanaEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaEnvironments
// 列出 Grafana 环境变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaEnvironments(request *DescribeGrafanaEnvironmentsRequest) (response *DescribeGrafanaEnvironmentsResponse, err error) {
    return c.DescribeGrafanaEnvironmentsWithContext(context.Background(), request)
}

// DescribeGrafanaEnvironments
// 列出 Grafana 环境变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaEnvironmentsWithContext(ctx context.Context, request *DescribeGrafanaEnvironmentsRequest) (response *DescribeGrafanaEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaEnvironmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaEnvironments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaInstancesRequest() (request *DescribeGrafanaInstancesRequest) {
    request = &DescribeGrafanaInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaInstances")
    
    
    return
}

func NewDescribeGrafanaInstancesResponse() (response *DescribeGrafanaInstancesResponse) {
    response = &DescribeGrafanaInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaInstances
// 列出用户所有的 Grafana 服务
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeGrafanaInstances(request *DescribeGrafanaInstancesRequest) (response *DescribeGrafanaInstancesResponse, err error) {
    return c.DescribeGrafanaInstancesWithContext(context.Background(), request)
}

// DescribeGrafanaInstances
// 列出用户所有的 Grafana 服务
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeGrafanaInstancesWithContext(ctx context.Context, request *DescribeGrafanaInstancesRequest) (response *DescribeGrafanaInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaIntegrationsRequest() (request *DescribeGrafanaIntegrationsRequest) {
    request = &DescribeGrafanaIntegrationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaIntegrations")
    
    
    return
}

func NewDescribeGrafanaIntegrationsResponse() (response *DescribeGrafanaIntegrationsResponse) {
    response = &DescribeGrafanaIntegrationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaIntegrations
// 列出 Grafana 已安装的集成
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaIntegrations(request *DescribeGrafanaIntegrationsRequest) (response *DescribeGrafanaIntegrationsResponse, err error) {
    return c.DescribeGrafanaIntegrationsWithContext(context.Background(), request)
}

// DescribeGrafanaIntegrations
// 列出 Grafana 已安装的集成
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaIntegrationsWithContext(ctx context.Context, request *DescribeGrafanaIntegrationsRequest) (response *DescribeGrafanaIntegrationsResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaIntegrationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaIntegrations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaIntegrations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaIntegrationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaNotificationChannelsRequest() (request *DescribeGrafanaNotificationChannelsRequest) {
    request = &DescribeGrafanaNotificationChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaNotificationChannels")
    
    
    return
}

func NewDescribeGrafanaNotificationChannelsResponse() (response *DescribeGrafanaNotificationChannelsResponse) {
    response = &DescribeGrafanaNotificationChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaNotificationChannels
// 列出 Grafana 告警通道
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaNotificationChannels(request *DescribeGrafanaNotificationChannelsRequest) (response *DescribeGrafanaNotificationChannelsResponse, err error) {
    return c.DescribeGrafanaNotificationChannelsWithContext(context.Background(), request)
}

// DescribeGrafanaNotificationChannels
// 列出 Grafana 告警通道
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaNotificationChannelsWithContext(ctx context.Context, request *DescribeGrafanaNotificationChannelsRequest) (response *DescribeGrafanaNotificationChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaNotificationChannelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaNotificationChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaNotificationChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaNotificationChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaWhiteListRequest() (request *DescribeGrafanaWhiteListRequest) {
    request = &DescribeGrafanaWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaWhiteList")
    
    
    return
}

func NewDescribeGrafanaWhiteListResponse() (response *DescribeGrafanaWhiteListResponse) {
    response = &DescribeGrafanaWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaWhiteList
// 列出 Grafana 白名单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaWhiteList(request *DescribeGrafanaWhiteListRequest) (response *DescribeGrafanaWhiteListResponse, err error) {
    return c.DescribeGrafanaWhiteListWithContext(context.Background(), request)
}

// DescribeGrafanaWhiteList
// 列出 Grafana 白名单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaWhiteListWithContext(ctx context.Context, request *DescribeGrafanaWhiteListRequest) (response *DescribeGrafanaWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaWhiteListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaWhiteList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstalledPluginsRequest() (request *DescribeInstalledPluginsRequest) {
    request = &DescribeInstalledPluginsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeInstalledPlugins")
    
    
    return
}

func NewDescribeInstalledPluginsResponse() (response *DescribeInstalledPluginsResponse) {
    response = &DescribeInstalledPluginsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstalledPlugins
// 列出实例已安装的插件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstalledPlugins(request *DescribeInstalledPluginsRequest) (response *DescribeInstalledPluginsResponse, err error) {
    return c.DescribeInstalledPluginsWithContext(context.Background(), request)
}

// DescribeInstalledPlugins
// 列出实例已安装的插件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstalledPluginsWithContext(ctx context.Context, request *DescribeInstalledPluginsRequest) (response *DescribeInstalledPluginsResponse, err error) {
    if request == nil {
        request = NewDescribeInstalledPluginsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeInstalledPlugins")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstalledPlugins require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstalledPluginsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorResourceInfoRequest() (request *DescribeMonitorResourceInfoRequest) {
    request = &DescribeMonitorResourceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeMonitorResourceInfo")
    
    
    return
}

func NewDescribeMonitorResourceInfoResponse() (response *DescribeMonitorResourceInfoResponse) {
    response = &DescribeMonitorResourceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMonitorResourceInfo
// 获取资源消耗页概览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeMonitorResourceInfo(request *DescribeMonitorResourceInfoRequest) (response *DescribeMonitorResourceInfoResponse, err error) {
    return c.DescribeMonitorResourceInfoWithContext(context.Background(), request)
}

// DescribeMonitorResourceInfo
// 获取资源消耗页概览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeMonitorResourceInfoWithContext(ctx context.Context, request *DescribeMonitorResourceInfoRequest) (response *DescribeMonitorResourceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorResourceInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeMonitorResourceInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMonitorResourceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMonitorResourceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorTypesRequest() (request *DescribeMonitorTypesRequest) {
    request = &DescribeMonitorTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeMonitorTypes")
    
    
    return
}

func NewDescribeMonitorTypesResponse() (response *DescribeMonitorTypesResponse) {
    response = &DescribeMonitorTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMonitorTypes
// 腾讯云可观测平台支持多种类型的监控，此接口列出支持的所有类型
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonitorTypes(request *DescribeMonitorTypesRequest) (response *DescribeMonitorTypesResponse, err error) {
    return c.DescribeMonitorTypesWithContext(context.Background(), request)
}

// DescribeMonitorTypes
// 腾讯云可观测平台支持多种类型的监控，此接口列出支持的所有类型
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonitorTypesWithContext(ctx context.Context, request *DescribeMonitorTypesRequest) (response *DescribeMonitorTypesResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorTypesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeMonitorTypes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMonitorTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMonitorTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePhoneAlarmFlowTotalCountRequest() (request *DescribePhoneAlarmFlowTotalCountRequest) {
    request = &DescribePhoneAlarmFlowTotalCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePhoneAlarmFlowTotalCount")
    
    
    return
}

func NewDescribePhoneAlarmFlowTotalCountResponse() (response *DescribePhoneAlarmFlowTotalCountResponse) {
    response = &DescribePhoneAlarmFlowTotalCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePhoneAlarmFlowTotalCount
// 查询周期内电话流水总数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePhoneAlarmFlowTotalCount(request *DescribePhoneAlarmFlowTotalCountRequest) (response *DescribePhoneAlarmFlowTotalCountResponse, err error) {
    return c.DescribePhoneAlarmFlowTotalCountWithContext(context.Background(), request)
}

// DescribePhoneAlarmFlowTotalCount
// 查询周期内电话流水总数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePhoneAlarmFlowTotalCountWithContext(ctx context.Context, request *DescribePhoneAlarmFlowTotalCountRequest) (response *DescribePhoneAlarmFlowTotalCountResponse, err error) {
    if request == nil {
        request = NewDescribePhoneAlarmFlowTotalCountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePhoneAlarmFlowTotalCount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePhoneAlarmFlowTotalCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePhoneAlarmFlowTotalCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginOverviewsRequest() (request *DescribePluginOverviewsRequest) {
    request = &DescribePluginOverviewsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePluginOverviews")
    
    
    return
}

func NewDescribePluginOverviewsResponse() (response *DescribePluginOverviewsResponse) {
    response = &DescribePluginOverviewsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePluginOverviews
// 列出可安装的所有 Grafana 插件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePluginOverviews(request *DescribePluginOverviewsRequest) (response *DescribePluginOverviewsResponse, err error) {
    return c.DescribePluginOverviewsWithContext(context.Background(), request)
}

// DescribePluginOverviews
// 列出可安装的所有 Grafana 插件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePluginOverviewsWithContext(ctx context.Context, request *DescribePluginOverviewsRequest) (response *DescribePluginOverviewsResponse, err error) {
    if request == nil {
        request = NewDescribePluginOverviewsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePluginOverviews")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePluginOverviews require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginOverviewsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyConditionListRequest() (request *DescribePolicyConditionListRequest) {
    request = &DescribePolicyConditionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyConditionList")
    
    
    return
}

func NewDescribePolicyConditionListResponse() (response *DescribePolicyConditionListResponse) {
    response = &DescribePolicyConditionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicyConditionList
// 获取基础告警策略条件
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyConditionList(request *DescribePolicyConditionListRequest) (response *DescribePolicyConditionListResponse, err error) {
    return c.DescribePolicyConditionListWithContext(context.Background(), request)
}

// DescribePolicyConditionList
// 获取基础告警策略条件
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyConditionListWithContext(ctx context.Context, request *DescribePolicyConditionListRequest) (response *DescribePolicyConditionListResponse, err error) {
    if request == nil {
        request = NewDescribePolicyConditionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePolicyConditionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyConditionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyConditionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyGroupInfoRequest() (request *DescribePolicyGroupInfoRequest) {
    request = &DescribePolicyGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyGroupInfo")
    
    
    return
}

func NewDescribePolicyGroupInfoResponse() (response *DescribePolicyGroupInfoResponse) {
    response = &DescribePolicyGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicyGroupInfo
// 获取基础策略组详情
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyGroupInfo(request *DescribePolicyGroupInfoRequest) (response *DescribePolicyGroupInfoResponse, err error) {
    return c.DescribePolicyGroupInfoWithContext(context.Background(), request)
}

// DescribePolicyGroupInfo
// 获取基础策略组详情
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyGroupInfoWithContext(ctx context.Context, request *DescribePolicyGroupInfoRequest) (response *DescribePolicyGroupInfoResponse, err error) {
    if request == nil {
        request = NewDescribePolicyGroupInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePolicyGroupInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyGroupListRequest() (request *DescribePolicyGroupListRequest) {
    request = &DescribePolicyGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyGroupList")
    
    
    return
}

func NewDescribePolicyGroupListResponse() (response *DescribePolicyGroupListResponse) {
    response = &DescribePolicyGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicyGroupList
// 获取基础策略告警组列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyGroupList(request *DescribePolicyGroupListRequest) (response *DescribePolicyGroupListResponse, err error) {
    return c.DescribePolicyGroupListWithContext(context.Background(), request)
}

// DescribePolicyGroupList
// 获取基础策略告警组列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyGroupListWithContext(ctx context.Context, request *DescribePolicyGroupListRequest) (response *DescribePolicyGroupListResponse, err error) {
    if request == nil {
        request = NewDescribePolicyGroupListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePolicyGroupList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyObjectCountRequest() (request *DescribePolicyObjectCountRequest) {
    request = &DescribePolicyObjectCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyObjectCount")
    
    
    return
}

func NewDescribePolicyObjectCountResponse() (response *DescribePolicyObjectCountResponse) {
    response = &DescribePolicyObjectCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicyObjectCount
// 查询策略组在每个地域下面绑定的对象数统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyObjectCount(request *DescribePolicyObjectCountRequest) (response *DescribePolicyObjectCountResponse, err error) {
    return c.DescribePolicyObjectCountWithContext(context.Background(), request)
}

// DescribePolicyObjectCount
// 查询策略组在每个地域下面绑定的对象数统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyObjectCountWithContext(ctx context.Context, request *DescribePolicyObjectCountRequest) (response *DescribePolicyObjectCountResponse, err error) {
    if request == nil {
        request = NewDescribePolicyObjectCountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePolicyObjectCount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyObjectCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyObjectCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductEventListRequest() (request *DescribeProductEventListRequest) {
    request = &DescribeProductEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeProductEventList")
    
    
    return
}

func NewDescribeProductEventListResponse() (response *DescribeProductEventListResponse) {
    response = &DescribeProductEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProductEventList
// 分页获取产品事件的列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductEventList(request *DescribeProductEventListRequest) (response *DescribeProductEventListResponse, err error) {
    return c.DescribeProductEventListWithContext(context.Background(), request)
}

// DescribeProductEventList
// 分页获取产品事件的列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductEventListWithContext(ctx context.Context, request *DescribeProductEventListRequest) (response *DescribeProductEventListResponse, err error) {
    if request == nil {
        request = NewDescribeProductEventListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeProductEventList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductListRequest() (request *DescribeProductListRequest) {
    request = &DescribeProductListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeProductList")
    
    
    return
}

func NewDescribeProductListResponse() (response *DescribeProductListResponse) {
    response = &DescribeProductListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProductList
// 查询腾讯云可观测平台云产品列表，支持云服务器CVM、云数据库、云消息队列、负载均衡、容器服务、专线等云产品。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductList(request *DescribeProductListRequest) (response *DescribeProductListResponse, err error) {
    return c.DescribeProductListWithContext(context.Background(), request)
}

// DescribeProductList
// 查询腾讯云可观测平台云产品列表，支持云服务器CVM、云数据库、云消息队列、负载均衡、容器服务、专线等云产品。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductListWithContext(ctx context.Context, request *DescribeProductListRequest) (response *DescribeProductListResponse, err error) {
    if request == nil {
        request = NewDescribeProductListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeProductList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusAgentInstancesRequest() (request *DescribePrometheusAgentInstancesRequest) {
    request = &DescribePrometheusAgentInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusAgentInstances")
    
    
    return
}

func NewDescribePrometheusAgentInstancesResponse() (response *DescribePrometheusAgentInstancesResponse) {
    response = &DescribePrometheusAgentInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusAgentInstances
// 获取关联目标集群的实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusAgentInstances(request *DescribePrometheusAgentInstancesRequest) (response *DescribePrometheusAgentInstancesResponse, err error) {
    return c.DescribePrometheusAgentInstancesWithContext(context.Background(), request)
}

// DescribePrometheusAgentInstances
// 获取关联目标集群的实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusAgentInstancesWithContext(ctx context.Context, request *DescribePrometheusAgentInstancesRequest) (response *DescribePrometheusAgentInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusAgentInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusAgentInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAgentInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAgentInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusAgentsRequest() (request *DescribePrometheusAgentsRequest) {
    request = &DescribePrometheusAgentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusAgents")
    
    
    return
}

func NewDescribePrometheusAgentsResponse() (response *DescribePrometheusAgentsResponse) {
    response = &DescribePrometheusAgentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusAgents
// 列出 Prometheus CVM Agent
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusAgents(request *DescribePrometheusAgentsRequest) (response *DescribePrometheusAgentsResponse, err error) {
    return c.DescribePrometheusAgentsWithContext(context.Background(), request)
}

// DescribePrometheusAgents
// 列出 Prometheus CVM Agent
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusAgentsWithContext(ctx context.Context, request *DescribePrometheusAgentsRequest) (response *DescribePrometheusAgentsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusAgentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusAgents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAgents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAgentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusAlertGroupsRequest() (request *DescribePrometheusAlertGroupsRequest) {
    request = &DescribePrometheusAlertGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusAlertGroups")
    
    
    return
}

func NewDescribePrometheusAlertGroupsResponse() (response *DescribePrometheusAlertGroupsResponse) {
    response = &DescribePrometheusAlertGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusAlertGroups
// 查询给定prometheus下的告警分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusAlertGroups(request *DescribePrometheusAlertGroupsRequest) (response *DescribePrometheusAlertGroupsResponse, err error) {
    return c.DescribePrometheusAlertGroupsWithContext(context.Background(), request)
}

// DescribePrometheusAlertGroups
// 查询给定prometheus下的告警分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusAlertGroupsWithContext(ctx context.Context, request *DescribePrometheusAlertGroupsRequest) (response *DescribePrometheusAlertGroupsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusAlertGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusAlertGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAlertGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAlertGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusAlertPolicyRequest() (request *DescribePrometheusAlertPolicyRequest) {
    request = &DescribePrometheusAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusAlertPolicy")
    
    
    return
}

func NewDescribePrometheusAlertPolicyResponse() (response *DescribePrometheusAlertPolicyResponse) {
    response = &DescribePrometheusAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusAlertPolicy
// 获取2.0实例告警策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusAlertPolicy(request *DescribePrometheusAlertPolicyRequest) (response *DescribePrometheusAlertPolicyResponse, err error) {
    return c.DescribePrometheusAlertPolicyWithContext(context.Background(), request)
}

// DescribePrometheusAlertPolicy
// 获取2.0实例告警策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusAlertPolicyWithContext(ctx context.Context, request *DescribePrometheusAlertPolicyRequest) (response *DescribePrometheusAlertPolicyResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusAlertPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusAlertPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAlertPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusClusterAgentsRequest() (request *DescribePrometheusClusterAgentsRequest) {
    request = &DescribePrometheusClusterAgentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusClusterAgents")
    
    
    return
}

func NewDescribePrometheusClusterAgentsResponse() (response *DescribePrometheusClusterAgentsResponse) {
    response = &DescribePrometheusClusterAgentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusClusterAgents
// 获取TMP实例关联集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DescribePrometheusClusterAgents(request *DescribePrometheusClusterAgentsRequest) (response *DescribePrometheusClusterAgentsResponse, err error) {
    return c.DescribePrometheusClusterAgentsWithContext(context.Background(), request)
}

// DescribePrometheusClusterAgents
// 获取TMP实例关联集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DescribePrometheusClusterAgentsWithContext(ctx context.Context, request *DescribePrometheusClusterAgentsRequest) (response *DescribePrometheusClusterAgentsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusClusterAgentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusClusterAgents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusClusterAgents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusClusterAgentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusConfigRequest() (request *DescribePrometheusConfigRequest) {
    request = &DescribePrometheusConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusConfig")
    
    
    return
}

func NewDescribePrometheusConfigResponse() (response *DescribePrometheusConfigResponse) {
    response = &DescribePrometheusConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusConfig
// 拉取Prometheus配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusConfig(request *DescribePrometheusConfigRequest) (response *DescribePrometheusConfigResponse, err error) {
    return c.DescribePrometheusConfigWithContext(context.Background(), request)
}

// DescribePrometheusConfig
// 拉取Prometheus配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusConfigWithContext(ctx context.Context, request *DescribePrometheusConfigRequest) (response *DescribePrometheusConfigResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusGlobalConfigRequest() (request *DescribePrometheusGlobalConfigRequest) {
    request = &DescribePrometheusGlobalConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusGlobalConfig")
    
    
    return
}

func NewDescribePrometheusGlobalConfigResponse() (response *DescribePrometheusGlobalConfigResponse) {
    response = &DescribePrometheusGlobalConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusGlobalConfig
// 获得实例级别抓取配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusGlobalConfig(request *DescribePrometheusGlobalConfigRequest) (response *DescribePrometheusGlobalConfigResponse, err error) {
    return c.DescribePrometheusGlobalConfigWithContext(context.Background(), request)
}

// DescribePrometheusGlobalConfig
// 获得实例级别抓取配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusGlobalConfigWithContext(ctx context.Context, request *DescribePrometheusGlobalConfigRequest) (response *DescribePrometheusGlobalConfigResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusGlobalConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusGlobalConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusGlobalConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusGlobalConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusGlobalNotificationRequest() (request *DescribePrometheusGlobalNotificationRequest) {
    request = &DescribePrometheusGlobalNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusGlobalNotification")
    
    
    return
}

func NewDescribePrometheusGlobalNotificationResponse() (response *DescribePrometheusGlobalNotificationResponse) {
    response = &DescribePrometheusGlobalNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusGlobalNotification
// 查询全局告警通知渠道
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusGlobalNotification(request *DescribePrometheusGlobalNotificationRequest) (response *DescribePrometheusGlobalNotificationResponse, err error) {
    return c.DescribePrometheusGlobalNotificationWithContext(context.Background(), request)
}

// DescribePrometheusGlobalNotification
// 查询全局告警通知渠道
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusGlobalNotificationWithContext(ctx context.Context, request *DescribePrometheusGlobalNotificationRequest) (response *DescribePrometheusGlobalNotificationResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusGlobalNotificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusGlobalNotification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusGlobalNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusGlobalNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstanceDetailRequest() (request *DescribePrometheusInstanceDetailRequest) {
    request = &DescribePrometheusInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusInstanceDetail")
    
    
    return
}

func NewDescribePrometheusInstanceDetailResponse() (response *DescribePrometheusInstanceDetailResponse) {
    response = &DescribePrometheusInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusInstanceDetail
// 获取TMP实例详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusInstanceDetail(request *DescribePrometheusInstanceDetailRequest) (response *DescribePrometheusInstanceDetailResponse, err error) {
    return c.DescribePrometheusInstanceDetailWithContext(context.Background(), request)
}

// DescribePrometheusInstanceDetail
// 获取TMP实例详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusInstanceDetailWithContext(ctx context.Context, request *DescribePrometheusInstanceDetailRequest) (response *DescribePrometheusInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstanceDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusInstanceDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstanceInitStatusRequest() (request *DescribePrometheusInstanceInitStatusRequest) {
    request = &DescribePrometheusInstanceInitStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusInstanceInitStatus")
    
    
    return
}

func NewDescribePrometheusInstanceInitStatusResponse() (response *DescribePrometheusInstanceInitStatusResponse) {
    response = &DescribePrometheusInstanceInitStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusInstanceInitStatus
// 获取2.0实例初始化任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusInstanceInitStatus(request *DescribePrometheusInstanceInitStatusRequest) (response *DescribePrometheusInstanceInitStatusResponse, err error) {
    return c.DescribePrometheusInstanceInitStatusWithContext(context.Background(), request)
}

// DescribePrometheusInstanceInitStatus
// 获取2.0实例初始化任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusInstanceInitStatusWithContext(ctx context.Context, request *DescribePrometheusInstanceInitStatusRequest) (response *DescribePrometheusInstanceInitStatusResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstanceInitStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusInstanceInitStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstanceInitStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstanceInitStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstanceUsageRequest() (request *DescribePrometheusInstanceUsageRequest) {
    request = &DescribePrometheusInstanceUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusInstanceUsage")
    
    
    return
}

func NewDescribePrometheusInstanceUsageResponse() (response *DescribePrometheusInstanceUsageResponse) {
    response = &DescribePrometheusInstanceUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusInstanceUsage
// 查询Prometheus按量实例用量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusInstanceUsage(request *DescribePrometheusInstanceUsageRequest) (response *DescribePrometheusInstanceUsageResponse, err error) {
    return c.DescribePrometheusInstanceUsageWithContext(context.Background(), request)
}

// DescribePrometheusInstanceUsage
// 查询Prometheus按量实例用量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusInstanceUsageWithContext(ctx context.Context, request *DescribePrometheusInstanceUsageRequest) (response *DescribePrometheusInstanceUsageResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstanceUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusInstanceUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstanceUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstanceUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstancesRequest() (request *DescribePrometheusInstancesRequest) {
    request = &DescribePrometheusInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusInstances")
    
    
    return
}

func NewDescribePrometheusInstancesResponse() (response *DescribePrometheusInstancesResponse) {
    response = &DescribePrometheusInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusInstances
// 本接口 (DescribePrometheusInstances) 用于查询一个或多个实例的详细信息。
//
// <ul>
//
// <li>可以根据实例ID、实例名称或者实例状态等信息来查询实例的详细信息</li>
//
// <li>如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的实例。</li>
//
// </ul>
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePrometheusInstances(request *DescribePrometheusInstancesRequest) (response *DescribePrometheusInstancesResponse, err error) {
    return c.DescribePrometheusInstancesWithContext(context.Background(), request)
}

// DescribePrometheusInstances
// 本接口 (DescribePrometheusInstances) 用于查询一个或多个实例的详细信息。
//
// <ul>
//
// <li>可以根据实例ID、实例名称或者实例状态等信息来查询实例的详细信息</li>
//
// <li>如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的实例。</li>
//
// </ul>
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePrometheusInstancesWithContext(ctx context.Context, request *DescribePrometheusInstancesRequest) (response *DescribePrometheusInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstancesOverviewRequest() (request *DescribePrometheusInstancesOverviewRequest) {
    request = &DescribePrometheusInstancesOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusInstancesOverview")
    
    
    return
}

func NewDescribePrometheusInstancesOverviewResponse() (response *DescribePrometheusInstancesOverviewResponse) {
    response = &DescribePrometheusInstancesOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusInstancesOverview
// 获取与 Prometheus 监控融合实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusInstancesOverview(request *DescribePrometheusInstancesOverviewRequest) (response *DescribePrometheusInstancesOverviewResponse, err error) {
    return c.DescribePrometheusInstancesOverviewWithContext(context.Background(), request)
}

// DescribePrometheusInstancesOverview
// 获取与 Prometheus 监控融合实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusInstancesOverviewWithContext(ctx context.Context, request *DescribePrometheusInstancesOverviewRequest) (response *DescribePrometheusInstancesOverviewResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstancesOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusInstancesOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstancesOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstancesOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusIntegrationMetricsRequest() (request *DescribePrometheusIntegrationMetricsRequest) {
    request = &DescribePrometheusIntegrationMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusIntegrationMetrics")
    
    
    return
}

func NewDescribePrometheusIntegrationMetricsResponse() (response *DescribePrometheusIntegrationMetricsResponse) {
    response = &DescribePrometheusIntegrationMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusIntegrationMetrics
// 获取prometheus集成指标
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribePrometheusIntegrationMetrics(request *DescribePrometheusIntegrationMetricsRequest) (response *DescribePrometheusIntegrationMetricsResponse, err error) {
    return c.DescribePrometheusIntegrationMetricsWithContext(context.Background(), request)
}

// DescribePrometheusIntegrationMetrics
// 获取prometheus集成指标
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribePrometheusIntegrationMetricsWithContext(ctx context.Context, request *DescribePrometheusIntegrationMetricsRequest) (response *DescribePrometheusIntegrationMetricsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusIntegrationMetricsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusIntegrationMetrics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusIntegrationMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusIntegrationMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusRecordRulesRequest() (request *DescribePrometheusRecordRulesRequest) {
    request = &DescribePrometheusRecordRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusRecordRules")
    
    
    return
}

func NewDescribePrometheusRecordRulesResponse() (response *DescribePrometheusRecordRulesResponse) {
    response = &DescribePrometheusRecordRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusRecordRules
// 获取聚合规则列表，包含关联集群内crd资源创建的record rule
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusRecordRules(request *DescribePrometheusRecordRulesRequest) (response *DescribePrometheusRecordRulesResponse, err error) {
    return c.DescribePrometheusRecordRulesWithContext(context.Background(), request)
}

// DescribePrometheusRecordRules
// 获取聚合规则列表，包含关联集群内crd资源创建的record rule
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusRecordRulesWithContext(ctx context.Context, request *DescribePrometheusRecordRulesRequest) (response *DescribePrometheusRecordRulesResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusRecordRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusRecordRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusRecordRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusRecordRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusRegionsRequest() (request *DescribePrometheusRegionsRequest) {
    request = &DescribePrometheusRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusRegions")
    
    
    return
}

func NewDescribePrometheusRegionsResponse() (response *DescribePrometheusRegionsResponse) {
    response = &DescribePrometheusRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusRegions
// 列出 Prometheus 服务所有可用的地域
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusRegions(request *DescribePrometheusRegionsRequest) (response *DescribePrometheusRegionsResponse, err error) {
    return c.DescribePrometheusRegionsWithContext(context.Background(), request)
}

// DescribePrometheusRegions
// 列出 Prometheus 服务所有可用的地域
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusRegionsWithContext(ctx context.Context, request *DescribePrometheusRegionsRequest) (response *DescribePrometheusRegionsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusRegionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusRegions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusScrapeJobsRequest() (request *DescribePrometheusScrapeJobsRequest) {
    request = &DescribePrometheusScrapeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusScrapeJobs")
    
    
    return
}

func NewDescribePrometheusScrapeJobsResponse() (response *DescribePrometheusScrapeJobsResponse) {
    response = &DescribePrometheusScrapeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusScrapeJobs
// 列出 Prometheus 抓取任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusScrapeJobs(request *DescribePrometheusScrapeJobsRequest) (response *DescribePrometheusScrapeJobsResponse, err error) {
    return c.DescribePrometheusScrapeJobsWithContext(context.Background(), request)
}

// DescribePrometheusScrapeJobs
// 列出 Prometheus 抓取任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusScrapeJobsWithContext(ctx context.Context, request *DescribePrometheusScrapeJobsRequest) (response *DescribePrometheusScrapeJobsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusScrapeJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusScrapeJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusScrapeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusScrapeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusScrapeStatisticsRequest() (request *DescribePrometheusScrapeStatisticsRequest) {
    request = &DescribePrometheusScrapeStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusScrapeStatistics")
    
    
    return
}

func NewDescribePrometheusScrapeStatisticsResponse() (response *DescribePrometheusScrapeStatisticsResponse) {
    response = &DescribePrometheusScrapeStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusScrapeStatistics
// 获取实例采集速率信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusScrapeStatistics(request *DescribePrometheusScrapeStatisticsRequest) (response *DescribePrometheusScrapeStatisticsResponse, err error) {
    return c.DescribePrometheusScrapeStatisticsWithContext(context.Background(), request)
}

// DescribePrometheusScrapeStatistics
// 获取实例采集速率信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusScrapeStatisticsWithContext(ctx context.Context, request *DescribePrometheusScrapeStatisticsRequest) (response *DescribePrometheusScrapeStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusScrapeStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusScrapeStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusScrapeStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusScrapeStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusTargetsTMPRequest() (request *DescribePrometheusTargetsTMPRequest) {
    request = &DescribePrometheusTargetsTMPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusTargetsTMP")
    
    
    return
}

func NewDescribePrometheusTargetsTMPResponse() (response *DescribePrometheusTargetsTMPResponse) {
    response = &DescribePrometheusTargetsTMPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusTargetsTMP
// 获取targets信息，在过滤条件中指定job名称时返回targets详情，否则仅返回数量
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusTargetsTMP(request *DescribePrometheusTargetsTMPRequest) (response *DescribePrometheusTargetsTMPResponse, err error) {
    return c.DescribePrometheusTargetsTMPWithContext(context.Background(), request)
}

// DescribePrometheusTargetsTMP
// 获取targets信息，在过滤条件中指定job名称时返回targets详情，否则仅返回数量
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusTargetsTMPWithContext(ctx context.Context, request *DescribePrometheusTargetsTMPRequest) (response *DescribePrometheusTargetsTMPResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusTargetsTMPRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusTargetsTMP")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusTargetsTMP require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusTargetsTMPResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusTempRequest() (request *DescribePrometheusTempRequest) {
    request = &DescribePrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusTemp")
    
    
    return
}

func NewDescribePrometheusTempResponse() (response *DescribePrometheusTempResponse) {
    response = &DescribePrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusTemp
// 拉取模板列表，默认模板将总是在最前面
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusTemp(request *DescribePrometheusTempRequest) (response *DescribePrometheusTempResponse, err error) {
    return c.DescribePrometheusTempWithContext(context.Background(), request)
}

// DescribePrometheusTemp
// 拉取模板列表，默认模板将总是在最前面
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusTempWithContext(ctx context.Context, request *DescribePrometheusTempRequest) (response *DescribePrometheusTempResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusTempRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusTemp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusTempSyncRequest() (request *DescribePrometheusTempSyncRequest) {
    request = &DescribePrometheusTempSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusTempSync")
    
    
    return
}

func NewDescribePrometheusTempSyncResponse() (response *DescribePrometheusTempSyncResponse) {
    response = &DescribePrometheusTempSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusTempSync
// 获取模板关联实例信息，针对V2版本实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusTempSync(request *DescribePrometheusTempSyncRequest) (response *DescribePrometheusTempSyncResponse, err error) {
    return c.DescribePrometheusTempSyncWithContext(context.Background(), request)
}

// DescribePrometheusTempSync
// 获取模板关联实例信息，针对V2版本实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusTempSyncWithContext(ctx context.Context, request *DescribePrometheusTempSyncRequest) (response *DescribePrometheusTempSyncResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusTempSyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusTempSync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusTempSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusTempSyncResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusZonesRequest() (request *DescribePrometheusZonesRequest) {
    request = &DescribePrometheusZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusZones")
    
    
    return
}

func NewDescribePrometheusZonesResponse() (response *DescribePrometheusZonesResponse) {
    response = &DescribePrometheusZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusZones
// 列出 Prometheus 服务可用区。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_REGIONUNAVAILABLE = "FailedOperation.RegionUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusZones(request *DescribePrometheusZonesRequest) (response *DescribePrometheusZonesResponse, err error) {
    return c.DescribePrometheusZonesWithContext(context.Background(), request)
}

// DescribePrometheusZones
// 列出 Prometheus 服务可用区。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_REGIONUNAVAILABLE = "FailedOperation.RegionUnavailable"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusZonesWithContext(ctx context.Context, request *DescribePrometheusZonesRequest) (response *DescribePrometheusZonesResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusZonesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordingRulesRequest() (request *DescribeRecordingRulesRequest) {
    request = &DescribeRecordingRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeRecordingRules")
    
    
    return
}

func NewDescribeRecordingRulesResponse() (response *DescribeRecordingRulesResponse) {
    response = &DescribeRecordingRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordingRules
// 根据条件查询 Prometheus 预聚合规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRecordingRules(request *DescribeRecordingRulesRequest) (response *DescribeRecordingRulesResponse, err error) {
    return c.DescribeRecordingRulesWithContext(context.Background(), request)
}

// DescribeRecordingRules
// 根据条件查询 Prometheus 预聚合规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRecordingRulesWithContext(ctx context.Context, request *DescribeRecordingRulesRequest) (response *DescribeRecordingRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRecordingRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeRecordingRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordingRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordingRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRemoteURLsRequest() (request *DescribeRemoteURLsRequest) {
    request = &DescribeRemoteURLsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeRemoteURLs")
    
    
    return
}

func NewDescribeRemoteURLsResponse() (response *DescribeRemoteURLsResponse) {
    response = &DescribeRemoteURLsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRemoteURLs
// 获取多写配置详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribeRemoteURLs(request *DescribeRemoteURLsRequest) (response *DescribeRemoteURLsResponse, err error) {
    return c.DescribeRemoteURLsWithContext(context.Background(), request)
}

// DescribeRemoteURLs
// 获取多写配置详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribeRemoteURLsWithContext(ctx context.Context, request *DescribeRemoteURLsRequest) (response *DescribeRemoteURLsResponse, err error) {
    if request == nil {
        request = NewDescribeRemoteURLsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeRemoteURLs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRemoteURLs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRemoteURLsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRemoteWritesRequest() (request *DescribeRemoteWritesRequest) {
    request = &DescribeRemoteWritesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeRemoteWrites")
    
    
    return
}

func NewDescribeRemoteWritesResponse() (response *DescribeRemoteWritesResponse) {
    response = &DescribeRemoteWritesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRemoteWrites
// 查询安装的 Agent 列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DescribeRemoteWrites(request *DescribeRemoteWritesRequest) (response *DescribeRemoteWritesResponse, err error) {
    return c.DescribeRemoteWritesWithContext(context.Background(), request)
}

// DescribeRemoteWrites
// 查询安装的 Agent 列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DescribeRemoteWritesWithContext(ctx context.Context, request *DescribeRemoteWritesRequest) (response *DescribeRemoteWritesResponse, err error) {
    if request == nil {
        request = NewDescribeRemoteWritesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeRemoteWrites")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRemoteWrites require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRemoteWritesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSSOAccountRequest() (request *DescribeSSOAccountRequest) {
    request = &DescribeSSOAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeSSOAccount")
    
    
    return
}

func NewDescribeSSOAccountResponse() (response *DescribeSSOAccountResponse) {
    response = &DescribeSSOAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSSOAccount
// 列出当前grafana实例的所有授权账号
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSSOAccount(request *DescribeSSOAccountRequest) (response *DescribeSSOAccountResponse, err error) {
    return c.DescribeSSOAccountWithContext(context.Background(), request)
}

// DescribeSSOAccount
// 列出当前grafana实例的所有授权账号
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSSOAccountWithContext(ctx context.Context, request *DescribeSSOAccountRequest) (response *DescribeSSOAccountResponse, err error) {
    if request == nil {
        request = NewDescribeSSOAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeSSOAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSSOAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSSOAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceDiscoveryRequest() (request *DescribeServiceDiscoveryRequest) {
    request = &DescribeServiceDiscoveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeServiceDiscovery")
    
    
    return
}

func NewDescribeServiceDiscoveryResponse() (response *DescribeServiceDiscoveryResponse) {
    response = &DescribeServiceDiscoveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceDiscovery
// 列出在腾讯云容器服务下创建的 Prometheus 服务发现。
//
// <p>注意：前提条件，已经通过 Prometheus 控制台集成了对应的腾讯云容器服务，具体请参考
//
// <a href="https://cloud.tencent.com/document/product/248/48859" target="_blank">Agent 安装</a>。</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeServiceDiscovery(request *DescribeServiceDiscoveryRequest) (response *DescribeServiceDiscoveryResponse, err error) {
    return c.DescribeServiceDiscoveryWithContext(context.Background(), request)
}

// DescribeServiceDiscovery
// 列出在腾讯云容器服务下创建的 Prometheus 服务发现。
//
// <p>注意：前提条件，已经通过 Prometheus 控制台集成了对应的腾讯云容器服务，具体请参考
//
// <a href="https://cloud.tencent.com/document/product/248/48859" target="_blank">Agent 安装</a>。</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeServiceDiscoveryWithContext(ctx context.Context, request *DescribeServiceDiscoveryRequest) (response *DescribeServiceDiscoveryResponse, err error) {
    if request == nil {
        request = NewDescribeServiceDiscoveryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeServiceDiscovery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceDiscovery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceDiscoveryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStatisticDataRequest() (request *DescribeStatisticDataRequest) {
    request = &DescribeStatisticDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeStatisticData")
    
    
    return
}

func NewDescribeStatisticDataResponse() (response *DescribeStatisticDataResponse) {
    response = &DescribeStatisticDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStatisticData
// 根据维度条件查询监控数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATACOLUMNNOTFOUND = "FailedOperation.DataColumnNotFound"
//  FAILEDOPERATION_DATAQUERYFAILED = "FailedOperation.DataQueryFailed"
//  FAILEDOPERATION_DATATABLENOTFOUND = "FailedOperation.DataTableNotFound"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLBACKFAIL = "InternalError.CallbackFail"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_DEPENDSMQ = "InternalError.DependsMq"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_TASKRESULTFORMAT = "InternalError.TaskResultFormat"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_MISSAKSK = "InvalidParameter.MissAKSK"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SECRETIDORSECRETKEYERROR = "InvalidParameter.SecretIdOrSecretKeyError"
//  INVALIDPARAMETER_UNSUPPORTEDPRODUCT = "InvalidParameter.UnsupportedProduct"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DASHBOARDNAMEEXISTS = "InvalidParameterValue.DashboardNameExists"
//  INVALIDPARAMETERVALUE_VERSIONMISMATCH = "InvalidParameterValue.VersionMismatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStatisticData(request *DescribeStatisticDataRequest) (response *DescribeStatisticDataResponse, err error) {
    return c.DescribeStatisticDataWithContext(context.Background(), request)
}

// DescribeStatisticData
// 根据维度条件查询监控数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATACOLUMNNOTFOUND = "FailedOperation.DataColumnNotFound"
//  FAILEDOPERATION_DATAQUERYFAILED = "FailedOperation.DataQueryFailed"
//  FAILEDOPERATION_DATATABLENOTFOUND = "FailedOperation.DataTableNotFound"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLBACKFAIL = "InternalError.CallbackFail"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_DEPENDSMQ = "InternalError.DependsMq"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_TASKRESULTFORMAT = "InternalError.TaskResultFormat"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_MISSAKSK = "InvalidParameter.MissAKSK"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SECRETIDORSECRETKEYERROR = "InvalidParameter.SecretIdOrSecretKeyError"
//  INVALIDPARAMETER_UNSUPPORTEDPRODUCT = "InvalidParameter.UnsupportedProduct"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DASHBOARDNAMEEXISTS = "InvalidParameterValue.DashboardNameExists"
//  INVALIDPARAMETERVALUE_VERSIONMISMATCH = "InvalidParameterValue.VersionMismatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStatisticDataWithContext(ctx context.Context, request *DescribeStatisticDataRequest) (response *DescribeStatisticDataResponse, err error) {
    if request == nil {
        request = NewDescribeStatisticDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeStatisticData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStatisticData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStatisticDataResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyPrometheusInstanceRequest() (request *DestroyPrometheusInstanceRequest) {
    request = &DestroyPrometheusInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DestroyPrometheusInstance")
    
    
    return
}

func NewDestroyPrometheusInstanceResponse() (response *DestroyPrometheusInstanceResponse) {
    response = &DestroyPrometheusInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyPrometheusInstance
// 彻底删除 Prometheus 实例相关数据，给定的实例必须先被 Terminate(该接口是异步接口，实例是否释放需要通过 DescribePrometheusInstances 接口返回的状态来判断)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTSNOTINUNINSTALLSTAGE = "FailedOperation.AgentsNotInUninstallStage"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DestroyPrometheusInstance(request *DestroyPrometheusInstanceRequest) (response *DestroyPrometheusInstanceResponse, err error) {
    return c.DestroyPrometheusInstanceWithContext(context.Background(), request)
}

// DestroyPrometheusInstance
// 彻底删除 Prometheus 实例相关数据，给定的实例必须先被 Terminate(该接口是异步接口，实例是否释放需要通过 DescribePrometheusInstances 接口返回的状态来判断)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTSNOTINUNINSTALLSTAGE = "FailedOperation.AgentsNotInUninstallStage"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DestroyPrometheusInstanceWithContext(ctx context.Context, request *DestroyPrometheusInstanceRequest) (response *DestroyPrometheusInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyPrometheusInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DestroyPrometheusInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyPrometheusInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyPrometheusInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewEnableGrafanaInternetRequest() (request *EnableGrafanaInternetRequest) {
    request = &EnableGrafanaInternetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "EnableGrafanaInternet")
    
    
    return
}

func NewEnableGrafanaInternetResponse() (response *EnableGrafanaInternetResponse) {
    response = &EnableGrafanaInternetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableGrafanaInternet
// 设置 Grafana 公网访问
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) EnableGrafanaInternet(request *EnableGrafanaInternetRequest) (response *EnableGrafanaInternetResponse, err error) {
    return c.EnableGrafanaInternetWithContext(context.Background(), request)
}

// EnableGrafanaInternet
// 设置 Grafana 公网访问
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) EnableGrafanaInternetWithContext(ctx context.Context, request *EnableGrafanaInternetRequest) (response *EnableGrafanaInternetResponse, err error) {
    if request == nil {
        request = NewEnableGrafanaInternetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "EnableGrafanaInternet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableGrafanaInternet require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableGrafanaInternetResponse()
    err = c.Send(request, response)
    return
}

func NewEnableGrafanaSSORequest() (request *EnableGrafanaSSORequest) {
    request = &EnableGrafanaSSORequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "EnableGrafanaSSO")
    
    
    return
}

func NewEnableGrafanaSSOResponse() (response *EnableGrafanaSSOResponse) {
    response = &EnableGrafanaSSOResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableGrafanaSSO
// 设置 Grafana 单点登录，使用腾讯云账号
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) EnableGrafanaSSO(request *EnableGrafanaSSORequest) (response *EnableGrafanaSSOResponse, err error) {
    return c.EnableGrafanaSSOWithContext(context.Background(), request)
}

// EnableGrafanaSSO
// 设置 Grafana 单点登录，使用腾讯云账号
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) EnableGrafanaSSOWithContext(ctx context.Context, request *EnableGrafanaSSORequest) (response *EnableGrafanaSSOResponse, err error) {
    if request == nil {
        request = NewEnableGrafanaSSORequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "EnableGrafanaSSO")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableGrafanaSSO require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableGrafanaSSOResponse()
    err = c.Send(request, response)
    return
}

func NewEnableSSOCamCheckRequest() (request *EnableSSOCamCheckRequest) {
    request = &EnableSSOCamCheckRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "EnableSSOCamCheck")
    
    
    return
}

func NewEnableSSOCamCheckResponse() (response *EnableSSOCamCheckResponse) {
    response = &EnableSSOCamCheckResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableSSOCamCheck
// SSO单点登录时，设置是否cam鉴权
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) EnableSSOCamCheck(request *EnableSSOCamCheckRequest) (response *EnableSSOCamCheckResponse, err error) {
    return c.EnableSSOCamCheckWithContext(context.Background(), request)
}

// EnableSSOCamCheck
// SSO单点登录时，设置是否cam鉴权
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) EnableSSOCamCheckWithContext(ctx context.Context, request *EnableSSOCamCheckRequest) (response *EnableSSOCamCheckResponse, err error) {
    if request == nil {
        request = NewEnableSSOCamCheckRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "EnableSSOCamCheck")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableSSOCamCheck require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableSSOCamCheckResponse()
    err = c.Send(request, response)
    return
}

func NewExportPrometheusReadOnlyDynamicAPIRequest() (request *ExportPrometheusReadOnlyDynamicAPIRequest) {
    request = &ExportPrometheusReadOnlyDynamicAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ExportPrometheusReadOnlyDynamicAPI")
    
    
    return
}

func NewExportPrometheusReadOnlyDynamicAPIResponse() (response *ExportPrometheusReadOnlyDynamicAPIResponse) {
    response = &ExportPrometheusReadOnlyDynamicAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportPrometheusReadOnlyDynamicAPI
// Prometheus 内部动态 api 代理，仅内部使用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATACOLUMNNOTFOUND = "FailedOperation.DataColumnNotFound"
//  FAILEDOPERATION_DATAQUERYFAILED = "FailedOperation.DataQueryFailed"
//  FAILEDOPERATION_DATATABLENOTFOUND = "FailedOperation.DataTableNotFound"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExportPrometheusReadOnlyDynamicAPI(request *ExportPrometheusReadOnlyDynamicAPIRequest) (response *ExportPrometheusReadOnlyDynamicAPIResponse, err error) {
    return c.ExportPrometheusReadOnlyDynamicAPIWithContext(context.Background(), request)
}

// ExportPrometheusReadOnlyDynamicAPI
// Prometheus 内部动态 api 代理，仅内部使用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATACOLUMNNOTFOUND = "FailedOperation.DataColumnNotFound"
//  FAILEDOPERATION_DATAQUERYFAILED = "FailedOperation.DataQueryFailed"
//  FAILEDOPERATION_DATATABLENOTFOUND = "FailedOperation.DataTableNotFound"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExportPrometheusReadOnlyDynamicAPIWithContext(ctx context.Context, request *ExportPrometheusReadOnlyDynamicAPIRequest) (response *ExportPrometheusReadOnlyDynamicAPIResponse, err error) {
    if request == nil {
        request = NewExportPrometheusReadOnlyDynamicAPIRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ExportPrometheusReadOnlyDynamicAPI")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportPrometheusReadOnlyDynamicAPI require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportPrometheusReadOnlyDynamicAPIResponse()
    err = c.Send(request, response)
    return
}

func NewGetMonitorDataRequest() (request *GetMonitorDataRequest) {
    request = &GetMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "GetMonitorData")
    
    
    return
}

func NewGetMonitorDataResponse() (response *GetMonitorDataResponse) {
    response = &GetMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetMonitorData
// 获取云产品的监控数据。此接口不适用于拉取容器服务监控数据，如需拉取容器服务监控数据，请使用[根据维度条件查询监控数据](https://cloud.tencent.com/document/product/248/51845)接口。
//
// 传入产品的命名空间、对象维度描述和监控指标即可获得相应的监控数据。
//
// 接口调用限制：单请求最多可支持批量拉取10个实例的监控数据，单请求的数据点数限制为1440个。
//
// 若您需要调用的指标、对象较多，可能存在因限频出现拉取失败的情况，建议尽量将请求按时间维度均摊。
//
// 拉取数据的粒度和统计方式的对应关系尽量在接入平台进行配置，如果没有配置对应统计方式，请提工单反馈。
//
// 
//
// >?
//
// >- 2022年9月1日起，腾讯云可观测平台开始对GetMonitorData接口计费。每个主账号每月可获得100万次免费请求额度，超过免费额度后如需继续调用接口需要开通 [API请求按量付费](https://buy.cloud.tencent.com/APIRequestBuy)。计费规则可查看[API计费文档](https://cloud.tencent.com/document/product/248/77914)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APITRANSLATEERROR = "FailedOperation.APITranslateError"
//  FAILEDOPERATION_ERRNOTOPEN = "FailedOperation.ErrNotOpen"
//  FAILEDOPERATION_ERROWED = "FailedOperation.ErrOwed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEDACCESS = "LimitExceeded.LimitedAccess"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetMonitorData(request *GetMonitorDataRequest) (response *GetMonitorDataResponse, err error) {
    return c.GetMonitorDataWithContext(context.Background(), request)
}

// GetMonitorData
// 获取云产品的监控数据。此接口不适用于拉取容器服务监控数据，如需拉取容器服务监控数据，请使用[根据维度条件查询监控数据](https://cloud.tencent.com/document/product/248/51845)接口。
//
// 传入产品的命名空间、对象维度描述和监控指标即可获得相应的监控数据。
//
// 接口调用限制：单请求最多可支持批量拉取10个实例的监控数据，单请求的数据点数限制为1440个。
//
// 若您需要调用的指标、对象较多，可能存在因限频出现拉取失败的情况，建议尽量将请求按时间维度均摊。
//
// 拉取数据的粒度和统计方式的对应关系尽量在接入平台进行配置，如果没有配置对应统计方式，请提工单反馈。
//
// 
//
// >?
//
// >- 2022年9月1日起，腾讯云可观测平台开始对GetMonitorData接口计费。每个主账号每月可获得100万次免费请求额度，超过免费额度后如需继续调用接口需要开通 [API请求按量付费](https://buy.cloud.tencent.com/APIRequestBuy)。计费规则可查看[API计费文档](https://cloud.tencent.com/document/product/248/77914)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APITRANSLATEERROR = "FailedOperation.APITranslateError"
//  FAILEDOPERATION_ERRNOTOPEN = "FailedOperation.ErrNotOpen"
//  FAILEDOPERATION_ERROWED = "FailedOperation.ErrOwed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEDACCESS = "LimitExceeded.LimitedAccess"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetMonitorDataWithContext(ctx context.Context, request *GetMonitorDataRequest) (response *GetMonitorDataResponse, err error) {
    if request == nil {
        request = NewGetMonitorDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "GetMonitorData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMonitorData require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMonitorDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetPrometheusAgentManagementCommandRequest() (request *GetPrometheusAgentManagementCommandRequest) {
    request = &GetPrometheusAgentManagementCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "GetPrometheusAgentManagementCommand")
    
    
    return
}

func NewGetPrometheusAgentManagementCommandResponse() (response *GetPrometheusAgentManagementCommandResponse) {
    response = &GetPrometheusAgentManagementCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetPrometheusAgentManagementCommand
// 获取 Prometheus Agent 管理相关的命令行
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) GetPrometheusAgentManagementCommand(request *GetPrometheusAgentManagementCommandRequest) (response *GetPrometheusAgentManagementCommandResponse, err error) {
    return c.GetPrometheusAgentManagementCommandWithContext(context.Background(), request)
}

// GetPrometheusAgentManagementCommand
// 获取 Prometheus Agent 管理相关的命令行
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) GetPrometheusAgentManagementCommandWithContext(ctx context.Context, request *GetPrometheusAgentManagementCommandRequest) (response *GetPrometheusAgentManagementCommandResponse, err error) {
    if request == nil {
        request = NewGetPrometheusAgentManagementCommandRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "GetPrometheusAgentManagementCommand")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPrometheusAgentManagementCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPrometheusAgentManagementCommandResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopNMonitorDataRequest() (request *GetTopNMonitorDataRequest) {
    request = &GetTopNMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "GetTopNMonitorData")
    
    
    return
}

func NewGetTopNMonitorDataResponse() (response *GetTopNMonitorDataResponse) {
    response = &GetTopNMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTopNMonitorData
// 支持TopN查询，对于给定的监控指标和时间区间，按照指标大小按序返回不同维度组合及数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRNOTOPEN = "FailedOperation.ErrNotOpen"
//  FAILEDOPERATION_ERROWED = "FailedOperation.ErrOwed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTopNMonitorData(request *GetTopNMonitorDataRequest) (response *GetTopNMonitorDataResponse, err error) {
    return c.GetTopNMonitorDataWithContext(context.Background(), request)
}

// GetTopNMonitorData
// 支持TopN查询，对于给定的监控指标和时间区间，按照指标大小按序返回不同维度组合及数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRNOTOPEN = "FailedOperation.ErrNotOpen"
//  FAILEDOPERATION_ERROWED = "FailedOperation.ErrOwed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTopNMonitorDataWithContext(ctx context.Context, request *GetTopNMonitorDataRequest) (response *GetTopNMonitorDataResponse, err error) {
    if request == nil {
        request = NewGetTopNMonitorDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "GetTopNMonitorData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTopNMonitorData require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTopNMonitorDataResponse()
    err = c.Send(request, response)
    return
}

func NewInstallPluginsRequest() (request *InstallPluginsRequest) {
    request = &InstallPluginsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "InstallPlugins")
    
    
    return
}

func NewInstallPluginsResponse() (response *InstallPluginsResponse) {
    response = &InstallPluginsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InstallPlugins
// 安装 Grafana Plugin
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) InstallPlugins(request *InstallPluginsRequest) (response *InstallPluginsResponse, err error) {
    return c.InstallPluginsWithContext(context.Background(), request)
}

// InstallPlugins
// 安装 Grafana Plugin
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) InstallPluginsWithContext(ctx context.Context, request *InstallPluginsRequest) (response *InstallPluginsResponse, err error) {
    if request == nil {
        request = NewInstallPluginsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "InstallPlugins")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstallPlugins require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstallPluginsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmNoticeRequest() (request *ModifyAlarmNoticeRequest) {
    request = &ModifyAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmNotice")
    
    
    return
}

func NewModifyAlarmNoticeResponse() (response *ModifyAlarmNoticeResponse) {
    response = &ModifyAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmNotice
// 编辑告警通知模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmNotice(request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
    return c.ModifyAlarmNoticeWithContext(context.Background(), request)
}

// ModifyAlarmNotice
// 编辑告警通知模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmNoticeWithContext(ctx context.Context, request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmNoticeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmNotice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyConditionRequest() (request *ModifyAlarmPolicyConditionRequest) {
    request = &ModifyAlarmPolicyConditionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyCondition")
    
    
    return
}

func NewModifyAlarmPolicyConditionResponse() (response *ModifyAlarmPolicyConditionResponse) {
    response = &ModifyAlarmPolicyConditionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmPolicyCondition
// 修改告警策略触发条件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyCondition(request *ModifyAlarmPolicyConditionRequest) (response *ModifyAlarmPolicyConditionResponse, err error) {
    return c.ModifyAlarmPolicyConditionWithContext(context.Background(), request)
}

// ModifyAlarmPolicyCondition
// 修改告警策略触发条件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyConditionWithContext(ctx context.Context, request *ModifyAlarmPolicyConditionRequest) (response *ModifyAlarmPolicyConditionResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyConditionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmPolicyCondition")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmPolicyCondition require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyConditionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyInfoRequest() (request *ModifyAlarmPolicyInfoRequest) {
    request = &ModifyAlarmPolicyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyInfo")
    
    
    return
}

func NewModifyAlarmPolicyInfoResponse() (response *ModifyAlarmPolicyInfoResponse) {
    response = &ModifyAlarmPolicyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmPolicyInfo
// 告警2.0编辑告警策略基本信息，包括策略名、备注
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyInfo(request *ModifyAlarmPolicyInfoRequest) (response *ModifyAlarmPolicyInfoResponse, err error) {
    return c.ModifyAlarmPolicyInfoWithContext(context.Background(), request)
}

// ModifyAlarmPolicyInfo
// 告警2.0编辑告警策略基本信息，包括策略名、备注
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyInfoWithContext(ctx context.Context, request *ModifyAlarmPolicyInfoRequest) (response *ModifyAlarmPolicyInfoResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmPolicyInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmPolicyInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyNoticeRequest() (request *ModifyAlarmPolicyNoticeRequest) {
    request = &ModifyAlarmPolicyNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyNotice")
    
    
    return
}

func NewModifyAlarmPolicyNoticeResponse() (response *ModifyAlarmPolicyNoticeResponse) {
    response = &ModifyAlarmPolicyNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmPolicyNotice
// 修改告警策略绑定的告警通知模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyNotice(request *ModifyAlarmPolicyNoticeRequest) (response *ModifyAlarmPolicyNoticeResponse, err error) {
    return c.ModifyAlarmPolicyNoticeWithContext(context.Background(), request)
}

// ModifyAlarmPolicyNotice
// 修改告警策略绑定的告警通知模板
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyNoticeWithContext(ctx context.Context, request *ModifyAlarmPolicyNoticeRequest) (response *ModifyAlarmPolicyNoticeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyNoticeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmPolicyNotice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmPolicyNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyStatusRequest() (request *ModifyAlarmPolicyStatusRequest) {
    request = &ModifyAlarmPolicyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyStatus")
    
    
    return
}

func NewModifyAlarmPolicyStatusResponse() (response *ModifyAlarmPolicyStatusResponse) {
    response = &ModifyAlarmPolicyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmPolicyStatus
// 启停告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyStatus(request *ModifyAlarmPolicyStatusRequest) (response *ModifyAlarmPolicyStatusResponse, err error) {
    return c.ModifyAlarmPolicyStatusWithContext(context.Background(), request)
}

// ModifyAlarmPolicyStatus
// 启停告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyStatusWithContext(ctx context.Context, request *ModifyAlarmPolicyStatusRequest) (response *ModifyAlarmPolicyStatusResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmPolicyStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmPolicyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyTasksRequest() (request *ModifyAlarmPolicyTasksRequest) {
    request = &ModifyAlarmPolicyTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyTasks")
    
    
    return
}

func NewModifyAlarmPolicyTasksResponse() (response *ModifyAlarmPolicyTasksResponse) {
    response = &ModifyAlarmPolicyTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmPolicyTasks
// 修改告警策略的触发任务，TriggerTasks字段放触发任务列表，TriggerTasks传空数组时，代表解绑该策略的所有触发任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmPolicyTasks(request *ModifyAlarmPolicyTasksRequest) (response *ModifyAlarmPolicyTasksResponse, err error) {
    return c.ModifyAlarmPolicyTasksWithContext(context.Background(), request)
}

// ModifyAlarmPolicyTasks
// 修改告警策略的触发任务，TriggerTasks字段放触发任务列表，TriggerTasks传空数组时，代表解绑该策略的所有触发任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmPolicyTasksWithContext(ctx context.Context, request *ModifyAlarmPolicyTasksRequest) (response *ModifyAlarmPolicyTasksResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmPolicyTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmPolicyTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyTasksResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmReceiversRequest() (request *ModifyAlarmReceiversRequest) {
    request = &ModifyAlarmReceiversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmReceivers")
    
    
    return
}

func NewModifyAlarmReceiversResponse() (response *ModifyAlarmReceiversResponse) {
    response = &ModifyAlarmReceiversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmReceivers
// 修改告警接收人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmReceivers(request *ModifyAlarmReceiversRequest) (response *ModifyAlarmReceiversResponse, err error) {
    return c.ModifyAlarmReceiversWithContext(context.Background(), request)
}

// ModifyAlarmReceivers
// 修改告警接收人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmReceiversWithContext(ctx context.Context, request *ModifyAlarmReceiversRequest) (response *ModifyAlarmReceiversResponse, err error) {
    if request == nil {
        request = NewModifyAlarmReceiversRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmReceivers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmReceivers require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmReceiversResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGrafanaInstanceRequest() (request *ModifyGrafanaInstanceRequest) {
    request = &ModifyGrafanaInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyGrafanaInstance")
    
    
    return
}

func NewModifyGrafanaInstanceResponse() (response *ModifyGrafanaInstanceResponse) {
    response = &ModifyGrafanaInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGrafanaInstance
// 修改 Grafana 实例属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) ModifyGrafanaInstance(request *ModifyGrafanaInstanceRequest) (response *ModifyGrafanaInstanceResponse, err error) {
    return c.ModifyGrafanaInstanceWithContext(context.Background(), request)
}

// ModifyGrafanaInstance
// 修改 Grafana 实例属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) ModifyGrafanaInstanceWithContext(ctx context.Context, request *ModifyGrafanaInstanceRequest) (response *ModifyGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewModifyGrafanaInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyGrafanaInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPolicyGroupRequest() (request *ModifyPolicyGroupRequest) {
    request = &ModifyPolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPolicyGroup")
    
    
    return
}

func NewModifyPolicyGroupResponse() (response *ModifyPolicyGroupResponse) {
    response = &ModifyPolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPolicyGroup
// 更新策略组
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPolicyGroup(request *ModifyPolicyGroupRequest) (response *ModifyPolicyGroupResponse, err error) {
    return c.ModifyPolicyGroupWithContext(context.Background(), request)
}

// ModifyPolicyGroup
// 更新策略组
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPolicyGroupWithContext(ctx context.Context, request *ModifyPolicyGroupRequest) (response *ModifyPolicyGroupResponse, err error) {
    if request == nil {
        request = NewModifyPolicyGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPolicyGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPolicyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusAgentExternalLabelsRequest() (request *ModifyPrometheusAgentExternalLabelsRequest) {
    request = &ModifyPrometheusAgentExternalLabelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusAgentExternalLabels")
    
    
    return
}

func NewModifyPrometheusAgentExternalLabelsResponse() (response *ModifyPrometheusAgentExternalLabelsResponse) {
    response = &ModifyPrometheusAgentExternalLabelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusAgentExternalLabels
// 修改被关联集群的external labels
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) ModifyPrometheusAgentExternalLabels(request *ModifyPrometheusAgentExternalLabelsRequest) (response *ModifyPrometheusAgentExternalLabelsResponse, err error) {
    return c.ModifyPrometheusAgentExternalLabelsWithContext(context.Background(), request)
}

// ModifyPrometheusAgentExternalLabels
// 修改被关联集群的external labels
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) ModifyPrometheusAgentExternalLabelsWithContext(ctx context.Context, request *ModifyPrometheusAgentExternalLabelsRequest) (response *ModifyPrometheusAgentExternalLabelsResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusAgentExternalLabelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusAgentExternalLabels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusAgentExternalLabels require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusAgentExternalLabelsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusAlertPolicyRequest() (request *ModifyPrometheusAlertPolicyRequest) {
    request = &ModifyPrometheusAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusAlertPolicy")
    
    
    return
}

func NewModifyPrometheusAlertPolicyResponse() (response *ModifyPrometheusAlertPolicyResponse) {
    response = &ModifyPrometheusAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusAlertPolicy
// 修改2.0实例告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusAlertPolicy(request *ModifyPrometheusAlertPolicyRequest) (response *ModifyPrometheusAlertPolicyResponse, err error) {
    return c.ModifyPrometheusAlertPolicyWithContext(context.Background(), request)
}

// ModifyPrometheusAlertPolicy
// 修改2.0实例告警策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusAlertPolicyWithContext(ctx context.Context, request *ModifyPrometheusAlertPolicyRequest) (response *ModifyPrometheusAlertPolicyResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusAlertPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusAlertPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusAlertPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusConfigRequest() (request *ModifyPrometheusConfigRequest) {
    request = &ModifyPrometheusConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusConfig")
    
    
    return
}

func NewModifyPrometheusConfigResponse() (response *ModifyPrometheusConfigResponse) {
    response = &ModifyPrometheusConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusConfig
// 修改prometheus采集配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPrometheusConfig(request *ModifyPrometheusConfigRequest) (response *ModifyPrometheusConfigResponse, err error) {
    return c.ModifyPrometheusConfigWithContext(context.Background(), request)
}

// ModifyPrometheusConfig
// 修改prometheus采集配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPrometheusConfigWithContext(ctx context.Context, request *ModifyPrometheusConfigRequest) (response *ModifyPrometheusConfigResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusGlobalNotificationRequest() (request *ModifyPrometheusGlobalNotificationRequest) {
    request = &ModifyPrometheusGlobalNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusGlobalNotification")
    
    
    return
}

func NewModifyPrometheusGlobalNotificationResponse() (response *ModifyPrometheusGlobalNotificationResponse) {
    response = &ModifyPrometheusGlobalNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusGlobalNotification
// 修改全局告警通知渠道
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusGlobalNotification(request *ModifyPrometheusGlobalNotificationRequest) (response *ModifyPrometheusGlobalNotificationResponse, err error) {
    return c.ModifyPrometheusGlobalNotificationWithContext(context.Background(), request)
}

// ModifyPrometheusGlobalNotification
// 修改全局告警通知渠道
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusGlobalNotificationWithContext(ctx context.Context, request *ModifyPrometheusGlobalNotificationRequest) (response *ModifyPrometheusGlobalNotificationResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusGlobalNotificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusGlobalNotification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusGlobalNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusGlobalNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusInstanceAttributesRequest() (request *ModifyPrometheusInstanceAttributesRequest) {
    request = &ModifyPrometheusInstanceAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusInstanceAttributes")
    
    
    return
}

func NewModifyPrometheusInstanceAttributesResponse() (response *ModifyPrometheusInstanceAttributesResponse) {
    response = &ModifyPrometheusInstanceAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusInstanceAttributes
// 修改 Prometheus 实例相关属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyPrometheusInstanceAttributes(request *ModifyPrometheusInstanceAttributesRequest) (response *ModifyPrometheusInstanceAttributesResponse, err error) {
    return c.ModifyPrometheusInstanceAttributesWithContext(context.Background(), request)
}

// ModifyPrometheusInstanceAttributes
// 修改 Prometheus 实例相关属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyPrometheusInstanceAttributesWithContext(ctx context.Context, request *ModifyPrometheusInstanceAttributesRequest) (response *ModifyPrometheusInstanceAttributesResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusInstanceAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusInstanceAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusInstanceAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusInstanceAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusRecordRuleYamlRequest() (request *ModifyPrometheusRecordRuleYamlRequest) {
    request = &ModifyPrometheusRecordRuleYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusRecordRuleYaml")
    
    
    return
}

func NewModifyPrometheusRecordRuleYamlResponse() (response *ModifyPrometheusRecordRuleYamlResponse) {
    response = &ModifyPrometheusRecordRuleYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusRecordRuleYaml
// 通过yaml的方式修改Prometheus预聚合规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusRecordRuleYaml(request *ModifyPrometheusRecordRuleYamlRequest) (response *ModifyPrometheusRecordRuleYamlResponse, err error) {
    return c.ModifyPrometheusRecordRuleYamlWithContext(context.Background(), request)
}

// ModifyPrometheusRecordRuleYaml
// 通过yaml的方式修改Prometheus预聚合规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusRecordRuleYamlWithContext(ctx context.Context, request *ModifyPrometheusRecordRuleYamlRequest) (response *ModifyPrometheusRecordRuleYamlResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusRecordRuleYamlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusRecordRuleYaml")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusRecordRuleYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusRecordRuleYamlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusTempRequest() (request *ModifyPrometheusTempRequest) {
    request = &ModifyPrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusTemp")
    
    
    return
}

func NewModifyPrometheusTempResponse() (response *ModifyPrometheusTempResponse) {
    response = &ModifyPrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusTemp
// 修改模板内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPrometheusTemp(request *ModifyPrometheusTempRequest) (response *ModifyPrometheusTempResponse, err error) {
    return c.ModifyPrometheusTempWithContext(context.Background(), request)
}

// ModifyPrometheusTemp
// 修改模板内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPrometheusTempWithContext(ctx context.Context, request *ModifyPrometheusTempRequest) (response *ModifyPrometheusTempResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusTempRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusTemp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRemoteURLsRequest() (request *ModifyRemoteURLsRequest) {
    request = &ModifyRemoteURLsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyRemoteURLs")
    
    
    return
}

func NewModifyRemoteURLsResponse() (response *ModifyRemoteURLsResponse) {
    response = &ModifyRemoteURLsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRemoteURLs
// 修改多写配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_RESOURCECONFLICT = "FailedOperation.ResourceConflict"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) ModifyRemoteURLs(request *ModifyRemoteURLsRequest) (response *ModifyRemoteURLsResponse, err error) {
    return c.ModifyRemoteURLsWithContext(context.Background(), request)
}

// ModifyRemoteURLs
// 修改多写配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_RESOURCECONFLICT = "FailedOperation.ResourceConflict"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) ModifyRemoteURLsWithContext(ctx context.Context, request *ModifyRemoteURLsRequest) (response *ModifyRemoteURLsResponse, err error) {
    if request == nil {
        request = NewModifyRemoteURLsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyRemoteURLs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRemoteURLs require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRemoteURLsResponse()
    err = c.Send(request, response)
    return
}

func NewResumeGrafanaInstanceRequest() (request *ResumeGrafanaInstanceRequest) {
    request = &ResumeGrafanaInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ResumeGrafanaInstance")
    
    
    return
}

func NewResumeGrafanaInstanceResponse() (response *ResumeGrafanaInstanceResponse) {
    response = &ResumeGrafanaInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeGrafanaInstance
// 本接口（ResumeGrafanaInstance）用于 Grafana 包年包月实例的停服续费，调用后按原版本续费一个月。仍在运行中的实例无法使用该接口进行续费。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_RESOURCECONFLICT = "FailedOperation.ResourceConflict"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) ResumeGrafanaInstance(request *ResumeGrafanaInstanceRequest) (response *ResumeGrafanaInstanceResponse, err error) {
    return c.ResumeGrafanaInstanceWithContext(context.Background(), request)
}

// ResumeGrafanaInstance
// 本接口（ResumeGrafanaInstance）用于 Grafana 包年包月实例的停服续费，调用后按原版本续费一个月。仍在运行中的实例无法使用该接口进行续费。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_RESOURCECONFLICT = "FailedOperation.ResourceConflict"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) ResumeGrafanaInstanceWithContext(ctx context.Context, request *ResumeGrafanaInstanceRequest) (response *ResumeGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewResumeGrafanaInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ResumeGrafanaInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRunPrometheusInstanceRequest() (request *RunPrometheusInstanceRequest) {
    request = &RunPrometheusInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "RunPrometheusInstance")
    
    
    return
}

func NewRunPrometheusInstanceResponse() (response *RunPrometheusInstanceResponse) {
    response = &RunPrometheusInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunPrometheusInstance
// 初始化TMP实例，开启集成中心时调用
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) RunPrometheusInstance(request *RunPrometheusInstanceRequest) (response *RunPrometheusInstanceResponse, err error) {
    return c.RunPrometheusInstanceWithContext(context.Background(), request)
}

// RunPrometheusInstance
// 初始化TMP实例，开启集成中心时调用
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) RunPrometheusInstanceWithContext(ctx context.Context, request *RunPrometheusInstanceRequest) (response *RunPrometheusInstanceResponse, err error) {
    if request == nil {
        request = NewRunPrometheusInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "RunPrometheusInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunPrometheusInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunPrometheusInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewSetDefaultAlarmPolicyRequest() (request *SetDefaultAlarmPolicyRequest) {
    request = &SetDefaultAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "SetDefaultAlarmPolicy")
    
    
    return
}

func NewSetDefaultAlarmPolicyResponse() (response *SetDefaultAlarmPolicyResponse) {
    response = &SetDefaultAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetDefaultAlarmPolicy
// 设置一个策略为该告警策略类型、该项目的默认告警策略。
//
// 同一项目下相同的告警策略类型，就会被设置为非默认。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetDefaultAlarmPolicy(request *SetDefaultAlarmPolicyRequest) (response *SetDefaultAlarmPolicyResponse, err error) {
    return c.SetDefaultAlarmPolicyWithContext(context.Background(), request)
}

// SetDefaultAlarmPolicy
// 设置一个策略为该告警策略类型、该项目的默认告警策略。
//
// 同一项目下相同的告警策略类型，就会被设置为非默认。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetDefaultAlarmPolicyWithContext(ctx context.Context, request *SetDefaultAlarmPolicyRequest) (response *SetDefaultAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewSetDefaultAlarmPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "SetDefaultAlarmPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetDefaultAlarmPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetDefaultAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewSyncPrometheusTempRequest() (request *SyncPrometheusTempRequest) {
    request = &SyncPrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "SyncPrometheusTemp")
    
    
    return
}

func NewSyncPrometheusTempResponse() (response *SyncPrometheusTempResponse) {
    response = &SyncPrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncPrometheusTemp
// 同步模板到实例或者集群，针对V2版本实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncPrometheusTemp(request *SyncPrometheusTempRequest) (response *SyncPrometheusTempResponse, err error) {
    return c.SyncPrometheusTempWithContext(context.Background(), request)
}

// SyncPrometheusTemp
// 同步模板到实例或者集群，针对V2版本实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncPrometheusTempWithContext(ctx context.Context, request *SyncPrometheusTempRequest) (response *SyncPrometheusTempResponse, err error) {
    if request == nil {
        request = NewSyncPrometheusTempRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "SyncPrometheusTemp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncPrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncPrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewTerminatePrometheusInstancesRequest() (request *TerminatePrometheusInstancesRequest) {
    request = &TerminatePrometheusInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "TerminatePrometheusInstances")
    
    
    return
}

func NewTerminatePrometheusInstancesResponse() (response *TerminatePrometheusInstancesResponse) {
    response = &TerminatePrometheusInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminatePrometheusInstances
// 销毁按量 Prometheus 实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTSNOTINUNINSTALLSTAGE = "FailedOperation.AgentsNotInUninstallStage"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) TerminatePrometheusInstances(request *TerminatePrometheusInstancesRequest) (response *TerminatePrometheusInstancesResponse, err error) {
    return c.TerminatePrometheusInstancesWithContext(context.Background(), request)
}

// TerminatePrometheusInstances
// 销毁按量 Prometheus 实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTSNOTINUNINSTALLSTAGE = "FailedOperation.AgentsNotInUninstallStage"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) TerminatePrometheusInstancesWithContext(ctx context.Context, request *TerminatePrometheusInstancesRequest) (response *TerminatePrometheusInstancesResponse, err error) {
    if request == nil {
        request = NewTerminatePrometheusInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "TerminatePrometheusInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminatePrometheusInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminatePrometheusInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindingAllPolicyObjectRequest() (request *UnBindingAllPolicyObjectRequest) {
    request = &UnBindingAllPolicyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UnBindingAllPolicyObject")
    
    
    return
}

func NewUnBindingAllPolicyObjectResponse() (response *UnBindingAllPolicyObjectResponse) {
    response = &UnBindingAllPolicyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnBindingAllPolicyObject
// 删除全部的关联对象
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnBindingAllPolicyObject(request *UnBindingAllPolicyObjectRequest) (response *UnBindingAllPolicyObjectResponse, err error) {
    return c.UnBindingAllPolicyObjectWithContext(context.Background(), request)
}

// UnBindingAllPolicyObject
// 删除全部的关联对象
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnBindingAllPolicyObjectWithContext(ctx context.Context, request *UnBindingAllPolicyObjectRequest) (response *UnBindingAllPolicyObjectResponse, err error) {
    if request == nil {
        request = NewUnBindingAllPolicyObjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UnBindingAllPolicyObject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindingAllPolicyObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindingAllPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindingPolicyObjectRequest() (request *UnBindingPolicyObjectRequest) {
    request = &UnBindingPolicyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UnBindingPolicyObject")
    
    
    return
}

func NewUnBindingPolicyObjectResponse() (response *UnBindingPolicyObjectResponse) {
    response = &UnBindingPolicyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnBindingPolicyObject
// 删除策略的关联对象
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnBindingPolicyObject(request *UnBindingPolicyObjectRequest) (response *UnBindingPolicyObjectResponse, err error) {
    return c.UnBindingPolicyObjectWithContext(context.Background(), request)
}

// UnBindingPolicyObject
// 删除策略的关联对象
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnBindingPolicyObjectWithContext(ctx context.Context, request *UnBindingPolicyObjectRequest) (response *UnBindingPolicyObjectResponse, err error) {
    if request == nil {
        request = NewUnBindingPolicyObjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UnBindingPolicyObject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindingPolicyObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindingPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindPrometheusManagedGrafanaRequest() (request *UnbindPrometheusManagedGrafanaRequest) {
    request = &UnbindPrometheusManagedGrafanaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UnbindPrometheusManagedGrafana")
    
    
    return
}

func NewUnbindPrometheusManagedGrafanaResponse() (response *UnbindPrometheusManagedGrafanaResponse) {
    response = &UnbindPrometheusManagedGrafanaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindPrometheusManagedGrafana
// 解除实例绑定的 Grafana 可视化实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UnbindPrometheusManagedGrafana(request *UnbindPrometheusManagedGrafanaRequest) (response *UnbindPrometheusManagedGrafanaResponse, err error) {
    return c.UnbindPrometheusManagedGrafanaWithContext(context.Background(), request)
}

// UnbindPrometheusManagedGrafana
// 解除实例绑定的 Grafana 可视化实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UnbindPrometheusManagedGrafanaWithContext(ctx context.Context, request *UnbindPrometheusManagedGrafanaRequest) (response *UnbindPrometheusManagedGrafanaResponse, err error) {
    if request == nil {
        request = NewUnbindPrometheusManagedGrafanaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UnbindPrometheusManagedGrafana")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindPrometheusManagedGrafana require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindPrometheusManagedGrafanaResponse()
    err = c.Send(request, response)
    return
}

func NewUninstallGrafanaDashboardRequest() (request *UninstallGrafanaDashboardRequest) {
    request = &UninstallGrafanaDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UninstallGrafanaDashboard")
    
    
    return
}

func NewUninstallGrafanaDashboardResponse() (response *UninstallGrafanaDashboardResponse) {
    response = &UninstallGrafanaDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UninstallGrafanaDashboard
// 删除 Grafana Dashboard
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UninstallGrafanaDashboard(request *UninstallGrafanaDashboardRequest) (response *UninstallGrafanaDashboardResponse, err error) {
    return c.UninstallGrafanaDashboardWithContext(context.Background(), request)
}

// UninstallGrafanaDashboard
// 删除 Grafana Dashboard
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UninstallGrafanaDashboardWithContext(ctx context.Context, request *UninstallGrafanaDashboardRequest) (response *UninstallGrafanaDashboardResponse, err error) {
    if request == nil {
        request = NewUninstallGrafanaDashboardRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UninstallGrafanaDashboard")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UninstallGrafanaDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewUninstallGrafanaDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewUninstallGrafanaPluginsRequest() (request *UninstallGrafanaPluginsRequest) {
    request = &UninstallGrafanaPluginsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UninstallGrafanaPlugins")
    
    
    return
}

func NewUninstallGrafanaPluginsResponse() (response *UninstallGrafanaPluginsResponse) {
    response = &UninstallGrafanaPluginsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UninstallGrafanaPlugins
// 删除已安装的插件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UninstallGrafanaPlugins(request *UninstallGrafanaPluginsRequest) (response *UninstallGrafanaPluginsResponse, err error) {
    return c.UninstallGrafanaPluginsWithContext(context.Background(), request)
}

// UninstallGrafanaPlugins
// 删除已安装的插件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UninstallGrafanaPluginsWithContext(ctx context.Context, request *UninstallGrafanaPluginsRequest) (response *UninstallGrafanaPluginsResponse, err error) {
    if request == nil {
        request = NewUninstallGrafanaPluginsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UninstallGrafanaPlugins")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UninstallGrafanaPlugins require credential")
    }

    request.SetContext(ctx)
    
    response = NewUninstallGrafanaPluginsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAlertRuleRequest() (request *UpdateAlertRuleRequest) {
    request = &UpdateAlertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateAlertRule")
    
    
    return
}

func NewUpdateAlertRuleResponse() (response *UpdateAlertRuleResponse) {
    response = &UpdateAlertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAlertRule
// 更新 Prometheus 的告警规则。
//
// 
//
// 请注意，**告警对象和告警消息是 Prometheus Rule Annotations 的特殊字段，需要通过 annotations 来传递，对应的 Key 分别为summary/description**，请参考 [Prometheus Rule更多配置请参考](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/)。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateAlertRule(request *UpdateAlertRuleRequest) (response *UpdateAlertRuleResponse, err error) {
    return c.UpdateAlertRuleWithContext(context.Background(), request)
}

// UpdateAlertRule
// 更新 Prometheus 的告警规则。
//
// 
//
// 请注意，**告警对象和告警消息是 Prometheus Rule Annotations 的特殊字段，需要通过 annotations 来传递，对应的 Key 分别为summary/description**，请参考 [Prometheus Rule更多配置请参考](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/)。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateAlertRuleWithContext(ctx context.Context, request *UpdateAlertRuleRequest) (response *UpdateAlertRuleResponse, err error) {
    if request == nil {
        request = NewUpdateAlertRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateAlertRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAlertRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAlertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAlertRuleStateRequest() (request *UpdateAlertRuleStateRequest) {
    request = &UpdateAlertRuleStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateAlertRuleState")
    
    
    return
}

func NewUpdateAlertRuleStateResponse() (response *UpdateAlertRuleStateResponse) {
    response = &UpdateAlertRuleStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAlertRuleState
// 更新 Prometheus 报警策略状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateAlertRuleState(request *UpdateAlertRuleStateRequest) (response *UpdateAlertRuleStateResponse, err error) {
    return c.UpdateAlertRuleStateWithContext(context.Background(), request)
}

// UpdateAlertRuleState
// 更新 Prometheus 报警策略状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateAlertRuleStateWithContext(ctx context.Context, request *UpdateAlertRuleStateRequest) (response *UpdateAlertRuleStateResponse, err error) {
    if request == nil {
        request = NewUpdateAlertRuleStateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateAlertRuleState")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAlertRuleState require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAlertRuleStateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDNSConfigRequest() (request *UpdateDNSConfigRequest) {
    request = &UpdateDNSConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateDNSConfig")
    
    
    return
}

func NewUpdateDNSConfigResponse() (response *UpdateDNSConfigResponse) {
    response = &UpdateDNSConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDNSConfig
// 更新 Grafana 的 DNS 配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateDNSConfig(request *UpdateDNSConfigRequest) (response *UpdateDNSConfigResponse, err error) {
    return c.UpdateDNSConfigWithContext(context.Background(), request)
}

// UpdateDNSConfig
// 更新 Grafana 的 DNS 配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateDNSConfigWithContext(ctx context.Context, request *UpdateDNSConfigRequest) (response *UpdateDNSConfigResponse, err error) {
    if request == nil {
        request = NewUpdateDNSConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateDNSConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDNSConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDNSConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateExporterIntegrationRequest() (request *UpdateExporterIntegrationRequest) {
    request = &UpdateExporterIntegrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateExporterIntegration")
    
    
    return
}

func NewUpdateExporterIntegrationResponse() (response *UpdateExporterIntegrationResponse) {
    response = &UpdateExporterIntegrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateExporterIntegration
// 更新 exporter 集成配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTNOTALLOWED = "FailedOperation.AgentNotAllowed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateExporterIntegration(request *UpdateExporterIntegrationRequest) (response *UpdateExporterIntegrationResponse, err error) {
    return c.UpdateExporterIntegrationWithContext(context.Background(), request)
}

// UpdateExporterIntegration
// 更新 exporter 集成配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTNOTALLOWED = "FailedOperation.AgentNotAllowed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateExporterIntegrationWithContext(ctx context.Context, request *UpdateExporterIntegrationRequest) (response *UpdateExporterIntegrationResponse, err error) {
    if request == nil {
        request = NewUpdateExporterIntegrationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateExporterIntegration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateExporterIntegration require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateExporterIntegrationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGrafanaConfigRequest() (request *UpdateGrafanaConfigRequest) {
    request = &UpdateGrafanaConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateGrafanaConfig")
    
    
    return
}

func NewUpdateGrafanaConfigResponse() (response *UpdateGrafanaConfigResponse) {
    response = &UpdateGrafanaConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGrafanaConfig
// 更新 Grafana 配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaConfig(request *UpdateGrafanaConfigRequest) (response *UpdateGrafanaConfigResponse, err error) {
    return c.UpdateGrafanaConfigWithContext(context.Background(), request)
}

// UpdateGrafanaConfig
// 更新 Grafana 配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaConfigWithContext(ctx context.Context, request *UpdateGrafanaConfigRequest) (response *UpdateGrafanaConfigResponse, err error) {
    if request == nil {
        request = NewUpdateGrafanaConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateGrafanaConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGrafanaConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGrafanaConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGrafanaEnvironmentsRequest() (request *UpdateGrafanaEnvironmentsRequest) {
    request = &UpdateGrafanaEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateGrafanaEnvironments")
    
    
    return
}

func NewUpdateGrafanaEnvironmentsResponse() (response *UpdateGrafanaEnvironmentsResponse) {
    response = &UpdateGrafanaEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGrafanaEnvironments
// 更新 Grafana 环境变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaEnvironments(request *UpdateGrafanaEnvironmentsRequest) (response *UpdateGrafanaEnvironmentsResponse, err error) {
    return c.UpdateGrafanaEnvironmentsWithContext(context.Background(), request)
}

// UpdateGrafanaEnvironments
// 更新 Grafana 环境变量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaEnvironmentsWithContext(ctx context.Context, request *UpdateGrafanaEnvironmentsRequest) (response *UpdateGrafanaEnvironmentsResponse, err error) {
    if request == nil {
        request = NewUpdateGrafanaEnvironmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateGrafanaEnvironments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGrafanaEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGrafanaEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGrafanaIntegrationRequest() (request *UpdateGrafanaIntegrationRequest) {
    request = &UpdateGrafanaIntegrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateGrafanaIntegration")
    
    
    return
}

func NewUpdateGrafanaIntegrationResponse() (response *UpdateGrafanaIntegrationResponse) {
    response = &UpdateGrafanaIntegrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGrafanaIntegration
// 更新 Grafana 集成配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaIntegration(request *UpdateGrafanaIntegrationRequest) (response *UpdateGrafanaIntegrationResponse, err error) {
    return c.UpdateGrafanaIntegrationWithContext(context.Background(), request)
}

// UpdateGrafanaIntegration
// 更新 Grafana 集成配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaIntegrationWithContext(ctx context.Context, request *UpdateGrafanaIntegrationRequest) (response *UpdateGrafanaIntegrationResponse, err error) {
    if request == nil {
        request = NewUpdateGrafanaIntegrationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateGrafanaIntegration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGrafanaIntegration require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGrafanaIntegrationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGrafanaNotificationChannelRequest() (request *UpdateGrafanaNotificationChannelRequest) {
    request = &UpdateGrafanaNotificationChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateGrafanaNotificationChannel")
    
    
    return
}

func NewUpdateGrafanaNotificationChannelResponse() (response *UpdateGrafanaNotificationChannelResponse) {
    response = &UpdateGrafanaNotificationChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGrafanaNotificationChannel
// 更新 Grafana 告警通道
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaNotificationChannel(request *UpdateGrafanaNotificationChannelRequest) (response *UpdateGrafanaNotificationChannelResponse, err error) {
    return c.UpdateGrafanaNotificationChannelWithContext(context.Background(), request)
}

// UpdateGrafanaNotificationChannel
// 更新 Grafana 告警通道
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaNotificationChannelWithContext(ctx context.Context, request *UpdateGrafanaNotificationChannelRequest) (response *UpdateGrafanaNotificationChannelResponse, err error) {
    if request == nil {
        request = NewUpdateGrafanaNotificationChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateGrafanaNotificationChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGrafanaNotificationChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGrafanaNotificationChannelResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGrafanaWhiteListRequest() (request *UpdateGrafanaWhiteListRequest) {
    request = &UpdateGrafanaWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateGrafanaWhiteList")
    
    
    return
}

func NewUpdateGrafanaWhiteListResponse() (response *UpdateGrafanaWhiteListResponse) {
    response = &UpdateGrafanaWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGrafanaWhiteList
// 更新 Grafana 白名单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaWhiteList(request *UpdateGrafanaWhiteListRequest) (response *UpdateGrafanaWhiteListResponse, err error) {
    return c.UpdateGrafanaWhiteListWithContext(context.Background(), request)
}

// UpdateGrafanaWhiteList
// 更新 Grafana 白名单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaWhiteListWithContext(ctx context.Context, request *UpdateGrafanaWhiteListRequest) (response *UpdateGrafanaWhiteListResponse, err error) {
    if request == nil {
        request = NewUpdateGrafanaWhiteListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateGrafanaWhiteList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGrafanaWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGrafanaWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePrometheusAgentStatusRequest() (request *UpdatePrometheusAgentStatusRequest) {
    request = &UpdatePrometheusAgentStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdatePrometheusAgentStatus")
    
    
    return
}

func NewUpdatePrometheusAgentStatusResponse() (response *UpdatePrometheusAgentStatusResponse) {
    response = &UpdatePrometheusAgentStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePrometheusAgentStatus
// 更新 Prometheus CVM Agent 状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdatePrometheusAgentStatus(request *UpdatePrometheusAgentStatusRequest) (response *UpdatePrometheusAgentStatusResponse, err error) {
    return c.UpdatePrometheusAgentStatusWithContext(context.Background(), request)
}

// UpdatePrometheusAgentStatus
// 更新 Prometheus CVM Agent 状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdatePrometheusAgentStatusWithContext(ctx context.Context, request *UpdatePrometheusAgentStatusRequest) (response *UpdatePrometheusAgentStatusResponse, err error) {
    if request == nil {
        request = NewUpdatePrometheusAgentStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdatePrometheusAgentStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePrometheusAgentStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePrometheusAgentStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePrometheusAlertGroupRequest() (request *UpdatePrometheusAlertGroupRequest) {
    request = &UpdatePrometheusAlertGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdatePrometheusAlertGroup")
    
    
    return
}

func NewUpdatePrometheusAlertGroupResponse() (response *UpdatePrometheusAlertGroupResponse) {
    response = &UpdatePrometheusAlertGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePrometheusAlertGroup
// 更新Prometheus告警规则分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdatePrometheusAlertGroup(request *UpdatePrometheusAlertGroupRequest) (response *UpdatePrometheusAlertGroupResponse, err error) {
    return c.UpdatePrometheusAlertGroupWithContext(context.Background(), request)
}

// UpdatePrometheusAlertGroup
// 更新Prometheus告警规则分组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdatePrometheusAlertGroupWithContext(ctx context.Context, request *UpdatePrometheusAlertGroupRequest) (response *UpdatePrometheusAlertGroupResponse, err error) {
    if request == nil {
        request = NewUpdatePrometheusAlertGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdatePrometheusAlertGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePrometheusAlertGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePrometheusAlertGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePrometheusAlertGroupStateRequest() (request *UpdatePrometheusAlertGroupStateRequest) {
    request = &UpdatePrometheusAlertGroupStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdatePrometheusAlertGroupState")
    
    
    return
}

func NewUpdatePrometheusAlertGroupStateResponse() (response *UpdatePrometheusAlertGroupStateResponse) {
    response = &UpdatePrometheusAlertGroupStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePrometheusAlertGroupState
// 批量更新告警分组状态，将分组中全部告警规则更新为目标状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdatePrometheusAlertGroupState(request *UpdatePrometheusAlertGroupStateRequest) (response *UpdatePrometheusAlertGroupStateResponse, err error) {
    return c.UpdatePrometheusAlertGroupStateWithContext(context.Background(), request)
}

// UpdatePrometheusAlertGroupState
// 批量更新告警分组状态，将分组中全部告警规则更新为目标状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdatePrometheusAlertGroupStateWithContext(ctx context.Context, request *UpdatePrometheusAlertGroupStateRequest) (response *UpdatePrometheusAlertGroupStateResponse, err error) {
    if request == nil {
        request = NewUpdatePrometheusAlertGroupStateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdatePrometheusAlertGroupState")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePrometheusAlertGroupState require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePrometheusAlertGroupStateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePrometheusScrapeJobRequest() (request *UpdatePrometheusScrapeJobRequest) {
    request = &UpdatePrometheusScrapeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdatePrometheusScrapeJob")
    
    
    return
}

func NewUpdatePrometheusScrapeJobResponse() (response *UpdatePrometheusScrapeJobResponse) {
    response = &UpdatePrometheusScrapeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePrometheusScrapeJob
// 更新 Prometheus Agent 抓取任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdatePrometheusScrapeJob(request *UpdatePrometheusScrapeJobRequest) (response *UpdatePrometheusScrapeJobResponse, err error) {
    return c.UpdatePrometheusScrapeJobWithContext(context.Background(), request)
}

// UpdatePrometheusScrapeJob
// 更新 Prometheus Agent 抓取任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdatePrometheusScrapeJobWithContext(ctx context.Context, request *UpdatePrometheusScrapeJobRequest) (response *UpdatePrometheusScrapeJobResponse, err error) {
    if request == nil {
        request = NewUpdatePrometheusScrapeJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdatePrometheusScrapeJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePrometheusScrapeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePrometheusScrapeJobResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRecordingRuleRequest() (request *UpdateRecordingRuleRequest) {
    request = &UpdateRecordingRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateRecordingRule")
    
    
    return
}

func NewUpdateRecordingRuleResponse() (response *UpdateRecordingRuleResponse) {
    response = &UpdateRecordingRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRecordingRule
// 更新 Prometheus 的预聚合规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateRecordingRule(request *UpdateRecordingRuleRequest) (response *UpdateRecordingRuleResponse, err error) {
    return c.UpdateRecordingRuleWithContext(context.Background(), request)
}

// UpdateRecordingRule
// 更新 Prometheus 的预聚合规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateRecordingRuleWithContext(ctx context.Context, request *UpdateRecordingRuleRequest) (response *UpdateRecordingRuleResponse, err error) {
    if request == nil {
        request = NewUpdateRecordingRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateRecordingRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRecordingRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRecordingRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSSOAccountRequest() (request *UpdateSSOAccountRequest) {
    request = &UpdateSSOAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateSSOAccount")
    
    
    return
}

func NewUpdateSSOAccountResponse() (response *UpdateSSOAccountResponse) {
    response = &UpdateSSOAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSSOAccount
// 更新已授权账号的备注、权限信息，会直接覆盖原有的信息，不传则不会更新。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UpdateSSOAccount(request *UpdateSSOAccountRequest) (response *UpdateSSOAccountResponse, err error) {
    return c.UpdateSSOAccountWithContext(context.Background(), request)
}

// UpdateSSOAccount
// 更新已授权账号的备注、权限信息，会直接覆盖原有的信息，不传则不会更新。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UpdateSSOAccountWithContext(ctx context.Context, request *UpdateSSOAccountRequest) (response *UpdateSSOAccountResponse, err error) {
    if request == nil {
        request = NewUpdateSSOAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateSSOAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSSOAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSSOAccountResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateServiceDiscoveryRequest() (request *UpdateServiceDiscoveryRequest) {
    request = &UpdateServiceDiscoveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateServiceDiscovery")
    
    
    return
}

func NewUpdateServiceDiscoveryResponse() (response *UpdateServiceDiscoveryResponse) {
    response = &UpdateServiceDiscoveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateServiceDiscovery
// 在腾讯云容器服务下更新 Prometheus 服务发现。
//
// <p>注意：前提条件，已经通过 Prometheus 控制台集成了对应的腾讯云容器服务，具体请参考
//
// <a href="https://cloud.tencent.com/document/product/248/48859" target="_blank">Agent 安装</a>。</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateServiceDiscovery(request *UpdateServiceDiscoveryRequest) (response *UpdateServiceDiscoveryResponse, err error) {
    return c.UpdateServiceDiscoveryWithContext(context.Background(), request)
}

// UpdateServiceDiscovery
// 在腾讯云容器服务下更新 Prometheus 服务发现。
//
// <p>注意：前提条件，已经通过 Prometheus 控制台集成了对应的腾讯云容器服务，具体请参考
//
// <a href="https://cloud.tencent.com/document/product/248/48859" target="_blank">Agent 安装</a>。</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateServiceDiscoveryWithContext(ctx context.Context, request *UpdateServiceDiscoveryRequest) (response *UpdateServiceDiscoveryResponse, err error) {
    if request == nil {
        request = NewUpdateServiceDiscoveryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateServiceDiscovery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateServiceDiscovery require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateServiceDiscoveryResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeGrafanaDashboardRequest() (request *UpgradeGrafanaDashboardRequest) {
    request = &UpgradeGrafanaDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpgradeGrafanaDashboard")
    
    
    return
}

func NewUpgradeGrafanaDashboardResponse() (response *UpgradeGrafanaDashboardResponse) {
    response = &UpgradeGrafanaDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeGrafanaDashboard
// 升级 Grafana Dashboard
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpgradeGrafanaDashboard(request *UpgradeGrafanaDashboardRequest) (response *UpgradeGrafanaDashboardResponse, err error) {
    return c.UpgradeGrafanaDashboardWithContext(context.Background(), request)
}

// UpgradeGrafanaDashboard
// 升级 Grafana Dashboard
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpgradeGrafanaDashboardWithContext(ctx context.Context, request *UpgradeGrafanaDashboardRequest) (response *UpgradeGrafanaDashboardResponse, err error) {
    if request == nil {
        request = NewUpgradeGrafanaDashboardRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpgradeGrafanaDashboard")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeGrafanaDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeGrafanaDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeGrafanaInstanceRequest() (request *UpgradeGrafanaInstanceRequest) {
    request = &UpgradeGrafanaInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpgradeGrafanaInstance")
    
    
    return
}

func NewUpgradeGrafanaInstanceResponse() (response *UpgradeGrafanaInstanceResponse) {
    response = &UpgradeGrafanaInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeGrafanaInstance
// 升级 Grafana 实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpgradeGrafanaInstance(request *UpgradeGrafanaInstanceRequest) (response *UpgradeGrafanaInstanceResponse, err error) {
    return c.UpgradeGrafanaInstanceWithContext(context.Background(), request)
}

// UpgradeGrafanaInstance
// 升级 Grafana 实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpgradeGrafanaInstanceWithContext(ctx context.Context, request *UpgradeGrafanaInstanceRequest) (response *UpgradeGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeGrafanaInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpgradeGrafanaInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}
