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

package v20190722

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-22"

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


func NewDissolveRoomRequest() (request *DissolveRoomRequest) {
    request = &DissolveRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "DissolveRoom")
    return
}

func NewDissolveRoomResponse() (response *DissolveRoomResponse) {
    response = &DissolveRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 接口说明：把房间所有用户从房间踢出，解散房间。
func (c *Client) DissolveRoom(request *DissolveRoomRequest) (response *DissolveRoomResponse, err error) {
    if request == nil {
        request = NewDissolveRoomRequest()
    }
    response = NewDissolveRoomResponse()
    err = c.Send(request, response)
    return
}

func NewKickOutUserRequest() (request *KickOutUserRequest) {
    request = &KickOutUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trtc", APIVersion, "KickOutUser")
    return
}

func NewKickOutUserResponse() (response *KickOutUserResponse) {
    response = &KickOutUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 接口说明：将用户从房间踢出。
func (c *Client) KickOutUser(request *KickOutUserRequest) (response *KickOutUserResponse, err error) {
    if request == nil {
        request = NewKickOutUserRequest()
    }
    response = NewKickOutUserResponse()
    err = c.Send(request, response)
    return
}
