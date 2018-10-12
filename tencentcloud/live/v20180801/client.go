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

func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithSecretId(credential.SecretId, credential.SecretKey).
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

// 对流设置延播
func (c *Client) AddDelayLiveStream(request *AddDelayLiveStreamRequest) (response *AddDelayLiveStreamResponse, err error) {
    if request == nil {
        request = NewAddDelayLiveStreamRequest()
    }
    response = NewAddDelayLiveStreamResponse()
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
