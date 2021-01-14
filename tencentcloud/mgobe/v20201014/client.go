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

package v20201014

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-10-14"

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


func NewChangeRoomPlayerProfileRequest() (request *ChangeRoomPlayerProfileRequest) {
    request = &ChangeRoomPlayerProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mgobe", APIVersion, "ChangeRoomPlayerProfile")
    return
}

func NewChangeRoomPlayerProfileResponse() (response *ChangeRoomPlayerProfileResponse) {
    response = &ChangeRoomPlayerProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改房间玩家自定义属性
func (c *Client) ChangeRoomPlayerProfile(request *ChangeRoomPlayerProfileRequest) (response *ChangeRoomPlayerProfileResponse, err error) {
    if request == nil {
        request = NewChangeRoomPlayerProfileRequest()
    }
    response = NewChangeRoomPlayerProfileResponse()
    err = c.Send(request, response)
    return
}

func NewChangeRoomPlayerStatusRequest() (request *ChangeRoomPlayerStatusRequest) {
    request = &ChangeRoomPlayerStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mgobe", APIVersion, "ChangeRoomPlayerStatus")
    return
}

func NewChangeRoomPlayerStatusResponse() (response *ChangeRoomPlayerStatusResponse) {
    response = &ChangeRoomPlayerStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改玩家自定义状态
func (c *Client) ChangeRoomPlayerStatus(request *ChangeRoomPlayerStatusRequest) (response *ChangeRoomPlayerStatusResponse, err error) {
    if request == nil {
        request = NewChangeRoomPlayerStatusRequest()
    }
    response = NewChangeRoomPlayerStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDismissRoomRequest() (request *DismissRoomRequest) {
    request = &DismissRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mgobe", APIVersion, "DismissRoom")
    return
}

func NewDismissRoomResponse() (response *DismissRoomResponse) {
    response = &DismissRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过game_id、room_id解散房间
func (c *Client) DismissRoom(request *DismissRoomRequest) (response *DismissRoomResponse, err error) {
    if request == nil {
        request = NewDismissRoomRequest()
    }
    response = NewDismissRoomResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoomRequest() (request *ModifyRoomRequest) {
    request = &ModifyRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mgobe", APIVersion, "ModifyRoom")
    return
}

func NewModifyRoomResponse() (response *ModifyRoomResponse) {
    response = &ModifyRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改房间
func (c *Client) ModifyRoom(request *ModifyRoomRequest) (response *ModifyRoomResponse, err error) {
    if request == nil {
        request = NewModifyRoomRequest()
    }
    response = NewModifyRoomResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveRoomPlayerRequest() (request *RemoveRoomPlayerRequest) {
    request = &RemoveRoomPlayerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mgobe", APIVersion, "RemoveRoomPlayer")
    return
}

func NewRemoveRoomPlayerResponse() (response *RemoveRoomPlayerResponse) {
    response = &RemoveRoomPlayerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 踢出房间玩家
func (c *Client) RemoveRoomPlayer(request *RemoveRoomPlayerRequest) (response *RemoveRoomPlayerResponse, err error) {
    if request == nil {
        request = NewRemoveRoomPlayerRequest()
    }
    response = NewRemoveRoomPlayerResponse()
    err = c.Send(request, response)
    return
}
