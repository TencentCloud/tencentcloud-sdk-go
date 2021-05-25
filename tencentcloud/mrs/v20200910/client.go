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

package v20200910

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-09-10"

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


func NewImageToClassRequest() (request *ImageToClassRequest) {
    request = &ImageToClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mrs", APIVersion, "ImageToClass")
    return
}

func NewImageToClassResponse() (response *ImageToClassResponse) {
    response = &ImageToClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 图片分类
func (c *Client) ImageToClass(request *ImageToClassRequest) (response *ImageToClassResponse, err error) {
    if request == nil {
        request = NewImageToClassRequest()
    }
    response = NewImageToClassResponse()
    err = c.Send(request, response)
    return
}

func NewImageToObjectRequest() (request *ImageToObjectRequest) {
    request = &ImageToObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mrs", APIVersion, "ImageToObject")
    return
}

func NewImageToObjectResponse() (response *ImageToObjectResponse) {
    response = &ImageToObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 图片转结构化对象
func (c *Client) ImageToObject(request *ImageToObjectRequest) (response *ImageToObjectResponse, err error) {
    if request == nil {
        request = NewImageToObjectRequest()
    }
    response = NewImageToObjectResponse()
    err = c.Send(request, response)
    return
}

func NewTextToClassRequest() (request *TextToClassRequest) {
    request = &TextToClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mrs", APIVersion, "TextToClass")
    return
}

func NewTextToClassResponse() (response *TextToClassResponse) {
    response = &TextToClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 文本分类
func (c *Client) TextToClass(request *TextToClassRequest) (response *TextToClassResponse, err error) {
    if request == nil {
        request = NewTextToClassRequest()
    }
    response = NewTextToClassResponse()
    err = c.Send(request, response)
    return
}

func NewTextToObjectRequest() (request *TextToObjectRequest) {
    request = &TextToObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mrs", APIVersion, "TextToObject")
    return
}

func NewTextToObjectResponse() (response *TextToObjectResponse) {
    response = &TextToObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 文本转结构化对象
func (c *Client) TextToObject(request *TextToObjectRequest) (response *TextToObjectResponse, err error) {
    if request == nil {
        request = NewTextToObjectRequest()
    }
    response = NewTextToObjectResponse()
    err = c.Send(request, response)
    return
}
