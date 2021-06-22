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

// AddDelayLiveStream
// 对流设置延播时间
//
// 注意：如果在推流前设置延播，需要提前5分钟设置。
//
// 目前该接口只支持流粒度的，域名及应用粒度功能支持当前开发中。
//
// 使用场景：对重要直播，避免出现突发状况，可通过设置延迟播放，提前做好把控。
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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

// AddLiveDomain
// 添加域名，一次只能提交一个域名。域名必须已备案。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEDOMAININLOCKEDTIME = "FailedOperation.DeleteDomainInLockedTime"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHINESECHARACTERDETECTED = "InternalError.ChineseCharacterDetected"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINALREADYEXIST = "InternalError.DomainAlreadyExist"
//  INTERNALERROR_DOMAINFORMATERROR = "InternalError.DomainFormatError"
//  INTERNALERROR_DOMAINGSLBFAIL = "InternalError.DomainGslbFail"
//  INTERNALERROR_DOMAINISFAMOUS = "InternalError.DomainIsFamous"
//  INTERNALERROR_DOMAINISLIMITED = "InternalError.DomainIsLimited"
//  INTERNALERROR_DOMAINNORECORD = "InternalError.DomainNoRecord"
//  INTERNALERROR_DOMAINTOOLONG = "InternalError.DomainTooLong"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER_DOMAINALREADYEXIST = "InvalidParameter.DomainAlreadyExist"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETER_DOMAINISFAMOUS = "InvalidParameter.DomainIsFamous"
//  INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"
//  INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"
//  INVALIDPARAMETER_MPHOSTDELETE = "InvalidParameter.MpHostDelete"
//  INVALIDPARAMETER_MPPLUGINNOUSE = "InvalidParameter.MpPluginNoUse"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
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

// AddLiveWatermark
// 添加水印，成功返回水印 ID 后，需要调用[CreateLiveWatermarkRule](/document/product/267/32629)接口将水印 ID 绑定到流使用。
//
// 水印数量上限 100，超过后需要先删除，再添加。
//
// 可能返回的错误码:
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_WATERMARKADDERROR = "InternalError.WatermarkAddError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_USERNOTFOUNT = "ResourceNotFound.UserNotFount"
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

// BindLiveDomainCert
// 域名绑定证书。
//
// 注意：需先调用添加证书接口进行证书添加。获取到证书Id后再调用该接口进行绑定。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CRTDATEINUSING = "InvalidParameter.CrtDateInUsing"
//  INVALIDPARAMETER_CRTDATENOTLEGAL = "InvalidParameter.CrtDateNotLegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
func (c *Client) BindLiveDomainCert(request *BindLiveDomainCertRequest) (response *BindLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewBindLiveDomainCertRequest()
    }
    response = NewBindLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewCancelCommonMixStreamRequest() (request *CancelCommonMixStreamRequest) {
    request = &CancelCommonMixStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CancelCommonMixStream")
    return
}

