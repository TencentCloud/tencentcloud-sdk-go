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

package v20200210

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-02-10"

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


func NewManageDeviceRiskRequest() (request *ManageDeviceRiskRequest) {
    request = &ManageDeviceRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("taf", APIVersion, "ManageDeviceRisk")
    
    
    return
}

func NewManageDeviceRiskResponse() (response *ManageDeviceRiskResponse) {
    response = &ManageDeviceRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ManageDeviceRisk
// oaid 设备风险接口
func (c *Client) ManageDeviceRisk(request *ManageDeviceRiskRequest) (response *ManageDeviceRiskResponse, err error) {
    return c.ManageDeviceRiskWithContext(context.Background(), request)
}

// ManageDeviceRisk
// oaid 设备风险接口
func (c *Client) ManageDeviceRiskWithContext(ctx context.Context, request *ManageDeviceRiskRequest) (response *ManageDeviceRiskResponse, err error) {
    if request == nil {
        request = NewManageDeviceRiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageDeviceRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewManageDeviceRiskResponse()
    err = c.Send(request, response)
    return
}

func NewManagePortraitRiskRequest() (request *ManagePortraitRiskRequest) {
    request = &ManagePortraitRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("taf", APIVersion, "ManagePortraitRisk")
    
    
    return
}

func NewManagePortraitRiskResponse() (response *ManagePortraitRiskResponse) {
    response = &ManagePortraitRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ManagePortraitRisk
// 虚假流量识别
func (c *Client) ManagePortraitRisk(request *ManagePortraitRiskRequest) (response *ManagePortraitRiskResponse, err error) {
    return c.ManagePortraitRiskWithContext(context.Background(), request)
}

// ManagePortraitRisk
// 虚假流量识别
func (c *Client) ManagePortraitRiskWithContext(ctx context.Context, request *ManagePortraitRiskRequest) (response *ManagePortraitRiskResponse, err error) {
    if request == nil {
        request = NewManagePortraitRiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManagePortraitRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewManagePortraitRiskResponse()
    err = c.Send(request, response)
    return
}
