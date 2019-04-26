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

package v20180801

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-08-01"

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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAddDelayLiveStreamRequest() (request *AddDelayLiveStreamRequest) {
    request = &AddDelayLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "AddDelayLiveStream")
    return
}

func NewAddDelayLiveStreamResponse() (response *AddDelayLiveStreamResponse) {
    response = &AddDelayLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对流设置延播时间
// 注意：如果在推流前设置延播，需要提前5分钟设置。
// 目前该接口只支持流粒度的，域名及应用粒度功能支持当前开发中。
func (c *Client) AddDelayLiveStream(request *AddDelayLiveStreamRequest) (response *AddDelayLiveStreamResponse, err error) {
    if request == nil {
        request = NewAddDelayLiveStreamRequest()
    }
    response = NewAddDelayLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewAddLiveDomainRequest() (request *AddLiveDomainRequest) {
    request = &AddLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "AddLiveDomain")
    return
}

func NewAddLiveDomainResponse() (response *AddLiveDomainResponse) {
    response = &AddLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加域名，一次只能提交一个域名。域名必须已备案。
func (c *Client) AddLiveDomain(request *AddLiveDomainRequest) (response *AddLiveDomainResponse, err error) {
    if request == nil {
        request = NewAddLiveDomainRequest()
    }
    response = NewAddLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewAddLiveWatermarkRequest() (request *AddLiveWatermarkRequest) {
    request = &AddLiveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "AddLiveWatermark")
    return
}

func NewAddLiveWatermarkResponse() (response *AddLiveWatermarkResponse) {
    response = &AddLiveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加水印，成功返回水印id后，需要调用[CreateLiveWatermarkRule](/document/product/267/32629)接口将水印id绑定到流使用。
func (c *Client) AddLiveWatermark(request *AddLiveWatermarkRequest) (response *AddLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewAddLiveWatermarkRequest()
    }
    response = NewAddLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewBindLiveDomainCertRequest() (request *BindLiveDomainCertRequest) {
    request = &BindLiveDomainCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "BindLiveDomainCert")
    return
}

func NewBindLiveDomainCertResponse() (response *BindLiveDomainCertResponse) {
    response = &BindLiveDomainCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 域名绑定证书
func (c *Client) BindLiveDomainCert(request *BindLiveDomainCertRequest) (response *BindLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewBindLiveDomainCertRequest()
    }
    response = NewBindLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveCallbackRuleRequest() (request *CreateLiveCallbackRuleRequest) {
    request = &CreateLiveCallbackRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveCallbackRule")
    return
}

func NewCreateLiveCallbackRuleResponse() (response *CreateLiveCallbackRuleResponse) {
    response = &CreateLiveCallbackRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建回调规则，需要先调用[CreateLiveCallbackTemplate](/document/product/267/32637)接口创建回调模板，将返回的模板id绑定到域名/路径进行使用。
// <br>回调协议相关文档：[事件消息通知](/document/product/267/32744)。
func (c *Client) CreateLiveCallbackRule(request *CreateLiveCallbackRuleRequest) (response *CreateLiveCallbackRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveCallbackRuleRequest()
    }
    response = NewCreateLiveCallbackRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveCallbackTemplateRequest() (request *CreateLiveCallbackTemplateRequest) {
    request = &CreateLiveCallbackTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveCallbackTemplate")
    return
}

func NewCreateLiveCallbackTemplateResponse() (response *CreateLiveCallbackTemplateResponse) {
    response = &CreateLiveCallbackTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建回调模板，成功返回模板id后，需要调用[CreateLiveCallbackRule](/document/product/267/32638)接口将模板id绑定到域名/路径使用。
// <br>回调协议相关文档：[事件消息通知](/document/product/267/32744)。
func (c *Client) CreateLiveCallbackTemplate(request *CreateLiveCallbackTemplateRequest) (response *CreateLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveCallbackTemplateRequest()
    }
    response = NewCreateLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveCertRequest() (request *CreateLiveCertRequest) {
    request = &CreateLiveCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveCert")
    return
}

func NewCreateLiveCertResponse() (response *CreateLiveCertResponse) {
    response = &CreateLiveCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加证书
func (c *Client) CreateLiveCert(request *CreateLiveCertRequest) (response *CreateLiveCertResponse, err error) {
    if request == nil {
        request = NewCreateLiveCertRequest()
    }
    response = NewCreateLiveCertResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveRecordRequest() (request *CreateLiveRecordRequest) {
    request = &CreateLiveRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveRecord")
    return
}

func NewCreateLiveRecordResponse() (response *CreateLiveRecordResponse) {
    response = &CreateLiveRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// - 使用前提
//   1. 录制文件存放于点播平台，所以用户如需使用录制功能，需首先自行开通点播服务。
//   2. 录制文件存放后相关费用（含存储以及下行播放流量）按照点播平台计费方式收取，具体请参考 [对应文档](https://cloud.tencent.com/document/product/266/2838)。
// 
// - 模式说明
//   该接口支持两种录制模式：
//   1. 定时录制模式【默认模式】。
//     需要传入开始时间与结束时间，录制任务根据时间自动开始与结束。
//   2. 实时视频录制模式。
//     忽略传入的开始时间，在录制任务创建后立即开始录制，录制时长支持最大为30分钟，如果传入的结束时间与当前时间差大于30分钟，则按30分钟计算，实时视频录制主要用于录制精彩视频场景，时长建议控制在5分钟以内。
// 
// - 注意事项
//   1. 调用接口超时设置应大于3秒，小于3秒重试以及频繁调用都有可能产生重复录制任务。
func (c *Client) CreateLiveRecord(request *CreateLiveRecordRequest) (response *CreateLiveRecordResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordRequest()
    }
    response = NewCreateLiveRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveRecordRuleRequest() (request *CreateLiveRecordRuleRequest) {
    request = &CreateLiveRecordRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveRecordRule")
    return
}

func NewCreateLiveRecordRuleResponse() (response *CreateLiveRecordRuleResponse) {
    response = &CreateLiveRecordRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建录制规则，需要先调用[CreateLiveRecordTemplate](/document/product/267/32614)接口创建录制模板，将返回的模板id绑定到流使用。
// <br>录制相关文档：[直播录制](/document/product/267/32739)。
func (c *Client) CreateLiveRecordRule(request *CreateLiveRecordRuleRequest) (response *CreateLiveRecordRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordRuleRequest()
    }
    response = NewCreateLiveRecordRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveRecordTemplateRequest() (request *CreateLiveRecordTemplateRequest) {
    request = &CreateLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveRecordTemplate")
    return
}

func NewCreateLiveRecordTemplateResponse() (response *CreateLiveRecordTemplateResponse) {
    response = &CreateLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建录制模板，成功返回模板id后，需要调用[CreateLiveRecordRule](/document/product/267/32615)接口，将模板id绑定到流进行使用。
// <br>录制相关文档：[直播录制](/document/product/267/32739)。
func (c *Client) CreateLiveRecordTemplate(request *CreateLiveRecordTemplateRequest) (response *CreateLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveRecordTemplateRequest()
    }
    response = NewCreateLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveSnapshotRuleRequest() (request *CreateLiveSnapshotRuleRequest) {
    request = &CreateLiveSnapshotRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveSnapshotRule")
    return
}

func NewCreateLiveSnapshotRuleResponse() (response *CreateLiveSnapshotRuleResponse) {
    response = &CreateLiveSnapshotRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建截图规则，需要先调用[CreateLiveSnapshotTemplate](/document/product/267/32624)接口创建截图模板，然后将返回的模板id绑定到流进行使用。
// <br>截图相关文档：[直播截图](/document/product/267/32737)。
func (c *Client) CreateLiveSnapshotRule(request *CreateLiveSnapshotRuleRequest) (response *CreateLiveSnapshotRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveSnapshotRuleRequest()
    }
    response = NewCreateLiveSnapshotRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveSnapshotTemplateRequest() (request *CreateLiveSnapshotTemplateRequest) {
    request = &CreateLiveSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveSnapshotTemplate")
    return
}

func NewCreateLiveSnapshotTemplateResponse() (response *CreateLiveSnapshotTemplateResponse) {
    response = &CreateLiveSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建截图模板，成功返回模板id后，需要调用[CreateLiveSnapshotRule](/document/product/267/32625)接口，将模板id绑定到流使用。
// <br>截图相关文档：[直播截图](/document/product/267/32737)。
func (c *Client) CreateLiveSnapshotTemplate(request *CreateLiveSnapshotTemplateRequest) (response *CreateLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveSnapshotTemplateRequest()
    }
    response = NewCreateLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveTranscodeRuleRequest() (request *CreateLiveTranscodeRuleRequest) {
    request = &CreateLiveTranscodeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveTranscodeRule")
    return
}

func NewCreateLiveTranscodeRuleResponse() (response *CreateLiveTranscodeRuleResponse) {
    response = &CreateLiveTranscodeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建转码规则，需要先调用[CreateLiveTranscodeTemplate](/document/product/267/32646)接口创建转码模板，将返回的模板id绑定到流使用。
// <br>转码相关文档：[直播转封装及转码](/document/product/267/32736)。
func (c *Client) CreateLiveTranscodeRule(request *CreateLiveTranscodeRuleRequest) (response *CreateLiveTranscodeRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveTranscodeRuleRequest()
    }
    response = NewCreateLiveTranscodeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveTranscodeTemplateRequest() (request *CreateLiveTranscodeTemplateRequest) {
    request = &CreateLiveTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveTranscodeTemplate")
    return
}

func NewCreateLiveTranscodeTemplateResponse() (response *CreateLiveTranscodeTemplateResponse) {
    response = &CreateLiveTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建转码模板，成功返回模板id后，需要调用[CreateLiveTranscodeRule](/document/product/267/32647)接口，将返回的模板id绑定到流使用。
// <br>转码相关文档：[直播转封装及转码](/document/product/267/32736)。
func (c *Client) CreateLiveTranscodeTemplate(request *CreateLiveTranscodeTemplateRequest) (response *CreateLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLiveTranscodeTemplateRequest()
    }
    response = NewCreateLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLiveWatermarkRuleRequest() (request *CreateLiveWatermarkRuleRequest) {
    request = &CreateLiveWatermarkRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLiveWatermarkRule")
    return
}

func NewCreateLiveWatermarkRuleResponse() (response *CreateLiveWatermarkRuleResponse) {
    response = &CreateLiveWatermarkRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建水印规则，需要先调用[AddLiveWatermark](/document/product/267/30154)接口添加水印，将返回的水印id绑定到流使用。
func (c *Client) CreateLiveWatermarkRule(request *CreateLiveWatermarkRuleRequest) (response *CreateLiveWatermarkRuleResponse, err error) {
    if request == nil {
        request = NewCreateLiveWatermarkRuleRequest()
    }
    response = NewCreateLiveWatermarkRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePullStreamConfigRequest() (request *CreatePullStreamConfigRequest) {
    request = &CreatePullStreamConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreatePullStreamConfig")
    return
}

func NewCreatePullStreamConfigResponse() (response *CreatePullStreamConfigResponse) {
    response = &CreatePullStreamConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加拉流配置，目前限制添加10条任务。
func (c *Client) CreatePullStreamConfig(request *CreatePullStreamConfigRequest) (response *CreatePullStreamConfigResponse, err error) {
    if request == nil {
        request = NewCreatePullStreamConfigRequest()
    }
    response = NewCreatePullStreamConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveCallbackRuleRequest() (request *DeleteLiveCallbackRuleRequest) {
    request = &DeleteLiveCallbackRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveCallbackRule")
    return
}

func NewDeleteLiveCallbackRuleResponse() (response *DeleteLiveCallbackRuleResponse) {
    response = &DeleteLiveCallbackRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除回调规则
func (c *Client) DeleteLiveCallbackRule(request *DeleteLiveCallbackRuleRequest) (response *DeleteLiveCallbackRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveCallbackRuleRequest()
    }
    response = NewDeleteLiveCallbackRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveCallbackTemplateRequest() (request *DeleteLiveCallbackTemplateRequest) {
    request = &DeleteLiveCallbackTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveCallbackTemplate")
    return
}

func NewDeleteLiveCallbackTemplateResponse() (response *DeleteLiveCallbackTemplateResponse) {
    response = &DeleteLiveCallbackTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除回调模板
func (c *Client) DeleteLiveCallbackTemplate(request *DeleteLiveCallbackTemplateRequest) (response *DeleteLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveCallbackTemplateRequest()
    }
    response = NewDeleteLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveCertRequest() (request *DeleteLiveCertRequest) {
    request = &DeleteLiveCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveCert")
    return
}

func NewDeleteLiveCertResponse() (response *DeleteLiveCertResponse) {
    response = &DeleteLiveCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除域名对应的证书
func (c *Client) DeleteLiveCert(request *DeleteLiveCertRequest) (response *DeleteLiveCertResponse, err error) {
    if request == nil {
        request = NewDeleteLiveCertRequest()
    }
    response = NewDeleteLiveCertResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveDomainRequest() (request *DeleteLiveDomainRequest) {
    request = &DeleteLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveDomain")
    return
}

func NewDeleteLiveDomainResponse() (response *DeleteLiveDomainResponse) {
    response = &DeleteLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除已添加的直播域名
func (c *Client) DeleteLiveDomain(request *DeleteLiveDomainRequest) (response *DeleteLiveDomainResponse, err error) {
    if request == nil {
        request = NewDeleteLiveDomainRequest()
    }
    response = NewDeleteLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveRecordRequest() (request *DeleteLiveRecordRequest) {
    request = &DeleteLiveRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveRecord")
    return
}

func NewDeleteLiveRecordResponse() (response *DeleteLiveRecordResponse) {
    response = &DeleteLiveRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于删除录制任务
func (c *Client) DeleteLiveRecord(request *DeleteLiveRecordRequest) (response *DeleteLiveRecordResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordRequest()
    }
    response = NewDeleteLiveRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveRecordRuleRequest() (request *DeleteLiveRecordRuleRequest) {
    request = &DeleteLiveRecordRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveRecordRule")
    return
}

func NewDeleteLiveRecordRuleResponse() (response *DeleteLiveRecordRuleResponse) {
    response = &DeleteLiveRecordRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除录制规则
func (c *Client) DeleteLiveRecordRule(request *DeleteLiveRecordRuleRequest) (response *DeleteLiveRecordRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordRuleRequest()
    }
    response = NewDeleteLiveRecordRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveRecordTemplateRequest() (request *DeleteLiveRecordTemplateRequest) {
    request = &DeleteLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveRecordTemplate")
    return
}

func NewDeleteLiveRecordTemplateResponse() (response *DeleteLiveRecordTemplateResponse) {
    response = &DeleteLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除录制模板
func (c *Client) DeleteLiveRecordTemplate(request *DeleteLiveRecordTemplateRequest) (response *DeleteLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveRecordTemplateRequest()
    }
    response = NewDeleteLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveSnapshotRuleRequest() (request *DeleteLiveSnapshotRuleRequest) {
    request = &DeleteLiveSnapshotRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveSnapshotRule")
    return
}

func NewDeleteLiveSnapshotRuleResponse() (response *DeleteLiveSnapshotRuleResponse) {
    response = &DeleteLiveSnapshotRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除截图规则
func (c *Client) DeleteLiveSnapshotRule(request *DeleteLiveSnapshotRuleRequest) (response *DeleteLiveSnapshotRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveSnapshotRuleRequest()
    }
    response = NewDeleteLiveSnapshotRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveSnapshotTemplateRequest() (request *DeleteLiveSnapshotTemplateRequest) {
    request = &DeleteLiveSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveSnapshotTemplate")
    return
}

func NewDeleteLiveSnapshotTemplateResponse() (response *DeleteLiveSnapshotTemplateResponse) {
    response = &DeleteLiveSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除截图模板
func (c *Client) DeleteLiveSnapshotTemplate(request *DeleteLiveSnapshotTemplateRequest) (response *DeleteLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveSnapshotTemplateRequest()
    }
    response = NewDeleteLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveTranscodeRuleRequest() (request *DeleteLiveTranscodeRuleRequest) {
    request = &DeleteLiveTranscodeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveTranscodeRule")
    return
}

func NewDeleteLiveTranscodeRuleResponse() (response *DeleteLiveTranscodeRuleResponse) {
    response = &DeleteLiveTranscodeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除转码规则
func (c *Client) DeleteLiveTranscodeRule(request *DeleteLiveTranscodeRuleRequest) (response *DeleteLiveTranscodeRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveTranscodeRuleRequest()
    }
    response = NewDeleteLiveTranscodeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveTranscodeTemplateRequest() (request *DeleteLiveTranscodeTemplateRequest) {
    request = &DeleteLiveTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveTranscodeTemplate")
    return
}

func NewDeleteLiveTranscodeTemplateResponse() (response *DeleteLiveTranscodeTemplateResponse) {
    response = &DeleteLiveTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除转码模板
func (c *Client) DeleteLiveTranscodeTemplate(request *DeleteLiveTranscodeTemplateRequest) (response *DeleteLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLiveTranscodeTemplateRequest()
    }
    response = NewDeleteLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveWatermarkRequest() (request *DeleteLiveWatermarkRequest) {
    request = &DeleteLiveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveWatermark")
    return
}

func NewDeleteLiveWatermarkResponse() (response *DeleteLiveWatermarkResponse) {
    response = &DeleteLiveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除水印
func (c *Client) DeleteLiveWatermark(request *DeleteLiveWatermarkRequest) (response *DeleteLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewDeleteLiveWatermarkRequest()
    }
    response = NewDeleteLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLiveWatermarkRuleRequest() (request *DeleteLiveWatermarkRuleRequest) {
    request = &DeleteLiveWatermarkRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLiveWatermarkRule")
    return
}

func NewDeleteLiveWatermarkRuleResponse() (response *DeleteLiveWatermarkRuleResponse) {
    response = &DeleteLiveWatermarkRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除水印规则
func (c *Client) DeleteLiveWatermarkRule(request *DeleteLiveWatermarkRuleRequest) (response *DeleteLiveWatermarkRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLiveWatermarkRuleRequest()
    }
    response = NewDeleteLiveWatermarkRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePullStreamConfigRequest() (request *DeletePullStreamConfigRequest) {
    request = &DeletePullStreamConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeletePullStreamConfig")
    return
}

func NewDeletePullStreamConfigResponse() (response *DeletePullStreamConfigResponse) {
    response = &DeletePullStreamConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除直播拉流配置
func (c *Client) DeletePullStreamConfig(request *DeletePullStreamConfigRequest) (response *DeletePullStreamConfigResponse, err error) {
    if request == nil {
        request = NewDeletePullStreamConfigRequest()
    }
    response = NewDeletePullStreamConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCallbackRulesRequest() (request *DescribeLiveCallbackRulesRequest) {
    request = &DescribeLiveCallbackRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCallbackRules")
    return
}

func NewDescribeLiveCallbackRulesResponse() (response *DescribeLiveCallbackRulesResponse) {
    response = &DescribeLiveCallbackRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取回调规则列表
func (c *Client) DescribeLiveCallbackRules(request *DescribeLiveCallbackRulesRequest) (response *DescribeLiveCallbackRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCallbackRulesRequest()
    }
    response = NewDescribeLiveCallbackRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCallbackTemplateRequest() (request *DescribeLiveCallbackTemplateRequest) {
    request = &DescribeLiveCallbackTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCallbackTemplate")
    return
}

func NewDescribeLiveCallbackTemplateResponse() (response *DescribeLiveCallbackTemplateResponse) {
    response = &DescribeLiveCallbackTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单个回调模板
func (c *Client) DescribeLiveCallbackTemplate(request *DescribeLiveCallbackTemplateRequest) (response *DescribeLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCallbackTemplateRequest()
    }
    response = NewDescribeLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCallbackTemplatesRequest() (request *DescribeLiveCallbackTemplatesRequest) {
    request = &DescribeLiveCallbackTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCallbackTemplates")
    return
}

func NewDescribeLiveCallbackTemplatesResponse() (response *DescribeLiveCallbackTemplatesResponse) {
    response = &DescribeLiveCallbackTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取回调模板列表
func (c *Client) DescribeLiveCallbackTemplates(request *DescribeLiveCallbackTemplatesRequest) (response *DescribeLiveCallbackTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCallbackTemplatesRequest()
    }
    response = NewDescribeLiveCallbackTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCertRequest() (request *DescribeLiveCertRequest) {
    request = &DescribeLiveCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCert")
    return
}

func NewDescribeLiveCertResponse() (response *DescribeLiveCertResponse) {
    response = &DescribeLiveCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取证书信息
func (c *Client) DescribeLiveCert(request *DescribeLiveCertRequest) (response *DescribeLiveCertResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCertRequest()
    }
    response = NewDescribeLiveCertResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveCertsRequest() (request *DescribeLiveCertsRequest) {
    request = &DescribeLiveCertsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveCerts")
    return
}

func NewDescribeLiveCertsResponse() (response *DescribeLiveCertsResponse) {
    response = &DescribeLiveCertsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取证书信息列表
func (c *Client) DescribeLiveCerts(request *DescribeLiveCertsRequest) (response *DescribeLiveCertsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCertsRequest()
    }
    response = NewDescribeLiveCertsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainRequest() (request *DescribeLiveDomainRequest) {
    request = &DescribeLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomain")
    return
}

func NewDescribeLiveDomainResponse() (response *DescribeLiveDomainResponse) {
    response = &DescribeLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询直播域名信息。
func (c *Client) DescribeLiveDomain(request *DescribeLiveDomainRequest) (response *DescribeLiveDomainResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainRequest()
    }
    response = NewDescribeLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainCertRequest() (request *DescribeLiveDomainCertRequest) {
    request = &DescribeLiveDomainCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomainCert")
    return
}

func NewDescribeLiveDomainCertResponse() (response *DescribeLiveDomainCertResponse) {
    response = &DescribeLiveDomainCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取域名证书信息
func (c *Client) DescribeLiveDomainCert(request *DescribeLiveDomainCertRequest) (response *DescribeLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainCertRequest()
    }
    response = NewDescribeLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainsRequest() (request *DescribeLiveDomainsRequest) {
    request = &DescribeLiveDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomains")
    return
}

func NewDescribeLiveDomainsResponse() (response *DescribeLiveDomainsResponse) {
    response = &DescribeLiveDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据域名状态、类型等信息查询用户的域名信息
func (c *Client) DescribeLiveDomains(request *DescribeLiveDomainsRequest) (response *DescribeLiveDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainsRequest()
    }
    response = NewDescribeLiveDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveForbidStreamListRequest() (request *DescribeLiveForbidStreamListRequest) {
    request = &DescribeLiveForbidStreamListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveForbidStreamList")
    return
}

func NewDescribeLiveForbidStreamListResponse() (response *DescribeLiveForbidStreamListResponse) {
    response = &DescribeLiveForbidStreamListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取禁推流列表
func (c *Client) DescribeLiveForbidStreamList(request *DescribeLiveForbidStreamListRequest) (response *DescribeLiveForbidStreamListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveForbidStreamListRequest()
    }
    response = NewDescribeLiveForbidStreamListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePlayAuthKeyRequest() (request *DescribeLivePlayAuthKeyRequest) {
    request = &DescribeLivePlayAuthKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePlayAuthKey")
    return
}

func NewDescribeLivePlayAuthKeyResponse() (response *DescribeLivePlayAuthKeyResponse) {
    response = &DescribeLivePlayAuthKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询播放鉴权key
func (c *Client) DescribeLivePlayAuthKey(request *DescribeLivePlayAuthKeyRequest) (response *DescribeLivePlayAuthKeyResponse, err error) {
    if request == nil {
        request = NewDescribeLivePlayAuthKeyRequest()
    }
    response = NewDescribeLivePlayAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePushAuthKeyRequest() (request *DescribeLivePushAuthKeyRequest) {
    request = &DescribeLivePushAuthKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePushAuthKey")
    return
}

func NewDescribeLivePushAuthKeyResponse() (response *DescribeLivePushAuthKeyResponse) {
    response = &DescribeLivePushAuthKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询直播推流鉴权key
func (c *Client) DescribeLivePushAuthKey(request *DescribeLivePushAuthKeyRequest) (response *DescribeLivePushAuthKeyResponse, err error) {
    if request == nil {
        request = NewDescribeLivePushAuthKeyRequest()
    }
    response = NewDescribeLivePushAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordRulesRequest() (request *DescribeLiveRecordRulesRequest) {
    request = &DescribeLiveRecordRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveRecordRules")
    return
}

func NewDescribeLiveRecordRulesResponse() (response *DescribeLiveRecordRulesResponse) {
    response = &DescribeLiveRecordRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取录制规则列表
func (c *Client) DescribeLiveRecordRules(request *DescribeLiveRecordRulesRequest) (response *DescribeLiveRecordRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordRulesRequest()
    }
    response = NewDescribeLiveRecordRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordTemplateRequest() (request *DescribeLiveRecordTemplateRequest) {
    request = &DescribeLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveRecordTemplate")
    return
}

func NewDescribeLiveRecordTemplateResponse() (response *DescribeLiveRecordTemplateResponse) {
    response = &DescribeLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单个录制模板
func (c *Client) DescribeLiveRecordTemplate(request *DescribeLiveRecordTemplateRequest) (response *DescribeLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordTemplateRequest()
    }
    response = NewDescribeLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveRecordTemplatesRequest() (request *DescribeLiveRecordTemplatesRequest) {
    request = &DescribeLiveRecordTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveRecordTemplates")
    return
}

func NewDescribeLiveRecordTemplatesResponse() (response *DescribeLiveRecordTemplatesResponse) {
    response = &DescribeLiveRecordTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取录制模板列表
func (c *Client) DescribeLiveRecordTemplates(request *DescribeLiveRecordTemplatesRequest) (response *DescribeLiveRecordTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveRecordTemplatesRequest()
    }
    response = NewDescribeLiveRecordTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveSnapshotRulesRequest() (request *DescribeLiveSnapshotRulesRequest) {
    request = &DescribeLiveSnapshotRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveSnapshotRules")
    return
}

func NewDescribeLiveSnapshotRulesResponse() (response *DescribeLiveSnapshotRulesResponse) {
    response = &DescribeLiveSnapshotRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取截图规则列表
func (c *Client) DescribeLiveSnapshotRules(request *DescribeLiveSnapshotRulesRequest) (response *DescribeLiveSnapshotRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveSnapshotRulesRequest()
    }
    response = NewDescribeLiveSnapshotRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveSnapshotTemplateRequest() (request *DescribeLiveSnapshotTemplateRequest) {
    request = &DescribeLiveSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveSnapshotTemplate")
    return
}

func NewDescribeLiveSnapshotTemplateResponse() (response *DescribeLiveSnapshotTemplateResponse) {
    response = &DescribeLiveSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单个截图模板
func (c *Client) DescribeLiveSnapshotTemplate(request *DescribeLiveSnapshotTemplateRequest) (response *DescribeLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveSnapshotTemplateRequest()
    }
    response = NewDescribeLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveSnapshotTemplatesRequest() (request *DescribeLiveSnapshotTemplatesRequest) {
    request = &DescribeLiveSnapshotTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveSnapshotTemplates")
    return
}

func NewDescribeLiveSnapshotTemplatesResponse() (response *DescribeLiveSnapshotTemplatesResponse) {
    response = &DescribeLiveSnapshotTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取截图模板列表
func (c *Client) DescribeLiveSnapshotTemplates(request *DescribeLiveSnapshotTemplatesRequest) (response *DescribeLiveSnapshotTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveSnapshotTemplatesRequest()
    }
    response = NewDescribeLiveSnapshotTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamEventListRequest() (request *DescribeLiveStreamEventListRequest) {
    request = &DescribeLiveStreamEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamEventList")
    return
}

func NewDescribeLiveStreamEventListResponse() (response *DescribeLiveStreamEventListResponse) {
    response = &DescribeLiveStreamEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询推断流事件
func (c *Client) DescribeLiveStreamEventList(request *DescribeLiveStreamEventListRequest) (response *DescribeLiveStreamEventListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamEventListRequest()
    }
    response = NewDescribeLiveStreamEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamOnlineInfoRequest() (request *DescribeLiveStreamOnlineInfoRequest) {
    request = &DescribeLiveStreamOnlineInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamOnlineInfo")
    return
}

func NewDescribeLiveStreamOnlineInfoResponse() (response *DescribeLiveStreamOnlineInfoResponse) {
    response = &DescribeLiveStreamOnlineInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询在线推流信息列表
func (c *Client) DescribeLiveStreamOnlineInfo(request *DescribeLiveStreamOnlineInfoRequest) (response *DescribeLiveStreamOnlineInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamOnlineInfoRequest()
    }
    response = NewDescribeLiveStreamOnlineInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamOnlineListRequest() (request *DescribeLiveStreamOnlineListRequest) {
    request = &DescribeLiveStreamOnlineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamOnlineList")
    return
}

func NewDescribeLiveStreamOnlineListResponse() (response *DescribeLiveStreamOnlineListResponse) {
    response = &DescribeLiveStreamOnlineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回正在直播中的流列表
func (c *Client) DescribeLiveStreamOnlineList(request *DescribeLiveStreamOnlineListRequest) (response *DescribeLiveStreamOnlineListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamOnlineListRequest()
    }
    response = NewDescribeLiveStreamOnlineListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamPublishedListRequest() (request *DescribeLiveStreamPublishedListRequest) {
    request = &DescribeLiveStreamPublishedListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamPublishedList")
    return
}

func NewDescribeLiveStreamPublishedListResponse() (response *DescribeLiveStreamPublishedListResponse) {
    response = &DescribeLiveStreamPublishedListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回已经推过流的流列表
func (c *Client) DescribeLiveStreamPublishedList(request *DescribeLiveStreamPublishedListRequest) (response *DescribeLiveStreamPublishedListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamPublishedListRequest()
    }
    response = NewDescribeLiveStreamPublishedListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamPushInfoListRequest() (request *DescribeLiveStreamPushInfoListRequest) {
    request = &DescribeLiveStreamPushInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamPushInfoList")
    return
}

func NewDescribeLiveStreamPushInfoListResponse() (response *DescribeLiveStreamPushInfoListResponse) {
    response = &DescribeLiveStreamPushInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询所有实时流的推流信息，包括客户端IP，服务端IP，帧率，码率，域名，开始推流时间。
func (c *Client) DescribeLiveStreamPushInfoList(request *DescribeLiveStreamPushInfoListRequest) (response *DescribeLiveStreamPushInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamPushInfoListRequest()
    }
    response = NewDescribeLiveStreamPushInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveStreamStateRequest() (request *DescribeLiveStreamStateRequest) {
    request = &DescribeLiveStreamStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveStreamState")
    return
}

func NewDescribeLiveStreamStateResponse() (response *DescribeLiveStreamStateResponse) {
    response = &DescribeLiveStreamStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回直播中、无推流或者禁播等状态
func (c *Client) DescribeLiveStreamState(request *DescribeLiveStreamStateRequest) (response *DescribeLiveStreamStateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamStateRequest()
    }
    response = NewDescribeLiveStreamStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeDetailInfoRequest() (request *DescribeLiveTranscodeDetailInfoRequest) {
    request = &DescribeLiveTranscodeDetailInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeDetailInfo")
    return
}

func NewDescribeLiveTranscodeDetailInfoResponse() (response *DescribeLiveTranscodeDetailInfoResponse) {
    response = &DescribeLiveTranscodeDetailInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 支持查询某天的转码详细信息。
// 注意：当前只支持查询近30天内某天的详细数据。
func (c *Client) DescribeLiveTranscodeDetailInfo(request *DescribeLiveTranscodeDetailInfoRequest) (response *DescribeLiveTranscodeDetailInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeDetailInfoRequest()
    }
    response = NewDescribeLiveTranscodeDetailInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeRulesRequest() (request *DescribeLiveTranscodeRulesRequest) {
    request = &DescribeLiveTranscodeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeRules")
    return
}

func NewDescribeLiveTranscodeRulesResponse() (response *DescribeLiveTranscodeRulesResponse) {
    response = &DescribeLiveTranscodeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取转码规则列表
func (c *Client) DescribeLiveTranscodeRules(request *DescribeLiveTranscodeRulesRequest) (response *DescribeLiveTranscodeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeRulesRequest()
    }
    response = NewDescribeLiveTranscodeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeTemplateRequest() (request *DescribeLiveTranscodeTemplateRequest) {
    request = &DescribeLiveTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeTemplate")
    return
}

func NewDescribeLiveTranscodeTemplateResponse() (response *DescribeLiveTranscodeTemplateResponse) {
    response = &DescribeLiveTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单个转码模板
func (c *Client) DescribeLiveTranscodeTemplate(request *DescribeLiveTranscodeTemplateRequest) (response *DescribeLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeTemplateRequest()
    }
    response = NewDescribeLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveTranscodeTemplatesRequest() (request *DescribeLiveTranscodeTemplatesRequest) {
    request = &DescribeLiveTranscodeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveTranscodeTemplates")
    return
}

func NewDescribeLiveTranscodeTemplatesResponse() (response *DescribeLiveTranscodeTemplatesResponse) {
    response = &DescribeLiveTranscodeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取转码模板列表
func (c *Client) DescribeLiveTranscodeTemplates(request *DescribeLiveTranscodeTemplatesRequest) (response *DescribeLiveTranscodeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveTranscodeTemplatesRequest()
    }
    response = NewDescribeLiveTranscodeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveWatermarkRequest() (request *DescribeLiveWatermarkRequest) {
    request = &DescribeLiveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveWatermark")
    return
}

func NewDescribeLiveWatermarkResponse() (response *DescribeLiveWatermarkResponse) {
    response = &DescribeLiveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单个水印信息
func (c *Client) DescribeLiveWatermark(request *DescribeLiveWatermarkRequest) (response *DescribeLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewDescribeLiveWatermarkRequest()
    }
    response = NewDescribeLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveWatermarkRulesRequest() (request *DescribeLiveWatermarkRulesRequest) {
    request = &DescribeLiveWatermarkRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveWatermarkRules")
    return
}

func NewDescribeLiveWatermarkRulesResponse() (response *DescribeLiveWatermarkRulesResponse) {
    response = &DescribeLiveWatermarkRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取水印规则列表
func (c *Client) DescribeLiveWatermarkRules(request *DescribeLiveWatermarkRulesRequest) (response *DescribeLiveWatermarkRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLiveWatermarkRulesRequest()
    }
    response = NewDescribeLiveWatermarkRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveWatermarksRequest() (request *DescribeLiveWatermarksRequest) {
    request = &DescribeLiveWatermarksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveWatermarks")
    return
}

func NewDescribeLiveWatermarksResponse() (response *DescribeLiveWatermarksResponse) {
    response = &DescribeLiveWatermarksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询水印列表
func (c *Client) DescribeLiveWatermarks(request *DescribeLiveWatermarksRequest) (response *DescribeLiveWatermarksResponse, err error) {
    if request == nil {
        request = NewDescribeLiveWatermarksRequest()
    }
    response = NewDescribeLiveWatermarksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogDownloadListRequest() (request *DescribeLogDownloadListRequest) {
    request = &DescribeLogDownloadListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLogDownloadList")
    return
}

func NewDescribeLogDownloadListResponse() (response *DescribeLogDownloadListResponse) {
    response = &DescribeLogDownloadListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量获取日志URL。
func (c *Client) DescribeLogDownloadList(request *DescribeLogDownloadListRequest) (response *DescribeLogDownloadListResponse, err error) {
    if request == nil {
        request = NewDescribeLogDownloadListRequest()
    }
    response = NewDescribeLogDownloadListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProIspPlaySumInfoListRequest() (request *DescribeProIspPlaySumInfoListRequest) {
    request = &DescribeProIspPlaySumInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeProIspPlaySumInfoList")
    return
}

func NewDescribeProIspPlaySumInfoListResponse() (response *DescribeProIspPlaySumInfoListResponse) {
    response = &DescribeProIspPlaySumInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询某段时间内每个省份每个运营商的平均每秒流量，总流量，总请求数信息。
func (c *Client) DescribeProIspPlaySumInfoList(request *DescribeProIspPlaySumInfoListRequest) (response *DescribeProIspPlaySumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeProIspPlaySumInfoListRequest()
    }
    response = NewDescribeProIspPlaySumInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProvinceIspPlayInfoListRequest() (request *DescribeProvinceIspPlayInfoListRequest) {
    request = &DescribeProvinceIspPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeProvinceIspPlayInfoList")
    return
}

func NewDescribeProvinceIspPlayInfoListResponse() (response *DescribeProvinceIspPlayInfoListResponse) {
    response = &DescribeProvinceIspPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询某省份某运营商下行播放数据，包括带宽，流量，请求数，并发连接数信息。
func (c *Client) DescribeProvinceIspPlayInfoList(request *DescribeProvinceIspPlayInfoListRequest) (response *DescribeProvinceIspPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeProvinceIspPlayInfoListRequest()
    }
    response = NewDescribeProvinceIspPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePullStreamConfigsRequest() (request *DescribePullStreamConfigsRequest) {
    request = &DescribePullStreamConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribePullStreamConfigs")
    return
}

func NewDescribePullStreamConfigsResponse() (response *DescribePullStreamConfigsResponse) {
    response = &DescribePullStreamConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询拉流配置
func (c *Client) DescribePullStreamConfigs(request *DescribePullStreamConfigsRequest) (response *DescribePullStreamConfigsResponse, err error) {
    if request == nil {
        request = NewDescribePullStreamConfigsRequest()
    }
    response = NewDescribePullStreamConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamDayPlayInfoListRequest() (request *DescribeStreamDayPlayInfoListRequest) {
    request = &DescribeStreamDayPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeStreamDayPlayInfoList")
    return
}

func NewDescribeStreamDayPlayInfoListResponse() (response *DescribeStreamDayPlayInfoListResponse) {
    response = &DescribeStreamDayPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询天维度每条流的播放数据，包括总流量等。
func (c *Client) DescribeStreamDayPlayInfoList(request *DescribeStreamDayPlayInfoListRequest) (response *DescribeStreamDayPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamDayPlayInfoListRequest()
    }
    response = NewDescribeStreamDayPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPlayInfoListRequest() (request *DescribeStreamPlayInfoListRequest) {
    request = &DescribeStreamPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeStreamPlayInfoList")
    return
}

func NewDescribeStreamPlayInfoListResponse() (response *DescribeStreamPlayInfoListResponse) {
    response = &DescribeStreamPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询播放数据，支持按流名称查询详细播放数据，也可按播放域名查询详细总数据。
func (c *Client) DescribeStreamPlayInfoList(request *DescribeStreamPlayInfoListRequest) (response *DescribeStreamPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPlayInfoListRequest()
    }
    response = NewDescribeStreamPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDropLiveStreamRequest() (request *DropLiveStreamRequest) {
    request = &DropLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DropLiveStream")
    return
}

func NewDropLiveStreamResponse() (response *DropLiveStreamResponse) {
    response = &DropLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 断开推流连接，但可以重新推流
func (c *Client) DropLiveStream(request *DropLiveStreamRequest) (response *DropLiveStreamResponse, err error) {
    if request == nil {
        request = NewDropLiveStreamRequest()
    }
    response = NewDropLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewEnableLiveDomainRequest() (request *EnableLiveDomainRequest) {
    request = &EnableLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "EnableLiveDomain")
    return
}

func NewEnableLiveDomainResponse() (response *EnableLiveDomainResponse) {
    response = &EnableLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 启用状态为停用的直播域名
func (c *Client) EnableLiveDomain(request *EnableLiveDomainRequest) (response *EnableLiveDomainResponse, err error) {
    if request == nil {
        request = NewEnableLiveDomainRequest()
    }
    response = NewEnableLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewForbidLiveDomainRequest() (request *ForbidLiveDomainRequest) {
    request = &ForbidLiveDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ForbidLiveDomain")
    return
}

func NewForbidLiveDomainResponse() (response *ForbidLiveDomainResponse) {
    response = &ForbidLiveDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停用使用某个直播域名
func (c *Client) ForbidLiveDomain(request *ForbidLiveDomainRequest) (response *ForbidLiveDomainResponse, err error) {
    if request == nil {
        request = NewForbidLiveDomainRequest()
    }
    response = NewForbidLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewForbidLiveStreamRequest() (request *ForbidLiveStreamRequest) {
    request = &ForbidLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ForbidLiveStream")
    return
}

func NewForbidLiveStreamResponse() (response *ForbidLiveStreamResponse) {
    response = &ForbidLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 禁止某条流的推送，可以预设某个时刻将流恢复。
func (c *Client) ForbidLiveStream(request *ForbidLiveStreamRequest) (response *ForbidLiveStreamResponse, err error) {
    if request == nil {
        request = NewForbidLiveStreamRequest()
    }
    response = NewForbidLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveCallbackTemplateRequest() (request *ModifyLiveCallbackTemplateRequest) {
    request = &ModifyLiveCallbackTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveCallbackTemplate")
    return
}

func NewModifyLiveCallbackTemplateResponse() (response *ModifyLiveCallbackTemplateResponse) {
    response = &ModifyLiveCallbackTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改回调模板
func (c *Client) ModifyLiveCallbackTemplate(request *ModifyLiveCallbackTemplateRequest) (response *ModifyLiveCallbackTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveCallbackTemplateRequest()
    }
    response = NewModifyLiveCallbackTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveCertRequest() (request *ModifyLiveCertRequest) {
    request = &ModifyLiveCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveCert")
    return
}

func NewModifyLiveCertResponse() (response *ModifyLiveCertResponse) {
    response = &ModifyLiveCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改证书
func (c *Client) ModifyLiveCert(request *ModifyLiveCertRequest) (response *ModifyLiveCertResponse, err error) {
    if request == nil {
        request = NewModifyLiveCertRequest()
    }
    response = NewModifyLiveCertResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveDomainCertRequest() (request *ModifyLiveDomainCertRequest) {
    request = &ModifyLiveDomainCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveDomainCert")
    return
}

func NewModifyLiveDomainCertResponse() (response *ModifyLiveDomainCertResponse) {
    response = &ModifyLiveDomainCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改域名和证书绑定信息
func (c *Client) ModifyLiveDomainCert(request *ModifyLiveDomainCertRequest) (response *ModifyLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewModifyLiveDomainCertRequest()
    }
    response = NewModifyLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePlayAuthKeyRequest() (request *ModifyLivePlayAuthKeyRequest) {
    request = &ModifyLivePlayAuthKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePlayAuthKey")
    return
}

func NewModifyLivePlayAuthKeyResponse() (response *ModifyLivePlayAuthKeyResponse) {
    response = &ModifyLivePlayAuthKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改播放鉴权key
func (c *Client) ModifyLivePlayAuthKey(request *ModifyLivePlayAuthKeyRequest) (response *ModifyLivePlayAuthKeyResponse, err error) {
    if request == nil {
        request = NewModifyLivePlayAuthKeyRequest()
    }
    response = NewModifyLivePlayAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePlayDomainRequest() (request *ModifyLivePlayDomainRequest) {
    request = &ModifyLivePlayDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePlayDomain")
    return
}

func NewModifyLivePlayDomainResponse() (response *ModifyLivePlayDomainResponse) {
    response = &ModifyLivePlayDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改播放域名信息
func (c *Client) ModifyLivePlayDomain(request *ModifyLivePlayDomainRequest) (response *ModifyLivePlayDomainResponse, err error) {
    if request == nil {
        request = NewModifyLivePlayDomainRequest()
    }
    response = NewModifyLivePlayDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePushAuthKeyRequest() (request *ModifyLivePushAuthKeyRequest) {
    request = &ModifyLivePushAuthKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePushAuthKey")
    return
}

func NewModifyLivePushAuthKeyResponse() (response *ModifyLivePushAuthKeyResponse) {
    response = &ModifyLivePushAuthKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改直播推流鉴权key
func (c *Client) ModifyLivePushAuthKey(request *ModifyLivePushAuthKeyRequest) (response *ModifyLivePushAuthKeyResponse, err error) {
    if request == nil {
        request = NewModifyLivePushAuthKeyRequest()
    }
    response = NewModifyLivePushAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveRecordTemplateRequest() (request *ModifyLiveRecordTemplateRequest) {
    request = &ModifyLiveRecordTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveRecordTemplate")
    return
}

func NewModifyLiveRecordTemplateResponse() (response *ModifyLiveRecordTemplateResponse) {
    response = &ModifyLiveRecordTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改录制模板配置
func (c *Client) ModifyLiveRecordTemplate(request *ModifyLiveRecordTemplateRequest) (response *ModifyLiveRecordTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveRecordTemplateRequest()
    }
    response = NewModifyLiveRecordTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveSnapshotTemplateRequest() (request *ModifyLiveSnapshotTemplateRequest) {
    request = &ModifyLiveSnapshotTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveSnapshotTemplate")
    return
}

func NewModifyLiveSnapshotTemplateResponse() (response *ModifyLiveSnapshotTemplateResponse) {
    response = &ModifyLiveSnapshotTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改截图模板配置
func (c *Client) ModifyLiveSnapshotTemplate(request *ModifyLiveSnapshotTemplateRequest) (response *ModifyLiveSnapshotTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveSnapshotTemplateRequest()
    }
    response = NewModifyLiveSnapshotTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveTranscodeTemplateRequest() (request *ModifyLiveTranscodeTemplateRequest) {
    request = &ModifyLiveTranscodeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveTranscodeTemplate")
    return
}

func NewModifyLiveTranscodeTemplateResponse() (response *ModifyLiveTranscodeTemplateResponse) {
    response = &ModifyLiveTranscodeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改转码模板配置
func (c *Client) ModifyLiveTranscodeTemplate(request *ModifyLiveTranscodeTemplateRequest) (response *ModifyLiveTranscodeTemplateResponse, err error) {
    if request == nil {
        request = NewModifyLiveTranscodeTemplateRequest()
    }
    response = NewModifyLiveTranscodeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPullStreamConfigRequest() (request *ModifyPullStreamConfigRequest) {
    request = &ModifyPullStreamConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyPullStreamConfig")
    return
}

func NewModifyPullStreamConfigResponse() (response *ModifyPullStreamConfigResponse) {
    response = &ModifyPullStreamConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新拉流配置
func (c *Client) ModifyPullStreamConfig(request *ModifyPullStreamConfigRequest) (response *ModifyPullStreamConfigResponse, err error) {
    if request == nil {
        request = NewModifyPullStreamConfigRequest()
    }
    response = NewModifyPullStreamConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPullStreamStatusRequest() (request *ModifyPullStreamStatusRequest) {
    request = &ModifyPullStreamStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyPullStreamStatus")
    return
}

func NewModifyPullStreamStatusResponse() (response *ModifyPullStreamStatusResponse) {
    response = &ModifyPullStreamStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改直播拉流配置状态
func (c *Client) ModifyPullStreamStatus(request *ModifyPullStreamStatusRequest) (response *ModifyPullStreamStatusResponse, err error) {
    if request == nil {
        request = NewModifyPullStreamStatusRequest()
    }
    response = NewModifyPullStreamStatusResponse()
    err = c.Send(request, response)
    return
}

func NewResumeDelayLiveStreamRequest() (request *ResumeDelayLiveStreamRequest) {
    request = &ResumeDelayLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ResumeDelayLiveStream")
    return
}

func NewResumeDelayLiveStreamResponse() (response *ResumeDelayLiveStreamResponse) {
    response = &ResumeDelayLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 恢复延迟播放设置
func (c *Client) ResumeDelayLiveStream(request *ResumeDelayLiveStreamRequest) (response *ResumeDelayLiveStreamResponse, err error) {
    if request == nil {
        request = NewResumeDelayLiveStreamRequest()
    }
    response = NewResumeDelayLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewResumeLiveStreamRequest() (request *ResumeLiveStreamRequest) {
    request = &ResumeLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ResumeLiveStream")
    return
}

func NewResumeLiveStreamResponse() (response *ResumeLiveStreamResponse) {
    response = &ResumeLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 恢复某条流的推送。
func (c *Client) ResumeLiveStream(request *ResumeLiveStreamRequest) (response *ResumeLiveStreamResponse, err error) {
    if request == nil {
        request = NewResumeLiveStreamRequest()
    }
    response = NewResumeLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewSetLiveWatermarkStatusRequest() (request *SetLiveWatermarkStatusRequest) {
    request = &SetLiveWatermarkStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "SetLiveWatermarkStatus")
    return
}

func NewSetLiveWatermarkStatusResponse() (response *SetLiveWatermarkStatusResponse) {
    response = &SetLiveWatermarkStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置水印是否启用
func (c *Client) SetLiveWatermarkStatus(request *SetLiveWatermarkStatusRequest) (response *SetLiveWatermarkStatusResponse, err error) {
    if request == nil {
        request = NewSetLiveWatermarkStatusRequest()
    }
    response = NewSetLiveWatermarkStatusResponse()
    err = c.Send(request, response)
    return
}

func NewStopLiveRecordRequest() (request *StopLiveRecordRequest) {
    request = &StopLiveRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "StopLiveRecord")
    return
}

func NewStopLiveRecordResponse() (response *StopLiveRecordResponse) {
    response = &StopLiveRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 说明：录制后的文件存放于点播平台。用户如需使用录制功能，需首先自行开通点播账号并确保账号可用。录制文件存放后，相关费用（含存储以及下行播放流量）按照点播平台计费方式收取，请参考对应文档。
func (c *Client) StopLiveRecord(request *StopLiveRecordRequest) (response *StopLiveRecordResponse, err error) {
    if request == nil {
        request = NewStopLiveRecordRequest()
    }
    response = NewStopLiveRecordResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindLiveDomainCertRequest() (request *UnBindLiveDomainCertRequest) {
    request = &UnBindLiveDomainCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "UnBindLiveDomainCert")
    return
}

func NewUnBindLiveDomainCertResponse() (response *UnBindLiveDomainCertResponse) {
    response = &UnBindLiveDomainCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 解绑域名证书
func (c *Client) UnBindLiveDomainCert(request *UnBindLiveDomainCertRequest) (response *UnBindLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewUnBindLiveDomainCertRequest()
    }
    response = NewUnBindLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateLiveWatermarkRequest() (request *UpdateLiveWatermarkRequest) {
    request = &UpdateLiveWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "UpdateLiveWatermark")
    return
}

func NewUpdateLiveWatermarkResponse() (response *UpdateLiveWatermarkResponse) {
    response = &UpdateLiveWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新水印
func (c *Client) UpdateLiveWatermark(request *UpdateLiveWatermarkRequest) (response *UpdateLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewUpdateLiveWatermarkRequest()
    }
    response = NewUpdateLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}