func NewCancelCommonMixStreamResponse() (response *CancelCommonMixStreamResponse) {
    response = &CancelCommonMixStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelCommonMixStream
// 该接口用来取消混流。用法与 mix_streamv2.cancel_mix_stream 基本一致。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRERROR = "FailedOperation.CallOtherSvrError"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_CANCELSESSIONNOTEXIST = "FailedOperation.CancelSessionNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CANCELSESSIONNOTEXIST = "InvalidParameter.CancelSessionNotExist"
//  INVALIDPARAMETER_OTHERERROR = "InvalidParameter.OtherError"
func (c *Client) CancelCommonMixStream(request *CancelCommonMixStreamRequest) (response *CancelCommonMixStreamResponse, err error) {
    if request == nil {
        request = NewCancelCommonMixStreamRequest()
    }
    response = NewCancelCommonMixStreamResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCommonMixStreamRequest() (request *CreateCommonMixStreamRequest) {
    request = &CreateCommonMixStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateCommonMixStream")
    return
}

func NewCreateCommonMixStreamResponse() (response *CreateCommonMixStreamResponse) {
    response = &CreateCommonMixStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCommonMixStream
// 该接口用来创建通用混流。用法与旧接口 mix_streamv2.start_mix_stream_advanced 基本一致。
//
// 注意：当前最多支持16路混流。
//
// 最佳实践：https://cloud.tencent.com/document/product/267/45566
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRERROR = "FailedOperation.CallOtherSvrError"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  FAILEDOPERATION_GETPICTUREURLERROR = "FailedOperation.GetPictureUrlError"
//  FAILEDOPERATION_GETSTREAMRESOLUTIONERROR = "FailedOperation.GetStreamResolutionError"
//  FAILEDOPERATION_PROCESSMIXERROR = "FailedOperation.ProcessMixError"
//  FAILEDOPERATION_STREAMNOTEXIST = "FailedOperation.StreamNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_JIFEIOTHERERROR = "InternalError.JiFeiOtherError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CANCELSESSIONNOTEXIST = "InvalidParameter.CancelSessionNotExist"
//  INVALIDPARAMETER_INPUTNUMLIMITEXCEEDED = "InvalidParameter.InputNumLimitExceeded"
//  INVALIDPARAMETER_INVALIDBACKGROUDRESOLUTION = "InvalidParameter.InvalidBackgroudResolution"
//  INVALIDPARAMETER_INVALIDBITRATE = "InvalidParameter.InvalidBitrate"
//  INVALIDPARAMETER_INVALIDCROPPARAM = "InvalidParameter.InvalidCropParam"
//  INVALIDPARAMETER_INVALIDLAYERPARAM = "InvalidParameter.InvalidLayerParam"
//  INVALIDPARAMETER_INVALIDOUTPUTSTREAMID = "InvalidParameter.InvalidOutputStreamID"
//  INVALIDPARAMETER_INVALIDOUTPUTTYPE = "InvalidParameter.InvalidOutputType"
//  INVALIDPARAMETER_INVALIDPICTUREID = "InvalidParameter.InvalidPictureID"
//  INVALIDPARAMETER_INVALIDROUNDRECTRADIUS = "InvalidParameter.InvalidRoundRectRadius"
//  INVALIDPARAMETER_OTHERERROR = "InvalidParameter.OtherError"
//  INVALIDPARAMETER_SESSIONOUTPUTSTREAMCHANGED = "InvalidParameter.SessionOutputStreamChanged"
//  INVALIDPARAMETER_TEMPLATENOTMATCHINPUTNUM = "InvalidParameter.TemplateNotMatchInputNum"
func (c *Client) CreateCommonMixStream(request *CreateCommonMixStreamRequest) (response *CreateCommonMixStreamResponse, err error) {
    if request == nil {
        request = NewCreateCommonMixStreamRequest()
    }
    response = NewCreateCommonMixStreamResponse()
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

// CreateLiveCallbackRule
// 创建回调规则，需要先调用[CreateLiveCallbackTemplate](/document/product/267/32637)接口创建回调模板，将返回的模板id绑定到域名/路径进行使用。
//
// <br>回调协议相关文档：[事件消息通知](/document/product/267/32744)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
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

// CreateLiveCallbackTemplate
// 创建回调模板，成功返回模板id后，需要调用[CreateLiveCallbackRule](/document/product/267/32638)接口将模板 ID 绑定到域名/路径使用。
//
// <br>回调协议相关文档：[事件消息通知](/document/product/267/32744)。
//
// 注意：至少填写一个回调 URL。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETER_URLNOTSAFE = "InvalidParameter.UrlNotSafe"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
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

// CreateLiveCert
// 添加证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CRTDATENOTLEGAL = "InternalError.CrtDateNotLegal"
//  INTERNALERROR_CRTDATEOVERDUE = "InternalError.CrtDateOverdue"
//  INTERNALERROR_CRTKEYNOTMATCH = "InternalError.CrtKeyNotMatch"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLOUDCRTIDERROR = "InvalidParameter.CloudCrtIdError"
//  INVALIDPARAMETER_CRTDATENOTLEGAL = "InvalidParameter.CrtDateNotLegal"
//  INVALIDPARAMETER_CRTDATEOVERDUE = "InvalidParameter.CrtDateOverdue"
//  INVALIDPARAMETER_CRTKEYNOTMATCH = "InvalidParameter.CrtKeyNotMatch"
//  INVALIDPARAMETER_CRTORKEYNOTEXIST = "InvalidParameter.CrtOrKeyNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateLiveCert(request *CreateLiveCertRequest) (response *CreateLiveCertResponse, err error) {
    if request == nil {
        request = NewCreateLiveCertRequest()
    }
    response = NewCreateLiveCertResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLivePullStreamTaskRequest() (request *CreateLivePullStreamTaskRequest) {
    request = &CreateLivePullStreamTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateLivePullStreamTask")
    return
}

func NewCreateLivePullStreamTaskResponse() (response *CreateLivePullStreamTaskResponse) {
    response = &CreateLivePullStreamTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLivePullStreamTask
// 创建直播拉流任务。支持将外部已有的点播文件，或者直播源拉取过来转推到直播系统。
//
// 注意：
//
// 1. 源流视频编码目前只支持: H264, H265。其他编码格式建议先进行转码处理。
//
// 2. 源流音频编码目前只支持: AAC。其他编码格式建议先进行转码处理。
//
// 3. 拉流转推功能为计费增值服务，计费规则详情可参见[计费文档](https://cloud.tencent.com/document/product/267/53308)。
//
// 4. 拉流转推功能仅提供内容拉取与推送服务，请确保内容已获得授权并符合内容传播相关的法律法规。若内容有侵权或违规相关问题，云直播会停止相关的功能服务并保留追究法律责任的权利。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCALLBACKURL = "InvalidParameter.InvalidCallbackUrl"
//  INVALIDPARAMETER_INVALIDSOURCEURL = "InvalidParameter.InvalidSourceUrl"
//  INVALIDPARAMETER_INVALIDTASKTIME = "InvalidParameter.InvalidTaskTime"
//  INVALIDPARAMETER_INVALIDTOURL = "InvalidParameter.InvalidToUrl"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  INVALIDPARAMETER_TASKNUMMORETHANLIMIT = "InvalidParameter.TaskNumMoreThanLimit"
//  INVALIDPARAMETER_TOURLNOPERMISSION = "InvalidParameter.ToUrlNoPermission"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateLivePullStreamTask(request *CreateLivePullStreamTaskRequest) (response *CreateLivePullStreamTaskResponse, err error) {
    if request == nil {
        request = NewCreateLivePullStreamTaskRequest()
    }
    response = NewCreateLivePullStreamTaskResponse()
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

// CreateLiveRecord
// - 使用前提
//
//   1. 录制文件存放于点播平台，所以用户如需使用录制功能，需首先自行开通点播服务。
//
//   2. 录制文件存放后相关费用（含存储以及下行播放流量）按照点播平台计费方式收取，具体请参考 [对应文档](https://cloud.tencent.com/document/product/266/2838)。
//
// 
//
// - 模式说明
//
//   该接口支持两种录制模式：
//
//   1. 定时录制模式【默认模式】。
//
//     需要传入开始时间与结束时间，录制任务根据起止时间自动开始与结束。在所设置结束时间过期之前（且未调用StopLiveRecord提前终止任务），录制任务都是有效的，期间多次断流然后重推都会启动录制任务。
//
//   2. 实时视频录制模式。
//
//     忽略传入的开始时间，在录制任务创建后立即开始录制，录制时长支持最大为30分钟，如果传入的结束时间与当前时间差大于30分钟，则按30分钟计算，实时视频录制主要用于录制精彩视频场景，时长建议控制在5分钟以内。
//
// 
//
// - 注意事项
//
//   1. 调用接口超时设置应大于3秒，小于3秒重试以及按不同起止时间调用都有可能产生重复录制任务，进而导致额外录制费用。
//
//   2. 受限于音视频文件格式（FLV/MP4/HLS）对编码类型的支持，视频编码类型支持 H.264，音频编码类型支持 AAC。
//
//   3. 为避免恶意或非主观的频繁 API 请求，对定时录制模式最大创建任务数做了限制：其中，当天可以创建的最大任务数不超过4000（不含已删除的任务）；当前时刻并发运行的任务数不超过400。有超出此限制的需要提工单申请。
//
//   4. 此调用方式暂时不支持海外推流录制。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  LIMITEXCEEDED_MAXIMUMTASK = "LimitExceeded.MaximumTask"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  RESOURCEUNAVAILABLE_STREAMNOTEXIST = "ResourceUnavailable.StreamNotExist"
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

// CreateLiveRecordRule
// 创建录制规则，需要先调用[CreateLiveRecordTemplate](/document/product/267/32614)接口创建录制模板，将返回的模板id绑定到流使用。
//
// <br>录制相关文档：[直播录制](/document/product/267/32739)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// CreateLiveRecordTemplate
// 创建录制模板，成功返回模板id后，需要调用[CreateLiveRecordRule](/document/product/267/32615)接口，将模板id绑定到流进行使用。
//
// <br>录制相关文档：[直播录制](/document/product/267/32739)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
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

// CreateLiveSnapshotRule
// 创建截图规则，需要先调用[CreateLiveSnapshotTemplate](/document/product/267/32624)接口创建截图模板，然后将返回的模板 ID 绑定到流进行使用。
//
// <br>截图相关文档：[直播截图](/document/product/267/32737)。
//
// 注意：单个域名仅支持关联一个截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// CreateLiveSnapshotTemplate
// 创建截图模板，成功返回模板id后，需要调用[CreateLiveSnapshotRule](/document/product/267/32625)接口，将模板id绑定到流使用。
//
// <br>截图相关文档：[直播截图](/document/product/267/32737)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
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

// CreateLiveTranscodeRule
// 创建转码规则，需要先调用[CreateLiveTranscodeTemplate](/document/product/267/32646)接口创建转码模板，将返回的模板id绑定到流使用。
//
// <br>转码相关文档：[直播转封装及转码](/document/product/267/32736)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// CreateLiveTranscodeTemplate
// 创建转码模板，成功返回模板id后，需要调用[CreateLiveTranscodeRule](/document/product/267/32647)接口，将返回的模板id绑定到流使用。
//
// <br>转码相关文档：[直播转封装及转码](/document/product/267/32736)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_PROCESSORALREADYEXIST = "InternalError.ProcessorAlreadyExist"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
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

// CreateLiveWatermarkRule
// 创建水印规则，需要先调用[AddLiveWatermark](/document/product/267/30154)接口添加水印，将返回的水印id绑定到流使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// CreatePullStreamConfig
// 创建临时拉流转推任务，目前限制添加10条任务。
//
// 
//
// 注意：该接口用于创建临时拉流转推任务，
//
// 拉流源地址即 FromUrl 可以是腾讯或非腾讯数据源，
//
// 但转推目标地址即 ToUrl 目前限制为已注册的腾讯直播域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreatePullStreamConfig(request *CreatePullStreamConfigRequest) (response *CreatePullStreamConfigResponse, err error) {
    if request == nil {
        request = NewCreatePullStreamConfigRequest()
    }
    response = NewCreatePullStreamConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordTaskRequest() (request *CreateRecordTaskRequest) {
    request = &CreateRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "CreateRecordTask")
    return
}

func NewCreateRecordTaskResponse() (response *CreateRecordTaskResponse) {
    response = &CreateRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRecordTask
// 创建一个在指定时间启动、结束的录制任务，并使用指定录制模板ID对应的配置进行录制。
//
// - 使用前提
//
// 1. 录制文件存放于点播平台，所以用户如需使用录制功能，需首先自行开通点播服务。
//
// 2. 录制文件存放后相关费用（含存储以及下行播放流量）按照点播平台计费方式收取，具体请参考 [对应文档](https://cloud.tencent.com/document/product/266/2837)。
//
// - 注意事项
//
// 1. 断流会结束当前录制并生成录制文件。在结束时间到达之前任务仍然有效，期间只要正常推流都会正常录制，与是否多次推、断流无关。
//
// 2. 使用上避免创建时间段相互重叠的录制任务。若同一条流当前存在多个时段重叠的任务，为避免重复录制系统将启动最多3个录制任务。
//
// 3. 创建的录制任务记录在平台侧只保留3个月。
//
// 4. 当前录制任务管理API（CreateRecordTask/StopRecordTask/DeleteRecordTask）与旧API（CreateLiveRecord/StopLiveRecord/DeleteLiveRecord）不兼容，两套接口不能混用。
//
// 5. 避免 创建录制任务 与 推流 操作同时进行，可能导致因录制任务未生效而引起任务延迟启动问题，两者操作间隔建议大于3秒。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRecordTask(request *CreateRecordTaskRequest) (response *CreateRecordTaskResponse, err error) {
    if request == nil {
        request = NewCreateRecordTaskRequest()
    }
    response = NewCreateRecordTaskResponse()
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

// DeleteLiveCallbackRule
// 删除回调规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DeleteLiveCallbackTemplate
// 删除回调模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DeleteLiveCert
// 删除域名对应的证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CRTDATEINUSING = "InternalError.CrtDateInUsing"
//  INTERNALERROR_CRTDATENOTFOUND = "InternalError.CrtDateNotFound"
//  INTERNALERROR_CRTDATENOTLEGAL = "InternalError.CrtDateNotLegal"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INVALIDPARAMETER_CRTDATEINUSING = "InvalidParameter.CrtDateInUsing"
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

// DeleteLiveDomain
// 删除已添加的直播域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEDOMAININLOCKEDTIME = "FailedOperation.DeleteDomainInLockedTime"
//  FAILEDOPERATION_JIFEINOENOUGHFUND = "FailedOperation.JiFeiNoEnoughFund"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
//  RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"
func (c *Client) DeleteLiveDomain(request *DeleteLiveDomainRequest) (response *DeleteLiveDomainResponse, err error) {
    if request == nil {
        request = NewDeleteLiveDomainRequest()
    }
    response = NewDeleteLiveDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLivePullStreamTaskRequest() (request *DeleteLivePullStreamTaskRequest) {
    request = &DeleteLivePullStreamTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteLivePullStreamTask")
    return
}

func NewDeleteLivePullStreamTaskResponse() (response *DeleteLivePullStreamTaskResponse) {
    response = &DeleteLivePullStreamTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLivePullStreamTask
// 删除接口 CreateLivePullStreamTask 创建的拉流任务。
//
// 注意：
//
// 1. 入参中的 TaskId 为 CreateLivePullStreamTask 接口创建时返回的TaskId。
//
// 2. 也可通过 DescribeLivePullStreamTasks 进行查询创建的任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteLivePullStreamTask(request *DeleteLivePullStreamTaskRequest) (response *DeleteLivePullStreamTaskResponse, err error) {
    if request == nil {
        request = NewDeleteLivePullStreamTaskRequest()
    }
    response = NewDeleteLivePullStreamTaskResponse()
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

// DeleteLiveRecord
// 注：DeleteLiveRecord 接口仅用于删除录制任务记录，不具备停止录制的功能，也不能删除正在进行中的录制。如果需要停止录制任务，请使用终止录制[StopLiveRecord](/document/product/267/30146) 接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ALTERTASKSTATE = "FailedOperation.AlterTaskState"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"
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

// DeleteLiveRecordRule
// 删除录制规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DeleteLiveRecordTemplate
// 删除录制模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DeleteLiveSnapshotRule
// 删除截图规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DeleteLiveSnapshotTemplate
// 删除截图模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DeleteLiveTranscodeRule
// 删除转码规则。
//
// DomainName+AppName+StreamName+TemplateId唯一标识单个转码规则，如需删除需要强匹配。其中TemplateId必填，其余参数为空时也需要传空字符串进行强匹配。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DeleteLiveTranscodeTemplate
// 删除转码模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DeleteLiveWatermark
// 删除水印。
//
// 可能返回的错误码:
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETWATERMARKERROR = "InternalError.GetWatermarkError"
//  INTERNALERROR_WATERMARKNOTEXIST = "InternalError.WatermarkNotExist"
//  RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"
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

// DeleteLiveWatermarkRule
// 删除水印规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DeletePullStreamConfig
// 删除直播拉流配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeletePullStreamConfig(request *DeletePullStreamConfigRequest) (response *DeletePullStreamConfigResponse, err error) {
    if request == nil {
        request = NewDeletePullStreamConfigRequest()
    }
    response = NewDeletePullStreamConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordTaskRequest() (request *DeleteRecordTaskRequest) {
    request = &DeleteRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DeleteRecordTask")
    return
}

func NewDeleteRecordTaskResponse() (response *DeleteRecordTaskResponse) {
    response = &DeleteRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecordTask
// 删除录制任务配置。删除操作不影响正在运行当中的任务，仅对删除之后新的推流有效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRecordTask(request *DeleteRecordTaskRequest) (response *DeleteRecordTaskResponse, err error) {
    if request == nil {
        request = NewDeleteRecordTaskRequest()
    }
    response = NewDeleteRecordTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllStreamPlayInfoListRequest() (request *DescribeAllStreamPlayInfoListRequest) {
    request = &DescribeAllStreamPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeAllStreamPlayInfoList")
    return
}

func NewDescribeAllStreamPlayInfoListResponse() (response *DescribeAllStreamPlayInfoListResponse) {
    response = &DescribeAllStreamPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAllStreamPlayInfoList
// 输入某个时间点（1分钟维度），查询该时间点所有流的下行信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAllStreamPlayInfoList(request *DescribeAllStreamPlayInfoListRequest) (response *DescribeAllStreamPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeAllStreamPlayInfoListRequest()
    }
    response = NewDescribeAllStreamPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAreaBillBandwidthAndFluxListRequest() (request *DescribeAreaBillBandwidthAndFluxListRequest) {
    request = &DescribeAreaBillBandwidthAndFluxListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeAreaBillBandwidthAndFluxList")
    return
}

func NewDescribeAreaBillBandwidthAndFluxListResponse() (response *DescribeAreaBillBandwidthAndFluxListResponse) {
    response = &DescribeAreaBillBandwidthAndFluxListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAreaBillBandwidthAndFluxList
// 海外分区直播计费带宽和流量数据查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAreaBillBandwidthAndFluxList(request *DescribeAreaBillBandwidthAndFluxListRequest) (response *DescribeAreaBillBandwidthAndFluxListResponse, err error) {
    if request == nil {
        request = NewDescribeAreaBillBandwidthAndFluxListRequest()
    }
    response = NewDescribeAreaBillBandwidthAndFluxListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillBandwidthAndFluxListRequest() (request *DescribeBillBandwidthAndFluxListRequest) {
    request = &DescribeBillBandwidthAndFluxListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeBillBandwidthAndFluxList")
    return
}

func NewDescribeBillBandwidthAndFluxListResponse() (response *DescribeBillBandwidthAndFluxListResponse) {
    response = &DescribeBillBandwidthAndFluxListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillBandwidthAndFluxList
// 直播计费带宽和流量数据查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBillBandwidthAndFluxList(request *DescribeBillBandwidthAndFluxListRequest) (response *DescribeBillBandwidthAndFluxListResponse, err error) {
    if request == nil {
        request = NewDescribeBillBandwidthAndFluxListRequest()
    }
    response = NewDescribeBillBandwidthAndFluxListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallbackRecordsListRequest() (request *DescribeCallbackRecordsListRequest) {
    request = &DescribeCallbackRecordsListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeCallbackRecordsList")
    return
}

func NewDescribeCallbackRecordsListResponse() (response *DescribeCallbackRecordsListResponse) {
    response = &DescribeCallbackRecordsListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCallbackRecordsList
// 回调事件查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeCallbackRecordsList(request *DescribeCallbackRecordsListRequest) (response *DescribeCallbackRecordsListResponse, err error) {
    if request == nil {
        request = NewDescribeCallbackRecordsListRequest()
    }
    response = NewDescribeCallbackRecordsListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrentRecordStreamNumRequest() (request *DescribeConcurrentRecordStreamNumRequest) {
    request = &DescribeConcurrentRecordStreamNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeConcurrentRecordStreamNum")
    return
}

func NewDescribeConcurrentRecordStreamNumResponse() (response *DescribeConcurrentRecordStreamNumResponse) {
    response = &DescribeConcurrentRecordStreamNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConcurrentRecordStreamNum
// 查询并发录制路数，对慢直播和普通直播适用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeConcurrentRecordStreamNum(request *DescribeConcurrentRecordStreamNumRequest) (response *DescribeConcurrentRecordStreamNumResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrentRecordStreamNumRequest()
    }
    response = NewDescribeConcurrentRecordStreamNumResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeliverBandwidthListRequest() (request *DescribeDeliverBandwidthListRequest) {
    request = &DescribeDeliverBandwidthListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeDeliverBandwidthList")
    return
}

func NewDescribeDeliverBandwidthListResponse() (response *DescribeDeliverBandwidthListResponse) {
    response = &DescribeDeliverBandwidthListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeliverBandwidthList
// 查询直播转推计费带宽，查询时间范围最大支持3个月内的数据，时间跨度最长31天。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeliverBandwidthList(request *DescribeDeliverBandwidthListRequest) (response *DescribeDeliverBandwidthListResponse, err error) {
    if request == nil {
        request = NewDescribeDeliverBandwidthListRequest()
    }
    response = NewDescribeDeliverBandwidthListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupProIspPlayInfoListRequest() (request *DescribeGroupProIspPlayInfoListRequest) {
    request = &DescribeGroupProIspPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeGroupProIspPlayInfoList")
    return
}

func NewDescribeGroupProIspPlayInfoListResponse() (response *DescribeGroupProIspPlayInfoListResponse) {
    response = &DescribeGroupProIspPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupProIspPlayInfoList
// 查询按省份和运营商分组的下行播放数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupProIspPlayInfoList(request *DescribeGroupProIspPlayInfoListRequest) (response *DescribeGroupProIspPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeGroupProIspPlayInfoListRequest()
    }
    response = NewDescribeGroupProIspPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHttpStatusInfoListRequest() (request *DescribeHttpStatusInfoListRequest) {
    request = &DescribeHttpStatusInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeHttpStatusInfoList")
    return
}

func NewDescribeHttpStatusInfoListResponse() (response *DescribeHttpStatusInfoListResponse) {
    response = &DescribeHttpStatusInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHttpStatusInfoList
// 查询某段时间内5分钟粒度的各播放http状态码的个数。
//
// 备注：数据延迟1小时，如10:00-10:59点的数据12点才能查到。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeHttpStatusInfoList(request *DescribeHttpStatusInfoListRequest) (response *DescribeHttpStatusInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeHttpStatusInfoListRequest()
    }
    response = NewDescribeHttpStatusInfoListResponse()
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

// DescribeLiveCallbackRules
// 获取回调规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
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

// DescribeLiveCallbackTemplate
// 获取单个回调模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribeLiveCallbackTemplates
// 获取回调模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
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

// DescribeLiveCert
// 获取证书信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CRTDOMAINNOTFOUND = "InternalError.CrtDomainNotFound"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
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

// DescribeLiveCerts
// 获取证书信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
func (c *Client) DescribeLiveCerts(request *DescribeLiveCertsRequest) (response *DescribeLiveCertsResponse, err error) {
    if request == nil {
        request = NewDescribeLiveCertsRequest()
    }
    response = NewDescribeLiveCertsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDelayInfoListRequest() (request *DescribeLiveDelayInfoListRequest) {
    request = &DescribeLiveDelayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDelayInfoList")
    return
}

func NewDescribeLiveDelayInfoListResponse() (response *DescribeLiveDelayInfoListResponse) {
    response = &DescribeLiveDelayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveDelayInfoList
// 获取直播延播列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveDelayInfoList(request *DescribeLiveDelayInfoListRequest) (response *DescribeLiveDelayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDelayInfoListRequest()
    }
    response = NewDescribeLiveDelayInfoListResponse()
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

// DescribeLiveDomain
// 查询直播域名信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
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

// DescribeLiveDomainCert
// 获取域名证书信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
func (c *Client) DescribeLiveDomainCert(request *DescribeLiveDomainCertRequest) (response *DescribeLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainCertRequest()
    }
    response = NewDescribeLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainPlayInfoListRequest() (request *DescribeLiveDomainPlayInfoListRequest) {
    request = &DescribeLiveDomainPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomainPlayInfoList")
    return
}

func NewDescribeLiveDomainPlayInfoListResponse() (response *DescribeLiveDomainPlayInfoListResponse) {
    response = &DescribeLiveDomainPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveDomainPlayInfoList
// 查询实时的域名维度下行播放数据，由于数据处理有耗时，接口默认查询4分钟前的准实时数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLiveDomainPlayInfoList(request *DescribeLiveDomainPlayInfoListRequest) (response *DescribeLiveDomainPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainPlayInfoListRequest()
    }
    response = NewDescribeLiveDomainPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveDomainRefererRequest() (request *DescribeLiveDomainRefererRequest) {
    request = &DescribeLiveDomainRefererRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLiveDomainReferer")
    return
}

func NewDescribeLiveDomainRefererResponse() (response *DescribeLiveDomainRefererResponse) {
    response = &DescribeLiveDomainRefererResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveDomainReferer
// 查询直播域名 Referer 黑白名单配置。
//
// 由于 Referer 信息包含在 http 协议中，在开启配置后，播放协议为 rtmp 或 webrtc 不会校验 Referer 配置，仍可正常播放。如需配置 Referer 鉴权建议使用 http-flv 或 http-hls 协议播放。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
func (c *Client) DescribeLiveDomainReferer(request *DescribeLiveDomainRefererRequest) (response *DescribeLiveDomainRefererResponse, err error) {
    if request == nil {
        request = NewDescribeLiveDomainRefererRequest()
    }
    response = NewDescribeLiveDomainRefererResponse()
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

// DescribeLiveDomains
// 根据域名状态、类型等信息查询用户的域名信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
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

// DescribeLiveForbidStreamList
// 获取禁推流列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveForbidStreamList(request *DescribeLiveForbidStreamListRequest) (response *DescribeLiveForbidStreamListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveForbidStreamListRequest()
    }
    response = NewDescribeLiveForbidStreamListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePackageInfoRequest() (request *DescribeLivePackageInfoRequest) {
    request = &DescribeLivePackageInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePackageInfo")
    return
}

func NewDescribeLivePackageInfoResponse() (response *DescribeLivePackageInfoResponse) {
    response = &DescribeLivePackageInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLivePackageInfo
// 查询用户套餐包总量、使用量、剩余量、包状态、购买时间和过期时间等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLivePackageInfo(request *DescribeLivePackageInfoRequest) (response *DescribeLivePackageInfoResponse, err error) {
    if request == nil {
        request = NewDescribeLivePackageInfoRequest()
    }
    response = NewDescribeLivePackageInfoResponse()
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

// DescribeLivePlayAuthKey
// 查询播放鉴权key。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PLAYDOMAINNORECORD = "InternalError.PlayDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"
func (c *Client) DescribeLivePlayAuthKey(request *DescribeLivePlayAuthKeyRequest) (response *DescribeLivePlayAuthKeyResponse, err error) {
    if request == nil {
        request = NewDescribeLivePlayAuthKeyRequest()
    }
    response = NewDescribeLivePlayAuthKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLivePullStreamTasksRequest() (request *DescribeLivePullStreamTasksRequest) {
    request = &DescribeLivePullStreamTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeLivePullStreamTasks")
    return
}

func NewDescribeLivePullStreamTasksResponse() (response *DescribeLivePullStreamTasksResponse) {
    response = &DescribeLivePullStreamTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLivePullStreamTasks
// 查询使用 CreateLivePullStreamTask 接口创建的直播拉流任务。
//
// 排序方式：默认按更新时间 倒序排列。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLivePullStreamTasks(request *DescribeLivePullStreamTasksRequest) (response *DescribeLivePullStreamTasksResponse, err error) {
    if request == nil {
        request = NewDescribeLivePullStreamTasksRequest()
    }
    response = NewDescribeLivePullStreamTasksResponse()
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

// DescribeLivePushAuthKey
// 查询直播推流鉴权key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PUSHDOMAINNORECORD = "InternalError.PushDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"
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

// DescribeLiveRecordRules
// 获取录制规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribeLiveRecordTemplate
// 获取单个录制模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribeLiveRecordTemplates
// 获取录制模板列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribeLiveSnapshotRules
// 获取截图规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribeLiveSnapshotTemplate
// 获取单个截图模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribeLiveSnapshotTemplates
// 获取截图模板列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
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

// DescribeLiveStreamEventList
// 用于查询推断流事件。<br>
//
// 
//
// 注意：该接口可通过使用IsFilter进行过滤，返回推流历史记录。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeLiveStreamEventList(request *DescribeLiveStreamEventListRequest) (response *DescribeLiveStreamEventListResponse, err error) {
    if request == nil {
        request = NewDescribeLiveStreamEventListRequest()
    }
    response = NewDescribeLiveStreamEventListResponse()
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

// DescribeLiveStreamOnlineList
// 返回正在直播中的流列表。适用于推流成功后查询在线流信息。
//
// 注意：该接口仅适用于流数少于2万路的情况，对于流数较大用户请联系售后。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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

// DescribeLiveStreamPublishedList
// 返回已经推过流的流列表。<br>
//
// 注意：分页最多支持查询1万条记录，可通过调整查询时间范围来获取更多数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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

// DescribeLiveStreamPushInfoList
// 查询所有实时流的推流信息，包括客户端IP，服务端IP，帧率，码率，域名，开始推流时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION_HASNOTLIVINGSTREAM = "FailedOperation.HasNotLivingStream"
//  FAILEDOPERATION_QUERYUPLOADINFOFAILED = "FailedOperation.QueryUploadInfoFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDREQUEST = "InternalError.InvalidRequest"
//  INTERNALERROR_QUERYPROISPPLAYINFOERROR = "InternalError.QueryProIspPlayInfoError"
//  INTERNALERROR_QUERYUPLOADINFOFAILED = "InternalError.QueryUploadInfoFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribeLiveStreamState
// 返回直播中、无推流或者禁播等状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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

// DescribeLiveTranscodeDetailInfo
// 支持查询某天或某段时间的转码详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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

// DescribeLiveTranscodeRules
// 获取转码规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribeLiveTranscodeTemplate
// 获取单个转码模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
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

// DescribeLiveTranscodeTemplates
// 获取转码模板列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
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

// DescribeLiveWatermark
// 获取单个水印信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribeLiveWatermarkRules
// 获取水印规则列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribeLiveWatermarks
// 查询水印列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribeLogDownloadList
// 批量获取日志URL。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVOKECDNAPIFAIL = "FailedOperation.InvokeCdnApiFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CDNLOGEMPTY = "ResourceNotFound.CdnLogEmpty"
//  RESOURCENOTFOUND_CDNTHEMEEMPTY = "ResourceNotFound.CdnThemeEmpty"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
func (c *Client) DescribeLogDownloadList(request *DescribeLogDownloadListRequest) (response *DescribeLogDownloadListResponse, err error) {
    if request == nil {
        request = NewDescribeLogDownloadListRequest()
    }
    response = NewDescribeLogDownloadListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlayErrorCodeDetailInfoListRequest() (request *DescribePlayErrorCodeDetailInfoListRequest) {
    request = &DescribePlayErrorCodeDetailInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribePlayErrorCodeDetailInfoList")
    return
}

func NewDescribePlayErrorCodeDetailInfoListResponse() (response *DescribePlayErrorCodeDetailInfoListResponse) {
    response = &DescribePlayErrorCodeDetailInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePlayErrorCodeDetailInfoList
// 查询下行播放错误码信息，某段时间内1分钟粒度的各http错误码出现的次数，包括4xx，5xx。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribePlayErrorCodeDetailInfoList(request *DescribePlayErrorCodeDetailInfoListRequest) (response *DescribePlayErrorCodeDetailInfoListResponse, err error) {
    if request == nil {
        request = NewDescribePlayErrorCodeDetailInfoListRequest()
    }
    response = NewDescribePlayErrorCodeDetailInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlayErrorCodeSumInfoListRequest() (request *DescribePlayErrorCodeSumInfoListRequest) {
    request = &DescribePlayErrorCodeSumInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribePlayErrorCodeSumInfoList")
    return
}

func NewDescribePlayErrorCodeSumInfoListResponse() (response *DescribePlayErrorCodeSumInfoListResponse) {
    response = &DescribePlayErrorCodeSumInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePlayErrorCodeSumInfoList
// 查询下行播放错误码信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribePlayErrorCodeSumInfoList(request *DescribePlayErrorCodeSumInfoListRequest) (response *DescribePlayErrorCodeSumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribePlayErrorCodeSumInfoListRequest()
    }
    response = NewDescribePlayErrorCodeSumInfoListResponse()
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

// DescribeProIspPlaySumInfoList
// 查询某段时间内每个国家地区每个省份每个运营商的平均每秒流量，总流量，总请求数信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribeProvinceIspPlayInfoList
// 查询某省份某运营商下行播放数据，包括带宽，流量，请求数，并发连接数信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_HASNOTLIVINGSTREAM = "InternalError.HasNotLivingStream"
//  INTERNALERROR_INVALIDREQUEST = "InternalError.InvalidRequest"
//  INTERNALERROR_QUERYPROISPPLAYINFOERROR = "InternalError.QueryProIspPlayInfoError"
//  INTERNALERROR_QUERYUPLOADINFOFAILED = "InternalError.QueryUploadInfoFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// DescribePullStreamConfigs
// 查询直播拉流配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribePullStreamConfigs(request *DescribePullStreamConfigsRequest) (response *DescribePullStreamConfigsResponse, err error) {
    if request == nil {
        request = NewDescribePullStreamConfigsRequest()
    }
    response = NewDescribePullStreamConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordTaskRequest() (request *DescribeRecordTaskRequest) {
    request = &DescribeRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeRecordTask")
    return
}

func NewDescribeRecordTaskResponse() (response *DescribeRecordTaskResponse) {
    response = &DescribeRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordTask
// 查询指定时间段范围内启动和结束的录制任务列表。
//
// - 使用前提
//
// 1. 仅用于查询由 CreateRecordTask 接口创建的录制任务。
//
// 2. 不能查询被 DeleteRecordTask 接口删除以及已过期（平台侧保留3个月）的录制任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRecordTask(request *DescribeRecordTaskRequest) (response *DescribeRecordTaskResponse, err error) {
    if request == nil {
        request = NewDescribeRecordTaskRequest()
    }
    response = NewDescribeRecordTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScreenShotSheetNumListRequest() (request *DescribeScreenShotSheetNumListRequest) {
    request = &DescribeScreenShotSheetNumListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeScreenShotSheetNumList")
    return
}

func NewDescribeScreenShotSheetNumListResponse() (response *DescribeScreenShotSheetNumListResponse) {
    response = &DescribeScreenShotSheetNumListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScreenShotSheetNumList
// 接口用来查询直播增值业务--截图的张数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeScreenShotSheetNumList(request *DescribeScreenShotSheetNumListRequest) (response *DescribeScreenShotSheetNumListResponse, err error) {
    if request == nil {
        request = NewDescribeScreenShotSheetNumListRequest()
    }
    response = NewDescribeScreenShotSheetNumListResponse()
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

// DescribeStreamDayPlayInfoList
// 查询天维度每条流的播放数据，包括总流量等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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

// DescribeStreamPlayInfoList
// 查询播放数据，支持按流名称查询详细播放数据，也可按播放域名查询详细总数据，数据延迟4分钟左右。
//
// 注意：按AppName查询请先联系工单申请，开通后配置生效预计需要5个工作日左右，具体时间以最终回复为准。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeStreamPlayInfoList(request *DescribeStreamPlayInfoListRequest) (response *DescribeStreamPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPlayInfoListRequest()
    }
    response = NewDescribeStreamPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamPushInfoListRequest() (request *DescribeStreamPushInfoListRequest) {
    request = &DescribeStreamPushInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeStreamPushInfoList")
    return
}

func NewDescribeStreamPushInfoListResponse() (response *DescribeStreamPushInfoListResponse) {
    response = &DescribeStreamPushInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStreamPushInfoList
// 查询流id的上行推流质量数据，包括音视频的帧率，码率，流逝时间，编码格式等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeStreamPushInfoList(request *DescribeStreamPushInfoListRequest) (response *DescribeStreamPushInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPushInfoListRequest()
    }
    response = NewDescribeStreamPushInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopClientIpSumInfoListRequest() (request *DescribeTopClientIpSumInfoListRequest) {
    request = &DescribeTopClientIpSumInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeTopClientIpSumInfoList")
    return
}

func NewDescribeTopClientIpSumInfoListResponse() (response *DescribeTopClientIpSumInfoListResponse) {
    response = &DescribeTopClientIpSumInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopClientIpSumInfoList
// 查询某段时间top n客户端ip汇总信息（暂支持top 1000）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeTopClientIpSumInfoList(request *DescribeTopClientIpSumInfoListRequest) (response *DescribeTopClientIpSumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeTopClientIpSumInfoListRequest()
    }
    response = NewDescribeTopClientIpSumInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUploadStreamNumsRequest() (request *DescribeUploadStreamNumsRequest) {
    request = &DescribeUploadStreamNumsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeUploadStreamNums")
    return
}

func NewDescribeUploadStreamNumsResponse() (response *DescribeUploadStreamNumsResponse) {
    response = &DescribeUploadStreamNumsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUploadStreamNums
// 直播上行路数查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeUploadStreamNums(request *DescribeUploadStreamNumsRequest) (response *DescribeUploadStreamNumsResponse, err error) {
    if request == nil {
        request = NewDescribeUploadStreamNumsRequest()
    }
    response = NewDescribeUploadStreamNumsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVisitTopSumInfoListRequest() (request *DescribeVisitTopSumInfoListRequest) {
    request = &DescribeVisitTopSumInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "DescribeVisitTopSumInfoList")
    return
}

func NewDescribeVisitTopSumInfoListResponse() (response *DescribeVisitTopSumInfoListResponse) {
    response = &DescribeVisitTopSumInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVisitTopSumInfoList
// 查询某时间段top n的域名或流id信息（暂支持top 1000）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeVisitTopSumInfoList(request *DescribeVisitTopSumInfoListRequest) (response *DescribeVisitTopSumInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeVisitTopSumInfoListRequest()
    }
    response = NewDescribeVisitTopSumInfoListResponse()
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

// DropLiveStream
// 断开推流连接，但可以重新推流。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_STREAMNOTALIVE = "ResourceNotFound.StreamNotAlive"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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

// EnableLiveDomain
// 启用状态为停用的直播域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SDKNOPACKAGE = "FailedOperation.SdkNoPackage"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_CLOUDDOMAINISSTOP = "InvalidParameter.CloudDomainIsStop"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"
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

// ForbidLiveDomain
// 停止使用某个直播域名。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  RESOURCENOTFOUND = "ResourceNotFound"
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

// ForbidLiveStream
// 禁止某条流的推送，可以预设某个时刻将流恢复。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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

// ModifyLiveCallbackTemplate
// 修改回调模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETER_URLNOTSAFE = "InvalidParameter.UrlNotSafe"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
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

// ModifyLiveCert
// 修改证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CRTDATEINUSING = "InternalError.CrtDateInUsing"
//  INTERNALERROR_CRTDATEOVERDUE = "InternalError.CrtDateOverdue"
//  INTERNALERROR_CRTKEYNOTMATCH = "InternalError.CrtKeyNotMatch"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CRTDATEINUSING = "InvalidParameter.CrtDateInUsing"
//  INVALIDPARAMETER_CRTDATEOVERDUE = "InvalidParameter.CrtDateOverdue"
//  INVALIDPARAMETER_CRTKEYNOTMATCH = "InvalidParameter.CrtKeyNotMatch"
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

// ModifyLiveDomainCert
// 修改域名和证书绑定信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CRTDATENOTFOUND = "InternalError.CrtDateNotFound"
//  INTERNALERROR_CRTDOMAINNOTFOUND = "InternalError.CrtDomainNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CRTDATENOTLEGAL = "InvalidParameter.CrtDateNotLegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"
func (c *Client) ModifyLiveDomainCert(request *ModifyLiveDomainCertRequest) (response *ModifyLiveDomainCertResponse, err error) {
    if request == nil {
        request = NewModifyLiveDomainCertRequest()
    }
    response = NewModifyLiveDomainCertResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLiveDomainRefererRequest() (request *ModifyLiveDomainRefererRequest) {
    request = &ModifyLiveDomainRefererRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLiveDomainReferer")
    return
}

func NewModifyLiveDomainRefererResponse() (response *ModifyLiveDomainRefererResponse) {
    response = &ModifyLiveDomainRefererResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLiveDomainReferer
// 设置直播域名 Referer 黑白名单。
//
// 由于 Referer 信息包含在 http 协议中，在开启配置后，播放协议为 rtmp 或 webrtc 不会校验 Referer 配置，仍可正常播放。如需配置 Referer 鉴权建议使用 http-flv 或 http-hls 协议播放。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
func (c *Client) ModifyLiveDomainReferer(request *ModifyLiveDomainRefererRequest) (response *ModifyLiveDomainRefererResponse, err error) {
    if request == nil {
        request = NewModifyLiveDomainRefererRequest()
    }
    response = NewModifyLiveDomainRefererResponse()
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

// ModifyLivePlayAuthKey
// 修改播放鉴权key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PLAYDOMAINNORECORD = "InternalError.PlayDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"
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

// ModifyLivePlayDomain
// 修改播放域名信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"
//  RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"
func (c *Client) ModifyLivePlayDomain(request *ModifyLivePlayDomainRequest) (response *ModifyLivePlayDomainResponse, err error) {
    if request == nil {
        request = NewModifyLivePlayDomainRequest()
    }
    response = NewModifyLivePlayDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLivePullStreamTaskRequest() (request *ModifyLivePullStreamTaskRequest) {
    request = &ModifyLivePullStreamTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "ModifyLivePullStreamTask")
    return
}

func NewModifyLivePullStreamTaskResponse() (response *ModifyLivePullStreamTaskResponse) {
    response = &ModifyLivePullStreamTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLivePullStreamTask
// 更新直播拉流任务。 
//
// 1. 不支持修改目标地址，如需推到新地址，请创建新任务。
//
// 2. 不支持修改任务类型，如需更换，请创建新任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCALLBACKURL = "InvalidParameter.InvalidCallbackUrl"
//  INVALIDPARAMETER_INVALIDSOURCEURL = "InvalidParameter.InvalidSourceUrl"
//  INVALIDPARAMETER_INVALIDTASKTIME = "InvalidParameter.InvalidTaskTime"
//  INVALIDPARAMETER_INVALIDTOURL = "InvalidParameter.InvalidToUrl"
//  INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"
//  INVALIDPARAMETER_TOURLNOPERMISSION = "InvalidParameter.ToUrlNoPermission"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ModifyLivePullStreamTask(request *ModifyLivePullStreamTaskRequest) (response *ModifyLivePullStreamTaskResponse, err error) {
    if request == nil {
        request = NewModifyLivePullStreamTaskRequest()
    }
    response = NewModifyLivePullStreamTaskResponse()
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

// ModifyLivePushAuthKey
// 修改直播推流鉴权key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PUSHDOMAINNORECORD = "InternalError.PushDomainNoRecord"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"
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

// ModifyLiveRecordTemplate
// 修改录制模板配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// ModifyLiveSnapshotTemplate
// 修改截图模板配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// ModifyLiveTranscodeTemplate
// 修改转码模板配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"
//  FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"
//  INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"
//  INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"
//  INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"
//  INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"
//  INTERNALERROR_NOTFOUND = "InternalError.NotFound"
//  INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"
//  INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"
//  INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"
//  INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// ModifyPullStreamConfig
// 更新拉流配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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

// ModifyPullStreamStatus
// 修改直播拉流配置的状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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

// ResumeDelayLiveStream
// 恢复延迟播放设置
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

// ResumeLiveStream
// 恢复某条流的推流。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"
//  INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"
//  INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"
//  INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"
//  INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"
//  INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ResumeLiveStream(request *ResumeLiveStreamRequest) (response *ResumeLiveStreamResponse, err error) {
    if request == nil {
        request = NewResumeLiveStreamRequest()
    }
    response = NewResumeLiveStreamResponse()
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

// StopLiveRecord
// 说明：录制后的文件存放于点播平台。用户如需使用录制功能，需首先自行开通点播账号并确保账号可用。录制文件存放后，相关费用（含存储以及下行播放流量）按照点播平台计费方式收取，请参考对应文档。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"
//  RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"
func (c *Client) StopLiveRecord(request *StopLiveRecordRequest) (response *StopLiveRecordResponse, err error) {
    if request == nil {
        request = NewStopLiveRecordRequest()
    }
    response = NewStopLiveRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStopRecordTaskRequest() (request *StopRecordTaskRequest) {
    request = &StopRecordTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("live", APIVersion, "StopRecordTask")
    return
}

func NewStopRecordTaskResponse() (response *StopRecordTaskResponse) {
    response = &StopRecordTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopRecordTask
// 提前结束录制，并中止运行中的录制任务。任务被成功终止后，本次任务将不再启动。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopRecordTask(request *StopRecordTaskRequest) (response *StopRecordTaskResponse, err error) {
    if request == nil {
        request = NewStopRecordTaskRequest()
    }
    response = NewStopRecordTaskResponse()
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

// UnBindLiveDomainCert
// 解绑域名证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
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

// UpdateLiveWatermark
// 更新水印。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"
//  INTERNALERROR_WATERMARKEDITERROR = "InternalError.WatermarkEditError"
//  INTERNALERROR_WATERMARKNOTEXIST = "InternalError.WatermarkNotExist"
//  RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"
func (c *Client) UpdateLiveWatermark(request *UpdateLiveWatermarkRequest) (response *UpdateLiveWatermarkResponse, err error) {
    if request == nil {
        request = NewUpdateLiveWatermarkRequest()
    }
    response = NewUpdateLiveWatermarkResponse()
    err = c.Send(request, response)
    return
}
