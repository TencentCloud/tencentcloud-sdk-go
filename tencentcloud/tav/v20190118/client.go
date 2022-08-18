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

package v20190118

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-01-18"

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


func NewGetLocalEngineRequest() (request *GetLocalEngineRequest) {
    request = &GetLocalEngineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tav", APIVersion, "GetLocalEngine")
    
    
    return
}

func NewGetLocalEngineResponse() (response *GetLocalEngineResponse) {
    response = &GetLocalEngineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetLocalEngine
// 获取TAV本地引擎
func (c *Client) GetLocalEngine(request *GetLocalEngineRequest) (response *GetLocalEngineResponse, err error) {
    return c.GetLocalEngineWithContext(context.Background(), request)
}

// GetLocalEngine
// 获取TAV本地引擎
func (c *Client) GetLocalEngineWithContext(ctx context.Context, request *GetLocalEngineRequest) (response *GetLocalEngineResponse, err error) {
    if request == nil {
        request = NewGetLocalEngineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLocalEngine require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLocalEngineResponse()
    err = c.Send(request, response)
    return
}

func NewGetScanResultRequest() (request *GetScanResultRequest) {
    request = &GetScanResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tav", APIVersion, "GetScanResult")
    
    
    return
}

func NewGetScanResultResponse() (response *GetScanResultResponse) {
    response = &GetScanResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetScanResult
// tav文件上传扫描结果查询
func (c *Client) GetScanResult(request *GetScanResultRequest) (response *GetScanResultResponse, err error) {
    return c.GetScanResultWithContext(context.Background(), request)
}

// GetScanResult
// tav文件上传扫描结果查询
func (c *Client) GetScanResultWithContext(ctx context.Context, request *GetScanResultRequest) (response *GetScanResultResponse, err error) {
    if request == nil {
        request = NewGetScanResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetScanResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetScanResultResponse()
    err = c.Send(request, response)
    return
}

func NewScanFileRequest() (request *ScanFileRequest) {
    request = &ScanFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tav", APIVersion, "ScanFile")
    
    
    return
}

func NewScanFileResponse() (response *ScanFileResponse) {
    response = &ScanFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanFile
// tav文件上传扫描
func (c *Client) ScanFile(request *ScanFileRequest) (response *ScanFileResponse, err error) {
    return c.ScanFileWithContext(context.Background(), request)
}

// ScanFile
// tav文件上传扫描
func (c *Client) ScanFileWithContext(ctx context.Context, request *ScanFileRequest) (response *ScanFileResponse, err error) {
    if request == nil {
        request = NewScanFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanFileResponse()
    err = c.Send(request, response)
    return
}

func NewScanFileHashRequest() (request *ScanFileHashRequest) {
    request = &ScanFileHashRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tav", APIVersion, "ScanFileHash")
    
    
    return
}

func NewScanFileHashResponse() (response *ScanFileHashResponse) {
    response = &ScanFileHashResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanFileHash
// 通过文件哈希值获取文件黑白属性
func (c *Client) ScanFileHash(request *ScanFileHashRequest) (response *ScanFileHashResponse, err error) {
    return c.ScanFileHashWithContext(context.Background(), request)
}

// ScanFileHash
// 通过文件哈希值获取文件黑白属性
func (c *Client) ScanFileHashWithContext(ctx context.Context, request *ScanFileHashRequest) (response *ScanFileHashResponse, err error) {
    if request == nil {
        request = NewScanFileHashRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanFileHash require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanFileHashResponse()
    err = c.Send(request, response)
    return
}
