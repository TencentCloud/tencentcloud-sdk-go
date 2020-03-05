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

package v20191213

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-12-13"

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


func NewBeautifyPicRequest() (request *BeautifyPicRequest) {
    request = &BeautifyPicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("fmu", APIVersion, "BeautifyPic")
    return
}

func NewBeautifyPicResponse() (response *BeautifyPicResponse) {
    response = &BeautifyPicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户上传一张人脸图片，精准定位五官，实现美肤、亮肤、祛痘等美颜功能。
func (c *Client) BeautifyPic(request *BeautifyPicRequest) (response *BeautifyPicResponse, err error) {
    if request == nil {
        request = NewBeautifyPicRequest()
    }
    response = NewBeautifyPicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModelRequest() (request *CreateModelRequest) {
    request = &CreateModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("fmu", APIVersion, "CreateModel")
    return
}

func NewCreateModelResponse() (response *CreateModelResponse) {
    response = &CreateModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 在使用LUT素材的modelid实现试唇色前，您需要先上传 LUT 格式的cube文件注册唇色ID。查看 [LUT文件的使用说明](https://cloud.tencent.com/document/product/1172/41701)。
// 
// 注：您也可以直接使用 [试唇色接口](https://cloud.tencent.com/document/product/1172/40706)，通过输入RGBA模型数值的方式指定唇色，更简单易用。
func (c *Client) CreateModel(request *CreateModelRequest) (response *CreateModelResponse, err error) {
    if request == nil {
        request = NewCreateModelRequest()
    }
    response = NewCreateModelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteModelRequest() (request *DeleteModelRequest) {
    request = &DeleteModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("fmu", APIVersion, "DeleteModel")
    return
}

func NewDeleteModelResponse() (response *DeleteModelResponse) {
    response = &DeleteModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除已注册的唇色素材。
func (c *Client) DeleteModel(request *DeleteModelRequest) (response *DeleteModelResponse, err error) {
    if request == nil {
        request = NewDeleteModelRequest()
    }
    response = NewDeleteModelResponse()
    err = c.Send(request, response)
    return
}

func NewGetModelListRequest() (request *GetModelListRequest) {
    request = &GetModelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("fmu", APIVersion, "GetModelList")
    return
}

func NewGetModelListResponse() (response *GetModelListResponse) {
    response = &GetModelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询已注册的唇色素材。
func (c *Client) GetModelList(request *GetModelListRequest) (response *GetModelListResponse, err error) {
    if request == nil {
        request = NewGetModelListRequest()
    }
    response = NewGetModelListResponse()
    err = c.Send(request, response)
    return
}

func NewTryLipstickPicRequest() (request *TryLipstickPicRequest) {
    request = &TryLipstickPicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("fmu", APIVersion, "TryLipstickPic")
    return
}

func NewTryLipstickPicResponse() (response *TryLipstickPicResponse) {
    response = &TryLipstickPicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对图片中的人脸嘴唇进行着色，最多支持同时对一张图中的3张人脸进行试唇色。
// 
// 您可以通过事先注册在腾讯云的唇色素材（LUT文件）改变图片中的人脸唇色，也可以输入RGBA模型数值。
// 
// 为了更好的效果，建议您使用事先注册在腾讯云的唇色素材（LUT文件）。
// 
// >     
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) TryLipstickPic(request *TryLipstickPicRequest) (response *TryLipstickPicResponse, err error) {
    if request == nil {
        request = NewTryLipstickPicRequest()
    }
    response = NewTryLipstickPicResponse()
    err = c.Send(request, response)
    return
}
