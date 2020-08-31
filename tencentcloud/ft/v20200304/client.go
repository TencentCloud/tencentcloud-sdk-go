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

package v20200304

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-03-04"

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


func NewCancelFaceMorphJobRequest() (request *CancelFaceMorphJobRequest) {
    request = &CancelFaceMorphJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ft", APIVersion, "CancelFaceMorphJob")
    return
}

func NewCancelFaceMorphJobResponse() (response *CancelFaceMorphJobResponse) {
    response = &CancelFaceMorphJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 撤销人像渐变任务请求
func (c *Client) CancelFaceMorphJob(request *CancelFaceMorphJobRequest) (response *CancelFaceMorphJobResponse, err error) {
    if request == nil {
        request = NewCancelFaceMorphJobRequest()
    }
    response = NewCancelFaceMorphJobResponse()
    err = c.Send(request, response)
    return
}

func NewChangeAgePicRequest() (request *ChangeAgePicRequest) {
    request = &ChangeAgePicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ft", APIVersion, "ChangeAgePic")
    return
}

func NewChangeAgePicResponse() (response *ChangeAgePicResponse) {
    response = &ChangeAgePicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户上传一张人脸图片，基于人脸编辑与生成算法，输出一张人脸变老或变年轻的图片，支持实现人脸不同年龄的变化。
func (c *Client) ChangeAgePic(request *ChangeAgePicRequest) (response *ChangeAgePicResponse, err error) {
    if request == nil {
        request = NewChangeAgePicRequest()
    }
    response = NewChangeAgePicResponse()
    err = c.Send(request, response)
    return
}

func NewFaceCartoonPicRequest() (request *FaceCartoonPicRequest) {
    request = &FaceCartoonPicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ft", APIVersion, "FaceCartoonPic")
    return
}

func NewFaceCartoonPicResponse() (response *FaceCartoonPicResponse) {
    response = &FaceCartoonPicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 人像动漫化
func (c *Client) FaceCartoonPic(request *FaceCartoonPicRequest) (response *FaceCartoonPicResponse, err error) {
    if request == nil {
        request = NewFaceCartoonPicRequest()
    }
    response = NewFaceCartoonPicResponse()
    err = c.Send(request, response)
    return
}

func NewMorphFaceRequest() (request *MorphFaceRequest) {
    request = &MorphFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ft", APIVersion, "MorphFace")
    return
}

func NewMorphFaceResponse() (response *MorphFaceResponse) {
    response = &MorphFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户上传2-5张人脸照片，输出一段人脸变换特效视频
func (c *Client) MorphFace(request *MorphFaceRequest) (response *MorphFaceResponse, err error) {
    if request == nil {
        request = NewMorphFaceRequest()
    }
    response = NewMorphFaceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFaceMorphJobRequest() (request *QueryFaceMorphJobRequest) {
    request = &QueryFaceMorphJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ft", APIVersion, "QueryFaceMorphJob")
    return
}

func NewQueryFaceMorphJobResponse() (response *QueryFaceMorphJobResponse) {
    response = &QueryFaceMorphJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询人像渐变处理进度
func (c *Client) QueryFaceMorphJob(request *QueryFaceMorphJobRequest) (response *QueryFaceMorphJobResponse, err error) {
    if request == nil {
        request = NewQueryFaceMorphJobRequest()
    }
    response = NewQueryFaceMorphJobResponse()
    err = c.Send(request, response)
    return
}

func NewSwapGenderPicRequest() (request *SwapGenderPicRequest) {
    request = &SwapGenderPicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ft", APIVersion, "SwapGenderPic")
    return
}

func NewSwapGenderPicResponse() (response *SwapGenderPicResponse) {
    response = &SwapGenderPicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户上传一张人脸图片，基于人脸编辑与生成算法，输出一张人脸性别转换的图片。男变女可实现美颜、淡妆、加刘海和长发的效果；女变男可实现加胡须、变短发的效果。 
func (c *Client) SwapGenderPic(request *SwapGenderPicRequest) (response *SwapGenderPicResponse, err error) {
    if request == nil {
        request = NewSwapGenderPicRequest()
    }
    response = NewSwapGenderPicResponse()
    err = c.Send(request, response)
    return
}
